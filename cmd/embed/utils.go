// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package embed

func mime(ext string) string {
	switch ext {
	case ".js":
		return "application/javascript"
	case ".css":
		return "text/css"
	case ".json":
		return "application/json"
	case ".html":
		return "text/html"
	case ".wasm":
		return "application/wasm"
	default:
		return "text/plain"
	}
}
