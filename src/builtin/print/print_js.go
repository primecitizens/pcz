// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build js

package stdprint

import (
	"github.com/primecitizens/std/core/mark"
	"github.com/primecitizens/std/ffi/js/bindings"
)

func gwrite(b []byte) {
	if len(b) == 0 {
		return
	}

	bindings.Print(mark.NoEscapeBytesString(b))
}
