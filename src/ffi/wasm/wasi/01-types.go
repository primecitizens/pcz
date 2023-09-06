// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package wasi

type FD int32
type uintptr32 = uint32
type Size uint32
type FDflag uint32
type FileSize uint64
type Filetype uint8
type LookupFlags uint32
type OpenFlags uint32
type Rights uint64
type Timestamp uint64
type DirCookie uint64
type FileDelta int64
type FstFlags uint32

type IOBuffer struct {
	Ptr uintptr32
	Len Size
}

const (
	LOOKUP_SYMLINK_FOLLOW = 0x00000001
)

const (
	OpenFlag_CREATE    = 0x0001
	OpenFlag_DIRECTORY = 0x0002
	OpenFlag_EXCL      = 0x0004
	OpenFlag_TRUNC     = 0x0008
)

const (
	FDflag_APPEND   FDflag = 0x0001
	FDflag_DSYNC    FDflag = 0x0002
	FDflag_NONBLOCK FDflag = 0x0004
	FDflag_RSYNC    FDflag = 0x0008
	FDflag_SYNC     FDflag = 0x0010
)

const (
	Right_FD_DATASYNC = 1 << iota
	Right_FD_READ
	Right_FD_SEEK
	Right_FDSTAT_SET_FLAGS
	Right_FD_SYNC
	Right_FD_TELL
	Right_FD_WRITE
	Right_FD_ADVISE
	Right_FD_ALLOCATE
	Right_PATH_CREATE_DIRECTORY
	Right_PATH_CREATE_FILE
	Right_PATH_LINK_SOURCE
	Right_PATH_LINK_TARGET
	Right_PATH_OPEN
	Right_FD_READDIR
	Right_PATH_READLINK
	Right_PATH_RENAME_SOURCE
	Right_PATH_RENAME_TARGET
	Right_PATH_FILESTAT_GET
	Right_PATH_FILESTAT_SET_SIZE
	Right_PATH_FILESTAT_SET_TIMES
	Right_FD_FILESTAT_GET
	Right_FD_FILESTAT_SET_SIZE
	Right_FD_FILESTAT_SET_TIMES
	Right_PATH_SYMLINK
	Right_PATH_REMOVE_DIRECTORY
	Right_PATH_UNLINK_FILE
	Right_POLL_FD_READWRITE
	Right_SOCK_SHUTDOWN
	Right_SOCK_ACCEPT
)

const (
	WHENCE_SET = 0
	WHENCE_CUR = 1
	WHENCE_END = 2
)

const (
	FILESTAT_SET_ATIM     = 0x0001
	FILESTAT_SET_ATIM_NOW = 0x0002
	FILESTAT_SET_MTIM     = 0x0004
	FILESTAT_SET_MTIM_NOW = 0x0008
)

const (
	// Despite the rights being defined as a 64 bits integer in the spec,
	// wasmtime crashes the program if we set any of the upper 32 bits.
	fullRights  Rights = ^Rights(0)
	readRights  Rights = Right_FD_READ | Right_FD_READDIR
	writeRights Rights = Right_FD_DATASYNC | Right_FD_WRITE | Right_FD_ALLOCATE | Right_PATH_FILESTAT_SET_SIZE

	// Some runtimes have very strict expectations when it comes to which
	// rights can be enabled on files opened by path_open. The fileRights
	// constant is used as a mask to retain only bits for operations that
	// are supported on files.
	fileRights Rights = Right_FD_DATASYNC |
		Right_FD_READ |
		Right_FD_SEEK |
		Right_FDSTAT_SET_FLAGS |
		Right_FD_SYNC |
		Right_FD_TELL |
		Right_FD_WRITE |
		Right_FD_ADVISE |
		Right_FD_ALLOCATE |
		Right_PATH_CREATE_DIRECTORY |
		Right_PATH_CREATE_FILE |
		Right_PATH_LINK_SOURCE |
		Right_PATH_LINK_TARGET |
		Right_PATH_OPEN |
		Right_FD_READDIR |
		Right_PATH_READLINK |
		Right_PATH_RENAME_SOURCE |
		Right_PATH_RENAME_TARGET |
		Right_PATH_FILESTAT_GET |
		Right_PATH_FILESTAT_SET_SIZE |
		Right_PATH_FILESTAT_SET_TIMES |
		Right_FD_FILESTAT_GET |
		Right_FD_FILESTAT_SET_SIZE |
		Right_FD_FILESTAT_SET_TIMES |
		Right_PATH_SYMLINK |
		Right_PATH_REMOVE_DIRECTORY |
		Right_PATH_UNLINK_FILE |
		Right_POLL_FD_READWRITE

	// Runtimes like wasmtime and wasmedge will refuse to open directories
	// if the rights requested by the application exceed the operations that
	// can be performed on a directory.
	dirRights Rights = Right_FD_SEEK |
		Right_FDSTAT_SET_FLAGS |
		Right_FD_SYNC |
		Right_PATH_CREATE_DIRECTORY |
		Right_PATH_CREATE_FILE |
		Right_PATH_LINK_SOURCE |
		Right_PATH_LINK_TARGET |
		Right_PATH_OPEN |
		Right_FD_READDIR |
		Right_PATH_READLINK |
		Right_PATH_RENAME_SOURCE |
		Right_PATH_RENAME_TARGET |
		Right_PATH_FILESTAT_GET |
		Right_PATH_FILESTAT_SET_SIZE |
		Right_PATH_FILESTAT_SET_TIMES |
		Right_FD_FILESTAT_GET |
		Right_FD_FILESTAT_SET_TIMES |
		Right_PATH_SYMLINK |
		Right_PATH_REMOVE_DIRECTORY |
		Right_PATH_UNLINK_FILE
)

