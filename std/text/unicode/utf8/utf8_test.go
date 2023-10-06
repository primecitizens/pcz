// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utf8_test

import (
	"bytes"
	"strings"
	"testing"

	stdstring "github.com/primecitizens/pcz/std/builtin/string"
	. "github.com/primecitizens/pcz/std/text/unicode/common"
	. "github.com/primecitizens/pcz/std/text/unicode/utf8"
)

type Utf8Map struct {
	r   rune
	str string
}

var utf8map = []Utf8Map{
	{0x0000, "\x00"},
	{0x0001, "\x01"},
	{0x007e, "\x7e"},
	{0x007f, "\x7f"},
	{0x0080, "\xc2\x80"},
	{0x0081, "\xc2\x81"},
	{0x00bf, "\xc2\xbf"},
	{0x00c0, "\xc3\x80"},
	{0x00c1, "\xc3\x81"},
	{0x00c8, "\xc3\x88"},
	{0x00d0, "\xc3\x90"},
	{0x00e0, "\xc3\xa0"},
	{0x00f0, "\xc3\xb0"},
	{0x00f8, "\xc3\xb8"},
	{0x00ff, "\xc3\xbf"},
	{0x0100, "\xc4\x80"},
	{0x07ff, "\xdf\xbf"},
	{0x0400, "\xd0\x80"},
	{0x0800, "\xe0\xa0\x80"},
	{0x0801, "\xe0\xa0\x81"},
	{0x1000, "\xe1\x80\x80"},
	{0xd000, "\xed\x80\x80"},
	{0xd7ff, "\xed\x9f\xbf"}, // last code point before surrogate half.
	{0xe000, "\xee\x80\x80"}, // first code point after surrogate half.
	{0xfffe, "\xef\xbf\xbe"},
	{0xffff, "\xef\xbf\xbf"},
	{0x10000, "\xf0\x90\x80\x80"},
	{0x10001, "\xf0\x90\x80\x81"},
	{0x40000, "\xf1\x80\x80\x80"},
	{0x10fffe, "\xf4\x8f\xbf\xbe"},
	{0x10ffff, "\xf4\x8f\xbf\xbf"},
	{0xFFFD, "\xef\xbf\xbd"},
}

var surrogateMap = []Utf8Map{
	{0xd800, "\xed\xa0\x80"}, // surrogate min decodes to (RuneError, 1)
	{0xdfff, "\xed\xbf\xbf"}, // surrogate max decodes to (RuneError, 1)
}

var testStrings = []string{
	"",
	"abcd",
	"☺☻☹",
	"日a本b語ç日ð本Ê語þ日¥本¼語i日©",
	"日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ日¥本¼語i日©",
	"\x80\x80\x80\x80",
}

func TestFullRune(t *testing.T) {
	for _, m := range utf8map {
		b := []byte(m.str)
		s := m.str
		if !FullRune(s) {
			t.Errorf("FullRuneInString(%q) (%U) = false, want true", s, m.r)
		}
		b1 := b[0 : len(b)-1]
		s1 := string(b1)
		if FullRune(s1) {
			t.Errorf("FullRune(%q) = true, want false", s1)
		}
	}
	for _, s := range []string{"\xc0", "\xc1"} {
		if !FullRune(s) {
			t.Errorf("FullRuneInString(%q) = false, want true", s)
		}
	}
}

func TestEncodeRune(t *testing.T) {
	for _, m := range utf8map {
		b := []byte(m.str)
		var buf [10]byte
		b1, _ := EncodeRune(buf[0:0], m.r)
		if !bytes.Equal(b, b1) {
			t.Errorf("EncodeRune(%#04x) = %q want %q", m.r, b1, b)
		}
	}
}

func TestAppendRune(t *testing.T) {
	for _, m := range utf8map {
		if buf := AppendRunes(nil, m.r); string(buf) != m.str {
			t.Errorf("AppendRune(nil, %#04x) = %s, want %s", m.r, buf, m.str)
		}
		if buf := AppendRunes([]byte("init"), m.r); string(buf) != "init"+m.str {
			t.Errorf("AppendRune(init, %#04x) = %s, want %s", m.r, buf, "init"+m.str)
		}
	}
}

