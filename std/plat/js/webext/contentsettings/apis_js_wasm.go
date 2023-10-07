// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package contentsettings

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/contentsettings/bindings"
)

type AutoVerifyContentSetting uint32

const (
	_ AutoVerifyContentSetting = iota

	AutoVerifyContentSetting_ALLOW
	AutoVerifyContentSetting_BLOCK
)

func (AutoVerifyContentSetting) FromRef(str js.Ref) AutoVerifyContentSetting {
	return AutoVerifyContentSetting(bindings.ConstOfAutoVerifyContentSetting(str))
}

func (x AutoVerifyContentSetting) String() (string, bool) {
	switch x {
	case AutoVerifyContentSetting_ALLOW:
		return "allow", true
	case AutoVerifyContentSetting_BLOCK:
		return "block", true
	default:
		return "", false
	}
}

type CameraContentSetting uint32

const (
	_ CameraContentSetting = iota

	CameraContentSetting_ALLOW
	CameraContentSetting_BLOCK
	CameraContentSetting_ASK
)

func (CameraContentSetting) FromRef(str js.Ref) CameraContentSetting {
	return CameraContentSetting(bindings.ConstOfCameraContentSetting(str))
}

func (x CameraContentSetting) String() (string, bool) {
	switch x {
	case CameraContentSetting_ALLOW:
		return "allow", true
	case CameraContentSetting_BLOCK:
		return "block", true
	case CameraContentSetting_ASK:
		return "ask", true
	default:
		return "", false
	}
}

type Scope uint32

const (
	_ Scope = iota

	Scope_REGULAR
	Scope_INCOGNITO_SESSION_ONLY
)

func (Scope) FromRef(str js.Ref) Scope {
	return Scope(bindings.ConstOfScope(str))
}

func (x Scope) String() (string, bool) {
	switch x {
	case Scope_REGULAR:
		return "regular", true
	case Scope_INCOGNITO_SESSION_ONLY:
		return "incognito_session_only", true
	default:
		return "", false
	}
}

