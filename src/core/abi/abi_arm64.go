// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64

package abi

const (
	// See abi_generic.go.

	// R0 - R15.
	IntArgRegs = 16

	// F0 - F15.
	FloatArgRegs = 16

	EffectiveFloatRegSize = 8
)
