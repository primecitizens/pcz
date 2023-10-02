// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && wasm

package runtime

import (
	"github.com/primecitizens/std/core/assert"
)

func wasmDiv()
func wasmTruncS()
func wasmTruncU()

func sigpanic() {
	assert.Throw("invalid", "memory", "address", "or", "nil", "pointer", "dereference")
}
