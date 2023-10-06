// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package time

type Duration int64

const (
	Nanosecond  = 1
	Microsecond = Nanosecond * 1000
	Millisecond = Microsecond * 1000
	Second      = Millisecond * 1000
	Minute      = Second * 60
	Hour        = Minute * 60
	Day         = Hour * 24
	Week        = Day * 7
)
