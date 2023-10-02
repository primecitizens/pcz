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
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_FontFacePalette_UsableWithLightBackground
//go:noescape

func GetFontFacePaletteUsableWithLightBackground(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontFacePalette_UsableWithDarkBackground
//go:noescape

func GetFontFacePaletteUsableWithDarkBackground(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_FontFacePalette_Get
//go:noescape

func CallFontFacePaletteGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_FontFacePalette_Get
//go:noescape

func FontFacePaletteGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_FontFacePalettes_Length
//go:noescape

func GetFontFacePalettesLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_FontFacePalettes_Get
//go:noescape

func CallFontFacePalettesGet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_FontFacePalettes_Get
//go:noescape

func FontFacePalettesGetFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_Family
//go:noescape

func SetFontFaceFamily(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Style
//go:noescape

func GetFontFaceStyle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_Style
//go:noescape

func SetFontFaceStyle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Weight
//go:noescape

func GetFontFaceWeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_Weight
//go:noescape

func SetFontFaceWeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Stretch
//go:noescape

func GetFontFaceStretch(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_Stretch
//go:noescape

func SetFontFaceStretch(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_UnicodeRange
//go:noescape

func GetFontFaceUnicodeRange(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_UnicodeRange
//go:noescape

func SetFontFaceUnicodeRange(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Variant
//go:noescape

func GetFontFaceVariant(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_Variant
//go:noescape

func SetFontFaceVariant(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_FeatureSettings
//go:noescape

func GetFontFaceFeatureSettings(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_FeatureSettings
//go:noescape

func SetFontFaceFeatureSettings(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_VariationSettings
//go:noescape

func GetFontFaceVariationSettings(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_VariationSettings
//go:noescape

func SetFontFaceVariationSettings(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Display
//go:noescape

func GetFontFaceDisplay(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_Display
//go:noescape

func SetFontFaceDisplay(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_AscentOverride
//go:noescape

func GetFontFaceAscentOverride(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_AscentOverride
//go:noescape

func SetFontFaceAscentOverride(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_DescentOverride
//go:noescape

func GetFontFaceDescentOverride(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_DescentOverride
//go:noescape

func SetFontFaceDescentOverride(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_LineGapOverride
//go:noescape

func GetFontFaceLineGapOverride(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_FontFace_LineGapOverride
//go:noescape

func SetFontFaceLineGapOverride(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Status
//go:noescape

func GetFontFaceStatus(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_FontFace_Loaded
//go:noescape

func GetFontFaceLoaded(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Features
//go:noescape

func GetFontFaceFeatures(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Variations
//go:noescape

func GetFontFaceVariations(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontFace_Palettes
//go:noescape

func GetFontFacePalettes(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_FontFace_Load
//go:noescape

func CallFontFaceLoad(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_FontFace_Load
//go:noescape

func FontFaceLoadFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontFaceSet_Status
//go:noescape

func GetFontFaceSetStatus(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_FontFaceSet_Add
//go:noescape

func CallFontFaceSetAdd(
	this js.Ref,
	ptr unsafe.Pointer,

	font js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Add
//go:noescape

func FontFaceSetAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Delete
//go:noescape

func CallFontFaceSetDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	font js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Delete
//go:noescape

func FontFaceSetDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Clear
//go:noescape

func CallFontFaceSetClear(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Clear
//go:noescape

func FontFaceSetClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Load
//go:noescape

func CallFontFaceSetLoad(
	this js.Ref,
	ptr unsafe.Pointer,

	font js.Ref,
	text js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Load
//go:noescape

func FontFaceSetLoadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Load1
//go:noescape

func CallFontFaceSetLoad1(
	this js.Ref,
	ptr unsafe.Pointer,

	font js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Load1
//go:noescape

func FontFaceSetLoad1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Check
//go:noescape

func CallFontFaceSetCheck(
	this js.Ref,
	ptr unsafe.Pointer,

	font js.Ref,
	text js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Check
//go:noescape

func FontFaceSetCheckFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_FontFaceSet_Check1
//go:noescape

func CallFontFaceSetCheck1(
	this js.Ref,
	ptr unsafe.Pointer,

	font js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_FontFaceSet_Check1
//go:noescape

func FontFaceSetCheck1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_Document_Implementation
//go:noescape

func GetDocumentImplementation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_URL
//go:noescape

func GetDocumentURL(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_DocumentURI
//go:noescape

func GetDocumentDocumentURI(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_CompatMode
//go:noescape

func GetDocumentCompatMode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_CharacterSet
//go:noescape

func GetDocumentCharacterSet(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Charset
//go:noescape

func GetDocumentCharset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_InputEncoding
//go:noescape

func GetDocumentInputEncoding(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_ContentType
//go:noescape

func GetDocumentContentType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Doctype
//go:noescape

func GetDocumentDoctype(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_DocumentElement
//go:noescape

func GetDocumentDocumentElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_FragmentDirective
//go:noescape

func GetDocumentFragmentDirective(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_PermissionsPolicy
//go:noescape

func GetDocumentPermissionsPolicy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_WasDiscarded
//go:noescape

func GetDocumentWasDiscarded(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_FullscreenEnabled
//go:noescape

func GetDocumentFullscreenEnabled(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Fullscreen
//go:noescape

func GetDocumentFullscreen(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Timeline
//go:noescape

func GetDocumentTimeline(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_PictureInPictureEnabled
//go:noescape

func GetDocumentPictureInPictureEnabled(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_NamedFlows
//go:noescape

func GetDocumentNamedFlows(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_ScrollingElement
//go:noescape

func GetDocumentScrollingElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_RootElement
//go:noescape

func GetDocumentRootElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Prerendering
//go:noescape

func GetDocumentPrerendering(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_FgColor
//go:noescape

func GetDocumentFgColor(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_FgColor
//go:noescape

func SetDocumentFgColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_LinkColor
//go:noescape

func GetDocumentLinkColor(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_LinkColor
//go:noescape

func SetDocumentLinkColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_VlinkColor
//go:noescape

func GetDocumentVlinkColor(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_VlinkColor
//go:noescape

func SetDocumentVlinkColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_AlinkColor
//go:noescape

func GetDocumentAlinkColor(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_AlinkColor
//go:noescape

func SetDocumentAlinkColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_BgColor
//go:noescape

func GetDocumentBgColor(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_BgColor
//go:noescape

func SetDocumentBgColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Anchors
//go:noescape

func GetDocumentAnchors(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Applets
//go:noescape

func GetDocumentApplets(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_All
//go:noescape

func GetDocumentAll(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Location
//go:noescape

func GetDocumentLocation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Domain
//go:noescape

func GetDocumentDomain(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_Domain
//go:noescape

func SetDocumentDomain(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Referrer
//go:noescape

func GetDocumentReferrer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Cookie
//go:noescape

func GetDocumentCookie(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_Cookie
//go:noescape

func SetDocumentCookie(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_LastModified
//go:noescape

func GetDocumentLastModified(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_ReadyState
//go:noescape

func GetDocumentReadyState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Document_Title
//go:noescape

func GetDocumentTitle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_Title
//go:noescape

func SetDocumentTitle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Dir
//go:noescape

func GetDocumentDir(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_Dir
//go:noescape

func SetDocumentDir(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Body
//go:noescape

func GetDocumentBody(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_Body
//go:noescape

func SetDocumentBody(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Head
//go:noescape

func GetDocumentHead(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Images
//go:noescape

func GetDocumentImages(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Embeds
//go:noescape

func GetDocumentEmbeds(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Plugins
//go:noescape

func GetDocumentPlugins(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Links
//go:noescape

func GetDocumentLinks(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Forms
//go:noescape

func GetDocumentForms(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Scripts
//go:noescape

func GetDocumentScripts(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_CurrentScript
//go:noescape

func GetDocumentCurrentScript(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_DefaultView
//go:noescape

func GetDocumentDefaultView(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_DesignMode
//go:noescape

func GetDocumentDesignMode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_DesignMode
//go:noescape

func SetDocumentDesignMode(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_Hidden
//go:noescape

func GetDocumentHidden(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_VisibilityState
//go:noescape

func GetDocumentVisibilityState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Document_FullscreenElement
//go:noescape

func GetDocumentFullscreenElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_StyleSheets
//go:noescape

func GetDocumentStyleSheets(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_AdoptedStyleSheets
//go:noescape

func GetDocumentAdoptedStyleSheets(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Document_AdoptedStyleSheets
//go:noescape

func SetDocumentAdoptedStyleSheets(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Document_PictureInPictureElement
//go:noescape

func GetDocumentPictureInPictureElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_ActiveElement
//go:noescape

func GetDocumentActiveElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_PointerLockElement
//go:noescape

func GetDocumentPointerLockElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_Children
//go:noescape

func GetDocumentChildren(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_FirstElementChild
//go:noescape

func GetDocumentFirstElementChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_LastElementChild
//go:noescape

func GetDocumentLastElementChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Document_ChildElementCount
//go:noescape

func GetDocumentChildElementCount(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Document_Fonts
//go:noescape

func GetDocumentFonts(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementsByTagName
//go:noescape

func CallDocumentGetElementsByTagName(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementsByTagName
//go:noescape

func DocumentGetElementsByTagNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementsByTagNameNS
//go:noescape

func CallDocumentGetElementsByTagNameNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementsByTagNameNS
//go:noescape

func DocumentGetElementsByTagNameNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementsByClassName
//go:noescape

func CallDocumentGetElementsByClassName(
	this js.Ref,
	ptr unsafe.Pointer,

	classNames js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementsByClassName
//go:noescape

func DocumentGetElementsByClassNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateElement
//go:noescape

func CallDocumentCreateElement(
	this js.Ref,
	ptr unsafe.Pointer,

	localName js.Ref,
	options js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateElement
//go:noescape

func DocumentCreateElementFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateElement1
//go:noescape

func CallDocumentCreateElement1(
	this js.Ref,
	ptr unsafe.Pointer,

	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateElement1
//go:noescape

func DocumentCreateElement1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateElementNS
//go:noescape

func CallDocumentCreateElementNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	qualifiedName js.Ref,
	options js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateElementNS
//go:noescape

func DocumentCreateElementNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateElementNS1
//go:noescape

func CallDocumentCreateElementNS1(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateElementNS1
//go:noescape

func DocumentCreateElementNS1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateDocumentFragment
//go:noescape

func CallDocumentCreateDocumentFragment(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_CreateDocumentFragment
//go:noescape

func DocumentCreateDocumentFragmentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateTextNode
//go:noescape

func CallDocumentCreateTextNode(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateTextNode
//go:noescape

func DocumentCreateTextNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateCDATASection
//go:noescape

func CallDocumentCreateCDATASection(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateCDATASection
//go:noescape

func DocumentCreateCDATASectionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateComment
//go:noescape

func CallDocumentCreateComment(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateComment
//go:noescape

func DocumentCreateCommentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateProcessingInstruction
//go:noescape

func CallDocumentCreateProcessingInstruction(
	this js.Ref,
	ptr unsafe.Pointer,

	target js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateProcessingInstruction
//go:noescape

func DocumentCreateProcessingInstructionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ImportNode
//go:noescape

func CallDocumentImportNode(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	deep js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_ImportNode
//go:noescape

func DocumentImportNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ImportNode1
//go:noescape

func CallDocumentImportNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_ImportNode1
//go:noescape

func DocumentImportNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_AdoptNode
//go:noescape

func CallDocumentAdoptNode(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_AdoptNode
//go:noescape

func DocumentAdoptNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateAttribute
//go:noescape

func CallDocumentCreateAttribute(
	this js.Ref,
	ptr unsafe.Pointer,

	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateAttribute
//go:noescape

func DocumentCreateAttributeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateAttributeNS
//go:noescape

func CallDocumentCreateAttributeNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateAttributeNS
//go:noescape

func DocumentCreateAttributeNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateEvent
//go:noescape

func CallDocumentCreateEvent(
	this js.Ref,
	ptr unsafe.Pointer,

	iface js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateEvent
//go:noescape

func DocumentCreateEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateRange
//go:noescape

func CallDocumentCreateRange(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_CreateRange
//go:noescape

func DocumentCreateRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateNodeIterator
//go:noescape

func CallDocumentCreateNodeIterator(
	this js.Ref,
	ptr unsafe.Pointer,

	root js.Ref,
	whatToShow uint32,
	filter js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateNodeIterator
//go:noescape

func DocumentCreateNodeIteratorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateNodeIterator1
//go:noescape

func CallDocumentCreateNodeIterator1(
	this js.Ref,
	ptr unsafe.Pointer,

	root js.Ref,
	whatToShow uint32,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateNodeIterator1
//go:noescape

func DocumentCreateNodeIterator1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateNodeIterator2
//go:noescape

func CallDocumentCreateNodeIterator2(
	this js.Ref,
	ptr unsafe.Pointer,

	root js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateNodeIterator2
//go:noescape

func DocumentCreateNodeIterator2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateTreeWalker
//go:noescape

func CallDocumentCreateTreeWalker(
	this js.Ref,
	ptr unsafe.Pointer,

	root js.Ref,
	whatToShow uint32,
	filter js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateTreeWalker
//go:noescape

func DocumentCreateTreeWalkerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateTreeWalker1
//go:noescape

func CallDocumentCreateTreeWalker1(
	this js.Ref,
	ptr unsafe.Pointer,

	root js.Ref,
	whatToShow uint32,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateTreeWalker1
//go:noescape

func DocumentCreateTreeWalker1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateTreeWalker2
//go:noescape

func CallDocumentCreateTreeWalker2(
	this js.Ref,
	ptr unsafe.Pointer,

	root js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateTreeWalker2
//go:noescape

func DocumentCreateTreeWalker2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_StartViewTransition
//go:noescape

func CallDocumentStartViewTransition(
	this js.Ref,
	ptr unsafe.Pointer,

	updateCallback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_StartViewTransition
//go:noescape

func DocumentStartViewTransitionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_StartViewTransition1
//go:noescape

func CallDocumentStartViewTransition1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_StartViewTransition1
//go:noescape

func DocumentStartViewTransition1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_MeasureElement
//go:noescape

func CallDocumentMeasureElement(
	this js.Ref,
	ptr unsafe.Pointer,

	element js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_MeasureElement
//go:noescape

func DocumentMeasureElementFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_MeasureText
//go:noescape

func CallDocumentMeasureText(
	this js.Ref,
	ptr unsafe.Pointer,

	text js.Ref,
	styleMap js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_MeasureText
//go:noescape

func DocumentMeasureTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetSelection
//go:noescape

func CallDocumentGetSelection(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_GetSelection
//go:noescape

func DocumentGetSelectionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExitPointerLock
//go:noescape

func CallDocumentExitPointerLock(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_ExitPointerLock
//go:noescape

func DocumentExitPointerLockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExitFullscreen
//go:noescape

func CallDocumentExitFullscreen(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_ExitFullscreen
//go:noescape

func DocumentExitFullscreenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_HasStorageAccess
//go:noescape

func CallDocumentHasStorageAccess(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_HasStorageAccess
//go:noescape

func DocumentHasStorageAccessFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_RequestStorageAccess
//go:noescape

func CallDocumentRequestStorageAccess(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_RequestStorageAccess
//go:noescape

func DocumentRequestStorageAccessFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExitPictureInPicture
//go:noescape

func CallDocumentExitPictureInPicture(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_ExitPictureInPicture
//go:noescape

func DocumentExitPictureInPictureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_HasPrivateTokens
//go:noescape

func CallDocumentHasPrivateTokens(
	this js.Ref,
	ptr unsafe.Pointer,

	issuer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_HasPrivateTokens
//go:noescape

func DocumentHasPrivateTokensFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_HasRedemptionRecord
//go:noescape

func CallDocumentHasRedemptionRecord(
	this js.Ref,
	ptr unsafe.Pointer,

	issuer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_HasRedemptionRecord
//go:noescape

func DocumentHasRedemptionRecordFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ElementFromPoint
//go:noescape

func CallDocumentElementFromPoint(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_Document_ElementFromPoint
//go:noescape

func DocumentElementFromPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ElementsFromPoint
//go:noescape

func CallDocumentElementsFromPoint(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_Document_ElementsFromPoint
//go:noescape

func DocumentElementsFromPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CaretPositionFromPoint
//go:noescape

func CallDocumentCaretPositionFromPoint(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_Document_CaretPositionFromPoint
//go:noescape

func DocumentCaretPositionFromPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Clear
//go:noescape

func CallDocumentClear(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_Clear
//go:noescape

func DocumentClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CaptureEvents
//go:noescape

func CallDocumentCaptureEvents(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_CaptureEvents
//go:noescape

func DocumentCaptureEventsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ReleaseEvents
//go:noescape

func CallDocumentReleaseEvents(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_ReleaseEvents
//go:noescape

func DocumentReleaseEventsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_RequestStorageAccessFor
//go:noescape

func CallDocumentRequestStorageAccessFor(
	this js.Ref,
	ptr unsafe.Pointer,

	requestedOrigin js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_RequestStorageAccessFor
//go:noescape

func DocumentRequestStorageAccessForFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Get
//go:noescape

func CallDocumentGet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_Get
//go:noescape

func DocumentGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementsByName
//go:noescape

func CallDocumentGetElementsByName(
	this js.Ref,
	ptr unsafe.Pointer,

	elementName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementsByName
//go:noescape

func DocumentGetElementsByNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Open
//go:noescape

func CallDocumentOpen(
	this js.Ref,
	ptr unsafe.Pointer,

	unused1 js.Ref,
	unused2 js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_Open
//go:noescape

func DocumentOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Open1
//go:noescape

func CallDocumentOpen1(
	this js.Ref,
	ptr unsafe.Pointer,

	unused1 js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_Open1
//go:noescape

func DocumentOpen1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Open2
//go:noescape

func CallDocumentOpen2(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_Open2
//go:noescape

func DocumentOpen2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Open3
//go:noescape

func CallDocumentOpen3(
	this js.Ref,
	ptr unsafe.Pointer,

	url js.Ref,
	name js.Ref,
	features js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_Open3
//go:noescape

func DocumentOpen3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Close
//go:noescape

func CallDocumentClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_Close
//go:noescape

func DocumentCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Write
//go:noescape

func CallDocumentWrite(
	this js.Ref,
	ptr unsafe.Pointer,

	textPtr unsafe.Pointer,
	textCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Document_Write
//go:noescape

func DocumentWriteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Writeln
//go:noescape

func CallDocumentWriteln(
	this js.Ref,
	ptr unsafe.Pointer,

	textPtr unsafe.Pointer,
	textCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Document_Writeln
//go:noescape

func DocumentWritelnFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_HasFocus
//go:noescape

func CallDocumentHasFocus(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_HasFocus
//go:noescape

func DocumentHasFocusFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExecCommand
//go:noescape

func CallDocumentExecCommand(
	this js.Ref,
	ptr unsafe.Pointer,

	commandId js.Ref,
	showUI js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_ExecCommand
//go:noescape

func DocumentExecCommandFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExecCommand1
//go:noescape

func CallDocumentExecCommand1(
	this js.Ref,
	ptr unsafe.Pointer,

	commandId js.Ref,
	showUI js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_ExecCommand1
//go:noescape

func DocumentExecCommand1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ExecCommand2
//go:noescape

func CallDocumentExecCommand2(
	this js.Ref,
	ptr unsafe.Pointer,

	commandId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_ExecCommand2
//go:noescape

func DocumentExecCommand2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandEnabled
//go:noescape

func CallDocumentQueryCommandEnabled(
	this js.Ref,
	ptr unsafe.Pointer,

	commandId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandEnabled
//go:noescape

func DocumentQueryCommandEnabledFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandIndeterm
//go:noescape

func CallDocumentQueryCommandIndeterm(
	this js.Ref,
	ptr unsafe.Pointer,

	commandId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandIndeterm
//go:noescape

func DocumentQueryCommandIndetermFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandState
//go:noescape

func CallDocumentQueryCommandState(
	this js.Ref,
	ptr unsafe.Pointer,

	commandId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandState
//go:noescape

func DocumentQueryCommandStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandSupported
//go:noescape

func CallDocumentQueryCommandSupported(
	this js.Ref,
	ptr unsafe.Pointer,

	commandId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandSupported
//go:noescape

func DocumentQueryCommandSupportedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QueryCommandValue
//go:noescape

func CallDocumentQueryCommandValue(
	this js.Ref,
	ptr unsafe.Pointer,

	commandId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_QueryCommandValue
//go:noescape

func DocumentQueryCommandValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetBoxQuads
//go:noescape

func CallDocumentGetBoxQuads(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Document_GetBoxQuads
//go:noescape

func DocumentGetBoxQuadsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetBoxQuads1
//go:noescape

func CallDocumentGetBoxQuads1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_GetBoxQuads1
//go:noescape

func DocumentGetBoxQuads1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertQuadFromNode
//go:noescape

func CallDocumentConvertQuadFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	quad unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertQuadFromNode
//go:noescape

func DocumentConvertQuadFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertQuadFromNode1
//go:noescape

func CallDocumentConvertQuadFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	quad unsafe.Pointer,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertQuadFromNode1
//go:noescape

func DocumentConvertQuadFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertRectFromNode
//go:noescape

func CallDocumentConvertRectFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertRectFromNode
//go:noescape

func DocumentConvertRectFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertRectFromNode1
//go:noescape

func CallDocumentConvertRectFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertRectFromNode1
//go:noescape

func DocumentConvertRectFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertPointFromNode
//go:noescape

func CallDocumentConvertPointFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	point unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertPointFromNode
//go:noescape

func DocumentConvertPointFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ConvertPointFromNode1
//go:noescape

func CallDocumentConvertPointFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	point unsafe.Pointer,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_ConvertPointFromNode1
//go:noescape

func DocumentConvertPointFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetElementById
//go:noescape

func CallDocumentGetElementById(
	this js.Ref,
	ptr unsafe.Pointer,

	elementId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_GetElementById
//go:noescape

func DocumentGetElementByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_GetAnimations
//go:noescape

func CallDocumentGetAnimations(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Document_GetAnimations
//go:noescape

func DocumentGetAnimationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Prepend
//go:noescape

func CallDocumentPrepend(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Document_Prepend
//go:noescape

func DocumentPrependFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Append
//go:noescape

func CallDocumentAppend(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Document_Append
//go:noescape

func DocumentAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_ReplaceChildren
//go:noescape

func CallDocumentReplaceChildren(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Document_ReplaceChildren
//go:noescape

func DocumentReplaceChildrenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QuerySelector
//go:noescape

func CallDocumentQuerySelector(
	this js.Ref,
	ptr unsafe.Pointer,

	selectors js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_QuerySelector
//go:noescape

func DocumentQuerySelectorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_QuerySelectorAll
//go:noescape

func CallDocumentQuerySelectorAll(
	this js.Ref,
	ptr unsafe.Pointer,

	selectors js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_QuerySelectorAll
//go:noescape

func DocumentQuerySelectorAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateExpression
//go:noescape

func CallDocumentCreateExpression(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	resolver js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateExpression
//go:noescape

func DocumentCreateExpressionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateExpression1
//go:noescape

func CallDocumentCreateExpression1(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateExpression1
//go:noescape

func DocumentCreateExpression1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_CreateNSResolver
//go:noescape

func CallDocumentCreateNSResolver(
	this js.Ref,
	ptr unsafe.Pointer,

	nodeResolver js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_CreateNSResolver
//go:noescape

func DocumentCreateNSResolverFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Evaluate
//go:noescape

func CallDocumentEvaluate(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32,
	result js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_Evaluate
//go:noescape

func DocumentEvaluateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Evaluate1
//go:noescape

func CallDocumentEvaluate1(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
	typ uint32,
) js.Ref

//go:wasmimport plat/js/web func_Document_Evaluate1
//go:noescape

func DocumentEvaluate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Evaluate2
//go:noescape

func CallDocumentEvaluate2(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	contextNode js.Ref,
	resolver js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_Evaluate2
//go:noescape

func DocumentEvaluate2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Document_Evaluate3
//go:noescape

func CallDocumentEvaluate3(
	this js.Ref,
	ptr unsafe.Pointer,

	expression js.Ref,
	contextNode js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Document_Evaluate3
//go:noescape

func DocumentEvaluate3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_Node_NodeType
//go:noescape

func GetNodeNodeType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Node_NodeName
//go:noescape

func GetNodeNodeName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_BaseURI
//go:noescape

func GetNodeBaseURI(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_IsConnected
//go:noescape

func GetNodeIsConnected(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_OwnerDocument
//go:noescape

func GetNodeOwnerDocument(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_ParentNode
//go:noescape

func GetNodeParentNode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_ParentElement
//go:noescape

func GetNodeParentElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_ChildNodes
//go:noescape

func GetNodeChildNodes(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_FirstChild
//go:noescape

func GetNodeFirstChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_LastChild
//go:noescape

func GetNodeLastChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_PreviousSibling
//go:noescape

func GetNodePreviousSibling(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_NextSibling
//go:noescape

func GetNodeNextSibling(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Node_NodeValue
//go:noescape

func GetNodeNodeValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Node_NodeValue
//go:noescape

func SetNodeNodeValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Node_TextContent
//go:noescape

func GetNodeTextContent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Node_TextContent
//go:noescape

func SetNodeTextContent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_Node_GetRootNode
//go:noescape

func CallNodeGetRootNode(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Node_GetRootNode
//go:noescape

func NodeGetRootNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_GetRootNode1
//go:noescape

func CallNodeGetRootNode1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Node_GetRootNode1
//go:noescape

func NodeGetRootNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_HasChildNodes
//go:noescape

func CallNodeHasChildNodes(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Node_HasChildNodes
//go:noescape

func NodeHasChildNodesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_Normalize
//go:noescape

func CallNodeNormalize(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Node_Normalize
//go:noescape

func NodeNormalizeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_CloneNode
//go:noescape

func CallNodeCloneNode(
	this js.Ref,
	ptr unsafe.Pointer,

	deep js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_CloneNode
//go:noescape

func NodeCloneNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_CloneNode1
//go:noescape

func CallNodeCloneNode1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Node_CloneNode1
//go:noescape

func NodeCloneNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_IsEqualNode
//go:noescape

func CallNodeIsEqualNode(
	this js.Ref,
	ptr unsafe.Pointer,

	otherNode js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_IsEqualNode
//go:noescape

func NodeIsEqualNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_IsSameNode
//go:noescape

func CallNodeIsSameNode(
	this js.Ref,
	ptr unsafe.Pointer,

	otherNode js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_IsSameNode
//go:noescape

func NodeIsSameNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_CompareDocumentPosition
//go:noescape

func CallNodeCompareDocumentPosition(
	this js.Ref,
	ptr unsafe.Pointer,

	other js.Ref,
) uint32

//go:wasmimport plat/js/web func_Node_CompareDocumentPosition
//go:noescape

func NodeCompareDocumentPositionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_Contains
//go:noescape

func CallNodeContains(
	this js.Ref,
	ptr unsafe.Pointer,

	other js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_Contains
//go:noescape

func NodeContainsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_LookupPrefix
//go:noescape

func CallNodeLookupPrefix(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_LookupPrefix
//go:noescape

func NodeLookupPrefixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_LookupNamespaceURI
//go:noescape

func CallNodeLookupNamespaceURI(
	this js.Ref,
	ptr unsafe.Pointer,

	prefix js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_LookupNamespaceURI
//go:noescape

func NodeLookupNamespaceURIFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_IsDefaultNamespace
//go:noescape

func CallNodeIsDefaultNamespace(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_IsDefaultNamespace
//go:noescape

func NodeIsDefaultNamespaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_InsertBefore
//go:noescape

func CallNodeInsertBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	child js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_InsertBefore
//go:noescape

func NodeInsertBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_AppendChild
//go:noescape

func CallNodeAppendChild(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_AppendChild
//go:noescape

func NodeAppendChildFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_ReplaceChild
//go:noescape

func CallNodeReplaceChild(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	child js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_ReplaceChild
//go:noescape

func NodeReplaceChildFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Node_RemoveChild
//go:noescape

func CallNodeRemoveChild(
	this js.Ref,
	ptr unsafe.Pointer,

	child js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Node_RemoveChild
//go:noescape

func NodeRemoveChildFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_AbstractRange_StartContainer
//go:noescape

func GetAbstractRangeStartContainer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AbstractRange_StartOffset
//go:noescape

func GetAbstractRangeStartOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_AbstractRange_EndContainer
//go:noescape

func GetAbstractRangeEndContainer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_AbstractRange_EndOffset
//go:noescape

func GetAbstractRangeEndOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_AbstractRange_Collapsed
//go:noescape

func GetAbstractRangeCollapsed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_Accelerometer_Y
//go:noescape

func GetAccelerometerY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_Accelerometer_Z
//go:noescape

func GetAccelerometerZ(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

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
