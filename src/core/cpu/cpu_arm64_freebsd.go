// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64

package cpu

func osInit() *ARM64Features {
	// Retrieve info from system register ID_AA64ISAR0_EL1.
	parseARM64SystemRegisters(getisar0())
	return &ARM64
}
