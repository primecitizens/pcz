// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package automation

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/automation/bindings"
)

type AccessibilityFocusCallbackFunc func(this js.Ref, focusedNode AutomationNode) js.Ref

func (fn AccessibilityFocusCallbackFunc) Register() js.Func[func(focusedNode AutomationNode)] {
	return js.RegisterCallback[func(focusedNode AutomationNode)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AccessibilityFocusCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		AutomationNode{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AccessibilityFocusCallback[T any] struct {
	Fn  func(arg T, this js.Ref, focusedNode AutomationNode) js.Ref
	Arg T
}

func (cb *AccessibilityFocusCallback[T]) Register() js.Func[func(focusedNode AutomationNode)] {
	return js.RegisterCallback[func(focusedNode AutomationNode)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AccessibilityFocusCallback[T]) DispatchCallback(
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

		AutomationNode{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type BoundsForRangeCallbackFunc func(this js.Ref, bounds *Rect) js.Ref

func (fn BoundsForRangeCallbackFunc) Register() js.Func[func(bounds *Rect)] {
	return js.RegisterCallback[func(bounds *Rect)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn BoundsForRangeCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Rect
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

type BoundsForRangeCallback[T any] struct {
	Fn  func(arg T, this js.Ref, bounds *Rect) js.Ref
	Arg T
}

func (cb *BoundsForRangeCallback[T]) Register() js.Func[func(bounds *Rect)] {
	return js.RegisterCallback[func(bounds *Rect)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *BoundsForRangeCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 Rect
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

type Rect struct {
	// Left is "Rect.left"
	//
	// Optional
	//
	// NOTE: FFI_USE_Left MUST be set to true to make this field effective.
	Left int32
	// Top is "Rect.top"
	//
	// Optional
	//
	// NOTE: FFI_USE_Top MUST be set to true to make this field effective.
	Top int32
	// Width is "Rect.width"
	//
	// Optional
	//
	// NOTE: FFI_USE_Width MUST be set to true to make this field effective.
	Width int32
	// Height is "Rect.height"
	//
	// Optional
	//
	// NOTE: FFI_USE_Height MUST be set to true to make this field effective.
	Height int32

	FFI_USE_Left   bool // for Left.
	FFI_USE_Top    bool // for Top.
	FFI_USE_Width  bool // for Width.
	FFI_USE_Height bool // for Height.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Rect with all fields set.
func (p Rect) FromRef(ref js.Ref) Rect {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Rect in the application heap.
func (p Rect) New() js.Ref {
	return bindings.RectJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Rect) UpdateFrom(ref js.Ref) {
	bindings.RectJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Rect) Update(ref js.Ref) {
	bindings.RectJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Rect) FreeMembers(recursive bool) {
}

type EventType uint32

const (
	_ EventType = iota

	EventType_ACCESS_KEY_CHANGED
	EventType_ACTIVE_DESCENDANT_CHANGED
	EventType_ALERT
	EventType_ARIA_ATTRIBUTE_CHANGED_DEPRECATED
	EventType_ARIA_CURRENT_CHANGED
	EventType_ATOMIC_CHANGED
	EventType_AUTO_COMPLETE_CHANGED
	EventType_AUTOCORRECTION_OCCURED
	EventType_AUTOFILL_AVAILABILITY_CHANGED
	EventType_BLUR
	EventType_BUSY_CHANGED
	EventType_CARET_BOUNDS_CHANGED
	EventType_CHECKED_STATE_CHANGED
	EventType_CHECKED_STATE_DESCRIPTION_CHANGED
	EventType_CHILDREN_CHANGED
	EventType_CLASS_NAME_CHANGED
	EventType_CLICKED
	EventType_COLLAPSED
	EventType_CONTROLS_CHANGED
	EventType_DETAILS_CHANGED
	EventType_DESCRIBED_BY_CHANGED
	EventType_DESCRIPTION_CHANGED
	EventType_DOCUMENT_SELECTION_CHANGED
	EventType_DOCUMENT_TITLE_CHANGED
	EventType_DROPEFFECT_CHANGED
	EventType_EDITABLE_TEXT_CHANGED
	EventType_ENABLED_CHANGED
	EventType_END_OF_TEST
	EventType_EXPANDED
	EventType_EXPANDED_CHANGED
	EventType_FLOW_FROM_CHANGED
	EventType_FLOW_TO_CHANGED
	EventType_FOCUS
	EventType_FOCUS_AFTER_MENU_CLOSE
	EventType_FOCUS_CHANGED
	EventType_FOCUS_CONTEXT
	EventType_GRABBED_CHANGED
	EventType_HASPOPUP_CHANGED
	EventType_HIDE
	EventType_HIERARCHICAL_LEVEL_CHANGED
	EventType_HIT_TEST_RESULT
	EventType_HOVER
	EventType_IGNORED_CHANGED
	EventType_IMAGE_ANNOTATION_CHANGED
	EventType_IMAGE_FRAME_UPDATED
	EventType_INVALID_STATUS_CHANGED
	EventType_KEY_SHORTCUTS_CHANGED
	EventType_LABELED_BY_CHANGED
	EventType_LANGUAGE_CHANGED
	EventType_LAYOUT_COMPLETE
	EventType_LAYOUT_INVALIDATED
	EventType_LIVE_REGION_CHANGED
	EventType_LIVE_REGION_CREATED
	EventType_LIVE_REGION_NODE_CHANGED
	EventType_LIVE_RELEVANT_CHANGED
	EventType_LIVE_STATUS_CHANGED
	EventType_LOAD_COMPLETE
	EventType_LOAD_START
	EventType_LOCATION_CHANGED
	EventType_MEDIA_STARTED_PLAYING
	EventType_MEDIA_STOPPED_PLAYING
	EventType_MENU_END
	EventType_MENU_ITEM_SELECTED
	EventType_MENU_LIST_VALUE_CHANGED
	EventType_MENU_POPUP_END
	EventType_MENU_POPUP_START
	EventType_MENU_START
	EventType_MOUSE_CANCELED
	EventType_MOUSE_DRAGGED
	EventType_MOUSE_MOVED
	EventType_MOUSE_PRESSED
	EventType_MOUSE_RELEASED
	EventType_MULTILINE_STATE_CHANGED
	EventType_MULTISELECTABLE_STATE_CHANGED
	EventType_NAME_CHANGED
	EventType_OBJECT_ATTRIBUTE_CHANGED
	EventType_ORIENTATION_CHANGED
	EventType_OTHER_ATTRIBUTE_CHANGED
	EventType_PARENT_CHANGED
	EventType_PLACEHOLDER_CHANGED
	EventType_PORTAL_ACTIVATED
	EventType_POSITION_IN_SET_CHANGED
	EventType_RANGE_VALUE_CHANGED
	EventType_RANGE_VALUE_MAX_CHANGED
	EventType_RANGE_VALUE_MIN_CHANGED
	EventType_RANGE_VALUE_STEP_CHANGED
	EventType_READONLY_CHANGED
	EventType_RELATED_NODE_CHANGED
	EventType_REQUIRED_STATE_CHANGED
	EventType_ROLE_CHANGED
	EventType_ROW_COLLAPSED
	EventType_ROW_COUNT_CHANGED
	EventType_ROW_EXPANDED
	EventType_SCROLL_HORIZONTAL_POSITION_CHANGED
	EventType_SCROLL_POSITION_CHANGED
	EventType_SCROLL_VERTICAL_POSITION_CHANGED
	EventType_SCROLLED_TO_ANCHOR
	EventType_SELECTED_CHANGED
	EventType_SELECTED_CHILDREN_CHANGED
	EventType_SELECTED_VALUE_CHANGED
	EventType_SELECTION
	EventType_SELECTION_ADD
	EventType_SELECTION_REMOVE
	EventType_SET_SIZE_CHANGED
	EventType_SHOW
	EventType_SORT_CHANGED
	EventType_STATE_CHANGED
	EventType_SUBTREE_CREATED
	EventType_TEXT_ATTRIBUTE_CHANGED
	EventType_TEXT_SELECTION_CHANGED
	EventType_TEXT_CHANGED
	EventType_TOOLTIP_CLOSED
	EventType_TOOLTIP_OPENED
	EventType_TREE_CHANGED
	EventType_VALUE_IN_TEXT_FIELD_CHANGED
	EventType_VALUE_CHANGED
	EventType_WINDOW_ACTIVATED
	EventType_WINDOW_DEACTIVATED
	EventType_WINDOW_VISIBILITY_CHANGED
)

func (EventType) FromRef(str js.Ref) EventType {
	return EventType(bindings.ConstOfEventType(str))
}

func (x EventType) String() (string, bool) {
	switch x {
	case EventType_ACCESS_KEY_CHANGED:
		return "accessKeyChanged", true
	case EventType_ACTIVE_DESCENDANT_CHANGED:
		return "activeDescendantChanged", true
	case EventType_ALERT:
		return "alert", true
	case EventType_ARIA_ATTRIBUTE_CHANGED_DEPRECATED:
		return "ariaAttributeChangedDeprecated", true
	case EventType_ARIA_CURRENT_CHANGED:
		return "ariaCurrentChanged", true
	case EventType_ATOMIC_CHANGED:
		return "atomicChanged", true
	case EventType_AUTO_COMPLETE_CHANGED:
		return "autoCompleteChanged", true
	case EventType_AUTOCORRECTION_OCCURED:
		return "autocorrectionOccured", true
	case EventType_AUTOFILL_AVAILABILITY_CHANGED:
		return "autofillAvailabilityChanged", true
	case EventType_BLUR:
		return "blur", true
	case EventType_BUSY_CHANGED:
		return "busyChanged", true
	case EventType_CARET_BOUNDS_CHANGED:
		return "caretBoundsChanged", true
	case EventType_CHECKED_STATE_CHANGED:
		return "checkedStateChanged", true
	case EventType_CHECKED_STATE_DESCRIPTION_CHANGED:
		return "checkedStateDescriptionChanged", true
	case EventType_CHILDREN_CHANGED:
		return "childrenChanged", true
	case EventType_CLASS_NAME_CHANGED:
		return "classNameChanged", true
	case EventType_CLICKED:
		return "clicked", true
	case EventType_COLLAPSED:
		return "collapsed", true
	case EventType_CONTROLS_CHANGED:
		return "controlsChanged", true
	case EventType_DETAILS_CHANGED:
		return "detailsChanged", true
	case EventType_DESCRIBED_BY_CHANGED:
		return "describedByChanged", true
	case EventType_DESCRIPTION_CHANGED:
		return "descriptionChanged", true
	case EventType_DOCUMENT_SELECTION_CHANGED:
		return "documentSelectionChanged", true
	case EventType_DOCUMENT_TITLE_CHANGED:
		return "documentTitleChanged", true
	case EventType_DROPEFFECT_CHANGED:
		return "dropeffectChanged", true
	case EventType_EDITABLE_TEXT_CHANGED:
		return "editableTextChanged", true
	case EventType_ENABLED_CHANGED:
		return "enabledChanged", true
	case EventType_END_OF_TEST:
		return "endOfTest", true
	case EventType_EXPANDED:
		return "expanded", true
	case EventType_EXPANDED_CHANGED:
		return "expandedChanged", true
	case EventType_FLOW_FROM_CHANGED:
		return "flowFromChanged", true
	case EventType_FLOW_TO_CHANGED:
		return "flowToChanged", true
	case EventType_FOCUS:
		return "focus", true
	case EventType_FOCUS_AFTER_MENU_CLOSE:
		return "focusAfterMenuClose", true
	case EventType_FOCUS_CHANGED:
		return "focusChanged", true
	case EventType_FOCUS_CONTEXT:
		return "focusContext", true
	case EventType_GRABBED_CHANGED:
		return "grabbedChanged", true
	case EventType_HASPOPUP_CHANGED:
		return "haspopupChanged", true
	case EventType_HIDE:
		return "hide", true
	case EventType_HIERARCHICAL_LEVEL_CHANGED:
		return "hierarchicalLevelChanged", true
	case EventType_HIT_TEST_RESULT:
		return "hitTestResult", true
	case EventType_HOVER:
		return "hover", true
	case EventType_IGNORED_CHANGED:
		return "ignoredChanged", true
	case EventType_IMAGE_ANNOTATION_CHANGED:
		return "imageAnnotationChanged", true
	case EventType_IMAGE_FRAME_UPDATED:
		return "imageFrameUpdated", true
	case EventType_INVALID_STATUS_CHANGED:
		return "invalidStatusChanged", true
	case EventType_KEY_SHORTCUTS_CHANGED:
		return "keyShortcutsChanged", true
	case EventType_LABELED_BY_CHANGED:
		return "labeledByChanged", true
	case EventType_LANGUAGE_CHANGED:
		return "languageChanged", true
	case EventType_LAYOUT_COMPLETE:
		return "layoutComplete", true
	case EventType_LAYOUT_INVALIDATED:
		return "layoutInvalidated", true
	case EventType_LIVE_REGION_CHANGED:
		return "liveRegionChanged", true
	case EventType_LIVE_REGION_CREATED:
		return "liveRegionCreated", true
	case EventType_LIVE_REGION_NODE_CHANGED:
		return "liveRegionNodeChanged", true
	case EventType_LIVE_RELEVANT_CHANGED:
		return "liveRelevantChanged", true
	case EventType_LIVE_STATUS_CHANGED:
		return "liveStatusChanged", true
	case EventType_LOAD_COMPLETE:
		return "loadComplete", true
	case EventType_LOAD_START:
		return "loadStart", true
	case EventType_LOCATION_CHANGED:
		return "locationChanged", true
	case EventType_MEDIA_STARTED_PLAYING:
		return "mediaStartedPlaying", true
	case EventType_MEDIA_STOPPED_PLAYING:
		return "mediaStoppedPlaying", true
	case EventType_MENU_END:
		return "menuEnd", true
	case EventType_MENU_ITEM_SELECTED:
		return "menuItemSelected", true
	case EventType_MENU_LIST_VALUE_CHANGED:
		return "menuListValueChanged", true
	case EventType_MENU_POPUP_END:
		return "menuPopupEnd", true
	case EventType_MENU_POPUP_START:
		return "menuPopupStart", true
	case EventType_MENU_START:
		return "menuStart", true
	case EventType_MOUSE_CANCELED:
		return "mouseCanceled", true
	case EventType_MOUSE_DRAGGED:
		return "mouseDragged", true
	case EventType_MOUSE_MOVED:
		return "mouseMoved", true
	case EventType_MOUSE_PRESSED:
		return "mousePressed", true
	case EventType_MOUSE_RELEASED:
		return "mouseReleased", true
	case EventType_MULTILINE_STATE_CHANGED:
		return "multilineStateChanged", true
	case EventType_MULTISELECTABLE_STATE_CHANGED:
		return "multiselectableStateChanged", true
	case EventType_NAME_CHANGED:
		return "nameChanged", true
	case EventType_OBJECT_ATTRIBUTE_CHANGED:
		return "objectAttributeChanged", true
	case EventType_ORIENTATION_CHANGED:
		return "orientationChanged", true
	case EventType_OTHER_ATTRIBUTE_CHANGED:
		return "otherAttributeChanged", true
	case EventType_PARENT_CHANGED:
		return "parentChanged", true
	case EventType_PLACEHOLDER_CHANGED:
		return "placeholderChanged", true
	case EventType_PORTAL_ACTIVATED:
		return "portalActivated", true
	case EventType_POSITION_IN_SET_CHANGED:
		return "positionInSetChanged", true
	case EventType_RANGE_VALUE_CHANGED:
		return "rangeValueChanged", true
	case EventType_RANGE_VALUE_MAX_CHANGED:
		return "rangeValueMaxChanged", true
	case EventType_RANGE_VALUE_MIN_CHANGED:
		return "rangeValueMinChanged", true
	case EventType_RANGE_VALUE_STEP_CHANGED:
		return "rangeValueStepChanged", true
	case EventType_READONLY_CHANGED:
		return "readonlyChanged", true
	case EventType_RELATED_NODE_CHANGED:
		return "relatedNodeChanged", true
	case EventType_REQUIRED_STATE_CHANGED:
		return "requiredStateChanged", true
	case EventType_ROLE_CHANGED:
		return "roleChanged", true
	case EventType_ROW_COLLAPSED:
		return "rowCollapsed", true
	case EventType_ROW_COUNT_CHANGED:
		return "rowCountChanged", true
	case EventType_ROW_EXPANDED:
		return "rowExpanded", true
	case EventType_SCROLL_HORIZONTAL_POSITION_CHANGED:
		return "scrollHorizontalPositionChanged", true
	case EventType_SCROLL_POSITION_CHANGED:
		return "scrollPositionChanged", true
	case EventType_SCROLL_VERTICAL_POSITION_CHANGED:
		return "scrollVerticalPositionChanged", true
	case EventType_SCROLLED_TO_ANCHOR:
		return "scrolledToAnchor", true
	case EventType_SELECTED_CHANGED:
		return "selectedChanged", true
	case EventType_SELECTED_CHILDREN_CHANGED:
		return "selectedChildrenChanged", true
	case EventType_SELECTED_VALUE_CHANGED:
		return "selectedValueChanged", true
	case EventType_SELECTION:
		return "selection", true
	case EventType_SELECTION_ADD:
		return "selectionAdd", true
	case EventType_SELECTION_REMOVE:
		return "selectionRemove", true
	case EventType_SET_SIZE_CHANGED:
		return "setSizeChanged", true
	case EventType_SHOW:
		return "show", true
	case EventType_SORT_CHANGED:
		return "sortChanged", true
	case EventType_STATE_CHANGED:
		return "stateChanged", true
	case EventType_SUBTREE_CREATED:
		return "subtreeCreated", true
	case EventType_TEXT_ATTRIBUTE_CHANGED:
		return "textAttributeChanged", true
	case EventType_TEXT_SELECTION_CHANGED:
		return "textSelectionChanged", true
	case EventType_TEXT_CHANGED:
		return "textChanged", true
	case EventType_TOOLTIP_CLOSED:
		return "tooltipClosed", true
	case EventType_TOOLTIP_OPENED:
		return "tooltipOpened", true
	case EventType_TREE_CHANGED:
		return "treeChanged", true
	case EventType_VALUE_IN_TEXT_FIELD_CHANGED:
		return "valueInTextFieldChanged", true
	case EventType_VALUE_CHANGED:
		return "valueChanged", true
	case EventType_WINDOW_ACTIVATED:
		return "windowActivated", true
	case EventType_WINDOW_DEACTIVATED:
		return "windowDeactivated", true
	case EventType_WINDOW_VISIBILITY_CHANGED:
		return "windowVisibilityChanged", true
	default:
		return "", false
	}
}

type PerformActionCallbackWithNodeFunc func(this js.Ref, node AutomationNode) js.Ref

func (fn PerformActionCallbackWithNodeFunc) Register() js.Func[func(node AutomationNode)] {
	return js.RegisterCallback[func(node AutomationNode)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PerformActionCallbackWithNodeFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		AutomationNode{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PerformActionCallbackWithNode[T any] struct {
	Fn  func(arg T, this js.Ref, node AutomationNode) js.Ref
	Arg T
}

func (cb *PerformActionCallbackWithNode[T]) Register() js.Func[func(node AutomationNode)] {
	return js.RegisterCallback[func(node AutomationNode)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PerformActionCallbackWithNode[T]) DispatchCallback(
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

		AutomationNode{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ActionType uint32

const (
	_ ActionType = iota

	ActionType_ANNOTATE_PAGE_IMAGES
	ActionType_BLUR
	ActionType_CLEAR_ACCESSIBILITY_FOCUS
	ActionType_COLLAPSE
	ActionType_CUSTOM_ACTION
	ActionType_DECREMENT
	ActionType_DO_DEFAULT
	ActionType_EXPAND
	ActionType_FOCUS
	ActionType_GET_IMAGE_DATA
	ActionType_GET_TEXT_LOCATION
	ActionType_HIDE_TOOLTIP
	ActionType_HIT_TEST
	ActionType_INCREMENT
	ActionType_INTERNAL_INVALIDATE_TREE
	ActionType_LOAD_INLINE_TEXT_BOXES
	ActionType_LONG_CLICK
	ActionType_REPLACE_SELECTED_TEXT
	ActionType_RESUME_MEDIA
	ActionType_SCROLL_BACKWARD
	ActionType_SCROLL_DOWN
	ActionType_SCROLL_FORWARD
	ActionType_SCROLL_LEFT
	ActionType_SCROLL_RIGHT
	ActionType_SCROLL_UP
	ActionType_SCROLL_TO_MAKE_VISIBLE
	ActionType_SCROLL_TO_POINT
	ActionType_SCROLL_TO_POSITION_AT_ROW_COLUMN
	ActionType_SET_ACCESSIBILITY_FOCUS
	ActionType_SET_SCROLL_OFFSET
	ActionType_SET_SELECTION
	ActionType_SET_SEQUENTIAL_FOCUS_NAVIGATION_STARTING_POINT
	ActionType_SET_VALUE
	ActionType_SHOW_CONTEXT_MENU
	ActionType_SIGNAL_END_OF_TEST
	ActionType_SHOW_TOOLTIP
	ActionType_START_DUCKING_MEDIA
	ActionType_STOP_DUCKING_MEDIA
	ActionType_SUSPEND_MEDIA
)

func (ActionType) FromRef(str js.Ref) ActionType {
	return ActionType(bindings.ConstOfActionType(str))
}

func (x ActionType) String() (string, bool) {
	switch x {
	case ActionType_ANNOTATE_PAGE_IMAGES:
		return "annotatePageImages", true
	case ActionType_BLUR:
		return "blur", true
	case ActionType_CLEAR_ACCESSIBILITY_FOCUS:
		return "clearAccessibilityFocus", true
	case ActionType_COLLAPSE:
		return "collapse", true
	case ActionType_CUSTOM_ACTION:
		return "customAction", true
	case ActionType_DECREMENT:
		return "decrement", true
	case ActionType_DO_DEFAULT:
		return "doDefault", true
	case ActionType_EXPAND:
		return "expand", true
	case ActionType_FOCUS:
		return "focus", true
	case ActionType_GET_IMAGE_DATA:
		return "getImageData", true
	case ActionType_GET_TEXT_LOCATION:
		return "getTextLocation", true
	case ActionType_HIDE_TOOLTIP:
		return "hideTooltip", true
	case ActionType_HIT_TEST:
		return "hitTest", true
	case ActionType_INCREMENT:
		return "increment", true
	case ActionType_INTERNAL_INVALIDATE_TREE:
		return "internalInvalidateTree", true
	case ActionType_LOAD_INLINE_TEXT_BOXES:
		return "loadInlineTextBoxes", true
	case ActionType_LONG_CLICK:
		return "longClick", true
	case ActionType_REPLACE_SELECTED_TEXT:
		return "replaceSelectedText", true
	case ActionType_RESUME_MEDIA:
		return "resumeMedia", true
	case ActionType_SCROLL_BACKWARD:
		return "scrollBackward", true
	case ActionType_SCROLL_DOWN:
		return "scrollDown", true
	case ActionType_SCROLL_FORWARD:
		return "scrollForward", true
	case ActionType_SCROLL_LEFT:
		return "scrollLeft", true
	case ActionType_SCROLL_RIGHT:
		return "scrollRight", true
	case ActionType_SCROLL_UP:
		return "scrollUp", true
	case ActionType_SCROLL_TO_MAKE_VISIBLE:
		return "scrollToMakeVisible", true
	case ActionType_SCROLL_TO_POINT:
		return "scrollToPoint", true
	case ActionType_SCROLL_TO_POSITION_AT_ROW_COLUMN:
		return "scrollToPositionAtRowColumn", true
	case ActionType_SET_ACCESSIBILITY_FOCUS:
		return "setAccessibilityFocus", true
	case ActionType_SET_SCROLL_OFFSET:
		return "setScrollOffset", true
	case ActionType_SET_SELECTION:
		return "setSelection", true
	case ActionType_SET_SEQUENTIAL_FOCUS_NAVIGATION_STARTING_POINT:
		return "setSequentialFocusNavigationStartingPoint", true
	case ActionType_SET_VALUE:
		return "setValue", true
	case ActionType_SHOW_CONTEXT_MENU:
		return "showContextMenu", true
	case ActionType_SIGNAL_END_OF_TEST:
		return "signalEndOfTest", true
	case ActionType_SHOW_TOOLTIP:
		return "showTooltip", true
	case ActionType_START_DUCKING_MEDIA:
		return "startDuckingMedia", true
	case ActionType_STOP_DUCKING_MEDIA:
		return "stopDuckingMedia", true
	case ActionType_SUSPEND_MEDIA:
		return "suspendMedia", true
	default:
		return "", false
	}
}

type PerformActionCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn PerformActionCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PerformActionCallbackFunc) DispatchCallback(
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

type PerformActionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *PerformActionCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PerformActionCallback[T]) DispatchCallback(
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

type AutomationListenerFunc func(this js.Ref, event AutomationEvent) js.Ref

func (fn AutomationListenerFunc) Register() js.Func[func(event AutomationEvent)] {
	return js.RegisterCallback[func(event AutomationEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AutomationListenerFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		AutomationEvent{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AutomationListener[T any] struct {
	Fn  func(arg T, this js.Ref, event AutomationEvent) js.Ref
	Arg T
}

func (cb *AutomationListener[T]) Register() js.Func[func(event AutomationEvent)] {
	return js.RegisterCallback[func(event AutomationEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AutomationListener[T]) DispatchCallback(
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

		AutomationEvent{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IntentCommandType uint32

const (
	_ IntentCommandType = iota

	IntentCommandType_CLEAR_SELECTION
	IntentCommandType_DELETE
	IntentCommandType_DICTATE
	IntentCommandType_EXTEND_SELECTION
	IntentCommandType_FORMAT
	IntentCommandType_HISTORY
	IntentCommandType_INSERT
	IntentCommandType_MARKER
	IntentCommandType_MOVE_SELECTION
	IntentCommandType_SET_SELECTION
)

func (IntentCommandType) FromRef(str js.Ref) IntentCommandType {
	return IntentCommandType(bindings.ConstOfIntentCommandType(str))
}

func (x IntentCommandType) String() (string, bool) {
	switch x {
	case IntentCommandType_CLEAR_SELECTION:
		return "clearSelection", true
	case IntentCommandType_DELETE:
		return "delete", true
	case IntentCommandType_DICTATE:
		return "dictate", true
	case IntentCommandType_EXTEND_SELECTION:
		return "extendSelection", true
	case IntentCommandType_FORMAT:
		return "format", true
	case IntentCommandType_HISTORY:
		return "history", true
	case IntentCommandType_INSERT:
		return "insert", true
	case IntentCommandType_MARKER:
		return "marker", true
	case IntentCommandType_MOVE_SELECTION:
		return "moveSelection", true
	case IntentCommandType_SET_SELECTION:
		return "setSelection", true
	default:
		return "", false
	}
}

type IntentTextBoundaryType uint32

const (
	_ IntentTextBoundaryType = iota

	IntentTextBoundaryType_CHARACTER
	IntentTextBoundaryType_FORMAT_END
	IntentTextBoundaryType_FORMAT_START
	IntentTextBoundaryType_FORMAT_START_OR_END
	IntentTextBoundaryType_LINE_END
	IntentTextBoundaryType_LINE_START
	IntentTextBoundaryType_LINE_START_OR_END
	IntentTextBoundaryType_OBJECT
	IntentTextBoundaryType_PAGE_END
	IntentTextBoundaryType_PAGE_START
	IntentTextBoundaryType_PAGE_START_OR_END
	IntentTextBoundaryType_PARAGRAPH_END
	IntentTextBoundaryType_PARAGRAPH_START
	IntentTextBoundaryType_PARAGRAPH_START_SKIPPING_EMPTY_PARAGRAPHS
	IntentTextBoundaryType_PARAGRAPH_START_OR_END
	IntentTextBoundaryType_SENTENCE_END
	IntentTextBoundaryType_SENTENCE_START
	IntentTextBoundaryType_SENTENCE_START_OR_END
	IntentTextBoundaryType_WEB_PAGE
	IntentTextBoundaryType_WORD_END
	IntentTextBoundaryType_WORD_START
	IntentTextBoundaryType_WORD_START_OR_END
)

func (IntentTextBoundaryType) FromRef(str js.Ref) IntentTextBoundaryType {
	return IntentTextBoundaryType(bindings.ConstOfIntentTextBoundaryType(str))
}

func (x IntentTextBoundaryType) String() (string, bool) {
	switch x {
	case IntentTextBoundaryType_CHARACTER:
		return "character", true
	case IntentTextBoundaryType_FORMAT_END:
		return "formatEnd", true
	case IntentTextBoundaryType_FORMAT_START:
		return "formatStart", true
	case IntentTextBoundaryType_FORMAT_START_OR_END:
		return "formatStartOrEnd", true
	case IntentTextBoundaryType_LINE_END:
		return "lineEnd", true
	case IntentTextBoundaryType_LINE_START:
		return "lineStart", true
	case IntentTextBoundaryType_LINE_START_OR_END:
		return "lineStartOrEnd", true
	case IntentTextBoundaryType_OBJECT:
		return "object", true
	case IntentTextBoundaryType_PAGE_END:
		return "pageEnd", true
	case IntentTextBoundaryType_PAGE_START:
		return "pageStart", true
	case IntentTextBoundaryType_PAGE_START_OR_END:
		return "pageStartOrEnd", true
	case IntentTextBoundaryType_PARAGRAPH_END:
		return "paragraphEnd", true
	case IntentTextBoundaryType_PARAGRAPH_START:
		return "paragraphStart", true
	case IntentTextBoundaryType_PARAGRAPH_START_SKIPPING_EMPTY_PARAGRAPHS:
		return "paragraphStartSkippingEmptyParagraphs", true
	case IntentTextBoundaryType_PARAGRAPH_START_OR_END:
		return "paragraphStartOrEnd", true
	case IntentTextBoundaryType_SENTENCE_END:
		return "sentenceEnd", true
	case IntentTextBoundaryType_SENTENCE_START:
		return "sentenceStart", true
	case IntentTextBoundaryType_SENTENCE_START_OR_END:
		return "sentenceStartOrEnd", true
	case IntentTextBoundaryType_WEB_PAGE:
		return "webPage", true
	case IntentTextBoundaryType_WORD_END:
		return "wordEnd", true
	case IntentTextBoundaryType_WORD_START:
		return "wordStart", true
	case IntentTextBoundaryType_WORD_START_OR_END:
		return "wordStartOrEnd", true
	default:
		return "", false
	}
}

type IntentMoveDirectionType uint32

const (
	_ IntentMoveDirectionType = iota

	IntentMoveDirectionType_BACKWARD
	IntentMoveDirectionType_FORWARD
)

func (IntentMoveDirectionType) FromRef(str js.Ref) IntentMoveDirectionType {
	return IntentMoveDirectionType(bindings.ConstOfIntentMoveDirectionType(str))
}

func (x IntentMoveDirectionType) String() (string, bool) {
	switch x {
	case IntentMoveDirectionType_BACKWARD:
		return "backward", true
	case IntentMoveDirectionType_FORWARD:
		return "forward", true
	default:
		return "", false
	}
}

type AutomationIntent struct {
	// Command is "AutomationIntent.command"
	//
	// Optional
	Command IntentCommandType
	// TextBoundary is "AutomationIntent.textBoundary"
	//
	// Optional
	TextBoundary IntentTextBoundaryType
	// MoveDirection is "AutomationIntent.moveDirection"
	//
	// Optional
	MoveDirection IntentMoveDirectionType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AutomationIntent with all fields set.
func (p AutomationIntent) FromRef(ref js.Ref) AutomationIntent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AutomationIntent in the application heap.
func (p AutomationIntent) New() js.Ref {
	return bindings.AutomationIntentJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AutomationIntent) UpdateFrom(ref js.Ref) {
	bindings.AutomationIntentJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AutomationIntent) Update(ref js.Ref) {
	bindings.AutomationIntentJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AutomationIntent) FreeMembers(recursive bool) {
}

type AutomationEvent struct {
	ref js.Ref
}

func (this AutomationEvent) Once() AutomationEvent {
	this.ref.Once()
	return this
}

func (this AutomationEvent) Ref() js.Ref {
	return this.ref
}

func (this AutomationEvent) FromRef(ref js.Ref) AutomationEvent {
	this.ref = ref
	return this
}

func (this AutomationEvent) Free() {
	this.ref.Free()
}

// Target returns the value of property "AutomationEvent.target".
//
// It returns ok=false if there is no such property.
func (this AutomationEvent) Target() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationEventTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "AutomationEvent.target" to val.
//
// It returns false if the property cannot be set.
func (this AutomationEvent) SetTarget(val AutomationNode) bool {
	return js.True == bindings.SetAutomationEventTarget(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "AutomationEvent.type".
//
// It returns ok=false if there is no such property.
func (this AutomationEvent) Type() (ret EventType, ok bool) {
	ok = js.True == bindings.GetAutomationEventType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "AutomationEvent.type" to val.
//
// It returns false if the property cannot be set.
func (this AutomationEvent) SetType(val EventType) bool {
	return js.True == bindings.SetAutomationEventType(
		this.ref,
		uint32(val),
	)
}

// EventFrom returns the value of property "AutomationEvent.eventFrom".
//
// It returns ok=false if there is no such property.
func (this AutomationEvent) EventFrom() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationEventEventFrom(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEventFrom sets the value of property "AutomationEvent.eventFrom" to val.
//
// It returns false if the property cannot be set.
func (this AutomationEvent) SetEventFrom(val js.String) bool {
	return js.True == bindings.SetAutomationEventEventFrom(
		this.ref,
		val.Ref(),
	)
}

// MouseX returns the value of property "AutomationEvent.mouseX".
//
// It returns ok=false if there is no such property.
func (this AutomationEvent) MouseX() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationEventMouseX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMouseX sets the value of property "AutomationEvent.mouseX" to val.
//
// It returns false if the property cannot be set.
func (this AutomationEvent) SetMouseX(val int32) bool {
	return js.True == bindings.SetAutomationEventMouseX(
		this.ref,
		int32(val),
	)
}

// MouseY returns the value of property "AutomationEvent.mouseY".
//
// It returns ok=false if there is no such property.
func (this AutomationEvent) MouseY() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationEventMouseY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMouseY sets the value of property "AutomationEvent.mouseY" to val.
//
// It returns false if the property cannot be set.
func (this AutomationEvent) SetMouseY(val int32) bool {
	return js.True == bindings.SetAutomationEventMouseY(
		this.ref,
		int32(val),
	)
}

// Intents returns the value of property "AutomationEvent.intents".
//
// It returns ok=false if there is no such property.
func (this AutomationEvent) Intents() (ret js.Array[AutomationIntent], ok bool) {
	ok = js.True == bindings.GetAutomationEventIntents(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIntents sets the value of property "AutomationEvent.intents" to val.
//
// It returns false if the property cannot be set.
func (this AutomationEvent) SetIntents(val js.Array[AutomationIntent]) bool {
	return js.True == bindings.SetAutomationEventIntents(
		this.ref,
		val.Ref(),
	)
}

// HasFuncStopPropagation returns true if the method "AutomationEvent.stopPropagation" exists.
func (this AutomationEvent) HasFuncStopPropagation() bool {
	return js.True == bindings.HasFuncAutomationEventStopPropagation(
		this.ref,
	)
}

// FuncStopPropagation returns the method "AutomationEvent.stopPropagation".
func (this AutomationEvent) FuncStopPropagation() (fn js.Func[func()]) {
	bindings.FuncAutomationEventStopPropagation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StopPropagation calls the method "AutomationEvent.stopPropagation".
func (this AutomationEvent) StopPropagation() (ret js.Void) {
	bindings.CallAutomationEventStopPropagation(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStopPropagation calls the method "AutomationEvent.stopPropagation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationEvent) TryStopPropagation() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationEventStopPropagation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RoleType uint32

const (
	_ RoleType = iota

	RoleType_ABBR
	RoleType_ALERT
	RoleType_ALERT_DIALOG
	RoleType_APPLICATION
	RoleType_ARTICLE
	RoleType_AUDIO
	RoleType_BANNER
	RoleType_BLOCKQUOTE
	RoleType_BUTTON
	RoleType_CANVAS
	RoleType_CAPTION
	RoleType_CARET
	RoleType_CELL
	RoleType_CHECK_BOX
	RoleType_CLIENT
	RoleType_CODE
	RoleType_COLOR_WELL
	RoleType_COLUMN
	RoleType_COLUMN_HEADER
	RoleType_COMBO_BOX_GROUPING
	RoleType_COMBO_BOX_MENU_BUTTON
	RoleType_COMBO_BOX_SELECT
	RoleType_COMMENT
	RoleType_COMPLEMENTARY
	RoleType_CONTENT_DELETION
	RoleType_CONTENT_INSERTION
	RoleType_CONTENT_INFO
	RoleType_DATE
	RoleType_DATE_TIME
	RoleType_DEFINITION
	RoleType_DESCRIPTION_LIST
	RoleType_DESCRIPTION_LIST_DETAIL
	RoleType_DESCRIPTION_LIST_TERM
	RoleType_DESKTOP
	RoleType_DETAILS
	RoleType_DIALOG
	RoleType_DIRECTORY
	RoleType_DISCLOSURE_TRIANGLE
	RoleType_DOC_ABSTRACT
	RoleType_DOC_ACKNOWLEDGMENTS
	RoleType_DOC_AFTERWORD
	RoleType_DOC_APPENDIX
	RoleType_DOC_BACK_LINK
	RoleType_DOC_BIBLIO_ENTRY
	RoleType_DOC_BIBLIOGRAPHY
	RoleType_DOC_BIBLIO_REF
	RoleType_DOC_CHAPTER
	RoleType_DOC_COLOPHON
	RoleType_DOC_CONCLUSION
	RoleType_DOC_COVER
	RoleType_DOC_CREDIT
	RoleType_DOC_CREDITS
	RoleType_DOC_DEDICATION
	RoleType_DOC_ENDNOTE
	RoleType_DOC_ENDNOTES
	RoleType_DOC_EPIGRAPH
	RoleType_DOC_EPILOGUE
	RoleType_DOC_ERRATA
	RoleType_DOC_EXAMPLE
	RoleType_DOC_FOOTNOTE
	RoleType_DOC_FOREWORD
	RoleType_DOC_GLOSSARY
	RoleType_DOC_GLOSS_REF
	RoleType_DOC_INDEX
	RoleType_DOC_INTRODUCTION
	RoleType_DOC_NOTE_REF
	RoleType_DOC_NOTICE
	RoleType_DOC_PAGE_BREAK
	RoleType_DOC_PAGE_FOOTER
	RoleType_DOC_PAGE_HEADER
	RoleType_DOC_PAGE_LIST
	RoleType_DOC_PART
	RoleType_DOC_PREFACE
	RoleType_DOC_PROLOGUE
	RoleType_DOC_PULLQUOTE
	RoleType_DOC_QNA
	RoleType_DOC_SUBTITLE
	RoleType_DOC_TIP
	RoleType_DOC_TOC
	RoleType_DOCUMENT
	RoleType_EMBEDDED_OBJECT
	RoleType_EMPHASIS
	RoleType_FEED
	RoleType_FIGCAPTION
	RoleType_FIGURE
	RoleType_FOOTER
	RoleType_FOOTER_AS_NON_LANDMARK
	RoleType_FORM
	RoleType_GENERIC_CONTAINER
	RoleType_GRAPHICS_DOCUMENT
	RoleType_GRAPHICS_OBJECT
	RoleType_GRAPHICS_SYMBOL
	RoleType_GRID
	RoleType_GROUP
	RoleType_HEADER
	RoleType_HEADER_AS_NON_LANDMARK
	RoleType_HEADING
	RoleType_IFRAME
	RoleType_IFRAME_PRESENTATIONAL
	RoleType_IMAGE
	RoleType_IME_CANDIDATE
	RoleType_INLINE_TEXT_BOX
	RoleType_INPUT_TIME
	RoleType_KEYBOARD
	RoleType_LABEL_TEXT
	RoleType_LAYOUT_TABLE
	RoleType_LAYOUT_TABLE_CELL
	RoleType_LAYOUT_TABLE_ROW
	RoleType_LEGEND
	RoleType_LINE_BREAK
	RoleType_LINK
	RoleType_LIST
	RoleType_LIST_BOX
	RoleType_LIST_BOX_OPTION
	RoleType_LIST_GRID
	RoleType_LIST_ITEM
	RoleType_LIST_MARKER
	RoleType_LOG
	RoleType_MAIN
	RoleType_MARK
	RoleType_MARQUEE
	RoleType_MATH
	RoleType_MATH_ML_FRACTION
	RoleType_MATH_ML_IDENTIFIER
	RoleType_MATH_ML_MATH
	RoleType_MATH_ML_MULTISCRIPTS
	RoleType_MATH_ML_NONE_SCRIPT
	RoleType_MATH_ML_NUMBER
	RoleType_MATH_ML_OPERATOR
	RoleType_MATH_ML_OVER
	RoleType_MATH_ML_PRESCRIPT_DELIMITER
	RoleType_MATH_ML_ROOT
	RoleType_MATH_ML_ROW
	RoleType_MATH_ML_SQUARE_ROOT
	RoleType_MATH_ML_STRING_LITERAL
	RoleType_MATH_ML_SUB
	RoleType_MATH_ML_SUB_SUP
	RoleType_MATH_ML_SUP
	RoleType_MATH_ML_TABLE
	RoleType_MATH_ML_TABLE_CELL
	RoleType_MATH_ML_TABLE_ROW
	RoleType_MATH_ML_TEXT
	RoleType_MATH_ML_UNDER
	RoleType_MATH_ML_UNDER_OVER
	RoleType_MENU
	RoleType_MENU_BAR
	RoleType_MENU_ITEM
	RoleType_MENU_ITEM_CHECK_BOX
	RoleType_MENU_ITEM_RADIO
	RoleType_MENU_LIST_OPTION
	RoleType_MENU_LIST_POPUP
	RoleType_METER
	RoleType_NAVIGATION
	RoleType_NOTE
	RoleType_PANE
	RoleType_PARAGRAPH
	RoleType_PDF_ACTIONABLE_HIGHLIGHT
	RoleType_PDF_ROOT
	RoleType_PLUGIN_OBJECT
	RoleType_POP_UP_BUTTON
	RoleType_PORTAL
	RoleType_PRE_DEPRECATED
	RoleType_PROGRESS_INDICATOR
	RoleType_RADIO_BUTTON
	RoleType_RADIO_GROUP
	RoleType_REGION
	RoleType_ROOT_WEB_AREA
	RoleType_ROW
	RoleType_ROW_GROUP
	RoleType_ROW_HEADER
	RoleType_RUBY
	RoleType_RUBY_ANNOTATION
	RoleType_SCROLL_BAR
	RoleType_SCROLL_VIEW
	RoleType_SEARCH
	RoleType_SEARCH_BOX
	RoleType_SECTION
	RoleType_SLIDER
	RoleType_SPIN_BUTTON
	RoleType_SPLITTER
	RoleType_STATIC_TEXT
	RoleType_STATUS
	RoleType_STRONG
	RoleType_SUBSCRIPT
	RoleType_SUGGESTION
	RoleType_SUPERSCRIPT
	RoleType_SVG_ROOT
	RoleType_SWITCH
	RoleType_TAB
	RoleType_TAB_LIST
	RoleType_TAB_PANEL
	RoleType_TABLE
	RoleType_TABLE_HEADER_CONTAINER
	RoleType_TERM
	RoleType_TEXT_FIELD
	RoleType_TEXT_FIELD_WITH_COMBO_BOX
	RoleType_TIME
	RoleType_TIMER
	RoleType_TITLE_BAR
	RoleType_TOGGLE_BUTTON
	RoleType_TOOLBAR
	RoleType_TOOLTIP
	RoleType_TREE
	RoleType_TREE_GRID
	RoleType_TREE_ITEM
	RoleType_UNKNOWN
	RoleType_VIDEO
	RoleType_WEB_VIEW
	RoleType_WINDOW
)

func (RoleType) FromRef(str js.Ref) RoleType {
	return RoleType(bindings.ConstOfRoleType(str))
}

func (x RoleType) String() (string, bool) {
	switch x {
	case RoleType_ABBR:
		return "abbr", true
	case RoleType_ALERT:
		return "alert", true
	case RoleType_ALERT_DIALOG:
		return "alertDialog", true
	case RoleType_APPLICATION:
		return "application", true
	case RoleType_ARTICLE:
		return "article", true
	case RoleType_AUDIO:
		return "audio", true
	case RoleType_BANNER:
		return "banner", true
	case RoleType_BLOCKQUOTE:
		return "blockquote", true
	case RoleType_BUTTON:
		return "button", true
	case RoleType_CANVAS:
		return "canvas", true
	case RoleType_CAPTION:
		return "caption", true
	case RoleType_CARET:
		return "caret", true
	case RoleType_CELL:
		return "cell", true
	case RoleType_CHECK_BOX:
		return "checkBox", true
	case RoleType_CLIENT:
		return "client", true
	case RoleType_CODE:
		return "code", true
	case RoleType_COLOR_WELL:
		return "colorWell", true
	case RoleType_COLUMN:
		return "column", true
	case RoleType_COLUMN_HEADER:
		return "columnHeader", true
	case RoleType_COMBO_BOX_GROUPING:
		return "comboBoxGrouping", true
	case RoleType_COMBO_BOX_MENU_BUTTON:
		return "comboBoxMenuButton", true
	case RoleType_COMBO_BOX_SELECT:
		return "comboBoxSelect", true
	case RoleType_COMMENT:
		return "comment", true
	case RoleType_COMPLEMENTARY:
		return "complementary", true
	case RoleType_CONTENT_DELETION:
		return "contentDeletion", true
	case RoleType_CONTENT_INSERTION:
		return "contentInsertion", true
	case RoleType_CONTENT_INFO:
		return "contentInfo", true
	case RoleType_DATE:
		return "date", true
	case RoleType_DATE_TIME:
		return "dateTime", true
	case RoleType_DEFINITION:
		return "definition", true
	case RoleType_DESCRIPTION_LIST:
		return "descriptionList", true
	case RoleType_DESCRIPTION_LIST_DETAIL:
		return "descriptionListDetail", true
	case RoleType_DESCRIPTION_LIST_TERM:
		return "descriptionListTerm", true
	case RoleType_DESKTOP:
		return "desktop", true
	case RoleType_DETAILS:
		return "details", true
	case RoleType_DIALOG:
		return "dialog", true
	case RoleType_DIRECTORY:
		return "directory", true
	case RoleType_DISCLOSURE_TRIANGLE:
		return "disclosureTriangle", true
	case RoleType_DOC_ABSTRACT:
		return "docAbstract", true
	case RoleType_DOC_ACKNOWLEDGMENTS:
		return "docAcknowledgments", true
	case RoleType_DOC_AFTERWORD:
		return "docAfterword", true
	case RoleType_DOC_APPENDIX:
		return "docAppendix", true
	case RoleType_DOC_BACK_LINK:
		return "docBackLink", true
	case RoleType_DOC_BIBLIO_ENTRY:
		return "docBiblioEntry", true
	case RoleType_DOC_BIBLIOGRAPHY:
		return "docBibliography", true
	case RoleType_DOC_BIBLIO_REF:
		return "docBiblioRef", true
	case RoleType_DOC_CHAPTER:
		return "docChapter", true
	case RoleType_DOC_COLOPHON:
		return "docColophon", true
	case RoleType_DOC_CONCLUSION:
		return "docConclusion", true
	case RoleType_DOC_COVER:
		return "docCover", true
	case RoleType_DOC_CREDIT:
		return "docCredit", true
	case RoleType_DOC_CREDITS:
		return "docCredits", true
	case RoleType_DOC_DEDICATION:
		return "docDedication", true
	case RoleType_DOC_ENDNOTE:
		return "docEndnote", true
	case RoleType_DOC_ENDNOTES:
		return "docEndnotes", true
	case RoleType_DOC_EPIGRAPH:
		return "docEpigraph", true
	case RoleType_DOC_EPILOGUE:
		return "docEpilogue", true
	case RoleType_DOC_ERRATA:
		return "docErrata", true
	case RoleType_DOC_EXAMPLE:
		return "docExample", true
	case RoleType_DOC_FOOTNOTE:
		return "docFootnote", true
	case RoleType_DOC_FOREWORD:
		return "docForeword", true
	case RoleType_DOC_GLOSSARY:
		return "docGlossary", true
	case RoleType_DOC_GLOSS_REF:
		return "docGlossRef", true
	case RoleType_DOC_INDEX:
		return "docIndex", true
	case RoleType_DOC_INTRODUCTION:
		return "docIntroduction", true
	case RoleType_DOC_NOTE_REF:
		return "docNoteRef", true
	case RoleType_DOC_NOTICE:
		return "docNotice", true
	case RoleType_DOC_PAGE_BREAK:
		return "docPageBreak", true
	case RoleType_DOC_PAGE_FOOTER:
		return "docPageFooter", true
	case RoleType_DOC_PAGE_HEADER:
		return "docPageHeader", true
	case RoleType_DOC_PAGE_LIST:
		return "docPageList", true
	case RoleType_DOC_PART:
		return "docPart", true
	case RoleType_DOC_PREFACE:
		return "docPreface", true
	case RoleType_DOC_PROLOGUE:
		return "docPrologue", true
	case RoleType_DOC_PULLQUOTE:
		return "docPullquote", true
	case RoleType_DOC_QNA:
		return "docQna", true
	case RoleType_DOC_SUBTITLE:
		return "docSubtitle", true
	case RoleType_DOC_TIP:
		return "docTip", true
	case RoleType_DOC_TOC:
		return "docToc", true
	case RoleType_DOCUMENT:
		return "document", true
	case RoleType_EMBEDDED_OBJECT:
		return "embeddedObject", true
	case RoleType_EMPHASIS:
		return "emphasis", true
	case RoleType_FEED:
		return "feed", true
	case RoleType_FIGCAPTION:
		return "figcaption", true
	case RoleType_FIGURE:
		return "figure", true
	case RoleType_FOOTER:
		return "footer", true
	case RoleType_FOOTER_AS_NON_LANDMARK:
		return "footerAsNonLandmark", true
	case RoleType_FORM:
		return "form", true
	case RoleType_GENERIC_CONTAINER:
		return "genericContainer", true
	case RoleType_GRAPHICS_DOCUMENT:
		return "graphicsDocument", true
	case RoleType_GRAPHICS_OBJECT:
		return "graphicsObject", true
	case RoleType_GRAPHICS_SYMBOL:
		return "graphicsSymbol", true
	case RoleType_GRID:
		return "grid", true
	case RoleType_GROUP:
		return "group", true
	case RoleType_HEADER:
		return "header", true
	case RoleType_HEADER_AS_NON_LANDMARK:
		return "headerAsNonLandmark", true
	case RoleType_HEADING:
		return "heading", true
	case RoleType_IFRAME:
		return "iframe", true
	case RoleType_IFRAME_PRESENTATIONAL:
		return "iframePresentational", true
	case RoleType_IMAGE:
		return "image", true
	case RoleType_IME_CANDIDATE:
		return "imeCandidate", true
	case RoleType_INLINE_TEXT_BOX:
		return "inlineTextBox", true
	case RoleType_INPUT_TIME:
		return "inputTime", true
	case RoleType_KEYBOARD:
		return "keyboard", true
	case RoleType_LABEL_TEXT:
		return "labelText", true
	case RoleType_LAYOUT_TABLE:
		return "layoutTable", true
	case RoleType_LAYOUT_TABLE_CELL:
		return "layoutTableCell", true
	case RoleType_LAYOUT_TABLE_ROW:
		return "layoutTableRow", true
	case RoleType_LEGEND:
		return "legend", true
	case RoleType_LINE_BREAK:
		return "lineBreak", true
	case RoleType_LINK:
		return "link", true
	case RoleType_LIST:
		return "list", true
	case RoleType_LIST_BOX:
		return "listBox", true
	case RoleType_LIST_BOX_OPTION:
		return "listBoxOption", true
	case RoleType_LIST_GRID:
		return "listGrid", true
	case RoleType_LIST_ITEM:
		return "listItem", true
	case RoleType_LIST_MARKER:
		return "listMarker", true
	case RoleType_LOG:
		return "log", true
	case RoleType_MAIN:
		return "main", true
	case RoleType_MARK:
		return "mark", true
	case RoleType_MARQUEE:
		return "marquee", true
	case RoleType_MATH:
		return "math", true
	case RoleType_MATH_ML_FRACTION:
		return "mathMLFraction", true
	case RoleType_MATH_ML_IDENTIFIER:
		return "mathMLIdentifier", true
	case RoleType_MATH_ML_MATH:
		return "mathMLMath", true
	case RoleType_MATH_ML_MULTISCRIPTS:
		return "mathMLMultiscripts", true
	case RoleType_MATH_ML_NONE_SCRIPT:
		return "mathMLNoneScript", true
	case RoleType_MATH_ML_NUMBER:
		return "mathMLNumber", true
	case RoleType_MATH_ML_OPERATOR:
		return "mathMLOperator", true
	case RoleType_MATH_ML_OVER:
		return "mathMLOver", true
	case RoleType_MATH_ML_PRESCRIPT_DELIMITER:
		return "mathMLPrescriptDelimiter", true
	case RoleType_MATH_ML_ROOT:
		return "mathMLRoot", true
	case RoleType_MATH_ML_ROW:
		return "mathMLRow", true
	case RoleType_MATH_ML_SQUARE_ROOT:
		return "mathMLSquareRoot", true
	case RoleType_MATH_ML_STRING_LITERAL:
		return "mathMLStringLiteral", true
	case RoleType_MATH_ML_SUB:
		return "mathMLSub", true
	case RoleType_MATH_ML_SUB_SUP:
		return "mathMLSubSup", true
	case RoleType_MATH_ML_SUP:
		return "mathMLSup", true
	case RoleType_MATH_ML_TABLE:
		return "mathMLTable", true
	case RoleType_MATH_ML_TABLE_CELL:
		return "mathMLTableCell", true
	case RoleType_MATH_ML_TABLE_ROW:
		return "mathMLTableRow", true
	case RoleType_MATH_ML_TEXT:
		return "mathMLText", true
	case RoleType_MATH_ML_UNDER:
		return "mathMLUnder", true
	case RoleType_MATH_ML_UNDER_OVER:
		return "mathMLUnderOver", true
	case RoleType_MENU:
		return "menu", true
	case RoleType_MENU_BAR:
		return "menuBar", true
	case RoleType_MENU_ITEM:
		return "menuItem", true
	case RoleType_MENU_ITEM_CHECK_BOX:
		return "menuItemCheckBox", true
	case RoleType_MENU_ITEM_RADIO:
		return "menuItemRadio", true
	case RoleType_MENU_LIST_OPTION:
		return "menuListOption", true
	case RoleType_MENU_LIST_POPUP:
		return "menuListPopup", true
	case RoleType_METER:
		return "meter", true
	case RoleType_NAVIGATION:
		return "navigation", true
	case RoleType_NOTE:
		return "note", true
	case RoleType_PANE:
		return "pane", true
	case RoleType_PARAGRAPH:
		return "paragraph", true
	case RoleType_PDF_ACTIONABLE_HIGHLIGHT:
		return "pdfActionableHighlight", true
	case RoleType_PDF_ROOT:
		return "pdfRoot", true
	case RoleType_PLUGIN_OBJECT:
		return "pluginObject", true
	case RoleType_POP_UP_BUTTON:
		return "popUpButton", true
	case RoleType_PORTAL:
		return "portal", true
	case RoleType_PRE_DEPRECATED:
		return "preDeprecated", true
	case RoleType_PROGRESS_INDICATOR:
		return "progressIndicator", true
	case RoleType_RADIO_BUTTON:
		return "radioButton", true
	case RoleType_RADIO_GROUP:
		return "radioGroup", true
	case RoleType_REGION:
		return "region", true
	case RoleType_ROOT_WEB_AREA:
		return "rootWebArea", true
	case RoleType_ROW:
		return "row", true
	case RoleType_ROW_GROUP:
		return "rowGroup", true
	case RoleType_ROW_HEADER:
		return "rowHeader", true
	case RoleType_RUBY:
		return "ruby", true
	case RoleType_RUBY_ANNOTATION:
		return "rubyAnnotation", true
	case RoleType_SCROLL_BAR:
		return "scrollBar", true
	case RoleType_SCROLL_VIEW:
		return "scrollView", true
	case RoleType_SEARCH:
		return "search", true
	case RoleType_SEARCH_BOX:
		return "searchBox", true
	case RoleType_SECTION:
		return "section", true
	case RoleType_SLIDER:
		return "slider", true
	case RoleType_SPIN_BUTTON:
		return "spinButton", true
	case RoleType_SPLITTER:
		return "splitter", true
	case RoleType_STATIC_TEXT:
		return "staticText", true
	case RoleType_STATUS:
		return "status", true
	case RoleType_STRONG:
		return "strong", true
	case RoleType_SUBSCRIPT:
		return "subscript", true
	case RoleType_SUGGESTION:
		return "suggestion", true
	case RoleType_SUPERSCRIPT:
		return "superscript", true
	case RoleType_SVG_ROOT:
		return "svgRoot", true
	case RoleType_SWITCH:
		return "switch", true
	case RoleType_TAB:
		return "tab", true
	case RoleType_TAB_LIST:
		return "tabList", true
	case RoleType_TAB_PANEL:
		return "tabPanel", true
	case RoleType_TABLE:
		return "table", true
	case RoleType_TABLE_HEADER_CONTAINER:
		return "tableHeaderContainer", true
	case RoleType_TERM:
		return "term", true
	case RoleType_TEXT_FIELD:
		return "textField", true
	case RoleType_TEXT_FIELD_WITH_COMBO_BOX:
		return "textFieldWithComboBox", true
	case RoleType_TIME:
		return "time", true
	case RoleType_TIMER:
		return "timer", true
	case RoleType_TITLE_BAR:
		return "titleBar", true
	case RoleType_TOGGLE_BUTTON:
		return "toggleButton", true
	case RoleType_TOOLBAR:
		return "toolbar", true
	case RoleType_TOOLTIP:
		return "tooltip", true
	case RoleType_TREE:
		return "tree", true
	case RoleType_TREE_GRID:
		return "treeGrid", true
	case RoleType_TREE_ITEM:
		return "treeItem", true
	case RoleType_UNKNOWN:
		return "unknown", true
	case RoleType_VIDEO:
		return "video", true
	case RoleType_WEB_VIEW:
		return "webView", true
	case RoleType_WINDOW:
		return "window", true
	default:
		return "", false
	}
}

type FindParams struct {
	// Role is "FindParams.role"
	//
	// Optional
	Role RoleType
	// State is "FindParams.state"
	//
	// Optional
	State js.Object
	// Attributes is "FindParams.attributes"
	//
	// Optional
	Attributes js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FindParams with all fields set.
func (p FindParams) FromRef(ref js.Ref) FindParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FindParams in the application heap.
func (p FindParams) New() js.Ref {
	return bindings.FindParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FindParams) UpdateFrom(ref js.Ref) {
	bindings.FindParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FindParams) Update(ref js.Ref) {
	bindings.FindParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FindParams) FreeMembers(recursive bool) {
	js.Free(
		p.State.Ref(),
		p.Attributes.Ref(),
	)
	p.State = p.State.FromRef(js.Undefined)
	p.Attributes = p.Attributes.FromRef(js.Undefined)
}

type AutomationPosition struct {
	ref js.Ref
}

func (this AutomationPosition) Once() AutomationPosition {
	this.ref.Once()
	return this
}

func (this AutomationPosition) Ref() js.Ref {
	return this.ref
}

func (this AutomationPosition) FromRef(ref js.Ref) AutomationPosition {
	this.ref = ref
	return this
}

func (this AutomationPosition) Free() {
	this.ref.Free()
}

// Node returns the value of property "AutomationPosition.node".
//
// It returns ok=false if there is no such property.
func (this AutomationPosition) Node() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationPositionNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNode sets the value of property "AutomationPosition.node" to val.
//
// It returns false if the property cannot be set.
func (this AutomationPosition) SetNode(val AutomationNode) bool {
	return js.True == bindings.SetAutomationPositionNode(
		this.ref,
		val.Ref(),
	)
}

// ChildIndex returns the value of property "AutomationPosition.childIndex".
//
// It returns ok=false if there is no such property.
func (this AutomationPosition) ChildIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationPositionChildIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChildIndex sets the value of property "AutomationPosition.childIndex" to val.
//
// It returns false if the property cannot be set.
func (this AutomationPosition) SetChildIndex(val int32) bool {
	return js.True == bindings.SetAutomationPositionChildIndex(
		this.ref,
		int32(val),
	)
}

// TextOffset returns the value of property "AutomationPosition.textOffset".
//
// It returns ok=false if there is no such property.
func (this AutomationPosition) TextOffset() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationPositionTextOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTextOffset sets the value of property "AutomationPosition.textOffset" to val.
//
// It returns false if the property cannot be set.
func (this AutomationPosition) SetTextOffset(val int32) bool {
	return js.True == bindings.SetAutomationPositionTextOffset(
		this.ref,
		int32(val),
	)
}

// Affinity returns the value of property "AutomationPosition.affinity".
//
// It returns ok=false if there is no such property.
func (this AutomationPosition) Affinity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationPositionAffinity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAffinity sets the value of property "AutomationPosition.affinity" to val.
//
// It returns false if the property cannot be set.
func (this AutomationPosition) SetAffinity(val js.String) bool {
	return js.True == bindings.SetAutomationPositionAffinity(
		this.ref,
		val.Ref(),
	)
}

// HasFuncIsNullPosition returns true if the method "AutomationPosition.isNullPosition" exists.
func (this AutomationPosition) HasFuncIsNullPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionIsNullPosition(
		this.ref,
	)
}

// FuncIsNullPosition returns the method "AutomationPosition.isNullPosition".
func (this AutomationPosition) FuncIsNullPosition() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionIsNullPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsNullPosition calls the method "AutomationPosition.isNullPosition".
func (this AutomationPosition) IsNullPosition() (ret bool) {
	bindings.CallAutomationPositionIsNullPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsNullPosition calls the method "AutomationPosition.isNullPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryIsNullPosition() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionIsNullPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsTreePosition returns true if the method "AutomationPosition.isTreePosition" exists.
func (this AutomationPosition) HasFuncIsTreePosition() bool {
	return js.True == bindings.HasFuncAutomationPositionIsTreePosition(
		this.ref,
	)
}

// FuncIsTreePosition returns the method "AutomationPosition.isTreePosition".
func (this AutomationPosition) FuncIsTreePosition() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionIsTreePosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsTreePosition calls the method "AutomationPosition.isTreePosition".
func (this AutomationPosition) IsTreePosition() (ret bool) {
	bindings.CallAutomationPositionIsTreePosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsTreePosition calls the method "AutomationPosition.isTreePosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryIsTreePosition() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionIsTreePosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsTextPosition returns true if the method "AutomationPosition.isTextPosition" exists.
func (this AutomationPosition) HasFuncIsTextPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionIsTextPosition(
		this.ref,
	)
}

// FuncIsTextPosition returns the method "AutomationPosition.isTextPosition".
func (this AutomationPosition) FuncIsTextPosition() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionIsTextPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsTextPosition calls the method "AutomationPosition.isTextPosition".
func (this AutomationPosition) IsTextPosition() (ret bool) {
	bindings.CallAutomationPositionIsTextPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsTextPosition calls the method "AutomationPosition.isTextPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryIsTextPosition() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionIsTextPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsLeafTextPosition returns true if the method "AutomationPosition.isLeafTextPosition" exists.
func (this AutomationPosition) HasFuncIsLeafTextPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionIsLeafTextPosition(
		this.ref,
	)
}

// FuncIsLeafTextPosition returns the method "AutomationPosition.isLeafTextPosition".
func (this AutomationPosition) FuncIsLeafTextPosition() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionIsLeafTextPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsLeafTextPosition calls the method "AutomationPosition.isLeafTextPosition".
func (this AutomationPosition) IsLeafTextPosition() (ret bool) {
	bindings.CallAutomationPositionIsLeafTextPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsLeafTextPosition calls the method "AutomationPosition.isLeafTextPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryIsLeafTextPosition() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionIsLeafTextPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtStartOfAnchor returns true if the method "AutomationPosition.atStartOfAnchor" exists.
func (this AutomationPosition) HasFuncAtStartOfAnchor() bool {
	return js.True == bindings.HasFuncAutomationPositionAtStartOfAnchor(
		this.ref,
	)
}

// FuncAtStartOfAnchor returns the method "AutomationPosition.atStartOfAnchor".
func (this AutomationPosition) FuncAtStartOfAnchor() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtStartOfAnchor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtStartOfAnchor calls the method "AutomationPosition.atStartOfAnchor".
func (this AutomationPosition) AtStartOfAnchor() (ret bool) {
	bindings.CallAutomationPositionAtStartOfAnchor(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtStartOfAnchor calls the method "AutomationPosition.atStartOfAnchor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtStartOfAnchor() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtStartOfAnchor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtEndOfAnchor returns true if the method "AutomationPosition.atEndOfAnchor" exists.
func (this AutomationPosition) HasFuncAtEndOfAnchor() bool {
	return js.True == bindings.HasFuncAutomationPositionAtEndOfAnchor(
		this.ref,
	)
}

// FuncAtEndOfAnchor returns the method "AutomationPosition.atEndOfAnchor".
func (this AutomationPosition) FuncAtEndOfAnchor() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtEndOfAnchor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtEndOfAnchor calls the method "AutomationPosition.atEndOfAnchor".
func (this AutomationPosition) AtEndOfAnchor() (ret bool) {
	bindings.CallAutomationPositionAtEndOfAnchor(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtEndOfAnchor calls the method "AutomationPosition.atEndOfAnchor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtEndOfAnchor() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtEndOfAnchor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtStartOfWord returns true if the method "AutomationPosition.atStartOfWord" exists.
func (this AutomationPosition) HasFuncAtStartOfWord() bool {
	return js.True == bindings.HasFuncAutomationPositionAtStartOfWord(
		this.ref,
	)
}

// FuncAtStartOfWord returns the method "AutomationPosition.atStartOfWord".
func (this AutomationPosition) FuncAtStartOfWord() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtStartOfWord(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtStartOfWord calls the method "AutomationPosition.atStartOfWord".
func (this AutomationPosition) AtStartOfWord() (ret bool) {
	bindings.CallAutomationPositionAtStartOfWord(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtStartOfWord calls the method "AutomationPosition.atStartOfWord"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtStartOfWord() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtStartOfWord(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtEndOfWord returns true if the method "AutomationPosition.atEndOfWord" exists.
func (this AutomationPosition) HasFuncAtEndOfWord() bool {
	return js.True == bindings.HasFuncAutomationPositionAtEndOfWord(
		this.ref,
	)
}

// FuncAtEndOfWord returns the method "AutomationPosition.atEndOfWord".
func (this AutomationPosition) FuncAtEndOfWord() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtEndOfWord(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtEndOfWord calls the method "AutomationPosition.atEndOfWord".
func (this AutomationPosition) AtEndOfWord() (ret bool) {
	bindings.CallAutomationPositionAtEndOfWord(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtEndOfWord calls the method "AutomationPosition.atEndOfWord"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtEndOfWord() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtEndOfWord(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtStartOfLine returns true if the method "AutomationPosition.atStartOfLine" exists.
func (this AutomationPosition) HasFuncAtStartOfLine() bool {
	return js.True == bindings.HasFuncAutomationPositionAtStartOfLine(
		this.ref,
	)
}

// FuncAtStartOfLine returns the method "AutomationPosition.atStartOfLine".
func (this AutomationPosition) FuncAtStartOfLine() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtStartOfLine(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtStartOfLine calls the method "AutomationPosition.atStartOfLine".
func (this AutomationPosition) AtStartOfLine() (ret bool) {
	bindings.CallAutomationPositionAtStartOfLine(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtStartOfLine calls the method "AutomationPosition.atStartOfLine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtStartOfLine() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtStartOfLine(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtEndOfLine returns true if the method "AutomationPosition.atEndOfLine" exists.
func (this AutomationPosition) HasFuncAtEndOfLine() bool {
	return js.True == bindings.HasFuncAutomationPositionAtEndOfLine(
		this.ref,
	)
}

// FuncAtEndOfLine returns the method "AutomationPosition.atEndOfLine".
func (this AutomationPosition) FuncAtEndOfLine() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtEndOfLine(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtEndOfLine calls the method "AutomationPosition.atEndOfLine".
func (this AutomationPosition) AtEndOfLine() (ret bool) {
	bindings.CallAutomationPositionAtEndOfLine(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtEndOfLine calls the method "AutomationPosition.atEndOfLine"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtEndOfLine() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtEndOfLine(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtStartOfParagraph returns true if the method "AutomationPosition.atStartOfParagraph" exists.
func (this AutomationPosition) HasFuncAtStartOfParagraph() bool {
	return js.True == bindings.HasFuncAutomationPositionAtStartOfParagraph(
		this.ref,
	)
}

// FuncAtStartOfParagraph returns the method "AutomationPosition.atStartOfParagraph".
func (this AutomationPosition) FuncAtStartOfParagraph() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtStartOfParagraph(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtStartOfParagraph calls the method "AutomationPosition.atStartOfParagraph".
func (this AutomationPosition) AtStartOfParagraph() (ret bool) {
	bindings.CallAutomationPositionAtStartOfParagraph(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtStartOfParagraph calls the method "AutomationPosition.atStartOfParagraph"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtStartOfParagraph() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtStartOfParagraph(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtEndOfParagraph returns true if the method "AutomationPosition.atEndOfParagraph" exists.
func (this AutomationPosition) HasFuncAtEndOfParagraph() bool {
	return js.True == bindings.HasFuncAutomationPositionAtEndOfParagraph(
		this.ref,
	)
}

// FuncAtEndOfParagraph returns the method "AutomationPosition.atEndOfParagraph".
func (this AutomationPosition) FuncAtEndOfParagraph() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtEndOfParagraph(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtEndOfParagraph calls the method "AutomationPosition.atEndOfParagraph".
func (this AutomationPosition) AtEndOfParagraph() (ret bool) {
	bindings.CallAutomationPositionAtEndOfParagraph(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtEndOfParagraph calls the method "AutomationPosition.atEndOfParagraph"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtEndOfParagraph() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtEndOfParagraph(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtStartOfPage returns true if the method "AutomationPosition.atStartOfPage" exists.
func (this AutomationPosition) HasFuncAtStartOfPage() bool {
	return js.True == bindings.HasFuncAutomationPositionAtStartOfPage(
		this.ref,
	)
}

// FuncAtStartOfPage returns the method "AutomationPosition.atStartOfPage".
func (this AutomationPosition) FuncAtStartOfPage() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtStartOfPage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtStartOfPage calls the method "AutomationPosition.atStartOfPage".
func (this AutomationPosition) AtStartOfPage() (ret bool) {
	bindings.CallAutomationPositionAtStartOfPage(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtStartOfPage calls the method "AutomationPosition.atStartOfPage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtStartOfPage() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtStartOfPage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtEndOfPage returns true if the method "AutomationPosition.atEndOfPage" exists.
func (this AutomationPosition) HasFuncAtEndOfPage() bool {
	return js.True == bindings.HasFuncAutomationPositionAtEndOfPage(
		this.ref,
	)
}

// FuncAtEndOfPage returns the method "AutomationPosition.atEndOfPage".
func (this AutomationPosition) FuncAtEndOfPage() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtEndOfPage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtEndOfPage calls the method "AutomationPosition.atEndOfPage".
func (this AutomationPosition) AtEndOfPage() (ret bool) {
	bindings.CallAutomationPositionAtEndOfPage(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtEndOfPage calls the method "AutomationPosition.atEndOfPage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtEndOfPage() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtEndOfPage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtStartOfFormat returns true if the method "AutomationPosition.atStartOfFormat" exists.
func (this AutomationPosition) HasFuncAtStartOfFormat() bool {
	return js.True == bindings.HasFuncAutomationPositionAtStartOfFormat(
		this.ref,
	)
}

// FuncAtStartOfFormat returns the method "AutomationPosition.atStartOfFormat".
func (this AutomationPosition) FuncAtStartOfFormat() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtStartOfFormat(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtStartOfFormat calls the method "AutomationPosition.atStartOfFormat".
func (this AutomationPosition) AtStartOfFormat() (ret bool) {
	bindings.CallAutomationPositionAtStartOfFormat(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtStartOfFormat calls the method "AutomationPosition.atStartOfFormat"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtStartOfFormat() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtStartOfFormat(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtEndOfFormat returns true if the method "AutomationPosition.atEndOfFormat" exists.
func (this AutomationPosition) HasFuncAtEndOfFormat() bool {
	return js.True == bindings.HasFuncAutomationPositionAtEndOfFormat(
		this.ref,
	)
}

// FuncAtEndOfFormat returns the method "AutomationPosition.atEndOfFormat".
func (this AutomationPosition) FuncAtEndOfFormat() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtEndOfFormat(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtEndOfFormat calls the method "AutomationPosition.atEndOfFormat".
func (this AutomationPosition) AtEndOfFormat() (ret bool) {
	bindings.CallAutomationPositionAtEndOfFormat(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtEndOfFormat calls the method "AutomationPosition.atEndOfFormat"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtEndOfFormat() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtEndOfFormat(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtStartOfDocument returns true if the method "AutomationPosition.atStartOfDocument" exists.
func (this AutomationPosition) HasFuncAtStartOfDocument() bool {
	return js.True == bindings.HasFuncAutomationPositionAtStartOfDocument(
		this.ref,
	)
}

// FuncAtStartOfDocument returns the method "AutomationPosition.atStartOfDocument".
func (this AutomationPosition) FuncAtStartOfDocument() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtStartOfDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtStartOfDocument calls the method "AutomationPosition.atStartOfDocument".
func (this AutomationPosition) AtStartOfDocument() (ret bool) {
	bindings.CallAutomationPositionAtStartOfDocument(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtStartOfDocument calls the method "AutomationPosition.atStartOfDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtStartOfDocument() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtStartOfDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAtEndOfDocument returns true if the method "AutomationPosition.atEndOfDocument" exists.
func (this AutomationPosition) HasFuncAtEndOfDocument() bool {
	return js.True == bindings.HasFuncAutomationPositionAtEndOfDocument(
		this.ref,
	)
}

// FuncAtEndOfDocument returns the method "AutomationPosition.atEndOfDocument".
func (this AutomationPosition) FuncAtEndOfDocument() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionAtEndOfDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AtEndOfDocument calls the method "AutomationPosition.atEndOfDocument".
func (this AutomationPosition) AtEndOfDocument() (ret bool) {
	bindings.CallAutomationPositionAtEndOfDocument(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAtEndOfDocument calls the method "AutomationPosition.atEndOfDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAtEndOfDocument() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAtEndOfDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAsTreePosition returns true if the method "AutomationPosition.asTreePosition" exists.
func (this AutomationPosition) HasFuncAsTreePosition() bool {
	return js.True == bindings.HasFuncAutomationPositionAsTreePosition(
		this.ref,
	)
}

// FuncAsTreePosition returns the method "AutomationPosition.asTreePosition".
func (this AutomationPosition) FuncAsTreePosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionAsTreePosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AsTreePosition calls the method "AutomationPosition.asTreePosition".
func (this AutomationPosition) AsTreePosition() (ret js.Void) {
	bindings.CallAutomationPositionAsTreePosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAsTreePosition calls the method "AutomationPosition.asTreePosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAsTreePosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAsTreePosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAsTextPosition returns true if the method "AutomationPosition.asTextPosition" exists.
func (this AutomationPosition) HasFuncAsTextPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionAsTextPosition(
		this.ref,
	)
}

// FuncAsTextPosition returns the method "AutomationPosition.asTextPosition".
func (this AutomationPosition) FuncAsTextPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionAsTextPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AsTextPosition calls the method "AutomationPosition.asTextPosition".
func (this AutomationPosition) AsTextPosition() (ret js.Void) {
	bindings.CallAutomationPositionAsTextPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAsTextPosition calls the method "AutomationPosition.asTextPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAsTextPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAsTextPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAsLeafTextPosition returns true if the method "AutomationPosition.asLeafTextPosition" exists.
func (this AutomationPosition) HasFuncAsLeafTextPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionAsLeafTextPosition(
		this.ref,
	)
}

// FuncAsLeafTextPosition returns the method "AutomationPosition.asLeafTextPosition".
func (this AutomationPosition) FuncAsLeafTextPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionAsLeafTextPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AsLeafTextPosition calls the method "AutomationPosition.asLeafTextPosition".
func (this AutomationPosition) AsLeafTextPosition() (ret js.Void) {
	bindings.CallAutomationPositionAsLeafTextPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAsLeafTextPosition calls the method "AutomationPosition.asLeafTextPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryAsLeafTextPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionAsLeafTextPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPositionAtStartOfAnchor returns true if the method "AutomationPosition.moveToPositionAtStartOfAnchor" exists.
func (this AutomationPosition) HasFuncMoveToPositionAtStartOfAnchor() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPositionAtStartOfAnchor(
		this.ref,
	)
}

// FuncMoveToPositionAtStartOfAnchor returns the method "AutomationPosition.moveToPositionAtStartOfAnchor".
func (this AutomationPosition) FuncMoveToPositionAtStartOfAnchor() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPositionAtStartOfAnchor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPositionAtStartOfAnchor calls the method "AutomationPosition.moveToPositionAtStartOfAnchor".
func (this AutomationPosition) MoveToPositionAtStartOfAnchor() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPositionAtStartOfAnchor(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPositionAtStartOfAnchor calls the method "AutomationPosition.moveToPositionAtStartOfAnchor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPositionAtStartOfAnchor() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPositionAtStartOfAnchor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPositionAtEndOfAnchor returns true if the method "AutomationPosition.moveToPositionAtEndOfAnchor" exists.
func (this AutomationPosition) HasFuncMoveToPositionAtEndOfAnchor() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPositionAtEndOfAnchor(
		this.ref,
	)
}

// FuncMoveToPositionAtEndOfAnchor returns the method "AutomationPosition.moveToPositionAtEndOfAnchor".
func (this AutomationPosition) FuncMoveToPositionAtEndOfAnchor() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPositionAtEndOfAnchor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPositionAtEndOfAnchor calls the method "AutomationPosition.moveToPositionAtEndOfAnchor".
func (this AutomationPosition) MoveToPositionAtEndOfAnchor() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPositionAtEndOfAnchor(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPositionAtEndOfAnchor calls the method "AutomationPosition.moveToPositionAtEndOfAnchor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPositionAtEndOfAnchor() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPositionAtEndOfAnchor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPositionAtStartOfDocument returns true if the method "AutomationPosition.moveToPositionAtStartOfDocument" exists.
func (this AutomationPosition) HasFuncMoveToPositionAtStartOfDocument() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPositionAtStartOfDocument(
		this.ref,
	)
}

// FuncMoveToPositionAtStartOfDocument returns the method "AutomationPosition.moveToPositionAtStartOfDocument".
func (this AutomationPosition) FuncMoveToPositionAtStartOfDocument() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPositionAtStartOfDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPositionAtStartOfDocument calls the method "AutomationPosition.moveToPositionAtStartOfDocument".
func (this AutomationPosition) MoveToPositionAtStartOfDocument() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPositionAtStartOfDocument(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPositionAtStartOfDocument calls the method "AutomationPosition.moveToPositionAtStartOfDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPositionAtStartOfDocument() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPositionAtStartOfDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPositionAtEndOfDocument returns true if the method "AutomationPosition.moveToPositionAtEndOfDocument" exists.
func (this AutomationPosition) HasFuncMoveToPositionAtEndOfDocument() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPositionAtEndOfDocument(
		this.ref,
	)
}

// FuncMoveToPositionAtEndOfDocument returns the method "AutomationPosition.moveToPositionAtEndOfDocument".
func (this AutomationPosition) FuncMoveToPositionAtEndOfDocument() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPositionAtEndOfDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPositionAtEndOfDocument calls the method "AutomationPosition.moveToPositionAtEndOfDocument".
func (this AutomationPosition) MoveToPositionAtEndOfDocument() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPositionAtEndOfDocument(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPositionAtEndOfDocument calls the method "AutomationPosition.moveToPositionAtEndOfDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPositionAtEndOfDocument() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPositionAtEndOfDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToParentPosition returns true if the method "AutomationPosition.moveToParentPosition" exists.
func (this AutomationPosition) HasFuncMoveToParentPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToParentPosition(
		this.ref,
	)
}

// FuncMoveToParentPosition returns the method "AutomationPosition.moveToParentPosition".
func (this AutomationPosition) FuncMoveToParentPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToParentPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToParentPosition calls the method "AutomationPosition.moveToParentPosition".
func (this AutomationPosition) MoveToParentPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToParentPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToParentPosition calls the method "AutomationPosition.moveToParentPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToParentPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToParentPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextLeafTreePosition returns true if the method "AutomationPosition.moveToNextLeafTreePosition" exists.
func (this AutomationPosition) HasFuncMoveToNextLeafTreePosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextLeafTreePosition(
		this.ref,
	)
}

// FuncMoveToNextLeafTreePosition returns the method "AutomationPosition.moveToNextLeafTreePosition".
func (this AutomationPosition) FuncMoveToNextLeafTreePosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextLeafTreePosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextLeafTreePosition calls the method "AutomationPosition.moveToNextLeafTreePosition".
func (this AutomationPosition) MoveToNextLeafTreePosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextLeafTreePosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextLeafTreePosition calls the method "AutomationPosition.moveToNextLeafTreePosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextLeafTreePosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextLeafTreePosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousLeafTreePosition returns true if the method "AutomationPosition.moveToPreviousLeafTreePosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousLeafTreePosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousLeafTreePosition(
		this.ref,
	)
}

// FuncMoveToPreviousLeafTreePosition returns the method "AutomationPosition.moveToPreviousLeafTreePosition".
func (this AutomationPosition) FuncMoveToPreviousLeafTreePosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousLeafTreePosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousLeafTreePosition calls the method "AutomationPosition.moveToPreviousLeafTreePosition".
func (this AutomationPosition) MoveToPreviousLeafTreePosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousLeafTreePosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousLeafTreePosition calls the method "AutomationPosition.moveToPreviousLeafTreePosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousLeafTreePosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousLeafTreePosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextLeafTextPosition returns true if the method "AutomationPosition.moveToNextLeafTextPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextLeafTextPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextLeafTextPosition(
		this.ref,
	)
}

// FuncMoveToNextLeafTextPosition returns the method "AutomationPosition.moveToNextLeafTextPosition".
func (this AutomationPosition) FuncMoveToNextLeafTextPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextLeafTextPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextLeafTextPosition calls the method "AutomationPosition.moveToNextLeafTextPosition".
func (this AutomationPosition) MoveToNextLeafTextPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextLeafTextPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextLeafTextPosition calls the method "AutomationPosition.moveToNextLeafTextPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextLeafTextPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextLeafTextPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousLeafTextPosition returns true if the method "AutomationPosition.moveToPreviousLeafTextPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousLeafTextPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousLeafTextPosition(
		this.ref,
	)
}

// FuncMoveToPreviousLeafTextPosition returns the method "AutomationPosition.moveToPreviousLeafTextPosition".
func (this AutomationPosition) FuncMoveToPreviousLeafTextPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousLeafTextPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousLeafTextPosition calls the method "AutomationPosition.moveToPreviousLeafTextPosition".
func (this AutomationPosition) MoveToPreviousLeafTextPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousLeafTextPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousLeafTextPosition calls the method "AutomationPosition.moveToPreviousLeafTextPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousLeafTextPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousLeafTextPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextCharacterPosition returns true if the method "AutomationPosition.moveToNextCharacterPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextCharacterPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextCharacterPosition(
		this.ref,
	)
}

// FuncMoveToNextCharacterPosition returns the method "AutomationPosition.moveToNextCharacterPosition".
func (this AutomationPosition) FuncMoveToNextCharacterPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextCharacterPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextCharacterPosition calls the method "AutomationPosition.moveToNextCharacterPosition".
func (this AutomationPosition) MoveToNextCharacterPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextCharacterPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextCharacterPosition calls the method "AutomationPosition.moveToNextCharacterPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextCharacterPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextCharacterPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousCharacterPosition returns true if the method "AutomationPosition.moveToPreviousCharacterPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousCharacterPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousCharacterPosition(
		this.ref,
	)
}

// FuncMoveToPreviousCharacterPosition returns the method "AutomationPosition.moveToPreviousCharacterPosition".
func (this AutomationPosition) FuncMoveToPreviousCharacterPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousCharacterPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousCharacterPosition calls the method "AutomationPosition.moveToPreviousCharacterPosition".
func (this AutomationPosition) MoveToPreviousCharacterPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousCharacterPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousCharacterPosition calls the method "AutomationPosition.moveToPreviousCharacterPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousCharacterPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousCharacterPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextWordStartPosition returns true if the method "AutomationPosition.moveToNextWordStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextWordStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextWordStartPosition(
		this.ref,
	)
}

// FuncMoveToNextWordStartPosition returns the method "AutomationPosition.moveToNextWordStartPosition".
func (this AutomationPosition) FuncMoveToNextWordStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextWordStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextWordStartPosition calls the method "AutomationPosition.moveToNextWordStartPosition".
func (this AutomationPosition) MoveToNextWordStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextWordStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextWordStartPosition calls the method "AutomationPosition.moveToNextWordStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextWordStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextWordStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousWordStartPosition returns true if the method "AutomationPosition.moveToPreviousWordStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousWordStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousWordStartPosition(
		this.ref,
	)
}

// FuncMoveToPreviousWordStartPosition returns the method "AutomationPosition.moveToPreviousWordStartPosition".
func (this AutomationPosition) FuncMoveToPreviousWordStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousWordStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousWordStartPosition calls the method "AutomationPosition.moveToPreviousWordStartPosition".
func (this AutomationPosition) MoveToPreviousWordStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousWordStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousWordStartPosition calls the method "AutomationPosition.moveToPreviousWordStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousWordStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousWordStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextWordEndPosition returns true if the method "AutomationPosition.moveToNextWordEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextWordEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextWordEndPosition(
		this.ref,
	)
}

// FuncMoveToNextWordEndPosition returns the method "AutomationPosition.moveToNextWordEndPosition".
func (this AutomationPosition) FuncMoveToNextWordEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextWordEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextWordEndPosition calls the method "AutomationPosition.moveToNextWordEndPosition".
func (this AutomationPosition) MoveToNextWordEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextWordEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextWordEndPosition calls the method "AutomationPosition.moveToNextWordEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextWordEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextWordEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousWordEndPosition returns true if the method "AutomationPosition.moveToPreviousWordEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousWordEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousWordEndPosition(
		this.ref,
	)
}

// FuncMoveToPreviousWordEndPosition returns the method "AutomationPosition.moveToPreviousWordEndPosition".
func (this AutomationPosition) FuncMoveToPreviousWordEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousWordEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousWordEndPosition calls the method "AutomationPosition.moveToPreviousWordEndPosition".
func (this AutomationPosition) MoveToPreviousWordEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousWordEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousWordEndPosition calls the method "AutomationPosition.moveToPreviousWordEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousWordEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousWordEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextLineStartPosition returns true if the method "AutomationPosition.moveToNextLineStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextLineStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextLineStartPosition(
		this.ref,
	)
}

// FuncMoveToNextLineStartPosition returns the method "AutomationPosition.moveToNextLineStartPosition".
func (this AutomationPosition) FuncMoveToNextLineStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextLineStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextLineStartPosition calls the method "AutomationPosition.moveToNextLineStartPosition".
func (this AutomationPosition) MoveToNextLineStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextLineStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextLineStartPosition calls the method "AutomationPosition.moveToNextLineStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextLineStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextLineStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousLineStartPosition returns true if the method "AutomationPosition.moveToPreviousLineStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousLineStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousLineStartPosition(
		this.ref,
	)
}

// FuncMoveToPreviousLineStartPosition returns the method "AutomationPosition.moveToPreviousLineStartPosition".
func (this AutomationPosition) FuncMoveToPreviousLineStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousLineStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousLineStartPosition calls the method "AutomationPosition.moveToPreviousLineStartPosition".
func (this AutomationPosition) MoveToPreviousLineStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousLineStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousLineStartPosition calls the method "AutomationPosition.moveToPreviousLineStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousLineStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousLineStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextLineEndPosition returns true if the method "AutomationPosition.moveToNextLineEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextLineEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextLineEndPosition(
		this.ref,
	)
}

// FuncMoveToNextLineEndPosition returns the method "AutomationPosition.moveToNextLineEndPosition".
func (this AutomationPosition) FuncMoveToNextLineEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextLineEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextLineEndPosition calls the method "AutomationPosition.moveToNextLineEndPosition".
func (this AutomationPosition) MoveToNextLineEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextLineEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextLineEndPosition calls the method "AutomationPosition.moveToNextLineEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextLineEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextLineEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousLineEndPosition returns true if the method "AutomationPosition.moveToPreviousLineEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousLineEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousLineEndPosition(
		this.ref,
	)
}

// FuncMoveToPreviousLineEndPosition returns the method "AutomationPosition.moveToPreviousLineEndPosition".
func (this AutomationPosition) FuncMoveToPreviousLineEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousLineEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousLineEndPosition calls the method "AutomationPosition.moveToPreviousLineEndPosition".
func (this AutomationPosition) MoveToPreviousLineEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousLineEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousLineEndPosition calls the method "AutomationPosition.moveToPreviousLineEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousLineEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousLineEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextFormatStartPosition returns true if the method "AutomationPosition.moveToNextFormatStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextFormatStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextFormatStartPosition(
		this.ref,
	)
}

// FuncMoveToNextFormatStartPosition returns the method "AutomationPosition.moveToNextFormatStartPosition".
func (this AutomationPosition) FuncMoveToNextFormatStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextFormatStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextFormatStartPosition calls the method "AutomationPosition.moveToNextFormatStartPosition".
func (this AutomationPosition) MoveToNextFormatStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextFormatStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextFormatStartPosition calls the method "AutomationPosition.moveToNextFormatStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextFormatStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextFormatStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousFormatStartPosition returns true if the method "AutomationPosition.moveToPreviousFormatStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousFormatStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousFormatStartPosition(
		this.ref,
	)
}

// FuncMoveToPreviousFormatStartPosition returns the method "AutomationPosition.moveToPreviousFormatStartPosition".
func (this AutomationPosition) FuncMoveToPreviousFormatStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousFormatStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousFormatStartPosition calls the method "AutomationPosition.moveToPreviousFormatStartPosition".
func (this AutomationPosition) MoveToPreviousFormatStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousFormatStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousFormatStartPosition calls the method "AutomationPosition.moveToPreviousFormatStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousFormatStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousFormatStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextFormatEndPosition returns true if the method "AutomationPosition.moveToNextFormatEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextFormatEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextFormatEndPosition(
		this.ref,
	)
}

// FuncMoveToNextFormatEndPosition returns the method "AutomationPosition.moveToNextFormatEndPosition".
func (this AutomationPosition) FuncMoveToNextFormatEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextFormatEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextFormatEndPosition calls the method "AutomationPosition.moveToNextFormatEndPosition".
func (this AutomationPosition) MoveToNextFormatEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextFormatEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextFormatEndPosition calls the method "AutomationPosition.moveToNextFormatEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextFormatEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextFormatEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousFormatEndPosition returns true if the method "AutomationPosition.moveToPreviousFormatEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousFormatEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousFormatEndPosition(
		this.ref,
	)
}

// FuncMoveToPreviousFormatEndPosition returns the method "AutomationPosition.moveToPreviousFormatEndPosition".
func (this AutomationPosition) FuncMoveToPreviousFormatEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousFormatEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousFormatEndPosition calls the method "AutomationPosition.moveToPreviousFormatEndPosition".
func (this AutomationPosition) MoveToPreviousFormatEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousFormatEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousFormatEndPosition calls the method "AutomationPosition.moveToPreviousFormatEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousFormatEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousFormatEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextParagraphStartPosition returns true if the method "AutomationPosition.moveToNextParagraphStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextParagraphStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextParagraphStartPosition(
		this.ref,
	)
}

// FuncMoveToNextParagraphStartPosition returns the method "AutomationPosition.moveToNextParagraphStartPosition".
func (this AutomationPosition) FuncMoveToNextParagraphStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextParagraphStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextParagraphStartPosition calls the method "AutomationPosition.moveToNextParagraphStartPosition".
func (this AutomationPosition) MoveToNextParagraphStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextParagraphStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextParagraphStartPosition calls the method "AutomationPosition.moveToNextParagraphStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextParagraphStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextParagraphStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousParagraphStartPosition returns true if the method "AutomationPosition.moveToPreviousParagraphStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousParagraphStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousParagraphStartPosition(
		this.ref,
	)
}

// FuncMoveToPreviousParagraphStartPosition returns the method "AutomationPosition.moveToPreviousParagraphStartPosition".
func (this AutomationPosition) FuncMoveToPreviousParagraphStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousParagraphStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousParagraphStartPosition calls the method "AutomationPosition.moveToPreviousParagraphStartPosition".
func (this AutomationPosition) MoveToPreviousParagraphStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousParagraphStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousParagraphStartPosition calls the method "AutomationPosition.moveToPreviousParagraphStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousParagraphStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousParagraphStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextParagraphEndPosition returns true if the method "AutomationPosition.moveToNextParagraphEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextParagraphEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextParagraphEndPosition(
		this.ref,
	)
}

// FuncMoveToNextParagraphEndPosition returns the method "AutomationPosition.moveToNextParagraphEndPosition".
func (this AutomationPosition) FuncMoveToNextParagraphEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextParagraphEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextParagraphEndPosition calls the method "AutomationPosition.moveToNextParagraphEndPosition".
func (this AutomationPosition) MoveToNextParagraphEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextParagraphEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextParagraphEndPosition calls the method "AutomationPosition.moveToNextParagraphEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextParagraphEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextParagraphEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousParagraphEndPosition returns true if the method "AutomationPosition.moveToPreviousParagraphEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousParagraphEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousParagraphEndPosition(
		this.ref,
	)
}

// FuncMoveToPreviousParagraphEndPosition returns the method "AutomationPosition.moveToPreviousParagraphEndPosition".
func (this AutomationPosition) FuncMoveToPreviousParagraphEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousParagraphEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousParagraphEndPosition calls the method "AutomationPosition.moveToPreviousParagraphEndPosition".
func (this AutomationPosition) MoveToPreviousParagraphEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousParagraphEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousParagraphEndPosition calls the method "AutomationPosition.moveToPreviousParagraphEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousParagraphEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousParagraphEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextPageStartPosition returns true if the method "AutomationPosition.moveToNextPageStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextPageStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextPageStartPosition(
		this.ref,
	)
}

// FuncMoveToNextPageStartPosition returns the method "AutomationPosition.moveToNextPageStartPosition".
func (this AutomationPosition) FuncMoveToNextPageStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextPageStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextPageStartPosition calls the method "AutomationPosition.moveToNextPageStartPosition".
func (this AutomationPosition) MoveToNextPageStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextPageStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextPageStartPosition calls the method "AutomationPosition.moveToNextPageStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextPageStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextPageStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousPageStartPosition returns true if the method "AutomationPosition.moveToPreviousPageStartPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousPageStartPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousPageStartPosition(
		this.ref,
	)
}

// FuncMoveToPreviousPageStartPosition returns the method "AutomationPosition.moveToPreviousPageStartPosition".
func (this AutomationPosition) FuncMoveToPreviousPageStartPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousPageStartPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousPageStartPosition calls the method "AutomationPosition.moveToPreviousPageStartPosition".
func (this AutomationPosition) MoveToPreviousPageStartPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousPageStartPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousPageStartPosition calls the method "AutomationPosition.moveToPreviousPageStartPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousPageStartPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousPageStartPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextPageEndPosition returns true if the method "AutomationPosition.moveToNextPageEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextPageEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextPageEndPosition(
		this.ref,
	)
}

// FuncMoveToNextPageEndPosition returns the method "AutomationPosition.moveToNextPageEndPosition".
func (this AutomationPosition) FuncMoveToNextPageEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextPageEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextPageEndPosition calls the method "AutomationPosition.moveToNextPageEndPosition".
func (this AutomationPosition) MoveToNextPageEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextPageEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextPageEndPosition calls the method "AutomationPosition.moveToNextPageEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextPageEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextPageEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousPageEndPosition returns true if the method "AutomationPosition.moveToPreviousPageEndPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousPageEndPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousPageEndPosition(
		this.ref,
	)
}

// FuncMoveToPreviousPageEndPosition returns the method "AutomationPosition.moveToPreviousPageEndPosition".
func (this AutomationPosition) FuncMoveToPreviousPageEndPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousPageEndPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousPageEndPosition calls the method "AutomationPosition.moveToPreviousPageEndPosition".
func (this AutomationPosition) MoveToPreviousPageEndPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousPageEndPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousPageEndPosition calls the method "AutomationPosition.moveToPreviousPageEndPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousPageEndPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousPageEndPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToNextAnchorPosition returns true if the method "AutomationPosition.moveToNextAnchorPosition" exists.
func (this AutomationPosition) HasFuncMoveToNextAnchorPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToNextAnchorPosition(
		this.ref,
	)
}

// FuncMoveToNextAnchorPosition returns the method "AutomationPosition.moveToNextAnchorPosition".
func (this AutomationPosition) FuncMoveToNextAnchorPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToNextAnchorPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToNextAnchorPosition calls the method "AutomationPosition.moveToNextAnchorPosition".
func (this AutomationPosition) MoveToNextAnchorPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToNextAnchorPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToNextAnchorPosition calls the method "AutomationPosition.moveToNextAnchorPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToNextAnchorPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToNextAnchorPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMoveToPreviousAnchorPosition returns true if the method "AutomationPosition.moveToPreviousAnchorPosition" exists.
func (this AutomationPosition) HasFuncMoveToPreviousAnchorPosition() bool {
	return js.True == bindings.HasFuncAutomationPositionMoveToPreviousAnchorPosition(
		this.ref,
	)
}

// FuncMoveToPreviousAnchorPosition returns the method "AutomationPosition.moveToPreviousAnchorPosition".
func (this AutomationPosition) FuncMoveToPreviousAnchorPosition() (fn js.Func[func()]) {
	bindings.FuncAutomationPositionMoveToPreviousAnchorPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MoveToPreviousAnchorPosition calls the method "AutomationPosition.moveToPreviousAnchorPosition".
func (this AutomationPosition) MoveToPreviousAnchorPosition() (ret js.Void) {
	bindings.CallAutomationPositionMoveToPreviousAnchorPosition(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMoveToPreviousAnchorPosition calls the method "AutomationPosition.moveToPreviousAnchorPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMoveToPreviousAnchorPosition() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMoveToPreviousAnchorPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncMaxTextOffset returns true if the method "AutomationPosition.maxTextOffset" exists.
func (this AutomationPosition) HasFuncMaxTextOffset() bool {
	return js.True == bindings.HasFuncAutomationPositionMaxTextOffset(
		this.ref,
	)
}

// FuncMaxTextOffset returns the method "AutomationPosition.maxTextOffset".
func (this AutomationPosition) FuncMaxTextOffset() (fn js.Func[func() int32]) {
	bindings.FuncAutomationPositionMaxTextOffset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MaxTextOffset calls the method "AutomationPosition.maxTextOffset".
func (this AutomationPosition) MaxTextOffset() (ret int32) {
	bindings.CallAutomationPositionMaxTextOffset(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMaxTextOffset calls the method "AutomationPosition.maxTextOffset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryMaxTextOffset() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionMaxTextOffset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsInLineBreak returns true if the method "AutomationPosition.isInLineBreak" exists.
func (this AutomationPosition) HasFuncIsInLineBreak() bool {
	return js.True == bindings.HasFuncAutomationPositionIsInLineBreak(
		this.ref,
	)
}

// FuncIsInLineBreak returns the method "AutomationPosition.isInLineBreak".
func (this AutomationPosition) FuncIsInLineBreak() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionIsInLineBreak(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsInLineBreak calls the method "AutomationPosition.isInLineBreak".
func (this AutomationPosition) IsInLineBreak() (ret bool) {
	bindings.CallAutomationPositionIsInLineBreak(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsInLineBreak calls the method "AutomationPosition.isInLineBreak"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryIsInLineBreak() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionIsInLineBreak(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsInTextObject returns true if the method "AutomationPosition.isInTextObject" exists.
func (this AutomationPosition) HasFuncIsInTextObject() bool {
	return js.True == bindings.HasFuncAutomationPositionIsInTextObject(
		this.ref,
	)
}

// FuncIsInTextObject returns the method "AutomationPosition.isInTextObject".
func (this AutomationPosition) FuncIsInTextObject() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionIsInTextObject(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsInTextObject calls the method "AutomationPosition.isInTextObject".
func (this AutomationPosition) IsInTextObject() (ret bool) {
	bindings.CallAutomationPositionIsInTextObject(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsInTextObject calls the method "AutomationPosition.isInTextObject"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryIsInTextObject() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionIsInTextObject(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsInWhiteSpace returns true if the method "AutomationPosition.isInWhiteSpace" exists.
func (this AutomationPosition) HasFuncIsInWhiteSpace() bool {
	return js.True == bindings.HasFuncAutomationPositionIsInWhiteSpace(
		this.ref,
	)
}

// FuncIsInWhiteSpace returns the method "AutomationPosition.isInWhiteSpace".
func (this AutomationPosition) FuncIsInWhiteSpace() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionIsInWhiteSpace(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsInWhiteSpace calls the method "AutomationPosition.isInWhiteSpace".
func (this AutomationPosition) IsInWhiteSpace() (ret bool) {
	bindings.CallAutomationPositionIsInWhiteSpace(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsInWhiteSpace calls the method "AutomationPosition.isInWhiteSpace"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryIsInWhiteSpace() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionIsInWhiteSpace(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsValid returns true if the method "AutomationPosition.isValid" exists.
func (this AutomationPosition) HasFuncIsValid() bool {
	return js.True == bindings.HasFuncAutomationPositionIsValid(
		this.ref,
	)
}

// FuncIsValid returns the method "AutomationPosition.isValid".
func (this AutomationPosition) FuncIsValid() (fn js.Func[func() bool]) {
	bindings.FuncAutomationPositionIsValid(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsValid calls the method "AutomationPosition.isValid".
func (this AutomationPosition) IsValid() (ret bool) {
	bindings.CallAutomationPositionIsValid(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsValid calls the method "AutomationPosition.isValid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryIsValid() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionIsValid(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetText returns true if the method "AutomationPosition.getText" exists.
func (this AutomationPosition) HasFuncGetText() bool {
	return js.True == bindings.HasFuncAutomationPositionGetText(
		this.ref,
	)
}

// FuncGetText returns the method "AutomationPosition.getText".
func (this AutomationPosition) FuncGetText() (fn js.Func[func() js.String]) {
	bindings.FuncAutomationPositionGetText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetText calls the method "AutomationPosition.getText".
func (this AutomationPosition) GetText() (ret js.String) {
	bindings.CallAutomationPositionGetText(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetText calls the method "AutomationPosition.getText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationPosition) TryGetText() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationPositionGetText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PositionType uint32

const (
	_ PositionType = iota

	PositionType_NULL
	PositionType_TEXT
	PositionType_TREE
)

func (PositionType) FromRef(str js.Ref) PositionType {
	return PositionType(bindings.ConstOfPositionType(str))
}

func (x PositionType) String() (string, bool) {
	switch x {
	case PositionType_NULL:
		return "null", true
	case PositionType_TEXT:
		return "text", true
	case PositionType_TREE:
		return "tree", true
	default:
		return "", false
	}
}

type NameFromType uint32

const (
	_ NameFromType = iota

	NameFromType_ATTRIBUTE
	NameFromType_ATTRIBUTE_EXPLICITLY_EMPTY
	NameFromType_CAPTION
	NameFromType_CONTENTS
	NameFromType_PLACEHOLDER
	NameFromType_POPOVER_ATTRIBUTE
	NameFromType_RELATED_ELEMENT
	NameFromType_TITLE
	NameFromType_VALUE
)

func (NameFromType) FromRef(str js.Ref) NameFromType {
	return NameFromType(bindings.ConstOfNameFromType(str))
}

func (x NameFromType) String() (string, bool) {
	switch x {
	case NameFromType_ATTRIBUTE:
		return "attribute", true
	case NameFromType_ATTRIBUTE_EXPLICITLY_EMPTY:
		return "attributeExplicitlyEmpty", true
	case NameFromType_CAPTION:
		return "caption", true
	case NameFromType_CONTENTS:
		return "contents", true
	case NameFromType_PLACEHOLDER:
		return "placeholder", true
	case NameFromType_POPOVER_ATTRIBUTE:
		return "popoverAttribute", true
	case NameFromType_RELATED_ELEMENT:
		return "relatedElement", true
	case NameFromType_TITLE:
		return "title", true
	case NameFromType_VALUE:
		return "value", true
	default:
		return "", false
	}
}

type CustomAction struct {
	// Id is "CustomAction.id"
	//
	// Optional
	//
	// NOTE: FFI_USE_Id MUST be set to true to make this field effective.
	Id int32
	// Description is "CustomAction.description"
	//
	// Optional
	Description js.String

	FFI_USE_Id bool // for Id.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CustomAction with all fields set.
func (p CustomAction) FromRef(ref js.Ref) CustomAction {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CustomAction in the application heap.
func (p CustomAction) New() js.Ref {
	return bindings.CustomActionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CustomAction) UpdateFrom(ref js.Ref) {
	bindings.CustomActionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CustomAction) Update(ref js.Ref) {
	bindings.CustomActionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CustomAction) FreeMembers(recursive bool) {
	js.Free(
		p.Description.Ref(),
	)
	p.Description = p.Description.FromRef(js.Undefined)
}

type DefaultActionVerb uint32

const (
	_ DefaultActionVerb = iota

	DefaultActionVerb_ACTIVATE
	DefaultActionVerb_CHECK
	DefaultActionVerb_CLICK
	DefaultActionVerb_CLICK_ANCESTOR
	DefaultActionVerb_JUMP
	DefaultActionVerb_OPEN
	DefaultActionVerb_PRESS
	DefaultActionVerb_SELECT
	DefaultActionVerb_UNCHECK
)

func (DefaultActionVerb) FromRef(str js.Ref) DefaultActionVerb {
	return DefaultActionVerb(bindings.ConstOfDefaultActionVerb(str))
}

func (x DefaultActionVerb) String() (string, bool) {
	switch x {
	case DefaultActionVerb_ACTIVATE:
		return "activate", true
	case DefaultActionVerb_CHECK:
		return "check", true
	case DefaultActionVerb_CLICK:
		return "click", true
	case DefaultActionVerb_CLICK_ANCESTOR:
		return "clickAncestor", true
	case DefaultActionVerb_JUMP:
		return "jump", true
	case DefaultActionVerb_OPEN:
		return "open", true
	case DefaultActionVerb_PRESS:
		return "press", true
	case DefaultActionVerb_SELECT:
		return "select", true
	case DefaultActionVerb_UNCHECK:
		return "uncheck", true
	default:
		return "", false
	}
}

type Marker struct {
	// StartOffset is "Marker.startOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_StartOffset MUST be set to true to make this field effective.
	StartOffset int32
	// EndOffset is "Marker.endOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_EndOffset MUST be set to true to make this field effective.
	EndOffset int32
	// Flags is "Marker.flags"
	//
	// Optional
	Flags js.Object

	FFI_USE_StartOffset bool // for StartOffset.
	FFI_USE_EndOffset   bool // for EndOffset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Marker with all fields set.
func (p Marker) FromRef(ref js.Ref) Marker {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Marker in the application heap.
func (p Marker) New() js.Ref {
	return bindings.MarkerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Marker) UpdateFrom(ref js.Ref) {
	bindings.MarkerJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Marker) Update(ref js.Ref) {
	bindings.MarkerJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Marker) FreeMembers(recursive bool) {
	js.Free(
		p.Flags.Ref(),
	)
	p.Flags = p.Flags.FromRef(js.Undefined)
}

type HasPopup uint32

const (
	_ HasPopup = iota

	HasPopup_FALSE
	HasPopup_TRUE
	HasPopup_MENU
	HasPopup_LISTBOX
	HasPopup_TREE
	HasPopup_GRID
	HasPopup_DIALOG
)

func (HasPopup) FromRef(str js.Ref) HasPopup {
	return HasPopup(bindings.ConstOfHasPopup(str))
}

func (x HasPopup) String() (string, bool) {
	switch x {
	case HasPopup_FALSE:
		return "false", true
	case HasPopup_TRUE:
		return "true", true
	case HasPopup_MENU:
		return "menu", true
	case HasPopup_LISTBOX:
		return "listbox", true
	case HasPopup_TREE:
		return "tree", true
	case HasPopup_GRID:
		return "grid", true
	case HasPopup_DIALOG:
		return "dialog", true
	default:
		return "", false
	}
}

type AriaCurrentState uint32

const (
	_ AriaCurrentState = iota

	AriaCurrentState_FALSE
	AriaCurrentState_TRUE
	AriaCurrentState_PAGE
	AriaCurrentState_STEP
	AriaCurrentState_LOCATION
	AriaCurrentState_DATE
	AriaCurrentState_TIME
)

func (AriaCurrentState) FromRef(str js.Ref) AriaCurrentState {
	return AriaCurrentState(bindings.ConstOfAriaCurrentState(str))
}

func (x AriaCurrentState) String() (string, bool) {
	switch x {
	case AriaCurrentState_FALSE:
		return "false", true
	case AriaCurrentState_TRUE:
		return "true", true
	case AriaCurrentState_PAGE:
		return "page", true
	case AriaCurrentState_STEP:
		return "step", true
	case AriaCurrentState_LOCATION:
		return "location", true
	case AriaCurrentState_DATE:
		return "date", true
	case AriaCurrentState_TIME:
		return "time", true
	default:
		return "", false
	}
}

type InvalidState uint32

const (
	_ InvalidState = iota

	InvalidState_FALSE
	InvalidState_TRUE
)

func (InvalidState) FromRef(str js.Ref) InvalidState {
	return InvalidState(bindings.ConstOfInvalidState(str))
}

func (x InvalidState) String() (string, bool) {
	switch x {
	case InvalidState_FALSE:
		return "false", true
	case InvalidState_TRUE:
		return "true", true
	default:
		return "", false
	}
}

type SortDirectionType uint32

const (
	_ SortDirectionType = iota

	SortDirectionType_UNSORTED
	SortDirectionType_ASCENDING
	SortDirectionType_DESCENDING
	SortDirectionType_OTHER
)

func (SortDirectionType) FromRef(str js.Ref) SortDirectionType {
	return SortDirectionType(bindings.ConstOfSortDirectionType(str))
}

func (x SortDirectionType) String() (string, bool) {
	switch x {
	case SortDirectionType_UNSORTED:
		return "unsorted", true
	case SortDirectionType_ASCENDING:
		return "ascending", true
	case SortDirectionType_DESCENDING:
		return "descending", true
	case SortDirectionType_OTHER:
		return "other", true
	default:
		return "", false
	}
}

type AutomationNode struct {
	ref js.Ref
}

func (this AutomationNode) Once() AutomationNode {
	this.ref.Once()
	return this
}

func (this AutomationNode) Ref() js.Ref {
	return this.ref
}

func (this AutomationNode) FromRef(ref js.Ref) AutomationNode {
	this.ref = ref
	return this
}

func (this AutomationNode) Free() {
	this.ref.Free()
}

// Root returns the value of property "AutomationNode.root".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Root() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeRoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRoot sets the value of property "AutomationNode.root" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetRoot(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeRoot(
		this.ref,
		val.Ref(),
	)
}

// IsRootNode returns the value of property "AutomationNode.isRootNode".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) IsRootNode() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeIsRootNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIsRootNode sets the value of property "AutomationNode.isRootNode" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetIsRootNode(val bool) bool {
	return js.True == bindings.SetAutomationNodeIsRootNode(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Role returns the value of property "AutomationNode.role".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Role() (ret RoleType, ok bool) {
	ok = js.True == bindings.GetAutomationNodeRole(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRole sets the value of property "AutomationNode.role" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetRole(val RoleType) bool {
	return js.True == bindings.SetAutomationNodeRole(
		this.ref,
		uint32(val),
	)
}

// State returns the value of property "AutomationNode.state".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) State() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetAutomationNodeState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetState sets the value of property "AutomationNode.state" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetState(val js.Object) bool {
	return js.True == bindings.SetAutomationNodeState(
		this.ref,
		val.Ref(),
	)
}

// Location returns the value of property "AutomationNode.location".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Location() (ret Rect, ok bool) {
	ok = js.True == bindings.GetAutomationNodeLocation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLocation sets the value of property "AutomationNode.location" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLocation(val Rect) bool {
	return js.True == bindings.SetAutomationNodeLocation(
		this.ref,
		js.Pointer(&val),
	)
}

// UnclippedLocation returns the value of property "AutomationNode.unclippedLocation".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) UnclippedLocation() (ret Rect, ok bool) {
	ok = js.True == bindings.GetAutomationNodeUnclippedLocation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUnclippedLocation sets the value of property "AutomationNode.unclippedLocation" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetUnclippedLocation(val Rect) bool {
	return js.True == bindings.SetAutomationNodeUnclippedLocation(
		this.ref,
		js.Pointer(&val),
	)
}

// Description returns the value of property "AutomationNode.description".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Description() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDescription sets the value of property "AutomationNode.description" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDescription(val js.String) bool {
	return js.True == bindings.SetAutomationNodeDescription(
		this.ref,
		val.Ref(),
	)
}

// CheckedStateDescription returns the value of property "AutomationNode.checkedStateDescription".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) CheckedStateDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeCheckedStateDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCheckedStateDescription sets the value of property "AutomationNode.checkedStateDescription" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetCheckedStateDescription(val js.String) bool {
	return js.True == bindings.SetAutomationNodeCheckedStateDescription(
		this.ref,
		val.Ref(),
	)
}

// Placeholder returns the value of property "AutomationNode.placeholder".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Placeholder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodePlaceholder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPlaceholder sets the value of property "AutomationNode.placeholder" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetPlaceholder(val js.String) bool {
	return js.True == bindings.SetAutomationNodePlaceholder(
		this.ref,
		val.Ref(),
	)
}

// RoleDescription returns the value of property "AutomationNode.roleDescription".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) RoleDescription() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeRoleDescription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRoleDescription sets the value of property "AutomationNode.roleDescription" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetRoleDescription(val js.String) bool {
	return js.True == bindings.SetAutomationNodeRoleDescription(
		this.ref,
		val.Ref(),
	)
}

// Name returns the value of property "AutomationNode.name".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "AutomationNode.name" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetName(val js.String) bool {
	return js.True == bindings.SetAutomationNodeName(
		this.ref,
		val.Ref(),
	)
}

// DoDefaultLabel returns the value of property "AutomationNode.doDefaultLabel".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DoDefaultLabel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeDoDefaultLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDoDefaultLabel sets the value of property "AutomationNode.doDefaultLabel" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDoDefaultLabel(val js.String) bool {
	return js.True == bindings.SetAutomationNodeDoDefaultLabel(
		this.ref,
		val.Ref(),
	)
}

// LongClickLabel returns the value of property "AutomationNode.longClickLabel".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) LongClickLabel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeLongClickLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLongClickLabel sets the value of property "AutomationNode.longClickLabel" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLongClickLabel(val js.String) bool {
	return js.True == bindings.SetAutomationNodeLongClickLabel(
		this.ref,
		val.Ref(),
	)
}

// Tooltip returns the value of property "AutomationNode.tooltip".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Tooltip() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTooltip(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTooltip sets the value of property "AutomationNode.tooltip" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTooltip(val js.String) bool {
	return js.True == bindings.SetAutomationNodeTooltip(
		this.ref,
		val.Ref(),
	)
}

// NameFrom returns the value of property "AutomationNode.nameFrom".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) NameFrom() (ret NameFromType, ok bool) {
	ok = js.True == bindings.GetAutomationNodeNameFrom(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNameFrom sets the value of property "AutomationNode.nameFrom" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetNameFrom(val NameFromType) bool {
	return js.True == bindings.SetAutomationNodeNameFrom(
		this.ref,
		uint32(val),
	)
}

// ImageAnnotation returns the value of property "AutomationNode.imageAnnotation".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ImageAnnotation() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeImageAnnotation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetImageAnnotation sets the value of property "AutomationNode.imageAnnotation" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetImageAnnotation(val js.String) bool {
	return js.True == bindings.SetAutomationNodeImageAnnotation(
		this.ref,
		val.Ref(),
	)
}

// Value returns the value of property "AutomationNode.value".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HtmlTag returns the value of property "AutomationNode.htmlTag".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) HtmlTag() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeHtmlTag(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHtmlTag sets the value of property "AutomationNode.htmlTag" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetHtmlTag(val js.String) bool {
	return js.True == bindings.SetAutomationNodeHtmlTag(
		this.ref,
		val.Ref(),
	)
}

// HierarchicalLevel returns the value of property "AutomationNode.hierarchicalLevel".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) HierarchicalLevel() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeHierarchicalLevel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHierarchicalLevel sets the value of property "AutomationNode.hierarchicalLevel" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetHierarchicalLevel(val int32) bool {
	return js.True == bindings.SetAutomationNodeHierarchicalLevel(
		this.ref,
		int32(val),
	)
}

// CaretBounds returns the value of property "AutomationNode.caretBounds".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) CaretBounds() (ret Rect, ok bool) {
	ok = js.True == bindings.GetAutomationNodeCaretBounds(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCaretBounds sets the value of property "AutomationNode.caretBounds" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetCaretBounds(val Rect) bool {
	return js.True == bindings.SetAutomationNodeCaretBounds(
		this.ref,
		js.Pointer(&val),
	)
}

// WordStarts returns the value of property "AutomationNode.wordStarts".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) WordStarts() (ret js.Array[int32], ok bool) {
	ok = js.True == bindings.GetAutomationNodeWordStarts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWordStarts sets the value of property "AutomationNode.wordStarts" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetWordStarts(val js.Array[int32]) bool {
	return js.True == bindings.SetAutomationNodeWordStarts(
		this.ref,
		val.Ref(),
	)
}

// WordEnds returns the value of property "AutomationNode.wordEnds".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) WordEnds() (ret js.Array[int32], ok bool) {
	ok = js.True == bindings.GetAutomationNodeWordEnds(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWordEnds sets the value of property "AutomationNode.wordEnds" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetWordEnds(val js.Array[int32]) bool {
	return js.True == bindings.SetAutomationNodeWordEnds(
		this.ref,
		val.Ref(),
	)
}

// SentenceStarts returns the value of property "AutomationNode.sentenceStarts".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SentenceStarts() (ret js.Array[int32], ok bool) {
	ok = js.True == bindings.GetAutomationNodeSentenceStarts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSentenceStarts sets the value of property "AutomationNode.sentenceStarts" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSentenceStarts(val js.Array[int32]) bool {
	return js.True == bindings.SetAutomationNodeSentenceStarts(
		this.ref,
		val.Ref(),
	)
}

// SentenceEnds returns the value of property "AutomationNode.sentenceEnds".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SentenceEnds() (ret js.Array[int32], ok bool) {
	ok = js.True == bindings.GetAutomationNodeSentenceEnds(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSentenceEnds sets the value of property "AutomationNode.sentenceEnds" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSentenceEnds(val js.Array[int32]) bool {
	return js.True == bindings.SetAutomationNodeSentenceEnds(
		this.ref,
		val.Ref(),
	)
}

// NonInlineTextWordStarts returns the value of property "AutomationNode.nonInlineTextWordStarts".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) NonInlineTextWordStarts() (ret js.Array[int32], ok bool) {
	ok = js.True == bindings.GetAutomationNodeNonInlineTextWordStarts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNonInlineTextWordStarts sets the value of property "AutomationNode.nonInlineTextWordStarts" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetNonInlineTextWordStarts(val js.Array[int32]) bool {
	return js.True == bindings.SetAutomationNodeNonInlineTextWordStarts(
		this.ref,
		val.Ref(),
	)
}

// NonInlineTextWordEnds returns the value of property "AutomationNode.nonInlineTextWordEnds".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) NonInlineTextWordEnds() (ret js.Array[int32], ok bool) {
	ok = js.True == bindings.GetAutomationNodeNonInlineTextWordEnds(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNonInlineTextWordEnds sets the value of property "AutomationNode.nonInlineTextWordEnds" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetNonInlineTextWordEnds(val js.Array[int32]) bool {
	return js.True == bindings.SetAutomationNodeNonInlineTextWordEnds(
		this.ref,
		val.Ref(),
	)
}

// Controls returns the value of property "AutomationNode.controls".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Controls() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeControls(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetControls sets the value of property "AutomationNode.controls" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetControls(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeControls(
		this.ref,
		val.Ref(),
	)
}

// DescribedBy returns the value of property "AutomationNode.describedBy".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DescribedBy() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeDescribedBy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDescribedBy sets the value of property "AutomationNode.describedBy" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDescribedBy(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeDescribedBy(
		this.ref,
		val.Ref(),
	)
}

// FlowTo returns the value of property "AutomationNode.flowTo".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) FlowTo() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeFlowTo(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFlowTo sets the value of property "AutomationNode.flowTo" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetFlowTo(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeFlowTo(
		this.ref,
		val.Ref(),
	)
}

// LabelledBy returns the value of property "AutomationNode.labelledBy".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) LabelledBy() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeLabelledBy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabelledBy sets the value of property "AutomationNode.labelledBy" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLabelledBy(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeLabelledBy(
		this.ref,
		val.Ref(),
	)
}

// ActiveDescendant returns the value of property "AutomationNode.activeDescendant".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ActiveDescendant() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeActiveDescendant(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetActiveDescendant sets the value of property "AutomationNode.activeDescendant" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetActiveDescendant(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeActiveDescendant(
		this.ref,
		val.Ref(),
	)
}

// ActiveDescendantFor returns the value of property "AutomationNode.activeDescendantFor".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ActiveDescendantFor() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeActiveDescendantFor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetActiveDescendantFor sets the value of property "AutomationNode.activeDescendantFor" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetActiveDescendantFor(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeActiveDescendantFor(
		this.ref,
		val.Ref(),
	)
}

// InPageLinkTarget returns the value of property "AutomationNode.inPageLinkTarget".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) InPageLinkTarget() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeInPageLinkTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInPageLinkTarget sets the value of property "AutomationNode.inPageLinkTarget" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetInPageLinkTarget(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeInPageLinkTarget(
		this.ref,
		val.Ref(),
	)
}

// Details returns the value of property "AutomationNode.details".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Details() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeDetails(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDetails sets the value of property "AutomationNode.details" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDetails(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeDetails(
		this.ref,
		val.Ref(),
	)
}

// ErrorMessages returns the value of property "AutomationNode.errorMessages".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ErrorMessages() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeErrorMessages(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetErrorMessages sets the value of property "AutomationNode.errorMessages" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetErrorMessages(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeErrorMessages(
		this.ref,
		val.Ref(),
	)
}

// DetailsFor returns the value of property "AutomationNode.detailsFor".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DetailsFor() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeDetailsFor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDetailsFor sets the value of property "AutomationNode.detailsFor" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDetailsFor(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeDetailsFor(
		this.ref,
		val.Ref(),
	)
}

// ErrorMessageFor returns the value of property "AutomationNode.errorMessageFor".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ErrorMessageFor() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeErrorMessageFor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetErrorMessageFor sets the value of property "AutomationNode.errorMessageFor" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetErrorMessageFor(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeErrorMessageFor(
		this.ref,
		val.Ref(),
	)
}

// ControlledBy returns the value of property "AutomationNode.controlledBy".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ControlledBy() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeControlledBy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetControlledBy sets the value of property "AutomationNode.controlledBy" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetControlledBy(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeControlledBy(
		this.ref,
		val.Ref(),
	)
}

// DescriptionFor returns the value of property "AutomationNode.descriptionFor".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DescriptionFor() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeDescriptionFor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDescriptionFor sets the value of property "AutomationNode.descriptionFor" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDescriptionFor(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeDescriptionFor(
		this.ref,
		val.Ref(),
	)
}

// FlowFrom returns the value of property "AutomationNode.flowFrom".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) FlowFrom() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeFlowFrom(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFlowFrom sets the value of property "AutomationNode.flowFrom" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetFlowFrom(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeFlowFrom(
		this.ref,
		val.Ref(),
	)
}

// LabelFor returns the value of property "AutomationNode.labelFor".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) LabelFor() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeLabelFor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabelFor sets the value of property "AutomationNode.labelFor" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLabelFor(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeLabelFor(
		this.ref,
		val.Ref(),
	)
}

// TableCellColumnHeaders returns the value of property "AutomationNode.tableCellColumnHeaders".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableCellColumnHeaders() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableCellColumnHeaders(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableCellColumnHeaders sets the value of property "AutomationNode.tableCellColumnHeaders" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableCellColumnHeaders(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeTableCellColumnHeaders(
		this.ref,
		val.Ref(),
	)
}

// TableCellRowHeaders returns the value of property "AutomationNode.tableCellRowHeaders".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableCellRowHeaders() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableCellRowHeaders(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableCellRowHeaders sets the value of property "AutomationNode.tableCellRowHeaders" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableCellRowHeaders(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeTableCellRowHeaders(
		this.ref,
		val.Ref(),
	)
}

// StandardActions returns the value of property "AutomationNode.standardActions".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) StandardActions() (ret js.Array[ActionType], ok bool) {
	ok = js.True == bindings.GetAutomationNodeStandardActions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStandardActions sets the value of property "AutomationNode.standardActions" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetStandardActions(val js.Array[ActionType]) bool {
	return js.True == bindings.SetAutomationNodeStandardActions(
		this.ref,
		val.Ref(),
	)
}

// CustomActions returns the value of property "AutomationNode.customActions".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) CustomActions() (ret js.Array[CustomAction], ok bool) {
	ok = js.True == bindings.GetAutomationNodeCustomActions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCustomActions sets the value of property "AutomationNode.customActions" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetCustomActions(val js.Array[CustomAction]) bool {
	return js.True == bindings.SetAutomationNodeCustomActions(
		this.ref,
		val.Ref(),
	)
}

// DefaultActionVerb returns the value of property "AutomationNode.defaultActionVerb".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DefaultActionVerb() (ret DefaultActionVerb, ok bool) {
	ok = js.True == bindings.GetAutomationNodeDefaultActionVerb(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDefaultActionVerb sets the value of property "AutomationNode.defaultActionVerb" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDefaultActionVerb(val DefaultActionVerb) bool {
	return js.True == bindings.SetAutomationNodeDefaultActionVerb(
		this.ref,
		uint32(val),
	)
}

// Url returns the value of property "AutomationNode.url".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUrl sets the value of property "AutomationNode.url" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetUrl(val js.String) bool {
	return js.True == bindings.SetAutomationNodeUrl(
		this.ref,
		val.Ref(),
	)
}

// DocUrl returns the value of property "AutomationNode.docUrl".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DocUrl() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeDocUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDocUrl sets the value of property "AutomationNode.docUrl" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDocUrl(val js.String) bool {
	return js.True == bindings.SetAutomationNodeDocUrl(
		this.ref,
		val.Ref(),
	)
}

// DocTitle returns the value of property "AutomationNode.docTitle".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DocTitle() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeDocTitle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDocTitle sets the value of property "AutomationNode.docTitle" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDocTitle(val js.String) bool {
	return js.True == bindings.SetAutomationNodeDocTitle(
		this.ref,
		val.Ref(),
	)
}

// DocLoaded returns the value of property "AutomationNode.docLoaded".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DocLoaded() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeDocLoaded(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDocLoaded sets the value of property "AutomationNode.docLoaded" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDocLoaded(val bool) bool {
	return js.True == bindings.SetAutomationNodeDocLoaded(
		this.ref,
		js.Bool(bool(val)),
	)
}

// DocLoadingProgress returns the value of property "AutomationNode.docLoadingProgress".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DocLoadingProgress() (ret float64, ok bool) {
	ok = js.True == bindings.GetAutomationNodeDocLoadingProgress(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDocLoadingProgress sets the value of property "AutomationNode.docLoadingProgress" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDocLoadingProgress(val float64) bool {
	return js.True == bindings.SetAutomationNodeDocLoadingProgress(
		this.ref,
		float64(val),
	)
}

// ScrollX returns the value of property "AutomationNode.scrollX".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ScrollX() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeScrollX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollX sets the value of property "AutomationNode.scrollX" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetScrollX(val int32) bool {
	return js.True == bindings.SetAutomationNodeScrollX(
		this.ref,
		int32(val),
	)
}

// ScrollXMin returns the value of property "AutomationNode.scrollXMin".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ScrollXMin() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeScrollXMin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollXMin sets the value of property "AutomationNode.scrollXMin" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetScrollXMin(val int32) bool {
	return js.True == bindings.SetAutomationNodeScrollXMin(
		this.ref,
		int32(val),
	)
}

// ScrollXMax returns the value of property "AutomationNode.scrollXMax".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ScrollXMax() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeScrollXMax(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollXMax sets the value of property "AutomationNode.scrollXMax" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetScrollXMax(val int32) bool {
	return js.True == bindings.SetAutomationNodeScrollXMax(
		this.ref,
		int32(val),
	)
}

// ScrollY returns the value of property "AutomationNode.scrollY".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ScrollY() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeScrollY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollY sets the value of property "AutomationNode.scrollY" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetScrollY(val int32) bool {
	return js.True == bindings.SetAutomationNodeScrollY(
		this.ref,
		int32(val),
	)
}

// ScrollYMin returns the value of property "AutomationNode.scrollYMin".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ScrollYMin() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeScrollYMin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollYMin sets the value of property "AutomationNode.scrollYMin" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetScrollYMin(val int32) bool {
	return js.True == bindings.SetAutomationNodeScrollYMin(
		this.ref,
		int32(val),
	)
}

// ScrollYMax returns the value of property "AutomationNode.scrollYMax".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ScrollYMax() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeScrollYMax(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollYMax sets the value of property "AutomationNode.scrollYMax" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetScrollYMax(val int32) bool {
	return js.True == bindings.SetAutomationNodeScrollYMax(
		this.ref,
		int32(val),
	)
}

// Scrollable returns the value of property "AutomationNode.scrollable".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Scrollable() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeScrollable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollable sets the value of property "AutomationNode.scrollable" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetScrollable(val bool) bool {
	return js.True == bindings.SetAutomationNodeScrollable(
		this.ref,
		js.Bool(bool(val)),
	)
}

// TextSelStart returns the value of property "AutomationNode.textSelStart".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TextSelStart() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTextSelStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTextSelStart sets the value of property "AutomationNode.textSelStart" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTextSelStart(val int32) bool {
	return js.True == bindings.SetAutomationNodeTextSelStart(
		this.ref,
		int32(val),
	)
}

// TextSelEnd returns the value of property "AutomationNode.textSelEnd".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TextSelEnd() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTextSelEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTextSelEnd sets the value of property "AutomationNode.textSelEnd" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTextSelEnd(val int32) bool {
	return js.True == bindings.SetAutomationNodeTextSelEnd(
		this.ref,
		int32(val),
	)
}

// Markers returns the value of property "AutomationNode.markers".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Markers() (ret js.Array[Marker], ok bool) {
	ok = js.True == bindings.GetAutomationNodeMarkers(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMarkers sets the value of property "AutomationNode.markers" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetMarkers(val js.Array[Marker]) bool {
	return js.True == bindings.SetAutomationNodeMarkers(
		this.ref,
		val.Ref(),
	)
}

// IsSelectionBackward returns the value of property "AutomationNode.isSelectionBackward".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) IsSelectionBackward() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeIsSelectionBackward(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIsSelectionBackward sets the value of property "AutomationNode.isSelectionBackward" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetIsSelectionBackward(val bool) bool {
	return js.True == bindings.SetAutomationNodeIsSelectionBackward(
		this.ref,
		js.Bool(bool(val)),
	)
}

// AnchorObject returns the value of property "AutomationNode.anchorObject".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AnchorObject() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAnchorObject(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAnchorObject sets the value of property "AutomationNode.anchorObject" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAnchorObject(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeAnchorObject(
		this.ref,
		val.Ref(),
	)
}

// AnchorOffset returns the value of property "AutomationNode.anchorOffset".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AnchorOffset() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAnchorOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAnchorOffset sets the value of property "AutomationNode.anchorOffset" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAnchorOffset(val int32) bool {
	return js.True == bindings.SetAutomationNodeAnchorOffset(
		this.ref,
		int32(val),
	)
}

// AnchorAffinity returns the value of property "AutomationNode.anchorAffinity".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AnchorAffinity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAnchorAffinity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAnchorAffinity sets the value of property "AutomationNode.anchorAffinity" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAnchorAffinity(val js.String) bool {
	return js.True == bindings.SetAutomationNodeAnchorAffinity(
		this.ref,
		val.Ref(),
	)
}

// FocusObject returns the value of property "AutomationNode.focusObject".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) FocusObject() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeFocusObject(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFocusObject sets the value of property "AutomationNode.focusObject" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetFocusObject(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeFocusObject(
		this.ref,
		val.Ref(),
	)
}

// FocusOffset returns the value of property "AutomationNode.focusOffset".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) FocusOffset() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeFocusOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFocusOffset sets the value of property "AutomationNode.focusOffset" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetFocusOffset(val int32) bool {
	return js.True == bindings.SetAutomationNodeFocusOffset(
		this.ref,
		int32(val),
	)
}

// FocusAffinity returns the value of property "AutomationNode.focusAffinity".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) FocusAffinity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeFocusAffinity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFocusAffinity sets the value of property "AutomationNode.focusAffinity" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetFocusAffinity(val js.String) bool {
	return js.True == bindings.SetAutomationNodeFocusAffinity(
		this.ref,
		val.Ref(),
	)
}

// SelectionStartObject returns the value of property "AutomationNode.selectionStartObject".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SelectionStartObject() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSelectionStartObject(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionStartObject sets the value of property "AutomationNode.selectionStartObject" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSelectionStartObject(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeSelectionStartObject(
		this.ref,
		val.Ref(),
	)
}

// SelectionStartOffset returns the value of property "AutomationNode.selectionStartOffset".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SelectionStartOffset() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSelectionStartOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionStartOffset sets the value of property "AutomationNode.selectionStartOffset" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSelectionStartOffset(val int32) bool {
	return js.True == bindings.SetAutomationNodeSelectionStartOffset(
		this.ref,
		int32(val),
	)
}

// SelectionStartAffinity returns the value of property "AutomationNode.selectionStartAffinity".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SelectionStartAffinity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSelectionStartAffinity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionStartAffinity sets the value of property "AutomationNode.selectionStartAffinity" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSelectionStartAffinity(val js.String) bool {
	return js.True == bindings.SetAutomationNodeSelectionStartAffinity(
		this.ref,
		val.Ref(),
	)
}

// SelectionEndObject returns the value of property "AutomationNode.selectionEndObject".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SelectionEndObject() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSelectionEndObject(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionEndObject sets the value of property "AutomationNode.selectionEndObject" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSelectionEndObject(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeSelectionEndObject(
		this.ref,
		val.Ref(),
	)
}

// SelectionEndOffset returns the value of property "AutomationNode.selectionEndOffset".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SelectionEndOffset() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSelectionEndOffset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionEndOffset sets the value of property "AutomationNode.selectionEndOffset" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSelectionEndOffset(val int32) bool {
	return js.True == bindings.SetAutomationNodeSelectionEndOffset(
		this.ref,
		int32(val),
	)
}

// SelectionEndAffinity returns the value of property "AutomationNode.selectionEndAffinity".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SelectionEndAffinity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSelectionEndAffinity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionEndAffinity sets the value of property "AutomationNode.selectionEndAffinity" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSelectionEndAffinity(val js.String) bool {
	return js.True == bindings.SetAutomationNodeSelectionEndAffinity(
		this.ref,
		val.Ref(),
	)
}

// NotUserSelectableStyle returns the value of property "AutomationNode.notUserSelectableStyle".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) NotUserSelectableStyle() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeNotUserSelectableStyle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNotUserSelectableStyle sets the value of property "AutomationNode.notUserSelectableStyle" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetNotUserSelectableStyle(val bool) bool {
	return js.True == bindings.SetAutomationNodeNotUserSelectableStyle(
		this.ref,
		js.Bool(bool(val)),
	)
}

// ValueForRange returns the value of property "AutomationNode.valueForRange".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ValueForRange() (ret float64, ok bool) {
	ok = js.True == bindings.GetAutomationNodeValueForRange(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValueForRange sets the value of property "AutomationNode.valueForRange" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetValueForRange(val float64) bool {
	return js.True == bindings.SetAutomationNodeValueForRange(
		this.ref,
		float64(val),
	)
}

// MinValueForRange returns the value of property "AutomationNode.minValueForRange".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) MinValueForRange() (ret float64, ok bool) {
	ok = js.True == bindings.GetAutomationNodeMinValueForRange(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMinValueForRange sets the value of property "AutomationNode.minValueForRange" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetMinValueForRange(val float64) bool {
	return js.True == bindings.SetAutomationNodeMinValueForRange(
		this.ref,
		float64(val),
	)
}

// MaxValueForRange returns the value of property "AutomationNode.maxValueForRange".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) MaxValueForRange() (ret float64, ok bool) {
	ok = js.True == bindings.GetAutomationNodeMaxValueForRange(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMaxValueForRange sets the value of property "AutomationNode.maxValueForRange" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetMaxValueForRange(val float64) bool {
	return js.True == bindings.SetAutomationNodeMaxValueForRange(
		this.ref,
		float64(val),
	)
}

// PosInSet returns the value of property "AutomationNode.posInSet".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) PosInSet() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodePosInSet(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPosInSet sets the value of property "AutomationNode.posInSet" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetPosInSet(val int32) bool {
	return js.True == bindings.SetAutomationNodePosInSet(
		this.ref,
		int32(val),
	)
}

// SetSize returns the value of property "AutomationNode.setSize".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SetSize() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSetSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSetSize sets the value of property "AutomationNode.setSize" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSetSize(val int32) bool {
	return js.True == bindings.SetAutomationNodeSetSize(
		this.ref,
		int32(val),
	)
}

// TableRowCount returns the value of property "AutomationNode.tableRowCount".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableRowCount() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableRowCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableRowCount sets the value of property "AutomationNode.tableRowCount" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableRowCount(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableRowCount(
		this.ref,
		int32(val),
	)
}

// AriaRowCount returns the value of property "AutomationNode.ariaRowCount".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AriaRowCount() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAriaRowCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaRowCount sets the value of property "AutomationNode.ariaRowCount" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAriaRowCount(val int32) bool {
	return js.True == bindings.SetAutomationNodeAriaRowCount(
		this.ref,
		int32(val),
	)
}

// TableColumnCount returns the value of property "AutomationNode.tableColumnCount".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableColumnCount() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableColumnCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableColumnCount sets the value of property "AutomationNode.tableColumnCount" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableColumnCount(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableColumnCount(
		this.ref,
		int32(val),
	)
}

// AriaColumnCount returns the value of property "AutomationNode.ariaColumnCount".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AriaColumnCount() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAriaColumnCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaColumnCount sets the value of property "AutomationNode.ariaColumnCount" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAriaColumnCount(val int32) bool {
	return js.True == bindings.SetAutomationNodeAriaColumnCount(
		this.ref,
		int32(val),
	)
}

// TableCellColumnIndex returns the value of property "AutomationNode.tableCellColumnIndex".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableCellColumnIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableCellColumnIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableCellColumnIndex sets the value of property "AutomationNode.tableCellColumnIndex" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableCellColumnIndex(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableCellColumnIndex(
		this.ref,
		int32(val),
	)
}

// TableCellAriaColumnIndex returns the value of property "AutomationNode.tableCellAriaColumnIndex".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableCellAriaColumnIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableCellAriaColumnIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableCellAriaColumnIndex sets the value of property "AutomationNode.tableCellAriaColumnIndex" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableCellAriaColumnIndex(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableCellAriaColumnIndex(
		this.ref,
		int32(val),
	)
}

// TableCellColumnSpan returns the value of property "AutomationNode.tableCellColumnSpan".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableCellColumnSpan() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableCellColumnSpan(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableCellColumnSpan sets the value of property "AutomationNode.tableCellColumnSpan" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableCellColumnSpan(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableCellColumnSpan(
		this.ref,
		int32(val),
	)
}

// TableCellRowIndex returns the value of property "AutomationNode.tableCellRowIndex".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableCellRowIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableCellRowIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableCellRowIndex sets the value of property "AutomationNode.tableCellRowIndex" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableCellRowIndex(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableCellRowIndex(
		this.ref,
		int32(val),
	)
}

// TableCellAriaRowIndex returns the value of property "AutomationNode.tableCellAriaRowIndex".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableCellAriaRowIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableCellAriaRowIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableCellAriaRowIndex sets the value of property "AutomationNode.tableCellAriaRowIndex" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableCellAriaRowIndex(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableCellAriaRowIndex(
		this.ref,
		int32(val),
	)
}