func TestDecodeRune(t *testing.T) {
	for _, m := range utf8map {
		b := []byte(m.str)
		s := m.str
		r, size := First(s)
		if r != m.r || size != len(b) {
			t.Errorf("DecodeRune(%q) = %#04x, %d want %#04x, %d", s, r, size, m.r, len(b))
		}

		s = m.str + "\x00"
		r, size = First(s)
		if r != m.r || size != len(b) {
			t.Errorf("DecodeRune(%q) = %#04x, %d want %#04x, %d", s, r, size, m.r, len(b))
		}

		// make sure missing bytes fail
		wantsize := 1
		if wantsize >= len(b) {
			wantsize = 0
		}
		s = m.str[0 : len(m.str)-1]
		r, size = First(s)
		if r != RuneError || size != wantsize {
			t.Errorf("DecodeRune(%q) = %#04x, %d want %#04x, %d", s, r, size, RuneError, wantsize)
		}

		// make sure bad sequences fail
		if len(b) == 1 {
			b[0] = 0x80
		} else {
			b[len(b)-1] = 0x7F
		}
		s = string(b)
		r, size = First(s)
		if r != RuneError || size != 1 {
			t.Errorf("DecodeRune(%q) = %#04x, %d want %#04x, %d", s, r, size, RuneError, 1)
		}

	}
}

func TestDecodeSurrogateRune(t *testing.T) {
	for _, m := range surrogateMap {
		b := []byte(m.str)
		s := m.str
		r, size := First(s)
		if r != RuneError || size != 1 {
			t.Errorf("DecodeRune(%q) = %x, %d want %x, %d", b, r, size, RuneError, 1)
		}
	}
}

// Check that DecodeRune and DecodeLastRune correspond to
// the equivalent range loop.
func TestSequencing(t *testing.T) {
	for _, ts := range testStrings {
		for _, m := range utf8map {
			for _, s := range []string{ts + m.str, m.str + ts, ts + m.str + ts} {
				testSequence(t, s)
			}
		}
	}
}

func runtimeRuneCount(s string) int {
	return len([]rune(s)) // Replaced by gc with call to runtime.countrunes(s).
}

// Check that a range loop, len([]rune(string)) optimization and
// []rune conversions visit the same runes.
// Not really a test of this package, but the assumption is used here and
// it's good to verify.
func TestRuntimeConversion(t *testing.T) {
	for _, ts := range testStrings {
		count := Count(ts)
		if n := runtimeRuneCount(ts); n != count {
			t.Errorf("%q: len([]rune()) counted %d runes; got %d from RuneCountInString", ts, n, count)
			break
		}

		runes := []rune(ts)
		if n := len(runes); n != count {
			t.Errorf("%q: []rune() has length %d; got %d from RuneCountInString", ts, n, count)
			break
		}
		i := 0
		for _, r := range ts {
			if r != runes[i] {
				t.Errorf("%q[%d]: expected %c (%U); got %c (%U)", ts, i, runes[i], runes[i], r, r)
			}
			i++
		}
	}
}

var invalidSequenceTests = []string{
	"\xed\xa0\x80\x80", // surrogate min
	"\xed\xbf\xbf\x80", // surrogate max

	// xx
	"\x91\x80\x80\x80",

	// s1
	"\xC2\x7F\x80\x80",
	"\xC2\xC0\x80\x80",
	"\xDF\x7F\x80\x80",
	"\xDF\xC0\x80\x80",

	// s2
	"\xE0\x9F\xBF\x80",
	"\xE0\xA0\x7F\x80",
	"\xE0\xBF\xC0\x80",
	"\xE0\xC0\x80\x80",

	// s3
	"\xE1\x7F\xBF\x80",
	"\xE1\x80\x7F\x80",
	"\xE1\xBF\xC0\x80",
	"\xE1\xC0\x80\x80",

	//s4
	"\xED\x7F\xBF\x80",
	"\xED\x80\x7F\x80",
	"\xED\x9F\xC0\x80",
	"\xED\xA0\x80\x80",

	// s5
	"\xF0\x8F\xBF\xBF",
	"\xF0\x90\x7F\xBF",
	"\xF0\x90\x80\x7F",
	"\xF0\xBF\xBF\xC0",
	"\xF0\xBF\xC0\x80",
	"\xF0\xC0\x80\x80",

	// s6
	"\xF1\x7F\xBF\xBF",
	"\xF1\x80\x7F\xBF",
	"\xF1\x80\x80\x7F",
	"\xF1\xBF\xBF\xC0",
	"\xF1\xBF\xC0\x80",
	"\xF1\xC0\x80\x80",

	// s7
	"\xF4\x7F\xBF\xBF",
	"\xF4\x80\x7F\xBF",
	"\xF4\x80\x80\x7F",
	"\xF4\x8F\xBF\xC0",
	"\xF4\x8F\xC0\x80",
	"\xF4\x90\x80\x80",
}

