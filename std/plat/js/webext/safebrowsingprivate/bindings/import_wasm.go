// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/webext/safebrowsingprivate store_DangerousDownloadInfo
//go:noescape
func DangerousDownloadInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate load_DangerousDownloadInfo
//go:noescape
func DangerousDownloadInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate constof_URLType
//go:noescape
func ConstOfURLType(str js.Ref) uint32

//go:wasmimport plat/js/webext/safebrowsingprivate store_ServerRedirect
//go:noescape
func ServerRedirectJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate load_ServerRedirect
//go:noescape
func ServerRedirectJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate constof_NavigationInitiation
//go:noescape
func ConstOfNavigationInitiation(str js.Ref) uint32

//go:wasmimport plat/js/webext/safebrowsingprivate store_ReferrerChainEntry
//go:noescape
func ReferrerChainEntryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate load_ReferrerChainEntry
//go:noescape
func ReferrerChainEntryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate store_InterstitialInfo
//go:noescape
func InterstitialInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate load_InterstitialInfo
//go:noescape
func InterstitialInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate store_PolicySpecifiedPasswordReuse
//go:noescape
func PolicySpecifiedPasswordReuseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate load_PolicySpecifiedPasswordReuse
//go:noescape
func PolicySpecifiedPasswordReuseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate has_GetReferrerChain
//go:noescape
func HasFuncGetReferrerChain() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_GetReferrerChain
//go:noescape
func FuncGetReferrerChain(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_GetReferrerChain
//go:noescape
func CallGetReferrerChain(
	retPtr unsafe.Pointer,
	tabId int32)

//go:wasmimport plat/js/webext/safebrowsingprivate try_GetReferrerChain
//go:noescape
func TryGetReferrerChain(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OnDangerousDownloadOpened
//go:noescape
func HasFuncOnDangerousDownloadOpened() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OnDangerousDownloadOpened
//go:noescape
func FuncOnDangerousDownloadOpened(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OnDangerousDownloadOpened
//go:noescape
func CallOnDangerousDownloadOpened(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OnDangerousDownloadOpened
//go:noescape
func TryOnDangerousDownloadOpened(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OffDangerousDownloadOpened
//go:noescape
func HasFuncOffDangerousDownloadOpened() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OffDangerousDownloadOpened
//go:noescape
func FuncOffDangerousDownloadOpened(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OffDangerousDownloadOpened
//go:noescape
func CallOffDangerousDownloadOpened(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OffDangerousDownloadOpened
//go:noescape
func TryOffDangerousDownloadOpened(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_HasOnDangerousDownloadOpened
//go:noescape
func HasFuncHasOnDangerousDownloadOpened() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_HasOnDangerousDownloadOpened
//go:noescape
func FuncHasOnDangerousDownloadOpened(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_HasOnDangerousDownloadOpened
//go:noescape
func CallHasOnDangerousDownloadOpened(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_HasOnDangerousDownloadOpened
//go:noescape
func TryHasOnDangerousDownloadOpened(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OnPolicySpecifiedPasswordChanged
//go:noescape
func HasFuncOnPolicySpecifiedPasswordChanged() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OnPolicySpecifiedPasswordChanged
//go:noescape
func FuncOnPolicySpecifiedPasswordChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OnPolicySpecifiedPasswordChanged
//go:noescape
func CallOnPolicySpecifiedPasswordChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OnPolicySpecifiedPasswordChanged
//go:noescape
func TryOnPolicySpecifiedPasswordChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OffPolicySpecifiedPasswordChanged
//go:noescape
func HasFuncOffPolicySpecifiedPasswordChanged() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OffPolicySpecifiedPasswordChanged
//go:noescape
func FuncOffPolicySpecifiedPasswordChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OffPolicySpecifiedPasswordChanged
//go:noescape
func CallOffPolicySpecifiedPasswordChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OffPolicySpecifiedPasswordChanged
//go:noescape
func TryOffPolicySpecifiedPasswordChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_HasOnPolicySpecifiedPasswordChanged
//go:noescape
func HasFuncHasOnPolicySpecifiedPasswordChanged() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_HasOnPolicySpecifiedPasswordChanged
//go:noescape
func FuncHasOnPolicySpecifiedPasswordChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_HasOnPolicySpecifiedPasswordChanged
//go:noescape
func CallHasOnPolicySpecifiedPasswordChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_HasOnPolicySpecifiedPasswordChanged
//go:noescape
func TryHasOnPolicySpecifiedPasswordChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OnPolicySpecifiedPasswordReuseDetected
//go:noescape
func HasFuncOnPolicySpecifiedPasswordReuseDetected() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OnPolicySpecifiedPasswordReuseDetected
//go:noescape
func FuncOnPolicySpecifiedPasswordReuseDetected(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OnPolicySpecifiedPasswordReuseDetected
//go:noescape
func CallOnPolicySpecifiedPasswordReuseDetected(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OnPolicySpecifiedPasswordReuseDetected
//go:noescape
func TryOnPolicySpecifiedPasswordReuseDetected(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OffPolicySpecifiedPasswordReuseDetected
//go:noescape
func HasFuncOffPolicySpecifiedPasswordReuseDetected() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OffPolicySpecifiedPasswordReuseDetected
//go:noescape
func FuncOffPolicySpecifiedPasswordReuseDetected(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OffPolicySpecifiedPasswordReuseDetected
//go:noescape
func CallOffPolicySpecifiedPasswordReuseDetected(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OffPolicySpecifiedPasswordReuseDetected
//go:noescape
func TryOffPolicySpecifiedPasswordReuseDetected(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_HasOnPolicySpecifiedPasswordReuseDetected
//go:noescape
func HasFuncHasOnPolicySpecifiedPasswordReuseDetected() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_HasOnPolicySpecifiedPasswordReuseDetected
//go:noescape
func FuncHasOnPolicySpecifiedPasswordReuseDetected(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_HasOnPolicySpecifiedPasswordReuseDetected
//go:noescape
func CallHasOnPolicySpecifiedPasswordReuseDetected(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_HasOnPolicySpecifiedPasswordReuseDetected
//go:noescape
func TryHasOnPolicySpecifiedPasswordReuseDetected(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OnSecurityInterstitialProceeded
//go:noescape
func HasFuncOnSecurityInterstitialProceeded() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OnSecurityInterstitialProceeded
//go:noescape
func FuncOnSecurityInterstitialProceeded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OnSecurityInterstitialProceeded
//go:noescape
func CallOnSecurityInterstitialProceeded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OnSecurityInterstitialProceeded
//go:noescape
func TryOnSecurityInterstitialProceeded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OffSecurityInterstitialProceeded
//go:noescape
func HasFuncOffSecurityInterstitialProceeded() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OffSecurityInterstitialProceeded
//go:noescape
func FuncOffSecurityInterstitialProceeded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OffSecurityInterstitialProceeded
//go:noescape
func CallOffSecurityInterstitialProceeded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OffSecurityInterstitialProceeded
//go:noescape
func TryOffSecurityInterstitialProceeded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_HasOnSecurityInterstitialProceeded
//go:noescape
func HasFuncHasOnSecurityInterstitialProceeded() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_HasOnSecurityInterstitialProceeded
//go:noescape
func FuncHasOnSecurityInterstitialProceeded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_HasOnSecurityInterstitialProceeded
//go:noescape
func CallHasOnSecurityInterstitialProceeded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_HasOnSecurityInterstitialProceeded
//go:noescape
func TryHasOnSecurityInterstitialProceeded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OnSecurityInterstitialShown
//go:noescape
func HasFuncOnSecurityInterstitialShown() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OnSecurityInterstitialShown
//go:noescape
func FuncOnSecurityInterstitialShown(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OnSecurityInterstitialShown
//go:noescape
func CallOnSecurityInterstitialShown(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OnSecurityInterstitialShown
//go:noescape
func TryOnSecurityInterstitialShown(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_OffSecurityInterstitialShown
//go:noescape
func HasFuncOffSecurityInterstitialShown() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_OffSecurityInterstitialShown
//go:noescape
func FuncOffSecurityInterstitialShown(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_OffSecurityInterstitialShown
//go:noescape
func CallOffSecurityInterstitialShown(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_OffSecurityInterstitialShown
//go:noescape
func TryOffSecurityInterstitialShown(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate has_HasOnSecurityInterstitialShown
//go:noescape
func HasFuncHasOnSecurityInterstitialShown() js.Ref

//go:wasmimport plat/js/webext/safebrowsingprivate func_HasOnSecurityInterstitialShown
//go:noescape
func FuncHasOnSecurityInterstitialShown(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/safebrowsingprivate call_HasOnSecurityInterstitialShown
//go:noescape
func CallHasOnSecurityInterstitialShown(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/safebrowsingprivate try_HasOnSecurityInterstitialShown
//go:noescape
func TryHasOnSecurityInterstitialShown(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
