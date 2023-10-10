// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webext

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/utils/fsutils"
)

var defaultPackMatch = []string{
	"./**/*.{html,css,js,json}",
	"./LICENSE*",
	"./README*",
}

var (
	packFlags = PackFlags{
		Match: defaultPackMatch,
	}
	cmdPack = cli.Cmd{
		Pattern:    "pack",
		BriefUsage: "Package web extension for distribution",
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &packFlags),
		Run:        packWebext,
	}
)

type PackFlags struct {
	Chdir  string   `cli:"c|chdir,#set work dir when packing"`
	Output string   `cli:"o|output,#set path to write the output zip file, if not set, write to stdout"`
	Match  []string `cli:"m|match,#add glob pattern to match expected files"`
}

func packWebext(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) error {
	var output io.Writer
	if dst := packFlags.Output; len(dst) != 0 {
		outputFile, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0640)
		if err != nil {
			return fmt.Errorf("failed to open %q for writting zip: %w", dst, err)
		}
		defer outputFile.Close()

		output = outputFile
	} else {
		output = opts.PickStdout(os.Stdout)
	}

	entries, err := Pack(output, packFlags.Chdir, packFlags.Match)
	if err != nil {
		return err
	}

	stderr := opts.PickStderr(os.Stderr)
	if len(entries) != 0 {
		fmt.Fprintf(stderr, "files in zip:\n")
		for _, f := range entries {
			fmt.Fprintf(stderr, "- { path: %q, size: %d }\n", f.Path, f.Size)
		}
	} else {
		fmt.Fprintf(stderr, "no files packed.\n")
	}

	return nil
}

func Pack(zipOut io.Writer, chdir string, globPatterns []string) (entries []ArchiveEntry, err error) {
	globs := globPatterns
	if len(chdir) != 0 {
		globs = make([]string, len(globPatterns))
		cwd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get current working dir: %w", err)
		}

		chdir1, err := filepath.Abs(chdir)
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute to %q", chdir, err)
		}

		base, err := filepath.Rel(cwd, chdir1)
		if err != nil {
			return nil, fmt.Errorf("failed to get relative path to %q: %w", chdir, err)
		}

		base = filepath.ToSlash(base)
		for i, ptn := range globPatterns {
			globs[i] = path.Join(base, ptn)
		}
	}

	src, err := fsutils.MatchFiles(globs...)
	if err != nil {
		return nil, err
	}

	files, err := collectFiles(src)
	if err != nil {
		return nil, err
	}

	return createZip(zipOut, files)
}

type entry struct {
	to   string
	from string
	info fs.FileInfo
}

// collectFiles to be archived
func collectFiles(src []string) (files []*entry, err error) {
	var (
		// swap first serves as the storage of all matched path info
		swap []*entry

		// key: in archive path, dir entry with trailing slash added
		// value: index into files
		inArchiveFiles = make(map[string]int)
	)

	slashMatches := make([]string, len(src))
	for i, sp := range src {
		if len(sp) == 0 {
			sp = "."
		}

		var info fs.FileInfo
		info, err = os.Lstat(sp)
		if err != nil {
			return
		}

		if info.Mode()&fs.ModeSymlink != 0 {
			err = fmt.Errorf("unsupported symlink file %q", sp)
			return
		}

		swap = append(swap, &entry{
			info: info,
			from: sp,
		})

		slashMatches[i] = filepath.ToSlash(sp)
		if info.IsDir() {
			slashMatches[i] += "/"
		}
	}

	szSwap := len(swap)

	nSlashMatches := len(slashMatches)
	switch nSlashMatches {
	case 0:
		err = fmt.Errorf("no file matched")
		return
	case 1:
		toAdd := swap[szSwap-1]
		// only one file match
		toAdd.to = filepath.Base(toAdd.from)

		inArchiveFiles[toAdd.to] = len(files)
		files = append(files, toAdd)
	default:
		prefix := lcpp(slashMatches)
		for i, sp := range slashMatches {
			if sp == prefix {
				continue
			}

			toAdd := swap[szSwap-nSlashMatches+i]
			toAdd.to = strings.TrimPrefix(sp, prefix)

			inArchiveFiles[toAdd.to] = len(files)
			files = append(files, toAdd)
		}
	}

	// add missing directories

	// good is the set of in archive paths having a viable parent
	// key: in archive path
	good := make(map[string]struct{})
	nGood := -1 // known good in last iteration
	for nGood != len(good) {
		nGood = len(good)
		for k, idx := range inArchiveFiles {
			dir := path.Dir(k)
			if dir == "." || dir == "/" {
				good[k] = struct{}{}
				continue
			}

			// TODO: why this exists?
			_, ok := inArchiveFiles[dir]
			if ok {
				good[k] = struct{}{}
				continue
			}

			_, ok = inArchiveFiles[dir+"/"]
			if ok {
				good[k] = struct{}{}
				continue
			}

			// no parent dir, add a fake one based on
			// actual parent of the file
			from := filepath.Dir(files[idx].from)
			info, err := os.Lstat(from)
			if err != nil {
				return nil, err
			}

			ent := &entry{
				from: from,
				info: info,
				to:   dir,
			}

			inArchiveFiles[dir] = len(files)
			files = append(files, ent)
		}
	}

	slices.SortFunc(files, func(a, b *entry) int {
		return strings.Compare(a.to, b.to)
	})

	return files, nil
}

type ArchiveEntry struct {
	Path string
	Size int64
}

func createZip(out io.Writer, files []*entry) (entries []ArchiveEntry, err error) {
	zw := zip.NewWriter(out)
	defer zw.Close()

	var method uint16
	method = zip.Deflate
	zw.RegisterCompressor(method, func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w, flate.BestCompression)
	})

	var (
		hdr  *zip.FileHeader
		mode fs.FileMode

		entryWriter io.Writer
	)

	for _, f := range files {
		hdr, err = zip.FileInfoHeader(f.info)
		if err != nil {
			return
		}

		hdr.Name = f.to
		hdr.Method = method

		mode = f.info.Mode()
		if mode.IsDir() && !strings.HasSuffix(hdr.Name, "/") {
			hdr.Name += "/"
		}

		entryWriter, err = zw.CreateHeader(hdr)
		if err != nil {
			return
		}

		var sz int64
		switch {
		case mode&fs.ModeSymlink != 0:
			err = fmt.Errorf("unsupported symlink file")
			return
		case mode.IsRegular():
			sz, err = copyFileContent(entryWriter, f.from)
			if err != nil {
				return
			}
		}

		entries = append(entries, ArchiveEntry{
			Path: hdr.Name,
			Size: sz,
		})
	}

	return
}

func copyFileContent(w io.Writer, file string) (int64, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return io.Copy(w, f)
}

// lcpp returns the longest common path prefix among the slash paths.
func lcpp(paths []string) string {
	if len(paths) == 0 {
		return ""
	}

	min, max := paths[0], paths[0]
	for _, s := range paths[1:] {
		switch {
		case s < min:
			min = s
		case s > max:
			max = s
		}
	}

	lastSlashAt := -1
	for i := 0; i < len(min) && i < len(max); i++ {
		if min[i] != max[i] {
			break
		}

		if min[i] == '/' {
			lastSlashAt = i
		}
	}

	if lastSlashAt != -1 {
		return min[:lastSlashAt+1]
	}

	return ""
}
