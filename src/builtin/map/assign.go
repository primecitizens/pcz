// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdmap

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/abi"
)

// TODO: implement

func MapAssign(t *abi.MapType, h *hmap, key unsafe.Pointer) unsafe.Pointer          { return nil }
func MapAssignFast32(t *abi.MapType, h *hmap, key uint32) unsafe.Pointer            { return nil }
func MapAssignFast32Ptr(t *abi.MapType, h *hmap, key unsafe.Pointer) unsafe.Pointer { return nil }
func MapAssignFast64(t *abi.MapType, h *hmap, key uint64) unsafe.Pointer            { return nil }
func MapAssignFast64Ptr(t *abi.MapType, h *hmap, key unsafe.Pointer) unsafe.Pointer { return nil }
func MapAssignFastString(t *abi.MapType, h *hmap, key string) unsafe.Pointer        { return nil }