// TableCellRowSpan returns the value of property "AutomationNode.tableCellRowSpan".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableCellRowSpan() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableCellRowSpan(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableCellRowSpan sets the value of property "AutomationNode.tableCellRowSpan" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableCellRowSpan(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableCellRowSpan(
		this.ref,
		int32(val),
	)
}

// TableColumnHeader returns the value of property "AutomationNode.tableColumnHeader".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableColumnHeader() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableColumnHeader(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableColumnHeader sets the value of property "AutomationNode.tableColumnHeader" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableColumnHeader(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeTableColumnHeader(
		this.ref,
		val.Ref(),
	)
}

// TableRowHeader returns the value of property "AutomationNode.tableRowHeader".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableRowHeader() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableRowHeader(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableRowHeader sets the value of property "AutomationNode.tableRowHeader" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableRowHeader(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeTableRowHeader(
		this.ref,
		val.Ref(),
	)
}

// TableColumnIndex returns the value of property "AutomationNode.tableColumnIndex".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableColumnIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableColumnIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableColumnIndex sets the value of property "AutomationNode.tableColumnIndex" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableColumnIndex(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableColumnIndex(
		this.ref,
		int32(val),
	)
}

// TableRowIndex returns the value of property "AutomationNode.tableRowIndex".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) TableRowIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeTableRowIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTableRowIndex sets the value of property "AutomationNode.tableRowIndex" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetTableRowIndex(val int32) bool {
	return js.True == bindings.SetAutomationNodeTableRowIndex(
		this.ref,
		int32(val),
	)
}

