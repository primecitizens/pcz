// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package msan

import "unsafe"

func Read(addr unsafe.Pointer, sz uintptr)  {}
func Write(addr unsafe.Pointer, sz uintptr) {}
func Move(dst, src, sz uintptr)             {}
