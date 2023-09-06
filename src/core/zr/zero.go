// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// Package zr provides zero values for strings, slices, and pointers to
// zero-size values.
package zr

import "unsafe"

const (
	// MaxZero
	//
	// must match value in cmd/compile/internal/gc/walk.go:zeroValSize
	MaxZero = 1024
)

var (
	//go:linkname zerobase runtime.zerobase
	zerobase uintptr

	zeroVal [MaxZero]byte
)

// UnsafePointer returns an unsafe.Pointer pointing to a non-nil zero value.
func UnsafePointer() unsafe.Pointer { return unsafe.Pointer(&zeroVal[0]) }

func Pointer[T any]() *T { return (*T)(UnsafePointer()) }

func ZeroSized() unsafe.Pointer {
	return unsafe.Pointer(&zerobase)
}
