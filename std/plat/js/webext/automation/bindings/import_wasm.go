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

//go:wasmimport plat/js/webext/automation store_Rect
//go:noescape
func RectJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automation load_Rect
//go:noescape
func RectJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation constof_EventType
//go:noescape
func ConstOfEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_ActionType
//go:noescape
func ConstOfActionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_IntentCommandType
//go:noescape
func ConstOfIntentCommandType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_IntentTextBoundaryType
//go:noescape
func ConstOfIntentTextBoundaryType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_IntentMoveDirectionType
//go:noescape
func ConstOfIntentMoveDirectionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation store_AutomationIntent
//go:noescape
func AutomationIntentJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automation load_AutomationIntent
//go:noescape
func AutomationIntentJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationEvent_Target
//go:noescape
func GetAutomationEventTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationEvent_Target
//go:noescape
func SetAutomationEventTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationEvent_Type
//go:noescape
func GetAutomationEventType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationEvent_Type
//go:noescape
func SetAutomationEventType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationEvent_EventFrom
//go:noescape
func GetAutomationEventEventFrom(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationEvent_EventFrom
//go:noescape
func SetAutomationEventEventFrom(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationEvent_MouseX
//go:noescape
func GetAutomationEventMouseX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationEvent_MouseX
//go:noescape
func SetAutomationEventMouseX(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationEvent_MouseY
//go:noescape
func GetAutomationEventMouseY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationEvent_MouseY
//go:noescape
func SetAutomationEventMouseY(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationEvent_Intents
//go:noescape
func GetAutomationEventIntents(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationEvent_Intents
//go:noescape
func SetAutomationEventIntents(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation has_AutomationEvent_StopPropagation
//go:noescape
func HasFuncAutomationEventStopPropagation(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationEvent_StopPropagation
//go:noescape
func FuncAutomationEventStopPropagation(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationEvent_StopPropagation
//go:noescape
func CallAutomationEventStopPropagation(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationEvent_StopPropagation
//go:noescape
func TryAutomationEventStopPropagation(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation constof_RoleType
//go:noescape
func ConstOfRoleType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation store_FindParams
//go:noescape
func FindParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automation load_FindParams
//go:noescape
func FindParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationPosition_Node
//go:noescape
func GetAutomationPositionNode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationPosition_Node
//go:noescape
func SetAutomationPositionNode(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationPosition_ChildIndex
//go:noescape
func GetAutomationPositionChildIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationPosition_ChildIndex
//go:noescape
func SetAutomationPositionChildIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationPosition_TextOffset
//go:noescape
func GetAutomationPositionTextOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationPosition_TextOffset
//go:noescape
func SetAutomationPositionTextOffset(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationPosition_Affinity
//go:noescape
func GetAutomationPositionAffinity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationPosition_Affinity
//go:noescape
func SetAutomationPositionAffinity(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation has_AutomationPosition_IsNullPosition
//go:noescape
func HasFuncAutomationPositionIsNullPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_IsNullPosition
//go:noescape
func FuncAutomationPositionIsNullPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_IsNullPosition
//go:noescape
func CallAutomationPositionIsNullPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_IsNullPosition
//go:noescape
func TryAutomationPositionIsNullPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_IsTreePosition
//go:noescape
func HasFuncAutomationPositionIsTreePosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_IsTreePosition
//go:noescape
func FuncAutomationPositionIsTreePosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_IsTreePosition
//go:noescape
func CallAutomationPositionIsTreePosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_IsTreePosition
//go:noescape
func TryAutomationPositionIsTreePosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_IsTextPosition
//go:noescape
func HasFuncAutomationPositionIsTextPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_IsTextPosition
//go:noescape
func FuncAutomationPositionIsTextPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_IsTextPosition
//go:noescape
func CallAutomationPositionIsTextPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_IsTextPosition
//go:noescape
func TryAutomationPositionIsTextPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_IsLeafTextPosition
//go:noescape
func HasFuncAutomationPositionIsLeafTextPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_IsLeafTextPosition
//go:noescape
func FuncAutomationPositionIsLeafTextPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_IsLeafTextPosition
//go:noescape
func CallAutomationPositionIsLeafTextPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_IsLeafTextPosition
//go:noescape
func TryAutomationPositionIsLeafTextPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtStartOfAnchor
//go:noescape
func HasFuncAutomationPositionAtStartOfAnchor(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtStartOfAnchor
//go:noescape
func FuncAutomationPositionAtStartOfAnchor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtStartOfAnchor
//go:noescape
func CallAutomationPositionAtStartOfAnchor(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtStartOfAnchor
//go:noescape
func TryAutomationPositionAtStartOfAnchor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtEndOfAnchor
//go:noescape
func HasFuncAutomationPositionAtEndOfAnchor(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtEndOfAnchor
//go:noescape
func FuncAutomationPositionAtEndOfAnchor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtEndOfAnchor
//go:noescape
func CallAutomationPositionAtEndOfAnchor(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtEndOfAnchor
//go:noescape
func TryAutomationPositionAtEndOfAnchor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtStartOfWord
//go:noescape
func HasFuncAutomationPositionAtStartOfWord(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtStartOfWord
//go:noescape
func FuncAutomationPositionAtStartOfWord(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtStartOfWord
//go:noescape
func CallAutomationPositionAtStartOfWord(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtStartOfWord
//go:noescape
func TryAutomationPositionAtStartOfWord(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtEndOfWord
//go:noescape
func HasFuncAutomationPositionAtEndOfWord(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtEndOfWord
//go:noescape
func FuncAutomationPositionAtEndOfWord(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtEndOfWord
//go:noescape
func CallAutomationPositionAtEndOfWord(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtEndOfWord
//go:noescape
func TryAutomationPositionAtEndOfWord(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtStartOfLine
//go:noescape
func HasFuncAutomationPositionAtStartOfLine(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtStartOfLine
//go:noescape
func FuncAutomationPositionAtStartOfLine(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtStartOfLine
//go:noescape
func CallAutomationPositionAtStartOfLine(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtStartOfLine
//go:noescape
func TryAutomationPositionAtStartOfLine(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtEndOfLine
//go:noescape
func HasFuncAutomationPositionAtEndOfLine(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtEndOfLine
//go:noescape
func FuncAutomationPositionAtEndOfLine(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtEndOfLine
//go:noescape
func CallAutomationPositionAtEndOfLine(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtEndOfLine
//go:noescape
func TryAutomationPositionAtEndOfLine(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtStartOfParagraph
//go:noescape
func HasFuncAutomationPositionAtStartOfParagraph(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtStartOfParagraph
//go:noescape
func FuncAutomationPositionAtStartOfParagraph(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtStartOfParagraph
//go:noescape
func CallAutomationPositionAtStartOfParagraph(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtStartOfParagraph
//go:noescape
func TryAutomationPositionAtStartOfParagraph(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtEndOfParagraph
//go:noescape
func HasFuncAutomationPositionAtEndOfParagraph(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtEndOfParagraph
//go:noescape
func FuncAutomationPositionAtEndOfParagraph(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtEndOfParagraph
//go:noescape
func CallAutomationPositionAtEndOfParagraph(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtEndOfParagraph
//go:noescape
func TryAutomationPositionAtEndOfParagraph(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtStartOfPage
//go:noescape
func HasFuncAutomationPositionAtStartOfPage(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtStartOfPage
//go:noescape
func FuncAutomationPositionAtStartOfPage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtStartOfPage
//go:noescape
func CallAutomationPositionAtStartOfPage(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtStartOfPage
//go:noescape
func TryAutomationPositionAtStartOfPage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtEndOfPage
//go:noescape
func HasFuncAutomationPositionAtEndOfPage(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtEndOfPage
//go:noescape
func FuncAutomationPositionAtEndOfPage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtEndOfPage
//go:noescape
func CallAutomationPositionAtEndOfPage(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtEndOfPage
//go:noescape
func TryAutomationPositionAtEndOfPage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtStartOfFormat
//go:noescape
func HasFuncAutomationPositionAtStartOfFormat(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtStartOfFormat
//go:noescape
func FuncAutomationPositionAtStartOfFormat(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtStartOfFormat
//go:noescape
func CallAutomationPositionAtStartOfFormat(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtStartOfFormat
//go:noescape
func TryAutomationPositionAtStartOfFormat(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtEndOfFormat
//go:noescape
func HasFuncAutomationPositionAtEndOfFormat(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtEndOfFormat
//go:noescape
func FuncAutomationPositionAtEndOfFormat(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtEndOfFormat
//go:noescape
func CallAutomationPositionAtEndOfFormat(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtEndOfFormat
//go:noescape
func TryAutomationPositionAtEndOfFormat(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtStartOfDocument
//go:noescape
func HasFuncAutomationPositionAtStartOfDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtStartOfDocument
//go:noescape
func FuncAutomationPositionAtStartOfDocument(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtStartOfDocument
//go:noescape
func CallAutomationPositionAtStartOfDocument(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtStartOfDocument
//go:noescape
func TryAutomationPositionAtStartOfDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AtEndOfDocument
//go:noescape
func HasFuncAutomationPositionAtEndOfDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AtEndOfDocument
//go:noescape
func FuncAutomationPositionAtEndOfDocument(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AtEndOfDocument
//go:noescape
func CallAutomationPositionAtEndOfDocument(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AtEndOfDocument
//go:noescape
func TryAutomationPositionAtEndOfDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AsTreePosition
//go:noescape
func HasFuncAutomationPositionAsTreePosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AsTreePosition
//go:noescape
func FuncAutomationPositionAsTreePosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AsTreePosition
//go:noescape
func CallAutomationPositionAsTreePosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AsTreePosition
//go:noescape
func TryAutomationPositionAsTreePosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AsTextPosition
//go:noescape
func HasFuncAutomationPositionAsTextPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AsTextPosition
//go:noescape
func FuncAutomationPositionAsTextPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AsTextPosition
//go:noescape
func CallAutomationPositionAsTextPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AsTextPosition
//go:noescape
func TryAutomationPositionAsTextPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_AsLeafTextPosition
//go:noescape
func HasFuncAutomationPositionAsLeafTextPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_AsLeafTextPosition
//go:noescape
func FuncAutomationPositionAsLeafTextPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_AsLeafTextPosition
//go:noescape
func CallAutomationPositionAsLeafTextPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_AsLeafTextPosition
//go:noescape
func TryAutomationPositionAsLeafTextPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPositionAtStartOfAnchor
//go:noescape
func HasFuncAutomationPositionMoveToPositionAtStartOfAnchor(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPositionAtStartOfAnchor
//go:noescape
func FuncAutomationPositionMoveToPositionAtStartOfAnchor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPositionAtStartOfAnchor
//go:noescape
func CallAutomationPositionMoveToPositionAtStartOfAnchor(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPositionAtStartOfAnchor
//go:noescape
func TryAutomationPositionMoveToPositionAtStartOfAnchor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPositionAtEndOfAnchor
//go:noescape
func HasFuncAutomationPositionMoveToPositionAtEndOfAnchor(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPositionAtEndOfAnchor
//go:noescape
func FuncAutomationPositionMoveToPositionAtEndOfAnchor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPositionAtEndOfAnchor
//go:noescape
func CallAutomationPositionMoveToPositionAtEndOfAnchor(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPositionAtEndOfAnchor
//go:noescape
func TryAutomationPositionMoveToPositionAtEndOfAnchor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPositionAtStartOfDocument
//go:noescape
func HasFuncAutomationPositionMoveToPositionAtStartOfDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPositionAtStartOfDocument
//go:noescape
func FuncAutomationPositionMoveToPositionAtStartOfDocument(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPositionAtStartOfDocument
//go:noescape
func CallAutomationPositionMoveToPositionAtStartOfDocument(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPositionAtStartOfDocument
//go:noescape
func TryAutomationPositionMoveToPositionAtStartOfDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPositionAtEndOfDocument
//go:noescape
func HasFuncAutomationPositionMoveToPositionAtEndOfDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPositionAtEndOfDocument
//go:noescape
func FuncAutomationPositionMoveToPositionAtEndOfDocument(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPositionAtEndOfDocument
//go:noescape
func CallAutomationPositionMoveToPositionAtEndOfDocument(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPositionAtEndOfDocument
//go:noescape
func TryAutomationPositionMoveToPositionAtEndOfDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToParentPosition
//go:noescape
func HasFuncAutomationPositionMoveToParentPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToParentPosition
//go:noescape
func FuncAutomationPositionMoveToParentPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToParentPosition
//go:noescape
func CallAutomationPositionMoveToParentPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToParentPosition
//go:noescape
func TryAutomationPositionMoveToParentPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextLeafTreePosition
//go:noescape
func HasFuncAutomationPositionMoveToNextLeafTreePosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextLeafTreePosition
//go:noescape
func FuncAutomationPositionMoveToNextLeafTreePosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextLeafTreePosition
//go:noescape
func CallAutomationPositionMoveToNextLeafTreePosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextLeafTreePosition
//go:noescape
func TryAutomationPositionMoveToNextLeafTreePosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousLeafTreePosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousLeafTreePosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousLeafTreePosition
//go:noescape
func FuncAutomationPositionMoveToPreviousLeafTreePosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousLeafTreePosition
//go:noescape
func CallAutomationPositionMoveToPreviousLeafTreePosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousLeafTreePosition
//go:noescape
func TryAutomationPositionMoveToPreviousLeafTreePosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextLeafTextPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextLeafTextPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextLeafTextPosition
//go:noescape
func FuncAutomationPositionMoveToNextLeafTextPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextLeafTextPosition
//go:noescape
func CallAutomationPositionMoveToNextLeafTextPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextLeafTextPosition
//go:noescape
func TryAutomationPositionMoveToNextLeafTextPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousLeafTextPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousLeafTextPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousLeafTextPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousLeafTextPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousLeafTextPosition
//go:noescape
func CallAutomationPositionMoveToPreviousLeafTextPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousLeafTextPosition
//go:noescape
func TryAutomationPositionMoveToPreviousLeafTextPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextCharacterPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextCharacterPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextCharacterPosition
//go:noescape
func FuncAutomationPositionMoveToNextCharacterPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextCharacterPosition
//go:noescape
func CallAutomationPositionMoveToNextCharacterPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextCharacterPosition
//go:noescape
func TryAutomationPositionMoveToNextCharacterPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousCharacterPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousCharacterPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousCharacterPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousCharacterPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousCharacterPosition
//go:noescape
func CallAutomationPositionMoveToPreviousCharacterPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousCharacterPosition
//go:noescape
func TryAutomationPositionMoveToPreviousCharacterPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextWordStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextWordStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextWordStartPosition
//go:noescape
func FuncAutomationPositionMoveToNextWordStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextWordStartPosition
//go:noescape
func CallAutomationPositionMoveToNextWordStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextWordStartPosition
//go:noescape
func TryAutomationPositionMoveToNextWordStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousWordStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousWordStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousWordStartPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousWordStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousWordStartPosition
//go:noescape
func CallAutomationPositionMoveToPreviousWordStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousWordStartPosition
//go:noescape
func TryAutomationPositionMoveToPreviousWordStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextWordEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextWordEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextWordEndPosition
//go:noescape
func FuncAutomationPositionMoveToNextWordEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextWordEndPosition
//go:noescape
func CallAutomationPositionMoveToNextWordEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextWordEndPosition
//go:noescape
func TryAutomationPositionMoveToNextWordEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousWordEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousWordEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousWordEndPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousWordEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousWordEndPosition
//go:noescape
func CallAutomationPositionMoveToPreviousWordEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousWordEndPosition
//go:noescape
func TryAutomationPositionMoveToPreviousWordEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextLineStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextLineStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextLineStartPosition
//go:noescape
func FuncAutomationPositionMoveToNextLineStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextLineStartPosition
//go:noescape
func CallAutomationPositionMoveToNextLineStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextLineStartPosition
//go:noescape
func TryAutomationPositionMoveToNextLineStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousLineStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousLineStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousLineStartPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousLineStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousLineStartPosition
//go:noescape
func CallAutomationPositionMoveToPreviousLineStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousLineStartPosition
//go:noescape
func TryAutomationPositionMoveToPreviousLineStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextLineEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextLineEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextLineEndPosition
//go:noescape
func FuncAutomationPositionMoveToNextLineEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextLineEndPosition
//go:noescape
func CallAutomationPositionMoveToNextLineEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextLineEndPosition
//go:noescape
func TryAutomationPositionMoveToNextLineEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousLineEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousLineEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousLineEndPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousLineEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousLineEndPosition
//go:noescape
func CallAutomationPositionMoveToPreviousLineEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousLineEndPosition
//go:noescape
func TryAutomationPositionMoveToPreviousLineEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextFormatStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextFormatStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextFormatStartPosition
//go:noescape
func FuncAutomationPositionMoveToNextFormatStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextFormatStartPosition
//go:noescape
func CallAutomationPositionMoveToNextFormatStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextFormatStartPosition
//go:noescape
func TryAutomationPositionMoveToNextFormatStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousFormatStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousFormatStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousFormatStartPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousFormatStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousFormatStartPosition
//go:noescape
func CallAutomationPositionMoveToPreviousFormatStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousFormatStartPosition
//go:noescape
func TryAutomationPositionMoveToPreviousFormatStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextFormatEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextFormatEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextFormatEndPosition
//go:noescape
func FuncAutomationPositionMoveToNextFormatEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextFormatEndPosition
//go:noescape
func CallAutomationPositionMoveToNextFormatEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextFormatEndPosition
//go:noescape
func TryAutomationPositionMoveToNextFormatEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousFormatEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousFormatEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousFormatEndPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousFormatEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousFormatEndPosition
//go:noescape
func CallAutomationPositionMoveToPreviousFormatEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousFormatEndPosition
//go:noescape
func TryAutomationPositionMoveToPreviousFormatEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextParagraphStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextParagraphStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextParagraphStartPosition
//go:noescape
func FuncAutomationPositionMoveToNextParagraphStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextParagraphStartPosition
//go:noescape
func CallAutomationPositionMoveToNextParagraphStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextParagraphStartPosition
//go:noescape
func TryAutomationPositionMoveToNextParagraphStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousParagraphStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousParagraphStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousParagraphStartPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousParagraphStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousParagraphStartPosition
//go:noescape
func CallAutomationPositionMoveToPreviousParagraphStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousParagraphStartPosition
//go:noescape
func TryAutomationPositionMoveToPreviousParagraphStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextParagraphEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextParagraphEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextParagraphEndPosition
//go:noescape
func FuncAutomationPositionMoveToNextParagraphEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextParagraphEndPosition
//go:noescape
func CallAutomationPositionMoveToNextParagraphEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextParagraphEndPosition
//go:noescape
func TryAutomationPositionMoveToNextParagraphEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousParagraphEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousParagraphEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousParagraphEndPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousParagraphEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousParagraphEndPosition
//go:noescape
func CallAutomationPositionMoveToPreviousParagraphEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousParagraphEndPosition
//go:noescape
func TryAutomationPositionMoveToPreviousParagraphEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextPageStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextPageStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextPageStartPosition
//go:noescape
func FuncAutomationPositionMoveToNextPageStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextPageStartPosition
//go:noescape
func CallAutomationPositionMoveToNextPageStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextPageStartPosition
//go:noescape
func TryAutomationPositionMoveToNextPageStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousPageStartPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousPageStartPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousPageStartPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousPageStartPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousPageStartPosition
//go:noescape
func CallAutomationPositionMoveToPreviousPageStartPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousPageStartPosition
//go:noescape
func TryAutomationPositionMoveToPreviousPageStartPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextPageEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextPageEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextPageEndPosition
//go:noescape
func FuncAutomationPositionMoveToNextPageEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextPageEndPosition
//go:noescape
func CallAutomationPositionMoveToNextPageEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextPageEndPosition
//go:noescape
func TryAutomationPositionMoveToNextPageEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousPageEndPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousPageEndPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousPageEndPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousPageEndPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousPageEndPosition
//go:noescape
func CallAutomationPositionMoveToPreviousPageEndPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousPageEndPosition
//go:noescape
func TryAutomationPositionMoveToPreviousPageEndPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToNextAnchorPosition
//go:noescape
func HasFuncAutomationPositionMoveToNextAnchorPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToNextAnchorPosition
//go:noescape
func FuncAutomationPositionMoveToNextAnchorPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToNextAnchorPosition
//go:noescape
func CallAutomationPositionMoveToNextAnchorPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToNextAnchorPosition
//go:noescape
func TryAutomationPositionMoveToNextAnchorPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MoveToPreviousAnchorPosition
//go:noescape
func HasFuncAutomationPositionMoveToPreviousAnchorPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MoveToPreviousAnchorPosition
//go:noescape
func FuncAutomationPositionMoveToPreviousAnchorPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MoveToPreviousAnchorPosition
//go:noescape
func CallAutomationPositionMoveToPreviousAnchorPosition(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MoveToPreviousAnchorPosition
//go:noescape
func TryAutomationPositionMoveToPreviousAnchorPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_MaxTextOffset
//go:noescape
func HasFuncAutomationPositionMaxTextOffset(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_MaxTextOffset
//go:noescape
func FuncAutomationPositionMaxTextOffset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_MaxTextOffset
//go:noescape
func CallAutomationPositionMaxTextOffset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_MaxTextOffset
//go:noescape
func TryAutomationPositionMaxTextOffset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_IsInLineBreak
//go:noescape
func HasFuncAutomationPositionIsInLineBreak(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_IsInLineBreak
//go:noescape
func FuncAutomationPositionIsInLineBreak(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_IsInLineBreak
//go:noescape
func CallAutomationPositionIsInLineBreak(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_IsInLineBreak
//go:noescape
func TryAutomationPositionIsInLineBreak(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_IsInTextObject
//go:noescape
func HasFuncAutomationPositionIsInTextObject(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_IsInTextObject
//go:noescape
func FuncAutomationPositionIsInTextObject(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_IsInTextObject
//go:noescape
func CallAutomationPositionIsInTextObject(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_IsInTextObject
//go:noescape
func TryAutomationPositionIsInTextObject(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_IsInWhiteSpace
//go:noescape
func HasFuncAutomationPositionIsInWhiteSpace(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_IsInWhiteSpace
//go:noescape
func FuncAutomationPositionIsInWhiteSpace(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_IsInWhiteSpace
//go:noescape
func CallAutomationPositionIsInWhiteSpace(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_IsInWhiteSpace
//go:noescape
func TryAutomationPositionIsInWhiteSpace(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_IsValid
//go:noescape
func HasFuncAutomationPositionIsValid(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_IsValid
//go:noescape
func FuncAutomationPositionIsValid(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_IsValid
//go:noescape
func CallAutomationPositionIsValid(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_IsValid
//go:noescape
func TryAutomationPositionIsValid(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationPosition_GetText
//go:noescape
func HasFuncAutomationPositionGetText(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationPosition_GetText
//go:noescape
func FuncAutomationPositionGetText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationPosition_GetText
//go:noescape
func CallAutomationPositionGetText(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationPosition_GetText
//go:noescape
func TryAutomationPositionGetText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation constof_PositionType
//go:noescape
func ConstOfPositionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_NameFromType
//go:noescape
func ConstOfNameFromType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation store_CustomAction
//go:noescape
func CustomActionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automation load_CustomAction
//go:noescape
func CustomActionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation constof_DefaultActionVerb
//go:noescape
func ConstOfDefaultActionVerb(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation store_Marker
//go:noescape
func MarkerJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automation load_Marker
//go:noescape
func MarkerJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation constof_HasPopup
//go:noescape
func ConstOfHasPopup(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_AriaCurrentState
//go:noescape
func ConstOfAriaCurrentState(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_InvalidState
//go:noescape
func ConstOfInvalidState(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_SortDirectionType
//go:noescape
func ConstOfSortDirectionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation get_AutomationNode_Root
//go:noescape
func GetAutomationNodeRoot(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Root
//go:noescape
func SetAutomationNodeRoot(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_IsRootNode
//go:noescape
func GetAutomationNodeIsRootNode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_IsRootNode
//go:noescape
func SetAutomationNodeIsRootNode(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Role
//go:noescape
func GetAutomationNodeRole(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Role
//go:noescape
func SetAutomationNodeRole(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_State
//go:noescape
func GetAutomationNodeState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_State
//go:noescape
func SetAutomationNodeState(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Location
//go:noescape
func GetAutomationNodeLocation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Location
//go:noescape
func SetAutomationNodeLocation(
	this js.Ref,
	val unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_UnclippedLocation
//go:noescape
func GetAutomationNodeUnclippedLocation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_UnclippedLocation
//go:noescape
func SetAutomationNodeUnclippedLocation(
	this js.Ref,
	val unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Description
//go:noescape
func GetAutomationNodeDescription(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Description
//go:noescape
func SetAutomationNodeDescription(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_CheckedStateDescription
//go:noescape
func GetAutomationNodeCheckedStateDescription(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_CheckedStateDescription
//go:noescape
func SetAutomationNodeCheckedStateDescription(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Placeholder
//go:noescape
func GetAutomationNodePlaceholder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Placeholder
//go:noescape
func SetAutomationNodePlaceholder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_RoleDescription
//go:noescape
func GetAutomationNodeRoleDescription(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_RoleDescription
//go:noescape
func SetAutomationNodeRoleDescription(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Name
//go:noescape
func GetAutomationNodeName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Name
//go:noescape
func SetAutomationNodeName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DoDefaultLabel
//go:noescape
func GetAutomationNodeDoDefaultLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DoDefaultLabel
//go:noescape
func SetAutomationNodeDoDefaultLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_LongClickLabel
//go:noescape
func GetAutomationNodeLongClickLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_LongClickLabel
//go:noescape
func SetAutomationNodeLongClickLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Tooltip
//go:noescape
func GetAutomationNodeTooltip(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Tooltip
//go:noescape
func SetAutomationNodeTooltip(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_NameFrom
//go:noescape
func GetAutomationNodeNameFrom(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_NameFrom
//go:noescape
func SetAutomationNodeNameFrom(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ImageAnnotation
//go:noescape
func GetAutomationNodeImageAnnotation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ImageAnnotation
//go:noescape
func SetAutomationNodeImageAnnotation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Value
//go:noescape
func GetAutomationNodeValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation get_AutomationNode_HtmlTag
//go:noescape
func GetAutomationNodeHtmlTag(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_HtmlTag
//go:noescape
func SetAutomationNodeHtmlTag(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_HierarchicalLevel
//go:noescape
func GetAutomationNodeHierarchicalLevel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_HierarchicalLevel
//go:noescape
func SetAutomationNodeHierarchicalLevel(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_CaretBounds
//go:noescape
func GetAutomationNodeCaretBounds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_CaretBounds
//go:noescape
func SetAutomationNodeCaretBounds(
	this js.Ref,
	val unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_WordStarts
//go:noescape
func GetAutomationNodeWordStarts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_WordStarts
//go:noescape
func SetAutomationNodeWordStarts(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_WordEnds
//go:noescape
func GetAutomationNodeWordEnds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_WordEnds
//go:noescape
func SetAutomationNodeWordEnds(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SentenceStarts
//go:noescape
func GetAutomationNodeSentenceStarts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SentenceStarts
//go:noescape
func SetAutomationNodeSentenceStarts(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SentenceEnds
//go:noescape
func GetAutomationNodeSentenceEnds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SentenceEnds
//go:noescape
func SetAutomationNodeSentenceEnds(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_NonInlineTextWordStarts
//go:noescape
func GetAutomationNodeNonInlineTextWordStarts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_NonInlineTextWordStarts
//go:noescape
func SetAutomationNodeNonInlineTextWordStarts(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_NonInlineTextWordEnds
//go:noescape
func GetAutomationNodeNonInlineTextWordEnds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_NonInlineTextWordEnds
//go:noescape
func SetAutomationNodeNonInlineTextWordEnds(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Controls
//go:noescape
func GetAutomationNodeControls(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Controls
//go:noescape
func SetAutomationNodeControls(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DescribedBy
//go:noescape
func GetAutomationNodeDescribedBy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DescribedBy
//go:noescape
func SetAutomationNodeDescribedBy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_FlowTo
//go:noescape
func GetAutomationNodeFlowTo(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_FlowTo
//go:noescape
func SetAutomationNodeFlowTo(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_LabelledBy
//go:noescape
func GetAutomationNodeLabelledBy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_LabelledBy
//go:noescape
func SetAutomationNodeLabelledBy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ActiveDescendant
//go:noescape
func GetAutomationNodeActiveDescendant(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ActiveDescendant
//go:noescape
func SetAutomationNodeActiveDescendant(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ActiveDescendantFor
//go:noescape
func GetAutomationNodeActiveDescendantFor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ActiveDescendantFor
//go:noescape
func SetAutomationNodeActiveDescendantFor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_InPageLinkTarget
//go:noescape
func GetAutomationNodeInPageLinkTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_InPageLinkTarget
//go:noescape
func SetAutomationNodeInPageLinkTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Details
//go:noescape
func GetAutomationNodeDetails(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Details
//go:noescape
func SetAutomationNodeDetails(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ErrorMessages
//go:noescape
func GetAutomationNodeErrorMessages(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ErrorMessages
//go:noescape
func SetAutomationNodeErrorMessages(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DetailsFor
//go:noescape
func GetAutomationNodeDetailsFor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DetailsFor
//go:noescape
func SetAutomationNodeDetailsFor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ErrorMessageFor
//go:noescape
func GetAutomationNodeErrorMessageFor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ErrorMessageFor
//go:noescape
func SetAutomationNodeErrorMessageFor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ControlledBy
//go:noescape
func GetAutomationNodeControlledBy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ControlledBy
//go:noescape
func SetAutomationNodeControlledBy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DescriptionFor
//go:noescape
func GetAutomationNodeDescriptionFor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DescriptionFor
//go:noescape
func SetAutomationNodeDescriptionFor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_FlowFrom
//go:noescape
func GetAutomationNodeFlowFrom(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_FlowFrom
//go:noescape
func SetAutomationNodeFlowFrom(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_LabelFor
//go:noescape
func GetAutomationNodeLabelFor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_LabelFor
//go:noescape
func SetAutomationNodeLabelFor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableCellColumnHeaders
//go:noescape
func GetAutomationNodeTableCellColumnHeaders(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableCellColumnHeaders
//go:noescape
func SetAutomationNodeTableCellColumnHeaders(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableCellRowHeaders
//go:noescape
func GetAutomationNodeTableCellRowHeaders(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableCellRowHeaders
//go:noescape
func SetAutomationNodeTableCellRowHeaders(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_StandardActions
//go:noescape
func GetAutomationNodeStandardActions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_StandardActions
//go:noescape
func SetAutomationNodeStandardActions(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_CustomActions
//go:noescape
func GetAutomationNodeCustomActions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_CustomActions
//go:noescape
func SetAutomationNodeCustomActions(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DefaultActionVerb
//go:noescape
func GetAutomationNodeDefaultActionVerb(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DefaultActionVerb
//go:noescape
func SetAutomationNodeDefaultActionVerb(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Url
//go:noescape
func GetAutomationNodeUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Url
//go:noescape
func SetAutomationNodeUrl(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DocUrl
//go:noescape
func GetAutomationNodeDocUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DocUrl
//go:noescape
func SetAutomationNodeDocUrl(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DocTitle
//go:noescape
func GetAutomationNodeDocTitle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DocTitle
//go:noescape
func SetAutomationNodeDocTitle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DocLoaded
//go:noescape
func GetAutomationNodeDocLoaded(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DocLoaded
//go:noescape
func SetAutomationNodeDocLoaded(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DocLoadingProgress
//go:noescape
func GetAutomationNodeDocLoadingProgress(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DocLoadingProgress
//go:noescape
func SetAutomationNodeDocLoadingProgress(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ScrollX
//go:noescape
func GetAutomationNodeScrollX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ScrollX
//go:noescape
func SetAutomationNodeScrollX(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ScrollXMin
//go:noescape
func GetAutomationNodeScrollXMin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ScrollXMin
//go:noescape
func SetAutomationNodeScrollXMin(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ScrollXMax
//go:noescape
func GetAutomationNodeScrollXMax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ScrollXMax
//go:noescape
func SetAutomationNodeScrollXMax(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ScrollY
//go:noescape
func GetAutomationNodeScrollY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ScrollY
//go:noescape
func SetAutomationNodeScrollY(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ScrollYMin
//go:noescape
func GetAutomationNodeScrollYMin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ScrollYMin
//go:noescape
func SetAutomationNodeScrollYMin(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ScrollYMax
//go:noescape
func GetAutomationNodeScrollYMax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ScrollYMax
//go:noescape
func SetAutomationNodeScrollYMax(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Scrollable
//go:noescape
func GetAutomationNodeScrollable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Scrollable
//go:noescape
func SetAutomationNodeScrollable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TextSelStart
//go:noescape
func GetAutomationNodeTextSelStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TextSelStart
//go:noescape
func SetAutomationNodeTextSelStart(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TextSelEnd
//go:noescape
func GetAutomationNodeTextSelEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TextSelEnd
//go:noescape
func SetAutomationNodeTextSelEnd(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Markers
//go:noescape
func GetAutomationNodeMarkers(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Markers
//go:noescape
func SetAutomationNodeMarkers(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_IsSelectionBackward
//go:noescape
func GetAutomationNodeIsSelectionBackward(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_IsSelectionBackward
//go:noescape
func SetAutomationNodeIsSelectionBackward(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AnchorObject
//go:noescape
func GetAutomationNodeAnchorObject(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AnchorObject
//go:noescape
func SetAutomationNodeAnchorObject(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AnchorOffset
//go:noescape
func GetAutomationNodeAnchorOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AnchorOffset
//go:noescape
func SetAutomationNodeAnchorOffset(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AnchorAffinity
//go:noescape
func GetAutomationNodeAnchorAffinity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AnchorAffinity
//go:noescape
func SetAutomationNodeAnchorAffinity(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_FocusObject
//go:noescape
func GetAutomationNodeFocusObject(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_FocusObject
//go:noescape
func SetAutomationNodeFocusObject(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_FocusOffset
//go:noescape
func GetAutomationNodeFocusOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_FocusOffset
//go:noescape
func SetAutomationNodeFocusOffset(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_FocusAffinity
//go:noescape
func GetAutomationNodeFocusAffinity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_FocusAffinity
//go:noescape
func SetAutomationNodeFocusAffinity(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SelectionStartObject
//go:noescape
func GetAutomationNodeSelectionStartObject(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SelectionStartObject
//go:noescape
func SetAutomationNodeSelectionStartObject(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SelectionStartOffset
//go:noescape
func GetAutomationNodeSelectionStartOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SelectionStartOffset
//go:noescape
func SetAutomationNodeSelectionStartOffset(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SelectionStartAffinity
//go:noescape
func GetAutomationNodeSelectionStartAffinity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SelectionStartAffinity
//go:noescape
func SetAutomationNodeSelectionStartAffinity(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SelectionEndObject
//go:noescape
func GetAutomationNodeSelectionEndObject(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SelectionEndObject
//go:noescape
func SetAutomationNodeSelectionEndObject(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SelectionEndOffset
//go:noescape
func GetAutomationNodeSelectionEndOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SelectionEndOffset
//go:noescape
func SetAutomationNodeSelectionEndOffset(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SelectionEndAffinity
//go:noescape
func GetAutomationNodeSelectionEndAffinity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SelectionEndAffinity
//go:noescape
func SetAutomationNodeSelectionEndAffinity(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_NotUserSelectableStyle
//go:noescape
func GetAutomationNodeNotUserSelectableStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_NotUserSelectableStyle
//go:noescape
func SetAutomationNodeNotUserSelectableStyle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ValueForRange
//go:noescape
func GetAutomationNodeValueForRange(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ValueForRange
//go:noescape
func SetAutomationNodeValueForRange(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_MinValueForRange
//go:noescape
func GetAutomationNodeMinValueForRange(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_MinValueForRange
//go:noescape
func SetAutomationNodeMinValueForRange(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_MaxValueForRange
//go:noescape
func GetAutomationNodeMaxValueForRange(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_MaxValueForRange
//go:noescape
func SetAutomationNodeMaxValueForRange(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_PosInSet
//go:noescape
func GetAutomationNodePosInSet(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_PosInSet
//go:noescape
func SetAutomationNodePosInSet(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SetSize
//go:noescape
func GetAutomationNodeSetSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SetSize
//go:noescape
func SetAutomationNodeSetSize(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableRowCount
//go:noescape
func GetAutomationNodeTableRowCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableRowCount
//go:noescape
func SetAutomationNodeTableRowCount(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AriaRowCount
//go:noescape
func GetAutomationNodeAriaRowCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AriaRowCount
//go:noescape
func SetAutomationNodeAriaRowCount(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableColumnCount
//go:noescape
func GetAutomationNodeTableColumnCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableColumnCount
//go:noescape
func SetAutomationNodeTableColumnCount(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AriaColumnCount
//go:noescape
func GetAutomationNodeAriaColumnCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AriaColumnCount
//go:noescape
func SetAutomationNodeAriaColumnCount(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableCellColumnIndex
//go:noescape
func GetAutomationNodeTableCellColumnIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableCellColumnIndex
//go:noescape
func SetAutomationNodeTableCellColumnIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableCellAriaColumnIndex
//go:noescape
func GetAutomationNodeTableCellAriaColumnIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableCellAriaColumnIndex
//go:noescape
func SetAutomationNodeTableCellAriaColumnIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableCellColumnSpan
//go:noescape
func GetAutomationNodeTableCellColumnSpan(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableCellColumnSpan
//go:noescape
func SetAutomationNodeTableCellColumnSpan(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableCellRowIndex
//go:noescape
func GetAutomationNodeTableCellRowIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableCellRowIndex
//go:noescape
func SetAutomationNodeTableCellRowIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableCellAriaRowIndex
//go:noescape
func GetAutomationNodeTableCellAriaRowIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableCellAriaRowIndex
//go:noescape
func SetAutomationNodeTableCellAriaRowIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableCellRowSpan
//go:noescape
func GetAutomationNodeTableCellRowSpan(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableCellRowSpan
//go:noescape
func SetAutomationNodeTableCellRowSpan(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableColumnHeader
//go:noescape
func GetAutomationNodeTableColumnHeader(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableColumnHeader
//go:noescape
func SetAutomationNodeTableColumnHeader(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableRowHeader
//go:noescape
func GetAutomationNodeTableRowHeader(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableRowHeader
//go:noescape
func SetAutomationNodeTableRowHeader(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableColumnIndex
//go:noescape
func GetAutomationNodeTableColumnIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableColumnIndex
//go:noescape
func SetAutomationNodeTableColumnIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_TableRowIndex
//go:noescape
func GetAutomationNodeTableRowIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_TableRowIndex
//go:noescape
func SetAutomationNodeTableRowIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_LiveStatus
//go:noescape
func GetAutomationNodeLiveStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_LiveStatus
//go:noescape
func SetAutomationNodeLiveStatus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_LiveRelevant
//go:noescape
func GetAutomationNodeLiveRelevant(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_LiveRelevant
//go:noescape
func SetAutomationNodeLiveRelevant(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_LiveAtomic
//go:noescape
func GetAutomationNodeLiveAtomic(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_LiveAtomic
//go:noescape
func SetAutomationNodeLiveAtomic(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Busy
//go:noescape
func GetAutomationNodeBusy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Busy
//go:noescape
func SetAutomationNodeBusy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ContainerLiveStatus
//go:noescape
func GetAutomationNodeContainerLiveStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ContainerLiveStatus
//go:noescape
func SetAutomationNodeContainerLiveStatus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ContainerLiveRelevant
//go:noescape
func GetAutomationNodeContainerLiveRelevant(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ContainerLiveRelevant
//go:noescape
func SetAutomationNodeContainerLiveRelevant(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ContainerLiveAtomic
//go:noescape
func GetAutomationNodeContainerLiveAtomic(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ContainerLiveAtomic
//go:noescape
func SetAutomationNodeContainerLiveAtomic(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ContainerLiveBusy
//go:noescape
func GetAutomationNodeContainerLiveBusy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ContainerLiveBusy
//go:noescape
func SetAutomationNodeContainerLiveBusy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_IsButton
//go:noescape
func GetAutomationNodeIsButton(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_IsButton
//go:noescape
func SetAutomationNodeIsButton(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_IsCheckBox
//go:noescape
func GetAutomationNodeIsCheckBox(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_IsCheckBox
//go:noescape
func SetAutomationNodeIsCheckBox(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_IsComboBox
//go:noescape
func GetAutomationNodeIsComboBox(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_IsComboBox
//go:noescape
func SetAutomationNodeIsComboBox(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_IsImage
//go:noescape
func GetAutomationNodeIsImage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_IsImage
//go:noescape
func SetAutomationNodeIsImage(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_HasHiddenOffscreenNodes
//go:noescape
func GetAutomationNodeHasHiddenOffscreenNodes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_HasHiddenOffscreenNodes
//go:noescape
func SetAutomationNodeHasHiddenOffscreenNodes(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AutoComplete
//go:noescape
func GetAutomationNodeAutoComplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AutoComplete
//go:noescape
func SetAutomationNodeAutoComplete(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ClassName
//go:noescape
func GetAutomationNodeClassName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ClassName
//go:noescape
func SetAutomationNodeClassName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Modal
//go:noescape
func GetAutomationNodeModal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Modal
//go:noescape
func SetAutomationNodeModal(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_HtmlAttributes
//go:noescape
func GetAutomationNodeHtmlAttributes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_HtmlAttributes
//go:noescape
func SetAutomationNodeHtmlAttributes(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_InputType
//go:noescape
func GetAutomationNodeInputType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_InputType
//go:noescape
func SetAutomationNodeInputType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AccessKey
//go:noescape
func GetAutomationNodeAccessKey(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AccessKey
//go:noescape
func SetAutomationNodeAccessKey(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AriaInvalidValue
//go:noescape
func GetAutomationNodeAriaInvalidValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AriaInvalidValue
//go:noescape
func SetAutomationNodeAriaInvalidValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Display
//go:noescape
func GetAutomationNodeDisplay(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Display
//go:noescape
func SetAutomationNodeDisplay(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ImageDataUrl
//go:noescape
func GetAutomationNodeImageDataUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ImageDataUrl
//go:noescape
func SetAutomationNodeImageDataUrl(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Language
//go:noescape
func GetAutomationNodeLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Language
//go:noescape
func SetAutomationNodeLanguage(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_DetectedLanguage
//go:noescape
func GetAutomationNodeDetectedLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_DetectedLanguage
//go:noescape
func SetAutomationNodeDetectedLanguage(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_HasPopup
//go:noescape
func GetAutomationNodeHasPopup(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_HasPopup
//go:noescape
func SetAutomationNodeHasPopup(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Restriction
//go:noescape
func GetAutomationNodeRestriction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Restriction
//go:noescape
func SetAutomationNodeRestriction(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Checked
//go:noescape
func GetAutomationNodeChecked(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Checked
//go:noescape
func SetAutomationNodeChecked(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_InnerHtml
//go:noescape
func GetAutomationNodeInnerHtml(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_InnerHtml
//go:noescape
func SetAutomationNodeInnerHtml(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Color
//go:noescape
func GetAutomationNodeColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Color
//go:noescape
func SetAutomationNodeColor(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_BackgroundColor
//go:noescape
func GetAutomationNodeBackgroundColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_BackgroundColor
//go:noescape
func SetAutomationNodeBackgroundColor(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_ColorValue
//go:noescape
func GetAutomationNodeColorValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_ColorValue
//go:noescape
func SetAutomationNodeColorValue(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Subscript
//go:noescape
func GetAutomationNodeSubscript(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Subscript
//go:noescape
func SetAutomationNodeSubscript(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Superscript
//go:noescape
func GetAutomationNodeSuperscript(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Superscript
//go:noescape
func SetAutomationNodeSuperscript(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Bold
//go:noescape
func GetAutomationNodeBold(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Bold
//go:noescape
func SetAutomationNodeBold(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Italic
//go:noescape
func GetAutomationNodeItalic(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Italic
//go:noescape
func SetAutomationNodeItalic(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Underline
//go:noescape
func GetAutomationNodeUnderline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Underline
//go:noescape
func SetAutomationNodeUnderline(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_LineThrough
//go:noescape
func GetAutomationNodeLineThrough(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_LineThrough
//go:noescape
func SetAutomationNodeLineThrough(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Selected
//go:noescape
func GetAutomationNodeSelected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Selected
//go:noescape
func SetAutomationNodeSelected(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_FontSize
//go:noescape
func GetAutomationNodeFontSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_FontSize
//go:noescape
func SetAutomationNodeFontSize(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_FontFamily
//go:noescape
func GetAutomationNodeFontFamily(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_FontFamily
//go:noescape
func SetAutomationNodeFontFamily(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_NonAtomicTextFieldRoot
//go:noescape
func GetAutomationNodeNonAtomicTextFieldRoot(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_NonAtomicTextFieldRoot
//go:noescape
func SetAutomationNodeNonAtomicTextFieldRoot(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AriaCurrentState
//go:noescape
func GetAutomationNodeAriaCurrentState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AriaCurrentState
//go:noescape
func SetAutomationNodeAriaCurrentState(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_InvalidState
//go:noescape
func GetAutomationNodeInvalidState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_InvalidState
//go:noescape
func SetAutomationNodeInvalidState(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_AppId
//go:noescape
func GetAutomationNodeAppId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_AppId
//go:noescape
func SetAutomationNodeAppId(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Children
//go:noescape
func GetAutomationNodeChildren(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Children
//go:noescape
func SetAutomationNodeChildren(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Parent
//go:noescape
func GetAutomationNodeParent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Parent
//go:noescape
func SetAutomationNodeParent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_FirstChild
//go:noescape
func GetAutomationNodeFirstChild(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_FirstChild
//go:noescape
func SetAutomationNodeFirstChild(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_LastChild
//go:noescape
func GetAutomationNodeLastChild(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_LastChild
//go:noescape
func SetAutomationNodeLastChild(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_PreviousSibling
//go:noescape
func GetAutomationNodePreviousSibling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_PreviousSibling
//go:noescape
func SetAutomationNodePreviousSibling(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_NextSibling
//go:noescape
func GetAutomationNodeNextSibling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_NextSibling
//go:noescape
func SetAutomationNodeNextSibling(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_PreviousOnLine
//go:noescape
func GetAutomationNodePreviousOnLine(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_PreviousOnLine
//go:noescape
func SetAutomationNodePreviousOnLine(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_NextOnLine
//go:noescape
func GetAutomationNodeNextOnLine(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_NextOnLine
//go:noescape
func SetAutomationNodeNextOnLine(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_PreviousFocus
//go:noescape
func GetAutomationNodePreviousFocus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_PreviousFocus
//go:noescape
func SetAutomationNodePreviousFocus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_NextFocus
//go:noescape
func GetAutomationNodeNextFocus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_NextFocus
//go:noescape
func SetAutomationNodeNextFocus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_PreviousWindowFocus
//go:noescape
func GetAutomationNodePreviousWindowFocus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_PreviousWindowFocus
//go:noescape
func SetAutomationNodePreviousWindowFocus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_NextWindowFocus
//go:noescape
func GetAutomationNodeNextWindowFocus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_NextWindowFocus
//go:noescape
func SetAutomationNodeNextWindowFocus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_IndexInParent
//go:noescape
func GetAutomationNodeIndexInParent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_IndexInParent
//go:noescape
func SetAutomationNodeIndexInParent(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_SortDirection
//go:noescape
func GetAutomationNodeSortDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_SortDirection
//go:noescape
func SetAutomationNodeSortDirection(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/webext/automation get_AutomationNode_Clickable
//go:noescape
func GetAutomationNodeClickable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation set_AutomationNode_Clickable
//go:noescape
func SetAutomationNodeClickable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/automation has_AutomationNode_BoundsForRange
//go:noescape
func HasFuncAutomationNodeBoundsForRange(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_BoundsForRange
//go:noescape
func FuncAutomationNodeBoundsForRange(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_BoundsForRange
//go:noescape
func CallAutomationNodeBoundsForRange(
	this js.Ref, retPtr unsafe.Pointer,
	startIndex int32,
	endIndex int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_BoundsForRange
//go:noescape
func TryAutomationNodeBoundsForRange(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	startIndex int32,
	endIndex int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_UnclippedBoundsForRange
//go:noescape
func HasFuncAutomationNodeUnclippedBoundsForRange(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_UnclippedBoundsForRange
//go:noescape
func FuncAutomationNodeUnclippedBoundsForRange(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_UnclippedBoundsForRange
//go:noescape
func CallAutomationNodeUnclippedBoundsForRange(
	this js.Ref, retPtr unsafe.Pointer,
	startIndex int32,
	endIndex int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_UnclippedBoundsForRange
//go:noescape
func TryAutomationNodeUnclippedBoundsForRange(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	startIndex int32,
	endIndex int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_DoDefault
//go:noescape
func HasFuncAutomationNodeDoDefault(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_DoDefault
//go:noescape
func FuncAutomationNodeDoDefault(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_DoDefault
//go:noescape
func CallAutomationNodeDoDefault(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_DoDefault
//go:noescape
func TryAutomationNodeDoDefault(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_Focus
//go:noescape
func HasFuncAutomationNodeFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_Focus
//go:noescape
func FuncAutomationNodeFocus(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_Focus
//go:noescape
func CallAutomationNodeFocus(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_Focus
//go:noescape
func TryAutomationNodeFocus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_GetImageData
//go:noescape
func HasFuncAutomationNodeGetImageData(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_GetImageData
//go:noescape
func FuncAutomationNodeGetImageData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_GetImageData
//go:noescape
func CallAutomationNodeGetImageData(
	this js.Ref, retPtr unsafe.Pointer,
	maxWidth int32,
	maxHeight int32)

//go:wasmimport plat/js/webext/automation try_AutomationNode_GetImageData
//go:noescape
func TryAutomationNodeGetImageData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	maxWidth int32,
	maxHeight int32) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_HitTest
//go:noescape
func HasFuncAutomationNodeHitTest(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_HitTest
//go:noescape
func FuncAutomationNodeHitTest(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_HitTest
//go:noescape
func CallAutomationNodeHitTest(
	this js.Ref, retPtr unsafe.Pointer,
	x int32,
	y int32,
	eventToFire uint32)

//go:wasmimport plat/js/webext/automation try_AutomationNode_HitTest
//go:noescape
func TryAutomationNodeHitTest(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x int32,
	y int32,
	eventToFire uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_HitTestWithReply
//go:noescape
func HasFuncAutomationNodeHitTestWithReply(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_HitTestWithReply
//go:noescape
func FuncAutomationNodeHitTestWithReply(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_HitTestWithReply
//go:noescape
func CallAutomationNodeHitTestWithReply(
	this js.Ref, retPtr unsafe.Pointer,
	x int32,
	y int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_HitTestWithReply
//go:noescape
func TryAutomationNodeHitTestWithReply(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x int32,
	y int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_MakeVisible
//go:noescape
func HasFuncAutomationNodeMakeVisible(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_MakeVisible
//go:noescape
func FuncAutomationNodeMakeVisible(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_MakeVisible
//go:noescape
func CallAutomationNodeMakeVisible(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_MakeVisible
//go:noescape
func TryAutomationNodeMakeVisible(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_PerformCustomAction
//go:noescape
func HasFuncAutomationNodePerformCustomAction(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_PerformCustomAction
//go:noescape
func FuncAutomationNodePerformCustomAction(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_PerformCustomAction
//go:noescape
func CallAutomationNodePerformCustomAction(
	this js.Ref, retPtr unsafe.Pointer,
	customActionId int32)

//go:wasmimport plat/js/webext/automation try_AutomationNode_PerformCustomAction
//go:noescape
func TryAutomationNodePerformCustomAction(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	customActionId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_PerformStandardAction
//go:noescape
func HasFuncAutomationNodePerformStandardAction(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_PerformStandardAction
//go:noescape
func FuncAutomationNodePerformStandardAction(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_PerformStandardAction
//go:noescape
func CallAutomationNodePerformStandardAction(
	this js.Ref, retPtr unsafe.Pointer,
	actionType uint32)

//go:wasmimport plat/js/webext/automation try_AutomationNode_PerformStandardAction
//go:noescape
func TryAutomationNodePerformStandardAction(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	actionType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ReplaceSelectedText
//go:noescape
func HasFuncAutomationNodeReplaceSelectedText(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ReplaceSelectedText
//go:noescape
func FuncAutomationNodeReplaceSelectedText(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ReplaceSelectedText
//go:noescape
func CallAutomationNodeReplaceSelectedText(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ReplaceSelectedText
//go:noescape
func TryAutomationNodeReplaceSelectedText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_SetAccessibilityFocus
//go:noescape
func HasFuncAutomationNodeSetAccessibilityFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_SetAccessibilityFocus
//go:noescape
func FuncAutomationNodeSetAccessibilityFocus(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_SetAccessibilityFocus
//go:noescape
func CallAutomationNodeSetAccessibilityFocus(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_SetAccessibilityFocus
//go:noescape
func TryAutomationNodeSetAccessibilityFocus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_SetSelection
//go:noescape
func HasFuncAutomationNodeSetSelection(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_SetSelection
//go:noescape
func FuncAutomationNodeSetSelection(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_SetSelection
//go:noescape
func CallAutomationNodeSetSelection(
	this js.Ref, retPtr unsafe.Pointer,
	startIndex int32,
	endIndex int32)

//go:wasmimport plat/js/webext/automation try_AutomationNode_SetSelection
//go:noescape
func TryAutomationNodeSetSelection(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	startIndex int32,
	endIndex int32) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_SetSequentialFocusNavigationStartingPoint
//go:noescape
func HasFuncAutomationNodeSetSequentialFocusNavigationStartingPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_SetSequentialFocusNavigationStartingPoint
//go:noescape
func FuncAutomationNodeSetSequentialFocusNavigationStartingPoint(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_SetSequentialFocusNavigationStartingPoint
//go:noescape
func CallAutomationNodeSetSequentialFocusNavigationStartingPoint(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_SetSequentialFocusNavigationStartingPoint
//go:noescape
func TryAutomationNodeSetSequentialFocusNavigationStartingPoint(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_SetValue
//go:noescape
func HasFuncAutomationNodeSetValue(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_SetValue
//go:noescape
func FuncAutomationNodeSetValue(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_SetValue
//go:noescape
func CallAutomationNodeSetValue(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_SetValue
//go:noescape
func TryAutomationNodeSetValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ShowContextMenu
//go:noescape
func HasFuncAutomationNodeShowContextMenu(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ShowContextMenu
//go:noescape
func FuncAutomationNodeShowContextMenu(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ShowContextMenu
//go:noescape
func CallAutomationNodeShowContextMenu(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ShowContextMenu
//go:noescape
func TryAutomationNodeShowContextMenu(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ResumeMedia
//go:noescape
func HasFuncAutomationNodeResumeMedia(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ResumeMedia
//go:noescape
func FuncAutomationNodeResumeMedia(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ResumeMedia
//go:noescape
func CallAutomationNodeResumeMedia(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ResumeMedia
//go:noescape
func TryAutomationNodeResumeMedia(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_StartDuckingMedia
//go:noescape
func HasFuncAutomationNodeStartDuckingMedia(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_StartDuckingMedia
//go:noescape
func FuncAutomationNodeStartDuckingMedia(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_StartDuckingMedia
//go:noescape
func CallAutomationNodeStartDuckingMedia(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_StartDuckingMedia
//go:noescape
func TryAutomationNodeStartDuckingMedia(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_StopDuckingMedia
//go:noescape
func HasFuncAutomationNodeStopDuckingMedia(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_StopDuckingMedia
//go:noescape
func FuncAutomationNodeStopDuckingMedia(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_StopDuckingMedia
//go:noescape
func CallAutomationNodeStopDuckingMedia(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_StopDuckingMedia
//go:noescape
func TryAutomationNodeStopDuckingMedia(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_SuspendMedia
//go:noescape
func HasFuncAutomationNodeSuspendMedia(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_SuspendMedia
//go:noescape
func FuncAutomationNodeSuspendMedia(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_SuspendMedia
//go:noescape
func CallAutomationNodeSuspendMedia(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_SuspendMedia
//go:noescape
func TryAutomationNodeSuspendMedia(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_LongClick
//go:noescape
func HasFuncAutomationNodeLongClick(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_LongClick
//go:noescape
func FuncAutomationNodeLongClick(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_LongClick
//go:noescape
func CallAutomationNodeLongClick(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_LongClick
//go:noescape
func TryAutomationNodeLongClick(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollBackward
//go:noescape
func HasFuncAutomationNodeScrollBackward(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollBackward
//go:noescape
func FuncAutomationNodeScrollBackward(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollBackward
//go:noescape
func CallAutomationNodeScrollBackward(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollBackward
//go:noescape
func TryAutomationNodeScrollBackward(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollBackward1
//go:noescape
func HasFuncAutomationNodeScrollBackward1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollBackward1
//go:noescape
func FuncAutomationNodeScrollBackward1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollBackward1
//go:noescape
func CallAutomationNodeScrollBackward1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollBackward1
//go:noescape
func TryAutomationNodeScrollBackward1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollForward
//go:noescape
func HasFuncAutomationNodeScrollForward(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollForward
//go:noescape
func FuncAutomationNodeScrollForward(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollForward
//go:noescape
func CallAutomationNodeScrollForward(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollForward
//go:noescape
func TryAutomationNodeScrollForward(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollForward1
//go:noescape
func HasFuncAutomationNodeScrollForward1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollForward1
//go:noescape
func FuncAutomationNodeScrollForward1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollForward1
//go:noescape
func CallAutomationNodeScrollForward1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollForward1
//go:noescape
func TryAutomationNodeScrollForward1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollUp
//go:noescape
func HasFuncAutomationNodeScrollUp(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollUp
//go:noescape
func FuncAutomationNodeScrollUp(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollUp
//go:noescape
func CallAutomationNodeScrollUp(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollUp
//go:noescape
func TryAutomationNodeScrollUp(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollUp1
//go:noescape
func HasFuncAutomationNodeScrollUp1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollUp1
//go:noescape
func FuncAutomationNodeScrollUp1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollUp1
//go:noescape
func CallAutomationNodeScrollUp1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollUp1
//go:noescape
func TryAutomationNodeScrollUp1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollDown
//go:noescape
func HasFuncAutomationNodeScrollDown(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollDown
//go:noescape
func FuncAutomationNodeScrollDown(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollDown
//go:noescape
func CallAutomationNodeScrollDown(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollDown
//go:noescape
func TryAutomationNodeScrollDown(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollDown1
//go:noescape
func HasFuncAutomationNodeScrollDown1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollDown1
//go:noescape
func FuncAutomationNodeScrollDown1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollDown1
//go:noescape
func CallAutomationNodeScrollDown1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollDown1
//go:noescape
func TryAutomationNodeScrollDown1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollLeft
//go:noescape
func HasFuncAutomationNodeScrollLeft(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollLeft
//go:noescape
func FuncAutomationNodeScrollLeft(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollLeft
//go:noescape
func CallAutomationNodeScrollLeft(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollLeft
//go:noescape
func TryAutomationNodeScrollLeft(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollLeft1
//go:noescape
func HasFuncAutomationNodeScrollLeft1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollLeft1
//go:noescape
func FuncAutomationNodeScrollLeft1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollLeft1
//go:noescape
func CallAutomationNodeScrollLeft1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollLeft1
//go:noescape
func TryAutomationNodeScrollLeft1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollRight
//go:noescape
func HasFuncAutomationNodeScrollRight(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollRight
//go:noescape
func FuncAutomationNodeScrollRight(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollRight
//go:noescape
func CallAutomationNodeScrollRight(
	this js.Ref, retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollRight
//go:noescape
func TryAutomationNodeScrollRight(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollRight1
//go:noescape
func HasFuncAutomationNodeScrollRight1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollRight1
//go:noescape
func FuncAutomationNodeScrollRight1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollRight1
//go:noescape
func CallAutomationNodeScrollRight1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollRight1
//go:noescape
func TryAutomationNodeScrollRight1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_ScrollToPoint
//go:noescape
func HasFuncAutomationNodeScrollToPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_ScrollToPoint
//go:noescape
func FuncAutomationNodeScrollToPoint(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_ScrollToPoint
//go:noescape
func CallAutomationNodeScrollToPoint(
	this js.Ref, retPtr unsafe.Pointer,
	x int32,
	y int32)

//go:wasmimport plat/js/webext/automation try_AutomationNode_ScrollToPoint
//go:noescape
func TryAutomationNodeScrollToPoint(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x int32,
	y int32) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_SetScrollOffset
//go:noescape
func HasFuncAutomationNodeSetScrollOffset(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_SetScrollOffset
//go:noescape
func FuncAutomationNodeSetScrollOffset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_SetScrollOffset
//go:noescape
func CallAutomationNodeSetScrollOffset(
	this js.Ref, retPtr unsafe.Pointer,
	x int32,
	y int32)

//go:wasmimport plat/js/webext/automation try_AutomationNode_SetScrollOffset
//go:noescape
func TryAutomationNodeSetScrollOffset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x int32,
	y int32) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_AddEventListener
//go:noescape
func HasFuncAutomationNodeAddEventListener(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_AddEventListener
//go:noescape
func FuncAutomationNodeAddEventListener(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_AddEventListener
//go:noescape
func CallAutomationNodeAddEventListener(
	this js.Ref, retPtr unsafe.Pointer,
	eventType uint32,
	listener js.Ref,
	capture js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_AddEventListener
//go:noescape
func TryAutomationNodeAddEventListener(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventType uint32,
	listener js.Ref,
	capture js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_RemoveEventListener
//go:noescape
func HasFuncAutomationNodeRemoveEventListener(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_RemoveEventListener
//go:noescape
func FuncAutomationNodeRemoveEventListener(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_RemoveEventListener
//go:noescape
func CallAutomationNodeRemoveEventListener(
	this js.Ref, retPtr unsafe.Pointer,
	eventType uint32,
	listener js.Ref,
	capture js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_RemoveEventListener
//go:noescape
func TryAutomationNodeRemoveEventListener(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	eventType uint32,
	listener js.Ref,
	capture js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_Find
//go:noescape
func HasFuncAutomationNodeFind(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_Find
//go:noescape
func FuncAutomationNodeFind(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_Find
//go:noescape
func CallAutomationNodeFind(
	this js.Ref, retPtr unsafe.Pointer,
	params unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_Find
//go:noescape
func TryAutomationNodeFind(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	params unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_FindAll
//go:noescape
func HasFuncAutomationNodeFindAll(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_FindAll
//go:noescape
func FuncAutomationNodeFindAll(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_FindAll
//go:noescape
func CallAutomationNodeFindAll(
	this js.Ref, retPtr unsafe.Pointer,
	params unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_FindAll
//go:noescape
func TryAutomationNodeFindAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	params unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_Matches
//go:noescape
func HasFuncAutomationNodeMatches(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_Matches
//go:noescape
func FuncAutomationNodeMatches(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_Matches
//go:noescape
func CallAutomationNodeMatches(
	this js.Ref, retPtr unsafe.Pointer,
	params unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_AutomationNode_Matches
//go:noescape
func TryAutomationNodeMatches(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	params unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_GetNextTextMatch
//go:noescape
func HasFuncAutomationNodeGetNextTextMatch(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_GetNextTextMatch
//go:noescape
func FuncAutomationNodeGetNextTextMatch(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_GetNextTextMatch
//go:noescape
func CallAutomationNodeGetNextTextMatch(
	this js.Ref, retPtr unsafe.Pointer,
	searchStr js.Ref,
	backward js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_GetNextTextMatch
//go:noescape
func TryAutomationNodeGetNextTextMatch(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	searchStr js.Ref,
	backward js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_CreatePosition
//go:noescape
func HasFuncAutomationNodeCreatePosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_CreatePosition
//go:noescape
func FuncAutomationNodeCreatePosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_CreatePosition
//go:noescape
func CallAutomationNodeCreatePosition(
	this js.Ref, retPtr unsafe.Pointer,
	typ uint32,
	offset int32,
	isUpstream js.Ref)

//go:wasmimport plat/js/webext/automation try_AutomationNode_CreatePosition
//go:noescape
func TryAutomationNodeCreatePosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32,
	offset int32,
	isUpstream js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_AutomationNode_CreatePosition1
//go:noescape
func HasFuncAutomationNodeCreatePosition1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation func_AutomationNode_CreatePosition1
//go:noescape
func FuncAutomationNodeCreatePosition1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AutomationNode_CreatePosition1
//go:noescape
func CallAutomationNodeCreatePosition1(
	this js.Ref, retPtr unsafe.Pointer,
	typ uint32,
	offset int32)

//go:wasmimport plat/js/webext/automation try_AutomationNode_CreatePosition1
//go:noescape
func TryAutomationNodeCreatePosition1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32,
	offset int32) (ok js.Ref)

//go:wasmimport plat/js/webext/automation constof_DescriptionFromType
//go:noescape
func ConstOfDescriptionFromType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_IntentInputEventType
//go:noescape
func ConstOfIntentInputEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_MarkerType
//go:noescape
func ConstOfMarkerType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_Restriction
//go:noescape
func ConstOfRestriction(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation store_SetDocumentSelectionParams
//go:noescape
func SetDocumentSelectionParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automation load_SetDocumentSelectionParams
//go:noescape
func SetDocumentSelectionParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation constof_StateType
//go:noescape
func ConstOfStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation constof_TreeChangeType
//go:noescape
func ConstOfTreeChangeType(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation store_TreeChange
//go:noescape
func TreeChangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/automation load_TreeChange
//go:noescape
func TreeChangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/automation constof_TreeChangeObserverFilter
//go:noescape
func ConstOfTreeChangeObserverFilter(str js.Ref) uint32

//go:wasmimport plat/js/webext/automation has_AddTreeChangeObserver
//go:noescape
func HasFuncAddTreeChangeObserver() js.Ref

//go:wasmimport plat/js/webext/automation func_AddTreeChangeObserver
//go:noescape
func FuncAddTreeChangeObserver(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_AddTreeChangeObserver
//go:noescape
func CallAddTreeChangeObserver(
	retPtr unsafe.Pointer,
	filter uint32,
	observer js.Ref)

//go:wasmimport plat/js/webext/automation try_AddTreeChangeObserver
//go:noescape
func TryAddTreeChangeObserver(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter uint32,
	observer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_GetAccessibilityFocus
//go:noescape
func HasFuncGetAccessibilityFocus() js.Ref

//go:wasmimport plat/js/webext/automation func_GetAccessibilityFocus
//go:noescape
func FuncGetAccessibilityFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_GetAccessibilityFocus
//go:noescape
func CallGetAccessibilityFocus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_GetAccessibilityFocus
//go:noescape
func TryGetAccessibilityFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_GetDesktop
//go:noescape
func HasFuncGetDesktop() js.Ref

//go:wasmimport plat/js/webext/automation func_GetDesktop
//go:noescape
func FuncGetDesktop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_GetDesktop
//go:noescape
func CallGetDesktop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_GetDesktop
//go:noescape
func TryGetDesktop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_GetFocus
//go:noescape
func HasFuncGetFocus() js.Ref

//go:wasmimport plat/js/webext/automation func_GetFocus
//go:noescape
func FuncGetFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_GetFocus
//go:noescape
func CallGetFocus(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_GetFocus
//go:noescape
func TryGetFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_GetTree
//go:noescape
func HasFuncGetTree() js.Ref

//go:wasmimport plat/js/webext/automation func_GetTree
//go:noescape
func FuncGetTree(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_GetTree
//go:noescape
func CallGetTree(
	retPtr unsafe.Pointer,
	tabId int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/automation try_GetTree
//go:noescape
func TryGetTree(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_RemoveTreeChangeObserver
//go:noescape
func HasFuncRemoveTreeChangeObserver() js.Ref

//go:wasmimport plat/js/webext/automation func_RemoveTreeChangeObserver
//go:noescape
func FuncRemoveTreeChangeObserver(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_RemoveTreeChangeObserver
//go:noescape
func CallRemoveTreeChangeObserver(
	retPtr unsafe.Pointer,
	observer js.Ref)

//go:wasmimport plat/js/webext/automation try_RemoveTreeChangeObserver
//go:noescape
func TryRemoveTreeChangeObserver(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	observer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/automation has_SetDocumentSelection
//go:noescape
func HasFuncSetDocumentSelection() js.Ref

//go:wasmimport plat/js/webext/automation func_SetDocumentSelection
//go:noescape
func FuncSetDocumentSelection(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/automation call_SetDocumentSelection
//go:noescape
func CallSetDocumentSelection(
	retPtr unsafe.Pointer,
	params unsafe.Pointer)

//go:wasmimport plat/js/webext/automation try_SetDocumentSelection
//go:noescape
func TrySetDocumentSelection(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	params unsafe.Pointer) (ok js.Ref)
