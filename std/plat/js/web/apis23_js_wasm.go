// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type IdleRequestCallbackFunc func(this js.Ref, deadline IdleDeadline) js.Ref

func (fn IdleRequestCallbackFunc) Register() js.Func[func(deadline IdleDeadline)] {
	return js.RegisterCallback[func(deadline IdleDeadline)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IdleRequestCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		IdleDeadline{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IdleRequestCallback[T any] struct {
	Fn  func(arg T, this js.Ref, deadline IdleDeadline) js.Ref
	Arg T
}

func (cb *IdleRequestCallback[T]) Register() js.Func[func(deadline IdleDeadline)] {
	return js.RegisterCallback[func(deadline IdleDeadline)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IdleRequestCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		IdleDeadline{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IdleDeadline struct {
	ref js.Ref
}

func (this IdleDeadline) Once() IdleDeadline {
	this.ref.Once()
	return this
}

func (this IdleDeadline) Ref() js.Ref {
	return this.ref
}

func (this IdleDeadline) FromRef(ref js.Ref) IdleDeadline {
	this.ref = ref
	return this
}

func (this IdleDeadline) Free() {
	this.ref.Free()
}

// DidTimeout returns the value of property "IdleDeadline.didTimeout".
//
// It returns ok=false if there is no such property.
func (this IdleDeadline) DidTimeout() (ret bool, ok bool) {
	ok = js.True == bindings.GetIdleDeadlineDidTimeout(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncTimeRemaining returns true if the method "IdleDeadline.timeRemaining" exists.
func (this IdleDeadline) HasFuncTimeRemaining() bool {
	return js.True == bindings.HasFuncIdleDeadlineTimeRemaining(
		this.ref,
	)
}

// FuncTimeRemaining returns the method "IdleDeadline.timeRemaining".
func (this IdleDeadline) FuncTimeRemaining() (fn js.Func[func() DOMHighResTimeStamp]) {
	bindings.FuncIdleDeadlineTimeRemaining(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TimeRemaining calls the method "IdleDeadline.timeRemaining".
func (this IdleDeadline) TimeRemaining() (ret DOMHighResTimeStamp) {
	bindings.CallIdleDeadlineTimeRemaining(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTimeRemaining calls the method "IdleDeadline.timeRemaining"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IdleDeadline) TryTimeRemaining() (ret DOMHighResTimeStamp, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdleDeadlineTimeRemaining(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type IdleRequestOptions struct {
	// Timeout is "IdleRequestOptions.timeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
	Timeout uint32

	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IdleRequestOptions with all fields set.
func (p IdleRequestOptions) FromRef(ref js.Ref) IdleRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IdleRequestOptions in the application heap.
func (p IdleRequestOptions) New() js.Ref {
	return bindings.IdleRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IdleRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.IdleRequestOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IdleRequestOptions) Update(ref js.Ref) {
	bindings.IdleRequestOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IdleRequestOptions) FreeMembers(recursive bool) {
}

type WriteCommandType uint32

const (
	_ WriteCommandType = iota

	WriteCommandType_WRITE
	WriteCommandType_SEEK
	WriteCommandType_TRUNCATE
)

func (WriteCommandType) FromRef(str js.Ref) WriteCommandType {
	return WriteCommandType(bindings.ConstOfWriteCommandType(str))
}

func (x WriteCommandType) String() (string, bool) {
	switch x {
	case WriteCommandType_WRITE:
		return "write", true
	case WriteCommandType_SEEK:
		return "seek", true
	case WriteCommandType_TRUNCATE:
		return "truncate", true
	default:
		return "", false
	}
}

type WriteParams struct {
	// Type is "WriteParams.type"
	//
	// Required
	Type WriteCommandType
	// Size is "WriteParams.size"
	//
	// Optional
	//
	// NOTE: FFI_USE_Size MUST be set to true to make this field effective.
	Size uint64
	// Position is "WriteParams.position"
	//
	// Optional
	//
	// NOTE: FFI_USE_Position MUST be set to true to make this field effective.
	Position uint64
	// Data is "WriteParams.data"
	//
	// Optional
	Data OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String

	FFI_USE_Size     bool // for Size.
	FFI_USE_Position bool // for Position.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WriteParams with all fields set.
func (p WriteParams) FromRef(ref js.Ref) WriteParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WriteParams in the application heap.
func (p WriteParams) New() js.Ref {
	return bindings.WriteParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WriteParams) UpdateFrom(ref js.Ref) {
	bindings.WriteParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WriteParams) Update(ref js.Ref) {
	bindings.WriteParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WriteParams) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) FromRef(ref js.Ref) OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams {
	return OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) Blob() Blob {
	return Blob{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams) WriteParams() WriteParams {
	var ret WriteParams
	ret.UpdateFrom(x.ref)
	return ret
}

type FileSystemWriteChunkType = OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_Blob_String_WriteParams

func NewFileSystemWritableFileStream(underlyingSink js.Object, strategy QueuingStrategy) (ret FileSystemWritableFileStream) {
	ret.ref = bindings.NewFileSystemWritableFileStreamByFileSystemWritableFileStream(
		underlyingSink.Ref(),
		js.Pointer(&strategy))
	return
}

func NewFileSystemWritableFileStreamByFileSystemWritableFileStream1(underlyingSink js.Object) (ret FileSystemWritableFileStream) {
	ret.ref = bindings.NewFileSystemWritableFileStreamByFileSystemWritableFileStream1(
		underlyingSink.Ref())
	return
}

func NewFileSystemWritableFileStreamByFileSystemWritableFileStream2() (ret FileSystemWritableFileStream) {
	ret.ref = bindings.NewFileSystemWritableFileStreamByFileSystemWritableFileStream2()
	return
}

type FileSystemWritableFileStream struct {
	WritableStream
}

func (this FileSystemWritableFileStream) Once() FileSystemWritableFileStream {
	this.ref.Once()
	return this
}

func (this FileSystemWritableFileStream) Ref() js.Ref {
	return this.WritableStream.Ref()
}

func (this FileSystemWritableFileStream) FromRef(ref js.Ref) FileSystemWritableFileStream {
	this.WritableStream = this.WritableStream.FromRef(ref)
	return this
}

func (this FileSystemWritableFileStream) Free() {
	this.ref.Free()
}

// HasFuncWrite returns true if the method "FileSystemWritableFileStream.write" exists.
func (this FileSystemWritableFileStream) HasFuncWrite() bool {
	return js.True == bindings.HasFuncFileSystemWritableFileStreamWrite(
		this.ref,
	)
}

// FuncWrite returns the method "FileSystemWritableFileStream.write".
func (this FileSystemWritableFileStream) FuncWrite() (fn js.Func[func(data FileSystemWriteChunkType) js.Promise[js.Void]]) {
	bindings.FuncFileSystemWritableFileStreamWrite(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Write calls the method "FileSystemWritableFileStream.write".
func (this FileSystemWritableFileStream) Write(data FileSystemWriteChunkType) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemWritableFileStreamWrite(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryWrite calls the method "FileSystemWritableFileStream.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemWritableFileStream) TryWrite(data FileSystemWriteChunkType) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemWritableFileStreamWrite(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSeek returns true if the method "FileSystemWritableFileStream.seek" exists.
func (this FileSystemWritableFileStream) HasFuncSeek() bool {
	return js.True == bindings.HasFuncFileSystemWritableFileStreamSeek(
		this.ref,
	)
}

// FuncSeek returns the method "FileSystemWritableFileStream.seek".
func (this FileSystemWritableFileStream) FuncSeek() (fn js.Func[func(position uint64) js.Promise[js.Void]]) {
	bindings.FuncFileSystemWritableFileStreamSeek(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Seek calls the method "FileSystemWritableFileStream.seek".
func (this FileSystemWritableFileStream) Seek(position uint64) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemWritableFileStreamSeek(
		this.ref, js.Pointer(&ret),
		float64(position),
	)

	return
}

// TrySeek calls the method "FileSystemWritableFileStream.seek"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemWritableFileStream) TrySeek(position uint64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemWritableFileStreamSeek(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(position),
	)

	return
}

// HasFuncTruncate returns true if the method "FileSystemWritableFileStream.truncate" exists.
func (this FileSystemWritableFileStream) HasFuncTruncate() bool {
	return js.True == bindings.HasFuncFileSystemWritableFileStreamTruncate(
		this.ref,
	)
}

// FuncTruncate returns the method "FileSystemWritableFileStream.truncate".
func (this FileSystemWritableFileStream) FuncTruncate() (fn js.Func[func(size uint64) js.Promise[js.Void]]) {
	bindings.FuncFileSystemWritableFileStreamTruncate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Truncate calls the method "FileSystemWritableFileStream.truncate".
func (this FileSystemWritableFileStream) Truncate(size uint64) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemWritableFileStreamTruncate(
		this.ref, js.Pointer(&ret),
		float64(size),
	)

	return
}

// TryTruncate calls the method "FileSystemWritableFileStream.truncate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemWritableFileStream) TryTruncate(size uint64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemWritableFileStreamTruncate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(size),
	)

	return
}

type FileSystemCreateWritableOptions struct {
	// KeepExistingData is "FileSystemCreateWritableOptions.keepExistingData"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_KeepExistingData MUST be set to true to make this field effective.
	KeepExistingData bool

	FFI_USE_KeepExistingData bool // for KeepExistingData.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemCreateWritableOptions with all fields set.
func (p FileSystemCreateWritableOptions) FromRef(ref js.Ref) FileSystemCreateWritableOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileSystemCreateWritableOptions in the application heap.
func (p FileSystemCreateWritableOptions) New() js.Ref {
	return bindings.FileSystemCreateWritableOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileSystemCreateWritableOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemCreateWritableOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemCreateWritableOptions) Update(ref js.Ref) {
	bindings.FileSystemCreateWritableOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemCreateWritableOptions) FreeMembers(recursive bool) {
}

type FileSystemReadWriteOptions struct {
	// At is "FileSystemReadWriteOptions.at"
	//
	// Optional
	//
	// NOTE: FFI_USE_At MUST be set to true to make this field effective.
	At uint64

	FFI_USE_At bool // for At.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemReadWriteOptions with all fields set.
func (p FileSystemReadWriteOptions) FromRef(ref js.Ref) FileSystemReadWriteOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileSystemReadWriteOptions in the application heap.
func (p FileSystemReadWriteOptions) New() js.Ref {
	return bindings.FileSystemReadWriteOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileSystemReadWriteOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemReadWriteOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemReadWriteOptions) Update(ref js.Ref) {
	bindings.FileSystemReadWriteOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemReadWriteOptions) FreeMembers(recursive bool) {
}

type FileSystemSyncAccessHandle struct {
	ref js.Ref
}

func (this FileSystemSyncAccessHandle) Once() FileSystemSyncAccessHandle {
	this.ref.Once()
	return this
}

func (this FileSystemSyncAccessHandle) Ref() js.Ref {
	return this.ref
}

func (this FileSystemSyncAccessHandle) FromRef(ref js.Ref) FileSystemSyncAccessHandle {
	this.ref = ref
	return this
}

func (this FileSystemSyncAccessHandle) Free() {
	this.ref.Free()
}

// HasFuncRead returns true if the method "FileSystemSyncAccessHandle.read" exists.
func (this FileSystemSyncAccessHandle) HasFuncRead() bool {
	return js.True == bindings.HasFuncFileSystemSyncAccessHandleRead(
		this.ref,
	)
}

// FuncRead returns the method "FileSystemSyncAccessHandle.read".
func (this FileSystemSyncAccessHandle) FuncRead() (fn js.Func[func(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) uint64]) {
	bindings.FuncFileSystemSyncAccessHandleRead(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Read calls the method "FileSystemSyncAccessHandle.read".
func (this FileSystemSyncAccessHandle) Read(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleRead(
		this.ref, js.Pointer(&ret),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryRead calls the method "FileSystemSyncAccessHandle.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemSyncAccessHandle) TryRead(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleRead(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncRead1 returns true if the method "FileSystemSyncAccessHandle.read" exists.
func (this FileSystemSyncAccessHandle) HasFuncRead1() bool {
	return js.True == bindings.HasFuncFileSystemSyncAccessHandleRead1(
		this.ref,
	)
}

// FuncRead1 returns the method "FileSystemSyncAccessHandle.read".
func (this FileSystemSyncAccessHandle) FuncRead1() (fn js.Func[func(buffer AllowSharedBufferSource) uint64]) {
	bindings.FuncFileSystemSyncAccessHandleRead1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Read1 calls the method "FileSystemSyncAccessHandle.read".
func (this FileSystemSyncAccessHandle) Read1(buffer AllowSharedBufferSource) (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleRead1(
		this.ref, js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryRead1 calls the method "FileSystemSyncAccessHandle.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemSyncAccessHandle) TryRead1(buffer AllowSharedBufferSource) (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleRead1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasFuncWrite returns true if the method "FileSystemSyncAccessHandle.write" exists.
func (this FileSystemSyncAccessHandle) HasFuncWrite() bool {
	return js.True == bindings.HasFuncFileSystemSyncAccessHandleWrite(
		this.ref,
	)
}

// FuncWrite returns the method "FileSystemSyncAccessHandle.write".
func (this FileSystemSyncAccessHandle) FuncWrite() (fn js.Func[func(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) uint64]) {
	bindings.FuncFileSystemSyncAccessHandleWrite(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Write calls the method "FileSystemSyncAccessHandle.write".
func (this FileSystemSyncAccessHandle) Write(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleWrite(
		this.ref, js.Pointer(&ret),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryWrite calls the method "FileSystemSyncAccessHandle.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemSyncAccessHandle) TryWrite(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleWrite(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncWrite1 returns true if the method "FileSystemSyncAccessHandle.write" exists.
func (this FileSystemSyncAccessHandle) HasFuncWrite1() bool {
	return js.True == bindings.HasFuncFileSystemSyncAccessHandleWrite1(
		this.ref,
	)
}

// FuncWrite1 returns the method "FileSystemSyncAccessHandle.write".
func (this FileSystemSyncAccessHandle) FuncWrite1() (fn js.Func[func(buffer AllowSharedBufferSource) uint64]) {
	bindings.FuncFileSystemSyncAccessHandleWrite1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Write1 calls the method "FileSystemSyncAccessHandle.write".
func (this FileSystemSyncAccessHandle) Write1(buffer AllowSharedBufferSource) (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleWrite1(
		this.ref, js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryWrite1 calls the method "FileSystemSyncAccessHandle.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemSyncAccessHandle) TryWrite1(buffer AllowSharedBufferSource) (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleWrite1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasFuncTruncate returns true if the method "FileSystemSyncAccessHandle.truncate" exists.
func (this FileSystemSyncAccessHandle) HasFuncTruncate() bool {
	return js.True == bindings.HasFuncFileSystemSyncAccessHandleTruncate(
		this.ref,
	)
}

// FuncTruncate returns the method "FileSystemSyncAccessHandle.truncate".
func (this FileSystemSyncAccessHandle) FuncTruncate() (fn js.Func[func(newSize uint64)]) {
	bindings.FuncFileSystemSyncAccessHandleTruncate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Truncate calls the method "FileSystemSyncAccessHandle.truncate".
func (this FileSystemSyncAccessHandle) Truncate(newSize uint64) (ret js.Void) {
	bindings.CallFileSystemSyncAccessHandleTruncate(
		this.ref, js.Pointer(&ret),
		float64(newSize),
	)

	return
}

// TryTruncate calls the method "FileSystemSyncAccessHandle.truncate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemSyncAccessHandle) TryTruncate(newSize uint64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleTruncate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(newSize),
	)

	return
}

// HasFuncGetSize returns true if the method "FileSystemSyncAccessHandle.getSize" exists.
func (this FileSystemSyncAccessHandle) HasFuncGetSize() bool {
	return js.True == bindings.HasFuncFileSystemSyncAccessHandleGetSize(
		this.ref,
	)
}

// FuncGetSize returns the method "FileSystemSyncAccessHandle.getSize".
func (this FileSystemSyncAccessHandle) FuncGetSize() (fn js.Func[func() uint64]) {
	bindings.FuncFileSystemSyncAccessHandleGetSize(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSize calls the method "FileSystemSyncAccessHandle.getSize".
func (this FileSystemSyncAccessHandle) GetSize() (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleGetSize(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSize calls the method "FileSystemSyncAccessHandle.getSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemSyncAccessHandle) TryGetSize() (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleGetSize(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFlush returns true if the method "FileSystemSyncAccessHandle.flush" exists.
func (this FileSystemSyncAccessHandle) HasFuncFlush() bool {
	return js.True == bindings.HasFuncFileSystemSyncAccessHandleFlush(
		this.ref,
	)
}

// FuncFlush returns the method "FileSystemSyncAccessHandle.flush".
func (this FileSystemSyncAccessHandle) FuncFlush() (fn js.Func[func()]) {
	bindings.FuncFileSystemSyncAccessHandleFlush(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Flush calls the method "FileSystemSyncAccessHandle.flush".
func (this FileSystemSyncAccessHandle) Flush() (ret js.Void) {
	bindings.CallFileSystemSyncAccessHandleFlush(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "FileSystemSyncAccessHandle.flush"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemSyncAccessHandle) TryFlush() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleFlush(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "FileSystemSyncAccessHandle.close" exists.
func (this FileSystemSyncAccessHandle) HasFuncClose() bool {
	return js.True == bindings.HasFuncFileSystemSyncAccessHandleClose(
		this.ref,
	)
}

// FuncClose returns the method "FileSystemSyncAccessHandle.close".
func (this FileSystemSyncAccessHandle) FuncClose() (fn js.Func[func()]) {
	bindings.FuncFileSystemSyncAccessHandleClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "FileSystemSyncAccessHandle.close".
func (this FileSystemSyncAccessHandle) Close() (ret js.Void) {
	bindings.CallFileSystemSyncAccessHandleClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "FileSystemSyncAccessHandle.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemSyncAccessHandle) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type FileSystemFileHandle struct {
	FileSystemHandle
}

func (this FileSystemFileHandle) Once() FileSystemFileHandle {
	this.ref.Once()
	return this
}

func (this FileSystemFileHandle) Ref() js.Ref {
	return this.FileSystemHandle.Ref()
}

func (this FileSystemFileHandle) FromRef(ref js.Ref) FileSystemFileHandle {
	this.FileSystemHandle = this.FileSystemHandle.FromRef(ref)
	return this
}

func (this FileSystemFileHandle) Free() {
	this.ref.Free()
}

// HasFuncGetFile returns true if the method "FileSystemFileHandle.getFile" exists.
func (this FileSystemFileHandle) HasFuncGetFile() bool {
	return js.True == bindings.HasFuncFileSystemFileHandleGetFile(
		this.ref,
	)
}

// FuncGetFile returns the method "FileSystemFileHandle.getFile".
func (this FileSystemFileHandle) FuncGetFile() (fn js.Func[func() js.Promise[File]]) {
	bindings.FuncFileSystemFileHandleGetFile(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFile calls the method "FileSystemFileHandle.getFile".
func (this FileSystemFileHandle) GetFile() (ret js.Promise[File]) {
	bindings.CallFileSystemFileHandleGetFile(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetFile calls the method "FileSystemFileHandle.getFile"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemFileHandle) TryGetFile() (ret js.Promise[File], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileHandleGetFile(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateWritable returns true if the method "FileSystemFileHandle.createWritable" exists.
func (this FileSystemFileHandle) HasFuncCreateWritable() bool {
	return js.True == bindings.HasFuncFileSystemFileHandleCreateWritable(
		this.ref,
	)
}

// FuncCreateWritable returns the method "FileSystemFileHandle.createWritable".
func (this FileSystemFileHandle) FuncCreateWritable() (fn js.Func[func(options FileSystemCreateWritableOptions) js.Promise[FileSystemWritableFileStream]]) {
	bindings.FuncFileSystemFileHandleCreateWritable(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateWritable calls the method "FileSystemFileHandle.createWritable".
func (this FileSystemFileHandle) CreateWritable(options FileSystemCreateWritableOptions) (ret js.Promise[FileSystemWritableFileStream]) {
	bindings.CallFileSystemFileHandleCreateWritable(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateWritable calls the method "FileSystemFileHandle.createWritable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemFileHandle) TryCreateWritable(options FileSystemCreateWritableOptions) (ret js.Promise[FileSystemWritableFileStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileHandleCreateWritable(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateWritable1 returns true if the method "FileSystemFileHandle.createWritable" exists.
func (this FileSystemFileHandle) HasFuncCreateWritable1() bool {
	return js.True == bindings.HasFuncFileSystemFileHandleCreateWritable1(
		this.ref,
	)
}

// FuncCreateWritable1 returns the method "FileSystemFileHandle.createWritable".
func (this FileSystemFileHandle) FuncCreateWritable1() (fn js.Func[func() js.Promise[FileSystemWritableFileStream]]) {
	bindings.FuncFileSystemFileHandleCreateWritable1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateWritable1 calls the method "FileSystemFileHandle.createWritable".
func (this FileSystemFileHandle) CreateWritable1() (ret js.Promise[FileSystemWritableFileStream]) {
	bindings.CallFileSystemFileHandleCreateWritable1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateWritable1 calls the method "FileSystemFileHandle.createWritable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemFileHandle) TryCreateWritable1() (ret js.Promise[FileSystemWritableFileStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileHandleCreateWritable1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateSyncAccessHandle returns true if the method "FileSystemFileHandle.createSyncAccessHandle" exists.
func (this FileSystemFileHandle) HasFuncCreateSyncAccessHandle() bool {
	return js.True == bindings.HasFuncFileSystemFileHandleCreateSyncAccessHandle(
		this.ref,
	)
}

// FuncCreateSyncAccessHandle returns the method "FileSystemFileHandle.createSyncAccessHandle".
func (this FileSystemFileHandle) FuncCreateSyncAccessHandle() (fn js.Func[func() js.Promise[FileSystemSyncAccessHandle]]) {
	bindings.FuncFileSystemFileHandleCreateSyncAccessHandle(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSyncAccessHandle calls the method "FileSystemFileHandle.createSyncAccessHandle".
func (this FileSystemFileHandle) CreateSyncAccessHandle() (ret js.Promise[FileSystemSyncAccessHandle]) {
	bindings.CallFileSystemFileHandleCreateSyncAccessHandle(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSyncAccessHandle calls the method "FileSystemFileHandle.createSyncAccessHandle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemFileHandle) TryCreateSyncAccessHandle() (ret js.Promise[FileSystemSyncAccessHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileHandleCreateSyncAccessHandle(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type FilePickerAcceptType struct {
	// Description is "FilePickerAcceptType.description"
	//
	// Optional, defaults to "".
	Description js.String
	// Accept is "FilePickerAcceptType.accept"
	//
	// Optional
	Accept js.Record[OneOf_String_ArrayString]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FilePickerAcceptType with all fields set.
func (p FilePickerAcceptType) FromRef(ref js.Ref) FilePickerAcceptType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FilePickerAcceptType in the application heap.
func (p FilePickerAcceptType) New() js.Ref {
	return bindings.FilePickerAcceptTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FilePickerAcceptType) UpdateFrom(ref js.Ref) {
	bindings.FilePickerAcceptTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FilePickerAcceptType) Update(ref js.Ref) {
	bindings.FilePickerAcceptTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FilePickerAcceptType) FreeMembers(recursive bool) {
	js.Free(
		p.Description.Ref(),
		p.Accept.Ref(),
	)
	p.Description = p.Description.FromRef(js.Undefined)
	p.Accept = p.Accept.FromRef(js.Undefined)
}

type WellKnownDirectory uint32

const (
	_ WellKnownDirectory = iota

	WellKnownDirectory_DESKTOP
	WellKnownDirectory_DOCUMENTS
	WellKnownDirectory_DOWNLOADS
	WellKnownDirectory_MUSIC
	WellKnownDirectory_PICTURES
	WellKnownDirectory_VIDEOS
)

func (WellKnownDirectory) FromRef(str js.Ref) WellKnownDirectory {
	return WellKnownDirectory(bindings.ConstOfWellKnownDirectory(str))
}

func (x WellKnownDirectory) String() (string, bool) {
	switch x {
	case WellKnownDirectory_DESKTOP:
		return "desktop", true
	case WellKnownDirectory_DOCUMENTS:
		return "documents", true
	case WellKnownDirectory_DOWNLOADS:
		return "downloads", true
	case WellKnownDirectory_MUSIC:
		return "music", true
	case WellKnownDirectory_PICTURES:
		return "pictures", true
	case WellKnownDirectory_VIDEOS:
		return "videos", true
	default:
		return "", false
	}
}

type OneOf_WellKnownDirectory_FileSystemHandle struct {
	ref js.Ref
}

func (x OneOf_WellKnownDirectory_FileSystemHandle) Ref() js.Ref {
	return x.ref
}

func (x OneOf_WellKnownDirectory_FileSystemHandle) Free() {
	x.ref.Free()
}

func (x OneOf_WellKnownDirectory_FileSystemHandle) FromRef(ref js.Ref) OneOf_WellKnownDirectory_FileSystemHandle {
	return OneOf_WellKnownDirectory_FileSystemHandle{
		ref: ref,
	}
}

func (x OneOf_WellKnownDirectory_FileSystemHandle) WellKnownDirectory() WellKnownDirectory {
	return WellKnownDirectory(0).FromRef(x.ref)
}

func (x OneOf_WellKnownDirectory_FileSystemHandle) FileSystemHandle() FileSystemHandle {
	return FileSystemHandle{}.FromRef(x.ref)
}

type StartInDirectory = OneOf_WellKnownDirectory_FileSystemHandle

type OpenFilePickerOptions struct {
	// Multiple is "OpenFilePickerOptions.multiple"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Multiple MUST be set to true to make this field effective.
	Multiple bool
	// Types is "OpenFilePickerOptions.types"
	//
	// Optional
	Types js.Array[FilePickerAcceptType]
	// ExcludeAcceptAllOption is "OpenFilePickerOptions.excludeAcceptAllOption"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ExcludeAcceptAllOption MUST be set to true to make this field effective.
	ExcludeAcceptAllOption bool
	// Id is "OpenFilePickerOptions.id"
	//
	// Optional
	Id js.String
	// StartIn is "OpenFilePickerOptions.startIn"
	//
	// Optional
	StartIn StartInDirectory

	FFI_USE_Multiple               bool // for Multiple.
	FFI_USE_ExcludeAcceptAllOption bool // for ExcludeAcceptAllOption.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OpenFilePickerOptions with all fields set.
func (p OpenFilePickerOptions) FromRef(ref js.Ref) OpenFilePickerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OpenFilePickerOptions in the application heap.
func (p OpenFilePickerOptions) New() js.Ref {
	return bindings.OpenFilePickerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OpenFilePickerOptions) UpdateFrom(ref js.Ref) {
	bindings.OpenFilePickerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OpenFilePickerOptions) Update(ref js.Ref) {
	bindings.OpenFilePickerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OpenFilePickerOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Types.Ref(),
		p.Id.Ref(),
		p.StartIn.Ref(),
	)
	p.Types = p.Types.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.StartIn = p.StartIn.FromRef(js.Undefined)
}

type SaveFilePickerOptions struct {
	// SuggestedName is "SaveFilePickerOptions.suggestedName"
	//
	// Optional
	SuggestedName js.String
	// Types is "SaveFilePickerOptions.types"
	//
	// Optional
	Types js.Array[FilePickerAcceptType]
	// ExcludeAcceptAllOption is "SaveFilePickerOptions.excludeAcceptAllOption"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ExcludeAcceptAllOption MUST be set to true to make this field effective.
	ExcludeAcceptAllOption bool
	// Id is "SaveFilePickerOptions.id"
	//
	// Optional
	Id js.String
	// StartIn is "SaveFilePickerOptions.startIn"
	//
	// Optional
	StartIn StartInDirectory

	FFI_USE_ExcludeAcceptAllOption bool // for ExcludeAcceptAllOption.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SaveFilePickerOptions with all fields set.
func (p SaveFilePickerOptions) FromRef(ref js.Ref) SaveFilePickerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SaveFilePickerOptions in the application heap.
func (p SaveFilePickerOptions) New() js.Ref {
	return bindings.SaveFilePickerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SaveFilePickerOptions) UpdateFrom(ref js.Ref) {
	bindings.SaveFilePickerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SaveFilePickerOptions) Update(ref js.Ref) {
	bindings.SaveFilePickerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SaveFilePickerOptions) FreeMembers(recursive bool) {
	js.Free(
		p.SuggestedName.Ref(),
		p.Types.Ref(),
		p.Id.Ref(),
		p.StartIn.Ref(),
	)
	p.SuggestedName = p.SuggestedName.FromRef(js.Undefined)
	p.Types = p.Types.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.StartIn = p.StartIn.FromRef(js.Undefined)
}

type FileSystemGetFileOptions struct {
	// Create is "FileSystemGetFileOptions.create"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Create MUST be set to true to make this field effective.
	Create bool

	FFI_USE_Create bool // for Create.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemGetFileOptions with all fields set.
func (p FileSystemGetFileOptions) FromRef(ref js.Ref) FileSystemGetFileOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileSystemGetFileOptions in the application heap.
func (p FileSystemGetFileOptions) New() js.Ref {
	return bindings.FileSystemGetFileOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileSystemGetFileOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemGetFileOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemGetFileOptions) Update(ref js.Ref) {
	bindings.FileSystemGetFileOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemGetFileOptions) FreeMembers(recursive bool) {
}

type FileSystemGetDirectoryOptions struct {
	// Create is "FileSystemGetDirectoryOptions.create"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Create MUST be set to true to make this field effective.
	Create bool

	FFI_USE_Create bool // for Create.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemGetDirectoryOptions with all fields set.
func (p FileSystemGetDirectoryOptions) FromRef(ref js.Ref) FileSystemGetDirectoryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileSystemGetDirectoryOptions in the application heap.
func (p FileSystemGetDirectoryOptions) New() js.Ref {
	return bindings.FileSystemGetDirectoryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileSystemGetDirectoryOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemGetDirectoryOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemGetDirectoryOptions) Update(ref js.Ref) {
	bindings.FileSystemGetDirectoryOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemGetDirectoryOptions) FreeMembers(recursive bool) {
}

type FileSystemRemoveOptions struct {
	// Recursive is "FileSystemRemoveOptions.recursive"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Recursive MUST be set to true to make this field effective.
	Recursive bool

	FFI_USE_Recursive bool // for Recursive.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemRemoveOptions with all fields set.
func (p FileSystemRemoveOptions) FromRef(ref js.Ref) FileSystemRemoveOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileSystemRemoveOptions in the application heap.
func (p FileSystemRemoveOptions) New() js.Ref {
	return bindings.FileSystemRemoveOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FileSystemRemoveOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemRemoveOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FileSystemRemoveOptions) Update(ref js.Ref) {
	bindings.FileSystemRemoveOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FileSystemRemoveOptions) FreeMembers(recursive bool) {
}

type FileSystemDirectoryHandle struct {
	FileSystemHandle
}

func (this FileSystemDirectoryHandle) Once() FileSystemDirectoryHandle {
	this.ref.Once()
	return this
}

func (this FileSystemDirectoryHandle) Ref() js.Ref {
	return this.FileSystemHandle.Ref()
}

func (this FileSystemDirectoryHandle) FromRef(ref js.Ref) FileSystemDirectoryHandle {
	this.FileSystemHandle = this.FileSystemHandle.FromRef(ref)
	return this
}

func (this FileSystemDirectoryHandle) Free() {
	this.ref.Free()
}

// HasFuncGetFileHandle returns true if the method "FileSystemDirectoryHandle.getFileHandle" exists.
func (this FileSystemDirectoryHandle) HasFuncGetFileHandle() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryHandleGetFileHandle(
		this.ref,
	)
}

// FuncGetFileHandle returns the method "FileSystemDirectoryHandle.getFileHandle".
func (this FileSystemDirectoryHandle) FuncGetFileHandle() (fn js.Func[func(name js.String, options FileSystemGetFileOptions) js.Promise[FileSystemFileHandle]]) {
	bindings.FuncFileSystemDirectoryHandleGetFileHandle(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFileHandle calls the method "FileSystemDirectoryHandle.getFileHandle".
func (this FileSystemDirectoryHandle) GetFileHandle(name js.String, options FileSystemGetFileOptions) (ret js.Promise[FileSystemFileHandle]) {
	bindings.CallFileSystemDirectoryHandleGetFileHandle(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryGetFileHandle calls the method "FileSystemDirectoryHandle.getFileHandle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryHandle) TryGetFileHandle(name js.String, options FileSystemGetFileOptions) (ret js.Promise[FileSystemFileHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleGetFileHandle(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetFileHandle1 returns true if the method "FileSystemDirectoryHandle.getFileHandle" exists.
func (this FileSystemDirectoryHandle) HasFuncGetFileHandle1() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryHandleGetFileHandle1(
		this.ref,
	)
}

// FuncGetFileHandle1 returns the method "FileSystemDirectoryHandle.getFileHandle".
func (this FileSystemDirectoryHandle) FuncGetFileHandle1() (fn js.Func[func(name js.String) js.Promise[FileSystemFileHandle]]) {
	bindings.FuncFileSystemDirectoryHandleGetFileHandle1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFileHandle1 calls the method "FileSystemDirectoryHandle.getFileHandle".
func (this FileSystemDirectoryHandle) GetFileHandle1(name js.String) (ret js.Promise[FileSystemFileHandle]) {
	bindings.CallFileSystemDirectoryHandleGetFileHandle1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetFileHandle1 calls the method "FileSystemDirectoryHandle.getFileHandle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryHandle) TryGetFileHandle1(name js.String) (ret js.Promise[FileSystemFileHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleGetFileHandle1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetDirectoryHandle returns true if the method "FileSystemDirectoryHandle.getDirectoryHandle" exists.
func (this FileSystemDirectoryHandle) HasFuncGetDirectoryHandle() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryHandleGetDirectoryHandle(
		this.ref,
	)
}

// FuncGetDirectoryHandle returns the method "FileSystemDirectoryHandle.getDirectoryHandle".
func (this FileSystemDirectoryHandle) FuncGetDirectoryHandle() (fn js.Func[func(name js.String, options FileSystemGetDirectoryOptions) js.Promise[FileSystemDirectoryHandle]]) {
	bindings.FuncFileSystemDirectoryHandleGetDirectoryHandle(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDirectoryHandle calls the method "FileSystemDirectoryHandle.getDirectoryHandle".
func (this FileSystemDirectoryHandle) GetDirectoryHandle(name js.String, options FileSystemGetDirectoryOptions) (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallFileSystemDirectoryHandleGetDirectoryHandle(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryGetDirectoryHandle calls the method "FileSystemDirectoryHandle.getDirectoryHandle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryHandle) TryGetDirectoryHandle(name js.String, options FileSystemGetDirectoryOptions) (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleGetDirectoryHandle(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncGetDirectoryHandle1 returns true if the method "FileSystemDirectoryHandle.getDirectoryHandle" exists.
func (this FileSystemDirectoryHandle) HasFuncGetDirectoryHandle1() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryHandleGetDirectoryHandle1(
		this.ref,
	)
}

// FuncGetDirectoryHandle1 returns the method "FileSystemDirectoryHandle.getDirectoryHandle".
func (this FileSystemDirectoryHandle) FuncGetDirectoryHandle1() (fn js.Func[func(name js.String) js.Promise[FileSystemDirectoryHandle]]) {
	bindings.FuncFileSystemDirectoryHandleGetDirectoryHandle1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDirectoryHandle1 calls the method "FileSystemDirectoryHandle.getDirectoryHandle".
func (this FileSystemDirectoryHandle) GetDirectoryHandle1(name js.String) (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallFileSystemDirectoryHandleGetDirectoryHandle1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetDirectoryHandle1 calls the method "FileSystemDirectoryHandle.getDirectoryHandle"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryHandle) TryGetDirectoryHandle1(name js.String) (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleGetDirectoryHandle1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncRemoveEntry returns true if the method "FileSystemDirectoryHandle.removeEntry" exists.
func (this FileSystemDirectoryHandle) HasFuncRemoveEntry() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryHandleRemoveEntry(
		this.ref,
	)
}

// FuncRemoveEntry returns the method "FileSystemDirectoryHandle.removeEntry".
func (this FileSystemDirectoryHandle) FuncRemoveEntry() (fn js.Func[func(name js.String, options FileSystemRemoveOptions) js.Promise[js.Void]]) {
	bindings.FuncFileSystemDirectoryHandleRemoveEntry(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveEntry calls the method "FileSystemDirectoryHandle.removeEntry".
func (this FileSystemDirectoryHandle) RemoveEntry(name js.String, options FileSystemRemoveOptions) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemDirectoryHandleRemoveEntry(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryRemoveEntry calls the method "FileSystemDirectoryHandle.removeEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryHandle) TryRemoveEntry(name js.String, options FileSystemRemoveOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleRemoveEntry(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncRemoveEntry1 returns true if the method "FileSystemDirectoryHandle.removeEntry" exists.
func (this FileSystemDirectoryHandle) HasFuncRemoveEntry1() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryHandleRemoveEntry1(
		this.ref,
	)
}

// FuncRemoveEntry1 returns the method "FileSystemDirectoryHandle.removeEntry".
func (this FileSystemDirectoryHandle) FuncRemoveEntry1() (fn js.Func[func(name js.String) js.Promise[js.Void]]) {
	bindings.FuncFileSystemDirectoryHandleRemoveEntry1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveEntry1 calls the method "FileSystemDirectoryHandle.removeEntry".
func (this FileSystemDirectoryHandle) RemoveEntry1(name js.String) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemDirectoryHandleRemoveEntry1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryRemoveEntry1 calls the method "FileSystemDirectoryHandle.removeEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryHandle) TryRemoveEntry1(name js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleRemoveEntry1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncResolve returns true if the method "FileSystemDirectoryHandle.resolve" exists.
func (this FileSystemDirectoryHandle) HasFuncResolve() bool {
	return js.True == bindings.HasFuncFileSystemDirectoryHandleResolve(
		this.ref,
	)
}

// FuncResolve returns the method "FileSystemDirectoryHandle.resolve".
func (this FileSystemDirectoryHandle) FuncResolve() (fn js.Func[func(possibleDescendant FileSystemHandle) js.Promise[js.Array[js.String]]]) {
	bindings.FuncFileSystemDirectoryHandleResolve(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Resolve calls the method "FileSystemDirectoryHandle.resolve".
func (this FileSystemDirectoryHandle) Resolve(possibleDescendant FileSystemHandle) (ret js.Promise[js.Array[js.String]]) {
	bindings.CallFileSystemDirectoryHandleResolve(
		this.ref, js.Pointer(&ret),
		possibleDescendant.Ref(),
	)

	return
}

// TryResolve calls the method "FileSystemDirectoryHandle.resolve"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemDirectoryHandle) TryResolve(possibleDescendant FileSystemHandle) (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleResolve(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		possibleDescendant.Ref(),
	)

	return
}

type DirectoryPickerOptions struct {
	// Id is "DirectoryPickerOptions.id"
	//
	// Optional
	Id js.String
	// StartIn is "DirectoryPickerOptions.startIn"
	//
	// Optional
	StartIn StartInDirectory
	// Mode is "DirectoryPickerOptions.mode"
	//
	// Optional, defaults to "read".
	Mode FileSystemPermissionMode

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DirectoryPickerOptions with all fields set.
func (p DirectoryPickerOptions) FromRef(ref js.Ref) DirectoryPickerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DirectoryPickerOptions in the application heap.
func (p DirectoryPickerOptions) New() js.Ref {
	return bindings.DirectoryPickerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DirectoryPickerOptions) UpdateFrom(ref js.Ref) {
	bindings.DirectoryPickerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DirectoryPickerOptions) Update(ref js.Ref) {
	bindings.DirectoryPickerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DirectoryPickerOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.StartIn.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.StartIn = p.StartIn.FromRef(js.Undefined)
}

type FontData struct {
	ref js.Ref
}

func (this FontData) Once() FontData {
	this.ref.Once()
	return this
}

func (this FontData) Ref() js.Ref {
	return this.ref
}

func (this FontData) FromRef(ref js.Ref) FontData {
	this.ref = ref
	return this
}

func (this FontData) Free() {
	this.ref.Free()
}

// PostscriptName returns the value of property "FontData.postscriptName".
//
// It returns ok=false if there is no such property.
func (this FontData) PostscriptName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontDataPostscriptName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FullName returns the value of property "FontData.fullName".
//
// It returns ok=false if there is no such property.
func (this FontData) FullName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontDataFullName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Family returns the value of property "FontData.family".
//
// It returns ok=false if there is no such property.
func (this FontData) Family() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontDataFamily(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Style returns the value of property "FontData.style".
//
// It returns ok=false if there is no such property.
func (this FontData) Style() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontDataStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncBlob returns true if the method "FontData.blob" exists.
func (this FontData) HasFuncBlob() bool {
	return js.True == bindings.HasFuncFontDataBlob(
		this.ref,
	)
}

// FuncBlob returns the method "FontData.blob".
func (this FontData) FuncBlob() (fn js.Func[func() js.Promise[Blob]]) {
	bindings.FuncFontDataBlob(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Blob calls the method "FontData.blob".
func (this FontData) Blob() (ret js.Promise[Blob]) {
	bindings.CallFontDataBlob(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBlob calls the method "FontData.blob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FontData) TryBlob() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontDataBlob(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type QueryOptions struct {
	// PostscriptNames is "QueryOptions.postscriptNames"
	//
	// Optional
	PostscriptNames js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueryOptions with all fields set.
func (p QueryOptions) FromRef(ref js.Ref) QueryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueryOptions in the application heap.
func (p QueryOptions) New() js.Ref {
	return bindings.QueryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *QueryOptions) UpdateFrom(ref js.Ref) {
	bindings.QueryOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueryOptions) Update(ref js.Ref) {
	bindings.QueryOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueryOptions) FreeMembers(recursive bool) {
	js.Free(
		p.PostscriptNames.Ref(),
	)
	p.PostscriptNames = p.PostscriptNames.FromRef(js.Undefined)
}

type ScreenDetails struct {
	EventTarget
}

func (this ScreenDetails) Once() ScreenDetails {
	this.ref.Once()
	return this
}

func (this ScreenDetails) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this ScreenDetails) FromRef(ref js.Ref) ScreenDetails {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this ScreenDetails) Free() {
	this.ref.Free()
}

// Screens returns the value of property "ScreenDetails.screens".
//
// It returns ok=false if there is no such property.
func (this ScreenDetails) Screens() (ret js.FrozenArray[ScreenDetailed], ok bool) {
	ok = js.True == bindings.GetScreenDetailsScreens(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CurrentScreen returns the value of property "ScreenDetails.currentScreen".
//
// It returns ok=false if there is no such property.
func (this ScreenDetails) CurrentScreen() (ret ScreenDetailed, ok bool) {
	ok = js.True == bindings.GetScreenDetailsCurrentScreen(
		this.ref, js.Pointer(&ret),
	)
	return
}

type ItemType uint32

const (
	_ ItemType = iota

	ItemType_PRODUCT
	ItemType_SUBSCRIPTION
)

func (ItemType) FromRef(str js.Ref) ItemType {
	return ItemType(bindings.ConstOfItemType(str))
}

func (x ItemType) String() (string, bool) {
	switch x {
	case ItemType_PRODUCT:
		return "product", true
	case ItemType_SUBSCRIPTION:
		return "subscription", true
	default:
		return "", false
	}
}

type ItemDetails struct {
	// ItemId is "ItemDetails.itemId"
	//
	// Required
	ItemId js.String
	// Title is "ItemDetails.title"
	//
	// Required
	Title js.String
	// Price is "ItemDetails.price"
	//
	// Required
	//
	// NOTE: Price.FFI_USE MUST be set to true to get Price used.
	Price PaymentCurrencyAmount
	// Type is "ItemDetails.type"
	//
	// Optional
	Type ItemType
	// Description is "ItemDetails.description"
	//
	// Optional
	Description js.String
	// IconURLs is "ItemDetails.iconURLs"
	//
	// Optional
	IconURLs js.Array[js.String]
	// SubscriptionPeriod is "ItemDetails.subscriptionPeriod"
	//
	// Optional
	SubscriptionPeriod js.String
	// FreeTrialPeriod is "ItemDetails.freeTrialPeriod"
	//
	// Optional
	FreeTrialPeriod js.String
	// IntroductoryPrice is "ItemDetails.introductoryPrice"
	//
	// Optional
	//
	// NOTE: IntroductoryPrice.FFI_USE MUST be set to true to get IntroductoryPrice used.
	IntroductoryPrice PaymentCurrencyAmount
	// IntroductoryPricePeriod is "ItemDetails.introductoryPricePeriod"
	//
	// Optional
	IntroductoryPricePeriod js.String
	// IntroductoryPriceCycles is "ItemDetails.introductoryPriceCycles"
	//
	// Optional
	//
	// NOTE: FFI_USE_IntroductoryPriceCycles MUST be set to true to make this field effective.
	IntroductoryPriceCycles uint64

	FFI_USE_IntroductoryPriceCycles bool // for IntroductoryPriceCycles.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ItemDetails with all fields set.
func (p ItemDetails) FromRef(ref js.Ref) ItemDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ItemDetails in the application heap.
func (p ItemDetails) New() js.Ref {
	return bindings.ItemDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ItemDetails) UpdateFrom(ref js.Ref) {
	bindings.ItemDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ItemDetails) Update(ref js.Ref) {
	bindings.ItemDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ItemDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ItemId.Ref(),
		p.Title.Ref(),
		p.Description.Ref(),
		p.IconURLs.Ref(),
		p.SubscriptionPeriod.Ref(),
		p.FreeTrialPeriod.Ref(),
		p.IntroductoryPricePeriod.Ref(),
	)
	p.ItemId = p.ItemId.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Description = p.Description.FromRef(js.Undefined)
	p.IconURLs = p.IconURLs.FromRef(js.Undefined)
	p.SubscriptionPeriod = p.SubscriptionPeriod.FromRef(js.Undefined)
	p.FreeTrialPeriod = p.FreeTrialPeriod.FromRef(js.Undefined)
	p.IntroductoryPricePeriod = p.IntroductoryPricePeriod.FromRef(js.Undefined)
	if recursive {
		p.Price.FreeMembers(true)
		p.IntroductoryPrice.FreeMembers(true)
	}
}

type PurchaseDetails struct {
	// ItemId is "PurchaseDetails.itemId"
	//
	// Required
	ItemId js.String
	// PurchaseToken is "PurchaseDetails.purchaseToken"
	//
	// Required
	PurchaseToken js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PurchaseDetails with all fields set.
func (p PurchaseDetails) FromRef(ref js.Ref) PurchaseDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PurchaseDetails in the application heap.
func (p PurchaseDetails) New() js.Ref {
	return bindings.PurchaseDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PurchaseDetails) UpdateFrom(ref js.Ref) {
	bindings.PurchaseDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PurchaseDetails) Update(ref js.Ref) {
	bindings.PurchaseDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PurchaseDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ItemId.Ref(),
		p.PurchaseToken.Ref(),
	)
	p.ItemId = p.ItemId.FromRef(js.Undefined)
	p.PurchaseToken = p.PurchaseToken.FromRef(js.Undefined)
}

type DigitalGoodsService struct {
	ref js.Ref
}

func (this DigitalGoodsService) Once() DigitalGoodsService {
	this.ref.Once()
	return this
}

func (this DigitalGoodsService) Ref() js.Ref {
	return this.ref
}

func (this DigitalGoodsService) FromRef(ref js.Ref) DigitalGoodsService {
	this.ref = ref
	return this
}

func (this DigitalGoodsService) Free() {
	this.ref.Free()
}

// HasFuncGetDetails returns true if the method "DigitalGoodsService.getDetails" exists.
func (this DigitalGoodsService) HasFuncGetDetails() bool {
	return js.True == bindings.HasFuncDigitalGoodsServiceGetDetails(
		this.ref,
	)
}

// FuncGetDetails returns the method "DigitalGoodsService.getDetails".
func (this DigitalGoodsService) FuncGetDetails() (fn js.Func[func(itemIds js.Array[js.String]) js.Promise[js.Array[ItemDetails]]]) {
	bindings.FuncDigitalGoodsServiceGetDetails(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetDetails calls the method "DigitalGoodsService.getDetails".
func (this DigitalGoodsService) GetDetails(itemIds js.Array[js.String]) (ret js.Promise[js.Array[ItemDetails]]) {
	bindings.CallDigitalGoodsServiceGetDetails(
		this.ref, js.Pointer(&ret),
		itemIds.Ref(),
	)

	return
}

// TryGetDetails calls the method "DigitalGoodsService.getDetails"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DigitalGoodsService) TryGetDetails(itemIds js.Array[js.String]) (ret js.Promise[js.Array[ItemDetails]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDigitalGoodsServiceGetDetails(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		itemIds.Ref(),
	)

	return
}

// HasFuncListPurchases returns true if the method "DigitalGoodsService.listPurchases" exists.
func (this DigitalGoodsService) HasFuncListPurchases() bool {
	return js.True == bindings.HasFuncDigitalGoodsServiceListPurchases(
		this.ref,
	)
}

// FuncListPurchases returns the method "DigitalGoodsService.listPurchases".
func (this DigitalGoodsService) FuncListPurchases() (fn js.Func[func() js.Promise[js.Array[PurchaseDetails]]]) {
	bindings.FuncDigitalGoodsServiceListPurchases(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ListPurchases calls the method "DigitalGoodsService.listPurchases".
func (this DigitalGoodsService) ListPurchases() (ret js.Promise[js.Array[PurchaseDetails]]) {
	bindings.CallDigitalGoodsServiceListPurchases(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryListPurchases calls the method "DigitalGoodsService.listPurchases"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DigitalGoodsService) TryListPurchases() (ret js.Promise[js.Array[PurchaseDetails]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDigitalGoodsServiceListPurchases(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncListPurchaseHistory returns true if the method "DigitalGoodsService.listPurchaseHistory" exists.
func (this DigitalGoodsService) HasFuncListPurchaseHistory() bool {
	return js.True == bindings.HasFuncDigitalGoodsServiceListPurchaseHistory(
		this.ref,
	)
}

// FuncListPurchaseHistory returns the method "DigitalGoodsService.listPurchaseHistory".
func (this DigitalGoodsService) FuncListPurchaseHistory() (fn js.Func[func() js.Promise[js.Array[PurchaseDetails]]]) {
	bindings.FuncDigitalGoodsServiceListPurchaseHistory(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ListPurchaseHistory calls the method "DigitalGoodsService.listPurchaseHistory".
func (this DigitalGoodsService) ListPurchaseHistory() (ret js.Promise[js.Array[PurchaseDetails]]) {
	bindings.CallDigitalGoodsServiceListPurchaseHistory(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryListPurchaseHistory calls the method "DigitalGoodsService.listPurchaseHistory"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DigitalGoodsService) TryListPurchaseHistory() (ret js.Promise[js.Array[PurchaseDetails]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDigitalGoodsServiceListPurchaseHistory(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncConsume returns true if the method "DigitalGoodsService.consume" exists.
func (this DigitalGoodsService) HasFuncConsume() bool {
	return js.True == bindings.HasFuncDigitalGoodsServiceConsume(
		this.ref,
	)
}

// FuncConsume returns the method "DigitalGoodsService.consume".
func (this DigitalGoodsService) FuncConsume() (fn js.Func[func(purchaseToken js.String) js.Promise[js.Void]]) {
	bindings.FuncDigitalGoodsServiceConsume(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Consume calls the method "DigitalGoodsService.consume".
func (this DigitalGoodsService) Consume(purchaseToken js.String) (ret js.Promise[js.Void]) {
	bindings.CallDigitalGoodsServiceConsume(
		this.ref, js.Pointer(&ret),
		purchaseToken.Ref(),
	)

	return
}

// TryConsume calls the method "DigitalGoodsService.consume"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DigitalGoodsService) TryConsume(purchaseToken js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDigitalGoodsServiceConsume(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		purchaseToken.Ref(),
	)

	return
}

type OneOf_String_FuncFunction struct {
	ref js.Ref
}

func (x OneOf_String_FuncFunction) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_FuncFunction) Free() {
	x.ref.Free()
}

func (x OneOf_String_FuncFunction) FromRef(ref js.Ref) OneOf_String_FuncFunction {
	return OneOf_String_FuncFunction{
		ref: ref,
	}
}

func (x OneOf_String_FuncFunction) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_FuncFunction) FuncFunction() js.Func[func(arguments ...js.Any) js.Any] {
	return js.Func[func(arguments ...js.Any) js.Any]{}.FromRef(x.ref)
}

type TimerHandler = OneOf_String_FuncFunction

type VoidFunctionFunc func(this js.Ref) js.Ref

func (fn VoidFunctionFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidFunctionFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VoidFunction[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidFunction[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidFunction[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ImageOrientation uint32

const (
	_ ImageOrientation = iota

	ImageOrientation_FROM_IMAGE
	ImageOrientation_FLIP_Y
)

func (ImageOrientation) FromRef(str js.Ref) ImageOrientation {
	return ImageOrientation(bindings.ConstOfImageOrientation(str))
}

func (x ImageOrientation) String() (string, bool) {
	switch x {
	case ImageOrientation_FROM_IMAGE:
		return "from-image", true
	case ImageOrientation_FLIP_Y:
		return "flipY", true
	default:
		return "", false
	}
}

type PremultiplyAlpha uint32

const (
	_ PremultiplyAlpha = iota

	PremultiplyAlpha_NONE
	PremultiplyAlpha_PREMULTIPLY
	PremultiplyAlpha_DEFAULT
)

func (PremultiplyAlpha) FromRef(str js.Ref) PremultiplyAlpha {
	return PremultiplyAlpha(bindings.ConstOfPremultiplyAlpha(str))
}

func (x PremultiplyAlpha) String() (string, bool) {
	switch x {
	case PremultiplyAlpha_NONE:
		return "none", true
	case PremultiplyAlpha_PREMULTIPLY:
		return "premultiply", true
	case PremultiplyAlpha_DEFAULT:
		return "default", true
	default:
		return "", false
	}
}

type ResizeQuality uint32

const (
	_ ResizeQuality = iota

	ResizeQuality_PIXELATED
	ResizeQuality_LOW
	ResizeQuality_MEDIUM
	ResizeQuality_HIGH
)

func (ResizeQuality) FromRef(str js.Ref) ResizeQuality {
	return ResizeQuality(bindings.ConstOfResizeQuality(str))
}

func (x ResizeQuality) String() (string, bool) {
	switch x {
	case ResizeQuality_PIXELATED:
		return "pixelated", true
	case ResizeQuality_LOW:
		return "low", true
	case ResizeQuality_MEDIUM:
		return "medium", true
	case ResizeQuality_HIGH:
		return "high", true
	default:
		return "", false
	}
}

type ImageBitmapOptions struct {
	// ImageOrientation is "ImageBitmapOptions.imageOrientation"
	//
	// Optional, defaults to "from-image".
	ImageOrientation ImageOrientation
	// PremultiplyAlpha is "ImageBitmapOptions.premultiplyAlpha"
	//
	// Optional, defaults to "default".
	PremultiplyAlpha PremultiplyAlpha
	// ColorSpaceConversion is "ImageBitmapOptions.colorSpaceConversion"
	//
	// Optional, defaults to "default".
	ColorSpaceConversion ColorSpaceConversion
	// ResizeWidth is "ImageBitmapOptions.resizeWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResizeWidth MUST be set to true to make this field effective.
	ResizeWidth uint32
	// ResizeHeight is "ImageBitmapOptions.resizeHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_ResizeHeight MUST be set to true to make this field effective.
	ResizeHeight uint32
	// ResizeQuality is "ImageBitmapOptions.resizeQuality"
	//
	// Optional, defaults to "low".
	ResizeQuality ResizeQuality

	FFI_USE_ResizeWidth  bool // for ResizeWidth.
	FFI_USE_ResizeHeight bool // for ResizeHeight.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImageBitmapOptions with all fields set.
func (p ImageBitmapOptions) FromRef(ref js.Ref) ImageBitmapOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImageBitmapOptions in the application heap.
func (p ImageBitmapOptions) New() js.Ref {
	return bindings.ImageBitmapOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ImageBitmapOptions) UpdateFrom(ref js.Ref) {
	bindings.ImageBitmapOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImageBitmapOptions) Update(ref js.Ref) {
	bindings.ImageBitmapOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImageBitmapOptions) FreeMembers(recursive bool) {
}

type FrameRequestCallbackFunc func(this js.Ref, time DOMHighResTimeStamp) js.Ref

func (fn FrameRequestCallbackFunc) Register() js.Func[func(time DOMHighResTimeStamp)] {
	return js.RegisterCallback[func(time DOMHighResTimeStamp)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FrameRequestCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[DOMHighResTimeStamp]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FrameRequestCallback[T any] struct {
	Fn  func(arg T, this js.Ref, time DOMHighResTimeStamp) js.Ref
	Arg T
}

func (cb *FrameRequestCallback[T]) Register() js.Func[func(time DOMHighResTimeStamp)] {
	return js.RegisterCallback[func(time DOMHighResTimeStamp)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FrameRequestCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[DOMHighResTimeStamp]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ScrollRestoration uint32

const (
	_ ScrollRestoration = iota

	ScrollRestoration_AUTO
	ScrollRestoration_MANUAL
)

func (ScrollRestoration) FromRef(str js.Ref) ScrollRestoration {
	return ScrollRestoration(bindings.ConstOfScrollRestoration(str))
}

func (x ScrollRestoration) String() (string, bool) {
	switch x {
	case ScrollRestoration_AUTO:
		return "auto", true
	case ScrollRestoration_MANUAL:
		return "manual", true
	default:
		return "", false
	}
}

type History struct {
	ref js.Ref
}

func (this History) Once() History {
	this.ref.Once()
	return this
}

func (this History) Ref() js.Ref {
	return this.ref
}

func (this History) FromRef(ref js.Ref) History {
	this.ref = ref
	return this
}

func (this History) Free() {
	this.ref.Free()
}

// Length returns the value of property "History.length".
//
// It returns ok=false if there is no such property.
func (this History) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHistoryLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScrollRestoration returns the value of property "History.scrollRestoration".
//
// It returns ok=false if there is no such property.
func (this History) ScrollRestoration() (ret ScrollRestoration, ok bool) {
	ok = js.True == bindings.GetHistoryScrollRestoration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollRestoration sets the value of property "History.scrollRestoration" to val.
//
// It returns false if the property cannot be set.
func (this History) SetScrollRestoration(val ScrollRestoration) bool {
	return js.True == bindings.SetHistoryScrollRestoration(
		this.ref,
		uint32(val),
	)
}

// State returns the value of property "History.state".
//
// It returns ok=false if there is no such property.
func (this History) State() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetHistoryState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGo returns true if the method "History.go" exists.
func (this History) HasFuncGo() bool {
	return js.True == bindings.HasFuncHistoryGo(
		this.ref,
	)
}

// FuncGo returns the method "History.go".
func (this History) FuncGo() (fn js.Func[func(delta int32)]) {
	bindings.FuncHistoryGo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Go calls the method "History.go".
func (this History) Go(delta int32) (ret js.Void) {
	bindings.CallHistoryGo(
		this.ref, js.Pointer(&ret),
		int32(delta),
	)

	return
}

// TryGo calls the method "History.go"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this History) TryGo(delta int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryGo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(delta),
	)

	return
}

// HasFuncGo1 returns true if the method "History.go" exists.
func (this History) HasFuncGo1() bool {
	return js.True == bindings.HasFuncHistoryGo1(
		this.ref,
	)
}

// FuncGo1 returns the method "History.go".
func (this History) FuncGo1() (fn js.Func[func()]) {
	bindings.FuncHistoryGo1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Go1 calls the method "History.go".
func (this History) Go1() (ret js.Void) {
	bindings.CallHistoryGo1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGo1 calls the method "History.go"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this History) TryGo1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryGo1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBack returns true if the method "History.back" exists.
func (this History) HasFuncBack() bool {
	return js.True == bindings.HasFuncHistoryBack(
		this.ref,
	)
}

// FuncBack returns the method "History.back".
func (this History) FuncBack() (fn js.Func[func()]) {
	bindings.FuncHistoryBack(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Back calls the method "History.back".
func (this History) Back() (ret js.Void) {
	bindings.CallHistoryBack(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBack calls the method "History.back"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this History) TryBack() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryBack(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncForward returns true if the method "History.forward" exists.
func (this History) HasFuncForward() bool {
	return js.True == bindings.HasFuncHistoryForward(
		this.ref,
	)
}

// FuncForward returns the method "History.forward".
func (this History) FuncForward() (fn js.Func[func()]) {
	bindings.FuncHistoryForward(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Forward calls the method "History.forward".
func (this History) Forward() (ret js.Void) {
	bindings.CallHistoryForward(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryForward calls the method "History.forward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this History) TryForward() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryForward(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPushState returns true if the method "History.pushState" exists.
func (this History) HasFuncPushState() bool {
	return js.True == bindings.HasFuncHistoryPushState(
		this.ref,
	)
}

// FuncPushState returns the method "History.pushState".
func (this History) FuncPushState() (fn js.Func[func(data js.Any, unused js.String, url js.String)]) {
	bindings.FuncHistoryPushState(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PushState calls the method "History.pushState".
func (this History) PushState(data js.Any, unused js.String, url js.String) (ret js.Void) {
	bindings.CallHistoryPushState(
		this.ref, js.Pointer(&ret),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	return
}

// TryPushState calls the method "History.pushState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this History) TryPushState(data js.Any, unused js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryPushState(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	return
}

// HasFuncPushState1 returns true if the method "History.pushState" exists.
func (this History) HasFuncPushState1() bool {
	return js.True == bindings.HasFuncHistoryPushState1(
		this.ref,
	)
}

// FuncPushState1 returns the method "History.pushState".
func (this History) FuncPushState1() (fn js.Func[func(data js.Any, unused js.String)]) {
	bindings.FuncHistoryPushState1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PushState1 calls the method "History.pushState".
func (this History) PushState1(data js.Any, unused js.String) (ret js.Void) {
	bindings.CallHistoryPushState1(
		this.ref, js.Pointer(&ret),
		data.Ref(),
		unused.Ref(),
	)

	return
}

// TryPushState1 calls the method "History.pushState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this History) TryPushState1(data js.Any, unused js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryPushState1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		unused.Ref(),
	)

	return
}

// HasFuncReplaceState returns true if the method "History.replaceState" exists.
func (this History) HasFuncReplaceState() bool {
	return js.True == bindings.HasFuncHistoryReplaceState(
		this.ref,
	)
}

// FuncReplaceState returns the method "History.replaceState".
func (this History) FuncReplaceState() (fn js.Func[func(data js.Any, unused js.String, url js.String)]) {
	bindings.FuncHistoryReplaceState(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceState calls the method "History.replaceState".
func (this History) ReplaceState(data js.Any, unused js.String, url js.String) (ret js.Void) {
	bindings.CallHistoryReplaceState(
		this.ref, js.Pointer(&ret),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	return
}

// TryReplaceState calls the method "History.replaceState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this History) TryReplaceState(data js.Any, unused js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryReplaceState(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	return
}

// HasFuncReplaceState1 returns true if the method "History.replaceState" exists.
func (this History) HasFuncReplaceState1() bool {
	return js.True == bindings.HasFuncHistoryReplaceState1(
		this.ref,
	)
}

// FuncReplaceState1 returns the method "History.replaceState".
func (this History) FuncReplaceState1() (fn js.Func[func(data js.Any, unused js.String)]) {
	bindings.FuncHistoryReplaceState1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceState1 calls the method "History.replaceState".
func (this History) ReplaceState1(data js.Any, unused js.String) (ret js.Void) {
	bindings.CallHistoryReplaceState1(
		this.ref, js.Pointer(&ret),
		data.Ref(),
		unused.Ref(),
	)

	return
}

// TryReplaceState1 calls the method "History.replaceState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this History) TryReplaceState1(data js.Any, unused js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryReplaceState1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		unused.Ref(),
	)

	return
}

type NavigationHistoryEntry struct {
	EventTarget
}

func (this NavigationHistoryEntry) Once() NavigationHistoryEntry {
	this.ref.Once()
	return this
}

func (this NavigationHistoryEntry) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this NavigationHistoryEntry) FromRef(ref js.Ref) NavigationHistoryEntry {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this NavigationHistoryEntry) Free() {
	this.ref.Free()
}

// Url returns the value of property "NavigationHistoryEntry.url".
//
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntryUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Key returns the value of property "NavigationHistoryEntry.key".
//
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) Key() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntryKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "NavigationHistoryEntry.id".
//
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntryId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Index returns the value of property "NavigationHistoryEntry.index".
//
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) Index() (ret int64, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntryIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SameDocument returns the value of property "NavigationHistoryEntry.sameDocument".
//
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) SameDocument() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntrySameDocument(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetState returns true if the method "NavigationHistoryEntry.getState" exists.
func (this NavigationHistoryEntry) HasFuncGetState() bool {
	return js.True == bindings.HasFuncNavigationHistoryEntryGetState(
		this.ref,
	)
}

// FuncGetState returns the method "NavigationHistoryEntry.getState".
func (this NavigationHistoryEntry) FuncGetState() (fn js.Func[func() js.Any]) {
	bindings.FuncNavigationHistoryEntryGetState(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetState calls the method "NavigationHistoryEntry.getState".
func (this NavigationHistoryEntry) GetState() (ret js.Any) {
	bindings.CallNavigationHistoryEntryGetState(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetState calls the method "NavigationHistoryEntry.getState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigationHistoryEntry) TryGetState() (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationHistoryEntryGetState(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type NavigationUpdateCurrentEntryOptions struct {
	// State is "NavigationUpdateCurrentEntryOptions.state"
	//
	// Required
	State js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigationUpdateCurrentEntryOptions with all fields set.
func (p NavigationUpdateCurrentEntryOptions) FromRef(ref js.Ref) NavigationUpdateCurrentEntryOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NavigationUpdateCurrentEntryOptions in the application heap.
func (p NavigationUpdateCurrentEntryOptions) New() js.Ref {
	return bindings.NavigationUpdateCurrentEntryOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NavigationUpdateCurrentEntryOptions) UpdateFrom(ref js.Ref) {
	bindings.NavigationUpdateCurrentEntryOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NavigationUpdateCurrentEntryOptions) Update(ref js.Ref) {
	bindings.NavigationUpdateCurrentEntryOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NavigationUpdateCurrentEntryOptions) FreeMembers(recursive bool) {
	js.Free(
		p.State.Ref(),
	)
	p.State = p.State.FromRef(js.Undefined)
}

type NavigationResult struct {
	// Committed is "NavigationResult.committed"
	//
	// Optional
	Committed js.Promise[NavigationHistoryEntry]
	// Finished is "NavigationResult.finished"
	//
	// Optional
	Finished js.Promise[NavigationHistoryEntry]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigationResult with all fields set.
func (p NavigationResult) FromRef(ref js.Ref) NavigationResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NavigationResult in the application heap.
func (p NavigationResult) New() js.Ref {
	return bindings.NavigationResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *NavigationResult) UpdateFrom(ref js.Ref) {
	bindings.NavigationResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NavigationResult) Update(ref js.Ref) {
	bindings.NavigationResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NavigationResult) FreeMembers(recursive bool) {
	js.Free(
		p.Committed.Ref(),
		p.Finished.Ref(),
	)
	p.Committed = p.Committed.FromRef(js.Undefined)
	p.Finished = p.Finished.FromRef(js.Undefined)
}