func runtimeDecodeRune(s string) rune {
	for _, r := range s {
		return r
	}
	return -1
}

func TestDecodeInvalidSequence(t *testing.T) {
	for _, s := range invalidSequenceTests {
		r2, _ := First(s)
		if want := RuneError; r2 != want {
			t.Errorf("DecodeRune(%q) = %#04x, want %#04x", s, r2, want)
			return
		}

		r3 := runtimeDecodeRune(s)
		if r2 != r3 {
			t.Errorf("DecodeRune(%q) = %#04x mismatch with runtime.decoderune(%q) = %#04x", s, r2, s, r3)
			return
		}
	}
}

func testSequence(t *testing.T, s string) {
	type info struct {
		index int
		r     rune
	}
	index := make([]info, len(s))
	si := 0
	j := 0
	for i, r := range s {
		if si != i {
			t.Errorf("Sequence(%q) mismatched index %d, want %d", s, si, i)
			return
		}
		index[j] = info{i, r}
		j++
		r2, size2 := First(s[i:])
		if r != r2 {
			t.Errorf("DecodeRune(%q) = %#04x, want %#04x", s[i:], r2, r)
			return
		}
		si += size2
	}
	j--
	for si = len(s); si > 0; {
		r2, size2 := Last(s[0:si])
		if r2 != index[j].r {
			t.Errorf("DecodeLastRune(%q, %d) = %#04x, want %#04x", s, si, r2, index[j].r)
			return
		}
		si -= size2
		if si != index[j].index {
			t.Errorf("DecodeLastRune(%q) index mismatch at %d, want %d", s, si, index[j].index)
			return
		}
		j--
	}
	if si != 0 {
		t.Errorf("DecodeLastRune(%q) finished at %d, not 0", s, si)
	}
}

// Check that negative runes encode as U+FFFD.
func TestNegativeRune(t *testing.T) {
	errorbuf := make([]byte, 0, MaxRuneLen)
	errorbuf, _ = EncodeRune(errorbuf, RuneError)
	buf := make([]byte, 0, MaxRuneLen)
	buf, _ = EncodeRune(buf, -1)
	if !bytes.Equal(buf, errorbuf) {
		t.Errorf("incorrect encoding [% x] for -1; expected [% x]", buf, errorbuf)
	}
}

type RuneCountTest struct {
	in  string
	out int
}

var runecounttests = []RuneCountTest{
	{"abcd", 4},
	{"☺☻☹", 3},
	{"1,2,3,4", 7},
	{"\xe2\x00", 2},
	{"\xe2\x80", 2},
	{"a\xe2\x80", 3},
}

func TestRuneCount(t *testing.T) {
	for _, tt := range runecounttests {
		if out := Count(tt.in); out != tt.out {
			t.Errorf("RuneCountInString(%q) = %d, want %d", tt.in, out, tt.out)
		}
	}
}

type RuneLenTest struct {
	r    rune
	size int
}

var runelentests = []RuneLenTest{
	{0, 1},
	{'e', 1},
	{'é', 2},
	{'☺', 3},
	{RuneError, 3},
	{MaxRune, 4},
	{0xD800, -1},
	{0xDFFF, -1},
	{MaxRune + 1, -1},
	{-1, -1},
}

func TestRuneLen(t *testing.T) {
	for _, tt := range runelentests {
		if size := RuneLen(tt.r); size != tt.size {
			t.Errorf("RuneLen(%#U) = %d, want %d", tt.r, size, tt.size)
		}
	}
}

type ValidTest struct {
	in  string
	out bool
}

