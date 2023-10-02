// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !pcz

package debug

import (
	"runtime"
)

func Abort() {
	runtime.Breakpoint()
}

func Return0() {}

func Breakpoint() {
	runtime.Breakpoint()
}

func GetCallerPC() uintptr {
	pc, _, _, _ := runtime.Caller(1)
	return pc
}

func GetCallerSP() uintptr {
	return 0
}
