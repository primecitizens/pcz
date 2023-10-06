// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || amd64

package cpu

import "unsafe"

const CacheLinePadSize = 64

// cpuid is implemented in cpu_x86.s.
func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)

// xgetbv with ecx = 0 is implemented in cpu_x86.s.
func xgetbv() (eax, edx uint32)

// getGOAMD64level is implemented in cpu_x86.s. Returns number in [1,4].
func getGOAMD64level() int32

const (
	// edx bits
	cpuid_SSE2 = 1 << 26

	// ecx bits
	cpuid_SSE3      = 1 << 0
	cpuid_PCLMULQDQ = 1 << 1
	cpuid_SSSE3     = 1 << 9
	cpuid_FMA       = 1 << 12
	cpuid_SSE41     = 1 << 19
	cpuid_SSE42     = 1 << 20
	cpuid_POPCNT    = 1 << 23
	cpuid_AES       = 1 << 25
	cpuid_OSXSAVE   = 1 << 27
	cpuid_AVX       = 1 << 28

	// ebx bits
	cpuid_BMI1 = 1 << 3
	cpuid_AVX2 = 1 << 5
	cpuid_BMI2 = 1 << 8
	cpuid_ERMS = 1 << 9
	cpuid_ADX  = 1 << 19

	// edx bits for CPUID 0x80000001
	cpuid_RDTSCP = 1 << 27
)

var (
	maxExtendedFunctionInformation uint32
)

func doinit() (ret *X86Features) {
	ret = &X86
	maxID, _, _, _ := cpuid(0, 0)

	if maxID < 1 {
		return
	}

	maxExtendedFunctionInformation, _, _, _ = cpuid(0x80000000, 0)

	_, _, ecx1, _ := cpuid(1, 0)

	if isSet(ecx1, cpuid_SSE3) {
		X86 |= X86Feature_sse3
	}
	if isSet(ecx1, cpuid_PCLMULQDQ) {
		X86 |= X86Feature_pclmulqdq
	}
	if isSet(ecx1, cpuid_SSSE3) {
		X86 |= X86Feature_ssse3
	}
	if isSet(ecx1, cpuid_SSE41) {
		X86 |= X86Feature_sse41
	}
	if isSet(ecx1, cpuid_SSE42) {
		X86 |= X86Feature_sse42
	}
	if isSet(ecx1, cpuid_POPCNT) {
		X86 |= X86Feature_popcnt
	}
	if isSet(ecx1, cpuid_AES) {
		X86 |= X86Feature_aes
	}

	// OSXSAVE can be false when using older Operating Systems
	// or when explicitly disabled on newer Operating Systems by
	// e.g. setting the xsavedisable boot option on Windows 10.
	HasOSXSAVE := isSet(ecx1, cpuid_OSXSAVE)

	// The FMA instruction set extension only has VEX prefixed instructions.
	// VEX prefixed instructions require OSXSAVE to be enabled.
	// See Intel 64 and IA-32 Architecture Software Developer’s Manual Volume 2
	// Section 2.4 "AVX and SSE Instruction Exception Specification"
	if HasOSXSAVE && isSet(ecx1, cpuid_FMA) {
		X86 |= X86Feature_fma
	}

	osSupportsAVX := false
	// For XGETBV, OSXSAVE bit is required and sufficient.
	if HasOSXSAVE {
		eax, _ := xgetbv()
		// Check if XMM and YMM registers have OS support.
		osSupportsAVX = isSet(eax, 1<<1) && isSet(eax, 1<<2)
	}

	if osSupportsAVX && isSet(ecx1, cpuid_AVX) {
		X86 |= X86Feature_avx
	}

	if maxID < 7 {
		return
	}

	_, ebx7, _, _ := cpuid(7, 0)
	if isSet(ebx7, cpuid_BMI1) {
		X86 |= X86Feature_bmi1
	}

	if osSupportsAVX && isSet(ebx7, cpuid_AVX2) {
		X86 |= X86Feature_avx2
	}

	if isSet(ebx7, cpuid_BMI2) {
		X86 |= X86Feature_bmi2
	}

	if isSet(ebx7, cpuid_ERMS) {
		X86 |= X86Feature_erms
	}

	if isSet(ebx7, cpuid_ADX) {
		X86 |= X86Feature_adx
	}

	var maxExtendedInformation uint32
	maxExtendedInformation, _, _, _ = cpuid(0x80000000, 0)

	if maxExtendedInformation < 0x80000001 {
		return
	}

	_, _, _, edxExt1 := cpuid(0x80000001, 0)
	if isSet(edxExt1, cpuid_RDTSCP) {
		X86 |= X86Feature_rdtscp
	}
	return
}

func isSet(hwc uint32, value uint32) bool {
	return hwc&value != 0
}

var nameBuf [3 * 4 * 4]byte

// Name returns the CPU name given by the vendor.
// If the CPU name can not be determined an
// empty string is returned.
func Name() string {
	if maxExtendedFunctionInformation < 0x80000004 {
		return ""
	}

	data := nameBuf[:]

	var eax, ebx, ecx, edx uint32
	eax, ebx, ecx, edx = cpuid(0x80000002, 0)
	data = appendBytes(data, eax, ebx, ecx, edx)
	eax, ebx, ecx, edx = cpuid(0x80000003, 0)
	data = appendBytes(data, eax, ebx, ecx, edx)
	eax, ebx, ecx, edx = cpuid(0x80000004, 0)
	data = appendBytes(data, eax, ebx, ecx, edx)

	// Trim leading spaces.
	for len(data) > 0 && data[0] == ' ' {
		data = data[1:]
	}

	// Trim tail after and including the first null byte.
	for i, c := range data {
		if c == '\x00' {
			data = data[:i]
			break
		}
	}

	return *(*string)(unsafe.Pointer(&data))
}

func appendBytes(b []byte, args ...uint32) []byte {
	for _, arg := range args {
		b = append(b,
			byte((arg >> 0)),
			byte((arg >> 8)),
			byte((arg >> 16)),
			byte((arg >> 24)))
	}
	return b
}
