// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package matrix

import (
	"testing"
)

func TestMat4_Mul(t *testing.T) {
	x := Mat4[uint8]{
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
	}

	expcted := Mat4[uint8]{
		10, 10, 10, 10,
		26, 26, 26, 26,
		42, 42, 42, 42,
		58, 58, 58, 58,
	}

	if *x.Mul(Mat4[uint8]{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}) != expcted {
		t.Error(x)
	}
}

func BenchmarkMat4_Mul(b *testing.B) {
	b.Run("Mul", func(b *testing.B) {
		var x Mat4[uint8]
		for i := 0; i < b.N; i++ {
			_ = x.Mul(x)
		}
	})

	b.Run("MulV", func(b *testing.B) {
		var x Mat4[uint8]
		for i := 0; i < b.N; i++ {
			_ = x.MulV(x)
		}
	})
}
