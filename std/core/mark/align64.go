// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mark

import (
	"github.com/primecitizens/pcz/std/core/mark/internal/atomic"
)

// Align64 may be added to structs that must be 64-bit aligned.
type Align64 = atomic.Align64
