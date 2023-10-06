// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz && arm64

package mem

import (
	"github.com/primecitizens/pcz/std/core/cpu"
)

var arm64UseAlignedLoads bool

func init() {
	arm64UseAlignedLoads = cpu.ARM64.HasAll(cpu.ARM64Feature_is_neoversen1 | cpu.ARM64Feature_is_zeus)
}
