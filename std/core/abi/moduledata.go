// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package abi

import (
	"unsafe"
	_ "unsafe" // for go:linkname

	stdstring "github.com/primecitizens/pcz/std/builtin/string"
	"github.com/primecitizens/pcz/std/core/arch"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/mark"
)

var (
	//go:linkname firstmoduledata runtime.firstmoduledata
	firstmoduledata ModuleData // linker symbol

	//go:linkname lastmoduledatap runtime.lastmoduledatap
	lastmoduledatap *ModuleData // linker symbol
)

// ModuleIter is an iterator for active modules.
type ModuleIter uintptr

func (iter *ModuleIter) Nth(i int) (md *ModuleData, ok bool) {
	if i == 0 {
		*iter = ModuleIter(uintptr(unsafe.Pointer(&firstmoduledata)))
	}

	for md := (*ModuleData)(unsafe.Pointer(uintptr(*iter))); md != nil; md = md.Next {
		if md.Bad {
			continue
		}

		*iter = ModuleIter(uintptr(unsafe.Pointer(md.Next)))
		return md, true
	}

	*iter = 0
	return nil, false
}

func findModuleForType(ptrInModule uintptr) (md *ModuleData) {
	for i, iter := 0, ModuleIter(0); ; i++ {
		md, ok := iter.Nth(i)
		if !ok {
			break
		}

		if ptrInModule >= md.Types && ptrInModule < md.ETypes {
			return md
		}
	}

	return nil
}

func resolveNameOff(ptrInModule unsafe.Pointer, off NameOff) Name {
	if off == 0 {
		return Name{}
	}

	md := findModuleForType(uintptr(ptrInModule))
	if md == nil {
		// TODO: support runtime name
		// 	// No module found. see if it is a run time name.
		// 	reflectOffsLock()
		// 	res, found := reflectOffs.m[int32(off)]
		// 	reflectOffsUnlock()
		// 	if !found {
		// 		println("runtime: nameOff", hex(off), "base", hex(base), "not in ranges:")
		// 		for next := &firstmoduledata; next != nil; next = next.next {
		// 			println("\ttypes", hex(next.types), "etypes", hex(next.etypes))
		// 		}
		// 		throw("runtime: name offset base pointer out of range")
		// 	}
		//
		// 	return Name{(*byte)(res)}
		return Name{}
	}

	res := md.Types + uintptr(off)
	if res > md.ETypes {
		println("runtime: nameOff", hex(off), "out of range", hex(md.Types), "-", hex(md.ETypes))
		assert.Throw("runtime:", "name", "offset", "out", "of", "range")
	}

	return Name{Bytes: (*byte)(unsafe.Pointer(res))}
}

// resolveTypeOff resolves an *rtype offset from a base type.
// The (*rtype).typeOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
func resolveTypeOff(ptrInModule unsafe.Pointer, off TypeOff) *Type {
	if off == 0 || off == -1 {
		// -1 is the sentinel value for unreachable code.
		// See cmd/link/internal/ld/data.go:relocsym.
		return nil
	}

	md := findModuleForType(uintptr(ptrInModule))
	if md == nil {
		// TODO: support runtime name
		// reflectOffsLock()
		// res := reflectOffs.m[int32(off)]
		// reflectOffsUnlock()
		// if res == nil {
		// 	println("runtime: typeOff", hex(off), "base", hex(base), "not in ranges:")
		// 	for next := &firstmoduledata; next != nil; next = next.next {
		// 		println("\ttypes", hex(next.types), "etypes", hex(next.etypes))
		// 	}
		// 	throw("runtime: type offset base pointer out of range")
		// }
		// return (*_type)(res)
		return nil
	}

	// TODO: initialize md.typemap (see $GOROOT/src/runtime/type.go#func.typelinksinit)
	// if t := md.typemap[off]; t != nil {
	// 	return t
	// }

	res := md.Types + uintptr(off)
	if res > md.ETypes {
		println("runtime: typeOff", hex(off), "out of range", hex(md.Types), "-", hex(md.ETypes))
		assert.Throw("runtime:", "type", "offset", "out", "of", "range")
	}
	return (*Type)(unsafe.Pointer(res))
}

