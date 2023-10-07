// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package filesystemproviderinternal

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/filesystemprovider"
	"github.com/primecitizens/pcz/std/plat/js/webext/filesystemproviderinternal/bindings"
)

// HasFuncGetActionsRequestedSuccess returns true if the function "WEBEXT.fileSystemProviderInternal.getActionsRequestedSuccess" exists.
func HasFuncGetActionsRequestedSuccess() bool {
	return js.True == bindings.HasFuncGetActionsRequestedSuccess()
}

// FuncGetActionsRequestedSuccess returns the function "WEBEXT.fileSystemProviderInternal.getActionsRequestedSuccess".
func FuncGetActionsRequestedSuccess() (fn js.Func[func(fileSystemId js.String, requestId int32, actions js.Array[filesystemprovider.Action], executionTime int32)]) {
	bindings.FuncGetActionsRequestedSuccess(
		js.Pointer(&fn),
	)
	return
}

// GetActionsRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.getActionsRequestedSuccess" directly.
func GetActionsRequestedSuccess(fileSystemId js.String, requestId int32, actions js.Array[filesystemprovider.Action], executionTime int32) (ret js.Void) {
	bindings.CallGetActionsRequestedSuccess(
		js.Pointer(&ret),
		fileSystemId.Ref(),
		int32(requestId),
		actions.Ref(),
		int32(executionTime),
	)

	return
}

// TryGetActionsRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.getActionsRequestedSuccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetActionsRequestedSuccess(fileSystemId js.String, requestId int32, actions js.Array[filesystemprovider.Action], executionTime int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetActionsRequestedSuccess(
		js.Pointer(&ret), js.Pointer(&exception),
		fileSystemId.Ref(),
		int32(requestId),
		actions.Ref(),
		int32(executionTime),
	)

	return
}

// HasFuncGetMetadataRequestedSuccess returns true if the function "WEBEXT.fileSystemProviderInternal.getMetadataRequestedSuccess" exists.
func HasFuncGetMetadataRequestedSuccess() bool {
	return js.True == bindings.HasFuncGetMetadataRequestedSuccess()
}

// FuncGetMetadataRequestedSuccess returns the function "WEBEXT.fileSystemProviderInternal.getMetadataRequestedSuccess".
func FuncGetMetadataRequestedSuccess() (fn js.Func[func(fileSystemId js.String, requestId int32, metadata filesystemprovider.EntryMetadata, executionTime int32)]) {
	bindings.FuncGetMetadataRequestedSuccess(
		js.Pointer(&fn),
	)
	return
}

// GetMetadataRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.getMetadataRequestedSuccess" directly.
func GetMetadataRequestedSuccess(fileSystemId js.String, requestId int32, metadata filesystemprovider.EntryMetadata, executionTime int32) (ret js.Void) {
	bindings.CallGetMetadataRequestedSuccess(
		js.Pointer(&ret),
		fileSystemId.Ref(),
		int32(requestId),
		js.Pointer(&metadata),
		int32(executionTime),
	)

	return
}

// TryGetMetadataRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.getMetadataRequestedSuccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMetadataRequestedSuccess(fileSystemId js.String, requestId int32, metadata filesystemprovider.EntryMetadata, executionTime int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMetadataRequestedSuccess(
		js.Pointer(&ret), js.Pointer(&exception),
		fileSystemId.Ref(),
		int32(requestId),
		js.Pointer(&metadata),
		int32(executionTime),
	)

	return
}

// HasFuncOperationRequestedError returns true if the function "WEBEXT.fileSystemProviderInternal.operationRequestedError" exists.
func HasFuncOperationRequestedError() bool {
	return js.True == bindings.HasFuncOperationRequestedError()
}

// FuncOperationRequestedError returns the function "WEBEXT.fileSystemProviderInternal.operationRequestedError".
func FuncOperationRequestedError() (fn js.Func[func(fileSystemId js.String, requestId int32, err filesystemprovider.ProviderError, executionTime int32)]) {
	bindings.FuncOperationRequestedError(
		js.Pointer(&fn),
	)
	return
}

