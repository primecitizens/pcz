// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package tmplutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanupSourceTemplate(t *testing.T) {
	for _, test := range []struct {
		tmpl     string
		expected string
	}{
		{"// + foo\nbar", "foo\n"},
		{"  // + foo\r\n//\n//\nbar\n", "  foo\r\n"},
		{"// +\nfoo\nbar", "bar"},
		{"  // {{ . }}", "  {{ . }}"},
	} {
		assert.EqualValues(t, test.expected, CleanupSourceTemplate(test.tmpl))
	}
}
