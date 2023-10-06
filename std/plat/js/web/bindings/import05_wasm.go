// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web constof_DocumentVisibilityState
//go:noescape
func ConstOfDocumentVisibilityState(str js.Ref) uint32

//go:wasmimport plat/js/web store_FontFaceDescriptors
//go:noescape
func FontFaceDescriptorsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FontFaceDescriptors
//go:noescape
func FontFaceDescriptorsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_FontFaceLoadStatus
//go:noescape
func ConstOfFontFaceLoadStatus(str js.Ref) uint32

//go:wasmimport plat/js/web get_FontFacePalette_Length
//go:noescape
func GetFontFacePaletteLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontFacePalette_UsableWithLightBackground
//go:noescape
func GetFontFacePaletteUsableWithLightBackground(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontFacePalette_UsableWithDarkBackground
//go:noescape
func GetFontFacePaletteUsableWithDarkBackground(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFacePalette_Get
//go:noescape
func HasFontFacePaletteGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFacePalette_Get
//go:noescape
func FontFacePaletteGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFacePalette_Get
//go:noescape
func CallFontFacePaletteGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_FontFacePalette_Get
//go:noescape
func TryFontFacePaletteGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web get_FontFacePalettes_Length
//go:noescape
func GetFontFacePalettesLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFacePalettes_Get
//go:noescape
func HasFontFacePalettesGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFacePalettes_Get
//go:noescape
func FontFacePalettesGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFacePalettes_Get
//go:noescape
func CallFontFacePalettesGet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_FontFacePalettes_Get
//go:noescape
func TryFontFacePalettesGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web new_FontFace_FontFace
//go:noescape
func NewFontFaceByFontFace(
	family js.Ref,
	source js.Ref,
	descriptors unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_FontFace_FontFace1
//go:noescape
func NewFontFaceByFontFace1(
	family js.Ref,
	source js.Ref) js.Ref

//go:wasmimport plat/js/web get_FontFace_Family
//go:noescape
func GetFontFaceFamily(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_Family
//go:noescape
func SetFontFaceFamily(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Style
//go:noescape
func GetFontFaceStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_Style
//go:noescape
func SetFontFaceStyle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Weight
//go:noescape
func GetFontFaceWeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_Weight
//go:noescape
func SetFontFaceWeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Stretch
//go:noescape
func GetFontFaceStretch(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_Stretch
//go:noescape
func SetFontFaceStretch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_UnicodeRange
//go:noescape
func GetFontFaceUnicodeRange(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_UnicodeRange
//go:noescape
func SetFontFaceUnicodeRange(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Variant
//go:noescape
func GetFontFaceVariant(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_Variant
//go:noescape
func SetFontFaceVariant(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_FeatureSettings
//go:noescape
func GetFontFaceFeatureSettings(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_FeatureSettings
//go:noescape
func SetFontFaceFeatureSettings(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_VariationSettings
//go:noescape
func GetFontFaceVariationSettings(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_VariationSettings
//go:noescape
func SetFontFaceVariationSettings(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Display
//go:noescape
func GetFontFaceDisplay(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_Display
//go:noescape
func SetFontFaceDisplay(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_AscentOverride
//go:noescape
func GetFontFaceAscentOverride(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_AscentOverride
//go:noescape
func SetFontFaceAscentOverride(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_DescentOverride
//go:noescape
func GetFontFaceDescentOverride(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_DescentOverride
//go:noescape
func SetFontFaceDescentOverride(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_LineGapOverride
//go:noescape
func GetFontFaceLineGapOverride(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_FontFace_LineGapOverride
//go:noescape
func SetFontFaceLineGapOverride(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Status
//go:noescape
func GetFontFaceStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontFace_Loaded
//go:noescape
func GetFontFaceLoaded(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontFace_Features
//go:noescape
func GetFontFaceFeatures(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontFace_Variations
//go:noescape
func GetFontFaceVariations(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontFace_Palettes
//go:noescape
func GetFontFacePalettes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFace_Load
//go:noescape
func HasFontFaceLoad(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFace_Load
//go:noescape
func FontFaceLoadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFace_Load
//go:noescape
func CallFontFaceLoad(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FontFace_Load
//go:noescape
func TryFontFaceLoad(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_FontFaceSetLoadStatus
//go:noescape
func ConstOfFontFaceSetLoadStatus(str js.Ref) uint32

//go:wasmimport plat/js/web new_FontFaceSet_FontFaceSet
//go:noescape
func NewFontFaceSetByFontFaceSet(
	initialFaces js.Ref) js.Ref

//go:wasmimport plat/js/web get_FontFaceSet_Ready
//go:noescape
func GetFontFaceSetReady(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontFaceSet_Status
//go:noescape
func GetFontFaceSetStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFaceSet_Add
//go:noescape
func HasFontFaceSetAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Add
//go:noescape
func FontFaceSetAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Add
//go:noescape
func CallFontFaceSetAdd(
	this js.Ref, retPtr unsafe.Pointer,
	font js.Ref)

//go:wasmimport plat/js/web try_FontFaceSet_Add
//go:noescape
func TryFontFaceSetAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	font js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFaceSet_Delete
//go:noescape
func HasFontFaceSetDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Delete
//go:noescape
func FontFaceSetDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Delete
//go:noescape
func CallFontFaceSetDelete(
	this js.Ref, retPtr unsafe.Pointer,
	font js.Ref)

//go:wasmimport plat/js/web try_FontFaceSet_Delete
//go:noescape
func TryFontFaceSetDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	font js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFaceSet_Clear
//go:noescape
func HasFontFaceSetClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Clear
//go:noescape
func FontFaceSetClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Clear
//go:noescape
func CallFontFaceSetClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_FontFaceSet_Clear
//go:noescape
func TryFontFaceSetClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFaceSet_Load
//go:noescape
func HasFontFaceSetLoad(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Load
//go:noescape
func FontFaceSetLoadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Load
//go:noescape
func CallFontFaceSetLoad(
	this js.Ref, retPtr unsafe.Pointer,
	font js.Ref,
	text js.Ref)

//go:wasmimport plat/js/web try_FontFaceSet_Load
//go:noescape
func TryFontFaceSetLoad(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	font js.Ref,
	text js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFaceSet_Load1
//go:noescape
func HasFontFaceSetLoad1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Load1
//go:noescape
func FontFaceSetLoad1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Load1
//go:noescape
func CallFontFaceSetLoad1(
	this js.Ref, retPtr unsafe.Pointer,
	font js.Ref)

//go:wasmimport plat/js/web try_FontFaceSet_Load1
//go:noescape
func TryFontFaceSetLoad1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	font js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFaceSet_Check
//go:noescape
func HasFontFaceSetCheck(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Check
//go:noescape
func FontFaceSetCheckFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Check
//go:noescape
func CallFontFaceSetCheck(
	this js.Ref, retPtr unsafe.Pointer,
	font js.Ref,
	text js.Ref)

//go:wasmimport plat/js/web try_FontFaceSet_Check
//go:noescape
func TryFontFaceSetCheck(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	font js.Ref,
	text js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_FontFaceSet_Check1
//go:noescape
func HasFontFaceSetCheck1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Check1
//go:noescape
func FontFaceSetCheck1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Check1
//go:noescape
func CallFontFaceSetCheck1(
	this js.Ref, retPtr unsafe.Pointer,
	font js.Ref)

//go:wasmimport plat/js/web try_FontFaceSet_Check1
//go:noescape
func TryFontFaceSetCheck1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	font js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Implementation
//go:noescape
func GetDocumentImplementation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_URL
//go:noescape
func GetDocumentURL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_DocumentURI
//go:noescape
func GetDocumentDocumentURI(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_CompatMode
//go:noescape
func GetDocumentCompatMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_CharacterSet
//go:noescape
func GetDocumentCharacterSet(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Charset
//go:noescape
func GetDocumentCharset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_InputEncoding
//go:noescape
func GetDocumentInputEncoding(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_ContentType
//go:noescape
func GetDocumentContentType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Doctype
//go:noescape
func GetDocumentDoctype(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_DocumentElement
//go:noescape
func GetDocumentDocumentElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_FragmentDirective
//go:noescape
func GetDocumentFragmentDirective(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_PermissionsPolicy
//go:noescape
func GetDocumentPermissionsPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_WasDiscarded
//go:noescape
func GetDocumentWasDiscarded(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_FullscreenEnabled
//go:noescape
func GetDocumentFullscreenEnabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Fullscreen
//go:noescape
func GetDocumentFullscreen(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Timeline
//go:noescape
func GetDocumentTimeline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_PictureInPictureEnabled
//go:noescape
func GetDocumentPictureInPictureEnabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_NamedFlows
//go:noescape
func GetDocumentNamedFlows(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_ScrollingElement
//go:noescape
func GetDocumentScrollingElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_RootElement
//go:noescape
func GetDocumentRootElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Prerendering
//go:noescape
func GetDocumentPrerendering(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_FgColor
//go:noescape
func GetDocumentFgColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_FgColor
//go:noescape
func SetDocumentFgColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_LinkColor
//go:noescape
func GetDocumentLinkColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_LinkColor
//go:noescape
func SetDocumentLinkColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_VlinkColor
//go:noescape
func GetDocumentVlinkColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_VlinkColor
//go:noescape
func SetDocumentVlinkColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_AlinkColor
//go:noescape
func GetDocumentAlinkColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_AlinkColor
//go:noescape
func SetDocumentAlinkColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_BgColor
//go:noescape
func GetDocumentBgColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_BgColor
//go:noescape
func SetDocumentBgColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Anchors
//go:noescape
func GetDocumentAnchors(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Applets
//go:noescape
func GetDocumentApplets(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_All
//go:noescape
func GetDocumentAll(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Location
//go:noescape
func GetDocumentLocation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Domain
//go:noescape
func GetDocumentDomain(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_Domain
//go:noescape
func SetDocumentDomain(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Referrer
//go:noescape
func GetDocumentReferrer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Cookie
//go:noescape
func GetDocumentCookie(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_Cookie
//go:noescape
func SetDocumentCookie(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_LastModified
//go:noescape
func GetDocumentLastModified(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_ReadyState
//go:noescape
func GetDocumentReadyState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Title
//go:noescape
func GetDocumentTitle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_Title
//go:noescape
func SetDocumentTitle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Dir
//go:noescape
func GetDocumentDir(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_Dir
//go:noescape
func SetDocumentDir(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Body
//go:noescape
func GetDocumentBody(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_Body
//go:noescape
func SetDocumentBody(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Head
//go:noescape
func GetDocumentHead(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Images
//go:noescape
func GetDocumentImages(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Embeds
//go:noescape
func GetDocumentEmbeds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Plugins
//go:noescape
func GetDocumentPlugins(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Links
//go:noescape
func GetDocumentLinks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Forms
//go:noescape
func GetDocumentForms(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Scripts
//go:noescape
func GetDocumentScripts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_CurrentScript
//go:noescape
func GetDocumentCurrentScript(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_DefaultView
//go:noescape
func GetDocumentDefaultView(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_DesignMode
//go:noescape
func GetDocumentDesignMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_DesignMode
//go:noescape
func SetDocumentDesignMode(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Hidden
//go:noescape
func GetDocumentHidden(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_VisibilityState
//go:noescape
func GetDocumentVisibilityState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_FullscreenElement
//go:noescape
func GetDocumentFullscreenElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_StyleSheets
//go:noescape
func GetDocumentStyleSheets(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_AdoptedStyleSheets
//go:noescape
func GetDocumentAdoptedStyleSheets(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Document_AdoptedStyleSheets
//go:noescape
func SetDocumentAdoptedStyleSheets(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_PictureInPictureElement
//go:noescape
func GetDocumentPictureInPictureElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_ActiveElement
//go:noescape
func GetDocumentActiveElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_PointerLockElement
//go:noescape
func GetDocumentPointerLockElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Children
//go:noescape
func GetDocumentChildren(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_FirstElementChild
//go:noescape
func GetDocumentFirstElementChild(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_LastElementChild
//go:noescape
func GetDocumentLastElementChild(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_ChildElementCount
//go:noescape
func GetDocumentChildElementCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Document_Fonts
//go:noescape
func GetDocumentFonts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_GetElementsByTagName
//go:noescape
func HasDocumentGetElementsByTagName(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementsByTagName
//go:noescape
func DocumentGetElementsByTagNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementsByTagName
//go:noescape
func CallDocumentGetElementsByTagName(
	this js.Ref, retPtr unsafe.Pointer,
	qualifiedName js.Ref)

//go:wasmimport plat/js/web try_Document_GetElementsByTagName
//go:noescape
func TryDocumentGetElementsByTagName(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	qualifiedName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_GetElementsByTagNameNS
//go:noescape
func HasDocumentGetElementsByTagNameNS(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementsByTagNameNS
//go:noescape
func DocumentGetElementsByTagNameNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementsByTagNameNS
//go:noescape
func CallDocumentGetElementsByTagNameNS(
	this js.Ref, retPtr unsafe.Pointer,
	namespace js.Ref,
	localName js.Ref)

//go:wasmimport plat/js/web try_Document_GetElementsByTagNameNS
//go:noescape
func TryDocumentGetElementsByTagNameNS(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespace js.Ref,
	localName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_GetElementsByClassName
//go:noescape
func HasDocumentGetElementsByClassName(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementsByClassName
//go:noescape
func DocumentGetElementsByClassNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementsByClassName
//go:noescape
func CallDocumentGetElementsByClassName(
	this js.Ref, retPtr unsafe.Pointer,
	classNames js.Ref)

//go:wasmimport plat/js/web try_Document_GetElementsByClassName
//go:noescape
func TryDocumentGetElementsByClassName(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	classNames js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateElement
//go:noescape
func HasDocumentCreateElement(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateElement
//go:noescape
func DocumentCreateElementFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateElement
//go:noescape
func CallDocumentCreateElement(
	this js.Ref, retPtr unsafe.Pointer,
	localName js.Ref,
	options js.Ref)

//go:wasmimport plat/js/web try_Document_CreateElement
//go:noescape
func TryDocumentCreateElement(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	localName js.Ref,
	options js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateElement1
//go:noescape
func HasDocumentCreateElement1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateElement1
//go:noescape
func DocumentCreateElement1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateElement1
//go:noescape
func CallDocumentCreateElement1(
	this js.Ref, retPtr unsafe.Pointer,
	localName js.Ref)

//go:wasmimport plat/js/web try_Document_CreateElement1
//go:noescape
func TryDocumentCreateElement1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	localName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateElementNS
//go:noescape
func HasDocumentCreateElementNS(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateElementNS
//go:noescape
func DocumentCreateElementNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateElementNS
//go:noescape
func CallDocumentCreateElementNS(
	this js.Ref, retPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref,
	options js.Ref)

//go:wasmimport plat/js/web try_Document_CreateElementNS
//go:noescape
func TryDocumentCreateElementNS(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref,
	options js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateElementNS1
//go:noescape
func HasDocumentCreateElementNS1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateElementNS1
//go:noescape
func DocumentCreateElementNS1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateElementNS1
//go:noescape
func CallDocumentCreateElementNS1(
	this js.Ref, retPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref)

//go:wasmimport plat/js/web try_Document_CreateElementNS1
//go:noescape
func TryDocumentCreateElementNS1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateDocumentFragment
//go:noescape
func HasDocumentCreateDocumentFragment(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateDocumentFragment
//go:noescape
func DocumentCreateDocumentFragmentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateDocumentFragment
//go:noescape
func CallDocumentCreateDocumentFragment(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_CreateDocumentFragment
//go:noescape
func TryDocumentCreateDocumentFragment(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateTextNode
//go:noescape
func HasDocumentCreateTextNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateTextNode
//go:noescape
func DocumentCreateTextNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateTextNode
//go:noescape
func CallDocumentCreateTextNode(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_Document_CreateTextNode
//go:noescape
func TryDocumentCreateTextNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateCDATASection
//go:noescape
func HasDocumentCreateCDATASection(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateCDATASection
//go:noescape
func DocumentCreateCDATASectionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateCDATASection
//go:noescape
func CallDocumentCreateCDATASection(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_Document_CreateCDATASection
//go:noescape
func TryDocumentCreateCDATASection(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateComment
//go:noescape
func HasDocumentCreateComment(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateComment
//go:noescape
func DocumentCreateCommentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateComment
//go:noescape
func CallDocumentCreateComment(
	this js.Ref, retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/web try_Document_CreateComment
//go:noescape
func TryDocumentCreateComment(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateProcessingInstruction
//go:noescape
func HasDocumentCreateProcessingInstruction(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateProcessingInstruction
//go:noescape
func DocumentCreateProcessingInstructionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateProcessingInstruction
//go:noescape
func CallDocumentCreateProcessingInstruction(
	this js.Ref, retPtr unsafe.Pointer,
	target js.Ref,
	data js.Ref)

//go:wasmimport plat/js/web try_Document_CreateProcessingInstruction
//go:noescape
func TryDocumentCreateProcessingInstruction(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	target js.Ref,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ImportNode
//go:noescape
func HasDocumentImportNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ImportNode
//go:noescape
func DocumentImportNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ImportNode
//go:noescape
func CallDocumentImportNode(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref,
	deep js.Ref)

//go:wasmimport plat/js/web try_Document_ImportNode
//go:noescape
func TryDocumentImportNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref,
	deep js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ImportNode1
//go:noescape
func HasDocumentImportNode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ImportNode1
//go:noescape
func DocumentImportNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ImportNode1
//go:noescape
func CallDocumentImportNode1(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref)

//go:wasmimport plat/js/web try_Document_ImportNode1
//go:noescape
func TryDocumentImportNode1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_AdoptNode
//go:noescape
func HasDocumentAdoptNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_AdoptNode
//go:noescape
func DocumentAdoptNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_AdoptNode
//go:noescape
func CallDocumentAdoptNode(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref)

//go:wasmimport plat/js/web try_Document_AdoptNode
//go:noescape
func TryDocumentAdoptNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateAttribute
//go:noescape
func HasDocumentCreateAttribute(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateAttribute
//go:noescape
func DocumentCreateAttributeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateAttribute
//go:noescape
func CallDocumentCreateAttribute(
	this js.Ref, retPtr unsafe.Pointer,
	localName js.Ref)

//go:wasmimport plat/js/web try_Document_CreateAttribute
//go:noescape
func TryDocumentCreateAttribute(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	localName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateAttributeNS
//go:noescape
func HasDocumentCreateAttributeNS(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateAttributeNS
//go:noescape
func DocumentCreateAttributeNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateAttributeNS
//go:noescape
func CallDocumentCreateAttributeNS(
	this js.Ref, retPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref)

//go:wasmimport plat/js/web try_Document_CreateAttributeNS
//go:noescape
func TryDocumentCreateAttributeNS(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateEvent
//go:noescape
func HasDocumentCreateEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateEvent
//go:noescape
func DocumentCreateEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateEvent
//go:noescape
func CallDocumentCreateEvent(
	this js.Ref, retPtr unsafe.Pointer,
	iface js.Ref)

//go:wasmimport plat/js/web try_Document_CreateEvent
//go:noescape
func TryDocumentCreateEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	iface js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateRange
//go:noescape
func HasDocumentCreateRange(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateRange
//go:noescape
func DocumentCreateRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateRange
//go:noescape
func CallDocumentCreateRange(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_CreateRange
//go:noescape
func TryDocumentCreateRange(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateNodeIterator
//go:noescape
func HasDocumentCreateNodeIterator(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateNodeIterator
//go:noescape
func DocumentCreateNodeIteratorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateNodeIterator
//go:noescape
func CallDocumentCreateNodeIterator(
	this js.Ref, retPtr unsafe.Pointer,
	root js.Ref,
	whatToShow uint32,
	filter js.Ref)

//go:wasmimport plat/js/web try_Document_CreateNodeIterator
//go:noescape
func TryDocumentCreateNodeIterator(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	root js.Ref,
	whatToShow uint32,
	filter js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateNodeIterator1
//go:noescape
func HasDocumentCreateNodeIterator1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateNodeIterator1
//go:noescape
func DocumentCreateNodeIterator1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateNodeIterator1
//go:noescape
func CallDocumentCreateNodeIterator1(
	this js.Ref, retPtr unsafe.Pointer,
	root js.Ref,
	whatToShow uint32)

//go:wasmimport plat/js/web try_Document_CreateNodeIterator1
//go:noescape
func TryDocumentCreateNodeIterator1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	root js.Ref,
	whatToShow uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateNodeIterator2
//go:noescape
func HasDocumentCreateNodeIterator2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateNodeIterator2
//go:noescape
func DocumentCreateNodeIterator2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateNodeIterator2
//go:noescape
func CallDocumentCreateNodeIterator2(
	this js.Ref, retPtr unsafe.Pointer,
	root js.Ref)

//go:wasmimport plat/js/web try_Document_CreateNodeIterator2
//go:noescape
func TryDocumentCreateNodeIterator2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	root js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateTreeWalker
//go:noescape
func HasDocumentCreateTreeWalker(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateTreeWalker
//go:noescape
func DocumentCreateTreeWalkerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateTreeWalker
//go:noescape
func CallDocumentCreateTreeWalker(
	this js.Ref, retPtr unsafe.Pointer,
	root js.Ref,
	whatToShow uint32,
	filter js.Ref)

//go:wasmimport plat/js/web try_Document_CreateTreeWalker
//go:noescape
func TryDocumentCreateTreeWalker(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	root js.Ref,
	whatToShow uint32,
	filter js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateTreeWalker1
//go:noescape
func HasDocumentCreateTreeWalker1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateTreeWalker1
//go:noescape
func DocumentCreateTreeWalker1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateTreeWalker1
//go:noescape
func CallDocumentCreateTreeWalker1(
	this js.Ref, retPtr unsafe.Pointer,
	root js.Ref,
	whatToShow uint32)

//go:wasmimport plat/js/web try_Document_CreateTreeWalker1
//go:noescape
func TryDocumentCreateTreeWalker1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	root js.Ref,
	whatToShow uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateTreeWalker2
//go:noescape
func HasDocumentCreateTreeWalker2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateTreeWalker2
//go:noescape
func DocumentCreateTreeWalker2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateTreeWalker2
//go:noescape
func CallDocumentCreateTreeWalker2(
	this js.Ref, retPtr unsafe.Pointer,
	root js.Ref)

//go:wasmimport plat/js/web try_Document_CreateTreeWalker2
//go:noescape
func TryDocumentCreateTreeWalker2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	root js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_StartViewTransition
//go:noescape
func HasDocumentStartViewTransition(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_StartViewTransition
//go:noescape
func DocumentStartViewTransitionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_StartViewTransition
//go:noescape
func CallDocumentStartViewTransition(
	this js.Ref, retPtr unsafe.Pointer,
	updateCallback js.Ref)

//go:wasmimport plat/js/web try_Document_StartViewTransition
//go:noescape
func TryDocumentStartViewTransition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	updateCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_StartViewTransition1
//go:noescape
func HasDocumentStartViewTransition1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_StartViewTransition1
//go:noescape
func DocumentStartViewTransition1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_StartViewTransition1
//go:noescape
func CallDocumentStartViewTransition1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_StartViewTransition1
//go:noescape
func TryDocumentStartViewTransition1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_MeasureElement
//go:noescape
func HasDocumentMeasureElement(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_MeasureElement
//go:noescape
func DocumentMeasureElementFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_MeasureElement
//go:noescape
func CallDocumentMeasureElement(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref)

//go:wasmimport plat/js/web try_Document_MeasureElement
//go:noescape
func TryDocumentMeasureElement(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_MeasureText
//go:noescape
func HasDocumentMeasureText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_MeasureText
//go:noescape
func DocumentMeasureTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_MeasureText
//go:noescape
func CallDocumentMeasureText(
	this js.Ref, retPtr unsafe.Pointer,
	text js.Ref,
	styleMap js.Ref)

//go:wasmimport plat/js/web try_Document_MeasureText
//go:noescape
func TryDocumentMeasureText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref,
	styleMap js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_GetSelection
//go:noescape
func HasDocumentGetSelection(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_GetSelection
//go:noescape
func DocumentGetSelectionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetSelection
//go:noescape
func CallDocumentGetSelection(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_GetSelection
//go:noescape
func TryDocumentGetSelection(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ExitPointerLock
//go:noescape
func HasDocumentExitPointerLock(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ExitPointerLock
//go:noescape
func DocumentExitPointerLockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExitPointerLock
//go:noescape
func CallDocumentExitPointerLock(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_ExitPointerLock
//go:noescape
func TryDocumentExitPointerLock(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ExitFullscreen
//go:noescape
func HasDocumentExitFullscreen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ExitFullscreen
//go:noescape
func DocumentExitFullscreenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExitFullscreen
//go:noescape
func CallDocumentExitFullscreen(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_ExitFullscreen
//go:noescape
func TryDocumentExitFullscreen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_HasStorageAccess
//go:noescape
func HasDocumentHasStorageAccess(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_HasStorageAccess
//go:noescape
func DocumentHasStorageAccessFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_HasStorageAccess
//go:noescape
func CallDocumentHasStorageAccess(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_HasStorageAccess
//go:noescape
func TryDocumentHasStorageAccess(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_RequestStorageAccess
//go:noescape
func HasDocumentRequestStorageAccess(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_RequestStorageAccess
//go:noescape
func DocumentRequestStorageAccessFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_RequestStorageAccess
//go:noescape
func CallDocumentRequestStorageAccess(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_RequestStorageAccess
//go:noescape
func TryDocumentRequestStorageAccess(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ExitPictureInPicture
//go:noescape
func HasDocumentExitPictureInPicture(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ExitPictureInPicture
//go:noescape
func DocumentExitPictureInPictureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExitPictureInPicture
//go:noescape
func CallDocumentExitPictureInPicture(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_ExitPictureInPicture
//go:noescape
func TryDocumentExitPictureInPicture(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_HasPrivateTokens
//go:noescape
func HasDocumentHasPrivateTokens(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_HasPrivateTokens
//go:noescape
func DocumentHasPrivateTokensFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_HasPrivateTokens
//go:noescape
func CallDocumentHasPrivateTokens(
	this js.Ref, retPtr unsafe.Pointer,
	issuer js.Ref)

//go:wasmimport plat/js/web try_Document_HasPrivateTokens
//go:noescape
func TryDocumentHasPrivateTokens(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	issuer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_HasRedemptionRecord
//go:noescape
func HasDocumentHasRedemptionRecord(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_HasRedemptionRecord
//go:noescape
func DocumentHasRedemptionRecordFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_HasRedemptionRecord
//go:noescape
func CallDocumentHasRedemptionRecord(
	this js.Ref, retPtr unsafe.Pointer,
	issuer js.Ref)

//go:wasmimport plat/js/web try_Document_HasRedemptionRecord
//go:noescape
func TryDocumentHasRedemptionRecord(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	issuer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ElementFromPoint
//go:noescape
func HasDocumentElementFromPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ElementFromPoint
//go:noescape
func DocumentElementFromPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ElementFromPoint
//go:noescape
func CallDocumentElementFromPoint(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Document_ElementFromPoint
//go:noescape
func TryDocumentElementFromPoint(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ElementsFromPoint
//go:noescape
func HasDocumentElementsFromPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ElementsFromPoint
//go:noescape
func DocumentElementsFromPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ElementsFromPoint
//go:noescape
func CallDocumentElementsFromPoint(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Document_ElementsFromPoint
//go:noescape
func TryDocumentElementsFromPoint(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CaretPositionFromPoint
//go:noescape
func HasDocumentCaretPositionFromPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CaretPositionFromPoint
//go:noescape
func DocumentCaretPositionFromPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CaretPositionFromPoint
//go:noescape
func CallDocumentCaretPositionFromPoint(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_Document_CaretPositionFromPoint
//go:noescape
func TryDocumentCaretPositionFromPoint(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Clear
//go:noescape
func HasDocumentClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Clear
//go:noescape
func DocumentClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Clear
//go:noescape
func CallDocumentClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_Clear
//go:noescape
func TryDocumentClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CaptureEvents
//go:noescape
func HasDocumentCaptureEvents(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CaptureEvents
//go:noescape
func DocumentCaptureEventsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CaptureEvents
//go:noescape
func CallDocumentCaptureEvents(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_CaptureEvents
//go:noescape
func TryDocumentCaptureEvents(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ReleaseEvents
//go:noescape
func HasDocumentReleaseEvents(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ReleaseEvents
//go:noescape
func DocumentReleaseEventsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ReleaseEvents
//go:noescape
func CallDocumentReleaseEvents(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_ReleaseEvents
//go:noescape
func TryDocumentReleaseEvents(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_RequestStorageAccessFor
//go:noescape
func HasDocumentRequestStorageAccessFor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_RequestStorageAccessFor
//go:noescape
func DocumentRequestStorageAccessForFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_RequestStorageAccessFor
//go:noescape
func CallDocumentRequestStorageAccessFor(
	this js.Ref, retPtr unsafe.Pointer,
	requestedOrigin js.Ref)

//go:wasmimport plat/js/web try_Document_RequestStorageAccessFor
//go:noescape
func TryDocumentRequestStorageAccessFor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	requestedOrigin js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Get
//go:noescape
func HasDocumentGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Get
//go:noescape
func DocumentGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Get
//go:noescape
func CallDocumentGet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_Document_Get
//go:noescape
func TryDocumentGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_GetElementsByName
//go:noescape
func HasDocumentGetElementsByName(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementsByName
//go:noescape
func DocumentGetElementsByNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementsByName
//go:noescape
func CallDocumentGetElementsByName(
	this js.Ref, retPtr unsafe.Pointer,
	elementName js.Ref)

//go:wasmimport plat/js/web try_Document_GetElementsByName
//go:noescape
func TryDocumentGetElementsByName(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	elementName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Open
//go:noescape
func HasDocumentOpen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Open
//go:noescape
func DocumentOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Open
//go:noescape
func CallDocumentOpen(
	this js.Ref, retPtr unsafe.Pointer,
	unused1 js.Ref,
	unused2 js.Ref)

//go:wasmimport plat/js/web try_Document_Open
//go:noescape
func TryDocumentOpen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	unused1 js.Ref,
	unused2 js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Open1
//go:noescape
func HasDocumentOpen1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Open1
//go:noescape
func DocumentOpen1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Open1
//go:noescape
func CallDocumentOpen1(
	this js.Ref, retPtr unsafe.Pointer,
	unused1 js.Ref)

//go:wasmimport plat/js/web try_Document_Open1
//go:noescape
func TryDocumentOpen1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	unused1 js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Open2
//go:noescape
func HasDocumentOpen2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Open2
//go:noescape
func DocumentOpen2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Open2
//go:noescape
func CallDocumentOpen2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_Open2
//go:noescape
func TryDocumentOpen2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Open3
//go:noescape
func HasDocumentOpen3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Open3
//go:noescape
func DocumentOpen3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Open3
//go:noescape
func CallDocumentOpen3(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref,
	name js.Ref,
	features js.Ref)

//go:wasmimport plat/js/web try_Document_Open3
//go:noescape
func TryDocumentOpen3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	name js.Ref,
	features js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Close
//go:noescape
func HasDocumentClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Close
//go:noescape
func DocumentCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Close
//go:noescape
func CallDocumentClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_Close
//go:noescape
func TryDocumentClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Write
//go:noescape
func HasDocumentWrite(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Write
//go:noescape
func DocumentWriteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Write
//go:noescape
func CallDocumentWrite(
	this js.Ref, retPtr unsafe.Pointer,
	textPtr unsafe.Pointer,
	textCount uint32)

//go:wasmimport plat/js/web try_Document_Write
//go:noescape
func TryDocumentWrite(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	textPtr unsafe.Pointer,
	textCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Writeln
//go:noescape
func HasDocumentWriteln(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Writeln
//go:noescape
func DocumentWritelnFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Writeln
//go:noescape
func CallDocumentWriteln(
	this js.Ref, retPtr unsafe.Pointer,
	textPtr unsafe.Pointer,
	textCount uint32)

//go:wasmimport plat/js/web try_Document_Writeln
//go:noescape
func TryDocumentWriteln(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	textPtr unsafe.Pointer,
	textCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_HasFocus
//go:noescape
func HasDocumentHasFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_HasFocus
//go:noescape
func DocumentHasFocusFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_HasFocus
//go:noescape
func CallDocumentHasFocus(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_HasFocus
//go:noescape
func TryDocumentHasFocus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ExecCommand
//go:noescape
func HasDocumentExecCommand(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ExecCommand
//go:noescape
func DocumentExecCommandFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExecCommand
//go:noescape
func CallDocumentExecCommand(
	this js.Ref, retPtr unsafe.Pointer,
	commandId js.Ref,
	showUI js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_Document_ExecCommand
//go:noescape
func TryDocumentExecCommand(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	commandId js.Ref,
	showUI js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ExecCommand1
//go:noescape
func HasDocumentExecCommand1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ExecCommand1
//go:noescape
func DocumentExecCommand1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExecCommand1
//go:noescape
func CallDocumentExecCommand1(
	this js.Ref, retPtr unsafe.Pointer,
	commandId js.Ref,
	showUI js.Ref)

//go:wasmimport plat/js/web try_Document_ExecCommand1
//go:noescape
func TryDocumentExecCommand1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	commandId js.Ref,
	showUI js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ExecCommand2
//go:noescape
func HasDocumentExecCommand2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ExecCommand2
//go:noescape
func DocumentExecCommand2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExecCommand2
//go:noescape
func CallDocumentExecCommand2(
	this js.Ref, retPtr unsafe.Pointer,
	commandId js.Ref)

//go:wasmimport plat/js/web try_Document_ExecCommand2
//go:noescape
func TryDocumentExecCommand2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	commandId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_QueryCommandEnabled
//go:noescape
func HasDocumentQueryCommandEnabled(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandEnabled
//go:noescape
func DocumentQueryCommandEnabledFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandEnabled
//go:noescape
func CallDocumentQueryCommandEnabled(
	this js.Ref, retPtr unsafe.Pointer,
	commandId js.Ref)

//go:wasmimport plat/js/web try_Document_QueryCommandEnabled
//go:noescape
func TryDocumentQueryCommandEnabled(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	commandId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_QueryCommandIndeterm
//go:noescape
func HasDocumentQueryCommandIndeterm(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandIndeterm
//go:noescape
func DocumentQueryCommandIndetermFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandIndeterm
//go:noescape
func CallDocumentQueryCommandIndeterm(
	this js.Ref, retPtr unsafe.Pointer,
	commandId js.Ref)

//go:wasmimport plat/js/web try_Document_QueryCommandIndeterm
//go:noescape
func TryDocumentQueryCommandIndeterm(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	commandId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_QueryCommandState
//go:noescape
func HasDocumentQueryCommandState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandState
//go:noescape
func DocumentQueryCommandStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandState
//go:noescape
func CallDocumentQueryCommandState(
	this js.Ref, retPtr unsafe.Pointer,
	commandId js.Ref)

//go:wasmimport plat/js/web try_Document_QueryCommandState
//go:noescape
func TryDocumentQueryCommandState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	commandId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_QueryCommandSupported
//go:noescape
func HasDocumentQueryCommandSupported(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandSupported
//go:noescape
func DocumentQueryCommandSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandSupported
//go:noescape
func CallDocumentQueryCommandSupported(
	this js.Ref, retPtr unsafe.Pointer,
	commandId js.Ref)

//go:wasmimport plat/js/web try_Document_QueryCommandSupported
//go:noescape
func TryDocumentQueryCommandSupported(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	commandId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_QueryCommandValue
//go:noescape
func HasDocumentQueryCommandValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandValue
//go:noescape
func DocumentQueryCommandValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandValue
//go:noescape
func CallDocumentQueryCommandValue(
	this js.Ref, retPtr unsafe.Pointer,
	commandId js.Ref)

//go:wasmimport plat/js/web try_Document_QueryCommandValue
//go:noescape
func TryDocumentQueryCommandValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	commandId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_GetBoxQuads
//go:noescape
func HasDocumentGetBoxQuads(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_GetBoxQuads
//go:noescape
func DocumentGetBoxQuadsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetBoxQuads
//go:noescape
func CallDocumentGetBoxQuads(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_GetBoxQuads
//go:noescape
func TryDocumentGetBoxQuads(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_GetBoxQuads1
//go:noescape
func HasDocumentGetBoxQuads1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_GetBoxQuads1
//go:noescape
func DocumentGetBoxQuads1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetBoxQuads1
//go:noescape
func CallDocumentGetBoxQuads1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_GetBoxQuads1
//go:noescape
func TryDocumentGetBoxQuads1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ConvertQuadFromNode
//go:noescape
func HasDocumentConvertQuadFromNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertQuadFromNode
//go:noescape
func DocumentConvertQuadFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertQuadFromNode
//go:noescape
func CallDocumentConvertQuadFromNode(
	this js.Ref, retPtr unsafe.Pointer,
	quad unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_ConvertQuadFromNode
//go:noescape
func TryDocumentConvertQuadFromNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	quad unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ConvertQuadFromNode1
//go:noescape
func HasDocumentConvertQuadFromNode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertQuadFromNode1
//go:noescape
func DocumentConvertQuadFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertQuadFromNode1
//go:noescape
func CallDocumentConvertQuadFromNode1(
	this js.Ref, retPtr unsafe.Pointer,
	quad unsafe.Pointer,
	from js.Ref)

//go:wasmimport plat/js/web try_Document_ConvertQuadFromNode1
//go:noescape
func TryDocumentConvertQuadFromNode1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	quad unsafe.Pointer,
	from js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ConvertRectFromNode
//go:noescape
func HasDocumentConvertRectFromNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertRectFromNode
//go:noescape
func DocumentConvertRectFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertRectFromNode
//go:noescape
func CallDocumentConvertRectFromNode(
	this js.Ref, retPtr unsafe.Pointer,
	rect js.Ref,
	from js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_ConvertRectFromNode
//go:noescape
func TryDocumentConvertRectFromNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rect js.Ref,
	from js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ConvertRectFromNode1
//go:noescape
func HasDocumentConvertRectFromNode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertRectFromNode1
//go:noescape
func DocumentConvertRectFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertRectFromNode1
//go:noescape
func CallDocumentConvertRectFromNode1(
	this js.Ref, retPtr unsafe.Pointer,
	rect js.Ref,
	from js.Ref)

//go:wasmimport plat/js/web try_Document_ConvertRectFromNode1
//go:noescape
func TryDocumentConvertRectFromNode1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rect js.Ref,
	from js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ConvertPointFromNode
//go:noescape
func HasDocumentConvertPointFromNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertPointFromNode
//go:noescape
func DocumentConvertPointFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertPointFromNode
//go:noescape
func CallDocumentConvertPointFromNode(
	this js.Ref, retPtr unsafe.Pointer,
	point unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_ConvertPointFromNode
//go:noescape
func TryDocumentConvertPointFromNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	point unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ConvertPointFromNode1
//go:noescape
func HasDocumentConvertPointFromNode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertPointFromNode1
//go:noescape
func DocumentConvertPointFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertPointFromNode1
//go:noescape
func CallDocumentConvertPointFromNode1(
	this js.Ref, retPtr unsafe.Pointer,
	point unsafe.Pointer,
	from js.Ref)

//go:wasmimport plat/js/web try_Document_ConvertPointFromNode1
//go:noescape
func TryDocumentConvertPointFromNode1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	point unsafe.Pointer,
	from js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_GetElementById
//go:noescape
func HasDocumentGetElementById(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementById
//go:noescape
func DocumentGetElementByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementById
//go:noescape
func CallDocumentGetElementById(
	this js.Ref, retPtr unsafe.Pointer,
	elementId js.Ref)

//go:wasmimport plat/js/web try_Document_GetElementById
//go:noescape
func TryDocumentGetElementById(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	elementId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_GetAnimations
//go:noescape
func HasDocumentGetAnimations(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_GetAnimations
//go:noescape
func DocumentGetAnimationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetAnimations
//go:noescape
func CallDocumentGetAnimations(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Document_GetAnimations
//go:noescape
func TryDocumentGetAnimations(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Prepend
//go:noescape
func HasDocumentPrepend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Prepend
//go:noescape
func DocumentPrependFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Prepend
//go:noescape
func CallDocumentPrepend(
	this js.Ref, retPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32)

//go:wasmimport plat/js/web try_Document_Prepend
//go:noescape
func TryDocumentPrepend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Append
//go:noescape
func HasDocumentAppend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Append
//go:noescape
func DocumentAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Append
//go:noescape
func CallDocumentAppend(
	this js.Ref, retPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32)

//go:wasmimport plat/js/web try_Document_Append
//go:noescape
func TryDocumentAppend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_ReplaceChildren
//go:noescape
func HasDocumentReplaceChildren(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_ReplaceChildren
//go:noescape
func DocumentReplaceChildrenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ReplaceChildren
//go:noescape
func CallDocumentReplaceChildren(
	this js.Ref, retPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32)

//go:wasmimport plat/js/web try_Document_ReplaceChildren
//go:noescape
func TryDocumentReplaceChildren(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_QuerySelector
//go:noescape
func HasDocumentQuerySelector(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_QuerySelector
//go:noescape
func DocumentQuerySelectorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QuerySelector
//go:noescape
func CallDocumentQuerySelector(
	this js.Ref, retPtr unsafe.Pointer,
	selectors js.Ref)

//go:wasmimport plat/js/web try_Document_QuerySelector
//go:noescape
func TryDocumentQuerySelector(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	selectors js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_QuerySelectorAll
//go:noescape
func HasDocumentQuerySelectorAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_QuerySelectorAll
//go:noescape
func DocumentQuerySelectorAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QuerySelectorAll
//go:noescape
func CallDocumentQuerySelectorAll(
	this js.Ref, retPtr unsafe.Pointer,
	selectors js.Ref)

//go:wasmimport plat/js/web try_Document_QuerySelectorAll
//go:noescape
func TryDocumentQuerySelectorAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	selectors js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateExpression
//go:noescape
func HasDocumentCreateExpression(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateExpression
//go:noescape
func DocumentCreateExpressionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateExpression
//go:noescape
func CallDocumentCreateExpression(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	resolver js.Ref)

//go:wasmimport plat/js/web try_Document_CreateExpression
//go:noescape
func TryDocumentCreateExpression(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	resolver js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateExpression1
//go:noescape
func HasDocumentCreateExpression1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateExpression1
//go:noescape
func DocumentCreateExpression1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateExpression1
//go:noescape
func CallDocumentCreateExpression1(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref)

//go:wasmimport plat/js/web try_Document_CreateExpression1
//go:noescape
func TryDocumentCreateExpression1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_CreateNSResolver
//go:noescape
func HasDocumentCreateNSResolver(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_CreateNSResolver
//go:noescape
func DocumentCreateNSResolverFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateNSResolver
//go:noescape
func CallDocumentCreateNSResolver(
	this js.Ref, retPtr unsafe.Pointer,
	nodeResolver js.Ref)

//go:wasmimport plat/js/web try_Document_CreateNSResolver
//go:noescape
func TryDocumentCreateNSResolver(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodeResolver js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Evaluate
//go:noescape
func HasDocumentEvaluate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Evaluate
//go:noescape
func DocumentEvaluateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Evaluate
//go:noescape
func CallDocumentEvaluate(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32,
	result js.Ref)

//go:wasmimport plat/js/web try_Document_Evaluate
//go:noescape
func TryDocumentEvaluate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32,
	result js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Evaluate1
//go:noescape
func HasDocumentEvaluate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Evaluate1
//go:noescape
func DocumentEvaluate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Evaluate1
//go:noescape
func CallDocumentEvaluate1(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32)

//go:wasmimport plat/js/web try_Document_Evaluate1
//go:noescape
func TryDocumentEvaluate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Evaluate2
//go:noescape
func HasDocumentEvaluate2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Evaluate2
//go:noescape
func DocumentEvaluate2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Evaluate2
//go:noescape
func CallDocumentEvaluate2(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref)

//go:wasmimport plat/js/web try_Document_Evaluate2
//go:noescape
func TryDocumentEvaluate2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Document_Evaluate3
//go:noescape
func HasDocumentEvaluate3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Document_Evaluate3
//go:noescape
func DocumentEvaluate3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Evaluate3
//go:noescape
func CallDocumentEvaluate3(
	this js.Ref, retPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref)

//go:wasmimport plat/js/web try_Document_Evaluate3
//go:noescape
func TryDocumentEvaluate3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	expression js.Ref,
	contextNode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_NodeType
//go:noescape
func GetNodeNodeType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_NodeName
//go:noescape
func GetNodeNodeName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_BaseURI
//go:noescape
func GetNodeBaseURI(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_IsConnected
//go:noescape
func GetNodeIsConnected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_OwnerDocument
//go:noescape
func GetNodeOwnerDocument(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_ParentNode
//go:noescape
func GetNodeParentNode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_ParentElement
//go:noescape
func GetNodeParentElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_ChildNodes
//go:noescape
func GetNodeChildNodes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_FirstChild
//go:noescape
func GetNodeFirstChild(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_LastChild
//go:noescape
func GetNodeLastChild(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_PreviousSibling
//go:noescape
func GetNodePreviousSibling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_NextSibling
//go:noescape
func GetNodeNextSibling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Node_NodeValue
//go:noescape
func GetNodeNodeValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Node_NodeValue
//go:noescape
func SetNodeNodeValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Node_TextContent
//go:noescape
func GetNodeTextContent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Node_TextContent
//go:noescape
func SetNodeTextContent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_Node_GetRootNode
//go:noescape
func HasNodeGetRootNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_GetRootNode
//go:noescape
func NodeGetRootNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_GetRootNode
//go:noescape
func CallNodeGetRootNode(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Node_GetRootNode
//go:noescape
func TryNodeGetRootNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_GetRootNode1
//go:noescape
func HasNodeGetRootNode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_GetRootNode1
//go:noescape
func NodeGetRootNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_GetRootNode1
//go:noescape
func CallNodeGetRootNode1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Node_GetRootNode1
//go:noescape
func TryNodeGetRootNode1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_HasChildNodes
//go:noescape
func HasNodeHasChildNodes(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_HasChildNodes
//go:noescape
func NodeHasChildNodesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_HasChildNodes
//go:noescape
func CallNodeHasChildNodes(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Node_HasChildNodes
//go:noescape
func TryNodeHasChildNodes(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_Normalize
//go:noescape
func HasNodeNormalize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_Normalize
//go:noescape
func NodeNormalizeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_Normalize
//go:noescape
func CallNodeNormalize(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Node_Normalize
//go:noescape
func TryNodeNormalize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_CloneNode
//go:noescape
func HasNodeCloneNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_CloneNode
//go:noescape
func NodeCloneNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_CloneNode
//go:noescape
func CallNodeCloneNode(
	this js.Ref, retPtr unsafe.Pointer,
	deep js.Ref)

//go:wasmimport plat/js/web try_Node_CloneNode
//go:noescape
func TryNodeCloneNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deep js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_CloneNode1
//go:noescape
func HasNodeCloneNode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_CloneNode1
//go:noescape
func NodeCloneNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_CloneNode1
//go:noescape
func CallNodeCloneNode1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Node_CloneNode1
//go:noescape
func TryNodeCloneNode1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_IsEqualNode
//go:noescape
func HasNodeIsEqualNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_IsEqualNode
//go:noescape
func NodeIsEqualNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_IsEqualNode
//go:noescape
func CallNodeIsEqualNode(
	this js.Ref, retPtr unsafe.Pointer,
	otherNode js.Ref)

//go:wasmimport plat/js/web try_Node_IsEqualNode
//go:noescape
func TryNodeIsEqualNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	otherNode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_IsSameNode
//go:noescape
func HasNodeIsSameNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_IsSameNode
//go:noescape
func NodeIsSameNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_IsSameNode
//go:noescape
func CallNodeIsSameNode(
	this js.Ref, retPtr unsafe.Pointer,
	otherNode js.Ref)

//go:wasmimport plat/js/web try_Node_IsSameNode
//go:noescape
func TryNodeIsSameNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	otherNode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_CompareDocumentPosition
//go:noescape
func HasNodeCompareDocumentPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_CompareDocumentPosition
//go:noescape
func NodeCompareDocumentPositionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_CompareDocumentPosition
//go:noescape
func CallNodeCompareDocumentPosition(
	this js.Ref, retPtr unsafe.Pointer,
	other js.Ref)

//go:wasmimport plat/js/web try_Node_CompareDocumentPosition
//go:noescape
func TryNodeCompareDocumentPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	other js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_Contains
//go:noescape
func HasNodeContains(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_Contains
//go:noescape
func NodeContainsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_Contains
//go:noescape
func CallNodeContains(
	this js.Ref, retPtr unsafe.Pointer,
	other js.Ref)

//go:wasmimport plat/js/web try_Node_Contains
//go:noescape
func TryNodeContains(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	other js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_LookupPrefix
//go:noescape
func HasNodeLookupPrefix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_LookupPrefix
//go:noescape
func NodeLookupPrefixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_LookupPrefix
//go:noescape
func CallNodeLookupPrefix(
	this js.Ref, retPtr unsafe.Pointer,
	namespace js.Ref)

//go:wasmimport plat/js/web try_Node_LookupPrefix
//go:noescape
func TryNodeLookupPrefix(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespace js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_LookupNamespaceURI
//go:noescape
func HasNodeLookupNamespaceURI(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_LookupNamespaceURI
//go:noescape
func NodeLookupNamespaceURIFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_LookupNamespaceURI
//go:noescape
func CallNodeLookupNamespaceURI(
	this js.Ref, retPtr unsafe.Pointer,
	prefix js.Ref)

//go:wasmimport plat/js/web try_Node_LookupNamespaceURI
//go:noescape
func TryNodeLookupNamespaceURI(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	prefix js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_IsDefaultNamespace
//go:noescape
func HasNodeIsDefaultNamespace(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_IsDefaultNamespace
//go:noescape
func NodeIsDefaultNamespaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_IsDefaultNamespace
//go:noescape
func CallNodeIsDefaultNamespace(
	this js.Ref, retPtr unsafe.Pointer,
	namespace js.Ref)

//go:wasmimport plat/js/web try_Node_IsDefaultNamespace
//go:noescape
func TryNodeIsDefaultNamespace(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespace js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_InsertBefore
//go:noescape
func HasNodeInsertBefore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_InsertBefore
//go:noescape
func NodeInsertBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_InsertBefore
//go:noescape
func CallNodeInsertBefore(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref,
	child js.Ref)

//go:wasmimport plat/js/web try_Node_InsertBefore
//go:noescape
func TryNodeInsertBefore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref,
	child js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_AppendChild
//go:noescape
func HasNodeAppendChild(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_AppendChild
//go:noescape
func NodeAppendChildFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_AppendChild
//go:noescape
func CallNodeAppendChild(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref)

//go:wasmimport plat/js/web try_Node_AppendChild
//go:noescape
func TryNodeAppendChild(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_ReplaceChild
//go:noescape
func HasNodeReplaceChild(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_ReplaceChild
//go:noescape
func NodeReplaceChildFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_ReplaceChild
//go:noescape
func CallNodeReplaceChild(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref,
	child js.Ref)

//go:wasmimport plat/js/web try_Node_ReplaceChild
//go:noescape
func TryNodeReplaceChild(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref,
	child js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Node_RemoveChild
//go:noescape
func HasNodeRemoveChild(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Node_RemoveChild
//go:noescape
func NodeRemoveChildFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_RemoveChild
//go:noescape
func CallNodeRemoveChild(
	this js.Ref, retPtr unsafe.Pointer,
	child js.Ref)

//go:wasmimport plat/js/web try_Node_RemoveChild
//go:noescape
func TryNodeRemoveChild(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	child js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_AbstractRange_StartContainer
//go:noescape
func GetAbstractRangeStartContainer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AbstractRange_StartOffset
//go:noescape
func GetAbstractRangeStartOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AbstractRange_EndContainer
//go:noescape
func GetAbstractRangeEndContainer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AbstractRange_EndOffset
//go:noescape
func GetAbstractRangeEndOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_AbstractRange_Collapsed
//go:noescape
func GetAbstractRangeCollapsed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_AccelerometerLocalCoordinateSystem
//go:noescape
func ConstOfAccelerometerLocalCoordinateSystem(str js.Ref) uint32

//go:wasmimport plat/js/web store_AccelerometerSensorOptions
//go:noescape
func AccelerometerSensorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AccelerometerSensorOptions
//go:noescape
func AccelerometerSensorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Accelerometer_Accelerometer
//go:noescape
func NewAccelerometerByAccelerometer(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Accelerometer_Accelerometer1
//go:noescape
func NewAccelerometerByAccelerometer1() js.Ref

//go:wasmimport plat/js/web get_Accelerometer_X
//go:noescape
func GetAccelerometerX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Accelerometer_Y
//go:noescape
func GetAccelerometerY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Accelerometer_Z
//go:noescape
func GetAccelerometerZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_AccelerometerReadingValues
//go:noescape
func AccelerometerReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AccelerometerReadingValues
//go:noescape
func AccelerometerReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AdRender
//go:noescape
func AdRenderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AdRender
//go:noescape
func AdRenderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AddressErrors
//go:noescape
func AddressErrorsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AddressErrors
//go:noescape
func AddressErrorsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AddressInit
//go:noescape
func AddressInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AddressInit
//go:noescape
func AddressInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AesCbcParams
//go:noescape
func AesCbcParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AesCbcParams
//go:noescape
func AesCbcParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AesCtrParams
//go:noescape
func AesCtrParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AesCtrParams
//go:noescape
func AesCtrParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AesDerivedKeyParams
//go:noescape
func AesDerivedKeyParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AesDerivedKeyParams
//go:noescape
func AesDerivedKeyParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AesGcmParams
//go:noescape
func AesGcmParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AesGcmParams
//go:noescape
func AesGcmParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AesKeyAlgorithm
//go:noescape
func AesKeyAlgorithmJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AesKeyAlgorithm
//go:noescape
func AesKeyAlgorithmJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AesKeyGenParams
//go:noescape
func AesKeyGenParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AesKeyGenParams
//go:noescape
func AesKeyGenParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_Algorithm
//go:noescape
func AlgorithmJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_Algorithm
//go:noescape
func AlgorithmJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_AlignSetting
//go:noescape
func ConstOfAlignSetting(str js.Ref) uint32

//go:wasmimport plat/js/web store_AllowedBluetoothDevice
//go:noescape
func AllowedBluetoothDeviceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AllowedBluetoothDevice
//go:noescape
func AllowedBluetoothDeviceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_AllowedUSBDevice
//go:noescape
func AllowedUSBDeviceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AllowedUSBDevice
//go:noescape
func AllowedUSBDeviceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_AlphaOption
//go:noescape
func ConstOfAlphaOption(str js.Ref) uint32

//go:wasmimport plat/js/web store_AmbientLightReadingValues
//go:noescape
func AmbientLightReadingValuesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AmbientLightReadingValues
//go:noescape
func AmbientLightReadingValuesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_SensorOptions
//go:noescape
func SensorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SensorOptions
//go:noescape
func SensorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_AmbientLightSensor_AmbientLightSensor
//go:noescape
func NewAmbientLightSensorByAmbientLightSensor(
	sensorOptions unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_AmbientLightSensor_AmbientLightSensor1
//go:noescape
func NewAmbientLightSensorByAmbientLightSensor1() js.Ref

//go:wasmimport plat/js/web get_AmbientLightSensor_Illuminance
//go:noescape
func GetAmbientLightSensorIlluminance(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)
