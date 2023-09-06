// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build bindgen

package js

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/clarkmcc/go-typescript"
	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/buildcfg"
	"github.com/primecitizens/std/ffi/js/bindings"
)

func runBindgen(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) error {
	bopts := route.Target().Extra.(*BindgenOptions)
	if len(posArgs) != 1 {
		panic(fmt.Errorf("expecting exactly one positional arg for main package (got %d)", len(posArgs)))
	}

	spec := bindgenSpec{
		mode:       bopts.Mode,
		es:         bopts.ES,
		moduleName: bopts.ModuleName,
	}
	if tfile := bopts.TemplateFile; len(tfile) != 0 {
		data, err := os.ReadFile(tfile)
		if err != nil {
			panic(err)
		}

		spec.template = unsafe.String(unsafe.SliceData(data), len(data))
	}

	bopts.buildOpts.Platform = bopts.Platform
	bopts.buildOpts.Tags = bopts.Tags

	tc := buildcfg.NewToolchain(&bopts.buildOpts)
	pkgs, err := tc.List(posArgs[0])
	if err != nil {
		panic(err)
	}

	var out io.Writer
	output := bopts.OutputFile
	if len(output) == 0 {
		out = opts.PickStdout(os.Stdout)
	} else {
		f, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0640)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		out = f
	}

	err = createBindings(out, pkgs, spec)
	if err != nil {
		panic(err)
	}

	return nil
}

func createBindings(out io.Writer, pkgs []buildcfg.Package, spec bindgenSpec) error {
	var imports []string
	for i := range pkgs {
		file := filepath.Join(pkgs[i].Dir, "ffi_bindings.ts")
		data, err := os.ReadFile(file)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			panic(err)
		}

		imports = append(imports, unsafe.String(unsafe.SliceData(data), len(data)))
	}

	switch spec.mode {
	case "raw":
		err := bindings.Generate(out, spec.template, imports...)
		if err != nil {
			return fmt.Errorf("failed to generate bindings: %w", err)
		}
	case "cjs", "commonjs":
		spec.mode = "commonjs"
		fallthrough
	case "amd", "umd", "system":
		var buf bytes.Buffer
		err := bindings.Generate(&buf, spec.template, imports...)
		if err != nil {
			return fmt.Errorf("failed to generate bindings: %w", err)
		}

		if len(spec.moduleName) == 0 { // in case user set it to empty
			spec.moduleName = "bindings"
		}

		script, err := typescript.Transpile(
			&buf,
			typescript.WithModuleName(spec.moduleName),
			typescript.WithCompileOptions(
				map[string]interface{}{
					"module":  spec.mode,
					"target":  spec.es,
					"newLine": "lf",
				},
			),
		)
		if err != nil {
			return err
		}

		// tsc generates umd for both "amd" and "cjs" but not for browser
		// we need to patch it to support browser globalThis
		if spec.mode == "umd" {
			err = patchUMDForBrowser(out, spec.moduleName, script)
		} else {
			_, err = io.WriteString(out, script)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func patchUMDForBrowser(out io.Writer, moduleName, tscUMD string) error {
	const AMDCheckLine = `else if (typeof define === "function" && define.amd) {` + "\n"

	i := strings.Index(tscUMD, AMDCheckLine)
	if i < 0 {
		panic(fmt.Errorf("unexpected amd.js checkline %q not found in script:\n\n%s", AMDCheckLine, tscUMD))
	}

	i += len(AMDCheckLine)
	// the tsc gnereated import block is like
	//
	// (function (factory) {
	//   if (typeof module === "object" && typeof module.exports === "object") {
	//     var v = factory(require, exports);
	//     if (v !== undefined) module.exports = v;
	//   }
	//   else if (typeof define === "function" && define.amd) {
	//     define("bindings", ["require", "exports"], factory);
	//   }
	// })(function (require, exports) {

	before, after, ok := strings.Cut(tscUMD[i:], "}\n")
	if !ok {
		panic(fmt.Errorf("unexpected amd.js define block end not found in script:\n\n%s", tscUMD))
	}
	i += len(before) + 1 /* "}" */

	_, err := io.WriteString(out, tscUMD[:i])
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(
		out,
		// TODO: handle indentation automatically by checking lines before
		"\n"+
			"    else {\n"+
			"        factory(undefined, (globalThis.%s = {}));\n"+
			"    }\n",
		moduleName,
	)
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(out, after)
	if err != nil {
		panic(err)
	}

	return nil
}
