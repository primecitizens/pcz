// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64

package abi

const (
	// See abi_generic.go.

	// RAX, RBX, RCX, RDI, RSI, R8, R9, R10, R11.
	IntArgRegs = 9

	// X0 -> X14.
	FloatArgRegs = 15

	// We use SSE2 registers which support 64-bit float operations.
	EffectiveFloatRegSize = 8
)
