// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toolchain

import (
	"go/build"
	"os"
	"path/filepath"
)

var (
	GO386    = EnvOr("GO386", defaultGO386)
	GOAMD64  = EnvOr("GOAMD64", defaultGOAMD64)
	GOARM    = EnvOr("GOARM", "7")
	GOMIPS   = EnvOr("GOMIPS", defaultGOMIPS)
	GOMIPS64 = EnvOr("GOMIPS64", defaultGOMIPS64)
	GOPPC64  = EnvOr("GOPPC64", defaultGOPPC64)
	GOWASM   = EnvOr("GOWASM", "")
	GO_LDSO  = defaultGO_LDSO

	GOTOOLDIR = EnvOr("GOTOOLDIR", build.ToolDir)

	Context = build.Default

	GoDefaultIncludeDir = filepath.Join(Context.GOROOT, "pkg", "include")
)

func EnvOr(key, def string) string {
	if x := os.Getenv(key); len(x) != 0 {
		return x
	}
	return def
}

func Getgoextlinkenabled() string {
	return EnvOr("GO_EXTLINK_ENABLED", defaultGO_EXTLINK_ENABLED)
}
