// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdconst

import "unsafe"

const (
	SizeStringType = unsafe.Sizeof(string(""))
	SizeSliceType  = unsafe.Sizeof([]byte(nil))
	SizeEfaceType  = unsafe.Sizeof(any((*byte)(nil)))
	SizeIfaceType  = unsafe.Sizeof(iface((*ifaceImpl)(nil)))

	SizeUnsafePointer = unsafe.Sizeof(unsafe.Pointer(nil))

	SizeUintType    = unsafe.Sizeof(uint(0))
	SizeUint8Type   = unsafe.Sizeof(uint8(0))
	SizeUint16Type  = unsafe.Sizeof(uint16(0))
	SizeUint32Type  = unsafe.Sizeof(uint32(0))
	SizeUint64Type  = unsafe.Sizeof(uint64(0))
	SizeUintptrType = unsafe.Sizeof(uintptr(0))

	SizeIntType   = unsafe.Sizeof(int(0))
	SizeInt8Type  = unsafe.Sizeof(int8(0))
	SizeInt16Type = unsafe.Sizeof(int16(0))
	SizeInt32Type = unsafe.Sizeof(int32(0))
	SizeInt64Type = unsafe.Sizeof(int64(0))

	SizeFloat32Type = unsafe.Sizeof(float32(0))
	SizeFloat64Type = unsafe.Sizeof(float64(0))

	SizeComplex64Type  = unsafe.Sizeof(complex64(0))
	SizeComplex128Type = unsafe.Sizeof(complex128(0))
)

type iface interface{ isiface() }

type ifaceImpl struct{}

func (*ifaceImpl) isiface()
