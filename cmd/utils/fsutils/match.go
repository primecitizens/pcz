// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package fsutils

import (
	"io"
	"os"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/primecitizens/cli"
)

func MatchFiles(globs ...string) ([]string, error) {
	files := make(map[string]int)
	for _, g := range globs {
		matches, err := doublestar.FilepathGlob(
			g,
			doublestar.WithFailOnIOErrors(),
			doublestar.WithFilesOnly(),
		)
		if err != nil {
			panic(err)
		}

		for _, m := range matches {
			_, ok := files[m]
			if !ok {
				files[m] = len(files)
				continue
			}
		}
	}

	ret := make([]string, len(files))
	for k, v := range files {
		ret[v] = k
	}

	return ret, nil
}

func WriteToFileOrStdout(opts *cli.CmdOptions, file string, do func(w io.Writer) error) error {
	if len(file) == 0 {
		return do(opts.PickStdout(os.Stdout))
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0640)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return do(f)
}
