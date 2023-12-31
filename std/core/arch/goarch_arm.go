// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm

package arch

const (
	_ArchFamily          = ARM
	_DefaultPhysPageSize = 65536
	_PCQuantum           = 4
	_MinFrameSize        = 4
	_StackAlign          = PtrSize
)