type hex = uint64

//go:linkname unreachableMethod runtime.unreachableMethod
func unreachableMethod()

// resolveTextOff resolves a function pointer offset from a base type.
// The (*rtype).textOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
func resolveTextOff(rtype unsafe.Pointer, off TextOff) unsafe.Pointer {
	if off == -1 {
		// -1 is the sentinel value for unreachable code.
		// See cmd/link/internal/ld/data.go:relocsym.
		return unsafe.Pointer(FuncPCABIInternal(unreachableMethod))
	}

	md := findModuleForType(uintptr(rtype))
	if md == nil {
		// TODO: support runtime type.
		// reflectOffsLock()
		// res := reflectOffs.m[int32(off)]
		// reflectOffsUnlock()
		// if res == nil {
		// 	println("runtime: textOff", hex(off), "base", hex(base), "not in ranges:")
		// 	for next := &firstmoduledata; next != nil; next = next.next {
		// 		println("\ttypes", hex(next.types), "etypes", hex(next.etypes))
		// 	}
		// 	throw("runtime: text offset base pointer out of range")
		// }
		// return res

		assert.Throw("runtime:", "text", "offset", "base", "pointer", "out", "of", "range")
	}

	res := md.textAddr(uint32(off))
	return unsafe.Pointer(res)
}

type GetItabResult uint8

const (
	GetItabResult_OK GetItabResult = iota
	GetItabResult_UncommonType
	GetItabResult_
)

func GetItab(inter *InterfaceType, typ *Type) (*Itab, GetItabResult) {
	if len(inter.Methods) == 0 {
		assert.Throw("internal", "error:", "misuse", "of", "itab")
	}

	// easy case
	if typ.TFlag&TFlagUncommon == 0 {
		return nil, GetItabResult_UncommonType
	}

	var m *Itab
	// 	// First, look in the existing table to see if we can find the itab we need.
	// 	// This is by far the most common case, so do it without locks.
	// 	// Use atomic to ensure we see any previous writes done by the thread
	// 	// that updates the itabTable field (with atomic.Storep in itabAdd).
	// 	t := (*itabTableType)(atomic.Loadp(unsafe.Pointer(&itabTable)))
	// 	if m = t.find(inter, typ); m != nil {
	// 		goto finish
	// 	}
	//
	// 	// Not found.  Grab the lock and try again.
	// 	lock(&itabLock)
	// 	if m = itabTable.find(inter, typ); m != nil {
	// 		unlock(&itabLock)
	// 		goto finish
	// 	}
	//
	// 	// Entry doesn't exist yet. Make a new entry & add it.
	// 	m = (*itab)(persistentalloc(unsafe.Sizeof(itab{})+uintptr(len(inter.Methods)-1)*goarch.PtrSize, 0, &memstats.other_sys))
	// 	m.inter = inter
	// 	m._type = typ
	// 	// The hash is used in type switches. However, compiler statically generates itab's
	// 	// for all interface/type pairs used in switches (which are added to itabTable
	// 	// in itabsinit). The dynamically-generated itab's never participate in type switches,
	// 	// and thus the hash is irrelevant.
	// 	// Note: m.hash is _not_ the hash used for the runtime itabTable hash table.
	// 	m.hash = 0
	// 	m.init()
	// 	itabAdd(m)
	// 	unlock(&itabLock)
	// finish:
	if m.Fun[0] != 0 {
		return m, GetItabResult_OK
	}

	// this can only happen if the conversion
	// was already done once using the , ok form
	// and we have a cached negative result.
	// The cached result doesn't record which
	// interface function was missing, so initialize
	// the itab again to get the missing function name.
	return m, GetItabResult_
}

