// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

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
		assert.Throw("invalid", "callback", "invocation")
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// DidTimeout returns the value of property "IdleDeadline.didTimeout".
//
// The returned bool will be false if there is no such property.
func (this IdleDeadline) DidTimeout() (bool, bool) {
	var _ok bool
	_ret := bindings.GetIdleDeadlineDidTimeout(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// TimeRemaining calls the method "IdleDeadline.timeRemaining".
//
// The returned bool will be false if there is no such method.
func (this IdleDeadline) TimeRemaining() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.CallIdleDeadlineTimeRemaining(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMHighResTimeStamp(_ret), _ok
}

// TimeRemainingFunc returns the method "IdleDeadline.timeRemaining".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IdleDeadline) TimeRemainingFunc() (fn js.Func[func() DOMHighResTimeStamp]) {
	return fn.FromRef(
		bindings.IdleDeadlineTimeRemainingFunc(
			this.Ref(),
		),
	)
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
func (p IdleRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.IdleRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IdleRequestOptions) Update(ref js.Ref) {
	bindings.IdleRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p WriteParams) UpdateFrom(ref js.Ref) {
	bindings.WriteParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WriteParams) Update(ref js.Ref) {
	bindings.WriteParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

func NewFileSystemWritableFileStream(underlyingSink js.Object, strategy QueuingStrategy) FileSystemWritableFileStream {
	return FileSystemWritableFileStream{}.FromRef(
		bindings.NewFileSystemWritableFileStreamByFileSystemWritableFileStream(
			underlyingSink.Ref(),
			js.Pointer(&strategy)),
	)
}

func NewFileSystemWritableFileStreamByFileSystemWritableFileStream1(underlyingSink js.Object) FileSystemWritableFileStream {
	return FileSystemWritableFileStream{}.FromRef(
		bindings.NewFileSystemWritableFileStreamByFileSystemWritableFileStream1(
			underlyingSink.Ref()),
	)
}

func NewFileSystemWritableFileStreamByFileSystemWritableFileStream2() FileSystemWritableFileStream {
	return FileSystemWritableFileStream{}.FromRef(
		bindings.NewFileSystemWritableFileStreamByFileSystemWritableFileStream2(),
	)
}

type FileSystemWritableFileStream struct {
	WritableStream
}

func (this FileSystemWritableFileStream) Once() FileSystemWritableFileStream {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Write calls the method "FileSystemWritableFileStream.write".
//
// The returned bool will be false if there is no such method.
func (this FileSystemWritableFileStream) Write(data FileSystemWriteChunkType) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemWritableFileStreamWrite(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WriteFunc returns the method "FileSystemWritableFileStream.write".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemWritableFileStream) WriteFunc() (fn js.Func[func(data FileSystemWriteChunkType) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemWritableFileStreamWriteFunc(
			this.Ref(),
		),
	)
}

// Seek calls the method "FileSystemWritableFileStream.seek".
//
// The returned bool will be false if there is no such method.
func (this FileSystemWritableFileStream) Seek(position uint64) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemWritableFileStreamSeek(
		this.Ref(), js.Pointer(&_ok),
		float64(position),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SeekFunc returns the method "FileSystemWritableFileStream.seek".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemWritableFileStream) SeekFunc() (fn js.Func[func(position uint64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemWritableFileStreamSeekFunc(
			this.Ref(),
		),
	)
}

// Truncate calls the method "FileSystemWritableFileStream.truncate".
//
// The returned bool will be false if there is no such method.
func (this FileSystemWritableFileStream) Truncate(size uint64) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemWritableFileStreamTruncate(
		this.Ref(), js.Pointer(&_ok),
		float64(size),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// TruncateFunc returns the method "FileSystemWritableFileStream.truncate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemWritableFileStream) TruncateFunc() (fn js.Func[func(size uint64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemWritableFileStreamTruncateFunc(
			this.Ref(),
		),
	)
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
func (p FileSystemCreateWritableOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemCreateWritableOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FileSystemCreateWritableOptions) Update(ref js.Ref) {
	bindings.FileSystemCreateWritableOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p FileSystemReadWriteOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemReadWriteOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FileSystemReadWriteOptions) Update(ref js.Ref) {
	bindings.FileSystemReadWriteOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type FileSystemSyncAccessHandle struct {
	ref js.Ref
}

func (this FileSystemSyncAccessHandle) Once() FileSystemSyncAccessHandle {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Read calls the method "FileSystemSyncAccessHandle.read".
//
// The returned bool will be false if there is no such method.
func (this FileSystemSyncAccessHandle) Read(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (uint64, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemSyncAccessHandleRead(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return uint64(_ret), _ok
}

// ReadFunc returns the method "FileSystemSyncAccessHandle.read".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemSyncAccessHandle) ReadFunc() (fn js.Func[func(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleReadFunc(
			this.Ref(),
		),
	)
}

// Read1 calls the method "FileSystemSyncAccessHandle.read".
//
// The returned bool will be false if there is no such method.
func (this FileSystemSyncAccessHandle) Read1(buffer AllowSharedBufferSource) (uint64, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemSyncAccessHandleRead1(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
	)

	return uint64(_ret), _ok
}

// Read1Func returns the method "FileSystemSyncAccessHandle.read".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemSyncAccessHandle) Read1Func() (fn js.Func[func(buffer AllowSharedBufferSource) uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleRead1Func(
			this.Ref(),
		),
	)
}

// Write calls the method "FileSystemSyncAccessHandle.write".
//
// The returned bool will be false if there is no such method.
func (this FileSystemSyncAccessHandle) Write(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (uint64, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemSyncAccessHandleWrite(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return uint64(_ret), _ok
}

// WriteFunc returns the method "FileSystemSyncAccessHandle.write".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemSyncAccessHandle) WriteFunc() (fn js.Func[func(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleWriteFunc(
			this.Ref(),
		),
	)
}

// Write1 calls the method "FileSystemSyncAccessHandle.write".
//
// The returned bool will be false if there is no such method.
func (this FileSystemSyncAccessHandle) Write1(buffer AllowSharedBufferSource) (uint64, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemSyncAccessHandleWrite1(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
	)

	return uint64(_ret), _ok
}

// Write1Func returns the method "FileSystemSyncAccessHandle.write".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemSyncAccessHandle) Write1Func() (fn js.Func[func(buffer AllowSharedBufferSource) uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleWrite1Func(
			this.Ref(),
		),
	)
}

