// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cpu implements processor feature detection
// used by the Go standard library.
package cpu

import (
	"github.com/primecitizens/std/core/os"
)

// DebugOptions is set to true by the runtime if the OS supports reading
// GODEBUG early in runtime startup.
// This should not be changed after it is initialized.
const DebugOptions = 0|
	os.IsAix|
	os.IsDarwin|
	os.IsIos|
	os.IsDragonfly|
	os.IsFreebsd|
	os.IsNetbsd|
	os.IsOpenbsd|
	os.IsIllumos|
	os.IsSolaris|
	os.IsLinux|
	0 != 0

type FeatureChecker[Self any] interface {
	// HasAll checks whether wanted features are all present
	HasAll(want Self) bool
}

type FeatureSet[Self any] interface {
	// Set `feat` state to `enable`
	//
	// known = false when the feat is unknown to this feature set
	// enabled = true
	Set(feat string, enable bool) (enabled, found bool)
}

var (
	ARM    ARMFeatures
	X86    X86Features
	ARM64  ARM64Features
	MIPS64 MIPS64Features
	PPC64  PPC64Features
	S390X  S390XFeatures
	WASM   WASMFeatures
)

type ARMFeatures uint32

const (
	ARMFeature_vfpv4 ARMFeatures = 1 << iota
	ARMFeature_idiva
)

func (x ARMFeatures) HasAll(want ARMFeatures) bool {
	return x&want == want
}

func (x *ARMFeatures) Set(feat string, enable bool) (enabled bool, found bool) {
	var want ARMFeatures
	switch feat {
	case "vfpv4":
		want = ARMFeature_vfpv4
	case "idiva":
		want = ARMFeature_idiva
	default:
		return
	}

	found = true
	if *x&want != 0 /* has this feat */ {
		if enable {
			enabled = true
		} else {
			*x &= ^want
		}
	}

	return
}

type X86Features uint32

const (
	X86Feature_adx = 1 << iota
	X86Feature_aes
	X86Feature_erms
	X86Feature_pclmulqdq
	X86Feature_rdtscp
	X86Feature_popcnt
	X86Feature_sse3
	X86Feature_sse41
	X86Feature_sse42
	X86Feature_ssse3
	X86Feature_avx
	X86Feature_avx2
	X86Feature_bmi1
	X86Feature_bmi2
	X86Feature_fma
)

func (x X86Features) HasAll(want X86Features) bool {
	return x&want == want
}

func (x *X86Features) Set(feat string, enable bool) (enabled, found bool) {
	var want X86Features
	switch feat {
	case "adx":
		want = X86Feature_adx
	case "aes":
		want = X86Feature_aes
	case "erms":
		want = X86Feature_erms
	case "pclmulqdq":
		want = X86Feature_pclmulqdq
	case "rdtscp":
		want = X86Feature_rdtscp
	case "popcnt":
		want = X86Feature_popcnt
	case "sse3":
		want = X86Feature_sse3
	case "sse41":
		want = X86Feature_sse41
	case "sse42":
		want = X86Feature_sse42
	case "ssse3":
		want = X86Feature_ssse3
	case "avx":
		want = X86Feature_avx
	case "avx2":
		want = X86Feature_avx2
	case "bmi1":
		want = X86Feature_bmi1
	case "bmi2":
		want = X86Feature_bmi2
	case "fma":
		want = X86Feature_fma
	default:
		return
	}

	found = true
	if *x&want != 0 /* has this feat */ {
		if enable {
			enabled = true
		} else {
			*x &= ^want
		}
	}

	return
}

type ARM64Features uint32

const (
	ARM64Feature_aes ARM64Features = 1 << iota
	ARM64Feature_pmull
	ARM64Feature_sha1
	ARM64Feature_sha2
	ARM64Feature_sha512
	ARM64Feature_crc32
	ARM64Feature_atomics
	ARM64Feature_cpuid
	ARM64Feature_is_neoversen1
	ARM64Feature_is_zeus
)

func (x ARM64Features) HasAll(want ARM64Features) bool {
	return x&want == want
}

