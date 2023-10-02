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

//go:wasmimport plat/js/web store_RTCRtpEncodingParameters
//go:noescape

func RTCRtpEncodingParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpEncodingParameters
//go:noescape

func RTCRtpEncodingParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpHeaderExtensionParameters
//go:noescape

func RTCRtpHeaderExtensionParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpHeaderExtensionParameters
//go:noescape

func RTCRtpHeaderExtensionParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtcpParameters
//go:noescape

func RTCRtcpParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtcpParameters
//go:noescape

func RTCRtcpParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpCodecParameters
//go:noescape

func RTCRtpCodecParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpCodecParameters
//go:noescape

func RTCRtpCodecParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpSendParameters
//go:noescape

func RTCRtpSendParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpSendParameters
//go:noescape

func RTCRtpSendParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCSetParameterOptions
//go:noescape

func RTCSetParameterOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCSetParameterOptions
//go:noescape

func RTCSetParameterOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_SFrameTransformRole
//go:noescape

func ConstOfSFrameTransformRole(str js.Ref) uint32

//go:wasmimport plat/js/web store_SFrameTransformOptions
//go:noescape

func SFrameTransformOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SFrameTransformOptions
//go:noescape

func SFrameTransformOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_SFrameTransform_SFrameTransform
//go:noescape

