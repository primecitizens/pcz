// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdgo

// Gosched yields the processor, allowing other goroutines to run. It does not
// suspend the current goroutine, so execution resumes automatically.
//
// If guarded is true, also checks for forbidden states and opts out of the
// yield in those cases.
//
//go:nosplit
func Sched(guarded bool) {
	// TODO: gopark
}
