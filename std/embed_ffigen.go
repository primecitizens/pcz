// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// NOTE: this file exits only to work around go:embed limitation
// of cross module embedding for the ffigen package.

package std

import (
	"embed"
	_ "embed" // for go:embed
)

//go:embed algo builtin core ffi rt0 runtime text
var coreFiles embed.FS

func FS() embed.FS {
	return coreFiles
}
