// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package internal

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/primecitizens/pcz/cmd/build/toolchain"
)

func NewDepWatch(tc *toolchain.Toolchain, initPkgs []*toolchain.Package, entryPkg string) *DepWatcher {
	dw := &DepWatcher{
		tc:       tc,
		entryPkg: entryPkg,
		dirs:     map[string]int{},
	}

	_ = dw.updateLocked(initPkgs)
	return dw
}

type DepWatcher struct {
	tc       *toolchain.Toolchain
	entryPkg string

	wts []*fsnotify.Watcher

	// key: local dir path
	// value: index into deps + 1
	dirs map[string]int
	deps []*toolchain.Package

	lastRefresh time.Time
	scheduled   bool
	mu          sync.Mutex
}

func (w *DepWatcher) CurrentDeps() []*toolchain.Package {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.deps
}

func (w *DepWatcher) Contains(dir string) *toolchain.Package {
	w.mu.Lock()
	defer w.mu.Unlock()

	if i := w.dirs[dir]; i != 0 {
		return w.deps[i-1]
	}

	return nil
}

func (w *DepWatcher) Build(output string) (err error) {
	defer func() {
		panic := recover()
		if panic != nil {
			err = panic.(error)
		}
	}()

	w.mu.Lock()
	pkgs := w.deps
	w.mu.Unlock()

	return w.tc.Build(pkgs, output)
}

func (w *DepWatcher) AddWatcher(wt *fsnotify.Watcher) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	for dir := range w.dirs {
		if err := wt.Add(dir); err != nil {
			return err
		}
	}

	w.wts = append(w.wts, wt)
	return nil
}

func (w *DepWatcher) Refresh() (err error) {
	w.mu.Lock()
	defer func() {
		w.mu.Unlock()

		panic := recover()
		if panic != nil {
			err = panic.(error)
		}
	}()

	if time.Since(w.lastRefresh) < time.Second {
		if !w.scheduled {
			w.scheduled = true
			time.AfterFunc(time.Second, func() {
				w.Refresh()
			})
		}

		return nil
	}

	w.lastRefresh = time.Now()
	pkgs, err := w.tc.List(w.entryPkg)
	if err != nil {
		return fmt.Errorf("failed to list dependent go packages of %s", w.entryPkg)
	}

	errs := w.updateLocked(pkgs)
	if len(errs) == 0 {
		return nil
	}

	return errors.Join(errs...)
}

func (w *DepWatcher) updateLocked(pkgs []*toolchain.Package) []error {
	var errs []error
	current := make(map[string]int, len(pkgs))
	for i := 0; i < len(pkgs); i++ {
		dir := pkgs[i].Dir
		current[dir] = i + 1
		if w.dirs[dir] == 0 {
			for _, wt := range w.wts {
				if err := wt.Add(dir); err != nil {
					errs = append(errs, err)
				}
			}
		}
	}

	for dir := range w.dirs {
		if current[dir] == 0 {
			for _, wt := range w.wts {
				if err := wt.Remove(dir); err != nil {
					errs = append(errs, err)
				}
			}
		}
	}

	w.dirs = current
	w.deps = pkgs
	return errs
}
