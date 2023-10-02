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

//go:wasmimport plat/js/web constof_XRLayerLayout
//go:noescape

func ConstOfXRLayerLayout(str js.Ref) uint32

//go:wasmimport plat/js/web constof_XRLayerQuality
//go:noescape

func ConstOfXRLayerQuality(str js.Ref) uint32

//go:wasmimport plat/js/web get_XRCompositionLayer_Layout
//go:noescape

func GetXRCompositionLayerLayout(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRCompositionLayer_BlendTextureSourceAlpha
//go:noescape

func GetXRCompositionLayerBlendTextureSourceAlpha(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRCompositionLayer_BlendTextureSourceAlpha
//go:noescape

func SetXRCompositionLayerBlendTextureSourceAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCompositionLayer_ForceMonoPresentation
//go:noescape

func GetXRCompositionLayerForceMonoPresentation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRCompositionLayer_ForceMonoPresentation
//go:noescape

func SetXRCompositionLayerForceMonoPresentation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCompositionLayer_Opacity
//go:noescape

func GetXRCompositionLayerOpacity(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XRCompositionLayer_Opacity
//go:noescape

func SetXRCompositionLayerOpacity(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRCompositionLayer_MipLevels
//go:noescape

func GetXRCompositionLayerMipLevels(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRCompositionLayer_Quality
//go:noescape

func GetXRCompositionLayerQuality(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_XRCompositionLayer_Quality
//go:noescape

func SetXRCompositionLayerQuality(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_XRCompositionLayer_NeedsRedraw
//go:noescape

func GetXRCompositionLayerNeedsRedraw(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_XRCompositionLayer_Destroy
//go:noescape

func CallXRCompositionLayerDestroy(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRCompositionLayer_Destroy
//go:noescape

func XRCompositionLayerDestroyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRCubeLayer_Space
//go:noescape

func GetXRCubeLayerSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRCubeLayer_Space
//go:noescape

func SetXRCubeLayerSpace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCubeLayer_Orientation
//go:noescape

func GetXRCubeLayerOrientation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRCubeLayer_Orientation
//go:noescape

func SetXRCubeLayerOrientation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_XRCubeLayerInit
//go:noescape

func XRCubeLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRCubeLayerInit
//go:noescape

func XRCubeLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRCylinderLayer_Space
//go:noescape

func GetXRCylinderLayerSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRCylinderLayer_Space
//go:noescape

func SetXRCylinderLayerSpace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCylinderLayer_Transform
//go:noescape

func GetXRCylinderLayerTransform(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRCylinderLayer_Transform
//go:noescape

func SetXRCylinderLayerTransform(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCylinderLayer_Radius
//go:noescape

func GetXRCylinderLayerRadius(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XRCylinderLayer_Radius
//go:noescape

func SetXRCylinderLayerRadius(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRCylinderLayer_CentralAngle
//go:noescape

func GetXRCylinderLayerCentralAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XRCylinderLayer_CentralAngle
//go:noescape

func SetXRCylinderLayerCentralAngle(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRCylinderLayer_AspectRatio
//go:noescape

func GetXRCylinderLayerAspectRatio(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XRCylinderLayer_AspectRatio
//go:noescape

func SetXRCylinderLayerAspectRatio(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web constof_XRTextureType
//go:noescape

func ConstOfXRTextureType(str js.Ref) uint32

//go:wasmimport plat/js/web store_XRCylinderLayerInit
//go:noescape

func XRCylinderLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRCylinderLayerInit
//go:noescape

func XRCylinderLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRDepthInformation_Width
//go:noescape

func GetXRDepthInformationWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRDepthInformation_Height
//go:noescape

func GetXRDepthInformationHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRDepthInformation_NormDepthBufferFromNormView
//go:noescape

func GetXRDepthInformationNormDepthBufferFromNormView(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRDepthInformation_RawValueToMeters
//go:noescape

func GetXRDepthInformationRawValueToMeters(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_XREquirectLayer_Space
//go:noescape

func GetXREquirectLayerSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XREquirectLayer_Space
//go:noescape

func SetXREquirectLayerSpace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_Transform
//go:noescape

func GetXREquirectLayerTransform(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XREquirectLayer_Transform
//go:noescape

func SetXREquirectLayerTransform(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_Radius
//go:noescape

func GetXREquirectLayerRadius(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XREquirectLayer_Radius
//go:noescape

func SetXREquirectLayerRadius(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_CentralHorizontalAngle
//go:noescape

func GetXREquirectLayerCentralHorizontalAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XREquirectLayer_CentralHorizontalAngle
//go:noescape

func SetXREquirectLayerCentralHorizontalAngle(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_UpperVerticalAngle
//go:noescape

func GetXREquirectLayerUpperVerticalAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XREquirectLayer_UpperVerticalAngle
//go:noescape

func SetXREquirectLayerUpperVerticalAngle(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_LowerVerticalAngle
//go:noescape

func GetXREquirectLayerLowerVerticalAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XREquirectLayer_LowerVerticalAngle
//go:noescape

func SetXREquirectLayerLowerVerticalAngle(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web store_XREquirectLayerInit
//go:noescape

func XREquirectLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XREquirectLayerInit
//go:noescape

func XREquirectLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_XRInputSourceEventInit
//go:noescape

func XRInputSourceEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRInputSourceEventInit
//go:noescape

func XRInputSourceEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_XRInputSourceEvent_XRInputSourceEvent
//go:noescape

func NewXRInputSourceEventByXRInputSourceEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_XRInputSourceEvent_Frame
//go:noescape

func GetXRInputSourceEventFrame(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRInputSourceEvent_InputSource
//go:noescape

func GetXRInputSourceEventInputSource(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_XRInputSourcesChangeEventInit
//go:noescape

func XRInputSourcesChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRInputSourcesChangeEventInit
//go:noescape

func XRInputSourcesChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_XRInputSourcesChangeEvent_XRInputSourcesChangeEvent
//go:noescape

func NewXRInputSourcesChangeEventByXRInputSourcesChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_XRInputSourcesChangeEvent_Session
//go:noescape

func GetXRInputSourcesChangeEventSession(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRInputSourcesChangeEvent_Added
//go:noescape

func GetXRInputSourcesChangeEventAdded(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRInputSourcesChangeEvent_Removed
//go:noescape

func GetXRInputSourcesChangeEventRemoved(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_XRLayerEventInit
//go:noescape

func XRLayerEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRLayerEventInit
//go:noescape

func XRLayerEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_XRLayerEvent_XRLayerEvent
//go:noescape

func NewXRLayerEventByXRLayerEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_XRLayerEvent_Layer
//go:noescape

func GetXRLayerEventLayer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_XRLayerInit
//go:noescape

func XRLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRLayerInit
//go:noescape

func XRLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRQuadLayer_Space
//go:noescape

func GetXRQuadLayerSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRQuadLayer_Space
//go:noescape

func SetXRQuadLayerSpace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRQuadLayer_Transform
//go:noescape

func GetXRQuadLayerTransform(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRQuadLayer_Transform
//go:noescape

func SetXRQuadLayerTransform(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRQuadLayer_Width
//go:noescape

func GetXRQuadLayerWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XRQuadLayer_Width
//go:noescape

func SetXRQuadLayerWidth(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRQuadLayer_Height
//go:noescape

func GetXRQuadLayerHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XRQuadLayer_Height
//go:noescape

func SetXRQuadLayerHeight(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web store_XRMediaQuadLayerInit
//go:noescape

func XRMediaQuadLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRMediaQuadLayerInit
//go:noescape

func XRMediaQuadLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_XRMediaCylinderLayerInit
//go:noescape

func XRMediaCylinderLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRMediaCylinderLayerInit
//go:noescape

func XRMediaCylinderLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_XRMediaEquirectLayerInit
//go:noescape

func XRMediaEquirectLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRMediaEquirectLayerInit
//go:noescape

func XRMediaEquirectLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_XRMediaBinding_XRMediaBinding
//go:noescape

func NewXRMediaBindingByXRMediaBinding(
	session js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRMediaBinding_CreateQuadLayer
//go:noescape

func CallXRMediaBindingCreateQuadLayer(
	this js.Ref,
	ptr unsafe.Pointer,

	video js.Ref,
	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateQuadLayer
//go:noescape

func XRMediaBindingCreateQuadLayerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRMediaBinding_CreateQuadLayer1
//go:noescape

func CallXRMediaBindingCreateQuadLayer1(
	this js.Ref,
	ptr unsafe.Pointer,

	video js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateQuadLayer1
//go:noescape

func XRMediaBindingCreateQuadLayer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRMediaBinding_CreateCylinderLayer
//go:noescape

func CallXRMediaBindingCreateCylinderLayer(
	this js.Ref,
	ptr unsafe.Pointer,

	video js.Ref,
	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateCylinderLayer
//go:noescape

func XRMediaBindingCreateCylinderLayerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRMediaBinding_CreateCylinderLayer1
//go:noescape

func CallXRMediaBindingCreateCylinderLayer1(
	this js.Ref,
	ptr unsafe.Pointer,

	video js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateCylinderLayer1
//go:noescape

func XRMediaBindingCreateCylinderLayer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRMediaBinding_CreateEquirectLayer
//go:noescape

func CallXRMediaBindingCreateEquirectLayer(
	this js.Ref,
	ptr unsafe.Pointer,

	video js.Ref,
	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateEquirectLayer
//go:noescape

func XRMediaBindingCreateEquirectLayerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRMediaBinding_CreateEquirectLayer1
//go:noescape

func CallXRMediaBindingCreateEquirectLayer1(
	this js.Ref,
	ptr unsafe.Pointer,

	video js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateEquirectLayer1
//go:noescape

func XRMediaBindingCreateEquirectLayer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_XRMediaLayerInit
//go:noescape

func XRMediaLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRMediaLayerInit
//go:noescape

func XRMediaLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRMesh_MeshSpace
//go:noescape

func GetXRMeshMeshSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRMesh_Vertices
//go:noescape

func GetXRMeshVertices(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRMesh_Indices
//go:noescape

func GetXRMeshIndices(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRMesh_LastChangedTime
//go:noescape

func GetXRMeshLastChangedTime(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_XRMesh_SemanticLabel
//go:noescape

func GetXRMeshSemanticLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_XRPermissionDescriptor
//go:noescape

func XRPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRPermissionDescriptor
//go:noescape

func XRPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRPermissionStatus_Granted
//go:noescape

func GetXRPermissionStatusGranted(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRPermissionStatus_Granted
//go:noescape

func SetXRPermissionStatusGranted(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRProjectionLayer_TextureWidth
//go:noescape

func GetXRProjectionLayerTextureWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRProjectionLayer_TextureHeight
//go:noescape

func GetXRProjectionLayerTextureHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRProjectionLayer_TextureArrayLength
//go:noescape

func GetXRProjectionLayerTextureArrayLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRProjectionLayer_IgnoreDepthValues
//go:noescape

func GetXRProjectionLayerIgnoreDepthValues(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRProjectionLayer_FixedFoveation
//go:noescape

func GetXRProjectionLayerFixedFoveation(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_XRProjectionLayer_FixedFoveation
//go:noescape

func SetXRProjectionLayerFixedFoveation(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRProjectionLayer_DeltaPose
//go:noescape

func GetXRProjectionLayerDeltaPose(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_XRProjectionLayer_DeltaPose
//go:noescape

func SetXRProjectionLayerDeltaPose(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_XRProjectionLayerInit
//go:noescape

func XRProjectionLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRProjectionLayerInit
//go:noescape

func XRProjectionLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_XRQuadLayerInit
//go:noescape

func XRQuadLayerInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRQuadLayerInit
//go:noescape

func XRQuadLayerInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_XRReferenceSpaceEventInit
//go:noescape

func XRReferenceSpaceEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRReferenceSpaceEventInit
//go:noescape

func XRReferenceSpaceEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_XRReferenceSpaceEvent_XRReferenceSpaceEvent
//go:noescape

func NewXRReferenceSpaceEventByXRReferenceSpaceEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_XRReferenceSpaceEvent_ReferenceSpace
//go:noescape

func GetXRReferenceSpaceEventReferenceSpace(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRReferenceSpaceEvent_Transform
//go:noescape

func GetXRReferenceSpaceEventTransform(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_XRSessionEventInit
//go:noescape

func XRSessionEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRSessionEventInit
//go:noescape

func XRSessionEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_XRSessionEvent_XRSessionEvent
//go:noescape

func NewXRSessionEventByXRSessionEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_XRSessionEvent_Session
//go:noescape

func GetXRSessionEventSession(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_XRSessionSupportedPermissionDescriptor
//go:noescape

func XRSessionSupportedPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_XRSessionSupportedPermissionDescriptor
//go:noescape

func XRSessionSupportedPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRSubImage_Viewport
//go:noescape

func GetXRSubImageViewport(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRWebGLSubImage_ColorTexture
//go:noescape

func GetXRWebGLSubImageColorTexture(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRWebGLSubImage_DepthStencilTexture
//go:noescape

func GetXRWebGLSubImageDepthStencilTexture(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRWebGLSubImage_MotionVectorTexture
//go:noescape

func GetXRWebGLSubImageMotionVectorTexture(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XRWebGLSubImage_ImageIndex
//go:noescape

func GetXRWebGLSubImageImageIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRWebGLSubImage_ColorTextureWidth
//go:noescape

func GetXRWebGLSubImageColorTextureWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRWebGLSubImage_ColorTextureHeight
//go:noescape

func GetXRWebGLSubImageColorTextureHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRWebGLSubImage_DepthStencilTextureWidth
//go:noescape

func GetXRWebGLSubImageDepthStencilTextureWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRWebGLSubImage_DepthStencilTextureHeight
//go:noescape

func GetXRWebGLSubImageDepthStencilTextureHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRWebGLSubImage_MotionVectorTextureWidth
//go:noescape

func GetXRWebGLSubImageMotionVectorTextureWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRWebGLSubImage_MotionVectorTextureHeight
//go:noescape

func GetXRWebGLSubImageMotionVectorTextureHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XRWebGLDepthInformation_Texture
//go:noescape

func GetXRWebGLDepthInformationTexture(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web new_XRWebGLBinding_XRWebGLBinding
//go:noescape

func NewXRWebGLBindingByXRWebGLBinding(
	session js.Ref,
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRWebGLBinding_NativeProjectionScaleFactor
//go:noescape

func GetXRWebGLBindingNativeProjectionScaleFactor(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_XRWebGLBinding_UsesDepthValues
//go:noescape

func GetXRWebGLBindingUsesDepthValues(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateProjectionLayer
//go:noescape

func CallXRWebGLBindingCreateProjectionLayer(
	this js.Ref,
	ptr unsafe.Pointer,

	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateProjectionLayer
//go:noescape

func XRWebGLBindingCreateProjectionLayerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateProjectionLayer1
//go:noescape

func CallXRWebGLBindingCreateProjectionLayer1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateProjectionLayer1
//go:noescape

func XRWebGLBindingCreateProjectionLayer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateQuadLayer
//go:noescape

func CallXRWebGLBindingCreateQuadLayer(
	this js.Ref,
	ptr unsafe.Pointer,

	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateQuadLayer
//go:noescape

func XRWebGLBindingCreateQuadLayerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateQuadLayer1
//go:noescape

func CallXRWebGLBindingCreateQuadLayer1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateQuadLayer1
//go:noescape

func XRWebGLBindingCreateQuadLayer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateCylinderLayer
//go:noescape

func CallXRWebGLBindingCreateCylinderLayer(
	this js.Ref,
	ptr unsafe.Pointer,

	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateCylinderLayer
//go:noescape

func XRWebGLBindingCreateCylinderLayerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateCylinderLayer1
//go:noescape

func CallXRWebGLBindingCreateCylinderLayer1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateCylinderLayer1
//go:noescape

func XRWebGLBindingCreateCylinderLayer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateEquirectLayer
//go:noescape

func CallXRWebGLBindingCreateEquirectLayer(
	this js.Ref,
	ptr unsafe.Pointer,

	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateEquirectLayer
//go:noescape

func XRWebGLBindingCreateEquirectLayerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateEquirectLayer1
//go:noescape

func CallXRWebGLBindingCreateEquirectLayer1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateEquirectLayer1
//go:noescape

func XRWebGLBindingCreateEquirectLayer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateCubeLayer
//go:noescape

func CallXRWebGLBindingCreateCubeLayer(
	this js.Ref,
	ptr unsafe.Pointer,

	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateCubeLayer
//go:noescape

func XRWebGLBindingCreateCubeLayerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateCubeLayer1
//go:noescape

func CallXRWebGLBindingCreateCubeLayer1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateCubeLayer1
//go:noescape

func XRWebGLBindingCreateCubeLayer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_GetSubImage
//go:noescape

func CallXRWebGLBindingGetSubImage(
	this js.Ref,
	ptr unsafe.Pointer,

	layer js.Ref,
	frame js.Ref,
	eye uint32,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetSubImage
//go:noescape

func XRWebGLBindingGetSubImageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_GetSubImage1
//go:noescape

func CallXRWebGLBindingGetSubImage1(
	this js.Ref,
	ptr unsafe.Pointer,

	layer js.Ref,
	frame js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetSubImage1
//go:noescape

func XRWebGLBindingGetSubImage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_GetViewSubImage
//go:noescape

func CallXRWebGLBindingGetViewSubImage(
	this js.Ref,
	ptr unsafe.Pointer,

	layer js.Ref,
	view js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetViewSubImage
//go:noescape

func XRWebGLBindingGetViewSubImageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_GetCameraImage
//go:noescape

func CallXRWebGLBindingGetCameraImage(
	this js.Ref,
	ptr unsafe.Pointer,

	camera js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetCameraImage
//go:noescape

func XRWebGLBindingGetCameraImageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_GetDepthInformation
//go:noescape

func CallXRWebGLBindingGetDepthInformation(
	this js.Ref,
	ptr unsafe.Pointer,

	view js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetDepthInformation
//go:noescape

func XRWebGLBindingGetDepthInformationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XRWebGLBinding_GetReflectionCubeMap
//go:noescape

func CallXRWebGLBindingGetReflectionCubeMap(
	this js.Ref,
	ptr unsafe.Pointer,

	lightProbe js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetReflectionCubeMap
//go:noescape

func XRWebGLBindingGetReflectionCubeMapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XSLTProcessor_ImportStylesheet
//go:noescape

func CallXSLTProcessorImportStylesheet(
	this js.Ref,
	ptr unsafe.Pointer,

	style js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_ImportStylesheet
//go:noescape

func XSLTProcessorImportStylesheetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XSLTProcessor_TransformToFragment
//go:noescape

func CallXSLTProcessorTransformToFragment(
	this js.Ref,
	ptr unsafe.Pointer,

	source js.Ref,
	output js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_TransformToFragment
//go:noescape

func XSLTProcessorTransformToFragmentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XSLTProcessor_TransformToDocument
//go:noescape

func CallXSLTProcessorTransformToDocument(
	this js.Ref,
	ptr unsafe.Pointer,

	source js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_TransformToDocument
//go:noescape

func XSLTProcessorTransformToDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XSLTProcessor_SetParameter
//go:noescape

func CallXSLTProcessorSetParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	namespaceURI js.Ref,
	localName js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_SetParameter
//go:noescape

func XSLTProcessorSetParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XSLTProcessor_GetParameter
//go:noescape

func CallXSLTProcessorGetParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	namespaceURI js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_GetParameter
//go:noescape

func XSLTProcessorGetParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XSLTProcessor_RemoveParameter
//go:noescape

func CallXSLTProcessorRemoveParameter(
	this js.Ref,
	ptr unsafe.Pointer,

	namespaceURI js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_RemoveParameter
//go:noescape

func XSLTProcessorRemoveParameterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XSLTProcessor_ClearParameters
//go:noescape

func CallXSLTProcessorClearParameters(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_ClearParameters
//go:noescape

func XSLTProcessorClearParametersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XSLTProcessor_Reset
//go:noescape

func CallXSLTProcessorReset(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_Reset
//go:noescape

func XSLTProcessorResetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Console_Assert
//go:noescape

func CallConsoleAssert(
	ptr unsafe.Pointer,

	condition js.Ref,
	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_Assert
//go:noescape

func ConsoleAssertFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Assert1
//go:noescape

func CallConsoleAssert1(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_Assert1
//go:noescape

func ConsoleAssert1Func() js.Ref

//go:wasmimport plat/js/web call_Console_Clear
//go:noescape

func CallConsoleClear(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_Clear
//go:noescape

func ConsoleClearFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Debug
//go:noescape

func CallConsoleDebug(
	ptr unsafe.Pointer,

	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_Debug
//go:noescape

func ConsoleDebugFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Error
//go:noescape

func CallConsoleError(
	ptr unsafe.Pointer,

	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_Error
//go:noescape

func ConsoleErrorFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Info
//go:noescape

func CallConsoleInfo(
	ptr unsafe.Pointer,

	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_Info
//go:noescape

func ConsoleInfoFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Log
//go:noescape

func CallConsoleLog(
	ptr unsafe.Pointer,

	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_Log
//go:noescape

func ConsoleLogFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Table
//go:noescape

func CallConsoleTable(
	ptr unsafe.Pointer,

	tabularData js.Ref,
	properties js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Console_Table
//go:noescape

func ConsoleTableFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Table1
//go:noescape

func CallConsoleTable1(
	ptr unsafe.Pointer,

	tabularData js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Console_Table1
//go:noescape

func ConsoleTable1Func() js.Ref

//go:wasmimport plat/js/web call_Console_Table2
//go:noescape

func CallConsoleTable2(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_Table2
//go:noescape

func ConsoleTable2Func() js.Ref

//go:wasmimport plat/js/web call_Console_Trace
//go:noescape

func CallConsoleTrace(
	ptr unsafe.Pointer,

	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_Trace
//go:noescape

func ConsoleTraceFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Warn
//go:noescape

func CallConsoleWarn(
	ptr unsafe.Pointer,

	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_Warn
//go:noescape

func ConsoleWarnFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Dir
//go:noescape

func CallConsoleDir(
	ptr unsafe.Pointer,

	item js.Ref,
	options js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Console_Dir
//go:noescape

func ConsoleDirFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Dir1
//go:noescape

func CallConsoleDir1(
	ptr unsafe.Pointer,

	item js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Console_Dir1
//go:noescape

func ConsoleDir1Func() js.Ref

//go:wasmimport plat/js/web call_Console_Dir2
//go:noescape

func CallConsoleDir2(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_Dir2
//go:noescape

func ConsoleDir2Func() js.Ref

//go:wasmimport plat/js/web call_Console_Dirxml
//go:noescape

func CallConsoleDirxml(
	ptr unsafe.Pointer,

	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_Dirxml
//go:noescape

func ConsoleDirxmlFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Count
//go:noescape

func CallConsoleCount(
	ptr unsafe.Pointer,

	label js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Console_Count
//go:noescape

func ConsoleCountFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Count1
//go:noescape

func CallConsoleCount1(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_Count1
//go:noescape

func ConsoleCount1Func() js.Ref

//go:wasmimport plat/js/web call_Console_CountReset
//go:noescape

func CallConsoleCountReset(
	ptr unsafe.Pointer,

	label js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Console_CountReset
//go:noescape

func ConsoleCountResetFunc() js.Ref

//go:wasmimport plat/js/web call_Console_CountReset1
//go:noescape

func CallConsoleCountReset1(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_CountReset1
//go:noescape

func ConsoleCountReset1Func() js.Ref

//go:wasmimport plat/js/web call_Console_Group
//go:noescape

func CallConsoleGroup(
	ptr unsafe.Pointer,

	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_Group
//go:noescape

func ConsoleGroupFunc() js.Ref

//go:wasmimport plat/js/web call_Console_GroupCollapsed
//go:noescape

func CallConsoleGroupCollapsed(
	ptr unsafe.Pointer,

	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_GroupCollapsed
//go:noescape

func ConsoleGroupCollapsedFunc() js.Ref

//go:wasmimport plat/js/web call_Console_GroupEnd
//go:noescape

func CallConsoleGroupEnd(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_GroupEnd
//go:noescape

func ConsoleGroupEndFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Time
//go:noescape

func CallConsoleTime(
	ptr unsafe.Pointer,

	label js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Console_Time
//go:noescape

func ConsoleTimeFunc() js.Ref

//go:wasmimport plat/js/web call_Console_Time1
//go:noescape

func CallConsoleTime1(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_Time1
//go:noescape

func ConsoleTime1Func() js.Ref

//go:wasmimport plat/js/web call_Console_TimeLog
//go:noescape

func CallConsoleTimeLog(
	ptr unsafe.Pointer,

	label js.Ref,
	dataPtr unsafe.Pointer,
	dataCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Console_TimeLog
//go:noescape

func ConsoleTimeLogFunc() js.Ref

//go:wasmimport plat/js/web call_Console_TimeLog1
//go:noescape

func CallConsoleTimeLog1(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_TimeLog1
//go:noescape

func ConsoleTimeLog1Func() js.Ref

//go:wasmimport plat/js/web call_Console_TimeEnd
//go:noescape

func CallConsoleTimeEnd(
	ptr unsafe.Pointer,

	label js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Console_TimeEnd
//go:noescape

func ConsoleTimeEndFunc() js.Ref

//go:wasmimport plat/js/web call_Console_TimeEnd1
//go:noescape

func CallConsoleTimeEnd1(
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Console_TimeEnd1
//go:noescape

func ConsoleTimeEnd1Func() js.Ref
