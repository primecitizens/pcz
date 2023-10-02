// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

func handleErr2[T any](arg T, err error) (T, error) {
	if err != nil {
		panic(err)
	}

	return arg, err
}

func handleErr(err error) error {
	// TODO: add option to not panic
	if err != nil {
		panic(err)
	}

	return nil
}