// LiveStatus returns the value of property "AutomationNode.liveStatus".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) LiveStatus() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeLiveStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLiveStatus sets the value of property "AutomationNode.liveStatus" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLiveStatus(val js.String) bool {
	return js.True == bindings.SetAutomationNodeLiveStatus(
		this.ref,
		val.Ref(),
	)
}

// LiveRelevant returns the value of property "AutomationNode.liveRelevant".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) LiveRelevant() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeLiveRelevant(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLiveRelevant sets the value of property "AutomationNode.liveRelevant" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLiveRelevant(val js.String) bool {
	return js.True == bindings.SetAutomationNodeLiveRelevant(
		this.ref,
		val.Ref(),
	)
}

// LiveAtomic returns the value of property "AutomationNode.liveAtomic".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) LiveAtomic() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeLiveAtomic(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLiveAtomic sets the value of property "AutomationNode.liveAtomic" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLiveAtomic(val bool) bool {
	return js.True == bindings.SetAutomationNodeLiveAtomic(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Busy returns the value of property "AutomationNode.busy".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Busy() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeBusy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBusy sets the value of property "AutomationNode.busy" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetBusy(val bool) bool {
	return js.True == bindings.SetAutomationNodeBusy(
		this.ref,
		js.Bool(bool(val)),
	)
}

// ContainerLiveStatus returns the value of property "AutomationNode.containerLiveStatus".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ContainerLiveStatus() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeContainerLiveStatus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetContainerLiveStatus sets the value of property "AutomationNode.containerLiveStatus" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetContainerLiveStatus(val js.String) bool {
	return js.True == bindings.SetAutomationNodeContainerLiveStatus(
		this.ref,
		val.Ref(),
	)
}

