//go:build bindgen

package bindings

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
)

const (
	ImportLine          = `import { heap, HEAP, addModuleImport, loadArgs } from "@ffi";` + "\n"
	TemplatePlaceholder = "//__TEMPLATE_INSERT_IMPORTS__"
)

func removeImports(moduleImports string) string {
	ret, ok := strings.CutPrefix(moduleImports, ImportLine)
	if !ok {
		return moduleImports
	}

	// deny any other import statements
	ret = strings.TrimSpace(ret)
	if strings.Index(ret, "\nimport ") != -1 {
		return moduleImports
	}

	return ret
}

//go:embed bindgen.ts
var bindgen_ts string

// Generate bindings in TypeScript for running the wasm program.
//
// NOTE: this is meant for packages using the official go std.
func Generate(out io.Writer, template string, moduleImports ...string) (err error) {
	if len(template) == 0 {
		template = bindgen_ts
	}

	i := strings.Index(template, TemplatePlaceholder)
	if i < 0 {
		panic("placeholder not found")
	}

	j := strings.LastIndexByte(template[:i], '\n')
	indent := template[j+1 : i]

	_, _ = io.WriteString(out, template[:i])
	for i, imp := range moduleImports {
		imp = removeImports(imp)
		if len(imp) == len(moduleImports[i]) {
			return fmt.Errorf("invalid module imports: invalid import statement\n\n%s\n", imp)
		}

		_, _ = io.WriteString(out, "\n")

		for _, line := range strings.SplitAfter(imp, "\n") {
			_, _ = io.WriteString(out, indent)
			_, _ = io.WriteString(out, line)
		}

		_, _ = io.WriteString(out, "\n")
	}

	_, err = io.WriteString(out, template[i+len(TemplatePlaceholder):])
	return
}
