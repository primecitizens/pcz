// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && windows

package sys

var (
	envs []string
)

func nthenv(i int) string {
	return envs[i]
}