// Truncate calls the method "FileSystemSyncAccessHandle.truncate".
//
// The returned bool will be false if there is no such method.
func (this FileSystemSyncAccessHandle) Truncate(newSize uint64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemSyncAccessHandleTruncate(
		this.Ref(), js.Pointer(&_ok),
		float64(newSize),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TruncateFunc returns the method "FileSystemSyncAccessHandle.truncate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemSyncAccessHandle) TruncateFunc() (fn js.Func[func(newSize uint64)]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleTruncateFunc(
			this.Ref(),
		),
	)
}

// GetSize calls the method "FileSystemSyncAccessHandle.getSize".
//
// The returned bool will be false if there is no such method.
func (this FileSystemSyncAccessHandle) GetSize() (uint64, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemSyncAccessHandleGetSize(
		this.Ref(), js.Pointer(&_ok),
	)

	return uint64(_ret), _ok
}

// GetSizeFunc returns the method "FileSystemSyncAccessHandle.getSize".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemSyncAccessHandle) GetSizeFunc() (fn js.Func[func() uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleGetSizeFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "FileSystemSyncAccessHandle.flush".
//
// The returned bool will be false if there is no such method.
func (this FileSystemSyncAccessHandle) Flush() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemSyncAccessHandleFlush(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FlushFunc returns the method "FileSystemSyncAccessHandle.flush".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemSyncAccessHandle) FlushFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleFlushFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "FileSystemSyncAccessHandle.close".
//
// The returned bool will be false if there is no such method.
func (this FileSystemSyncAccessHandle) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemSyncAccessHandleClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "FileSystemSyncAccessHandle.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemSyncAccessHandle) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleCloseFunc(
			this.Ref(),
		),
	)
}

type FileSystemFileHandle struct {
	FileSystemHandle
}

func (this FileSystemFileHandle) Once() FileSystemFileHandle {
	this.Ref().Once()
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
	this.Ref().Free()
}

