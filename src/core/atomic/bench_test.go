// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	"testing"

	"github.com/primecitizens/pcz/std/core/atomic"
)

var sink any

func BenchmarkAtomicLoad64(b *testing.B) {
	var x uint64
	sink = &x
	for i := 0; i < b.N; i++ {
		_ = atomic.Load64(&x)
	}
}

func BenchmarkAtomicStore64(b *testing.B) {
	var x uint64
	sink = &x
	for i := 0; i < b.N; i++ {
		atomic.Store64(&x, 0)
	}
}

func BenchmarkAtomicLoad32(b *testing.B) {
	var x uint32
	sink = &x
	for i := 0; i < b.N; i++ {
		_ = atomic.Load32(&x)
	}
}

func BenchmarkStore32(b *testing.B) {
	var x uint32
	sink = &x
	for i := 0; i < b.N; i++ {
		atomic.Store32(&x, 0)
	}
}

func BenchmarkAnd8(b *testing.B) {
	var x [512]uint8 // give byte its own cache line
	sink = &x
	for i := 0; i < b.N; i++ {
		atomic.And8(&x[255], uint8(i))
	}
}

func BenchmarkAnd32(b *testing.B) {
	var x [128]uint32 // give x its own cache line
	sink = &x
	for i := 0; i < b.N; i++ {
		atomic.And32(&x[63], uint32(i))
	}
}

func BenchmarkAnd8Parallel(b *testing.B) {
	var x [512]uint8 // give byte its own cache line
	sink = &x
	b.RunParallel(func(pb *testing.PB) {
		i := uint8(0)
		for pb.Next() {
			atomic.And8(&x[255], i)
			i++
		}
	})
}

func BenchmarkAnd32Parallel(b *testing.B) {
	var x [128]uint32 // give x its own cache line
	sink = &x
	b.RunParallel(func(pb *testing.PB) {
		i := uint32(0)
		for pb.Next() {
			atomic.And32(&x[63], i)
			i++
		}
	})
}

func BenchmarkOr8(b *testing.B) {
	var x [512]uint8 // give byte its own cache line
	sink = &x
	for i := 0; i < b.N; i++ {
		atomic.Or8(&x[255], uint8(i))
	}
}

func BenchmarkOr32(b *testing.B) {
	var x [128]uint32 // give x its own cache line
	sink = &x
	for i := 0; i < b.N; i++ {
		atomic.Or32(&x[63], uint32(i))
	}
}

func BenchmarkOr8Parallel(b *testing.B) {
	var x [512]uint8 // give byte its own cache line
	sink = &x
	b.RunParallel(func(pb *testing.PB) {
		i := uint8(0)
		for pb.Next() {
			atomic.Or8(&x[255], i)
			i++
		}
	})
}

func BenchmarkOr32Parallel(b *testing.B) {
	var x [128]uint32 // give x its own cache line
	sink = &x
	b.RunParallel(func(pb *testing.PB) {
		i := uint32(0)
		for pb.Next() {
			atomic.Or32(&x[63], i)
			i++
		}
	})
}

func BenchmarkAdd32(b *testing.B) {
	var x uint32
	ptr := &x
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.Add32(ptr, 1)
		}
	})
}

func BenchmarkAdd64(b *testing.B) {
	var x uint64
	ptr := &x
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.Add64(ptr, 1)
		}
	})
}

func BenchmarkCas(b *testing.B) {
	var x uint32
	x = 1
	ptr := &x
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.Cas32(ptr, 1, 0)
			atomic.Cas32(ptr, 0, 1)
		}
	})
}

func BenchmarkCas64(b *testing.B) {
	var x uint64
	x = 1
	ptr := &x
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.Cas64(ptr, 1, 0)
			atomic.Cas64(ptr, 0, 1)
		}
	})
}
func BenchmarkSwap32(b *testing.B) {
	var x uint32
	x = 1
	ptr := &x
	b.RunParallel(func(pb *testing.PB) {
		var y uint32
		y = 1
		for pb.Next() {
			y = atomic.Swap32(ptr, y)
			y += 1
		}
	})
}

func BenchmarkSwap64(b *testing.B) {
	var x uint64
	x = 1
	ptr := &x
	b.RunParallel(func(pb *testing.PB) {
		var y uint64
		y = 1
		for pb.Next() {
			y = atomic.Swap64(ptr, y)
			y += 1
		}
	})
}