// ModuleData records information about the layout of the executable
// image. It is written by the linker. Any changes here must be
// matched changes to the code in cmd/link/internal/ld/symtab.go:symtab.
// moduledata is stored in statically allocated non-pointer memory;
// none of the pointers here are visible to the garbage collector.
//
// See $GOROOT/src/runtime/symtab.go#type:moduledata
type ModuleData struct {
	_ mark.NotInHeap // Only in static data

	PCHeader     *PCHeader
	FuncNameTab  []byte
	CuTab        []uint32
	FileTab      []byte
	PCTab        []byte
	PCLnTable    []byte
	FTab         []FuncTab
	FindFuncTab  uintptr
	MinPC, MaxPC uintptr

	Text, Etext           uintptr
	NOPTRdata, ENOPTRdata uintptr
	Data, Edata           uintptr
	Bss, Ebss             uintptr
	NOPTRbss, ENOPTRbss   uintptr
	CovCtrs, ECovCtrs     uintptr // since go1.20
	End, GCdata, GCbss    uintptr
	Types, ETypes         uintptr
	ROData                uintptr
	GoFunc                uintptr // go.func.*

	TextSectMap []TextSect

	// .Types + .TypeLinks[i] = uintptr(unsafe.Pointer(*Type))
	//
	// The linker will leave a table of all the typelinks for
	// types in the binary, so the runtime can find them.
	//
	// When buildmode=shared, all types are in typelinks so the
	// runtime can deduplicate type pointers.
	//
	// Ref: ${GOROOT}/src/cmd/compile/internal/reflectdata/reflect.go#func:writeType
	TypeLinks []int32 // offsets from types
	ITabLinks []*Itab

	PTab []PTabEntry

	PluginPath string
	PkgHashes  []ModuleHash

	// This slice records the initializing tasks that need to be
	// done to start up the program. It is built by the linker.
	InitTasks []*InitTask // since go1.21

	ModuleName   string
	ModuleHashes []ModuleHash

	HasMain uint8 // 1 if module contains the main function, 0 otherwise

	GCdataMask, GCbssMask BitVector

	TypeMap map[TypeOff]*Type // offset to *_rtype in previous module

	Bad bool // module failed to load and should be ignored

	Next *ModuleData
}

// funcName returns the string at nameOff in the function name table.
func (md *ModuleData) funcName(off NameOff) string {
	if off == 0 {
		return ""
	}

	return stdstring.FromByteArray(&md.FuncNameTab[off])
}

// textAddr returns md.text + off, with special handling for multiple text sections.
// off is a (virtual) offset computed at internal linking time,
// before the external linker adjusts the sections' base addresses.
//
// The text, or instruction stream is generated as one large buffer.
// The off (offset) for a function is its offset within this buffer.
// If the total text size gets too large, there can be issues on platforms like ppc64
// if the target of calls are too far for the call instruction.
// To resolve the large text issue, the text is split into multiple text sections
// to allow the linker to generate long calls when necessary.
// When this happens, the vaddr for each text section is set to its offset within the text.
// Each function's offset is compared against the section vaddrs and ends to determine the containing section.
// Then the section relative offset is added to the section's
// relocated baseaddr to compute the function address.
//
// It is nosplit because it is part of the findfunc implementation.
//
//go:nosplit
func (md *ModuleData) textAddr(off32 uint32) uintptr {
	off := uintptr(off32)
	res := md.Text + off
	if len(md.TextSectMap) > 1 {
		for i, sect := range md.TextSectMap {
			// For the last section, include the end address (etext), as it is included in the functab.
			if off >= sect.VAddr && off < sect.End || (i == len(md.TextSectMap)-1 && off == sect.End) {
				res = sect.BaseAddr + off - sect.VAddr
				break
			}
		}
		if res > md.Etext && arch.IsWasm != 0 { // on wasm, functions do not live in the same address space as the linear memory
			println("runtime: textAddr", hex(res), "out of range", hex(md.Text), "-", hex(md.Etext))
			assert.Throw("runtime:", "text", "offset", "out", "of", "range")
		}
	}

	return res
}

