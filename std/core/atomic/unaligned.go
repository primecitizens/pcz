// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import (
	"github.com/primecitizens/pcz/std/core/assert"
)

func panicUnaligned() {
	assert.Throw("unaligned", "64-bit", "atomic", "operation")
}
