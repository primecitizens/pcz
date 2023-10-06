// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package common

const (
	RuneError = '\uFFFD'     // The "error" Rune or "Unicode replacement character"
	MaxRune   = '\U0010FFFF' // Maximum valid Unicode code point.

	SurrogateMin = 0xD800
	SurrogateMax = 0xDFFF
)