// textOff is the opposite of textAddr. It converts a PC to a (virtual) offset
// to md.text, and returns if the PC is in any Go text section.
//
// It is nosplit because it is part of the findfunc implementation.
//
//go:nosplit
func (md *ModuleData) textOff(pc uintptr) (uint32, bool) {
	res := uint32(pc - md.Text)
	if len(md.TextSectMap) > 1 {
		for i, sect := range md.TextSectMap {
			if sect.BaseAddr > pc {
				// pc is not in any section.
				return 0, false
			}
			end := sect.BaseAddr + (sect.End - sect.VAddr)
			// For the last section, include the end address (etext), as it is included in the functab.
			if i == len(md.TextSectMap) {
				end++
			}
			if pc < end {
				res = uint32(pc - sect.BaseAddr + sect.VAddr)
				break
			}
		}
	}
	return res, true
}

// PCHeader holds data used by the pclntab lookups.
type PCHeader struct {
	// 0xfffffff1 (go1.20), 0xfffffff0 (go1.18) 0xfffffffa (go1.16) 0xfffffffb (go1.2)
	Magic          uint32
	pad1, pad2     uint8   // 0,0
	MinLC          uint8   // min instruction size
	PtrSize        uint8   // size of a ptr in bytes
	NFunc          int     // number of functions in the module
	NFiles         uint    // number of entries in the file tab
	TextStart      uintptr // base for function entry PC offsets in this module, equal to moduledata.text
	FuncNameOffset uintptr // offset to the funcnametab variable from pcHeader
	CuOffset       uintptr // offset to the cutab variable from pcHeader
	FileTabOffset  uintptr // offset to the filetab variable from pcHeader
	PCTabOffset    uintptr // offset to the pctab variable from pcHeader
	PCLnOffset     uintptr // offset to the pclntab variable from pcHeader
}

type FuncTab struct {
	Entryoff uint32 // relative to runtime.text
	Funcoff  uint32
}

// Mapping information for secondary text sections
type TextSect struct {
	VAddr    uintptr // prelinked section vaddr
	End      uintptr // vaddr + section length
	BaseAddr uintptr // relocated section address
}

// A PTabEntry is generated by the compiler for each exported function
// and global variable in the main package of a plugin. It is used to
// initialize the plugin module's symbol map.
type PTabEntry struct {
	Name NameOff
	Type TypeOff
}

const minfunc = 16                 // minimum function size
const pcbucketsize = 256 * minfunc // size of bucket in the pc->func lookup table

// A ModuleHash is used to compare the ABI of a new module or a
// package in a new module with the loaded program.
//
// For each shared library a module links against, the linker creates an entry in the
// moduledata.modulehashes slice containing the name of the module, the abi hash seen
// at link time and a pointer to the runtime abi hash. These are checked in
// moduledataverify1 below.
//
// For each loaded plugin, the pkghashes slice has a ModuleHash of the
// newly loaded package that can be used to check the plugin's version of
// a package against any previously loaded version of the package.
// This is done in plugin.lastmoduleinit.
type ModuleHash struct {
	ModuleName   string
	LinktimeHash string
	RuntimeHash  *string
}

// Information from the compiler about the layout of stack frames.
// Note: this type must agree with reflect.bitVector.
type BitVector struct {
	N        int32 // # of bits
	ByteData *uint8
}

// layout of Itab known to compilers
// allocated in non-garbage-collected memory
// Needs to be in sync with
// ../cmd/compile/internal/reflectdata/reflect.go:/^func.WriteTabs.
type Itab struct {
	Inter *InterfaceType
	Type  *Type
	Hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte

	// actual length of Fun = len(.Inter.Methods)
	// fun[0]==0 means _type does not implement inter.
	Fun [1]uintptr
}

