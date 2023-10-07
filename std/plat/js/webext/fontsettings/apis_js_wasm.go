// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package fontsettings

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/fontsettings/bindings"
)

type ClearDefaultFixedFontSizeArgDetails struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClearDefaultFixedFontSizeArgDetails with all fields set.
func (p ClearDefaultFixedFontSizeArgDetails) FromRef(ref js.Ref) ClearDefaultFixedFontSizeArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClearDefaultFixedFontSizeArgDetails in the application heap.
func (p ClearDefaultFixedFontSizeArgDetails) New() js.Ref {
	return bindings.ClearDefaultFixedFontSizeArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClearDefaultFixedFontSizeArgDetails) UpdateFrom(ref js.Ref) {
	bindings.ClearDefaultFixedFontSizeArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClearDefaultFixedFontSizeArgDetails) Update(ref js.Ref) {
	bindings.ClearDefaultFixedFontSizeArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClearDefaultFixedFontSizeArgDetails) FreeMembers(recursive bool) {
}

type ClearDefaultFontSizeArgDetails struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClearDefaultFontSizeArgDetails with all fields set.
func (p ClearDefaultFontSizeArgDetails) FromRef(ref js.Ref) ClearDefaultFontSizeArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClearDefaultFontSizeArgDetails in the application heap.
func (p ClearDefaultFontSizeArgDetails) New() js.Ref {
	return bindings.ClearDefaultFontSizeArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClearDefaultFontSizeArgDetails) UpdateFrom(ref js.Ref) {
	bindings.ClearDefaultFontSizeArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClearDefaultFontSizeArgDetails) Update(ref js.Ref) {
	bindings.ClearDefaultFontSizeArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClearDefaultFontSizeArgDetails) FreeMembers(recursive bool) {
}

type GenericFamily uint32

const (
	_ GenericFamily = iota

	GenericFamily_STANDARD
	GenericFamily_SANSSERIF
	GenericFamily_SERIF
	GenericFamily_FIXED
	GenericFamily_CURSIVE
	GenericFamily_FANTASY
	GenericFamily_MATH
)

func (GenericFamily) FromRef(str js.Ref) GenericFamily {
	return GenericFamily(bindings.ConstOfGenericFamily(str))
}

func (x GenericFamily) String() (string, bool) {
	switch x {
	case GenericFamily_STANDARD:
		return "standard", true
	case GenericFamily_SANSSERIF:
		return "sansserif", true
	case GenericFamily_SERIF:
		return "serif", true
	case GenericFamily_FIXED:
		return "fixed", true
	case GenericFamily_CURSIVE:
		return "cursive", true
	case GenericFamily_FANTASY:
		return "fantasy", true
	case GenericFamily_MATH:
		return "math", true
	default:
		return "", false
	}
}

type ScriptCode uint32

const (
	_ ScriptCode = iota

	ScriptCode_AFAK
	ScriptCode_ARAB
	ScriptCode_ARMI
	ScriptCode_ARMN
	ScriptCode_AVST
	ScriptCode_BALI
	ScriptCode_BAMU
	ScriptCode_BASS
	ScriptCode_BATK
	ScriptCode_BENG
	ScriptCode_BLIS
	ScriptCode_BOPO
	ScriptCode_BRAH
	ScriptCode_BRAI
	ScriptCode_BUGI
	ScriptCode_BUHD
	ScriptCode_CAKM
	ScriptCode_CANS
	ScriptCode_CARI
	ScriptCode_CHAM
	ScriptCode_CHER
	ScriptCode_CIRT
	ScriptCode_COPT
	ScriptCode_CPRT
	ScriptCode_CYRL
	ScriptCode_CYRS
	ScriptCode_DEVA
	ScriptCode_DSRT
	ScriptCode_DUPL
	ScriptCode_EGYD
	ScriptCode_EGYH
	ScriptCode_EGYP
	ScriptCode_ELBA
	ScriptCode_ETHI
	ScriptCode_GEOR
	ScriptCode_GEOK
	ScriptCode_GLAG
	ScriptCode_GOTH
	ScriptCode_GRAN
	ScriptCode_GREK
	ScriptCode_GUJR
	ScriptCode_GURU
	ScriptCode_HANG
	ScriptCode_HANI
	ScriptCode_HANO
	ScriptCode_HANS
	ScriptCode_HANT
	ScriptCode_HEBR
	ScriptCode_HLUW
	ScriptCode_HMNG
	ScriptCode_HUNG
	ScriptCode_INDS
	ScriptCode_ITAL
	ScriptCode_JAVA
	ScriptCode_JPAN
	ScriptCode_JURC
	ScriptCode_KALI
	ScriptCode_KHAR
	ScriptCode_KHMR
	ScriptCode_KHOJ
	ScriptCode_KNDA
	ScriptCode_KPEL
	ScriptCode_KTHI
	ScriptCode_LANA
	ScriptCode_LAOO
	ScriptCode_LATF
	ScriptCode_LATG
	ScriptCode_LATN
	ScriptCode_LEPC
	ScriptCode_LIMB
	ScriptCode_LINA
	ScriptCode_LINB
	ScriptCode_LISU
	ScriptCode_LOMA
	ScriptCode_LYCI
	ScriptCode_LYDI
	ScriptCode_MAND
	ScriptCode_MANI
	ScriptCode_MAYA
	ScriptCode_MEND
	ScriptCode_MERC
	ScriptCode_MERO
	ScriptCode_MLYM
	ScriptCode_MOON
	ScriptCode_MONG
	ScriptCode_MROO
	ScriptCode_MTEI
	ScriptCode_MYMR
	ScriptCode_NARB
	ScriptCode_NBAT
	ScriptCode_NKGB
	ScriptCode_NKOO
	ScriptCode_NSHU
	ScriptCode_OGAM
	ScriptCode_OLCK
	ScriptCode_ORKH
	ScriptCode_ORYA
	ScriptCode_OSMA
	ScriptCode_PALM
	ScriptCode_PERM
	ScriptCode_PHAG
	ScriptCode_PHLI
	ScriptCode_PHLP
	ScriptCode_PHLV
	ScriptCode_PHNX
	ScriptCode_PLRD
	ScriptCode_PRTI
	ScriptCode_RJNG
	ScriptCode_RORO
	ScriptCode_RUNR
	ScriptCode_SAMR
	ScriptCode_SARA
	ScriptCode_SARB
	ScriptCode_SAUR
	ScriptCode_SGNW
	ScriptCode_SHAW
	ScriptCode_SHRD
	ScriptCode_SIND
	ScriptCode_SINH
	ScriptCode_SORA
	ScriptCode_SUND
	ScriptCode_SYLO
	ScriptCode_SYRC
	ScriptCode_SYRE
	ScriptCode_SYRJ
	ScriptCode_SYRN
	ScriptCode_TAGB
	ScriptCode_TAKR
	ScriptCode_TALE
	ScriptCode_TALU
	ScriptCode_TAML
	ScriptCode_TANG
	ScriptCode_TAVT
	ScriptCode_TELU
	ScriptCode_TENG
	ScriptCode_TFNG
	ScriptCode_TGLG
	ScriptCode_THAA
	ScriptCode_THAI
	ScriptCode_TIBT
	ScriptCode_TIRH
	ScriptCode_UGAR
	ScriptCode_VAII
	ScriptCode_VISP
	ScriptCode_WARA
	ScriptCode_WOLE
	ScriptCode_XPEO
	ScriptCode_XSUX
	ScriptCode_YIII
	ScriptCode_ZMTH
	ScriptCode_ZSYM
	ScriptCode_ZYYY
)

