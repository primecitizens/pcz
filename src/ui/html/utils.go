// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build js

package html

import (
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web"
)

type ElementFinder interface {
	GetElementById(id js.String) (web.Element, bool)
}

// FindElement by id.
func FindElement[
	T interface{ FromRef(js.Ref) T },
	F ElementFinder,
](out *T, root F, id string) bool {
	ref := js.Must(root.GetElementById(js.NewString(id).Once())).Ref()
	if ref.Undefined() {
		return false
	}

	*out = (*out).FromRef(ref)
	return true
}
