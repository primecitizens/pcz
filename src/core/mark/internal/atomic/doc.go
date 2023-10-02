// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package atomic

// NOTE: `align64` MUST be in either runtime/internal/atomic or sync/atomic,
// as compiler special case this type.
//
// See $GOROOT/src/cmd/compile/internal/types/size.go#func:calcStructOffset
//
//pcz:importpath sync/atomic

// align64 may be added to structs that must be 64-bit aligned.
// This struct is recognized by a special case in the compiler
// and will not work if copied to any other package.
type align64 struct{}

type Align64 = align64
