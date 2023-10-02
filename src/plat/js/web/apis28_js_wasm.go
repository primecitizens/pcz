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

type ServiceWorkerState uint32

const (
	_ ServiceWorkerState = iota

	ServiceWorkerState_PARSED
	ServiceWorkerState_INSTALLING
	ServiceWorkerState_INSTALLED
	ServiceWorkerState_ACTIVATING
	ServiceWorkerState_ACTIVATED
	ServiceWorkerState_REDUNDANT
)

func (ServiceWorkerState) FromRef(str js.Ref) ServiceWorkerState {
	return ServiceWorkerState(bindings.ConstOfServiceWorkerState(str))
}

func (x ServiceWorkerState) String() (string, bool) {
	switch x {
	case ServiceWorkerState_PARSED:
		return "parsed", true
	case ServiceWorkerState_INSTALLING:
		return "installing", true
	case ServiceWorkerState_INSTALLED:
		return "installed", true
	case ServiceWorkerState_ACTIVATING:
		return "activating", true
	case ServiceWorkerState_ACTIVATED:
		return "activated", true
	case ServiceWorkerState_REDUNDANT:
		return "redundant", true
	default:
		return "", false
	}
}

type ServiceWorker struct {
	EventTarget
}

func (this ServiceWorker) Once() ServiceWorker {
	this.Ref().Once()
	return this
}

func (this ServiceWorker) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this ServiceWorker) FromRef(ref js.Ref) ServiceWorker {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this ServiceWorker) Free() {
	this.Ref().Free()
}

// ScriptURL returns the value of property "ServiceWorker.scriptURL".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorker) ScriptURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerScriptURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// State returns the value of property "ServiceWorker.state".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorker) State() (ServiceWorkerState, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerState(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorkerState(_ret), _ok
}

// PostMessage calls the method "ServiceWorker.postMessage".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorker) PostMessage(message js.Any, transfer js.Array[js.Object]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerPostMessage(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		transfer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessageFunc returns the method "ServiceWorker.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorker) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.ServiceWorkerPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "ServiceWorker.postMessage".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorker) PostMessage1(message js.Any, options StructuredSerializeOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerPostMessage1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage1Func returns the method "ServiceWorker.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorker) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.ServiceWorkerPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "ServiceWorker.postMessage".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorker) PostMessage2(message js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerPostMessage2(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage2Func returns the method "ServiceWorker.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorker) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.ServiceWorkerPostMessage2Func(
			this.Ref(),
		),
	)
}

type NavigationPreloadState struct {
	// Enabled is "NavigationPreloadState.enabled"
	//
	// Optional, defaults to false.
	Enabled bool
	// HeaderValue is "NavigationPreloadState.headerValue"
	//
	// Optional
	HeaderValue js.String

	FFI_USE_Enabled bool // for Enabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigationPreloadState with all fields set.
func (p NavigationPreloadState) FromRef(ref js.Ref) NavigationPreloadState {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NavigationPreloadState NavigationPreloadState [// NavigationPreloadState] [0x140007befa0 0x140007bf0e0 0x140007bf040] 0x14000780c48 {0 0}} in the application heap.
func (p NavigationPreloadState) New() js.Ref {
	return bindings.NavigationPreloadStateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NavigationPreloadState) UpdateFrom(ref js.Ref) {
	bindings.NavigationPreloadStateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigationPreloadState) Update(ref js.Ref) {
	bindings.NavigationPreloadStateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NavigationPreloadManager struct {
	ref js.Ref
}

func (this NavigationPreloadManager) Once() NavigationPreloadManager {
	this.Ref().Once()
	return this
}

func (this NavigationPreloadManager) Ref() js.Ref {
	return this.ref
}

func (this NavigationPreloadManager) FromRef(ref js.Ref) NavigationPreloadManager {
	this.ref = ref
	return this
}

func (this NavigationPreloadManager) Free() {
	this.Ref().Free()
}

// Enable calls the method "NavigationPreloadManager.enable".
//
// The returned bool will be false if there is no such method.
func (this NavigationPreloadManager) Enable() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigationPreloadManagerEnable(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// EnableFunc returns the method "NavigationPreloadManager.enable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigationPreloadManager) EnableFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigationPreloadManagerEnableFunc(
			this.Ref(),
		),
	)
}

// Disable calls the method "NavigationPreloadManager.disable".
//
// The returned bool will be false if there is no such method.
func (this NavigationPreloadManager) Disable() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigationPreloadManagerDisable(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// DisableFunc returns the method "NavigationPreloadManager.disable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigationPreloadManager) DisableFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigationPreloadManagerDisableFunc(
			this.Ref(),
		),
	)
}

