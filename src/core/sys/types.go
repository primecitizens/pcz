// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package sys

import (
	"github.com/primecitizens/std/core/bytealg"
)

type (
	// ArgIter is the iterator implementation for system args.
	ArgIter int

	// EnvIter is the iterator implementation for system environ.
	EnvIter int

	// AuxvIter is the iterator implementation for auxvs.
	AuxvIter int
)

func Arg(i int) (string, bool) {
	return ArgIter(0).Nth(i)
}

func LookupEnv(name string) (value string, found bool) {
	return EnvIter(0).Lookup(name)
}

func GetEnv(name string) string {
	return EnvIter(0).Get(name)
}

func (it EnvIter) Lookup(name string) (value string, found bool) {
	var end int
	for it, end = 0, 0; ; it++ {
		if value, found = it.Nth(0); !found {
			break
		}

		end = bytealg.IndexByte(value, '=')
		if end < 0 {
			if value == name {
				return "", true
			}
		} else {
			if value[:end] == name {
				return value[end+1:], true
			}
		}
	}

	return
}

func (it EnvIter) Get(name string) (value string) {
	value, _ = it.Lookup(name)
	return
}
