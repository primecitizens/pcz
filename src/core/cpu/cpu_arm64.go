// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

// CacheLinePadSize is used to prevent false sharing of cache lines.
// We choose 128 because Apple Silicon, a.k.a. M1, has 128-byte cache line size.
// It doesn't cost much and is much more future-proof.
const CacheLinePadSize = 128

func doinit() *ARM64Features {
	// arm64 uses different ways to detect CPU features at runtime depending on the operating system.
	return osInit()
}

func getisar0() uint64

func getMIDR() uint64

func extractBits(data uint64, start, end uint) uint {
	return (uint)(data>>start) & ((1 << (end - start + 1)) - 1)
}

func parseARM64SystemRegisters(isar0 uint64) {
	// ID_AA64ISAR0_EL1
	switch extractBits(isar0, 4, 7) {
	case 1:
		ARM64 |= ARM64Feature_aes
	case 2:
		ARM64 |= ARM64Feature_aes | ARM64Feature_pmull
	}

	switch extractBits(isar0, 8, 11) {
	case 1:
		ARM64 |= ARM64Feature_sha1
	}

	switch extractBits(isar0, 12, 15) {
	case 1:
		ARM64 |= ARM64Feature_sha2
	case 2:
		ARM64 |= ARM64Feature_sha2 | ARM64Feature_sha512
	}

	switch extractBits(isar0, 16, 19) {
	case 1:
		ARM64 |= ARM64Feature_crc32
	}

	switch extractBits(isar0, 20, 23) {
	case 2:
		ARM64 |= ARM64Feature_atomics
	}
}
