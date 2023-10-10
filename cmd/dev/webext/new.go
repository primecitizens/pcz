// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webext

import (
	"bytes"
	_ "embed" // for go:embed
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"text/template"

	"github.com/huandu/xstrings"
	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/utils/tmplutils"
)

var (
	projOpts = ProjectOptions{
		ModuleName: "bindings",
	}
	cmdNew = cli.Cmd{
		Pattern:    "new [options] [DIR]",
		BriefUsage: "Create a new web extension project",
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &projOpts),
		Run:        createNewProject,
	}
)

type ProjectOptions struct {
	Name string `cli:"name,#set name of the extension"`
	URL  string `cli:"url,#set url to the project home"`

	ModuleName string `cli:"module-name,#set module name for js bindings"`

	All           bool `cli:"all,#add all components"`
	Popup         bool `cli:"popup,#add popup component"`
	ServiceWorker bool `cli:"service-worker|background,#add service-worker component"`
	Settings      bool `cli:"settings,#add extension settings component (options_page)"`
	SidePanel     bool `cli:"sidepanel,#add sidepanel component"`
}

func createNewProject(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) error {
	var dir string
	switch len(posArgs) {
	case 0:
		if len(projOpts.Name) == 0 {
			return fmt.Errorf("at lease one of project name and dir path should be set, got none")
		}

		dir = xstrings.ToKebabCase(projOpts.Name)
		_, err := os.Stat(dir)
		if err == nil {
			return fmt.Errorf(
				"path %q already exits, please add a potional arg"+
					" to specify the desired directory",
				dir,
			)
		}
	case 1:
		dir = posArgs[0]

		if len(projOpts.Name) == 0 {
			projOpts.Name = filepath.Base(dir)
		}
	default:
		return fmt.Errorf("expecting at most one positional arg, got %d", len(posArgs))
	}

	files := CreateProject(dir, projOpts)

	_ = os.MkdirAll(dir, 0750)
	for _, f := range files {
		err := os.WriteFile(f.Path, f.Content, 0640)
		if err != nil {
			panic(fmt.Errorf("failed to write file %q: %w", f.Path, err))
		}
	}

	return nil
}

type File struct {
	Path    string
	Content []byte
}

func CreateProject(dir string, opts ProjectOptions) []File {
	if opts.All {
		opts.Popup = true
		opts.ServiceWorker = true
		opts.Settings = true
		opts.SidePanel = true
	}

	if len(opts.ModuleName) == 0 {
		opts.ModuleName = "bindings"
	}

	if len(opts.URL) == 0 {
		opts.URL = "https://"
	}

	manifest, err := CreateManifest(opts)
	if err != nil {
		panic(fmt.Errorf("failed to create manifest.json: %w", err))
	}

	gocode, err := format.Source(createGoMain(opts))
	if err != nil {
		panic(err)
	}

	files := []File{
		{
			Path:    filepath.Join(dir, "README.md"),
			Content: createReadme(opts),
		},
		{
			Path:    filepath.Join(dir, ".gitignore"),
			Content: gitignoreTemplate,
		},
		{
			Path:    filepath.Join(dir, "manifest.json"),
			Content: manifest,
		},
		{
			Path:    filepath.Join(dir, "main.go"),
			Content: gocode,
		},
		{
			Path:    filepath.Join(dir, "go.mod"),
			Content: gomodTemplate,
		},
	}

	if opts.Popup {
		html, js := CreatePopupPage(opts.ModuleName, "")
		files = append(files,
			File{
				Path:    filepath.Join(dir, "popup.html"),
				Content: html,
			},
			File{
				Path:    filepath.Join(dir, "popup.js"),
				Content: js,
			},
			File{
				Path:    filepath.Join(dir, "popup.css"),
				Content: nil,
			},
		)
	}

	if opts.ServiceWorker {
		js := CreateServiceWorker(opts.ModuleName, "")
		files = append(files,
			File{
				Path:    filepath.Join(dir, "service_worker.js"),
				Content: js,
			},
		)
	}

	if opts.Settings {
		html, js := CreateSettingsPage(opts.ModuleName, "")
		files = append(files,
			File{
				Path:    filepath.Join(dir, "settings.html"),
				Content: html,
			},
			File{
				Path:    filepath.Join(dir, "settings.js"),
				Content: js,
			},
			File{
				Path:    filepath.Join(dir, "settings.css"),
				Content: nil,
			},
		)
	}

	if opts.SidePanel {
		html, js := CreateSidePanel(opts.ModuleName, "")
		files = append(files,
			File{
				Path:    filepath.Join(dir, "sidepanel.html"),
				Content: html,
			},
			File{
				Path:    filepath.Join(dir, "sidepanel.js"),
				Content: js,
			},
			File{
				Path:    filepath.Join(dir, "sidepanel.css"),
				Content: nil,
			},
		)
	}

	return files
}

func CreatePopupPage(moduleName, wasmEmbed string) (html, js []byte) {
	opts := &loaderOptions{
		Comp:       "popup",
		ModuleName: moduleName,
		WasmEmbed:  wasmEmbed,
	}
	html = createHTML(opts)
	js = createJS(opts)
	return
}

func CreateSettingsPage(moduleName, wasmEmbed string) (html, js []byte) {
	opts := &loaderOptions{
		Comp:       "settings",
		ModuleName: moduleName,
		WasmEmbed:  wasmEmbed,
	}
	html = createHTML(opts)
	js = createJS(opts)
	return
}

