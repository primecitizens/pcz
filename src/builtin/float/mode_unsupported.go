// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !(amd64 || arm64 || 386)

package stdfloat

type FPUMode uint8

func SetFPUMode(mode FPUMode) {}
func GetFPUMode() FPUMode     { return 0 }
