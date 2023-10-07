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

//go:wasmimport plat/js/webext/filesystemproviderinternal has_GetActionsRequestedSuccess
//go:noescape
func HasFuncGetActionsRequestedSuccess() js.Ref

//go:wasmimport plat/js/webext/filesystemproviderinternal func_GetActionsRequestedSuccess
//go:noescape
func FuncGetActionsRequestedSuccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemproviderinternal call_GetActionsRequestedSuccess
//go:noescape
func CallGetActionsRequestedSuccess(
	retPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	actions js.Ref,
	executionTime int32)

//go:wasmimport plat/js/webext/filesystemproviderinternal try_GetActionsRequestedSuccess
//go:noescape
func TryGetActionsRequestedSuccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	actions js.Ref,
	executionTime int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemproviderinternal has_GetMetadataRequestedSuccess
//go:noescape
func HasFuncGetMetadataRequestedSuccess() js.Ref

//go:wasmimport plat/js/webext/filesystemproviderinternal func_GetMetadataRequestedSuccess
//go:noescape
func FuncGetMetadataRequestedSuccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemproviderinternal call_GetMetadataRequestedSuccess
//go:noescape
func CallGetMetadataRequestedSuccess(
	retPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	metadata unsafe.Pointer,
	executionTime int32)

//go:wasmimport plat/js/webext/filesystemproviderinternal try_GetMetadataRequestedSuccess
//go:noescape
func TryGetMetadataRequestedSuccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	metadata unsafe.Pointer,
	executionTime int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemproviderinternal has_OperationRequestedError
//go:noescape
func HasFuncOperationRequestedError() js.Ref

//go:wasmimport plat/js/webext/filesystemproviderinternal func_OperationRequestedError
//go:noescape
func FuncOperationRequestedError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemproviderinternal call_OperationRequestedError
//go:noescape
func CallOperationRequestedError(
	retPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	err uint32,
	executionTime int32)

//go:wasmimport plat/js/webext/filesystemproviderinternal try_OperationRequestedError
//go:noescape
func TryOperationRequestedError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	err uint32,
	executionTime int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemproviderinternal has_OperationRequestedSuccess
//go:noescape
func HasFuncOperationRequestedSuccess() js.Ref

//go:wasmimport plat/js/webext/filesystemproviderinternal func_OperationRequestedSuccess
//go:noescape
func FuncOperationRequestedSuccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemproviderinternal call_OperationRequestedSuccess
//go:noescape
func CallOperationRequestedSuccess(
	retPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	executionTime int32)

//go:wasmimport plat/js/webext/filesystemproviderinternal try_OperationRequestedSuccess
//go:noescape
func TryOperationRequestedSuccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	executionTime int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemproviderinternal has_ReadDirectoryRequestedSuccess
//go:noescape
func HasFuncReadDirectoryRequestedSuccess() js.Ref

//go:wasmimport plat/js/webext/filesystemproviderinternal func_ReadDirectoryRequestedSuccess
//go:noescape
func FuncReadDirectoryRequestedSuccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemproviderinternal call_ReadDirectoryRequestedSuccess
//go:noescape
func CallReadDirectoryRequestedSuccess(
	retPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	entries js.Ref,
	hasMore js.Ref,
	executionTime int32)

//go:wasmimport plat/js/webext/filesystemproviderinternal try_ReadDirectoryRequestedSuccess
//go:noescape
func TryReadDirectoryRequestedSuccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	entries js.Ref,
	hasMore js.Ref,
	executionTime int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemproviderinternal has_ReadFileRequestedSuccess
//go:noescape
func HasFuncReadFileRequestedSuccess() js.Ref

//go:wasmimport plat/js/webext/filesystemproviderinternal func_ReadFileRequestedSuccess
//go:noescape
func FuncReadFileRequestedSuccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemproviderinternal call_ReadFileRequestedSuccess
//go:noescape
func CallReadFileRequestedSuccess(
	retPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	data js.Ref,
	hasMore js.Ref,
	executionTime int32)

//go:wasmimport plat/js/webext/filesystemproviderinternal try_ReadFileRequestedSuccess
//go:noescape
func TryReadFileRequestedSuccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	data js.Ref,
	hasMore js.Ref,
	executionTime int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemproviderinternal has_RespondToMountRequest
//go:noescape
func HasFuncRespondToMountRequest() js.Ref

//go:wasmimport plat/js/webext/filesystemproviderinternal func_RespondToMountRequest
//go:noescape
func FuncRespondToMountRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemproviderinternal call_RespondToMountRequest
//go:noescape
func CallRespondToMountRequest(
	retPtr unsafe.Pointer,
	requestId int32,
	err uint32,
	executionTime int32)

//go:wasmimport plat/js/webext/filesystemproviderinternal try_RespondToMountRequest
//go:noescape
func TryRespondToMountRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestId int32,
	err uint32,
	executionTime int32) (ok js.Ref)

//go:wasmimport plat/js/webext/filesystemproviderinternal has_UnmountRequestedSuccess
//go:noescape
func HasFuncUnmountRequestedSuccess() js.Ref

//go:wasmimport plat/js/webext/filesystemproviderinternal func_UnmountRequestedSuccess
//go:noescape
func FuncUnmountRequestedSuccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/filesystemproviderinternal call_UnmountRequestedSuccess
//go:noescape
func CallUnmountRequestedSuccess(
	retPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	executionTime int32)

//go:wasmimport plat/js/webext/filesystemproviderinternal try_UnmountRequestedSuccess
//go:noescape
func TryUnmountRequestedSuccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fileSystemId js.Ref,
	requestId int32,
	executionTime int32) (ok js.Ref)
