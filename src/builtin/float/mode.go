// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz && (amd64 || arm64 || 386)

package stdfloat

func SetFPUMode(mode FPUMode)
func GetFPUMode() FPUMode
