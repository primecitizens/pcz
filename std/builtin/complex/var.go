// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdcomplex

import (
	_ "unsafe" // for go:linkname

	_ "github.com/primecitizens/pcz/std/builtin/float" // import symbol "inf"
)

//go:linkname inf std/builtin/stdfloat.inf
var inf float64
