// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package wasi

type GroupID uint32

func Getgid() GroupID {
	return 1
}

func Getegid() GroupID {
	return 1
}

func Getgroups(buf []GroupID, offset Size) (Size, Errno) {
	if len(buf) == 0 || offset > 0 {
		return 0, 0
	}

	buf[0] = 1
	return 1, 0
}