// SetHeaderValue calls the method "NavigationPreloadManager.setHeaderValue".
//
// The returned bool will be false if there is no such method.
func (this NavigationPreloadManager) SetHeaderValue(value js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNavigationPreloadManagerSetHeaderValue(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SetHeaderValueFunc returns the method "NavigationPreloadManager.setHeaderValue".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigationPreloadManager) SetHeaderValueFunc() (fn js.Func[func(value js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigationPreloadManagerSetHeaderValueFunc(
			this.Ref(),
		),
	)
}

// GetState calls the method "NavigationPreloadManager.getState".
//
// The returned bool will be false if there is no such method.
func (this NavigationPreloadManager) GetState() (js.Promise[NavigationPreloadState], bool) {
	var _ok bool
	_ret := bindings.CallNavigationPreloadManagerGetState(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[NavigationPreloadState]{}.FromRef(_ret), _ok
}

// GetStateFunc returns the method "NavigationPreloadManager.getState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigationPreloadManager) GetStateFunc() (fn js.Func[func() js.Promise[NavigationPreloadState]]) {
	return fn.FromRef(
		bindings.NavigationPreloadManagerGetStateFunc(
			this.Ref(),
		),
	)
}

type ServiceWorkerUpdateViaCache uint32

const (
	_ ServiceWorkerUpdateViaCache = iota

	ServiceWorkerUpdateViaCache_IMPORTS
	ServiceWorkerUpdateViaCache_ALL
	ServiceWorkerUpdateViaCache_NONE
)

func (ServiceWorkerUpdateViaCache) FromRef(str js.Ref) ServiceWorkerUpdateViaCache {
	return ServiceWorkerUpdateViaCache(bindings.ConstOfServiceWorkerUpdateViaCache(str))
}

func (x ServiceWorkerUpdateViaCache) String() (string, bool) {
	switch x {
	case ServiceWorkerUpdateViaCache_IMPORTS:
		return "imports", true
	case ServiceWorkerUpdateViaCache_ALL:
		return "all", true
	case ServiceWorkerUpdateViaCache_NONE:
		return "none", true
	default:
		return "", false
	}
}

type SyncManager struct {
	ref js.Ref
}

func (this SyncManager) Once() SyncManager {
	this.Ref().Once()
	return this
}

func (this SyncManager) Ref() js.Ref {
	return this.ref
}

func (this SyncManager) FromRef(ref js.Ref) SyncManager {
	this.ref = ref
	return this
}

func (this SyncManager) Free() {
	this.Ref().Free()
}

// Register calls the method "SyncManager.register".
//
// The returned bool will be false if there is no such method.
func (this SyncManager) Register(tag js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallSyncManagerRegister(
		this.Ref(), js.Pointer(&_ok),
		tag.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RegisterFunc returns the method "SyncManager.register".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SyncManager) RegisterFunc() (fn js.Func[func(tag js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SyncManagerRegisterFunc(
			this.Ref(),
		),
	)
}

// GetTags calls the method "SyncManager.getTags".
//
// The returned bool will be false if there is no such method.
func (this SyncManager) GetTags() (js.Promise[js.Array[js.String]], bool) {
	var _ok bool
	_ret := bindings.CallSyncManagerGetTags(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[js.String]]{}.FromRef(_ret), _ok
}

// GetTagsFunc returns the method "SyncManager.getTags".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SyncManager) GetTagsFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.SyncManagerGetTagsFunc(
			this.Ref(),
		),
	)
}

type PushEncryptionKeyName uint32

const (
	_ PushEncryptionKeyName = iota

	PushEncryptionKeyName_P_256DH
	PushEncryptionKeyName_AUTH
)

func (PushEncryptionKeyName) FromRef(str js.Ref) PushEncryptionKeyName {
	return PushEncryptionKeyName(bindings.ConstOfPushEncryptionKeyName(str))
}

func (x PushEncryptionKeyName) String() (string, bool) {
	switch x {
	case PushEncryptionKeyName_P_256DH:
		return "p256dh", true
	case PushEncryptionKeyName_AUTH:
		return "auth", true
	default:
		return "", false
	}
}

type PushSubscriptionJSON struct {
	// Endpoint is "PushSubscriptionJSON.endpoint"
	//
	// Optional
	Endpoint js.String
	// ExpirationTime is "PushSubscriptionJSON.expirationTime"
	//
	// Optional, defaults to null.
	ExpirationTime EpochTimeStamp
	// Keys is "PushSubscriptionJSON.keys"
	//
	// Optional
	Keys js.Record[js.String]

	FFI_USE_ExpirationTime bool // for ExpirationTime.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PushSubscriptionJSON with all fields set.
func (p PushSubscriptionJSON) FromRef(ref js.Ref) PushSubscriptionJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PushSubscriptionJSON PushSubscriptionJSON [// PushSubscriptionJSON] [0x1400083a000 0x1400083a0a0 0x1400083a1e0 0x1400083a140] 0x1400081e030 {0 0}} in the application heap.
func (p PushSubscriptionJSON) New() js.Ref {
	return bindings.PushSubscriptionJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PushSubscriptionJSON) UpdateFrom(ref js.Ref) {
	bindings.PushSubscriptionJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PushSubscriptionJSON) Update(ref js.Ref) {
	bindings.PushSubscriptionJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PushSubscriptionOptions struct {
	ref js.Ref
}

func (this PushSubscriptionOptions) Once() PushSubscriptionOptions {
	this.Ref().Once()
	return this
}

func (this PushSubscriptionOptions) Ref() js.Ref {
	return this.ref
}

func (this PushSubscriptionOptions) FromRef(ref js.Ref) PushSubscriptionOptions {
	this.ref = ref
	return this
}

func (this PushSubscriptionOptions) Free() {
	this.Ref().Free()
}

// UserVisibleOnly returns the value of property "PushSubscriptionOptions.userVisibleOnly".
//
// The returned bool will be false if there is no such property.
func (this PushSubscriptionOptions) UserVisibleOnly() (bool, bool) {
	var _ok bool
	_ret := bindings.GetPushSubscriptionOptionsUserVisibleOnly(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ApplicationServerKey returns the value of property "PushSubscriptionOptions.applicationServerKey".
//
// The returned bool will be false if there is no such property.
func (this PushSubscriptionOptions) ApplicationServerKey() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.GetPushSubscriptionOptionsApplicationServerKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

type PushSubscription struct {
	ref js.Ref
}

func (this PushSubscription) Once() PushSubscription {
	this.Ref().Once()
	return this
}

func (this PushSubscription) Ref() js.Ref {
	return this.ref
}

func (this PushSubscription) FromRef(ref js.Ref) PushSubscription {
	this.ref = ref
	return this
}

func (this PushSubscription) Free() {
	this.Ref().Free()
}

// Endpoint returns the value of property "PushSubscription.endpoint".
//
// The returned bool will be false if there is no such property.
func (this PushSubscription) Endpoint() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPushSubscriptionEndpoint(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ExpirationTime returns the value of property "PushSubscription.expirationTime".
//
// The returned bool will be false if there is no such property.
func (this PushSubscription) ExpirationTime() (EpochTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPushSubscriptionExpirationTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return EpochTimeStamp(_ret), _ok
}

// Options returns the value of property "PushSubscription.options".
//
// The returned bool will be false if there is no such property.
func (this PushSubscription) Options() (PushSubscriptionOptions, bool) {
	var _ok bool
	_ret := bindings.GetPushSubscriptionOptions(
		this.Ref(), js.Pointer(&_ok),
	)
	return PushSubscriptionOptions{}.FromRef(_ret), _ok
}

// GetKey calls the method "PushSubscription.getKey".
//
// The returned bool will be false if there is no such method.
func (this PushSubscription) GetKey(name PushEncryptionKeyName) (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.CallPushSubscriptionGetKey(
		this.Ref(), js.Pointer(&_ok),
		uint32(name),
	)

	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// GetKeyFunc returns the method "PushSubscription.getKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushSubscription) GetKeyFunc() (fn js.Func[func(name PushEncryptionKeyName) js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.PushSubscriptionGetKeyFunc(
			this.Ref(),
		),
	)
}

// Unsubscribe calls the method "PushSubscription.unsubscribe".
//
// The returned bool will be false if there is no such method.
func (this PushSubscription) Unsubscribe() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallPushSubscriptionUnsubscribe(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// UnsubscribeFunc returns the method "PushSubscription.unsubscribe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushSubscription) UnsubscribeFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.PushSubscriptionUnsubscribeFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PushSubscription.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PushSubscription) ToJSON() (PushSubscriptionJSON, bool) {
	var _ret PushSubscriptionJSON
	_ok := js.True == bindings.CallPushSubscriptionToJSON(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// ToJSONFunc returns the method "PushSubscription.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushSubscription) ToJSONFunc() (fn js.Func[func() PushSubscriptionJSON]) {
	return fn.FromRef(
		bindings.PushSubscriptionToJSONFunc(
			this.Ref(),
		),
	)
}

type OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) FromRef(ref js.Ref) OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String {
	return OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type PushSubscriptionOptionsInit struct {
	// UserVisibleOnly is "PushSubscriptionOptionsInit.userVisibleOnly"
	//
	// Optional, defaults to false.
	UserVisibleOnly bool
	// ApplicationServerKey is "PushSubscriptionOptionsInit.applicationServerKey"
	//
	// Optional, defaults to null.
	ApplicationServerKey OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String

	FFI_USE_UserVisibleOnly bool // for UserVisibleOnly.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PushSubscriptionOptionsInit with all fields set.
func (p PushSubscriptionOptionsInit) FromRef(ref js.Ref) PushSubscriptionOptionsInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PushSubscriptionOptionsInit PushSubscriptionOptionsInit [// PushSubscriptionOptionsInit] [0x1400083a3c0 0x1400083a500 0x1400083a460] 0x1400081e090 {0 0}} in the application heap.
func (p PushSubscriptionOptionsInit) New() js.Ref {
	return bindings.PushSubscriptionOptionsInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PushSubscriptionOptionsInit) UpdateFrom(ref js.Ref) {
	bindings.PushSubscriptionOptionsInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PushSubscriptionOptionsInit) Update(ref js.Ref) {
	bindings.PushSubscriptionOptionsInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PushManager struct {
	ref js.Ref
}

func (this PushManager) Once() PushManager {
	this.Ref().Once()
	return this
}

func (this PushManager) Ref() js.Ref {
	return this.ref
}

func (this PushManager) FromRef(ref js.Ref) PushManager {
	this.ref = ref
	return this
}

func (this PushManager) Free() {
	this.Ref().Free()
}

// SupportedContentEncodings returns the value of property "PushManager.supportedContentEncodings".
//
// The returned bool will be false if there is no such property.
func (this PushManager) SupportedContentEncodings() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetPushManagerSupportedContentEncodings(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// Subscribe calls the method "PushManager.subscribe".
//
// The returned bool will be false if there is no such method.
func (this PushManager) Subscribe(options PushSubscriptionOptionsInit) (js.Promise[PushSubscription], bool) {
	var _ok bool
	_ret := bindings.CallPushManagerSubscribe(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[PushSubscription]{}.FromRef(_ret), _ok
}

// SubscribeFunc returns the method "PushManager.subscribe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushManager) SubscribeFunc() (fn js.Func[func(options PushSubscriptionOptionsInit) js.Promise[PushSubscription]]) {
	return fn.FromRef(
		bindings.PushManagerSubscribeFunc(
			this.Ref(),
		),
	)
}

// Subscribe1 calls the method "PushManager.subscribe".
//
// The returned bool will be false if there is no such method.
func (this PushManager) Subscribe1() (js.Promise[PushSubscription], bool) {
	var _ok bool
	_ret := bindings.CallPushManagerSubscribe1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PushSubscription]{}.FromRef(_ret), _ok
}

// Subscribe1Func returns the method "PushManager.subscribe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushManager) Subscribe1Func() (fn js.Func[func() js.Promise[PushSubscription]]) {
	return fn.FromRef(
		bindings.PushManagerSubscribe1Func(
			this.Ref(),
		),
	)
}

// GetSubscription calls the method "PushManager.getSubscription".
//
// The returned bool will be false if there is no such method.
func (this PushManager) GetSubscription() (js.Promise[PushSubscription], bool) {
	var _ok bool
	_ret := bindings.CallPushManagerGetSubscription(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PushSubscription]{}.FromRef(_ret), _ok
}

// GetSubscriptionFunc returns the method "PushManager.getSubscription".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushManager) GetSubscriptionFunc() (fn js.Func[func() js.Promise[PushSubscription]]) {
	return fn.FromRef(
		bindings.PushManagerGetSubscriptionFunc(
			this.Ref(),
		),
	)
}

// PermissionState calls the method "PushManager.permissionState".
//
// The returned bool will be false if there is no such method.
func (this PushManager) PermissionState(options PushSubscriptionOptionsInit) (js.Promise[PermissionState], bool) {
	var _ok bool
	_ret := bindings.CallPushManagerPermissionState(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[PermissionState]{}.FromRef(_ret), _ok
}

// PermissionStateFunc returns the method "PushManager.permissionState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushManager) PermissionStateFunc() (fn js.Func[func(options PushSubscriptionOptionsInit) js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.PushManagerPermissionStateFunc(
			this.Ref(),
		),
	)
}

// PermissionState1 calls the method "PushManager.permissionState".
//
// The returned bool will be false if there is no such method.
func (this PushManager) PermissionState1() (js.Promise[PermissionState], bool) {
	var _ok bool
	_ret := bindings.CallPushManagerPermissionState1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PermissionState]{}.FromRef(_ret), _ok
}

// PermissionState1Func returns the method "PushManager.permissionState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushManager) PermissionState1Func() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.PushManagerPermissionState1Func(
			this.Ref(),
		),
	)
}

type PaymentDelegation uint32

const (
	_ PaymentDelegation = iota

	PaymentDelegation_SHIPPING_ADDRESS
	PaymentDelegation_PAYER_NAME
	PaymentDelegation_PAYER_PHONE
	PaymentDelegation_PAYER_EMAIL
)

func (PaymentDelegation) FromRef(str js.Ref) PaymentDelegation {
	return PaymentDelegation(bindings.ConstOfPaymentDelegation(str))
}

func (x PaymentDelegation) String() (string, bool) {
	switch x {
	case PaymentDelegation_SHIPPING_ADDRESS:
		return "shippingAddress", true
	case PaymentDelegation_PAYER_NAME:
		return "payerName", true
	case PaymentDelegation_PAYER_PHONE:
		return "payerPhone", true
	case PaymentDelegation_PAYER_EMAIL:
		return "payerEmail", true
	default:
		return "", false
	}
}

type PaymentManager struct {
	ref js.Ref
}

func (this PaymentManager) Once() PaymentManager {
	this.Ref().Once()
	return this
}

func (this PaymentManager) Ref() js.Ref {
	return this.ref
}

func (this PaymentManager) FromRef(ref js.Ref) PaymentManager {
	this.ref = ref
	return this
}

func (this PaymentManager) Free() {
	this.Ref().Free()
}

// UserHint returns the value of property "PaymentManager.userHint".
//
// The returned bool will be false if there is no such property.
func (this PaymentManager) UserHint() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaymentManagerUserHint(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UserHint sets the value of property "PaymentManager.userHint" to val.
//
// It returns false if the property cannot be set.
func (this PaymentManager) SetUserHint(val js.String) bool {
	return js.True == bindings.SetPaymentManagerUserHint(
		this.Ref(),
		val.Ref(),
	)
}

// EnableDelegations calls the method "PaymentManager.enableDelegations".
//
// The returned bool will be false if there is no such method.
func (this PaymentManager) EnableDelegations(delegations js.Array[PaymentDelegation]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPaymentManagerEnableDelegations(
		this.Ref(), js.Pointer(&_ok),
		delegations.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// EnableDelegationsFunc returns the method "PaymentManager.enableDelegations".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentManager) EnableDelegationsFunc() (fn js.Func[func(delegations js.Array[PaymentDelegation]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentManagerEnableDelegationsFunc(
			this.Ref(),
		),
	)
}

type PeriodicSyncManager struct {
	ref js.Ref
}

func (this PeriodicSyncManager) Once() PeriodicSyncManager {
	this.Ref().Once()
	return this
}

func (this PeriodicSyncManager) Ref() js.Ref {
	return this.ref
}

func (this PeriodicSyncManager) FromRef(ref js.Ref) PeriodicSyncManager {
	this.ref = ref
	return this
}

func (this PeriodicSyncManager) Free() {
	this.Ref().Free()
}

// Register calls the method "PeriodicSyncManager.register".
//
// The returned bool will be false if there is no such method.
func (this PeriodicSyncManager) Register(tag js.String, options BackgroundSyncOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPeriodicSyncManagerRegister(
		this.Ref(), js.Pointer(&_ok),
		tag.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RegisterFunc returns the method "PeriodicSyncManager.register".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PeriodicSyncManager) RegisterFunc() (fn js.Func[func(tag js.String, options BackgroundSyncOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PeriodicSyncManagerRegisterFunc(
			this.Ref(),
		),
	)
}

// Register1 calls the method "PeriodicSyncManager.register".
//
// The returned bool will be false if there is no such method.
func (this PeriodicSyncManager) Register1(tag js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPeriodicSyncManagerRegister1(
		this.Ref(), js.Pointer(&_ok),
		tag.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Register1Func returns the method "PeriodicSyncManager.register".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PeriodicSyncManager) Register1Func() (fn js.Func[func(tag js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PeriodicSyncManagerRegister1Func(
			this.Ref(),
		),
	)
}

// GetTags calls the method "PeriodicSyncManager.getTags".
//
// The returned bool will be false if there is no such method.
func (this PeriodicSyncManager) GetTags() (js.Promise[js.Array[js.String]], bool) {
	var _ok bool
	_ret := bindings.CallPeriodicSyncManagerGetTags(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[js.String]]{}.FromRef(_ret), _ok
}

// GetTagsFunc returns the method "PeriodicSyncManager.getTags".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PeriodicSyncManager) GetTagsFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.PeriodicSyncManagerGetTagsFunc(
			this.Ref(),
		),
	)
}

// Unregister calls the method "PeriodicSyncManager.unregister".
//
// The returned bool will be false if there is no such method.
func (this PeriodicSyncManager) Unregister(tag js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPeriodicSyncManagerUnregister(
		this.Ref(), js.Pointer(&_ok),
		tag.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// UnregisterFunc returns the method "PeriodicSyncManager.unregister".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PeriodicSyncManager) UnregisterFunc() (fn js.Func[func(tag js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PeriodicSyncManagerUnregisterFunc(
			this.Ref(),
		),
	)
}

type CookieStoreGetOptions struct {
	// Name is "CookieStoreGetOptions.name"
	//
	// Optional
	Name js.String
	// Url is "CookieStoreGetOptions.url"
	//
	// Optional
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CookieStoreGetOptions with all fields set.
func (p CookieStoreGetOptions) FromRef(ref js.Ref) CookieStoreGetOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 CookieStoreGetOptions CookieStoreGetOptions [// CookieStoreGetOptions] [0x1400083a5a0 0x1400083a640] 0x1400081e1c8 {0 0}} in the application heap.
func (p CookieStoreGetOptions) New() js.Ref {
	return bindings.CookieStoreGetOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CookieStoreGetOptions) UpdateFrom(ref js.Ref) {
	bindings.CookieStoreGetOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CookieStoreGetOptions) Update(ref js.Ref) {
	bindings.CookieStoreGetOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CookieStoreManager struct {
	ref js.Ref
}

func (this CookieStoreManager) Once() CookieStoreManager {
	this.Ref().Once()
	return this
}

func (this CookieStoreManager) Ref() js.Ref {
	return this.ref
}

func (this CookieStoreManager) FromRef(ref js.Ref) CookieStoreManager {
	this.ref = ref
	return this
}

func (this CookieStoreManager) Free() {
	this.Ref().Free()
}

// Subscribe calls the method "CookieStoreManager.subscribe".
//
// The returned bool will be false if there is no such method.
func (this CookieStoreManager) Subscribe(subscriptions js.Array[CookieStoreGetOptions]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreManagerSubscribe(
		this.Ref(), js.Pointer(&_ok),
		subscriptions.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SubscribeFunc returns the method "CookieStoreManager.subscribe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStoreManager) SubscribeFunc() (fn js.Func[func(subscriptions js.Array[CookieStoreGetOptions]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreManagerSubscribeFunc(
			this.Ref(),
		),
	)
}

// GetSubscriptions calls the method "CookieStoreManager.getSubscriptions".
//
// The returned bool will be false if there is no such method.
func (this CookieStoreManager) GetSubscriptions() (js.Promise[js.Array[CookieStoreGetOptions]], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreManagerGetSubscriptions(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[CookieStoreGetOptions]]{}.FromRef(_ret), _ok
}

// GetSubscriptionsFunc returns the method "CookieStoreManager.getSubscriptions".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStoreManager) GetSubscriptionsFunc() (fn js.Func[func() js.Promise[js.Array[CookieStoreGetOptions]]]) {
	return fn.FromRef(
		bindings.CookieStoreManagerGetSubscriptionsFunc(
			this.Ref(),
		),
	)
}

// Unsubscribe calls the method "CookieStoreManager.unsubscribe".
//
// The returned bool will be false if there is no such method.
func (this CookieStoreManager) Unsubscribe(subscriptions js.Array[CookieStoreGetOptions]) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallCookieStoreManagerUnsubscribe(
		this.Ref(), js.Pointer(&_ok),
		subscriptions.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// UnsubscribeFunc returns the method "CookieStoreManager.unsubscribe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CookieStoreManager) UnsubscribeFunc() (fn js.Func[func(subscriptions js.Array[CookieStoreGetOptions]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreManagerUnsubscribeFunc(
			this.Ref(),
		),
	)
}

