// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
func (p *Transformer) UpdateFrom(ref js.Ref) {
	bindings.TransformerJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Transformer) Update(ref js.Ref) {
	bindings.TransformerJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Transformer) FreeMembers(recursive bool) {
	js.Free(
		p.Start.Ref(),
		p.Transform.Ref(),
		p.Flush.Ref(),
		p.ReadableType.Ref(),
		p.WritableType.Ref(),
	)
	p.Start = p.Start.FromRef(js.Undefined)
	p.Transform = p.Transform.FromRef(js.Undefined)
	p.Flush = p.Flush.FromRef(js.Undefined)
	p.ReadableType = p.ReadableType.FromRef(js.Undefined)
	p.WritableType = p.WritableType.FromRef(js.Undefined)
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
func (p *TransitionEventInit) UpdateFrom(ref js.Ref) {
	bindings.TransitionEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TransitionEventInit) Update(ref js.Ref) {
	bindings.TransitionEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TransitionEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.PropertyName.Ref(),
		p.PseudoElement.Ref(),
	)
	p.PropertyName = p.PropertyName.FromRef(js.Undefined)
	p.PseudoElement = p.PseudoElement.FromRef(js.Undefined)
}

func NewTransitionEvent(typ js.String, transitionEventInitDict TransitionEventInit) (ret TransitionEvent) {
	ret.ref = bindings.NewTransitionEventByTransitionEvent(
		typ.Ref(),
		js.Pointer(&transitionEventInitDict))
	return
}

func NewTransitionEventByTransitionEvent1(typ js.String) (ret TransitionEvent) {
	ret.ref = bindings.NewTransitionEventByTransitionEvent1(
		typ.Ref())
	return
}

type TransitionEvent struct {
	Event
}

func (this TransitionEvent) Once() TransitionEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// PropertyName returns the value of property "TransitionEvent.propertyName".
//
// It returns ok=false if there is no such property.
func (this TransitionEvent) PropertyName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTransitionEventPropertyName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ElapsedTime returns the value of property "TransitionEvent.elapsedTime".
//
// It returns ok=false if there is no such property.
func (this TransitionEvent) ElapsedTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetTransitionEventElapsedTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PseudoElement returns the value of property "TransitionEvent.pseudoElement".
//
// It returns ok=false if there is no such property.
func (this TransitionEvent) PseudoElement() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTransitionEventPseudoElement(
		this.ref, js.Pointer(&ret),
	)
	return
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
func (p *UIEventInit) UpdateFrom(ref js.Ref) {
	bindings.UIEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UIEventInit) Update(ref js.Ref) {
	bindings.UIEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UIEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.View.Ref(),
		p.SourceCapabilities.Ref(),
	)
	p.View = p.View.FromRef(js.Undefined)
	p.SourceCapabilities = p.SourceCapabilities.FromRef(js.Undefined)
}

