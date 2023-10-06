// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build js

package sysclock

import (
	"github.com/primecitizens/pcz/std/time"
	"github.com/primecitizens/pcz/std/time/sysclock/bindings"
)

var (
	timeOriginNS int64
)

func init() {
	timeOriginNS = int64(bindings.TimeOriginMS() * time.Millisecond)
}

func Walltime() (sec int64, nsec int32) {
	ms := bindings.Walltime()
	return int64(ms) / 1000, int32((int64(ms) % 1000) * 1000_000)
}

func Nanotime() int64 {
	ms := bindings.Millitime()
	return timeOriginNS + int64(ms*time.Millisecond)
}