// GetFile calls the method "FileSystemFileHandle.getFile".
//
// The returned bool will be false if there is no such method.
func (this FileSystemFileHandle) GetFile() (js.Promise[File], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemFileHandleGetFile(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[File]{}.FromRef(_ret), _ok
}

// GetFileFunc returns the method "FileSystemFileHandle.getFile".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemFileHandle) GetFileFunc() (fn js.Func[func() js.Promise[File]]) {
	return fn.FromRef(
		bindings.FileSystemFileHandleGetFileFunc(
			this.Ref(),
		),
	)
}

// CreateWritable calls the method "FileSystemFileHandle.createWritable".
//
// The returned bool will be false if there is no such method.
func (this FileSystemFileHandle) CreateWritable(options FileSystemCreateWritableOptions) (js.Promise[FileSystemWritableFileStream], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemFileHandleCreateWritable(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[FileSystemWritableFileStream]{}.FromRef(_ret), _ok
}

// CreateWritableFunc returns the method "FileSystemFileHandle.createWritable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemFileHandle) CreateWritableFunc() (fn js.Func[func(options FileSystemCreateWritableOptions) js.Promise[FileSystemWritableFileStream]]) {
	return fn.FromRef(
		bindings.FileSystemFileHandleCreateWritableFunc(
			this.Ref(),
		),
	)
}

// CreateWritable1 calls the method "FileSystemFileHandle.createWritable".
//
// The returned bool will be false if there is no such method.
func (this FileSystemFileHandle) CreateWritable1() (js.Promise[FileSystemWritableFileStream], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemFileHandleCreateWritable1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FileSystemWritableFileStream]{}.FromRef(_ret), _ok
}

// CreateWritable1Func returns the method "FileSystemFileHandle.createWritable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemFileHandle) CreateWritable1Func() (fn js.Func[func() js.Promise[FileSystemWritableFileStream]]) {
	return fn.FromRef(
		bindings.FileSystemFileHandleCreateWritable1Func(
			this.Ref(),
		),
	)
}

// CreateSyncAccessHandle calls the method "FileSystemFileHandle.createSyncAccessHandle".
//
// The returned bool will be false if there is no such method.
func (this FileSystemFileHandle) CreateSyncAccessHandle() (js.Promise[FileSystemSyncAccessHandle], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemFileHandleCreateSyncAccessHandle(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[FileSystemSyncAccessHandle]{}.FromRef(_ret), _ok
}

// CreateSyncAccessHandleFunc returns the method "FileSystemFileHandle.createSyncAccessHandle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemFileHandle) CreateSyncAccessHandleFunc() (fn js.Func[func() js.Promise[FileSystemSyncAccessHandle]]) {
	return fn.FromRef(
		bindings.FileSystemFileHandleCreateSyncAccessHandleFunc(
			this.Ref(),
		),
	)
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
func (p FilePickerAcceptType) UpdateFrom(ref js.Ref) {
	bindings.FilePickerAcceptTypeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FilePickerAcceptType) Update(ref js.Ref) {
	bindings.FilePickerAcceptTypeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p OpenFilePickerOptions) UpdateFrom(ref js.Ref) {
	bindings.OpenFilePickerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p OpenFilePickerOptions) Update(ref js.Ref) {
	bindings.OpenFilePickerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p SaveFilePickerOptions) UpdateFrom(ref js.Ref) {
	bindings.SaveFilePickerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SaveFilePickerOptions) Update(ref js.Ref) {
	bindings.SaveFilePickerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p FileSystemGetFileOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemGetFileOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FileSystemGetFileOptions) Update(ref js.Ref) {
	bindings.FileSystemGetFileOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p FileSystemGetDirectoryOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemGetDirectoryOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FileSystemGetDirectoryOptions) Update(ref js.Ref) {
	bindings.FileSystemGetDirectoryOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p FileSystemRemoveOptions) UpdateFrom(ref js.Ref) {
	bindings.FileSystemRemoveOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FileSystemRemoveOptions) Update(ref js.Ref) {
	bindings.FileSystemRemoveOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type FileSystemDirectoryHandle struct {
	FileSystemHandle
}

func (this FileSystemDirectoryHandle) Once() FileSystemDirectoryHandle {
	this.Ref().Once()
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
	this.Ref().Free()
}

// GetFileHandle calls the method "FileSystemDirectoryHandle.getFileHandle".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryHandle) GetFileHandle(name js.String, options FileSystemGetFileOptions) (js.Promise[FileSystemFileHandle], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryHandleGetFileHandle(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[FileSystemFileHandle]{}.FromRef(_ret), _ok
}

// GetFileHandleFunc returns the method "FileSystemDirectoryHandle.getFileHandle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryHandle) GetFileHandleFunc() (fn js.Func[func(name js.String, options FileSystemGetFileOptions) js.Promise[FileSystemFileHandle]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleGetFileHandleFunc(
			this.Ref(),
		),
	)
}

