// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package sys_test

import (
	"github.com/primecitizens/pcz/std/algo/sort"
	"github.com/primecitizens/pcz/std/core/iter"
	"github.com/primecitizens/pcz/std/core/sys"
)

var (
	_ iter.Interface[string, sys.ArgIter]    = sys.ArgIter(0)
	_ iter.Interface[string, sys.EnvIter]    = sys.EnvIter(0)
	_ iter.Interface[sys.Auxv, sys.AuxvIter] = sys.AuxvIter(0)

	_ sort.Interface = sys.EnvIter(0)
)

var (
	ExecutablePath func() string = sys.ExecutablePath
)
