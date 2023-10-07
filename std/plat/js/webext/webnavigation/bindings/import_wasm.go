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

//go:wasmimport plat/js/webext/webnavigation store_GetAllFramesArgDetails
//go:noescape
func GetAllFramesArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_GetAllFramesArgDetails
//go:noescape
func GetAllFramesArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_GetAllFramesReturnTypeElem
//go:noescape
func GetAllFramesReturnTypeElemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_GetAllFramesReturnTypeElem
//go:noescape
func GetAllFramesReturnTypeElemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_GetFrameArgDetails
//go:noescape
func GetFrameArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_GetFrameArgDetails
//go:noescape
func GetFrameArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_GetFrameReturnType
//go:noescape
func GetFrameReturnTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_GetFrameReturnType
//go:noescape
func GetFrameReturnTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_OnBeforeNavigateArgDetails
//go:noescape
func OnBeforeNavigateArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_OnBeforeNavigateArgDetails
//go:noescape
func OnBeforeNavigateArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation constof_TransitionQualifier
//go:noescape
func ConstOfTransitionQualifier(str js.Ref) uint32

//go:wasmimport plat/js/webext/webnavigation constof_TransitionType
//go:noescape
func ConstOfTransitionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/webnavigation store_OnCommittedArgDetails
//go:noescape
func OnCommittedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_OnCommittedArgDetails
//go:noescape
func OnCommittedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_OnCompletedArgDetails
//go:noescape
func OnCompletedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_OnCompletedArgDetails
//go:noescape
func OnCompletedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_OnCreatedNavigationTargetArgDetails
//go:noescape
func OnCreatedNavigationTargetArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_OnCreatedNavigationTargetArgDetails
//go:noescape
func OnCreatedNavigationTargetArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_OnDOMContentLoadedArgDetails
//go:noescape
func OnDOMContentLoadedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_OnDOMContentLoadedArgDetails
//go:noescape
func OnDOMContentLoadedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_OnErrorOccurredArgDetails
//go:noescape
func OnErrorOccurredArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_OnErrorOccurredArgDetails
//go:noescape
func OnErrorOccurredArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_OnHistoryStateUpdatedArgDetails
//go:noescape
func OnHistoryStateUpdatedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_OnHistoryStateUpdatedArgDetails
//go:noescape
func OnHistoryStateUpdatedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_OnReferenceFragmentUpdatedArgDetails
//go:noescape
func OnReferenceFragmentUpdatedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_OnReferenceFragmentUpdatedArgDetails
//go:noescape
func OnReferenceFragmentUpdatedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation store_OnTabReplacedArgDetails
//go:noescape
func OnTabReplacedArgDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webnavigation load_OnTabReplacedArgDetails
//go:noescape
func OnTabReplacedArgDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webnavigation has_GetAllFrames
//go:noescape
func HasFuncGetAllFrames() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_GetAllFrames
//go:noescape
func FuncGetAllFrames(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_GetAllFrames
//go:noescape
func CallGetAllFrames(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation try_GetAllFrames
//go:noescape
func TryGetAllFrames(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_GetFrame
//go:noescape
func HasFuncGetFrame() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_GetFrame
//go:noescape
func FuncGetFrame(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_GetFrame
//go:noescape
func CallGetFrame(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation try_GetFrame
//go:noescape
func TryGetFrame(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OnBeforeNavigate
//go:noescape
func HasFuncOnBeforeNavigate() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OnBeforeNavigate
//go:noescape
func FuncOnBeforeNavigate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OnBeforeNavigate
//go:noescape
func CallOnBeforeNavigate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OnBeforeNavigate
//go:noescape
func TryOnBeforeNavigate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OffBeforeNavigate
//go:noescape
func HasFuncOffBeforeNavigate() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OffBeforeNavigate
//go:noescape
func FuncOffBeforeNavigate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OffBeforeNavigate
//go:noescape
func CallOffBeforeNavigate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OffBeforeNavigate
//go:noescape
func TryOffBeforeNavigate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_HasOnBeforeNavigate
//go:noescape
func HasFuncHasOnBeforeNavigate() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_HasOnBeforeNavigate
//go:noescape
func FuncHasOnBeforeNavigate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_HasOnBeforeNavigate
//go:noescape
func CallHasOnBeforeNavigate(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_HasOnBeforeNavigate
//go:noescape
func TryHasOnBeforeNavigate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OnCommitted
//go:noescape
func HasFuncOnCommitted() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OnCommitted
//go:noescape
func FuncOnCommitted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OnCommitted
//go:noescape
func CallOnCommitted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OnCommitted
//go:noescape
func TryOnCommitted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OffCommitted
//go:noescape
func HasFuncOffCommitted() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OffCommitted
//go:noescape
func FuncOffCommitted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OffCommitted
//go:noescape
func CallOffCommitted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OffCommitted
//go:noescape
func TryOffCommitted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_HasOnCommitted
//go:noescape
func HasFuncHasOnCommitted() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_HasOnCommitted
//go:noescape
func FuncHasOnCommitted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_HasOnCommitted
//go:noescape
func CallHasOnCommitted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_HasOnCommitted
//go:noescape
func TryHasOnCommitted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OnCompleted
//go:noescape
func HasFuncOnCompleted() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OnCompleted
//go:noescape
func FuncOnCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OnCompleted
//go:noescape
func CallOnCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OnCompleted
//go:noescape
func TryOnCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OffCompleted
//go:noescape
func HasFuncOffCompleted() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OffCompleted
//go:noescape
func FuncOffCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OffCompleted
//go:noescape
func CallOffCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OffCompleted
//go:noescape
func TryOffCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_HasOnCompleted
//go:noescape
func HasFuncHasOnCompleted() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_HasOnCompleted
//go:noescape
func FuncHasOnCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_HasOnCompleted
//go:noescape
func CallHasOnCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_HasOnCompleted
//go:noescape
func TryHasOnCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OnCreatedNavigationTarget
//go:noescape
func HasFuncOnCreatedNavigationTarget() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OnCreatedNavigationTarget
//go:noescape
func FuncOnCreatedNavigationTarget(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OnCreatedNavigationTarget
//go:noescape
func CallOnCreatedNavigationTarget(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OnCreatedNavigationTarget
//go:noescape
func TryOnCreatedNavigationTarget(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OffCreatedNavigationTarget
//go:noescape
func HasFuncOffCreatedNavigationTarget() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OffCreatedNavigationTarget
//go:noescape
func FuncOffCreatedNavigationTarget(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OffCreatedNavigationTarget
//go:noescape
func CallOffCreatedNavigationTarget(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OffCreatedNavigationTarget
//go:noescape
func TryOffCreatedNavigationTarget(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_HasOnCreatedNavigationTarget
//go:noescape
func HasFuncHasOnCreatedNavigationTarget() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_HasOnCreatedNavigationTarget
//go:noescape
func FuncHasOnCreatedNavigationTarget(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_HasOnCreatedNavigationTarget
//go:noescape
func CallHasOnCreatedNavigationTarget(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_HasOnCreatedNavigationTarget
//go:noescape
func TryHasOnCreatedNavigationTarget(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OnDOMContentLoaded
//go:noescape
func HasFuncOnDOMContentLoaded() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OnDOMContentLoaded
//go:noescape
func FuncOnDOMContentLoaded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OnDOMContentLoaded
//go:noescape
func CallOnDOMContentLoaded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OnDOMContentLoaded
//go:noescape
func TryOnDOMContentLoaded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OffDOMContentLoaded
//go:noescape
func HasFuncOffDOMContentLoaded() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OffDOMContentLoaded
//go:noescape
func FuncOffDOMContentLoaded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OffDOMContentLoaded
//go:noescape
func CallOffDOMContentLoaded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OffDOMContentLoaded
//go:noescape
func TryOffDOMContentLoaded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_HasOnDOMContentLoaded
//go:noescape
func HasFuncHasOnDOMContentLoaded() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_HasOnDOMContentLoaded
//go:noescape
func FuncHasOnDOMContentLoaded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_HasOnDOMContentLoaded
//go:noescape
func CallHasOnDOMContentLoaded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_HasOnDOMContentLoaded
//go:noescape
func TryHasOnDOMContentLoaded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OnErrorOccurred
//go:noescape
func HasFuncOnErrorOccurred() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OnErrorOccurred
//go:noescape
func FuncOnErrorOccurred(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OnErrorOccurred
//go:noescape
func CallOnErrorOccurred(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OnErrorOccurred
//go:noescape
func TryOnErrorOccurred(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OffErrorOccurred
//go:noescape
func HasFuncOffErrorOccurred() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OffErrorOccurred
//go:noescape
func FuncOffErrorOccurred(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OffErrorOccurred
//go:noescape
func CallOffErrorOccurred(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OffErrorOccurred
//go:noescape
func TryOffErrorOccurred(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_HasOnErrorOccurred
//go:noescape
func HasFuncHasOnErrorOccurred() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_HasOnErrorOccurred
//go:noescape
func FuncHasOnErrorOccurred(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_HasOnErrorOccurred
//go:noescape
func CallHasOnErrorOccurred(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_HasOnErrorOccurred
//go:noescape
func TryHasOnErrorOccurred(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OnHistoryStateUpdated
//go:noescape
func HasFuncOnHistoryStateUpdated() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OnHistoryStateUpdated
//go:noescape
func FuncOnHistoryStateUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OnHistoryStateUpdated
//go:noescape
func CallOnHistoryStateUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OnHistoryStateUpdated
//go:noescape
func TryOnHistoryStateUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OffHistoryStateUpdated
//go:noescape
func HasFuncOffHistoryStateUpdated() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OffHistoryStateUpdated
//go:noescape
func FuncOffHistoryStateUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OffHistoryStateUpdated
//go:noescape
func CallOffHistoryStateUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OffHistoryStateUpdated
//go:noescape
func TryOffHistoryStateUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_HasOnHistoryStateUpdated
//go:noescape
func HasFuncHasOnHistoryStateUpdated() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_HasOnHistoryStateUpdated
//go:noescape
func FuncHasOnHistoryStateUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_HasOnHistoryStateUpdated
//go:noescape
func CallHasOnHistoryStateUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_HasOnHistoryStateUpdated
//go:noescape
func TryHasOnHistoryStateUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OnReferenceFragmentUpdated
//go:noescape
func HasFuncOnReferenceFragmentUpdated() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OnReferenceFragmentUpdated
//go:noescape
func FuncOnReferenceFragmentUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OnReferenceFragmentUpdated
//go:noescape
func CallOnReferenceFragmentUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OnReferenceFragmentUpdated
//go:noescape
func TryOnReferenceFragmentUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OffReferenceFragmentUpdated
//go:noescape
func HasFuncOffReferenceFragmentUpdated() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OffReferenceFragmentUpdated
//go:noescape
func FuncOffReferenceFragmentUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OffReferenceFragmentUpdated
//go:noescape
func CallOffReferenceFragmentUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OffReferenceFragmentUpdated
//go:noescape
func TryOffReferenceFragmentUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_HasOnReferenceFragmentUpdated
//go:noescape
func HasFuncHasOnReferenceFragmentUpdated() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_HasOnReferenceFragmentUpdated
//go:noescape
func FuncHasOnReferenceFragmentUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_HasOnReferenceFragmentUpdated
//go:noescape
func CallHasOnReferenceFragmentUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_HasOnReferenceFragmentUpdated
//go:noescape
func TryHasOnReferenceFragmentUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OnTabReplaced
//go:noescape
func HasFuncOnTabReplaced() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OnTabReplaced
//go:noescape
func FuncOnTabReplaced(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OnTabReplaced
//go:noescape
func CallOnTabReplaced(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OnTabReplaced
//go:noescape
func TryOnTabReplaced(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_OffTabReplaced
//go:noescape
func HasFuncOffTabReplaced() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_OffTabReplaced
//go:noescape
func FuncOffTabReplaced(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_OffTabReplaced
//go:noescape
func CallOffTabReplaced(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_OffTabReplaced
//go:noescape
func TryOffTabReplaced(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webnavigation has_HasOnTabReplaced
//go:noescape
func HasFuncHasOnTabReplaced() js.Ref

//go:wasmimport plat/js/webext/webnavigation func_HasOnTabReplaced
//go:noescape
func FuncHasOnTabReplaced(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webnavigation call_HasOnTabReplaced
//go:noescape
func CallHasOnTabReplaced(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webnavigation try_HasOnTabReplaced
//go:noescape
func TryHasOnTabReplaced(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)
