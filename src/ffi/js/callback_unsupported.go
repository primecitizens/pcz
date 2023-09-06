//go:build noos || !(js && wasm)

package js

import (
	"github.com/primecitizens/std/core/assert"
)

func callHandler(
	recv, targetPC uintptr, ctx *CallbackContext, fn dispFunc[uintptr],
) {
	assert.Throw("unsupported", "operation")
}
