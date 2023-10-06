// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cover

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
)

// CovMetaBlob is a container for holding the meta-data symbol (an
// RODATA variable) for an instrumented Go package. Here "p" points to
// the symbol itself, "len" is the length of the sym in bytes, and
// "hash" is an md5sum for the sym computed by the compiler. When
// the init function for a coverage-instrumented package executes, it
// will make a call into the runtime which will create a covMetaBlob
// object for the package and chain it onto a global list.
type CovMetaBlob struct {
	P                  *byte
	Len                uint32
	Hash               [16]byte
	PkgPath            string
	PkgID              int
	CounterMode        uint8 // coverage.CounterMode
	CounterGranularity uint8 // coverage.CounterGranularity
}

// CovCounterBlob is a container for encapsulating a counter section
// (BSS variable) for an instrumented Go module. Here "counters"
// points to the counter payload and "len" is the number of uint32
// entries in the section.
type CovCounterBlob struct {
	Counters *uint32
	Len      uint64
}

var (
	// covMeta is the top-level container for bits of state related to
	// code coverage meta-data in the runtime.
	covMeta struct {
		// metaList contains the list of currently registered meta-data
		// blobs for the running program.
		metaList []CovMetaBlob

		// pkgMap records mappings from hard-coded package IDs to
		// slots in the covMetaList above.
		pkgMap map[int]int

		// Set to true if we discover a package mapping glitch.
		hardCodedListNeedsUpdating bool
	}
)

func GetCovMetaList() []CovMetaBlob {
	return covMeta.metaList
}

func GetCovPkgMap() map[int]int {
	return covMeta.pkgMap
}

func AppendCovCounterList(res []CovCounterBlob) {
	const (
		u32sz = unsafe.Sizeof(uint32(0))
	)

	for i, iter := 0, abi.ModuleIter(0); ; i++ {
		md, more := iter.Nth(i)
		if !more {
			break
		}

		if md.CovCtrs == md.ECovCtrs {
			continue
		}

		res = append(res, CovCounterBlob{
			Counters: (*uint32)(unsafe.Pointer(md.CovCtrs)),
			Len:      uint64((md.ECovCtrs - md.CovCtrs) / u32sz),
		})
	}
}

// AddCovMeta is invoked during package "init" functions by the
// compiler when compiling for coverage instrumentation; here 'p' is a
// meta-data blob of length 'dlen' for the package in question, 'hash'
// is a compiler-computed md5.sum for the blob, 'pkpath' is the
// package path, 'pkid' is the hard-coded ID that the compiler is
// using for the package (or -1 if the compiler doesn't think a
// hard-coded ID is needed), and 'cmode'/'cgran' are the coverage
// counter mode and granularity requested by the user. Return value is
// the ID for the package for use by the package code itself.
func AddCovMeta(p unsafe.Pointer, dlen uint32, hash [16]byte, pkpath string, pkgid int, cmode uint8, cgran uint8) uint32 {
	slot := len(covMeta.metaList)
	covMeta.metaList = append(covMeta.metaList,
		CovMetaBlob{
			P:                  (*byte)(p),
			Len:                dlen,
			Hash:               hash,
			PkgPath:            pkpath,
			PkgID:              pkgid,
			CounterMode:        cmode,
			CounterGranularity: cgran,
		},
	)
	if pkgid != -1 {
		if covMeta.pkgMap == nil {
			covMeta.pkgMap = make(map[int]int)
		}
		if _, ok := covMeta.pkgMap[pkgid]; ok {
			assert.Throw("runtime.addCovMeta:", "coverage", "package", "map", "collision")
		}
		// Record the real slot (position on meta-list) for this
		// package; we'll use the map to fix things up later on.
		covMeta.pkgMap[pkgid] = slot
	}

	// ID zero is reserved as invalid.
	return uint32(slot + 1)
}
