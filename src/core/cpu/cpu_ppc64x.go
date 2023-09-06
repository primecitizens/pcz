// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64 || ppc64le

package cpu

const CacheLinePadSize = 128

// For ppc64(le), it is safe to check only for ISA level starting on ISA v3.00,
// since there are no optional categories. There are some exceptions that also
// require kernel support to work (darn, scv), so there are feature bits for
// those as well. The minimum processor requirement is POWER8 (ISA 2.07).
// The struct is padded to avoid false sharing.
var (
	options = []option{
		{Name: "darn", Feature: &PPC64.HasDARN},
		{Name: "scv", Feature: &PPC64.HasSCV},
		{Name: "power9", Feature: &PPC64.IsPOWER9},
	}
)

func doinit() {
	osinit()
}

func isSet(hwc uint, value uint) bool {
	return hwc&value != 0
}

func Name() string {
	switch {
	case PPC64.IsPOWER10:
		return "POWER10"
	case PPC64.IsPOWER9:
		return "POWER9"
	case PPC64.IsPOWER8:
		return "POWER8"
	}
	return ""
}
