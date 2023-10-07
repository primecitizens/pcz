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

//go:wasmimport plat/js/web constof_XRLayerLayout
//go:noescape
func ConstOfXRLayerLayout(str js.Ref) uint32

//go:wasmimport plat/js/web constof_XRLayerQuality
//go:noescape
func ConstOfXRLayerQuality(str js.Ref) uint32

//go:wasmimport plat/js/web get_XRCompositionLayer_Layout
//go:noescape
func GetXRCompositionLayerLayout(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRCompositionLayer_BlendTextureSourceAlpha
//go:noescape
func GetXRCompositionLayerBlendTextureSourceAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRCompositionLayer_BlendTextureSourceAlpha
//go:noescape
func SetXRCompositionLayerBlendTextureSourceAlpha(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCompositionLayer_ForceMonoPresentation
//go:noescape
func GetXRCompositionLayerForceMonoPresentation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRCompositionLayer_ForceMonoPresentation
//go:noescape
func SetXRCompositionLayerForceMonoPresentation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCompositionLayer_Opacity
//go:noescape
func GetXRCompositionLayerOpacity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRCompositionLayer_Opacity
//go:noescape
func SetXRCompositionLayerOpacity(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRCompositionLayer_MipLevels
//go:noescape
func GetXRCompositionLayerMipLevels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRCompositionLayer_Quality
//go:noescape
func GetXRCompositionLayerQuality(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRCompositionLayer_Quality
//go:noescape
func SetXRCompositionLayerQuality(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_XRCompositionLayer_NeedsRedraw
//go:noescape
func GetXRCompositionLayerNeedsRedraw(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRCompositionLayer_Destroy
//go:noescape
func HasFuncXRCompositionLayerDestroy(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRCompositionLayer_Destroy
//go:noescape
func FuncXRCompositionLayerDestroy(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRCompositionLayer_Destroy
//go:noescape
func CallXRCompositionLayerDestroy(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRCompositionLayer_Destroy
//go:noescape
func TryXRCompositionLayerDestroy(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRCubeLayer_Space
//go:noescape
func GetXRCubeLayerSpace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRCubeLayer_Space
//go:noescape
func SetXRCubeLayerSpace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCubeLayer_Orientation
//go:noescape
func GetXRCubeLayerOrientation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRCylinderLayer_Space
//go:noescape
func SetXRCylinderLayerSpace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCylinderLayer_Transform
//go:noescape
func GetXRCylinderLayerTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRCylinderLayer_Transform
//go:noescape
func SetXRCylinderLayerTransform(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRCylinderLayer_Radius
//go:noescape
func GetXRCylinderLayerRadius(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRCylinderLayer_Radius
//go:noescape
func SetXRCylinderLayerRadius(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRCylinderLayer_CentralAngle
//go:noescape
func GetXRCylinderLayerCentralAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRCylinderLayer_CentralAngle
//go:noescape
func SetXRCylinderLayerCentralAngle(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRCylinderLayer_AspectRatio
//go:noescape
func GetXRCylinderLayerAspectRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRDepthInformation_Height
//go:noescape
func GetXRDepthInformationHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRDepthInformation_NormDepthBufferFromNormView
//go:noescape
func GetXRDepthInformationNormDepthBufferFromNormView(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRDepthInformation_RawValueToMeters
//go:noescape
func GetXRDepthInformationRawValueToMeters(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XREquirectLayer_Space
//go:noescape
func GetXREquirectLayerSpace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XREquirectLayer_Space
//go:noescape
func SetXREquirectLayerSpace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_Transform
//go:noescape
func GetXREquirectLayerTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XREquirectLayer_Transform
//go:noescape
func SetXREquirectLayerTransform(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_Radius
//go:noescape
func GetXREquirectLayerRadius(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XREquirectLayer_Radius
//go:noescape
func SetXREquirectLayerRadius(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_CentralHorizontalAngle
//go:noescape
func GetXREquirectLayerCentralHorizontalAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XREquirectLayer_CentralHorizontalAngle
//go:noescape
func SetXREquirectLayerCentralHorizontalAngle(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_UpperVerticalAngle
//go:noescape
func GetXREquirectLayerUpperVerticalAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XREquirectLayer_UpperVerticalAngle
//go:noescape
func SetXREquirectLayerUpperVerticalAngle(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XREquirectLayer_LowerVerticalAngle
//go:noescape
func GetXREquirectLayerLowerVerticalAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSourceEvent_InputSource
//go:noescape
func GetXRInputSourceEventInputSource(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSourcesChangeEvent_Added
//go:noescape
func GetXRInputSourcesChangeEventAdded(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRInputSourcesChangeEvent_Removed
//go:noescape
func GetXRInputSourcesChangeEventRemoved(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRQuadLayer_Space
//go:noescape
func SetXRQuadLayerSpace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRQuadLayer_Transform
//go:noescape
func GetXRQuadLayerTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRQuadLayer_Transform
//go:noescape
func SetXRQuadLayerTransform(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRQuadLayer_Width
//go:noescape
func GetXRQuadLayerWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRQuadLayer_Width
//go:noescape
func SetXRQuadLayerWidth(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRQuadLayer_Height
//go:noescape
func GetXRQuadLayerHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_XRMediaBinding_CreateQuadLayer
//go:noescape
func HasFuncXRMediaBindingCreateQuadLayer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateQuadLayer
//go:noescape
func FuncXRMediaBindingCreateQuadLayer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRMediaBinding_CreateQuadLayer
//go:noescape
func CallXRMediaBindingCreateQuadLayer(
	this js.Ref, retPtr unsafe.Pointer,
	video js.Ref,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_XRMediaBinding_CreateQuadLayer
//go:noescape
func TryXRMediaBindingCreateQuadLayer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	video js.Ref,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRMediaBinding_CreateQuadLayer1
//go:noescape
func HasFuncXRMediaBindingCreateQuadLayer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateQuadLayer1
//go:noescape
func FuncXRMediaBindingCreateQuadLayer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRMediaBinding_CreateQuadLayer1
//go:noescape
func CallXRMediaBindingCreateQuadLayer1(
	this js.Ref, retPtr unsafe.Pointer,
	video js.Ref)

//go:wasmimport plat/js/web try_XRMediaBinding_CreateQuadLayer1
//go:noescape
func TryXRMediaBindingCreateQuadLayer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	video js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRMediaBinding_CreateCylinderLayer
//go:noescape
func HasFuncXRMediaBindingCreateCylinderLayer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateCylinderLayer
//go:noescape
func FuncXRMediaBindingCreateCylinderLayer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRMediaBinding_CreateCylinderLayer
//go:noescape
func CallXRMediaBindingCreateCylinderLayer(
	this js.Ref, retPtr unsafe.Pointer,
	video js.Ref,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_XRMediaBinding_CreateCylinderLayer
//go:noescape
func TryXRMediaBindingCreateCylinderLayer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	video js.Ref,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRMediaBinding_CreateCylinderLayer1
//go:noescape
func HasFuncXRMediaBindingCreateCylinderLayer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateCylinderLayer1
//go:noescape
func FuncXRMediaBindingCreateCylinderLayer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRMediaBinding_CreateCylinderLayer1
//go:noescape
func CallXRMediaBindingCreateCylinderLayer1(
	this js.Ref, retPtr unsafe.Pointer,
	video js.Ref)

//go:wasmimport plat/js/web try_XRMediaBinding_CreateCylinderLayer1
//go:noescape
func TryXRMediaBindingCreateCylinderLayer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	video js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRMediaBinding_CreateEquirectLayer
//go:noescape
func HasFuncXRMediaBindingCreateEquirectLayer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateEquirectLayer
//go:noescape
func FuncXRMediaBindingCreateEquirectLayer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRMediaBinding_CreateEquirectLayer
//go:noescape
func CallXRMediaBindingCreateEquirectLayer(
	this js.Ref, retPtr unsafe.Pointer,
	video js.Ref,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_XRMediaBinding_CreateEquirectLayer
//go:noescape
func TryXRMediaBindingCreateEquirectLayer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	video js.Ref,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRMediaBinding_CreateEquirectLayer1
//go:noescape
func HasFuncXRMediaBindingCreateEquirectLayer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRMediaBinding_CreateEquirectLayer1
//go:noescape
func FuncXRMediaBindingCreateEquirectLayer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRMediaBinding_CreateEquirectLayer1
//go:noescape
func CallXRMediaBindingCreateEquirectLayer1(
	this js.Ref, retPtr unsafe.Pointer,
	video js.Ref)

//go:wasmimport plat/js/web try_XRMediaBinding_CreateEquirectLayer1
//go:noescape
func TryXRMediaBindingCreateEquirectLayer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	video js.Ref) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRMesh_Vertices
//go:noescape
func GetXRMeshVertices(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRMesh_Indices
//go:noescape
func GetXRMeshIndices(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRMesh_LastChangedTime
//go:noescape
func GetXRMeshLastChangedTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRMesh_SemanticLabel
//go:noescape
func GetXRMeshSemanticLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRPermissionStatus_Granted
//go:noescape
func SetXRPermissionStatusGranted(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_XRProjectionLayer_TextureWidth
//go:noescape
func GetXRProjectionLayerTextureWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRProjectionLayer_TextureHeight
//go:noescape
func GetXRProjectionLayerTextureHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRProjectionLayer_TextureArrayLength
//go:noescape
func GetXRProjectionLayerTextureArrayLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRProjectionLayer_IgnoreDepthValues
//go:noescape
func GetXRProjectionLayerIgnoreDepthValues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRProjectionLayer_FixedFoveation
//go:noescape
func GetXRProjectionLayerFixedFoveation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_XRProjectionLayer_FixedFoveation
//go:noescape
func SetXRProjectionLayerFixedFoveation(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_XRProjectionLayer_DeltaPose
//go:noescape
func GetXRProjectionLayerDeltaPose(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRReferenceSpaceEvent_Transform
//go:noescape
func GetXRReferenceSpaceEventTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_ColorTexture
//go:noescape
func GetXRWebGLSubImageColorTexture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_DepthStencilTexture
//go:noescape
func GetXRWebGLSubImageDepthStencilTexture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_MotionVectorTexture
//go:noescape
func GetXRWebGLSubImageMotionVectorTexture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_ImageIndex
//go:noescape
func GetXRWebGLSubImageImageIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_ColorTextureWidth
//go:noescape
func GetXRWebGLSubImageColorTextureWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_ColorTextureHeight
//go:noescape
func GetXRWebGLSubImageColorTextureHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_DepthStencilTextureWidth
//go:noescape
func GetXRWebGLSubImageDepthStencilTextureWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_DepthStencilTextureHeight
//go:noescape
func GetXRWebGLSubImageDepthStencilTextureHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_MotionVectorTextureWidth
//go:noescape
func GetXRWebGLSubImageMotionVectorTextureWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLSubImage_MotionVectorTextureHeight
//go:noescape
func GetXRWebGLSubImageMotionVectorTextureHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLDepthInformation_Texture
//go:noescape
func GetXRWebGLDepthInformationTexture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_XRWebGLBinding_XRWebGLBinding
//go:noescape
func NewXRWebGLBindingByXRWebGLBinding(
	session js.Ref,
	context js.Ref) js.Ref

//go:wasmimport plat/js/web get_XRWebGLBinding_NativeProjectionScaleFactor
//go:noescape
func GetXRWebGLBindingNativeProjectionScaleFactor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XRWebGLBinding_UsesDepthValues
//go:noescape
func GetXRWebGLBindingUsesDepthValues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateProjectionLayer
//go:noescape
func HasFuncXRWebGLBindingCreateProjectionLayer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateProjectionLayer
//go:noescape
func FuncXRWebGLBindingCreateProjectionLayer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateProjectionLayer
//go:noescape
func CallXRWebGLBindingCreateProjectionLayer(
	this js.Ref, retPtr unsafe.Pointer,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateProjectionLayer
//go:noescape
func TryXRWebGLBindingCreateProjectionLayer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateProjectionLayer1
//go:noescape
func HasFuncXRWebGLBindingCreateProjectionLayer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateProjectionLayer1
//go:noescape
func FuncXRWebGLBindingCreateProjectionLayer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateProjectionLayer1
//go:noescape
func CallXRWebGLBindingCreateProjectionLayer1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateProjectionLayer1
//go:noescape
func TryXRWebGLBindingCreateProjectionLayer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateQuadLayer
//go:noescape
func HasFuncXRWebGLBindingCreateQuadLayer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateQuadLayer
//go:noescape
func FuncXRWebGLBindingCreateQuadLayer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateQuadLayer
//go:noescape
func CallXRWebGLBindingCreateQuadLayer(
	this js.Ref, retPtr unsafe.Pointer,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateQuadLayer
//go:noescape
func TryXRWebGLBindingCreateQuadLayer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateQuadLayer1
//go:noescape
func HasFuncXRWebGLBindingCreateQuadLayer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateQuadLayer1
//go:noescape
func FuncXRWebGLBindingCreateQuadLayer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateQuadLayer1
//go:noescape
func CallXRWebGLBindingCreateQuadLayer1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateQuadLayer1
//go:noescape
func TryXRWebGLBindingCreateQuadLayer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateCylinderLayer
//go:noescape
func HasFuncXRWebGLBindingCreateCylinderLayer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateCylinderLayer
//go:noescape
func FuncXRWebGLBindingCreateCylinderLayer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateCylinderLayer
//go:noescape
func CallXRWebGLBindingCreateCylinderLayer(
	this js.Ref, retPtr unsafe.Pointer,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateCylinderLayer
//go:noescape
func TryXRWebGLBindingCreateCylinderLayer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateCylinderLayer1
//go:noescape
func HasFuncXRWebGLBindingCreateCylinderLayer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateCylinderLayer1
//go:noescape
func FuncXRWebGLBindingCreateCylinderLayer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateCylinderLayer1
//go:noescape
func CallXRWebGLBindingCreateCylinderLayer1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateCylinderLayer1
//go:noescape
func TryXRWebGLBindingCreateCylinderLayer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateEquirectLayer
//go:noescape
func HasFuncXRWebGLBindingCreateEquirectLayer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateEquirectLayer
//go:noescape
func FuncXRWebGLBindingCreateEquirectLayer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateEquirectLayer
//go:noescape
func CallXRWebGLBindingCreateEquirectLayer(
	this js.Ref, retPtr unsafe.Pointer,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateEquirectLayer
//go:noescape
func TryXRWebGLBindingCreateEquirectLayer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateEquirectLayer1
//go:noescape
func HasFuncXRWebGLBindingCreateEquirectLayer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateEquirectLayer1
//go:noescape
func FuncXRWebGLBindingCreateEquirectLayer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateEquirectLayer1
//go:noescape
func CallXRWebGLBindingCreateEquirectLayer1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateEquirectLayer1
//go:noescape
func TryXRWebGLBindingCreateEquirectLayer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateCubeLayer
//go:noescape
func HasFuncXRWebGLBindingCreateCubeLayer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateCubeLayer
//go:noescape
func FuncXRWebGLBindingCreateCubeLayer(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateCubeLayer
//go:noescape
func CallXRWebGLBindingCreateCubeLayer(
	this js.Ref, retPtr unsafe.Pointer,
	init unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateCubeLayer
//go:noescape
func TryXRWebGLBindingCreateCubeLayer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	init unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_CreateCubeLayer1
//go:noescape
func HasFuncXRWebGLBindingCreateCubeLayer1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_CreateCubeLayer1
//go:noescape
func FuncXRWebGLBindingCreateCubeLayer1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_CreateCubeLayer1
//go:noescape
func CallXRWebGLBindingCreateCubeLayer1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XRWebGLBinding_CreateCubeLayer1
//go:noescape
func TryXRWebGLBindingCreateCubeLayer1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_GetSubImage
//go:noescape
func HasFuncXRWebGLBindingGetSubImage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetSubImage
//go:noescape
func FuncXRWebGLBindingGetSubImage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_GetSubImage
//go:noescape
func CallXRWebGLBindingGetSubImage(
	this js.Ref, retPtr unsafe.Pointer,
	layer js.Ref,
	frame js.Ref,
	eye uint32)

//go:wasmimport plat/js/web try_XRWebGLBinding_GetSubImage
//go:noescape
func TryXRWebGLBindingGetSubImage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	layer js.Ref,
	frame js.Ref,
	eye uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_GetSubImage1
//go:noescape
func HasFuncXRWebGLBindingGetSubImage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetSubImage1
//go:noescape
func FuncXRWebGLBindingGetSubImage1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_GetSubImage1
//go:noescape
func CallXRWebGLBindingGetSubImage1(
	this js.Ref, retPtr unsafe.Pointer,
	layer js.Ref,
	frame js.Ref)

//go:wasmimport plat/js/web try_XRWebGLBinding_GetSubImage1
//go:noescape
func TryXRWebGLBindingGetSubImage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	layer js.Ref,
	frame js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_GetViewSubImage
//go:noescape
func HasFuncXRWebGLBindingGetViewSubImage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetViewSubImage
//go:noescape
func FuncXRWebGLBindingGetViewSubImage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_GetViewSubImage
//go:noescape
func CallXRWebGLBindingGetViewSubImage(
	this js.Ref, retPtr unsafe.Pointer,
	layer js.Ref,
	view js.Ref)

//go:wasmimport plat/js/web try_XRWebGLBinding_GetViewSubImage
//go:noescape
func TryXRWebGLBindingGetViewSubImage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	layer js.Ref,
	view js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_GetCameraImage
//go:noescape
func HasFuncXRWebGLBindingGetCameraImage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetCameraImage
//go:noescape
func FuncXRWebGLBindingGetCameraImage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_GetCameraImage
//go:noescape
func CallXRWebGLBindingGetCameraImage(
	this js.Ref, retPtr unsafe.Pointer,
	camera js.Ref)

//go:wasmimport plat/js/web try_XRWebGLBinding_GetCameraImage
//go:noescape
func TryXRWebGLBindingGetCameraImage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	camera js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_GetDepthInformation
//go:noescape
func HasFuncXRWebGLBindingGetDepthInformation(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetDepthInformation
//go:noescape
func FuncXRWebGLBindingGetDepthInformation(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_GetDepthInformation
//go:noescape
func CallXRWebGLBindingGetDepthInformation(
	this js.Ref, retPtr unsafe.Pointer,
	view js.Ref)

//go:wasmimport plat/js/web try_XRWebGLBinding_GetDepthInformation
//go:noescape
func TryXRWebGLBindingGetDepthInformation(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	view js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XRWebGLBinding_GetReflectionCubeMap
//go:noescape
func HasFuncXRWebGLBindingGetReflectionCubeMap(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XRWebGLBinding_GetReflectionCubeMap
//go:noescape
func FuncXRWebGLBindingGetReflectionCubeMap(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XRWebGLBinding_GetReflectionCubeMap
//go:noescape
func CallXRWebGLBindingGetReflectionCubeMap(
	this js.Ref, retPtr unsafe.Pointer,
	lightProbe js.Ref)

//go:wasmimport plat/js/web try_XRWebGLBinding_GetReflectionCubeMap
//go:noescape
func TryXRWebGLBindingGetReflectionCubeMap(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	lightProbe js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XSLTProcessor_ImportStylesheet
//go:noescape
func HasFuncXSLTProcessorImportStylesheet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_ImportStylesheet
//go:noescape
func FuncXSLTProcessorImportStylesheet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XSLTProcessor_ImportStylesheet
//go:noescape
func CallXSLTProcessorImportStylesheet(
	this js.Ref, retPtr unsafe.Pointer,
	style js.Ref)

//go:wasmimport plat/js/web try_XSLTProcessor_ImportStylesheet
//go:noescape
func TryXSLTProcessorImportStylesheet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	style js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XSLTProcessor_TransformToFragment
//go:noescape
func HasFuncXSLTProcessorTransformToFragment(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_TransformToFragment
//go:noescape
func FuncXSLTProcessorTransformToFragment(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XSLTProcessor_TransformToFragment
//go:noescape
func CallXSLTProcessorTransformToFragment(
	this js.Ref, retPtr unsafe.Pointer,
	source js.Ref,
	output js.Ref)

//go:wasmimport plat/js/web try_XSLTProcessor_TransformToFragment
//go:noescape
func TryXSLTProcessorTransformToFragment(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref,
	output js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XSLTProcessor_TransformToDocument
//go:noescape
func HasFuncXSLTProcessorTransformToDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_TransformToDocument
//go:noescape
func FuncXSLTProcessorTransformToDocument(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XSLTProcessor_TransformToDocument
//go:noescape
func CallXSLTProcessorTransformToDocument(
	this js.Ref, retPtr unsafe.Pointer,
	source js.Ref)

//go:wasmimport plat/js/web try_XSLTProcessor_TransformToDocument
//go:noescape
func TryXSLTProcessorTransformToDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	source js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XSLTProcessor_SetParameter
//go:noescape
func HasFuncXSLTProcessorSetParameter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_SetParameter
//go:noescape
func FuncXSLTProcessorSetParameter(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XSLTProcessor_SetParameter
//go:noescape
func CallXSLTProcessorSetParameter(
	this js.Ref, retPtr unsafe.Pointer,
	namespaceURI js.Ref,
	localName js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_XSLTProcessor_SetParameter
//go:noescape
func TryXSLTProcessorSetParameter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespaceURI js.Ref,
	localName js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XSLTProcessor_GetParameter
//go:noescape
func HasFuncXSLTProcessorGetParameter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_GetParameter
//go:noescape
func FuncXSLTProcessorGetParameter(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XSLTProcessor_GetParameter
//go:noescape
func CallXSLTProcessorGetParameter(
	this js.Ref, retPtr unsafe.Pointer,
	namespaceURI js.Ref,
	localName js.Ref)

//go:wasmimport plat/js/web try_XSLTProcessor_GetParameter
//go:noescape
func TryXSLTProcessorGetParameter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespaceURI js.Ref,
	localName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XSLTProcessor_RemoveParameter
//go:noescape
func HasFuncXSLTProcessorRemoveParameter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_RemoveParameter
//go:noescape
func FuncXSLTProcessorRemoveParameter(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XSLTProcessor_RemoveParameter
//go:noescape
func CallXSLTProcessorRemoveParameter(
	this js.Ref, retPtr unsafe.Pointer,
	namespaceURI js.Ref,
	localName js.Ref)

//go:wasmimport plat/js/web try_XSLTProcessor_RemoveParameter
//go:noescape
func TryXSLTProcessorRemoveParameter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespaceURI js.Ref,
	localName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XSLTProcessor_ClearParameters
//go:noescape
func HasFuncXSLTProcessorClearParameters(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_ClearParameters
//go:noescape
func FuncXSLTProcessorClearParameters(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XSLTProcessor_ClearParameters
//go:noescape
func CallXSLTProcessorClearParameters(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XSLTProcessor_ClearParameters
//go:noescape
func TryXSLTProcessorClearParameters(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XSLTProcessor_Reset
//go:noescape
func HasFuncXSLTProcessorReset(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XSLTProcessor_Reset
//go:noescape
func FuncXSLTProcessorReset(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_XSLTProcessor_Reset
//go:noescape
func CallXSLTProcessorReset(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XSLTProcessor_Reset
//go:noescape
func TryXSLTProcessorReset(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Console_Assert
//go:noescape
func HasFuncConsoleAssert() js.Ref

//go:wasmimport plat/js/web func_Console_Assert
//go:noescape
func FuncConsoleAssert(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Assert
//go:noescape
func CallConsoleAssert(
	retPtr unsafe.Pointer,
	condition js.Ref,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_Assert
//go:noescape
func TryConsoleAssert(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	condition js.Ref,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_Assert1
//go:noescape
func HasFuncConsoleAssert1() js.Ref

//go:wasmimport plat/js/web func_Console_Assert1
//go:noescape
func FuncConsoleAssert1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Assert1
//go:noescape
func CallConsoleAssert1(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_Assert1
//go:noescape
func TryConsoleAssert1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_Console_Clear
//go:noescape
func HasFuncConsoleClear() js.Ref

//go:wasmimport plat/js/web func_Console_Clear
//go:noescape
func FuncConsoleClear(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Clear
//go:noescape
func CallConsoleClear(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_Clear
//go:noescape
func TryConsoleClear(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_Console_Debug
//go:noescape
func HasFuncConsoleDebug() js.Ref

//go:wasmimport plat/js/web func_Console_Debug
//go:noescape
func FuncConsoleDebug(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Debug
//go:noescape
func CallConsoleDebug(
	retPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_Debug
//go:noescape
func TryConsoleDebug(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_Error
//go:noescape
func HasFuncConsoleError() js.Ref

//go:wasmimport plat/js/web func_Console_Error
//go:noescape
func FuncConsoleError(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Error
//go:noescape
func CallConsoleError(
	retPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_Error
//go:noescape
func TryConsoleError(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_Info
//go:noescape
func HasFuncConsoleInfo() js.Ref

//go:wasmimport plat/js/web func_Console_Info
//go:noescape
func FuncConsoleInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Info
//go:noescape
func CallConsoleInfo(
	retPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_Info
//go:noescape
func TryConsoleInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_Log
//go:noescape
func HasFuncConsoleLog() js.Ref

//go:wasmimport plat/js/web func_Console_Log
//go:noescape
func FuncConsoleLog(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Log
//go:noescape
func CallConsoleLog(
	retPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_Log
//go:noescape
func TryConsoleLog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_Table
//go:noescape
func HasFuncConsoleTable() js.Ref

//go:wasmimport plat/js/web func_Console_Table
//go:noescape
func FuncConsoleTable(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Table
//go:noescape
func CallConsoleTable(
	retPtr unsafe.Pointer,
	tabularData js.Ref,
	properties js.Ref)

//go:wasmimport plat/js/web try_Console_Table
//go:noescape
func TryConsoleTable(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabularData js.Ref,
	properties js.Ref) js.Ref

//go:wasmimport plat/js/web has_Console_Table1
//go:noescape
func HasFuncConsoleTable1() js.Ref

//go:wasmimport plat/js/web func_Console_Table1
//go:noescape
func FuncConsoleTable1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Table1
//go:noescape
func CallConsoleTable1(
	retPtr unsafe.Pointer,
	tabularData js.Ref)

//go:wasmimport plat/js/web try_Console_Table1
//go:noescape
func TryConsoleTable1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabularData js.Ref) js.Ref

//go:wasmimport plat/js/web has_Console_Table2
//go:noescape
func HasFuncConsoleTable2() js.Ref

//go:wasmimport plat/js/web func_Console_Table2
//go:noescape
func FuncConsoleTable2(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Table2
//go:noescape
func CallConsoleTable2(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_Table2
//go:noescape
func TryConsoleTable2(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_Console_Trace
//go:noescape
func HasFuncConsoleTrace() js.Ref

//go:wasmimport plat/js/web func_Console_Trace
//go:noescape
func FuncConsoleTrace(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Trace
//go:noescape
func CallConsoleTrace(
	retPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_Trace
//go:noescape
func TryConsoleTrace(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_Warn
//go:noescape
func HasFuncConsoleWarn() js.Ref

//go:wasmimport plat/js/web func_Console_Warn
//go:noescape
func FuncConsoleWarn(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Warn
//go:noescape
func CallConsoleWarn(
	retPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_Warn
//go:noescape
func TryConsoleWarn(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_Dir
//go:noescape
func HasFuncConsoleDir() js.Ref

//go:wasmimport plat/js/web func_Console_Dir
//go:noescape
func FuncConsoleDir(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Dir
//go:noescape
func CallConsoleDir(
	retPtr unsafe.Pointer,
	item js.Ref,
	options js.Ref)

//go:wasmimport plat/js/web try_Console_Dir
//go:noescape
func TryConsoleDir(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	item js.Ref,
	options js.Ref) js.Ref

//go:wasmimport plat/js/web has_Console_Dir1
//go:noescape
func HasFuncConsoleDir1() js.Ref

//go:wasmimport plat/js/web func_Console_Dir1
//go:noescape
func FuncConsoleDir1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Dir1
//go:noescape
func CallConsoleDir1(
	retPtr unsafe.Pointer,
	item js.Ref)

//go:wasmimport plat/js/web try_Console_Dir1
//go:noescape
func TryConsoleDir1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	item js.Ref) js.Ref

//go:wasmimport plat/js/web has_Console_Dir2
//go:noescape
func HasFuncConsoleDir2() js.Ref

//go:wasmimport plat/js/web func_Console_Dir2
//go:noescape
func FuncConsoleDir2(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Dir2
//go:noescape
func CallConsoleDir2(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_Dir2
//go:noescape
func TryConsoleDir2(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_Console_Dirxml
//go:noescape
func HasFuncConsoleDirxml() js.Ref

//go:wasmimport plat/js/web func_Console_Dirxml
//go:noescape
func FuncConsoleDirxml(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Dirxml
//go:noescape
func CallConsoleDirxml(
	retPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_Dirxml
//go:noescape
func TryConsoleDirxml(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_Count
//go:noescape
func HasFuncConsoleCount() js.Ref

//go:wasmimport plat/js/web func_Console_Count
//go:noescape
func FuncConsoleCount(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Count
//go:noescape
func CallConsoleCount(
	retPtr unsafe.Pointer,
	label js.Ref)

//go:wasmimport plat/js/web try_Console_Count
//go:noescape
func TryConsoleCount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	label js.Ref) js.Ref

//go:wasmimport plat/js/web has_Console_Count1
//go:noescape
func HasFuncConsoleCount1() js.Ref

//go:wasmimport plat/js/web func_Console_Count1
//go:noescape
func FuncConsoleCount1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Count1
//go:noescape
func CallConsoleCount1(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_Count1
//go:noescape
func TryConsoleCount1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_Console_CountReset
//go:noescape
func HasFuncConsoleCountReset() js.Ref

//go:wasmimport plat/js/web func_Console_CountReset
//go:noescape
func FuncConsoleCountReset(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_CountReset
//go:noescape
func CallConsoleCountReset(
	retPtr unsafe.Pointer,
	label js.Ref)

//go:wasmimport plat/js/web try_Console_CountReset
//go:noescape
func TryConsoleCountReset(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	label js.Ref) js.Ref

//go:wasmimport plat/js/web has_Console_CountReset1
//go:noescape
func HasFuncConsoleCountReset1() js.Ref

//go:wasmimport plat/js/web func_Console_CountReset1
//go:noescape
func FuncConsoleCountReset1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_CountReset1
//go:noescape
func CallConsoleCountReset1(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_CountReset1
//go:noescape
func TryConsoleCountReset1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_Console_Group
//go:noescape
func HasFuncConsoleGroup() js.Ref

//go:wasmimport plat/js/web func_Console_Group
//go:noescape
func FuncConsoleGroup(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Group
//go:noescape
func CallConsoleGroup(
	retPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_Group
//go:noescape
func TryConsoleGroup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_GroupCollapsed
//go:noescape
func HasFuncConsoleGroupCollapsed() js.Ref

//go:wasmimport plat/js/web func_Console_GroupCollapsed
//go:noescape
func FuncConsoleGroupCollapsed(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_GroupCollapsed
//go:noescape
func CallConsoleGroupCollapsed(
	retPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_GroupCollapsed
//go:noescape
func TryConsoleGroupCollapsed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_GroupEnd
//go:noescape
func HasFuncConsoleGroupEnd() js.Ref

//go:wasmimport plat/js/web func_Console_GroupEnd
//go:noescape
func FuncConsoleGroupEnd(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_GroupEnd
//go:noescape
func CallConsoleGroupEnd(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_GroupEnd
//go:noescape
func TryConsoleGroupEnd(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_Console_Time
//go:noescape
func HasFuncConsoleTime() js.Ref

//go:wasmimport plat/js/web func_Console_Time
//go:noescape
func FuncConsoleTime(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Time
//go:noescape
func CallConsoleTime(
	retPtr unsafe.Pointer,
	label js.Ref)

//go:wasmimport plat/js/web try_Console_Time
//go:noescape
func TryConsoleTime(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	label js.Ref) js.Ref

//go:wasmimport plat/js/web has_Console_Time1
//go:noescape
func HasFuncConsoleTime1() js.Ref

//go:wasmimport plat/js/web func_Console_Time1
//go:noescape
func FuncConsoleTime1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_Time1
//go:noescape
func CallConsoleTime1(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_Time1
//go:noescape
func TryConsoleTime1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_Console_TimeLog
//go:noescape
func HasFuncConsoleTimeLog() js.Ref

//go:wasmimport plat/js/web func_Console_TimeLog
//go:noescape
func FuncConsoleTimeLog(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_TimeLog
//go:noescape
func CallConsoleTimeLog(
	retPtr unsafe.Pointer,
	label js.Ref,
	dataPtr unsafe.Pointer,
	dataCount uint32)

//go:wasmimport plat/js/web try_Console_TimeLog
//go:noescape
func TryConsoleTimeLog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	label js.Ref,
	dataPtr unsafe.Pointer,
	dataCount uint32) js.Ref

//go:wasmimport plat/js/web has_Console_TimeLog1
//go:noescape
func HasFuncConsoleTimeLog1() js.Ref

//go:wasmimport plat/js/web func_Console_TimeLog1
//go:noescape
func FuncConsoleTimeLog1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_TimeLog1
//go:noescape
func CallConsoleTimeLog1(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_TimeLog1
//go:noescape
func TryConsoleTimeLog1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_Console_TimeEnd
//go:noescape
func HasFuncConsoleTimeEnd() js.Ref

//go:wasmimport plat/js/web func_Console_TimeEnd
//go:noescape
func FuncConsoleTimeEnd(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_TimeEnd
//go:noescape
func CallConsoleTimeEnd(
	retPtr unsafe.Pointer,
	label js.Ref)

//go:wasmimport plat/js/web try_Console_TimeEnd
//go:noescape
func TryConsoleTimeEnd(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	label js.Ref) js.Ref

//go:wasmimport plat/js/web has_Console_TimeEnd1
//go:noescape
func HasFuncConsoleTimeEnd1() js.Ref

//go:wasmimport plat/js/web func_Console_TimeEnd1
//go:noescape
func FuncConsoleTimeEnd1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Console_TimeEnd1
//go:noescape
func CallConsoleTimeEnd1(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Console_TimeEnd1
//go:noescape
func TryConsoleTimeEnd1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) js.Ref
