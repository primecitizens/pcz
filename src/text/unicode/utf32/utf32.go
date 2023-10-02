package utf32

import (
	"math/bits"
	"unsafe"

	"github.com/primecitizens/std/core/arch"
)

func AsString(s []uint32) String {
	if arch.BigEndian {
		for i, x := range s {
			s[i] = bits.ReverseBytes32(x)
		}
	}

	return String(
		unsafe.String(
			(*byte)(unsafe.Pointer(unsafe.SliceData(s))), len(s)*4,
		),
	)
}

type String string

func (s String) Slice() []uint32 {
	ret := unsafe.Slice(
		(*uint32)(unsafe.Pointer(unsafe.StringData(string(s)))), len(s)/4,
	)

	if arch.BigEndian {
		for i, x := range ret {
			ret[i] = bits.ReverseBytes32(x)
		}
	}

	return ret
}
