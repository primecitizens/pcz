package js

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/primecitizens/pcz/cmd/build/toolchain"
	"github.com/primecitizens/pcz/cmd/utils/jsutils"
	"github.com/primecitizens/std"
)

type BindgenOptions struct {
	Mode          string
	ES            string
	ModuleName    string
	CustomWrapper string
	Deps          []*toolchain.Package
	Filter        WasmImportFilter
}

func (s *BindgenOptions) discard(module, field string) bool {
	return s.Filter != nil && !s.Filter.Keep(module, field)
}

func CreateJSBindings(spec BindgenOptions) (code string, err error) {
	var imports []string
	for _, pkg := range spec.Deps {
		file := filepath.Join(pkg.Dir, "ffi_bindings.ts")
		data, err := os.ReadFile(file)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			panic(err)
		}

		// TODO: use strings.Cut(str, "\n") to read & write line by line
		var (
			str     = unsafe.String(unsafe.SliceData(data), len(data))
			lines   = strings.SplitAfter(str, "\n")
			cur     int
			discard = true // discard statements before MODULE_PREFIX
			module  string

			totalFields     int
			discardedFields int
		)

		const (
			MODULE_PREFIX = `importModule("`
			FIELD_PREFIX  = `    "`
		)
		for i := 0; i < len(lines); i++ {
			switch {
			case strings.HasPrefix(lines[i], MODULE_PREFIX): // importModule("module", {
				module = lines[i][len(MODULE_PREFIX):]
				end := strings.IndexByte(module, '"')
				if end < 0 {
					panic(fmt.Errorf("failed to find module name end: %q", lines[i]))
				}

				module = module[:end]
				discard = false
			case strings.HasPrefix(lines[i], FIELD_PREFIX): // "field": ...
				field := lines[i][len(FIELD_PREFIX):]
				end := strings.IndexByte(field, '"')
				if end < 0 {
					panic(fmt.Errorf("failed to find field end: %q", lines[i]))
				}

				totalFields++
				if discard = spec.discard(module, field[:end]); discard {
					discardedFields++
					continue
				}
			case strings.HasPrefix(lines[i], "});"): // end of module import
				discard = true
			default:
				if discard {
					continue
				}
			}

			lines[cur] = lines[i]
			cur++
		}

		if totalFields == discardedFields {
			// this module is not used
			continue
		}

		imports = append(imports, strings.Join(lines[:cur], ""))
	}

	switch spec.Mode {
	case "raw", "":
		var sb strings.Builder
		err := assembleImports(&sb, spec.CustomWrapper, imports...)
		if err != nil {
			return "", fmt.Errorf("failed to generate bindings: %w", err)
		}

		return sb.String(), nil
	case "cjs", "commonjs":
		spec.Mode = "commonjs"
		fallthrough
	case "amd", "umd", "system":
		var sb strings.Builder
		err := assembleImports(&sb, spec.CustomWrapper, imports...)
		if err != nil {
			return "", fmt.Errorf("failed to generate bindings: %w", err)
		}

		if len(spec.ModuleName) == 0 { // in case user set it to empty
			spec.ModuleName = "bindings"
		}

		return transpileTS(sb.String(), spec)
	default:
		panic(fmt.Errorf("unsupported bindgen mode: %q", spec.Mode))
	}
}

func transpileTS(tsSource string, spec BindgenOptions) (string, error) {
	var sb strings.Builder

	script, err := jsutils.Transpile(
		jsutils.TranspileOptions{
			Src:        tsSource,
			ModuleName: spec.ModuleName,
			CompileOptions: map[string]interface{}{
				"module":  spec.Mode,
				"target":  spec.ES,
				"newLine": "lf",
			},
		},
	)
	if err != nil {
		return "", err
	}

	// ts generates umd for both "amd" and "cjs" but not for browser
	// we need to patch it to support browser globalThis
	if spec.Mode == "umd" {
		err = patchUMDForBrowser(&sb, spec.ModuleName, script)
	} else {
		_, err = io.WriteString(&sb, script)
	}

	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

const (
	TemplatePlaceholder = "//__TEMPLATE_INSERT_IMPORTS__"
)

// assembleImports bindings in TypeScript for running the wasm program.
//
// NOTE: this is meant for packages using the official go std.
func assembleImports(out io.Writer, template string, moduleImports ...string) (err error) {
	if len(template) == 0 {
		var data []byte
		data, err = fs.ReadFile(std.FS(), "ffi/js/bindings/bindgen.ts")
		if err != nil {
			panic(err)
		}

		template = string(data)
	}

	i := strings.Index(template, TemplatePlaceholder)
	if i < 0 {
		panic("placeholder not found")
	}

	j := strings.LastIndexByte(template[:i], '\n')
	indent := template[j+1 : i]

	io.WriteString(out, template[:i])
	for _, imp := range moduleImports {
		io.WriteString(out, "\n")

		for _, line := range strings.SplitAfter(imp, "\n") {
			io.WriteString(out, indent)
			io.WriteString(out, line)
		}

		io.WriteString(out, "\n")
	}

	_, err = io.WriteString(out, template[i+len(TemplatePlaceholder):])
	return
}

func patchUMDForBrowser(out io.Writer, moduleName, tscUMD string) error {
	const AMD_CHECK_LINE = `else if (typeof define === "function" && define.amd) {` + "\n"

	i := strings.Index(tscUMD, AMD_CHECK_LINE)
	if i < 0 {
		panic(fmt.Errorf("unexpected amd.js checkline %q not found in script:\n\n%s", AMD_CHECK_LINE, tscUMD))
	}

	i += len(AMD_CHECK_LINE)
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