// OperationRequestedError calls the function "WEBEXT.fileSystemProviderInternal.operationRequestedError" directly.
func OperationRequestedError(fileSystemId js.String, requestId int32, err filesystemprovider.ProviderError, executionTime int32) (ret js.Void) {
	bindings.CallOperationRequestedError(
		js.Pointer(&ret),
		fileSystemId.Ref(),
		int32(requestId),
		uint32(err),
		int32(executionTime),
	)

	return
}

// TryOperationRequestedError calls the function "WEBEXT.fileSystemProviderInternal.operationRequestedError"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOperationRequestedError(fileSystemId js.String, requestId int32, err filesystemprovider.ProviderError, executionTime int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOperationRequestedError(
		js.Pointer(&ret), js.Pointer(&exception),
		fileSystemId.Ref(),
		int32(requestId),
		uint32(err),
		int32(executionTime),
	)

	return
}

// HasFuncOperationRequestedSuccess returns true if the function "WEBEXT.fileSystemProviderInternal.operationRequestedSuccess" exists.
func HasFuncOperationRequestedSuccess() bool {
	return js.True == bindings.HasFuncOperationRequestedSuccess()
}

// FuncOperationRequestedSuccess returns the function "WEBEXT.fileSystemProviderInternal.operationRequestedSuccess".
func FuncOperationRequestedSuccess() (fn js.Func[func(fileSystemId js.String, requestId int32, executionTime int32)]) {
	bindings.FuncOperationRequestedSuccess(
		js.Pointer(&fn),
	)
	return
}

// OperationRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.operationRequestedSuccess" directly.
func OperationRequestedSuccess(fileSystemId js.String, requestId int32, executionTime int32) (ret js.Void) {
	bindings.CallOperationRequestedSuccess(
		js.Pointer(&ret),
		fileSystemId.Ref(),
		int32(requestId),
		int32(executionTime),
	)

	return
}

// TryOperationRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.operationRequestedSuccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOperationRequestedSuccess(fileSystemId js.String, requestId int32, executionTime int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOperationRequestedSuccess(
		js.Pointer(&ret), js.Pointer(&exception),
		fileSystemId.Ref(),
		int32(requestId),
		int32(executionTime),
	)

	return
}

// HasFuncReadDirectoryRequestedSuccess returns true if the function "WEBEXT.fileSystemProviderInternal.readDirectoryRequestedSuccess" exists.
func HasFuncReadDirectoryRequestedSuccess() bool {
	return js.True == bindings.HasFuncReadDirectoryRequestedSuccess()
}

// FuncReadDirectoryRequestedSuccess returns the function "WEBEXT.fileSystemProviderInternal.readDirectoryRequestedSuccess".
func FuncReadDirectoryRequestedSuccess() (fn js.Func[func(fileSystemId js.String, requestId int32, entries js.Array[filesystemprovider.EntryMetadata], hasMore bool, executionTime int32)]) {
	bindings.FuncReadDirectoryRequestedSuccess(
		js.Pointer(&fn),
	)
	return
}

// ReadDirectoryRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.readDirectoryRequestedSuccess" directly.
func ReadDirectoryRequestedSuccess(fileSystemId js.String, requestId int32, entries js.Array[filesystemprovider.EntryMetadata], hasMore bool, executionTime int32) (ret js.Void) {
	bindings.CallReadDirectoryRequestedSuccess(
		js.Pointer(&ret),
		fileSystemId.Ref(),
		int32(requestId),
		entries.Ref(),
		js.Bool(bool(hasMore)),
		int32(executionTime),
	)

	return
}

// TryReadDirectoryRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.readDirectoryRequestedSuccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReadDirectoryRequestedSuccess(fileSystemId js.String, requestId int32, entries js.Array[filesystemprovider.EntryMetadata], hasMore bool, executionTime int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadDirectoryRequestedSuccess(
		js.Pointer(&ret), js.Pointer(&exception),
		fileSystemId.Ref(),
		int32(requestId),
		entries.Ref(),
		js.Bool(bool(hasMore)),
		int32(executionTime),
	)

	return
}

// HasFuncReadFileRequestedSuccess returns true if the function "WEBEXT.fileSystemProviderInternal.readFileRequestedSuccess" exists.
func HasFuncReadFileRequestedSuccess() bool {
	return js.True == bindings.HasFuncReadFileRequestedSuccess()
}