// type opendir struct {
// 	fd   int32
// 	name string
// }
//
// // List of preopen directories that were exposed by the runtime. The first one
// // is assumed to the be root directory of the file system, and others are seen
// // as mount points at sub paths of the root.
// var preopens []opendir
//
// // Current working directory. We maintain this as a string and resolve paths in
// // the code because wasmtime does not allow relative path lookups outside of the
// // scope of a directory; a previous approach we tried consisted in maintaining
// // open a file descriptor to the current directory so we could perform relative
// // path lookups from that location, but it resulted in breaking path resolution
// // from the current directory to its parent.
// var cwd string
//
// func init() {
// 	dirNameBuf := make([]byte, 256)
// 	// We start looking for preopens at fd=3 because 0, 1, and 2 are reserved
// 	// for standard input and outputs.
// 	for preopenFd := int32(3); ; preopenFd++ {
// 		var prestat prestat
//
// 		errno := fd_prestat_get(preopenFd, unsafe.Pointer(&prestat))
// 		if errno == EBADF {
// 			break
// 		}
// 		if errno == ENOTDIR || prestat.typ != preopentypeDir {
// 			continue
// 		}
// 		if errno != 0 {
// 			panic("fd_prestat: " + errno.Error())
// 		}
// 		if int(prestat.dir.prNameLen) > len(dirNameBuf) {
// 			dirNameBuf = make([]byte, prestat.dir.prNameLen)
// 		}
//
// 		errno = fd_prestat_dir_name(preopenFd, unsafe.Pointer(&dirNameBuf[0]), prestat.dir.prNameLen)
// 		if errno != 0 {
// 			panic("fd_prestat_dir_name: " + errno.Error())
// 		}
//
// 		preopens = append(preopens, opendir{
// 			fd:   preopenFd,
// 			name: string(dirNameBuf[:prestat.dir.prNameLen]),
// 		})
// 	}
//
// 	if cwd, _ = Getenv("PWD"); cwd != "" {
// 		cwd = joinPath("/", cwd)
// 	} else if len(preopens) > 0 {
// 		cwd = preopens[0].name
// 	}
// }
//
// //go:nosplit
// func appendCleanPath(buf []byte, path string, lookupParent bool) ([]byte, bool) {
// 	i := 0
// 	for i < len(path) {
// 		for i < len(path) && path[i] == '/' {
// 			i++
// 		}
//
// 		j := i
// 		for j < len(path) && path[j] != '/' {
// 			j++
// 		}
//
// 		s := path[i:j]
// 		i = j
//
// 		switch s {
// 		case "":
// 			continue
// 		case ".":
// 			continue
// 		case "..":
// 			if !lookupParent {
// 				k := len(buf)
// 				for k > 0 && buf[k-1] != '/' {
// 					k--
// 				}
// 				for k > 1 && buf[k-1] == '/' {
// 					k--
// 				}
// 				buf = buf[:k]
// 				if k == 0 {
// 					lookupParent = true
// 				} else {
// 					s = ""
// 					continue
// 				}
// 			}
// 		default:
// 			lookupParent = false
// 		}
//
// 		if len(buf) > 0 && buf[len(buf)-1] != '/' {
// 			buf = append(buf, '/')
// 		}
// 		buf = append(buf, s...)
// 	}
// 	return buf, lookupParent
// }
//
// // joinPath concatenates dir and file paths, producing a cleaned path where
// // "." and ".." have been removed, unless dir is relative and the references
// // to parent directories in file represented a location relative to a parent
// // of dir.
// //
// // This function is used for path resolution of all wasi functions expecting
// // a path argument; the returned string is heap allocated, which we may want
// // to optimize in the future. Instead of returning a string, the function
// // could append the result to an output buffer that the functions in this
// // file can manage to have allocated on the stack (e.g. initializing to a
// // fixed capacity). Since it will significantly increase code complexity,
// // we prefer to optimize for readability and maintainability at this time.
// func joinPath(dir, file string) string {
// 	buf := make([]byte, 0, len(dir)+len(file)+1)
// 	if isAbs(dir) {
// 		buf = append(buf, '/')
// 	}
// 	buf, lookupParent := appendCleanPath(buf, dir, false)
// 	buf, _ = appendCleanPath(buf, file, lookupParent)
// 	// The appendCleanPath function cleans the path so it does not inject
// 	// references to the current directory. If both the dir and file args
// 	// were ".", this results in the output buffer being empty so we handle
// 	// this condition here.
// 	if len(buf) == 0 {
// 		buf = append(buf, '.')
// 	}
// 	// If the file ended with a '/' we make sure that the output also ends
// 	// with a '/'. This is needed to ensure that programs have a mechanism
// 	// to represent dereferencing symbolic links pointing to directories.
// 	if buf[len(buf)-1] != '/' && isDir(file) {
// 		buf = append(buf, '/')
// 	}
// 	return unsafe.String(&buf[0], len(buf))
// }
//
// func isAbs(path string) bool {
// 	return hasPrefix(path, "/")
// }
//
// func isDir(path string) bool {
// 	return hasSuffix(path, "/")
// }
//
// func hasPrefix(s, p string) bool {
// 	return len(s) >= len(p) && s[:len(p)] == p
// }
//
// func hasSuffix(s, x string) bool {
// 	return len(s) >= len(x) && s[len(s)-len(x):] == x
// }
//
// // preparePath returns the preopen file descriptor of the directory to perform
// // path resolution from, along with the pair of pointer and length for the
// // relative expression of path from the directory.
// //
// // If the path argument is not absolute, it is first appended to the current
// // working directory before resolution.
// func preparePath(path string) (int32, unsafe.Pointer, Size) {
// 	var dirFd = int32(-1)
// 	var dirName string
//
// 	dir := "/"
// 	if !isAbs(path) {
// 		dir = cwd
// 	}
// 	path = joinPath(dir, path)
//
// 	for _, p := range preopens {
// 		if len(p.name) > len(dirName) && hasPrefix(path, p.name) {
// 			dirFd, dirName = p.fd, p.name
// 		}
// 	}
//
// 	path = path[len(dirName):]
// 	for isAbs(path) {
// 		path = path[1:]
// 	}
// 	if len(path) == 0 {
// 		path = "."
// 	}
//
// 	return dirFd, stringPointer(path), Size(len(path))
// }
//
// func Chmod(path string, mode uint32) Errno {
// 	var stat Stat_t
// 	return Stat(path, mark.NoEscape(&stat))
// }
//
// func Fchmod(fd FD, mode uint32) Errno {
// 	var stat Stat_t
// 	return Fstat(fd, mark.NoEscape(&stat))
// }
//
// func Getwd() (string, error) {
// 	return cwd, nil
// }
//
// func Chdir(path string) Errno {
// 	if len(path) == 0 {
// 		return EINVAL
// 	}
//
// 	dir := "/"
// 	if !isAbs(path) {
// 		dir = cwd
// 	}
// 	path = joinPath(dir, path)
//
// 	var stat Stat_t
// 	dirFd, pathPtr, pathLen := preparePath(path)
// 	errno := path_filestat_get(dirFd, LOOKUP_SYMLINK_FOLLOW, pathPtr, pathLen, unsafe.Pointer(&stat))
// 	if errno != 0 {
// 		return errno
// 	}
// 	if stat.Filetype != FILETYPE_DIRECTORY {
// 		return ENOTDIR
// 	}
// 	cwd = path
// 	return 0
// }
//
