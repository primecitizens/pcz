// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdcomplex

import (
	stdfloat "github.com/primecitizens/pcz/std/builtin/float"
)

func Complex128Div(n, m complex128) complex128 {
	var e, f float64 // complex(e, f) = n/m

	// Algorithm for robust complex division as described in
	// Robert L. Smith: Algorithm 116: Complex division. Commun. ACM 5(8): 435 (1962).
	if stdfloat.Abs(real(m)) >= stdfloat.Abs(imag(m)) {
		ratio := imag(m) / real(m)
		denom := real(m) + ratio*imag(m)
		e = (real(n) + imag(n)*ratio) / denom
		f = (imag(n) - real(n)*ratio) / denom
	} else {
		ratio := real(m) / imag(m)
		denom := imag(m) + ratio*real(m)
		e = (real(n)*ratio + imag(n)) / denom
		f = (imag(n)*ratio - real(n)) / denom
	}

	if stdfloat.IsNaN(e) && stdfloat.IsNaN(f) {
		// Correct final result to infinities and zeros if applicable.
		// Matches C99: ISO/IEC 9899:1999 - G.5.1  Multiplicative operators.

		a, b := real(n), imag(n)
		c, d := real(m), imag(m)

		switch {
		case m == 0 && (!stdfloat.IsNaN(a) || !stdfloat.IsNaN(b)):
			e = stdfloat.CopySign(inf, c) * a
			f = stdfloat.CopySign(inf, c) * b

		case (stdfloat.IsInf(a) || stdfloat.IsInf(b)) && stdfloat.IsFinite(c) && stdfloat.IsFinite(d):
			a = inf2one(a)
			b = inf2one(b)
			e = inf * (a*c + b*d)
			f = inf * (b*c - a*d)

		case (stdfloat.IsInf(c) || stdfloat.IsInf(d)) && stdfloat.IsFinite(a) && stdfloat.IsFinite(b):
			c = inf2one(c)
			d = inf2one(d)
			e = 0 * (a*c + b*d)
			f = 0 * (b*c - a*d)
		}
	}

	return complex(e, f)
}

// inf2one returns a signed 1 if f is an infinity and a signed 0 otherwise.
// The sign of the result is the sign of f.
func inf2one(f float64) float64 {
	g := 0.0
	if stdfloat.IsInf(f) {
		g = 1.0
	}
	return stdfloat.CopySign(g, f)
}
