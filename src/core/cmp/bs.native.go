// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && (386 || amd64 || s390x || arm || arm64 || ppc64 || ppc64le || mips || mipsle || wasm || mips64 || mips64le || riscv64)

package cmp

// NOTE: code originally found in ${GOROOT}/src/internal/bytealg/compare*

//go:noescape
func Bytes(a, b []byte) int

//go:noescape
func String(a, b string) int
