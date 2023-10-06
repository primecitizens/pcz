// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package jsutils

import (
	"errors"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

type MinifyOptions struct {
	SourceCode string
	ES         api.Target
}

// Minify javascript source code using esbuild (no typescript support).
func Minify(opts MinifyOptions) (code []byte, sourceMap []byte, err error) {
	result := api.Transform(opts.SourceCode, api.TransformOptions{
		Loader: api.LoaderJS,
		Target: opts.ES,

		Sourcemap: api.SourceMapExternal,

		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,

		Color:        api.ColorNever,
		LogLevel:     api.LogLevelSilent,
		MangleQuoted: api.MangleQuotedFalse,
		TreeShaking:  api.TreeShakingTrue,
		Charset:      api.CharsetASCII,

		Format: api.FormatDefault, // we should have decided what formats to use when transpiling

		Banner: "" +
			"// SPDX-License-Identifier: Apache-2.0\n" +
			"// Copyright 2023 The Prime Citizens\n",
	})

	code = result.Code
	sourceMap = result.Map

	if len(result.Errors) > 0 {
		var sb strings.Builder
		for _, msg := range result.Errors {
			sb.WriteString(msg.Text)
			sb.WriteString("\n")
		}
		err = errors.New(sb.String())
	}

	return
}
