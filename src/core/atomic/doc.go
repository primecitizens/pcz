// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package atomic provides atomic operations.
//
// On most platforms, the compiler is aware of the functions defined
// in this package, and they're replaced with platform-specific intrinsics.
// On other platforms, generic implementations are made available.
//
// Unless otherwise noted, operations defined in this package are sequentially
// consistent across threads with respect to the values they manipulate. More
// specifically, operations that happen in a specific order on one thread,
// will always be observed to happen in exactly that order by another thread.
package atomic

// NOTE: this package MUST be in the exact package runtime/internal/atomic to
// get functions repalced by compiler intrinsics.
//
// See $GOROOT/src/cmd/compile/internal/ssagen/ssa.go#func:InitTables
//
//pcz:importpath runtime/internal/atomic
