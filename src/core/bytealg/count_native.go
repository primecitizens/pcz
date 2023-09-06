// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm || arm64 || ppc64le || ppc64 || riscv64 || s390x

package bytealg

//go:noescape
func Count(b []byte, c byte) int

//go:noescape
func CountString(s string, c byte) int
