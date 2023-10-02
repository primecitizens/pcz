// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"bytes"
	_ "embed" // for go:embed
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"sync"
	"text/template" // use text/template instead of html template for easy javascript templating
	"time"
	"unsafe"

	"github.com/fsnotify/fsnotify"
	"github.com/primecitizens/pcz/cmd/dev/internal"
	"github.com/primecitizens/pcz/cmd/dev/internal/js"
	"github.com/primecitizens/pcz/cmd/utils/tmplutils"
)

type resourceUpdate struct {
	Kind string
	Data []byte
}

type serveResouces struct {
	html     []byte
	wasm     []byte
	bindings []byte

	mu sync.Mutex
}

func sendBindings(ctx *serveContext, data []byte) bool {
	return ctx.send("bindings", data)
}

func sendHTML(ctx *serveContext, data []byte) bool {
	return ctx.send("html", data)
}

func sendWASM(ctx *serveContext, data []byte) bool {
	return ctx.send("wasm", data)
}

func (r *serveResouces) Update(update resourceUpdate) {
	r.mu.Lock()
	defer r.mu.Unlock()

	switch update.Kind {
	case "html":
		r.html = update.Data
	case "wasm":
		r.wasm = update.Data
	case "bindings":
		r.bindings = update.Data
	default:
		panic("unreachable")
	}
}

func (r *serveResouces) HTML() []byte {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.html
}

func (r *serveResouces) WASM() []byte {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.wasm
}

func (r *serveResouces) Bindings() []byte {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.bindings
}

type serveContext struct {
	stop   <-chan struct{}
	update chan<- resourceUpdate

	stderr io.Writer
}

func (ctx *serveContext) Stderr(format string, args ...any) {
	fmt.Fprintf(ctx.stderr, format, args...)
}

func (ctx *serveContext) send(kind string, data []byte) bool {
	select {
	case <-ctx.stop:
		return false
	case ctx.update <- resourceUpdate{
		Kind: kind,
		Data: data,
	}:
		return true
	}
}

func (ctx *serveContext) watchCustom(w *fsnotify.Watcher, file string, update func(ctx *serveContext, data []byte) bool) (err error) {
	file, err = filepath.Abs(file)
	if err != nil {
		return fmt.Errorf("cannot get absolute path to custom file %s", file)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read custom file %q: %w", file, err)
	}

	if !update(ctx, data) {
		// stopped
		return nil
	}

	err = w.Add(file)
	if err != nil {
		return fmt.Errorf("failed to watch custom file %s", file)
	}

	for {
		select {
		case <-ctx.stop:
			return nil
		case ev := <-w.Events:
			if ev.Op == fsnotify.Remove {
				continue
			}

			data, err := os.ReadFile(file)
			if err != nil {
				ctx.Stderr("failed to read custom file %s on update", file)
			}

			if !update(ctx, data) {
				return nil
			}
		}
	}
}

