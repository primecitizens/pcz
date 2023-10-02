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

//go:wasmimport plat/js/web store_BluetoothServiceDataFilterInit
//go:noescape

func BluetoothServiceDataFilterInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothServiceDataFilterInit
//go:noescape

func BluetoothServiceDataFilterInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_BluetoothLEScanFilterInit
//go:noescape

func BluetoothLEScanFilterInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothLEScanFilterInit
//go:noescape

func BluetoothLEScanFilterInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RequestDeviceOptions
//go:noescape

func RequestDeviceOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RequestDeviceOptions
//go:noescape

func RequestDeviceOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_Bluetooth_ReferringDevice
//go:noescape

func GetBluetoothReferringDevice(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Bluetooth_GetAvailability
//go:noescape

func CallBluetoothGetAvailability(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_GetAvailability
//go:noescape

func BluetoothGetAvailabilityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Bluetooth_GetDevices
//go:noescape

func CallBluetoothGetDevices(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_GetDevices
//go:noescape

func BluetoothGetDevicesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Bluetooth_RequestDevice
//go:noescape

func CallBluetoothRequestDevice(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_RequestDevice
//go:noescape

func BluetoothRequestDeviceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Bluetooth_RequestDevice1
//go:noescape

func CallBluetoothRequestDevice1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_RequestDevice1
//go:noescape

func BluetoothRequestDevice1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_BluetoothAdvertisingEventInit
//go:noescape

func BluetoothAdvertisingEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothAdvertisingEventInit
//go:noescape

func BluetoothAdvertisingEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_BluetoothAdvertisingEvent_BluetoothAdvertisingEvent
//go:noescape

func NewBluetoothAdvertisingEventByBluetoothAdvertisingEvent(
	typ js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Device
//go:noescape

func GetBluetoothAdvertisingEventDevice(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Uuids
//go:noescape

func GetBluetoothAdvertisingEventUuids(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Name
//go:noescape

func GetBluetoothAdvertisingEventName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Appearance
//go:noescape

func GetBluetoothAdvertisingEventAppearance(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_TxPower
//go:noescape

func GetBluetoothAdvertisingEventTxPower(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Rssi
//go:noescape

func GetBluetoothAdvertisingEventRssi(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_ManufacturerData
//go:noescape

func GetBluetoothAdvertisingEventManufacturerData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_ServiceData
//go:noescape

func GetBluetoothAdvertisingEventServiceData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_BluetoothDataFilterInit
//go:noescape

func BluetoothDataFilterInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothDataFilterInit
//go:noescape

func BluetoothDataFilterInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_BluetoothPermissionDescriptor
//go:noescape

func BluetoothPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothPermissionDescriptor
//go:noescape

func BluetoothPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_BluetoothPermissionResult_Devices
//go:noescape

func GetBluetoothPermissionResultDevices(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_BluetoothPermissionResult_Devices
//go:noescape

func SetBluetoothPermissionResultDevices(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_BluetoothPermissionStorage
//go:noescape

func BluetoothPermissionStorageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothPermissionStorage
//go:noescape

func BluetoothPermissionStorageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothUUID_GetService
//go:noescape

func CallBluetoothUUIDGetService(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_GetService
//go:noescape

func BluetoothUUIDGetServiceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothUUID_GetCharacteristic
//go:noescape

func CallBluetoothUUIDGetCharacteristic(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_GetCharacteristic
//go:noescape

func BluetoothUUIDGetCharacteristicFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothUUID_GetDescriptor
//go:noescape

func CallBluetoothUUIDGetDescriptor(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_GetDescriptor
//go:noescape

func BluetoothUUIDGetDescriptorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BluetoothUUID_CanonicalUUID
//go:noescape

func CallBluetoothUUIDCanonicalUUID(
	this js.Ref,
	ptr unsafe.Pointer,

	alias uint32,
) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_CanonicalUUID
//go:noescape

func BluetoothUUIDCanonicalUUIDFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_BreakType
//go:noescape

func ConstOfBreakType(str js.Ref) uint32

//go:wasmimport plat/js/web get_IntrinsicSizes_MinContentSize
//go:noescape

func GetIntrinsicSizesMinContentSize(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_IntrinsicSizes_MaxContentSize
//go:noescape

func GetIntrinsicSizesMaxContentSize(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_LayoutFragment_InlineSize
//go:noescape

func GetLayoutFragmentInlineSize(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_LayoutFragment_BlockSize
//go:noescape

func GetLayoutFragmentBlockSize(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_LayoutFragment_InlineOffset
//go:noescape

func GetLayoutFragmentInlineOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_LayoutFragment_InlineOffset
//go:noescape

func SetLayoutFragmentInlineOffset(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_LayoutFragment_BlockOffset
//go:noescape

func GetLayoutFragmentBlockOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_LayoutFragment_BlockOffset
//go:noescape

func SetLayoutFragmentBlockOffset(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_LayoutFragment_Data
//go:noescape

func GetLayoutFragmentData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_LayoutFragment_BreakToken
//go:noescape

func GetLayoutFragmentBreakToken(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_LayoutConstraintsOptions
//go:noescape

func LayoutConstraintsOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_LayoutConstraintsOptions
//go:noescape

func LayoutConstraintsOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_LayoutChild_StyleMap
//go:noescape

func GetLayoutChildStyleMap(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_LayoutChild_IntrinsicSizes
//go:noescape

func CallLayoutChildIntrinsicSizes(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_LayoutChild_IntrinsicSizes
//go:noescape

func LayoutChildIntrinsicSizesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_LayoutChild_LayoutNextFragment
//go:noescape

func CallLayoutChildLayoutNextFragment(
	this js.Ref,
	ptr unsafe.Pointer,

	constraints unsafe.Pointer,
	breakToken js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_LayoutChild_LayoutNextFragment
//go:noescape

func LayoutChildLayoutNextFragmentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_ChildBreakToken_BreakType
//go:noescape

func GetChildBreakTokenBreakType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_ChildBreakToken_Child
//go:noescape

func GetChildBreakTokenChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_BreakToken_ChildBreakTokens
//go:noescape

func GetBreakTokenChildBreakTokens(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_BreakToken_Data
//go:noescape

func GetBreakTokenData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_BreakTokenOptions
//go:noescape

func BreakTokenOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BreakTokenOptions
//go:noescape

func BreakTokenOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_BroadcastChannel_BroadcastChannel
//go:noescape

func NewBroadcastChannelByBroadcastChannel(
	name js.Ref) js.Ref

//go:wasmimport plat/js/web get_BroadcastChannel_Name
//go:noescape

func GetBroadcastChannelName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_BroadcastChannel_PostMessage
//go:noescape

func CallBroadcastChannelPostMessage(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_BroadcastChannel_PostMessage
//go:noescape

func BroadcastChannelPostMessageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BroadcastChannel_Close
//go:noescape

func CallBroadcastChannelClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_BroadcastChannel_Close
//go:noescape

func BroadcastChannelCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CropTarget_FromElement
//go:noescape

func CallCropTargetFromElement(
	this js.Ref,
	ptr unsafe.Pointer,

	element js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CropTarget_FromElement
//go:noescape

func CropTargetFromElementFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RestrictionTarget_FromElement
//go:noescape

func CallRestrictionTargetFromElement(
	this js.Ref,
	ptr unsafe.Pointer,

	element js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RestrictionTarget_FromElement
//go:noescape

func RestrictionTargetFromElementFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BrowserCaptureMediaStreamTrack_CropTo
//go:noescape

func CallBrowserCaptureMediaStreamTrackCropTo(
	this js.Ref,
	ptr unsafe.Pointer,

	cropTarget js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_BrowserCaptureMediaStreamTrack_CropTo
//go:noescape

func BrowserCaptureMediaStreamTrackCropToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BrowserCaptureMediaStreamTrack_Clone
//go:noescape

func CallBrowserCaptureMediaStreamTrackClone(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_BrowserCaptureMediaStreamTrack_Clone
//go:noescape

func BrowserCaptureMediaStreamTrackCloneFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_BrowserCaptureMediaStreamTrack_RestrictTo
//go:noescape

func CallBrowserCaptureMediaStreamTrackRestrictTo(
	this js.Ref,
	ptr unsafe.Pointer,

	RestrictionTarget js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_BrowserCaptureMediaStreamTrack_RestrictTo
//go:noescape

func BrowserCaptureMediaStreamTrackRestrictToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_QueuingStrategyInit
//go:noescape

func QueuingStrategyInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_QueuingStrategyInit
//go:noescape

func QueuingStrategyInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ByteLengthQueuingStrategy_ByteLengthQueuingStrategy
//go:noescape

func NewByteLengthQueuingStrategyByByteLengthQueuingStrategy(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_ByteLengthQueuingStrategy_HighWaterMark
//go:noescape

func GetByteLengthQueuingStrategyHighWaterMark(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_ByteLengthQueuingStrategy_Size
//go:noescape

func GetByteLengthQueuingStrategySize(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_SecurityPolicyViolationEventDisposition
//go:noescape

func ConstOfSecurityPolicyViolationEventDisposition(str js.Ref) uint32

//go:wasmimport plat/js/web get_CSPViolationReportBody_DocumentURL
//go:noescape

func GetCSPViolationReportBodyDocumentURL(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSPViolationReportBody_Referrer
//go:noescape

func GetCSPViolationReportBodyReferrer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSPViolationReportBody_BlockedURL
//go:noescape

func GetCSPViolationReportBodyBlockedURL(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSPViolationReportBody_EffectiveDirective
//go:noescape

func GetCSPViolationReportBodyEffectiveDirective(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSPViolationReportBody_OriginalPolicy
//go:noescape

func GetCSPViolationReportBodyOriginalPolicy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSPViolationReportBody_SourceFile
//go:noescape

func GetCSPViolationReportBodySourceFile(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSPViolationReportBody_Sample
//go:noescape

func GetCSPViolationReportBodySample(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSPViolationReportBody_Disposition
//go:noescape

func GetCSPViolationReportBodyDisposition(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_CSPViolationReportBody_StatusCode
//go:noescape

func GetCSPViolationReportBodyStatusCode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_CSPViolationReportBody_LineNumber
//go:noescape

func GetCSPViolationReportBodyLineNumber(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_CSPViolationReportBody_ColumnNumber
//go:noescape

func GetCSPViolationReportBodyColumnNumber(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_CSPViolationReportBody_ToJSON
//go:noescape

func CallCSPViolationReportBodyToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CSPViolationReportBody_ToJSON
//go:noescape

func CSPViolationReportBodyToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_PropertyDefinition
//go:noescape

func PropertyDefinitionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PropertyDefinition
//go:noescape

func PropertyDefinitionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CSSParserOptions
//go:noescape

func CSSParserOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CSSParserOptions
//go:noescape

func CSSParserOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSParserDeclaration_CSSParserDeclaration
//go:noescape

func NewCSSParserDeclarationByCSSParserDeclaration(
	name js.Ref,
	body js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSParserDeclaration_CSSParserDeclaration1
//go:noescape

func NewCSSParserDeclarationByCSSParserDeclaration1(
	name js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSParserDeclaration_Name
//go:noescape

func GetCSSParserDeclarationName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSParserDeclaration_Body
//go:noescape

func GetCSSParserDeclarationBody(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_CSSParserDeclaration_ToString
//go:noescape

func CallCSSParserDeclarationToString(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CSSParserDeclaration_ToString
//go:noescape

func CSSParserDeclarationToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_WorkletOptions
//go:noescape

func WorkletOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WorkletOptions
//go:noescape

func WorkletOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_Worklet_AddModule
//go:noescape

func CallWorkletAddModule(
	this js.Ref,
	ptr unsafe.Pointer,

	moduleURL js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Worklet_AddModule
//go:noescape

func WorkletAddModuleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Worklet_AddModule1
//go:noescape

func CallWorkletAddModule1(
	this js.Ref,
	ptr unsafe.Pointer,

	moduleURL js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Worklet_AddModule1
//go:noescape

func WorkletAddModule1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSS_ElementSources
//go:noescape

func GetCSSElementSources(
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSS_AnimationWorklet
//go:noescape

func GetCSSAnimationWorklet(
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSS_PaintWorklet
//go:noescape

func GetCSSPaintWorklet(
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSS_LayoutWorklet
//go:noescape

func GetCSSLayoutWorklet(
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSS_Highlights
//go:noescape

func GetCSSHighlights(
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_CSS_Escape
//go:noescape

func CallCSSEscape(
	ptr unsafe.Pointer,

	ident js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Escape
//go:noescape

func CSSEscapeFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_RegisterProperty
//go:noescape

func CallCSSRegisterProperty(
	ptr unsafe.Pointer,

	definition unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSS_RegisterProperty
//go:noescape

func CSSRegisterPropertyFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Supports
//go:noescape

func CallCSSSupports(
	ptr unsafe.Pointer,

	property js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Supports
//go:noescape

func CSSSupportsFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Supports1
//go:noescape

func CallCSSSupports1(
	ptr unsafe.Pointer,

	conditionText js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Supports1
//go:noescape

func CSSSupports1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_Number
//go:noescape

func CallCSSNumber(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Number
//go:noescape

func CSSNumberFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Percent
//go:noescape

func CallCSSPercent(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Percent
//go:noescape

func CSSPercentFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cap
//go:noescape

func CallCSSCap(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Cap
//go:noescape

func CSSCapFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ch
//go:noescape

func CallCSSCh(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Ch
//go:noescape

func CSSChFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Em
//go:noescape

func CallCSSEm(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Em
//go:noescape

func CSSEmFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ex
//go:noescape

func CallCSSEx(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Ex
//go:noescape

func CSSExFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ic
//go:noescape

func CallCSSIc(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Ic
//go:noescape

func CSSIcFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lh
//go:noescape

func CallCSSLh(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Lh
//go:noescape

func CSSLhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rcap
//go:noescape

func CallCSSRcap(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Rcap
//go:noescape

func CSSRcapFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rch
//go:noescape

func CallCSSRch(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Rch
//go:noescape

func CSSRchFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rem
//go:noescape

func CallCSSRem(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Rem
//go:noescape

func CSSRemFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rex
//go:noescape

func CallCSSRex(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Rex
//go:noescape

func CSSRexFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ric
//go:noescape

func CallCSSRic(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Ric
//go:noescape

func CSSRicFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rlh
//go:noescape

func CallCSSRlh(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Rlh
//go:noescape

func CSSRlhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vw
//go:noescape

func CallCSSVw(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Vw
//go:noescape

func CSSVwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vh
//go:noescape

func CallCSSVh(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Vh
//go:noescape

func CSSVhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vi
//go:noescape

func CallCSSVi(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Vi
//go:noescape

func CSSViFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vb
//go:noescape

func CallCSSVb(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Vb
//go:noescape

func CSSVbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vmin
//go:noescape

func CallCSSVmin(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Vmin
//go:noescape

func CSSVminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vmax
//go:noescape

func CallCSSVmax(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Vmax
//go:noescape

func CSSVmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svw
//go:noescape

func CallCSSSvw(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Svw
//go:noescape

func CSSSvwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svh
//go:noescape

func CallCSSSvh(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Svh
//go:noescape

func CSSSvhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svi
//go:noescape

func CallCSSSvi(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Svi
//go:noescape

func CSSSviFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svb
//go:noescape

func CallCSSSvb(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Svb
//go:noescape

func CSSSvbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svmin
//go:noescape

func CallCSSSvmin(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Svmin
//go:noescape

func CSSSvminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svmax
//go:noescape

func CallCSSSvmax(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Svmax
//go:noescape

func CSSSvmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvw
//go:noescape

func CallCSSLvw(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Lvw
//go:noescape

func CSSLvwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvh
//go:noescape

func CallCSSLvh(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Lvh
//go:noescape

func CSSLvhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvi
//go:noescape

func CallCSSLvi(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Lvi
//go:noescape

func CSSLviFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvb
//go:noescape

func CallCSSLvb(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Lvb
//go:noescape

func CSSLvbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvmin
//go:noescape

func CallCSSLvmin(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Lvmin
//go:noescape

func CSSLvminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvmax
//go:noescape

func CallCSSLvmax(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Lvmax
//go:noescape

func CSSLvmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvw
//go:noescape

func CallCSSDvw(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Dvw
//go:noescape

func CSSDvwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvh
//go:noescape

func CallCSSDvh(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Dvh
//go:noescape

func CSSDvhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvi
//go:noescape

func CallCSSDvi(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Dvi
//go:noescape

func CSSDviFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvb
//go:noescape

func CallCSSDvb(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Dvb
//go:noescape

func CSSDvbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvmin
//go:noescape

func CallCSSDvmin(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Dvmin
//go:noescape

func CSSDvminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvmax
//go:noescape

func CallCSSDvmax(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Dvmax
//go:noescape

func CSSDvmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqw
//go:noescape

func CallCSSCqw(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Cqw
//go:noescape

func CSSCqwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqh
//go:noescape

func CallCSSCqh(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Cqh
//go:noescape

func CSSCqhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqi
//go:noescape

func CallCSSCqi(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Cqi
//go:noescape

func CSSCqiFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqb
//go:noescape

func CallCSSCqb(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Cqb
//go:noescape

func CSSCqbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqmin
//go:noescape

func CallCSSCqmin(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Cqmin
//go:noescape

func CSSCqminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqmax
//go:noescape

func CallCSSCqmax(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Cqmax
//go:noescape

func CSSCqmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cm
//go:noescape

func CallCSSCm(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Cm
//go:noescape

func CSSCmFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Mm
//go:noescape

func CallCSSMm(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Mm
//go:noescape

func CSSMmFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Q
//go:noescape

func CallCSSQ(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Q
//go:noescape

func CSSQFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_In
//go:noescape

func CallCSSIn(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_In
//go:noescape

func CSSInFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Pt
//go:noescape

func CallCSSPt(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Pt
//go:noescape

func CSSPtFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Pc
//go:noescape

func CallCSSPc(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Pc
//go:noescape

func CSSPcFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Px
//go:noescape

func CallCSSPx(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Px
//go:noescape

func CSSPxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Deg
//go:noescape

func CallCSSDeg(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Deg
//go:noescape

func CSSDegFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Grad
//go:noescape

func CallCSSGrad(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Grad
//go:noescape

func CSSGradFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rad
//go:noescape

func CallCSSRad(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Rad
//go:noescape

func CSSRadFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Turn
//go:noescape

func CallCSSTurn(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Turn
//go:noescape

func CSSTurnFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_S
//go:noescape

func CallCSSS(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_S
//go:noescape

func CSSSFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ms
//go:noescape

func CallCSSMs(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Ms
//go:noescape

func CSSMsFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Hz
//go:noescape

func CallCSSHz(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Hz
//go:noescape

func CSSHzFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_KHz
//go:noescape

func CallCSSKHz(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_KHz
//go:noescape

func CSSKHzFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dpi
//go:noescape

func CallCSSDpi(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Dpi
//go:noescape

func CSSDpiFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dpcm
//go:noescape

func CallCSSDpcm(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Dpcm
//go:noescape

func CSSDpcmFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dppx
//go:noescape

func CallCSSDppx(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Dppx
//go:noescape

func CSSDppxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Fr
//go:noescape

func CallCSSFr(
	ptr unsafe.Pointer,

	value float64,
) js.Ref

//go:wasmimport plat/js/web func_CSS_Fr
//go:noescape

func CSSFrFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseStylesheet
//go:noescape

func CallCSSParseStylesheet(
	ptr unsafe.Pointer,

	css js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseStylesheet
//go:noescape

func CSSParseStylesheetFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseStylesheet1
//go:noescape

func CallCSSParseStylesheet1(
	ptr unsafe.Pointer,

	css js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseStylesheet1
//go:noescape

func CSSParseStylesheet1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseRuleList
//go:noescape

func CallCSSParseRuleList(
	ptr unsafe.Pointer,

	css js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRuleList
//go:noescape

func CSSParseRuleListFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseRuleList1
//go:noescape

func CallCSSParseRuleList1(
	ptr unsafe.Pointer,

	css js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRuleList1
//go:noescape

func CSSParseRuleList1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseRule
//go:noescape

func CallCSSParseRule(
	ptr unsafe.Pointer,

	css js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRule
//go:noescape

func CSSParseRuleFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseRule1
//go:noescape

func CallCSSParseRule1(
	ptr unsafe.Pointer,

	css js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRule1
//go:noescape

func CSSParseRule1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseDeclarationList
//go:noescape

func CallCSSParseDeclarationList(
	ptr unsafe.Pointer,

	css js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclarationList
//go:noescape

func CSSParseDeclarationListFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseDeclarationList1
//go:noescape

func CallCSSParseDeclarationList1(
	ptr unsafe.Pointer,

	css js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclarationList1
//go:noescape

func CSSParseDeclarationList1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseDeclaration
//go:noescape

func CallCSSParseDeclaration(
	ptr unsafe.Pointer,

	css js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclaration
//go:noescape

func CSSParseDeclarationFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseDeclaration1
//go:noescape

func CallCSSParseDeclaration1(
	ptr unsafe.Pointer,

	css js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclaration1
//go:noescape

func CSSParseDeclaration1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseValue
//go:noescape

func CallCSSParseValue(
	ptr unsafe.Pointer,

	css js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseValue
//go:noescape

func CSSParseValueFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseValueList
//go:noescape

func CallCSSParseValueList(
	ptr unsafe.Pointer,

	css js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseValueList
//go:noescape

func CSSParseValueListFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseCommaValueList
//go:noescape

func CallCSSParseCommaValueList(
	ptr unsafe.Pointer,

	css js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSS_ParseCommaValueList
//go:noescape

func CSSParseCommaValueListFunc() js.Ref

//go:wasmimport plat/js/web new_CSSAnimation_CSSAnimation
//go:noescape

func NewCSSAnimationByCSSAnimation(
	effect js.Ref,
	timeline js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSAnimation_CSSAnimation1
//go:noescape

func NewCSSAnimationByCSSAnimation1(
	effect js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSAnimation_CSSAnimation2
//go:noescape

func NewCSSAnimationByCSSAnimation2() js.Ref

//go:wasmimport plat/js/web get_CSSAnimation_AnimationName
//go:noescape

func GetCSSAnimationAnimationName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref
