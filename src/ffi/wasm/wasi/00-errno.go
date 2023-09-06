// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package wasi

// An Errno is an unsigned number describing an error condition.
// The zero Errno is by convention a non-error.
type Errno uint32

func (e Errno) Temporary() bool {
	return e == EINTR || e == EMFILE || e.Timeout()
}

func (e Errno) Timeout() bool {
	return e == EAGAIN || e == ETIMEDOUT
}

// TODO: Auto-generate some day. (Hard-coded in binaries so not likely to change.)
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
	// needed by src/net/error_unix_test.go
	EOPNOTSUPP = ENOTSUP
)
