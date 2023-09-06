// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips64 || mips64le

package cpu

const CacheLinePadSize = 32

// HWCAP bits. These are exposed by the Linux kernel 5.4.
const (
	// CPU features
	hwcap_MIPS_MSA = 1 << 1
)

var (
	// This is initialized by archauxv and should not be changed after it is
	// initialized.
	HWCap uint

	MIPS64X struct {
		_      CacheLinePad
		HasMSA bool // MIPS SIMD architecture
		_      CacheLinePad
	}

	options = []option{
		{Name: "msa", Feature: &MIPS64X.HasMSA},
	}
)

func doinit() {
	// HWCAP feature bits
	MIPS64X.HasMSA = isSet(HWCap, hwcap_MIPS_MSA)
}

func isSet(hwc uint, value uint) bool {
	return hwc&value != 0
}
