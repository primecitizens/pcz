// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package embed

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/utils/fsutils"
)

var (
	jsEmbedOptions = JSEmbedOptions{}
	jsEmbedCmd     = cli.Cmd{
		Pattern:    "js",
		BriefUsage: "Create new js source file with data embedded",
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &jsEmbedOptions),
		FlagRule:   nil,
		Run:        runJSEmbed,
	}
)

type JSEmbedOptions struct {
	// for cli only, not used by the EmbedJS func
	Output string `cli:"o|output,#set path to write embedded content, if empty, write to stdout"`

	Mode  string            `cli:"m|mode,def=uint8array,comp=uint8array,comp=base64,comp=data-url,#set embedding mode"`
	Files map[string]string `cli:"f|file,#add embedded file in format \"[export] {var|const|let} <var-name>=<file-path>\""`
}

func runJSEmbed(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) error {
	var buf bytes.Buffer
	if err := EmbedJS(&buf, jsEmbedOptions); err != nil {
		return err
	}

	return fsutils.WriteToFileOrStdout(opts, jsEmbedOptions.Output, func(w io.Writer) error {
		_, err := buf.WriteTo(w)
		return err
	})
}

func EmbedJS(out io.StringWriter, opts JSEmbedOptions) error {
	var encodeContent func(path string, data []byte) error
	switch mode := opts.Mode; mode {
	case "uint8array":
		encodeContent = func(path string, data []byte) error {
			out.WriteString("new Uint8Array([")
			for i, b := range data {
				if i%64 == 0 {
					out.WriteString("\n  ")
				}

				out.WriteString(strconv.FormatUint(uint64(b), 10))
				out.WriteString(", ")
			}
			if len(data) > 0 {
				out.WriteString("\n")
			}
			_, err := out.WriteString("]);")
			return err
		}
	case "base64":
		encodeContent = func(path string, data []byte) error {
			out.WriteString(`"`)
			out.WriteString(base64.StdEncoding.EncodeToString(data))
			_, err := out.WriteString(`";`)
			return err
		}
	case "data-url":
		encodeContent = func(path string, data []byte) error {
			out.WriteString(`"data:`)

			contentType := "text/plain"
			if i := strings.LastIndexByte(path, '.'); i >= 0 {
				contentType = mime(path[i:])
			}

			out.WriteString(contentType)
			out.WriteString(",")
			out.WriteString(base64.StdEncoding.EncodeToString(data))
			_, err := out.WriteString(`"`)
			return err
		}
	default:
		return fmt.Errorf("unsupported embedding mode %q", mode)
	}

	type EmbedVar struct {
		Name       string
		SourcePath string
		Content    []byte
	}

	i := 0
	variables := make([]EmbedVar, len(opts.Files))
	for decl, file := range opts.Files {
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		variables[i] = EmbedVar{
			Name:       strings.TrimSpace(decl),
			SourcePath: file,
			Content:    data,
		}
	}

	slices.SortFunc(variables, func(a, b EmbedVar) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, v := range variables {
		out.WriteString(v.Name)
		out.WriteString(" = ")
		err := encodeContent(v.SourcePath, v.Content)
		if err != nil {
			return fmt.Errorf("failed to encode embedded content for %q (source: %q): %w", v.Name, v.SourcePath, err)
		}
		out.WriteString("\n")
	}

	return nil
}
