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

//go:wasmimport plat/js/web constof_RTCIceGathererState
//go:noescape
func ConstOfRTCIceGathererState(str js.Ref) uint32

//go:wasmimport plat/js/web get_RTCIceTransport_Role
//go:noescape
func GetRTCIceTransportRole(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCIceTransport_Component
//go:noescape
func GetRTCIceTransportComponent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCIceTransport_State
//go:noescape
func GetRTCIceTransportState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCIceTransport_GatheringState
//go:noescape
func GetRTCIceTransportGatheringState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_GetLocalCandidates
//go:noescape
func HasFuncRTCIceTransportGetLocalCandidates(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetLocalCandidates
//go:noescape
func FuncRTCIceTransportGetLocalCandidates(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_GetLocalCandidates
//go:noescape
func CallRTCIceTransportGetLocalCandidates(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_GetLocalCandidates
//go:noescape
func TryRTCIceTransportGetLocalCandidates(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_GetRemoteCandidates
//go:noescape
func HasFuncRTCIceTransportGetRemoteCandidates(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetRemoteCandidates
//go:noescape
func FuncRTCIceTransportGetRemoteCandidates(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_GetRemoteCandidates
//go:noescape
func CallRTCIceTransportGetRemoteCandidates(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_GetRemoteCandidates
//go:noescape
func TryRTCIceTransportGetRemoteCandidates(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_GetSelectedCandidatePair
//go:noescape
func HasFuncRTCIceTransportGetSelectedCandidatePair(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetSelectedCandidatePair
//go:noescape
func FuncRTCIceTransportGetSelectedCandidatePair(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_GetSelectedCandidatePair
//go:noescape
func CallRTCIceTransportGetSelectedCandidatePair(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_GetSelectedCandidatePair
//go:noescape
func TryRTCIceTransportGetSelectedCandidatePair(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_GetLocalParameters
//go:noescape
func HasFuncRTCIceTransportGetLocalParameters(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetLocalParameters
//go:noescape
func FuncRTCIceTransportGetLocalParameters(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_GetLocalParameters
//go:noescape
func CallRTCIceTransportGetLocalParameters(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_GetLocalParameters
//go:noescape
func TryRTCIceTransportGetLocalParameters(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_GetRemoteParameters
//go:noescape
func HasFuncRTCIceTransportGetRemoteParameters(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetRemoteParameters
//go:noescape
func FuncRTCIceTransportGetRemoteParameters(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_GetRemoteParameters
//go:noescape
func CallRTCIceTransportGetRemoteParameters(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_GetRemoteParameters
//go:noescape
func TryRTCIceTransportGetRemoteParameters(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_Gather
//go:noescape
func HasFuncRTCIceTransportGather(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Gather
//go:noescape
func FuncRTCIceTransportGather(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_Gather
//go:noescape
func CallRTCIceTransportGather(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_Gather
//go:noescape
func TryRTCIceTransportGather(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_Gather1
//go:noescape
func HasFuncRTCIceTransportGather1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Gather1
//go:noescape
func FuncRTCIceTransportGather1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_Gather1
//go:noescape
func CallRTCIceTransportGather1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_Gather1
//go:noescape
func TryRTCIceTransportGather1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_Start
//go:noescape
func HasFuncRTCIceTransportStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Start
//go:noescape
func FuncRTCIceTransportStart(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_Start
//go:noescape
func CallRTCIceTransportStart(
	this js.Ref, retPtr unsafe.Pointer,
	remoteParameters unsafe.Pointer,
	role uint32)

//go:wasmimport plat/js/web try_RTCIceTransport_Start
//go:noescape
func TryRTCIceTransportStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	remoteParameters unsafe.Pointer,
	role uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_Start1
//go:noescape
func HasFuncRTCIceTransportStart1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Start1
//go:noescape
func FuncRTCIceTransportStart1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_Start1
//go:noescape
func CallRTCIceTransportStart1(
	this js.Ref, retPtr unsafe.Pointer,
	remoteParameters unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_Start1
//go:noescape
func TryRTCIceTransportStart1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	remoteParameters unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_Start2
//go:noescape
func HasFuncRTCIceTransportStart2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Start2
//go:noescape
func FuncRTCIceTransportStart2(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_Start2
//go:noescape
func CallRTCIceTransportStart2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_Start2
//go:noescape
func TryRTCIceTransportStart2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_Stop
//go:noescape
func HasFuncRTCIceTransportStop(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Stop
//go:noescape
func FuncRTCIceTransportStop(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_Stop
//go:noescape
func CallRTCIceTransportStop(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_Stop
//go:noescape
func TryRTCIceTransportStop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_AddRemoteCandidate
//go:noescape
func HasFuncRTCIceTransportAddRemoteCandidate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_AddRemoteCandidate
//go:noescape
func FuncRTCIceTransportAddRemoteCandidate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_AddRemoteCandidate
//go:noescape
func CallRTCIceTransportAddRemoteCandidate(
	this js.Ref, retPtr unsafe.Pointer,
	remoteCandidate unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_AddRemoteCandidate
//go:noescape
func TryRTCIceTransportAddRemoteCandidate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	remoteCandidate unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCIceTransport_AddRemoteCandidate1
//go:noescape
func HasFuncRTCIceTransportAddRemoteCandidate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_AddRemoteCandidate1
//go:noescape
func FuncRTCIceTransportAddRemoteCandidate1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIceTransport_AddRemoteCandidate1
//go:noescape
func CallRTCIceTransportAddRemoteCandidate1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIceTransport_AddRemoteCandidate1
//go:noescape
func TryRTCIceTransportAddRemoteCandidate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_RTCDtlsTransportState
//go:noescape
func ConstOfRTCDtlsTransportState(str js.Ref) uint32

//go:wasmimport plat/js/web get_RTCDtlsTransport_IceTransport
//go:noescape
func GetRTCDtlsTransportIceTransport(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCDtlsTransport_State
//go:noescape
func GetRTCDtlsTransportState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_RTCDtlsTransport_GetRemoteCertificates
//go:noescape
func HasFuncRTCDtlsTransportGetRemoteCertificates(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCDtlsTransport_GetRemoteCertificates
//go:noescape
func FuncRTCDtlsTransportGetRemoteCertificates(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCDtlsTransport_GetRemoteCertificates
//go:noescape
func CallRTCDtlsTransportGetRemoteCertificates(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCDtlsTransport_GetRemoteCertificates
//go:noescape
func TryRTCDtlsTransportGetRemoteCertificates(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_RTCEncodedAudioFrameMetadata
//go:noescape
func RTCEncodedAudioFrameMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCEncodedAudioFrameMetadata
//go:noescape
func RTCEncodedAudioFrameMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCEncodedAudioFrame_Timestamp
//go:noescape
func GetRTCEncodedAudioFrameTimestamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCEncodedAudioFrame_Data
//go:noescape
func GetRTCEncodedAudioFrameData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_RTCEncodedAudioFrame_Data
//go:noescape
func SetRTCEncodedAudioFrameData(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_RTCEncodedAudioFrame_GetMetadata
//go:noescape
func HasFuncRTCEncodedAudioFrameGetMetadata(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCEncodedAudioFrame_GetMetadata
//go:noescape
func FuncRTCEncodedAudioFrameGetMetadata(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCEncodedAudioFrame_GetMetadata
//go:noescape
func CallRTCEncodedAudioFrameGetMetadata(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCEncodedAudioFrame_GetMetadata
//go:noescape
func TryRTCEncodedAudioFrameGetMetadata(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_RTCEncodedVideoFrameMetadata
//go:noescape
func RTCEncodedVideoFrameMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCEncodedVideoFrameMetadata
//go:noescape
func RTCEncodedVideoFrameMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCEncodedVideoFrameType
//go:noescape
func ConstOfRTCEncodedVideoFrameType(str js.Ref) uint32

//go:wasmimport plat/js/web get_RTCEncodedVideoFrame_Type
//go:noescape
func GetRTCEncodedVideoFrameType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCEncodedVideoFrame_Timestamp
//go:noescape
func GetRTCEncodedVideoFrameTimestamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCEncodedVideoFrame_Data
//go:noescape
func GetRTCEncodedVideoFrameData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_RTCEncodedVideoFrame_Data
//go:noescape
func SetRTCEncodedVideoFrameData(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_RTCEncodedVideoFrame_GetMetadata
//go:noescape
func HasFuncRTCEncodedVideoFrameGetMetadata(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCEncodedVideoFrame_GetMetadata
//go:noescape
func FuncRTCEncodedVideoFrameGetMetadata(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCEncodedVideoFrame_GetMetadata
//go:noescape
func CallRTCEncodedVideoFrameGetMetadata(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCEncodedVideoFrame_GetMetadata
//go:noescape
func TryRTCEncodedVideoFrameGetMetadata(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_RTCErrorDetailType
//go:noescape
func ConstOfRTCErrorDetailType(str js.Ref) uint32

//go:wasmimport plat/js/web store_RTCErrorInit
//go:noescape
func RTCErrorInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCErrorInit
//go:noescape
func RTCErrorInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCError_RTCError
//go:noescape
func NewRTCErrorByRTCError(
	init unsafe.Pointer,
	message js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCError_RTCError1
//go:noescape
func NewRTCErrorByRTCError1(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_RTCError_ErrorDetail
//go:noescape
func GetRTCErrorErrorDetail(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCError_SdpLineNumber
//go:noescape
func GetRTCErrorSdpLineNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCError_SctpCauseCode
//go:noescape
func GetRTCErrorSctpCauseCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCError_ReceivedAlert
//go:noescape
func GetRTCErrorReceivedAlert(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCError_SentAlert
//go:noescape
func GetRTCErrorSentAlert(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCError_HttpRequestStatusCode
//go:noescape
func GetRTCErrorHttpRequestStatusCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_RTCErrorDetailTypeIdp
//go:noescape
func ConstOfRTCErrorDetailTypeIdp(str js.Ref) uint32

//go:wasmimport plat/js/web store_RTCErrorEventInit
//go:noescape
func RTCErrorEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCErrorEventInit
//go:noescape
func RTCErrorEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCErrorEvent_RTCErrorEvent
//go:noescape
func NewRTCErrorEventByRTCErrorEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_RTCErrorEvent_Error
//go:noescape
func GetRTCErrorEventError(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_RTCStatsIceCandidatePairState
//go:noescape
func ConstOfRTCStatsIceCandidatePairState(str js.Ref) uint32

//go:wasmimport plat/js/web store_RTCIceCandidatePairStats
//go:noescape
func RTCIceCandidatePairStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIceCandidatePairStats
//go:noescape
func RTCIceCandidatePairStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCIceCandidateStats
//go:noescape
func RTCIceCandidateStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIceCandidateStats
//go:noescape
func RTCIceCandidateStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCIceConnectionState
//go:noescape
func ConstOfRTCIceConnectionState(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCIceGatheringState
//go:noescape
func ConstOfRTCIceGatheringState(str js.Ref) uint32

//go:wasmimport plat/js/web new_RTCIdentityAssertion_RTCIdentityAssertion
//go:noescape
func NewRTCIdentityAssertionByRTCIdentityAssertion(
	idp js.Ref,
	name js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCIdentityAssertion_Idp
//go:noescape
func GetRTCIdentityAssertionIdp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_RTCIdentityAssertion_Idp
//go:noescape
func SetRTCIdentityAssertionIdp(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_RTCIdentityAssertion_Name
//go:noescape
func GetRTCIdentityAssertionName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_RTCIdentityAssertion_Name
//go:noescape
func SetRTCIdentityAssertionName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_RTCIdentityValidationResult
//go:noescape
func RTCIdentityValidationResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIdentityValidationResult
//go:noescape
func RTCIdentityValidationResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCIdentityProvider
//go:noescape
func RTCIdentityProviderJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIdentityProvider
//go:noescape
func RTCIdentityProviderJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_RTCIdentityProviderRegistrar_Register
//go:noescape
func HasFuncRTCIdentityProviderRegistrarRegister(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RTCIdentityProviderRegistrar_Register
//go:noescape
func FuncRTCIdentityProviderRegistrarRegister(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RTCIdentityProviderRegistrar_Register
//go:noescape
func CallRTCIdentityProviderRegistrarRegister(
	this js.Ref, retPtr unsafe.Pointer,
	idp unsafe.Pointer)

//go:wasmimport plat/js/web try_RTCIdentityProviderRegistrar_Register
//go:noescape
func TryRTCIdentityProviderRegistrarRegister(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	idp unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_RTCIdentityProviderGlobalScope_RtcIdentityProvider
//go:noescape
func GetRTCIdentityProviderGlobalScopeRtcIdentityProvider(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_RTCInboundRtpStreamStats
//go:noescape
func RTCInboundRtpStreamStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCInboundRtpStreamStats
//go:noescape
func RTCInboundRtpStreamStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCSdpType
//go:noescape
func ConstOfRTCSdpType(str js.Ref) uint32

//go:wasmimport plat/js/web store_RTCLocalSessionDescriptionInit
//go:noescape
func RTCLocalSessionDescriptionInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCLocalSessionDescriptionInit
//go:noescape
func RTCLocalSessionDescriptionInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCMediaSourceStats
//go:noescape
func RTCMediaSourceStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCMediaSourceStats
//go:noescape
func RTCMediaSourceStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCOfferAnswerOptions
//go:noescape
func RTCOfferAnswerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCOfferAnswerOptions
//go:noescape
func RTCOfferAnswerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCOfferOptions
//go:noescape
func RTCOfferOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCOfferOptions
//go:noescape
func RTCOfferOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCQualityLimitationReason
//go:noescape
func ConstOfRTCQualityLimitationReason(str js.Ref) uint32

//go:wasmimport plat/js/web store_RTCOutboundRtpStreamStats
//go:noescape
func RTCOutboundRtpStreamStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCOutboundRtpStreamStats
//go:noescape
func RTCOutboundRtpStreamStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCSessionDescriptionInit
//go:noescape
func RTCSessionDescriptionInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCSessionDescriptionInit
//go:noescape
func RTCSessionDescriptionInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpCodecCapability
//go:noescape
func RTCRtpCodecCapabilityJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpCodecCapability
//go:noescape
func RTCRtpCodecCapabilityJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpHeaderExtensionCapability
//go:noescape
func RTCRtpHeaderExtensionCapabilityJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpHeaderExtensionCapability
//go:noescape
func RTCRtpHeaderExtensionCapabilityJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpCapabilities
//go:noescape
func RTCRtpCapabilitiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpCapabilities
//go:noescape
func RTCRtpCapabilitiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