func (ctx *serveContext) routineWASM(customWASM string, dw *internal.DepWatcher, cacheDir string, spec js.BindgenOptions) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		panic("failed to create fs watcher for wasm blob")
	}
	defer w.Close()

	if len(customWASM) != 0 {
		if err = ctx.watchCustom(w, customWASM, sendWASM); err != nil {
			panic(err)
		}
		return
	}

	err = dw.AddWatcher(w)
	if err != nil {
		panic(fmt.Errorf("failed to add watcher for wasm for the first time: %w", err))
	}

	err = os.MkdirAll(cacheDir, 0750)
	if err != nil && !os.IsExist(err) {
		panic(fmt.Errorf("failed to create cache-dir: %w", err))
	}

	blobFile := filepath.Join(cacheDir, "dev.wasm")
	bindingsFile := filepath.Join(cacheDir, "bindings.js")

	w.Add(blobFile)
	w.Add(bindingsFile)

	bindingsTsk := internal.NewSingleTask(time.Second, func() {
		ctx.Stderr("updating bindings...\n")
		data, err := os.ReadFile(blobFile)
		if err != nil {
			ctx.Stderr("failed to read the built wasm for bindings: %v\n", err)
			return
		}

		filter, err := js.CreateSimpleFilterFromWasm(data)
		if err != nil {
			ctx.Stderr("failed to collect imported functions from the built wasm: %v\n", err)
			return
		}

		code, err := js.CreateJSBindings(js.BindgenOptions{
			Mode:          spec.Mode,
			ES:            spec.ES,
			ModuleName:    spec.ModuleName,
			CustomWrapper: spec.CustomWrapper,
			Deps:          dw.CurrentDeps(),
			Filter:        filter,
		})
		if err != nil {
			ctx.Stderr("failed to generate bindings: %v\n", err)
			return
		}

		data = unsafe.Slice(unsafe.StringData(code), len(code))
		_ = os.WriteFile(bindingsFile, data, 0640)
		sendBindings(ctx, data)
	})

	wasmTsk := internal.NewSingleTask(time.Second, func() {
		ctx.Stderr("building application...\n")
		err = dw.Build(blobFile)
		if err != nil {
			ctx.Stderr("failed to build application: %v\n", err)
			return
		}

		data, err := os.ReadFile(blobFile)
		if err != nil {
			ctx.Stderr("failed to read the built wasm: %v\n", err)
			return
		}

		bindingsTsk.Run()

		sendWASM(ctx, data)
	})

	wasmTsk.Run()
	for {
		select {
		case <-ctx.stop:
			return
		case ev := <-w.Events:
			switch ev.Op {
			case fsnotify.Remove,
				fsnotify.Write:
			default:
				continue
			}

			pkg := dw.Contains(filepath.Dir(ev.Name))
			if pkg == nil {
				switch ev.Name {
				case blobFile, bindingsFile:
					data, err := os.ReadFile(ev.Name)
					if err != nil {
						ctx.Stderr("cannot read updated cache (%s): %v\n", ev.Name, err)
						break
					}

					if ev.Name == blobFile {
						sendWASM(ctx, data)
					} else {
						sendBindings(ctx, data)
					}
				}

				continue
			}

			switch file := filepath.Base(ev.Name); {
			case file == "ffi_bindings.ts":
				go bindingsTsk.Run()
				continue
			case slices.Contains(pkg.GoFiles, file):
				// TODO: diff package deps in dw?
				// may have changed import statements
				dw.Refresh()
			case slices.Contains(pkg.SFiles, file),
				slices.Contains(pkg.EmbedFiles, file):
			default:
				continue
			}

			go wasmTsk.Run()
		case <-w.Errors:
		}
	}
}

func (ctx *serveContext) routineCustomBindings(file string) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		panic("failed to create fs watcher for js bindings")
	}
	defer w.Close()

	if len(file) != 0 {
		if err = ctx.watchCustom(w, file, sendBindings); err != nil {
			panic(err)
		}
		return
	}
}

func (ctx *serveContext) routineHTML(custom string, tdCh <-chan *TemplateData, cacheDir string) {
	htmlFile := filepath.Join(cacheDir, "index.html")
	err := os.MkdirAll(cacheDir, 0750)
	if err != nil && !os.IsExist(err) {
		panic(fmt.Errorf("failed to create cache-dir: %w", err))
	}

	writeHTMLFile := func(data []byte) {
		os.WriteFile(htmlFile, data, 0640)
	}

	if len(custom) == 0 {
	Loop:
		for {
			select {
			case <-ctx.stop:
				break Loop
			case td := <-tdCh:
				data, err := execTemplate(defaultHTML, td)
				if err != nil {
					// built-in template, can not change
					panic(err)
				}

				writeHTMLFile(data)
				sendHTML(ctx, data)
			}
		}

		return
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		panic("failed to create fs watcher for html page")
	}
	defer w.Close()

	tmplCh := make(chan string, 1)
	go func() {
		var (
			tmpl string
			td   *TemplateData
		)

		for {
			select {
			case <-ctx.stop:
				return
			case tmpl = <-tmplCh:
				if td == nil {
					continue
				}
			case td = <-tdCh:
				if len(tmpl) == 0 {
					continue
				}
			}

			data, err := execTemplate(tmpl, td)
			if err != nil {
				ctx.Stderr("%v\n", err)
				continue
			}

			writeHTMLFile(data)
			sendHTML(ctx, data)
		}
	}()

	if err = ctx.watchCustom(w, custom, func(ctx *serveContext, data []byte) bool {
		tmplCh <- unsafe.String(unsafe.SliceData(data), len(data))
		return true
	}); err != nil {
		panic(err)
	}
}

//go:embed serve.html
var defaultHTML string

type TemplateData struct {
	OS         string
	ModuleName string
	ManualRun  bool

	Args    []string
	Environ []string
}

func execTemplate(src string, td *TemplateData) ([]byte, error) {
	tmpl, err := template.New("").Parse(tmplutils.CleanupSourceTemplate(src))
	if err != nil {
		return nil, fmt.Errorf("html template contains error: %w", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, td)
	if err != nil {
		return nil, fmt.Errorf("failed to execute html template: %w", err)
	}

	return buf.Bytes(), nil
}
