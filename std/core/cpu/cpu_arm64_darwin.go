// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64 && darwin && !ios

package cpu

import (
	"unsafe"
)

func osInit() *ARM64Features {
	if sysctlEnabled("hw.optional.armv8_1_atomics\x00") {
		ARM64 |= ARM64Feature_atomics
	}

	if sysctlEnabled("hw.optional.armv8_crc32\x00") {
		ARM64 |= ARM64Feature_crc32
	}

	// There are no hw.optional sysctl values for the below features on Mac OS 11.0
	// to detect their supported state dynamically. Assume the CPU features that
	// Apple Silicon M1 supports to be available as a minimal set of features
	// to all Go programs running on darwin/arm64.
	ARM64 |= ARM64Feature_aes | ARM64Feature_pmull | ARM64Feature_sha1 | ARM64Feature_sha2
	return &ARM64
}

type stringHeader struct {
	str *byte
	sz  int
}

func getsysctlbyname(name string) (ret int32, out int32) {
	nout := unsafe.Sizeof(out)
	ret = sysctlByName((*stringHeader)(unsafe.Pointer(&name)).str, (*byte)(unsafe.Pointer(&out)), &nout, nil, 0)
	return
}

// we cannot call libc.SysctlByName directly due to import cycle as this package
// is used in core/atomic to determine cpu features, and core/atomic is used in ffi/cgo
// ffi/cgo is used in platform/libc
//
//go:noescape
//go:linkname sysctlByName platform/libc.SysctlByName
func sysctlByName(name *byte, oldp *byte, oldlenp *uintptr, newp *byte, newlen uintptr) int32

func sysctlEnabled(name string) bool {
	ret, value := getsysctlbyname(name)
	if ret < 0 {
		return false
	}
	return value > 0
}
