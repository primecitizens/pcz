// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm

package arch

const (
	_ArchFamily          = WASM
	_DefaultPhysPageSize = 65536
	_PCQuantum           = 1
	_MinFrameSize        = 0
	_StackAlign          = PtrSize
)
