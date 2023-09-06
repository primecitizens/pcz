// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build ffigen

package builtin

import "embed"

//go:embed int/*.go
var _stdint_pkg_files embed.FS

func StdintPackageFiles() embed.FS {
	return _stdint_pkg_files
}

//go:embed float/*.go
var _stdfloat_pkg_files embed.FS

func StdfloatPackageFiles() embed.FS {
	return _stdfloat_pkg_files
}

//go:embed complex/*.go
var _stdcomplex_pkg_files embed.FS

func StdcomplexPackageFiles() embed.FS {
	return _stdcomplex_pkg_files
}

//go:embed fixed/*.go
var _stdfixed_pkg_files embed.FS

func StdfixedPackageFiles() embed.FS {
	return _stdfixed_pkg_files
}
