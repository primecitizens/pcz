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

//go:wasmimport plat/js/web constof_RTCIceGathererState
//go:noescape

func ConstOfRTCIceGathererState(str js.Ref) uint32

//go:wasmimport plat/js/web get_RTCIceTransport_Role
//go:noescape

func GetRTCIceTransportRole(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceTransport_Component
//go:noescape

func GetRTCIceTransportComponent(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceTransport_State
//go:noescape

func GetRTCIceTransportState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceTransport_GatheringState
//go:noescape

func GetRTCIceTransportGatheringState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_RTCIceTransport_GetLocalCandidates
//go:noescape

func CallRTCIceTransportGetLocalCandidates(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetLocalCandidates
//go:noescape

func RTCIceTransportGetLocalCandidatesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_GetRemoteCandidates
//go:noescape

func CallRTCIceTransportGetRemoteCandidates(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetRemoteCandidates
//go:noescape

func RTCIceTransportGetRemoteCandidatesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_GetSelectedCandidatePair
//go:noescape

func CallRTCIceTransportGetSelectedCandidatePair(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetSelectedCandidatePair
//go:noescape

func RTCIceTransportGetSelectedCandidatePairFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_GetLocalParameters
//go:noescape

func CallRTCIceTransportGetLocalParameters(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetLocalParameters
//go:noescape

func RTCIceTransportGetLocalParametersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_GetRemoteParameters
//go:noescape

func CallRTCIceTransportGetRemoteParameters(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_GetRemoteParameters
//go:noescape

func RTCIceTransportGetRemoteParametersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_Gather
//go:noescape

func CallRTCIceTransportGather(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Gather
//go:noescape

func RTCIceTransportGatherFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_Gather1
//go:noescape

func CallRTCIceTransportGather1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Gather1
//go:noescape

func RTCIceTransportGather1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_Start
//go:noescape

func CallRTCIceTransportStart(
	this js.Ref,
	ptr unsafe.Pointer,

	remoteParameters unsafe.Pointer,
	role uint32,
) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Start
//go:noescape

func RTCIceTransportStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_Start1
//go:noescape

func CallRTCIceTransportStart1(
	this js.Ref,
	ptr unsafe.Pointer,

	remoteParameters unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Start1
//go:noescape

func RTCIceTransportStart1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_Start2
//go:noescape

func CallRTCIceTransportStart2(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Start2
//go:noescape

func RTCIceTransportStart2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_Stop
//go:noescape

func CallRTCIceTransportStop(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_Stop
//go:noescape

func RTCIceTransportStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_AddRemoteCandidate
//go:noescape

func CallRTCIceTransportAddRemoteCandidate(
	this js.Ref,
	ptr unsafe.Pointer,

	remoteCandidate unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_AddRemoteCandidate
//go:noescape

func RTCIceTransportAddRemoteCandidateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCIceTransport_AddRemoteCandidate1
//go:noescape

func CallRTCIceTransportAddRemoteCandidate1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceTransport_AddRemoteCandidate1
//go:noescape

func RTCIceTransportAddRemoteCandidate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCDtlsTransportState
//go:noescape

func ConstOfRTCDtlsTransportState(str js.Ref) uint32

//go:wasmimport plat/js/web get_RTCDtlsTransport_IceTransport
//go:noescape

func GetRTCDtlsTransportIceTransport(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCDtlsTransport_State
//go:noescape

func GetRTCDtlsTransportState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_RTCDtlsTransport_GetRemoteCertificates
//go:noescape

func CallRTCDtlsTransportGetRemoteCertificates(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCDtlsTransport_GetRemoteCertificates
//go:noescape

func RTCDtlsTransportGetRemoteCertificatesFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCEncodedAudioFrame_Data
//go:noescape

func GetRTCEncodedAudioFrameData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_RTCEncodedAudioFrame_Data
//go:noescape

func SetRTCEncodedAudioFrameData(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_RTCEncodedAudioFrame_GetMetadata
//go:noescape

func CallRTCEncodedAudioFrameGetMetadata(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCEncodedAudioFrame_GetMetadata
//go:noescape

func RTCEncodedAudioFrameGetMetadataFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCEncodedVideoFrame_Timestamp
//go:noescape

func GetRTCEncodedVideoFrameTimestamp(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCEncodedVideoFrame_Data
//go:noescape

func GetRTCEncodedVideoFrameData(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_RTCEncodedVideoFrame_Data
//go:noescape

func SetRTCEncodedVideoFrameData(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_RTCEncodedVideoFrame_GetMetadata
//go:noescape

func CallRTCEncodedVideoFrameGetMetadata(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCEncodedVideoFrame_GetMetadata
//go:noescape

func RTCEncodedVideoFrameGetMetadataFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCError_SdpLineNumber
//go:noescape

func GetRTCErrorSdpLineNumber(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_RTCError_SctpCauseCode
//go:noescape

func GetRTCErrorSctpCauseCode(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_RTCError_ReceivedAlert
//go:noescape

func GetRTCErrorReceivedAlert(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCError_SentAlert
//go:noescape

func GetRTCErrorSentAlert(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCError_HttpRequestStatusCode
//go:noescape

func GetRTCErrorHttpRequestStatusCode(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_RTCIdentityAssertion_Idp
//go:noescape

func SetRTCIdentityAssertionIdp(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_RTCIdentityAssertion_Name
//go:noescape

func GetRTCIdentityAssertionName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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

//go:wasmimport plat/js/web call_RTCIdentityProviderRegistrar_Register
//go:noescape

func CallRTCIdentityProviderRegistrarRegister(
	this js.Ref,
	ptr unsafe.Pointer,

	idp unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCIdentityProviderRegistrar_Register
//go:noescape

func RTCIdentityProviderRegistrarRegisterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCIdentityProviderGlobalScope_RtcIdentityProvider
//go:noescape

func GetRTCIdentityProviderGlobalScopeRtcIdentityProvider(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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
