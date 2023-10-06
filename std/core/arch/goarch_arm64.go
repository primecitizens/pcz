// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64

package arch

const (
	_ArchFamily          = ARM64
	_DefaultPhysPageSize = 65536
	_PCQuantum           = 4
	_MinFrameSize        = 8
	_StackAlign          = 16
)