func (ScriptCode) FromRef(str js.Ref) ScriptCode {
	return ScriptCode(bindings.ConstOfScriptCode(str))
}

func (x ScriptCode) String() (string, bool) {
	switch x {
	case ScriptCode_AFAK:
		return "Afak", true
	case ScriptCode_ARAB:
		return "Arab", true
	case ScriptCode_ARMI:
		return "Armi", true
	case ScriptCode_ARMN:
		return "Armn", true
	case ScriptCode_AVST:
		return "Avst", true
	case ScriptCode_BALI:
		return "Bali", true
	case ScriptCode_BAMU:
		return "Bamu", true
	case ScriptCode_BASS:
		return "Bass", true
	case ScriptCode_BATK:
		return "Batk", true
	case ScriptCode_BENG:
		return "Beng", true
	case ScriptCode_BLIS:
		return "Blis", true
	case ScriptCode_BOPO:
		return "Bopo", true
	case ScriptCode_BRAH:
		return "Brah", true
	case ScriptCode_BRAI:
		return "Brai", true
	case ScriptCode_BUGI:
		return "Bugi", true
	case ScriptCode_BUHD:
		return "Buhd", true
	case ScriptCode_CAKM:
		return "Cakm", true
	case ScriptCode_CANS:
		return "Cans", true
	case ScriptCode_CARI:
		return "Cari", true
	case ScriptCode_CHAM:
		return "Cham", true
	case ScriptCode_CHER:
		return "Cher", true
	case ScriptCode_CIRT:
		return "Cirt", true
	case ScriptCode_COPT:
		return "Copt", true
	case ScriptCode_CPRT:
		return "Cprt", true
	case ScriptCode_CYRL:
		return "Cyrl", true
	case ScriptCode_CYRS:
		return "Cyrs", true
	case ScriptCode_DEVA:
		return "Deva", true
	case ScriptCode_DSRT:
		return "Dsrt", true
	case ScriptCode_DUPL:
		return "Dupl", true
	case ScriptCode_EGYD:
		return "Egyd", true
	case ScriptCode_EGYH:
		return "Egyh", true
	case ScriptCode_EGYP:
		return "Egyp", true
	case ScriptCode_ELBA:
		return "Elba", true
	case ScriptCode_ETHI:
		return "Ethi", true
	case ScriptCode_GEOR:
		return "Geor", true
	case ScriptCode_GEOK:
		return "Geok", true
	case ScriptCode_GLAG:
		return "Glag", true
	case ScriptCode_GOTH:
		return "Goth", true
	case ScriptCode_GRAN:
		return "Gran", true
	case ScriptCode_GREK:
		return "Grek", true
	case ScriptCode_GUJR:
		return "Gujr", true
	case ScriptCode_GURU:
		return "Guru", true
	case ScriptCode_HANG:
		return "Hang", true
	case ScriptCode_HANI:
		return "Hani", true
	case ScriptCode_HANO:
		return "Hano", true
	case ScriptCode_HANS:
		return "Hans", true
	case ScriptCode_HANT:
		return "Hant", true
	case ScriptCode_HEBR:
		return "Hebr", true
	case ScriptCode_HLUW:
		return "Hluw", true
	case ScriptCode_HMNG:
		return "Hmng", true
	case ScriptCode_HUNG:
		return "Hung", true
	case ScriptCode_INDS:
		return "Inds", true
	case ScriptCode_ITAL:
		return "Ital", true
	case ScriptCode_JAVA:
		return "Java", true
	case ScriptCode_JPAN:
		return "Jpan", true
	case ScriptCode_JURC:
		return "Jurc", true
	case ScriptCode_KALI:
		return "Kali", true
	case ScriptCode_KHAR:
		return "Khar", true
	case ScriptCode_KHMR:
		return "Khmr", true
	case ScriptCode_KHOJ:
		return "Khoj", true
	case ScriptCode_KNDA:
		return "Knda", true
	case ScriptCode_KPEL:
		return "Kpel", true
	case ScriptCode_KTHI:
		return "Kthi", true
	case ScriptCode_LANA:
		return "Lana", true
	case ScriptCode_LAOO:
		return "Laoo", true
	case ScriptCode_LATF:
		return "Latf", true
	case ScriptCode_LATG:
		return "Latg", true
	case ScriptCode_LATN:
		return "Latn", true
	case ScriptCode_LEPC:
		return "Lepc", true
	case ScriptCode_LIMB:
		return "Limb", true
	case ScriptCode_LINA:
		return "Lina", true
	case ScriptCode_LINB:
		return "Linb", true
	case ScriptCode_LISU:
		return "Lisu", true
	case ScriptCode_LOMA:
		return "Loma", true
	case ScriptCode_LYCI:
		return "Lyci", true
	case ScriptCode_LYDI:
		return "Lydi", true
	case ScriptCode_MAND:
		return "Mand", true
	case ScriptCode_MANI:
		return "Mani", true
	case ScriptCode_MAYA:
		return "Maya", true
	case ScriptCode_MEND:
		return "Mend", true
	case ScriptCode_MERC:
		return "Merc", true
	case ScriptCode_MERO:
		return "Mero", true
	case ScriptCode_MLYM:
		return "Mlym", true
	case ScriptCode_MOON:
		return "Moon", true
	case ScriptCode_MONG:
		return "Mong", true
	case ScriptCode_MROO:
		return "Mroo", true
	case ScriptCode_MTEI:
		return "Mtei", true
	case ScriptCode_MYMR:
		return "Mymr", true
	case ScriptCode_NARB:
		return "Narb", true
	case ScriptCode_NBAT:
		return "Nbat", true
	case ScriptCode_NKGB:
		return "Nkgb", true
	case ScriptCode_NKOO:
		return "Nkoo", true
	case ScriptCode_NSHU:
		return "Nshu", true
	case ScriptCode_OGAM:
		return "Ogam", true
	case ScriptCode_OLCK:
		return "Olck", true
	case ScriptCode_ORKH:
		return "Orkh", true
	case ScriptCode_ORYA:
		return "Orya", true
	case ScriptCode_OSMA:
		return "Osma", true
	case ScriptCode_PALM:
		return "Palm", true
	case ScriptCode_PERM:
		return "Perm", true
	case ScriptCode_PHAG:
		return "Phag", true
	case ScriptCode_PHLI:
		return "Phli", true
	case ScriptCode_PHLP:
		return "Phlp", true
	case ScriptCode_PHLV:
		return "Phlv", true
	case ScriptCode_PHNX:
		return "Phnx", true
	case ScriptCode_PLRD:
		return "Plrd", true
	case ScriptCode_PRTI:
		return "Prti", true
	case ScriptCode_RJNG:
		return "Rjng", true
	case ScriptCode_RORO:
		return "Roro", true
	case ScriptCode_RUNR:
		return "Runr", true
	case ScriptCode_SAMR:
		return "Samr", true
	case ScriptCode_SARA:
		return "Sara", true
	case ScriptCode_SARB:
		return "Sarb", true
	case ScriptCode_SAUR:
		return "Saur", true
	case ScriptCode_SGNW:
		return "Sgnw", true
	case ScriptCode_SHAW:
		return "Shaw", true
	case ScriptCode_SHRD:
		return "Shrd", true
	case ScriptCode_SIND:
		return "Sind", true
	case ScriptCode_SINH:
		return "Sinh", true
	case ScriptCode_SORA:
		return "Sora", true
	case ScriptCode_SUND:
		return "Sund", true
	case ScriptCode_SYLO:
		return "Sylo", true
	case ScriptCode_SYRC:
		return "Syrc", true
	case ScriptCode_SYRE:
		return "Syre", true
	case ScriptCode_SYRJ:
		return "Syrj", true
	case ScriptCode_SYRN:
		return "Syrn", true
	case ScriptCode_TAGB:
		return "Tagb", true
	case ScriptCode_TAKR:
		return "Takr", true
	case ScriptCode_TALE:
		return "Tale", true
	case ScriptCode_TALU:
		return "Talu", true
	case ScriptCode_TAML:
		return "Taml", true
	case ScriptCode_TANG:
		return "Tang", true
	case ScriptCode_TAVT:
		return "Tavt", true
	case ScriptCode_TELU:
		return "Telu", true
	case ScriptCode_TENG:
		return "Teng", true
	case ScriptCode_TFNG:
		return "Tfng", true
	case ScriptCode_TGLG:
		return "Tglg", true
	case ScriptCode_THAA:
		return "Thaa", true
	case ScriptCode_THAI:
		return "Thai", true
	case ScriptCode_TIBT:
		return "Tibt", true
	case ScriptCode_TIRH:
		return "Tirh", true
	case ScriptCode_UGAR:
		return "Ugar", true
	case ScriptCode_VAII:
		return "Vaii", true
	case ScriptCode_VISP:
		return "Visp", true
	case ScriptCode_WARA:
		return "Wara", true
	case ScriptCode_WOLE:
		return "Wole", true
	case ScriptCode_XPEO:
		return "Xpeo", true
	case ScriptCode_XSUX:
		return "Xsux", true
	case ScriptCode_YIII:
		return "Yiii", true
	case ScriptCode_ZMTH:
		return "Zmth", true
	case ScriptCode_ZSYM:
		return "Zsym", true
	case ScriptCode_ZYYY:
		return "Zyyy", true
	default:
		return "", false
	}
}

