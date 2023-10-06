// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package opt

// Interface defines required methods for an optional value.
type Interface[Value any] interface {
	// Set the value.
	Set(Value)

	// Get the value, IsNone() MUST be checked before calling this method
	Get() Value

	// Erase the existing value and returns true if there was value.
	//
	// After calling this, IsNone() will return true until Set(T) is called.
	Erase() (hadValue bool)

	// IsNone returns true if there is no value set.
	IsNone() bool

	// IsSome returns true if there is value set.
	IsSome() bool
}
