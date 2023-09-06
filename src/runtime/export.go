// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package runtime

import (
	_ "unsafe" // for go:linkname

	stdgo "github.com/primecitizens/std/builtin/go"
)

//go:linkname _getg
//go:nosplit
func _getg() *stdgo.GHead {
	return getg()
}