var validTests = []ValidTest{
	{"", true},
	{"a", true},
	{"abc", true},
	{"Ж", true},
	{"ЖЖ", true},
	{"брэд-ЛГТМ", true},
	{"☺☻☹", true},
	{"aa\xe2", false},
	{string([]byte{66, 250}), false},
	{string([]byte{66, 250, 67}), false},
	{"a\uFFFDb", true},
	{string("\xF4\x8F\xBF\xBF"), true},      // U+10FFFF
	{string("\xF4\x90\x80\x80"), false},     // U+10FFFF+1; out of range
	{string("\xF7\xBF\xBF\xBF"), false},     // 0x1FFFFF; out of range
	{string("\xFB\xBF\xBF\xBF\xBF"), false}, // 0x3FFFFFF; out of range
	{string("\xc0\x80"), false},             // U+0000 encoded in two bytes: incorrect
	{string("\xed\xa0\x80"), false},         // U+D800 high surrogate (sic)
	{string("\xed\xbf\xbf"), false},         // U+DFFF low surrogate (sic)
}

func TestValid(t *testing.T) {
	for _, tt := range validTests {
		if Valid(tt.in) != tt.out {
			t.Errorf("ValidString(%q) = %v; want %v", tt.in, !tt.out, tt.out)
		}
	}
}

type ValidRuneTest struct {
	r  rune
	ok bool
}

var validrunetests = []ValidRuneTest{
	{0, true},
	{'e', true},
	{'é', true},
	{'☺', true},
	{RuneError, true},
	{MaxRune, true},
	{0xD7FF, true},
	{0xD800, false},
	{0xDFFF, false},
	{0xE000, true},
	{MaxRune + 1, false},
	{-1, false},
}

func TestValidRune(t *testing.T) {
	for _, tt := range validrunetests {
		if ok := RuneValid(tt.r); ok != tt.ok {
			t.Errorf("ValidRune(%#U) = %t, want %t", tt.r, ok, tt.ok)
		}
	}
}

func BenchmarkRuneCountTenASCIIChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("0123456789")
	}
}

func BenchmarkRuneCountTenJapaneseChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("日本語日本語日本語日")
	}
}

var ascii100000 = strings.Repeat("0123456789", 10000)

func BenchmarkValidStringTenASCIIChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid("0123456789")
	}
}

func BenchmarkValidString100KASCIIChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid(ascii100000)
	}
}

func BenchmarkValidStringTenJapaneseChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid("日本語日本語日本語日")
	}
}

func BenchmarkValidStringLongMostlyASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid(longStringMostlyASCII)
	}
}

func BenchmarkValidStringLongJapanese(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid(longStringJapanese)
	}
}

var longStringMostlyASCII string // ~100KB, ~97% ASCII
var longStringJapanese string    // ~100KB, non-ASCII

func init() {
	const japanese = "日本語日本語日本語日"
	var b strings.Builder
	for i := 0; b.Len() < 100_000; i++ {
		if i%100 == 0 {
			b.WriteString(japanese)
		} else {
			b.WriteString("0123456789")
		}
	}
	longStringMostlyASCII = b.String()
	longStringJapanese = strings.Repeat(japanese, 100_000/len(japanese))
}

func BenchmarkEncodeASCIIRune(b *testing.B) {
	buf := make([]byte, 0, MaxRuneLen)
	for i := 0; i < b.N; i++ {
		EncodeRune(buf, 'a')
	}
}

func BenchmarkEncodeJapaneseRune(b *testing.B) {
	buf := make([]byte, 0, MaxRuneLen)
	for i := 0; i < b.N; i++ {
		EncodeRune(buf, '本')
	}
}

func BenchmarkAppendASCIIRune(b *testing.B) {
	buf := make([]byte, MaxRuneLen)
	for i := 0; i < b.N; i++ {
		AppendRunes(buf[:0], 'a')
	}
}

func BenchmarkAppendJapaneseRune(b *testing.B) {
	buf := make([]byte, MaxRuneLen)
	for i := 0; i < b.N; i++ {
		AppendRunes(buf[:0], '本')
	}
}

func BenchmarkDecodeASCIIRune(b *testing.B) {
	a := "a"
	for i := 0; i < b.N; i++ {
		First(a)
	}
}

func BenchmarkDecodeJapaneseRune(b *testing.B) {
	nihon := "本"
	for i := 0; i < b.N; i++ {
		First(nihon)
	}
}

// boolSink is used to reference the return value of benchmarked
// functions to avoid dead code elimination.
var boolSink bool

func BenchmarkFullRune(b *testing.B) {
	benchmarks := []struct {
		name string
		data []byte
	}{
		{"ASCII", []byte("a")},
		{"Incomplete", []byte("\xf0\x90\x80")},
		{"Japanese", []byte("本")},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				boolSink = FullRune(stdstring.FromBytes(bm.data))
			}
		})
	}
}