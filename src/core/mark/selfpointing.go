// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mark

// SelfPointing marks a struct field as a pointer may pointing to data inside
// the same struct.
//
// When using, add an extra field right before the pointer whose name
// MUST be "_" + <name of the pointed field>.
//
// For example:
//
//	type T struct {
//		_storage mark.SelfPointing
//		buf      []byte
//		_foo     mark.SelfPointing
//		f        *uint
//
//		foo     uint
//		storage [1024]byte
//	}
//
// T.buf is a slice whose first field is a pointer may be pointing
// at storage, which is:
//
//	unsafe.Pointer(unsafe.SliceData(T.buf)) == unsafe.Pointer(&storage)
type SelfPointing struct{}