type ContentCategory uint32

const (
	_ ContentCategory = iota

	ContentCategory_
	ContentCategory_HOMEPAGE
	ContentCategory_ARTICLE
	ContentCategory_VIDEO
	ContentCategory_AUDIO
)

func (ContentCategory) FromRef(str js.Ref) ContentCategory {
	return ContentCategory(bindings.ConstOfContentCategory(str))
}

func (x ContentCategory) String() (string, bool) {
	switch x {
	case ContentCategory_:
		return "", true
	case ContentCategory_HOMEPAGE:
		return "homepage", true
	case ContentCategory_ARTICLE:
		return "article", true
	case ContentCategory_VIDEO:
		return "video", true
	case ContentCategory_AUDIO:
		return "audio", true
	default:
		return "", false
	}
}

type ContentDescription struct {
	// Id is "ContentDescription.id"
	//
	// Required
	Id js.String
	// Title is "ContentDescription.title"
	//
	// Required
	Title js.String
	// Description is "ContentDescription.description"
	//
	// Required
	Description js.String
	// Category is "ContentDescription.category"
	//
	// Optional, defaults to "".
	Category ContentCategory
	// Icons is "ContentDescription.icons"
	//
	// Optional, defaults to [].
	Icons js.Array[ImageResource]
	// Url is "ContentDescription.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContentDescription with all fields set.
func (p ContentDescription) FromRef(ref js.Ref) ContentDescription {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ContentDescription ContentDescription [// ContentDescription] [0x1400083a6e0 0x1400083a780 0x1400083a820 0x1400083a8c0 0x1400083a960 0x1400083aa00] 0x1400081e240 {0 0}} in the application heap.
func (p ContentDescription) New() js.Ref {
	return bindings.ContentDescriptionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ContentDescription) UpdateFrom(ref js.Ref) {
	bindings.ContentDescriptionJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ContentDescription) Update(ref js.Ref) {
	bindings.ContentDescriptionJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ContentIndex struct {
	ref js.Ref
}

func (this ContentIndex) Once() ContentIndex {
	this.Ref().Once()
	return this
}

func (this ContentIndex) Ref() js.Ref {
	return this.ref
}

func (this ContentIndex) FromRef(ref js.Ref) ContentIndex {
	this.ref = ref
	return this
}

func (this ContentIndex) Free() {
	this.Ref().Free()
}

// Add calls the method "ContentIndex.add".
//
// The returned bool will be false if there is no such method.
func (this ContentIndex) Add(description ContentDescription) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallContentIndexAdd(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&description),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AddFunc returns the method "ContentIndex.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ContentIndex) AddFunc() (fn js.Func[func(description ContentDescription) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ContentIndexAddFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "ContentIndex.delete".
//
// The returned bool will be false if there is no such method.
func (this ContentIndex) Delete(id js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallContentIndexDelete(
		this.Ref(), js.Pointer(&_ok),
		id.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// DeleteFunc returns the method "ContentIndex.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ContentIndex) DeleteFunc() (fn js.Func[func(id js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ContentIndexDeleteFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "ContentIndex.getAll".
//
// The returned bool will be false if there is no such method.
func (this ContentIndex) GetAll() (js.Promise[js.Array[ContentDescription]], bool) {
	var _ok bool
	_ret := bindings.CallContentIndexGetAll(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[ContentDescription]]{}.FromRef(_ret), _ok
}

// GetAllFunc returns the method "ContentIndex.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ContentIndex) GetAllFunc() (fn js.Func[func() js.Promise[js.Array[ContentDescription]]]) {
	return fn.FromRef(
		bindings.ContentIndexGetAllFunc(
			this.Ref(),
		),
	)
}

type ServiceWorkerRegistration struct {
	EventTarget
}

func (this ServiceWorkerRegistration) Once() ServiceWorkerRegistration {
	this.Ref().Once()
	return this
}

func (this ServiceWorkerRegistration) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this ServiceWorkerRegistration) FromRef(ref js.Ref) ServiceWorkerRegistration {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this ServiceWorkerRegistration) Free() {
	this.Ref().Free()
}

// Installing returns the value of property "ServiceWorkerRegistration.installing".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) Installing() (ServiceWorker, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationInstalling(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorker{}.FromRef(_ret), _ok
}

// Waiting returns the value of property "ServiceWorkerRegistration.waiting".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) Waiting() (ServiceWorker, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationWaiting(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorker{}.FromRef(_ret), _ok
}

// Active returns the value of property "ServiceWorkerRegistration.active".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) Active() (ServiceWorker, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationActive(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorker{}.FromRef(_ret), _ok
}

// NavigationPreload returns the value of property "ServiceWorkerRegistration.navigationPreload".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) NavigationPreload() (NavigationPreloadManager, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationNavigationPreload(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationPreloadManager{}.FromRef(_ret), _ok
}