func (m *Itab) FuncPCs() []unsafe.Pointer {
	return unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(&m.Fun[0])), len(m.Inter.Methods))
}

// init fills in the m.fun array with all the code pointers for
// the m.inter/m._type pair. If the type does not implement the interface,
// it sets m.fun[0] to 0 and returns the name of an interface function that is missing.
// It is ok to call this multiple times on the same m, even concurrently.
func (m *Itab) Init() string {
	inter := m.Inter
	typ := m.Type
	x := typ.Uncommon()

	// both inter and typ have method sorted by name,
	// and interface names are unique,
	// so can iterate over both in lock step;
	// the loop is O(ni+nt) not O(ni*nt).
	ni := len(inter.Methods)
	nt := int(x.Mcount)
	xmhdr := unsafe.Slice((*Method)(unsafe.Add(unsafe.Pointer(x), uintptr(x.Moff))), nt)
	j := 0
	methods := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(&m.Fun[0])), ni)
	var fun0 unsafe.Pointer
imethods:
	for k := 0; k < ni; k++ {
		i := &inter.Methods[k]
		itype := inter.Type.TypeOff(i.Typ)
		name := inter.Type.NameOff(i.Name)
		iname := name.Name()
		ipkg := pkgPath(name)
		if len(ipkg) == 0 {
			ipkg = inter.PkgPath.Name()
		}
		for ; j < nt; j++ {
			t := &xmhdr[j]
			tname := typ.NameOff(t.Name)
			if typ.TypeOff(t.Mtyp) == itype && tname.Name() == iname {
				pkgPath := pkgPath(tname)
				if len(pkgPath) == 0 {
					pkgPath = typ.NameOff(x.PkgPath).Name()
				}
				if tname.IsExported() || pkgPath == ipkg {
					if m != nil {
						ifn := typ.TextOff(t.Ifn)
						if k == 0 {
							fun0 = ifn // we'll set m.fun[0] at the end
						} else {
							methods[k] = ifn
						}
					}
					continue imethods
				}
			}
		}
		// didn't find method
		m.Fun[0] = 0
		return iname
	}
	m.Fun[0] = uintptr(fun0)
	return ""
}

func pkgPath(n Name) string {
	if n.Bytes == nil || *n.Data(0)&(1<<2) == 0 {
		return ""
	}
	i, l := n.ReadVarint(1)
	off := 1 + i + l
	if *n.Data(0)&(1<<1) != 0 {
		i2, l2 := n.ReadVarint(off)
		off += i2 + l2
	}
	var nameOff NameOff
	copy((*[4]byte)(unsafe.Pointer(&nameOff))[:], (*[4]byte)(unsafe.Pointer(n.Data(off)))[:])
	pkgPathName := resolveNameOff(unsafe.Pointer(n.Bytes), nameOff)
	return pkgPathName.Name()
}

// An InitTask represents the set of initializations that need to be done for a package.
// Keep in sync with ../../test/noinit.go:InitTask
//
// see $GOROOT/src/runtime/proc.go#type:initTask
type InitTask struct {
	State uint32 // 0 = uninitialized, 1 = in progress, 2 = done
	NFns  uint32
	// followed by nfns pcs, uintptr sized, one per init function to run
	_fns [0]uintptr // len(_fns) = NFns
}

// Nth returns the i-th init() function in the InitTask.
//
// i MUST be in range [0, tsk.NFns).
func (tsk *InitTask) Nth(i uint32) func() {
	p := unsafe.Add(
		unsafe.Pointer(tsk), unsafe.Offsetof(tsk._fns)+uintptr(i)*arch.PtrSize,
	)
	return *(*func())(unsafe.Pointer(&p))
}
