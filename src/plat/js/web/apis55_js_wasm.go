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

type Transformer struct {
	// Start is "Transformer.start"
	//
	// Optional
	Start js.Func[func(controller TransformStreamDefaultController) js.Any]
	// Transform is "Transformer.transform"
	//
	// Optional
	Transform js.Func[func(chunk js.Any, controller TransformStreamDefaultController) js.Promise[js.Void]]
	// Flush is "Transformer.flush"
	//
	// Optional
	Flush js.Func[func(controller TransformStreamDefaultController) js.Promise[js.Void]]
	// ReadableType is "Transformer.readableType"
	//
	// Optional
	ReadableType js.Any
	// WritableType is "Transformer.writableType"
	//
	// Optional
	WritableType js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Transformer with all fields set.
func (p Transformer) FromRef(ref js.Ref) Transformer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Transformer in the application heap.
func (p Transformer) New() js.Ref {
	return bindings.TransformerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p Transformer) UpdateFrom(ref js.Ref) {
	bindings.TransformerJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p Transformer) Update(ref js.Ref) {
	bindings.TransformerJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type TransitionEventInit struct {
	// PropertyName is "TransitionEventInit.propertyName"
	//
	// Optional, defaults to "".
	PropertyName js.String
	// ElapsedTime is "TransitionEventInit.elapsedTime"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_ElapsedTime MUST be set to true to make this field effective.
	ElapsedTime float64
	// PseudoElement is "TransitionEventInit.pseudoElement"
	//
	// Optional, defaults to "".
	PseudoElement js.String
	// Bubbles is "TransitionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "TransitionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "TransitionEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_ElapsedTime bool // for ElapsedTime.
	FFI_USE_Bubbles     bool // for Bubbles.
	FFI_USE_Cancelable  bool // for Cancelable.
	FFI_USE_Composed    bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TransitionEventInit with all fields set.
func (p TransitionEventInit) FromRef(ref js.Ref) TransitionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TransitionEventInit in the application heap.
func (p TransitionEventInit) New() js.Ref {
	return bindings.TransitionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p TransitionEventInit) UpdateFrom(ref js.Ref) {
	bindings.TransitionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p TransitionEventInit) Update(ref js.Ref) {
	bindings.TransitionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewTransitionEvent(typ js.String, transitionEventInitDict TransitionEventInit) TransitionEvent {
	return TransitionEvent{}.FromRef(
		bindings.NewTransitionEventByTransitionEvent(
			typ.Ref(),
			js.Pointer(&transitionEventInitDict)),
	)
}

func NewTransitionEventByTransitionEvent1(typ js.String) TransitionEvent {
	return TransitionEvent{}.FromRef(
		bindings.NewTransitionEventByTransitionEvent1(
			typ.Ref()),
	)
}

type TransitionEvent struct {
	Event
}

func (this TransitionEvent) Once() TransitionEvent {
	this.Ref().Once()
	return this
}

func (this TransitionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this TransitionEvent) FromRef(ref js.Ref) TransitionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this TransitionEvent) Free() {
	this.Ref().Free()
}

// PropertyName returns the value of property "TransitionEvent.propertyName".
//
// The returned bool will be false if there is no such property.
func (this TransitionEvent) PropertyName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTransitionEventPropertyName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ElapsedTime returns the value of property "TransitionEvent.elapsedTime".
//
// The returned bool will be false if there is no such property.
func (this TransitionEvent) ElapsedTime() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTransitionEventElapsedTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PseudoElement returns the value of property "TransitionEvent.pseudoElement".
//
// The returned bool will be false if there is no such property.
func (this TransitionEvent) PseudoElement() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTransitionEventPseudoElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type OneOf_TrustedHTML_TrustedScript_TrustedScriptURL struct {
	ref js.Ref
}

func (x OneOf_TrustedHTML_TrustedScript_TrustedScriptURL) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TrustedHTML_TrustedScript_TrustedScriptURL) Free() {
	x.ref.Free()
}

func (x OneOf_TrustedHTML_TrustedScript_TrustedScriptURL) FromRef(ref js.Ref) OneOf_TrustedHTML_TrustedScript_TrustedScriptURL {
	return OneOf_TrustedHTML_TrustedScript_TrustedScriptURL{
		ref: ref,
	}
}

func (x OneOf_TrustedHTML_TrustedScript_TrustedScriptURL) TrustedHTML() TrustedHTML {
	return TrustedHTML{}.FromRef(x.ref)
}

func (x OneOf_TrustedHTML_TrustedScript_TrustedScriptURL) TrustedScript() TrustedScript {
	return TrustedScript{}.FromRef(x.ref)
}

func (x OneOf_TrustedHTML_TrustedScript_TrustedScriptURL) TrustedScriptURL() TrustedScriptURL {
	return TrustedScriptURL{}.FromRef(x.ref)
}

type TrustedType = OneOf_TrustedHTML_TrustedScript_TrustedScriptURL

type UIEventInit struct {
	// View is "UIEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "UIEventInit.detail"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "UIEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "UIEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "UIEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool
	// SourceCapabilities is "UIEventInit.sourceCapabilities"
	//
	// Optional, defaults to null.
	SourceCapabilities InputDeviceCapabilities
	// Which is "UIEventInit.which"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Which MUST be set to true to make this field effective.
	Which uint32

	FFI_USE_Detail     bool // for Detail.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.
	FFI_USE_Which      bool // for Which.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UIEventInit with all fields set.
func (p UIEventInit) FromRef(ref js.Ref) UIEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UIEventInit in the application heap.
func (p UIEventInit) New() js.Ref {
	return bindings.UIEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p UIEventInit) UpdateFrom(ref js.Ref) {
	bindings.UIEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p UIEventInit) Update(ref js.Ref) {
	bindings.UIEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewUIEvent(typ js.String, eventInitDict UIEventInit) UIEvent {
	return UIEvent{}.FromRef(
		bindings.NewUIEventByUIEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewUIEventByUIEvent1(typ js.String) UIEvent {
	return UIEvent{}.FromRef(
		bindings.NewUIEventByUIEvent1(
			typ.Ref()),
	)
}

type UIEvent struct {
	Event
}

func (this UIEvent) Once() UIEvent {
	this.Ref().Once()
	return this
}

func (this UIEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this UIEvent) FromRef(ref js.Ref) UIEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this UIEvent) Free() {
	this.Ref().Free()
}

// View returns the value of property "UIEvent.view".
//
// The returned bool will be false if there is no such property.
func (this UIEvent) View() (Window, bool) {
	var _ok bool
	_ret := bindings.GetUIEventView(
		this.Ref(), js.Pointer(&_ok),
	)
	return Window{}.FromRef(_ret), _ok
}

// Detail returns the value of property "UIEvent.detail".
//
// The returned bool will be false if there is no such property.
func (this UIEvent) Detail() (int32, bool) {
	var _ok bool
	_ret := bindings.GetUIEventDetail(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Which returns the value of property "UIEvent.which".
//
// The returned bool will be false if there is no such property.
func (this UIEvent) Which() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetUIEventWhich(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SourceCapabilities returns the value of property "UIEvent.sourceCapabilities".
//
// The returned bool will be false if there is no such property.
func (this UIEvent) SourceCapabilities() (InputDeviceCapabilities, bool) {
	var _ok bool
	_ret := bindings.GetUIEventSourceCapabilities(
		this.Ref(), js.Pointer(&_ok),
	)
	return InputDeviceCapabilities{}.FromRef(_ret), _ok
}

// InitUIEvent calls the method "UIEvent.initUIEvent".
//
// The returned bool will be false if there is no such method.
func (this UIEvent) InitUIEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallUIEventInitUIEvent(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitUIEventFunc returns the method "UIEvent.initUIEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this UIEvent) InitUIEventFunc() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32)]) {
	return fn.FromRef(
		bindings.UIEventInitUIEventFunc(
			this.Ref(),
		),
	)
}

// InitUIEvent1 calls the method "UIEvent.initUIEvent".
//
// The returned bool will be false if there is no such method.
func (this UIEvent) InitUIEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallUIEventInitUIEvent1(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitUIEvent1Func returns the method "UIEvent.initUIEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this UIEvent) InitUIEvent1Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window)]) {
	return fn.FromRef(
		bindings.UIEventInitUIEvent1Func(
			this.Ref(),
		),
	)
}

// InitUIEvent2 calls the method "UIEvent.initUIEvent".
//
// The returned bool will be false if there is no such method.
func (this UIEvent) InitUIEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallUIEventInitUIEvent2(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitUIEvent2Func returns the method "UIEvent.initUIEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this UIEvent) InitUIEvent2Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	return fn.FromRef(
		bindings.UIEventInitUIEvent2Func(
			this.Ref(),
		),
	)
}

// InitUIEvent3 calls the method "UIEvent.initUIEvent".
//
// The returned bool will be false if there is no such method.
func (this UIEvent) InitUIEvent3(typeArg js.String, bubblesArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallUIEventInitUIEvent3(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitUIEvent3Func returns the method "UIEvent.initUIEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this UIEvent) InitUIEvent3Func() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	return fn.FromRef(
		bindings.UIEventInitUIEvent3Func(
			this.Ref(),
		),
	)
}

// InitUIEvent4 calls the method "UIEvent.initUIEvent".
//
// The returned bool will be false if there is no such method.
func (this UIEvent) InitUIEvent4(typeArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallUIEventInitUIEvent4(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitUIEvent4Func returns the method "UIEvent.initUIEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this UIEvent) InitUIEvent4Func() (fn js.Func[func(typeArg js.String)]) {
	return fn.FromRef(
		bindings.UIEventInitUIEvent4Func(
			this.Ref(),
		),
	)
}

type OneOf_Blob_MediaSource struct {
	ref js.Ref
}

func (x OneOf_Blob_MediaSource) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Blob_MediaSource) Free() {
	x.ref.Free()
}

func (x OneOf_Blob_MediaSource) FromRef(ref js.Ref) OneOf_Blob_MediaSource {
	return OneOf_Blob_MediaSource{
		ref: ref,
	}
}

func (x OneOf_Blob_MediaSource) Blob() Blob {
	return Blob{}.FromRef(x.ref)
}

func (x OneOf_Blob_MediaSource) MediaSource() MediaSource {
	return MediaSource{}.FromRef(x.ref)
}

func NewURL(url js.String, base js.String) URL {
	return URL{}.FromRef(
		bindings.NewURLByURL(
			url.Ref(),
			base.Ref()),
	)
}

func NewURLByURL1(url js.String) URL {
	return URL{}.FromRef(
		bindings.NewURLByURL1(
			url.Ref()),
	)
}

type URL struct {
	ref js.Ref
}

func (this URL) Once() URL {
	this.Ref().Once()
	return this
}

func (this URL) Ref() js.Ref {
	return this.ref
}

func (this URL) FromRef(ref js.Ref) URL {
	this.ref = ref
	return this
}

func (this URL) Free() {
	this.Ref().Free()
}

// Href returns the value of property "URL.href".
//
// The returned bool will be false if there is no such property.
func (this URL) Href() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Href sets the value of property "URL.href" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetHref(val js.String) bool {
	return js.True == bindings.SetURLHref(
		this.Ref(),
		val.Ref(),
	)
}

// Origin returns the value of property "URL.origin".
//
// The returned bool will be false if there is no such property.
func (this URL) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol returns the value of property "URL.protocol".
//
// The returned bool will be false if there is no such property.
func (this URL) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol sets the value of property "URL.protocol" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetProtocol(val js.String) bool {
	return js.True == bindings.SetURLProtocol(
		this.Ref(),
		val.Ref(),
	)
}

// Username returns the value of property "URL.username".
//
// The returned bool will be false if there is no such property.
func (this URL) Username() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLUsername(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Username sets the value of property "URL.username" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetUsername(val js.String) bool {
	return js.True == bindings.SetURLUsername(
		this.Ref(),
		val.Ref(),
	)
}

// Password returns the value of property "URL.password".
//
// The returned bool will be false if there is no such property.
func (this URL) Password() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPassword(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Password sets the value of property "URL.password" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetPassword(val js.String) bool {
	return js.True == bindings.SetURLPassword(
		this.Ref(),
		val.Ref(),
	)
}

// Host returns the value of property "URL.host".
//
// The returned bool will be false if there is no such property.
func (this URL) Host() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLHost(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Host sets the value of property "URL.host" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetHost(val js.String) bool {
	return js.True == bindings.SetURLHost(
		this.Ref(),
		val.Ref(),
	)
}

// Hostname returns the value of property "URL.hostname".
//
// The returned bool will be false if there is no such property.
func (this URL) Hostname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLHostname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hostname sets the value of property "URL.hostname" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetHostname(val js.String) bool {
	return js.True == bindings.SetURLHostname(
		this.Ref(),
		val.Ref(),
	)
}

// Port returns the value of property "URL.port".
//
// The returned bool will be false if there is no such property.
func (this URL) Port() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Port sets the value of property "URL.port" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetPort(val js.String) bool {
	return js.True == bindings.SetURLPort(
		this.Ref(),
		val.Ref(),
	)
}

// Pathname returns the value of property "URL.pathname".
//
// The returned bool will be false if there is no such property.
func (this URL) Pathname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPathname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Pathname sets the value of property "URL.pathname" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetPathname(val js.String) bool {
	return js.True == bindings.SetURLPathname(
		this.Ref(),
		val.Ref(),
	)
}

// Search returns the value of property "URL.search".
//
// The returned bool will be false if there is no such property.
func (this URL) Search() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLSearch(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Search sets the value of property "URL.search" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetSearch(val js.String) bool {
	return js.True == bindings.SetURLSearch(
		this.Ref(),
		val.Ref(),
	)
}

// SearchParams returns the value of property "URL.searchParams".
//
// The returned bool will be false if there is no such property.
func (this URL) SearchParams() (URLSearchParams, bool) {
	var _ok bool
	_ret := bindings.GetURLSearchParams(
		this.Ref(), js.Pointer(&_ok),
	)
	return URLSearchParams{}.FromRef(_ret), _ok
}

// Hash returns the value of property "URL.hash".
//
// The returned bool will be false if there is no such property.
func (this URL) Hash() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLHash(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hash sets the value of property "URL.hash" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetHash(val js.String) bool {
	return js.True == bindings.SetURLHash(
		this.Ref(),
		val.Ref(),
	)
}

// CanParse calls the staticmethod "URL.canParse".
//
// The returned bool will be false if there is no such method.
func (this URL) CanParse(url js.String, base js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallURLCanParse(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
		base.Ref(),
	)

	return _ret == js.True, _ok
}

// CanParseFunc returns the staticmethod "URL.canParse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URL) CanParseFunc() (fn js.Func[func(url js.String, base js.String) bool]) {
	return fn.FromRef(
		bindings.URLCanParseFunc(
			this.Ref(),
		),
	)
}

// CanParse1 calls the staticmethod "URL.canParse".
//
// The returned bool will be false if there is no such method.
func (this URL) CanParse1(url js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallURLCanParse1(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	return _ret == js.True, _ok
}

// CanParse1Func returns the staticmethod "URL.canParse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URL) CanParse1Func() (fn js.Func[func(url js.String) bool]) {
	return fn.FromRef(
		bindings.URLCanParse1Func(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "URL.toJSON".
//
// The returned bool will be false if there is no such method.
func (this URL) ToJSON() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallURLToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "URL.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URL) ToJSONFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.URLToJSONFunc(
			this.Ref(),
		),
	)
}

// CreateObjectURL calls the staticmethod "URL.createObjectURL".
//
// The returned bool will be false if there is no such method.
func (this URL) CreateObjectURL(obj OneOf_Blob_MediaSource) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallURLCreateObjectURL(
		this.Ref(), js.Pointer(&_ok),
		obj.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// CreateObjectURLFunc returns the staticmethod "URL.createObjectURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URL) CreateObjectURLFunc() (fn js.Func[func(obj OneOf_Blob_MediaSource) js.String]) {
	return fn.FromRef(
		bindings.URLCreateObjectURLFunc(
			this.Ref(),
		),
	)
}

// RevokeObjectURL calls the staticmethod "URL.revokeObjectURL".
//
// The returned bool will be false if there is no such method.
func (this URL) RevokeObjectURL(url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallURLRevokeObjectURL(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RevokeObjectURLFunc returns the staticmethod "URL.revokeObjectURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URL) RevokeObjectURLFunc() (fn js.Func[func(url js.String)]) {
	return fn.FromRef(
		bindings.URLRevokeObjectURLFunc(
			this.Ref(),
		),
	)
}

type URLPatternInit struct {
	// Protocol is "URLPatternInit.protocol"
	//
	// Optional
	Protocol js.String
	// Username is "URLPatternInit.username"
	//
	// Optional
	Username js.String
	// Password is "URLPatternInit.password"
	//
	// Optional
	Password js.String
	// Hostname is "URLPatternInit.hostname"
	//
	// Optional
	Hostname js.String
	// Port is "URLPatternInit.port"
	//
	// Optional
	Port js.String
	// Pathname is "URLPatternInit.pathname"
	//
	// Optional
	Pathname js.String
	// Search is "URLPatternInit.search"
	//
	// Optional
	Search js.String
	// Hash is "URLPatternInit.hash"
	//
	// Optional
	Hash js.String
	// BaseURL is "URLPatternInit.baseURL"
	//
	// Optional
	BaseURL js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a URLPatternInit with all fields set.
func (p URLPatternInit) FromRef(ref js.Ref) URLPatternInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new URLPatternInit in the application heap.
func (p URLPatternInit) New() js.Ref {
	return bindings.URLPatternInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p URLPatternInit) UpdateFrom(ref js.Ref) {
	bindings.URLPatternInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p URLPatternInit) Update(ref js.Ref) {
	bindings.URLPatternInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_String_URLPatternInit struct {
	ref js.Ref
}

func (x OneOf_String_URLPatternInit) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_URLPatternInit) Free() {
	x.ref.Free()
}

func (x OneOf_String_URLPatternInit) FromRef(ref js.Ref) OneOf_String_URLPatternInit {
	return OneOf_String_URLPatternInit{
		ref: ref,
	}
}

func (x OneOf_String_URLPatternInit) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_URLPatternInit) URLPatternInit() URLPatternInit {
	var ret URLPatternInit
	ret.UpdateFrom(x.ref)
	return ret
}

type URLPatternInput = OneOf_String_URLPatternInit

type URLPatternOptions struct {
	// IgnoreCase is "URLPatternOptions.ignoreCase"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreCase MUST be set to true to make this field effective.
	IgnoreCase bool

	FFI_USE_IgnoreCase bool // for IgnoreCase.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a URLPatternOptions with all fields set.
func (p URLPatternOptions) FromRef(ref js.Ref) URLPatternOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new URLPatternOptions in the application heap.
func (p URLPatternOptions) New() js.Ref {
	return bindings.URLPatternOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p URLPatternOptions) UpdateFrom(ref js.Ref) {
	bindings.URLPatternOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p URLPatternOptions) Update(ref js.Ref) {
	bindings.URLPatternOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_String_undefined struct {
	ref js.Ref
}

func (x OneOf_String_undefined) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_undefined) Free() {
	x.ref.Free()
}

func (x OneOf_String_undefined) FromRef(ref js.Ref) OneOf_String_undefined {
	return OneOf_String_undefined{
		ref: ref,
	}
}

func (x OneOf_String_undefined) Undefined() bool {
	return x.ref == js.Undefined
}

func (x OneOf_String_undefined) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type URLPatternComponentResult struct {
	// Input is "URLPatternComponentResult.input"
	//
	// Optional
	Input js.String
	// Groups is "URLPatternComponentResult.groups"
	//
	// Optional
	Groups js.Record[OneOf_String_undefined]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a URLPatternComponentResult with all fields set.
func (p URLPatternComponentResult) FromRef(ref js.Ref) URLPatternComponentResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new URLPatternComponentResult in the application heap.
func (p URLPatternComponentResult) New() js.Ref {
	return bindings.URLPatternComponentResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p URLPatternComponentResult) UpdateFrom(ref js.Ref) {
	bindings.URLPatternComponentResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p URLPatternComponentResult) Update(ref js.Ref) {
	bindings.URLPatternComponentResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type URLPatternResult struct {
	// Inputs is "URLPatternResult.inputs"
	//
	// Optional
	Inputs js.Array[URLPatternInput]
	// Protocol is "URLPatternResult.protocol"
	//
	// Optional
	Protocol URLPatternComponentResult
	// Username is "URLPatternResult.username"
	//
	// Optional
	Username URLPatternComponentResult
	// Password is "URLPatternResult.password"
	//
	// Optional
	Password URLPatternComponentResult
	// Hostname is "URLPatternResult.hostname"
	//
	// Optional
	Hostname URLPatternComponentResult
	// Port is "URLPatternResult.port"
	//
	// Optional
	Port URLPatternComponentResult
	// Pathname is "URLPatternResult.pathname"
	//
	// Optional
	Pathname URLPatternComponentResult
	// Search is "URLPatternResult.search"
	//
	// Optional
	Search URLPatternComponentResult
	// Hash is "URLPatternResult.hash"
	//
	// Optional
	Hash URLPatternComponentResult

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a URLPatternResult with all fields set.
func (p URLPatternResult) FromRef(ref js.Ref) URLPatternResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new URLPatternResult in the application heap.
func (p URLPatternResult) New() js.Ref {
	return bindings.URLPatternResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p URLPatternResult) UpdateFrom(ref js.Ref) {
	bindings.URLPatternResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p URLPatternResult) Update(ref js.Ref) {
	bindings.URLPatternResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewURLPattern(input URLPatternInput, baseURL js.String, options URLPatternOptions) URLPattern {
	return URLPattern{}.FromRef(
		bindings.NewURLPatternByURLPattern(
			input.Ref(),
			baseURL.Ref(),
			js.Pointer(&options)),
	)
}

func NewURLPatternByURLPattern1(input URLPatternInput, baseURL js.String) URLPattern {
	return URLPattern{}.FromRef(
		bindings.NewURLPatternByURLPattern1(
			input.Ref(),
			baseURL.Ref()),
	)
}

func NewURLPatternByURLPattern2(input URLPatternInput, options URLPatternOptions) URLPattern {
	return URLPattern{}.FromRef(
		bindings.NewURLPatternByURLPattern2(
			input.Ref(),
			js.Pointer(&options)),
	)
}

func NewURLPatternByURLPattern3(input URLPatternInput) URLPattern {
	return URLPattern{}.FromRef(
		bindings.NewURLPatternByURLPattern3(
			input.Ref()),
	)
}

func NewURLPatternByURLPattern4() URLPattern {
	return URLPattern{}.FromRef(
		bindings.NewURLPatternByURLPattern4(),
	)
}

type URLPattern struct {
	ref js.Ref
}

func (this URLPattern) Once() URLPattern {
	this.Ref().Once()
	return this
}

func (this URLPattern) Ref() js.Ref {
	return this.ref
}

func (this URLPattern) FromRef(ref js.Ref) URLPattern {
	this.ref = ref
	return this
}

func (this URLPattern) Free() {
	this.Ref().Free()
}

// Protocol returns the value of property "URLPattern.protocol".
//
// The returned bool will be false if there is no such property.
func (this URLPattern) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPatternProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Username returns the value of property "URLPattern.username".
//
// The returned bool will be false if there is no such property.
func (this URLPattern) Username() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPatternUsername(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Password returns the value of property "URLPattern.password".
//
// The returned bool will be false if there is no such property.
func (this URLPattern) Password() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPatternPassword(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hostname returns the value of property "URLPattern.hostname".
//
// The returned bool will be false if there is no such property.
func (this URLPattern) Hostname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPatternHostname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Port returns the value of property "URLPattern.port".
//
// The returned bool will be false if there is no such property.
func (this URLPattern) Port() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPatternPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Pathname returns the value of property "URLPattern.pathname".
//
// The returned bool will be false if there is no such property.
func (this URLPattern) Pathname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPatternPathname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Search returns the value of property "URLPattern.search".
//
// The returned bool will be false if there is no such property.
func (this URLPattern) Search() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPatternSearch(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hash returns the value of property "URLPattern.hash".
//
// The returned bool will be false if there is no such property.
func (this URLPattern) Hash() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetURLPatternHash(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Test calls the method "URLPattern.test".
//
// The returned bool will be false if there is no such method.
func (this URLPattern) Test(input URLPatternInput, baseURL js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallURLPatternTest(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		baseURL.Ref(),
	)

	return _ret == js.True, _ok
}

// TestFunc returns the method "URLPattern.test".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URLPattern) TestFunc() (fn js.Func[func(input URLPatternInput, baseURL js.String) bool]) {
	return fn.FromRef(
		bindings.URLPatternTestFunc(
			this.Ref(),
		),
	)
}

// Test1 calls the method "URLPattern.test".
//
// The returned bool will be false if there is no such method.
func (this URLPattern) Test1(input URLPatternInput) (bool, bool) {
	var _ok bool
	_ret := bindings.CallURLPatternTest1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return _ret == js.True, _ok
}

// Test1Func returns the method "URLPattern.test".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URLPattern) Test1Func() (fn js.Func[func(input URLPatternInput) bool]) {
	return fn.FromRef(
		bindings.URLPatternTest1Func(
			this.Ref(),
		),
	)
}

// Test2 calls the method "URLPattern.test".
//
// The returned bool will be false if there is no such method.
func (this URLPattern) Test2() (bool, bool) {
	var _ok bool
	_ret := bindings.CallURLPatternTest2(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// Test2Func returns the method "URLPattern.test".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URLPattern) Test2Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.URLPatternTest2Func(
			this.Ref(),
		),
	)
}

// Exec calls the method "URLPattern.exec".
//
// The returned bool will be false if there is no such method.
func (this URLPattern) Exec(input URLPatternInput, baseURL js.String) (URLPatternResult, bool) {
	var _ret URLPatternResult
	_ok := js.True == bindings.CallURLPatternExec(
		this.Ref(), js.Pointer(&_ret),
		input.Ref(),
		baseURL.Ref(),
	)

	return _ret, _ok
}

// ExecFunc returns the method "URLPattern.exec".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URLPattern) ExecFunc() (fn js.Func[func(input URLPatternInput, baseURL js.String) URLPatternResult]) {
	return fn.FromRef(
		bindings.URLPatternExecFunc(
			this.Ref(),
		),
	)
}

// Exec1 calls the method "URLPattern.exec".
//
// The returned bool will be false if there is no such method.
func (this URLPattern) Exec1(input URLPatternInput) (URLPatternResult, bool) {
	var _ret URLPatternResult
	_ok := js.True == bindings.CallURLPatternExec1(
		this.Ref(), js.Pointer(&_ret),
		input.Ref(),
	)

	return _ret, _ok
}

// Exec1Func returns the method "URLPattern.exec".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URLPattern) Exec1Func() (fn js.Func[func(input URLPatternInput) URLPatternResult]) {
	return fn.FromRef(
		bindings.URLPatternExec1Func(
			this.Ref(),
		),
	)
}

// Exec2 calls the method "URLPattern.exec".
//
// The returned bool will be false if there is no such method.
func (this URLPattern) Exec2() (URLPatternResult, bool) {
	var _ret URLPatternResult
	_ok := js.True == bindings.CallURLPatternExec2(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// Exec2Func returns the method "URLPattern.exec".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this URLPattern) Exec2Func() (fn js.Func[func() URLPatternResult]) {
	return fn.FromRef(
		bindings.URLPatternExec2Func(
			this.Ref(),
		),
	)
}

type USBConnectionEventInit struct {
	// Device is "USBConnectionEventInit.device"
	//
	// Required
	Device USBDevice
	// Bubbles is "USBConnectionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "USBConnectionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "USBConnectionEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a USBConnectionEventInit with all fields set.
func (p USBConnectionEventInit) FromRef(ref js.Ref) USBConnectionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new USBConnectionEventInit in the application heap.
func (p USBConnectionEventInit) New() js.Ref {
	return bindings.USBConnectionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p USBConnectionEventInit) UpdateFrom(ref js.Ref) {
	bindings.USBConnectionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p USBConnectionEventInit) Update(ref js.Ref) {
	bindings.USBConnectionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewUSBConnectionEvent(typ js.String, eventInitDict USBConnectionEventInit) USBConnectionEvent {
	return USBConnectionEvent{}.FromRef(
		bindings.NewUSBConnectionEventByUSBConnectionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type USBConnectionEvent struct {
	Event
}

func (this USBConnectionEvent) Once() USBConnectionEvent {
	this.Ref().Once()
	return this
}

func (this USBConnectionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this USBConnectionEvent) FromRef(ref js.Ref) USBConnectionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this USBConnectionEvent) Free() {
	this.Ref().Free()
}

// Device returns the value of property "USBConnectionEvent.device".
//
// The returned bool will be false if there is no such property.
func (this USBConnectionEvent) Device() (USBDevice, bool) {
	var _ok bool
	_ret := bindings.GetUSBConnectionEventDevice(
		this.Ref(), js.Pointer(&_ok),
	)
	return USBDevice{}.FromRef(_ret), _ok
}

type USBPermissionDescriptor struct {
	// Filters is "USBPermissionDescriptor.filters"
	//
	// Optional
	Filters js.Array[USBDeviceFilter]
	// ExclusionFilters is "USBPermissionDescriptor.exclusionFilters"
	//
	// Optional
	ExclusionFilters js.Array[USBDeviceFilter]
	// Name is "USBPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a USBPermissionDescriptor with all fields set.
func (p USBPermissionDescriptor) FromRef(ref js.Ref) USBPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new USBPermissionDescriptor in the application heap.
func (p USBPermissionDescriptor) New() js.Ref {
	return bindings.USBPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p USBPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.USBPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p USBPermissionDescriptor) Update(ref js.Ref) {
	bindings.USBPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type USBPermissionResult struct {
	PermissionStatus
}

func (this USBPermissionResult) Once() USBPermissionResult {
	this.Ref().Once()
	return this
}

func (this USBPermissionResult) Ref() js.Ref {
	return this.PermissionStatus.Ref()
}

func (this USBPermissionResult) FromRef(ref js.Ref) USBPermissionResult {
	this.PermissionStatus = this.PermissionStatus.FromRef(ref)
	return this
}

func (this USBPermissionResult) Free() {
	this.Ref().Free()
}

// Devices returns the value of property "USBPermissionResult.devices".
//
// The returned bool will be false if there is no such property.
func (this USBPermissionResult) Devices() (js.FrozenArray[USBDevice], bool) {
	var _ok bool
	_ret := bindings.GetUSBPermissionResultDevices(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[USBDevice]{}.FromRef(_ret), _ok
}

// Devices sets the value of property "USBPermissionResult.devices" to val.
//
// It returns false if the property cannot be set.
func (this USBPermissionResult) SetDevices(val js.FrozenArray[USBDevice]) bool {
	return js.True == bindings.SetUSBPermissionResultDevices(
		this.Ref(),
		val.Ref(),
	)
}

type USBPermissionStorage struct {
	// AllowedDevices is "USBPermissionStorage.allowedDevices"
	//
	// Optional, defaults to [].
	AllowedDevices js.Array[AllowedUSBDevice]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a USBPermissionStorage with all fields set.
func (p USBPermissionStorage) FromRef(ref js.Ref) USBPermissionStorage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new USBPermissionStorage in the application heap.
func (p USBPermissionStorage) New() js.Ref {
	return bindings.USBPermissionStorageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p USBPermissionStorage) UpdateFrom(ref js.Ref) {
	bindings.USBPermissionStorageJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p USBPermissionStorage) Update(ref js.Ref) {
	bindings.USBPermissionStorageJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewUncalibratedMagnetometer(sensorOptions MagnetometerSensorOptions) UncalibratedMagnetometer {
	return UncalibratedMagnetometer{}.FromRef(
		bindings.NewUncalibratedMagnetometerByUncalibratedMagnetometer(
			js.Pointer(&sensorOptions)),
	)
}

func NewUncalibratedMagnetometerByUncalibratedMagnetometer1() UncalibratedMagnetometer {
	return UncalibratedMagnetometer{}.FromRef(
		bindings.NewUncalibratedMagnetometerByUncalibratedMagnetometer1(),
	)
}

type UncalibratedMagnetometer struct {
	Sensor
}

func (this UncalibratedMagnetometer) Once() UncalibratedMagnetometer {
	this.Ref().Once()
	return this
}

func (this UncalibratedMagnetometer) Ref() js.Ref {
	return this.Sensor.Ref()
}

func (this UncalibratedMagnetometer) FromRef(ref js.Ref) UncalibratedMagnetometer {
	this.Sensor = this.Sensor.FromRef(ref)
	return this
}

func (this UncalibratedMagnetometer) Free() {
	this.Ref().Free()
}

// X returns the value of property "UncalibratedMagnetometer.x".
//
// The returned bool will be false if there is no such property.
func (this UncalibratedMagnetometer) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetUncalibratedMagnetometerX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y returns the value of property "UncalibratedMagnetometer.y".
//
// The returned bool will be false if there is no such property.
func (this UncalibratedMagnetometer) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetUncalibratedMagnetometerY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Z returns the value of property "UncalibratedMagnetometer.z".
//
// The returned bool will be false if there is no such property.
func (this UncalibratedMagnetometer) Z() (float64, bool) {
	var _ok bool
	_ret := bindings.GetUncalibratedMagnetometerZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// XBias returns the value of property "UncalibratedMagnetometer.xBias".
//
// The returned bool will be false if there is no such property.
func (this UncalibratedMagnetometer) XBias() (float64, bool) {
	var _ok bool
	_ret := bindings.GetUncalibratedMagnetometerXBias(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// YBias returns the value of property "UncalibratedMagnetometer.yBias".
//
// The returned bool will be false if there is no such property.
func (this UncalibratedMagnetometer) YBias() (float64, bool) {
	var _ok bool
	_ret := bindings.GetUncalibratedMagnetometerYBias(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ZBias returns the value of property "UncalibratedMagnetometer.zBias".
//
// The returned bool will be false if there is no such property.
func (this UncalibratedMagnetometer) ZBias() (float64, bool) {
	var _ok bool
	_ret := bindings.GetUncalibratedMagnetometerZBias(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type UncalibratedMagnetometerReadingValues struct {
	// X is "UncalibratedMagnetometerReadingValues.x"
	//
	// Required
	X float64
	// Y is "UncalibratedMagnetometerReadingValues.y"
	//
	// Required
	Y float64
	// Z is "UncalibratedMagnetometerReadingValues.z"
	//
	// Required
	Z float64
	// XBias is "UncalibratedMagnetometerReadingValues.xBias"
	//
	// Required
	XBias float64
	// YBias is "UncalibratedMagnetometerReadingValues.yBias"
	//
	// Required
	YBias float64
	// ZBias is "UncalibratedMagnetometerReadingValues.zBias"
	//
	// Required
	ZBias float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UncalibratedMagnetometerReadingValues with all fields set.
func (p UncalibratedMagnetometerReadingValues) FromRef(ref js.Ref) UncalibratedMagnetometerReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UncalibratedMagnetometerReadingValues in the application heap.
func (p UncalibratedMagnetometerReadingValues) New() js.Ref {
	return bindings.UncalibratedMagnetometerReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p UncalibratedMagnetometerReadingValues) UpdateFrom(ref js.Ref) {
	bindings.UncalibratedMagnetometerReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p UncalibratedMagnetometerReadingValues) Update(ref js.Ref) {
	bindings.UncalibratedMagnetometerReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type UnderlyingSinkStartCallbackFunc func(this js.Ref, controller WritableStreamDefaultController) js.Ref

func (fn UnderlyingSinkStartCallbackFunc) Register() js.Func[func(controller WritableStreamDefaultController) js.Any] {
	return js.RegisterCallback[func(controller WritableStreamDefaultController) js.Any](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnderlyingSinkStartCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		WritableStreamDefaultController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSinkStartCallback[T any] struct {
	Fn  func(arg T, this js.Ref, controller WritableStreamDefaultController) js.Ref
	Arg T
}

func (cb *UnderlyingSinkStartCallback[T]) Register() js.Func[func(controller WritableStreamDefaultController) js.Any] {
	return js.RegisterCallback[func(controller WritableStreamDefaultController) js.Any](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnderlyingSinkStartCallback[T]) DispatchCallback(
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

		WritableStreamDefaultController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type WritableStreamDefaultController struct {
	ref js.Ref
}

func (this WritableStreamDefaultController) Once() WritableStreamDefaultController {
	this.Ref().Once()
	return this
}

func (this WritableStreamDefaultController) Ref() js.Ref {
	return this.ref
}

func (this WritableStreamDefaultController) FromRef(ref js.Ref) WritableStreamDefaultController {
	this.ref = ref
	return this
}

func (this WritableStreamDefaultController) Free() {
	this.Ref().Free()
}

// Signal returns the value of property "WritableStreamDefaultController.signal".
//
// The returned bool will be false if there is no such property.
func (this WritableStreamDefaultController) Signal() (AbortSignal, bool) {
	var _ok bool
	_ret := bindings.GetWritableStreamDefaultControllerSignal(
		this.Ref(), js.Pointer(&_ok),
	)
	return AbortSignal{}.FromRef(_ret), _ok
}

// Error calls the method "WritableStreamDefaultController.error".
//
// The returned bool will be false if there is no such method.
func (this WritableStreamDefaultController) Error(e js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamDefaultControllerError(
		this.Ref(), js.Pointer(&_ok),
		e.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ErrorFunc returns the method "WritableStreamDefaultController.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStreamDefaultController) ErrorFunc() (fn js.Func[func(e js.Any)]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultControllerErrorFunc(
			this.Ref(),
		),
	)
}

// Error1 calls the method "WritableStreamDefaultController.error".
//
// The returned bool will be false if there is no such method.
func (this WritableStreamDefaultController) Error1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallWritableStreamDefaultControllerError1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Error1Func returns the method "WritableStreamDefaultController.error".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WritableStreamDefaultController) Error1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.WritableStreamDefaultControllerError1Func(
			this.Ref(),
		),
	)
}

type UnderlyingSinkWriteCallbackFunc func(this js.Ref, chunk js.Any, controller WritableStreamDefaultController) js.Ref

func (fn UnderlyingSinkWriteCallbackFunc) Register() js.Func[func(chunk js.Any, controller WritableStreamDefaultController) js.Promise[js.Void]] {
	return js.RegisterCallback[func(chunk js.Any, controller WritableStreamDefaultController) js.Promise[js.Void]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnderlyingSinkWriteCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
		WritableStreamDefaultController{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSinkWriteCallback[T any] struct {
	Fn  func(arg T, this js.Ref, chunk js.Any, controller WritableStreamDefaultController) js.Ref
	Arg T
}

func (cb *UnderlyingSinkWriteCallback[T]) Register() js.Func[func(chunk js.Any, controller WritableStreamDefaultController) js.Promise[js.Void]] {
	return js.RegisterCallback[func(chunk js.Any, controller WritableStreamDefaultController) js.Promise[js.Void]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnderlyingSinkWriteCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
		WritableStreamDefaultController{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSinkCloseCallbackFunc func(this js.Ref) js.Ref

func (fn UnderlyingSinkCloseCallbackFunc) Register() js.Func[func() js.Promise[js.Void]] {
	return js.RegisterCallback[func() js.Promise[js.Void]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnderlyingSinkCloseCallbackFunc) DispatchCallback(
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

type UnderlyingSinkCloseCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *UnderlyingSinkCloseCallback[T]) Register() js.Func[func() js.Promise[js.Void]] {
	return js.RegisterCallback[func() js.Promise[js.Void]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnderlyingSinkCloseCallback[T]) DispatchCallback(
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

type UnderlyingSinkAbortCallbackFunc func(this js.Ref, reason js.Any) js.Ref

func (fn UnderlyingSinkAbortCallbackFunc) Register() js.Func[func(reason js.Any) js.Promise[js.Void]] {
	return js.RegisterCallback[func(reason js.Any) js.Promise[js.Void]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnderlyingSinkAbortCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSinkAbortCallback[T any] struct {
	Fn  func(arg T, this js.Ref, reason js.Any) js.Ref
	Arg T
}

func (cb *UnderlyingSinkAbortCallback[T]) Register() js.Func[func(reason js.Any) js.Promise[js.Void]] {
	return js.RegisterCallback[func(reason js.Any) js.Promise[js.Void]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnderlyingSinkAbortCallback[T]) DispatchCallback(
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

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSink struct {
	// Start is "UnderlyingSink.start"
	//
	// Optional
	Start js.Func[func(controller WritableStreamDefaultController) js.Any]
	// Write is "UnderlyingSink.write"
	//
	// Optional
	Write js.Func[func(chunk js.Any, controller WritableStreamDefaultController) js.Promise[js.Void]]
	// Close is "UnderlyingSink.close"
	//
	// Optional
	Close js.Func[func() js.Promise[js.Void]]
	// Abort is "UnderlyingSink.abort"
	//
	// Optional
	Abort js.Func[func(reason js.Any) js.Promise[js.Void]]
	// Type is "UnderlyingSink.type"
	//
	// Optional
	Type js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UnderlyingSink with all fields set.
func (p UnderlyingSink) FromRef(ref js.Ref) UnderlyingSink {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UnderlyingSink in the application heap.
func (p UnderlyingSink) New() js.Ref {
	return bindings.UnderlyingSinkJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p UnderlyingSink) UpdateFrom(ref js.Ref) {
	bindings.UnderlyingSinkJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p UnderlyingSink) Update(ref js.Ref) {
	bindings.UnderlyingSinkJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type UnderlyingSourceStartCallbackFunc func(this js.Ref, controller ReadableStreamController) js.Ref

func (fn UnderlyingSourceStartCallbackFunc) Register() js.Func[func(controller ReadableStreamController) js.Any] {
	return js.RegisterCallback[func(controller ReadableStreamController) js.Any](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnderlyingSourceStartCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		ReadableStreamController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSourceStartCallback[T any] struct {
	Fn  func(arg T, this js.Ref, controller ReadableStreamController) js.Ref
	Arg T
}

func (cb *UnderlyingSourceStartCallback[T]) Register() js.Func[func(controller ReadableStreamController) js.Any] {
	return js.RegisterCallback[func(controller ReadableStreamController) js.Any](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnderlyingSourceStartCallback[T]) DispatchCallback(
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

		ReadableStreamController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSourcePullCallbackFunc func(this js.Ref, controller ReadableStreamController) js.Ref

func (fn UnderlyingSourcePullCallbackFunc) Register() js.Func[func(controller ReadableStreamController) js.Promise[js.Void]] {
	return js.RegisterCallback[func(controller ReadableStreamController) js.Promise[js.Void]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnderlyingSourcePullCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		ReadableStreamController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSourcePullCallback[T any] struct {
	Fn  func(arg T, this js.Ref, controller ReadableStreamController) js.Ref
	Arg T
}

func (cb *UnderlyingSourcePullCallback[T]) Register() js.Func[func(controller ReadableStreamController) js.Promise[js.Void]] {
	return js.RegisterCallback[func(controller ReadableStreamController) js.Promise[js.Void]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnderlyingSourcePullCallback[T]) DispatchCallback(
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

		ReadableStreamController{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSourceCancelCallbackFunc func(this js.Ref, reason js.Any) js.Ref

func (fn UnderlyingSourceCancelCallbackFunc) Register() js.Func[func(reason js.Any) js.Promise[js.Void]] {
	return js.RegisterCallback[func(reason js.Any) js.Promise[js.Void]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnderlyingSourceCancelCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSourceCancelCallback[T any] struct {
	Fn  func(arg T, this js.Ref, reason js.Any) js.Ref
	Arg T
}

func (cb *UnderlyingSourceCancelCallback[T]) Register() js.Func[func(reason js.Any) js.Promise[js.Void]] {
	return js.RegisterCallback[func(reason js.Any) js.Promise[js.Void]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnderlyingSourceCancelCallback[T]) DispatchCallback(
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

		js.Any{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnderlyingSource struct {
	// Start is "UnderlyingSource.start"
	//
	// Optional
	Start js.Func[func(controller ReadableStreamController) js.Any]
	// Pull is "UnderlyingSource.pull"
	//
	// Optional
	Pull js.Func[func(controller ReadableStreamController) js.Promise[js.Void]]
	// Cancel is "UnderlyingSource.cancel"
	//
	// Optional
	Cancel js.Func[func(reason js.Any) js.Promise[js.Void]]
	// Type is "UnderlyingSource.type"
	//
	// Optional
	Type ReadableStreamType
	// AutoAllocateChunkSize is "UnderlyingSource.autoAllocateChunkSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoAllocateChunkSize MUST be set to true to make this field effective.
	AutoAllocateChunkSize uint64

	FFI_USE_AutoAllocateChunkSize bool // for AutoAllocateChunkSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UnderlyingSource with all fields set.
func (p UnderlyingSource) FromRef(ref js.Ref) UnderlyingSource {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UnderlyingSource in the application heap.
func (p UnderlyingSource) New() js.Ref {
	return bindings.UnderlyingSourceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p UnderlyingSource) UpdateFrom(ref js.Ref) {
	bindings.UnderlyingSourceJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p UnderlyingSource) Update(ref js.Ref) {
	bindings.UnderlyingSourceJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type UserVerificationRequirement uint32

const (
	_ UserVerificationRequirement = iota

	UserVerificationRequirement_REQUIRED
	UserVerificationRequirement_PREFERRED
	UserVerificationRequirement_DISCOURAGED
)

func (UserVerificationRequirement) FromRef(str js.Ref) UserVerificationRequirement {
	return UserVerificationRequirement(bindings.ConstOfUserVerificationRequirement(str))
}

func (x UserVerificationRequirement) String() (string, bool) {
	switch x {
	case UserVerificationRequirement_REQUIRED:
		return "required", true
	case UserVerificationRequirement_PREFERRED:
		return "preferred", true
	case UserVerificationRequirement_DISCOURAGED:
		return "discouraged", true
	default:
		return "", false
	}
}

type VTTRegion struct {
	ref js.Ref
}

func (this VTTRegion) Once() VTTRegion {
	this.Ref().Once()
	return this
}

func (this VTTRegion) Ref() js.Ref {
	return this.ref
}

func (this VTTRegion) FromRef(ref js.Ref) VTTRegion {
	this.ref = ref
	return this
}

func (this VTTRegion) Free() {
	this.Ref().Free()
}

// Id returns the value of property "VTTRegion.id".
//
// The returned bool will be false if there is no such property.
func (this VTTRegion) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetVTTRegionId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id sets the value of property "VTTRegion.id" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetId(val js.String) bool {
	return js.True == bindings.SetVTTRegionId(
		this.Ref(),
		val.Ref(),
	)
}

// Width returns the value of property "VTTRegion.width".
//
// The returned bool will be false if there is no such property.
func (this VTTRegion) Width() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVTTRegionWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Width sets the value of property "VTTRegion.width" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetWidth(val float64) bool {
	return js.True == bindings.SetVTTRegionWidth(
		this.Ref(),
		float64(val),
	)
}

// Lines returns the value of property "VTTRegion.lines".
//
// The returned bool will be false if there is no such property.
func (this VTTRegion) Lines() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVTTRegionLines(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Lines sets the value of property "VTTRegion.lines" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetLines(val uint32) bool {
	return js.True == bindings.SetVTTRegionLines(
		this.Ref(),
		uint32(val),
	)
}

// RegionAnchorX returns the value of property "VTTRegion.regionAnchorX".
//
// The returned bool will be false if there is no such property.
func (this VTTRegion) RegionAnchorX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVTTRegionRegionAnchorX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// RegionAnchorX sets the value of property "VTTRegion.regionAnchorX" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetRegionAnchorX(val float64) bool {
	return js.True == bindings.SetVTTRegionRegionAnchorX(
		this.Ref(),
		float64(val),
	)
}

// RegionAnchorY returns the value of property "VTTRegion.regionAnchorY".
//
// The returned bool will be false if there is no such property.
func (this VTTRegion) RegionAnchorY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVTTRegionRegionAnchorY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// RegionAnchorY sets the value of property "VTTRegion.regionAnchorY" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetRegionAnchorY(val float64) bool {
	return js.True == bindings.SetVTTRegionRegionAnchorY(
		this.Ref(),
		float64(val),
	)
}

// ViewportAnchorX returns the value of property "VTTRegion.viewportAnchorX".
//
// The returned bool will be false if there is no such property.
func (this VTTRegion) ViewportAnchorX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVTTRegionViewportAnchorX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ViewportAnchorX sets the value of property "VTTRegion.viewportAnchorX" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetViewportAnchorX(val float64) bool {
	return js.True == bindings.SetVTTRegionViewportAnchorX(
		this.Ref(),
		float64(val),
	)
}

// ViewportAnchorY returns the value of property "VTTRegion.viewportAnchorY".
//
// The returned bool will be false if there is no such property.
func (this VTTRegion) ViewportAnchorY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVTTRegionViewportAnchorY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ViewportAnchorY sets the value of property "VTTRegion.viewportAnchorY" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetViewportAnchorY(val float64) bool {
	return js.True == bindings.SetVTTRegionViewportAnchorY(
		this.Ref(),
		float64(val),
	)
}

// Scroll returns the value of property "VTTRegion.scroll".
//
// The returned bool will be false if there is no such property.
func (this VTTRegion) Scroll() (ScrollSetting, bool) {
	var _ok bool
	_ret := bindings.GetVTTRegionScroll(
		this.Ref(), js.Pointer(&_ok),
	)
	return ScrollSetting(_ret), _ok
}

// Scroll sets the value of property "VTTRegion.scroll" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetScroll(val ScrollSetting) bool {
	return js.True == bindings.SetVTTRegionScroll(
		this.Ref(),
		uint32(val),
	)
}

func NewVTTCue(startTime float64, endTime float64, text js.String) VTTCue {
	return VTTCue{}.FromRef(
		bindings.NewVTTCueByVTTCue(
			float64(startTime),
			float64(endTime),
			text.Ref()),
	)
}

type VTTCue struct {
	TextTrackCue
}

func (this VTTCue) Once() VTTCue {
	this.Ref().Once()
	return this
}

func (this VTTCue) Ref() js.Ref {
	return this.TextTrackCue.Ref()
}

func (this VTTCue) FromRef(ref js.Ref) VTTCue {
	this.TextTrackCue = this.TextTrackCue.FromRef(ref)
	return this
}

func (this VTTCue) Free() {
	this.Ref().Free()
}

// Region returns the value of property "VTTCue.region".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) Region() (VTTRegion, bool) {
	var _ok bool
	_ret := bindings.GetVTTCueRegion(
		this.Ref(), js.Pointer(&_ok),
	)
	return VTTRegion{}.FromRef(_ret), _ok
}

// Region sets the value of property "VTTCue.region" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetRegion(val VTTRegion) bool {
	return js.True == bindings.SetVTTCueRegion(
		this.Ref(),
		val.Ref(),
	)
}

// Vertical returns the value of property "VTTCue.vertical".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) Vertical() (DirectionSetting, bool) {
	var _ok bool
	_ret := bindings.GetVTTCueVertical(
		this.Ref(), js.Pointer(&_ok),
	)
	return DirectionSetting(_ret), _ok
}

// Vertical sets the value of property "VTTCue.vertical" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetVertical(val DirectionSetting) bool {
	return js.True == bindings.SetVTTCueVertical(
		this.Ref(),
		uint32(val),
	)
}

// SnapToLines returns the value of property "VTTCue.snapToLines".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) SnapToLines() (bool, bool) {
	var _ok bool
	_ret := bindings.GetVTTCueSnapToLines(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// SnapToLines sets the value of property "VTTCue.snapToLines" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetSnapToLines(val bool) bool {
	return js.True == bindings.SetVTTCueSnapToLines(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Line returns the value of property "VTTCue.line".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) Line() (LineAndPositionSetting, bool) {
	var _ok bool
	_ret := bindings.GetVTTCueLine(
		this.Ref(), js.Pointer(&_ok),
	)
	return LineAndPositionSetting{}.FromRef(_ret), _ok
}

// Line sets the value of property "VTTCue.line" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetLine(val LineAndPositionSetting) bool {
	return js.True == bindings.SetVTTCueLine(
		this.Ref(),
		val.Ref(),
	)
}

// LineAlign returns the value of property "VTTCue.lineAlign".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) LineAlign() (LineAlignSetting, bool) {
	var _ok bool
	_ret := bindings.GetVTTCueLineAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return LineAlignSetting(_ret), _ok
}

// LineAlign sets the value of property "VTTCue.lineAlign" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetLineAlign(val LineAlignSetting) bool {
	return js.True == bindings.SetVTTCueLineAlign(
		this.Ref(),
		uint32(val),
	)
}

// Position returns the value of property "VTTCue.position".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) Position() (LineAndPositionSetting, bool) {
	var _ok bool
	_ret := bindings.GetVTTCuePosition(
		this.Ref(), js.Pointer(&_ok),
	)
	return LineAndPositionSetting{}.FromRef(_ret), _ok
}

// Position sets the value of property "VTTCue.position" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetPosition(val LineAndPositionSetting) bool {
	return js.True == bindings.SetVTTCuePosition(
		this.Ref(),
		val.Ref(),
	)
}

// PositionAlign returns the value of property "VTTCue.positionAlign".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) PositionAlign() (PositionAlignSetting, bool) {
	var _ok bool
	_ret := bindings.GetVTTCuePositionAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return PositionAlignSetting(_ret), _ok
}

// PositionAlign sets the value of property "VTTCue.positionAlign" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetPositionAlign(val PositionAlignSetting) bool {
	return js.True == bindings.SetVTTCuePositionAlign(
		this.Ref(),
		uint32(val),
	)
}

// Size returns the value of property "VTTCue.size".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) Size() (float64, bool) {
	var _ok bool
	_ret := bindings.GetVTTCueSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Size sets the value of property "VTTCue.size" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetSize(val float64) bool {
	return js.True == bindings.SetVTTCueSize(
		this.Ref(),
		float64(val),
	)
}

// Align returns the value of property "VTTCue.align".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) Align() (AlignSetting, bool) {
	var _ok bool
	_ret := bindings.GetVTTCueAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return AlignSetting(_ret), _ok
}

// Align sets the value of property "VTTCue.align" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetAlign(val AlignSetting) bool {
	return js.True == bindings.SetVTTCueAlign(
		this.Ref(),
		uint32(val),
	)
}

// Text returns the value of property "VTTCue.text".
//
// The returned bool will be false if there is no such property.
func (this VTTCue) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetVTTCueText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Text sets the value of property "VTTCue.text" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetText(val js.String) bool {
	return js.True == bindings.SetVTTCueText(
		this.Ref(),
		val.Ref(),
	)
}

// GetCueAsHTML calls the method "VTTCue.getCueAsHTML".
//
// The returned bool will be false if there is no such method.
func (this VTTCue) GetCueAsHTML() (DocumentFragment, bool) {
	var _ok bool
	_ret := bindings.CallVTTCueGetCueAsHTML(
		this.Ref(), js.Pointer(&_ok),
	)

	return DocumentFragment{}.FromRef(_ret), _ok
}

// GetCueAsHTMLFunc returns the method "VTTCue.getCueAsHTML".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VTTCue) GetCueAsHTMLFunc() (fn js.Func[func() DocumentFragment]) {
	return fn.FromRef(
		bindings.VTTCueGetCueAsHTMLFunc(
			this.Ref(),
		),
	)
}

type ValueEventInit struct {
	// Value is "ValueEventInit.value"
	//
	// Optional, defaults to null.
	Value js.Any
	// Bubbles is "ValueEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ValueEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ValueEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ValueEventInit with all fields set.
func (p ValueEventInit) FromRef(ref js.Ref) ValueEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ValueEventInit in the application heap.
func (p ValueEventInit) New() js.Ref {
	return bindings.ValueEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ValueEventInit) UpdateFrom(ref js.Ref) {
	bindings.ValueEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ValueEventInit) Update(ref js.Ref) {
	bindings.ValueEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewValueEvent(typ js.String, initDict ValueEventInit) ValueEvent {
	return ValueEvent{}.FromRef(
		bindings.NewValueEventByValueEvent(
			typ.Ref(),
			js.Pointer(&initDict)),
	)
}

func NewValueEventByValueEvent1(typ js.String) ValueEvent {
	return ValueEvent{}.FromRef(
		bindings.NewValueEventByValueEvent1(
			typ.Ref()),
	)
}

type ValueEvent struct {
	Event
}

func (this ValueEvent) Once() ValueEvent {
	this.Ref().Once()
	return this
}

func (this ValueEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this ValueEvent) FromRef(ref js.Ref) ValueEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this ValueEvent) Free() {
	this.Ref().Free()
}

// Value returns the value of property "ValueEvent.value".
//
// The returned bool will be false if there is no such property.
func (this ValueEvent) Value() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetValueEventValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

type VideoFrameOutputCallbackFunc func(this js.Ref, output VideoFrame) js.Ref

func (fn VideoFrameOutputCallbackFunc) Register() js.Func[func(output VideoFrame)] {
	return js.RegisterCallback[func(output VideoFrame)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VideoFrameOutputCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		VideoFrame{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VideoFrameOutputCallback[T any] struct {
	Fn  func(arg T, this js.Ref, output VideoFrame) js.Ref
	Arg T
}

func (cb *VideoFrameOutputCallback[T]) Register() js.Func[func(output VideoFrame)] {
	return js.RegisterCallback[func(output VideoFrame)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VideoFrameOutputCallback[T]) DispatchCallback(
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

		VideoFrame{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type VideoDecoderInit struct {
	// Output is "VideoDecoderInit.output"
	//
	// Required
	Output js.Func[func(output VideoFrame)]
	// Error is "VideoDecoderInit.error"
	//
	// Required
	Error js.Func[func(err DOMException)]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoDecoderInit with all fields set.
func (p VideoDecoderInit) FromRef(ref js.Ref) VideoDecoderInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoDecoderInit in the application heap.
func (p VideoDecoderInit) New() js.Ref {
	return bindings.VideoDecoderInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoDecoderInit) UpdateFrom(ref js.Ref) {
	bindings.VideoDecoderInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoDecoderInit) Update(ref js.Ref) {
	bindings.VideoDecoderInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoDecoderSupport struct {
	// Supported is "VideoDecoderSupport.supported"
	//
	// Optional
	//
	// NOTE: FFI_USE_Supported MUST be set to true to make this field effective.
	Supported bool
	// Config is "VideoDecoderSupport.config"
	//
	// Optional
	Config VideoDecoderConfig

	FFI_USE_Supported bool // for Supported.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoDecoderSupport with all fields set.
func (p VideoDecoderSupport) FromRef(ref js.Ref) VideoDecoderSupport {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoDecoderSupport in the application heap.
func (p VideoDecoderSupport) New() js.Ref {
	return bindings.VideoDecoderSupportJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoDecoderSupport) UpdateFrom(ref js.Ref) {
	bindings.VideoDecoderSupportJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoDecoderSupport) Update(ref js.Ref) {
	bindings.VideoDecoderSupportJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewVideoDecoder(init VideoDecoderInit) VideoDecoder {
	return VideoDecoder{}.FromRef(
		bindings.NewVideoDecoderByVideoDecoder(
			js.Pointer(&init)),
	)
}

type VideoDecoder struct {
	EventTarget
}

func (this VideoDecoder) Once() VideoDecoder {
	this.Ref().Once()
	return this
}

func (this VideoDecoder) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this VideoDecoder) FromRef(ref js.Ref) VideoDecoder {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this VideoDecoder) Free() {
	this.Ref().Free()
}

// State returns the value of property "VideoDecoder.state".
//
// The returned bool will be false if there is no such property.
func (this VideoDecoder) State() (CodecState, bool) {
	var _ok bool
	_ret := bindings.GetVideoDecoderState(
		this.Ref(), js.Pointer(&_ok),
	)
	return CodecState(_ret), _ok
}

// DecodeQueueSize returns the value of property "VideoDecoder.decodeQueueSize".
//
// The returned bool will be false if there is no such property.
func (this VideoDecoder) DecodeQueueSize() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetVideoDecoderDecodeQueueSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Configure calls the method "VideoDecoder.configure".
//
// The returned bool will be false if there is no such method.
func (this VideoDecoder) Configure(config VideoDecoderConfig) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallVideoDecoderConfigure(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&config),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ConfigureFunc returns the method "VideoDecoder.configure".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoDecoder) ConfigureFunc() (fn js.Func[func(config VideoDecoderConfig)]) {
	return fn.FromRef(
		bindings.VideoDecoderConfigureFunc(
			this.Ref(),
		),
	)
}

// Decode calls the method "VideoDecoder.decode".
//
// The returned bool will be false if there is no such method.
func (this VideoDecoder) Decode(chunk EncodedVideoChunk) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallVideoDecoderDecode(
		this.Ref(), js.Pointer(&_ok),
		chunk.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DecodeFunc returns the method "VideoDecoder.decode".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoDecoder) DecodeFunc() (fn js.Func[func(chunk EncodedVideoChunk)]) {
	return fn.FromRef(
		bindings.VideoDecoderDecodeFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "VideoDecoder.flush".
//
// The returned bool will be false if there is no such method.
func (this VideoDecoder) Flush() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallVideoDecoderFlush(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// FlushFunc returns the method "VideoDecoder.flush".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoDecoder) FlushFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.VideoDecoderFlushFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "VideoDecoder.reset".
//
// The returned bool will be false if there is no such method.
func (this VideoDecoder) Reset() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallVideoDecoderReset(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetFunc returns the method "VideoDecoder.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoDecoder) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VideoDecoderResetFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "VideoDecoder.close".
//
// The returned bool will be false if there is no such method.
func (this VideoDecoder) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallVideoDecoderClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "VideoDecoder.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoDecoder) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.VideoDecoderCloseFunc(
			this.Ref(),
		),
	)
}

// IsConfigSupported calls the staticmethod "VideoDecoder.isConfigSupported".
//
// The returned bool will be false if there is no such method.
func (this VideoDecoder) IsConfigSupported(config VideoDecoderConfig) (js.Promise[VideoDecoderSupport], bool) {
	var _ok bool
	_ret := bindings.CallVideoDecoderIsConfigSupported(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&config),
	)

	return js.Promise[VideoDecoderSupport]{}.FromRef(_ret), _ok
}

// IsConfigSupportedFunc returns the staticmethod "VideoDecoder.isConfigSupported".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this VideoDecoder) IsConfigSupportedFunc() (fn js.Func[func(config VideoDecoderConfig) js.Promise[VideoDecoderSupport]]) {
	return fn.FromRef(
		bindings.VideoDecoderIsConfigSupportedFunc(
			this.Ref(),
		),
	)
}

type VideoEncoderInit struct {
	// Output is "VideoEncoderInit.output"
	//
	// Required
	Output js.Func[func(chunk EncodedVideoChunk, metadata EncodedVideoChunkMetadata)]
	// Error is "VideoEncoderInit.error"
	//
	// Required
	Error js.Func[func(err DOMException)]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoEncoderInit with all fields set.
func (p VideoEncoderInit) FromRef(ref js.Ref) VideoEncoderInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoEncoderInit in the application heap.
func (p VideoEncoderInit) New() js.Ref {
	return bindings.VideoEncoderInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoEncoderInit) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoEncoderInit) Update(ref js.Ref) {
	bindings.VideoEncoderInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoEncoderBitrateMode uint32

const (
	_ VideoEncoderBitrateMode = iota

	VideoEncoderBitrateMode_CONSTANT
	VideoEncoderBitrateMode_VARIABLE
	VideoEncoderBitrateMode_QUANTIZER
)

func (VideoEncoderBitrateMode) FromRef(str js.Ref) VideoEncoderBitrateMode {
	return VideoEncoderBitrateMode(bindings.ConstOfVideoEncoderBitrateMode(str))
}

func (x VideoEncoderBitrateMode) String() (string, bool) {
	switch x {
	case VideoEncoderBitrateMode_CONSTANT:
		return "constant", true
	case VideoEncoderBitrateMode_VARIABLE:
		return "variable", true
	case VideoEncoderBitrateMode_QUANTIZER:
		return "quantizer", true
	default:
		return "", false
	}
}

type VideoEncoderConfig struct {
	// Codec is "VideoEncoderConfig.codec"
	//
	// Required
	Codec js.String
	// Width is "VideoEncoderConfig.width"
	//
	// Required
	Width uint32
	// Height is "VideoEncoderConfig.height"
	//
	// Required
	Height uint32
	// DisplayWidth is "VideoEncoderConfig.displayWidth"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayWidth MUST be set to true to make this field effective.
	DisplayWidth uint32
	// DisplayHeight is "VideoEncoderConfig.displayHeight"
	//
	// Optional
	//
	// NOTE: FFI_USE_DisplayHeight MUST be set to true to make this field effective.
	DisplayHeight uint32
	// Bitrate is "VideoEncoderConfig.bitrate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Bitrate MUST be set to true to make this field effective.
	Bitrate uint64
	// Framerate is "VideoEncoderConfig.framerate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Framerate MUST be set to true to make this field effective.
	Framerate float64
	// HardwareAcceleration is "VideoEncoderConfig.hardwareAcceleration"
	//
	// Optional, defaults to "no-preference".
	HardwareAcceleration HardwareAcceleration
	// Alpha is "VideoEncoderConfig.alpha"
	//
	// Optional, defaults to "discard".
	Alpha AlphaOption
	// ScalabilityMode is "VideoEncoderConfig.scalabilityMode"
	//
	// Optional
	ScalabilityMode js.String
	// BitrateMode is "VideoEncoderConfig.bitrateMode"
	//
	// Optional, defaults to "variable".
	BitrateMode VideoEncoderBitrateMode
	// LatencyMode is "VideoEncoderConfig.latencyMode"
	//
	// Optional, defaults to "quality".
	LatencyMode LatencyMode
	// Hevc is "VideoEncoderConfig.hevc"
	//
	// Optional
	Hevc HevcEncoderConfig
	// Avc is "VideoEncoderConfig.avc"
	//
	// Optional
	Avc AvcEncoderConfig
	// Av1 is "VideoEncoderConfig.av1"
	//
	// Optional
	Av1 AV1EncoderConfig

	FFI_USE_DisplayWidth  bool // for DisplayWidth.
	FFI_USE_DisplayHeight bool // for DisplayHeight.
	FFI_USE_Bitrate       bool // for Bitrate.
	FFI_USE_Framerate     bool // for Framerate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoEncoderConfig with all fields set.
func (p VideoEncoderConfig) FromRef(ref js.Ref) VideoEncoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoEncoderConfig in the application heap.
func (p VideoEncoderConfig) New() js.Ref {
	return bindings.VideoEncoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoEncoderConfig) Update(ref js.Ref) {
	bindings.VideoEncoderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoEncoderEncodeOptionsForHevc struct {
	// Quantizer is "VideoEncoderEncodeOptionsForHevc.quantizer"
	//
	// Optional
	//
	// NOTE: FFI_USE_Quantizer MUST be set to true to make this field effective.
	Quantizer uint16

	FFI_USE_Quantizer bool // for Quantizer.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoEncoderEncodeOptionsForHevc with all fields set.
func (p VideoEncoderEncodeOptionsForHevc) FromRef(ref js.Ref) VideoEncoderEncodeOptionsForHevc {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoEncoderEncodeOptionsForHevc in the application heap.
func (p VideoEncoderEncodeOptionsForHevc) New() js.Ref {
	return bindings.VideoEncoderEncodeOptionsForHevcJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoEncoderEncodeOptionsForHevc) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForHevcJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoEncoderEncodeOptionsForHevc) Update(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForHevcJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type VideoEncoderEncodeOptionsForVp9 struct {
	// Quantizer is "VideoEncoderEncodeOptionsForVp9.quantizer"
	//
	// Optional
	//
	// NOTE: FFI_USE_Quantizer MUST be set to true to make this field effective.
	Quantizer uint16

	FFI_USE_Quantizer bool // for Quantizer.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VideoEncoderEncodeOptionsForVp9 with all fields set.
func (p VideoEncoderEncodeOptionsForVp9) FromRef(ref js.Ref) VideoEncoderEncodeOptionsForVp9 {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VideoEncoderEncodeOptionsForVp9 in the application heap.
func (p VideoEncoderEncodeOptionsForVp9) New() js.Ref {
	return bindings.VideoEncoderEncodeOptionsForVp9JSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p VideoEncoderEncodeOptionsForVp9) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForVp9JSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p VideoEncoderEncodeOptionsForVp9) Update(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForVp9JSLoad(
		js.Pointer(&p), js.False, ref,
	)
}
