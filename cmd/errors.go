// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package cmd

type ErrFeatureNotBuilt struct {
	Feature string
	Tag     string
}

func (err *ErrFeatureNotBuilt) Error() string {
	return "feature `" + err.Feature + "` is not built into this pcz, " +
		"try to build pcz with go build tag `" + err.Tag + "`"
}
