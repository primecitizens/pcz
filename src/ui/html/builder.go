// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build js

package html

import (
	"github.com/primecitizens/std/core/mark"
	"github.com/primecitizens/std/ffi/js"
)

// Builder builds HTMLElement.innerHTML
type Builder struct {
	elem js.Ref

	_buffer mark.SelfPointing
	buf     []byte

	buffer [256]byte
}

func (b *Builder) Reset(targetElem js.Ref) *Builder {
	b.elem = targetElem
	b.buf = b.buffer[:0:len(b.buffer)]
	reset(targetElem)
	return b
}

func (b *Builder) Close() {
	free(b.elem)
}

func (b *Builder) Int(v int64, base int) *Builder {
	b.clearLocalBuf()
	integer(b.elem, js.True, js.Pointer(&v), uint32(base))
	return b
}

func (b *Builder) Uint(v uint64, base int) *Builder {
	b.clearLocalBuf()
	integer(b.elem, js.False, js.Pointer(&v), uint32(base))
	return b
}

func (b *Builder) Float(v float64) *Builder {
	b.clearLocalBuf()
	float(b.elem, v)
	return b
}

func (b *Builder) Text(text ...string) *Builder {
	var (
		x       string
		replace string
	)

	for _, s := range text {
		last := 0
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '&':
				replace = "&amp;"
			case '\'':
				replace = "&#39;"
			case '<':
				replace = "&lt;"
			case '>':
				replace = "&gt;"
			case '"':
				replace = "&#34;"
			default:
				continue
			}

			if x = s[last:i]; len(x) != 0 {
				b.write(x)
			}
			last = i + 1
			b.write(replace)
		}

		if x = s[last:]; len(x) != 0 {
			b.write(x)
		}
	}

	return b
}

func (b *Builder) HTML(html ...string) *Builder {
	for _, s := range html {
		b.write(s)
	}

	return b
}

func (b *Builder) TextBlock(htmlBegin, htmlEnd string, text ...string) *Builder {
	b.write(htmlBegin)
	b.Text(text...)
	b.write(htmlEnd)
	return b
}

func (b *Builder) Discard() *Builder {
	b.Reset(b.elem)
	return b
}

// Flush buffered content to innerHTML.
func (b *Builder) Flush(append bool) *Builder {
	b.clearLocalBuf()
	flush(b.elem, js.Bool(append))
	return b
}

func (b *Builder) clearLocalBuf() {
	if buf := b.buf; len(buf) > 0 {
		bufRaw(b.elem, js.SliceData(buf), js.SizeU(len(buf)))
		b.buf = buf[:0]
	}
}

func (b *Builder) write(s string) {
	buf := b.buf
	for len(s) > 0 {
		if write := min(cap(buf)-len(buf), len(s)); write > 0 {
			buf = append(buf, s[:write]...)
			s = s[write:]
		}

		if len(s) > 0 {
			bufRaw(b.elem, js.SliceData(buf), js.SizeU(len(buf)))
			buf = buf[:0]
		}
	}

	b.buf = buf
}
