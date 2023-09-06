// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package buildcfg

import "runtime"

const (
	defaultGO386              = `sse2`
	defaultGOAMD64            = `v1`
	defaultGOARM              = `5`
	defaultGOMIPS             = `hardfloat`
	defaultGOMIPS64           = `hardfloat`
	defaultGOPPC64            = `power8`
	defaultGOEXPERIMENT       = ``
	defaultGO_EXTLINK_ENABLED = `1`
	defaultGO_LDSO            = ``
	defaultGOOS               = runtime.GOOS
	defaultGOARCH             = runtime.GOARCH
)
