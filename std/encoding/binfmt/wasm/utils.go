// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package wasm

func readU8[T ~uint8 | ~int8](data []byte) (v T, next []byte, ok bool) {
	if len(data) == 0 {
		return
	}

	return T(data[0]), data[1:], true
}

func readU32(data []byte) (v uint32, next []byte, ok bool) {
	var shift uint32
	for len(data) > 0 {
		b := data[0]
		data = data[1:]

		v |= uint32(b&0x7F) << shift
		if (b & 0x80) == 0 {
			ok = true
			break
		}
		shift += 7
	}

	next = data
	return
}
