// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdtype

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
)

func mallocgc(sz uintptr, typ *abi.Type, needzero bool) unsafe.Pointer {
	assert.TODO()
	return nil
}