// Scope returns the value of property "ServiceWorkerRegistration.scope".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) Scope() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationScope(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UpdateViaCache returns the value of property "ServiceWorkerRegistration.updateViaCache".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) UpdateViaCache() (ServiceWorkerUpdateViaCache, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationUpdateViaCache(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorkerUpdateViaCache(_ret), _ok
}

// Sync returns the value of property "ServiceWorkerRegistration.sync".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) Sync() (SyncManager, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationSync(
		this.Ref(), js.Pointer(&_ok),
	)
	return SyncManager{}.FromRef(_ret), _ok
}

// BackgroundFetch returns the value of property "ServiceWorkerRegistration.backgroundFetch".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) BackgroundFetch() (BackgroundFetchManager, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationBackgroundFetch(
		this.Ref(), js.Pointer(&_ok),
	)
	return BackgroundFetchManager{}.FromRef(_ret), _ok
}

// PushManager returns the value of property "ServiceWorkerRegistration.pushManager".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) PushManager() (PushManager, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationPushManager(
		this.Ref(), js.Pointer(&_ok),
	)
	return PushManager{}.FromRef(_ret), _ok
}

// PaymentManager returns the value of property "ServiceWorkerRegistration.paymentManager".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) PaymentManager() (PaymentManager, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationPaymentManager(
		this.Ref(), js.Pointer(&_ok),
	)
	return PaymentManager{}.FromRef(_ret), _ok
}

// PeriodicSync returns the value of property "ServiceWorkerRegistration.periodicSync".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) PeriodicSync() (PeriodicSyncManager, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationPeriodicSync(
		this.Ref(), js.Pointer(&_ok),
	)
	return PeriodicSyncManager{}.FromRef(_ret), _ok
}

// Cookies returns the value of property "ServiceWorkerRegistration.cookies".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) Cookies() (CookieStoreManager, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationCookies(
		this.Ref(), js.Pointer(&_ok),
	)
	return CookieStoreManager{}.FromRef(_ret), _ok
}

// Index returns the value of property "ServiceWorkerRegistration.index".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerRegistration) Index() (ContentIndex, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerRegistrationIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return ContentIndex{}.FromRef(_ret), _ok
}

// Update calls the method "ServiceWorkerRegistration.update".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerRegistration) Update() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerRegistrationUpdate(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// UpdateFunc returns the method "ServiceWorkerRegistration.update".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerRegistration) UpdateFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationUpdateFunc(
			this.Ref(),
		),
	)
}

// Unregister calls the method "ServiceWorkerRegistration.unregister".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerRegistration) Unregister() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerRegistrationUnregister(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// UnregisterFunc returns the method "ServiceWorkerRegistration.unregister".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerRegistration) UnregisterFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationUnregisterFunc(
			this.Ref(),
		),
	)
}

// ShowNotification calls the method "ServiceWorkerRegistration.showNotification".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerRegistration) ShowNotification(title js.String, options NotificationOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerRegistrationShowNotification(
		this.Ref(), js.Pointer(&_ok),
		title.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ShowNotificationFunc returns the method "ServiceWorkerRegistration.showNotification".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerRegistration) ShowNotificationFunc() (fn js.Func[func(title js.String, options NotificationOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationShowNotificationFunc(
			this.Ref(),
		),
	)
}

// ShowNotification1 calls the method "ServiceWorkerRegistration.showNotification".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerRegistration) ShowNotification1(title js.String) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerRegistrationShowNotification1(
		this.Ref(), js.Pointer(&_ok),
		title.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ShowNotification1Func returns the method "ServiceWorkerRegistration.showNotification".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerRegistration) ShowNotification1Func() (fn js.Func[func(title js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationShowNotification1Func(
			this.Ref(),
		),
	)
}

// GetNotifications calls the method "ServiceWorkerRegistration.getNotifications".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerRegistration) GetNotifications(filter GetNotificationOptions) (js.Promise[js.Array[Notification]], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerRegistrationGetNotifications(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&filter),
	)

	return js.Promise[js.Array[Notification]]{}.FromRef(_ret), _ok
}

