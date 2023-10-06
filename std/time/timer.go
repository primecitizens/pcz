// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package time

type Timer interface {
	Reset(Duration) bool
}
