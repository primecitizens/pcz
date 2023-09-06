// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdtype

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/cerr"
)

type error = cerr.E

// A TypeAssertionError explains a failed type assertion.
type TypeAssertionError struct {
	_interface    *abi.Type
	concrete      *abi.Type
	asserted      *abi.Type
	missingMethod string // one method needed by Interface, missing from Concrete
}

func (*TypeAssertionError) RuntimeError() {}

// WriteErr implements cerr.E
func (e *TypeAssertionError) WriteErr(w cerr.Writer) int {
	inter := "interface"
	if e._interface != nil {
		inter = e._interface.String()
	}

	n := cerr.WriteJoinS(w, " ", "interface", "conversion") + w.Write(": ")
	as := e.asserted.String()
	if e.concrete == nil {
		return n + cerr.WriteJoinS(w, " ", inter, "is", "nil") + w.Write(", ") +
			cerr.WriteJoinS(w, " ", "not", as)
	}

	cs := e.concrete.String()
	if len(e.missingMethod) == 0 {
		n += cerr.WriteJoinS(w, " ", inter, "is", cs) + w.Write(", ") +
			cerr.WriteJoinS(w, " ", "not", as)

		if cs == as {
			// provide slightly clearer error message
			n += w.Write(" (")

			if e.concrete.PkgPath() != e.asserted.String() {
				n += cerr.WriteJoinS(w, " ", "types", "from", "different", "packages")
			} else {
				n += cerr.WriteJoinS(w, " ", "types", "from", "different", "scopes")
			}

			n += w.Write(")")
		}

		return n
	}

	return n + cerr.WriteJoinS(w, " ", cs, "is", "not", as) + w.Write(": ") +
		cerr.WriteJoinS(w, " ", "missing", "method", e.missingMethod)
}