// GetNotificationsFunc returns the method "ServiceWorkerRegistration.getNotifications".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerRegistration) GetNotificationsFunc() (fn js.Func[func(filter GetNotificationOptions) js.Promise[js.Array[Notification]]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationGetNotificationsFunc(
			this.Ref(),
		),
	)
}

// GetNotifications1 calls the method "ServiceWorkerRegistration.getNotifications".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerRegistration) GetNotifications1() (js.Promise[js.Array[Notification]], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerRegistrationGetNotifications1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[Notification]]{}.FromRef(_ret), _ok
}

// GetNotifications1Func returns the method "ServiceWorkerRegistration.getNotifications".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerRegistration) GetNotifications1Func() (fn js.Func[func() js.Promise[js.Array[Notification]]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationGetNotifications1Func(
			this.Ref(),
		),
	)
}

type WorkerType uint32

const (
	_ WorkerType = iota

	WorkerType_CLASSIC
	WorkerType_MODULE
)

func (WorkerType) FromRef(str js.Ref) WorkerType {
	return WorkerType(bindings.ConstOfWorkerType(str))
}

func (x WorkerType) String() (string, bool) {
	switch x {
	case WorkerType_CLASSIC:
		return "classic", true
	case WorkerType_MODULE:
		return "module", true
	default:
		return "", false
	}
}

type RegistrationOptions struct {
	// Scope is "RegistrationOptions.scope"
	//
	// Optional
	Scope js.String
	// Type is "RegistrationOptions.type"
	//
	// Optional, defaults to "classic".
	Type WorkerType
	// UpdateViaCache is "RegistrationOptions.updateViaCache"
	//
	// Optional, defaults to "imports".
	UpdateViaCache ServiceWorkerUpdateViaCache

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RegistrationOptions with all fields set.
func (p RegistrationOptions) FromRef(ref js.Ref) RegistrationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RegistrationOptions RegistrationOptions [// RegistrationOptions] [0x1400083aaa0 0x1400083ab40 0x1400083abe0] 0x1400081e2b8 {0 0}} in the application heap.
func (p RegistrationOptions) New() js.Ref {
	return bindings.RegistrationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RegistrationOptions) UpdateFrom(ref js.Ref) {
	bindings.RegistrationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RegistrationOptions) Update(ref js.Ref) {
	bindings.RegistrationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_ServiceWorkerRegistration_undefined struct {
	ref js.Ref
}

func (x OneOf_ServiceWorkerRegistration_undefined) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ServiceWorkerRegistration_undefined) Free() {
	x.ref.Free()
}

func (x OneOf_ServiceWorkerRegistration_undefined) FromRef(ref js.Ref) OneOf_ServiceWorkerRegistration_undefined {
	return OneOf_ServiceWorkerRegistration_undefined{
		ref: ref,
	}
}

func (x OneOf_ServiceWorkerRegistration_undefined) Undefined() bool {
	return x.ref == js.Undefined
}

func (x OneOf_ServiceWorkerRegistration_undefined) ServiceWorkerRegistration() ServiceWorkerRegistration {
	return ServiceWorkerRegistration{}.FromRef(x.ref)
}

type ServiceWorkerContainer struct {
	EventTarget
}

func (this ServiceWorkerContainer) Once() ServiceWorkerContainer {
	this.Ref().Once()
	return this
}

func (this ServiceWorkerContainer) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this ServiceWorkerContainer) FromRef(ref js.Ref) ServiceWorkerContainer {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this ServiceWorkerContainer) Free() {
	this.Ref().Free()
}

// Controller returns the value of property "ServiceWorkerContainer.controller".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerContainer) Controller() (ServiceWorker, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerContainerController(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorker{}.FromRef(_ret), _ok
}

// Ready returns the value of property "ServiceWorkerContainer.ready".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerContainer) Ready() (js.Promise[ServiceWorkerRegistration], bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerContainerReady(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[ServiceWorkerRegistration]{}.FromRef(_ret), _ok
}

// Register calls the method "ServiceWorkerContainer.register".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerContainer) Register(scriptURL js.String, options RegistrationOptions) (js.Promise[ServiceWorkerRegistration], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerContainerRegister(
		this.Ref(), js.Pointer(&_ok),
		scriptURL.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[ServiceWorkerRegistration]{}.FromRef(_ret), _ok
}

// RegisterFunc returns the method "ServiceWorkerContainer.register".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerContainer) RegisterFunc() (fn js.Func[func(scriptURL js.String, options RegistrationOptions) js.Promise[ServiceWorkerRegistration]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerRegisterFunc(
			this.Ref(),
		),
	)
}

// Register1 calls the method "ServiceWorkerContainer.register".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerContainer) Register1(scriptURL js.String) (js.Promise[ServiceWorkerRegistration], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerContainerRegister1(
		this.Ref(), js.Pointer(&_ok),
		scriptURL.Ref(),
	)

	return js.Promise[ServiceWorkerRegistration]{}.FromRef(_ret), _ok
}

// Register1Func returns the method "ServiceWorkerContainer.register".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerContainer) Register1Func() (fn js.Func[func(scriptURL js.String) js.Promise[ServiceWorkerRegistration]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerRegister1Func(
			this.Ref(),
		),
	)
}

// GetRegistration calls the method "ServiceWorkerContainer.getRegistration".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerContainer) GetRegistration(clientURL js.String) (js.Promise[OneOf_ServiceWorkerRegistration_undefined], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerContainerGetRegistration(
		this.Ref(), js.Pointer(&_ok),
		clientURL.Ref(),
	)

	return js.Promise[OneOf_ServiceWorkerRegistration_undefined]{}.FromRef(_ret), _ok
}

// GetRegistrationFunc returns the method "ServiceWorkerContainer.getRegistration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerContainer) GetRegistrationFunc() (fn js.Func[func(clientURL js.String) js.Promise[OneOf_ServiceWorkerRegistration_undefined]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerGetRegistrationFunc(
			this.Ref(),
		),
	)
}

// GetRegistration1 calls the method "ServiceWorkerContainer.getRegistration".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerContainer) GetRegistration1() (js.Promise[OneOf_ServiceWorkerRegistration_undefined], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerContainerGetRegistration1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[OneOf_ServiceWorkerRegistration_undefined]{}.FromRef(_ret), _ok
}

// GetRegistration1Func returns the method "ServiceWorkerContainer.getRegistration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerContainer) GetRegistration1Func() (fn js.Func[func() js.Promise[OneOf_ServiceWorkerRegistration_undefined]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerGetRegistration1Func(
			this.Ref(),
		),
	)
}

// GetRegistrations calls the method "ServiceWorkerContainer.getRegistrations".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerContainer) GetRegistrations() (js.Promise[js.FrozenArray[ServiceWorkerRegistration]], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerContainerGetRegistrations(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.FrozenArray[ServiceWorkerRegistration]]{}.FromRef(_ret), _ok
}

// GetRegistrationsFunc returns the method "ServiceWorkerContainer.getRegistrations".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerContainer) GetRegistrationsFunc() (fn js.Func[func() js.Promise[js.FrozenArray[ServiceWorkerRegistration]]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerGetRegistrationsFunc(
			this.Ref(),
		),
	)
}

// StartMessages calls the method "ServiceWorkerContainer.startMessages".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerContainer) StartMessages() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerContainerStartMessages(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StartMessagesFunc returns the method "ServiceWorkerContainer.startMessages".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerContainer) StartMessagesFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerStartMessagesFunc(
			this.Ref(),
		),
	)
}

type MediaDeviceKind uint32

const (
	_ MediaDeviceKind = iota

	MediaDeviceKind_AUDIOINPUT
	MediaDeviceKind_AUDIOOUTPUT
	MediaDeviceKind_VIDEOINPUT
)

func (MediaDeviceKind) FromRef(str js.Ref) MediaDeviceKind {
	return MediaDeviceKind(bindings.ConstOfMediaDeviceKind(str))
}

func (x MediaDeviceKind) String() (string, bool) {
	switch x {
	case MediaDeviceKind_AUDIOINPUT:
		return "audioinput", true
	case MediaDeviceKind_AUDIOOUTPUT:
		return "audiooutput", true
	case MediaDeviceKind_VIDEOINPUT:
		return "videoinput", true
	default:
		return "", false
	}
}

type MediaDeviceInfo struct {
	ref js.Ref
}

func (this MediaDeviceInfo) Once() MediaDeviceInfo {
	this.Ref().Once()
	return this
}

func (this MediaDeviceInfo) Ref() js.Ref {
	return this.ref
}

func (this MediaDeviceInfo) FromRef(ref js.Ref) MediaDeviceInfo {
	this.ref = ref
	return this
}

func (this MediaDeviceInfo) Free() {
	this.Ref().Free()
}

