// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package binary

import (
	"github.com/primecitizens/pcz/std/core/arch"
	"github.com/primecitizens/pcz/std/core/bits"
)

type (
	U16le[T ~uint16] uint16
	U32le[T ~uint32] uint32
	U64le[T ~uint64] uint64

	I16le[T ~uint16] uint16
	I32le[T ~uint32] uint32
	I64le[T ~uint64] uint64
)

func (x U16le[T]) Get() T {
	if arch.BigEndian {
		return T(bits.ReverseBytes16(uint16(x)))
	}

	return T(x)
}

func (x *U16le[T]) Set(v T) {
	if arch.BigEndian {
		*x = U16le[T](bits.ReverseBytes16(uint16(v)))
	} else {
		*x = U16le[T](v)
	}
}

func (x I16le[T]) Get() T {
	if arch.BigEndian {
		return T(bits.ReverseBytes16(uint16(x)))
	}

	return T(x)
}

func (x *I16le[T]) Set(v T) {
	if arch.BigEndian {
		*x = I16le[T](bits.ReverseBytes16(uint16(v)))
	} else {
		*x = I16le[T](v)
	}
}

func (x U32le[T]) Get() T {
	if arch.BigEndian {
		return T(bits.ReverseBytes32(uint32(x)))
	}

	return T(x)
}

func (x *U32le[T]) Set(v T) {
	if arch.BigEndian {
		*x = U32le[T](bits.ReverseBytes32(uint32(v)))
	} else {
		*x = U32le[T](v)
	}
}

func (x I32le[T]) Get() T {
	if arch.BigEndian {
		return T(bits.ReverseBytes32(uint32(x)))
	}

	return T(x)
}

func (x *I32le[T]) Set(v T) {
	if arch.BigEndian {
		*x = I32le[T](bits.ReverseBytes32(uint32(v)))
	} else {
		*x = I32le[T](v)
	}
}

func (x U64le[T]) Get() T {
	if arch.BigEndian {
		return T(bits.ReverseBytes64(uint64(x)))
	}

	return T(x)
}

func (x *U64le[T]) Set(v T) {
	if arch.BigEndian {
		*x = U64le[T](bits.ReverseBytes64(uint64(v)))
	} else {
		*x = U64le[T](v)
	}
}

func (x I64le[T]) Get() T {
	if arch.BigEndian {
		return T(bits.ReverseBytes64(uint64(x)))
	}

	return T(x)
}

func (x *I64le[T]) Set(v T) {
	if arch.BigEndian {
		*x = I64le[T](bits.ReverseBytes64(uint64(v)))
	} else {
		*x = I64le[T](v)
	}
}

type (
	U16be[T ~uint16] uint16
	U32be[T ~uint32] uint32
	U64be[T ~uint64] uint64

	I16be[T ~uint16] uint16
	I32be[T ~uint32] uint32
	I64be[T ~uint64] uint64
)

func (x U16be[T]) Get() T {
	if arch.LittleEndian {
		return T(bits.ReverseBytes16(uint16(x)))
	}

	return T(x)
}

func (x *U16be[T]) Set(v T) {
	if arch.LittleEndian {
		*x = U16be[T](bits.ReverseBytes16(uint16(v)))
	} else {
		*x = U16be[T](v)
	}
}

func (x I16be[T]) Get() T {
	if arch.LittleEndian {
		return T(bits.ReverseBytes16(uint16(x)))
	}

	return T(x)
}

func (x *I16be[T]) Set(v T) {
	if arch.LittleEndian {
		*x = I16be[T](bits.ReverseBytes16(uint16(v)))
	} else {
		*x = I16be[T](v)
	}
}

func (x U32be[T]) Get() T {
	if arch.LittleEndian {
		return T(bits.ReverseBytes32(uint32(x)))
	}

	return T(x)
}

func (x *U32be[T]) Set(v T) {
	if arch.LittleEndian {
		*x = U32be[T](bits.ReverseBytes32(uint32(v)))
	} else {
		*x = U32be[T](v)
	}
}

func (x I32be[T]) Get() T {
	if arch.LittleEndian {
		return T(bits.ReverseBytes32(uint32(x)))
	}

	return T(x)
}

func (x *I32be[T]) Set(v T) {
	if arch.LittleEndian {
		*x = I32be[T](bits.ReverseBytes32(uint32(v)))
	} else {
		*x = I32be[T](v)
	}
}

func (x U64be[T]) Get() T {
	if arch.LittleEndian {
		return T(bits.ReverseBytes64(uint64(x)))
	}

	return T(x)
}

func (x *U64be[T]) Set(v T) {
	if arch.LittleEndian {
		*x = U64be[T](bits.ReverseBytes64(uint64(v)))
	} else {
		*x = U64be[T](v)
	}
}

func (x I64be[T]) Get() T {
	if arch.LittleEndian {
		return T(bits.ReverseBytes64(uint64(x)))
	}

	return T(x)
}

func (x *I64be[T]) Set(v T) {
	if arch.LittleEndian {
		*x = I64be[T](bits.ReverseBytes64(uint64(v)))
	} else {
		*x = I64be[T](v)
	}
}
