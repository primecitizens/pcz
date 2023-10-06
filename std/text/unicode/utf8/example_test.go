// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utf8_test

import (
	"fmt"

	"github.com/primecitizens/pcz/std/text/unicode/common"
	"github.com/primecitizens/pcz/std/text/unicode/utf8"
)

func ExampleLast() {
	str := "Hello, 世界"

	for len(str) > 0 {
		r, size := utf8.Last(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[:len(str)-size]
	}
	// Output:
	// 界 3
	// 世 3
	//   1
	// , 1
	// o 1
	// l 1
	// l 1
	// e 1
	// H 1

}

func ExampleFirst() {
	str := "Hello, 世界"

	for len(str) > 0 {
		r, size := utf8.First(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[size:]
	}
	// Output:
	// H 1
	// e 1
	// l 1
	// l 1
	// o 1
	// , 1
	//   1
	// 世 3
	// 界 3
}

func ExampleEncodeRune() {
	r := '世'
	buf := make([]byte, 0, 3)

	buf, n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)
	// Output:
	// [228 184 150]
	// 3
}

func ExampleEncodeRune_outOfRange() {
	runes := []rune{
		// Less than 0, out of range.
		-1,
		// Greater than 0x10FFFF, out of range.
		0x110000,
		// The Unicode replacement character.
		common.RuneError,
	}
	for i, c := range runes {
		buf := make([]byte, 0, 3)
		buf, size := utf8.EncodeRune(buf, c)
		fmt.Printf("%d: %d %[2]s %d\n", i, buf, size)
	}
	// Output:
	// 0: [239 191 189] � 3
	// 1: [239 191 189] � 3
	// 2: [239 191 189] � 3
}

func ExampleFullRune() {
	str := "世"
	fmt.Println(utf8.FullRune(str))
	fmt.Println(utf8.FullRune(str[:2]))
	// Output:
	// true
	// false
}

func ExampleCount() {
	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.Count(str))
	// Output:
	// bytes = 13
	// runes = 9
}

func ExampleRuneLen() {
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('界'))
	// Output:
	// 1
	// 3
}

func ExampleRuneStart() {
	buf := []byte("a界")
	fmt.Println(utf8.RuneStart(buf[0]))
	fmt.Println(utf8.RuneStart(buf[1]))
	fmt.Println(utf8.RuneStart(buf[2]))
	// Output:
	// true
	// true
	// false
}

func ExampleRuneValid() {
	valid := 'a'
	invalid := rune(0xfffffff)

	fmt.Println(utf8.RuneValid(valid))
	fmt.Println(utf8.RuneValid(invalid))
	// Output:
	// true
	// false
}

func ExampleValid() {
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd})

	fmt.Println(utf8.Valid(valid))
	fmt.Println(utf8.Valid(invalid))
	// Output:
	// true
	// false
}

func ExampleAppendRunes() {
	buf1 := utf8.AppendRunes(nil, 0x10000)
	buf2 := utf8.AppendRunes([]byte("init"), 0x10000)
	fmt.Println(string(buf1))
	fmt.Println(string(buf2))
	// Output:
	// 𐀀
	// init𐀀
}
