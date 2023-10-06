// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package alloc

// A for arena creations.
type A interface {
	NewArena() Arena
}

// Arena defines the arena interface.
type Arena interface {
	M
	// Discard this arena, all allocated data will be gone.
	Discard()
}
