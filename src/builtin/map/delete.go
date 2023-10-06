// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdmap

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/abi"
)

// TODO: implement

func MapDelete(t *abi.MapType, h *hmap, key unsafe.Pointer)   {}
func MapDeleteFast32(t *abi.MapType, h *hmap, key uint32)     {}
func MapDeleteFast64(t *abi.MapType, h *hmap, key uint64)     {}
func MapDeleteFastString(t *abi.MapType, h *hmap, key string) {}

func MapClear(t *abi.MapType, h *hmap) {}