func (x *ARM64Features) Set(feat string, enable bool) (enabled bool, found bool) {
	var want ARM64Features
	switch feat {
	case "aes":
		want = ARM64Feature_aes
	case "pmull":
		want = ARM64Feature_pmull
	case "sha1":
		want = ARM64Feature_sha1
	case "sha2":
		want = ARM64Feature_sha2
	case "sha512":
		want = ARM64Feature_sha512
	case "crc32":
		want = ARM64Feature_crc32
	case "atomics":
		want = ARM64Feature_atomics
	case "cpuid":
		want = ARM64Feature_cpuid
	case "isNeoverseN1":
		want = ARM64Feature_is_neoversen1
	case "isZeus":
		want = ARM64Feature_is_zeus
	default:
		return
	}

	found = true
	if *x&want != 0 /* has this feat */ {
		if enable {
			enabled = true
		} else {
			*x &= ^want
		}
	}

	return
}

type MIPS64Features uint32

const (
	MIPS64Feature_msa MIPS64Features = 1 << iota // MIPS SIMD architecture
)

func (x MIPS64Features) HasAll(want MIPS64Features) bool {
	return x&want == want
}

func (x *MIPS64Features) Set(feat string, enable bool) (enabled bool, found bool) {
	var want MIPS64Features
	switch feat {
	case "msa":
		want = MIPS64Feature_msa
	default:
		return
	}

	found = true
	if *x&want != 0 /* has this feat */ {
		if enable {
			enabled = true
		} else {
			*x &= ^want
		}
	}

	return
}

// For ppc64(le), it is safe to check only for ISA level starting on ISA v3.00,
// since there are no optional categories. There are some exceptions that also
// require kernel support to work (darn, scv), so there are feature bits for
// those as well. The minimum processor requirement is POWER8 (ISA 2.07).
// The struct is padded to avoid false sharing.
type PPC64Features uint32

const (
	PPC64Feature_darn       PPC64Features = 1 << iota // Hardware random number generator (requires kernel enablement)
	PPC64Feature_scv                                  // Syscall vectored (requires kernel enablement)
	PPC64Feature_is_power8                            // ISA v2.07 (POWER8)
	PPC64Feature_is_power9                            // ISA v3.00 (POWER9)
	PPC64Feature_is_power10                           // ISA v3.1  (POWER10)
)

func (x PPC64Features) HasAll(want PPC64Features) bool {
	return x&want == want
}

func (x *PPC64Features) Set(feat string, enable bool) (enabled bool, found bool) {
	var want PPC64Features
	switch feat {
	case "darn":
		want = PPC64Feature_darn
	case "scv":
		want = PPC64Feature_scv
	case "power8":
		want = PPC64Feature_is_power8
	case "power9":
		want = PPC64Feature_is_power9
	case "power10":
		want = PPC64Feature_is_power10
	default:
		return
	}

	found = true
	if *x&want != 0 /* has this feat */ {
		if enable {
			enabled = true
		} else {
			*x &= ^want
		}
	}

	return
}

type S390XFeatures uint32

const (
	S390XFeature_zarch  S390XFeatures = 1 << iota // z architecture mode is active [mandatory]
	S390XFeature_stfle                            // store facility list extended [mandatory]
	S390XFeature_ldisp                            // long (20-bit) displacements [mandatory]
	S390XFeature_eimm                             // 32-bit immediates [mandatory]
	S390XFeature_dfp                              // decimal floating point
	S390XFeature_etf3eh                           // ETF-3 enhanced
	S390XFeature_msa                              // message security assist (CPACF)
	S390XFeature_aes                              // KM-AES{128,192,256} functions
	S390XFeature_aescbc                           // KMC-AES{128,192,256} functions
	S390XFeature_aesctr                           // KMCTR-AES{128,192,256} functions
	S390XFeature_aesgcm                           // KMA-GCM-AES{128,192,256} functions
	S390XFeature_ghash                            // KIMD-GHASH function
	S390XFeature_sha1                             // K{I,L}MD-SHA-1 functions
	S390XFeature_sha256                           // K{I,L}MD-SHA-256 functions
	S390XFeature_sha512                           // K{I,L}MD-SHA-512 functions
	S390XFeature_sha3                             // K{I,L}MD-SHA3-{224,256,384,512} and K{I,L}MD-SHAKE-{128,256} functions
	S390XFeature_vx                               // vector facility. Note: the runtime sets this when it processes auxv records.
	S390XFeature_vxe                              // vector-enhancements facility 1
	S390XFeature_kdsa                             // elliptic curve functions
	S390XFeature_ecdsa                            // NIST curves
	S390XFeature_eddsa                            // Edwards curves
)

