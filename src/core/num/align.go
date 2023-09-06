// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package num

// AlignUp rounds n up to multiple of a. a must be a power of 2.
func AlignUp[T Uint](n, a T) T {
	return (n + a - 1) & (^(a - 1))
}

// AlignDown rounds n down to multiple of a. a must be a power of 2.
func AlignDown[T Uint](n, a T) T {
	return n & (^(a - 1))
}
