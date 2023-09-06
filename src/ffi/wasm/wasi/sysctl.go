// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package wasi

func Sysctl(key string) (value string, errno Errno) {
	switch key {
	case "kern.hostname":
		return "wasip1", 0
	default:
		errno = ENOSYS
		return
	}
}
