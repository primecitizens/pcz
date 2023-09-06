// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noos && js && wasm

package js

// pause sets SP to newsp and pauses the execution of Go's WebAssembly
// code until an event is triggered.
func pause(newsp uintptr)
