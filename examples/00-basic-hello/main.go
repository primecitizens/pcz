// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/pcz/std/core/iter"
	"github.com/primecitizens/pcz/std/core/sys"

	// unless your application contains symbols expected by the go linker
	// add a runtime package
	_ "github.com/primecitizens/pcz/std/runtime"
)

func main() {
	if _, ok := sys.Arg(0); ok {
		print("Hello")
		iter.Each(sys.ArgIter(0), func(i int, v string) bool {
			print(" ", v)
			return true
		})
		println()
	} else {
		println("Hello 🥳")
	}
}
