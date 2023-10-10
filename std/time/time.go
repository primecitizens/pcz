// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package time

type Time struct {
	Sec  int64
	NSec int32

	Loc *Location
}
