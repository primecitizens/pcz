// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !pcz

package runtime

import "runtime"

var (
	buildVersion        = runtime.Version()
	defaultGOROOT       = runtime.GOROOT()
	goarm         uint8 = 0
)
