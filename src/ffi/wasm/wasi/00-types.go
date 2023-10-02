// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package wasi

import (
	"unsafe"

	"github.com/primecitizens/std/io"
)

// An Errno is an unsigned number describing an error condition.
// The zero Errno is by convention a non-error.
type Errno uint32

func (e Errno) Temporary() bool {
	return e == EINTR || e == EMFILE || e.Timeout()
}

func (e Errno) Timeout() bool {
	return e == EAGAIN || e == ETIMEDOUT
}

const (
	E2BIG           Errno = 1
	EACCES          Errno = 2
	EADDRINUSE      Errno = 3
	EADDRNOTAVAIL   Errno = 4
	EAFNOSUPPORT    Errno = 5
	EAGAIN          Errno = 6
	EALREADY        Errno = 7
	EBADF           Errno = 8
	EBADMSG         Errno = 9
	EBUSY           Errno = 10
	ECANCELED       Errno = 11
	ECHILD          Errno = 12
	ECONNABORTED    Errno = 13
	ECONNREFUSED    Errno = 14
	ECONNRESET      Errno = 15
	EDEADLK         Errno = 16
	EDESTADDRREQ    Errno = 17
	EDOM            Errno = 18
	EDQUOT          Errno = 19
	EEXIST          Errno = 20
	EFAULT          Errno = 21
	EFBIG           Errno = 22
	EHOSTUNREACH    Errno = 23
	EIDRM           Errno = 24
	EILSEQ          Errno = 25
	EINPROGRESS     Errno = 26
	EINTR           Errno = 27
	EINVAL          Errno = 28
	EIO             Errno = 29
	EISCONN         Errno = 30
	EISDIR          Errno = 31
	ELOOP           Errno = 32
	EMFILE          Errno = 33
	EMLINK          Errno = 34
	EMSGSIZE        Errno = 35
	EMULTIHOP       Errno = 36
	ENAMETOOLONG    Errno = 37
	ENETDOWN        Errno = 38
	ENETRESET       Errno = 39
	ENETUNREACH     Errno = 40
	ENFILE          Errno = 41
	ENOBUFS         Errno = 42
	ENODEV          Errno = 43
	ENOENT          Errno = 44
	ENOEXEC         Errno = 45
	ENOLCK          Errno = 46
	ENOLINK         Errno = 47
	ENOMEM          Errno = 48
	ENOMSG          Errno = 49
	ENOPROTOOPT     Errno = 50
	ENOSPC          Errno = 51
	ENOSYS          Errno = 52
	ENOTCONN        Errno = 53
	ENOTDIR         Errno = 54
	ENOTEMPTY       Errno = 55
	ENOTRECOVERABLE Errno = 56
	ENOTSOCK        Errno = 57
	ENOTSUP         Errno = 58
	ENOTTY          Errno = 59
	ENXIO           Errno = 60
	EOVERFLOW       Errno = 61
	EOWNERDEAD      Errno = 62
	EPERM           Errno = 63
	EPIPE           Errno = 64
	EPROTO          Errno = 65
	EPROTONOSUPPORT Errno = 66
	EPROTOTYPE      Errno = 67
	ERANGE          Errno = 68
	EROFS           Errno = 69
	ESPIPE          Errno = 70
	ESRCH           Errno = 71
	ESTALE          Errno = 72
	ETIMEDOUT       Errno = 73
	ETXTBSY         Errno = 74
	EXDEV           Errno = 75
	ENOTCAPABLE     Errno = 76
)

func (errno Errno) IOErr() io.Status {
	if errno == 0 {
		return io.StatusOK
	}

	switch errno {
	case ENOTDIR:
		return io.StatusNotDir
	case ENOTSUP:
		return io.StatusUnsupported
	default:
		return io.StatusUnkown
	}
}

type (
	FD          int32
	uintptr32   = uint32
	Size        uint32
	FDflags     uint32
	Filesize    uint64
	Filetype    uint8
	LookupFlags uint32
	OFlags      uint32
	Rights      uint64
	Timestamp   uint64
	Dircookie   uint64
	FileDelta   int64
	FstFlags    uint32
)