// FuncReadFileRequestedSuccess returns the function "WEBEXT.fileSystemProviderInternal.readFileRequestedSuccess".
func FuncReadFileRequestedSuccess() (fn js.Func[func(fileSystemId js.String, requestId int32, data js.ArrayBuffer, hasMore bool, executionTime int32)]) {
	bindings.FuncReadFileRequestedSuccess(
		js.Pointer(&fn),
	)
	return
}

// ReadFileRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.readFileRequestedSuccess" directly.
func ReadFileRequestedSuccess(fileSystemId js.String, requestId int32, data js.ArrayBuffer, hasMore bool, executionTime int32) (ret js.Void) {
	bindings.CallReadFileRequestedSuccess(
		js.Pointer(&ret),
		fileSystemId.Ref(),
		int32(requestId),
		data.Ref(),
		js.Bool(bool(hasMore)),
		int32(executionTime),
	)

	return
}

// TryReadFileRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.readFileRequestedSuccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReadFileRequestedSuccess(fileSystemId js.String, requestId int32, data js.ArrayBuffer, hasMore bool, executionTime int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReadFileRequestedSuccess(
		js.Pointer(&ret), js.Pointer(&exception),
		fileSystemId.Ref(),
		int32(requestId),
		data.Ref(),
		js.Bool(bool(hasMore)),
		int32(executionTime),
	)

	return
}

// HasFuncRespondToMountRequest returns true if the function "WEBEXT.fileSystemProviderInternal.respondToMountRequest" exists.
func HasFuncRespondToMountRequest() bool {
	return js.True == bindings.HasFuncRespondToMountRequest()
}

// FuncRespondToMountRequest returns the function "WEBEXT.fileSystemProviderInternal.respondToMountRequest".
func FuncRespondToMountRequest() (fn js.Func[func(requestId int32, err filesystemprovider.ProviderError, executionTime int32)]) {
	bindings.FuncRespondToMountRequest(
		js.Pointer(&fn),
	)
	return
}

// RespondToMountRequest calls the function "WEBEXT.fileSystemProviderInternal.respondToMountRequest" directly.
func RespondToMountRequest(requestId int32, err filesystemprovider.ProviderError, executionTime int32) (ret js.Void) {
	bindings.CallRespondToMountRequest(
		js.Pointer(&ret),
		int32(requestId),
		uint32(err),
		int32(executionTime),
	)

	return
}

// TryRespondToMountRequest calls the function "WEBEXT.fileSystemProviderInternal.respondToMountRequest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRespondToMountRequest(requestId int32, err filesystemprovider.ProviderError, executionTime int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRespondToMountRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(requestId),
		uint32(err),
		int32(executionTime),
	)

	return
}

// HasFuncUnmountRequestedSuccess returns true if the function "WEBEXT.fileSystemProviderInternal.unmountRequestedSuccess" exists.
func HasFuncUnmountRequestedSuccess() bool {
	return js.True == bindings.HasFuncUnmountRequestedSuccess()
}

// FuncUnmountRequestedSuccess returns the function "WEBEXT.fileSystemProviderInternal.unmountRequestedSuccess".
func FuncUnmountRequestedSuccess() (fn js.Func[func(fileSystemId js.String, requestId int32, executionTime int32)]) {
	bindings.FuncUnmountRequestedSuccess(
		js.Pointer(&fn),
	)
	return
}

// UnmountRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.unmountRequestedSuccess" directly.
func UnmountRequestedSuccess(fileSystemId js.String, requestId int32, executionTime int32) (ret js.Void) {
	bindings.CallUnmountRequestedSuccess(
		js.Pointer(&ret),
		fileSystemId.Ref(),
		int32(requestId),
		int32(executionTime),
	)

	return
}

// TryUnmountRequestedSuccess calls the function "WEBEXT.fileSystemProviderInternal.unmountRequestedSuccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnmountRequestedSuccess(fileSystemId js.String, requestId int32, executionTime int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnmountRequestedSuccess(
		js.Pointer(&ret), js.Pointer(&exception),
		fileSystemId.Ref(),
		int32(requestId),
		int32(executionTime),
	)

	return
}