// DeviceId returns the value of property "MediaDeviceInfo.deviceId".
//
// The returned bool will be false if there is no such property.
func (this MediaDeviceInfo) DeviceId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaDeviceInfoDeviceId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Kind returns the value of property "MediaDeviceInfo.kind".
//
// The returned bool will be false if there is no such property.
func (this MediaDeviceInfo) Kind() (MediaDeviceKind, bool) {
	var _ok bool
	_ret := bindings.GetMediaDeviceInfoKind(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaDeviceKind(_ret), _ok
}

// Label returns the value of property "MediaDeviceInfo.label".
//
// The returned bool will be false if there is no such property.
func (this MediaDeviceInfo) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaDeviceInfoLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GroupId returns the value of property "MediaDeviceInfo.groupId".
//
// The returned bool will be false if there is no such property.
func (this MediaDeviceInfo) GroupId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaDeviceInfoGroupId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ToJSON calls the method "MediaDeviceInfo.toJSON".
//
// The returned bool will be false if there is no such method.
func (this MediaDeviceInfo) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallMediaDeviceInfoToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "MediaDeviceInfo.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDeviceInfo) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.MediaDeviceInfoToJSONFunc(
			this.Ref(),
		),
	)
}

type SelfCapturePreferenceEnum uint32

const (
	_ SelfCapturePreferenceEnum = iota

	SelfCapturePreferenceEnum_INCLUDE
	SelfCapturePreferenceEnum_EXCLUDE
)

func (SelfCapturePreferenceEnum) FromRef(str js.Ref) SelfCapturePreferenceEnum {
	return SelfCapturePreferenceEnum(bindings.ConstOfSelfCapturePreferenceEnum(str))
}

func (x SelfCapturePreferenceEnum) String() (string, bool) {
	switch x {
	case SelfCapturePreferenceEnum_INCLUDE:
		return "include", true
	case SelfCapturePreferenceEnum_EXCLUDE:
		return "exclude", true
	default:
		return "", false
	}
}

type SystemAudioPreferenceEnum uint32

const (
	_ SystemAudioPreferenceEnum = iota

	SystemAudioPreferenceEnum_INCLUDE
	SystemAudioPreferenceEnum_EXCLUDE
)

func (SystemAudioPreferenceEnum) FromRef(str js.Ref) SystemAudioPreferenceEnum {
	return SystemAudioPreferenceEnum(bindings.ConstOfSystemAudioPreferenceEnum(str))
}

func (x SystemAudioPreferenceEnum) String() (string, bool) {
	switch x {
	case SystemAudioPreferenceEnum_INCLUDE:
		return "include", true
	case SystemAudioPreferenceEnum_EXCLUDE:
		return "exclude", true
	default:
		return "", false
	}
}

type SurfaceSwitchingPreferenceEnum uint32

const (
	_ SurfaceSwitchingPreferenceEnum = iota

	SurfaceSwitchingPreferenceEnum_INCLUDE
	SurfaceSwitchingPreferenceEnum_EXCLUDE
)

func (SurfaceSwitchingPreferenceEnum) FromRef(str js.Ref) SurfaceSwitchingPreferenceEnum {
	return SurfaceSwitchingPreferenceEnum(bindings.ConstOfSurfaceSwitchingPreferenceEnum(str))
}

func (x SurfaceSwitchingPreferenceEnum) String() (string, bool) {
	switch x {
	case SurfaceSwitchingPreferenceEnum_INCLUDE:
		return "include", true
	case SurfaceSwitchingPreferenceEnum_EXCLUDE:
		return "exclude", true
	default:
		return "", false
	}
}

type MonitorTypeSurfacesEnum uint32

const (
	_ MonitorTypeSurfacesEnum = iota

	MonitorTypeSurfacesEnum_INCLUDE
	MonitorTypeSurfacesEnum_EXCLUDE
)

func (MonitorTypeSurfacesEnum) FromRef(str js.Ref) MonitorTypeSurfacesEnum {
	return MonitorTypeSurfacesEnum(bindings.ConstOfMonitorTypeSurfacesEnum(str))
}

func (x MonitorTypeSurfacesEnum) String() (string, bool) {
	switch x {
	case MonitorTypeSurfacesEnum_INCLUDE:
		return "include", true
	case MonitorTypeSurfacesEnum_EXCLUDE:
		return "exclude", true
	default:
		return "", false
	}
}

