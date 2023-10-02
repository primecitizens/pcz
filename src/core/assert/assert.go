// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package assert

func Panic[T any](v ...T) {
	Throw("panic")
}

func Unreachable() {
	Throw("unreachable")
}

func TODO(s ...string) {
	Throw("TODO")
}
