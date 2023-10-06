// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package jsutils

import (
	_ "embed" // for go:embed
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
)

//go:embed typescript-5.2.2.js
var tsSource string

type TranspileOptions struct {
	// Src is the typescript source code to be transpiled.
	Src            string
	ModuleName     string
	CompileOptions map[string]any
}

func Transpile(opts TranspileOptions) (js string, err error) {
	var (
		vm       = goja.New()
		optsJSON []byte
	)
	vm.SetParserOptions(parser.WithDisableSourceMaps)

	if optsJSON, err = json.Marshal(opts.CompileOptions); err != nil {
		panic(err)
	} else if _, err = vm.RunString(tsSource); err != nil {
		panic(err)
	} else if err = vm.GlobalObject().Set("src", opts.Src); err != nil {
		panic(err)
	} else if err = vm.GlobalObject().Set("modulename", opts.ModuleName); err != nil {
		panic(err)
	}

	// function transpile(
	//   input: string,
	//   compilerOptions?: CompilerOptions,
	//   fileName?: string,
	//   diagnostics?: Diagnostic[],
	//   moduleName?: string,
	// ): string;
	script := fmt.Sprintf(
		`ts.transpile(src, %s, undefined, undefined, modulename)`,
		unsafe.String(unsafe.SliceData(optsJSON), len(optsJSON)),
	)

	val, err := vm.RunString(script)
	if err != nil {
		panic(err)
	}

	return val.String(), nil
}
