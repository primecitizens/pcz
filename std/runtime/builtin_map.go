// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package runtime

import (
	"unsafe"

	stdmap "github.com/primecitizens/pcz/std/builtin/map"
	"github.com/primecitizens/pcz/std/core/abi"
)

//
// map
//

func makemap64(t *abi.MapType, hint int64, mapbuf *stdmap.HashMap) *stdmap.HashMap {
	if int64(int(hint)) != hint {
		hint = 0
	}

	return stdmap.Make(t, int(hint), mapbuf)
}

func makemap(t *abi.MapType, hint int, mapbuf *stdmap.HashMap) *stdmap.HashMap {
	return stdmap.Make(t, hint, mapbuf)
}

func makemap_small() *stdmap.HashMap { return stdmap.MakeSmall() }

func mapaccess1(t *abi.MapType, h *stdmap.HashMap, key unsafe.Pointer) unsafe.Pointer {
	return stdmap.MapAccess1(t, h, key)
}

func mapaccess1_fast32(t *abi.MapType, h *stdmap.HashMap, key uint32) unsafe.Pointer {
	return stdmap.MapAccess1Fast32(t, h, key)
}

func mapaccess1_fast64(t *abi.MapType, h *stdmap.HashMap, key uint64) unsafe.Pointer {
	return stdmap.MapAccess1Fast64(t, h, key)
}

func mapaccess1_faststr(t *abi.MapType, h *stdmap.HashMap, key string) unsafe.Pointer {
	return stdmap.MapAccess1FastString(t, h, key)
}

func mapaccess1_fat(t *abi.MapType, h *stdmap.HashMap, key unsafe.Pointer, zero *abi.Type) unsafe.Pointer {
	return stdmap.MapAccess1Fat(t, h, key, zero)
}

func mapaccess2(t *abi.MapType, h *stdmap.HashMap, key unsafe.Pointer) (unsafe.Pointer, bool) {
	return stdmap.MapAccess2(t, h, key)
}

func mapaccess2_fast32(t *abi.MapType, h *stdmap.HashMap, key uint32) (unsafe.Pointer, bool) {
	return stdmap.MapAccess2Fast32(t, h, key)
}

func mapaccess2_fast64(t *abi.MapType, h *stdmap.HashMap, key uint64) (unsafe.Pointer, bool) {
	return stdmap.MapAccess2Fast64(t, h, key)
}

func mapaccess2_faststr(t *abi.MapType, h *stdmap.HashMap, key string) (unsafe.Pointer, bool) {
	return stdmap.MapAccess2FastString(t, h, key)
}

func mapaccess2_fat(t *abi.MapType, h *stdmap.HashMap, key unsafe.Pointer, zero *abi.Type) (unsafe.Pointer, bool) {
	return stdmap.MapAccess2Fat(t, h, key, zero)
}

func mapassign(t *abi.MapType, h *stdmap.HashMap, key unsafe.Pointer) unsafe.Pointer {
	return stdmap.MapAssign(t, h, key)
}

func mapassign_fast32(t *abi.MapType, h *stdmap.HashMap, key uint32) unsafe.Pointer {
	return stdmap.MapAssignFast32(t, h, key)
}

func mapassign_fast32ptr(t *abi.MapType, h *stdmap.HashMap, key unsafe.Pointer) unsafe.Pointer {
	return stdmap.MapAssignFast32Ptr(t, h, key)
}

func mapassign_fast64(t *abi.MapType, h *stdmap.HashMap, key uint64) unsafe.Pointer {
	return stdmap.MapAssignFast64(t, h, key)
}

func mapassign_fast64ptr(t *abi.MapType, h *stdmap.HashMap, key unsafe.Pointer) unsafe.Pointer {
	return stdmap.MapAssignFast64Ptr(t, h, key)
}

func mapassign_faststr(t *abi.MapType, h *stdmap.HashMap, key string) unsafe.Pointer {
	return stdmap.MapAssignFastString(t, h, key)
}

func mapdelete(t *abi.MapType, h *stdmap.HashMap, key unsafe.Pointer) {
	stdmap.MapDelete(t, h, key)
}

func mapdelete_fast32(t *abi.MapType, h *stdmap.HashMap, key uint32) {
	stdmap.MapDeleteFast32(t, h, key)
}

func mapdelete_fast64(t *abi.MapType, h *stdmap.HashMap, key uint64) {
	stdmap.MapDeleteFast64(t, h, key)
}

func mapdelete_faststr(t *abi.MapType, h *stdmap.HashMap, key string) {
	stdmap.MapDeleteFastString(t, h, key)
}

func mapiterinit(t *abi.MapType, h *stdmap.HashMap, it *hiter) {
	stdmap.MapIterInit(t, h, it)
}

func mapiternext(it *hiter) {
	stdmap.MapIterNext(it)
}

func mapclear(t *abi.MapType, h *stdmap.HashMap) {
	stdmap.MapClear(t, h)
}