// GetFileHandle1 calls the method "FileSystemDirectoryHandle.getFileHandle".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryHandle) GetFileHandle1(name js.String) (js.Promise[FileSystemFileHandle], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryHandleGetFileHandle1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[FileSystemFileHandle]{}.FromRef(_ret), _ok
}

// GetFileHandle1Func returns the method "FileSystemDirectoryHandle.getFileHandle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryHandle) GetFileHandle1Func() (fn js.Func[func(name js.String) js.Promise[FileSystemFileHandle]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleGetFileHandle1Func(
			this.Ref(),
		),
	)
}

// GetDirectoryHandle calls the method "FileSystemDirectoryHandle.getDirectoryHandle".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryHandle) GetDirectoryHandle(name js.String, options FileSystemGetDirectoryOptions) (js.Promise[FileSystemDirectoryHandle], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryHandleGetDirectoryHandle(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[FileSystemDirectoryHandle]{}.FromRef(_ret), _ok
}

// GetDirectoryHandleFunc returns the method "FileSystemDirectoryHandle.getDirectoryHandle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryHandle) GetDirectoryHandleFunc() (fn js.Func[func(name js.String, options FileSystemGetDirectoryOptions) js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleGetDirectoryHandleFunc(
			this.Ref(),
		),
	)
}

// GetDirectoryHandle1 calls the method "FileSystemDirectoryHandle.getDirectoryHandle".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryHandle) GetDirectoryHandle1(name js.String) (js.Promise[FileSystemDirectoryHandle], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryHandleGetDirectoryHandle1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[FileSystemDirectoryHandle]{}.FromRef(_ret), _ok
}

// GetDirectoryHandle1Func returns the method "FileSystemDirectoryHandle.getDirectoryHandle".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryHandle) GetDirectoryHandle1Func() (fn js.Func[func(name js.String) js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleGetDirectoryHandle1Func(
			this.Ref(),
		),
	)
}

// RemoveEntry calls the method "FileSystemDirectoryHandle.removeEntry".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryHandle) RemoveEntry(name js.String, options FileSystemRemoveOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryHandleRemoveEntry(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RemoveEntryFunc returns the method "FileSystemDirectoryHandle.removeEntry".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryHandle) RemoveEntryFunc() (fn js.Func[func(name js.String, options FileSystemRemoveOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleRemoveEntryFunc(
			this.Ref(),
		),
	)
}

// RemoveEntry1 calls the method "FileSystemDirectoryHandle.removeEntry".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryHandle) RemoveEntry1(name js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryHandleRemoveEntry1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RemoveEntry1Func returns the method "FileSystemDirectoryHandle.removeEntry".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryHandle) RemoveEntry1Func() (fn js.Func[func(name js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleRemoveEntry1Func(
			this.Ref(),
		),
	)
}

// Resolve calls the method "FileSystemDirectoryHandle.resolve".
//
// The returned bool will be false if there is no such method.
func (this FileSystemDirectoryHandle) Resolve(possibleDescendant FileSystemHandle) (js.Promise[js.Array[js.String]], bool) {
	var _ok bool
	_ret := bindings.CallFileSystemDirectoryHandleResolve(
		this.Ref(), js.Pointer(&_ok),
		possibleDescendant.Ref(),
	)

	return js.Promise[js.Array[js.String]]{}.FromRef(_ret), _ok
}

