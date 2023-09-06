// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdprint

import (
	"unsafe"

	stdtype "github.com/primecitizens/std/builtin/type"
)

func PrintPointer(p unsafe.Pointer) { PrintHex(uint64(uintptr(p))) }
func PrintUintptr(p uintptr)        { PrintHex(uint64(p)) }

func PrintString(s string) {
	if string_endpadding > 0 {
		var (
			buf    [128]byte
			offset int
			n      int
		)

		for offset < len(s) {
			n = copy(buf[:127], s[offset:])
			if n != 127 {
				break
			}

			gwrite(buf[:128])
			offset += n + 1
		}

		buf[n] = 0 // may have previous data left
		gwrite(buf[:n+1])
		return
	}

	gwrite(unsafe.Slice(unsafe.StringData(s), len(s)))
}

func PrintEface(e stdtype.Eface) {
	print(str_left_round_bracket, e.Type, str_comma, e.Data, str_right_round_bracket)
}

func PrintIface(i stdtype.Iface) {
	print(str_left_round_bracket, i.Itab, str_comma, i.Data, str_right_round_bracket)
}

func PrintLock()    { /* TODO */ }
func PrintUnlock()  { /* TODO */ }
func PrintSpace()   { PrintString(str_space) }
func PrintNewline() { PrintString(str_newline) }
func PrintComplex(c complex128) {
	print(str_left_round_bracket, real(c), imag(c), str_i_right_round_bracket)
}

func PrintBool(v bool) {
	if v {
		PrintString(str_true)
	} else {
		PrintString(str_false)
	}
}

func PrintFloat(v float64) {
	switch {
	case v != v:
		PrintString(str_nan)
		return
	case v+v == v && v > 0:
		PrintString(str_positive_inf)
		return
	case v+v == v && v < 0:
		PrintString(str_negative_inf)
		return
	}

	const n = 7 // digits printed
	var buf [n + float_digits_ext]byte
	buf[0] = '+'
	e := 0 // exp
	if v == 0 {
		if 1/v < 0 {
			buf[0] = '-'
		}
	} else {
		if v < 0 {
			v = -v
			buf[0] = '-'
		}

		// normalize
		for v >= 10 {
			e++
			v /= 10
		}
		for v < 1 {
			e--
			v *= 10
		}

		// round
		h := 5.0
		for i := 0; i < n; i++ {
			h /= 10
		}
		v += h
		if v >= 10 {
			e++
			v /= 10
		}
	}

	// format +d.dddd+edd
	for i := 0; i < n; i++ {
		s := int(v)
		buf[i+2] = byte(s + '0')
		v -= float64(s)
		v *= 10
	}
	buf[1] = buf[2]
	buf[2] = '.'

	buf[n+2] = 'e'
	buf[n+3] = '+'
	if e < 0 {
		e = -e
		buf[n+3] = '-'
	}

	buf[n+4] = byte(e/100) + '0'
	buf[n+5] = byte(e/10)%10 + '0'
	buf[n+6] = byte(e%10) + '0'
	gwrite(buf[:])
}

func PrintUint(v uint64) {
	var buf [100]byte
	i := len(buf)
	for i--; i > 0; i-- {
		buf[i] = byte(v%10 + '0')
		if v < 10 {
			break
		}
		v /= 10
	}
	gwrite(buf[i:])
}

func PrintInt(v int64) {
	if v < 0 {
		PrintString(str_minus)
		v = -v
	}
	PrintUint(uint64(v))
}

const minhexdigits = 0

func PrintHex(v uint64) {
	const dig = "0123456789abcdef"
	var buf [100]byte
	i := len(buf) - uint_padding
	for i--; i > 0; i-- {
		buf[i] = dig[v%16]
		if v < 16 && len(buf)-i >= minhexdigits {
			break
		}
		v /= 16
	}
	i--
	buf[i] = 'x'
	i--
	buf[i] = '0'
	gwrite(buf[i:])
}

func PrintSlice(s []byte) {
	print(str_left_square_bracket, len(s), str_slash, cap(s), str_right_square_bracket)
	PrintPointer(unsafe.Pointer(unsafe.SliceData(s)))
}
