// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdcomplex

import (
	stdfloat "github.com/primecitizens/std/builtin/float"
)

type Complex256 struct {
	Real stdfloat.Float128
	Imag stdfloat.Float128
}