func CreateServiceWorker(moduleName, wasmEmbed string) (js []byte) {
	opts := &loaderOptions{
		Comp:       "service-worker",
		ModuleName: moduleName,
		WasmEmbed:  wasmEmbed,
	}
	js = createJS(opts)
	return
}

func CreateSidePanel(moduleName, wasmEmbed string) (html, js []byte) {
	opts := &loaderOptions{
		Comp:       "sidepanel",
		ModuleName: moduleName,
		WasmEmbed:  wasmEmbed,
	}
	html = createHTML(opts)
	js = createJS(opts)
	return
}

func CreateManifest(opts ProjectOptions) ([]byte, error) {
	x := &ManifestBasicSchema{
		ManifestVersion: 3,
		Name:            opts.Name,
		HomepageURL:     opts.URL,
		Version:         "0.1.0",
		CSP: ManifestCSP{
			ExtensionPages: `default-src 'self' 'wasm-unsafe-eval'`,
		},
	}

	if opts.Settings {
		x.OptionsPage = "options.html"
	}

	if opts.Popup {
		x.Action = &ManifestAction{
			DefaultTitle: "My Popup Action",
			DefaultPopup: "popup.html",
		}
		x.Permissions = append(x.Permissions, "action")
	}

	if opts.ServiceWorker {
		x.Background = &ManifestBackground{
			ServiceWorker: "service_worker.js",
		}
	}

	if opts.SidePanel {
		x.SidePanel = &ManifestSidePanel{
			DefaultPath: "sidepanel.html",
		}
		x.Permissions = append(x.Permissions, "sidePanel")
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")

	err := enc.Encode(x)
	return buf.Bytes(), err
}

type ManifestBasicSchema struct {
	ManifestVersion int    `json:"manifest_version"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	HomepageURL     string `json:"homepage_url"`
	Version         string `json:"version"`

	Icons struct {
		N16  string `json:"16"`
		N48  string `json:"48"`
		N128 string `json:"128"`
		N256 string `json:"256"`
	} `json:"icons"`

	OptionsPage string   `json:"options_page,omitempty"`
	Permissions []string `json:"permissions"`

	CSP ManifestCSP `json:"content_security_policy"`

	SidePanel  *ManifestSidePanel  `json:"side_panel,omitempty"`
	Background *ManifestBackground `json:"background,omitempty"`
	Action     *ManifestAction     `json:"action,omitempty"`
}

type ManifestSidePanel struct {
	DefaultPath string `json:"default_path"`
}

type ManifestBackground struct {
	ServiceWorker string `json:"service_worker"`
}

type ManifestAction struct {
	DefaultTitle string `json:"default_title"`
	DefaultPopup string `json:"default_popup"`
	DefaultIcon  struct {
		N16 string `json:"16"`
		N24 string `json:"24"`
		N32 string `json:"32"`
	} `json:"default_icon"`
}

type ManifestCSP struct {
	ExtensionPages string `json:"extension_pages"`
}

var (
	//go:embed templates/loader.html
	htmlTemplate string

	//go:embed templates/loader.js
	jsTemplate string

	//go:embed templates/main.go
	goTemplate string

	//go:embed templates/gomod
	gomodTemplate []byte

	//go:embed templates/README.md
	readmeTemplate string

	//go:embed templates/gitignore
	gitignoreTemplate []byte
)

type loaderOptions struct {
	// Comp can be one of
	//  - service-worker
	//  - sidepanel
	//  - popup
	//  - settings
	Comp string

	// ModuleName is the module name of the js bindings.
	ModuleName string

	// WasmEmbed is the variable reference to the imported wasm
	//
	// defaults to "wasmBlob.buffer"
	WasmEmbed string
}

func createReadme(data ProjectOptions) []byte {
	ret, err := execTmpl(readmeTemplate, &data)
	if err != nil {
		panic(fmt.Errorf("failed to create readme file: %w", err))
	}

	return ret
}

func createHTML(data *loaderOptions) []byte {
	if len(data.ModuleName) == 0 {
		data.ModuleName = "bindings"
	}

	if len(data.WasmEmbed) == 0 {
		data.WasmEmbed = "wasmBlob.buffer"
	}

	ret, err := execTmpl(htmlTemplate, data)
	if err != nil {
		panic(fmt.Errorf("failed to create html file: %w", err))
	}

	return ret
}

func createJS(data *loaderOptions) []byte {
	if len(data.ModuleName) == 0 {
		data.ModuleName = "bindings"
	}

	if len(data.WasmEmbed) == 0 {
		data.WasmEmbed = "wasmBlob.buffer"
	}

	ret, err := execTmpl(jsTemplate, data)
	if err != nil {
		panic(fmt.Errorf("failed to create js loader file: %w", err))
	}

	return ret
}

func createGoMain(data ProjectOptions) []byte {
	ret, err := execTmpl(goTemplate, data)
	if err != nil {
		panic(fmt.Errorf("failed to create go file: %w", err))
	}

	return ret
}

func execTmpl(t string, data any) (ret []byte, err error) {
	var buf bytes.Buffer

	tmpl, err := template.New("").Parse(tmplutils.CleanupSourceTemplate(t))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.Bytes(), nil
}
