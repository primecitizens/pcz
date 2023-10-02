// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package runtime

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
)

func explicit_mallocgc(typ *abi.Type, n uintptr, needszero bool) unsafe.Pointer {
	return getg().G().DefaultAlloc().Malloc(typ, n, needszero)
}
