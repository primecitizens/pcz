// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package fsutils

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

func Chmod(path string, perm fs.FileMode, ignoreNotExist bool) error {
	err := os.Chmod(path, perm)
	if err != nil {
		if os.IsNotExist(err) && ignoreNotExist {
			return nil
		}

		return err
	}

	return nil
}

func WriteReadOnly(path string, data []byte) (err error) {
	if err = Chmod(path, 0600, true); err != nil {
		return
	} else if err = os.Remove(path); err != nil && !os.IsNotExist(err) {
		return
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0440)
	if err != nil {
		return
	}
	defer f.Close()

	n, err := f.Write(data)
	if err != nil {
		return
	}

	if n != len(data) {
		panic(fmt.Errorf("not all data wrote (%d/%d): %w", n, len(data), io.ErrShortWrite))
	}

	return nil
}

func TemporaryWritable(do func() error, files ...string) (err error) {
	defer func() {
		var xerr []error
		for _, f := range files {
			if len(f) == 0 {
				continue
			}

			if x := Chmod(f, 0440, false); x != nil {
				xerr = append(xerr, x)
			}
		}

		err = errors.Join(xerr...)
	}()

	for _, f := range files {
		if len(f) == 0 {
			continue
		}

		if err = Chmod(f, 0600, true); err != nil {
			return
		}
	}

	if err = do(); err != nil {
		return
	}

	return nil
}
