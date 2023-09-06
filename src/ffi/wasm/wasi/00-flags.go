// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package wasi

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
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

const (
	O_RDONLY = 0
	O_WRONLY = 1
	O_RDWR   = 2

	O_CREAT  = 0100
	O_CREATE = O_CREAT
	O_TRUNC  = 01000
	O_APPEND = 02000
	O_EXCL   = 0200
	O_SYNC   = 010000

	O_CLOEXEC = 0
)

const (
	F_DUPFD   = 0
	F_GETFD   = 1
	F_SETFD   = 2
	F_GETFL   = 3
	F_SETFL   = 4
	F_GETOWN  = 5
	F_SETOWN  = 6
	F_GETLK   = 7
	F_SETLK   = 8
	F_SETLKW  = 9
	F_RGETLK  = 10
	F_RSETLK  = 11
	F_CNVT    = 12
	F_RSETLKW = 13

	F_RDLCK   = 1
	F_WRLCK   = 2
	F_UNLCK   = 3
	F_UNLKSYS = 4
)

const (
	S_IFMT        = 0000370000
	S_IFSHM_SYSV  = 0000300000
	S_IFSEMA      = 0000270000
	S_IFCOND      = 0000260000
	S_IFMUTEX     = 0000250000
	S_IFSHM       = 0000240000
	S_IFBOUNDSOCK = 0000230000
	S_IFSOCKADDR  = 0000220000
	S_IFDSOCK     = 0000210000

	S_IFSOCK = 0000140000
	S_IFLNK  = 0000120000
	S_IFREG  = 0000100000
	S_IFBLK  = 0000060000
	S_IFDIR  = 0000040000
	S_IFCHR  = 0000020000
	S_IFIFO  = 0000010000

	S_UNSUP = 0000370000

	S_ISUID = 0004000
	S_ISGID = 0002000
	S_ISVTX = 0001000

	S_IREAD  = 0400
	S_IWRITE = 0200
	S_IEXEC  = 0100

	S_IRWXU = 0700
	S_IRUSR = 0400
	S_IWUSR = 0200
	S_IXUSR = 0100

	S_IRWXG = 070
	S_IRGRP = 040
	S_IWGRP = 020
	S_IXGRP = 010

	S_IRWXO = 07
	S_IROTH = 04
	S_IWOTH = 02
	S_IXOTH = 01
)
