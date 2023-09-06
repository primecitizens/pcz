// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm && wasip1

package sysalloc

func resetMemoryDataView() {
	// This function is a no-op on WASI, it is only used to notify the browser
	// that its view of the WASM memory needs to be updated when compiling for
	// GOOS=js.
}
