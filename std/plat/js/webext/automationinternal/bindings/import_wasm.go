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

//go:wasmimport plat/js/webext/automationinternal store_AXEventParams
//go:noescape
func AXEventParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_AXEventParams
//go:noescape
func AXEventParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_AXTextLocationParams
//go:noescape
func AXTextLocationParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_AXTextLocationParams
//go:noescape
func AXTextLocationParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_GetImageDataParams
//go:noescape
func GetImageDataParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_GetImageDataParams
//go:noescape
func GetImageDataParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_GetTextLocationDataParams
//go:noescape
func GetTextLocationDataParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_GetTextLocationDataParams
//go:noescape
func GetTextLocationDataParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_HitTestParams
//go:noescape
func HitTestParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_HitTestParams
//go:noescape
func HitTestParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_PerformActionRequiredParams
//go:noescape
func PerformActionRequiredParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_PerformActionRequiredParams
//go:noescape
func PerformActionRequiredParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_PerformCustomActionParams
//go:noescape
func PerformCustomActionParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_PerformCustomActionParams
//go:noescape
func PerformCustomActionParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_ReplaceSelectedTextParams
//go:noescape
func ReplaceSelectedTextParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_ReplaceSelectedTextParams
//go:noescape
func ReplaceSelectedTextParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_ScrollToPointParams
//go:noescape
func ScrollToPointParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_ScrollToPointParams
//go:noescape
func ScrollToPointParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_ScrollToPositionAtRowColumnParams
//go:noescape
func ScrollToPositionAtRowColumnParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_ScrollToPositionAtRowColumnParams
//go:noescape
func ScrollToPositionAtRowColumnParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_SetScrollOffsetParams
//go:noescape
func SetScrollOffsetParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_SetScrollOffsetParams
//go:noescape
func SetScrollOffsetParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_SetSelectionParams
//go:noescape
func SetSelectionParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_SetSelectionParams
//go:noescape
func SetSelectionParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal store_SetValueParams
//go:noescape
func SetValueParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automationinternal load_SetValueParams
//go:noescape
func SetValueParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automationinternal has_DisableDesktop
//go:noescape
func HasFuncDisableDesktop() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_DisableDesktop
//go:noescape
func FuncDisableDesktop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_DisableDesktop
//go:noescape
func CallDisableDesktop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_DisableDesktop
//go:noescape
func TryDisableDesktop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_EnableDesktop
//go:noescape
func HasFuncEnableDesktop() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_EnableDesktop
//go:noescape
func FuncEnableDesktop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_EnableDesktop
//go:noescape
func CallEnableDesktop(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal try_EnableDesktop
//go:noescape
func TryEnableDesktop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_EnableTree
//go:noescape
func HasFuncEnableTree() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_EnableTree
//go:noescape
func FuncEnableTree(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_EnableTree
//go:noescape
func CallEnableTree(
	retPtr unsafe.Pointer,
	tree_id js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_EnableTree
//go:noescape
func TryEnableTree(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tree_id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OnAccessibilityEvent
//go:noescape
func HasFuncOnAccessibilityEvent() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OnAccessibilityEvent
//go:noescape
func FuncOnAccessibilityEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OnAccessibilityEvent
//go:noescape
func CallOnAccessibilityEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OnAccessibilityEvent
//go:noescape
func TryOnAccessibilityEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OffAccessibilityEvent
//go:noescape
func HasFuncOffAccessibilityEvent() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OffAccessibilityEvent
//go:noescape
func FuncOffAccessibilityEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OffAccessibilityEvent
//go:noescape
func CallOffAccessibilityEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OffAccessibilityEvent
//go:noescape
func TryOffAccessibilityEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_HasOnAccessibilityEvent
//go:noescape
func HasFuncHasOnAccessibilityEvent() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_HasOnAccessibilityEvent
//go:noescape
func FuncHasOnAccessibilityEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_HasOnAccessibilityEvent
//go:noescape
func CallHasOnAccessibilityEvent(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_HasOnAccessibilityEvent
//go:noescape
func TryHasOnAccessibilityEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OnAccessibilityTreeDestroyed
//go:noescape
func HasFuncOnAccessibilityTreeDestroyed() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OnAccessibilityTreeDestroyed
//go:noescape
func FuncOnAccessibilityTreeDestroyed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OnAccessibilityTreeDestroyed
//go:noescape
func CallOnAccessibilityTreeDestroyed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OnAccessibilityTreeDestroyed
//go:noescape
func TryOnAccessibilityTreeDestroyed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OffAccessibilityTreeDestroyed
//go:noescape
func HasFuncOffAccessibilityTreeDestroyed() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OffAccessibilityTreeDestroyed
//go:noescape
func FuncOffAccessibilityTreeDestroyed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OffAccessibilityTreeDestroyed
//go:noescape
func CallOffAccessibilityTreeDestroyed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OffAccessibilityTreeDestroyed
//go:noescape
func TryOffAccessibilityTreeDestroyed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_HasOnAccessibilityTreeDestroyed
//go:noescape
func HasFuncHasOnAccessibilityTreeDestroyed() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_HasOnAccessibilityTreeDestroyed
//go:noescape
func FuncHasOnAccessibilityTreeDestroyed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_HasOnAccessibilityTreeDestroyed
//go:noescape
func CallHasOnAccessibilityTreeDestroyed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_HasOnAccessibilityTreeDestroyed
//go:noescape
func TryHasOnAccessibilityTreeDestroyed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OnAccessibilityTreeSerializationError
//go:noescape
func HasFuncOnAccessibilityTreeSerializationError() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OnAccessibilityTreeSerializationError
//go:noescape
func FuncOnAccessibilityTreeSerializationError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OnAccessibilityTreeSerializationError
//go:noescape
func CallOnAccessibilityTreeSerializationError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OnAccessibilityTreeSerializationError
//go:noescape
func TryOnAccessibilityTreeSerializationError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OffAccessibilityTreeSerializationError
//go:noescape
func HasFuncOffAccessibilityTreeSerializationError() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OffAccessibilityTreeSerializationError
//go:noescape
func FuncOffAccessibilityTreeSerializationError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OffAccessibilityTreeSerializationError
//go:noescape
func CallOffAccessibilityTreeSerializationError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OffAccessibilityTreeSerializationError
//go:noescape
func TryOffAccessibilityTreeSerializationError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_HasOnAccessibilityTreeSerializationError
//go:noescape
func HasFuncHasOnAccessibilityTreeSerializationError() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_HasOnAccessibilityTreeSerializationError
//go:noescape
func FuncHasOnAccessibilityTreeSerializationError(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_HasOnAccessibilityTreeSerializationError
//go:noescape
func CallHasOnAccessibilityTreeSerializationError(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_HasOnAccessibilityTreeSerializationError
//go:noescape
func TryHasOnAccessibilityTreeSerializationError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OnActionResult
//go:noescape
func HasFuncOnActionResult() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OnActionResult
//go:noescape
func FuncOnActionResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OnActionResult
//go:noescape
func CallOnActionResult(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OnActionResult
//go:noescape
func TryOnActionResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OffActionResult
//go:noescape
func HasFuncOffActionResult() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OffActionResult
//go:noescape
func FuncOffActionResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OffActionResult
//go:noescape
func CallOffActionResult(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OffActionResult
//go:noescape
func TryOffActionResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_HasOnActionResult
//go:noescape
func HasFuncHasOnActionResult() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_HasOnActionResult
//go:noescape
func FuncHasOnActionResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_HasOnActionResult
//go:noescape
func CallHasOnActionResult(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_HasOnActionResult
//go:noescape
func TryHasOnActionResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OnAllAutomationEventListenersRemoved
//go:noescape
func HasFuncOnAllAutomationEventListenersRemoved() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OnAllAutomationEventListenersRemoved
//go:noescape
func FuncOnAllAutomationEventListenersRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OnAllAutomationEventListenersRemoved
//go:noescape
func CallOnAllAutomationEventListenersRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OnAllAutomationEventListenersRemoved
//go:noescape
func TryOnAllAutomationEventListenersRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OffAllAutomationEventListenersRemoved
//go:noescape
func HasFuncOffAllAutomationEventListenersRemoved() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OffAllAutomationEventListenersRemoved
//go:noescape
func FuncOffAllAutomationEventListenersRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OffAllAutomationEventListenersRemoved
//go:noescape
func CallOffAllAutomationEventListenersRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OffAllAutomationEventListenersRemoved
//go:noescape
func TryOffAllAutomationEventListenersRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_HasOnAllAutomationEventListenersRemoved
//go:noescape
func HasFuncHasOnAllAutomationEventListenersRemoved() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_HasOnAllAutomationEventListenersRemoved
//go:noescape
func FuncHasOnAllAutomationEventListenersRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_HasOnAllAutomationEventListenersRemoved
//go:noescape
func CallHasOnAllAutomationEventListenersRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_HasOnAllAutomationEventListenersRemoved
//go:noescape
func TryHasOnAllAutomationEventListenersRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OnChildTreeID
//go:noescape
func HasFuncOnChildTreeID() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OnChildTreeID
//go:noescape
func FuncOnChildTreeID(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OnChildTreeID
//go:noescape
func CallOnChildTreeID(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OnChildTreeID
//go:noescape
func TryOnChildTreeID(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OffChildTreeID
//go:noescape
func HasFuncOffChildTreeID() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OffChildTreeID
//go:noescape
func FuncOffChildTreeID(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OffChildTreeID
//go:noescape
func CallOffChildTreeID(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OffChildTreeID
//go:noescape
func TryOffChildTreeID(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_HasOnChildTreeID
//go:noescape
func HasFuncHasOnChildTreeID() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_HasOnChildTreeID
//go:noescape
func FuncHasOnChildTreeID(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_HasOnChildTreeID
//go:noescape
func CallHasOnChildTreeID(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_HasOnChildTreeID
//go:noescape
func TryHasOnChildTreeID(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OnGetTextLocationResult
//go:noescape
func HasFuncOnGetTextLocationResult() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OnGetTextLocationResult
//go:noescape
func FuncOnGetTextLocationResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OnGetTextLocationResult
//go:noescape
func CallOnGetTextLocationResult(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OnGetTextLocationResult
//go:noescape
func TryOnGetTextLocationResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OffGetTextLocationResult
//go:noescape
func HasFuncOffGetTextLocationResult() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OffGetTextLocationResult
//go:noescape
func FuncOffGetTextLocationResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OffGetTextLocationResult
//go:noescape
func CallOffGetTextLocationResult(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OffGetTextLocationResult
//go:noescape
func TryOffGetTextLocationResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_HasOnGetTextLocationResult
//go:noescape
func HasFuncHasOnGetTextLocationResult() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_HasOnGetTextLocationResult
//go:noescape
func FuncHasOnGetTextLocationResult(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_HasOnGetTextLocationResult
//go:noescape
func CallHasOnGetTextLocationResult(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_HasOnGetTextLocationResult
//go:noescape
func TryHasOnGetTextLocationResult(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OnNodesRemoved
//go:noescape
func HasFuncOnNodesRemoved() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OnNodesRemoved
//go:noescape
func FuncOnNodesRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OnNodesRemoved
//go:noescape
func CallOnNodesRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OnNodesRemoved
//go:noescape
func TryOnNodesRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OffNodesRemoved
//go:noescape
func HasFuncOffNodesRemoved() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OffNodesRemoved
//go:noescape
func FuncOffNodesRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OffNodesRemoved
//go:noescape
func CallOffNodesRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OffNodesRemoved
//go:noescape
func TryOffNodesRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_HasOnNodesRemoved
//go:noescape
func HasFuncHasOnNodesRemoved() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_HasOnNodesRemoved
//go:noescape
func FuncHasOnNodesRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_HasOnNodesRemoved
//go:noescape
func CallHasOnNodesRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_HasOnNodesRemoved
//go:noescape
func TryHasOnNodesRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OnTreeChange
//go:noescape
func HasFuncOnTreeChange() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OnTreeChange
//go:noescape
func FuncOnTreeChange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OnTreeChange
//go:noescape
func CallOnTreeChange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OnTreeChange
//go:noescape
func TryOnTreeChange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_OffTreeChange
//go:noescape
func HasFuncOffTreeChange() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_OffTreeChange
//go:noescape
func FuncOffTreeChange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_OffTreeChange
//go:noescape
func CallOffTreeChange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_OffTreeChange
//go:noescape
func TryOffTreeChange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_HasOnTreeChange
//go:noescape
func HasFuncHasOnTreeChange() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_HasOnTreeChange
//go:noescape
func FuncHasOnTreeChange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_HasOnTreeChange
//go:noescape
func CallHasOnTreeChange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_HasOnTreeChange
//go:noescape
func TryHasOnTreeChange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automationinternal has_PerformAction
//go:noescape
func HasFuncPerformAction() js.Ref

//go:wasmimport plat/js/webext/automationinternal func_PerformAction
//go:noescape
func FuncPerformAction(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automationinternal call_PerformAction
//go:noescape
func CallPerformAction(
	retPtr unsafe.Pointer,
	args unsafe.Pointer,
	opt_args js.Ref)

//go:wasmimport plat/js/webext/automationinternal try_PerformAction
//go:noescape
func TryPerformAction(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	args unsafe.Pointer,
	opt_args js.Ref) (ok js.Ref)