// ContainerLiveRelevant returns the value of property "AutomationNode.containerLiveRelevant".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ContainerLiveRelevant() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeContainerLiveRelevant(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetContainerLiveRelevant sets the value of property "AutomationNode.containerLiveRelevant" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetContainerLiveRelevant(val js.String) bool {
	return js.True == bindings.SetAutomationNodeContainerLiveRelevant(
		this.ref,
		val.Ref(),
	)
}

// ContainerLiveAtomic returns the value of property "AutomationNode.containerLiveAtomic".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ContainerLiveAtomic() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeContainerLiveAtomic(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetContainerLiveAtomic sets the value of property "AutomationNode.containerLiveAtomic" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetContainerLiveAtomic(val bool) bool {
	return js.True == bindings.SetAutomationNodeContainerLiveAtomic(
		this.ref,
		js.Bool(bool(val)),
	)
}

// ContainerLiveBusy returns the value of property "AutomationNode.containerLiveBusy".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ContainerLiveBusy() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeContainerLiveBusy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetContainerLiveBusy sets the value of property "AutomationNode.containerLiveBusy" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetContainerLiveBusy(val bool) bool {
	return js.True == bindings.SetAutomationNodeContainerLiveBusy(
		this.ref,
		js.Bool(bool(val)),
	)
}

