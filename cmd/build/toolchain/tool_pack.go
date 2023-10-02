// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"os"

	"github.com/primecitizens/pcz/cmd/utils/fsutils"
)

func (t *Toolchain) Pack(archiveFile string, objFiles ...string) (newArchiveFile string, err error) {
	newArchiveFile = archiveFile
	if len(objFiles) == 0 {
		return
	}

	err = fsutils.TemporaryWritable(
		func() error {
			return t.run("", nil, os.Stdout, os.Stderr,
				append([]string{t.binPack, "r", archiveFile}, objFiles...)...,
			)
		},
		archiveFile,
	)

	return
}
