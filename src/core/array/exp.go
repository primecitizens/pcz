// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package array

type (
	Exp1To65536[T any] interface {
		Exp1To64[T] | Exp128To4096[T] | Exp8192To65536[T]
	}

	Exp1To64[T any] interface {
		~[exp0]T | ~[exp1]T | ~[exp2]T | ~[exp3]T | ~[exp4]T | ~[exp5]T | ~[exp6]T
	}

	Exp128To4096[T any] interface {
		~[exp7]T | ~[exp8]T | ~[exp9]T | ~[exp10]T | ~[exp11]T | ~[exp12]T
	}

	Exp8192To65536[T any] interface {
		~[exp13]T | ~[exp14]T | ~[exp15]T | ~[exp16]T
	}
)

const (
	exp0 = 1 << iota
	exp1
	exp2
	exp3
	exp4
	exp5
	exp6
	exp7
	exp8
	exp9
	exp10
	exp11
	exp12
	exp13
	exp14
	exp15
	exp16
)
