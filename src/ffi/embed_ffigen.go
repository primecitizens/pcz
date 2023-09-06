// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build ffigen

package ffi

import (
	"embed"
	_ "embed" // for go:embed
)

//go:embed js/*.go
var _js_pkg_files_ embed.FS

func JSPackageFiles() embed.FS {
	return _js_pkg_files_
}
