// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasip1

package sysclock

import (
	"github.com/primecitizens/std/core/mark"
	"github.com/primecitizens/std/ffi/wasm/wasi"
)

func Walltime() (sec int64, nsec int32) {
	var time wasi.Timestamp
	errno := wasi.ClockTimeGet(wasi.ClockRealtime, 0, mark.NoEscapePointer(&time))
	if errno != 0 {
		return 0, 0
	}

	return int64(time / 1000_000_000), int32(time % 1000_000_000)
}

func Nanotime() int64 {
	var time wasi.Timestamp
	errno := wasi.ClockTimeGet(wasi.ClockMonotonic, 0, mark.NoEscapePointer(&time))
	if errno != 0 {
		return 0
	}

	return int64(time)
}
