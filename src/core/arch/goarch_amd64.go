// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64

package arch

const (
	_ArchFamily          = AMD64
	_DefaultPhysPageSize = 4096
	_PCQuantum           = 1
	_MinFrameSize        = 0
	_StackAlign          = PtrSize
)
