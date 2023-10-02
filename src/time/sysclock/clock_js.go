// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build js

package sysclock

import (
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/time"
	"github.com/primecitizens/std/time/sysclock/bindings"
)

var (
	timeOriginNS int64
)

func init() {
	bindings.TimeOriginMS(js.Pointer(&timeOriginNS))
	timeOriginNS *= time.Millisecond
}

func Walltime() (sec int64, nsec int32) {
	var ms int64
	bindings.Walltime(js.Pointer(&ms))
	return ms / 1000, int32((ms % 1000) * 1000_000)
}

func Nanotime() int64 {
	var ms int64
	bindings.Millitime(js.Pointer(&ms))
	return timeOriginNS + ms*time.Millisecond
}