type DisplayMediaStreamOptions struct {
	// Video is "DisplayMediaStreamOptions.video"
	//
	// Optional, defaults to true.
	Video OneOf_Bool_MediaTrackConstraints
	// Audio is "DisplayMediaStreamOptions.audio"
	//
	// Optional, defaults to false.
	Audio OneOf_Bool_MediaTrackConstraints
	// Controller is "DisplayMediaStreamOptions.controller"
	//
	// Optional
	Controller CaptureController
	// SelfBrowserSurface is "DisplayMediaStreamOptions.selfBrowserSurface"
	//
	// Optional
	SelfBrowserSurface SelfCapturePreferenceEnum
	// SystemAudio is "DisplayMediaStreamOptions.systemAudio"
	//
	// Optional
	SystemAudio SystemAudioPreferenceEnum
	// SurfaceSwitching is "DisplayMediaStreamOptions.surfaceSwitching"
	//
	// Optional
	SurfaceSwitching SurfaceSwitchingPreferenceEnum
	// MonitorTypeSurfaces is "DisplayMediaStreamOptions.monitorTypeSurfaces"
	//
	// Optional
	MonitorTypeSurfaces MonitorTypeSurfacesEnum

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DisplayMediaStreamOptions with all fields set.
func (p DisplayMediaStreamOptions) FromRef(ref js.Ref) DisplayMediaStreamOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 DisplayMediaStreamOptions DisplayMediaStreamOptions [// DisplayMediaStreamOptions] [0x1400083adc0 0x1400083ae60 0x1400083af00 0x1400083afa0 0x1400083b040 0x1400083b0e0 0x1400083b180] 0x1400081e318 {0 0}} in the application heap.
func (p DisplayMediaStreamOptions) New() js.Ref {
	return bindings.DisplayMediaStreamOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p DisplayMediaStreamOptions) UpdateFrom(ref js.Ref) {
	bindings.DisplayMediaStreamOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p DisplayMediaStreamOptions) Update(ref js.Ref) {
	bindings.DisplayMediaStreamOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ViewportMediaStreamConstraints struct {
	// Video is "ViewportMediaStreamConstraints.video"
	//
	// Optional, defaults to true.
	Video OneOf_Bool_MediaTrackConstraints
	// Audio is "ViewportMediaStreamConstraints.audio"
	//
	// Optional, defaults to false.
	Audio OneOf_Bool_MediaTrackConstraints

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ViewportMediaStreamConstraints with all fields set.
func (p ViewportMediaStreamConstraints) FromRef(ref js.Ref) ViewportMediaStreamConstraints {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ViewportMediaStreamConstraints ViewportMediaStreamConstraints [// ViewportMediaStreamConstraints] [0x1400083b220 0x1400083b2c0] 0x1400081e3a8 {0 0}} in the application heap.
func (p ViewportMediaStreamConstraints) New() js.Ref {
	return bindings.ViewportMediaStreamConstraintsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ViewportMediaStreamConstraints) UpdateFrom(ref js.Ref) {
	bindings.ViewportMediaStreamConstraintsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ViewportMediaStreamConstraints) Update(ref js.Ref) {
	bindings.ViewportMediaStreamConstraintsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaTrackSupportedConstraints struct {
	// Width is "MediaTrackSupportedConstraints.width"
	//
	// Optional, defaults to true.
	Width bool
	// Height is "MediaTrackSupportedConstraints.height"
	//
	// Optional, defaults to true.
	Height bool
	// AspectRatio is "MediaTrackSupportedConstraints.aspectRatio"
	//
	// Optional, defaults to true.
	AspectRatio bool
	// FrameRate is "MediaTrackSupportedConstraints.frameRate"
	//
	// Optional, defaults to true.
	FrameRate bool
	// FacingMode is "MediaTrackSupportedConstraints.facingMode"
	//
	// Optional, defaults to true.
	FacingMode bool
	// ResizeMode is "MediaTrackSupportedConstraints.resizeMode"
	//
	// Optional, defaults to true.
	ResizeMode bool
	// SampleRate is "MediaTrackSupportedConstraints.sampleRate"
	//
	// Optional, defaults to true.
	SampleRate bool
	// SampleSize is "MediaTrackSupportedConstraints.sampleSize"
	//
	// Optional, defaults to true.
	SampleSize bool
	// EchoCancellation is "MediaTrackSupportedConstraints.echoCancellation"
	//
	// Optional, defaults to true.
	EchoCancellation bool
	// AutoGainControl is "MediaTrackSupportedConstraints.autoGainControl"
	//
	// Optional, defaults to true.
	AutoGainControl bool
	// NoiseSuppression is "MediaTrackSupportedConstraints.noiseSuppression"
	//
	// Optional, defaults to true.
	NoiseSuppression bool
	// Latency is "MediaTrackSupportedConstraints.latency"
	//
	// Optional, defaults to true.
	Latency bool
	// ChannelCount is "MediaTrackSupportedConstraints.channelCount"
	//
	// Optional, defaults to true.
	ChannelCount bool
	// DeviceId is "MediaTrackSupportedConstraints.deviceId"
	//
	// Optional, defaults to true.
	DeviceId bool
	// GroupId is "MediaTrackSupportedConstraints.groupId"
	//
	// Optional, defaults to true.
	GroupId bool
	// WhiteBalanceMode is "MediaTrackSupportedConstraints.whiteBalanceMode"
	//
	// Optional, defaults to true.
	WhiteBalanceMode bool
	// ExposureMode is "MediaTrackSupportedConstraints.exposureMode"
	//
	// Optional, defaults to true.
	ExposureMode bool
	// FocusMode is "MediaTrackSupportedConstraints.focusMode"
	//
	// Optional, defaults to true.
	FocusMode bool
	// PointsOfInterest is "MediaTrackSupportedConstraints.pointsOfInterest"
	//
	// Optional, defaults to true.
	PointsOfInterest bool
	// ExposureCompensation is "MediaTrackSupportedConstraints.exposureCompensation"
	//
	// Optional, defaults to true.
	ExposureCompensation bool
	// ExposureTime is "MediaTrackSupportedConstraints.exposureTime"
	//
	// Optional, defaults to true.
	ExposureTime bool
	// ColorTemperature is "MediaTrackSupportedConstraints.colorTemperature"
	//
	// Optional, defaults to true.
	ColorTemperature bool
	// Iso is "MediaTrackSupportedConstraints.iso"
	//
	// Optional, defaults to true.
	Iso bool
	// Brightness is "MediaTrackSupportedConstraints.brightness"
	//
	// Optional, defaults to true.
	Brightness bool
	// Contrast is "MediaTrackSupportedConstraints.contrast"
	//
	// Optional, defaults to true.
	Contrast bool
	// Pan is "MediaTrackSupportedConstraints.pan"
	//
	// Optional, defaults to true.
	Pan bool
	// Saturation is "MediaTrackSupportedConstraints.saturation"
	//
	// Optional, defaults to true.
	Saturation bool
	// Sharpness is "MediaTrackSupportedConstraints.sharpness"
	//
	// Optional, defaults to true.
	Sharpness bool
	// FocusDistance is "MediaTrackSupportedConstraints.focusDistance"
	//
	// Optional, defaults to true.
	FocusDistance bool
	// Tilt is "MediaTrackSupportedConstraints.tilt"
	//
	// Optional, defaults to true.
	Tilt bool
	// Zoom is "MediaTrackSupportedConstraints.zoom"
	//
	// Optional, defaults to true.
	Zoom bool
	// Torch is "MediaTrackSupportedConstraints.torch"
	//
	// Optional, defaults to true.
	Torch bool
	// DisplaySurface is "MediaTrackSupportedConstraints.displaySurface"
	//
	// Optional, defaults to true.
	DisplaySurface bool
	// LogicalSurface is "MediaTrackSupportedConstraints.logicalSurface"
	//
	// Optional, defaults to true.
	LogicalSurface bool
	// Cursor is "MediaTrackSupportedConstraints.cursor"
	//
	// Optional, defaults to true.
	Cursor bool
	// RestrictOwnAudio is "MediaTrackSupportedConstraints.restrictOwnAudio"
	//
	// Optional, defaults to true.
	RestrictOwnAudio bool
	// SuppressLocalAudioPlayback is "MediaTrackSupportedConstraints.suppressLocalAudioPlayback"
	//
	// Optional, defaults to true.
	SuppressLocalAudioPlayback bool

	FFI_USE_Width                      bool // for Width.
	FFI_USE_Height                     bool // for Height.
	FFI_USE_AspectRatio                bool // for AspectRatio.
	FFI_USE_FrameRate                  bool // for FrameRate.
	FFI_USE_FacingMode                 bool // for FacingMode.
	FFI_USE_ResizeMode                 bool // for ResizeMode.
	FFI_USE_SampleRate                 bool // for SampleRate.
	FFI_USE_SampleSize                 bool // for SampleSize.
	FFI_USE_EchoCancellation           bool // for EchoCancellation.
	FFI_USE_AutoGainControl            bool // for AutoGainControl.
	FFI_USE_NoiseSuppression           bool // for NoiseSuppression.
	FFI_USE_Latency                    bool // for Latency.
	FFI_USE_ChannelCount               bool // for ChannelCount.
	FFI_USE_DeviceId                   bool // for DeviceId.
	FFI_USE_GroupId                    bool // for GroupId.
	FFI_USE_WhiteBalanceMode           bool // for WhiteBalanceMode.
	FFI_USE_ExposureMode               bool // for ExposureMode.
	FFI_USE_FocusMode                  bool // for FocusMode.
	FFI_USE_PointsOfInterest           bool // for PointsOfInterest.
	FFI_USE_ExposureCompensation       bool // for ExposureCompensation.
	FFI_USE_ExposureTime               bool // for ExposureTime.
	FFI_USE_ColorTemperature           bool // for ColorTemperature.
	FFI_USE_Iso                        bool // for Iso.
	FFI_USE_Brightness                 bool // for Brightness.
	FFI_USE_Contrast                   bool // for Contrast.
	FFI_USE_Pan                        bool // for Pan.
	FFI_USE_Saturation                 bool // for Saturation.
	FFI_USE_Sharpness                  bool // for Sharpness.
	FFI_USE_FocusDistance              bool // for FocusDistance.
	FFI_USE_Tilt                       bool // for Tilt.
	FFI_USE_Zoom                       bool // for Zoom.
	FFI_USE_Torch                      bool // for Torch.
	FFI_USE_DisplaySurface             bool // for DisplaySurface.
	FFI_USE_LogicalSurface             bool // for LogicalSurface.
	FFI_USE_Cursor                     bool // for Cursor.
	FFI_USE_RestrictOwnAudio           bool // for RestrictOwnAudio.
	FFI_USE_SuppressLocalAudioPlayback bool // for SuppressLocalAudioPlayback.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaTrackSupportedConstraints with all fields set.
func (p MediaTrackSupportedConstraints) FromRef(ref js.Ref) MediaTrackSupportedConstraints {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MediaTrackSupportedConstraints MediaTrackSupportedConstraints [// MediaTrackSupportedConstraints] [0x1400083b360 0x1400083b4a0 0x1400083b5e0 0x1400083b720 0x1400083b860 0x1400083b9a0 0x1400083bae0 0x1400083bc20 0x1400083bd60 0x1400083bea0 0x140007bf180 0x140007bf2c0 0x140007bf400 0x140007bf540 0x140007bf680 0x140007bf7c0 0x140007bf900 0x140007bfa40 0x140007bfb80 0x140007bfcc0 0x140007bfe00 0x140007bff40 0x14000d010e0 0x14000d13720 0x14000d13900 0x14000d13d60 0x14000be8000 0x14000be8140 0x14000be8b40 0x14000ca21e0 0x14000ca2500 0x14000ca2640 0x14000ca2be0 0x14000ca2d20 0x14000ca32c0 0x14000ca34a0 0x14000d181e0 0x1400083b400 0x1400083b540 0x1400083b680 0x1400083b7c0 0x1400083b900 0x1400083ba40 0x1400083bb80 0x1400083bcc0 0x1400083be00 0x1400083bf40 0x140007bf220 0x140007bf360 0x140007bf4a0 0x140007bf5e0 0x140007bf720 0x140007bf860 0x140007bf9a0 0x140007bfae0 0x140007bfc20 0x140007bfd60 0x140007bfea0 0x14000d00c80 0x14000d01180 0x14000d137c0 0x14000d139a0 0x14000d13e00 0x14000be80a0 0x14000be83c0 0x14000be92c0 0x14000ca2460 0x14000ca25a0 0x14000ca2aa0 0x14000ca2c80 0x14000ca3220 0x14000ca3360 0x14000ca3540 0x14000d197c0] 0x1400081e3d8 {0 0}} in the application heap.
func (p MediaTrackSupportedConstraints) New() js.Ref {
	return bindings.MediaTrackSupportedConstraintsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaTrackSupportedConstraints) UpdateFrom(ref js.Ref) {
	bindings.MediaTrackSupportedConstraintsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaTrackSupportedConstraints) Update(ref js.Ref) {
	bindings.MediaTrackSupportedConstraintsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaDevices struct {
	EventTarget
}

func (this MediaDevices) Once() MediaDevices {
	this.Ref().Once()
	return this
}

func (this MediaDevices) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MediaDevices) FromRef(ref js.Ref) MediaDevices {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MediaDevices) Free() {
	this.Ref().Free()
}

// EnumerateDevices calls the method "MediaDevices.enumerateDevices".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) EnumerateDevices() (js.Promise[js.Array[MediaDeviceInfo]], bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesEnumerateDevices(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Array[MediaDeviceInfo]]{}.FromRef(_ret), _ok
}

// EnumerateDevicesFunc returns the method "MediaDevices.enumerateDevices".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) EnumerateDevicesFunc() (fn js.Func[func() js.Promise[js.Array[MediaDeviceInfo]]]) {
	return fn.FromRef(
		bindings.MediaDevicesEnumerateDevicesFunc(
			this.Ref(),
		),
	)
}

// GetDisplayMedia calls the method "MediaDevices.getDisplayMedia".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) GetDisplayMedia(options DisplayMediaStreamOptions) (js.Promise[MediaStream], bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesGetDisplayMedia(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[MediaStream]{}.FromRef(_ret), _ok
}

// GetDisplayMediaFunc returns the method "MediaDevices.getDisplayMedia".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) GetDisplayMediaFunc() (fn js.Func[func(options DisplayMediaStreamOptions) js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetDisplayMediaFunc(
			this.Ref(),
		),
	)
}

// GetDisplayMedia1 calls the method "MediaDevices.getDisplayMedia".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) GetDisplayMedia1() (js.Promise[MediaStream], bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesGetDisplayMedia1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MediaStream]{}.FromRef(_ret), _ok
}

// GetDisplayMedia1Func returns the method "MediaDevices.getDisplayMedia".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) GetDisplayMedia1Func() (fn js.Func[func() js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetDisplayMedia1Func(
			this.Ref(),
		),
	)
}

