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
// It returns ok=false if there is no such property.
func (this IdleDeadline) DidTimeout() (ret bool, ok bool) {
	ok = js.True == bindings.GetIdleDeadlineDidTimeout(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasTimeRemaining returns true if the method "IdleDeadline.timeRemaining" exists.
func (this IdleDeadline) HasTimeRemaining() bool {
	return js.True == bindings.HasIdleDeadlineTimeRemaining(
		this.Ref(),
	)
}

// TimeRemainingFunc returns the method "IdleDeadline.timeRemaining".
func (this IdleDeadline) TimeRemainingFunc() (fn js.Func[func() DOMHighResTimeStamp]) {
	return fn.FromRef(
		bindings.IdleDeadlineTimeRemainingFunc(
			this.Ref(),
		),
	)
}

// TimeRemaining calls the method "IdleDeadline.timeRemaining".
func (this IdleDeadline) TimeRemaining() (ret DOMHighResTimeStamp) {
	bindings.CallIdleDeadlineTimeRemaining(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTimeRemaining calls the method "IdleDeadline.timeRemaining"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IdleDeadline) TryTimeRemaining() (ret DOMHighResTimeStamp, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIdleDeadlineTimeRemaining(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

// HasWrite returns true if the method "FileSystemWritableFileStream.write" exists.
func (this FileSystemWritableFileStream) HasWrite() bool {
	return js.True == bindings.HasFileSystemWritableFileStreamWrite(
		this.Ref(),
	)
}

// WriteFunc returns the method "FileSystemWritableFileStream.write".
func (this FileSystemWritableFileStream) WriteFunc() (fn js.Func[func(data FileSystemWriteChunkType) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemWritableFileStreamWriteFunc(
			this.Ref(),
		),
	)
}

// Write calls the method "FileSystemWritableFileStream.write".
func (this FileSystemWritableFileStream) Write(data FileSystemWriteChunkType) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemWritableFileStreamWrite(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryWrite calls the method "FileSystemWritableFileStream.write"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemWritableFileStream) TryWrite(data FileSystemWriteChunkType) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemWritableFileStreamWrite(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasSeek returns true if the method "FileSystemWritableFileStream.seek" exists.
func (this FileSystemWritableFileStream) HasSeek() bool {
	return js.True == bindings.HasFileSystemWritableFileStreamSeek(
		this.Ref(),
	)
}

// SeekFunc returns the method "FileSystemWritableFileStream.seek".
func (this FileSystemWritableFileStream) SeekFunc() (fn js.Func[func(position uint64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemWritableFileStreamSeekFunc(
			this.Ref(),
		),
	)
}

// Seek calls the method "FileSystemWritableFileStream.seek".
func (this FileSystemWritableFileStream) Seek(position uint64) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemWritableFileStreamSeek(
		this.Ref(), js.Pointer(&ret),
		float64(position),
	)

	return
}

// TrySeek calls the method "FileSystemWritableFileStream.seek"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemWritableFileStream) TrySeek(position uint64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemWritableFileStreamSeek(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(position),
	)

	return
}

// HasTruncate returns true if the method "FileSystemWritableFileStream.truncate" exists.
func (this FileSystemWritableFileStream) HasTruncate() bool {
	return js.True == bindings.HasFileSystemWritableFileStreamTruncate(
		this.Ref(),
	)
}

// TruncateFunc returns the method "FileSystemWritableFileStream.truncate".
func (this FileSystemWritableFileStream) TruncateFunc() (fn js.Func[func(size uint64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemWritableFileStreamTruncateFunc(
			this.Ref(),
		),
	)
}

// Truncate calls the method "FileSystemWritableFileStream.truncate".
func (this FileSystemWritableFileStream) Truncate(size uint64) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemWritableFileStreamTruncate(
		this.Ref(), js.Pointer(&ret),
		float64(size),
	)

	return
}

// TryTruncate calls the method "FileSystemWritableFileStream.truncate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemWritableFileStream) TryTruncate(size uint64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemWritableFileStreamTruncate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

// HasRead returns true if the method "FileSystemSyncAccessHandle.read" exists.
func (this FileSystemSyncAccessHandle) HasRead() bool {
	return js.True == bindings.HasFileSystemSyncAccessHandleRead(
		this.Ref(),
	)
}

// ReadFunc returns the method "FileSystemSyncAccessHandle.read".
func (this FileSystemSyncAccessHandle) ReadFunc() (fn js.Func[func(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleReadFunc(
			this.Ref(),
		),
	)
}

// Read calls the method "FileSystemSyncAccessHandle.read".
func (this FileSystemSyncAccessHandle) Read(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleRead(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryRead calls the method "FileSystemSyncAccessHandle.read"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemSyncAccessHandle) TryRead(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleRead(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasRead1 returns true if the method "FileSystemSyncAccessHandle.read" exists.
func (this FileSystemSyncAccessHandle) HasRead1() bool {
	return js.True == bindings.HasFileSystemSyncAccessHandleRead1(
		this.Ref(),
	)
}

// Read1Func returns the method "FileSystemSyncAccessHandle.read".
func (this FileSystemSyncAccessHandle) Read1Func() (fn js.Func[func(buffer AllowSharedBufferSource) uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleRead1Func(
			this.Ref(),
		),
	)
}

// Read1 calls the method "FileSystemSyncAccessHandle.read".
func (this FileSystemSyncAccessHandle) Read1(buffer AllowSharedBufferSource) (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleRead1(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryRead1 calls the method "FileSystemSyncAccessHandle.read"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemSyncAccessHandle) TryRead1(buffer AllowSharedBufferSource) (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleRead1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasWrite returns true if the method "FileSystemSyncAccessHandle.write" exists.
func (this FileSystemSyncAccessHandle) HasWrite() bool {
	return js.True == bindings.HasFileSystemSyncAccessHandleWrite(
		this.Ref(),
	)
}

// WriteFunc returns the method "FileSystemSyncAccessHandle.write".
func (this FileSystemSyncAccessHandle) WriteFunc() (fn js.Func[func(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleWriteFunc(
			this.Ref(),
		),
	)
}

// Write calls the method "FileSystemSyncAccessHandle.write".
func (this FileSystemSyncAccessHandle) Write(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleWrite(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryWrite calls the method "FileSystemSyncAccessHandle.write"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemSyncAccessHandle) TryWrite(buffer AllowSharedBufferSource, options FileSystemReadWriteOptions) (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleWrite(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasWrite1 returns true if the method "FileSystemSyncAccessHandle.write" exists.
func (this FileSystemSyncAccessHandle) HasWrite1() bool {
	return js.True == bindings.HasFileSystemSyncAccessHandleWrite1(
		this.Ref(),
	)
}

// Write1Func returns the method "FileSystemSyncAccessHandle.write".
func (this FileSystemSyncAccessHandle) Write1Func() (fn js.Func[func(buffer AllowSharedBufferSource) uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleWrite1Func(
			this.Ref(),
		),
	)
}

// Write1 calls the method "FileSystemSyncAccessHandle.write".
func (this FileSystemSyncAccessHandle) Write1(buffer AllowSharedBufferSource) (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleWrite1(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryWrite1 calls the method "FileSystemSyncAccessHandle.write"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemSyncAccessHandle) TryWrite1(buffer AllowSharedBufferSource) (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleWrite1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasTruncate returns true if the method "FileSystemSyncAccessHandle.truncate" exists.
func (this FileSystemSyncAccessHandle) HasTruncate() bool {
	return js.True == bindings.HasFileSystemSyncAccessHandleTruncate(
		this.Ref(),
	)
}

// TruncateFunc returns the method "FileSystemSyncAccessHandle.truncate".
func (this FileSystemSyncAccessHandle) TruncateFunc() (fn js.Func[func(newSize uint64)]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleTruncateFunc(
			this.Ref(),
		),
	)
}

// Truncate calls the method "FileSystemSyncAccessHandle.truncate".
func (this FileSystemSyncAccessHandle) Truncate(newSize uint64) (ret js.Void) {
	bindings.CallFileSystemSyncAccessHandleTruncate(
		this.Ref(), js.Pointer(&ret),
		float64(newSize),
	)

	return
}

// TryTruncate calls the method "FileSystemSyncAccessHandle.truncate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemSyncAccessHandle) TryTruncate(newSize uint64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleTruncate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(newSize),
	)

	return
}

// HasGetSize returns true if the method "FileSystemSyncAccessHandle.getSize" exists.
func (this FileSystemSyncAccessHandle) HasGetSize() bool {
	return js.True == bindings.HasFileSystemSyncAccessHandleGetSize(
		this.Ref(),
	)
}

// GetSizeFunc returns the method "FileSystemSyncAccessHandle.getSize".
func (this FileSystemSyncAccessHandle) GetSizeFunc() (fn js.Func[func() uint64]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleGetSizeFunc(
			this.Ref(),
		),
	)
}

// GetSize calls the method "FileSystemSyncAccessHandle.getSize".
func (this FileSystemSyncAccessHandle) GetSize() (ret uint64) {
	bindings.CallFileSystemSyncAccessHandleGetSize(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSize calls the method "FileSystemSyncAccessHandle.getSize"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemSyncAccessHandle) TryGetSize() (ret uint64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleGetSize(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFlush returns true if the method "FileSystemSyncAccessHandle.flush" exists.
func (this FileSystemSyncAccessHandle) HasFlush() bool {
	return js.True == bindings.HasFileSystemSyncAccessHandleFlush(
		this.Ref(),
	)
}

// FlushFunc returns the method "FileSystemSyncAccessHandle.flush".
func (this FileSystemSyncAccessHandle) FlushFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleFlushFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "FileSystemSyncAccessHandle.flush".
func (this FileSystemSyncAccessHandle) Flush() (ret js.Void) {
	bindings.CallFileSystemSyncAccessHandleFlush(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "FileSystemSyncAccessHandle.flush"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemSyncAccessHandle) TryFlush() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleFlush(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "FileSystemSyncAccessHandle.close" exists.
func (this FileSystemSyncAccessHandle) HasClose() bool {
	return js.True == bindings.HasFileSystemSyncAccessHandleClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "FileSystemSyncAccessHandle.close".
func (this FileSystemSyncAccessHandle) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileSystemSyncAccessHandleCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "FileSystemSyncAccessHandle.close".
func (this FileSystemSyncAccessHandle) Close() (ret js.Void) {
	bindings.CallFileSystemSyncAccessHandleClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "FileSystemSyncAccessHandle.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemSyncAccessHandle) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemSyncAccessHandleClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasGetFile returns true if the method "FileSystemFileHandle.getFile" exists.
func (this FileSystemFileHandle) HasGetFile() bool {
	return js.True == bindings.HasFileSystemFileHandleGetFile(
		this.Ref(),
	)
}

// GetFileFunc returns the method "FileSystemFileHandle.getFile".
func (this FileSystemFileHandle) GetFileFunc() (fn js.Func[func() js.Promise[File]]) {
	return fn.FromRef(
		bindings.FileSystemFileHandleGetFileFunc(
			this.Ref(),
		),
	)
}

// GetFile calls the method "FileSystemFileHandle.getFile".
func (this FileSystemFileHandle) GetFile() (ret js.Promise[File]) {
	bindings.CallFileSystemFileHandleGetFile(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetFile calls the method "FileSystemFileHandle.getFile"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemFileHandle) TryGetFile() (ret js.Promise[File], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileHandleGetFile(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateWritable returns true if the method "FileSystemFileHandle.createWritable" exists.
func (this FileSystemFileHandle) HasCreateWritable() bool {
	return js.True == bindings.HasFileSystemFileHandleCreateWritable(
		this.Ref(),
	)
}

// CreateWritableFunc returns the method "FileSystemFileHandle.createWritable".
func (this FileSystemFileHandle) CreateWritableFunc() (fn js.Func[func(options FileSystemCreateWritableOptions) js.Promise[FileSystemWritableFileStream]]) {
	return fn.FromRef(
		bindings.FileSystemFileHandleCreateWritableFunc(
			this.Ref(),
		),
	)
}

// CreateWritable calls the method "FileSystemFileHandle.createWritable".
func (this FileSystemFileHandle) CreateWritable(options FileSystemCreateWritableOptions) (ret js.Promise[FileSystemWritableFileStream]) {
	bindings.CallFileSystemFileHandleCreateWritable(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryCreateWritable calls the method "FileSystemFileHandle.createWritable"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemFileHandle) TryCreateWritable(options FileSystemCreateWritableOptions) (ret js.Promise[FileSystemWritableFileStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileHandleCreateWritable(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasCreateWritable1 returns true if the method "FileSystemFileHandle.createWritable" exists.
func (this FileSystemFileHandle) HasCreateWritable1() bool {
	return js.True == bindings.HasFileSystemFileHandleCreateWritable1(
		this.Ref(),
	)
}

// CreateWritable1Func returns the method "FileSystemFileHandle.createWritable".
func (this FileSystemFileHandle) CreateWritable1Func() (fn js.Func[func() js.Promise[FileSystemWritableFileStream]]) {
	return fn.FromRef(
		bindings.FileSystemFileHandleCreateWritable1Func(
			this.Ref(),
		),
	)
}

// CreateWritable1 calls the method "FileSystemFileHandle.createWritable".
func (this FileSystemFileHandle) CreateWritable1() (ret js.Promise[FileSystemWritableFileStream]) {
	bindings.CallFileSystemFileHandleCreateWritable1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateWritable1 calls the method "FileSystemFileHandle.createWritable"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemFileHandle) TryCreateWritable1() (ret js.Promise[FileSystemWritableFileStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileHandleCreateWritable1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateSyncAccessHandle returns true if the method "FileSystemFileHandle.createSyncAccessHandle" exists.
func (this FileSystemFileHandle) HasCreateSyncAccessHandle() bool {
	return js.True == bindings.HasFileSystemFileHandleCreateSyncAccessHandle(
		this.Ref(),
	)
}

// CreateSyncAccessHandleFunc returns the method "FileSystemFileHandle.createSyncAccessHandle".
func (this FileSystemFileHandle) CreateSyncAccessHandleFunc() (fn js.Func[func() js.Promise[FileSystemSyncAccessHandle]]) {
	return fn.FromRef(
		bindings.FileSystemFileHandleCreateSyncAccessHandleFunc(
			this.Ref(),
		),
	)
}

// CreateSyncAccessHandle calls the method "FileSystemFileHandle.createSyncAccessHandle".
func (this FileSystemFileHandle) CreateSyncAccessHandle() (ret js.Promise[FileSystemSyncAccessHandle]) {
	bindings.CallFileSystemFileHandleCreateSyncAccessHandle(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSyncAccessHandle calls the method "FileSystemFileHandle.createSyncAccessHandle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemFileHandle) TryCreateSyncAccessHandle() (ret js.Promise[FileSystemSyncAccessHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileHandleCreateSyncAccessHandle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

// HasGetFileHandle returns true if the method "FileSystemDirectoryHandle.getFileHandle" exists.
func (this FileSystemDirectoryHandle) HasGetFileHandle() bool {
	return js.True == bindings.HasFileSystemDirectoryHandleGetFileHandle(
		this.Ref(),
	)
}

// GetFileHandleFunc returns the method "FileSystemDirectoryHandle.getFileHandle".
func (this FileSystemDirectoryHandle) GetFileHandleFunc() (fn js.Func[func(name js.String, options FileSystemGetFileOptions) js.Promise[FileSystemFileHandle]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleGetFileHandleFunc(
			this.Ref(),
		),
	)
}

// GetFileHandle calls the method "FileSystemDirectoryHandle.getFileHandle".
func (this FileSystemDirectoryHandle) GetFileHandle(name js.String, options FileSystemGetFileOptions) (ret js.Promise[FileSystemFileHandle]) {
	bindings.CallFileSystemDirectoryHandleGetFileHandle(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryGetFileHandle calls the method "FileSystemDirectoryHandle.getFileHandle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemDirectoryHandle) TryGetFileHandle(name js.String, options FileSystemGetFileOptions) (ret js.Promise[FileSystemFileHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleGetFileHandle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasGetFileHandle1 returns true if the method "FileSystemDirectoryHandle.getFileHandle" exists.
func (this FileSystemDirectoryHandle) HasGetFileHandle1() bool {
	return js.True == bindings.HasFileSystemDirectoryHandleGetFileHandle1(
		this.Ref(),
	)
}

// GetFileHandle1Func returns the method "FileSystemDirectoryHandle.getFileHandle".
func (this FileSystemDirectoryHandle) GetFileHandle1Func() (fn js.Func[func(name js.String) js.Promise[FileSystemFileHandle]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleGetFileHandle1Func(
			this.Ref(),
		),
	)
}

// GetFileHandle1 calls the method "FileSystemDirectoryHandle.getFileHandle".
func (this FileSystemDirectoryHandle) GetFileHandle1(name js.String) (ret js.Promise[FileSystemFileHandle]) {
	bindings.CallFileSystemDirectoryHandleGetFileHandle1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetFileHandle1 calls the method "FileSystemDirectoryHandle.getFileHandle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemDirectoryHandle) TryGetFileHandle1(name js.String) (ret js.Promise[FileSystemFileHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleGetFileHandle1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetDirectoryHandle returns true if the method "FileSystemDirectoryHandle.getDirectoryHandle" exists.
func (this FileSystemDirectoryHandle) HasGetDirectoryHandle() bool {
	return js.True == bindings.HasFileSystemDirectoryHandleGetDirectoryHandle(
		this.Ref(),
	)
}

// GetDirectoryHandleFunc returns the method "FileSystemDirectoryHandle.getDirectoryHandle".
func (this FileSystemDirectoryHandle) GetDirectoryHandleFunc() (fn js.Func[func(name js.String, options FileSystemGetDirectoryOptions) js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleGetDirectoryHandleFunc(
			this.Ref(),
		),
	)
}

// GetDirectoryHandle calls the method "FileSystemDirectoryHandle.getDirectoryHandle".
func (this FileSystemDirectoryHandle) GetDirectoryHandle(name js.String, options FileSystemGetDirectoryOptions) (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallFileSystemDirectoryHandleGetDirectoryHandle(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryGetDirectoryHandle calls the method "FileSystemDirectoryHandle.getDirectoryHandle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemDirectoryHandle) TryGetDirectoryHandle(name js.String, options FileSystemGetDirectoryOptions) (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleGetDirectoryHandle(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasGetDirectoryHandle1 returns true if the method "FileSystemDirectoryHandle.getDirectoryHandle" exists.
func (this FileSystemDirectoryHandle) HasGetDirectoryHandle1() bool {
	return js.True == bindings.HasFileSystemDirectoryHandleGetDirectoryHandle1(
		this.Ref(),
	)
}

// GetDirectoryHandle1Func returns the method "FileSystemDirectoryHandle.getDirectoryHandle".
func (this FileSystemDirectoryHandle) GetDirectoryHandle1Func() (fn js.Func[func(name js.String) js.Promise[FileSystemDirectoryHandle]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleGetDirectoryHandle1Func(
			this.Ref(),
		),
	)
}

// GetDirectoryHandle1 calls the method "FileSystemDirectoryHandle.getDirectoryHandle".
func (this FileSystemDirectoryHandle) GetDirectoryHandle1(name js.String) (ret js.Promise[FileSystemDirectoryHandle]) {
	bindings.CallFileSystemDirectoryHandleGetDirectoryHandle1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetDirectoryHandle1 calls the method "FileSystemDirectoryHandle.getDirectoryHandle"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemDirectoryHandle) TryGetDirectoryHandle1(name js.String) (ret js.Promise[FileSystemDirectoryHandle], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleGetDirectoryHandle1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasRemoveEntry returns true if the method "FileSystemDirectoryHandle.removeEntry" exists.
func (this FileSystemDirectoryHandle) HasRemoveEntry() bool {
	return js.True == bindings.HasFileSystemDirectoryHandleRemoveEntry(
		this.Ref(),
	)
}

// RemoveEntryFunc returns the method "FileSystemDirectoryHandle.removeEntry".
func (this FileSystemDirectoryHandle) RemoveEntryFunc() (fn js.Func[func(name js.String, options FileSystemRemoveOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleRemoveEntryFunc(
			this.Ref(),
		),
	)
}

// RemoveEntry calls the method "FileSystemDirectoryHandle.removeEntry".
func (this FileSystemDirectoryHandle) RemoveEntry(name js.String, options FileSystemRemoveOptions) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemDirectoryHandleRemoveEntry(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryRemoveEntry calls the method "FileSystemDirectoryHandle.removeEntry"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemDirectoryHandle) TryRemoveEntry(name js.String, options FileSystemRemoveOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleRemoveEntry(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasRemoveEntry1 returns true if the method "FileSystemDirectoryHandle.removeEntry" exists.
func (this FileSystemDirectoryHandle) HasRemoveEntry1() bool {
	return js.True == bindings.HasFileSystemDirectoryHandleRemoveEntry1(
		this.Ref(),
	)
}

// RemoveEntry1Func returns the method "FileSystemDirectoryHandle.removeEntry".
func (this FileSystemDirectoryHandle) RemoveEntry1Func() (fn js.Func[func(name js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleRemoveEntry1Func(
			this.Ref(),
		),
	)
}

// RemoveEntry1 calls the method "FileSystemDirectoryHandle.removeEntry".
func (this FileSystemDirectoryHandle) RemoveEntry1(name js.String) (ret js.Promise[js.Void]) {
	bindings.CallFileSystemDirectoryHandleRemoveEntry1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryRemoveEntry1 calls the method "FileSystemDirectoryHandle.removeEntry"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemDirectoryHandle) TryRemoveEntry1(name js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleRemoveEntry1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasResolve returns true if the method "FileSystemDirectoryHandle.resolve" exists.
func (this FileSystemDirectoryHandle) HasResolve() bool {
	return js.True == bindings.HasFileSystemDirectoryHandleResolve(
		this.Ref(),
	)
}

// ResolveFunc returns the method "FileSystemDirectoryHandle.resolve".
func (this FileSystemDirectoryHandle) ResolveFunc() (fn js.Func[func(possibleDescendant FileSystemHandle) js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.FileSystemDirectoryHandleResolveFunc(
			this.Ref(),
		),
	)
}

// Resolve calls the method "FileSystemDirectoryHandle.resolve".
func (this FileSystemDirectoryHandle) Resolve(possibleDescendant FileSystemHandle) (ret js.Promise[js.Array[js.String]]) {
	bindings.CallFileSystemDirectoryHandleResolve(
		this.Ref(), js.Pointer(&ret),
		possibleDescendant.Ref(),
	)

	return
}

// TryResolve calls the method "FileSystemDirectoryHandle.resolve"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FileSystemDirectoryHandle) TryResolve(possibleDescendant FileSystemHandle) (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemDirectoryHandleResolve(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
// It returns ok=false if there is no such property.
func (this FontData) PostscriptName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontDataPostscriptName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FullName returns the value of property "FontData.fullName".
//
// It returns ok=false if there is no such property.
func (this FontData) FullName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontDataFullName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Family returns the value of property "FontData.family".
//
// It returns ok=false if there is no such property.
func (this FontData) Family() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontDataFamily(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Style returns the value of property "FontData.style".
//
// It returns ok=false if there is no such property.
func (this FontData) Style() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontDataStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasBlob returns true if the method "FontData.blob" exists.
func (this FontData) HasBlob() bool {
	return js.True == bindings.HasFontDataBlob(
		this.Ref(),
	)
}

// BlobFunc returns the method "FontData.blob".
func (this FontData) BlobFunc() (fn js.Func[func() js.Promise[Blob]]) {
	return fn.FromRef(
		bindings.FontDataBlobFunc(
			this.Ref(),
		),
	)
}

// Blob calls the method "FontData.blob".
func (this FontData) Blob() (ret js.Promise[Blob]) {
	bindings.CallFontDataBlob(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBlob calls the method "FontData.blob"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this FontData) TryBlob() (ret js.Promise[Blob], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFontDataBlob(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
// It returns ok=false if there is no such property.
func (this ScreenDetails) Screens() (ret js.FrozenArray[ScreenDetailed], ok bool) {
	ok = js.True == bindings.GetScreenDetailsScreens(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentScreen returns the value of property "ScreenDetails.currentScreen".
//
// It returns ok=false if there is no such property.
func (this ScreenDetails) CurrentScreen() (ret ScreenDetailed, ok bool) {
	ok = js.True == bindings.GetScreenDetailsCurrentScreen(
		this.Ref(), js.Pointer(&ret),
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

// HasGetDetails returns true if the method "DigitalGoodsService.getDetails" exists.
func (this DigitalGoodsService) HasGetDetails() bool {
	return js.True == bindings.HasDigitalGoodsServiceGetDetails(
		this.Ref(),
	)
}

// GetDetailsFunc returns the method "DigitalGoodsService.getDetails".
func (this DigitalGoodsService) GetDetailsFunc() (fn js.Func[func(itemIds js.Array[js.String]) js.Promise[js.Array[ItemDetails]]]) {
	return fn.FromRef(
		bindings.DigitalGoodsServiceGetDetailsFunc(
			this.Ref(),
		),
	)
}

// GetDetails calls the method "DigitalGoodsService.getDetails".
func (this DigitalGoodsService) GetDetails(itemIds js.Array[js.String]) (ret js.Promise[js.Array[ItemDetails]]) {
	bindings.CallDigitalGoodsServiceGetDetails(
		this.Ref(), js.Pointer(&ret),
		itemIds.Ref(),
	)

	return
}

// TryGetDetails calls the method "DigitalGoodsService.getDetails"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DigitalGoodsService) TryGetDetails(itemIds js.Array[js.String]) (ret js.Promise[js.Array[ItemDetails]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDigitalGoodsServiceGetDetails(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		itemIds.Ref(),
	)

	return
}

// HasListPurchases returns true if the method "DigitalGoodsService.listPurchases" exists.
func (this DigitalGoodsService) HasListPurchases() bool {
	return js.True == bindings.HasDigitalGoodsServiceListPurchases(
		this.Ref(),
	)
}

// ListPurchasesFunc returns the method "DigitalGoodsService.listPurchases".
func (this DigitalGoodsService) ListPurchasesFunc() (fn js.Func[func() js.Promise[js.Array[PurchaseDetails]]]) {
	return fn.FromRef(
		bindings.DigitalGoodsServiceListPurchasesFunc(
			this.Ref(),
		),
	)
}

// ListPurchases calls the method "DigitalGoodsService.listPurchases".
func (this DigitalGoodsService) ListPurchases() (ret js.Promise[js.Array[PurchaseDetails]]) {
	bindings.CallDigitalGoodsServiceListPurchases(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryListPurchases calls the method "DigitalGoodsService.listPurchases"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DigitalGoodsService) TryListPurchases() (ret js.Promise[js.Array[PurchaseDetails]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDigitalGoodsServiceListPurchases(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasListPurchaseHistory returns true if the method "DigitalGoodsService.listPurchaseHistory" exists.
func (this DigitalGoodsService) HasListPurchaseHistory() bool {
	return js.True == bindings.HasDigitalGoodsServiceListPurchaseHistory(
		this.Ref(),
	)
}

// ListPurchaseHistoryFunc returns the method "DigitalGoodsService.listPurchaseHistory".
func (this DigitalGoodsService) ListPurchaseHistoryFunc() (fn js.Func[func() js.Promise[js.Array[PurchaseDetails]]]) {
	return fn.FromRef(
		bindings.DigitalGoodsServiceListPurchaseHistoryFunc(
			this.Ref(),
		),
	)
}

// ListPurchaseHistory calls the method "DigitalGoodsService.listPurchaseHistory".
func (this DigitalGoodsService) ListPurchaseHistory() (ret js.Promise[js.Array[PurchaseDetails]]) {
	bindings.CallDigitalGoodsServiceListPurchaseHistory(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryListPurchaseHistory calls the method "DigitalGoodsService.listPurchaseHistory"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DigitalGoodsService) TryListPurchaseHistory() (ret js.Promise[js.Array[PurchaseDetails]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDigitalGoodsServiceListPurchaseHistory(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasConsume returns true if the method "DigitalGoodsService.consume" exists.
func (this DigitalGoodsService) HasConsume() bool {
	return js.True == bindings.HasDigitalGoodsServiceConsume(
		this.Ref(),
	)
}

// ConsumeFunc returns the method "DigitalGoodsService.consume".
func (this DigitalGoodsService) ConsumeFunc() (fn js.Func[func(purchaseToken js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.DigitalGoodsServiceConsumeFunc(
			this.Ref(),
		),
	)
}

// Consume calls the method "DigitalGoodsService.consume".
func (this DigitalGoodsService) Consume(purchaseToken js.String) (ret js.Promise[js.Void]) {
	bindings.CallDigitalGoodsServiceConsume(
		this.Ref(), js.Pointer(&ret),
		purchaseToken.Ref(),
	)

	return
}

// TryConsume calls the method "DigitalGoodsService.consume"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this DigitalGoodsService) TryConsume(purchaseToken js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDigitalGoodsServiceConsume(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
// It returns ok=false if there is no such property.
func (this History) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHistoryLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScrollRestoration returns the value of property "History.scrollRestoration".
//
// It returns ok=false if there is no such property.
func (this History) ScrollRestoration() (ret ScrollRestoration, ok bool) {
	ok = js.True == bindings.GetHistoryScrollRestoration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetScrollRestoration sets the value of property "History.scrollRestoration" to val.
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
// It returns ok=false if there is no such property.
func (this History) State() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetHistoryState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGo returns true if the method "History.go" exists.
func (this History) HasGo() bool {
	return js.True == bindings.HasHistoryGo(
		this.Ref(),
	)
}

// GoFunc returns the method "History.go".
func (this History) GoFunc() (fn js.Func[func(delta int32)]) {
	return fn.FromRef(
		bindings.HistoryGoFunc(
			this.Ref(),
		),
	)
}

// Go calls the method "History.go".
func (this History) Go(delta int32) (ret js.Void) {
	bindings.CallHistoryGo(
		this.Ref(), js.Pointer(&ret),
		int32(delta),
	)

	return
}

// TryGo calls the method "History.go"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this History) TryGo(delta int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryGo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(delta),
	)

	return
}

// HasGo1 returns true if the method "History.go" exists.
func (this History) HasGo1() bool {
	return js.True == bindings.HasHistoryGo1(
		this.Ref(),
	)
}

// Go1Func returns the method "History.go".
func (this History) Go1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HistoryGo1Func(
			this.Ref(),
		),
	)
}

// Go1 calls the method "History.go".
func (this History) Go1() (ret js.Void) {
	bindings.CallHistoryGo1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGo1 calls the method "History.go"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this History) TryGo1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryGo1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasBack returns true if the method "History.back" exists.
func (this History) HasBack() bool {
	return js.True == bindings.HasHistoryBack(
		this.Ref(),
	)
}

// BackFunc returns the method "History.back".
func (this History) BackFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HistoryBackFunc(
			this.Ref(),
		),
	)
}

// Back calls the method "History.back".
func (this History) Back() (ret js.Void) {
	bindings.CallHistoryBack(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBack calls the method "History.back"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this History) TryBack() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryBack(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasForward returns true if the method "History.forward" exists.
func (this History) HasForward() bool {
	return js.True == bindings.HasHistoryForward(
		this.Ref(),
	)
}

// ForwardFunc returns the method "History.forward".
func (this History) ForwardFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HistoryForwardFunc(
			this.Ref(),
		),
	)
}

// Forward calls the method "History.forward".
func (this History) Forward() (ret js.Void) {
	bindings.CallHistoryForward(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryForward calls the method "History.forward"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this History) TryForward() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryForward(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPushState returns true if the method "History.pushState" exists.
func (this History) HasPushState() bool {
	return js.True == bindings.HasHistoryPushState(
		this.Ref(),
	)
}

// PushStateFunc returns the method "History.pushState".
func (this History) PushStateFunc() (fn js.Func[func(data js.Any, unused js.String, url js.String)]) {
	return fn.FromRef(
		bindings.HistoryPushStateFunc(
			this.Ref(),
		),
	)
}

// PushState calls the method "History.pushState".
func (this History) PushState(data js.Any, unused js.String, url js.String) (ret js.Void) {
	bindings.CallHistoryPushState(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	return
}

// TryPushState calls the method "History.pushState"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this History) TryPushState(data js.Any, unused js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryPushState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	return
}

// HasPushState1 returns true if the method "History.pushState" exists.
func (this History) HasPushState1() bool {
	return js.True == bindings.HasHistoryPushState1(
		this.Ref(),
	)
}

// PushState1Func returns the method "History.pushState".
func (this History) PushState1Func() (fn js.Func[func(data js.Any, unused js.String)]) {
	return fn.FromRef(
		bindings.HistoryPushState1Func(
			this.Ref(),
		),
	)
}

// PushState1 calls the method "History.pushState".
func (this History) PushState1(data js.Any, unused js.String) (ret js.Void) {
	bindings.CallHistoryPushState1(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
		unused.Ref(),
	)

	return
}

// TryPushState1 calls the method "History.pushState"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this History) TryPushState1(data js.Any, unused js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryPushState1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		unused.Ref(),
	)

	return
}

// HasReplaceState returns true if the method "History.replaceState" exists.
func (this History) HasReplaceState() bool {
	return js.True == bindings.HasHistoryReplaceState(
		this.Ref(),
	)
}

// ReplaceStateFunc returns the method "History.replaceState".
func (this History) ReplaceStateFunc() (fn js.Func[func(data js.Any, unused js.String, url js.String)]) {
	return fn.FromRef(
		bindings.HistoryReplaceStateFunc(
			this.Ref(),
		),
	)
}

// ReplaceState calls the method "History.replaceState".
func (this History) ReplaceState(data js.Any, unused js.String, url js.String) (ret js.Void) {
	bindings.CallHistoryReplaceState(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	return
}

// TryReplaceState calls the method "History.replaceState"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this History) TryReplaceState(data js.Any, unused js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryReplaceState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		unused.Ref(),
		url.Ref(),
	)

	return
}

// HasReplaceState1 returns true if the method "History.replaceState" exists.
func (this History) HasReplaceState1() bool {
	return js.True == bindings.HasHistoryReplaceState1(
		this.Ref(),
	)
}

// ReplaceState1Func returns the method "History.replaceState".
func (this History) ReplaceState1Func() (fn js.Func[func(data js.Any, unused js.String)]) {
	return fn.FromRef(
		bindings.HistoryReplaceState1Func(
			this.Ref(),
		),
	)
}

// ReplaceState1 calls the method "History.replaceState".
func (this History) ReplaceState1(data js.Any, unused js.String) (ret js.Void) {
	bindings.CallHistoryReplaceState1(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
		unused.Ref(),
	)

	return
}

// TryReplaceState1 calls the method "History.replaceState"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this History) TryReplaceState1(data js.Any, unused js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHistoryReplaceState1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		unused.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntryUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Key returns the value of property "NavigationHistoryEntry.key".
//
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) Key() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntryKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "NavigationHistoryEntry.id".
//
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntryId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Index returns the value of property "NavigationHistoryEntry.index".
//
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) Index() (ret int64, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntryIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SameDocument returns the value of property "NavigationHistoryEntry.sameDocument".
//
// It returns ok=false if there is no such property.
func (this NavigationHistoryEntry) SameDocument() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigationHistoryEntrySameDocument(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetState returns true if the method "NavigationHistoryEntry.getState" exists.
func (this NavigationHistoryEntry) HasGetState() bool {
	return js.True == bindings.HasNavigationHistoryEntryGetState(
		this.Ref(),
	)
}

// GetStateFunc returns the method "NavigationHistoryEntry.getState".
func (this NavigationHistoryEntry) GetStateFunc() (fn js.Func[func() js.Any]) {
	return fn.FromRef(
		bindings.NavigationHistoryEntryGetStateFunc(
			this.Ref(),
		),
	)
}

// GetState calls the method "NavigationHistoryEntry.getState".
func (this NavigationHistoryEntry) GetState() (ret js.Any) {
	bindings.CallNavigationHistoryEntryGetState(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetState calls the method "NavigationHistoryEntry.getState"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NavigationHistoryEntry) TryGetState() (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationHistoryEntryGetState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