func NewUIEvent(typ js.String, eventInitDict UIEventInit) (ret UIEvent) {
	ret.ref = bindings.NewUIEventByUIEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewUIEventByUIEvent1(typ js.String) (ret UIEvent) {
	ret.ref = bindings.NewUIEventByUIEvent1(
		typ.Ref())
	return
}

type UIEvent struct {
	Event
}

func (this UIEvent) Once() UIEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// View returns the value of property "UIEvent.view".
//
// It returns ok=false if there is no such property.
func (this UIEvent) View() (ret Window, ok bool) {
	ok = js.True == bindings.GetUIEventView(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Detail returns the value of property "UIEvent.detail".
//
// It returns ok=false if there is no such property.
func (this UIEvent) Detail() (ret int32, ok bool) {
	ok = js.True == bindings.GetUIEventDetail(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Which returns the value of property "UIEvent.which".
//
// It returns ok=false if there is no such property.
func (this UIEvent) Which() (ret uint32, ok bool) {
	ok = js.True == bindings.GetUIEventWhich(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SourceCapabilities returns the value of property "UIEvent.sourceCapabilities".
//
// It returns ok=false if there is no such property.
func (this UIEvent) SourceCapabilities() (ret InputDeviceCapabilities, ok bool) {
	ok = js.True == bindings.GetUIEventSourceCapabilities(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncInitUIEvent returns true if the method "UIEvent.initUIEvent" exists.
func (this UIEvent) HasFuncInitUIEvent() bool {
	return js.True == bindings.HasFuncUIEventInitUIEvent(
		this.ref,
	)
}

// FuncInitUIEvent returns the method "UIEvent.initUIEvent".
func (this UIEvent) FuncInitUIEvent() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32)]) {
	bindings.FuncUIEventInitUIEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitUIEvent calls the method "UIEvent.initUIEvent".
func (this UIEvent) InitUIEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32) (ret js.Void) {
	bindings.CallUIEventInitUIEvent(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
	)

	return
}

// TryInitUIEvent calls the method "UIEvent.initUIEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this UIEvent) TryInitUIEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUIEventInitUIEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
	)

	return
}

// HasFuncInitUIEvent1 returns true if the method "UIEvent.initUIEvent" exists.
func (this UIEvent) HasFuncInitUIEvent1() bool {
	return js.True == bindings.HasFuncUIEventInitUIEvent1(
		this.ref,
	)
}

// FuncInitUIEvent1 returns the method "UIEvent.initUIEvent".
func (this UIEvent) FuncInitUIEvent1() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window)]) {
	bindings.FuncUIEventInitUIEvent1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitUIEvent1 calls the method "UIEvent.initUIEvent".
func (this UIEvent) InitUIEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (ret js.Void) {
	bindings.CallUIEventInitUIEvent1(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// TryInitUIEvent1 calls the method "UIEvent.initUIEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this UIEvent) TryInitUIEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUIEventInitUIEvent1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// HasFuncInitUIEvent2 returns true if the method "UIEvent.initUIEvent" exists.
func (this UIEvent) HasFuncInitUIEvent2() bool {
	return js.True == bindings.HasFuncUIEventInitUIEvent2(
		this.ref,
	)
}

// FuncInitUIEvent2 returns the method "UIEvent.initUIEvent".
func (this UIEvent) FuncInitUIEvent2() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	bindings.FuncUIEventInitUIEvent2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitUIEvent2 calls the method "UIEvent.initUIEvent".
func (this UIEvent) InitUIEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void) {
	bindings.CallUIEventInitUIEvent2(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// TryInitUIEvent2 calls the method "UIEvent.initUIEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this UIEvent) TryInitUIEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUIEventInitUIEvent2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// HasFuncInitUIEvent3 returns true if the method "UIEvent.initUIEvent" exists.
func (this UIEvent) HasFuncInitUIEvent3() bool {
	return js.True == bindings.HasFuncUIEventInitUIEvent3(
		this.ref,
	)
}

// FuncInitUIEvent3 returns the method "UIEvent.initUIEvent".
func (this UIEvent) FuncInitUIEvent3() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	bindings.FuncUIEventInitUIEvent3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitUIEvent3 calls the method "UIEvent.initUIEvent".
func (this UIEvent) InitUIEvent3(typeArg js.String, bubblesArg bool) (ret js.Void) {
	bindings.CallUIEventInitUIEvent3(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// TryInitUIEvent3 calls the method "UIEvent.initUIEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this UIEvent) TryInitUIEvent3(typeArg js.String, bubblesArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUIEventInitUIEvent3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// HasFuncInitUIEvent4 returns true if the method "UIEvent.initUIEvent" exists.
func (this UIEvent) HasFuncInitUIEvent4() bool {
	return js.True == bindings.HasFuncUIEventInitUIEvent4(
		this.ref,
	)
}

// FuncInitUIEvent4 returns the method "UIEvent.initUIEvent".
func (this UIEvent) FuncInitUIEvent4() (fn js.Func[func(typeArg js.String)]) {
	bindings.FuncUIEventInitUIEvent4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitUIEvent4 calls the method "UIEvent.initUIEvent".
func (this UIEvent) InitUIEvent4(typeArg js.String) (ret js.Void) {
	bindings.CallUIEventInitUIEvent4(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
	)

	return
}

// TryInitUIEvent4 calls the method "UIEvent.initUIEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this UIEvent) TryInitUIEvent4(typeArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUIEventInitUIEvent4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
	)

	return
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

func NewURL(url js.String, base js.String) (ret URL) {
	ret.ref = bindings.NewURLByURL(
		url.Ref(),
		base.Ref())
	return
}

func NewURLByURL1(url js.String) (ret URL) {
	ret.ref = bindings.NewURLByURL1(
		url.Ref())
	return
}

type URL struct {
	ref js.Ref
}

func (this URL) Once() URL {
	this.ref.Once()
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
	this.ref.Free()
}

// Href returns the value of property "URL.href".
//
// It returns ok=false if there is no such property.
func (this URL) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHref sets the value of property "URL.href" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetHref(val js.String) bool {
	return js.True == bindings.SetURLHref(
		this.ref,
		val.Ref(),
	)
}

// Origin returns the value of property "URL.origin".
//
// It returns ok=false if there is no such property.
func (this URL) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "URL.protocol".
//
// It returns ok=false if there is no such property.
func (this URL) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetProtocol sets the value of property "URL.protocol" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetProtocol(val js.String) bool {
	return js.True == bindings.SetURLProtocol(
		this.ref,
		val.Ref(),
	)
}

// Username returns the value of property "URL.username".
//
// It returns ok=false if there is no such property.
func (this URL) Username() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLUsername(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUsername sets the value of property "URL.username" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetUsername(val js.String) bool {
	return js.True == bindings.SetURLUsername(
		this.ref,
		val.Ref(),
	)
}

// Password returns the value of property "URL.password".
//
// It returns ok=false if there is no such property.
func (this URL) Password() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPassword(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPassword sets the value of property "URL.password" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetPassword(val js.String) bool {
	return js.True == bindings.SetURLPassword(
		this.ref,
		val.Ref(),
	)
}

// Host returns the value of property "URL.host".
//
// It returns ok=false if there is no such property.
func (this URL) Host() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLHost(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHost sets the value of property "URL.host" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetHost(val js.String) bool {
	return js.True == bindings.SetURLHost(
		this.ref,
		val.Ref(),
	)
}

// Hostname returns the value of property "URL.hostname".
//
// It returns ok=false if there is no such property.
func (this URL) Hostname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLHostname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHostname sets the value of property "URL.hostname" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetHostname(val js.String) bool {
	return js.True == bindings.SetURLHostname(
		this.ref,
		val.Ref(),
	)
}

// Port returns the value of property "URL.port".
//
// It returns ok=false if there is no such property.
func (this URL) Port() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPort sets the value of property "URL.port" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetPort(val js.String) bool {
	return js.True == bindings.SetURLPort(
		this.ref,
		val.Ref(),
	)
}

// Pathname returns the value of property "URL.pathname".
//
// It returns ok=false if there is no such property.
func (this URL) Pathname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPathname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPathname sets the value of property "URL.pathname" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetPathname(val js.String) bool {
	return js.True == bindings.SetURLPathname(
		this.ref,
		val.Ref(),
	)
}

// Search returns the value of property "URL.search".
//
// It returns ok=false if there is no such property.
func (this URL) Search() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLSearch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSearch sets the value of property "URL.search" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetSearch(val js.String) bool {
	return js.True == bindings.SetURLSearch(
		this.ref,
		val.Ref(),
	)
}

// SearchParams returns the value of property "URL.searchParams".
//
// It returns ok=false if there is no such property.
func (this URL) SearchParams() (ret URLSearchParams, ok bool) {
	ok = js.True == bindings.GetURLSearchParams(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Hash returns the value of property "URL.hash".
//
// It returns ok=false if there is no such property.
func (this URL) Hash() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLHash(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHash sets the value of property "URL.hash" to val.
//
// It returns false if the property cannot be set.
func (this URL) SetHash(val js.String) bool {
	return js.True == bindings.SetURLHash(
		this.ref,
		val.Ref(),
	)
}

// HasFuncCanParse returns true if the static method "URL.canParse" exists.
func (this URL) HasFuncCanParse() bool {
	return js.True == bindings.HasFuncURLCanParse(
		this.ref,
	)
}

// FuncCanParse returns the static method "URL.canParse".
func (this URL) FuncCanParse() (fn js.Func[func(url js.String, base js.String) bool]) {
	bindings.FuncURLCanParse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CanParse calls the static method "URL.canParse".
func (this URL) CanParse(url js.String, base js.String) (ret bool) {
	bindings.CallURLCanParse(
		this.ref, js.Pointer(&ret),
		url.Ref(),
		base.Ref(),
	)

	return
}

// TryCanParse calls the static method "URL.canParse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URL) TryCanParse(url js.String, base js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLCanParse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		base.Ref(),
	)

	return
}

// HasFuncCanParse1 returns true if the static method "URL.canParse" exists.
func (this URL) HasFuncCanParse1() bool {
	return js.True == bindings.HasFuncURLCanParse1(
		this.ref,
	)
}

// FuncCanParse1 returns the static method "URL.canParse".
func (this URL) FuncCanParse1() (fn js.Func[func(url js.String) bool]) {
	bindings.FuncURLCanParse1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CanParse1 calls the static method "URL.canParse".
func (this URL) CanParse1(url js.String) (ret bool) {
	bindings.CallURLCanParse1(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryCanParse1 calls the static method "URL.canParse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URL) TryCanParse1(url js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLCanParse1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncToJSON returns true if the method "URL.toJSON" exists.
func (this URL) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncURLToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "URL.toJSON".
func (this URL) FuncToJSON() (fn js.Func[func() js.String]) {
	bindings.FuncURLToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "URL.toJSON".
func (this URL) ToJSON() (ret js.String) {
	bindings.CallURLToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "URL.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URL) TryToJSON() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateObjectURL returns true if the static method "URL.createObjectURL" exists.
func (this URL) HasFuncCreateObjectURL() bool {
	return js.True == bindings.HasFuncURLCreateObjectURL(
		this.ref,
	)
}

// FuncCreateObjectURL returns the static method "URL.createObjectURL".
func (this URL) FuncCreateObjectURL() (fn js.Func[func(obj OneOf_Blob_MediaSource) js.String]) {
	bindings.FuncURLCreateObjectURL(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateObjectURL calls the static method "URL.createObjectURL".
func (this URL) CreateObjectURL(obj OneOf_Blob_MediaSource) (ret js.String) {
	bindings.CallURLCreateObjectURL(
		this.ref, js.Pointer(&ret),
		obj.Ref(),
	)

	return
}

// TryCreateObjectURL calls the static method "URL.createObjectURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URL) TryCreateObjectURL(obj OneOf_Blob_MediaSource) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLCreateObjectURL(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		obj.Ref(),
	)

	return
}

// HasFuncRevokeObjectURL returns true if the static method "URL.revokeObjectURL" exists.
func (this URL) HasFuncRevokeObjectURL() bool {
	return js.True == bindings.HasFuncURLRevokeObjectURL(
		this.ref,
	)
}

// FuncRevokeObjectURL returns the static method "URL.revokeObjectURL".
func (this URL) FuncRevokeObjectURL() (fn js.Func[func(url js.String)]) {
	bindings.FuncURLRevokeObjectURL(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RevokeObjectURL calls the static method "URL.revokeObjectURL".
func (this URL) RevokeObjectURL(url js.String) (ret js.Void) {
	bindings.CallURLRevokeObjectURL(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryRevokeObjectURL calls the static method "URL.revokeObjectURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URL) TryRevokeObjectURL(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLRevokeObjectURL(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
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
func (p *URLPatternInit) UpdateFrom(ref js.Ref) {
	bindings.URLPatternInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *URLPatternInit) Update(ref js.Ref) {
	bindings.URLPatternInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *URLPatternInit) FreeMembers(recursive bool) {
	js.Free(
		p.Protocol.Ref(),
		p.Username.Ref(),
		p.Password.Ref(),
		p.Hostname.Ref(),
		p.Port.Ref(),
		p.Pathname.Ref(),
		p.Search.Ref(),
		p.Hash.Ref(),
		p.BaseURL.Ref(),
	)
	p.Protocol = p.Protocol.FromRef(js.Undefined)
	p.Username = p.Username.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
	p.Hostname = p.Hostname.FromRef(js.Undefined)
	p.Port = p.Port.FromRef(js.Undefined)
	p.Pathname = p.Pathname.FromRef(js.Undefined)
	p.Search = p.Search.FromRef(js.Undefined)
	p.Hash = p.Hash.FromRef(js.Undefined)
	p.BaseURL = p.BaseURL.FromRef(js.Undefined)
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
func (p *URLPatternOptions) UpdateFrom(ref js.Ref) {
	bindings.URLPatternOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *URLPatternOptions) Update(ref js.Ref) {
	bindings.URLPatternOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *URLPatternOptions) FreeMembers(recursive bool) {
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
func (p *URLPatternComponentResult) UpdateFrom(ref js.Ref) {
	bindings.URLPatternComponentResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *URLPatternComponentResult) Update(ref js.Ref) {
	bindings.URLPatternComponentResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *URLPatternComponentResult) FreeMembers(recursive bool) {
	js.Free(
		p.Input.Ref(),
		p.Groups.Ref(),
	)
	p.Input = p.Input.FromRef(js.Undefined)
	p.Groups = p.Groups.FromRef(js.Undefined)
}

type URLPatternResult struct {
	// Inputs is "URLPatternResult.inputs"
	//
	// Optional
	Inputs js.Array[URLPatternInput]
	// Protocol is "URLPatternResult.protocol"
	//
	// Optional
	//
	// NOTE: Protocol.FFI_USE MUST be set to true to get Protocol used.
	Protocol URLPatternComponentResult
	// Username is "URLPatternResult.username"
	//
	// Optional
	//
	// NOTE: Username.FFI_USE MUST be set to true to get Username used.
	Username URLPatternComponentResult
	// Password is "URLPatternResult.password"
	//
	// Optional
	//
	// NOTE: Password.FFI_USE MUST be set to true to get Password used.
	Password URLPatternComponentResult
	// Hostname is "URLPatternResult.hostname"
	//
	// Optional
	//
	// NOTE: Hostname.FFI_USE MUST be set to true to get Hostname used.
	Hostname URLPatternComponentResult
	// Port is "URLPatternResult.port"
	//
	// Optional
	//
	// NOTE: Port.FFI_USE MUST be set to true to get Port used.
	Port URLPatternComponentResult
	// Pathname is "URLPatternResult.pathname"
	//
	// Optional
	//
	// NOTE: Pathname.FFI_USE MUST be set to true to get Pathname used.
	Pathname URLPatternComponentResult
	// Search is "URLPatternResult.search"
	//
	// Optional
	//
	// NOTE: Search.FFI_USE MUST be set to true to get Search used.
	Search URLPatternComponentResult
	// Hash is "URLPatternResult.hash"
	//
	// Optional
	//
	// NOTE: Hash.FFI_USE MUST be set to true to get Hash used.
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
func (p *URLPatternResult) UpdateFrom(ref js.Ref) {
	bindings.URLPatternResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *URLPatternResult) Update(ref js.Ref) {
	bindings.URLPatternResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *URLPatternResult) FreeMembers(recursive bool) {
	js.Free(
		p.Inputs.Ref(),
	)
	p.Inputs = p.Inputs.FromRef(js.Undefined)
	if recursive {
		p.Protocol.FreeMembers(true)
		p.Username.FreeMembers(true)
		p.Password.FreeMembers(true)
		p.Hostname.FreeMembers(true)
		p.Port.FreeMembers(true)
		p.Pathname.FreeMembers(true)
		p.Search.FreeMembers(true)
		p.Hash.FreeMembers(true)
	}
}

func NewURLPattern(input URLPatternInput, baseURL js.String, options URLPatternOptions) (ret URLPattern) {
	ret.ref = bindings.NewURLPatternByURLPattern(
		input.Ref(),
		baseURL.Ref(),
		js.Pointer(&options))
	return
}

func NewURLPatternByURLPattern1(input URLPatternInput, baseURL js.String) (ret URLPattern) {
	ret.ref = bindings.NewURLPatternByURLPattern1(
		input.Ref(),
		baseURL.Ref())
	return
}

func NewURLPatternByURLPattern2(input URLPatternInput, options URLPatternOptions) (ret URLPattern) {
	ret.ref = bindings.NewURLPatternByURLPattern2(
		input.Ref(),
		js.Pointer(&options))
	return
}

func NewURLPatternByURLPattern3(input URLPatternInput) (ret URLPattern) {
	ret.ref = bindings.NewURLPatternByURLPattern3(
		input.Ref())
	return
}

func NewURLPatternByURLPattern4() (ret URLPattern) {
	ret.ref = bindings.NewURLPatternByURLPattern4()
	return
}

type URLPattern struct {
	ref js.Ref
}

func (this URLPattern) Once() URLPattern {
	this.ref.Once()
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
	this.ref.Free()
}

// Protocol returns the value of property "URLPattern.protocol".
//
// It returns ok=false if there is no such property.
func (this URLPattern) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPatternProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Username returns the value of property "URLPattern.username".
//
// It returns ok=false if there is no such property.
func (this URLPattern) Username() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPatternUsername(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Password returns the value of property "URLPattern.password".
//
// It returns ok=false if there is no such property.
func (this URLPattern) Password() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPatternPassword(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Hostname returns the value of property "URLPattern.hostname".
//
// It returns ok=false if there is no such property.
func (this URLPattern) Hostname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPatternHostname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Port returns the value of property "URLPattern.port".
//
// It returns ok=false if there is no such property.
func (this URLPattern) Port() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPatternPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Pathname returns the value of property "URLPattern.pathname".
//
// It returns ok=false if there is no such property.
func (this URLPattern) Pathname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPatternPathname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Search returns the value of property "URLPattern.search".
//
// It returns ok=false if there is no such property.
func (this URLPattern) Search() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPatternSearch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Hash returns the value of property "URLPattern.hash".
//
// It returns ok=false if there is no such property.
func (this URLPattern) Hash() (ret js.String, ok bool) {
	ok = js.True == bindings.GetURLPatternHash(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncTest returns true if the method "URLPattern.test" exists.
func (this URLPattern) HasFuncTest() bool {
	return js.True == bindings.HasFuncURLPatternTest(
		this.ref,
	)
}

// FuncTest returns the method "URLPattern.test".
func (this URLPattern) FuncTest() (fn js.Func[func(input URLPatternInput, baseURL js.String) bool]) {
	bindings.FuncURLPatternTest(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Test calls the method "URLPattern.test".
func (this URLPattern) Test(input URLPatternInput, baseURL js.String) (ret bool) {
	bindings.CallURLPatternTest(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		baseURL.Ref(),
	)

	return
}

// TryTest calls the method "URLPattern.test"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URLPattern) TryTest(input URLPatternInput, baseURL js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLPatternTest(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		baseURL.Ref(),
	)

	return
}

// HasFuncTest1 returns true if the method "URLPattern.test" exists.
func (this URLPattern) HasFuncTest1() bool {
	return js.True == bindings.HasFuncURLPatternTest1(
		this.ref,
	)
}

// FuncTest1 returns the method "URLPattern.test".
func (this URLPattern) FuncTest1() (fn js.Func[func(input URLPatternInput) bool]) {
	bindings.FuncURLPatternTest1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Test1 calls the method "URLPattern.test".
func (this URLPattern) Test1(input URLPatternInput) (ret bool) {
	bindings.CallURLPatternTest1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryTest1 calls the method "URLPattern.test"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URLPattern) TryTest1(input URLPatternInput) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLPatternTest1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncTest2 returns true if the method "URLPattern.test" exists.
func (this URLPattern) HasFuncTest2() bool {
	return js.True == bindings.HasFuncURLPatternTest2(
		this.ref,
	)
}

// FuncTest2 returns the method "URLPattern.test".
func (this URLPattern) FuncTest2() (fn js.Func[func() bool]) {
	bindings.FuncURLPatternTest2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Test2 calls the method "URLPattern.test".
func (this URLPattern) Test2() (ret bool) {
	bindings.CallURLPatternTest2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTest2 calls the method "URLPattern.test"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URLPattern) TryTest2() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLPatternTest2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExec returns true if the method "URLPattern.exec" exists.
func (this URLPattern) HasFuncExec() bool {
	return js.True == bindings.HasFuncURLPatternExec(
		this.ref,
	)
}

// FuncExec returns the method "URLPattern.exec".
func (this URLPattern) FuncExec() (fn js.Func[func(input URLPatternInput, baseURL js.String) URLPatternResult]) {
	bindings.FuncURLPatternExec(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Exec calls the method "URLPattern.exec".
func (this URLPattern) Exec(input URLPatternInput, baseURL js.String) (ret URLPatternResult) {
	bindings.CallURLPatternExec(
		this.ref, js.Pointer(&ret),
		input.Ref(),
		baseURL.Ref(),
	)

	return
}

// TryExec calls the method "URLPattern.exec"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URLPattern) TryExec(input URLPatternInput, baseURL js.String) (ret URLPatternResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLPatternExec(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		baseURL.Ref(),
	)

	return
}

// HasFuncExec1 returns true if the method "URLPattern.exec" exists.
func (this URLPattern) HasFuncExec1() bool {
	return js.True == bindings.HasFuncURLPatternExec1(
		this.ref,
	)
}

// FuncExec1 returns the method "URLPattern.exec".
func (this URLPattern) FuncExec1() (fn js.Func[func(input URLPatternInput) URLPatternResult]) {
	bindings.FuncURLPatternExec1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Exec1 calls the method "URLPattern.exec".
func (this URLPattern) Exec1(input URLPatternInput) (ret URLPatternResult) {
	bindings.CallURLPatternExec1(
		this.ref, js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryExec1 calls the method "URLPattern.exec"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URLPattern) TryExec1(input URLPatternInput) (ret URLPatternResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLPatternExec1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFuncExec2 returns true if the method "URLPattern.exec" exists.
func (this URLPattern) HasFuncExec2() bool {
	return js.True == bindings.HasFuncURLPatternExec2(
		this.ref,
	)
}

// FuncExec2 returns the method "URLPattern.exec".
func (this URLPattern) FuncExec2() (fn js.Func[func() URLPatternResult]) {
	bindings.FuncURLPatternExec2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Exec2 calls the method "URLPattern.exec".
func (this URLPattern) Exec2() (ret URLPatternResult) {
	bindings.CallURLPatternExec2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryExec2 calls the method "URLPattern.exec"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this URLPattern) TryExec2() (ret URLPatternResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryURLPatternExec2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
func (p *USBConnectionEventInit) UpdateFrom(ref js.Ref) {
	bindings.USBConnectionEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *USBConnectionEventInit) Update(ref js.Ref) {
	bindings.USBConnectionEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *USBConnectionEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Device.Ref(),
	)
	p.Device = p.Device.FromRef(js.Undefined)
}

func NewUSBConnectionEvent(typ js.String, eventInitDict USBConnectionEventInit) (ret USBConnectionEvent) {
	ret.ref = bindings.NewUSBConnectionEventByUSBConnectionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type USBConnectionEvent struct {
	Event
}

func (this USBConnectionEvent) Once() USBConnectionEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Device returns the value of property "USBConnectionEvent.device".
//
// It returns ok=false if there is no such property.
func (this USBConnectionEvent) Device() (ret USBDevice, ok bool) {
	ok = js.True == bindings.GetUSBConnectionEventDevice(
		this.ref, js.Pointer(&ret),
	)
	return
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
func (p *USBPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.USBPermissionDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *USBPermissionDescriptor) Update(ref js.Ref) {
	bindings.USBPermissionDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *USBPermissionDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Filters.Ref(),
		p.ExclusionFilters.Ref(),
		p.Name.Ref(),
	)
	p.Filters = p.Filters.FromRef(js.Undefined)
	p.ExclusionFilters = p.ExclusionFilters.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

type USBPermissionResult struct {
	PermissionStatus
}

func (this USBPermissionResult) Once() USBPermissionResult {
	this.ref.Once()
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
	this.ref.Free()
}

// Devices returns the value of property "USBPermissionResult.devices".
//
// It returns ok=false if there is no such property.
func (this USBPermissionResult) Devices() (ret js.FrozenArray[USBDevice], ok bool) {
	ok = js.True == bindings.GetUSBPermissionResultDevices(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDevices sets the value of property "USBPermissionResult.devices" to val.
//
// It returns false if the property cannot be set.
func (this USBPermissionResult) SetDevices(val js.FrozenArray[USBDevice]) bool {
	return js.True == bindings.SetUSBPermissionResultDevices(
		this.ref,
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
func (p *USBPermissionStorage) UpdateFrom(ref js.Ref) {
	bindings.USBPermissionStorageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *USBPermissionStorage) Update(ref js.Ref) {
	bindings.USBPermissionStorageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *USBPermissionStorage) FreeMembers(recursive bool) {
	js.Free(
		p.AllowedDevices.Ref(),
	)
	p.AllowedDevices = p.AllowedDevices.FromRef(js.Undefined)
}

func NewUncalibratedMagnetometer(sensorOptions MagnetometerSensorOptions) (ret UncalibratedMagnetometer) {
	ret.ref = bindings.NewUncalibratedMagnetometerByUncalibratedMagnetometer(
		js.Pointer(&sensorOptions))
	return
}

func NewUncalibratedMagnetometerByUncalibratedMagnetometer1() (ret UncalibratedMagnetometer) {
	ret.ref = bindings.NewUncalibratedMagnetometerByUncalibratedMagnetometer1()
	return
}

type UncalibratedMagnetometer struct {
	Sensor
}

func (this UncalibratedMagnetometer) Once() UncalibratedMagnetometer {
	this.ref.Once()
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
	this.ref.Free()
}

// X returns the value of property "UncalibratedMagnetometer.x".
//
// It returns ok=false if there is no such property.
func (this UncalibratedMagnetometer) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetUncalibratedMagnetometerX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "UncalibratedMagnetometer.y".
//
// It returns ok=false if there is no such property.
func (this UncalibratedMagnetometer) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetUncalibratedMagnetometerY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "UncalibratedMagnetometer.z".
//
// It returns ok=false if there is no such property.
func (this UncalibratedMagnetometer) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetUncalibratedMagnetometerZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// XBias returns the value of property "UncalibratedMagnetometer.xBias".
//
// It returns ok=false if there is no such property.
func (this UncalibratedMagnetometer) XBias() (ret float64, ok bool) {
	ok = js.True == bindings.GetUncalibratedMagnetometerXBias(
		this.ref, js.Pointer(&ret),
	)
	return
}

// YBias returns the value of property "UncalibratedMagnetometer.yBias".
//
// It returns ok=false if there is no such property.
func (this UncalibratedMagnetometer) YBias() (ret float64, ok bool) {
	ok = js.True == bindings.GetUncalibratedMagnetometerYBias(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ZBias returns the value of property "UncalibratedMagnetometer.zBias".
//
// It returns ok=false if there is no such property.
func (this UncalibratedMagnetometer) ZBias() (ret float64, ok bool) {
	ok = js.True == bindings.GetUncalibratedMagnetometerZBias(
		this.ref, js.Pointer(&ret),
	)
	return
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
func (p *UncalibratedMagnetometerReadingValues) UpdateFrom(ref js.Ref) {
	bindings.UncalibratedMagnetometerReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UncalibratedMagnetometerReadingValues) Update(ref js.Ref) {
	bindings.UncalibratedMagnetometerReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UncalibratedMagnetometerReadingValues) FreeMembers(recursive bool) {
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// Signal returns the value of property "WritableStreamDefaultController.signal".
//
// It returns ok=false if there is no such property.
func (this WritableStreamDefaultController) Signal() (ret AbortSignal, ok bool) {
	ok = js.True == bindings.GetWritableStreamDefaultControllerSignal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncError returns true if the method "WritableStreamDefaultController.error" exists.
func (this WritableStreamDefaultController) HasFuncError() bool {
	return js.True == bindings.HasFuncWritableStreamDefaultControllerError(
		this.ref,
	)
}

// FuncError returns the method "WritableStreamDefaultController.error".
func (this WritableStreamDefaultController) FuncError() (fn js.Func[func(e js.Any)]) {
	bindings.FuncWritableStreamDefaultControllerError(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Error calls the method "WritableStreamDefaultController.error".
func (this WritableStreamDefaultController) Error(e js.Any) (ret js.Void) {
	bindings.CallWritableStreamDefaultControllerError(
		this.ref, js.Pointer(&ret),
		e.Ref(),
	)

	return
}

// TryError calls the method "WritableStreamDefaultController.error"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultController) TryError(e js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultControllerError(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		e.Ref(),
	)

	return
}

// HasFuncError1 returns true if the method "WritableStreamDefaultController.error" exists.
func (this WritableStreamDefaultController) HasFuncError1() bool {
	return js.True == bindings.HasFuncWritableStreamDefaultControllerError1(
		this.ref,
	)
}

// FuncError1 returns the method "WritableStreamDefaultController.error".
func (this WritableStreamDefaultController) FuncError1() (fn js.Func[func()]) {
	bindings.FuncWritableStreamDefaultControllerError1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Error1 calls the method "WritableStreamDefaultController.error".
func (this WritableStreamDefaultController) Error1() (ret js.Void) {
	bindings.CallWritableStreamDefaultControllerError1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryError1 calls the method "WritableStreamDefaultController.error"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WritableStreamDefaultController) TryError1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryWritableStreamDefaultControllerError1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
		js.ThrowInvalidCallbackInvocation()
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
		js.ThrowInvalidCallbackInvocation()
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
func (p *UnderlyingSink) UpdateFrom(ref js.Ref) {
	bindings.UnderlyingSinkJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UnderlyingSink) Update(ref js.Ref) {
	bindings.UnderlyingSinkJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UnderlyingSink) FreeMembers(recursive bool) {
	js.Free(
		p.Start.Ref(),
		p.Write.Ref(),
		p.Close.Ref(),
		p.Abort.Ref(),
		p.Type.Ref(),
	)
	p.Start = p.Start.FromRef(js.Undefined)
	p.Write = p.Write.FromRef(js.Undefined)
	p.Close = p.Close.FromRef(js.Undefined)
	p.Abort = p.Abort.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
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
		js.ThrowInvalidCallbackInvocation()
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
		js.ThrowInvalidCallbackInvocation()
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
		js.ThrowInvalidCallbackInvocation()
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
func (p *UnderlyingSource) UpdateFrom(ref js.Ref) {
	bindings.UnderlyingSourceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UnderlyingSource) Update(ref js.Ref) {
	bindings.UnderlyingSourceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UnderlyingSource) FreeMembers(recursive bool) {
	js.Free(
		p.Start.Ref(),
		p.Pull.Ref(),
		p.Cancel.Ref(),
	)
	p.Start = p.Start.FromRef(js.Undefined)
	p.Pull = p.Pull.FromRef(js.Undefined)
	p.Cancel = p.Cancel.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Id returns the value of property "VTTRegion.id".
//
// It returns ok=false if there is no such property.
func (this VTTRegion) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetVTTRegionId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetId sets the value of property "VTTRegion.id" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetId(val js.String) bool {
	return js.True == bindings.SetVTTRegionId(
		this.ref,
		val.Ref(),
	)
}

// Width returns the value of property "VTTRegion.width".
//
// It returns ok=false if there is no such property.
func (this VTTRegion) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetVTTRegionWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "VTTRegion.width" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetWidth(val float64) bool {
	return js.True == bindings.SetVTTRegionWidth(
		this.ref,
		float64(val),
	)
}

// Lines returns the value of property "VTTRegion.lines".
//
// It returns ok=false if there is no such property.
func (this VTTRegion) Lines() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVTTRegionLines(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLines sets the value of property "VTTRegion.lines" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetLines(val uint32) bool {
	return js.True == bindings.SetVTTRegionLines(
		this.ref,
		uint32(val),
	)
}

// RegionAnchorX returns the value of property "VTTRegion.regionAnchorX".
//
// It returns ok=false if there is no such property.
func (this VTTRegion) RegionAnchorX() (ret float64, ok bool) {
	ok = js.True == bindings.GetVTTRegionRegionAnchorX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRegionAnchorX sets the value of property "VTTRegion.regionAnchorX" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetRegionAnchorX(val float64) bool {
	return js.True == bindings.SetVTTRegionRegionAnchorX(
		this.ref,
		float64(val),
	)
}

// RegionAnchorY returns the value of property "VTTRegion.regionAnchorY".
//
// It returns ok=false if there is no such property.
func (this VTTRegion) RegionAnchorY() (ret float64, ok bool) {
	ok = js.True == bindings.GetVTTRegionRegionAnchorY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRegionAnchorY sets the value of property "VTTRegion.regionAnchorY" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetRegionAnchorY(val float64) bool {
	return js.True == bindings.SetVTTRegionRegionAnchorY(
		this.ref,
		float64(val),
	)
}

// ViewportAnchorX returns the value of property "VTTRegion.viewportAnchorX".
//
// It returns ok=false if there is no such property.
func (this VTTRegion) ViewportAnchorX() (ret float64, ok bool) {
	ok = js.True == bindings.GetVTTRegionViewportAnchorX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetViewportAnchorX sets the value of property "VTTRegion.viewportAnchorX" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetViewportAnchorX(val float64) bool {
	return js.True == bindings.SetVTTRegionViewportAnchorX(
		this.ref,
		float64(val),
	)
}

// ViewportAnchorY returns the value of property "VTTRegion.viewportAnchorY".
//
// It returns ok=false if there is no such property.
func (this VTTRegion) ViewportAnchorY() (ret float64, ok bool) {
	ok = js.True == bindings.GetVTTRegionViewportAnchorY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetViewportAnchorY sets the value of property "VTTRegion.viewportAnchorY" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetViewportAnchorY(val float64) bool {
	return js.True == bindings.SetVTTRegionViewportAnchorY(
		this.ref,
		float64(val),
	)
}

// Scroll returns the value of property "VTTRegion.scroll".
//
// It returns ok=false if there is no such property.
func (this VTTRegion) Scroll() (ret ScrollSetting, ok bool) {
	ok = js.True == bindings.GetVTTRegionScroll(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScroll sets the value of property "VTTRegion.scroll" to val.
//
// It returns false if the property cannot be set.
func (this VTTRegion) SetScroll(val ScrollSetting) bool {
	return js.True == bindings.SetVTTRegionScroll(
		this.ref,
		uint32(val),
	)
}

func NewVTTCue(startTime float64, endTime float64, text js.String) (ret VTTCue) {
	ret.ref = bindings.NewVTTCueByVTTCue(
		float64(startTime),
		float64(endTime),
		text.Ref())
	return
}

type VTTCue struct {
	TextTrackCue
}

func (this VTTCue) Once() VTTCue {
	this.ref.Once()
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
	this.ref.Free()
}

// Region returns the value of property "VTTCue.region".
//
// It returns ok=false if there is no such property.
func (this VTTCue) Region() (ret VTTRegion, ok bool) {
	ok = js.True == bindings.GetVTTCueRegion(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRegion sets the value of property "VTTCue.region" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetRegion(val VTTRegion) bool {
	return js.True == bindings.SetVTTCueRegion(
		this.ref,
		val.Ref(),
	)
}

// Vertical returns the value of property "VTTCue.vertical".
//
// It returns ok=false if there is no such property.
func (this VTTCue) Vertical() (ret DirectionSetting, ok bool) {
	ok = js.True == bindings.GetVTTCueVertical(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVertical sets the value of property "VTTCue.vertical" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetVertical(val DirectionSetting) bool {
	return js.True == bindings.SetVTTCueVertical(
		this.ref,
		uint32(val),
	)
}

// SnapToLines returns the value of property "VTTCue.snapToLines".
//
// It returns ok=false if there is no such property.
func (this VTTCue) SnapToLines() (ret bool, ok bool) {
	ok = js.True == bindings.GetVTTCueSnapToLines(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSnapToLines sets the value of property "VTTCue.snapToLines" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetSnapToLines(val bool) bool {
	return js.True == bindings.SetVTTCueSnapToLines(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Line returns the value of property "VTTCue.line".
//
// It returns ok=false if there is no such property.
func (this VTTCue) Line() (ret LineAndPositionSetting, ok bool) {
	ok = js.True == bindings.GetVTTCueLine(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLine sets the value of property "VTTCue.line" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetLine(val LineAndPositionSetting) bool {
	return js.True == bindings.SetVTTCueLine(
		this.ref,
		val.Ref(),
	)
}

// LineAlign returns the value of property "VTTCue.lineAlign".
//
// It returns ok=false if there is no such property.
func (this VTTCue) LineAlign() (ret LineAlignSetting, ok bool) {
	ok = js.True == bindings.GetVTTCueLineAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLineAlign sets the value of property "VTTCue.lineAlign" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetLineAlign(val LineAlignSetting) bool {
	return js.True == bindings.SetVTTCueLineAlign(
		this.ref,
		uint32(val),
	)
}

// Position returns the value of property "VTTCue.position".
//
// It returns ok=false if there is no such property.
func (this VTTCue) Position() (ret LineAndPositionSetting, ok bool) {
	ok = js.True == bindings.GetVTTCuePosition(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPosition sets the value of property "VTTCue.position" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetPosition(val LineAndPositionSetting) bool {
	return js.True == bindings.SetVTTCuePosition(
		this.ref,
		val.Ref(),
	)
}

// PositionAlign returns the value of property "VTTCue.positionAlign".
//
// It returns ok=false if there is no such property.
func (this VTTCue) PositionAlign() (ret PositionAlignSetting, ok bool) {
	ok = js.True == bindings.GetVTTCuePositionAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPositionAlign sets the value of property "VTTCue.positionAlign" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetPositionAlign(val PositionAlignSetting) bool {
	return js.True == bindings.SetVTTCuePositionAlign(
		this.ref,
		uint32(val),
	)
}

// Size returns the value of property "VTTCue.size".
//
// It returns ok=false if there is no such property.
func (this VTTCue) Size() (ret float64, ok bool) {
	ok = js.True == bindings.GetVTTCueSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSize sets the value of property "VTTCue.size" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetSize(val float64) bool {
	return js.True == bindings.SetVTTCueSize(
		this.ref,
		float64(val),
	)
}

// Align returns the value of property "VTTCue.align".
//
// It returns ok=false if there is no such property.
func (this VTTCue) Align() (ret AlignSetting, ok bool) {
	ok = js.True == bindings.GetVTTCueAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "VTTCue.align" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetAlign(val AlignSetting) bool {
	return js.True == bindings.SetVTTCueAlign(
		this.ref,
		uint32(val),
	)
}

// Text returns the value of property "VTTCue.text".
//
// It returns ok=false if there is no such property.
func (this VTTCue) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetVTTCueText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "VTTCue.text" to val.
//
// It returns false if the property cannot be set.
func (this VTTCue) SetText(val js.String) bool {
	return js.True == bindings.SetVTTCueText(
		this.ref,
		val.Ref(),
	)
}

// HasFuncGetCueAsHTML returns true if the method "VTTCue.getCueAsHTML" exists.
func (this VTTCue) HasFuncGetCueAsHTML() bool {
	return js.True == bindings.HasFuncVTTCueGetCueAsHTML(
		this.ref,
	)
}

// FuncGetCueAsHTML returns the method "VTTCue.getCueAsHTML".
func (this VTTCue) FuncGetCueAsHTML() (fn js.Func[func() DocumentFragment]) {
	bindings.FuncVTTCueGetCueAsHTML(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCueAsHTML calls the method "VTTCue.getCueAsHTML".
func (this VTTCue) GetCueAsHTML() (ret DocumentFragment) {
	bindings.CallVTTCueGetCueAsHTML(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetCueAsHTML calls the method "VTTCue.getCueAsHTML"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VTTCue) TryGetCueAsHTML() (ret DocumentFragment, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVTTCueGetCueAsHTML(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
func (p *ValueEventInit) UpdateFrom(ref js.Ref) {
	bindings.ValueEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ValueEventInit) Update(ref js.Ref) {
	bindings.ValueEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ValueEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

func NewValueEvent(typ js.String, initDict ValueEventInit) (ret ValueEvent) {
	ret.ref = bindings.NewValueEventByValueEvent(
		typ.Ref(),
		js.Pointer(&initDict))
	return
}

func NewValueEventByValueEvent1(typ js.String) (ret ValueEvent) {
	ret.ref = bindings.NewValueEventByValueEvent1(
		typ.Ref())
	return
}

type ValueEvent struct {
	Event
}

func (this ValueEvent) Once() ValueEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "ValueEvent.value".
//
// It returns ok=false if there is no such property.
func (this ValueEvent) Value() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetValueEventValue(
		this.ref, js.Pointer(&ret),
	)
	return
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
		js.ThrowInvalidCallbackInvocation()
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
func (p *VideoDecoderInit) UpdateFrom(ref js.Ref) {
	bindings.VideoDecoderInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoDecoderInit) Update(ref js.Ref) {
	bindings.VideoDecoderInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoDecoderInit) FreeMembers(recursive bool) {
	js.Free(
		p.Output.Ref(),
		p.Error.Ref(),
	)
	p.Output = p.Output.FromRef(js.Undefined)
	p.Error = p.Error.FromRef(js.Undefined)
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
	//
	// NOTE: Config.FFI_USE MUST be set to true to get Config used.
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
func (p *VideoDecoderSupport) UpdateFrom(ref js.Ref) {
	bindings.VideoDecoderSupportJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoDecoderSupport) Update(ref js.Ref) {
	bindings.VideoDecoderSupportJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoDecoderSupport) FreeMembers(recursive bool) {
	if recursive {
		p.Config.FreeMembers(true)
	}
}

func NewVideoDecoder(init VideoDecoderInit) (ret VideoDecoder) {
	ret.ref = bindings.NewVideoDecoderByVideoDecoder(
		js.Pointer(&init))
	return
}

type VideoDecoder struct {
	EventTarget
}

func (this VideoDecoder) Once() VideoDecoder {
	this.ref.Once()
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
	this.ref.Free()
}

// State returns the value of property "VideoDecoder.state".
//
// It returns ok=false if there is no such property.
func (this VideoDecoder) State() (ret CodecState, ok bool) {
	ok = js.True == bindings.GetVideoDecoderState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DecodeQueueSize returns the value of property "VideoDecoder.decodeQueueSize".
//
// It returns ok=false if there is no such property.
func (this VideoDecoder) DecodeQueueSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetVideoDecoderDecodeQueueSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncConfigure returns true if the method "VideoDecoder.configure" exists.
func (this VideoDecoder) HasFuncConfigure() bool {
	return js.True == bindings.HasFuncVideoDecoderConfigure(
		this.ref,
	)
}

// FuncConfigure returns the method "VideoDecoder.configure".
func (this VideoDecoder) FuncConfigure() (fn js.Func[func(config VideoDecoderConfig)]) {
	bindings.FuncVideoDecoderConfigure(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Configure calls the method "VideoDecoder.configure".
func (this VideoDecoder) Configure(config VideoDecoderConfig) (ret js.Void) {
	bindings.CallVideoDecoderConfigure(
		this.ref, js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryConfigure calls the method "VideoDecoder.configure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoDecoder) TryConfigure(config VideoDecoderConfig) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoDecoderConfigure(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

// HasFuncDecode returns true if the method "VideoDecoder.decode" exists.
func (this VideoDecoder) HasFuncDecode() bool {
	return js.True == bindings.HasFuncVideoDecoderDecode(
		this.ref,
	)
}

// FuncDecode returns the method "VideoDecoder.decode".
func (this VideoDecoder) FuncDecode() (fn js.Func[func(chunk EncodedVideoChunk)]) {
	bindings.FuncVideoDecoderDecode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Decode calls the method "VideoDecoder.decode".
func (this VideoDecoder) Decode(chunk EncodedVideoChunk) (ret js.Void) {
	bindings.CallVideoDecoderDecode(
		this.ref, js.Pointer(&ret),
		chunk.Ref(),
	)

	return
}

// TryDecode calls the method "VideoDecoder.decode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoDecoder) TryDecode(chunk EncodedVideoChunk) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoDecoderDecode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		chunk.Ref(),
	)

	return
}

// HasFuncFlush returns true if the method "VideoDecoder.flush" exists.
func (this VideoDecoder) HasFuncFlush() bool {
	return js.True == bindings.HasFuncVideoDecoderFlush(
		this.ref,
	)
}

// FuncFlush returns the method "VideoDecoder.flush".
func (this VideoDecoder) FuncFlush() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncVideoDecoderFlush(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Flush calls the method "VideoDecoder.flush".
func (this VideoDecoder) Flush() (ret js.Promise[js.Void]) {
	bindings.CallVideoDecoderFlush(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "VideoDecoder.flush"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoDecoder) TryFlush() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoDecoderFlush(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReset returns true if the method "VideoDecoder.reset" exists.
func (this VideoDecoder) HasFuncReset() bool {
	return js.True == bindings.HasFuncVideoDecoderReset(
		this.ref,
	)
}

// FuncReset returns the method "VideoDecoder.reset".
func (this VideoDecoder) FuncReset() (fn js.Func[func()]) {
	bindings.FuncVideoDecoderReset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reset calls the method "VideoDecoder.reset".
func (this VideoDecoder) Reset() (ret js.Void) {
	bindings.CallVideoDecoderReset(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "VideoDecoder.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoDecoder) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoDecoderReset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "VideoDecoder.close" exists.
func (this VideoDecoder) HasFuncClose() bool {
	return js.True == bindings.HasFuncVideoDecoderClose(
		this.ref,
	)
}

// FuncClose returns the method "VideoDecoder.close".
func (this VideoDecoder) FuncClose() (fn js.Func[func()]) {
	bindings.FuncVideoDecoderClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "VideoDecoder.close".
func (this VideoDecoder) Close() (ret js.Void) {
	bindings.CallVideoDecoderClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "VideoDecoder.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoDecoder) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoDecoderClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsConfigSupported returns true if the static method "VideoDecoder.isConfigSupported" exists.
func (this VideoDecoder) HasFuncIsConfigSupported() bool {
	return js.True == bindings.HasFuncVideoDecoderIsConfigSupported(
		this.ref,
	)
}

// FuncIsConfigSupported returns the static method "VideoDecoder.isConfigSupported".
func (this VideoDecoder) FuncIsConfigSupported() (fn js.Func[func(config VideoDecoderConfig) js.Promise[VideoDecoderSupport]]) {
	bindings.FuncVideoDecoderIsConfigSupported(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsConfigSupported calls the static method "VideoDecoder.isConfigSupported".
func (this VideoDecoder) IsConfigSupported(config VideoDecoderConfig) (ret js.Promise[VideoDecoderSupport]) {
	bindings.CallVideoDecoderIsConfigSupported(
		this.ref, js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryIsConfigSupported calls the static method "VideoDecoder.isConfigSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this VideoDecoder) TryIsConfigSupported(config VideoDecoderConfig) (ret js.Promise[VideoDecoderSupport], exception js.Any, ok bool) {
	ok = js.True == bindings.TryVideoDecoderIsConfigSupported(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

type VideoEncoderInit struct {
	// Output is "VideoEncoderInit.output"
	//
	// Required
	Output js.Func[func(chunk EncodedVideoChunk, metadata *EncodedVideoChunkMetadata)]
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
func (p *VideoEncoderInit) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoEncoderInit) Update(ref js.Ref) {
	bindings.VideoEncoderInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoEncoderInit) FreeMembers(recursive bool) {
	js.Free(
		p.Output.Ref(),
		p.Error.Ref(),
	)
	p.Output = p.Output.FromRef(js.Undefined)
	p.Error = p.Error.FromRef(js.Undefined)
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
	//
	// NOTE: Hevc.FFI_USE MUST be set to true to get Hevc used.
	Hevc HevcEncoderConfig
	// Avc is "VideoEncoderConfig.avc"
	//
	// Optional
	//
	// NOTE: Avc.FFI_USE MUST be set to true to get Avc used.
	Avc AvcEncoderConfig
	// Av1 is "VideoEncoderConfig.av1"
	//
	// Optional
	//
	// NOTE: Av1.FFI_USE MUST be set to true to get Av1 used.
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
func (p *VideoEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoEncoderConfig) Update(ref js.Ref) {
	bindings.VideoEncoderConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoEncoderConfig) FreeMembers(recursive bool) {
	js.Free(
		p.Codec.Ref(),
		p.ScalabilityMode.Ref(),
	)
	p.Codec = p.Codec.FromRef(js.Undefined)
	p.ScalabilityMode = p.ScalabilityMode.FromRef(js.Undefined)
	if recursive {
		p.Hevc.FreeMembers(true)
		p.Avc.FreeMembers(true)
		p.Av1.FreeMembers(true)
	}
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
func (p *VideoEncoderEncodeOptionsForHevc) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForHevcJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoEncoderEncodeOptionsForHevc) Update(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForHevcJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoEncoderEncodeOptionsForHevc) FreeMembers(recursive bool) {
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
func (p *VideoEncoderEncodeOptionsForVp9) UpdateFrom(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForVp9JSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VideoEncoderEncodeOptionsForVp9) Update(ref js.Ref) {
	bindings.VideoEncoderEncodeOptionsForVp9JSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VideoEncoderEncodeOptionsForVp9) FreeMembers(recursive bool) {
}
