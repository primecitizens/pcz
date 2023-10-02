// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web store_OptionalEffectTiming
//go:noescape

func OptionalEffectTimingJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_OptionalEffectTiming
//go:noescape

func OptionalEffectTimingJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_AnimationNodeList_Length
//go:noescape

func GetAnimationNodeListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_AnimationNodeList_Item
//go:noescape

func CallAnimationNodeListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_AnimationNodeList_Item
//go:noescape

func AnimationNodeListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_GroupEffect_GroupEffect
//go:noescape

func NewGroupEffectByGroupEffect(
	children js.Ref,
	timing js.Ref) js.Ref

//go:wasmimport plat/js/web new_GroupEffect_GroupEffect1
//go:noescape

func NewGroupEffectByGroupEffect1(
	children js.Ref) js.Ref

//go:wasmimport plat/js/web get_GroupEffect_Children
//go:noescape

func GetGroupEffectChildren(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_GroupEffect_FirstChild
//go:noescape

func GetGroupEffectFirstChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_GroupEffect_LastChild
//go:noescape

func GetGroupEffectLastChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_GroupEffect_Clone
//go:noescape

func CallGroupEffectClone(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_GroupEffect_Clone
//go:noescape

func GroupEffectCloneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GroupEffect_Prepend
//go:noescape

func CallGroupEffectPrepend(
	this js.Ref,
	ptr unsafe.Pointer,

	effectsPtr unsafe.Pointer,
	effectsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GroupEffect_Prepend
//go:noescape

func GroupEffectPrependFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GroupEffect_Append
//go:noescape

func CallGroupEffectAppend(
	this js.Ref,
	ptr unsafe.Pointer,

	effectsPtr unsafe.Pointer,
	effectsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_GroupEffect_Append
//go:noescape

func GroupEffectAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_AnimationEffect_Parent
//go:noescape

func GetAnimationEffectParent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AnimationEffect_PreviousSibling
//go:noescape

func GetAnimationEffectPreviousSibling(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AnimationEffect_NextSibling
//go:noescape

func GetAnimationEffectNextSibling(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_AnimationEffect_GetTiming
//go:noescape

func CallAnimationEffectGetTiming(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AnimationEffect_GetTiming
//go:noescape

func AnimationEffectGetTimingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AnimationEffect_GetComputedTiming
//go:noescape

func CallAnimationEffectGetComputedTiming(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AnimationEffect_GetComputedTiming
//go:noescape

func AnimationEffectGetComputedTimingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AnimationEffect_UpdateTiming
//go:noescape

func CallAnimationEffectUpdateTiming(
	this js.Ref,
	ptr unsafe.Pointer,

	timing unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_AnimationEffect_UpdateTiming
//go:noescape

func AnimationEffectUpdateTimingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AnimationEffect_UpdateTiming1
//go:noescape

func CallAnimationEffectUpdateTiming1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AnimationEffect_UpdateTiming1
//go:noescape

func AnimationEffectUpdateTiming1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AnimationEffect_Before
//go:noescape

func CallAnimationEffectBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	effectsPtr unsafe.Pointer,
	effectsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_AnimationEffect_Before
//go:noescape

func AnimationEffectBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AnimationEffect_After
//go:noescape

func CallAnimationEffectAfter(
	this js.Ref,
	ptr unsafe.Pointer,

	effectsPtr unsafe.Pointer,
	effectsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_AnimationEffect_After
//go:noescape

func AnimationEffectAfterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AnimationEffect_Replace
//go:noescape

func CallAnimationEffectReplace(
	this js.Ref,
	ptr unsafe.Pointer,

	effectsPtr unsafe.Pointer,
	effectsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_AnimationEffect_Replace
//go:noescape

func AnimationEffectReplaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AnimationEffect_Remove
//go:noescape

func CallAnimationEffectRemove(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AnimationEffect_Remove
//go:noescape

func AnimationEffectRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_AnimationTimeline_CurrentTime
//go:noescape

func GetAnimationTimelineCurrentTime(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AnimationTimeline_Duration
//go:noescape

func GetAnimationTimelineDuration(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_AnimationTimeline_Play
//go:noescape

func CallAnimationTimelinePlay(
	this js.Ref,
	ptr unsafe.Pointer,

	effect js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_AnimationTimeline_Play
//go:noescape

func AnimationTimelinePlayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AnimationTimeline_Play1
//go:noescape

func CallAnimationTimelinePlay1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AnimationTimeline_Play1
//go:noescape

func AnimationTimelinePlay1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_AnimationPlayState
//go:noescape

func ConstOfAnimationPlayState(str js.Ref) uint32

//go:wasmimport plat/js/web constof_AnimationReplaceState
//go:noescape

func ConstOfAnimationReplaceState(str js.Ref) uint32

//go:wasmimport plat/js/web new_Animation_Animation
//go:noescape

func NewAnimationByAnimation(
	effect js.Ref,
	timeline js.Ref) js.Ref

//go:wasmimport plat/js/web new_Animation_Animation1
//go:noescape

func NewAnimationByAnimation1(
	effect js.Ref) js.Ref

//go:wasmimport plat/js/web new_Animation_Animation2
//go:noescape

func NewAnimationByAnimation2() js.Ref

//go:wasmimport plat/js/web get_Animation_Id
//go:noescape

func GetAnimationId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Animation_Id
//go:noescape

func SetAnimationId(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Animation_Effect
//go:noescape

func GetAnimationEffect(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Animation_Effect
//go:noescape

func SetAnimationEffect(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Animation_Timeline
//go:noescape

func GetAnimationTimeline(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Animation_Timeline
//go:noescape

func SetAnimationTimeline(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Animation_PlaybackRate
//go:noescape

func GetAnimationPlaybackRate(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_Animation_PlaybackRate
//go:noescape

func SetAnimationPlaybackRate(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_Animation_PlayState
//go:noescape

func GetAnimationPlayState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Animation_ReplaceState
//go:noescape

func GetAnimationReplaceState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Animation_Pending
//go:noescape

func GetAnimationPending(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Animation_Ready
//go:noescape

func GetAnimationReady(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Animation_Finished
//go:noescape

func GetAnimationFinished(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Animation_StartTime
//go:noescape

func GetAnimationStartTime(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Animation_StartTime
//go:noescape

func SetAnimationStartTime(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Animation_CurrentTime
//go:noescape

func GetAnimationCurrentTime(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Animation_CurrentTime
//go:noescape

func SetAnimationCurrentTime(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_Animation_Cancel
//go:noescape

func CallAnimationCancel(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Animation_Cancel
//go:noescape

func AnimationCancelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Animation_Finish
//go:noescape

func CallAnimationFinish(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Animation_Finish
//go:noescape

func AnimationFinishFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Animation_Play
//go:noescape

func CallAnimationPlay(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Animation_Play
//go:noescape

func AnimationPlayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Animation_Pause
//go:noescape

func CallAnimationPause(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Animation_Pause
//go:noescape

func AnimationPauseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Animation_UpdatePlaybackRate
//go:noescape

func CallAnimationUpdatePlaybackRate(
	this js.Ref,
	ptr unsafe.Pointer,

	playbackRate float64,
) js.Ref

//go:wasmimport plat/js/web func_Animation_UpdatePlaybackRate
//go:noescape

func AnimationUpdatePlaybackRateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Animation_Reverse
//go:noescape

func CallAnimationReverse(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Animation_Reverse
//go:noescape

func AnimationReverseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Animation_Persist
//go:noescape

func CallAnimationPersist(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Animation_Persist
//go:noescape

func AnimationPersistFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Animation_CommitStyles
//go:noescape

func CallAnimationCommitStyles(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Animation_CommitStyles
//go:noescape

func AnimationCommitStylesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ShadowRootMode
//go:noescape

func ConstOfShadowRootMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_SlotAssignmentMode
//go:noescape

func ConstOfSlotAssignmentMode(str js.Ref) uint32

//go:wasmimport plat/js/web get_MediaList_MediaText
//go:noescape

func GetMediaListMediaText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_MediaList_MediaText
//go:noescape

func SetMediaListMediaText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaList_Length
//go:noescape

func GetMediaListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_MediaList_Item
//go:noescape

func CallMediaListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_MediaList_Item
//go:noescape

func MediaListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaList_AppendMedium
//go:noescape

func CallMediaListAppendMedium(
	this js.Ref,
	ptr unsafe.Pointer,

	medium js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaList_AppendMedium
//go:noescape

func MediaListAppendMediumFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaList_DeleteMedium
//go:noescape

func CallMediaListDeleteMedium(
	this js.Ref,
	ptr unsafe.Pointer,

	medium js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaList_DeleteMedium
//go:noescape

func MediaListDeleteMediumFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_CSSStyleSheetInit
//go:noescape

func CSSStyleSheetInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CSSStyleSheetInit
//go:noescape

func CSSStyleSheetInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSRule_CssText
//go:noescape

func GetCSSRuleCssText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_CSSRule_CssText
//go:noescape

func SetCSSRuleCssText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSRule_ParentRule
//go:noescape

func GetCSSRuleParentRule(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSRule_ParentStyleSheet
//go:noescape

func GetCSSRuleParentStyleSheet(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSRule_Type
//go:noescape

func GetCSSRuleType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_CSSRuleList_Length
//go:noescape

func GetCSSRuleListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_CSSRuleList_Item
//go:noescape

func CallCSSRuleListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSRuleList_Item
//go:noescape

func CSSRuleListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSStyleSheet_CSSStyleSheet
//go:noescape

func NewCSSStyleSheetByCSSStyleSheet(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_CSSStyleSheet_CSSStyleSheet1
//go:noescape

func NewCSSStyleSheetByCSSStyleSheet1() js.Ref

//go:wasmimport plat/js/web get_CSSStyleSheet_OwnerRule
//go:noescape

func GetCSSStyleSheetOwnerRule(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSStyleSheet_CssRules
//go:noescape

func GetCSSStyleSheetCssRules(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSStyleSheet_Rules
//go:noescape

func GetCSSStyleSheetRules(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_InsertRule
//go:noescape

func CallCSSStyleSheetInsertRule(
	this js.Ref,
	ptr unsafe.Pointer,

	rule js.Ref,
	index uint32,
) uint32

//go:wasmimport plat/js/web func_CSSStyleSheet_InsertRule
//go:noescape

func CSSStyleSheetInsertRuleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_InsertRule1
//go:noescape

func CallCSSStyleSheetInsertRule1(
	this js.Ref,
	ptr unsafe.Pointer,

	rule js.Ref,
) uint32

//go:wasmimport plat/js/web func_CSSStyleSheet_InsertRule1
//go:noescape

func CSSStyleSheetInsertRule1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_DeleteRule
//go:noescape

func CallCSSStyleSheetDeleteRule(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleSheet_DeleteRule
//go:noescape

func CSSStyleSheetDeleteRuleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_Replace
//go:noescape

func CallCSSStyleSheetReplace(
	this js.Ref,
	ptr unsafe.Pointer,

	text js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleSheet_Replace
//go:noescape

func CSSStyleSheetReplaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_ReplaceSync
//go:noescape

func CallCSSStyleSheetReplaceSync(
	this js.Ref,
	ptr unsafe.Pointer,

	text js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleSheet_ReplaceSync
//go:noescape

func CSSStyleSheetReplaceSyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_AddRule
//go:noescape

func CallCSSStyleSheetAddRule(
	this js.Ref,
	ptr unsafe.Pointer,

	selector js.Ref,
	style js.Ref,
	index uint32,
) int32

//go:wasmimport plat/js/web func_CSSStyleSheet_AddRule
//go:noescape

func CSSStyleSheetAddRuleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_AddRule1
//go:noescape

func CallCSSStyleSheetAddRule1(
	this js.Ref,
	ptr unsafe.Pointer,

	selector js.Ref,
	style js.Ref,
) int32

//go:wasmimport plat/js/web func_CSSStyleSheet_AddRule1
//go:noescape

func CSSStyleSheetAddRule1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_AddRule2
//go:noescape

func CallCSSStyleSheetAddRule2(
	this js.Ref,
	ptr unsafe.Pointer,

	selector js.Ref,
) int32

//go:wasmimport plat/js/web func_CSSStyleSheet_AddRule2
//go:noescape

func CSSStyleSheetAddRule2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_AddRule3
//go:noescape

func CallCSSStyleSheetAddRule3(
	this js.Ref,
	ptr unsafe.Pointer,

) int32

//go:wasmimport plat/js/web func_CSSStyleSheet_AddRule3
//go:noescape

func CSSStyleSheetAddRule3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_RemoveRule
//go:noescape

func CallCSSStyleSheetRemoveRule(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleSheet_RemoveRule
//go:noescape

func CSSStyleSheetRemoveRuleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleSheet_RemoveRule1
//go:noescape

func CallCSSStyleSheetRemoveRule1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CSSStyleSheet_RemoveRule1
//go:noescape

func CSSStyleSheetRemoveRule1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_StyleSheetList_Length
//go:noescape

func GetStyleSheetListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_StyleSheetList_Item
//go:noescape

func CallStyleSheetListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_StyleSheetList_Item
//go:noescape

func StyleSheetListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_ShadowRoot_Mode
//go:noescape

func GetShadowRootMode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_ShadowRoot_DelegatesFocus
//go:noescape

func GetShadowRootDelegatesFocus(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ShadowRoot_SlotAssignment
//go:noescape

func GetShadowRootSlotAssignment(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_ShadowRoot_Host
//go:noescape

func GetShadowRootHost(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ShadowRoot_InnerHTML
//go:noescape

func GetShadowRootInnerHTML(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ShadowRoot_InnerHTML
//go:noescape

func SetShadowRootInnerHTML(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ShadowRoot_FullscreenElement
//go:noescape

func GetShadowRootFullscreenElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ShadowRoot_StyleSheets
//go:noescape

func GetShadowRootStyleSheets(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ShadowRoot_AdoptedStyleSheets
//go:noescape

func GetShadowRootAdoptedStyleSheets(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_ShadowRoot_AdoptedStyleSheets
//go:noescape

func SetShadowRootAdoptedStyleSheets(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_ShadowRoot_PictureInPictureElement
//go:noescape

func GetShadowRootPictureInPictureElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ShadowRoot_ActiveElement
//go:noescape

func GetShadowRootActiveElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ShadowRoot_PointerLockElement
//go:noescape

func GetShadowRootPointerLockElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_ShadowRoot_GetAnimations
//go:noescape

func CallShadowRootGetAnimations(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ShadowRoot_GetAnimations
//go:noescape

func ShadowRootGetAnimationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_ShadowRootInit
//go:noescape

func ShadowRootInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ShadowRootInit
//go:noescape

func ShadowRootInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_FullscreenNavigationUI
//go:noescape

func ConstOfFullscreenNavigationUI(str js.Ref) uint32

//go:wasmimport plat/js/web get_ScreenDetailed_AvailLeft
//go:noescape

func GetScreenDetailedAvailLeft(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_ScreenDetailed_AvailTop
//go:noescape

func GetScreenDetailedAvailTop(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_ScreenDetailed_Left
//go:noescape

func GetScreenDetailedLeft(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_ScreenDetailed_Top
//go:noescape

func GetScreenDetailedTop(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_ScreenDetailed_IsPrimary
//go:noescape

func GetScreenDetailedIsPrimary(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ScreenDetailed_IsInternal
//go:noescape

func GetScreenDetailedIsInternal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ScreenDetailed_DevicePixelRatio
//go:noescape

func GetScreenDetailedDevicePixelRatio(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_ScreenDetailed_Label
//go:noescape

func GetScreenDetailedLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_FullscreenOptions
//go:noescape

func FullscreenOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FullscreenOptions
//go:noescape

func FullscreenOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleValue_ToString
//go:noescape

func CallCSSStyleValueToString(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CSSStyleValue_ToString
//go:noescape

func CSSStyleValueToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleValue_Parse
//go:noescape

func CallCSSStyleValueParse(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
	cssText js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleValue_Parse
//go:noescape

func CSSStyleValueParseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleValue_ParseAll
//go:noescape

func CallCSSStyleValueParseAll(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
	cssText js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleValue_ParseAll
//go:noescape

func CSSStyleValueParseAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_StylePropertyMapReadOnly_Size
//go:noescape

func GetStylePropertyMapReadOnlySize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_StylePropertyMapReadOnly_Get
//go:noescape

func CallStylePropertyMapReadOnlyGet(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMapReadOnly_Get
//go:noescape

func StylePropertyMapReadOnlyGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMapReadOnly_GetAll
//go:noescape

func CallStylePropertyMapReadOnlyGetAll(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMapReadOnly_GetAll
//go:noescape

func StylePropertyMapReadOnlyGetAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMapReadOnly_Has
//go:noescape

func CallStylePropertyMapReadOnlyHas(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMapReadOnly_Has
//go:noescape

func StylePropertyMapReadOnlyHasFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_FocusableAreaSearchMode
//go:noescape

func ConstOfFocusableAreaSearchMode(str js.Ref) uint32

//go:wasmimport plat/js/web store_FocusableAreasOption
//go:noescape

func FocusableAreasOptionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FocusableAreasOption
//go:noescape

func FocusableAreasOptionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_SpatialNavigationDirection
//go:noescape

func ConstOfSpatialNavigationDirection(str js.Ref) uint32

//go:wasmimport plat/js/web store_SpatialNavigationSearchOptions
//go:noescape

func SpatialNavigationSearchOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SpatialNavigationSearchOptions
//go:noescape

func SpatialNavigationSearchOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DOMPointInit
//go:noescape

func DOMPointInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DOMPointInit
//go:noescape

func DOMPointInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DOMRectInit
//go:noescape

func DOMRectInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DOMRectInit
//go:noescape

func DOMRectInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_DOMQuadInit
//go:noescape

func DOMQuadInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DOMQuadInit
//go:noescape

func DOMQuadInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_DOMRect_DOMRect
//go:noescape

func NewDOMRectByDOMRect(
	x float64,
	y float64,
	width float64,
	height float64) js.Ref

//go:wasmimport plat/js/web new_DOMRect_DOMRect1
//go:noescape

func NewDOMRectByDOMRect1(
	x float64,
	y float64,
	width float64) js.Ref

//go:wasmimport plat/js/web new_DOMRect_DOMRect2
//go:noescape

func NewDOMRectByDOMRect2(
	x float64,
	y float64) js.Ref

//go:wasmimport plat/js/web new_DOMRect_DOMRect3
//go:noescape

func NewDOMRectByDOMRect3(
	x float64) js.Ref

//go:wasmimport plat/js/web new_DOMRect_DOMRect4
//go:noescape

func NewDOMRectByDOMRect4() js.Ref

//go:wasmimport plat/js/web get_DOMRect_X
//go:noescape

func GetDOMRectX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMRect_X
//go:noescape

func SetDOMRectX(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMRect_Y
//go:noescape

func GetDOMRectY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMRect_Y
//go:noescape

func SetDOMRectY(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMRect_Width
//go:noescape

func GetDOMRectWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMRect_Width
//go:noescape

func SetDOMRectWidth(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMRect_Height
//go:noescape

func GetDOMRectHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMRect_Height
//go:noescape

func SetDOMRectHeight(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web call_DOMRect_FromRect
//go:noescape

func CallDOMRectFromRect(
	this js.Ref,
	ptr unsafe.Pointer,

	other unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMRect_FromRect
//go:noescape

func DOMRectFromRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMRect_FromRect1
//go:noescape

func CallDOMRectFromRect1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMRect_FromRect1
//go:noescape

func DOMRectFromRect1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_DOMPoint_DOMPoint
//go:noescape

func NewDOMPointByDOMPoint(
	x float64,
	y float64,
	z float64,
	w float64) js.Ref

//go:wasmimport plat/js/web new_DOMPoint_DOMPoint1
//go:noescape

func NewDOMPointByDOMPoint1(
	x float64,
	y float64,
	z float64) js.Ref

//go:wasmimport plat/js/web new_DOMPoint_DOMPoint2
//go:noescape

func NewDOMPointByDOMPoint2(
	x float64,
	y float64) js.Ref

//go:wasmimport plat/js/web new_DOMPoint_DOMPoint3
//go:noescape

func NewDOMPointByDOMPoint3(
	x float64) js.Ref

//go:wasmimport plat/js/web new_DOMPoint_DOMPoint4
//go:noescape

func NewDOMPointByDOMPoint4() js.Ref

//go:wasmimport plat/js/web get_DOMPoint_X
//go:noescape

func GetDOMPointX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMPoint_X
//go:noescape

func SetDOMPointX(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMPoint_Y
//go:noescape

func GetDOMPointY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMPoint_Y
//go:noescape

func SetDOMPointY(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMPoint_Z
//go:noescape

func GetDOMPointZ(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMPoint_Z
//go:noescape

func SetDOMPointZ(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMPoint_W
//go:noescape

func GetDOMPointW(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMPoint_W
//go:noescape

func SetDOMPointW(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web call_DOMPoint_FromPoint
//go:noescape

func CallDOMPointFromPoint(
	this js.Ref,
	ptr unsafe.Pointer,

	other unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMPoint_FromPoint
//go:noescape

func DOMPointFromPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPoint_FromPoint1
//go:noescape

func CallDOMPointFromPoint1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMPoint_FromPoint1
//go:noescape

func DOMPointFromPoint1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_DOMQuad_DOMQuad
//go:noescape

func NewDOMQuadByDOMQuad(
	p1 unsafe.Pointer,
	p2 unsafe.Pointer,
	p3 unsafe.Pointer,
	p4 unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DOMQuad_DOMQuad1
//go:noescape

func NewDOMQuadByDOMQuad1(
	p1 unsafe.Pointer,
	p2 unsafe.Pointer,
	p3 unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DOMQuad_DOMQuad2
//go:noescape

func NewDOMQuadByDOMQuad2(
	p1 unsafe.Pointer,
	p2 unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DOMQuad_DOMQuad3
//go:noescape

func NewDOMQuadByDOMQuad3(
	p1 unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DOMQuad_DOMQuad4
//go:noescape

func NewDOMQuadByDOMQuad4() js.Ref

//go:wasmimport plat/js/web get_DOMQuad_P1
//go:noescape

func GetDOMQuadP1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DOMQuad_P2
//go:noescape

func GetDOMQuadP2(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DOMQuad_P3
//go:noescape

func GetDOMQuadP3(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DOMQuad_P4
//go:noescape

func GetDOMQuadP4(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_DOMQuad_FromRect
//go:noescape

func CallDOMQuadFromRect(
	this js.Ref,
	ptr unsafe.Pointer,

	other unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMQuad_FromRect
//go:noescape

func DOMQuadFromRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMQuad_FromRect1
//go:noescape

func CallDOMQuadFromRect1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMQuad_FromRect1
//go:noescape

func DOMQuadFromRect1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMQuad_FromQuad
//go:noescape

func CallDOMQuadFromQuad(
	this js.Ref,
	ptr unsafe.Pointer,

	other unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMQuad_FromQuad
//go:noescape

func DOMQuadFromQuadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMQuad_FromQuad1
//go:noescape

func CallDOMQuadFromQuad1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMQuad_FromQuad1
//go:noescape

func DOMQuadFromQuad1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMQuad_GetBounds
//go:noescape

func CallDOMQuadGetBounds(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMQuad_GetBounds
//go:noescape

func DOMQuadGetBoundsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMQuad_ToJSON
//go:noescape

func CallDOMQuadToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMQuad_ToJSON
//go:noescape

func DOMQuadToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_CSSBoxType
//go:noescape

func ConstOfCSSBoxType(str js.Ref) uint32

//go:wasmimport plat/js/web store_ConvertCoordinateOptions
//go:noescape

func ConvertCoordinateOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ConvertCoordinateOptions
//go:noescape

func ConvertCoordinateOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_DOMRectReadOnly_DOMRectReadOnly
//go:noescape

func NewDOMRectReadOnlyByDOMRectReadOnly(
	x float64,
	y float64,
	width float64,
	height float64) js.Ref

//go:wasmimport plat/js/web new_DOMRectReadOnly_DOMRectReadOnly1
//go:noescape

func NewDOMRectReadOnlyByDOMRectReadOnly1(
	x float64,
	y float64,
	width float64) js.Ref

//go:wasmimport plat/js/web new_DOMRectReadOnly_DOMRectReadOnly2
//go:noescape

func NewDOMRectReadOnlyByDOMRectReadOnly2(
	x float64,
	y float64) js.Ref

//go:wasmimport plat/js/web new_DOMRectReadOnly_DOMRectReadOnly3
//go:noescape

func NewDOMRectReadOnlyByDOMRectReadOnly3(
	x float64) js.Ref

//go:wasmimport plat/js/web new_DOMRectReadOnly_DOMRectReadOnly4
//go:noescape

func NewDOMRectReadOnlyByDOMRectReadOnly4() js.Ref

//go:wasmimport plat/js/web get_DOMRectReadOnly_X
//go:noescape

func GetDOMRectReadOnlyX(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMRectReadOnly_Y
//go:noescape

func GetDOMRectReadOnlyY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMRectReadOnly_Width
//go:noescape

func GetDOMRectReadOnlyWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMRectReadOnly_Height
//go:noescape

func GetDOMRectReadOnlyHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMRectReadOnly_Top
//go:noescape

func GetDOMRectReadOnlyTop(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMRectReadOnly_Right
//go:noescape

func GetDOMRectReadOnlyRight(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMRectReadOnly_Bottom
//go:noescape

func GetDOMRectReadOnlyBottom(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMRectReadOnly_Left
//go:noescape

func GetDOMRectReadOnlyLeft(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web call_DOMRectReadOnly_FromRect
//go:noescape

func CallDOMRectReadOnlyFromRect(
	this js.Ref,
	ptr unsafe.Pointer,

	other unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMRectReadOnly_FromRect
//go:noescape

func DOMRectReadOnlyFromRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMRectReadOnly_FromRect1
//go:noescape

func CallDOMRectReadOnlyFromRect1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMRectReadOnly_FromRect1
//go:noescape

func DOMRectReadOnlyFromRect1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMRectReadOnly_ToJSON
//go:noescape

func CallDOMRectReadOnlyToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMRectReadOnly_ToJSON
//go:noescape

func DOMRectReadOnlyToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_AssignedNodesOptions
//go:noescape

func AssignedNodesOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AssignedNodesOptions
//go:noescape

func AssignedNodesOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
