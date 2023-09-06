// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package cmd

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSingleTask(t *testing.T) {
	var sb strings.Builder
	st := NewSingleTask(50*time.Millisecond, func() {
		sb.WriteString("foo")
		time.Sleep(100 * time.Millisecond)
	})

	for i := 0; i < 10; i++ {
		go st.Run()
	}

	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		go st.Run()
	}

	time.Sleep(time.Second)
	assert.EqualValues(t, "foofoofoo", sb.String())
}
