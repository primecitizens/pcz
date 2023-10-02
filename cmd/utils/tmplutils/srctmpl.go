// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package tmplutils

import "strings"

// CleanupSourceTemplate cleans up special comments in a source template.
//
// A source template is a file containing source code in target langauge
// but contains special double slash comments for templating:
//
//   - `// +` will be replaced with empty string and the line next to it will
//     be removed.
//   - `// {{` will be replaced with `{{`
//
// This makes the template syntactically correct in the target language, which
// improves readibility.
func CleanupSourceTemplate(s string) string {
	var (
		cur   int
		lines = strings.SplitAfter(s, "\n")
	)

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		switch {
		case strings.HasPrefix(line, "// + "): // remove next line (skip empty comment lines) and prefix
			before, after, _ := strings.Cut(lines[i], "// + ")
			lines[cur] = before + after
			i++
			for strings.TrimSpace(lines[i]) == "//" {
				i++
			}
		case line == "// +":
			i++
			for strings.TrimSpace(lines[i]) == "//" {
				i++
			}
			cur--
		case strings.HasPrefix(line, "// {{"): // remove prefix before {{
			before, after, _ := strings.Cut(lines[i], "// {{")
			lines[cur] = before + "{{" + after
		default:
			lines[cur] = lines[i]
		}

		cur++
	}

	s = strings.Join(lines[:cur], "")
	return s
}
