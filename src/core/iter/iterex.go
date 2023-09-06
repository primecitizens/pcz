// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package iter

// CoreEx is like Core but accepts an extra argument.
type CoreEx[Arg, Elem any] interface {
	NthEx(Arg, int) (Elem, bool)
}

// FiniteEx is like Finite but takes an extra argument.
type FiniteEx[Arg any] interface {
	LenEx(Arg) int
}

// DivisibleEx is Divisible with an extra argument.
type DivisibleEx[Arg, Self any] interface {
	SliceFromEx(arg Arg, start int) Self
}

type InterfaceEx[Arg, Elem, Self any] interface {
	CoreEx[Arg, Elem]
	FiniteEx[Arg]
	DivisibleEx[Arg, Self]
}

// FuncEx implements CoreEx for function.
type FuncEx[Arg, Elem any] func(Arg, int) (Elem, bool)

func (fn FuncEx[Arg, Elem]) NthEx(arg Arg, i int) (Elem, bool) {
	return fn(arg, i)
}
