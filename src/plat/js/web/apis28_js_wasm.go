// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
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
// It returns ok=false if there is no such property.
func (this ServiceWorker) ScriptURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetServiceWorkerScriptURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// State returns the value of property "ServiceWorker.state".
//
// It returns ok=false if there is no such property.
func (this ServiceWorker) State() (ret ServiceWorkerState, ok bool) {
	ok = js.True == bindings.GetServiceWorkerState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPostMessage returns true if the method "ServiceWorker.postMessage" exists.
func (this ServiceWorker) HasPostMessage() bool {
	return js.True == bindings.HasServiceWorkerPostMessage(
		this.Ref(),
	)
}

// PostMessageFunc returns the method "ServiceWorker.postMessage".
func (this ServiceWorker) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.ServiceWorkerPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "ServiceWorker.postMessage".
func (this ServiceWorker) PostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void) {
	bindings.CallServiceWorkerPostMessage(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// TryPostMessage calls the method "ServiceWorker.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorker) TryPostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerPostMessage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// HasPostMessage1 returns true if the method "ServiceWorker.postMessage" exists.
func (this ServiceWorker) HasPostMessage1() bool {
	return js.True == bindings.HasServiceWorkerPostMessage1(
		this.Ref(),
	)
}

// PostMessage1Func returns the method "ServiceWorker.postMessage".
func (this ServiceWorker) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.ServiceWorkerPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "ServiceWorker.postMessage".
func (this ServiceWorker) PostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallServiceWorkerPostMessage1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage1 calls the method "ServiceWorker.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorker) TryPostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerPostMessage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPostMessage2 returns true if the method "ServiceWorker.postMessage" exists.
func (this ServiceWorker) HasPostMessage2() bool {
	return js.True == bindings.HasServiceWorkerPostMessage2(
		this.Ref(),
	)
}

// PostMessage2Func returns the method "ServiceWorker.postMessage".
func (this ServiceWorker) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.ServiceWorkerPostMessage2Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "ServiceWorker.postMessage".
func (this ServiceWorker) PostMessage2(message js.Any) (ret js.Void) {
	bindings.CallServiceWorkerPostMessage2(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage2 calls the method "ServiceWorker.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorker) TryPostMessage2(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerPostMessage2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

type NavigationPreloadState struct {
	// Enabled is "NavigationPreloadState.enabled"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
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

// New creates a new NavigationPreloadState in the application heap.
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

// HasEnable returns true if the method "NavigationPreloadManager.enable" exists.
func (this NavigationPreloadManager) HasEnable() bool {
	return js.True == bindings.HasNavigationPreloadManagerEnable(
		this.Ref(),
	)
}

// EnableFunc returns the method "NavigationPreloadManager.enable".
func (this NavigationPreloadManager) EnableFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigationPreloadManagerEnableFunc(
			this.Ref(),
		),
	)
}

// Enable calls the method "NavigationPreloadManager.enable".
func (this NavigationPreloadManager) Enable() (ret js.Promise[js.Void]) {
	bindings.CallNavigationPreloadManagerEnable(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEnable calls the method "NavigationPreloadManager.enable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigationPreloadManager) TryEnable() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationPreloadManagerEnable(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDisable returns true if the method "NavigationPreloadManager.disable" exists.
func (this NavigationPreloadManager) HasDisable() bool {
	return js.True == bindings.HasNavigationPreloadManagerDisable(
		this.Ref(),
	)
}

// DisableFunc returns the method "NavigationPreloadManager.disable".
func (this NavigationPreloadManager) DisableFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigationPreloadManagerDisableFunc(
			this.Ref(),
		),
	)
}

// Disable calls the method "NavigationPreloadManager.disable".
func (this NavigationPreloadManager) Disable() (ret js.Promise[js.Void]) {
	bindings.CallNavigationPreloadManagerDisable(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDisable calls the method "NavigationPreloadManager.disable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigationPreloadManager) TryDisable() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationPreloadManagerDisable(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetHeaderValue returns true if the method "NavigationPreloadManager.setHeaderValue" exists.
func (this NavigationPreloadManager) HasSetHeaderValue() bool {
	return js.True == bindings.HasNavigationPreloadManagerSetHeaderValue(
		this.Ref(),
	)
}

// SetHeaderValueFunc returns the method "NavigationPreloadManager.setHeaderValue".
func (this NavigationPreloadManager) SetHeaderValueFunc() (fn js.Func[func(value js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NavigationPreloadManagerSetHeaderValueFunc(
			this.Ref(),
		),
	)
}

// SetHeaderValue calls the method "NavigationPreloadManager.setHeaderValue".
func (this NavigationPreloadManager) SetHeaderValue(value js.String) (ret js.Promise[js.Void]) {
	bindings.CallNavigationPreloadManagerSetHeaderValue(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TrySetHeaderValue calls the method "NavigationPreloadManager.setHeaderValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigationPreloadManager) TrySetHeaderValue(value js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationPreloadManagerSetHeaderValue(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasGetState returns true if the method "NavigationPreloadManager.getState" exists.
func (this NavigationPreloadManager) HasGetState() bool {
	return js.True == bindings.HasNavigationPreloadManagerGetState(
		this.Ref(),
	)
}

// GetStateFunc returns the method "NavigationPreloadManager.getState".
func (this NavigationPreloadManager) GetStateFunc() (fn js.Func[func() js.Promise[NavigationPreloadState]]) {
	return fn.FromRef(
		bindings.NavigationPreloadManagerGetStateFunc(
			this.Ref(),
		),
	)
}

// GetState calls the method "NavigationPreloadManager.getState".
func (this NavigationPreloadManager) GetState() (ret js.Promise[NavigationPreloadState]) {
	bindings.CallNavigationPreloadManagerGetState(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetState calls the method "NavigationPreloadManager.getState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigationPreloadManager) TryGetState() (ret js.Promise[NavigationPreloadState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationPreloadManagerGetState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasRegister returns true if the method "SyncManager.register" exists.
func (this SyncManager) HasRegister() bool {
	return js.True == bindings.HasSyncManagerRegister(
		this.Ref(),
	)
}

// RegisterFunc returns the method "SyncManager.register".
func (this SyncManager) RegisterFunc() (fn js.Func[func(tag js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SyncManagerRegisterFunc(
			this.Ref(),
		),
	)
}

// Register calls the method "SyncManager.register".
func (this SyncManager) Register(tag js.String) (ret js.Promise[js.Void]) {
	bindings.CallSyncManagerRegister(
		this.Ref(), js.Pointer(&ret),
		tag.Ref(),
	)

	return
}

// TryRegister calls the method "SyncManager.register"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SyncManager) TryRegister(tag js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySyncManagerRegister(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tag.Ref(),
	)

	return
}

// HasGetTags returns true if the method "SyncManager.getTags" exists.
func (this SyncManager) HasGetTags() bool {
	return js.True == bindings.HasSyncManagerGetTags(
		this.Ref(),
	)
}

// GetTagsFunc returns the method "SyncManager.getTags".
func (this SyncManager) GetTagsFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.SyncManagerGetTagsFunc(
			this.Ref(),
		),
	)
}

// GetTags calls the method "SyncManager.getTags".
func (this SyncManager) GetTags() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallSyncManagerGetTags(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTags calls the method "SyncManager.getTags"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SyncManager) TryGetTags() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySyncManagerGetTags(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
	//
	// NOTE: FFI_USE_ExpirationTime MUST be set to true to make this field effective.
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

// New creates a new PushSubscriptionJSON in the application heap.
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
// It returns ok=false if there is no such property.
func (this PushSubscriptionOptions) UserVisibleOnly() (ret bool, ok bool) {
	ok = js.True == bindings.GetPushSubscriptionOptionsUserVisibleOnly(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ApplicationServerKey returns the value of property "PushSubscriptionOptions.applicationServerKey".
//
// It returns ok=false if there is no such property.
func (this PushSubscriptionOptions) ApplicationServerKey() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetPushSubscriptionOptionsApplicationServerKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this PushSubscription) Endpoint() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPushSubscriptionEndpoint(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ExpirationTime returns the value of property "PushSubscription.expirationTime".
//
// It returns ok=false if there is no such property.
func (this PushSubscription) ExpirationTime() (ret EpochTimeStamp, ok bool) {
	ok = js.True == bindings.GetPushSubscriptionExpirationTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Options returns the value of property "PushSubscription.options".
//
// It returns ok=false if there is no such property.
func (this PushSubscription) Options() (ret PushSubscriptionOptions, ok bool) {
	ok = js.True == bindings.GetPushSubscriptionOptions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetKey returns true if the method "PushSubscription.getKey" exists.
func (this PushSubscription) HasGetKey() bool {
	return js.True == bindings.HasPushSubscriptionGetKey(
		this.Ref(),
	)
}

// GetKeyFunc returns the method "PushSubscription.getKey".
func (this PushSubscription) GetKeyFunc() (fn js.Func[func(name PushEncryptionKeyName) js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.PushSubscriptionGetKeyFunc(
			this.Ref(),
		),
	)
}

// GetKey calls the method "PushSubscription.getKey".
func (this PushSubscription) GetKey(name PushEncryptionKeyName) (ret js.ArrayBuffer) {
	bindings.CallPushSubscriptionGetKey(
		this.Ref(), js.Pointer(&ret),
		uint32(name),
	)

	return
}

// TryGetKey calls the method "PushSubscription.getKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushSubscription) TryGetKey(name PushEncryptionKeyName) (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushSubscriptionGetKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(name),
	)

	return
}

// HasUnsubscribe returns true if the method "PushSubscription.unsubscribe" exists.
func (this PushSubscription) HasUnsubscribe() bool {
	return js.True == bindings.HasPushSubscriptionUnsubscribe(
		this.Ref(),
	)
}

// UnsubscribeFunc returns the method "PushSubscription.unsubscribe".
func (this PushSubscription) UnsubscribeFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.PushSubscriptionUnsubscribeFunc(
			this.Ref(),
		),
	)
}

// Unsubscribe calls the method "PushSubscription.unsubscribe".
func (this PushSubscription) Unsubscribe() (ret js.Promise[js.Boolean]) {
	bindings.CallPushSubscriptionUnsubscribe(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUnsubscribe calls the method "PushSubscription.unsubscribe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushSubscription) TryUnsubscribe() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushSubscriptionUnsubscribe(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasToJSON returns true if the method "PushSubscription.toJSON" exists.
func (this PushSubscription) HasToJSON() bool {
	return js.True == bindings.HasPushSubscriptionToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PushSubscription.toJSON".
func (this PushSubscription) ToJSONFunc() (fn js.Func[func() PushSubscriptionJSON]) {
	return fn.FromRef(
		bindings.PushSubscriptionToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PushSubscription.toJSON".
func (this PushSubscription) ToJSON() (ret PushSubscriptionJSON) {
	bindings.CallPushSubscriptionToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PushSubscription.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushSubscription) TryToJSON() (ret PushSubscriptionJSON, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushSubscriptionToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
	//
	// NOTE: FFI_USE_UserVisibleOnly MUST be set to true to make this field effective.
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

// New creates a new PushSubscriptionOptionsInit in the application heap.
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
// It returns ok=false if there is no such property.
func (this PushManager) SupportedContentEncodings() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetPushManagerSupportedContentEncodings(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSubscribe returns true if the method "PushManager.subscribe" exists.
func (this PushManager) HasSubscribe() bool {
	return js.True == bindings.HasPushManagerSubscribe(
		this.Ref(),
	)
}

// SubscribeFunc returns the method "PushManager.subscribe".
func (this PushManager) SubscribeFunc() (fn js.Func[func(options PushSubscriptionOptionsInit) js.Promise[PushSubscription]]) {
	return fn.FromRef(
		bindings.PushManagerSubscribeFunc(
			this.Ref(),
		),
	)
}

// Subscribe calls the method "PushManager.subscribe".
func (this PushManager) Subscribe(options PushSubscriptionOptionsInit) (ret js.Promise[PushSubscription]) {
	bindings.CallPushManagerSubscribe(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySubscribe calls the method "PushManager.subscribe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushManager) TrySubscribe(options PushSubscriptionOptionsInit) (ret js.Promise[PushSubscription], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushManagerSubscribe(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasSubscribe1 returns true if the method "PushManager.subscribe" exists.
func (this PushManager) HasSubscribe1() bool {
	return js.True == bindings.HasPushManagerSubscribe1(
		this.Ref(),
	)
}

// Subscribe1Func returns the method "PushManager.subscribe".
func (this PushManager) Subscribe1Func() (fn js.Func[func() js.Promise[PushSubscription]]) {
	return fn.FromRef(
		bindings.PushManagerSubscribe1Func(
			this.Ref(),
		),
	)
}

// Subscribe1 calls the method "PushManager.subscribe".
func (this PushManager) Subscribe1() (ret js.Promise[PushSubscription]) {
	bindings.CallPushManagerSubscribe1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySubscribe1 calls the method "PushManager.subscribe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushManager) TrySubscribe1() (ret js.Promise[PushSubscription], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushManagerSubscribe1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSubscription returns true if the method "PushManager.getSubscription" exists.
func (this PushManager) HasGetSubscription() bool {
	return js.True == bindings.HasPushManagerGetSubscription(
		this.Ref(),
	)
}

// GetSubscriptionFunc returns the method "PushManager.getSubscription".
func (this PushManager) GetSubscriptionFunc() (fn js.Func[func() js.Promise[PushSubscription]]) {
	return fn.FromRef(
		bindings.PushManagerGetSubscriptionFunc(
			this.Ref(),
		),
	)
}

// GetSubscription calls the method "PushManager.getSubscription".
func (this PushManager) GetSubscription() (ret js.Promise[PushSubscription]) {
	bindings.CallPushManagerGetSubscription(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSubscription calls the method "PushManager.getSubscription"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushManager) TryGetSubscription() (ret js.Promise[PushSubscription], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushManagerGetSubscription(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPermissionState returns true if the method "PushManager.permissionState" exists.
func (this PushManager) HasPermissionState() bool {
	return js.True == bindings.HasPushManagerPermissionState(
		this.Ref(),
	)
}

// PermissionStateFunc returns the method "PushManager.permissionState".
func (this PushManager) PermissionStateFunc() (fn js.Func[func(options PushSubscriptionOptionsInit) js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.PushManagerPermissionStateFunc(
			this.Ref(),
		),
	)
}

// PermissionState calls the method "PushManager.permissionState".
func (this PushManager) PermissionState(options PushSubscriptionOptionsInit) (ret js.Promise[PermissionState]) {
	bindings.CallPushManagerPermissionState(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryPermissionState calls the method "PushManager.permissionState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushManager) TryPermissionState(options PushSubscriptionOptionsInit) (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushManagerPermissionState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasPermissionState1 returns true if the method "PushManager.permissionState" exists.
func (this PushManager) HasPermissionState1() bool {
	return js.True == bindings.HasPushManagerPermissionState1(
		this.Ref(),
	)
}

// PermissionState1Func returns the method "PushManager.permissionState".
func (this PushManager) PermissionState1Func() (fn js.Func[func() js.Promise[PermissionState]]) {
	return fn.FromRef(
		bindings.PushManagerPermissionState1Func(
			this.Ref(),
		),
	)
}

// PermissionState1 calls the method "PushManager.permissionState".
func (this PushManager) PermissionState1() (ret js.Promise[PermissionState]) {
	bindings.CallPushManagerPermissionState1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPermissionState1 calls the method "PushManager.permissionState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushManager) TryPermissionState1() (ret js.Promise[PermissionState], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushManagerPermissionState1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this PaymentManager) UserHint() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaymentManagerUserHint(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetUserHint sets the value of property "PaymentManager.userHint" to val.
//
// It returns false if the property cannot be set.
func (this PaymentManager) SetUserHint(val js.String) bool {
	return js.True == bindings.SetPaymentManagerUserHint(
		this.Ref(),
		val.Ref(),
	)
}

// HasEnableDelegations returns true if the method "PaymentManager.enableDelegations" exists.
func (this PaymentManager) HasEnableDelegations() bool {
	return js.True == bindings.HasPaymentManagerEnableDelegations(
		this.Ref(),
	)
}

// EnableDelegationsFunc returns the method "PaymentManager.enableDelegations".
func (this PaymentManager) EnableDelegationsFunc() (fn js.Func[func(delegations js.Array[PaymentDelegation]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentManagerEnableDelegationsFunc(
			this.Ref(),
		),
	)
}

// EnableDelegations calls the method "PaymentManager.enableDelegations".
func (this PaymentManager) EnableDelegations(delegations js.Array[PaymentDelegation]) (ret js.Promise[js.Void]) {
	bindings.CallPaymentManagerEnableDelegations(
		this.Ref(), js.Pointer(&ret),
		delegations.Ref(),
	)

	return
}

// TryEnableDelegations calls the method "PaymentManager.enableDelegations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentManager) TryEnableDelegations(delegations js.Array[PaymentDelegation]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentManagerEnableDelegations(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		delegations.Ref(),
	)

	return
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

// HasRegister returns true if the method "PeriodicSyncManager.register" exists.
func (this PeriodicSyncManager) HasRegister() bool {
	return js.True == bindings.HasPeriodicSyncManagerRegister(
		this.Ref(),
	)
}

// RegisterFunc returns the method "PeriodicSyncManager.register".
func (this PeriodicSyncManager) RegisterFunc() (fn js.Func[func(tag js.String, options BackgroundSyncOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PeriodicSyncManagerRegisterFunc(
			this.Ref(),
		),
	)
}

// Register calls the method "PeriodicSyncManager.register".
func (this PeriodicSyncManager) Register(tag js.String, options BackgroundSyncOptions) (ret js.Promise[js.Void]) {
	bindings.CallPeriodicSyncManagerRegister(
		this.Ref(), js.Pointer(&ret),
		tag.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryRegister calls the method "PeriodicSyncManager.register"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PeriodicSyncManager) TryRegister(tag js.String, options BackgroundSyncOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPeriodicSyncManagerRegister(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tag.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasRegister1 returns true if the method "PeriodicSyncManager.register" exists.
func (this PeriodicSyncManager) HasRegister1() bool {
	return js.True == bindings.HasPeriodicSyncManagerRegister1(
		this.Ref(),
	)
}

// Register1Func returns the method "PeriodicSyncManager.register".
func (this PeriodicSyncManager) Register1Func() (fn js.Func[func(tag js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PeriodicSyncManagerRegister1Func(
			this.Ref(),
		),
	)
}

// Register1 calls the method "PeriodicSyncManager.register".
func (this PeriodicSyncManager) Register1(tag js.String) (ret js.Promise[js.Void]) {
	bindings.CallPeriodicSyncManagerRegister1(
		this.Ref(), js.Pointer(&ret),
		tag.Ref(),
	)

	return
}

// TryRegister1 calls the method "PeriodicSyncManager.register"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PeriodicSyncManager) TryRegister1(tag js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPeriodicSyncManagerRegister1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tag.Ref(),
	)

	return
}

// HasGetTags returns true if the method "PeriodicSyncManager.getTags" exists.
func (this PeriodicSyncManager) HasGetTags() bool {
	return js.True == bindings.HasPeriodicSyncManagerGetTags(
		this.Ref(),
	)
}

// GetTagsFunc returns the method "PeriodicSyncManager.getTags".
func (this PeriodicSyncManager) GetTagsFunc() (fn js.Func[func() js.Promise[js.Array[js.String]]]) {
	return fn.FromRef(
		bindings.PeriodicSyncManagerGetTagsFunc(
			this.Ref(),
		),
	)
}

// GetTags calls the method "PeriodicSyncManager.getTags".
func (this PeriodicSyncManager) GetTags() (ret js.Promise[js.Array[js.String]]) {
	bindings.CallPeriodicSyncManagerGetTags(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTags calls the method "PeriodicSyncManager.getTags"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PeriodicSyncManager) TryGetTags() (ret js.Promise[js.Array[js.String]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPeriodicSyncManagerGetTags(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasUnregister returns true if the method "PeriodicSyncManager.unregister" exists.
func (this PeriodicSyncManager) HasUnregister() bool {
	return js.True == bindings.HasPeriodicSyncManagerUnregister(
		this.Ref(),
	)
}

// UnregisterFunc returns the method "PeriodicSyncManager.unregister".
func (this PeriodicSyncManager) UnregisterFunc() (fn js.Func[func(tag js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PeriodicSyncManagerUnregisterFunc(
			this.Ref(),
		),
	)
}

// Unregister calls the method "PeriodicSyncManager.unregister".
func (this PeriodicSyncManager) Unregister(tag js.String) (ret js.Promise[js.Void]) {
	bindings.CallPeriodicSyncManagerUnregister(
		this.Ref(), js.Pointer(&ret),
		tag.Ref(),
	)

	return
}

// TryUnregister calls the method "PeriodicSyncManager.unregister"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PeriodicSyncManager) TryUnregister(tag js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPeriodicSyncManagerUnregister(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		tag.Ref(),
	)

	return
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

// New creates a new CookieStoreGetOptions in the application heap.
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

// HasSubscribe returns true if the method "CookieStoreManager.subscribe" exists.
func (this CookieStoreManager) HasSubscribe() bool {
	return js.True == bindings.HasCookieStoreManagerSubscribe(
		this.Ref(),
	)
}

// SubscribeFunc returns the method "CookieStoreManager.subscribe".
func (this CookieStoreManager) SubscribeFunc() (fn js.Func[func(subscriptions js.Array[CookieStoreGetOptions]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreManagerSubscribeFunc(
			this.Ref(),
		),
	)
}

// Subscribe calls the method "CookieStoreManager.subscribe".
func (this CookieStoreManager) Subscribe(subscriptions js.Array[CookieStoreGetOptions]) (ret js.Promise[js.Void]) {
	bindings.CallCookieStoreManagerSubscribe(
		this.Ref(), js.Pointer(&ret),
		subscriptions.Ref(),
	)

	return
}

// TrySubscribe calls the method "CookieStoreManager.subscribe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CookieStoreManager) TrySubscribe(subscriptions js.Array[CookieStoreGetOptions]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreManagerSubscribe(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		subscriptions.Ref(),
	)

	return
}

// HasGetSubscriptions returns true if the method "CookieStoreManager.getSubscriptions" exists.
func (this CookieStoreManager) HasGetSubscriptions() bool {
	return js.True == bindings.HasCookieStoreManagerGetSubscriptions(
		this.Ref(),
	)
}

// GetSubscriptionsFunc returns the method "CookieStoreManager.getSubscriptions".
func (this CookieStoreManager) GetSubscriptionsFunc() (fn js.Func[func() js.Promise[js.Array[CookieStoreGetOptions]]]) {
	return fn.FromRef(
		bindings.CookieStoreManagerGetSubscriptionsFunc(
			this.Ref(),
		),
	)
}

// GetSubscriptions calls the method "CookieStoreManager.getSubscriptions".
func (this CookieStoreManager) GetSubscriptions() (ret js.Promise[js.Array[CookieStoreGetOptions]]) {
	bindings.CallCookieStoreManagerGetSubscriptions(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSubscriptions calls the method "CookieStoreManager.getSubscriptions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CookieStoreManager) TryGetSubscriptions() (ret js.Promise[js.Array[CookieStoreGetOptions]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreManagerGetSubscriptions(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasUnsubscribe returns true if the method "CookieStoreManager.unsubscribe" exists.
func (this CookieStoreManager) HasUnsubscribe() bool {
	return js.True == bindings.HasCookieStoreManagerUnsubscribe(
		this.Ref(),
	)
}

// UnsubscribeFunc returns the method "CookieStoreManager.unsubscribe".
func (this CookieStoreManager) UnsubscribeFunc() (fn js.Func[func(subscriptions js.Array[CookieStoreGetOptions]) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.CookieStoreManagerUnsubscribeFunc(
			this.Ref(),
		),
	)
}

// Unsubscribe calls the method "CookieStoreManager.unsubscribe".
func (this CookieStoreManager) Unsubscribe(subscriptions js.Array[CookieStoreGetOptions]) (ret js.Promise[js.Void]) {
	bindings.CallCookieStoreManagerUnsubscribe(
		this.Ref(), js.Pointer(&ret),
		subscriptions.Ref(),
	)

	return
}

// TryUnsubscribe calls the method "CookieStoreManager.unsubscribe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CookieStoreManager) TryUnsubscribe(subscriptions js.Array[CookieStoreGetOptions]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCookieStoreManagerUnsubscribe(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		subscriptions.Ref(),
	)

	return
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

// New creates a new ContentDescription in the application heap.
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

// HasAdd returns true if the method "ContentIndex.add" exists.
func (this ContentIndex) HasAdd() bool {
	return js.True == bindings.HasContentIndexAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "ContentIndex.add".
func (this ContentIndex) AddFunc() (fn js.Func[func(description ContentDescription) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ContentIndexAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "ContentIndex.add".
func (this ContentIndex) Add(description ContentDescription) (ret js.Promise[js.Void]) {
	bindings.CallContentIndexAdd(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&description),
	)

	return
}

// TryAdd calls the method "ContentIndex.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContentIndex) TryAdd(description ContentDescription) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContentIndexAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&description),
	)

	return
}

// HasDelete returns true if the method "ContentIndex.delete" exists.
func (this ContentIndex) HasDelete() bool {
	return js.True == bindings.HasContentIndexDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "ContentIndex.delete".
func (this ContentIndex) DeleteFunc() (fn js.Func[func(id js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ContentIndexDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "ContentIndex.delete".
func (this ContentIndex) Delete(id js.String) (ret js.Promise[js.Void]) {
	bindings.CallContentIndexDelete(
		this.Ref(), js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryDelete calls the method "ContentIndex.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContentIndex) TryDelete(id js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContentIndexDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasGetAll returns true if the method "ContentIndex.getAll" exists.
func (this ContentIndex) HasGetAll() bool {
	return js.True == bindings.HasContentIndexGetAll(
		this.Ref(),
	)
}

// GetAllFunc returns the method "ContentIndex.getAll".
func (this ContentIndex) GetAllFunc() (fn js.Func[func() js.Promise[js.Array[ContentDescription]]]) {
	return fn.FromRef(
		bindings.ContentIndexGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "ContentIndex.getAll".
func (this ContentIndex) GetAll() (ret js.Promise[js.Array[ContentDescription]]) {
	bindings.CallContentIndexGetAll(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAll calls the method "ContentIndex.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContentIndex) TryGetAll() (ret js.Promise[js.Array[ContentDescription]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContentIndexGetAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) Installing() (ret ServiceWorker, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationInstalling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Waiting returns the value of property "ServiceWorkerRegistration.waiting".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) Waiting() (ret ServiceWorker, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationWaiting(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Active returns the value of property "ServiceWorkerRegistration.active".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) Active() (ret ServiceWorker, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationActive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NavigationPreload returns the value of property "ServiceWorkerRegistration.navigationPreload".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) NavigationPreload() (ret NavigationPreloadManager, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationNavigationPreload(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Scope returns the value of property "ServiceWorkerRegistration.scope".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) Scope() (ret js.String, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationScope(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UpdateViaCache returns the value of property "ServiceWorkerRegistration.updateViaCache".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) UpdateViaCache() (ret ServiceWorkerUpdateViaCache, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationUpdateViaCache(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Sync returns the value of property "ServiceWorkerRegistration.sync".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) Sync() (ret SyncManager, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationSync(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BackgroundFetch returns the value of property "ServiceWorkerRegistration.backgroundFetch".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) BackgroundFetch() (ret BackgroundFetchManager, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationBackgroundFetch(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PushManager returns the value of property "ServiceWorkerRegistration.pushManager".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) PushManager() (ret PushManager, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationPushManager(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PaymentManager returns the value of property "ServiceWorkerRegistration.paymentManager".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) PaymentManager() (ret PaymentManager, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationPaymentManager(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PeriodicSync returns the value of property "ServiceWorkerRegistration.periodicSync".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) PeriodicSync() (ret PeriodicSyncManager, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationPeriodicSync(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Cookies returns the value of property "ServiceWorkerRegistration.cookies".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) Cookies() (ret CookieStoreManager, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationCookies(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Index returns the value of property "ServiceWorkerRegistration.index".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerRegistration) Index() (ret ContentIndex, ok bool) {
	ok = js.True == bindings.GetServiceWorkerRegistrationIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasUpdate returns true if the method "ServiceWorkerRegistration.update" exists.
func (this ServiceWorkerRegistration) HasUpdate() bool {
	return js.True == bindings.HasServiceWorkerRegistrationUpdate(
		this.Ref(),
	)
}

// UpdateFunc returns the method "ServiceWorkerRegistration.update".
func (this ServiceWorkerRegistration) UpdateFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationUpdateFunc(
			this.Ref(),
		),
	)
}

// Update calls the method "ServiceWorkerRegistration.update".
func (this ServiceWorkerRegistration) Update() (ret js.Promise[js.Void]) {
	bindings.CallServiceWorkerRegistrationUpdate(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUpdate calls the method "ServiceWorkerRegistration.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerRegistration) TryUpdate() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerRegistrationUpdate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasUnregister returns true if the method "ServiceWorkerRegistration.unregister" exists.
func (this ServiceWorkerRegistration) HasUnregister() bool {
	return js.True == bindings.HasServiceWorkerRegistrationUnregister(
		this.Ref(),
	)
}

// UnregisterFunc returns the method "ServiceWorkerRegistration.unregister".
func (this ServiceWorkerRegistration) UnregisterFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationUnregisterFunc(
			this.Ref(),
		),
	)
}

// Unregister calls the method "ServiceWorkerRegistration.unregister".
func (this ServiceWorkerRegistration) Unregister() (ret js.Promise[js.Boolean]) {
	bindings.CallServiceWorkerRegistrationUnregister(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUnregister calls the method "ServiceWorkerRegistration.unregister"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerRegistration) TryUnregister() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerRegistrationUnregister(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasShowNotification returns true if the method "ServiceWorkerRegistration.showNotification" exists.
func (this ServiceWorkerRegistration) HasShowNotification() bool {
	return js.True == bindings.HasServiceWorkerRegistrationShowNotification(
		this.Ref(),
	)
}

// ShowNotificationFunc returns the method "ServiceWorkerRegistration.showNotification".
func (this ServiceWorkerRegistration) ShowNotificationFunc() (fn js.Func[func(title js.String, options NotificationOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationShowNotificationFunc(
			this.Ref(),
		),
	)
}

// ShowNotification calls the method "ServiceWorkerRegistration.showNotification".
func (this ServiceWorkerRegistration) ShowNotification(title js.String, options NotificationOptions) (ret js.Promise[js.Void]) {
	bindings.CallServiceWorkerRegistrationShowNotification(
		this.Ref(), js.Pointer(&ret),
		title.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryShowNotification calls the method "ServiceWorkerRegistration.showNotification"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerRegistration) TryShowNotification(title js.String, options NotificationOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerRegistrationShowNotification(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		title.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasShowNotification1 returns true if the method "ServiceWorkerRegistration.showNotification" exists.
func (this ServiceWorkerRegistration) HasShowNotification1() bool {
	return js.True == bindings.HasServiceWorkerRegistrationShowNotification1(
		this.Ref(),
	)
}

// ShowNotification1Func returns the method "ServiceWorkerRegistration.showNotification".
func (this ServiceWorkerRegistration) ShowNotification1Func() (fn js.Func[func(title js.String) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationShowNotification1Func(
			this.Ref(),
		),
	)
}

// ShowNotification1 calls the method "ServiceWorkerRegistration.showNotification".
func (this ServiceWorkerRegistration) ShowNotification1(title js.String) (ret js.Promise[js.Void]) {
	bindings.CallServiceWorkerRegistrationShowNotification1(
		this.Ref(), js.Pointer(&ret),
		title.Ref(),
	)

	return
}

// TryShowNotification1 calls the method "ServiceWorkerRegistration.showNotification"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerRegistration) TryShowNotification1(title js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerRegistrationShowNotification1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		title.Ref(),
	)

	return
}

// HasGetNotifications returns true if the method "ServiceWorkerRegistration.getNotifications" exists.
func (this ServiceWorkerRegistration) HasGetNotifications() bool {
	return js.True == bindings.HasServiceWorkerRegistrationGetNotifications(
		this.Ref(),
	)
}

// GetNotificationsFunc returns the method "ServiceWorkerRegistration.getNotifications".
func (this ServiceWorkerRegistration) GetNotificationsFunc() (fn js.Func[func(filter GetNotificationOptions) js.Promise[js.Array[Notification]]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationGetNotificationsFunc(
			this.Ref(),
		),
	)
}

// GetNotifications calls the method "ServiceWorkerRegistration.getNotifications".
func (this ServiceWorkerRegistration) GetNotifications(filter GetNotificationOptions) (ret js.Promise[js.Array[Notification]]) {
	bindings.CallServiceWorkerRegistrationGetNotifications(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetNotifications calls the method "ServiceWorkerRegistration.getNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerRegistration) TryGetNotifications(filter GetNotificationOptions) (ret js.Promise[js.Array[Notification]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerRegistrationGetNotifications(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasGetNotifications1 returns true if the method "ServiceWorkerRegistration.getNotifications" exists.
func (this ServiceWorkerRegistration) HasGetNotifications1() bool {
	return js.True == bindings.HasServiceWorkerRegistrationGetNotifications1(
		this.Ref(),
	)
}

// GetNotifications1Func returns the method "ServiceWorkerRegistration.getNotifications".
func (this ServiceWorkerRegistration) GetNotifications1Func() (fn js.Func[func() js.Promise[js.Array[Notification]]]) {
	return fn.FromRef(
		bindings.ServiceWorkerRegistrationGetNotifications1Func(
			this.Ref(),
		),
	)
}

// GetNotifications1 calls the method "ServiceWorkerRegistration.getNotifications".
func (this ServiceWorkerRegistration) GetNotifications1() (ret js.Promise[js.Array[Notification]]) {
	bindings.CallServiceWorkerRegistrationGetNotifications1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetNotifications1 calls the method "ServiceWorkerRegistration.getNotifications"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerRegistration) TryGetNotifications1() (ret js.Promise[js.Array[Notification]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerRegistrationGetNotifications1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new RegistrationOptions in the application heap.
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
// It returns ok=false if there is no such property.
func (this ServiceWorkerContainer) Controller() (ret ServiceWorker, ok bool) {
	ok = js.True == bindings.GetServiceWorkerContainerController(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ready returns the value of property "ServiceWorkerContainer.ready".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerContainer) Ready() (ret js.Promise[ServiceWorkerRegistration], ok bool) {
	ok = js.True == bindings.GetServiceWorkerContainerReady(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRegister returns true if the method "ServiceWorkerContainer.register" exists.
func (this ServiceWorkerContainer) HasRegister() bool {
	return js.True == bindings.HasServiceWorkerContainerRegister(
		this.Ref(),
	)
}

// RegisterFunc returns the method "ServiceWorkerContainer.register".
func (this ServiceWorkerContainer) RegisterFunc() (fn js.Func[func(scriptURL js.String, options RegistrationOptions) js.Promise[ServiceWorkerRegistration]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerRegisterFunc(
			this.Ref(),
		),
	)
}

// Register calls the method "ServiceWorkerContainer.register".
func (this ServiceWorkerContainer) Register(scriptURL js.String, options RegistrationOptions) (ret js.Promise[ServiceWorkerRegistration]) {
	bindings.CallServiceWorkerContainerRegister(
		this.Ref(), js.Pointer(&ret),
		scriptURL.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryRegister calls the method "ServiceWorkerContainer.register"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerContainer) TryRegister(scriptURL js.String, options RegistrationOptions) (ret js.Promise[ServiceWorkerRegistration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerContainerRegister(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		scriptURL.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasRegister1 returns true if the method "ServiceWorkerContainer.register" exists.
func (this ServiceWorkerContainer) HasRegister1() bool {
	return js.True == bindings.HasServiceWorkerContainerRegister1(
		this.Ref(),
	)
}

// Register1Func returns the method "ServiceWorkerContainer.register".
func (this ServiceWorkerContainer) Register1Func() (fn js.Func[func(scriptURL js.String) js.Promise[ServiceWorkerRegistration]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerRegister1Func(
			this.Ref(),
		),
	)
}

// Register1 calls the method "ServiceWorkerContainer.register".
func (this ServiceWorkerContainer) Register1(scriptURL js.String) (ret js.Promise[ServiceWorkerRegistration]) {
	bindings.CallServiceWorkerContainerRegister1(
		this.Ref(), js.Pointer(&ret),
		scriptURL.Ref(),
	)

	return
}

// TryRegister1 calls the method "ServiceWorkerContainer.register"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerContainer) TryRegister1(scriptURL js.String) (ret js.Promise[ServiceWorkerRegistration], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerContainerRegister1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		scriptURL.Ref(),
	)

	return
}

// HasGetRegistration returns true if the method "ServiceWorkerContainer.getRegistration" exists.
func (this ServiceWorkerContainer) HasGetRegistration() bool {
	return js.True == bindings.HasServiceWorkerContainerGetRegistration(
		this.Ref(),
	)
}

// GetRegistrationFunc returns the method "ServiceWorkerContainer.getRegistration".
func (this ServiceWorkerContainer) GetRegistrationFunc() (fn js.Func[func(clientURL js.String) js.Promise[OneOf_ServiceWorkerRegistration_undefined]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerGetRegistrationFunc(
			this.Ref(),
		),
	)
}

// GetRegistration calls the method "ServiceWorkerContainer.getRegistration".
func (this ServiceWorkerContainer) GetRegistration(clientURL js.String) (ret js.Promise[OneOf_ServiceWorkerRegistration_undefined]) {
	bindings.CallServiceWorkerContainerGetRegistration(
		this.Ref(), js.Pointer(&ret),
		clientURL.Ref(),
	)

	return
}

// TryGetRegistration calls the method "ServiceWorkerContainer.getRegistration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerContainer) TryGetRegistration(clientURL js.String) (ret js.Promise[OneOf_ServiceWorkerRegistration_undefined], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerContainerGetRegistration(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		clientURL.Ref(),
	)

	return
}

// HasGetRegistration1 returns true if the method "ServiceWorkerContainer.getRegistration" exists.
func (this ServiceWorkerContainer) HasGetRegistration1() bool {
	return js.True == bindings.HasServiceWorkerContainerGetRegistration1(
		this.Ref(),
	)
}

// GetRegistration1Func returns the method "ServiceWorkerContainer.getRegistration".
func (this ServiceWorkerContainer) GetRegistration1Func() (fn js.Func[func() js.Promise[OneOf_ServiceWorkerRegistration_undefined]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerGetRegistration1Func(
			this.Ref(),
		),
	)
}

// GetRegistration1 calls the method "ServiceWorkerContainer.getRegistration".
func (this ServiceWorkerContainer) GetRegistration1() (ret js.Promise[OneOf_ServiceWorkerRegistration_undefined]) {
	bindings.CallServiceWorkerContainerGetRegistration1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetRegistration1 calls the method "ServiceWorkerContainer.getRegistration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerContainer) TryGetRegistration1() (ret js.Promise[OneOf_ServiceWorkerRegistration_undefined], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerContainerGetRegistration1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetRegistrations returns true if the method "ServiceWorkerContainer.getRegistrations" exists.
func (this ServiceWorkerContainer) HasGetRegistrations() bool {
	return js.True == bindings.HasServiceWorkerContainerGetRegistrations(
		this.Ref(),
	)
}

// GetRegistrationsFunc returns the method "ServiceWorkerContainer.getRegistrations".
func (this ServiceWorkerContainer) GetRegistrationsFunc() (fn js.Func[func() js.Promise[js.FrozenArray[ServiceWorkerRegistration]]]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerGetRegistrationsFunc(
			this.Ref(),
		),
	)
}

// GetRegistrations calls the method "ServiceWorkerContainer.getRegistrations".
func (this ServiceWorkerContainer) GetRegistrations() (ret js.Promise[js.FrozenArray[ServiceWorkerRegistration]]) {
	bindings.CallServiceWorkerContainerGetRegistrations(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetRegistrations calls the method "ServiceWorkerContainer.getRegistrations"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerContainer) TryGetRegistrations() (ret js.Promise[js.FrozenArray[ServiceWorkerRegistration]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerContainerGetRegistrations(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStartMessages returns true if the method "ServiceWorkerContainer.startMessages" exists.
func (this ServiceWorkerContainer) HasStartMessages() bool {
	return js.True == bindings.HasServiceWorkerContainerStartMessages(
		this.Ref(),
	)
}

// StartMessagesFunc returns the method "ServiceWorkerContainer.startMessages".
func (this ServiceWorkerContainer) StartMessagesFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.ServiceWorkerContainerStartMessagesFunc(
			this.Ref(),
		),
	)
}

// StartMessages calls the method "ServiceWorkerContainer.startMessages".
func (this ServiceWorkerContainer) StartMessages() (ret js.Void) {
	bindings.CallServiceWorkerContainerStartMessages(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStartMessages calls the method "ServiceWorkerContainer.startMessages"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerContainer) TryStartMessages() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerContainerStartMessages(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this MediaDeviceInfo) DeviceId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaDeviceInfoDeviceId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Kind returns the value of property "MediaDeviceInfo.kind".
//
// It returns ok=false if there is no such property.
func (this MediaDeviceInfo) Kind() (ret MediaDeviceKind, ok bool) {
	ok = js.True == bindings.GetMediaDeviceInfoKind(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "MediaDeviceInfo.label".
//
// It returns ok=false if there is no such property.
func (this MediaDeviceInfo) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaDeviceInfoLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// GroupId returns the value of property "MediaDeviceInfo.groupId".
//
// It returns ok=false if there is no such property.
func (this MediaDeviceInfo) GroupId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaDeviceInfoGroupId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "MediaDeviceInfo.toJSON" exists.
func (this MediaDeviceInfo) HasToJSON() bool {
	return js.True == bindings.HasMediaDeviceInfoToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "MediaDeviceInfo.toJSON".
func (this MediaDeviceInfo) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.MediaDeviceInfoToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "MediaDeviceInfo.toJSON".
func (this MediaDeviceInfo) ToJSON() (ret js.Object) {
	bindings.CallMediaDeviceInfoToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "MediaDeviceInfo.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDeviceInfo) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDeviceInfoToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new DisplayMediaStreamOptions in the application heap.
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

// New creates a new ViewportMediaStreamConstraints in the application heap.
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
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width bool
	// Height is "MediaTrackSupportedConstraints.height"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height bool
	// AspectRatio is "MediaTrackSupportedConstraints.aspectRatio"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_AspectRatio MUST be set to true to make this field effective.
	AspectRatio bool
	// FrameRate is "MediaTrackSupportedConstraints.frameRate"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_FrameRate MUST be set to true to make this field effective.
	FrameRate bool
	// FacingMode is "MediaTrackSupportedConstraints.facingMode"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_FacingMode MUST be set to true to make this field effective.
	FacingMode bool
	// ResizeMode is "MediaTrackSupportedConstraints.resizeMode"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ResizeMode MUST be set to true to make this field effective.
	ResizeMode bool
	// SampleRate is "MediaTrackSupportedConstraints.sampleRate"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_SampleRate MUST be set to true to make this field effective.
	SampleRate bool
	// SampleSize is "MediaTrackSupportedConstraints.sampleSize"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_SampleSize MUST be set to true to make this field effective.
	SampleSize bool
	// EchoCancellation is "MediaTrackSupportedConstraints.echoCancellation"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_EchoCancellation MUST be set to true to make this field effective.
	EchoCancellation bool
	// AutoGainControl is "MediaTrackSupportedConstraints.autoGainControl"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_AutoGainControl MUST be set to true to make this field effective.
	AutoGainControl bool
	// NoiseSuppression is "MediaTrackSupportedConstraints.noiseSuppression"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_NoiseSuppression MUST be set to true to make this field effective.
	NoiseSuppression bool
	// Latency is "MediaTrackSupportedConstraints.latency"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Latency MUST be set to true to make this field effective.
	Latency bool
	// ChannelCount is "MediaTrackSupportedConstraints.channelCount"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
	ChannelCount bool
	// DeviceId is "MediaTrackSupportedConstraints.deviceId"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_DeviceId MUST be set to true to make this field effective.
	DeviceId bool
	// GroupId is "MediaTrackSupportedConstraints.groupId"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_GroupId MUST be set to true to make this field effective.
	GroupId bool
	// WhiteBalanceMode is "MediaTrackSupportedConstraints.whiteBalanceMode"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_WhiteBalanceMode MUST be set to true to make this field effective.
	WhiteBalanceMode bool
	// ExposureMode is "MediaTrackSupportedConstraints.exposureMode"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ExposureMode MUST be set to true to make this field effective.
	ExposureMode bool
	// FocusMode is "MediaTrackSupportedConstraints.focusMode"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_FocusMode MUST be set to true to make this field effective.
	FocusMode bool
	// PointsOfInterest is "MediaTrackSupportedConstraints.pointsOfInterest"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_PointsOfInterest MUST be set to true to make this field effective.
	PointsOfInterest bool
	// ExposureCompensation is "MediaTrackSupportedConstraints.exposureCompensation"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ExposureCompensation MUST be set to true to make this field effective.
	ExposureCompensation bool
	// ExposureTime is "MediaTrackSupportedConstraints.exposureTime"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ExposureTime MUST be set to true to make this field effective.
	ExposureTime bool
	// ColorTemperature is "MediaTrackSupportedConstraints.colorTemperature"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ColorTemperature MUST be set to true to make this field effective.
	ColorTemperature bool
	// Iso is "MediaTrackSupportedConstraints.iso"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Iso MUST be set to true to make this field effective.
	Iso bool
	// Brightness is "MediaTrackSupportedConstraints.brightness"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Brightness MUST be set to true to make this field effective.
	Brightness bool
	// Contrast is "MediaTrackSupportedConstraints.contrast"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Contrast MUST be set to true to make this field effective.
	Contrast bool
	// Pan is "MediaTrackSupportedConstraints.pan"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Pan MUST be set to true to make this field effective.
	Pan bool
	// Saturation is "MediaTrackSupportedConstraints.saturation"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Saturation MUST be set to true to make this field effective.
	Saturation bool
	// Sharpness is "MediaTrackSupportedConstraints.sharpness"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Sharpness MUST be set to true to make this field effective.
	Sharpness bool
	// FocusDistance is "MediaTrackSupportedConstraints.focusDistance"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_FocusDistance MUST be set to true to make this field effective.
	FocusDistance bool
	// Tilt is "MediaTrackSupportedConstraints.tilt"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Tilt MUST be set to true to make this field effective.
	Tilt bool
	// Zoom is "MediaTrackSupportedConstraints.zoom"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Zoom MUST be set to true to make this field effective.
	Zoom bool
	// Torch is "MediaTrackSupportedConstraints.torch"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Torch MUST be set to true to make this field effective.
	Torch bool
	// DisplaySurface is "MediaTrackSupportedConstraints.displaySurface"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_DisplaySurface MUST be set to true to make this field effective.
	DisplaySurface bool
	// LogicalSurface is "MediaTrackSupportedConstraints.logicalSurface"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_LogicalSurface MUST be set to true to make this field effective.
	LogicalSurface bool
	// Cursor is "MediaTrackSupportedConstraints.cursor"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Cursor MUST be set to true to make this field effective.
	Cursor bool
	// RestrictOwnAudio is "MediaTrackSupportedConstraints.restrictOwnAudio"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_RestrictOwnAudio MUST be set to true to make this field effective.
	RestrictOwnAudio bool
	// SuppressLocalAudioPlayback is "MediaTrackSupportedConstraints.suppressLocalAudioPlayback"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_SuppressLocalAudioPlayback MUST be set to true to make this field effective.
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

// New creates a new MediaTrackSupportedConstraints in the application heap.
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

// HasEnumerateDevices returns true if the method "MediaDevices.enumerateDevices" exists.
func (this MediaDevices) HasEnumerateDevices() bool {
	return js.True == bindings.HasMediaDevicesEnumerateDevices(
		this.Ref(),
	)
}

// EnumerateDevicesFunc returns the method "MediaDevices.enumerateDevices".
func (this MediaDevices) EnumerateDevicesFunc() (fn js.Func[func() js.Promise[js.Array[MediaDeviceInfo]]]) {
	return fn.FromRef(
		bindings.MediaDevicesEnumerateDevicesFunc(
			this.Ref(),
		),
	)
}

// EnumerateDevices calls the method "MediaDevices.enumerateDevices".
func (this MediaDevices) EnumerateDevices() (ret js.Promise[js.Array[MediaDeviceInfo]]) {
	bindings.CallMediaDevicesEnumerateDevices(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEnumerateDevices calls the method "MediaDevices.enumerateDevices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TryEnumerateDevices() (ret js.Promise[js.Array[MediaDeviceInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesEnumerateDevices(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetDisplayMedia returns true if the method "MediaDevices.getDisplayMedia" exists.
func (this MediaDevices) HasGetDisplayMedia() bool {
	return js.True == bindings.HasMediaDevicesGetDisplayMedia(
		this.Ref(),
	)
}

// GetDisplayMediaFunc returns the method "MediaDevices.getDisplayMedia".
func (this MediaDevices) GetDisplayMediaFunc() (fn js.Func[func(options DisplayMediaStreamOptions) js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetDisplayMediaFunc(
			this.Ref(),
		),
	)
}

// GetDisplayMedia calls the method "MediaDevices.getDisplayMedia".
func (this MediaDevices) GetDisplayMedia(options DisplayMediaStreamOptions) (ret js.Promise[MediaStream]) {
	bindings.CallMediaDevicesGetDisplayMedia(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetDisplayMedia calls the method "MediaDevices.getDisplayMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TryGetDisplayMedia(options DisplayMediaStreamOptions) (ret js.Promise[MediaStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesGetDisplayMedia(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasGetDisplayMedia1 returns true if the method "MediaDevices.getDisplayMedia" exists.
func (this MediaDevices) HasGetDisplayMedia1() bool {
	return js.True == bindings.HasMediaDevicesGetDisplayMedia1(
		this.Ref(),
	)
}

// GetDisplayMedia1Func returns the method "MediaDevices.getDisplayMedia".
func (this MediaDevices) GetDisplayMedia1Func() (fn js.Func[func() js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetDisplayMedia1Func(
			this.Ref(),
		),
	)
}

// GetDisplayMedia1 calls the method "MediaDevices.getDisplayMedia".
func (this MediaDevices) GetDisplayMedia1() (ret js.Promise[MediaStream]) {
	bindings.CallMediaDevicesGetDisplayMedia1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetDisplayMedia1 calls the method "MediaDevices.getDisplayMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TryGetDisplayMedia1() (ret js.Promise[MediaStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesGetDisplayMedia1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetSupportedCaptureActions returns true if the method "MediaDevices.setSupportedCaptureActions" exists.
func (this MediaDevices) HasSetSupportedCaptureActions() bool {
	return js.True == bindings.HasMediaDevicesSetSupportedCaptureActions(
		this.Ref(),
	)
}

// SetSupportedCaptureActionsFunc returns the method "MediaDevices.setSupportedCaptureActions".
func (this MediaDevices) SetSupportedCaptureActionsFunc() (fn js.Func[func(actions js.Array[js.String])]) {
	return fn.FromRef(
		bindings.MediaDevicesSetSupportedCaptureActionsFunc(
			this.Ref(),
		),
	)
}

// SetSupportedCaptureActions calls the method "MediaDevices.setSupportedCaptureActions".
func (this MediaDevices) SetSupportedCaptureActions(actions js.Array[js.String]) (ret js.Void) {
	bindings.CallMediaDevicesSetSupportedCaptureActions(
		this.Ref(), js.Pointer(&ret),
		actions.Ref(),
	)

	return
}

// TrySetSupportedCaptureActions calls the method "MediaDevices.setSupportedCaptureActions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TrySetSupportedCaptureActions(actions js.Array[js.String]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesSetSupportedCaptureActions(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		actions.Ref(),
	)

	return
}

// HasSelectAudioOutput returns true if the method "MediaDevices.selectAudioOutput" exists.
func (this MediaDevices) HasSelectAudioOutput() bool {
	return js.True == bindings.HasMediaDevicesSelectAudioOutput(
		this.Ref(),
	)
}

// SelectAudioOutputFunc returns the method "MediaDevices.selectAudioOutput".
func (this MediaDevices) SelectAudioOutputFunc() (fn js.Func[func(options AudioOutputOptions) js.Promise[MediaDeviceInfo]]) {
	return fn.FromRef(
		bindings.MediaDevicesSelectAudioOutputFunc(
			this.Ref(),
		),
	)
}

// SelectAudioOutput calls the method "MediaDevices.selectAudioOutput".
func (this MediaDevices) SelectAudioOutput(options AudioOutputOptions) (ret js.Promise[MediaDeviceInfo]) {
	bindings.CallMediaDevicesSelectAudioOutput(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TrySelectAudioOutput calls the method "MediaDevices.selectAudioOutput"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TrySelectAudioOutput(options AudioOutputOptions) (ret js.Promise[MediaDeviceInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesSelectAudioOutput(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasSelectAudioOutput1 returns true if the method "MediaDevices.selectAudioOutput" exists.
func (this MediaDevices) HasSelectAudioOutput1() bool {
	return js.True == bindings.HasMediaDevicesSelectAudioOutput1(
		this.Ref(),
	)
}

// SelectAudioOutput1Func returns the method "MediaDevices.selectAudioOutput".
func (this MediaDevices) SelectAudioOutput1Func() (fn js.Func[func() js.Promise[MediaDeviceInfo]]) {
	return fn.FromRef(
		bindings.MediaDevicesSelectAudioOutput1Func(
			this.Ref(),
		),
	)
}

// SelectAudioOutput1 calls the method "MediaDevices.selectAudioOutput".
func (this MediaDevices) SelectAudioOutput1() (ret js.Promise[MediaDeviceInfo]) {
	bindings.CallMediaDevicesSelectAudioOutput1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySelectAudioOutput1 calls the method "MediaDevices.selectAudioOutput"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TrySelectAudioOutput1() (ret js.Promise[MediaDeviceInfo], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesSelectAudioOutput1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCaptureHandleConfig returns true if the method "MediaDevices.setCaptureHandleConfig" exists.
func (this MediaDevices) HasSetCaptureHandleConfig() bool {
	return js.True == bindings.HasMediaDevicesSetCaptureHandleConfig(
		this.Ref(),
	)
}

// SetCaptureHandleConfigFunc returns the method "MediaDevices.setCaptureHandleConfig".
func (this MediaDevices) SetCaptureHandleConfigFunc() (fn js.Func[func(config CaptureHandleConfig)]) {
	return fn.FromRef(
		bindings.MediaDevicesSetCaptureHandleConfigFunc(
			this.Ref(),
		),
	)
}

// SetCaptureHandleConfig calls the method "MediaDevices.setCaptureHandleConfig".
func (this MediaDevices) SetCaptureHandleConfig(config CaptureHandleConfig) (ret js.Void) {
	bindings.CallMediaDevicesSetCaptureHandleConfig(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TrySetCaptureHandleConfig calls the method "MediaDevices.setCaptureHandleConfig"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TrySetCaptureHandleConfig(config CaptureHandleConfig) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesSetCaptureHandleConfig(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

// HasSetCaptureHandleConfig1 returns true if the method "MediaDevices.setCaptureHandleConfig" exists.
func (this MediaDevices) HasSetCaptureHandleConfig1() bool {
	return js.True == bindings.HasMediaDevicesSetCaptureHandleConfig1(
		this.Ref(),
	)
}

// SetCaptureHandleConfig1Func returns the method "MediaDevices.setCaptureHandleConfig".
func (this MediaDevices) SetCaptureHandleConfig1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaDevicesSetCaptureHandleConfig1Func(
			this.Ref(),
		),
	)
}

// SetCaptureHandleConfig1 calls the method "MediaDevices.setCaptureHandleConfig".
func (this MediaDevices) SetCaptureHandleConfig1() (ret js.Void) {
	bindings.CallMediaDevicesSetCaptureHandleConfig1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetCaptureHandleConfig1 calls the method "MediaDevices.setCaptureHandleConfig"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TrySetCaptureHandleConfig1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesSetCaptureHandleConfig1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetViewportMedia returns true if the method "MediaDevices.getViewportMedia" exists.
func (this MediaDevices) HasGetViewportMedia() bool {
	return js.True == bindings.HasMediaDevicesGetViewportMedia(
		this.Ref(),
	)
}

// GetViewportMediaFunc returns the method "MediaDevices.getViewportMedia".
func (this MediaDevices) GetViewportMediaFunc() (fn js.Func[func(constraints ViewportMediaStreamConstraints) js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetViewportMediaFunc(
			this.Ref(),
		),
	)
}

// GetViewportMedia calls the method "MediaDevices.getViewportMedia".
func (this MediaDevices) GetViewportMedia(constraints ViewportMediaStreamConstraints) (ret js.Promise[MediaStream]) {
	bindings.CallMediaDevicesGetViewportMedia(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&constraints),
	)

	return
}

// TryGetViewportMedia calls the method "MediaDevices.getViewportMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TryGetViewportMedia(constraints ViewportMediaStreamConstraints) (ret js.Promise[MediaStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesGetViewportMedia(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&constraints),
	)

	return
}

// HasGetViewportMedia1 returns true if the method "MediaDevices.getViewportMedia" exists.
func (this MediaDevices) HasGetViewportMedia1() bool {
	return js.True == bindings.HasMediaDevicesGetViewportMedia1(
		this.Ref(),
	)
}

// GetViewportMedia1Func returns the method "MediaDevices.getViewportMedia".
func (this MediaDevices) GetViewportMedia1Func() (fn js.Func[func() js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetViewportMedia1Func(
			this.Ref(),
		),
	)
}

// GetViewportMedia1 calls the method "MediaDevices.getViewportMedia".
func (this MediaDevices) GetViewportMedia1() (ret js.Promise[MediaStream]) {
	bindings.CallMediaDevicesGetViewportMedia1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetViewportMedia1 calls the method "MediaDevices.getViewportMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TryGetViewportMedia1() (ret js.Promise[MediaStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesGetViewportMedia1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetSupportedConstraints returns true if the method "MediaDevices.getSupportedConstraints" exists.
func (this MediaDevices) HasGetSupportedConstraints() bool {
	return js.True == bindings.HasMediaDevicesGetSupportedConstraints(
		this.Ref(),
	)
}

// GetSupportedConstraintsFunc returns the method "MediaDevices.getSupportedConstraints".
func (this MediaDevices) GetSupportedConstraintsFunc() (fn js.Func[func() MediaTrackSupportedConstraints]) {
	return fn.FromRef(
		bindings.MediaDevicesGetSupportedConstraintsFunc(
			this.Ref(),
		),
	)
}

// GetSupportedConstraints calls the method "MediaDevices.getSupportedConstraints".
func (this MediaDevices) GetSupportedConstraints() (ret MediaTrackSupportedConstraints) {
	bindings.CallMediaDevicesGetSupportedConstraints(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSupportedConstraints calls the method "MediaDevices.getSupportedConstraints"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TryGetSupportedConstraints() (ret MediaTrackSupportedConstraints, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesGetSupportedConstraints(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetUserMedia returns true if the method "MediaDevices.getUserMedia" exists.
func (this MediaDevices) HasGetUserMedia() bool {
	return js.True == bindings.HasMediaDevicesGetUserMedia(
		this.Ref(),
	)
}

// GetUserMediaFunc returns the method "MediaDevices.getUserMedia".
func (this MediaDevices) GetUserMediaFunc() (fn js.Func[func(constraints MediaStreamConstraints) js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetUserMediaFunc(
			this.Ref(),
		),
	)
}

// GetUserMedia calls the method "MediaDevices.getUserMedia".
func (this MediaDevices) GetUserMedia(constraints MediaStreamConstraints) (ret js.Promise[MediaStream]) {
	bindings.CallMediaDevicesGetUserMedia(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&constraints),
	)

	return
}

// TryGetUserMedia calls the method "MediaDevices.getUserMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TryGetUserMedia(constraints MediaStreamConstraints) (ret js.Promise[MediaStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesGetUserMedia(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&constraints),
	)

	return
}

// HasGetUserMedia1 returns true if the method "MediaDevices.getUserMedia" exists.
func (this MediaDevices) HasGetUserMedia1() bool {
	return js.True == bindings.HasMediaDevicesGetUserMedia1(
		this.Ref(),
	)
}

// GetUserMedia1Func returns the method "MediaDevices.getUserMedia".
func (this MediaDevices) GetUserMedia1Func() (fn js.Func[func() js.Promise[MediaStream]]) {
	return fn.FromRef(
		bindings.MediaDevicesGetUserMedia1Func(
			this.Ref(),
		),
	)
}

// GetUserMedia1 calls the method "MediaDevices.getUserMedia".
func (this MediaDevices) GetUserMedia1() (ret js.Promise[MediaStream]) {
	bindings.CallMediaDevicesGetUserMedia1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetUserMedia1 calls the method "MediaDevices.getUserMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaDevices) TryGetUserMedia1() (ret js.Promise[MediaStream], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaDevicesGetUserMedia1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SerialPortInfo struct {
	// UsbVendorId is "SerialPortInfo.usbVendorId"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsbVendorId MUST be set to true to make this field effective.
	UsbVendorId uint16
	// UsbProductId is "SerialPortInfo.usbProductId"
	//
	// Optional
	//
	// NOTE: FFI_USE_UsbProductId MUST be set to true to make this field effective.
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

// New creates a new SerialPortInfo in the application heap.
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
