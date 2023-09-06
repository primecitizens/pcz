// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"github.com/primecitizens/std/core/num"
	"github.com/primecitizens/std/ffi/js/bindings"
)

type Number = float64

func NewNumber[T num.Real](x T) Ref {
	maxCache := smallIntCacheSize
	if x < T(0) || // negative number
		x >= T(maxCache) || // uncached
		x-T(uint64(x)) != 0 { // floating point
		return Ref(bindings.NewNumber(float64(x)))
	}

	return Ref(x) + firstSmallIntCache
}
