// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package requirements

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/requirements/bindings"
)

type ThreeDFeature uint32

const (
	_ ThreeDFeature = iota

	ThreeDFeature_WEBGL
	ThreeDFeature_CSS_3D
)

func (ThreeDFeature) FromRef(str js.Ref) ThreeDFeature {
	return ThreeDFeature(bindings.ConstOfThreeDFeature(str))
}

func (x ThreeDFeature) String() (string, bool) {
	switch x {
	case ThreeDFeature_WEBGL:
		return "webgl", true
	case ThreeDFeature_CSS_3D:
		return "css3d", true
	default:
		return "", false
	}
}
