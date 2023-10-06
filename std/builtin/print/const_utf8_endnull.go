// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build android

package stdprint

const (
	str_true                  = "true\x00"
	str_false                 = "false\x00"
	str_space                 = " \x00"
	str_slash                 = "/\x00"
	str_newline               = "\n\x00"
	str_comma                 = ",\x00"
	str_left_round_bracket    = "(\x00"
	str_right_round_bracket   = ")\x00"
	str_i_right_round_bracket = "i)\x00"
	str_left_square_bracket   = "[\x00"
	str_right_square_bracket  = "]\x00"
	str_nan                   = "NaN\x00"
	str_positive_inf          = "+Inf\x00"
	str_negative_inf          = "-Inf\x00"
	str_minus                 = "-\x00"

	float_digits_ext  = 8
	uint_padding      = 1
	string_endpadding = 1
)