// IsButton returns the value of property "AutomationNode.isButton".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) IsButton() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeIsButton(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIsButton sets the value of property "AutomationNode.isButton" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetIsButton(val bool) bool {
	return js.True == bindings.SetAutomationNodeIsButton(
		this.ref,
		js.Bool(bool(val)),
	)
}

// IsCheckBox returns the value of property "AutomationNode.isCheckBox".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) IsCheckBox() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeIsCheckBox(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIsCheckBox sets the value of property "AutomationNode.isCheckBox" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetIsCheckBox(val bool) bool {
	return js.True == bindings.SetAutomationNodeIsCheckBox(
		this.ref,
		js.Bool(bool(val)),
	)
}

// IsComboBox returns the value of property "AutomationNode.isComboBox".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) IsComboBox() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeIsComboBox(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIsComboBox sets the value of property "AutomationNode.isComboBox" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetIsComboBox(val bool) bool {
	return js.True == bindings.SetAutomationNodeIsComboBox(
		this.ref,
		js.Bool(bool(val)),
	)
}

// IsImage returns the value of property "AutomationNode.isImage".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) IsImage() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeIsImage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIsImage sets the value of property "AutomationNode.isImage" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetIsImage(val bool) bool {
	return js.True == bindings.SetAutomationNodeIsImage(
		this.ref,
		js.Bool(bool(val)),
	)
}

