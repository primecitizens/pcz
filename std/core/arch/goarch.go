// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arch contains GOARCH-specific constants.
package arch

type ArchFamilyType int

const (
	AMD64 ArchFamilyType = iota
	ARM
	ARM64
	I386
	MIPS
	MIPS64
	PPC64
	RISCV64
	S390X
	WASM
)

const (
	// Is64bit = 1 on 64-bit systems, 0 on 32-bit systems
	Is64bit = 1 << (^uintptr(0) >> 63) / 2

	// PtrSize is the size of a pointer in bytes - unsafe.Sizeof(uintptr(0)) but as an ideal constant.
	// It is also the size of the machine's native word size (that is, 4 on 32-bit systems, 8 on 64-bit).
	PtrSize = 4 << (^uintptr(0) >> 63)

	// 32 on 32-bit systems, 64 on 64-bit.
	UintBits = 32 << (^uint(0) >> 63)

	// ArchFamily is the architecture family (AMD64, ARM, ...)
	ArchFamily ArchFamilyType = _ArchFamily

	// DefaultPhysPageSize is the default physical page size.
	DefaultPhysPageSize = _DefaultPhysPageSize

	// PCQuantum is the minimal unit for a program counter (1 on x86, 4 on most other systems).
	// The various PC tables record PC deltas pre-divided by PCQuantum.
	PCQuantum = _PCQuantum

	// Int64Align is the required alignment for a 64-bit integer (4 on 32-bit systems, 8 on 64-bit).
	Int64Align = PtrSize

	// MinFrameSize is the size of the system-reserved words at the bottom
	// of a frame (just above the architectural stack pointer).
	// It is zero on x86 and PtrSize on most non-x86 (LR-based) systems.
	// On PowerPC it is larger, to cover three more reserved words:
	// the compiler word, the link editor word, and the TOC save word.
	MinFrameSize = _MinFrameSize

	// StackAlign is the required alignment of the SP register.
	// The stack must be at least word aligned, but some architectures require more.
	StackAlign = _StackAlign
)

const (
	PageShift = 13
	PageSize  = 1 << PageShift
)

const (
	// BigEndian reports whether the architecture is big-endian.
	BigEndian    = IsArmbe|IsArm64be|IsMips|IsMips64|IsMips64p32|IsPpc|IsPpc64|IsS390|IsS390x|IsSparc|IsSparc64 != 0
	LittleEndian = !BigEndian
)

const (
	FramePointerEnabled = IsAmd64|IsArm64 != 0
)