func Uintptr(p unsafe.Pointer) uintptr32 {
	return uintptr32(uintptr(p))
}

type IOBuffer struct {
	Ptr uintptr32
	Len Size
}

const (
	LOOKUP_SYMLINK_FOLLOW LookupFlags = 0x00000001
)

const (
	OFlag_CREAT     OFlags = 0x0001 // Create file if it does not exist.
	OFlag_DIRECTORY OFlags = 0x0002 // Fail if not a directory.
	OFlag_EXCL      OFlags = 0x0004 // Fail if file already exists.
	OFlag_TRUNC     OFlags = 0x0008 // Truncate file to size 0.
)

const (
	FDflag_APPEND   FDflags = 0x0001 //
	FDflag_DSYNC    FDflags = 0x0002 //
	FDflag_NONBLOCK FDflags = 0x0004 //
	FDflag_RSYNC    FDflags = 0x0008 //
	FDflag_SYNC     FDflags = 0x0010 //
)

const (
	Right_FD_DATASYNC             Rights = 1 << iota //
	Right_FD_READ                                    //
	Right_FD_SEEK                                    //
	Right_FDSTAT_SET_FLAGS                           //
	Right_FD_SYNC                                    //
	Right_FD_TELL                                    //
	Right_FD_WRITE                                   //
	Right_FD_ADVISE                                  //
	Right_FD_ALLOCATE                                //
	Right_PATH_CREATE_DIRECTORY                      //
	Right_PATH_CREATE_FILE                           //
	Right_PATH_LINK_SOURCE                           //
	Right_PATH_LINK_TARGET                           //
	Right_PATH_OPEN                                  //
	Right_FD_READDIR                                 //
	Right_PATH_READLINK                              //
	Right_PATH_RENAME_SOURCE                         //
	Right_PATH_RENAME_TARGET                         //
	Right_PATH_FILESTAT_GET                          //
	Right_PATH_FILESTAT_SET_SIZE                     //
	Right_PATH_FILESTAT_SET_TIMES                    //
	Right_FD_FILESTAT_GET                            //
	Right_FD_FILESTAT_SET_SIZE                       //
	Right_FD_FILESTAT_SET_TIMES                      //
	Right_PATH_SYMLINK                               //
	Right_PATH_REMOVE_DIRECTORY                      //
	Right_PATH_UNLINK_FILE                           //
	Right_POLL_FD_READWRITE                          //
	Right_SOCK_SHUTDOWN                              //
	Right_SOCK_ACCEPT                                //
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
	FullRights  Rights = ^Rights(0)
	ReadRights  Rights = Right_FD_READ | Right_FD_READDIR
	WriteRights Rights = Right_FD_DATASYNC | Right_FD_WRITE | Right_FD_ALLOCATE | Right_PATH_FILESTAT_SET_SIZE

	// Some runtimes have very strict expectations when it comes to which
	// rights can be enabled on files opened by path_open. The fileRights
	// constant is used as a mask to retain only bits for operations that
	// are supported on files.
	FileRights Rights = 0 |
		Right_FD_DATASYNC |
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
		Right_POLL_FD_READWRITE |
		0

	// Runtimes like wasmtime and wasmedge will refuse to open directories
	// if the rights requested by the application exceed the operations that
	// can be performed on a directory.
	DirRights Rights = 0 |
		Right_FD_SEEK |
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
		Right_PATH_UNLINK_FILE |
		0
)

const (
	Filetype_UNKNOWN          Filetype = iota
	Filetype_BLOCK_DEVICE              //
	Filetype_CHARACTER_DEVICE          //
	Filetype_DIRECTORY                 //
	Filetype_REGULAR_FILE              //
	Filetype_SOCKET_DGRAM              //
	Filetype_SOCKET_STREAM             //
	Filetype_SYMBOLIC_LINK             //
)

const (
	Stdin  FD = 0
	Stdout FD = 1
	Stderr FD = 2
)

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
