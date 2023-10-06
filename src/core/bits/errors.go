// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bits

import (
	"github.com/primecitizens/pcz/std/core/cerr"
)

type error = cerr.E

type ErrOverflow struct{}

func (ErrOverflow) WriteErr(w cerr.Writer) int { return w.Write("overflow") }

var _ error = ErrDivideByZero{}

type ErrDivideByZero struct{}

func (ErrDivideByZero) WriteErr(w cerr.Writer) int {
	return cerr.WriteJoinS(w, " ", "divide", "by", "zero")
}
