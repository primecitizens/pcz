// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"github.com/primecitizens/std/core/iter"
	"github.com/primecitizens/std/core/sys"

	// unless your application contains symbols expected the go linker
	// add a runtime package
	_ "github.com/primecitizens/std/runtime"
)

func main() {
	var args sys.ArgIter
	_, ok := args.Nth(0)
	if ok {
		print("hello")
		iter.Each(args, func(i int, v string) bool {
			print(" ", v)
			return true
		})
		println()
	} else {
		println("hello world")
	}
}
