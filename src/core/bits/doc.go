// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bits

// NOTE: this package MUST be in the exact package math/bits to get functions
// repalced by compiler intrinsics.
//
// See $GOROOT/src/cmd/compile/internal/ssagen/ssa.go#func:InitTables
//
//pcz:importpath math/bits
