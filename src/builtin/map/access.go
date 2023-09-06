// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdmap

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
)

// TODO: implement

func MapAccess1(t *abi.MapType, h *hmap, key unsafe.Pointer) unsafe.Pointer   { return nil }
func MapAccess1Fast32(t *abi.MapType, h *hmap, key uint32) unsafe.Pointer     { return nil }
func MapAccess1Fast64(t *abi.MapType, h *hmap, key uint64) unsafe.Pointer     { return nil }
func MapAccess1FastString(t *abi.MapType, h *hmap, key string) unsafe.Pointer { return nil }
func MapAccess1Fat(t *abi.MapType, h *hmap, key unsafe.Pointer, zero *abi.Type) unsafe.Pointer {
	return nil
}

func MapAccess2(t *abi.MapType, h *hmap, key unsafe.Pointer) (unsafe.Pointer, bool) {
	return nil, false
}
func MapAccess2Fast32(t *abi.MapType, h *hmap, key uint32) (unsafe.Pointer, bool) { return nil, false }
func MapAccess2Fast64(t *abi.MapType, h *hmap, key uint64) (unsafe.Pointer, bool) { return nil, false }
func MapAccess2FastString(t *abi.MapType, h *hmap, key string) (unsafe.Pointer, bool) {
	return nil, false
}
func MapAccess2Fat(t *abi.MapType, h *hmap, key unsafe.Pointer, zero *abi.Type) (unsafe.Pointer, bool) {
	return nil, false
}
