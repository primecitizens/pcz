// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package toolchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"time"
)

func (tc *Toolchain) CreateCacheRecord(now time.Time, buildDir string, p *Package) (_ *CacheRecord, err error) {
	cache := CacheRecord{
		Tags: tc.opts.GetBuildTags(),
		Src: FileSet{
			Dir: p.Dir,
		},
		Build: FileSet{
			Dir: buildDir,
		},
	}
	for _, fileSet := range [][]string{
		p.GoFiles,
		p.SFiles,
		p.EmbedFiles,
	} {
		err = cache.Src.Add(now, p.Dir, fileSet...)
		if err != nil {
			panic(err)
		}
	}

	return &cache, nil
}

// TODO: generate user friendly path for build cache from tags.
// currently it rebuilds everything when tags changed.

type CacheRecord struct {
	Tags []string `yaml:"tags"`

	Src   FileSet `yaml:"src"`
	Build FileSet `yaml:"build"`
}

func (old *CacheRecord) SourceUnchanged(newc *CacheRecord) bool {
	if len(old.Tags) != len(newc.Tags) {
		return false
	}

	for _, tg := range old.Tags {
		if !slices.Contains(newc.Tags, tg) {
			return false
		}
	}

	for _, tg := range newc.Tags {
		if !slices.Contains(old.Tags, tg) {
			return false
		}
	}

	return old.Src.NoChanges(&newc.Src)
}

type FileSet struct {
	Dir   string  `yaml:"dir"`
	Files []*File `yaml:"files"`

	done bool
}

func (fset *FileSet) Find(name string) *File {
	i := slices.IndexFunc(fset.Files, func(f *File) bool {
		return f.Name == name
	})
	if i < 0 {
		return nil
	}

	return fset.Files[i]
}

func (fset *FileSet) NoChanges(newfset *FileSet) bool {
	if len(fset.Files) != len(newfset.Files) {
		return false
	}

	for _, old := range fset.Files {
		if !slices.ContainsFunc(newfset.Files, func(f *File) bool {
			return f.SHA256Hex == old.SHA256Hex
		}) {
			return false
		}
	}

	for _, newf := range newfset.Files {
		if !slices.ContainsFunc(fset.Files, func(f *File) bool {
			return f.SHA256Hex == newf.SHA256Hex
		}) {
			return false
		}
	}

	return true
}

func (fset *FileSet) Add(now time.Time, dir string, filenames ...string) (err error) {
	if fset.done {
		panic(fmt.Errorf("adding file after FileSet closed"))
	}

	if fset.Dir != dir {
		panic(fmt.Errorf("adding files in different dir: dir = %q, got = %q", fset.Dir, dir))
	}

	for _, name := range filenames {
		path := name
		if !filepath.IsAbs(name) {
			path = filepath.Join(dir, name)
		}

		if slices.ContainsFunc(fset.Files, func(f *File) bool {
			return name == f.Name
		}) {
			panic(fmt.Errorf("duplicate file: %q", path))
		}

		info, err := os.Stat(path)
		if err != nil {
			panic(err)
		}

		mtime := info.ModTime()
		if mtime.IsZero() || mtime.Unix() == 0 {
			// mtime unknown
			mtime = now
		} else if mtime.After(now) {
			// bad mtime
			panic(fmt.Errorf("unexpected mtime after current time: %q", path))
		}

		sz := info.Size()

		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		var (
			n int64
			h = sha256.New()
		)
		func() {
			defer f.Close()

			n, err = io.Copy(h, f)
			if err != nil {
				panic(err)
			}
		}()

		if sz != n {
			panic(fmt.Errorf("unexpected copied size: want %d, got %d", sz, n))
		}

		fset.Files = append(fset.Files, &File{
			Name:      name,
			SHA256Hex: hex.EncodeToString(h.Sum(nil)),
			ModTime:   mtime.Unix(),
		})
	}

	return nil
}

// func (fset *FileSet) Finish() {
// 	if fset.done {
// 		panic(fmt.Errorf("already finished"))
// 	}
//
// 	var (
// 		h     = sha256.New()
// 		mtime int64
// 	)
// 	for _, f := range fset.Files {
// 		_, err := io.WriteString(h, f.SHA256Hex)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		mtime = max(mtime, f.ModTime)
// 	}
//
// 	fset.SHA256Hex = hex.EncodeToString(h.Sum(nil))
// 	fset.ModTime = mtime
// 	fset.done = true
// 	return
// }

type File struct {
	Name string `yaml:"name"`

	// SHA256Hex
	SHA256Hex string `yaml:"sha256"`

	ModTime int64 `yaml:"mtime"`
}

func (f *File) VerifyUnchanged(dir string) bool {
	if f == nil {
		return false
	}

	x, err := os.Open(filepath.Join(dir, f.Name))
	if err != nil {
		return false
	}
	defer x.Close()

	h := sha256.New()
	_, err = io.Copy(h, x)
	if err != nil {
		return false
	}

	return hex.EncodeToString(h.Sum(nil)) == f.SHA256Hex
}
