// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build js

package stdprint

import (
	"github.com/primecitizens/pcz/std/builtin/print/bindings"
	"github.com/primecitizens/pcz/std/ffi/js"
)

func gwrite(b []byte) {
	if len(b) == 0 {
		return
	}

	bindings.Print(js.SliceData(b), js.SizeU(len(b)))
}