// ResolveFunc returns the method "FileSystemDirectoryHandle.resolve".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemDirectoryHandle) ResolveFunc() (fn js.Func[func(possibleDescendant FileSystemHandle) js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleResolveFunc(
			this.Ref(),
		),
	)
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
func (p DirectoryPickerOptions) UpdateFrom(ref js.Ref) {
	bindings.DirectoryPickerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DirectoryPickerOptions) Update(ref js.Ref) {
	bindings.DirectoryPickerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type FontData struct {
	ref js.Ref
}

func (this FontData) Once() FontData {
	this.Ref().Once()
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
	this.Ref().Free()
}

// PostscriptName returns the value of property "FontData.postscriptName".
//
// The returned bool will be false if there is no such property.
func (this FontData) PostscriptName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontDataPostscriptName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FullName returns the value of property "FontData.fullName".
//
// The returned bool will be false if there is no such property.
func (this FontData) FullName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontDataFullName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Family returns the value of property "FontData.family".
//
// The returned bool will be false if there is no such property.
func (this FontData) Family() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontDataFamily(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Style returns the value of property "FontData.style".
//
// The returned bool will be false if there is no such property.
func (this FontData) Style() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontDataStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Blob calls the method "FontData.blob".
//
// The returned bool will be false if there is no such method.
func (this FontData) Blob() (js.Promise[Blob], bool) {
	var _ok bool
	_ret := bindings.CallFontDataBlob(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[Blob]{}.FromRef(_ret), _ok
}

// BlobFunc returns the method "FontData.blob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FontData) BlobFunc() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.FontDataBlobFunc(
			this.Ref(),
		),
	)
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
func (p QueryOptions) UpdateFrom(ref js.Ref) {
	bindings.QueryOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p QueryOptions) Update(ref js.Ref) {
	bindings.QueryOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ScreenDetails struct {
	EventTarget
}

func (this ScreenDetails) Once() ScreenDetails {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Screens returns the value of property "ScreenDetails.screens".
//
// The returned bool will be false if there is no such property.
func (this ScreenDetails) Screens() (js.FrozenArray[ScreenDetailed], bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailsScreens(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[ScreenDetailed]{}.FromRef(_ret), _ok
}

// CurrentScreen returns the value of property "ScreenDetails.currentScreen".
//
// The returned bool will be false if there is no such property.
func (this ScreenDetails) CurrentScreen() (ScreenDetailed, bool) {
	var _ok bool
	_ret := bindings.GetScreenDetailsCurrentScreen(
		this.Ref(), js.Pointer(&_ok),
	)
	return ScreenDetailed{}.FromRef(_ret), _ok
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
func (p ItemDetails) UpdateFrom(ref js.Ref) {
	bindings.ItemDetailsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ItemDetails) Update(ref js.Ref) {
	bindings.ItemDetailsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p PurchaseDetails) UpdateFrom(ref js.Ref) {
	bindings.PurchaseDetailsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PurchaseDetails) Update(ref js.Ref) {
	bindings.PurchaseDetailsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type DigitalGoodsService struct {
	ref js.Ref
}

func (this DigitalGoodsService) Once() DigitalGoodsService {
	this.Ref().Once()
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
	this.Ref().Free()
}

// GetDetails calls the method "DigitalGoodsService.getDetails".
//
// The returned bool will be false if there is no such method.
func (this DigitalGoodsService) GetDetails(itemIds js.Array[js.String]) (js.Promise[js.Array[ItemDetails]], bool) {
	var _ok bool
	_ret := bindings.CallDigitalGoodsServiceGetDetails(
		this.Ref(), js.Pointer(&_ok),
		itemIds.Ref(),
	)

	return js.Promise[js.Array[ItemDetails]]{}.FromRef(_ret), _ok
}

// GetDetailsFunc returns the method "DigitalGoodsService.getDetails".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DigitalGoodsService) GetDetailsFunc() (fn js.Func[func(itemIds js.Array[js.String]) js.Promise[js.Array[ItemDetails]]]) {
	return fn.FromRef(
		bindings.DigitalGoodsServiceGetDetailsFunc(
			this.Ref(),
		),
	)
}

// ListPurchases calls the method "DigitalGoodsService.listPurchases".
//
// The returned bool will be false if there is no such method.
func (this DigitalGoodsService) ListPurchases() (js.Promise[js.Array[PurchaseDetails]], bool) {
	var _ok bool
	_ret := bindings.CallDigitalGoodsServiceListPurchases(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[PurchaseDetails]]{}.FromRef(_ret), _ok
}

// ListPurchasesFunc returns the method "DigitalGoodsService.listPurchases".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DigitalGoodsService) ListPurchasesFunc() (fn js.Func[func() js.Promise[js.Array[PurchaseDetails]]]) {
	return fn.FromRef(
		bindings.DigitalGoodsServiceListPurchasesFunc(
			this.Ref(),
		),
	)
}

// ListPurchaseHistory calls the method "DigitalGoodsService.listPurchaseHistory".
//
// The returned bool will be false if there is no such method.
func (this DigitalGoodsService) ListPurchaseHistory() (js.Promise[js.Array[PurchaseDetails]], bool) {
	var _ok bool
	_ret := bindings.CallDigitalGoodsServiceListPurchaseHistory(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[PurchaseDetails]]{}.FromRef(_ret), _ok
}

// ListPurchaseHistoryFunc returns the method "DigitalGoodsService.listPurchaseHistory".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DigitalGoodsService) ListPurchaseHistoryFunc() (fn js.Func[func() js.Promise[js.Array[PurchaseDetails]]]) {
	return fn.FromRef(
		bindings.DigitalGoodsServiceListPurchaseHistoryFunc(
			this.Ref(),
		),
	)
}

// Consume calls the method "DigitalGoodsService.consume".
//
// The returned bool will be false if there is no such method.
func (this DigitalGoodsService) Consume(purchaseToken js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallDigitalGoodsServiceConsume(
		this.Ref(), js.Pointer(&_ok),
		purchaseToken.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ConsumeFunc returns the method "DigitalGoodsService.consume".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this DigitalGoodsService) ConsumeFunc() (fn js.Func[func(purchaseToken js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DigitalGoodsServiceConsumeFunc(
			this.Ref(),
		),
	)
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
		assert.Throw("invalid", "callback", "invocation")
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
func (p ImageBitmapOptions) UpdateFrom(ref js.Ref) {
	bindings.ImageBitmapOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ImageBitmapOptions) Update(ref js.Ref) {
	bindings.ImageBitmapOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
		assert.Throw("invalid", "callback", "invocation")
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Length returns the value of property "History.length".
//
// The returned bool will be false if there is no such property.
func (this History) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHistoryLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ScrollRestoration returns the value of property "History.scrollRestoration".
//
// The returned bool will be false if there is no such property.
func (this History) ScrollRestoration() (ScrollRestoration, bool) {
	var _ok bool
	_ret := bindings.GetHistoryScrollRestoration(
		this.Ref(), js.Pointer(&_ok),
	)
	return ScrollRestoration(_ret), _ok
}

// ScrollRestoration sets the value of property "History.scrollRestoration" to val.
//
// It returns false if the property cannot be set.
func (this History) SetScrollRestoration(val ScrollRestoration) bool {
	return js.True == bindings.SetHistoryScrollRestoration(
		this.Ref(),
		uint32(val),
	)
}

// State returns the value of property "History.state".
//
// The returned bool will be false if there is no such property.
func (this History) State() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetHistoryState(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Go calls the method "History.go".
//
// The returned bool will be false if there is no such method.
func (this History) Go(delta int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHistoryGo(
		this.Ref(), js.Pointer(&_ok),
		int32(delta),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GoFunc returns the method "History.go".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this History) GoFunc() (fn js.Func[func(delta int32)]) {
	return fn.FromRef(
		bindings.HistoryGoFunc(
			this.Ref(),
		),
	)
}

// Go1 calls the method "History.go".
//
// The returned bool will be false if there is no such method.
func (this History) Go1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHistoryGo1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Go1Func returns the method "History.go".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this History) Go1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HistoryGo1Func(
			this.Ref(),
		),
	)
}

// Back calls the method "History.back".
//
// The returned bool will be false if there is no such method.
func (this History) Back() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHistoryBack(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BackFunc returns the method "History.back".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this History) BackFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HistoryBackFunc(
			this.Ref(),
		),
	)
}

// Forward calls the method "History.forward".
//
// The returned bool will be false if there is no such method.
func (this History) Forward() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHistoryForward(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ForwardFunc returns the method "History.forward".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this History) ForwardFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HistoryForwardFunc(
			this.Ref(),
		),
	)
}

// PushState calls the method "History.pushState".
//
// The returned bool will be false if there is no such method.
func (this History) PushState(data js.Any, unused js.String, url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHistoryPushState(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PushStateFunc returns the method "History.pushState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this History) PushStateFunc() (fn js.Func[func(data js.Any, unused js.String, url js.String)]) {
	return fn.FromRef(
		bindings.HistoryPushStateFunc(
			this.Ref(),
		),
	)
}

// PushState1 calls the method "History.pushState".
//
// The returned bool will be false if there is no such method.
func (this History) PushState1(data js.Any, unused js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHistoryPushState1(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
		unused.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PushState1Func returns the method "History.pushState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this History) PushState1Func() (fn js.Func[func(data js.Any, unused js.String)]) {
	return fn.FromRef(
		bindings.HistoryPushState1Func(
			this.Ref(),
		),
	)
}

// ReplaceState calls the method "History.replaceState".
//
// The returned bool will be false if there is no such method.
func (this History) ReplaceState(data js.Any, unused js.String, url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHistoryReplaceState(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceStateFunc returns the method "History.replaceState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this History) ReplaceStateFunc() (fn js.Func[func(data js.Any, unused js.String, url js.String)]) {
	return fn.FromRef(
		bindings.HistoryReplaceStateFunc(
			this.Ref(),
		),
	)
}

// ReplaceState1 calls the method "History.replaceState".
//
// The returned bool will be false if there is no such method.
func (this History) ReplaceState1(data js.Any, unused js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHistoryReplaceState1(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
		unused.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReplaceState1Func returns the method "History.replaceState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this History) ReplaceState1Func() (fn js.Func[func(data js.Any, unused js.String)]) {
	return fn.FromRef(
		bindings.HistoryReplaceState1Func(
			this.Ref(),
		),
	)
}

type NavigationHistoryEntry struct {
	EventTarget
}

func (this NavigationHistoryEntry) Once() NavigationHistoryEntry {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Url returns the value of property "NavigationHistoryEntry.url".
//
// The returned bool will be false if there is no such property.
func (this NavigationHistoryEntry) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigationHistoryEntryUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Key returns the value of property "NavigationHistoryEntry.key".
//
// The returned bool will be false if there is no such property.
func (this NavigationHistoryEntry) Key() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigationHistoryEntryKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id returns the value of property "NavigationHistoryEntry.id".
//
// The returned bool will be false if there is no such property.
func (this NavigationHistoryEntry) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigationHistoryEntryId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Index returns the value of property "NavigationHistoryEntry.index".
//
// The returned bool will be false if there is no such property.
func (this NavigationHistoryEntry) Index() (int64, bool) {
	var _ok bool
	_ret := bindings.GetNavigationHistoryEntryIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int64(_ret), _ok
}

// SameDocument returns the value of property "NavigationHistoryEntry.sameDocument".
//
// The returned bool will be false if there is no such property.
func (this NavigationHistoryEntry) SameDocument() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigationHistoryEntrySameDocument(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// GetState calls the method "NavigationHistoryEntry.getState".
//
// The returned bool will be false if there is no such method.
func (this NavigationHistoryEntry) GetState() (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallNavigationHistoryEntryGetState(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetStateFunc returns the method "NavigationHistoryEntry.getState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigationHistoryEntry) GetStateFunc() (fn js.Func[func() js.Any]) {
	return fn.FromRef(
		bindings.NavigationHistoryEntryGetStateFunc(
			this.Ref(),
		),
	)
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
func (p NavigationUpdateCurrentEntryOptions) UpdateFrom(ref js.Ref) {
	bindings.NavigationUpdateCurrentEntryOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigationUpdateCurrentEntryOptions) Update(ref js.Ref) {
	bindings.NavigationUpdateCurrentEntryOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p NavigationResult) UpdateFrom(ref js.Ref) {
	bindings.NavigationResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigationResult) Update(ref js.Ref) {
	bindings.NavigationResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}