func NewSFrameTransformBySFrameTransform(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_SFrameTransform_SFrameTransform1
//go:noescape

func NewSFrameTransformBySFrameTransform1() js.Ref

//go:wasmimport plat/js/web get_SFrameTransform_Readable
//go:noescape

func GetSFrameTransformReadable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SFrameTransform_Writable
//go:noescape

func GetSFrameTransformWritable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SFrameTransform_SetEncryptionKey
//go:noescape

func CallSFrameTransformSetEncryptionKey(
	this js.Ref,
	ptr unsafe.Pointer,

	key js.Ref,
	keyID js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SFrameTransform_SetEncryptionKey
//go:noescape

func SFrameTransformSetEncryptionKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SFrameTransform_SetEncryptionKey1
//go:noescape

func CallSFrameTransformSetEncryptionKey1(
	this js.Ref,
	ptr unsafe.Pointer,

	key js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SFrameTransform_SetEncryptionKey1
//go:noescape

func SFrameTransformSetEncryptionKey1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_WorkerOptions
//go:noescape

func WorkerOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WorkerOptions
//go:noescape

func WorkerOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Worker_Worker
//go:noescape

func NewWorkerByWorker(
	scriptURL js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Worker_Worker1
//go:noescape

func NewWorkerByWorker1(
	scriptURL js.Ref) js.Ref

//go:wasmimport plat/js/web call_Worker_Terminate
//go:noescape

func CallWorkerTerminate(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Worker_Terminate
//go:noescape

func WorkerTerminateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Worker_PostMessage
//go:noescape

func CallWorkerPostMessage(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
	transfer js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Worker_PostMessage
//go:noescape

func WorkerPostMessageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Worker_PostMessage1
//go:noescape

func CallWorkerPostMessage1(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Worker_PostMessage1
//go:noescape

func WorkerPostMessage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Worker_PostMessage2
//go:noescape

func CallWorkerPostMessage2(
	this js.Ref,
	ptr unsafe.Pointer,

	message js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Worker_PostMessage2
//go:noescape

func WorkerPostMessage2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCRtpScriptTransform_RTCRtpScriptTransform
//go:noescape

func NewRTCRtpScriptTransformByRTCRtpScriptTransform(
	worker js.Ref,
	options js.Ref,
	transfer js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCRtpScriptTransform_RTCRtpScriptTransform1
//go:noescape

func NewRTCRtpScriptTransformByRTCRtpScriptTransform1(
	worker js.Ref,
	options js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCRtpScriptTransform_RTCRtpScriptTransform2
//go:noescape

func NewRTCRtpScriptTransformByRTCRtpScriptTransform2(
	worker js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCRtpSender_Track
//go:noescape

func GetRTCRtpSenderTrack(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpSender_Transport
//go:noescape

func GetRTCRtpSenderTransport(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpSender_Dtmf
//go:noescape

func GetRTCRtpSenderDtmf(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpSender_Transform
//go:noescape

func GetRTCRtpSenderTransform(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_RTCRtpSender_Transform
//go:noescape

func SetRTCRtpSenderTransform(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_RTCRtpSender_GetCapabilities
//go:noescape

func CallRTCRtpSenderGetCapabilities(
	this js.Ref,
	ptr unsafe.Pointer,

	kind js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCRtpSender_GetCapabilities
//go:noescape

func RTCRtpSenderGetCapabilitiesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpSender_SetParameters
//go:noescape

func CallRTCRtpSenderSetParameters(
	this js.Ref,
	ptr unsafe.Pointer,

	parameters unsafe.Pointer,
	setParameterOptions unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCRtpSender_SetParameters
//go:noescape

func RTCRtpSenderSetParametersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpSender_SetParameters1
//go:noescape

func CallRTCRtpSenderSetParameters1(
	this js.Ref,
	ptr unsafe.Pointer,

	parameters unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCRtpSender_SetParameters1
//go:noescape

func RTCRtpSenderSetParameters1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpSender_GetParameters
//go:noescape

func CallRTCRtpSenderGetParameters(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpSender_GetParameters
//go:noescape

func RTCRtpSenderGetParametersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpSender_ReplaceTrack
//go:noescape

func CallRTCRtpSenderReplaceTrack(
	this js.Ref,
	ptr unsafe.Pointer,

	withTrack js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCRtpSender_ReplaceTrack
//go:noescape

func RTCRtpSenderReplaceTrackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpSender_SetStreams
//go:noescape

func CallRTCRtpSenderSetStreams(
	this js.Ref,
	ptr unsafe.Pointer,

	streamsPtr unsafe.Pointer,
	streamsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_RTCRtpSender_SetStreams
//go:noescape

func RTCRtpSenderSetStreamsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpSender_GetStats
//go:noescape

func CallRTCRtpSenderGetStats(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpSender_GetStats
//go:noescape

func RTCRtpSenderGetStatsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpSender_GenerateKeyFrame
//go:noescape

func CallRTCRtpSenderGenerateKeyFrame(
	this js.Ref,
	ptr unsafe.Pointer,

	rids js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCRtpSender_GenerateKeyFrame
//go:noescape

func RTCRtpSenderGenerateKeyFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpSender_GenerateKeyFrame1
//go:noescape

func CallRTCRtpSenderGenerateKeyFrame1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpSender_GenerateKeyFrame1
//go:noescape

func RTCRtpSenderGenerateKeyFrame1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpReceiveParameters
//go:noescape

func RTCRtpReceiveParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpReceiveParameters
//go:noescape

func RTCRtpReceiveParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpContributingSource
//go:noescape

func RTCRtpContributingSourceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpContributingSource
//go:noescape

func RTCRtpContributingSourceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpSynchronizationSource
//go:noescape

func RTCRtpSynchronizationSourceJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpSynchronizationSource
//go:noescape

func RTCRtpSynchronizationSourceJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCRtpReceiver_Track
//go:noescape

func GetRTCRtpReceiverTrack(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpReceiver_Transport
//go:noescape

func GetRTCRtpReceiverTransport(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpReceiver_Transform
//go:noescape

func GetRTCRtpReceiverTransform(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_RTCRtpReceiver_Transform
//go:noescape

func SetRTCRtpReceiverTransform(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_RTCRtpReceiver_GetCapabilities
//go:noescape

func CallRTCRtpReceiverGetCapabilities(
	this js.Ref,
	ptr unsafe.Pointer,

	kind js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCRtpReceiver_GetCapabilities
//go:noescape

func RTCRtpReceiverGetCapabilitiesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpReceiver_GetParameters
//go:noescape

func CallRTCRtpReceiverGetParameters(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpReceiver_GetParameters
//go:noescape

func RTCRtpReceiverGetParametersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpReceiver_GetContributingSources
//go:noescape

func CallRTCRtpReceiverGetContributingSources(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpReceiver_GetContributingSources
//go:noescape

func RTCRtpReceiverGetContributingSourcesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpReceiver_GetSynchronizationSources
//go:noescape

func CallRTCRtpReceiverGetSynchronizationSources(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpReceiver_GetSynchronizationSources
//go:noescape

func RTCRtpReceiverGetSynchronizationSourcesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpReceiver_GetStats
//go:noescape

func CallRTCRtpReceiverGetStats(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpReceiver_GetStats
//go:noescape

func RTCRtpReceiverGetStatsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCRtpTransceiverDirection
//go:noescape

func ConstOfRTCRtpTransceiverDirection(str js.Ref) uint32

//go:wasmimport plat/js/web get_RTCRtpTransceiver_Mid
//go:noescape

func GetRTCRtpTransceiverMid(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpTransceiver_Sender
//go:noescape

func GetRTCRtpTransceiverSender(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpTransceiver_Receiver
//go:noescape

func GetRTCRtpTransceiverReceiver(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpTransceiver_Direction
//go:noescape

func GetRTCRtpTransceiverDirection(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_RTCRtpTransceiver_Direction
//go:noescape

func SetRTCRtpTransceiverDirection(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpTransceiver_CurrentDirection
//go:noescape

func GetRTCRtpTransceiverCurrentDirection(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_RTCRtpTransceiver_Stop
//go:noescape

func CallRTCRtpTransceiverStop(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpTransceiver_Stop
//go:noescape

func RTCRtpTransceiverStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpTransceiver_SetCodecPreferences
//go:noescape

func CallRTCRtpTransceiverSetCodecPreferences(
	this js.Ref,
	ptr unsafe.Pointer,

	codecs js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCRtpTransceiver_SetCodecPreferences
//go:noescape

func RTCRtpTransceiverSetCodecPreferencesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpTransceiverInit
//go:noescape

func RTCRtpTransceiverInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpTransceiverInit
//go:noescape

func RTCRtpTransceiverInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCSessionDescription_RTCSessionDescription
//go:noescape

func NewRTCSessionDescriptionByRTCSessionDescription(
	descriptionInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_RTCSessionDescription_Type
//go:noescape

func GetRTCSessionDescriptionType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCSessionDescription_Sdp
//go:noescape

func GetRTCSessionDescriptionSdp(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_RTCSessionDescription_ToJSON
//go:noescape

func CallRTCSessionDescriptionToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCSessionDescription_ToJSON
//go:noescape

func RTCSessionDescriptionToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RTCSignalingState
//go:noescape

func ConstOfRTCSignalingState(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCPeerConnectionState
//go:noescape

func ConstOfRTCPeerConnectionState(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RTCSctpTransportState
//go:noescape

func ConstOfRTCSctpTransportState(str js.Ref) uint32

//go:wasmimport plat/js/web get_RTCSctpTransport_Transport
//go:noescape

func GetRTCSctpTransportTransport(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCSctpTransport_State
//go:noescape

func GetRTCSctpTransportState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCSctpTransport_MaxMessageSize
//go:noescape

func GetRTCSctpTransportMaxMessageSize(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_RTCSctpTransport_MaxChannels
//go:noescape

func GetRTCSctpTransportMaxChannels(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web new_RTCPeerConnection_RTCPeerConnection
//go:noescape

func NewRTCPeerConnectionByRTCPeerConnection(
	configuration unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_RTCPeerConnection_RTCPeerConnection1
//go:noescape

func NewRTCPeerConnectionByRTCPeerConnection1() js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_LocalDescription
//go:noescape

func GetRTCPeerConnectionLocalDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_CurrentLocalDescription
//go:noescape

func GetRTCPeerConnectionCurrentLocalDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_PendingLocalDescription
//go:noescape

func GetRTCPeerConnectionPendingLocalDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_RemoteDescription
//go:noescape

func GetRTCPeerConnectionRemoteDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_CurrentRemoteDescription
//go:noescape

func GetRTCPeerConnectionCurrentRemoteDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_PendingRemoteDescription
//go:noescape

func GetRTCPeerConnectionPendingRemoteDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_SignalingState
//go:noescape

func GetRTCPeerConnectionSignalingState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCPeerConnection_IceGatheringState
//go:noescape

func GetRTCPeerConnectionIceGatheringState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCPeerConnection_IceConnectionState
//go:noescape

func GetRTCPeerConnectionIceConnectionState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCPeerConnection_ConnectionState
//go:noescape

func GetRTCPeerConnectionConnectionState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCPeerConnection_CanTrickleIceCandidates
//go:noescape

func GetRTCPeerConnectionCanTrickleIceCandidates(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_Sctp
//go:noescape

func GetRTCPeerConnectionSctp(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_PeerIdentity
//go:noescape

func GetRTCPeerConnectionPeerIdentity(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_IdpLoginUrl
//go:noescape

func GetRTCPeerConnectionIdpLoginUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnection_IdpErrorInfo
//go:noescape

func GetRTCPeerConnectionIdpErrorInfo(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_CreateOffer
//go:noescape

func CallRTCPeerConnectionCreateOffer(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_CreateOffer
//go:noescape

func RTCPeerConnectionCreateOfferFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_CreateOffer1
//go:noescape

func CallRTCPeerConnectionCreateOffer1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_CreateOffer1
//go:noescape

func RTCPeerConnectionCreateOffer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_CreateAnswer
//go:noescape

func CallRTCPeerConnectionCreateAnswer(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_CreateAnswer
//go:noescape

func RTCPeerConnectionCreateAnswerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_CreateAnswer1
//go:noescape

func CallRTCPeerConnectionCreateAnswer1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_CreateAnswer1
//go:noescape

func RTCPeerConnectionCreateAnswer1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_SetLocalDescription
//go:noescape

func CallRTCPeerConnectionSetLocalDescription(
	this js.Ref,
	ptr unsafe.Pointer,

	description unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_SetLocalDescription
//go:noescape

func RTCPeerConnectionSetLocalDescriptionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_SetLocalDescription1
//go:noescape

func CallRTCPeerConnectionSetLocalDescription1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_SetLocalDescription1
//go:noescape

func RTCPeerConnectionSetLocalDescription1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_SetRemoteDescription
//go:noescape

func CallRTCPeerConnectionSetRemoteDescription(
	this js.Ref,
	ptr unsafe.Pointer,

	description unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_SetRemoteDescription
//go:noescape

func RTCPeerConnectionSetRemoteDescriptionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_AddIceCandidate
//go:noescape

func CallRTCPeerConnectionAddIceCandidate(
	this js.Ref,
	ptr unsafe.Pointer,

	candidate unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_AddIceCandidate
//go:noescape

func RTCPeerConnectionAddIceCandidateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_AddIceCandidate1
//go:noescape

func CallRTCPeerConnectionAddIceCandidate1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_AddIceCandidate1
//go:noescape

func RTCPeerConnectionAddIceCandidate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_RestartIce
//go:noescape

func CallRTCPeerConnectionRestartIce(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_RestartIce
//go:noescape

func RTCPeerConnectionRestartIceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_GetConfiguration
//go:noescape

func CallRTCPeerConnectionGetConfiguration(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_GetConfiguration
//go:noescape

func RTCPeerConnectionGetConfigurationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_SetConfiguration
//go:noescape

func CallRTCPeerConnectionSetConfiguration(
	this js.Ref,
	ptr unsafe.Pointer,

	configuration unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_SetConfiguration
//go:noescape

func RTCPeerConnectionSetConfigurationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_SetConfiguration1
//go:noescape

func CallRTCPeerConnectionSetConfiguration1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_SetConfiguration1
//go:noescape

func RTCPeerConnectionSetConfiguration1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_Close
//go:noescape

func CallRTCPeerConnectionClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_Close
//go:noescape

func RTCPeerConnectionCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_CreateOffer2
//go:noescape

func CallRTCPeerConnectionCreateOffer2(
	this js.Ref,
	ptr unsafe.Pointer,

	successCallback js.Ref,
	failureCallback js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_CreateOffer2
//go:noescape

func RTCPeerConnectionCreateOffer2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_CreateOffer3
//go:noescape

func CallRTCPeerConnectionCreateOffer3(
	this js.Ref,
	ptr unsafe.Pointer,

	successCallback js.Ref,
	failureCallback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_CreateOffer3
//go:noescape

func RTCPeerConnectionCreateOffer3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_SetLocalDescription2
//go:noescape

func CallRTCPeerConnectionSetLocalDescription2(
	this js.Ref,
	ptr unsafe.Pointer,

	description unsafe.Pointer,
	successCallback js.Ref,
	failureCallback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_SetLocalDescription2
//go:noescape

func RTCPeerConnectionSetLocalDescription2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_CreateAnswer2
//go:noescape

func CallRTCPeerConnectionCreateAnswer2(
	this js.Ref,
	ptr unsafe.Pointer,

	successCallback js.Ref,
	failureCallback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_CreateAnswer2
//go:noescape

func RTCPeerConnectionCreateAnswer2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_SetRemoteDescription1
//go:noescape

func CallRTCPeerConnectionSetRemoteDescription1(
	this js.Ref,
	ptr unsafe.Pointer,

	description unsafe.Pointer,
	successCallback js.Ref,
	failureCallback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_SetRemoteDescription1
//go:noescape

func RTCPeerConnectionSetRemoteDescription1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_AddIceCandidate2
//go:noescape

func CallRTCPeerConnectionAddIceCandidate2(
	this js.Ref,
	ptr unsafe.Pointer,

	candidate unsafe.Pointer,
	successCallback js.Ref,
	failureCallback js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_AddIceCandidate2
//go:noescape

func RTCPeerConnectionAddIceCandidate2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_CreateDataChannel
//go:noescape

func CallRTCPeerConnectionCreateDataChannel(
	this js.Ref,
	ptr unsafe.Pointer,

	label js.Ref,
	dataChannelDict unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_CreateDataChannel
//go:noescape

func RTCPeerConnectionCreateDataChannelFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_CreateDataChannel1
//go:noescape

func CallRTCPeerConnectionCreateDataChannel1(
	this js.Ref,
	ptr unsafe.Pointer,

	label js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_CreateDataChannel1
//go:noescape

func RTCPeerConnectionCreateDataChannel1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_GetSenders
//go:noescape

func CallRTCPeerConnectionGetSenders(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_GetSenders
//go:noescape

func RTCPeerConnectionGetSendersFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_GetReceivers
//go:noescape

func CallRTCPeerConnectionGetReceivers(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_GetReceivers
//go:noescape

func RTCPeerConnectionGetReceiversFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_GetTransceivers
//go:noescape

func CallRTCPeerConnectionGetTransceivers(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_GetTransceivers
//go:noescape

func RTCPeerConnectionGetTransceiversFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_AddTrack
//go:noescape

func CallRTCPeerConnectionAddTrack(
	this js.Ref,
	ptr unsafe.Pointer,

	track js.Ref,
	streamsPtr unsafe.Pointer,
	streamsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_AddTrack
//go:noescape

func RTCPeerConnectionAddTrackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_RemoveTrack
//go:noescape

func CallRTCPeerConnectionRemoveTrack(
	this js.Ref,
	ptr unsafe.Pointer,

	sender js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_RemoveTrack
//go:noescape

func RTCPeerConnectionRemoveTrackFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_AddTransceiver
//go:noescape

func CallRTCPeerConnectionAddTransceiver(
	this js.Ref,
	ptr unsafe.Pointer,

	trackOrKind js.Ref,
	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_AddTransceiver
//go:noescape

func RTCPeerConnectionAddTransceiverFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_AddTransceiver1
//go:noescape

func CallRTCPeerConnectionAddTransceiver1(
	this js.Ref,
	ptr unsafe.Pointer,

	trackOrKind js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_AddTransceiver1
//go:noescape

func RTCPeerConnectionAddTransceiver1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_GetStats
//go:noescape

func CallRTCPeerConnectionGetStats(
	this js.Ref,
	ptr unsafe.Pointer,

	selector js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_GetStats
//go:noescape

func RTCPeerConnectionGetStatsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_GetStats1
//go:noescape

func CallRTCPeerConnectionGetStats1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_GetStats1
//go:noescape

func RTCPeerConnectionGetStats1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_GenerateCertificate
//go:noescape

func CallRTCPeerConnectionGenerateCertificate(
	this js.Ref,
	ptr unsafe.Pointer,

	keygenAlgorithm js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_GenerateCertificate
//go:noescape

func RTCPeerConnectionGenerateCertificateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_SetIdentityProvider
//go:noescape

func CallRTCPeerConnectionSetIdentityProvider(
	this js.Ref,
	ptr unsafe.Pointer,

	provider js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_SetIdentityProvider
//go:noescape

func RTCPeerConnectionSetIdentityProviderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_SetIdentityProvider1
//go:noescape

func CallRTCPeerConnectionSetIdentityProvider1(
	this js.Ref,
	ptr unsafe.Pointer,

	provider js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_SetIdentityProvider1
//go:noescape

func RTCPeerConnectionSetIdentityProvider1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCPeerConnection_GetIdentityAssertion
//go:noescape

func CallRTCPeerConnectionGetIdentityAssertion(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCPeerConnection_GetIdentityAssertion
//go:noescape

func RTCPeerConnectionGetIdentityAssertionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCPeerConnectionIceErrorEventInit
//go:noescape

func RTCPeerConnectionIceErrorEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCPeerConnectionIceErrorEventInit
//go:noescape

func RTCPeerConnectionIceErrorEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCPeerConnectionIceErrorEvent_RTCPeerConnectionIceErrorEvent
//go:noescape

func NewRTCPeerConnectionIceErrorEventByRTCPeerConnectionIceErrorEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnectionIceErrorEvent_Address
//go:noescape

func GetRTCPeerConnectionIceErrorEventAddress(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnectionIceErrorEvent_Port
//go:noescape

func GetRTCPeerConnectionIceErrorEventPort(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCPeerConnectionIceErrorEvent_Url
//go:noescape

func GetRTCPeerConnectionIceErrorEventUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnectionIceErrorEvent_ErrorCode
//go:noescape

func GetRTCPeerConnectionIceErrorEventErrorCode(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_RTCPeerConnectionIceErrorEvent_ErrorText
//go:noescape

func GetRTCPeerConnectionIceErrorEventErrorText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_RTCPeerConnectionIceEventInit
//go:noescape

func RTCPeerConnectionIceEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCPeerConnectionIceEventInit
//go:noescape

func RTCPeerConnectionIceEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_RTCPeerConnectionIceEvent_RTCPeerConnectionIceEvent
//go:noescape

func NewRTCPeerConnectionIceEventByRTCPeerConnectionIceEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_RTCPeerConnectionIceEvent_RTCPeerConnectionIceEvent1
//go:noescape

func NewRTCPeerConnectionIceEventByRTCPeerConnectionIceEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnectionIceEvent_Candidate
//go:noescape

func GetRTCPeerConnectionIceEventCandidate(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCPeerConnectionIceEvent_Url
//go:noescape

func GetRTCPeerConnectionIceEventUrl(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_RTCPeerConnectionStats
//go:noescape

func RTCPeerConnectionStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCPeerConnectionStats
//go:noescape

func RTCPeerConnectionStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCReceivedRtpStreamStats
//go:noescape

func RTCReceivedRtpStreamStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCReceivedRtpStreamStats
//go:noescape

func RTCReceivedRtpStreamStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRemoteInboundRtpStreamStats
//go:noescape

func RTCRemoteInboundRtpStreamStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRemoteInboundRtpStreamStats
//go:noescape

func RTCRemoteInboundRtpStreamStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRemoteOutboundRtpStreamStats
//go:noescape

func RTCRemoteOutboundRtpStreamStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRemoteOutboundRtpStreamStats
//go:noescape

func RTCRemoteOutboundRtpStreamStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpCodec
//go:noescape

func RTCRtpCodecJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpCodec
//go:noescape

func RTCRtpCodecJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpCodingParameters
//go:noescape

func RTCRtpCodingParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpCodingParameters
//go:noescape

func RTCRtpCodingParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpParameters
//go:noescape

func RTCRtpParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpParameters
//go:noescape

func RTCRtpParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_RTCRtpScriptTransformer_Readable
//go:noescape

func GetRTCRtpScriptTransformerReadable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpScriptTransformer_Writable
//go:noescape

func GetRTCRtpScriptTransformerWritable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_RTCRtpScriptTransformer_Options
//go:noescape

func GetRTCRtpScriptTransformerOptions(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_RTCRtpScriptTransformer_GenerateKeyFrame
//go:noescape

func CallRTCRtpScriptTransformerGenerateKeyFrame(
	this js.Ref,
	ptr unsafe.Pointer,

	rid js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_RTCRtpScriptTransformer_GenerateKeyFrame
//go:noescape

func RTCRtpScriptTransformerGenerateKeyFrameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpScriptTransformer_GenerateKeyFrame1
//go:noescape

func CallRTCRtpScriptTransformerGenerateKeyFrame1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpScriptTransformer_GenerateKeyFrame1
//go:noescape

func RTCRtpScriptTransformerGenerateKeyFrame1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_RTCRtpScriptTransformer_SendKeyFrameRequest
//go:noescape

func CallRTCRtpScriptTransformerSendKeyFrameRequest(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_RTCRtpScriptTransformer_SendKeyFrameRequest
//go:noescape

func RTCRtpScriptTransformerSendKeyFrameRequestFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCRtpStreamStats
//go:noescape

func RTCRtpStreamStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCRtpStreamStats
//go:noescape

func RTCRtpStreamStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCSentRtpStreamStats
//go:noescape

func RTCSentRtpStreamStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCSentRtpStreamStats
//go:noescape

func RTCSentRtpStreamStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RTCStats
//go:noescape

func RTCStatsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RTCStats
//go:noescape

func RTCStatsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