type ClearFontArgDetails struct {
	// GenericFamily is "ClearFontArgDetails.genericFamily"
	//
	// Required
	GenericFamily GenericFamily
	// Script is "ClearFontArgDetails.script"
	//
	// Optional
	Script ScriptCode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClearFontArgDetails with all fields set.
func (p ClearFontArgDetails) FromRef(ref js.Ref) ClearFontArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClearFontArgDetails in the application heap.
func (p ClearFontArgDetails) New() js.Ref {
	return bindings.ClearFontArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClearFontArgDetails) UpdateFrom(ref js.Ref) {
	bindings.ClearFontArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClearFontArgDetails) Update(ref js.Ref) {
	bindings.ClearFontArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClearFontArgDetails) FreeMembers(recursive bool) {
}

type ClearMinimumFontSizeArgDetails struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClearMinimumFontSizeArgDetails with all fields set.
func (p ClearMinimumFontSizeArgDetails) FromRef(ref js.Ref) ClearMinimumFontSizeArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClearMinimumFontSizeArgDetails in the application heap.
func (p ClearMinimumFontSizeArgDetails) New() js.Ref {
	return bindings.ClearMinimumFontSizeArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClearMinimumFontSizeArgDetails) UpdateFrom(ref js.Ref) {
	bindings.ClearMinimumFontSizeArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClearMinimumFontSizeArgDetails) Update(ref js.Ref) {
	bindings.ClearMinimumFontSizeArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClearMinimumFontSizeArgDetails) FreeMembers(recursive bool) {
}

type FontName struct {
	// DisplayName is "FontName.displayName"
	//
	// Required
	DisplayName js.String
	// FontId is "FontName.fontId"
	//
	// Required
	FontId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FontName with all fields set.
func (p FontName) FromRef(ref js.Ref) FontName {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FontName in the application heap.
func (p FontName) New() js.Ref {
	return bindings.FontNameJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FontName) UpdateFrom(ref js.Ref) {
	bindings.FontNameJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FontName) Update(ref js.Ref) {
	bindings.FontNameJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FontName) FreeMembers(recursive bool) {
	js.Free(
		p.DisplayName.Ref(),
		p.FontId.Ref(),
	)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.FontId = p.FontId.FromRef(js.Undefined)
}

type GetDefaultFixedFontSizeArgDetails struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetDefaultFixedFontSizeArgDetails with all fields set.
func (p GetDefaultFixedFontSizeArgDetails) FromRef(ref js.Ref) GetDefaultFixedFontSizeArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetDefaultFixedFontSizeArgDetails in the application heap.
func (p GetDefaultFixedFontSizeArgDetails) New() js.Ref {
	return bindings.GetDefaultFixedFontSizeArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetDefaultFixedFontSizeArgDetails) UpdateFrom(ref js.Ref) {
	bindings.GetDefaultFixedFontSizeArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetDefaultFixedFontSizeArgDetails) Update(ref js.Ref) {
	bindings.GetDefaultFixedFontSizeArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetDefaultFixedFontSizeArgDetails) FreeMembers(recursive bool) {
}

type LevelOfControl uint32

const (
	_ LevelOfControl = iota

	LevelOfControl_NOT_CONTROLLABLE
	LevelOfControl_CONTROLLED_BY_OTHER_EXTENSIONS
	LevelOfControl_CONTROLLABLE_BY_THIS_EXTENSION
	LevelOfControl_CONTROLLED_BY_THIS_EXTENSION
)

func (LevelOfControl) FromRef(str js.Ref) LevelOfControl {
	return LevelOfControl(bindings.ConstOfLevelOfControl(str))
}

func (x LevelOfControl) String() (string, bool) {
	switch x {
	case LevelOfControl_NOT_CONTROLLABLE:
		return "not_controllable", true
	case LevelOfControl_CONTROLLED_BY_OTHER_EXTENSIONS:
		return "controlled_by_other_extensions", true
	case LevelOfControl_CONTROLLABLE_BY_THIS_EXTENSION:
		return "controllable_by_this_extension", true
	case LevelOfControl_CONTROLLED_BY_THIS_EXTENSION:
		return "controlled_by_this_extension", true
	default:
		return "", false
	}
}

type GetDefaultFixedFontSizeReturnType struct {
	// LevelOfControl is "GetDefaultFixedFontSizeReturnType.levelOfControl"
	//
	// Required
	LevelOfControl LevelOfControl
	// PixelSize is "GetDefaultFixedFontSizeReturnType.pixelSize"
	//
	// Required
	PixelSize int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetDefaultFixedFontSizeReturnType with all fields set.
func (p GetDefaultFixedFontSizeReturnType) FromRef(ref js.Ref) GetDefaultFixedFontSizeReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetDefaultFixedFontSizeReturnType in the application heap.
func (p GetDefaultFixedFontSizeReturnType) New() js.Ref {
	return bindings.GetDefaultFixedFontSizeReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetDefaultFixedFontSizeReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetDefaultFixedFontSizeReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetDefaultFixedFontSizeReturnType) Update(ref js.Ref) {
	bindings.GetDefaultFixedFontSizeReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetDefaultFixedFontSizeReturnType) FreeMembers(recursive bool) {
}

type GetDefaultFontSizeArgDetails struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetDefaultFontSizeArgDetails with all fields set.
func (p GetDefaultFontSizeArgDetails) FromRef(ref js.Ref) GetDefaultFontSizeArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetDefaultFontSizeArgDetails in the application heap.
func (p GetDefaultFontSizeArgDetails) New() js.Ref {
	return bindings.GetDefaultFontSizeArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetDefaultFontSizeArgDetails) UpdateFrom(ref js.Ref) {
	bindings.GetDefaultFontSizeArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetDefaultFontSizeArgDetails) Update(ref js.Ref) {
	bindings.GetDefaultFontSizeArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetDefaultFontSizeArgDetails) FreeMembers(recursive bool) {
}

type GetDefaultFontSizeReturnType struct {
	// LevelOfControl is "GetDefaultFontSizeReturnType.levelOfControl"
	//
	// Required
	LevelOfControl LevelOfControl
	// PixelSize is "GetDefaultFontSizeReturnType.pixelSize"
	//
	// Required
	PixelSize int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetDefaultFontSizeReturnType with all fields set.
func (p GetDefaultFontSizeReturnType) FromRef(ref js.Ref) GetDefaultFontSizeReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetDefaultFontSizeReturnType in the application heap.
func (p GetDefaultFontSizeReturnType) New() js.Ref {
	return bindings.GetDefaultFontSizeReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetDefaultFontSizeReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetDefaultFontSizeReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetDefaultFontSizeReturnType) Update(ref js.Ref) {
	bindings.GetDefaultFontSizeReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetDefaultFontSizeReturnType) FreeMembers(recursive bool) {
}

type GetFontArgDetails struct {
	// GenericFamily is "GetFontArgDetails.genericFamily"
	//
	// Required
	GenericFamily GenericFamily
	// Script is "GetFontArgDetails.script"
	//
	// Optional
	Script ScriptCode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetFontArgDetails with all fields set.
func (p GetFontArgDetails) FromRef(ref js.Ref) GetFontArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetFontArgDetails in the application heap.
func (p GetFontArgDetails) New() js.Ref {
	return bindings.GetFontArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetFontArgDetails) UpdateFrom(ref js.Ref) {
	bindings.GetFontArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetFontArgDetails) Update(ref js.Ref) {
	bindings.GetFontArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetFontArgDetails) FreeMembers(recursive bool) {
}

type GetFontReturnType struct {
	// FontId is "GetFontReturnType.fontId"
	//
	// Required
	FontId js.String
	// LevelOfControl is "GetFontReturnType.levelOfControl"
	//
	// Required
	LevelOfControl LevelOfControl

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetFontReturnType with all fields set.
func (p GetFontReturnType) FromRef(ref js.Ref) GetFontReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetFontReturnType in the application heap.
func (p GetFontReturnType) New() js.Ref {
	return bindings.GetFontReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetFontReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetFontReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetFontReturnType) Update(ref js.Ref) {
	bindings.GetFontReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetFontReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.FontId.Ref(),
	)
	p.FontId = p.FontId.FromRef(js.Undefined)
}

type GetMinimumFontSizeArgDetails struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetMinimumFontSizeArgDetails with all fields set.
func (p GetMinimumFontSizeArgDetails) FromRef(ref js.Ref) GetMinimumFontSizeArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetMinimumFontSizeArgDetails in the application heap.
func (p GetMinimumFontSizeArgDetails) New() js.Ref {
	return bindings.GetMinimumFontSizeArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetMinimumFontSizeArgDetails) UpdateFrom(ref js.Ref) {
	bindings.GetMinimumFontSizeArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetMinimumFontSizeArgDetails) Update(ref js.Ref) {
	bindings.GetMinimumFontSizeArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetMinimumFontSizeArgDetails) FreeMembers(recursive bool) {
}

type GetMinimumFontSizeReturnType struct {
	// LevelOfControl is "GetMinimumFontSizeReturnType.levelOfControl"
	//
	// Required
	LevelOfControl LevelOfControl
	// PixelSize is "GetMinimumFontSizeReturnType.pixelSize"
	//
	// Required
	PixelSize int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetMinimumFontSizeReturnType with all fields set.
func (p GetMinimumFontSizeReturnType) FromRef(ref js.Ref) GetMinimumFontSizeReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetMinimumFontSizeReturnType in the application heap.
func (p GetMinimumFontSizeReturnType) New() js.Ref {
	return bindings.GetMinimumFontSizeReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetMinimumFontSizeReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetMinimumFontSizeReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetMinimumFontSizeReturnType) Update(ref js.Ref) {
	bindings.GetMinimumFontSizeReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetMinimumFontSizeReturnType) FreeMembers(recursive bool) {
}

type OnDefaultFixedFontSizeChangedArgDetails struct {
	// LevelOfControl is "OnDefaultFixedFontSizeChangedArgDetails.levelOfControl"
	//
	// Required
	LevelOfControl LevelOfControl
	// PixelSize is "OnDefaultFixedFontSizeChangedArgDetails.pixelSize"
	//
	// Required
	PixelSize int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnDefaultFixedFontSizeChangedArgDetails with all fields set.
func (p OnDefaultFixedFontSizeChangedArgDetails) FromRef(ref js.Ref) OnDefaultFixedFontSizeChangedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnDefaultFixedFontSizeChangedArgDetails in the application heap.
func (p OnDefaultFixedFontSizeChangedArgDetails) New() js.Ref {
	return bindings.OnDefaultFixedFontSizeChangedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnDefaultFixedFontSizeChangedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnDefaultFixedFontSizeChangedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnDefaultFixedFontSizeChangedArgDetails) Update(ref js.Ref) {
	bindings.OnDefaultFixedFontSizeChangedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnDefaultFixedFontSizeChangedArgDetails) FreeMembers(recursive bool) {
}

type OnDefaultFontSizeChangedArgDetails struct {
	// LevelOfControl is "OnDefaultFontSizeChangedArgDetails.levelOfControl"
	//
	// Required
	LevelOfControl LevelOfControl
	// PixelSize is "OnDefaultFontSizeChangedArgDetails.pixelSize"
	//
	// Required
	PixelSize int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnDefaultFontSizeChangedArgDetails with all fields set.
func (p OnDefaultFontSizeChangedArgDetails) FromRef(ref js.Ref) OnDefaultFontSizeChangedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnDefaultFontSizeChangedArgDetails in the application heap.
func (p OnDefaultFontSizeChangedArgDetails) New() js.Ref {
	return bindings.OnDefaultFontSizeChangedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnDefaultFontSizeChangedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnDefaultFontSizeChangedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnDefaultFontSizeChangedArgDetails) Update(ref js.Ref) {
	bindings.OnDefaultFontSizeChangedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnDefaultFontSizeChangedArgDetails) FreeMembers(recursive bool) {
}

type OnFontChangedArgDetails struct {
	// FontId is "OnFontChangedArgDetails.fontId"
	//
	// Required
	FontId js.String
	// GenericFamily is "OnFontChangedArgDetails.genericFamily"
	//
	// Required
	GenericFamily GenericFamily
	// LevelOfControl is "OnFontChangedArgDetails.levelOfControl"
	//
	// Required
	LevelOfControl LevelOfControl
	// Script is "OnFontChangedArgDetails.script"
	//
	// Optional
	Script ScriptCode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnFontChangedArgDetails with all fields set.
func (p OnFontChangedArgDetails) FromRef(ref js.Ref) OnFontChangedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnFontChangedArgDetails in the application heap.
func (p OnFontChangedArgDetails) New() js.Ref {
	return bindings.OnFontChangedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnFontChangedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnFontChangedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnFontChangedArgDetails) Update(ref js.Ref) {
	bindings.OnFontChangedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnFontChangedArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.FontId.Ref(),
	)
	p.FontId = p.FontId.FromRef(js.Undefined)
}

type OnMinimumFontSizeChangedArgDetails struct {
	// LevelOfControl is "OnMinimumFontSizeChangedArgDetails.levelOfControl"
	//
	// Required
	LevelOfControl LevelOfControl
	// PixelSize is "OnMinimumFontSizeChangedArgDetails.pixelSize"
	//
	// Required
	PixelSize int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnMinimumFontSizeChangedArgDetails with all fields set.
func (p OnMinimumFontSizeChangedArgDetails) FromRef(ref js.Ref) OnMinimumFontSizeChangedArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnMinimumFontSizeChangedArgDetails in the application heap.
func (p OnMinimumFontSizeChangedArgDetails) New() js.Ref {
	return bindings.OnMinimumFontSizeChangedArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnMinimumFontSizeChangedArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnMinimumFontSizeChangedArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnMinimumFontSizeChangedArgDetails) Update(ref js.Ref) {
	bindings.OnMinimumFontSizeChangedArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnMinimumFontSizeChangedArgDetails) FreeMembers(recursive bool) {
}

type SetDefaultFixedFontSizeArgDetails struct {
	// PixelSize is "SetDefaultFixedFontSizeArgDetails.pixelSize"
	//
	// Required
	PixelSize int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetDefaultFixedFontSizeArgDetails with all fields set.
func (p SetDefaultFixedFontSizeArgDetails) FromRef(ref js.Ref) SetDefaultFixedFontSizeArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetDefaultFixedFontSizeArgDetails in the application heap.
func (p SetDefaultFixedFontSizeArgDetails) New() js.Ref {
	return bindings.SetDefaultFixedFontSizeArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetDefaultFixedFontSizeArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetDefaultFixedFontSizeArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetDefaultFixedFontSizeArgDetails) Update(ref js.Ref) {
	bindings.SetDefaultFixedFontSizeArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetDefaultFixedFontSizeArgDetails) FreeMembers(recursive bool) {
}

type SetDefaultFontSizeArgDetails struct {
	// PixelSize is "SetDefaultFontSizeArgDetails.pixelSize"
	//
	// Required
	PixelSize int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetDefaultFontSizeArgDetails with all fields set.
func (p SetDefaultFontSizeArgDetails) FromRef(ref js.Ref) SetDefaultFontSizeArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetDefaultFontSizeArgDetails in the application heap.
func (p SetDefaultFontSizeArgDetails) New() js.Ref {
	return bindings.SetDefaultFontSizeArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetDefaultFontSizeArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetDefaultFontSizeArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetDefaultFontSizeArgDetails) Update(ref js.Ref) {
	bindings.SetDefaultFontSizeArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetDefaultFontSizeArgDetails) FreeMembers(recursive bool) {
}

type SetFontArgDetails struct {
	// FontId is "SetFontArgDetails.fontId"
	//
	// Required
	FontId js.String
	// GenericFamily is "SetFontArgDetails.genericFamily"
	//
	// Required
	GenericFamily GenericFamily
	// Script is "SetFontArgDetails.script"
	//
	// Optional
	Script ScriptCode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetFontArgDetails with all fields set.
func (p SetFontArgDetails) FromRef(ref js.Ref) SetFontArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetFontArgDetails in the application heap.
func (p SetFontArgDetails) New() js.Ref {
	return bindings.SetFontArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetFontArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetFontArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetFontArgDetails) Update(ref js.Ref) {
	bindings.SetFontArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetFontArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.FontId.Ref(),
	)
	p.FontId = p.FontId.FromRef(js.Undefined)
}

type SetMinimumFontSizeArgDetails struct {
	// PixelSize is "SetMinimumFontSizeArgDetails.pixelSize"
	//
	// Required
	PixelSize int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetMinimumFontSizeArgDetails with all fields set.
func (p SetMinimumFontSizeArgDetails) FromRef(ref js.Ref) SetMinimumFontSizeArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetMinimumFontSizeArgDetails in the application heap.
func (p SetMinimumFontSizeArgDetails) New() js.Ref {
	return bindings.SetMinimumFontSizeArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetMinimumFontSizeArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetMinimumFontSizeArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetMinimumFontSizeArgDetails) Update(ref js.Ref) {
	bindings.SetMinimumFontSizeArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetMinimumFontSizeArgDetails) FreeMembers(recursive bool) {
}

// HasFuncClearDefaultFixedFontSize returns true if the function "WEBEXT.fontSettings.clearDefaultFixedFontSize" exists.
func HasFuncClearDefaultFixedFontSize() bool {
	return js.True == bindings.HasFuncClearDefaultFixedFontSize()
}

// FuncClearDefaultFixedFontSize returns the function "WEBEXT.fontSettings.clearDefaultFixedFontSize".
func FuncClearDefaultFixedFontSize() (fn js.Func[func(details ClearDefaultFixedFontSizeArgDetails) js.Promise[js.Void]]) {
	bindings.FuncClearDefaultFixedFontSize(
		js.Pointer(&fn),
	)
	return
}

// ClearDefaultFixedFontSize calls the function "WEBEXT.fontSettings.clearDefaultFixedFontSize" directly.
func ClearDefaultFixedFontSize(details ClearDefaultFixedFontSizeArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallClearDefaultFixedFontSize(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryClearDefaultFixedFontSize calls the function "WEBEXT.fontSettings.clearDefaultFixedFontSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearDefaultFixedFontSize(details ClearDefaultFixedFontSizeArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearDefaultFixedFontSize(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncClearDefaultFontSize returns true if the function "WEBEXT.fontSettings.clearDefaultFontSize" exists.
func HasFuncClearDefaultFontSize() bool {
	return js.True == bindings.HasFuncClearDefaultFontSize()
}

// FuncClearDefaultFontSize returns the function "WEBEXT.fontSettings.clearDefaultFontSize".
func FuncClearDefaultFontSize() (fn js.Func[func(details ClearDefaultFontSizeArgDetails) js.Promise[js.Void]]) {
	bindings.FuncClearDefaultFontSize(
		js.Pointer(&fn),
	)
	return
}

// ClearDefaultFontSize calls the function "WEBEXT.fontSettings.clearDefaultFontSize" directly.
func ClearDefaultFontSize(details ClearDefaultFontSizeArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallClearDefaultFontSize(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryClearDefaultFontSize calls the function "WEBEXT.fontSettings.clearDefaultFontSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearDefaultFontSize(details ClearDefaultFontSizeArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearDefaultFontSize(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncClearFont returns true if the function "WEBEXT.fontSettings.clearFont" exists.
func HasFuncClearFont() bool {
	return js.True == bindings.HasFuncClearFont()
}

// FuncClearFont returns the function "WEBEXT.fontSettings.clearFont".
func FuncClearFont() (fn js.Func[func(details ClearFontArgDetails) js.Promise[js.Void]]) {
	bindings.FuncClearFont(
		js.Pointer(&fn),
	)
	return
}

// ClearFont calls the function "WEBEXT.fontSettings.clearFont" directly.
func ClearFont(details ClearFontArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallClearFont(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryClearFont calls the function "WEBEXT.fontSettings.clearFont"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearFont(details ClearFontArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearFont(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncClearMinimumFontSize returns true if the function "WEBEXT.fontSettings.clearMinimumFontSize" exists.
func HasFuncClearMinimumFontSize() bool {
	return js.True == bindings.HasFuncClearMinimumFontSize()
}

// FuncClearMinimumFontSize returns the function "WEBEXT.fontSettings.clearMinimumFontSize".
func FuncClearMinimumFontSize() (fn js.Func[func(details ClearMinimumFontSizeArgDetails) js.Promise[js.Void]]) {
	bindings.FuncClearMinimumFontSize(
		js.Pointer(&fn),
	)
	return
}

// ClearMinimumFontSize calls the function "WEBEXT.fontSettings.clearMinimumFontSize" directly.
func ClearMinimumFontSize(details ClearMinimumFontSizeArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallClearMinimumFontSize(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryClearMinimumFontSize calls the function "WEBEXT.fontSettings.clearMinimumFontSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearMinimumFontSize(details ClearMinimumFontSizeArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearMinimumFontSize(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetDefaultFixedFontSize returns true if the function "WEBEXT.fontSettings.getDefaultFixedFontSize" exists.
func HasFuncGetDefaultFixedFontSize() bool {
	return js.True == bindings.HasFuncGetDefaultFixedFontSize()
}

// FuncGetDefaultFixedFontSize returns the function "WEBEXT.fontSettings.getDefaultFixedFontSize".
func FuncGetDefaultFixedFontSize() (fn js.Func[func(details GetDefaultFixedFontSizeArgDetails) js.Promise[GetDefaultFixedFontSizeReturnType]]) {
	bindings.FuncGetDefaultFixedFontSize(
		js.Pointer(&fn),
	)
	return
}

// GetDefaultFixedFontSize calls the function "WEBEXT.fontSettings.getDefaultFixedFontSize" directly.
func GetDefaultFixedFontSize(details GetDefaultFixedFontSizeArgDetails) (ret js.Promise[GetDefaultFixedFontSizeReturnType]) {
	bindings.CallGetDefaultFixedFontSize(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetDefaultFixedFontSize calls the function "WEBEXT.fontSettings.getDefaultFixedFontSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDefaultFixedFontSize(details GetDefaultFixedFontSizeArgDetails) (ret js.Promise[GetDefaultFixedFontSizeReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDefaultFixedFontSize(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetDefaultFontSize returns true if the function "WEBEXT.fontSettings.getDefaultFontSize" exists.
func HasFuncGetDefaultFontSize() bool {
	return js.True == bindings.HasFuncGetDefaultFontSize()
}

// FuncGetDefaultFontSize returns the function "WEBEXT.fontSettings.getDefaultFontSize".
func FuncGetDefaultFontSize() (fn js.Func[func(details GetDefaultFontSizeArgDetails) js.Promise[GetDefaultFontSizeReturnType]]) {
	bindings.FuncGetDefaultFontSize(
		js.Pointer(&fn),
	)
	return
}

// GetDefaultFontSize calls the function "WEBEXT.fontSettings.getDefaultFontSize" directly.
func GetDefaultFontSize(details GetDefaultFontSizeArgDetails) (ret js.Promise[GetDefaultFontSizeReturnType]) {
	bindings.CallGetDefaultFontSize(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetDefaultFontSize calls the function "WEBEXT.fontSettings.getDefaultFontSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDefaultFontSize(details GetDefaultFontSizeArgDetails) (ret js.Promise[GetDefaultFontSizeReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDefaultFontSize(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetFont returns true if the function "WEBEXT.fontSettings.getFont" exists.
func HasFuncGetFont() bool {
	return js.True == bindings.HasFuncGetFont()
}

// FuncGetFont returns the function "WEBEXT.fontSettings.getFont".
func FuncGetFont() (fn js.Func[func(details GetFontArgDetails) js.Promise[GetFontReturnType]]) {
	bindings.FuncGetFont(
		js.Pointer(&fn),
	)
	return
}

// GetFont calls the function "WEBEXT.fontSettings.getFont" directly.
func GetFont(details GetFontArgDetails) (ret js.Promise[GetFontReturnType]) {
	bindings.CallGetFont(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetFont calls the function "WEBEXT.fontSettings.getFont"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetFont(details GetFontArgDetails) (ret js.Promise[GetFontReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetFont(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetFontList returns true if the function "WEBEXT.fontSettings.getFontList" exists.
func HasFuncGetFontList() bool {
	return js.True == bindings.HasFuncGetFontList()
}

// FuncGetFontList returns the function "WEBEXT.fontSettings.getFontList".
func FuncGetFontList() (fn js.Func[func() js.Promise[js.Array[FontName]]]) {
	bindings.FuncGetFontList(
		js.Pointer(&fn),
	)
	return
}

// GetFontList calls the function "WEBEXT.fontSettings.getFontList" directly.
func GetFontList() (ret js.Promise[js.Array[FontName]]) {
	bindings.CallGetFontList(
		js.Pointer(&ret),
	)

	return
}

// TryGetFontList calls the function "WEBEXT.fontSettings.getFontList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetFontList() (ret js.Promise[js.Array[FontName]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetFontList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetMinimumFontSize returns true if the function "WEBEXT.fontSettings.getMinimumFontSize" exists.
func HasFuncGetMinimumFontSize() bool {
	return js.True == bindings.HasFuncGetMinimumFontSize()
}

// FuncGetMinimumFontSize returns the function "WEBEXT.fontSettings.getMinimumFontSize".
func FuncGetMinimumFontSize() (fn js.Func[func(details GetMinimumFontSizeArgDetails) js.Promise[GetMinimumFontSizeReturnType]]) {
	bindings.FuncGetMinimumFontSize(
		js.Pointer(&fn),
	)
	return
}

// GetMinimumFontSize calls the function "WEBEXT.fontSettings.getMinimumFontSize" directly.
func GetMinimumFontSize(details GetMinimumFontSizeArgDetails) (ret js.Promise[GetMinimumFontSizeReturnType]) {
	bindings.CallGetMinimumFontSize(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetMinimumFontSize calls the function "WEBEXT.fontSettings.getMinimumFontSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMinimumFontSize(details GetMinimumFontSizeArgDetails) (ret js.Promise[GetMinimumFontSizeReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMinimumFontSize(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

type OnDefaultFixedFontSizeChangedEventCallbackFunc func(this js.Ref, details *OnDefaultFixedFontSizeChangedArgDetails) js.Ref

func (fn OnDefaultFixedFontSizeChangedEventCallbackFunc) Register() js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)] {
	return js.RegisterCallback[func(details *OnDefaultFixedFontSizeChangedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDefaultFixedFontSizeChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnDefaultFixedFontSizeChangedArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDefaultFixedFontSizeChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnDefaultFixedFontSizeChangedArgDetails) js.Ref
	Arg T
}

func (cb *OnDefaultFixedFontSizeChangedEventCallback[T]) Register() js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)] {
	return js.RegisterCallback[func(details *OnDefaultFixedFontSizeChangedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDefaultFixedFontSizeChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnDefaultFixedFontSizeChangedArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDefaultFixedFontSizeChanged returns true if the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.addListener" exists.
func HasFuncOnDefaultFixedFontSizeChanged() bool {
	return js.True == bindings.HasFuncOnDefaultFixedFontSizeChanged()
}

// FuncOnDefaultFixedFontSizeChanged returns the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.addListener".
func FuncOnDefaultFixedFontSizeChanged() (fn js.Func[func(callback js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)])]) {
	bindings.FuncOnDefaultFixedFontSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDefaultFixedFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.addListener" directly.
func OnDefaultFixedFontSizeChanged(callback js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)]) (ret js.Void) {
	bindings.CallOnDefaultFixedFontSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDefaultFixedFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDefaultFixedFontSizeChanged(callback js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDefaultFixedFontSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDefaultFixedFontSizeChanged returns true if the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.removeListener" exists.
func HasFuncOffDefaultFixedFontSizeChanged() bool {
	return js.True == bindings.HasFuncOffDefaultFixedFontSizeChanged()
}

// FuncOffDefaultFixedFontSizeChanged returns the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.removeListener".
func FuncOffDefaultFixedFontSizeChanged() (fn js.Func[func(callback js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)])]) {
	bindings.FuncOffDefaultFixedFontSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDefaultFixedFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.removeListener" directly.
func OffDefaultFixedFontSizeChanged(callback js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)]) (ret js.Void) {
	bindings.CallOffDefaultFixedFontSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDefaultFixedFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDefaultFixedFontSizeChanged(callback js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDefaultFixedFontSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDefaultFixedFontSizeChanged returns true if the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.hasListener" exists.
func HasFuncHasOnDefaultFixedFontSizeChanged() bool {
	return js.True == bindings.HasFuncHasOnDefaultFixedFontSizeChanged()
}

// FuncHasOnDefaultFixedFontSizeChanged returns the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.hasListener".
func FuncHasOnDefaultFixedFontSizeChanged() (fn js.Func[func(callback js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)]) bool]) {
	bindings.FuncHasOnDefaultFixedFontSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDefaultFixedFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.hasListener" directly.
func HasOnDefaultFixedFontSizeChanged(callback js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)]) (ret bool) {
	bindings.CallHasOnDefaultFixedFontSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDefaultFixedFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFixedFontSizeChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDefaultFixedFontSizeChanged(callback js.Func[func(details *OnDefaultFixedFontSizeChangedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDefaultFixedFontSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDefaultFontSizeChangedEventCallbackFunc func(this js.Ref, details *OnDefaultFontSizeChangedArgDetails) js.Ref

func (fn OnDefaultFontSizeChangedEventCallbackFunc) Register() js.Func[func(details *OnDefaultFontSizeChangedArgDetails)] {
	return js.RegisterCallback[func(details *OnDefaultFontSizeChangedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDefaultFontSizeChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnDefaultFontSizeChangedArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnDefaultFontSizeChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnDefaultFontSizeChangedArgDetails) js.Ref
	Arg T
}

func (cb *OnDefaultFontSizeChangedEventCallback[T]) Register() js.Func[func(details *OnDefaultFontSizeChangedArgDetails)] {
	return js.RegisterCallback[func(details *OnDefaultFontSizeChangedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDefaultFontSizeChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnDefaultFontSizeChangedArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDefaultFontSizeChanged returns true if the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.addListener" exists.
func HasFuncOnDefaultFontSizeChanged() bool {
	return js.True == bindings.HasFuncOnDefaultFontSizeChanged()
}

// FuncOnDefaultFontSizeChanged returns the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.addListener".
func FuncOnDefaultFontSizeChanged() (fn js.Func[func(callback js.Func[func(details *OnDefaultFontSizeChangedArgDetails)])]) {
	bindings.FuncOnDefaultFontSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// OnDefaultFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.addListener" directly.
func OnDefaultFontSizeChanged(callback js.Func[func(details *OnDefaultFontSizeChangedArgDetails)]) (ret js.Void) {
	bindings.CallOnDefaultFontSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDefaultFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDefaultFontSizeChanged(callback js.Func[func(details *OnDefaultFontSizeChangedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDefaultFontSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDefaultFontSizeChanged returns true if the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.removeListener" exists.
func HasFuncOffDefaultFontSizeChanged() bool {
	return js.True == bindings.HasFuncOffDefaultFontSizeChanged()
}

// FuncOffDefaultFontSizeChanged returns the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.removeListener".
func FuncOffDefaultFontSizeChanged() (fn js.Func[func(callback js.Func[func(details *OnDefaultFontSizeChangedArgDetails)])]) {
	bindings.FuncOffDefaultFontSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// OffDefaultFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.removeListener" directly.
func OffDefaultFontSizeChanged(callback js.Func[func(details *OnDefaultFontSizeChangedArgDetails)]) (ret js.Void) {
	bindings.CallOffDefaultFontSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDefaultFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDefaultFontSizeChanged(callback js.Func[func(details *OnDefaultFontSizeChangedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDefaultFontSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDefaultFontSizeChanged returns true if the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.hasListener" exists.
func HasFuncHasOnDefaultFontSizeChanged() bool {
	return js.True == bindings.HasFuncHasOnDefaultFontSizeChanged()
}

// FuncHasOnDefaultFontSizeChanged returns the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.hasListener".
func FuncHasOnDefaultFontSizeChanged() (fn js.Func[func(callback js.Func[func(details *OnDefaultFontSizeChangedArgDetails)]) bool]) {
	bindings.FuncHasOnDefaultFontSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnDefaultFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.hasListener" directly.
func HasOnDefaultFontSizeChanged(callback js.Func[func(details *OnDefaultFontSizeChangedArgDetails)]) (ret bool) {
	bindings.CallHasOnDefaultFontSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDefaultFontSizeChanged calls the function "WEBEXT.fontSettings.onDefaultFontSizeChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDefaultFontSizeChanged(callback js.Func[func(details *OnDefaultFontSizeChangedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDefaultFontSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnFontChangedEventCallbackFunc func(this js.Ref, details *OnFontChangedArgDetails) js.Ref

func (fn OnFontChangedEventCallbackFunc) Register() js.Func[func(details *OnFontChangedArgDetails)] {
	return js.RegisterCallback[func(details *OnFontChangedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnFontChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnFontChangedArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnFontChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnFontChangedArgDetails) js.Ref
	Arg T
}

func (cb *OnFontChangedEventCallback[T]) Register() js.Func[func(details *OnFontChangedArgDetails)] {
	return js.RegisterCallback[func(details *OnFontChangedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnFontChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnFontChangedArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnFontChanged returns true if the function "WEBEXT.fontSettings.onFontChanged.addListener" exists.
func HasFuncOnFontChanged() bool {
	return js.True == bindings.HasFuncOnFontChanged()
}

// FuncOnFontChanged returns the function "WEBEXT.fontSettings.onFontChanged.addListener".
func FuncOnFontChanged() (fn js.Func[func(callback js.Func[func(details *OnFontChangedArgDetails)])]) {
	bindings.FuncOnFontChanged(
		js.Pointer(&fn),
	)
	return
}

// OnFontChanged calls the function "WEBEXT.fontSettings.onFontChanged.addListener" directly.
func OnFontChanged(callback js.Func[func(details *OnFontChangedArgDetails)]) (ret js.Void) {
	bindings.CallOnFontChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnFontChanged calls the function "WEBEXT.fontSettings.onFontChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnFontChanged(callback js.Func[func(details *OnFontChangedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnFontChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffFontChanged returns true if the function "WEBEXT.fontSettings.onFontChanged.removeListener" exists.
func HasFuncOffFontChanged() bool {
	return js.True == bindings.HasFuncOffFontChanged()
}

// FuncOffFontChanged returns the function "WEBEXT.fontSettings.onFontChanged.removeListener".
func FuncOffFontChanged() (fn js.Func[func(callback js.Func[func(details *OnFontChangedArgDetails)])]) {
	bindings.FuncOffFontChanged(
		js.Pointer(&fn),
	)
	return
}

// OffFontChanged calls the function "WEBEXT.fontSettings.onFontChanged.removeListener" directly.
func OffFontChanged(callback js.Func[func(details *OnFontChangedArgDetails)]) (ret js.Void) {
	bindings.CallOffFontChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffFontChanged calls the function "WEBEXT.fontSettings.onFontChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffFontChanged(callback js.Func[func(details *OnFontChangedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffFontChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnFontChanged returns true if the function "WEBEXT.fontSettings.onFontChanged.hasListener" exists.
func HasFuncHasOnFontChanged() bool {
	return js.True == bindings.HasFuncHasOnFontChanged()
}

// FuncHasOnFontChanged returns the function "WEBEXT.fontSettings.onFontChanged.hasListener".
func FuncHasOnFontChanged() (fn js.Func[func(callback js.Func[func(details *OnFontChangedArgDetails)]) bool]) {
	bindings.FuncHasOnFontChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnFontChanged calls the function "WEBEXT.fontSettings.onFontChanged.hasListener" directly.
func HasOnFontChanged(callback js.Func[func(details *OnFontChangedArgDetails)]) (ret bool) {
	bindings.CallHasOnFontChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnFontChanged calls the function "WEBEXT.fontSettings.onFontChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnFontChanged(callback js.Func[func(details *OnFontChangedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnFontChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnMinimumFontSizeChangedEventCallbackFunc func(this js.Ref, details *OnMinimumFontSizeChangedArgDetails) js.Ref

func (fn OnMinimumFontSizeChangedEventCallbackFunc) Register() js.Func[func(details *OnMinimumFontSizeChangedArgDetails)] {
	return js.RegisterCallback[func(details *OnMinimumFontSizeChangedArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnMinimumFontSizeChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnMinimumFontSizeChangedArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnMinimumFontSizeChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnMinimumFontSizeChangedArgDetails) js.Ref
	Arg T
}

func (cb *OnMinimumFontSizeChangedEventCallback[T]) Register() js.Func[func(details *OnMinimumFontSizeChangedArgDetails)] {
	return js.RegisterCallback[func(details *OnMinimumFontSizeChangedArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnMinimumFontSizeChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnMinimumFontSizeChangedArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnMinimumFontSizeChanged returns true if the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.addListener" exists.
func HasFuncOnMinimumFontSizeChanged() bool {
	return js.True == bindings.HasFuncOnMinimumFontSizeChanged()
}

// FuncOnMinimumFontSizeChanged returns the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.addListener".
func FuncOnMinimumFontSizeChanged() (fn js.Func[func(callback js.Func[func(details *OnMinimumFontSizeChangedArgDetails)])]) {
	bindings.FuncOnMinimumFontSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// OnMinimumFontSizeChanged calls the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.addListener" directly.
func OnMinimumFontSizeChanged(callback js.Func[func(details *OnMinimumFontSizeChangedArgDetails)]) (ret js.Void) {
	bindings.CallOnMinimumFontSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnMinimumFontSizeChanged calls the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnMinimumFontSizeChanged(callback js.Func[func(details *OnMinimumFontSizeChangedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnMinimumFontSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffMinimumFontSizeChanged returns true if the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.removeListener" exists.
func HasFuncOffMinimumFontSizeChanged() bool {
	return js.True == bindings.HasFuncOffMinimumFontSizeChanged()
}

// FuncOffMinimumFontSizeChanged returns the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.removeListener".
func FuncOffMinimumFontSizeChanged() (fn js.Func[func(callback js.Func[func(details *OnMinimumFontSizeChangedArgDetails)])]) {
	bindings.FuncOffMinimumFontSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// OffMinimumFontSizeChanged calls the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.removeListener" directly.
func OffMinimumFontSizeChanged(callback js.Func[func(details *OnMinimumFontSizeChangedArgDetails)]) (ret js.Void) {
	bindings.CallOffMinimumFontSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffMinimumFontSizeChanged calls the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffMinimumFontSizeChanged(callback js.Func[func(details *OnMinimumFontSizeChangedArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffMinimumFontSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnMinimumFontSizeChanged returns true if the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.hasListener" exists.
func HasFuncHasOnMinimumFontSizeChanged() bool {
	return js.True == bindings.HasFuncHasOnMinimumFontSizeChanged()
}

// FuncHasOnMinimumFontSizeChanged returns the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.hasListener".
func FuncHasOnMinimumFontSizeChanged() (fn js.Func[func(callback js.Func[func(details *OnMinimumFontSizeChangedArgDetails)]) bool]) {
	bindings.FuncHasOnMinimumFontSizeChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnMinimumFontSizeChanged calls the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.hasListener" directly.
func HasOnMinimumFontSizeChanged(callback js.Func[func(details *OnMinimumFontSizeChangedArgDetails)]) (ret bool) {
	bindings.CallHasOnMinimumFontSizeChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnMinimumFontSizeChanged calls the function "WEBEXT.fontSettings.onMinimumFontSizeChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnMinimumFontSizeChanged(callback js.Func[func(details *OnMinimumFontSizeChangedArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnMinimumFontSizeChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetDefaultFixedFontSize returns true if the function "WEBEXT.fontSettings.setDefaultFixedFontSize" exists.
func HasFuncSetDefaultFixedFontSize() bool {
	return js.True == bindings.HasFuncSetDefaultFixedFontSize()
}

// FuncSetDefaultFixedFontSize returns the function "WEBEXT.fontSettings.setDefaultFixedFontSize".
func FuncSetDefaultFixedFontSize() (fn js.Func[func(details SetDefaultFixedFontSizeArgDetails) js.Promise[js.Void]]) {
	bindings.FuncSetDefaultFixedFontSize(
		js.Pointer(&fn),
	)
	return
}

// SetDefaultFixedFontSize calls the function "WEBEXT.fontSettings.setDefaultFixedFontSize" directly.
func SetDefaultFixedFontSize(details SetDefaultFixedFontSizeArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetDefaultFixedFontSize(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetDefaultFixedFontSize calls the function "WEBEXT.fontSettings.setDefaultFixedFontSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDefaultFixedFontSize(details SetDefaultFixedFontSizeArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDefaultFixedFontSize(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSetDefaultFontSize returns true if the function "WEBEXT.fontSettings.setDefaultFontSize" exists.
func HasFuncSetDefaultFontSize() bool {
	return js.True == bindings.HasFuncSetDefaultFontSize()
}

// FuncSetDefaultFontSize returns the function "WEBEXT.fontSettings.setDefaultFontSize".
func FuncSetDefaultFontSize() (fn js.Func[func(details SetDefaultFontSizeArgDetails) js.Promise[js.Void]]) {
	bindings.FuncSetDefaultFontSize(
		js.Pointer(&fn),
	)
	return
}

// SetDefaultFontSize calls the function "WEBEXT.fontSettings.setDefaultFontSize" directly.
func SetDefaultFontSize(details SetDefaultFontSizeArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetDefaultFontSize(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetDefaultFontSize calls the function "WEBEXT.fontSettings.setDefaultFontSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDefaultFontSize(details SetDefaultFontSizeArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDefaultFontSize(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSetFont returns true if the function "WEBEXT.fontSettings.setFont" exists.
func HasFuncSetFont() bool {
	return js.True == bindings.HasFuncSetFont()
}

// FuncSetFont returns the function "WEBEXT.fontSettings.setFont".
func FuncSetFont() (fn js.Func[func(details SetFontArgDetails) js.Promise[js.Void]]) {
	bindings.FuncSetFont(
		js.Pointer(&fn),
	)
	return
}

// SetFont calls the function "WEBEXT.fontSettings.setFont" directly.
func SetFont(details SetFontArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetFont(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetFont calls the function "WEBEXT.fontSettings.setFont"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetFont(details SetFontArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetFont(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSetMinimumFontSize returns true if the function "WEBEXT.fontSettings.setMinimumFontSize" exists.
func HasFuncSetMinimumFontSize() bool {
	return js.True == bindings.HasFuncSetMinimumFontSize()
}

// FuncSetMinimumFontSize returns the function "WEBEXT.fontSettings.setMinimumFontSize".
func FuncSetMinimumFontSize() (fn js.Func[func(details SetMinimumFontSizeArgDetails) js.Promise[js.Void]]) {
	bindings.FuncSetMinimumFontSize(
		js.Pointer(&fn),
	)
	return
}

// SetMinimumFontSize calls the function "WEBEXT.fontSettings.setMinimumFontSize" directly.
func SetMinimumFontSize(details SetMinimumFontSizeArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallSetMinimumFontSize(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySetMinimumFontSize calls the function "WEBEXT.fontSettings.setMinimumFontSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetMinimumFontSize(details SetMinimumFontSizeArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetMinimumFontSize(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}
