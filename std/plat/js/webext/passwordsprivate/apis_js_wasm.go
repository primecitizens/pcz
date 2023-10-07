// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package passwordsprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/passwordsprivate/bindings"
)

type AddPasswordOptions struct {
	// Url is "AddPasswordOptions.url"
	//
	// Optional
	Url js.String
	// Username is "AddPasswordOptions.username"
	//
	// Optional
	Username js.String
	// Password is "AddPasswordOptions.password"
	//
	// Optional
	Password js.String
	// Note is "AddPasswordOptions.note"
	//
	// Optional
	Note js.String
	// UseAccountStore is "AddPasswordOptions.useAccountStore"
	//
	// Optional
	//
	// NOTE: FFI_USE_UseAccountStore MUST be set to true to make this field effective.
	UseAccountStore bool

	FFI_USE_UseAccountStore bool // for UseAccountStore.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddPasswordOptions with all fields set.
func (p AddPasswordOptions) FromRef(ref js.Ref) AddPasswordOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddPasswordOptions in the application heap.
func (p AddPasswordOptions) New() js.Ref {
	return bindings.AddPasswordOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AddPasswordOptions) UpdateFrom(ref js.Ref) {
	bindings.AddPasswordOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddPasswordOptions) Update(ref js.Ref) {
	bindings.AddPasswordOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddPasswordOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Username.Ref(),
		p.Password.Ref(),
		p.Note.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Username = p.Username.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
	p.Note = p.Note.FromRef(js.Undefined)
}

type CompromiseType uint32

const (
	_ CompromiseType = iota

	CompromiseType_LEAKED
	CompromiseType_PHISHED
	CompromiseType_REUSED
	CompromiseType_WEAK
)

func (CompromiseType) FromRef(str js.Ref) CompromiseType {
	return CompromiseType(bindings.ConstOfCompromiseType(str))
}

func (x CompromiseType) String() (string, bool) {
	switch x {
	case CompromiseType_LEAKED:
		return "LEAKED", true
	case CompromiseType_PHISHED:
		return "PHISHED", true
	case CompromiseType_REUSED:
		return "REUSED", true
	case CompromiseType_WEAK:
		return "WEAK", true
	default:
		return "", false
	}
}

type CompromisedInfo struct {
	// CompromiseTime is "CompromisedInfo.compromiseTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_CompromiseTime MUST be set to true to make this field effective.
	CompromiseTime float64
	// ElapsedTimeSinceCompromise is "CompromisedInfo.elapsedTimeSinceCompromise"
	//
	// Optional
	ElapsedTimeSinceCompromise js.String
	// CompromiseTypes is "CompromisedInfo.compromiseTypes"
	//
	// Optional
	CompromiseTypes js.Array[CompromiseType]
	// IsMuted is "CompromisedInfo.isMuted"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsMuted MUST be set to true to make this field effective.
	IsMuted bool

	FFI_USE_CompromiseTime bool // for CompromiseTime.
	FFI_USE_IsMuted        bool // for IsMuted.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CompromisedInfo with all fields set.
func (p CompromisedInfo) FromRef(ref js.Ref) CompromisedInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CompromisedInfo in the application heap.
func (p CompromisedInfo) New() js.Ref {
	return bindings.CompromisedInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CompromisedInfo) UpdateFrom(ref js.Ref) {
	bindings.CompromisedInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CompromisedInfo) Update(ref js.Ref) {
	bindings.CompromisedInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CompromisedInfo) FreeMembers(recursive bool) {
	js.Free(
		p.ElapsedTimeSinceCompromise.Ref(),
		p.CompromiseTypes.Ref(),
	)
	p.ElapsedTimeSinceCompromise = p.ElapsedTimeSinceCompromise.FromRef(js.Undefined)
	p.CompromiseTypes = p.CompromiseTypes.FromRef(js.Undefined)
}

type DomainInfo struct {
	// Name is "DomainInfo.name"
	//
	// Optional
	Name js.String
	// Url is "DomainInfo.url"
	//
	// Optional
	Url js.String
	// SignonRealm is "DomainInfo.signonRealm"
	//
	// Optional
	SignonRealm js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DomainInfo with all fields set.
func (p DomainInfo) FromRef(ref js.Ref) DomainInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DomainInfo in the application heap.
func (p DomainInfo) New() js.Ref {
	return bindings.DomainInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DomainInfo) UpdateFrom(ref js.Ref) {
	bindings.DomainInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DomainInfo) Update(ref js.Ref) {
	bindings.DomainInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DomainInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Url.Ref(),
		p.SignonRealm.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.SignonRealm = p.SignonRealm.FromRef(js.Undefined)
}

type PasswordStoreSet uint32

const (
	_ PasswordStoreSet = iota

	PasswordStoreSet_DEVICE
	PasswordStoreSet_ACCOUNT
	PasswordStoreSet_DEVICE_AND_ACCOUNT
)

func (PasswordStoreSet) FromRef(str js.Ref) PasswordStoreSet {
	return PasswordStoreSet(bindings.ConstOfPasswordStoreSet(str))
}

func (x PasswordStoreSet) String() (string, bool) {
	switch x {
	case PasswordStoreSet_DEVICE:
		return "DEVICE", true
	case PasswordStoreSet_ACCOUNT:
		return "ACCOUNT", true
	case PasswordStoreSet_DEVICE_AND_ACCOUNT:
		return "DEVICE_AND_ACCOUNT", true
	default:
		return "", false
	}
}

type PasswordUiEntry struct {
	// AffiliatedDomains is "PasswordUiEntry.affiliatedDomains"
	//
	// Optional
	AffiliatedDomains js.Array[DomainInfo]
	// Username is "PasswordUiEntry.username"
	//
	// Optional
	Username js.String
	// DisplayName is "PasswordUiEntry.displayName"
	//
	// Optional
	DisplayName js.String
	// Password is "PasswordUiEntry.password"
	//
	// Optional
	Password js.String
	// FederationText is "PasswordUiEntry.federationText"
	//
	// Optional
	FederationText js.String
	// Id is "PasswordUiEntry.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// StoredIn is "PasswordUiEntry.storedIn"
	//
	// Optional
	StoredIn PasswordStoreSet
	// IsPasskey is "PasswordUiEntry.isPasskey"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsPasskey MUST be set to true to make this field effective.
	IsPasskey bool
	// Note is "PasswordUiEntry.note"
	//
	// Optional
	Note js.String
	// ChangePasswordUrl is "PasswordUiEntry.changePasswordUrl"
	//
	// Optional
	ChangePasswordUrl js.String
	// CompromisedInfo is "PasswordUiEntry.compromisedInfo"
	//
	// Optional
	//
	// NOTE: CompromisedInfo.FFI_USE MUST be set to true to get CompromisedInfo used.
	CompromisedInfo CompromisedInfo

	FFI_USE_Id        bool // for Id.
	FFI_USE_IsPasskey bool // for IsPasskey.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PasswordUiEntry with all fields set.
func (p PasswordUiEntry) FromRef(ref js.Ref) PasswordUiEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PasswordUiEntry in the application heap.
func (p PasswordUiEntry) New() js.Ref {
	return bindings.PasswordUiEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PasswordUiEntry) UpdateFrom(ref js.Ref) {
	bindings.PasswordUiEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PasswordUiEntry) Update(ref js.Ref) {
	bindings.PasswordUiEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PasswordUiEntry) FreeMembers(recursive bool) {
	js.Free(
		p.AffiliatedDomains.Ref(),
		p.Username.Ref(),
		p.DisplayName.Ref(),
		p.Password.Ref(),
		p.FederationText.Ref(),
		p.Note.Ref(),
		p.ChangePasswordUrl.Ref(),
	)
	p.AffiliatedDomains = p.AffiliatedDomains.FromRef(js.Undefined)
	p.Username = p.Username.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
	p.FederationText = p.FederationText.FromRef(js.Undefined)
	p.Note = p.Note.FromRef(js.Undefined)
	p.ChangePasswordUrl = p.ChangePasswordUrl.FromRef(js.Undefined)
	if recursive {
		p.CompromisedInfo.FreeMembers(true)
	}
}

type CredentialGroup struct {
	// Name is "CredentialGroup.name"
	//
	// Optional
	Name js.String
	// IconUrl is "CredentialGroup.iconUrl"
	//
	// Optional
	IconUrl js.String
	// Entries is "CredentialGroup.entries"
	//
	// Optional
	Entries js.Array[PasswordUiEntry]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CredentialGroup with all fields set.
func (p CredentialGroup) FromRef(ref js.Ref) CredentialGroup {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CredentialGroup in the application heap.
func (p CredentialGroup) New() js.Ref {
	return bindings.CredentialGroupJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CredentialGroup) UpdateFrom(ref js.Ref) {
	bindings.CredentialGroupJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CredentialGroup) Update(ref js.Ref) {
	bindings.CredentialGroupJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CredentialGroup) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.IconUrl.Ref(),
		p.Entries.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.IconUrl = p.IconUrl.FromRef(js.Undefined)
	p.Entries = p.Entries.FromRef(js.Undefined)
}

type CredentialsGroupCallbackFunc func(this js.Ref, entries js.Array[CredentialGroup]) js.Ref

func (fn CredentialsGroupCallbackFunc) Register() js.Func[func(entries js.Array[CredentialGroup])] {
	return js.RegisterCallback[func(entries js.Array[CredentialGroup])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CredentialsGroupCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[CredentialGroup]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CredentialsGroupCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[CredentialGroup]) js.Ref
	Arg T
}

func (cb *CredentialsGroupCallback[T]) Register() js.Func[func(entries js.Array[CredentialGroup])] {
	return js.RegisterCallback[func(entries js.Array[CredentialGroup])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CredentialsGroupCallback[T]) DispatchCallback(
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

		js.Array[CredentialGroup]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CredentialsWithReusedPasswordCallbackFunc func(this js.Ref, entries js.Array[PasswordUiEntryList]) js.Ref

func (fn CredentialsWithReusedPasswordCallbackFunc) Register() js.Func[func(entries js.Array[PasswordUiEntryList])] {
	return js.RegisterCallback[func(entries js.Array[PasswordUiEntryList])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CredentialsWithReusedPasswordCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[PasswordUiEntryList]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CredentialsWithReusedPasswordCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[PasswordUiEntryList]) js.Ref
	Arg T
}

func (cb *CredentialsWithReusedPasswordCallback[T]) Register() js.Func[func(entries js.Array[PasswordUiEntryList])] {
	return js.RegisterCallback[func(entries js.Array[PasswordUiEntryList])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CredentialsWithReusedPasswordCallback[T]) DispatchCallback(
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

		js.Array[PasswordUiEntryList]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PasswordUiEntryList struct {
	// Entries is "PasswordUiEntryList.entries"
	//
	// Optional
	Entries js.Array[PasswordUiEntry]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PasswordUiEntryList with all fields set.
func (p PasswordUiEntryList) FromRef(ref js.Ref) PasswordUiEntryList {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PasswordUiEntryList in the application heap.
func (p PasswordUiEntryList) New() js.Ref {
	return bindings.PasswordUiEntryListJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PasswordUiEntryList) UpdateFrom(ref js.Ref) {
	bindings.PasswordUiEntryListJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PasswordUiEntryList) Update(ref js.Ref) {
	bindings.PasswordUiEntryListJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PasswordUiEntryList) FreeMembers(recursive bool) {
	js.Free(
		p.Entries.Ref(),
	)
	p.Entries = p.Entries.FromRef(js.Undefined)
}

type UrlCollection struct {
	// SignonRealm is "UrlCollection.signonRealm"
	//
	// Optional
	SignonRealm js.String
	// Shown is "UrlCollection.shown"
	//
	// Optional
	Shown js.String
	// Link is "UrlCollection.link"
	//
	// Optional
	Link js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UrlCollection with all fields set.
func (p UrlCollection) FromRef(ref js.Ref) UrlCollection {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UrlCollection in the application heap.
func (p UrlCollection) New() js.Ref {
	return bindings.UrlCollectionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UrlCollection) UpdateFrom(ref js.Ref) {
	bindings.UrlCollectionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UrlCollection) Update(ref js.Ref) {
	bindings.UrlCollectionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UrlCollection) FreeMembers(recursive bool) {
	js.Free(
		p.SignonRealm.Ref(),
		p.Shown.Ref(),
		p.Link.Ref(),
	)
	p.SignonRealm = p.SignonRealm.FromRef(js.Undefined)
	p.Shown = p.Shown.FromRef(js.Undefined)
	p.Link = p.Link.FromRef(js.Undefined)
}

type ExceptionEntry struct {
	// Urls is "ExceptionEntry.urls"
	//
	// Optional
	//
	// NOTE: Urls.FFI_USE MUST be set to true to get Urls used.
	Urls UrlCollection
	// Id is "ExceptionEntry.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32

	FFI_USE_Id bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExceptionEntry with all fields set.
func (p ExceptionEntry) FromRef(ref js.Ref) ExceptionEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExceptionEntry in the application heap.
func (p ExceptionEntry) New() js.Ref {
	return bindings.ExceptionEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExceptionEntry) UpdateFrom(ref js.Ref) {
	bindings.ExceptionEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExceptionEntry) Update(ref js.Ref) {
	bindings.ExceptionEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExceptionEntry) FreeMembers(recursive bool) {
	if recursive {
		p.Urls.FreeMembers(true)
	}
}

type ExceptionListCallbackFunc func(this js.Ref, exceptions js.Array[ExceptionEntry]) js.Ref

func (fn ExceptionListCallbackFunc) Register() js.Func[func(exceptions js.Array[ExceptionEntry])] {
	return js.RegisterCallback[func(exceptions js.Array[ExceptionEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ExceptionListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ExceptionEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExceptionListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, exceptions js.Array[ExceptionEntry]) js.Ref
	Arg T
}

func (cb *ExceptionListCallback[T]) Register() js.Func[func(exceptions js.Array[ExceptionEntry])] {
	return js.RegisterCallback[func(exceptions js.Array[ExceptionEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ExceptionListCallback[T]) DispatchCallback(
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

		js.Array[ExceptionEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExportProgressStatus uint32

const (
	_ ExportProgressStatus = iota

	ExportProgressStatus_NOT_STARTED
	ExportProgressStatus_IN_PROGRESS
	ExportProgressStatus_SUCCEEDED
	ExportProgressStatus_FAILED_CANCELLED
	ExportProgressStatus_FAILED_WRITE_FAILED
)

func (ExportProgressStatus) FromRef(str js.Ref) ExportProgressStatus {
	return ExportProgressStatus(bindings.ConstOfExportProgressStatus(str))
}

func (x ExportProgressStatus) String() (string, bool) {
	switch x {
	case ExportProgressStatus_NOT_STARTED:
		return "NOT_STARTED", true
	case ExportProgressStatus_IN_PROGRESS:
		return "IN_PROGRESS", true
	case ExportProgressStatus_SUCCEEDED:
		return "SUCCEEDED", true
	case ExportProgressStatus_FAILED_CANCELLED:
		return "FAILED_CANCELLED", true
	case ExportProgressStatus_FAILED_WRITE_FAILED:
		return "FAILED_WRITE_FAILED", true
	default:
		return "", false
	}
}

type ExportProgressStatusCallbackFunc func(this js.Ref, status ExportProgressStatus) js.Ref

func (fn ExportProgressStatusCallbackFunc) Register() js.Func[func(status ExportProgressStatus)] {
	return js.RegisterCallback[func(status ExportProgressStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ExportProgressStatusCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		ExportProgressStatus(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExportProgressStatusCallback[T any] struct {
	Fn  func(arg T, this js.Ref, status ExportProgressStatus) js.Ref
	Arg T
}

func (cb *ExportProgressStatusCallback[T]) Register() js.Func[func(status ExportProgressStatus)] {
	return js.RegisterCallback[func(status ExportProgressStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ExportProgressStatusCallback[T]) DispatchCallback(
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

		ExportProgressStatus(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FamilyFetchStatus uint32

const (
	_ FamilyFetchStatus = iota

	FamilyFetchStatus_UNKNOWN_ERROR
	FamilyFetchStatus_NO_MEMBERS
	FamilyFetchStatus_SUCCESS
)

func (FamilyFetchStatus) FromRef(str js.Ref) FamilyFetchStatus {
	return FamilyFetchStatus(bindings.ConstOfFamilyFetchStatus(str))
}

func (x FamilyFetchStatus) String() (string, bool) {
	switch x {
	case FamilyFetchStatus_UNKNOWN_ERROR:
		return "UNKNOWN_ERROR", true
	case FamilyFetchStatus_NO_MEMBERS:
		return "NO_MEMBERS", true
	case FamilyFetchStatus_SUCCESS:
		return "SUCCESS", true
	default:
		return "", false
	}
}

type PublicKey struct {
	// Value is "PublicKey.value"
	//
	// Optional
	Value js.String
	// Version is "PublicKey.version"
	//
	// Optional
	//
	// NOTE: FFI_USE_Version MUST be set to true to make this field effective.
	Version int32

	FFI_USE_Version bool // for Version.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKey with all fields set.
func (p PublicKey) FromRef(ref js.Ref) PublicKey {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PublicKey in the application heap.
func (p PublicKey) New() js.Ref {
	return bindings.PublicKeyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKey) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKey) Update(ref js.Ref) {
	bindings.PublicKeyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKey) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

type RecipientInfo struct {
	// UserId is "RecipientInfo.userId"
	//
	// Optional
	UserId js.String
	// Email is "RecipientInfo.email"
	//
	// Optional
	Email js.String
	// DisplayName is "RecipientInfo.displayName"
	//
	// Optional
	DisplayName js.String
	// ProfileImageUrl is "RecipientInfo.profileImageUrl"
	//
	// Optional
	ProfileImageUrl js.String
	// IsEligible is "RecipientInfo.isEligible"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsEligible MUST be set to true to make this field effective.
	IsEligible bool
	// PublicKey is "RecipientInfo.publicKey"
	//
	// Optional
	//
	// NOTE: PublicKey.FFI_USE MUST be set to true to get PublicKey used.
	PublicKey PublicKey

	FFI_USE_IsEligible bool // for IsEligible.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RecipientInfo with all fields set.
func (p RecipientInfo) FromRef(ref js.Ref) RecipientInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RecipientInfo in the application heap.
func (p RecipientInfo) New() js.Ref {
	return bindings.RecipientInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RecipientInfo) UpdateFrom(ref js.Ref) {
	bindings.RecipientInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RecipientInfo) Update(ref js.Ref) {
	bindings.RecipientInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RecipientInfo) FreeMembers(recursive bool) {
	js.Free(
		p.UserId.Ref(),
		p.Email.Ref(),
		p.DisplayName.Ref(),
		p.ProfileImageUrl.Ref(),
	)
	p.UserId = p.UserId.FromRef(js.Undefined)
	p.Email = p.Email.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.ProfileImageUrl = p.ProfileImageUrl.FromRef(js.Undefined)
	if recursive {
		p.PublicKey.FreeMembers(true)
	}
}

type FamilyFetchResults struct {
	// Status is "FamilyFetchResults.status"
	//
	// Optional
	Status FamilyFetchStatus
	// FamilyMembers is "FamilyFetchResults.familyMembers"
	//
	// Optional
	FamilyMembers js.Array[RecipientInfo]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FamilyFetchResults with all fields set.
func (p FamilyFetchResults) FromRef(ref js.Ref) FamilyFetchResults {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FamilyFetchResults in the application heap.
func (p FamilyFetchResults) New() js.Ref {
	return bindings.FamilyFetchResultsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FamilyFetchResults) UpdateFrom(ref js.Ref) {
	bindings.FamilyFetchResultsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FamilyFetchResults) Update(ref js.Ref) {
	bindings.FamilyFetchResultsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FamilyFetchResults) FreeMembers(recursive bool) {
	js.Free(
		p.FamilyMembers.Ref(),
	)
	p.FamilyMembers = p.FamilyMembers.FromRef(js.Undefined)
}

type FetchFamilyResultsCallbackFunc func(this js.Ref, results *FamilyFetchResults) js.Ref

func (fn FetchFamilyResultsCallbackFunc) Register() js.Func[func(results *FamilyFetchResults)] {
	return js.RegisterCallback[func(results *FamilyFetchResults)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FetchFamilyResultsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FamilyFetchResults
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FetchFamilyResultsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, results *FamilyFetchResults) js.Ref
	Arg T
}

func (cb *FetchFamilyResultsCallback[T]) Register() js.Func[func(results *FamilyFetchResults)] {
	return js.RegisterCallback[func(results *FamilyFetchResults)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FetchFamilyResultsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FamilyFetchResults
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetUrlCollectionCallbackFunc func(this js.Ref, urlCollection *UrlCollection) js.Ref

func (fn GetUrlCollectionCallbackFunc) Register() js.Func[func(urlCollection *UrlCollection)] {
	return js.RegisterCallback[func(urlCollection *UrlCollection)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetUrlCollectionCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UrlCollection
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetUrlCollectionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, urlCollection *UrlCollection) js.Ref
	Arg T
}

func (cb *GetUrlCollectionCallback[T]) Register() js.Func[func(urlCollection *UrlCollection)] {
	return js.RegisterCallback[func(urlCollection *UrlCollection)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetUrlCollectionCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 UrlCollection
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ImportEntryStatus uint32

const (
	_ ImportEntryStatus = iota

	ImportEntryStatus_UNKNOWN_ERROR
	ImportEntryStatus_MISSING_PASSWORD
	ImportEntryStatus_MISSING_URL
	ImportEntryStatus_INVALID_URL
	ImportEntryStatus_NON_ASCII_URL
	ImportEntryStatus_LONG_URL
	ImportEntryStatus_LONG_PASSWORD
	ImportEntryStatus_LONG_USERNAME
	ImportEntryStatus_CONFLICT_PROFILE
	ImportEntryStatus_CONFLICT_ACCOUNT
	ImportEntryStatus_LONG_NOTE
	ImportEntryStatus_LONG_CONCATENATED_NOTE
	ImportEntryStatus_VALID
)

func (ImportEntryStatus) FromRef(str js.Ref) ImportEntryStatus {
	return ImportEntryStatus(bindings.ConstOfImportEntryStatus(str))
}

func (x ImportEntryStatus) String() (string, bool) {
	switch x {
	case ImportEntryStatus_UNKNOWN_ERROR:
		return "UNKNOWN_ERROR", true
	case ImportEntryStatus_MISSING_PASSWORD:
		return "MISSING_PASSWORD", true
	case ImportEntryStatus_MISSING_URL:
		return "MISSING_URL", true
	case ImportEntryStatus_INVALID_URL:
		return "INVALID_URL", true
	case ImportEntryStatus_NON_ASCII_URL:
		return "NON_ASCII_URL", true
	case ImportEntryStatus_LONG_URL:
		return "LONG_URL", true
	case ImportEntryStatus_LONG_PASSWORD:
		return "LONG_PASSWORD", true
	case ImportEntryStatus_LONG_USERNAME:
		return "LONG_USERNAME", true
	case ImportEntryStatus_CONFLICT_PROFILE:
		return "CONFLICT_PROFILE", true
	case ImportEntryStatus_CONFLICT_ACCOUNT:
		return "CONFLICT_ACCOUNT", true
	case ImportEntryStatus_LONG_NOTE:
		return "LONG_NOTE", true
	case ImportEntryStatus_LONG_CONCATENATED_NOTE:
		return "LONG_CONCATENATED_NOTE", true
	case ImportEntryStatus_VALID:
		return "VALID", true
	default:
		return "", false
	}
}

type ImportEntry struct {
	// Status is "ImportEntry.status"
	//
	// Optional
	Status ImportEntryStatus
	// Url is "ImportEntry.url"
	//
	// Optional
	Url js.String
	// Username is "ImportEntry.username"
	//
	// Optional
	Username js.String
	// Password is "ImportEntry.password"
	//
	// Optional
	Password js.String
	// Id is "ImportEntry.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32

	FFI_USE_Id bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImportEntry with all fields set.
func (p ImportEntry) FromRef(ref js.Ref) ImportEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImportEntry in the application heap.
func (p ImportEntry) New() js.Ref {
	return bindings.ImportEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ImportEntry) UpdateFrom(ref js.Ref) {
	bindings.ImportEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImportEntry) Update(ref js.Ref) {
	bindings.ImportEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImportEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Username.Ref(),
		p.Password.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Username = p.Username.FromRef(js.Undefined)
	p.Password = p.Password.FromRef(js.Undefined)
}

type ImportPasswordsCallbackFunc func(this js.Ref, results *ImportResults) js.Ref

func (fn ImportPasswordsCallbackFunc) Register() js.Func[func(results *ImportResults)] {
	return js.RegisterCallback[func(results *ImportResults)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ImportPasswordsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ImportResults
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ImportPasswordsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, results *ImportResults) js.Ref
	Arg T
}

func (cb *ImportPasswordsCallback[T]) Register() js.Func[func(results *ImportResults)] {
	return js.RegisterCallback[func(results *ImportResults)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ImportPasswordsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ImportResults
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ImportResultsStatus uint32

const (
	_ ImportResultsStatus = iota

	ImportResultsStatus_UNKNOWN_ERROR
	ImportResultsStatus_SUCCESS
	ImportResultsStatus_IO_ERROR
	ImportResultsStatus_BAD_FORMAT
	ImportResultsStatus_DISMISSED
	ImportResultsStatus_MAX_FILE_SIZE
	ImportResultsStatus_IMPORT_ALREADY_ACTIVE
	ImportResultsStatus_NUM_PASSWORDS_EXCEEDED
	ImportResultsStatus_CONFLICTS
)

func (ImportResultsStatus) FromRef(str js.Ref) ImportResultsStatus {
	return ImportResultsStatus(bindings.ConstOfImportResultsStatus(str))
}

func (x ImportResultsStatus) String() (string, bool) {
	switch x {
	case ImportResultsStatus_UNKNOWN_ERROR:
		return "UNKNOWN_ERROR", true
	case ImportResultsStatus_SUCCESS:
		return "SUCCESS", true
	case ImportResultsStatus_IO_ERROR:
		return "IO_ERROR", true
	case ImportResultsStatus_BAD_FORMAT:
		return "BAD_FORMAT", true
	case ImportResultsStatus_DISMISSED:
		return "DISMISSED", true
	case ImportResultsStatus_MAX_FILE_SIZE:
		return "MAX_FILE_SIZE", true
	case ImportResultsStatus_IMPORT_ALREADY_ACTIVE:
		return "IMPORT_ALREADY_ACTIVE", true
	case ImportResultsStatus_NUM_PASSWORDS_EXCEEDED:
		return "NUM_PASSWORDS_EXCEEDED", true
	case ImportResultsStatus_CONFLICTS:
		return "CONFLICTS", true
	default:
		return "", false
	}
}

type ImportResults struct {
	// Status is "ImportResults.status"
	//
	// Optional
	Status ImportResultsStatus
	// NumberImported is "ImportResults.numberImported"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumberImported MUST be set to true to make this field effective.
	NumberImported int32
	// DisplayedEntries is "ImportResults.displayedEntries"
	//
	// Optional
	DisplayedEntries js.Array[ImportEntry]
	// FileName is "ImportResults.fileName"
	//
	// Optional
	FileName js.String

	FFI_USE_NumberImported bool // for NumberImported.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ImportResults with all fields set.
func (p ImportResults) FromRef(ref js.Ref) ImportResults {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ImportResults in the application heap.
func (p ImportResults) New() js.Ref {
	return bindings.ImportResultsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ImportResults) UpdateFrom(ref js.Ref) {
	bindings.ImportResultsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ImportResults) Update(ref js.Ref) {
	bindings.ImportResultsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ImportResults) FreeMembers(recursive bool) {
	js.Free(
		p.DisplayedEntries.Ref(),
		p.FileName.Ref(),
	)
	p.DisplayedEntries = p.DisplayedEntries.FromRef(js.Undefined)
	p.FileName = p.FileName.FromRef(js.Undefined)
}

type IsAccountStoreDefaultCallbackFunc func(this js.Ref, isDefault bool) js.Ref

func (fn IsAccountStoreDefaultCallbackFunc) Register() js.Func[func(isDefault bool)] {
	return js.RegisterCallback[func(isDefault bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsAccountStoreDefaultCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IsAccountStoreDefaultCallback[T any] struct {
	Fn  func(arg T, this js.Ref, isDefault bool) js.Ref
	Arg T
}

func (cb *IsAccountStoreDefaultCallback[T]) Register() js.Func[func(isDefault bool)] {
	return js.RegisterCallback[func(isDefault bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsAccountStoreDefaultCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OptInCallbackFunc func(this js.Ref, optedIn bool) js.Ref

func (fn OptInCallbackFunc) Register() js.Func[func(optedIn bool)] {
	return js.RegisterCallback[func(optedIn bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OptInCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OptInCallback[T any] struct {
	Fn  func(arg T, this js.Ref, optedIn bool) js.Ref
	Arg T
}

func (cb *OptInCallback[T]) Register() js.Func[func(optedIn bool)] {
	return js.RegisterCallback[func(optedIn bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OptInCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PasswordCheckState uint32

const (
	_ PasswordCheckState = iota

	PasswordCheckState_IDLE
	PasswordCheckState_RUNNING
	PasswordCheckState_CANCELED
	PasswordCheckState_OFFLINE
	PasswordCheckState_SIGNED_OUT
	PasswordCheckState_NO_PASSWORDS
	PasswordCheckState_QUOTA_LIMIT
	PasswordCheckState_OTHER_ERROR
)

func (PasswordCheckState) FromRef(str js.Ref) PasswordCheckState {
	return PasswordCheckState(bindings.ConstOfPasswordCheckState(str))
}

func (x PasswordCheckState) String() (string, bool) {
	switch x {
	case PasswordCheckState_IDLE:
		return "IDLE", true
	case PasswordCheckState_RUNNING:
		return "RUNNING", true
	case PasswordCheckState_CANCELED:
		return "CANCELED", true
	case PasswordCheckState_OFFLINE:
		return "OFFLINE", true
	case PasswordCheckState_SIGNED_OUT:
		return "SIGNED_OUT", true
	case PasswordCheckState_NO_PASSWORDS:
		return "NO_PASSWORDS", true
	case PasswordCheckState_QUOTA_LIMIT:
		return "QUOTA_LIMIT", true
	case PasswordCheckState_OTHER_ERROR:
		return "OTHER_ERROR", true
	default:
		return "", false
	}
}

type PasswordCheckStatus struct {
	// State is "PasswordCheckStatus.state"
	//
	// Optional
	State PasswordCheckState
	// TotalNumberOfPasswords is "PasswordCheckStatus.totalNumberOfPasswords"
	//
	// Optional
	//
	// NOTE: FFI_USE_TotalNumberOfPasswords MUST be set to true to make this field effective.
	TotalNumberOfPasswords int32
	// AlreadyProcessed is "PasswordCheckStatus.alreadyProcessed"
	//
	// Optional
	//
	// NOTE: FFI_USE_AlreadyProcessed MUST be set to true to make this field effective.
	AlreadyProcessed int32
	// RemainingInQueue is "PasswordCheckStatus.remainingInQueue"
	//
	// Optional
	//
	// NOTE: FFI_USE_RemainingInQueue MUST be set to true to make this field effective.
	RemainingInQueue int32
	// ElapsedTimeSinceLastCheck is "PasswordCheckStatus.elapsedTimeSinceLastCheck"
	//
	// Optional
	ElapsedTimeSinceLastCheck js.String

	FFI_USE_TotalNumberOfPasswords bool // for TotalNumberOfPasswords.
	FFI_USE_AlreadyProcessed       bool // for AlreadyProcessed.
	FFI_USE_RemainingInQueue       bool // for RemainingInQueue.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PasswordCheckStatus with all fields set.
func (p PasswordCheckStatus) FromRef(ref js.Ref) PasswordCheckStatus {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PasswordCheckStatus in the application heap.
func (p PasswordCheckStatus) New() js.Ref {
	return bindings.PasswordCheckStatusJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PasswordCheckStatus) UpdateFrom(ref js.Ref) {
	bindings.PasswordCheckStatusJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PasswordCheckStatus) Update(ref js.Ref) {
	bindings.PasswordCheckStatusJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PasswordCheckStatus) FreeMembers(recursive bool) {
	js.Free(
		p.ElapsedTimeSinceLastCheck.Ref(),
	)
	p.ElapsedTimeSinceLastCheck = p.ElapsedTimeSinceLastCheck.FromRef(js.Undefined)
}

type PasswordCheckStatusCallbackFunc func(this js.Ref, status *PasswordCheckStatus) js.Ref

func (fn PasswordCheckStatusCallbackFunc) Register() js.Func[func(status *PasswordCheckStatus)] {
	return js.RegisterCallback[func(status *PasswordCheckStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PasswordCheckStatusCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PasswordCheckStatus
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PasswordCheckStatusCallback[T any] struct {
	Fn  func(arg T, this js.Ref, status *PasswordCheckStatus) js.Ref
	Arg T
}

func (cb *PasswordCheckStatusCallback[T]) Register() js.Func[func(status *PasswordCheckStatus)] {
	return js.RegisterCallback[func(status *PasswordCheckStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PasswordCheckStatusCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PasswordCheckStatus
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PasswordExportProgress struct {
	// Status is "PasswordExportProgress.status"
	//
	// Optional
	Status ExportProgressStatus
	// FilePath is "PasswordExportProgress.filePath"
	//
	// Optional
	FilePath js.String
	// FolderName is "PasswordExportProgress.folderName"
	//
	// Optional
	FolderName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PasswordExportProgress with all fields set.
func (p PasswordExportProgress) FromRef(ref js.Ref) PasswordExportProgress {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PasswordExportProgress in the application heap.
func (p PasswordExportProgress) New() js.Ref {
	return bindings.PasswordExportProgressJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PasswordExportProgress) UpdateFrom(ref js.Ref) {
	bindings.PasswordExportProgressJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PasswordExportProgress) Update(ref js.Ref) {
	bindings.PasswordExportProgressJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PasswordExportProgress) FreeMembers(recursive bool) {
	js.Free(
		p.FilePath.Ref(),
		p.FolderName.Ref(),
	)
	p.FilePath = p.FilePath.FromRef(js.Undefined)
	p.FolderName = p.FolderName.FromRef(js.Undefined)
}

type PasswordListCallbackFunc func(this js.Ref, entries js.Array[PasswordUiEntry]) js.Ref

func (fn PasswordListCallbackFunc) Register() js.Func[func(entries js.Array[PasswordUiEntry])] {
	return js.RegisterCallback[func(entries js.Array[PasswordUiEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PasswordListCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[PasswordUiEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PasswordListCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[PasswordUiEntry]) js.Ref
	Arg T
}

func (cb *PasswordListCallback[T]) Register() js.Func[func(entries js.Array[PasswordUiEntry])] {
	return js.RegisterCallback[func(entries js.Array[PasswordUiEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PasswordListCallback[T]) DispatchCallback(
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

		js.Array[PasswordUiEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PlaintextPasswordCallbackFunc func(this js.Ref, password js.String) js.Ref

func (fn PlaintextPasswordCallbackFunc) Register() js.Func[func(password js.String)] {
	return js.RegisterCallback[func(password js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PlaintextPasswordCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PlaintextPasswordCallback[T any] struct {
	Fn  func(arg T, this js.Ref, password js.String) js.Ref
	Arg T
}

func (cb *PlaintextPasswordCallback[T]) Register() js.Func[func(password js.String)] {
	return js.RegisterCallback[func(password js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PlaintextPasswordCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PlaintextReason uint32

const (
	_ PlaintextReason = iota

	PlaintextReason_VIEW
	PlaintextReason_COPY
	PlaintextReason_EDIT
)

func (PlaintextReason) FromRef(str js.Ref) PlaintextReason {
	return PlaintextReason(bindings.ConstOfPlaintextReason(str))
}

func (x PlaintextReason) String() (string, bool) {
	switch x {
	case PlaintextReason_VIEW:
		return "VIEW", true
	case PlaintextReason_COPY:
		return "COPY", true
	case PlaintextReason_EDIT:
		return "EDIT", true
	default:
		return "", false
	}
}

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
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

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
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

// HasFuncAddPassword returns true if the function "WEBEXT.passwordsPrivate.addPassword" exists.
func HasFuncAddPassword() bool {
	return js.True == bindings.HasFuncAddPassword()
}

// FuncAddPassword returns the function "WEBEXT.passwordsPrivate.addPassword".
func FuncAddPassword() (fn js.Func[func(options AddPasswordOptions) js.Promise[js.Void]]) {
	bindings.FuncAddPassword(
		js.Pointer(&fn),
	)
	return
}

// AddPassword calls the function "WEBEXT.passwordsPrivate.addPassword" directly.
func AddPassword(options AddPasswordOptions) (ret js.Promise[js.Void]) {
	bindings.CallAddPassword(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryAddPassword calls the function "WEBEXT.passwordsPrivate.addPassword"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddPassword(options AddPasswordOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddPassword(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncChangeCredential returns true if the function "WEBEXT.passwordsPrivate.changeCredential" exists.
func HasFuncChangeCredential() bool {
	return js.True == bindings.HasFuncChangeCredential()
}

// FuncChangeCredential returns the function "WEBEXT.passwordsPrivate.changeCredential".
func FuncChangeCredential() (fn js.Func[func(credential PasswordUiEntry) js.Promise[js.Void]]) {
	bindings.FuncChangeCredential(
		js.Pointer(&fn),
	)
	return
}

// ChangeCredential calls the function "WEBEXT.passwordsPrivate.changeCredential" directly.
func ChangeCredential(credential PasswordUiEntry) (ret js.Promise[js.Void]) {
	bindings.CallChangeCredential(
		js.Pointer(&ret),
		js.Pointer(&credential),
	)

	return
}

// TryChangeCredential calls the function "WEBEXT.passwordsPrivate.changeCredential"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChangeCredential(credential PasswordUiEntry) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryChangeCredential(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&credential),
	)

	return
}

// HasFuncContinueImport returns true if the function "WEBEXT.passwordsPrivate.continueImport" exists.
func HasFuncContinueImport() bool {
	return js.True == bindings.HasFuncContinueImport()
}

// FuncContinueImport returns the function "WEBEXT.passwordsPrivate.continueImport".
func FuncContinueImport() (fn js.Func[func(selectedIds js.Array[int32]) js.Promise[ImportResults]]) {
	bindings.FuncContinueImport(
		js.Pointer(&fn),
	)
	return
}

// ContinueImport calls the function "WEBEXT.passwordsPrivate.continueImport" directly.
func ContinueImport(selectedIds js.Array[int32]) (ret js.Promise[ImportResults]) {
	bindings.CallContinueImport(
		js.Pointer(&ret),
		selectedIds.Ref(),
	)

	return
}

// TryContinueImport calls the function "WEBEXT.passwordsPrivate.continueImport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryContinueImport(selectedIds js.Array[int32]) (ret js.Promise[ImportResults], exception js.Any, ok bool) {
	ok = js.True == bindings.TryContinueImport(
		js.Pointer(&ret), js.Pointer(&exception),
		selectedIds.Ref(),
	)

	return
}

// HasFuncExportPasswords returns true if the function "WEBEXT.passwordsPrivate.exportPasswords" exists.
func HasFuncExportPasswords() bool {
	return js.True == bindings.HasFuncExportPasswords()
}

// FuncExportPasswords returns the function "WEBEXT.passwordsPrivate.exportPasswords".
func FuncExportPasswords() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncExportPasswords(
		js.Pointer(&fn),
	)
	return
}

// ExportPasswords calls the function "WEBEXT.passwordsPrivate.exportPasswords" directly.
func ExportPasswords() (ret js.Promise[js.Void]) {
	bindings.CallExportPasswords(
		js.Pointer(&ret),
	)

	return
}

// TryExportPasswords calls the function "WEBEXT.passwordsPrivate.exportPasswords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExportPasswords() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExportPasswords(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExtendAuthValidity returns true if the function "WEBEXT.passwordsPrivate.extendAuthValidity" exists.
func HasFuncExtendAuthValidity() bool {
	return js.True == bindings.HasFuncExtendAuthValidity()
}

// FuncExtendAuthValidity returns the function "WEBEXT.passwordsPrivate.extendAuthValidity".
func FuncExtendAuthValidity() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncExtendAuthValidity(
		js.Pointer(&fn),
	)
	return
}

// ExtendAuthValidity calls the function "WEBEXT.passwordsPrivate.extendAuthValidity" directly.
func ExtendAuthValidity() (ret js.Promise[js.Void]) {
	bindings.CallExtendAuthValidity(
		js.Pointer(&ret),
	)

	return
}

// TryExtendAuthValidity calls the function "WEBEXT.passwordsPrivate.extendAuthValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExtendAuthValidity() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExtendAuthValidity(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFetchFamilyMembers returns true if the function "WEBEXT.passwordsPrivate.fetchFamilyMembers" exists.
func HasFuncFetchFamilyMembers() bool {
	return js.True == bindings.HasFuncFetchFamilyMembers()
}

// FuncFetchFamilyMembers returns the function "WEBEXT.passwordsPrivate.fetchFamilyMembers".
func FuncFetchFamilyMembers() (fn js.Func[func() js.Promise[FamilyFetchResults]]) {
	bindings.FuncFetchFamilyMembers(
		js.Pointer(&fn),
	)
	return
}

// FetchFamilyMembers calls the function "WEBEXT.passwordsPrivate.fetchFamilyMembers" directly.
func FetchFamilyMembers() (ret js.Promise[FamilyFetchResults]) {
	bindings.CallFetchFamilyMembers(
		js.Pointer(&ret),
	)

	return
}

// TryFetchFamilyMembers calls the function "WEBEXT.passwordsPrivate.fetchFamilyMembers"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFetchFamilyMembers() (ret js.Promise[FamilyFetchResults], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFetchFamilyMembers(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCredentialGroups returns true if the function "WEBEXT.passwordsPrivate.getCredentialGroups" exists.
func HasFuncGetCredentialGroups() bool {
	return js.True == bindings.HasFuncGetCredentialGroups()
}

// FuncGetCredentialGroups returns the function "WEBEXT.passwordsPrivate.getCredentialGroups".
func FuncGetCredentialGroups() (fn js.Func[func() js.Promise[js.Array[CredentialGroup]]]) {
	bindings.FuncGetCredentialGroups(
		js.Pointer(&fn),
	)
	return
}

// GetCredentialGroups calls the function "WEBEXT.passwordsPrivate.getCredentialGroups" directly.
func GetCredentialGroups() (ret js.Promise[js.Array[CredentialGroup]]) {
	bindings.CallGetCredentialGroups(
		js.Pointer(&ret),
	)

	return
}

// TryGetCredentialGroups calls the function "WEBEXT.passwordsPrivate.getCredentialGroups"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCredentialGroups() (ret js.Promise[js.Array[CredentialGroup]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCredentialGroups(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCredentialsWithReusedPassword returns true if the function "WEBEXT.passwordsPrivate.getCredentialsWithReusedPassword" exists.
func HasFuncGetCredentialsWithReusedPassword() bool {
	return js.True == bindings.HasFuncGetCredentialsWithReusedPassword()
}

// FuncGetCredentialsWithReusedPassword returns the function "WEBEXT.passwordsPrivate.getCredentialsWithReusedPassword".
func FuncGetCredentialsWithReusedPassword() (fn js.Func[func() js.Promise[js.Array[PasswordUiEntryList]]]) {
	bindings.FuncGetCredentialsWithReusedPassword(
		js.Pointer(&fn),
	)
	return
}

// GetCredentialsWithReusedPassword calls the function "WEBEXT.passwordsPrivate.getCredentialsWithReusedPassword" directly.
func GetCredentialsWithReusedPassword() (ret js.Promise[js.Array[PasswordUiEntryList]]) {
	bindings.CallGetCredentialsWithReusedPassword(
		js.Pointer(&ret),
	)

	return
}

// TryGetCredentialsWithReusedPassword calls the function "WEBEXT.passwordsPrivate.getCredentialsWithReusedPassword"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCredentialsWithReusedPassword() (ret js.Promise[js.Array[PasswordUiEntryList]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCredentialsWithReusedPassword(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetInsecureCredentials returns true if the function "WEBEXT.passwordsPrivate.getInsecureCredentials" exists.
func HasFuncGetInsecureCredentials() bool {
	return js.True == bindings.HasFuncGetInsecureCredentials()
}

// FuncGetInsecureCredentials returns the function "WEBEXT.passwordsPrivate.getInsecureCredentials".
func FuncGetInsecureCredentials() (fn js.Func[func() js.Promise[js.Array[PasswordUiEntry]]]) {
	bindings.FuncGetInsecureCredentials(
		js.Pointer(&fn),
	)
	return
}

// GetInsecureCredentials calls the function "WEBEXT.passwordsPrivate.getInsecureCredentials" directly.
func GetInsecureCredentials() (ret js.Promise[js.Array[PasswordUiEntry]]) {
	bindings.CallGetInsecureCredentials(
		js.Pointer(&ret),
	)

	return
}

// TryGetInsecureCredentials calls the function "WEBEXT.passwordsPrivate.getInsecureCredentials"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetInsecureCredentials() (ret js.Promise[js.Array[PasswordUiEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetInsecureCredentials(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPasswordCheckStatus returns true if the function "WEBEXT.passwordsPrivate.getPasswordCheckStatus" exists.
func HasFuncGetPasswordCheckStatus() bool {
	return js.True == bindings.HasFuncGetPasswordCheckStatus()
}

// FuncGetPasswordCheckStatus returns the function "WEBEXT.passwordsPrivate.getPasswordCheckStatus".
func FuncGetPasswordCheckStatus() (fn js.Func[func() js.Promise[PasswordCheckStatus]]) {
	bindings.FuncGetPasswordCheckStatus(
		js.Pointer(&fn),
	)
	return
}

// GetPasswordCheckStatus calls the function "WEBEXT.passwordsPrivate.getPasswordCheckStatus" directly.
func GetPasswordCheckStatus() (ret js.Promise[PasswordCheckStatus]) {
	bindings.CallGetPasswordCheckStatus(
		js.Pointer(&ret),
	)

	return
}

// TryGetPasswordCheckStatus calls the function "WEBEXT.passwordsPrivate.getPasswordCheckStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPasswordCheckStatus() (ret js.Promise[PasswordCheckStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPasswordCheckStatus(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetPasswordExceptionList returns true if the function "WEBEXT.passwordsPrivate.getPasswordExceptionList" exists.
func HasFuncGetPasswordExceptionList() bool {
	return js.True == bindings.HasFuncGetPasswordExceptionList()
}

// FuncGetPasswordExceptionList returns the function "WEBEXT.passwordsPrivate.getPasswordExceptionList".
func FuncGetPasswordExceptionList() (fn js.Func[func() js.Promise[js.Array[ExceptionEntry]]]) {
	bindings.FuncGetPasswordExceptionList(
		js.Pointer(&fn),
	)
	return
}

// GetPasswordExceptionList calls the function "WEBEXT.passwordsPrivate.getPasswordExceptionList" directly.
func GetPasswordExceptionList() (ret js.Promise[js.Array[ExceptionEntry]]) {
	bindings.CallGetPasswordExceptionList(
		js.Pointer(&ret),
	)

	return
}

// TryGetPasswordExceptionList calls the function "WEBEXT.passwordsPrivate.getPasswordExceptionList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPasswordExceptionList() (ret js.Promise[js.Array[ExceptionEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPasswordExceptionList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSavedPasswordList returns true if the function "WEBEXT.passwordsPrivate.getSavedPasswordList" exists.
func HasFuncGetSavedPasswordList() bool {
	return js.True == bindings.HasFuncGetSavedPasswordList()
}

// FuncGetSavedPasswordList returns the function "WEBEXT.passwordsPrivate.getSavedPasswordList".
func FuncGetSavedPasswordList() (fn js.Func[func() js.Promise[js.Array[PasswordUiEntry]]]) {
	bindings.FuncGetSavedPasswordList(
		js.Pointer(&fn),
	)
	return
}

// GetSavedPasswordList calls the function "WEBEXT.passwordsPrivate.getSavedPasswordList" directly.
func GetSavedPasswordList() (ret js.Promise[js.Array[PasswordUiEntry]]) {
	bindings.CallGetSavedPasswordList(
		js.Pointer(&ret),
	)

	return
}

// TryGetSavedPasswordList calls the function "WEBEXT.passwordsPrivate.getSavedPasswordList"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSavedPasswordList() (ret js.Promise[js.Array[PasswordUiEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSavedPasswordList(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetUrlCollection returns true if the function "WEBEXT.passwordsPrivate.getUrlCollection" exists.
func HasFuncGetUrlCollection() bool {
	return js.True == bindings.HasFuncGetUrlCollection()
}

// FuncGetUrlCollection returns the function "WEBEXT.passwordsPrivate.getUrlCollection".
func FuncGetUrlCollection() (fn js.Func[func(url js.String) js.Promise[UrlCollection]]) {
	bindings.FuncGetUrlCollection(
		js.Pointer(&fn),
	)
	return
}

// GetUrlCollection calls the function "WEBEXT.passwordsPrivate.getUrlCollection" directly.
func GetUrlCollection(url js.String) (ret js.Promise[UrlCollection]) {
	bindings.CallGetUrlCollection(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryGetUrlCollection calls the function "WEBEXT.passwordsPrivate.getUrlCollection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUrlCollection(url js.String) (ret js.Promise[UrlCollection], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUrlCollection(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncImportPasswords returns true if the function "WEBEXT.passwordsPrivate.importPasswords" exists.
func HasFuncImportPasswords() bool {
	return js.True == bindings.HasFuncImportPasswords()
}

// FuncImportPasswords returns the function "WEBEXT.passwordsPrivate.importPasswords".
func FuncImportPasswords() (fn js.Func[func(toStore PasswordStoreSet) js.Promise[ImportResults]]) {
	bindings.FuncImportPasswords(
		js.Pointer(&fn),
	)
	return
}

// ImportPasswords calls the function "WEBEXT.passwordsPrivate.importPasswords" directly.
func ImportPasswords(toStore PasswordStoreSet) (ret js.Promise[ImportResults]) {
	bindings.CallImportPasswords(
		js.Pointer(&ret),
		uint32(toStore),
	)

	return
}

// TryImportPasswords calls the function "WEBEXT.passwordsPrivate.importPasswords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryImportPasswords(toStore PasswordStoreSet) (ret js.Promise[ImportResults], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImportPasswords(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(toStore),
	)

	return
}

// HasFuncIsAccountStoreDefault returns true if the function "WEBEXT.passwordsPrivate.isAccountStoreDefault" exists.
func HasFuncIsAccountStoreDefault() bool {
	return js.True == bindings.HasFuncIsAccountStoreDefault()
}

// FuncIsAccountStoreDefault returns the function "WEBEXT.passwordsPrivate.isAccountStoreDefault".
func FuncIsAccountStoreDefault() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsAccountStoreDefault(
		js.Pointer(&fn),
	)
	return
}

// IsAccountStoreDefault calls the function "WEBEXT.passwordsPrivate.isAccountStoreDefault" directly.
func IsAccountStoreDefault() (ret js.Promise[js.Boolean]) {
	bindings.CallIsAccountStoreDefault(
		js.Pointer(&ret),
	)

	return
}

// TryIsAccountStoreDefault calls the function "WEBEXT.passwordsPrivate.isAccountStoreDefault"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsAccountStoreDefault() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsAccountStoreDefault(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsOptedInForAccountStorage returns true if the function "WEBEXT.passwordsPrivate.isOptedInForAccountStorage" exists.
func HasFuncIsOptedInForAccountStorage() bool {
	return js.True == bindings.HasFuncIsOptedInForAccountStorage()
}

// FuncIsOptedInForAccountStorage returns the function "WEBEXT.passwordsPrivate.isOptedInForAccountStorage".
func FuncIsOptedInForAccountStorage() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsOptedInForAccountStorage(
		js.Pointer(&fn),
	)
	return
}

// IsOptedInForAccountStorage calls the function "WEBEXT.passwordsPrivate.isOptedInForAccountStorage" directly.
func IsOptedInForAccountStorage() (ret js.Promise[js.Boolean]) {
	bindings.CallIsOptedInForAccountStorage(
		js.Pointer(&ret),
	)

	return
}

// TryIsOptedInForAccountStorage calls the function "WEBEXT.passwordsPrivate.isOptedInForAccountStorage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsOptedInForAccountStorage() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsOptedInForAccountStorage(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMovePasswordsToAccount returns true if the function "WEBEXT.passwordsPrivate.movePasswordsToAccount" exists.
func HasFuncMovePasswordsToAccount() bool {
	return js.True == bindings.HasFuncMovePasswordsToAccount()
}

// FuncMovePasswordsToAccount returns the function "WEBEXT.passwordsPrivate.movePasswordsToAccount".
func FuncMovePasswordsToAccount() (fn js.Func[func(ids js.Array[int32])]) {
	bindings.FuncMovePasswordsToAccount(
		js.Pointer(&fn),
	)
	return
}

// MovePasswordsToAccount calls the function "WEBEXT.passwordsPrivate.movePasswordsToAccount" directly.
func MovePasswordsToAccount(ids js.Array[int32]) (ret js.Void) {
	bindings.CallMovePasswordsToAccount(
		js.Pointer(&ret),
		ids.Ref(),
	)

	return
}

// TryMovePasswordsToAccount calls the function "WEBEXT.passwordsPrivate.movePasswordsToAccount"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMovePasswordsToAccount(ids js.Array[int32]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMovePasswordsToAccount(
		js.Pointer(&ret), js.Pointer(&exception),
		ids.Ref(),
	)

	return
}

// HasFuncMuteInsecureCredential returns true if the function "WEBEXT.passwordsPrivate.muteInsecureCredential" exists.
func HasFuncMuteInsecureCredential() bool {
	return js.True == bindings.HasFuncMuteInsecureCredential()
}

// FuncMuteInsecureCredential returns the function "WEBEXT.passwordsPrivate.muteInsecureCredential".
func FuncMuteInsecureCredential() (fn js.Func[func(credential PasswordUiEntry) js.Promise[js.Void]]) {
	bindings.FuncMuteInsecureCredential(
		js.Pointer(&fn),
	)
	return
}

// MuteInsecureCredential calls the function "WEBEXT.passwordsPrivate.muteInsecureCredential" directly.
func MuteInsecureCredential(credential PasswordUiEntry) (ret js.Promise[js.Void]) {
	bindings.CallMuteInsecureCredential(
		js.Pointer(&ret),
		js.Pointer(&credential),
	)

	return
}

// TryMuteInsecureCredential calls the function "WEBEXT.passwordsPrivate.muteInsecureCredential"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryMuteInsecureCredential(credential PasswordUiEntry) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMuteInsecureCredential(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&credential),
	)

	return
}

type OnAccountStorageOptInStateChangedEventCallbackFunc func(this js.Ref, optedIn bool) js.Ref

func (fn OnAccountStorageOptInStateChangedEventCallbackFunc) Register() js.Func[func(optedIn bool)] {
	return js.RegisterCallback[func(optedIn bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnAccountStorageOptInStateChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnAccountStorageOptInStateChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, optedIn bool) js.Ref
	Arg T
}

func (cb *OnAccountStorageOptInStateChangedEventCallback[T]) Register() js.Func[func(optedIn bool)] {
	return js.RegisterCallback[func(optedIn bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnAccountStorageOptInStateChangedEventCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnAccountStorageOptInStateChanged returns true if the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.addListener" exists.
func HasFuncOnAccountStorageOptInStateChanged() bool {
	return js.True == bindings.HasFuncOnAccountStorageOptInStateChanged()
}

// FuncOnAccountStorageOptInStateChanged returns the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.addListener".
func FuncOnAccountStorageOptInStateChanged() (fn js.Func[func(callback js.Func[func(optedIn bool)])]) {
	bindings.FuncOnAccountStorageOptInStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OnAccountStorageOptInStateChanged calls the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.addListener" directly.
func OnAccountStorageOptInStateChanged(callback js.Func[func(optedIn bool)]) (ret js.Void) {
	bindings.CallOnAccountStorageOptInStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnAccountStorageOptInStateChanged calls the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnAccountStorageOptInStateChanged(callback js.Func[func(optedIn bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnAccountStorageOptInStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffAccountStorageOptInStateChanged returns true if the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.removeListener" exists.
func HasFuncOffAccountStorageOptInStateChanged() bool {
	return js.True == bindings.HasFuncOffAccountStorageOptInStateChanged()
}

// FuncOffAccountStorageOptInStateChanged returns the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.removeListener".
func FuncOffAccountStorageOptInStateChanged() (fn js.Func[func(callback js.Func[func(optedIn bool)])]) {
	bindings.FuncOffAccountStorageOptInStateChanged(
		js.Pointer(&fn),
	)
	return
}

// OffAccountStorageOptInStateChanged calls the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.removeListener" directly.
func OffAccountStorageOptInStateChanged(callback js.Func[func(optedIn bool)]) (ret js.Void) {
	bindings.CallOffAccountStorageOptInStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffAccountStorageOptInStateChanged calls the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffAccountStorageOptInStateChanged(callback js.Func[func(optedIn bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffAccountStorageOptInStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnAccountStorageOptInStateChanged returns true if the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.hasListener" exists.
func HasFuncHasOnAccountStorageOptInStateChanged() bool {
	return js.True == bindings.HasFuncHasOnAccountStorageOptInStateChanged()
}

// FuncHasOnAccountStorageOptInStateChanged returns the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.hasListener".
func FuncHasOnAccountStorageOptInStateChanged() (fn js.Func[func(callback js.Func[func(optedIn bool)]) bool]) {
	bindings.FuncHasOnAccountStorageOptInStateChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnAccountStorageOptInStateChanged calls the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.hasListener" directly.
func HasOnAccountStorageOptInStateChanged(callback js.Func[func(optedIn bool)]) (ret bool) {
	bindings.CallHasOnAccountStorageOptInStateChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnAccountStorageOptInStateChanged calls the function "WEBEXT.passwordsPrivate.onAccountStorageOptInStateChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnAccountStorageOptInStateChanged(callback js.Func[func(optedIn bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnAccountStorageOptInStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnInsecureCredentialsChangedEventCallbackFunc func(this js.Ref, insecureCredentials js.Array[PasswordUiEntry]) js.Ref

func (fn OnInsecureCredentialsChangedEventCallbackFunc) Register() js.Func[func(insecureCredentials js.Array[PasswordUiEntry])] {
	return js.RegisterCallback[func(insecureCredentials js.Array[PasswordUiEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnInsecureCredentialsChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[PasswordUiEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnInsecureCredentialsChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, insecureCredentials js.Array[PasswordUiEntry]) js.Ref
	Arg T
}

func (cb *OnInsecureCredentialsChangedEventCallback[T]) Register() js.Func[func(insecureCredentials js.Array[PasswordUiEntry])] {
	return js.RegisterCallback[func(insecureCredentials js.Array[PasswordUiEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnInsecureCredentialsChangedEventCallback[T]) DispatchCallback(
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

		js.Array[PasswordUiEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnInsecureCredentialsChanged returns true if the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.addListener" exists.
func HasFuncOnInsecureCredentialsChanged() bool {
	return js.True == bindings.HasFuncOnInsecureCredentialsChanged()
}

// FuncOnInsecureCredentialsChanged returns the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.addListener".
func FuncOnInsecureCredentialsChanged() (fn js.Func[func(callback js.Func[func(insecureCredentials js.Array[PasswordUiEntry])])]) {
	bindings.FuncOnInsecureCredentialsChanged(
		js.Pointer(&fn),
	)
	return
}

// OnInsecureCredentialsChanged calls the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.addListener" directly.
func OnInsecureCredentialsChanged(callback js.Func[func(insecureCredentials js.Array[PasswordUiEntry])]) (ret js.Void) {
	bindings.CallOnInsecureCredentialsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnInsecureCredentialsChanged calls the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnInsecureCredentialsChanged(callback js.Func[func(insecureCredentials js.Array[PasswordUiEntry])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnInsecureCredentialsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffInsecureCredentialsChanged returns true if the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.removeListener" exists.
func HasFuncOffInsecureCredentialsChanged() bool {
	return js.True == bindings.HasFuncOffInsecureCredentialsChanged()
}

// FuncOffInsecureCredentialsChanged returns the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.removeListener".
func FuncOffInsecureCredentialsChanged() (fn js.Func[func(callback js.Func[func(insecureCredentials js.Array[PasswordUiEntry])])]) {
	bindings.FuncOffInsecureCredentialsChanged(
		js.Pointer(&fn),
	)
	return
}

// OffInsecureCredentialsChanged calls the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.removeListener" directly.
func OffInsecureCredentialsChanged(callback js.Func[func(insecureCredentials js.Array[PasswordUiEntry])]) (ret js.Void) {
	bindings.CallOffInsecureCredentialsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffInsecureCredentialsChanged calls the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffInsecureCredentialsChanged(callback js.Func[func(insecureCredentials js.Array[PasswordUiEntry])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffInsecureCredentialsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnInsecureCredentialsChanged returns true if the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.hasListener" exists.
func HasFuncHasOnInsecureCredentialsChanged() bool {
	return js.True == bindings.HasFuncHasOnInsecureCredentialsChanged()
}

// FuncHasOnInsecureCredentialsChanged returns the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.hasListener".
func FuncHasOnInsecureCredentialsChanged() (fn js.Func[func(callback js.Func[func(insecureCredentials js.Array[PasswordUiEntry])]) bool]) {
	bindings.FuncHasOnInsecureCredentialsChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnInsecureCredentialsChanged calls the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.hasListener" directly.
func HasOnInsecureCredentialsChanged(callback js.Func[func(insecureCredentials js.Array[PasswordUiEntry])]) (ret bool) {
	bindings.CallHasOnInsecureCredentialsChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnInsecureCredentialsChanged calls the function "WEBEXT.passwordsPrivate.onInsecureCredentialsChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnInsecureCredentialsChanged(callback js.Func[func(insecureCredentials js.Array[PasswordUiEntry])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnInsecureCredentialsChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPasswordCheckStatusChangedEventCallbackFunc func(this js.Ref, status *PasswordCheckStatus) js.Ref

func (fn OnPasswordCheckStatusChangedEventCallbackFunc) Register() js.Func[func(status *PasswordCheckStatus)] {
	return js.RegisterCallback[func(status *PasswordCheckStatus)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPasswordCheckStatusChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PasswordCheckStatus
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPasswordCheckStatusChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, status *PasswordCheckStatus) js.Ref
	Arg T
}

func (cb *OnPasswordCheckStatusChangedEventCallback[T]) Register() js.Func[func(status *PasswordCheckStatus)] {
	return js.RegisterCallback[func(status *PasswordCheckStatus)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPasswordCheckStatusChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PasswordCheckStatus
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPasswordCheckStatusChanged returns true if the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.addListener" exists.
func HasFuncOnPasswordCheckStatusChanged() bool {
	return js.True == bindings.HasFuncOnPasswordCheckStatusChanged()
}

// FuncOnPasswordCheckStatusChanged returns the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.addListener".
func FuncOnPasswordCheckStatusChanged() (fn js.Func[func(callback js.Func[func(status *PasswordCheckStatus)])]) {
	bindings.FuncOnPasswordCheckStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPasswordCheckStatusChanged calls the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.addListener" directly.
func OnPasswordCheckStatusChanged(callback js.Func[func(status *PasswordCheckStatus)]) (ret js.Void) {
	bindings.CallOnPasswordCheckStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPasswordCheckStatusChanged calls the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPasswordCheckStatusChanged(callback js.Func[func(status *PasswordCheckStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPasswordCheckStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPasswordCheckStatusChanged returns true if the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.removeListener" exists.
func HasFuncOffPasswordCheckStatusChanged() bool {
	return js.True == bindings.HasFuncOffPasswordCheckStatusChanged()
}

// FuncOffPasswordCheckStatusChanged returns the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.removeListener".
func FuncOffPasswordCheckStatusChanged() (fn js.Func[func(callback js.Func[func(status *PasswordCheckStatus)])]) {
	bindings.FuncOffPasswordCheckStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPasswordCheckStatusChanged calls the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.removeListener" directly.
func OffPasswordCheckStatusChanged(callback js.Func[func(status *PasswordCheckStatus)]) (ret js.Void) {
	bindings.CallOffPasswordCheckStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPasswordCheckStatusChanged calls the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPasswordCheckStatusChanged(callback js.Func[func(status *PasswordCheckStatus)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPasswordCheckStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPasswordCheckStatusChanged returns true if the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.hasListener" exists.
func HasFuncHasOnPasswordCheckStatusChanged() bool {
	return js.True == bindings.HasFuncHasOnPasswordCheckStatusChanged()
}

// FuncHasOnPasswordCheckStatusChanged returns the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.hasListener".
func FuncHasOnPasswordCheckStatusChanged() (fn js.Func[func(callback js.Func[func(status *PasswordCheckStatus)]) bool]) {
	bindings.FuncHasOnPasswordCheckStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPasswordCheckStatusChanged calls the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.hasListener" directly.
func HasOnPasswordCheckStatusChanged(callback js.Func[func(status *PasswordCheckStatus)]) (ret bool) {
	bindings.CallHasOnPasswordCheckStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPasswordCheckStatusChanged calls the function "WEBEXT.passwordsPrivate.onPasswordCheckStatusChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPasswordCheckStatusChanged(callback js.Func[func(status *PasswordCheckStatus)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPasswordCheckStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPasswordExceptionsListChangedEventCallbackFunc func(this js.Ref, exceptions js.Array[ExceptionEntry]) js.Ref

func (fn OnPasswordExceptionsListChangedEventCallbackFunc) Register() js.Func[func(exceptions js.Array[ExceptionEntry])] {
	return js.RegisterCallback[func(exceptions js.Array[ExceptionEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPasswordExceptionsListChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ExceptionEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPasswordExceptionsListChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, exceptions js.Array[ExceptionEntry]) js.Ref
	Arg T
}

func (cb *OnPasswordExceptionsListChangedEventCallback[T]) Register() js.Func[func(exceptions js.Array[ExceptionEntry])] {
	return js.RegisterCallback[func(exceptions js.Array[ExceptionEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPasswordExceptionsListChangedEventCallback[T]) DispatchCallback(
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

		js.Array[ExceptionEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPasswordExceptionsListChanged returns true if the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.addListener" exists.
func HasFuncOnPasswordExceptionsListChanged() bool {
	return js.True == bindings.HasFuncOnPasswordExceptionsListChanged()
}

// FuncOnPasswordExceptionsListChanged returns the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.addListener".
func FuncOnPasswordExceptionsListChanged() (fn js.Func[func(callback js.Func[func(exceptions js.Array[ExceptionEntry])])]) {
	bindings.FuncOnPasswordExceptionsListChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPasswordExceptionsListChanged calls the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.addListener" directly.
func OnPasswordExceptionsListChanged(callback js.Func[func(exceptions js.Array[ExceptionEntry])]) (ret js.Void) {
	bindings.CallOnPasswordExceptionsListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPasswordExceptionsListChanged calls the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPasswordExceptionsListChanged(callback js.Func[func(exceptions js.Array[ExceptionEntry])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPasswordExceptionsListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPasswordExceptionsListChanged returns true if the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.removeListener" exists.
func HasFuncOffPasswordExceptionsListChanged() bool {
	return js.True == bindings.HasFuncOffPasswordExceptionsListChanged()
}

// FuncOffPasswordExceptionsListChanged returns the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.removeListener".
func FuncOffPasswordExceptionsListChanged() (fn js.Func[func(callback js.Func[func(exceptions js.Array[ExceptionEntry])])]) {
	bindings.FuncOffPasswordExceptionsListChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPasswordExceptionsListChanged calls the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.removeListener" directly.
func OffPasswordExceptionsListChanged(callback js.Func[func(exceptions js.Array[ExceptionEntry])]) (ret js.Void) {
	bindings.CallOffPasswordExceptionsListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPasswordExceptionsListChanged calls the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPasswordExceptionsListChanged(callback js.Func[func(exceptions js.Array[ExceptionEntry])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPasswordExceptionsListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPasswordExceptionsListChanged returns true if the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.hasListener" exists.
func HasFuncHasOnPasswordExceptionsListChanged() bool {
	return js.True == bindings.HasFuncHasOnPasswordExceptionsListChanged()
}

// FuncHasOnPasswordExceptionsListChanged returns the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.hasListener".
func FuncHasOnPasswordExceptionsListChanged() (fn js.Func[func(callback js.Func[func(exceptions js.Array[ExceptionEntry])]) bool]) {
	bindings.FuncHasOnPasswordExceptionsListChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPasswordExceptionsListChanged calls the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.hasListener" directly.
func HasOnPasswordExceptionsListChanged(callback js.Func[func(exceptions js.Array[ExceptionEntry])]) (ret bool) {
	bindings.CallHasOnPasswordExceptionsListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPasswordExceptionsListChanged calls the function "WEBEXT.passwordsPrivate.onPasswordExceptionsListChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPasswordExceptionsListChanged(callback js.Func[func(exceptions js.Array[ExceptionEntry])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPasswordExceptionsListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPasswordManagerAuthTimeoutEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnPasswordManagerAuthTimeoutEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPasswordManagerAuthTimeoutEventCallbackFunc) DispatchCallback(
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

type OnPasswordManagerAuthTimeoutEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnPasswordManagerAuthTimeoutEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPasswordManagerAuthTimeoutEventCallback[T]) DispatchCallback(
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

// HasFuncOnPasswordManagerAuthTimeout returns true if the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.addListener" exists.
func HasFuncOnPasswordManagerAuthTimeout() bool {
	return js.True == bindings.HasFuncOnPasswordManagerAuthTimeout()
}

// FuncOnPasswordManagerAuthTimeout returns the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.addListener".
func FuncOnPasswordManagerAuthTimeout() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnPasswordManagerAuthTimeout(
		js.Pointer(&fn),
	)
	return
}

// OnPasswordManagerAuthTimeout calls the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.addListener" directly.
func OnPasswordManagerAuthTimeout(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnPasswordManagerAuthTimeout(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPasswordManagerAuthTimeout calls the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPasswordManagerAuthTimeout(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPasswordManagerAuthTimeout(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPasswordManagerAuthTimeout returns true if the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.removeListener" exists.
func HasFuncOffPasswordManagerAuthTimeout() bool {
	return js.True == bindings.HasFuncOffPasswordManagerAuthTimeout()
}

// FuncOffPasswordManagerAuthTimeout returns the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.removeListener".
func FuncOffPasswordManagerAuthTimeout() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffPasswordManagerAuthTimeout(
		js.Pointer(&fn),
	)
	return
}

// OffPasswordManagerAuthTimeout calls the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.removeListener" directly.
func OffPasswordManagerAuthTimeout(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffPasswordManagerAuthTimeout(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPasswordManagerAuthTimeout calls the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPasswordManagerAuthTimeout(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPasswordManagerAuthTimeout(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPasswordManagerAuthTimeout returns true if the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.hasListener" exists.
func HasFuncHasOnPasswordManagerAuthTimeout() bool {
	return js.True == bindings.HasFuncHasOnPasswordManagerAuthTimeout()
}

// FuncHasOnPasswordManagerAuthTimeout returns the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.hasListener".
func FuncHasOnPasswordManagerAuthTimeout() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnPasswordManagerAuthTimeout(
		js.Pointer(&fn),
	)
	return
}

// HasOnPasswordManagerAuthTimeout calls the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.hasListener" directly.
func HasOnPasswordManagerAuthTimeout(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnPasswordManagerAuthTimeout(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPasswordManagerAuthTimeout calls the function "WEBEXT.passwordsPrivate.onPasswordManagerAuthTimeout.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPasswordManagerAuthTimeout(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPasswordManagerAuthTimeout(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPasswordsFileExportProgressEventCallbackFunc func(this js.Ref, status *PasswordExportProgress) js.Ref

func (fn OnPasswordsFileExportProgressEventCallbackFunc) Register() js.Func[func(status *PasswordExportProgress)] {
	return js.RegisterCallback[func(status *PasswordExportProgress)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPasswordsFileExportProgressEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PasswordExportProgress
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPasswordsFileExportProgressEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, status *PasswordExportProgress) js.Ref
	Arg T
}

func (cb *OnPasswordsFileExportProgressEventCallback[T]) Register() js.Func[func(status *PasswordExportProgress)] {
	return js.RegisterCallback[func(status *PasswordExportProgress)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPasswordsFileExportProgressEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PasswordExportProgress
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPasswordsFileExportProgress returns true if the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.addListener" exists.
func HasFuncOnPasswordsFileExportProgress() bool {
	return js.True == bindings.HasFuncOnPasswordsFileExportProgress()
}

// FuncOnPasswordsFileExportProgress returns the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.addListener".
func FuncOnPasswordsFileExportProgress() (fn js.Func[func(callback js.Func[func(status *PasswordExportProgress)])]) {
	bindings.FuncOnPasswordsFileExportProgress(
		js.Pointer(&fn),
	)
	return
}

// OnPasswordsFileExportProgress calls the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.addListener" directly.
func OnPasswordsFileExportProgress(callback js.Func[func(status *PasswordExportProgress)]) (ret js.Void) {
	bindings.CallOnPasswordsFileExportProgress(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPasswordsFileExportProgress calls the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPasswordsFileExportProgress(callback js.Func[func(status *PasswordExportProgress)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPasswordsFileExportProgress(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPasswordsFileExportProgress returns true if the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.removeListener" exists.
func HasFuncOffPasswordsFileExportProgress() bool {
	return js.True == bindings.HasFuncOffPasswordsFileExportProgress()
}

// FuncOffPasswordsFileExportProgress returns the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.removeListener".
func FuncOffPasswordsFileExportProgress() (fn js.Func[func(callback js.Func[func(status *PasswordExportProgress)])]) {
	bindings.FuncOffPasswordsFileExportProgress(
		js.Pointer(&fn),
	)
	return
}

// OffPasswordsFileExportProgress calls the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.removeListener" directly.
func OffPasswordsFileExportProgress(callback js.Func[func(status *PasswordExportProgress)]) (ret js.Void) {
	bindings.CallOffPasswordsFileExportProgress(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPasswordsFileExportProgress calls the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPasswordsFileExportProgress(callback js.Func[func(status *PasswordExportProgress)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPasswordsFileExportProgress(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPasswordsFileExportProgress returns true if the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.hasListener" exists.
func HasFuncHasOnPasswordsFileExportProgress() bool {
	return js.True == bindings.HasFuncHasOnPasswordsFileExportProgress()
}

// FuncHasOnPasswordsFileExportProgress returns the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.hasListener".
func FuncHasOnPasswordsFileExportProgress() (fn js.Func[func(callback js.Func[func(status *PasswordExportProgress)]) bool]) {
	bindings.FuncHasOnPasswordsFileExportProgress(
		js.Pointer(&fn),
	)
	return
}

// HasOnPasswordsFileExportProgress calls the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.hasListener" directly.
func HasOnPasswordsFileExportProgress(callback js.Func[func(status *PasswordExportProgress)]) (ret bool) {
	bindings.CallHasOnPasswordsFileExportProgress(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPasswordsFileExportProgress calls the function "WEBEXT.passwordsPrivate.onPasswordsFileExportProgress.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPasswordsFileExportProgress(callback js.Func[func(status *PasswordExportProgress)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPasswordsFileExportProgress(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSavedPasswordsListChangedEventCallbackFunc func(this js.Ref, entries js.Array[PasswordUiEntry]) js.Ref

func (fn OnSavedPasswordsListChangedEventCallbackFunc) Register() js.Func[func(entries js.Array[PasswordUiEntry])] {
	return js.RegisterCallback[func(entries js.Array[PasswordUiEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSavedPasswordsListChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[PasswordUiEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSavedPasswordsListChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[PasswordUiEntry]) js.Ref
	Arg T
}

func (cb *OnSavedPasswordsListChangedEventCallback[T]) Register() js.Func[func(entries js.Array[PasswordUiEntry])] {
	return js.RegisterCallback[func(entries js.Array[PasswordUiEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSavedPasswordsListChangedEventCallback[T]) DispatchCallback(
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

		js.Array[PasswordUiEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSavedPasswordsListChanged returns true if the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.addListener" exists.
func HasFuncOnSavedPasswordsListChanged() bool {
	return js.True == bindings.HasFuncOnSavedPasswordsListChanged()
}

// FuncOnSavedPasswordsListChanged returns the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.addListener".
func FuncOnSavedPasswordsListChanged() (fn js.Func[func(callback js.Func[func(entries js.Array[PasswordUiEntry])])]) {
	bindings.FuncOnSavedPasswordsListChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSavedPasswordsListChanged calls the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.addListener" directly.
func OnSavedPasswordsListChanged(callback js.Func[func(entries js.Array[PasswordUiEntry])]) (ret js.Void) {
	bindings.CallOnSavedPasswordsListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSavedPasswordsListChanged calls the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSavedPasswordsListChanged(callback js.Func[func(entries js.Array[PasswordUiEntry])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSavedPasswordsListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSavedPasswordsListChanged returns true if the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.removeListener" exists.
func HasFuncOffSavedPasswordsListChanged() bool {
	return js.True == bindings.HasFuncOffSavedPasswordsListChanged()
}

// FuncOffSavedPasswordsListChanged returns the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.removeListener".
func FuncOffSavedPasswordsListChanged() (fn js.Func[func(callback js.Func[func(entries js.Array[PasswordUiEntry])])]) {
	bindings.FuncOffSavedPasswordsListChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSavedPasswordsListChanged calls the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.removeListener" directly.
func OffSavedPasswordsListChanged(callback js.Func[func(entries js.Array[PasswordUiEntry])]) (ret js.Void) {
	bindings.CallOffSavedPasswordsListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSavedPasswordsListChanged calls the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSavedPasswordsListChanged(callback js.Func[func(entries js.Array[PasswordUiEntry])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSavedPasswordsListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSavedPasswordsListChanged returns true if the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.hasListener" exists.
func HasFuncHasOnSavedPasswordsListChanged() bool {
	return js.True == bindings.HasFuncHasOnSavedPasswordsListChanged()
}

// FuncHasOnSavedPasswordsListChanged returns the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.hasListener".
func FuncHasOnSavedPasswordsListChanged() (fn js.Func[func(callback js.Func[func(entries js.Array[PasswordUiEntry])]) bool]) {
	bindings.FuncHasOnSavedPasswordsListChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSavedPasswordsListChanged calls the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.hasListener" directly.
func HasOnSavedPasswordsListChanged(callback js.Func[func(entries js.Array[PasswordUiEntry])]) (ret bool) {
	bindings.CallHasOnSavedPasswordsListChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSavedPasswordsListChanged calls the function "WEBEXT.passwordsPrivate.onSavedPasswordsListChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSavedPasswordsListChanged(callback js.Func[func(entries js.Array[PasswordUiEntry])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSavedPasswordsListChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOptInForAccountStorage returns true if the function "WEBEXT.passwordsPrivate.optInForAccountStorage" exists.
func HasFuncOptInForAccountStorage() bool {
	return js.True == bindings.HasFuncOptInForAccountStorage()
}

// FuncOptInForAccountStorage returns the function "WEBEXT.passwordsPrivate.optInForAccountStorage".
func FuncOptInForAccountStorage() (fn js.Func[func(optIn bool)]) {
	bindings.FuncOptInForAccountStorage(
		js.Pointer(&fn),
	)
	return
}

// OptInForAccountStorage calls the function "WEBEXT.passwordsPrivate.optInForAccountStorage" directly.
func OptInForAccountStorage(optIn bool) (ret js.Void) {
	bindings.CallOptInForAccountStorage(
		js.Pointer(&ret),
		js.Bool(bool(optIn)),
	)

	return
}

// TryOptInForAccountStorage calls the function "WEBEXT.passwordsPrivate.optInForAccountStorage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOptInForAccountStorage(optIn bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOptInForAccountStorage(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(optIn)),
	)

	return
}

// HasFuncRecordPasswordsPageAccessInSettings returns true if the function "WEBEXT.passwordsPrivate.recordPasswordsPageAccessInSettings" exists.
func HasFuncRecordPasswordsPageAccessInSettings() bool {
	return js.True == bindings.HasFuncRecordPasswordsPageAccessInSettings()
}

// FuncRecordPasswordsPageAccessInSettings returns the function "WEBEXT.passwordsPrivate.recordPasswordsPageAccessInSettings".
func FuncRecordPasswordsPageAccessInSettings() (fn js.Func[func()]) {
	bindings.FuncRecordPasswordsPageAccessInSettings(
		js.Pointer(&fn),
	)
	return
}

// RecordPasswordsPageAccessInSettings calls the function "WEBEXT.passwordsPrivate.recordPasswordsPageAccessInSettings" directly.
func RecordPasswordsPageAccessInSettings() (ret js.Void) {
	bindings.CallRecordPasswordsPageAccessInSettings(
		js.Pointer(&ret),
	)

	return
}

// TryRecordPasswordsPageAccessInSettings calls the function "WEBEXT.passwordsPrivate.recordPasswordsPageAccessInSettings"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRecordPasswordsPageAccessInSettings() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRecordPasswordsPageAccessInSettings(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemoveCredential returns true if the function "WEBEXT.passwordsPrivate.removeCredential" exists.
func HasFuncRemoveCredential() bool {
	return js.True == bindings.HasFuncRemoveCredential()
}

// FuncRemoveCredential returns the function "WEBEXT.passwordsPrivate.removeCredential".
func FuncRemoveCredential() (fn js.Func[func(id int32, fromStores PasswordStoreSet)]) {
	bindings.FuncRemoveCredential(
		js.Pointer(&fn),
	)
	return
}

// RemoveCredential calls the function "WEBEXT.passwordsPrivate.removeCredential" directly.
func RemoveCredential(id int32, fromStores PasswordStoreSet) (ret js.Void) {
	bindings.CallRemoveCredential(
		js.Pointer(&ret),
		int32(id),
		uint32(fromStores),
	)

	return
}

// TryRemoveCredential calls the function "WEBEXT.passwordsPrivate.removeCredential"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveCredential(id int32, fromStores PasswordStoreSet) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveCredential(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
		uint32(fromStores),
	)

	return
}

// HasFuncRemovePasswordException returns true if the function "WEBEXT.passwordsPrivate.removePasswordException" exists.
func HasFuncRemovePasswordException() bool {
	return js.True == bindings.HasFuncRemovePasswordException()
}

// FuncRemovePasswordException returns the function "WEBEXT.passwordsPrivate.removePasswordException".
func FuncRemovePasswordException() (fn js.Func[func(id int32)]) {
	bindings.FuncRemovePasswordException(
		js.Pointer(&fn),
	)
	return
}

// RemovePasswordException calls the function "WEBEXT.passwordsPrivate.removePasswordException" directly.
func RemovePasswordException(id int32) (ret js.Void) {
	bindings.CallRemovePasswordException(
		js.Pointer(&ret),
		int32(id),
	)

	return
}

// TryRemovePasswordException calls the function "WEBEXT.passwordsPrivate.removePasswordException"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemovePasswordException(id int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemovePasswordException(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
	)

	return
}

// HasFuncRequestCredentialsDetails returns true if the function "WEBEXT.passwordsPrivate.requestCredentialsDetails" exists.
func HasFuncRequestCredentialsDetails() bool {
	return js.True == bindings.HasFuncRequestCredentialsDetails()
}

// FuncRequestCredentialsDetails returns the function "WEBEXT.passwordsPrivate.requestCredentialsDetails".
func FuncRequestCredentialsDetails() (fn js.Func[func(ids js.Array[int32]) js.Promise[js.Array[PasswordUiEntry]]]) {
	bindings.FuncRequestCredentialsDetails(
		js.Pointer(&fn),
	)
	return
}

// RequestCredentialsDetails calls the function "WEBEXT.passwordsPrivate.requestCredentialsDetails" directly.
func RequestCredentialsDetails(ids js.Array[int32]) (ret js.Promise[js.Array[PasswordUiEntry]]) {
	bindings.CallRequestCredentialsDetails(
		js.Pointer(&ret),
		ids.Ref(),
	)

	return
}

// TryRequestCredentialsDetails calls the function "WEBEXT.passwordsPrivate.requestCredentialsDetails"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestCredentialsDetails(ids js.Array[int32]) (ret js.Promise[js.Array[PasswordUiEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestCredentialsDetails(
		js.Pointer(&ret), js.Pointer(&exception),
		ids.Ref(),
	)

	return
}

// HasFuncRequestExportProgressStatus returns true if the function "WEBEXT.passwordsPrivate.requestExportProgressStatus" exists.
func HasFuncRequestExportProgressStatus() bool {
	return js.True == bindings.HasFuncRequestExportProgressStatus()
}

// FuncRequestExportProgressStatus returns the function "WEBEXT.passwordsPrivate.requestExportProgressStatus".
func FuncRequestExportProgressStatus() (fn js.Func[func() js.Promise[ExportProgressStatus]]) {
	bindings.FuncRequestExportProgressStatus(
		js.Pointer(&fn),
	)
	return
}

// RequestExportProgressStatus calls the function "WEBEXT.passwordsPrivate.requestExportProgressStatus" directly.
func RequestExportProgressStatus() (ret js.Promise[ExportProgressStatus]) {
	bindings.CallRequestExportProgressStatus(
		js.Pointer(&ret),
	)

	return
}

// TryRequestExportProgressStatus calls the function "WEBEXT.passwordsPrivate.requestExportProgressStatus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestExportProgressStatus() (ret js.Promise[ExportProgressStatus], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestExportProgressStatus(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestPlaintextPassword returns true if the function "WEBEXT.passwordsPrivate.requestPlaintextPassword" exists.
func HasFuncRequestPlaintextPassword() bool {
	return js.True == bindings.HasFuncRequestPlaintextPassword()
}

// FuncRequestPlaintextPassword returns the function "WEBEXT.passwordsPrivate.requestPlaintextPassword".
func FuncRequestPlaintextPassword() (fn js.Func[func(id int32, reason PlaintextReason) js.Promise[js.String]]) {
	bindings.FuncRequestPlaintextPassword(
		js.Pointer(&fn),
	)
	return
}

// RequestPlaintextPassword calls the function "WEBEXT.passwordsPrivate.requestPlaintextPassword" directly.
func RequestPlaintextPassword(id int32, reason PlaintextReason) (ret js.Promise[js.String]) {
	bindings.CallRequestPlaintextPassword(
		js.Pointer(&ret),
		int32(id),
		uint32(reason),
	)

	return
}

// TryRequestPlaintextPassword calls the function "WEBEXT.passwordsPrivate.requestPlaintextPassword"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRequestPlaintextPassword(id int32, reason PlaintextReason) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRequestPlaintextPassword(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
		uint32(reason),
	)

	return
}

// HasFuncResetImporter returns true if the function "WEBEXT.passwordsPrivate.resetImporter" exists.
func HasFuncResetImporter() bool {
	return js.True == bindings.HasFuncResetImporter()
}

// FuncResetImporter returns the function "WEBEXT.passwordsPrivate.resetImporter".
func FuncResetImporter() (fn js.Func[func(deleteFile bool) js.Promise[js.Void]]) {
	bindings.FuncResetImporter(
		js.Pointer(&fn),
	)
	return
}

// ResetImporter calls the function "WEBEXT.passwordsPrivate.resetImporter" directly.
func ResetImporter(deleteFile bool) (ret js.Promise[js.Void]) {
	bindings.CallResetImporter(
		js.Pointer(&ret),
		js.Bool(bool(deleteFile)),
	)

	return
}

// TryResetImporter calls the function "WEBEXT.passwordsPrivate.resetImporter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResetImporter(deleteFile bool) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryResetImporter(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(deleteFile)),
	)

	return
}

// HasFuncSharePassword returns true if the function "WEBEXT.passwordsPrivate.sharePassword" exists.
func HasFuncSharePassword() bool {
	return js.True == bindings.HasFuncSharePassword()
}

// FuncSharePassword returns the function "WEBEXT.passwordsPrivate.sharePassword".
func FuncSharePassword() (fn js.Func[func(id int32, recipients js.Array[RecipientInfo]) js.Promise[js.Void]]) {
	bindings.FuncSharePassword(
		js.Pointer(&fn),
	)
	return
}

// SharePassword calls the function "WEBEXT.passwordsPrivate.sharePassword" directly.
func SharePassword(id int32, recipients js.Array[RecipientInfo]) (ret js.Promise[js.Void]) {
	bindings.CallSharePassword(
		js.Pointer(&ret),
		int32(id),
		recipients.Ref(),
	)

	return
}

// TrySharePassword calls the function "WEBEXT.passwordsPrivate.sharePassword"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySharePassword(id int32, recipients js.Array[RecipientInfo]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharePassword(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(id),
		recipients.Ref(),
	)

	return
}

// HasFuncShowAddShortcutDialog returns true if the function "WEBEXT.passwordsPrivate.showAddShortcutDialog" exists.
func HasFuncShowAddShortcutDialog() bool {
	return js.True == bindings.HasFuncShowAddShortcutDialog()
}

// FuncShowAddShortcutDialog returns the function "WEBEXT.passwordsPrivate.showAddShortcutDialog".
func FuncShowAddShortcutDialog() (fn js.Func[func()]) {
	bindings.FuncShowAddShortcutDialog(
		js.Pointer(&fn),
	)
	return
}

// ShowAddShortcutDialog calls the function "WEBEXT.passwordsPrivate.showAddShortcutDialog" directly.
func ShowAddShortcutDialog() (ret js.Void) {
	bindings.CallShowAddShortcutDialog(
		js.Pointer(&ret),
	)

	return
}

// TryShowAddShortcutDialog calls the function "WEBEXT.passwordsPrivate.showAddShortcutDialog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowAddShortcutDialog() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowAddShortcutDialog(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncShowExportedFileInShell returns true if the function "WEBEXT.passwordsPrivate.showExportedFileInShell" exists.
func HasFuncShowExportedFileInShell() bool {
	return js.True == bindings.HasFuncShowExportedFileInShell()
}

// FuncShowExportedFileInShell returns the function "WEBEXT.passwordsPrivate.showExportedFileInShell".
func FuncShowExportedFileInShell() (fn js.Func[func(file_path js.String)]) {
	bindings.FuncShowExportedFileInShell(
		js.Pointer(&fn),
	)
	return
}

// ShowExportedFileInShell calls the function "WEBEXT.passwordsPrivate.showExportedFileInShell" directly.
func ShowExportedFileInShell(file_path js.String) (ret js.Void) {
	bindings.CallShowExportedFileInShell(
		js.Pointer(&ret),
		file_path.Ref(),
	)

	return
}

// TryShowExportedFileInShell calls the function "WEBEXT.passwordsPrivate.showExportedFileInShell"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryShowExportedFileInShell(file_path js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryShowExportedFileInShell(
		js.Pointer(&ret), js.Pointer(&exception),
		file_path.Ref(),
	)

	return
}

// HasFuncStartPasswordCheck returns true if the function "WEBEXT.passwordsPrivate.startPasswordCheck" exists.
func HasFuncStartPasswordCheck() bool {
	return js.True == bindings.HasFuncStartPasswordCheck()
}

// FuncStartPasswordCheck returns the function "WEBEXT.passwordsPrivate.startPasswordCheck".
func FuncStartPasswordCheck() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncStartPasswordCheck(
		js.Pointer(&fn),
	)
	return
}

// StartPasswordCheck calls the function "WEBEXT.passwordsPrivate.startPasswordCheck" directly.
func StartPasswordCheck() (ret js.Promise[js.Void]) {
	bindings.CallStartPasswordCheck(
		js.Pointer(&ret),
	)

	return
}

// TryStartPasswordCheck calls the function "WEBEXT.passwordsPrivate.startPasswordCheck"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartPasswordCheck() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartPasswordCheck(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSwitchBiometricAuthBeforeFillingState returns true if the function "WEBEXT.passwordsPrivate.switchBiometricAuthBeforeFillingState" exists.
func HasFuncSwitchBiometricAuthBeforeFillingState() bool {
	return js.True == bindings.HasFuncSwitchBiometricAuthBeforeFillingState()
}

// FuncSwitchBiometricAuthBeforeFillingState returns the function "WEBEXT.passwordsPrivate.switchBiometricAuthBeforeFillingState".
func FuncSwitchBiometricAuthBeforeFillingState() (fn js.Func[func()]) {
	bindings.FuncSwitchBiometricAuthBeforeFillingState(
		js.Pointer(&fn),
	)
	return
}

// SwitchBiometricAuthBeforeFillingState calls the function "WEBEXT.passwordsPrivate.switchBiometricAuthBeforeFillingState" directly.
func SwitchBiometricAuthBeforeFillingState() (ret js.Void) {
	bindings.CallSwitchBiometricAuthBeforeFillingState(
		js.Pointer(&ret),
	)

	return
}

// TrySwitchBiometricAuthBeforeFillingState calls the function "WEBEXT.passwordsPrivate.switchBiometricAuthBeforeFillingState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySwitchBiometricAuthBeforeFillingState() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySwitchBiometricAuthBeforeFillingState(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncUndoRemoveSavedPasswordOrException returns true if the function "WEBEXT.passwordsPrivate.undoRemoveSavedPasswordOrException" exists.
func HasFuncUndoRemoveSavedPasswordOrException() bool {
	return js.True == bindings.HasFuncUndoRemoveSavedPasswordOrException()
}

// FuncUndoRemoveSavedPasswordOrException returns the function "WEBEXT.passwordsPrivate.undoRemoveSavedPasswordOrException".
func FuncUndoRemoveSavedPasswordOrException() (fn js.Func[func()]) {
	bindings.FuncUndoRemoveSavedPasswordOrException(
		js.Pointer(&fn),
	)
	return
}

// UndoRemoveSavedPasswordOrException calls the function "WEBEXT.passwordsPrivate.undoRemoveSavedPasswordOrException" directly.
func UndoRemoveSavedPasswordOrException() (ret js.Void) {
	bindings.CallUndoRemoveSavedPasswordOrException(
		js.Pointer(&ret),
	)

	return
}

// TryUndoRemoveSavedPasswordOrException calls the function "WEBEXT.passwordsPrivate.undoRemoveSavedPasswordOrException"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUndoRemoveSavedPasswordOrException() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUndoRemoveSavedPasswordOrException(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncUnmuteInsecureCredential returns true if the function "WEBEXT.passwordsPrivate.unmuteInsecureCredential" exists.
func HasFuncUnmuteInsecureCredential() bool {
	return js.True == bindings.HasFuncUnmuteInsecureCredential()
}

// FuncUnmuteInsecureCredential returns the function "WEBEXT.passwordsPrivate.unmuteInsecureCredential".
func FuncUnmuteInsecureCredential() (fn js.Func[func(credential PasswordUiEntry) js.Promise[js.Void]]) {
	bindings.FuncUnmuteInsecureCredential(
		js.Pointer(&fn),
	)
	return
}

// UnmuteInsecureCredential calls the function "WEBEXT.passwordsPrivate.unmuteInsecureCredential" directly.
func UnmuteInsecureCredential(credential PasswordUiEntry) (ret js.Promise[js.Void]) {
	bindings.CallUnmuteInsecureCredential(
		js.Pointer(&ret),
		js.Pointer(&credential),
	)

	return
}

// TryUnmuteInsecureCredential calls the function "WEBEXT.passwordsPrivate.unmuteInsecureCredential"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnmuteInsecureCredential(credential PasswordUiEntry) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnmuteInsecureCredential(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&credential),
	)

	return
}
