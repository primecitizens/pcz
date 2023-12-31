// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64 && openbsd

package cpu

const (
	// From OpenBSD's sys/sysctl.h.
	_CTL_MACHDEP = 7

	// From OpenBSD's machine/cpu.h.
	_CPU_ID_AA64ISAR0 = 2
	_CPU_ID_AA64ISAR1 = 3
)

//go:noescape
func sysctlUint64(mib []uint32) (uint64, bool)

func osInit() *ARM64Features {
	// Get ID_AA64ISAR0 from sysctl.
	isar0, ok := sysctlUint64([]uint32{_CTL_MACHDEP, _CPU_ID_AA64ISAR0})
	if !ok {
		return &ARM64
	}

	parseARM64SystemRegisters(isar0)
	return &ARM64
}
