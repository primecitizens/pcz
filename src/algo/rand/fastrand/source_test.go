// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package fastrand

import (
	"github.com/primecitizens/std/algo/rand"
)

var _ rand.Core = &Source{}
