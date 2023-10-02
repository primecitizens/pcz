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

//go:wasmimport plat/js/web store_RTCAudioPlayoutStats
//go:noescape

func RTCAudioPlayoutStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCAudioPlayoutStats
//go:noescape

func RTCAudioPlayoutStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCAudioSourceStats
//go:noescape

func RTCAudioSourceStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCAudioSourceStats
//go:noescape

func RTCAudioSourceStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCBundlePolicy
//go:noescape

func ConstOfRTCBundlePolicy(str js.Ref) uint32

//go:wasmimport plat/js/web store_RTCDtlsFingerprint
//go:noescape

func RTCDtlsFingerprintJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCDtlsFingerprint
//go:noescape

func RTCDtlsFingerprintJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCCertificate_Expires
//go:noescape

func GetRTCCertificateExpires(
	this js.Ref,
	ptr unsafe.Pointer,
) uint64

//go:wasmimport plat/js/web call_RTCCertificate_GetFingerprints
//go:noescape

func CallRTCCertificateGetFingerprints(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCCertificate_GetFingerprints
//go:noescape

func RTCCertificateGetFingerprintsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCCertificateExpiration
//go:noescape

func RTCCertificateExpirationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCCertificateExpiration
//go:noescape

func RTCCertificateExpirationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCCertificateStats
//go:noescape

func RTCCertificateStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCCertificateStats
//go:noescape

func RTCCertificateStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCCodecStats
//go:noescape

func RTCCodecStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCCodecStats
//go:noescape

func RTCCodecStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCIceServer
//go:noescape

func RTCIceServerJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIceServer
//go:noescape

func RTCIceServerJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCIceTransportPolicy
//go:noescape

func ConstOfRTCIceTransportPolicy(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCRtcpMuxPolicy
//go:noescape

func ConstOfRTCRtcpMuxPolicy(str js.Ref) uint32

//go:wasmimport plat/js/web store_RTCConfiguration
//go:noescape

func RTCConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCConfiguration
//go:noescape

func RTCConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCDTMFSender_CanInsertDTMF
//go:noescape

func GetRTCDTMFSenderCanInsertDTMF(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCDTMFSender_ToneBuffer
//go:noescape

func GetRTCDTMFSenderToneBuffer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_RTCDTMFSender_InsertDTMF
//go:noescape

func CallRTCDTMFSenderInsertDTMF(
	this js.Ref,
	ptr unsafe.Pointer,

	tones js.Ref,
	duration uint32,
	interToneGap uint32,
) js.Ref

//go:wasmimport plat/js/web func_RTCDTMFSender_InsertDTMF
//go:noescape

func RTCDTMFSenderInsertDTMFFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCDTMFSender_InsertDTMF1
//go:noescape

func CallRTCDTMFSenderInsertDTMF1(
	this js.Ref,
	ptr unsafe.Pointer,

	tones js.Ref,
	duration uint32,
) js.Ref

//go:wasmimport plat/js/web func_RTCDTMFSender_InsertDTMF1
//go:noescape

func RTCDTMFSenderInsertDTMF1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCDTMFSender_InsertDTMF2
//go:noescape

func CallRTCDTMFSenderInsertDTMF2(
	this js.Ref,
	ptr unsafe.Pointer,

	tones js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCDTMFSender_InsertDTMF2
//go:noescape

func RTCDTMFSenderInsertDTMF2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCDTMFToneChangeEventInit
//go:noescape

func RTCDTMFToneChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCDTMFToneChangeEventInit
//go:noescape

func RTCDTMFToneChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCDTMFToneChangeEvent_RTCDTMFToneChangeEvent
//go:noescape

func NewRTCDTMFToneChangeEventByRTCDTMFToneChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_RTCDTMFToneChangeEvent_RTCDTMFToneChangeEvent1
//go:noescape

func NewRTCDTMFToneChangeEventByRTCDTMFToneChangeEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCDTMFToneChangeEvent_Tone
//go:noescape

func GetRTCDTMFToneChangeEventTone(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web constof_RTCDataChannelState
//go:noescape

func ConstOfRTCDataChannelState(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCPriorityType
//go:noescape

func ConstOfRTCPriorityType(str js.Ref) uint32

//go:wasmimport plat/js/web get_RTCDataChannel_Label
//go:noescape

func GetRTCDataChannelLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCDataChannel_Ordered
//go:noescape

func GetRTCDataChannelOrdered(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCDataChannel_MaxPacketLifeTime
//go:noescape

func GetRTCDataChannelMaxPacketLifeTime(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCDataChannel_MaxRetransmits
//go:noescape

func GetRTCDataChannelMaxRetransmits(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCDataChannel_Protocol
//go:noescape

func GetRTCDataChannelProtocol(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCDataChannel_Negotiated
//go:noescape

func GetRTCDataChannelNegotiated(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCDataChannel_Id
//go:noescape

func GetRTCDataChannelId(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCDataChannel_ReadyState
//go:noescape

func GetRTCDataChannelReadyState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCDataChannel_BufferedAmount
//go:noescape

func GetRTCDataChannelBufferedAmount(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCDataChannel_BufferedAmountLowThreshold
//go:noescape

func GetRTCDataChannelBufferedAmountLowThreshold(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_RTCDataChannel_BufferedAmountLowThreshold
//go:noescape

func SetRTCDataChannelBufferedAmountLowThreshold(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_RTCDataChannel_BinaryType
//go:noescape

func GetRTCDataChannelBinaryType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_RTCDataChannel_BinaryType
//go:noescape

func SetRTCDataChannelBinaryType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_RTCDataChannel_Priority
//go:noescape

func GetRTCDataChannelPriority(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_RTCDataChannel_Close
//go:noescape

func CallRTCDataChannelClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCDataChannel_Close
//go:noescape

func RTCDataChannelCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCDataChannel_Send
//go:noescape

func CallRTCDataChannelSend(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCDataChannel_Send
//go:noescape

func RTCDataChannelSendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCDataChannel_Send1
//go:noescape

func CallRTCDataChannelSend1(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCDataChannel_Send1
//go:noescape

func RTCDataChannelSend1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCDataChannel_Send2
//go:noescape

func CallRTCDataChannelSend2(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCDataChannel_Send2
//go:noescape

func RTCDataChannelSend2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCDataChannel_Send3
//go:noescape

func CallRTCDataChannelSend3(
	this js.Ref,
	ptr unsafe.Pointer,

	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCDataChannel_Send3
//go:noescape

func RTCDataChannelSend3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCDataChannelEventInit
//go:noescape

func RTCDataChannelEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCDataChannelEventInit
//go:noescape

func RTCDataChannelEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCDataChannelEvent_RTCDataChannelEvent
//go:noescape

func NewRTCDataChannelEventByRTCDataChannelEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_RTCDataChannelEvent_Channel
//go:noescape

func GetRTCDataChannelEventChannel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_RTCDataChannelInit
//go:noescape

func RTCDataChannelInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCDataChannelInit
//go:noescape

func RTCDataChannelInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCDataChannelStats
//go:noescape

func RTCDataChannelStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCDataChannelStats
//go:noescape

func RTCDataChannelStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCDegradationPreference
//go:noescape

func ConstOfRTCDegradationPreference(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCDtlsRole
//go:noescape

func ConstOfRTCDtlsRole(str js.Ref) uint32

//go:wasmimport plat/js/web store_RTCIceCandidateInit
//go:noescape

func RTCIceCandidateInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIceCandidateInit
//go:noescape

func RTCIceCandidateInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCIceComponent
//go:noescape

func ConstOfRTCIceComponent(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCIceProtocol
//go:noescape

func ConstOfRTCIceProtocol(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCIceCandidateType
//go:noescape

func ConstOfRTCIceCandidateType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCIceTcpCandidateType
//go:noescape

func ConstOfRTCIceTcpCandidateType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCIceServerTransportProtocol
//go:noescape

func ConstOfRTCIceServerTransportProtocol(str js.Ref) uint32

//go:wasmimport plat/js/web new_RTCIceCandidate_RTCIceCandidate
//go:noescape

func NewRTCIceCandidateByRTCIceCandidate(
	candidateInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_RTCIceCandidate_RTCIceCandidate1
//go:noescape

func NewRTCIceCandidateByRTCIceCandidate1() js.Ref

//go:wasmimport plat/js/web get_RTCIceCandidate_Candidate
//go:noescape

func GetRTCIceCandidateCandidate(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCIceCandidate_SdpMid
//go:noescape

func GetRTCIceCandidateSdpMid(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCIceCandidate_SdpMLineIndex
//go:noescape

func GetRTCIceCandidateSdpMLineIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceCandidate_Foundation
//go:noescape

func GetRTCIceCandidateFoundation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCIceCandidate_Component
//go:noescape

func GetRTCIceCandidateComponent(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceCandidate_Priority
//go:noescape

func GetRTCIceCandidatePriority(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceCandidate_Address
//go:noescape

func GetRTCIceCandidateAddress(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCIceCandidate_Protocol
//go:noescape

func GetRTCIceCandidateProtocol(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceCandidate_Port
//go:noescape

func GetRTCIceCandidatePort(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceCandidate_Type
//go:noescape

func GetRTCIceCandidateType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceCandidate_TcpType
//go:noescape

func GetRTCIceCandidateTcpType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceCandidate_RelatedAddress
//go:noescape

func GetRTCIceCandidateRelatedAddress(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCIceCandidate_RelatedPort
//go:noescape

func GetRTCIceCandidateRelatedPort(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceCandidate_UsernameFragment
//go:noescape

func GetRTCIceCandidateUsernameFragment(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCIceCandidate_RelayProtocol
//go:noescape

func GetRTCIceCandidateRelayProtocol(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCIceCandidate_Url
//go:noescape

func GetRTCIceCandidateUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_RTCIceCandidate_ToJSON
//go:noescape

func CallRTCIceCandidateToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCIceCandidate_ToJSON
//go:noescape

func RTCIceCandidateToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCIceCandidatePair
//go:noescape

func RTCIceCandidatePairJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIceCandidatePair
//go:noescape

func RTCIceCandidatePairJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCIceParameters
//go:noescape

func RTCIceParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIceParameters
//go:noescape

func RTCIceParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCIceGatherOptions
//go:noescape

func RTCIceGatherOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCIceGatherOptions
//go:noescape

func RTCIceGatherOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCIceRole
//go:noescape

func ConstOfRTCIceRole(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCIceTransportState
//go:noescape

func ConstOfRTCIceTransportState(str js.Ref) uint32
