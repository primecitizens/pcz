// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package mem

import (
	"unsafe"
)

// Equal
//
// See ${GOROOT}/src/runtime/stubs.go#func:memequal
//
//go:noescape
func Equal(p, q unsafe.Pointer, sz uintptr) bool