// SetSupportedCaptureActions calls the method "MediaDevices.setSupportedCaptureActions".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) SetSupportedCaptureActions(actions js.Array[js.String]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesSetSupportedCaptureActions(
		this.Ref(), js.Pointer(&_ok),
		actions.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetSupportedCaptureActionsFunc returns the method "MediaDevices.setSupportedCaptureActions".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) SetSupportedCaptureActionsFunc() (fn js.Func[func(actions js.Array[js.String])]) {
	return fn.FromRef(
		bindings.MediaDevicesSetSupportedCaptureActionsFunc(
			this.Ref(),
		),
	)
}

// SelectAudioOutput calls the method "MediaDevices.selectAudioOutput".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) SelectAudioOutput(options AudioOutputOptions) (js.Promise[MediaDeviceInfo], bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesSelectAudioOutput(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[MediaDeviceInfo]{}.FromRef(_ret), _ok
}

// SelectAudioOutputFunc returns the method "MediaDevices.selectAudioOutput".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) SelectAudioOutputFunc() (fn js.Func[func(options AudioOutputOptions) js.Promise[MediaDeviceInfo]]) {
	return fn.FromRef(
		bindings.MediaDevicesSelectAudioOutputFunc(
			this.Ref(),
		),
	)
}

// SelectAudioOutput1 calls the method "MediaDevices.selectAudioOutput".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) SelectAudioOutput1() (js.Promise[MediaDeviceInfo], bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesSelectAudioOutput1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MediaDeviceInfo]{}.FromRef(_ret), _ok
}

// SelectAudioOutput1Func returns the method "MediaDevices.selectAudioOutput".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) SelectAudioOutput1Func() (fn js.Func[func() js.Promise[MediaDeviceInfo]]) {
	return fn.FromRef(
		bindings.MediaDevicesSelectAudioOutput1Func(
			this.Ref(),
		),
	)
}

// SetCaptureHandleConfig calls the method "MediaDevices.setCaptureHandleConfig".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) SetCaptureHandleConfig(config CaptureHandleConfig) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesSetCaptureHandleConfig(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&config),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCaptureHandleConfigFunc returns the method "MediaDevices.setCaptureHandleConfig".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) SetCaptureHandleConfigFunc() (fn js.Func[func(config CaptureHandleConfig)]) {
	return fn.FromRef(
		bindings.MediaDevicesSetCaptureHandleConfigFunc(
			this.Ref(),
		),
	)
}

// SetCaptureHandleConfig1 calls the method "MediaDevices.setCaptureHandleConfig".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) SetCaptureHandleConfig1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesSetCaptureHandleConfig1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCaptureHandleConfig1Func returns the method "MediaDevices.setCaptureHandleConfig".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) SetCaptureHandleConfig1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaDevicesSetCaptureHandleConfig1Func(
			this.Ref(),
		),
	)
}

// GetViewportMedia calls the method "MediaDevices.getViewportMedia".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) GetViewportMedia(constraints ViewportMediaStreamConstraints) (js.Promise[MediaStream], bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesGetViewportMedia(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&constraints),
	)

	return js.Promise[MediaStream]{}.FromRef(_ret), _ok
}

// GetViewportMediaFunc returns the method "MediaDevices.getViewportMedia".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) GetViewportMediaFunc() (fn js.Func[func(constraints ViewportMediaStreamConstraints) js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetViewportMediaFunc(
			this.Ref(),
		),
	)
}

// GetViewportMedia1 calls the method "MediaDevices.getViewportMedia".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) GetViewportMedia1() (js.Promise[MediaStream], bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesGetViewportMedia1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MediaStream]{}.FromRef(_ret), _ok
}

// GetViewportMedia1Func returns the method "MediaDevices.getViewportMedia".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) GetViewportMedia1Func() (fn js.Func[func() js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetViewportMedia1Func(
			this.Ref(),
		),
	)
}

// GetSupportedConstraints calls the method "MediaDevices.getSupportedConstraints".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) GetSupportedConstraints() (MediaTrackSupportedConstraints, bool) {
	var _ret MediaTrackSupportedConstraints
	_ok := js.True == bindings.CallMediaDevicesGetSupportedConstraints(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetSupportedConstraintsFunc returns the method "MediaDevices.getSupportedConstraints".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) GetSupportedConstraintsFunc() (fn js.Func[func() MediaTrackSupportedConstraints]) {
	return fn.FromRef(
		bindings.MediaDevicesGetSupportedConstraintsFunc(
			this.Ref(),
		),
	)
}

// GetUserMedia calls the method "MediaDevices.getUserMedia".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) GetUserMedia(constraints MediaStreamConstraints) (js.Promise[MediaStream], bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesGetUserMedia(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&constraints),
	)

	return js.Promise[MediaStream]{}.FromRef(_ret), _ok
}

// GetUserMediaFunc returns the method "MediaDevices.getUserMedia".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) GetUserMediaFunc() (fn js.Func[func(constraints MediaStreamConstraints) js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetUserMediaFunc(
			this.Ref(),
		),
	)
}

// GetUserMedia1 calls the method "MediaDevices.getUserMedia".
//
// The returned bool will be false if there is no such method.
func (this MediaDevices) GetUserMedia1() (js.Promise[MediaStream], bool) {
	var _ok bool
	_ret := bindings.CallMediaDevicesGetUserMedia1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MediaStream]{}.FromRef(_ret), _ok
}

// GetUserMedia1Func returns the method "MediaDevices.getUserMedia".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaDevices) GetUserMedia1Func() (fn js.Func[func() js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetUserMedia1Func(
			this.Ref(),
		),
	)
}

type SerialPortInfo struct {
	// UsbVendorId is "SerialPortInfo.usbVendorId"
	//
	// Optional
	UsbVendorId uint16
	// UsbProductId is "SerialPortInfo.usbProductId"
	//
	// Optional
	UsbProductId uint16
	// BluetoothServiceClassId is "SerialPortInfo.bluetoothServiceClassId"
	//
	// Optional
	BluetoothServiceClassId BluetoothServiceUUID

	FFI_USE_UsbVendorId  bool // for UsbVendorId.
	FFI_USE_UsbProductId bool // for UsbProductId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SerialPortInfo with all fields set.
func (p SerialPortInfo) FromRef(ref js.Ref) SerialPortInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SerialPortInfo SerialPortInfo [// SerialPortInfo] [0x14000e05360 0x14000e054a0 0x14000e05e00 0x14000e05400 0x14000e05720] 0x1400081e6a8 {0 0}} in the application heap.
func (p SerialPortInfo) New() js.Ref {
	return bindings.SerialPortInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SerialPortInfo) UpdateFrom(ref js.Ref) {
	bindings.SerialPortInfoJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SerialPortInfo) Update(ref js.Ref) {
	bindings.SerialPortInfoJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ParityType uint32

const (
	_ ParityType = iota

	ParityType_NONE
	ParityType_EVEN
	ParityType_ODD
)

func (ParityType) FromRef(str js.Ref) ParityType {
	return ParityType(bindings.ConstOfParityType(str))
}

func (x ParityType) String() (string, bool) {
	switch x {
	case ParityType_NONE:
		return "none", true
	case ParityType_EVEN:
		return "even", true
	case ParityType_ODD:
		return "odd", true
	default:
		return "", false
	}
}