func (x S390XFeatures) HasAll(want S390XFeatures) bool {
	return x&want == want
}

func (x *S390XFeatures) Set(feat string, enable bool) (enabled bool, found bool) {
	var want S390XFeatures
	switch feat {
	case "zarch":
		want = S390XFeature_zarch
	case "stfle":
		want = S390XFeature_stfle
	case "ldisp":
		want = S390XFeature_ldisp
	case "eimm":
		want = S390XFeature_eimm
	case "dfp":
		want = S390XFeature_dfp
	case "etf3eh":
		want = S390XFeature_etf3eh
	case "msa":
		want = S390XFeature_msa
	case "aes":
		want = S390XFeature_aes
	case "aescbc":
		want = S390XFeature_aescbc
	case "aesctr":
		want = S390XFeature_aesctr
	case "aesgcm":
		want = S390XFeature_aesgcm
	case "ghash":
		want = S390XFeature_ghash
	case "sha1":
		want = S390XFeature_sha1
	case "sha256":
		want = S390XFeature_sha256
	case "sha512":
		want = S390XFeature_sha512
	case "sha3":
		want = S390XFeature_sha3
	case "vx":
		want = S390XFeature_vx
	case "vxe":
		want = S390XFeature_vxe
	case "kdsa":
		want = S390XFeature_kdsa
	case "ecdsa":
		want = S390XFeature_ecdsa
	case "eddsa":
		want = S390XFeature_eddsa
	default:
		return
	}

	found = true
	if *x&want != 0 /* has this feat */ {
		if enable {
			enabled = true
		} else {
			*x &= ^want
		}
	}

	return
}

type WASMFeatures uint32

func (x WASMFeatures) HasAll(want WASMFeatures) bool {
	return x&want == want
}

func (x *WASMFeatures) Set(feat string, enable bool) (enabled, found bool) {
	return
}

// Initialize examines the processor and sets the relevant variables above.
// This is called by the runtime package early in program initialization,
// before normal init functions are run. env is set by runtime if the OS supports
// cpu feature options in GODEBUG.
func Initialize(env string) {
	processOptions(doinit(), env)
}

// processOptions enables or disables CPU feature values based on the parsed env string.
// The env string is expected to be of the form cpu.feature1=value1,cpu.feature2=value2...
// where feature names is one of the architecture specific list stored in the
// cpu packages options variable and values are either 'on' or 'off'.
// If env contains cpu.all=off then all cpu features referenced through the options
// variable are disabled. Other feature names and values result in warning messages.
func processOptions[T FeatureSet[T]](fset T, godebug string) {
	for len(godebug) != 0 {
		field := ""
		i := indexByte(godebug, ',')
		if i < 0 {
			field, godebug = godebug, ""
		} else {
			field, godebug = godebug[:i], godebug[i+1:]
		}
		if len(field) < 4 || field[:4] != "cpu." {
			continue
		}
		i = indexByte(field, '=')
		if i < 0 {
			print("GODEBUG: no value specified for \"", field, "\"\n")
			continue
		}
		key, value := field[4:i], field[i+1:] // e.g. "SSE2", "on"

		var enable bool
		switch value {
		case "on":
			enable = true
		case "off":
			enable = false
		default:
			print("GODEBUG: value \"", value, "\" not supported for cpu option \"", key, "\"\n")
			continue
		}

		enabled, found := fset.Set(key, enable)
		if !found {
			print("GODEBUG: unknown cpu feature \"", key, "\"\n")
			continue
		}

		if enable && !enabled {
			print("GODEBUG: can not enable \"", key, "\", missing CPU support\n")
		}
	}
}

// indexByte returns the index of the first instance of c in s,
// or -1 if c is not present in s.
func indexByte(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}
