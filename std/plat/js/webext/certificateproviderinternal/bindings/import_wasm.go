// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/webext/certificateproviderinternal has_ReportCertificates
//go:noescape
func HasFuncReportCertificates() js.Ref

//go:wasmimport plat/js/webext/certificateproviderinternal func_ReportCertificates
//go:noescape
func FuncReportCertificates(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateproviderinternal call_ReportCertificates
//go:noescape
func CallReportCertificates(
	retPtr unsafe.Pointer,
	requestId int32,
	certificates js.Ref)

//go:wasmimport plat/js/webext/certificateproviderinternal try_ReportCertificates
//go:noescape
func TryReportCertificates(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	certificates js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateproviderinternal has_ReportSignature
//go:noescape
func HasFuncReportSignature() js.Ref

//go:wasmimport plat/js/webext/certificateproviderinternal func_ReportSignature
//go:noescape
func FuncReportSignature(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateproviderinternal call_ReportSignature
//go:noescape
func CallReportSignature(
	retPtr unsafe.Pointer,
	requestId int32,
	signature js.Ref)

//go:wasmimport plat/js/webext/certificateproviderinternal try_ReportSignature
//go:noescape
func TryReportSignature(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	signature js.Ref) (ok js.Ref)
