// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// NOTE: `nih` MUST be in this exact package (runtime/internal/sys), as compiler
// special case this type.
//
// See $GOROOT/src/cmd/compile/internal/types/size.go#func:CalcSize
//
//pcz:package runtime/internal/sys
package sys

// NOTE: keep in sync with cmd/compile/internal/types.CalcSize
// to make the compiler recognize this as an intrinsic type.
type nih struct{}

type NotInHeap struct{ _ nih }