// HasHiddenOffscreenNodes returns the value of property "AutomationNode.hasHiddenOffscreenNodes".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) HasHiddenOffscreenNodes() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeHasHiddenOffscreenNodes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHasHiddenOffscreenNodes sets the value of property "AutomationNode.hasHiddenOffscreenNodes" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetHasHiddenOffscreenNodes(val bool) bool {
	return js.True == bindings.SetAutomationNodeHasHiddenOffscreenNodes(
		this.ref,
		js.Bool(bool(val)),
	)
}

// AutoComplete returns the value of property "AutomationNode.autoComplete".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AutoComplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAutoComplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutoComplete sets the value of property "AutomationNode.autoComplete" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAutoComplete(val js.String) bool {
	return js.True == bindings.SetAutomationNodeAutoComplete(
		this.ref,
		val.Ref(),
	)
}

// ClassName returns the value of property "AutomationNode.className".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ClassName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeClassName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetClassName sets the value of property "AutomationNode.className" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetClassName(val js.String) bool {
	return js.True == bindings.SetAutomationNodeClassName(
		this.ref,
		val.Ref(),
	)
}

// Modal returns the value of property "AutomationNode.modal".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Modal() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeModal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetModal sets the value of property "AutomationNode.modal" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetModal(val bool) bool {
	return js.True == bindings.SetAutomationNodeModal(
		this.ref,
		js.Bool(bool(val)),
	)
}

// HtmlAttributes returns the value of property "AutomationNode.htmlAttributes".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) HtmlAttributes() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetAutomationNodeHtmlAttributes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHtmlAttributes sets the value of property "AutomationNode.htmlAttributes" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetHtmlAttributes(val js.Object) bool {
	return js.True == bindings.SetAutomationNodeHtmlAttributes(
		this.ref,
		val.Ref(),
	)
}

// InputType returns the value of property "AutomationNode.inputType".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) InputType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeInputType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInputType sets the value of property "AutomationNode.inputType" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetInputType(val js.String) bool {
	return js.True == bindings.SetAutomationNodeInputType(
		this.ref,
		val.Ref(),
	)
}

// AccessKey returns the value of property "AutomationNode.accessKey".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AccessKey() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAccessKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAccessKey sets the value of property "AutomationNode.accessKey" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAccessKey(val js.String) bool {
	return js.True == bindings.SetAutomationNodeAccessKey(
		this.ref,
		val.Ref(),
	)
}

// AriaInvalidValue returns the value of property "AutomationNode.ariaInvalidValue".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AriaInvalidValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAriaInvalidValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaInvalidValue sets the value of property "AutomationNode.ariaInvalidValue" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAriaInvalidValue(val js.String) bool {
	return js.True == bindings.SetAutomationNodeAriaInvalidValue(
		this.ref,
		val.Ref(),
	)
}

// Display returns the value of property "AutomationNode.display".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Display() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeDisplay(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisplay sets the value of property "AutomationNode.display" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDisplay(val js.String) bool {
	return js.True == bindings.SetAutomationNodeDisplay(
		this.ref,
		val.Ref(),
	)
}

// ImageDataUrl returns the value of property "AutomationNode.imageDataUrl".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ImageDataUrl() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeImageDataUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetImageDataUrl sets the value of property "AutomationNode.imageDataUrl" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetImageDataUrl(val js.String) bool {
	return js.True == bindings.SetAutomationNodeImageDataUrl(
		this.ref,
		val.Ref(),
	)
}

// Language returns the value of property "AutomationNode.language".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Language() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeLanguage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLanguage sets the value of property "AutomationNode.language" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLanguage(val js.String) bool {
	return js.True == bindings.SetAutomationNodeLanguage(
		this.ref,
		val.Ref(),
	)
}

// DetectedLanguage returns the value of property "AutomationNode.detectedLanguage".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) DetectedLanguage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeDetectedLanguage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDetectedLanguage sets the value of property "AutomationNode.detectedLanguage" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetDetectedLanguage(val js.String) bool {
	return js.True == bindings.SetAutomationNodeDetectedLanguage(
		this.ref,
		val.Ref(),
	)
}

// HasPopup returns the value of property "AutomationNode.hasPopup".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) HasPopup() (ret HasPopup, ok bool) {
	ok = js.True == bindings.GetAutomationNodeHasPopup(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHasPopup sets the value of property "AutomationNode.hasPopup" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetHasPopup(val HasPopup) bool {
	return js.True == bindings.SetAutomationNodeHasPopup(
		this.ref,
		uint32(val),
	)
}

// Restriction returns the value of property "AutomationNode.restriction".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Restriction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeRestriction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRestriction sets the value of property "AutomationNode.restriction" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetRestriction(val js.String) bool {
	return js.True == bindings.SetAutomationNodeRestriction(
		this.ref,
		val.Ref(),
	)
}

// Checked returns the value of property "AutomationNode.checked".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Checked() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeChecked(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChecked sets the value of property "AutomationNode.checked" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetChecked(val js.String) bool {
	return js.True == bindings.SetAutomationNodeChecked(
		this.ref,
		val.Ref(),
	)
}

// InnerHtml returns the value of property "AutomationNode.innerHtml".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) InnerHtml() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeInnerHtml(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInnerHtml sets the value of property "AutomationNode.innerHtml" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetInnerHtml(val js.String) bool {
	return js.True == bindings.SetAutomationNodeInnerHtml(
		this.ref,
		val.Ref(),
	)
}

// Color returns the value of property "AutomationNode.color".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Color() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetColor sets the value of property "AutomationNode.color" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetColor(val int32) bool {
	return js.True == bindings.SetAutomationNodeColor(
		this.ref,
		int32(val),
	)
}

// BackgroundColor returns the value of property "AutomationNode.backgroundColor".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) BackgroundColor() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeBackgroundColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBackgroundColor sets the value of property "AutomationNode.backgroundColor" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetBackgroundColor(val int32) bool {
	return js.True == bindings.SetAutomationNodeBackgroundColor(
		this.ref,
		int32(val),
	)
}

// ColorValue returns the value of property "AutomationNode.colorValue".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) ColorValue() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeColorValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetColorValue sets the value of property "AutomationNode.colorValue" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetColorValue(val int32) bool {
	return js.True == bindings.SetAutomationNodeColorValue(
		this.ref,
		int32(val),
	)
}

// Subscript returns the value of property "AutomationNode.subscript".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Subscript() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSubscript(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSubscript sets the value of property "AutomationNode.subscript" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSubscript(val bool) bool {
	return js.True == bindings.SetAutomationNodeSubscript(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Superscript returns the value of property "AutomationNode.superscript".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Superscript() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSuperscript(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSuperscript sets the value of property "AutomationNode.superscript" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSuperscript(val bool) bool {
	return js.True == bindings.SetAutomationNodeSuperscript(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Bold returns the value of property "AutomationNode.bold".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Bold() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeBold(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBold sets the value of property "AutomationNode.bold" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetBold(val bool) bool {
	return js.True == bindings.SetAutomationNodeBold(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Italic returns the value of property "AutomationNode.italic".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Italic() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeItalic(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetItalic sets the value of property "AutomationNode.italic" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetItalic(val bool) bool {
	return js.True == bindings.SetAutomationNodeItalic(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Underline returns the value of property "AutomationNode.underline".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Underline() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeUnderline(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUnderline sets the value of property "AutomationNode.underline" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetUnderline(val bool) bool {
	return js.True == bindings.SetAutomationNodeUnderline(
		this.ref,
		js.Bool(bool(val)),
	)
}

// LineThrough returns the value of property "AutomationNode.lineThrough".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) LineThrough() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeLineThrough(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLineThrough sets the value of property "AutomationNode.lineThrough" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLineThrough(val bool) bool {
	return js.True == bindings.SetAutomationNodeLineThrough(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Selected returns the value of property "AutomationNode.selected".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Selected() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSelected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelected sets the value of property "AutomationNode.selected" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSelected(val bool) bool {
	return js.True == bindings.SetAutomationNodeSelected(
		this.ref,
		js.Bool(bool(val)),
	)
}

// FontSize returns the value of property "AutomationNode.fontSize".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) FontSize() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeFontSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFontSize sets the value of property "AutomationNode.fontSize" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetFontSize(val int32) bool {
	return js.True == bindings.SetAutomationNodeFontSize(
		this.ref,
		int32(val),
	)
}

// FontFamily returns the value of property "AutomationNode.fontFamily".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) FontFamily() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeFontFamily(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFontFamily sets the value of property "AutomationNode.fontFamily" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetFontFamily(val js.String) bool {
	return js.True == bindings.SetAutomationNodeFontFamily(
		this.ref,
		val.Ref(),
	)
}

// NonAtomicTextFieldRoot returns the value of property "AutomationNode.nonAtomicTextFieldRoot".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) NonAtomicTextFieldRoot() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeNonAtomicTextFieldRoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNonAtomicTextFieldRoot sets the value of property "AutomationNode.nonAtomicTextFieldRoot" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetNonAtomicTextFieldRoot(val bool) bool {
	return js.True == bindings.SetAutomationNodeNonAtomicTextFieldRoot(
		this.ref,
		js.Bool(bool(val)),
	)
}

// AriaCurrentState returns the value of property "AutomationNode.ariaCurrentState".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AriaCurrentState() (ret AriaCurrentState, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAriaCurrentState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAriaCurrentState sets the value of property "AutomationNode.ariaCurrentState" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAriaCurrentState(val AriaCurrentState) bool {
	return js.True == bindings.SetAutomationNodeAriaCurrentState(
		this.ref,
		uint32(val),
	)
}

// InvalidState returns the value of property "AutomationNode.invalidState".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) InvalidState() (ret InvalidState, ok bool) {
	ok = js.True == bindings.GetAutomationNodeInvalidState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInvalidState sets the value of property "AutomationNode.invalidState" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetInvalidState(val InvalidState) bool {
	return js.True == bindings.SetAutomationNodeInvalidState(
		this.ref,
		uint32(val),
	)
}

// AppId returns the value of property "AutomationNode.appId".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) AppId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAutomationNodeAppId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAppId sets the value of property "AutomationNode.appId" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetAppId(val js.String) bool {
	return js.True == bindings.SetAutomationNodeAppId(
		this.ref,
		val.Ref(),
	)
}

// Children returns the value of property "AutomationNode.children".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Children() (ret js.Array[AutomationNode], ok bool) {
	ok = js.True == bindings.GetAutomationNodeChildren(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChildren sets the value of property "AutomationNode.children" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetChildren(val js.Array[AutomationNode]) bool {
	return js.True == bindings.SetAutomationNodeChildren(
		this.ref,
		val.Ref(),
	)
}

// Parent returns the value of property "AutomationNode.parent".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Parent() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeParent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetParent sets the value of property "AutomationNode.parent" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetParent(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeParent(
		this.ref,
		val.Ref(),
	)
}

// FirstChild returns the value of property "AutomationNode.firstChild".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) FirstChild() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeFirstChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFirstChild sets the value of property "AutomationNode.firstChild" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetFirstChild(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeFirstChild(
		this.ref,
		val.Ref(),
	)
}

// LastChild returns the value of property "AutomationNode.lastChild".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) LastChild() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeLastChild(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLastChild sets the value of property "AutomationNode.lastChild" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetLastChild(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeLastChild(
		this.ref,
		val.Ref(),
	)
}

// PreviousSibling returns the value of property "AutomationNode.previousSibling".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) PreviousSibling() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodePreviousSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPreviousSibling sets the value of property "AutomationNode.previousSibling" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetPreviousSibling(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodePreviousSibling(
		this.ref,
		val.Ref(),
	)
}

// NextSibling returns the value of property "AutomationNode.nextSibling".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) NextSibling() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeNextSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNextSibling sets the value of property "AutomationNode.nextSibling" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetNextSibling(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeNextSibling(
		this.ref,
		val.Ref(),
	)
}

// PreviousOnLine returns the value of property "AutomationNode.previousOnLine".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) PreviousOnLine() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodePreviousOnLine(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPreviousOnLine sets the value of property "AutomationNode.previousOnLine" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetPreviousOnLine(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodePreviousOnLine(
		this.ref,
		val.Ref(),
	)
}

// NextOnLine returns the value of property "AutomationNode.nextOnLine".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) NextOnLine() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeNextOnLine(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNextOnLine sets the value of property "AutomationNode.nextOnLine" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetNextOnLine(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeNextOnLine(
		this.ref,
		val.Ref(),
	)
}

// PreviousFocus returns the value of property "AutomationNode.previousFocus".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) PreviousFocus() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodePreviousFocus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPreviousFocus sets the value of property "AutomationNode.previousFocus" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetPreviousFocus(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodePreviousFocus(
		this.ref,
		val.Ref(),
	)
}

// NextFocus returns the value of property "AutomationNode.nextFocus".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) NextFocus() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeNextFocus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNextFocus sets the value of property "AutomationNode.nextFocus" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetNextFocus(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeNextFocus(
		this.ref,
		val.Ref(),
	)
}

// PreviousWindowFocus returns the value of property "AutomationNode.previousWindowFocus".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) PreviousWindowFocus() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodePreviousWindowFocus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPreviousWindowFocus sets the value of property "AutomationNode.previousWindowFocus" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetPreviousWindowFocus(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodePreviousWindowFocus(
		this.ref,
		val.Ref(),
	)
}

// NextWindowFocus returns the value of property "AutomationNode.nextWindowFocus".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) NextWindowFocus() (ret AutomationNode, ok bool) {
	ok = js.True == bindings.GetAutomationNodeNextWindowFocus(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNextWindowFocus sets the value of property "AutomationNode.nextWindowFocus" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetNextWindowFocus(val AutomationNode) bool {
	return js.True == bindings.SetAutomationNodeNextWindowFocus(
		this.ref,
		val.Ref(),
	)
}

// IndexInParent returns the value of property "AutomationNode.indexInParent".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) IndexInParent() (ret int32, ok bool) {
	ok = js.True == bindings.GetAutomationNodeIndexInParent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIndexInParent sets the value of property "AutomationNode.indexInParent" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetIndexInParent(val int32) bool {
	return js.True == bindings.SetAutomationNodeIndexInParent(
		this.ref,
		int32(val),
	)
}

// SortDirection returns the value of property "AutomationNode.sortDirection".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) SortDirection() (ret SortDirectionType, ok bool) {
	ok = js.True == bindings.GetAutomationNodeSortDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSortDirection sets the value of property "AutomationNode.sortDirection" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetSortDirection(val SortDirectionType) bool {
	return js.True == bindings.SetAutomationNodeSortDirection(
		this.ref,
		uint32(val),
	)
}

// Clickable returns the value of property "AutomationNode.clickable".
//
// It returns ok=false if there is no such property.
func (this AutomationNode) Clickable() (ret bool, ok bool) {
	ok = js.True == bindings.GetAutomationNodeClickable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetClickable sets the value of property "AutomationNode.clickable" to val.
//
// It returns false if the property cannot be set.
func (this AutomationNode) SetClickable(val bool) bool {
	return js.True == bindings.SetAutomationNodeClickable(
		this.ref,
		js.Bool(bool(val)),
	)
}

// HasFuncBoundsForRange returns true if the method "AutomationNode.boundsForRange" exists.
func (this AutomationNode) HasFuncBoundsForRange() bool {
	return js.True == bindings.HasFuncAutomationNodeBoundsForRange(
		this.ref,
	)
}

// FuncBoundsForRange returns the method "AutomationNode.boundsForRange".
func (this AutomationNode) FuncBoundsForRange() (fn js.Func[func(startIndex int32, endIndex int32, callback js.Func[func(bounds *Rect)])]) {
	bindings.FuncAutomationNodeBoundsForRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BoundsForRange calls the method "AutomationNode.boundsForRange".
func (this AutomationNode) BoundsForRange(startIndex int32, endIndex int32, callback js.Func[func(bounds *Rect)]) (ret js.Void) {
	bindings.CallAutomationNodeBoundsForRange(
		this.ref, js.Pointer(&ret),
		int32(startIndex),
		int32(endIndex),
		callback.Ref(),
	)

	return
}

// TryBoundsForRange calls the method "AutomationNode.boundsForRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryBoundsForRange(startIndex int32, endIndex int32, callback js.Func[func(bounds *Rect)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeBoundsForRange(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(startIndex),
		int32(endIndex),
		callback.Ref(),
	)

	return
}

// HasFuncUnclippedBoundsForRange returns true if the method "AutomationNode.unclippedBoundsForRange" exists.
func (this AutomationNode) HasFuncUnclippedBoundsForRange() bool {
	return js.True == bindings.HasFuncAutomationNodeUnclippedBoundsForRange(
		this.ref,
	)
}

// FuncUnclippedBoundsForRange returns the method "AutomationNode.unclippedBoundsForRange".
func (this AutomationNode) FuncUnclippedBoundsForRange() (fn js.Func[func(startIndex int32, endIndex int32, callback js.Func[func(bounds *Rect)])]) {
	bindings.FuncAutomationNodeUnclippedBoundsForRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UnclippedBoundsForRange calls the method "AutomationNode.unclippedBoundsForRange".
func (this AutomationNode) UnclippedBoundsForRange(startIndex int32, endIndex int32, callback js.Func[func(bounds *Rect)]) (ret js.Void) {
	bindings.CallAutomationNodeUnclippedBoundsForRange(
		this.ref, js.Pointer(&ret),
		int32(startIndex),
		int32(endIndex),
		callback.Ref(),
	)

	return
}

// TryUnclippedBoundsForRange calls the method "AutomationNode.unclippedBoundsForRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryUnclippedBoundsForRange(startIndex int32, endIndex int32, callback js.Func[func(bounds *Rect)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeUnclippedBoundsForRange(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(startIndex),
		int32(endIndex),
		callback.Ref(),
	)

	return
}

// HasFuncDoDefault returns true if the method "AutomationNode.doDefault" exists.
func (this AutomationNode) HasFuncDoDefault() bool {
	return js.True == bindings.HasFuncAutomationNodeDoDefault(
		this.ref,
	)
}