type ClearArgDetails struct {
	// Scope is "ClearArgDetails.scope"
	//
	// Optional
	Scope Scope

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClearArgDetails with all fields set.
func (p ClearArgDetails) FromRef(ref js.Ref) ClearArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClearArgDetails in the application heap.
func (p ClearArgDetails) New() js.Ref {
	return bindings.ClearArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClearArgDetails) UpdateFrom(ref js.Ref) {
	bindings.ClearArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClearArgDetails) Update(ref js.Ref) {
	bindings.ClearArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClearArgDetails) FreeMembers(recursive bool) {
}

type GetReturnType struct {
	// Setting is "GetReturnType.setting"
	//
	// Required
	Setting js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetReturnType with all fields set.
func (p GetReturnType) FromRef(ref js.Ref) GetReturnType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetReturnType in the application heap.
func (p GetReturnType) New() js.Ref {
	return bindings.GetReturnTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetReturnType) UpdateFrom(ref js.Ref) {
	bindings.GetReturnTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetReturnType) Update(ref js.Ref) {
	bindings.GetReturnTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetReturnType) FreeMembers(recursive bool) {
	js.Free(
		p.Setting.Ref(),
	)
	p.Setting = p.Setting.FromRef(js.Undefined)
}

type ResourceIdentifier struct {
	// Description is "ResourceIdentifier.description"
	//
	// Optional
	Description js.String
	// Id is "ResourceIdentifier.id"
	//
	// Required
	Id js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ResourceIdentifier with all fields set.
func (p ResourceIdentifier) FromRef(ref js.Ref) ResourceIdentifier {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ResourceIdentifier in the application heap.
func (p ResourceIdentifier) New() js.Ref {
	return bindings.ResourceIdentifierJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ResourceIdentifier) UpdateFrom(ref js.Ref) {
	bindings.ResourceIdentifierJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ResourceIdentifier) Update(ref js.Ref) {
	bindings.ResourceIdentifierJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ResourceIdentifier) FreeMembers(recursive bool) {
	js.Free(
		p.Description.Ref(),
		p.Id.Ref(),
	)
	p.Description = p.Description.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
}

type GetArgDetails struct {
	// Incognito is "GetArgDetails.incognito"
	//
	// Optional
	//
	// NOTE: FFI_USE_Incognito MUST be set to true to make this field effective.
	Incognito bool
	// PrimaryUrl is "GetArgDetails.primaryUrl"
	//
	// Required
	PrimaryUrl js.String
	// ResourceIdentifier is "GetArgDetails.resourceIdentifier"
	//
	// Optional
	//
	// NOTE: ResourceIdentifier.FFI_USE MUST be set to true to get ResourceIdentifier used.
	ResourceIdentifier ResourceIdentifier
	// SecondaryUrl is "GetArgDetails.secondaryUrl"
	//
	// Optional
	SecondaryUrl js.String

	FFI_USE_Incognito bool // for Incognito.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetArgDetails with all fields set.
func (p GetArgDetails) FromRef(ref js.Ref) GetArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetArgDetails in the application heap.
func (p GetArgDetails) New() js.Ref {
	return bindings.GetArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetArgDetails) UpdateFrom(ref js.Ref) {
	bindings.GetArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetArgDetails) Update(ref js.Ref) {
	bindings.GetArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.PrimaryUrl.Ref(),
		p.SecondaryUrl.Ref(),
	)
	p.PrimaryUrl = p.PrimaryUrl.FromRef(js.Undefined)
	p.SecondaryUrl = p.SecondaryUrl.FromRef(js.Undefined)
	if recursive {
		p.ResourceIdentifier.FreeMembers(true)
	}
}

type SetArgDetails struct {
	// PrimaryPattern is "SetArgDetails.primaryPattern"
	//
	// Required
	PrimaryPattern js.String
	// ResourceIdentifier is "SetArgDetails.resourceIdentifier"
	//
	// Optional
	//
	// NOTE: ResourceIdentifier.FFI_USE MUST be set to true to get ResourceIdentifier used.
	ResourceIdentifier ResourceIdentifier
	// Scope is "SetArgDetails.scope"
	//
	// Optional
	Scope Scope
	// SecondaryPattern is "SetArgDetails.secondaryPattern"
	//
	// Optional
	SecondaryPattern js.String
	// Setting is "SetArgDetails.setting"
	//
	// Required
	Setting js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetArgDetails with all fields set.
func (p SetArgDetails) FromRef(ref js.Ref) SetArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetArgDetails in the application heap.
func (p SetArgDetails) New() js.Ref {
	return bindings.SetArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetArgDetails) UpdateFrom(ref js.Ref) {
	bindings.SetArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetArgDetails) Update(ref js.Ref) {
	bindings.SetArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.PrimaryPattern.Ref(),
		p.SecondaryPattern.Ref(),
		p.Setting.Ref(),
	)
	p.PrimaryPattern = p.PrimaryPattern.FromRef(js.Undefined)
	p.SecondaryPattern = p.SecondaryPattern.FromRef(js.Undefined)
	p.Setting = p.Setting.FromRef(js.Undefined)
	if recursive {
		p.ResourceIdentifier.FreeMembers(true)
	}
}

type ContentSetting struct {
	ref js.Ref
}

func (this ContentSetting) Once() ContentSetting {
	this.ref.Once()
	return this
}

func (this ContentSetting) Ref() js.Ref {
	return this.ref
}

func (this ContentSetting) FromRef(ref js.Ref) ContentSetting {
	this.ref = ref
	return this
}

func (this ContentSetting) Free() {
	this.ref.Free()
}

// HasFuncClear returns true if the method "ContentSetting.clear" exists.
func (this ContentSetting) HasFuncClear() bool {
	return js.True == bindings.HasFuncContentSettingClear(
		this.ref,
	)
}

// FuncClear returns the method "ContentSetting.clear".
func (this ContentSetting) FuncClear() (fn js.Func[func(details ClearArgDetails) js.Promise[js.Void]]) {
	bindings.FuncContentSettingClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "ContentSetting.clear".
func (this ContentSetting) Clear(details ClearArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallContentSettingClear(
		this.ref, js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryClear calls the method "ContentSetting.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContentSetting) TryClear(details ClearArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContentSettingClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGet returns true if the method "ContentSetting.get" exists.
func (this ContentSetting) HasFuncGet() bool {
	return js.True == bindings.HasFuncContentSettingGet(
		this.ref,
	)
}

// FuncGet returns the method "ContentSetting.get".
func (this ContentSetting) FuncGet() (fn js.Func[func(details GetArgDetails) js.Promise[GetReturnType]]) {
	bindings.FuncContentSettingGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "ContentSetting.get".
func (this ContentSetting) Get(details GetArgDetails) (ret js.Promise[GetReturnType]) {
	bindings.CallContentSettingGet(
		this.ref, js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGet calls the method "ContentSetting.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContentSetting) TryGet(details GetArgDetails) (ret js.Promise[GetReturnType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContentSettingGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncSet returns true if the method "ContentSetting.set" exists.
func (this ContentSetting) HasFuncSet() bool {
	return js.True == bindings.HasFuncContentSettingSet(
		this.ref,
	)
}

// FuncSet returns the method "ContentSetting.set".
func (this ContentSetting) FuncSet() (fn js.Func[func(details SetArgDetails) js.Promise[js.Void]]) {
	bindings.FuncContentSettingSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "ContentSetting.set".
func (this ContentSetting) Set(details SetArgDetails) (ret js.Promise[js.Void]) {
	bindings.CallContentSettingSet(
		this.ref, js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TrySet calls the method "ContentSetting.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContentSetting) TrySet(details SetArgDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContentSettingSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetResourceIdentifiers returns true if the method "ContentSetting.getResourceIdentifiers" exists.
func (this ContentSetting) HasFuncGetResourceIdentifiers() bool {
	return js.True == bindings.HasFuncContentSettingGetResourceIdentifiers(
		this.ref,
	)
}

// FuncGetResourceIdentifiers returns the method "ContentSetting.getResourceIdentifiers".
func (this ContentSetting) FuncGetResourceIdentifiers() (fn js.Func[func() js.Promise[js.Array[ResourceIdentifier]]]) {
	bindings.FuncContentSettingGetResourceIdentifiers(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetResourceIdentifiers calls the method "ContentSetting.getResourceIdentifiers".
func (this ContentSetting) GetResourceIdentifiers() (ret js.Promise[js.Array[ResourceIdentifier]]) {
	bindings.CallContentSettingGetResourceIdentifiers(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetResourceIdentifiers calls the method "ContentSetting.getResourceIdentifiers"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContentSetting) TryGetResourceIdentifiers() (ret js.Promise[js.Array[ResourceIdentifier]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContentSettingGetResourceIdentifiers(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type CookiesContentSetting uint32

const (
	_ CookiesContentSetting = iota

	CookiesContentSetting_ALLOW
	CookiesContentSetting_BLOCK
	CookiesContentSetting_SESSION_ONLY
)

func (CookiesContentSetting) FromRef(str js.Ref) CookiesContentSetting {
	return CookiesContentSetting(bindings.ConstOfCookiesContentSetting(str))
}

func (x CookiesContentSetting) String() (string, bool) {
	switch x {
	case CookiesContentSetting_ALLOW:
		return "allow", true
	case CookiesContentSetting_BLOCK:
		return "block", true
	case CookiesContentSetting_SESSION_ONLY:
		return "session_only", true
	default:
		return "", false
	}
}

type FullscreenContentSetting uint32

const (
	_ FullscreenContentSetting = iota

	FullscreenContentSetting_ALLOW
)

func (FullscreenContentSetting) FromRef(str js.Ref) FullscreenContentSetting {
	return FullscreenContentSetting(bindings.ConstOfFullscreenContentSetting(str))
}

func (x FullscreenContentSetting) String() (string, bool) {
	switch x {
	case FullscreenContentSetting_ALLOW:
		return "allow", true
	default:
		return "", false
	}
}

type ImagesContentSetting uint32

const (
	_ ImagesContentSetting = iota

	ImagesContentSetting_ALLOW
	ImagesContentSetting_BLOCK
)

func (ImagesContentSetting) FromRef(str js.Ref) ImagesContentSetting {
	return ImagesContentSetting(bindings.ConstOfImagesContentSetting(str))
}

func (x ImagesContentSetting) String() (string, bool) {
	switch x {
	case ImagesContentSetting_ALLOW:
		return "allow", true
	case ImagesContentSetting_BLOCK:
		return "block", true
	default:
		return "", false
	}
}

type JavascriptContentSetting uint32

const (
	_ JavascriptContentSetting = iota

	JavascriptContentSetting_ALLOW
	JavascriptContentSetting_BLOCK
)

func (JavascriptContentSetting) FromRef(str js.Ref) JavascriptContentSetting {
	return JavascriptContentSetting(bindings.ConstOfJavascriptContentSetting(str))
}

func (x JavascriptContentSetting) String() (string, bool) {
	switch x {
	case JavascriptContentSetting_ALLOW:
		return "allow", true
	case JavascriptContentSetting_BLOCK:
		return "block", true
	default:
		return "", false
	}
}

type LocationContentSetting uint32

const (
	_ LocationContentSetting = iota

	LocationContentSetting_ALLOW
	LocationContentSetting_BLOCK
	LocationContentSetting_ASK
)

func (LocationContentSetting) FromRef(str js.Ref) LocationContentSetting {
	return LocationContentSetting(bindings.ConstOfLocationContentSetting(str))
}

func (x LocationContentSetting) String() (string, bool) {
	switch x {
	case LocationContentSetting_ALLOW:
		return "allow", true
	case LocationContentSetting_BLOCK:
		return "block", true
	case LocationContentSetting_ASK:
		return "ask", true
	default:
		return "", false
	}
}

type MicrophoneContentSetting uint32

const (
	_ MicrophoneContentSetting = iota

	MicrophoneContentSetting_ALLOW
	MicrophoneContentSetting_BLOCK
	MicrophoneContentSetting_ASK
)

func (MicrophoneContentSetting) FromRef(str js.Ref) MicrophoneContentSetting {
	return MicrophoneContentSetting(bindings.ConstOfMicrophoneContentSetting(str))
}

func (x MicrophoneContentSetting) String() (string, bool) {
	switch x {
	case MicrophoneContentSetting_ALLOW:
		return "allow", true
	case MicrophoneContentSetting_BLOCK:
		return "block", true
	case MicrophoneContentSetting_ASK:
		return "ask", true
	default:
		return "", false
	}
}

type MouselockContentSetting uint32

const (
	_ MouselockContentSetting = iota

	MouselockContentSetting_ALLOW
)

func (MouselockContentSetting) FromRef(str js.Ref) MouselockContentSetting {
	return MouselockContentSetting(bindings.ConstOfMouselockContentSetting(str))
}

func (x MouselockContentSetting) String() (string, bool) {
	switch x {
	case MouselockContentSetting_ALLOW:
		return "allow", true
	default:
		return "", false
	}
}

type MultipleAutomaticDownloadsContentSetting uint32

const (
	_ MultipleAutomaticDownloadsContentSetting = iota

	MultipleAutomaticDownloadsContentSetting_ALLOW
	MultipleAutomaticDownloadsContentSetting_BLOCK
	MultipleAutomaticDownloadsContentSetting_ASK
)

func (MultipleAutomaticDownloadsContentSetting) FromRef(str js.Ref) MultipleAutomaticDownloadsContentSetting {
	return MultipleAutomaticDownloadsContentSetting(bindings.ConstOfMultipleAutomaticDownloadsContentSetting(str))
}

func (x MultipleAutomaticDownloadsContentSetting) String() (string, bool) {
	switch x {
	case MultipleAutomaticDownloadsContentSetting_ALLOW:
		return "allow", true
	case MultipleAutomaticDownloadsContentSetting_BLOCK:
		return "block", true
	case MultipleAutomaticDownloadsContentSetting_ASK:
		return "ask", true
	default:
		return "", false
	}
}

type NotificationsContentSetting uint32

const (
	_ NotificationsContentSetting = iota

	NotificationsContentSetting_ALLOW
	NotificationsContentSetting_BLOCK
	NotificationsContentSetting_ASK
)

func (NotificationsContentSetting) FromRef(str js.Ref) NotificationsContentSetting {
	return NotificationsContentSetting(bindings.ConstOfNotificationsContentSetting(str))
}

func (x NotificationsContentSetting) String() (string, bool) {
	switch x {
	case NotificationsContentSetting_ALLOW:
		return "allow", true
	case NotificationsContentSetting_BLOCK:
		return "block", true
	case NotificationsContentSetting_ASK:
		return "ask", true
	default:
		return "", false
	}
}

type PluginsContentSetting uint32

const (
	_ PluginsContentSetting = iota

	PluginsContentSetting_BLOCK
)

func (PluginsContentSetting) FromRef(str js.Ref) PluginsContentSetting {
	return PluginsContentSetting(bindings.ConstOfPluginsContentSetting(str))
}

func (x PluginsContentSetting) String() (string, bool) {
	switch x {
	case PluginsContentSetting_BLOCK:
		return "block", true
	default:
		return "", false
	}
}

type PopupsContentSetting uint32

const (
	_ PopupsContentSetting = iota

	PopupsContentSetting_ALLOW
	PopupsContentSetting_BLOCK
)

func (PopupsContentSetting) FromRef(str js.Ref) PopupsContentSetting {
	return PopupsContentSetting(bindings.ConstOfPopupsContentSetting(str))
}

func (x PopupsContentSetting) String() (string, bool) {
	switch x {
	case PopupsContentSetting_ALLOW:
		return "allow", true
	case PopupsContentSetting_BLOCK:
		return "block", true
	default:
		return "", false
	}
}

type PpapiBrokerContentSetting uint32

const (
	_ PpapiBrokerContentSetting = iota

	PpapiBrokerContentSetting_BLOCK
)

func (PpapiBrokerContentSetting) FromRef(str js.Ref) PpapiBrokerContentSetting {
	return PpapiBrokerContentSetting(bindings.ConstOfPpapiBrokerContentSetting(str))
}

func (x PpapiBrokerContentSetting) String() (string, bool) {
	switch x {
	case PpapiBrokerContentSetting_BLOCK:
		return "block", true
	default:
		return "", false
	}
}

// AutoVerify returns the value of property "WEBEXT.contentSettings.autoVerify".
//
// The returned bool will be false if there is no such property.
func AutoVerify() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetAutoVerify(
		js.Pointer(&ret),
	)

	return
}

// SetAutoVerify sets the value of property "WEBEXT.contentSettings.autoVerify" to val.
//
// It returns false if the property cannot be set.
func SetAutoVerify(val ContentSetting) bool {
	return js.True == bindings.SetAutoVerify(
		val.Ref())
}

// AutomaticDownloads returns the value of property "WEBEXT.contentSettings.automaticDownloads".
//
// The returned bool will be false if there is no such property.
func AutomaticDownloads() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetAutomaticDownloads(
		js.Pointer(&ret),
	)

	return
}

// SetAutomaticDownloads sets the value of property "WEBEXT.contentSettings.automaticDownloads" to val.
//
// It returns false if the property cannot be set.
func SetAutomaticDownloads(val ContentSetting) bool {
	return js.True == bindings.SetAutomaticDownloads(
		val.Ref())
}

// Camera returns the value of property "WEBEXT.contentSettings.camera".
//
// The returned bool will be false if there is no such property.
func Camera() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetCamera(
		js.Pointer(&ret),
	)

	return
}

// SetCamera sets the value of property "WEBEXT.contentSettings.camera" to val.
//
// It returns false if the property cannot be set.
func SetCamera(val ContentSetting) bool {
	return js.True == bindings.SetCamera(
		val.Ref())
}

// Cookies returns the value of property "WEBEXT.contentSettings.cookies".
//
// The returned bool will be false if there is no such property.
func Cookies() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetCookies(
		js.Pointer(&ret),
	)

	return
}

// SetCookies sets the value of property "WEBEXT.contentSettings.cookies" to val.
//
// It returns false if the property cannot be set.
func SetCookies(val ContentSetting) bool {
	return js.True == bindings.SetCookies(
		val.Ref())
}

// Fullscreen returns the value of property "WEBEXT.contentSettings.fullscreen".
//
// The returned bool will be false if there is no such property.
func Fullscreen() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetFullscreen(
		js.Pointer(&ret),
	)

	return
}

// SetFullscreen sets the value of property "WEBEXT.contentSettings.fullscreen" to val.
//
// It returns false if the property cannot be set.
func SetFullscreen(val ContentSetting) bool {
	return js.True == bindings.SetFullscreen(
		val.Ref())
}

// Images returns the value of property "WEBEXT.contentSettings.images".
//
// The returned bool will be false if there is no such property.
func Images() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetImages(
		js.Pointer(&ret),
	)

	return
}

// SetImages sets the value of property "WEBEXT.contentSettings.images" to val.
//
// It returns false if the property cannot be set.
func SetImages(val ContentSetting) bool {
	return js.True == bindings.SetImages(
		val.Ref())
}

// Javascript returns the value of property "WEBEXT.contentSettings.javascript".
//
// The returned bool will be false if there is no such property.
func Javascript() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetJavascript(
		js.Pointer(&ret),
	)

	return
}

// SetJavascript sets the value of property "WEBEXT.contentSettings.javascript" to val.
//
// It returns false if the property cannot be set.
func SetJavascript(val ContentSetting) bool {
	return js.True == bindings.SetJavascript(
		val.Ref())
}

// Location returns the value of property "WEBEXT.contentSettings.location".
//
// The returned bool will be false if there is no such property.
func Location() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetLocation(
		js.Pointer(&ret),
	)

	return
}

// SetLocation sets the value of property "WEBEXT.contentSettings.location" to val.
//
// It returns false if the property cannot be set.
func SetLocation(val ContentSetting) bool {
	return js.True == bindings.SetLocation(
		val.Ref())
}

// Microphone returns the value of property "WEBEXT.contentSettings.microphone".
//
// The returned bool will be false if there is no such property.
func Microphone() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetMicrophone(
		js.Pointer(&ret),
	)

	return
}

// SetMicrophone sets the value of property "WEBEXT.contentSettings.microphone" to val.
//
// It returns false if the property cannot be set.
func SetMicrophone(val ContentSetting) bool {
	return js.True == bindings.SetMicrophone(
		val.Ref())
}

// Mouselock returns the value of property "WEBEXT.contentSettings.mouselock".
//
// The returned bool will be false if there is no such property.
func Mouselock() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetMouselock(
		js.Pointer(&ret),
	)

	return
}

// SetMouselock sets the value of property "WEBEXT.contentSettings.mouselock" to val.
//
// It returns false if the property cannot be set.
func SetMouselock(val ContentSetting) bool {
	return js.True == bindings.SetMouselock(
		val.Ref())
}

// Notifications returns the value of property "WEBEXT.contentSettings.notifications".
//
// The returned bool will be false if there is no such property.
func Notifications() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetNotifications(
		js.Pointer(&ret),
	)

	return
}

// SetNotifications sets the value of property "WEBEXT.contentSettings.notifications" to val.
//
// It returns false if the property cannot be set.
func SetNotifications(val ContentSetting) bool {
	return js.True == bindings.SetNotifications(
		val.Ref())
}

// Plugins returns the value of property "WEBEXT.contentSettings.plugins".
//
// The returned bool will be false if there is no such property.
func Plugins() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetPlugins(
		js.Pointer(&ret),
	)

	return
}

// SetPlugins sets the value of property "WEBEXT.contentSettings.plugins" to val.
//
// It returns false if the property cannot be set.
func SetPlugins(val ContentSetting) bool {
	return js.True == bindings.SetPlugins(
		val.Ref())
}

// Popups returns the value of property "WEBEXT.contentSettings.popups".
//
// The returned bool will be false if there is no such property.
func Popups() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetPopups(
		js.Pointer(&ret),
	)

	return
}

// SetPopups sets the value of property "WEBEXT.contentSettings.popups" to val.
//
// It returns false if the property cannot be set.
func SetPopups(val ContentSetting) bool {
	return js.True == bindings.SetPopups(
		val.Ref())
}

// UnsandboxedPlugins returns the value of property "WEBEXT.contentSettings.unsandboxedPlugins".
//
// The returned bool will be false if there is no such property.
func UnsandboxedPlugins() (ret ContentSetting, ok bool) {
	ok = js.True == bindings.GetUnsandboxedPlugins(
		js.Pointer(&ret),
	)

	return
}

// SetUnsandboxedPlugins sets the value of property "WEBEXT.contentSettings.unsandboxedPlugins" to val.
//
// It returns false if the property cannot be set.
func SetUnsandboxedPlugins(val ContentSetting) bool {
	return js.True == bindings.SetUnsandboxedPlugins(
		val.Ref())
}
