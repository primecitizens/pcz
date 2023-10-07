// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package incognito

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/incognito/bindings"
)

type IncognitoMode uint32

const (
	_ IncognitoMode = iota

	IncognitoMode_SPLIT
	IncognitoMode_SPANNING
	IncognitoMode_NOT_ALLOWED
)

func (IncognitoMode) FromRef(str js.Ref) IncognitoMode {
	return IncognitoMode(bindings.ConstOfIncognitoMode(str))
}

func (x IncognitoMode) String() (string, bool) {
	switch x {
	case IncognitoMode_SPLIT:
		return "split", true
	case IncognitoMode_SPANNING:
		return "spanning", true
	case IncognitoMode_NOT_ALLOWED:
		return "not_allowed", true
	default:
		return "", false
	}
}