// FuncDoDefault returns the method "AutomationNode.doDefault".
func (this AutomationNode) FuncDoDefault() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeDoDefault(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DoDefault calls the method "AutomationNode.doDefault".
func (this AutomationNode) DoDefault() (ret js.Void) {
	bindings.CallAutomationNodeDoDefault(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDoDefault calls the method "AutomationNode.doDefault"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryDoDefault() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeDoDefault(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncFocus returns true if the method "AutomationNode.focus" exists.
func (this AutomationNode) HasFuncFocus() bool {
	return js.True == bindings.HasFuncAutomationNodeFocus(
		this.ref,
	)
}

// FuncFocus returns the method "AutomationNode.focus".
func (this AutomationNode) FuncFocus() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeFocus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Focus calls the method "AutomationNode.focus".
func (this AutomationNode) Focus() (ret js.Void) {
	bindings.CallAutomationNodeFocus(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFocus calls the method "AutomationNode.focus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryFocus() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeFocus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetImageData returns true if the method "AutomationNode.getImageData" exists.
func (this AutomationNode) HasFuncGetImageData() bool {
	return js.True == bindings.HasFuncAutomationNodeGetImageData(
		this.ref,
	)
}

// FuncGetImageData returns the method "AutomationNode.getImageData".
func (this AutomationNode) FuncGetImageData() (fn js.Func[func(maxWidth int32, maxHeight int32)]) {
	bindings.FuncAutomationNodeGetImageData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetImageData calls the method "AutomationNode.getImageData".
func (this AutomationNode) GetImageData(maxWidth int32, maxHeight int32) (ret js.Void) {
	bindings.CallAutomationNodeGetImageData(
		this.ref, js.Pointer(&ret),
		int32(maxWidth),
		int32(maxHeight),
	)

	return
}

// TryGetImageData calls the method "AutomationNode.getImageData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryGetImageData(maxWidth int32, maxHeight int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeGetImageData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(maxWidth),
		int32(maxHeight),
	)

	return
}

// HasFuncHitTest returns true if the method "AutomationNode.hitTest" exists.
func (this AutomationNode) HasFuncHitTest() bool {
	return js.True == bindings.HasFuncAutomationNodeHitTest(
		this.ref,
	)
}

// FuncHitTest returns the method "AutomationNode.hitTest".
func (this AutomationNode) FuncHitTest() (fn js.Func[func(x int32, y int32, eventToFire EventType)]) {
	bindings.FuncAutomationNodeHitTest(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HitTest calls the method "AutomationNode.hitTest".
func (this AutomationNode) HitTest(x int32, y int32, eventToFire EventType) (ret js.Void) {
	bindings.CallAutomationNodeHitTest(
		this.ref, js.Pointer(&ret),
		int32(x),
		int32(y),
		uint32(eventToFire),
	)

	return
}

// TryHitTest calls the method "AutomationNode.hitTest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryHitTest(x int32, y int32, eventToFire EventType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeHitTest(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		uint32(eventToFire),
	)

	return
}

// HasFuncHitTestWithReply returns true if the method "AutomationNode.hitTestWithReply" exists.
func (this AutomationNode) HasFuncHitTestWithReply() bool {
	return js.True == bindings.HasFuncAutomationNodeHitTestWithReply(
		this.ref,
	)
}

// FuncHitTestWithReply returns the method "AutomationNode.hitTestWithReply".
func (this AutomationNode) FuncHitTestWithReply() (fn js.Func[func(x int32, y int32, callback js.Func[func(node AutomationNode)])]) {
	bindings.FuncAutomationNodeHitTestWithReply(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HitTestWithReply calls the method "AutomationNode.hitTestWithReply".
func (this AutomationNode) HitTestWithReply(x int32, y int32, callback js.Func[func(node AutomationNode)]) (ret js.Void) {
	bindings.CallAutomationNodeHitTestWithReply(
		this.ref, js.Pointer(&ret),
		int32(x),
		int32(y),
		callback.Ref(),
	)

	return
}

// TryHitTestWithReply calls the method "AutomationNode.hitTestWithReply"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryHitTestWithReply(x int32, y int32, callback js.Func[func(node AutomationNode)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeHitTestWithReply(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
		callback.Ref(),
	)

	return
}

// HasFuncMakeVisible returns true if the method "AutomationNode.makeVisible" exists.
func (this AutomationNode) HasFuncMakeVisible() bool {
	return js.True == bindings.HasFuncAutomationNodeMakeVisible(
		this.ref,
	)
}

// FuncMakeVisible returns the method "AutomationNode.makeVisible".
func (this AutomationNode) FuncMakeVisible() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeMakeVisible(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MakeVisible calls the method "AutomationNode.makeVisible".
func (this AutomationNode) MakeVisible() (ret js.Void) {
	bindings.CallAutomationNodeMakeVisible(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMakeVisible calls the method "AutomationNode.makeVisible"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryMakeVisible() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeMakeVisible(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPerformCustomAction returns true if the method "AutomationNode.performCustomAction" exists.
func (this AutomationNode) HasFuncPerformCustomAction() bool {
	return js.True == bindings.HasFuncAutomationNodePerformCustomAction(
		this.ref,
	)
}

// FuncPerformCustomAction returns the method "AutomationNode.performCustomAction".
func (this AutomationNode) FuncPerformCustomAction() (fn js.Func[func(customActionId int32)]) {
	bindings.FuncAutomationNodePerformCustomAction(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PerformCustomAction calls the method "AutomationNode.performCustomAction".
func (this AutomationNode) PerformCustomAction(customActionId int32) (ret js.Void) {
	bindings.CallAutomationNodePerformCustomAction(
		this.ref, js.Pointer(&ret),
		int32(customActionId),
	)

	return
}

// TryPerformCustomAction calls the method "AutomationNode.performCustomAction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryPerformCustomAction(customActionId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodePerformCustomAction(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(customActionId),
	)

	return
}

// HasFuncPerformStandardAction returns true if the method "AutomationNode.performStandardAction" exists.
func (this AutomationNode) HasFuncPerformStandardAction() bool {
	return js.True == bindings.HasFuncAutomationNodePerformStandardAction(
		this.ref,
	)
}

// FuncPerformStandardAction returns the method "AutomationNode.performStandardAction".
func (this AutomationNode) FuncPerformStandardAction() (fn js.Func[func(actionType ActionType)]) {
	bindings.FuncAutomationNodePerformStandardAction(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PerformStandardAction calls the method "AutomationNode.performStandardAction".
func (this AutomationNode) PerformStandardAction(actionType ActionType) (ret js.Void) {
	bindings.CallAutomationNodePerformStandardAction(
		this.ref, js.Pointer(&ret),
		uint32(actionType),
	)

	return
}

// TryPerformStandardAction calls the method "AutomationNode.performStandardAction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryPerformStandardAction(actionType ActionType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodePerformStandardAction(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(actionType),
	)

	return
}

// HasFuncReplaceSelectedText returns true if the method "AutomationNode.replaceSelectedText" exists.
func (this AutomationNode) HasFuncReplaceSelectedText() bool {
	return js.True == bindings.HasFuncAutomationNodeReplaceSelectedText(
		this.ref,
	)
}

// FuncReplaceSelectedText returns the method "AutomationNode.replaceSelectedText".
func (this AutomationNode) FuncReplaceSelectedText() (fn js.Func[func(value js.String)]) {
	bindings.FuncAutomationNodeReplaceSelectedText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReplaceSelectedText calls the method "AutomationNode.replaceSelectedText".
func (this AutomationNode) ReplaceSelectedText(value js.String) (ret js.Void) {
	bindings.CallAutomationNodeReplaceSelectedText(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryReplaceSelectedText calls the method "AutomationNode.replaceSelectedText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryReplaceSelectedText(value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeReplaceSelectedText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncSetAccessibilityFocus returns true if the method "AutomationNode.setAccessibilityFocus" exists.
func (this AutomationNode) HasFuncSetAccessibilityFocus() bool {
	return js.True == bindings.HasFuncAutomationNodeSetAccessibilityFocus(
		this.ref,
	)
}

// FuncSetAccessibilityFocus returns the method "AutomationNode.setAccessibilityFocus".
func (this AutomationNode) FuncSetAccessibilityFocus() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeSetAccessibilityFocus(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetAccessibilityFocus calls the method "AutomationNode.setAccessibilityFocus".
func (this AutomationNode) SetAccessibilityFocus() (ret js.Void) {
	bindings.CallAutomationNodeSetAccessibilityFocus(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySetAccessibilityFocus calls the method "AutomationNode.setAccessibilityFocus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TrySetAccessibilityFocus() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeSetAccessibilityFocus(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetSelection returns true if the method "AutomationNode.setSelection" exists.
func (this AutomationNode) HasFuncSetSelection() bool {
	return js.True == bindings.HasFuncAutomationNodeSetSelection(
		this.ref,
	)
}

// FuncSetSelection returns the method "AutomationNode.setSelection".
func (this AutomationNode) FuncSetSelection() (fn js.Func[func(startIndex int32, endIndex int32)]) {
	bindings.FuncAutomationNodeSetSelection(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSelection calls the method "AutomationNode.setSelection".
func (this AutomationNode) SetSelection(startIndex int32, endIndex int32) (ret js.Void) {
	bindings.CallAutomationNodeSetSelection(
		this.ref, js.Pointer(&ret),
		int32(startIndex),
		int32(endIndex),
	)

	return
}

// TrySetSelection calls the method "AutomationNode.setSelection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TrySetSelection(startIndex int32, endIndex int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeSetSelection(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(startIndex),
		int32(endIndex),
	)

	return
}

// HasFuncSetSequentialFocusNavigationStartingPoint returns true if the method "AutomationNode.setSequentialFocusNavigationStartingPoint" exists.
func (this AutomationNode) HasFuncSetSequentialFocusNavigationStartingPoint() bool {
	return js.True == bindings.HasFuncAutomationNodeSetSequentialFocusNavigationStartingPoint(
		this.ref,
	)
}

// FuncSetSequentialFocusNavigationStartingPoint returns the method "AutomationNode.setSequentialFocusNavigationStartingPoint".
func (this AutomationNode) FuncSetSequentialFocusNavigationStartingPoint() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeSetSequentialFocusNavigationStartingPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSequentialFocusNavigationStartingPoint calls the method "AutomationNode.setSequentialFocusNavigationStartingPoint".
func (this AutomationNode) SetSequentialFocusNavigationStartingPoint() (ret js.Void) {
	bindings.CallAutomationNodeSetSequentialFocusNavigationStartingPoint(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySetSequentialFocusNavigationStartingPoint calls the method "AutomationNode.setSequentialFocusNavigationStartingPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TrySetSequentialFocusNavigationStartingPoint() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeSetSequentialFocusNavigationStartingPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetValue returns true if the method "AutomationNode.setValue" exists.
func (this AutomationNode) HasFuncSetValue() bool {
	return js.True == bindings.HasFuncAutomationNodeSetValue(
		this.ref,
	)
}

// FuncSetValue returns the method "AutomationNode.setValue".
func (this AutomationNode) FuncSetValue() (fn js.Func[func(value js.String)]) {
	bindings.FuncAutomationNodeSetValue(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetValue calls the method "AutomationNode.setValue".
func (this AutomationNode) SetValue(value js.String) (ret js.Void) {
	bindings.CallAutomationNodeSetValue(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TrySetValue calls the method "AutomationNode.setValue"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TrySetValue(value js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeSetValue(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncShowContextMenu returns true if the method "AutomationNode.showContextMenu" exists.
func (this AutomationNode) HasFuncShowContextMenu() bool {
	return js.True == bindings.HasFuncAutomationNodeShowContextMenu(
		this.ref,
	)
}

// FuncShowContextMenu returns the method "AutomationNode.showContextMenu".
func (this AutomationNode) FuncShowContextMenu() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeShowContextMenu(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowContextMenu calls the method "AutomationNode.showContextMenu".
func (this AutomationNode) ShowContextMenu() (ret js.Void) {
	bindings.CallAutomationNodeShowContextMenu(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShowContextMenu calls the method "AutomationNode.showContextMenu"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryShowContextMenu() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeShowContextMenu(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncResumeMedia returns true if the method "AutomationNode.resumeMedia" exists.
func (this AutomationNode) HasFuncResumeMedia() bool {
	return js.True == bindings.HasFuncAutomationNodeResumeMedia(
		this.ref,
	)
}

// FuncResumeMedia returns the method "AutomationNode.resumeMedia".
func (this AutomationNode) FuncResumeMedia() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeResumeMedia(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ResumeMedia calls the method "AutomationNode.resumeMedia".
func (this AutomationNode) ResumeMedia() (ret js.Void) {
	bindings.CallAutomationNodeResumeMedia(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryResumeMedia calls the method "AutomationNode.resumeMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryResumeMedia() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeResumeMedia(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStartDuckingMedia returns true if the method "AutomationNode.startDuckingMedia" exists.
func (this AutomationNode) HasFuncStartDuckingMedia() bool {
	return js.True == bindings.HasFuncAutomationNodeStartDuckingMedia(
		this.ref,
	)
}

// FuncStartDuckingMedia returns the method "AutomationNode.startDuckingMedia".
func (this AutomationNode) FuncStartDuckingMedia() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeStartDuckingMedia(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StartDuckingMedia calls the method "AutomationNode.startDuckingMedia".
func (this AutomationNode) StartDuckingMedia() (ret js.Void) {
	bindings.CallAutomationNodeStartDuckingMedia(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStartDuckingMedia calls the method "AutomationNode.startDuckingMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryStartDuckingMedia() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeStartDuckingMedia(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStopDuckingMedia returns true if the method "AutomationNode.stopDuckingMedia" exists.
func (this AutomationNode) HasFuncStopDuckingMedia() bool {
	return js.True == bindings.HasFuncAutomationNodeStopDuckingMedia(
		this.ref,
	)
}

// FuncStopDuckingMedia returns the method "AutomationNode.stopDuckingMedia".
func (this AutomationNode) FuncStopDuckingMedia() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeStopDuckingMedia(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StopDuckingMedia calls the method "AutomationNode.stopDuckingMedia".
func (this AutomationNode) StopDuckingMedia() (ret js.Void) {
	bindings.CallAutomationNodeStopDuckingMedia(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStopDuckingMedia calls the method "AutomationNode.stopDuckingMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryStopDuckingMedia() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeStopDuckingMedia(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSuspendMedia returns true if the method "AutomationNode.suspendMedia" exists.
func (this AutomationNode) HasFuncSuspendMedia() bool {
	return js.True == bindings.HasFuncAutomationNodeSuspendMedia(
		this.ref,
	)
}

// FuncSuspendMedia returns the method "AutomationNode.suspendMedia".
func (this AutomationNode) FuncSuspendMedia() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeSuspendMedia(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SuspendMedia calls the method "AutomationNode.suspendMedia".
func (this AutomationNode) SuspendMedia() (ret js.Void) {
	bindings.CallAutomationNodeSuspendMedia(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySuspendMedia calls the method "AutomationNode.suspendMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TrySuspendMedia() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeSuspendMedia(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLongClick returns true if the method "AutomationNode.longClick" exists.
func (this AutomationNode) HasFuncLongClick() bool {
	return js.True == bindings.HasFuncAutomationNodeLongClick(
		this.ref,
	)
}

// FuncLongClick returns the method "AutomationNode.longClick".
func (this AutomationNode) FuncLongClick() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeLongClick(
		this.ref, js.Pointer(&fn),
	)
	return
}

// LongClick calls the method "AutomationNode.longClick".
func (this AutomationNode) LongClick() (ret js.Void) {
	bindings.CallAutomationNodeLongClick(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryLongClick calls the method "AutomationNode.longClick"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryLongClick() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeLongClick(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollBackward returns true if the method "AutomationNode.scrollBackward" exists.
func (this AutomationNode) HasFuncScrollBackward() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollBackward(
		this.ref,
	)
}

// FuncScrollBackward returns the method "AutomationNode.scrollBackward".
func (this AutomationNode) FuncScrollBackward() (fn js.Func[func(callback js.Func[func(result bool)])]) {
	bindings.FuncAutomationNodeScrollBackward(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollBackward calls the method "AutomationNode.scrollBackward".
func (this AutomationNode) ScrollBackward(callback js.Func[func(result bool)]) (ret js.Void) {
	bindings.CallAutomationNodeScrollBackward(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryScrollBackward calls the method "AutomationNode.scrollBackward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollBackward(callback js.Func[func(result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollBackward(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncScrollBackward1 returns true if the method "AutomationNode.scrollBackward" exists.
func (this AutomationNode) HasFuncScrollBackward1() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollBackward1(
		this.ref,
	)
}

// FuncScrollBackward1 returns the method "AutomationNode.scrollBackward".
func (this AutomationNode) FuncScrollBackward1() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeScrollBackward1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollBackward1 calls the method "AutomationNode.scrollBackward".
func (this AutomationNode) ScrollBackward1() (ret js.Void) {
	bindings.CallAutomationNodeScrollBackward1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollBackward1 calls the method "AutomationNode.scrollBackward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollBackward1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollBackward1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollForward returns true if the method "AutomationNode.scrollForward" exists.
func (this AutomationNode) HasFuncScrollForward() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollForward(
		this.ref,
	)
}

// FuncScrollForward returns the method "AutomationNode.scrollForward".
func (this AutomationNode) FuncScrollForward() (fn js.Func[func(callback js.Func[func(result bool)])]) {
	bindings.FuncAutomationNodeScrollForward(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollForward calls the method "AutomationNode.scrollForward".
func (this AutomationNode) ScrollForward(callback js.Func[func(result bool)]) (ret js.Void) {
	bindings.CallAutomationNodeScrollForward(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryScrollForward calls the method "AutomationNode.scrollForward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollForward(callback js.Func[func(result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollForward(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncScrollForward1 returns true if the method "AutomationNode.scrollForward" exists.
func (this AutomationNode) HasFuncScrollForward1() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollForward1(
		this.ref,
	)
}

// FuncScrollForward1 returns the method "AutomationNode.scrollForward".
func (this AutomationNode) FuncScrollForward1() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeScrollForward1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollForward1 calls the method "AutomationNode.scrollForward".
func (this AutomationNode) ScrollForward1() (ret js.Void) {
	bindings.CallAutomationNodeScrollForward1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollForward1 calls the method "AutomationNode.scrollForward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollForward1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollForward1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollUp returns true if the method "AutomationNode.scrollUp" exists.
func (this AutomationNode) HasFuncScrollUp() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollUp(
		this.ref,
	)
}

// FuncScrollUp returns the method "AutomationNode.scrollUp".
func (this AutomationNode) FuncScrollUp() (fn js.Func[func(callback js.Func[func(result bool)])]) {
	bindings.FuncAutomationNodeScrollUp(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollUp calls the method "AutomationNode.scrollUp".
func (this AutomationNode) ScrollUp(callback js.Func[func(result bool)]) (ret js.Void) {
	bindings.CallAutomationNodeScrollUp(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryScrollUp calls the method "AutomationNode.scrollUp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollUp(callback js.Func[func(result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollUp(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncScrollUp1 returns true if the method "AutomationNode.scrollUp" exists.
func (this AutomationNode) HasFuncScrollUp1() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollUp1(
		this.ref,
	)
}

// FuncScrollUp1 returns the method "AutomationNode.scrollUp".
func (this AutomationNode) FuncScrollUp1() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeScrollUp1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollUp1 calls the method "AutomationNode.scrollUp".
func (this AutomationNode) ScrollUp1() (ret js.Void) {
	bindings.CallAutomationNodeScrollUp1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollUp1 calls the method "AutomationNode.scrollUp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollUp1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollUp1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollDown returns true if the method "AutomationNode.scrollDown" exists.
func (this AutomationNode) HasFuncScrollDown() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollDown(
		this.ref,
	)
}

// FuncScrollDown returns the method "AutomationNode.scrollDown".
func (this AutomationNode) FuncScrollDown() (fn js.Func[func(callback js.Func[func(result bool)])]) {
	bindings.FuncAutomationNodeScrollDown(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollDown calls the method "AutomationNode.scrollDown".
func (this AutomationNode) ScrollDown(callback js.Func[func(result bool)]) (ret js.Void) {
	bindings.CallAutomationNodeScrollDown(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryScrollDown calls the method "AutomationNode.scrollDown"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollDown(callback js.Func[func(result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollDown(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncScrollDown1 returns true if the method "AutomationNode.scrollDown" exists.
func (this AutomationNode) HasFuncScrollDown1() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollDown1(
		this.ref,
	)
}

// FuncScrollDown1 returns the method "AutomationNode.scrollDown".
func (this AutomationNode) FuncScrollDown1() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeScrollDown1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollDown1 calls the method "AutomationNode.scrollDown".
func (this AutomationNode) ScrollDown1() (ret js.Void) {
	bindings.CallAutomationNodeScrollDown1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollDown1 calls the method "AutomationNode.scrollDown"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollDown1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollDown1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollLeft returns true if the method "AutomationNode.scrollLeft" exists.
func (this AutomationNode) HasFuncScrollLeft() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollLeft(
		this.ref,
	)
}

// FuncScrollLeft returns the method "AutomationNode.scrollLeft".
func (this AutomationNode) FuncScrollLeft() (fn js.Func[func(callback js.Func[func(result bool)])]) {
	bindings.FuncAutomationNodeScrollLeft(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollLeft calls the method "AutomationNode.scrollLeft".
func (this AutomationNode) ScrollLeft(callback js.Func[func(result bool)]) (ret js.Void) {
	bindings.CallAutomationNodeScrollLeft(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryScrollLeft calls the method "AutomationNode.scrollLeft"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollLeft(callback js.Func[func(result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollLeft(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncScrollLeft1 returns true if the method "AutomationNode.scrollLeft" exists.
func (this AutomationNode) HasFuncScrollLeft1() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollLeft1(
		this.ref,
	)
}

// FuncScrollLeft1 returns the method "AutomationNode.scrollLeft".
func (this AutomationNode) FuncScrollLeft1() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeScrollLeft1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollLeft1 calls the method "AutomationNode.scrollLeft".
func (this AutomationNode) ScrollLeft1() (ret js.Void) {
	bindings.CallAutomationNodeScrollLeft1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollLeft1 calls the method "AutomationNode.scrollLeft"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollLeft1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollLeft1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollRight returns true if the method "AutomationNode.scrollRight" exists.
func (this AutomationNode) HasFuncScrollRight() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollRight(
		this.ref,
	)
}

// FuncScrollRight returns the method "AutomationNode.scrollRight".
func (this AutomationNode) FuncScrollRight() (fn js.Func[func(callback js.Func[func(result bool)])]) {
	bindings.FuncAutomationNodeScrollRight(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollRight calls the method "AutomationNode.scrollRight".
func (this AutomationNode) ScrollRight(callback js.Func[func(result bool)]) (ret js.Void) {
	bindings.CallAutomationNodeScrollRight(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryScrollRight calls the method "AutomationNode.scrollRight"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollRight(callback js.Func[func(result bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollRight(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncScrollRight1 returns true if the method "AutomationNode.scrollRight" exists.
func (this AutomationNode) HasFuncScrollRight1() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollRight1(
		this.ref,
	)
}

// FuncScrollRight1 returns the method "AutomationNode.scrollRight".
func (this AutomationNode) FuncScrollRight1() (fn js.Func[func()]) {
	bindings.FuncAutomationNodeScrollRight1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollRight1 calls the method "AutomationNode.scrollRight".
func (this AutomationNode) ScrollRight1() (ret js.Void) {
	bindings.CallAutomationNodeScrollRight1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScrollRight1 calls the method "AutomationNode.scrollRight"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollRight1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollRight1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncScrollToPoint returns true if the method "AutomationNode.scrollToPoint" exists.
func (this AutomationNode) HasFuncScrollToPoint() bool {
	return js.True == bindings.HasFuncAutomationNodeScrollToPoint(
		this.ref,
	)
}

// FuncScrollToPoint returns the method "AutomationNode.scrollToPoint".
func (this AutomationNode) FuncScrollToPoint() (fn js.Func[func(x int32, y int32)]) {
	bindings.FuncAutomationNodeScrollToPoint(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ScrollToPoint calls the method "AutomationNode.scrollToPoint".
func (this AutomationNode) ScrollToPoint(x int32, y int32) (ret js.Void) {
	bindings.CallAutomationNodeScrollToPoint(
		this.ref, js.Pointer(&ret),
		int32(x),
		int32(y),
	)

	return
}

// TryScrollToPoint calls the method "AutomationNode.scrollToPoint"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryScrollToPoint(x int32, y int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeScrollToPoint(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
	)

	return
}

// HasFuncSetScrollOffset returns true if the method "AutomationNode.setScrollOffset" exists.
func (this AutomationNode) HasFuncSetScrollOffset() bool {
	return js.True == bindings.HasFuncAutomationNodeSetScrollOffset(
		this.ref,
	)
}

// FuncSetScrollOffset returns the method "AutomationNode.setScrollOffset".
func (this AutomationNode) FuncSetScrollOffset() (fn js.Func[func(x int32, y int32)]) {
	bindings.FuncAutomationNodeSetScrollOffset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetScrollOffset calls the method "AutomationNode.setScrollOffset".
func (this AutomationNode) SetScrollOffset(x int32, y int32) (ret js.Void) {
	bindings.CallAutomationNodeSetScrollOffset(
		this.ref, js.Pointer(&ret),
		int32(x),
		int32(y),
	)

	return
}

// TrySetScrollOffset calls the method "AutomationNode.setScrollOffset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TrySetScrollOffset(x int32, y int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeSetScrollOffset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(x),
		int32(y),
	)

	return
}

// HasFuncAddEventListener returns true if the method "AutomationNode.addEventListener" exists.
func (this AutomationNode) HasFuncAddEventListener() bool {
	return js.True == bindings.HasFuncAutomationNodeAddEventListener(
		this.ref,
	)
}

// FuncAddEventListener returns the method "AutomationNode.addEventListener".
func (this AutomationNode) FuncAddEventListener() (fn js.Func[func(eventType EventType, listener js.Func[func(event AutomationEvent)], capture bool)]) {
	bindings.FuncAutomationNodeAddEventListener(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddEventListener calls the method "AutomationNode.addEventListener".
func (this AutomationNode) AddEventListener(eventType EventType, listener js.Func[func(event AutomationEvent)], capture bool) (ret js.Void) {
	bindings.CallAutomationNodeAddEventListener(
		this.ref, js.Pointer(&ret),
		uint32(eventType),
		listener.Ref(),
		js.Bool(bool(capture)),
	)

	return
}

// TryAddEventListener calls the method "AutomationNode.addEventListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryAddEventListener(eventType EventType, listener js.Func[func(event AutomationEvent)], capture bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeAddEventListener(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(eventType),
		listener.Ref(),
		js.Bool(bool(capture)),
	)

	return
}

// HasFuncRemoveEventListener returns true if the method "AutomationNode.removeEventListener" exists.
func (this AutomationNode) HasFuncRemoveEventListener() bool {
	return js.True == bindings.HasFuncAutomationNodeRemoveEventListener(
		this.ref,
	)
}

// FuncRemoveEventListener returns the method "AutomationNode.removeEventListener".
func (this AutomationNode) FuncRemoveEventListener() (fn js.Func[func(eventType EventType, listener js.Func[func(event AutomationEvent)], capture bool)]) {
	bindings.FuncAutomationNodeRemoveEventListener(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveEventListener calls the method "AutomationNode.removeEventListener".
func (this AutomationNode) RemoveEventListener(eventType EventType, listener js.Func[func(event AutomationEvent)], capture bool) (ret js.Void) {
	bindings.CallAutomationNodeRemoveEventListener(
		this.ref, js.Pointer(&ret),
		uint32(eventType),
		listener.Ref(),
		js.Bool(bool(capture)),
	)

	return
}

// TryRemoveEventListener calls the method "AutomationNode.removeEventListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryRemoveEventListener(eventType EventType, listener js.Func[func(event AutomationEvent)], capture bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeRemoveEventListener(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(eventType),
		listener.Ref(),
		js.Bool(bool(capture)),
	)

	return
}

// HasFuncFind returns true if the method "AutomationNode.find" exists.
func (this AutomationNode) HasFuncFind() bool {
	return js.True == bindings.HasFuncAutomationNodeFind(
		this.ref,
	)
}

// FuncFind returns the method "AutomationNode.find".
func (this AutomationNode) FuncFind() (fn js.Func[func(params FindParams) AutomationNode]) {
	bindings.FuncAutomationNodeFind(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Find calls the method "AutomationNode.find".
func (this AutomationNode) Find(params FindParams) (ret AutomationNode) {
	bindings.CallAutomationNodeFind(
		this.ref, js.Pointer(&ret),
		js.Pointer(&params),
	)

	return
}

// TryFind calls the method "AutomationNode.find"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryFind(params FindParams) (ret AutomationNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeFind(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&params),
	)

	return
}

// HasFuncFindAll returns true if the method "AutomationNode.findAll" exists.
func (this AutomationNode) HasFuncFindAll() bool {
	return js.True == bindings.HasFuncAutomationNodeFindAll(
		this.ref,
	)
}

// FuncFindAll returns the method "AutomationNode.findAll".
func (this AutomationNode) FuncFindAll() (fn js.Func[func(params FindParams) js.Array[AutomationNode]]) {
	bindings.FuncAutomationNodeFindAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// FindAll calls the method "AutomationNode.findAll".
func (this AutomationNode) FindAll(params FindParams) (ret js.Array[AutomationNode]) {
	bindings.CallAutomationNodeFindAll(
		this.ref, js.Pointer(&ret),
		js.Pointer(&params),
	)

	return
}

// TryFindAll calls the method "AutomationNode.findAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryFindAll(params FindParams) (ret js.Array[AutomationNode], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeFindAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&params),
	)

	return
}

// HasFuncMatches returns true if the method "AutomationNode.matches" exists.
func (this AutomationNode) HasFuncMatches() bool {
	return js.True == bindings.HasFuncAutomationNodeMatches(
		this.ref,
	)
}

// FuncMatches returns the method "AutomationNode.matches".
func (this AutomationNode) FuncMatches() (fn js.Func[func(params FindParams) bool]) {
	bindings.FuncAutomationNodeMatches(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Matches calls the method "AutomationNode.matches".
func (this AutomationNode) Matches(params FindParams) (ret bool) {
	bindings.CallAutomationNodeMatches(
		this.ref, js.Pointer(&ret),
		js.Pointer(&params),
	)

	return
}

// TryMatches calls the method "AutomationNode.matches"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryMatches(params FindParams) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeMatches(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&params),
	)

	return
}

// HasFuncGetNextTextMatch returns true if the method "AutomationNode.getNextTextMatch" exists.
func (this AutomationNode) HasFuncGetNextTextMatch() bool {
	return js.True == bindings.HasFuncAutomationNodeGetNextTextMatch(
		this.ref,
	)
}

// FuncGetNextTextMatch returns the method "AutomationNode.getNextTextMatch".
func (this AutomationNode) FuncGetNextTextMatch() (fn js.Func[func(searchStr js.String, backward bool) AutomationNode]) {
	bindings.FuncAutomationNodeGetNextTextMatch(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetNextTextMatch calls the method "AutomationNode.getNextTextMatch".
func (this AutomationNode) GetNextTextMatch(searchStr js.String, backward bool) (ret AutomationNode) {
	bindings.CallAutomationNodeGetNextTextMatch(
		this.ref, js.Pointer(&ret),
		searchStr.Ref(),
		js.Bool(bool(backward)),
	)

	return
}

// TryGetNextTextMatch calls the method "AutomationNode.getNextTextMatch"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryGetNextTextMatch(searchStr js.String, backward bool) (ret AutomationNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeGetNextTextMatch(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		searchStr.Ref(),
		js.Bool(bool(backward)),
	)

	return
}

// HasFuncCreatePosition returns true if the method "AutomationNode.createPosition" exists.
func (this AutomationNode) HasFuncCreatePosition() bool {
	return js.True == bindings.HasFuncAutomationNodeCreatePosition(
		this.ref,
	)
}

// FuncCreatePosition returns the method "AutomationNode.createPosition".
func (this AutomationNode) FuncCreatePosition() (fn js.Func[func(typ PositionType, offset int32, isUpstream bool) AutomationPosition]) {
	bindings.FuncAutomationNodeCreatePosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreatePosition calls the method "AutomationNode.createPosition".
func (this AutomationNode) CreatePosition(typ PositionType, offset int32, isUpstream bool) (ret AutomationPosition) {
	bindings.CallAutomationNodeCreatePosition(
		this.ref, js.Pointer(&ret),
		uint32(typ),
		int32(offset),
		js.Bool(bool(isUpstream)),
	)

	return
}

// TryCreatePosition calls the method "AutomationNode.createPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryCreatePosition(typ PositionType, offset int32, isUpstream bool) (ret AutomationPosition, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeCreatePosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
		int32(offset),
		js.Bool(bool(isUpstream)),
	)

	return
}

// HasFuncCreatePosition1 returns true if the method "AutomationNode.createPosition" exists.
func (this AutomationNode) HasFuncCreatePosition1() bool {
	return js.True == bindings.HasFuncAutomationNodeCreatePosition1(
		this.ref,
	)
}

// FuncCreatePosition1 returns the method "AutomationNode.createPosition".
func (this AutomationNode) FuncCreatePosition1() (fn js.Func[func(typ PositionType, offset int32) AutomationPosition]) {
	bindings.FuncAutomationNodeCreatePosition1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreatePosition1 calls the method "AutomationNode.createPosition".
func (this AutomationNode) CreatePosition1(typ PositionType, offset int32) (ret AutomationPosition) {
	bindings.CallAutomationNodeCreatePosition1(
		this.ref, js.Pointer(&ret),
		uint32(typ),
		int32(offset),
	)

	return
}

// TryCreatePosition1 calls the method "AutomationNode.createPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AutomationNode) TryCreatePosition1(typ PositionType, offset int32) (ret AutomationPosition, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAutomationNodeCreatePosition1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
		int32(offset),
	)

	return
}

type DescriptionFromType uint32

const (
	_ DescriptionFromType = iota

	DescriptionFromType_ARIA_DESCRIPTION
	DescriptionFromType_ATTRIBUTE_EXPLICITLY_EMPTY
	DescriptionFromType_BUTTON_LABEL
	DescriptionFromType_POPOVER_ATTRIBUTE
	DescriptionFromType_RELATED_ELEMENT
	DescriptionFromType_RUBY_ANNOTATION
	DescriptionFromType_SUMMARY
	DescriptionFromType_SVG_DESC_ELEMENT
	DescriptionFromType_TABLE_CAPTION
	DescriptionFromType_TITLE
)

func (DescriptionFromType) FromRef(str js.Ref) DescriptionFromType {
	return DescriptionFromType(bindings.ConstOfDescriptionFromType(str))
}

func (x DescriptionFromType) String() (string, bool) {
	switch x {
	case DescriptionFromType_ARIA_DESCRIPTION:
		return "ariaDescription", true
	case DescriptionFromType_ATTRIBUTE_EXPLICITLY_EMPTY:
		return "attributeExplicitlyEmpty", true
	case DescriptionFromType_BUTTON_LABEL:
		return "buttonLabel", true
	case DescriptionFromType_POPOVER_ATTRIBUTE:
		return "popoverAttribute", true
	case DescriptionFromType_RELATED_ELEMENT:
		return "relatedElement", true
	case DescriptionFromType_RUBY_ANNOTATION:
		return "rubyAnnotation", true
	case DescriptionFromType_SUMMARY:
		return "summary", true
	case DescriptionFromType_SVG_DESC_ELEMENT:
		return "svgDescElement", true
	case DescriptionFromType_TABLE_CAPTION:
		return "tableCaption", true
	case DescriptionFromType_TITLE:
		return "title", true
	default:
		return "", false
	}
}

type FocusCallbackFunc func(this js.Ref, focusedNode AutomationNode) js.Ref

func (fn FocusCallbackFunc) Register() js.Func[func(focusedNode AutomationNode)] {
	return js.RegisterCallback[func(focusedNode AutomationNode)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FocusCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		AutomationNode{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FocusCallback[T any] struct {
	Fn  func(arg T, this js.Ref, focusedNode AutomationNode) js.Ref
	Arg T
}

func (cb *FocusCallback[T]) Register() js.Func[func(focusedNode AutomationNode)] {
	return js.RegisterCallback[func(focusedNode AutomationNode)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FocusCallback[T]) DispatchCallback(
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

		AutomationNode{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IntentInputEventType uint32

const (
	_ IntentInputEventType = iota

	IntentInputEventType_INSERT_TEXT
	IntentInputEventType_INSERT_LINE_BREAK
	IntentInputEventType_INSERT_PARAGRAPH
	IntentInputEventType_INSERT_ORDERED_LIST
	IntentInputEventType_INSERT_UNORDERED_LIST
	IntentInputEventType_INSERT_HORIZONTAL_RULE
	IntentInputEventType_INSERT_FROM_PASTE
	IntentInputEventType_INSERT_FROM_DROP
	IntentInputEventType_INSERT_FROM_YANK
	IntentInputEventType_INSERT_TRANSPOSE
	IntentInputEventType_INSERT_REPLACEMENT_TEXT
	IntentInputEventType_INSERT_COMPOSITION_TEXT
	IntentInputEventType_DELETE_WORD_BACKWARD
	IntentInputEventType_DELETE_WORD_FORWARD
	IntentInputEventType_DELETE_SOFT_LINE_BACKWARD
	IntentInputEventType_DELETE_SOFT_LINE_FORWARD
	IntentInputEventType_DELETE_HARD_LINE_BACKWARD
	IntentInputEventType_DELETE_HARD_LINE_FORWARD
	IntentInputEventType_DELETE_CONTENT_BACKWARD
	IntentInputEventType_DELETE_CONTENT_FORWARD
	IntentInputEventType_DELETE_BY_CUT
	IntentInputEventType_DELETE_BY_DRAG
	IntentInputEventType_HISTORY_UNDO
	IntentInputEventType_HISTORY_REDO
	IntentInputEventType_FORMAT_BOLD
	IntentInputEventType_FORMAT_ITALIC
	IntentInputEventType_FORMAT_UNDERLINE
	IntentInputEventType_FORMAT_STRIKE_THROUGH
	IntentInputEventType_FORMAT_SUPERSCRIPT
	IntentInputEventType_FORMAT_SUBSCRIPT
	IntentInputEventType_FORMAT_JUSTIFY_CENTER
	IntentInputEventType_FORMAT_JUSTIFY_FULL
	IntentInputEventType_FORMAT_JUSTIFY_RIGHT
	IntentInputEventType_FORMAT_JUSTIFY_LEFT
	IntentInputEventType_FORMAT_INDENT
	IntentInputEventType_FORMAT_OUTDENT
	IntentInputEventType_FORMAT_REMOVE
	IntentInputEventType_FORMAT_SET_BLOCK_TEXT_DIRECTION
)

func (IntentInputEventType) FromRef(str js.Ref) IntentInputEventType {
	return IntentInputEventType(bindings.ConstOfIntentInputEventType(str))
}

func (x IntentInputEventType) String() (string, bool) {
	switch x {
	case IntentInputEventType_INSERT_TEXT:
		return "insertText", true
	case IntentInputEventType_INSERT_LINE_BREAK:
		return "insertLineBreak", true
	case IntentInputEventType_INSERT_PARAGRAPH:
		return "insertParagraph", true
	case IntentInputEventType_INSERT_ORDERED_LIST:
		return "insertOrderedList", true
	case IntentInputEventType_INSERT_UNORDERED_LIST:
		return "insertUnorderedList", true
	case IntentInputEventType_INSERT_HORIZONTAL_RULE:
		return "insertHorizontalRule", true
	case IntentInputEventType_INSERT_FROM_PASTE:
		return "insertFromPaste", true
	case IntentInputEventType_INSERT_FROM_DROP:
		return "insertFromDrop", true
	case IntentInputEventType_INSERT_FROM_YANK:
		return "insertFromYank", true
	case IntentInputEventType_INSERT_TRANSPOSE:
		return "insertTranspose", true
	case IntentInputEventType_INSERT_REPLACEMENT_TEXT:
		return "insertReplacementText", true
	case IntentInputEventType_INSERT_COMPOSITION_TEXT:
		return "insertCompositionText", true
	case IntentInputEventType_DELETE_WORD_BACKWARD:
		return "deleteWordBackward", true
	case IntentInputEventType_DELETE_WORD_FORWARD:
		return "deleteWordForward", true
	case IntentInputEventType_DELETE_SOFT_LINE_BACKWARD:
		return "deleteSoftLineBackward", true
	case IntentInputEventType_DELETE_SOFT_LINE_FORWARD:
		return "deleteSoftLineForward", true
	case IntentInputEventType_DELETE_HARD_LINE_BACKWARD:
		return "deleteHardLineBackward", true
	case IntentInputEventType_DELETE_HARD_LINE_FORWARD:
		return "deleteHardLineForward", true
	case IntentInputEventType_DELETE_CONTENT_BACKWARD:
		return "deleteContentBackward", true
	case IntentInputEventType_DELETE_CONTENT_FORWARD:
		return "deleteContentForward", true
	case IntentInputEventType_DELETE_BY_CUT:
		return "deleteByCut", true
	case IntentInputEventType_DELETE_BY_DRAG:
		return "deleteByDrag", true
	case IntentInputEventType_HISTORY_UNDO:
		return "historyUndo", true
	case IntentInputEventType_HISTORY_REDO:
		return "historyRedo", true
	case IntentInputEventType_FORMAT_BOLD:
		return "formatBold", true
	case IntentInputEventType_FORMAT_ITALIC:
		return "formatItalic", true
	case IntentInputEventType_FORMAT_UNDERLINE:
		return "formatUnderline", true
	case IntentInputEventType_FORMAT_STRIKE_THROUGH:
		return "formatStrikeThrough", true
	case IntentInputEventType_FORMAT_SUPERSCRIPT:
		return "formatSuperscript", true
	case IntentInputEventType_FORMAT_SUBSCRIPT:
		return "formatSubscript", true
	case IntentInputEventType_FORMAT_JUSTIFY_CENTER:
		return "formatJustifyCenter", true
	case IntentInputEventType_FORMAT_JUSTIFY_FULL:
		return "formatJustifyFull", true
	case IntentInputEventType_FORMAT_JUSTIFY_RIGHT:
		return "formatJustifyRight", true
	case IntentInputEventType_FORMAT_JUSTIFY_LEFT:
		return "formatJustifyLeft", true
	case IntentInputEventType_FORMAT_INDENT:
		return "formatIndent", true
	case IntentInputEventType_FORMAT_OUTDENT:
		return "formatOutdent", true
	case IntentInputEventType_FORMAT_REMOVE:
		return "formatRemove", true
	case IntentInputEventType_FORMAT_SET_BLOCK_TEXT_DIRECTION:
		return "formatSetBlockTextDirection", true
	default:
		return "", false
	}
}

type MarkerType uint32

const (
	_ MarkerType = iota

	MarkerType_SPELLING
	MarkerType_GRAMMAR
	MarkerType_TEXT_MATCH
	MarkerType_ACTIVE_SUGGESTION
	MarkerType_SUGGESTION
	MarkerType_HIGHLIGHT
)

func (MarkerType) FromRef(str js.Ref) MarkerType {
	return MarkerType(bindings.ConstOfMarkerType(str))
}

func (x MarkerType) String() (string, bool) {
	switch x {
	case MarkerType_SPELLING:
		return "spelling", true
	case MarkerType_GRAMMAR:
		return "grammar", true
	case MarkerType_TEXT_MATCH:
		return "textMatch", true
	case MarkerType_ACTIVE_SUGGESTION:
		return "activeSuggestion", true
	case MarkerType_SUGGESTION:
		return "suggestion", true
	case MarkerType_HIGHLIGHT:
		return "highlight", true
	default:
		return "", false
	}
}

type Restriction uint32

const (
	_ Restriction = iota

	Restriction_DISABLED
	Restriction_READ_ONLY
)

func (Restriction) FromRef(str js.Ref) Restriction {
	return Restriction(bindings.ConstOfRestriction(str))
}

func (x Restriction) String() (string, bool) {
	switch x {
	case Restriction_DISABLED:
		return "disabled", true
	case Restriction_READ_ONLY:
		return "readOnly", true
	default:
		return "", false
	}
}

type RootCallbackFunc func(this js.Ref, rootNode AutomationNode) js.Ref

func (fn RootCallbackFunc) Register() js.Func[func(rootNode AutomationNode)] {
	return js.RegisterCallback[func(rootNode AutomationNode)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RootCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		AutomationNode{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RootCallback[T any] struct {
	Fn  func(arg T, this js.Ref, rootNode AutomationNode) js.Ref
	Arg T
}

func (cb *RootCallback[T]) Register() js.Func[func(rootNode AutomationNode)] {
	return js.RegisterCallback[func(rootNode AutomationNode)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RootCallback[T]) DispatchCallback(
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

		AutomationNode{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SetDocumentSelectionParams struct {
	// AnchorObject is "SetDocumentSelectionParams.anchorObject"
	//
	// Optional
	AnchorObject js.Object
	// AnchorOffset is "SetDocumentSelectionParams.anchorOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_AnchorOffset MUST be set to true to make this field effective.
	AnchorOffset int32
	// FocusObject is "SetDocumentSelectionParams.focusObject"
	//
	// Optional
	FocusObject js.Object
	// FocusOffset is "SetDocumentSelectionParams.focusOffset"
	//
	// Optional
	//
	// NOTE: FFI_USE_FocusOffset MUST be set to true to make this field effective.
	FocusOffset int32

	FFI_USE_AnchorOffset bool // for AnchorOffset.
	FFI_USE_FocusOffset  bool // for FocusOffset.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetDocumentSelectionParams with all fields set.
func (p SetDocumentSelectionParams) FromRef(ref js.Ref) SetDocumentSelectionParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetDocumentSelectionParams in the application heap.
func (p SetDocumentSelectionParams) New() js.Ref {
	return bindings.SetDocumentSelectionParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetDocumentSelectionParams) UpdateFrom(ref js.Ref) {
	bindings.SetDocumentSelectionParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetDocumentSelectionParams) Update(ref js.Ref) {
	bindings.SetDocumentSelectionParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetDocumentSelectionParams) FreeMembers(recursive bool) {
	js.Free(
		p.AnchorObject.Ref(),
		p.FocusObject.Ref(),
	)
	p.AnchorObject = p.AnchorObject.FromRef(js.Undefined)
	p.FocusObject = p.FocusObject.FromRef(js.Undefined)
}

type StateType uint32

const (
	_ StateType = iota

	StateType_AUTOFILL_AVAILABLE
	StateType_COLLAPSED
	StateType_DEFAULT
	StateType_EDITABLE
	StateType_EXPANDED
	StateType_FOCUSABLE
	StateType_FOCUSED
	StateType_HORIZONTAL
	StateType_HOVERED
	StateType_IGNORED
	StateType_INVISIBLE
	StateType_LINKED
	StateType_MULTILINE
	StateType_MULTISELECTABLE
	StateType_OFFSCREEN
	StateType_PROTECTED
	StateType_REQUIRED
	StateType_RICHLY_EDITABLE
	StateType_VERTICAL
	StateType_VISITED
)

func (StateType) FromRef(str js.Ref) StateType {
	return StateType(bindings.ConstOfStateType(str))
}

func (x StateType) String() (string, bool) {
	switch x {
	case StateType_AUTOFILL_AVAILABLE:
		return "autofillAvailable", true
	case StateType_COLLAPSED:
		return "collapsed", true
	case StateType_DEFAULT:
		return "default", true
	case StateType_EDITABLE:
		return "editable", true
	case StateType_EXPANDED:
		return "expanded", true
	case StateType_FOCUSABLE:
		return "focusable", true
	case StateType_FOCUSED:
		return "focused", true
	case StateType_HORIZONTAL:
		return "horizontal", true
	case StateType_HOVERED:
		return "hovered", true
	case StateType_IGNORED:
		return "ignored", true
	case StateType_INVISIBLE:
		return "invisible", true
	case StateType_LINKED:
		return "linked", true
	case StateType_MULTILINE:
		return "multiline", true
	case StateType_MULTISELECTABLE:
		return "multiselectable", true
	case StateType_OFFSCREEN:
		return "offscreen", true
	case StateType_PROTECTED:
		return "protected", true
	case StateType_REQUIRED:
		return "required", true
	case StateType_RICHLY_EDITABLE:
		return "richlyEditable", true
	case StateType_VERTICAL:
		return "vertical", true
	case StateType_VISITED:
		return "visited", true
	default:
		return "", false
	}
}

type TreeChangeType uint32

const (
	_ TreeChangeType = iota

	TreeChangeType_NODE_CREATED
	TreeChangeType_SUBTREE_CREATED
	TreeChangeType_NODE_CHANGED
	TreeChangeType_TEXT_CHANGED
	TreeChangeType_NODE_REMOVED
	TreeChangeType_SUBTREE_UPDATE_END
)

func (TreeChangeType) FromRef(str js.Ref) TreeChangeType {
	return TreeChangeType(bindings.ConstOfTreeChangeType(str))
}

func (x TreeChangeType) String() (string, bool) {
	switch x {
	case TreeChangeType_NODE_CREATED:
		return "nodeCreated", true
	case TreeChangeType_SUBTREE_CREATED:
		return "subtreeCreated", true
	case TreeChangeType_NODE_CHANGED:
		return "nodeChanged", true
	case TreeChangeType_TEXT_CHANGED:
		return "textChanged", true
	case TreeChangeType_NODE_REMOVED:
		return "nodeRemoved", true
	case TreeChangeType_SUBTREE_UPDATE_END:
		return "subtreeUpdateEnd", true
	default:
		return "", false
	}
}

type TreeChange struct {
	// Target is "TreeChange.target"
	//
	// Optional
	Target AutomationNode
	// Type is "TreeChange.type"
	//
	// Optional
	Type TreeChangeType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TreeChange with all fields set.
func (p TreeChange) FromRef(ref js.Ref) TreeChange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TreeChange in the application heap.
func (p TreeChange) New() js.Ref {
	return bindings.TreeChangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TreeChange) UpdateFrom(ref js.Ref) {
	bindings.TreeChangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TreeChange) Update(ref js.Ref) {
	bindings.TreeChangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TreeChange) FreeMembers(recursive bool) {
	js.Free(
		p.Target.Ref(),
	)
	p.Target = p.Target.FromRef(js.Undefined)
}

type TreeChangeObserverFunc func(this js.Ref, treeChange *TreeChange) js.Ref

func (fn TreeChangeObserverFunc) Register() js.Func[func(treeChange *TreeChange)] {
	return js.RegisterCallback[func(treeChange *TreeChange)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TreeChangeObserverFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TreeChange
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

type TreeChangeObserver[T any] struct {
	Fn  func(arg T, this js.Ref, treeChange *TreeChange) js.Ref
	Arg T
}

func (cb *TreeChangeObserver[T]) Register() js.Func[func(treeChange *TreeChange)] {
	return js.RegisterCallback[func(treeChange *TreeChange)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TreeChangeObserver[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TreeChange
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

type TreeChangeObserverFilter uint32

const (
	_ TreeChangeObserverFilter = iota

	TreeChangeObserverFilter_NO_TREE_CHANGES
	TreeChangeObserverFilter_LIVE_REGION_TREE_CHANGES
	TreeChangeObserverFilter_TEXT_MARKER_CHANGES
	TreeChangeObserverFilter_ALL_TREE_CHANGES
)

func (TreeChangeObserverFilter) FromRef(str js.Ref) TreeChangeObserverFilter {
	return TreeChangeObserverFilter(bindings.ConstOfTreeChangeObserverFilter(str))
}

func (x TreeChangeObserverFilter) String() (string, bool) {
	switch x {
	case TreeChangeObserverFilter_NO_TREE_CHANGES:
		return "noTreeChanges", true
	case TreeChangeObserverFilter_LIVE_REGION_TREE_CHANGES:
		return "liveRegionTreeChanges", true
	case TreeChangeObserverFilter_TEXT_MARKER_CHANGES:
		return "textMarkerChanges", true
	case TreeChangeObserverFilter_ALL_TREE_CHANGES:
		return "allTreeChanges", true
	default:
		return "", false
	}
}

// HasFuncAddTreeChangeObserver returns true if the function "WEBEXT.automation.addTreeChangeObserver" exists.
func HasFuncAddTreeChangeObserver() bool {
	return js.True == bindings.HasFuncAddTreeChangeObserver()
}

// FuncAddTreeChangeObserver returns the function "WEBEXT.automation.addTreeChangeObserver".
func FuncAddTreeChangeObserver() (fn js.Func[func(filter TreeChangeObserverFilter, observer js.Func[func(treeChange *TreeChange)])]) {
	bindings.FuncAddTreeChangeObserver(
		js.Pointer(&fn),
	)
	return
}

// AddTreeChangeObserver calls the function "WEBEXT.automation.addTreeChangeObserver" directly.
func AddTreeChangeObserver(filter TreeChangeObserverFilter, observer js.Func[func(treeChange *TreeChange)]) (ret js.Void) {
	bindings.CallAddTreeChangeObserver(
		js.Pointer(&ret),
		uint32(filter),
		observer.Ref(),
	)

	return
}

// TryAddTreeChangeObserver calls the function "WEBEXT.automation.addTreeChangeObserver"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddTreeChangeObserver(filter TreeChangeObserverFilter, observer js.Func[func(treeChange *TreeChange)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddTreeChangeObserver(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(filter),
		observer.Ref(),
	)

	return
}

// HasFuncGetAccessibilityFocus returns true if the function "WEBEXT.automation.getAccessibilityFocus" exists.
func HasFuncGetAccessibilityFocus() bool {
	return js.True == bindings.HasFuncGetAccessibilityFocus()
}

// FuncGetAccessibilityFocus returns the function "WEBEXT.automation.getAccessibilityFocus".
func FuncGetAccessibilityFocus() (fn js.Func[func(callback js.Func[func(focusedNode AutomationNode)])]) {
	bindings.FuncGetAccessibilityFocus(
		js.Pointer(&fn),
	)
	return
}

// GetAccessibilityFocus calls the function "WEBEXT.automation.getAccessibilityFocus" directly.
func GetAccessibilityFocus(callback js.Func[func(focusedNode AutomationNode)]) (ret js.Void) {
	bindings.CallGetAccessibilityFocus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetAccessibilityFocus calls the function "WEBEXT.automation.getAccessibilityFocus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAccessibilityFocus(callback js.Func[func(focusedNode AutomationNode)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAccessibilityFocus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetDesktop returns true if the function "WEBEXT.automation.getDesktop" exists.
func HasFuncGetDesktop() bool {
	return js.True == bindings.HasFuncGetDesktop()
}

// FuncGetDesktop returns the function "WEBEXT.automation.getDesktop".
func FuncGetDesktop() (fn js.Func[func(callback js.Func[func(rootNode AutomationNode)])]) {
	bindings.FuncGetDesktop(
		js.Pointer(&fn),
	)
	return
}

// GetDesktop calls the function "WEBEXT.automation.getDesktop" directly.
func GetDesktop(callback js.Func[func(rootNode AutomationNode)]) (ret js.Void) {
	bindings.CallGetDesktop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetDesktop calls the function "WEBEXT.automation.getDesktop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDesktop(callback js.Func[func(rootNode AutomationNode)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDesktop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetFocus returns true if the function "WEBEXT.automation.getFocus" exists.
func HasFuncGetFocus() bool {
	return js.True == bindings.HasFuncGetFocus()
}

// FuncGetFocus returns the function "WEBEXT.automation.getFocus".
func FuncGetFocus() (fn js.Func[func(callback js.Func[func(focusedNode AutomationNode)])]) {
	bindings.FuncGetFocus(
		js.Pointer(&fn),
	)
	return
}

// GetFocus calls the function "WEBEXT.automation.getFocus" directly.
func GetFocus(callback js.Func[func(focusedNode AutomationNode)]) (ret js.Void) {
	bindings.CallGetFocus(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetFocus calls the function "WEBEXT.automation.getFocus"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetFocus(callback js.Func[func(focusedNode AutomationNode)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetFocus(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetTree returns true if the function "WEBEXT.automation.getTree" exists.
func HasFuncGetTree() bool {
	return js.True == bindings.HasFuncGetTree()
}

// FuncGetTree returns the function "WEBEXT.automation.getTree".
func FuncGetTree() (fn js.Func[func(tabId int32, callback js.Func[func(rootNode AutomationNode)])]) {
	bindings.FuncGetTree(
		js.Pointer(&fn),
	)
	return
}

// GetTree calls the function "WEBEXT.automation.getTree" directly.
func GetTree(tabId int32, callback js.Func[func(rootNode AutomationNode)]) (ret js.Void) {
	bindings.CallGetTree(
		js.Pointer(&ret),
		int32(tabId),
		callback.Ref(),
	)

	return
}

// TryGetTree calls the function "WEBEXT.automation.getTree"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetTree(tabId int32, callback js.Func[func(rootNode AutomationNode)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetTree(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(tabId),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveTreeChangeObserver returns true if the function "WEBEXT.automation.removeTreeChangeObserver" exists.
func HasFuncRemoveTreeChangeObserver() bool {
	return js.True == bindings.HasFuncRemoveTreeChangeObserver()
}

// FuncRemoveTreeChangeObserver returns the function "WEBEXT.automation.removeTreeChangeObserver".
func FuncRemoveTreeChangeObserver() (fn js.Func[func(observer js.Func[func(treeChange *TreeChange)])]) {
	bindings.FuncRemoveTreeChangeObserver(
		js.Pointer(&fn),
	)
	return
}

// RemoveTreeChangeObserver calls the function "WEBEXT.automation.removeTreeChangeObserver" directly.
func RemoveTreeChangeObserver(observer js.Func[func(treeChange *TreeChange)]) (ret js.Void) {
	bindings.CallRemoveTreeChangeObserver(
		js.Pointer(&ret),
		observer.Ref(),
	)

	return
}

// TryRemoveTreeChangeObserver calls the function "WEBEXT.automation.removeTreeChangeObserver"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveTreeChangeObserver(observer js.Func[func(treeChange *TreeChange)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveTreeChangeObserver(
		js.Pointer(&ret), js.Pointer(&exception),
		observer.Ref(),
	)

	return
}

// HasFuncSetDocumentSelection returns true if the function "WEBEXT.automation.setDocumentSelection" exists.
func HasFuncSetDocumentSelection() bool {
	return js.True == bindings.HasFuncSetDocumentSelection()
}

// FuncSetDocumentSelection returns the function "WEBEXT.automation.setDocumentSelection".
func FuncSetDocumentSelection() (fn js.Func[func(params SetDocumentSelectionParams)]) {
	bindings.FuncSetDocumentSelection(
		js.Pointer(&fn),
	)
	return
}

// SetDocumentSelection calls the function "WEBEXT.automation.setDocumentSelection" directly.
func SetDocumentSelection(params SetDocumentSelectionParams) (ret js.Void) {
	bindings.CallSetDocumentSelection(
		js.Pointer(&ret),
		js.Pointer(&params),
	)

	return
}

// TrySetDocumentSelection calls the function "WEBEXT.automation.setDocumentSelection"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetDocumentSelection(params SetDocumentSelectionParams) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetDocumentSelection(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&params),
	)

	return
}
