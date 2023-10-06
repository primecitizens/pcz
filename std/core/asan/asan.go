// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package asan

import "unsafe"

func Read(addr unsafe.Pointer, size uintptr)  {}
func Write(addr unsafe.Pointer, size uintptr) {}
