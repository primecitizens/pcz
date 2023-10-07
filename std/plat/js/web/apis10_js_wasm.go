// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type AudioEncoderInit struct {
	// Output is "AudioEncoderInit.output"
	//
	// Required
	Output js.Func[func(output EncodedAudioChunk, metadata *EncodedAudioChunkMetadata)]
	// Error is "AudioEncoderInit.error"
	//
	// Required
	Error js.Func[func(err DOMException)]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioEncoderInit with all fields set.
func (p AudioEncoderInit) FromRef(ref js.Ref) AudioEncoderInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioEncoderInit in the application heap.
func (p AudioEncoderInit) New() js.Ref {
	return bindings.AudioEncoderInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioEncoderInit) UpdateFrom(ref js.Ref) {
	bindings.AudioEncoderInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioEncoderInit) Update(ref js.Ref) {
	bindings.AudioEncoderInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioEncoderInit) FreeMembers(recursive bool) {
	js.Free(
		p.Output.Ref(),
		p.Error.Ref(),
	)
	p.Output = p.Output.FromRef(js.Undefined)
	p.Error = p.Error.FromRef(js.Undefined)
}

type BitrateMode uint32

const (
	_ BitrateMode = iota

	BitrateMode_CONSTANT
	BitrateMode_VARIABLE
)

func (BitrateMode) FromRef(str js.Ref) BitrateMode {
	return BitrateMode(bindings.ConstOfBitrateMode(str))
}

func (x BitrateMode) String() (string, bool) {
	switch x {
	case BitrateMode_CONSTANT:
		return "constant", true
	case BitrateMode_VARIABLE:
		return "variable", true
	default:
		return "", false
	}
}

type FlacEncoderConfig struct {
	// BlockSize is "FlacEncoderConfig.blockSize"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_BlockSize MUST be set to true to make this field effective.
	BlockSize uint32
	// CompressLevel is "FlacEncoderConfig.compressLevel"
	//
	// Optional, defaults to 5.
	//
	// NOTE: FFI_USE_CompressLevel MUST be set to true to make this field effective.
	CompressLevel uint32

	FFI_USE_BlockSize     bool // for BlockSize.
	FFI_USE_CompressLevel bool // for CompressLevel.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FlacEncoderConfig with all fields set.
func (p FlacEncoderConfig) FromRef(ref js.Ref) FlacEncoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FlacEncoderConfig in the application heap.
func (p FlacEncoderConfig) New() js.Ref {
	return bindings.FlacEncoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FlacEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.FlacEncoderConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FlacEncoderConfig) Update(ref js.Ref) {
	bindings.FlacEncoderConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FlacEncoderConfig) FreeMembers(recursive bool) {
}

type OpusBitstreamFormat uint32

const (
	_ OpusBitstreamFormat = iota

	OpusBitstreamFormat_OPUS
	OpusBitstreamFormat_OGG
)

func (OpusBitstreamFormat) FromRef(str js.Ref) OpusBitstreamFormat {
	return OpusBitstreamFormat(bindings.ConstOfOpusBitstreamFormat(str))
}

func (x OpusBitstreamFormat) String() (string, bool) {
	switch x {
	case OpusBitstreamFormat_OPUS:
		return "opus", true
	case OpusBitstreamFormat_OGG:
		return "ogg", true
	default:
		return "", false
	}
}

type OpusEncoderConfig struct {
	// Format is "OpusEncoderConfig.format"
	//
	// Optional, defaults to "opus".
	Format OpusBitstreamFormat
	// FrameDuration is "OpusEncoderConfig.frameDuration"
	//
	// Optional, defaults to 20000.
	//
	// NOTE: FFI_USE_FrameDuration MUST be set to true to make this field effective.
	FrameDuration uint64
	// Complexity is "OpusEncoderConfig.complexity"
	//
	// Optional
	//
	// NOTE: FFI_USE_Complexity MUST be set to true to make this field effective.
	Complexity uint32
	// Packetlossperc is "OpusEncoderConfig.packetlossperc"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Packetlossperc MUST be set to true to make this field effective.
	Packetlossperc uint32
	// Useinbandfec is "OpusEncoderConfig.useinbandfec"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Useinbandfec MUST be set to true to make this field effective.
	Useinbandfec bool
	// Usedtx is "OpusEncoderConfig.usedtx"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Usedtx MUST be set to true to make this field effective.
	Usedtx bool

	FFI_USE_FrameDuration  bool // for FrameDuration.
	FFI_USE_Complexity     bool // for Complexity.
	FFI_USE_Packetlossperc bool // for Packetlossperc.
	FFI_USE_Useinbandfec   bool // for Useinbandfec.
	FFI_USE_Usedtx         bool // for Usedtx.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OpusEncoderConfig with all fields set.
func (p OpusEncoderConfig) FromRef(ref js.Ref) OpusEncoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OpusEncoderConfig in the application heap.
func (p OpusEncoderConfig) New() js.Ref {
	return bindings.OpusEncoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OpusEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.OpusEncoderConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OpusEncoderConfig) Update(ref js.Ref) {
	bindings.OpusEncoderConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OpusEncoderConfig) FreeMembers(recursive bool) {
}

type AudioEncoderConfig struct {
	// Codec is "AudioEncoderConfig.codec"
	//
	// Required
	Codec js.String
	// SampleRate is "AudioEncoderConfig.sampleRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_SampleRate MUST be set to true to make this field effective.
	SampleRate uint32
	// NumberOfChannels is "AudioEncoderConfig.numberOfChannels"
	//
	// Optional
	//
	// NOTE: FFI_USE_NumberOfChannels MUST be set to true to make this field effective.
	NumberOfChannels uint32
	// Bitrate is "AudioEncoderConfig.bitrate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Bitrate MUST be set to true to make this field effective.
	Bitrate uint64
	// BitrateMode is "AudioEncoderConfig.bitrateMode"
	//
	// Optional, defaults to "variable".
	BitrateMode BitrateMode
	// Flac is "AudioEncoderConfig.flac"
	//
	// Optional
	//
	// NOTE: Flac.FFI_USE MUST be set to true to get Flac used.
	Flac FlacEncoderConfig
	// Opus is "AudioEncoderConfig.opus"
	//
	// Optional
	//
	// NOTE: Opus.FFI_USE MUST be set to true to get Opus used.
	Opus OpusEncoderConfig
	// Aac is "AudioEncoderConfig.aac"
	//
	// Optional
	//
	// NOTE: Aac.FFI_USE MUST be set to true to get Aac used.
	Aac AacEncoderConfig

	FFI_USE_SampleRate       bool // for SampleRate.
	FFI_USE_NumberOfChannels bool // for NumberOfChannels.
	FFI_USE_Bitrate          bool // for Bitrate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioEncoderConfig with all fields set.
func (p AudioEncoderConfig) FromRef(ref js.Ref) AudioEncoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioEncoderConfig in the application heap.
func (p AudioEncoderConfig) New() js.Ref {
	return bindings.AudioEncoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.AudioEncoderConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioEncoderConfig) Update(ref js.Ref) {
	bindings.AudioEncoderConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioEncoderConfig) FreeMembers(recursive bool) {
	js.Free(
		p.Codec.Ref(),
	)
	p.Codec = p.Codec.FromRef(js.Undefined)
	if recursive {
		p.Flac.FreeMembers(true)
		p.Opus.FreeMembers(true)
		p.Aac.FreeMembers(true)
	}
}

type AudioEncoderSupport struct {
	// Supported is "AudioEncoderSupport.supported"
	//
	// Optional
	//
	// NOTE: FFI_USE_Supported MUST be set to true to make this field effective.
	Supported bool
	// Config is "AudioEncoderSupport.config"
	//
	// Optional
	//
	// NOTE: Config.FFI_USE MUST be set to true to get Config used.
	Config AudioEncoderConfig

	FFI_USE_Supported bool // for Supported.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioEncoderSupport with all fields set.
func (p AudioEncoderSupport) FromRef(ref js.Ref) AudioEncoderSupport {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioEncoderSupport in the application heap.
func (p AudioEncoderSupport) New() js.Ref {
	return bindings.AudioEncoderSupportJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioEncoderSupport) UpdateFrom(ref js.Ref) {
	bindings.AudioEncoderSupportJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioEncoderSupport) Update(ref js.Ref) {
	bindings.AudioEncoderSupportJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioEncoderSupport) FreeMembers(recursive bool) {
	if recursive {
		p.Config.FreeMembers(true)
	}
}

func NewAudioEncoder(init AudioEncoderInit) (ret AudioEncoder) {
	ret.ref = bindings.NewAudioEncoderByAudioEncoder(
		js.Pointer(&init))
	return
}

type AudioEncoder struct {
	EventTarget
}

func (this AudioEncoder) Once() AudioEncoder {
	this.ref.Once()
	return this
}

func (this AudioEncoder) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this AudioEncoder) FromRef(ref js.Ref) AudioEncoder {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this AudioEncoder) Free() {
	this.ref.Free()
}

// State returns the value of property "AudioEncoder.state".
//
// It returns ok=false if there is no such property.
func (this AudioEncoder) State() (ret CodecState, ok bool) {
	ok = js.True == bindings.GetAudioEncoderState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EncodeQueueSize returns the value of property "AudioEncoder.encodeQueueSize".
//
// It returns ok=false if there is no such property.
func (this AudioEncoder) EncodeQueueSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioEncoderEncodeQueueSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncConfigure returns true if the method "AudioEncoder.configure" exists.
func (this AudioEncoder) HasFuncConfigure() bool {
	return js.True == bindings.HasFuncAudioEncoderConfigure(
		this.ref,
	)
}

// FuncConfigure returns the method "AudioEncoder.configure".
func (this AudioEncoder) FuncConfigure() (fn js.Func[func(config AudioEncoderConfig)]) {
	bindings.FuncAudioEncoderConfigure(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Configure calls the method "AudioEncoder.configure".
func (this AudioEncoder) Configure(config AudioEncoderConfig) (ret js.Void) {
	bindings.CallAudioEncoderConfigure(
		this.ref, js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryConfigure calls the method "AudioEncoder.configure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioEncoder) TryConfigure(config AudioEncoderConfig) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderConfigure(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

// HasFuncEncode returns true if the method "AudioEncoder.encode" exists.
func (this AudioEncoder) HasFuncEncode() bool {
	return js.True == bindings.HasFuncAudioEncoderEncode(
		this.ref,
	)
}

// FuncEncode returns the method "AudioEncoder.encode".
func (this AudioEncoder) FuncEncode() (fn js.Func[func(data AudioData)]) {
	bindings.FuncAudioEncoderEncode(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Encode calls the method "AudioEncoder.encode".
func (this AudioEncoder) Encode(data AudioData) (ret js.Void) {
	bindings.CallAudioEncoderEncode(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryEncode calls the method "AudioEncoder.encode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioEncoder) TryEncode(data AudioData) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderEncode(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncFlush returns true if the method "AudioEncoder.flush" exists.
func (this AudioEncoder) HasFuncFlush() bool {
	return js.True == bindings.HasFuncAudioEncoderFlush(
		this.ref,
	)
}

// FuncFlush returns the method "AudioEncoder.flush".
func (this AudioEncoder) FuncFlush() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncAudioEncoderFlush(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Flush calls the method "AudioEncoder.flush".
func (this AudioEncoder) Flush() (ret js.Promise[js.Void]) {
	bindings.CallAudioEncoderFlush(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "AudioEncoder.flush"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioEncoder) TryFlush() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderFlush(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReset returns true if the method "AudioEncoder.reset" exists.
func (this AudioEncoder) HasFuncReset() bool {
	return js.True == bindings.HasFuncAudioEncoderReset(
		this.ref,
	)
}

// FuncReset returns the method "AudioEncoder.reset".
func (this AudioEncoder) FuncReset() (fn js.Func[func()]) {
	bindings.FuncAudioEncoderReset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reset calls the method "AudioEncoder.reset".
func (this AudioEncoder) Reset() (ret js.Void) {
	bindings.CallAudioEncoderReset(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "AudioEncoder.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioEncoder) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderReset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "AudioEncoder.close" exists.
func (this AudioEncoder) HasFuncClose() bool {
	return js.True == bindings.HasFuncAudioEncoderClose(
		this.ref,
	)
}

// FuncClose returns the method "AudioEncoder.close".
func (this AudioEncoder) FuncClose() (fn js.Func[func()]) {
	bindings.FuncAudioEncoderClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "AudioEncoder.close".
func (this AudioEncoder) Close() (ret js.Void) {
	bindings.CallAudioEncoderClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "AudioEncoder.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioEncoder) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsConfigSupported returns true if the static method "AudioEncoder.isConfigSupported" exists.
func (this AudioEncoder) HasFuncIsConfigSupported() bool {
	return js.True == bindings.HasFuncAudioEncoderIsConfigSupported(
		this.ref,
	)
}

// FuncIsConfigSupported returns the static method "AudioEncoder.isConfigSupported".
func (this AudioEncoder) FuncIsConfigSupported() (fn js.Func[func(config AudioEncoderConfig) js.Promise[AudioEncoderSupport]]) {
	bindings.FuncAudioEncoderIsConfigSupported(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsConfigSupported calls the static method "AudioEncoder.isConfigSupported".
func (this AudioEncoder) IsConfigSupported(config AudioEncoderConfig) (ret js.Promise[AudioEncoderSupport]) {
	bindings.CallAudioEncoderIsConfigSupported(
		this.ref, js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryIsConfigSupported calls the static method "AudioEncoder.isConfigSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioEncoder) TryIsConfigSupported(config AudioEncoderConfig) (ret js.Promise[AudioEncoderSupport], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderIsConfigSupported(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

type AudioNode struct {
	EventTarget
}

func (this AudioNode) Once() AudioNode {
	this.ref.Once()
	return this
}

func (this AudioNode) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this AudioNode) FromRef(ref js.Ref) AudioNode {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this AudioNode) Free() {
	this.ref.Free()
}

// Context returns the value of property "AudioNode.context".
//
// It returns ok=false if there is no such property.
func (this AudioNode) Context() (ret BaseAudioContext, ok bool) {
	ok = js.True == bindings.GetAudioNodeContext(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NumberOfInputs returns the value of property "AudioNode.numberOfInputs".
//
// It returns ok=false if there is no such property.
func (this AudioNode) NumberOfInputs() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioNodeNumberOfInputs(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NumberOfOutputs returns the value of property "AudioNode.numberOfOutputs".
//
// It returns ok=false if there is no such property.
func (this AudioNode) NumberOfOutputs() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioNodeNumberOfOutputs(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ChannelCount returns the value of property "AudioNode.channelCount".
//
// It returns ok=false if there is no such property.
func (this AudioNode) ChannelCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioNodeChannelCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChannelCount sets the value of property "AudioNode.channelCount" to val.
//
// It returns false if the property cannot be set.
func (this AudioNode) SetChannelCount(val uint32) bool {
	return js.True == bindings.SetAudioNodeChannelCount(
		this.ref,
		uint32(val),
	)
}

// ChannelCountMode returns the value of property "AudioNode.channelCountMode".
//
// It returns ok=false if there is no such property.
func (this AudioNode) ChannelCountMode() (ret ChannelCountMode, ok bool) {
	ok = js.True == bindings.GetAudioNodeChannelCountMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChannelCountMode sets the value of property "AudioNode.channelCountMode" to val.
//
// It returns false if the property cannot be set.
func (this AudioNode) SetChannelCountMode(val ChannelCountMode) bool {
	return js.True == bindings.SetAudioNodeChannelCountMode(
		this.ref,
		uint32(val),
	)
}

// ChannelInterpretation returns the value of property "AudioNode.channelInterpretation".
//
// It returns ok=false if there is no such property.
func (this AudioNode) ChannelInterpretation() (ret ChannelInterpretation, ok bool) {
	ok = js.True == bindings.GetAudioNodeChannelInterpretation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChannelInterpretation sets the value of property "AudioNode.channelInterpretation" to val.
//
// It returns false if the property cannot be set.
func (this AudioNode) SetChannelInterpretation(val ChannelInterpretation) bool {
	return js.True == bindings.SetAudioNodeChannelInterpretation(
		this.ref,
		uint32(val),
	)
}

// HasFuncConnect returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasFuncConnect() bool {
	return js.True == bindings.HasFuncAudioNodeConnect(
		this.ref,
	)
}

// FuncConnect returns the method "AudioNode.connect".
func (this AudioNode) FuncConnect() (fn js.Func[func(destinationNode AudioNode, output uint32, input uint32) AudioNode]) {
	bindings.FuncAudioNodeConnect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Connect calls the method "AudioNode.connect".
func (this AudioNode) Connect(destinationNode AudioNode, output uint32, input uint32) (ret AudioNode) {
	bindings.CallAudioNodeConnect(
		this.ref, js.Pointer(&ret),
		destinationNode.Ref(),
		uint32(output),
		uint32(input),
	)

	return
}

// TryConnect calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryConnect(destinationNode AudioNode, output uint32, input uint32) (ret AudioNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
		uint32(output),
		uint32(input),
	)

	return
}

// HasFuncConnect1 returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasFuncConnect1() bool {
	return js.True == bindings.HasFuncAudioNodeConnect1(
		this.ref,
	)
}

// FuncConnect1 returns the method "AudioNode.connect".
func (this AudioNode) FuncConnect1() (fn js.Func[func(destinationNode AudioNode, output uint32) AudioNode]) {
	bindings.FuncAudioNodeConnect1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Connect1 calls the method "AudioNode.connect".
func (this AudioNode) Connect1(destinationNode AudioNode, output uint32) (ret AudioNode) {
	bindings.CallAudioNodeConnect1(
		this.ref, js.Pointer(&ret),
		destinationNode.Ref(),
		uint32(output),
	)

	return
}

// TryConnect1 calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryConnect1(destinationNode AudioNode, output uint32) (ret AudioNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
		uint32(output),
	)

	return
}

// HasFuncConnect2 returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasFuncConnect2() bool {
	return js.True == bindings.HasFuncAudioNodeConnect2(
		this.ref,
	)
}

// FuncConnect2 returns the method "AudioNode.connect".
func (this AudioNode) FuncConnect2() (fn js.Func[func(destinationNode AudioNode) AudioNode]) {
	bindings.FuncAudioNodeConnect2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Connect2 calls the method "AudioNode.connect".
func (this AudioNode) Connect2(destinationNode AudioNode) (ret AudioNode) {
	bindings.CallAudioNodeConnect2(
		this.ref, js.Pointer(&ret),
		destinationNode.Ref(),
	)

	return
}

// TryConnect2 calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryConnect2(destinationNode AudioNode) (ret AudioNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
	)

	return
}

// HasFuncConnect3 returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasFuncConnect3() bool {
	return js.True == bindings.HasFuncAudioNodeConnect3(
		this.ref,
	)
}

// FuncConnect3 returns the method "AudioNode.connect".
func (this AudioNode) FuncConnect3() (fn js.Func[func(destinationParam AudioParam, output uint32)]) {
	bindings.FuncAudioNodeConnect3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Connect3 calls the method "AudioNode.connect".
func (this AudioNode) Connect3(destinationParam AudioParam, output uint32) (ret js.Void) {
	bindings.CallAudioNodeConnect3(
		this.ref, js.Pointer(&ret),
		destinationParam.Ref(),
		uint32(output),
	)

	return
}

// TryConnect3 calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryConnect3(destinationParam AudioParam, output uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationParam.Ref(),
		uint32(output),
	)

	return
}

// HasFuncConnect4 returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasFuncConnect4() bool {
	return js.True == bindings.HasFuncAudioNodeConnect4(
		this.ref,
	)
}

// FuncConnect4 returns the method "AudioNode.connect".
func (this AudioNode) FuncConnect4() (fn js.Func[func(destinationParam AudioParam)]) {
	bindings.FuncAudioNodeConnect4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Connect4 calls the method "AudioNode.connect".
func (this AudioNode) Connect4(destinationParam AudioParam) (ret js.Void) {
	bindings.CallAudioNodeConnect4(
		this.ref, js.Pointer(&ret),
		destinationParam.Ref(),
	)

	return
}

// TryConnect4 calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryConnect4(destinationParam AudioParam) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationParam.Ref(),
	)

	return
}

// HasFuncDisconnect returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncAudioNodeDisconnect(
		this.ref,
	)
}

// FuncDisconnect returns the method "AudioNode.disconnect".
func (this AudioNode) FuncDisconnect() (fn js.Func[func()]) {
	bindings.FuncAudioNodeDisconnect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect() (ret js.Void) {
	bindings.CallAudioNodeDisconnect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDisconnect1 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasFuncDisconnect1() bool {
	return js.True == bindings.HasFuncAudioNodeDisconnect1(
		this.ref,
	)
}

// FuncDisconnect1 returns the method "AudioNode.disconnect".
func (this AudioNode) FuncDisconnect1() (fn js.Func[func(output uint32)]) {
	bindings.FuncAudioNodeDisconnect1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect1 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect1(output uint32) (ret js.Void) {
	bindings.CallAudioNodeDisconnect1(
		this.ref, js.Pointer(&ret),
		uint32(output),
	)

	return
}

// TryDisconnect1 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryDisconnect1(output uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(output),
	)

	return
}

// HasFuncDisconnect2 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasFuncDisconnect2() bool {
	return js.True == bindings.HasFuncAudioNodeDisconnect2(
		this.ref,
	)
}

// FuncDisconnect2 returns the method "AudioNode.disconnect".
func (this AudioNode) FuncDisconnect2() (fn js.Func[func(destinationNode AudioNode)]) {
	bindings.FuncAudioNodeDisconnect2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect2 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect2(destinationNode AudioNode) (ret js.Void) {
	bindings.CallAudioNodeDisconnect2(
		this.ref, js.Pointer(&ret),
		destinationNode.Ref(),
	)

	return
}

// TryDisconnect2 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryDisconnect2(destinationNode AudioNode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
	)

	return
}

// HasFuncDisconnect3 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasFuncDisconnect3() bool {
	return js.True == bindings.HasFuncAudioNodeDisconnect3(
		this.ref,
	)
}

// FuncDisconnect3 returns the method "AudioNode.disconnect".
func (this AudioNode) FuncDisconnect3() (fn js.Func[func(destinationNode AudioNode, output uint32)]) {
	bindings.FuncAudioNodeDisconnect3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect3 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect3(destinationNode AudioNode, output uint32) (ret js.Void) {
	bindings.CallAudioNodeDisconnect3(
		this.ref, js.Pointer(&ret),
		destinationNode.Ref(),
		uint32(output),
	)

	return
}

// TryDisconnect3 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryDisconnect3(destinationNode AudioNode, output uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
		uint32(output),
	)

	return
}

// HasFuncDisconnect4 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasFuncDisconnect4() bool {
	return js.True == bindings.HasFuncAudioNodeDisconnect4(
		this.ref,
	)
}

// FuncDisconnect4 returns the method "AudioNode.disconnect".
func (this AudioNode) FuncDisconnect4() (fn js.Func[func(destinationNode AudioNode, output uint32, input uint32)]) {
	bindings.FuncAudioNodeDisconnect4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect4 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect4(destinationNode AudioNode, output uint32, input uint32) (ret js.Void) {
	bindings.CallAudioNodeDisconnect4(
		this.ref, js.Pointer(&ret),
		destinationNode.Ref(),
		uint32(output),
		uint32(input),
	)

	return
}

// TryDisconnect4 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryDisconnect4(destinationNode AudioNode, output uint32, input uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
		uint32(output),
		uint32(input),
	)

	return
}

// HasFuncDisconnect5 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasFuncDisconnect5() bool {
	return js.True == bindings.HasFuncAudioNodeDisconnect5(
		this.ref,
	)
}

// FuncDisconnect5 returns the method "AudioNode.disconnect".
func (this AudioNode) FuncDisconnect5() (fn js.Func[func(destinationParam AudioParam)]) {
	bindings.FuncAudioNodeDisconnect5(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect5 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect5(destinationParam AudioParam) (ret js.Void) {
	bindings.CallAudioNodeDisconnect5(
		this.ref, js.Pointer(&ret),
		destinationParam.Ref(),
	)

	return
}

// TryDisconnect5 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryDisconnect5(destinationParam AudioParam) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect5(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationParam.Ref(),
	)

	return
}

// HasFuncDisconnect6 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasFuncDisconnect6() bool {
	return js.True == bindings.HasFuncAudioNodeDisconnect6(
		this.ref,
	)
}

// FuncDisconnect6 returns the method "AudioNode.disconnect".
func (this AudioNode) FuncDisconnect6() (fn js.Func[func(destinationParam AudioParam, output uint32)]) {
	bindings.FuncAudioNodeDisconnect6(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect6 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect6(destinationParam AudioParam, output uint32) (ret js.Void) {
	bindings.CallAudioNodeDisconnect6(
		this.ref, js.Pointer(&ret),
		destinationParam.Ref(),
		uint32(output),
	)

	return
}

// TryDisconnect6 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioNode) TryDisconnect6(destinationParam AudioParam, output uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect6(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		destinationParam.Ref(),
		uint32(output),
	)

	return
}

type AudioOutputOptions struct {
	// DeviceId is "AudioOutputOptions.deviceId"
	//
	// Optional, defaults to "".
	DeviceId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioOutputOptions with all fields set.
func (p AudioOutputOptions) FromRef(ref js.Ref) AudioOutputOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioOutputOptions in the application heap.
func (p AudioOutputOptions) New() js.Ref {
	return bindings.AudioOutputOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioOutputOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioOutputOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioOutputOptions) Update(ref js.Ref) {
	bindings.AudioOutputOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioOutputOptions) FreeMembers(recursive bool) {
	js.Free(
		p.DeviceId.Ref(),
	)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
}

type AudioParamDescriptor struct {
	// Name is "AudioParamDescriptor.name"
	//
	// Required
	Name js.String
	// DefaultValue is "AudioParamDescriptor.defaultValue"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_DefaultValue MUST be set to true to make this field effective.
	DefaultValue float32
	// MinValue is "AudioParamDescriptor.minValue"
	//
	// Optional, defaults to -3.4028235e38.
	//
	// NOTE: FFI_USE_MinValue MUST be set to true to make this field effective.
	MinValue float32
	// MaxValue is "AudioParamDescriptor.maxValue"
	//
	// Optional, defaults to 3.4028235e38.
	//
	// NOTE: FFI_USE_MaxValue MUST be set to true to make this field effective.
	MaxValue float32
	// AutomationRate is "AudioParamDescriptor.automationRate"
	//
	// Optional, defaults to "a-rate".
	AutomationRate AutomationRate

	FFI_USE_DefaultValue bool // for DefaultValue.
	FFI_USE_MinValue     bool // for MinValue.
	FFI_USE_MaxValue     bool // for MaxValue.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioParamDescriptor with all fields set.
func (p AudioParamDescriptor) FromRef(ref js.Ref) AudioParamDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioParamDescriptor in the application heap.
func (p AudioParamDescriptor) New() js.Ref {
	return bindings.AudioParamDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioParamDescriptor) UpdateFrom(ref js.Ref) {
	bindings.AudioParamDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioParamDescriptor) Update(ref js.Ref) {
	bindings.AudioParamDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioParamDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type AudioParamMap struct {
	ref js.Ref
}

func (this AudioParamMap) Once() AudioParamMap {
	this.ref.Once()
	return this
}

func (this AudioParamMap) Ref() js.Ref {
	return this.ref
}

func (this AudioParamMap) FromRef(ref js.Ref) AudioParamMap {
	this.ref = ref
	return this
}

func (this AudioParamMap) Free() {
	this.ref.Free()
}

type AudioProcessingEventInit struct {
	// PlaybackTime is "AudioProcessingEventInit.playbackTime"
	//
	// Required
	PlaybackTime float64
	// InputBuffer is "AudioProcessingEventInit.inputBuffer"
	//
	// Required
	InputBuffer AudioBuffer
	// OutputBuffer is "AudioProcessingEventInit.outputBuffer"
	//
	// Required
	OutputBuffer AudioBuffer
	// Bubbles is "AudioProcessingEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "AudioProcessingEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "AudioProcessingEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioProcessingEventInit with all fields set.
func (p AudioProcessingEventInit) FromRef(ref js.Ref) AudioProcessingEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioProcessingEventInit in the application heap.
func (p AudioProcessingEventInit) New() js.Ref {
	return bindings.AudioProcessingEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioProcessingEventInit) UpdateFrom(ref js.Ref) {
	bindings.AudioProcessingEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioProcessingEventInit) Update(ref js.Ref) {
	bindings.AudioProcessingEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioProcessingEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.InputBuffer.Ref(),
		p.OutputBuffer.Ref(),
	)
	p.InputBuffer = p.InputBuffer.FromRef(js.Undefined)
	p.OutputBuffer = p.OutputBuffer.FromRef(js.Undefined)
}

func NewAudioProcessingEvent(typ js.String, eventInitDict AudioProcessingEventInit) (ret AudioProcessingEvent) {
	ret.ref = bindings.NewAudioProcessingEventByAudioProcessingEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type AudioProcessingEvent struct {
	Event
}

func (this AudioProcessingEvent) Once() AudioProcessingEvent {
	this.ref.Once()
	return this
}

func (this AudioProcessingEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this AudioProcessingEvent) FromRef(ref js.Ref) AudioProcessingEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this AudioProcessingEvent) Free() {
	this.ref.Free()
}

// PlaybackTime returns the value of property "AudioProcessingEvent.playbackTime".
//
// It returns ok=false if there is no such property.
func (this AudioProcessingEvent) PlaybackTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioProcessingEventPlaybackTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// InputBuffer returns the value of property "AudioProcessingEvent.inputBuffer".
//
// It returns ok=false if there is no such property.
func (this AudioProcessingEvent) InputBuffer() (ret AudioBuffer, ok bool) {
	ok = js.True == bindings.GetAudioProcessingEventInputBuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OutputBuffer returns the value of property "AudioProcessingEvent.outputBuffer".
//
// It returns ok=false if there is no such property.
func (this AudioProcessingEvent) OutputBuffer() (ret AudioBuffer, ok bool) {
	ok = js.True == bindings.GetAudioProcessingEventOutputBuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

type AudioRenderCapacityEventInit struct {
	// Timestamp is "AudioRenderCapacityEventInit.timestamp"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp float64
	// AverageLoad is "AudioRenderCapacityEventInit.averageLoad"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_AverageLoad MUST be set to true to make this field effective.
	AverageLoad float64
	// PeakLoad is "AudioRenderCapacityEventInit.peakLoad"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_PeakLoad MUST be set to true to make this field effective.
	PeakLoad float64
	// UnderrunRatio is "AudioRenderCapacityEventInit.underrunRatio"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_UnderrunRatio MUST be set to true to make this field effective.
	UnderrunRatio float64
	// Bubbles is "AudioRenderCapacityEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "AudioRenderCapacityEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "AudioRenderCapacityEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Timestamp     bool // for Timestamp.
	FFI_USE_AverageLoad   bool // for AverageLoad.
	FFI_USE_PeakLoad      bool // for PeakLoad.
	FFI_USE_UnderrunRatio bool // for UnderrunRatio.
	FFI_USE_Bubbles       bool // for Bubbles.
	FFI_USE_Cancelable    bool // for Cancelable.
	FFI_USE_Composed      bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioRenderCapacityEventInit with all fields set.
func (p AudioRenderCapacityEventInit) FromRef(ref js.Ref) AudioRenderCapacityEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioRenderCapacityEventInit in the application heap.
func (p AudioRenderCapacityEventInit) New() js.Ref {
	return bindings.AudioRenderCapacityEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioRenderCapacityEventInit) UpdateFrom(ref js.Ref) {
	bindings.AudioRenderCapacityEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioRenderCapacityEventInit) Update(ref js.Ref) {
	bindings.AudioRenderCapacityEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioRenderCapacityEventInit) FreeMembers(recursive bool) {
}

func NewAudioRenderCapacityEvent(typ js.String, eventInitDict AudioRenderCapacityEventInit) (ret AudioRenderCapacityEvent) {
	ret.ref = bindings.NewAudioRenderCapacityEventByAudioRenderCapacityEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewAudioRenderCapacityEventByAudioRenderCapacityEvent1(typ js.String) (ret AudioRenderCapacityEvent) {
	ret.ref = bindings.NewAudioRenderCapacityEventByAudioRenderCapacityEvent1(
		typ.Ref())
	return
}

type AudioRenderCapacityEvent struct {
	Event
}

func (this AudioRenderCapacityEvent) Once() AudioRenderCapacityEvent {
	this.ref.Once()
	return this
}

func (this AudioRenderCapacityEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this AudioRenderCapacityEvent) FromRef(ref js.Ref) AudioRenderCapacityEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this AudioRenderCapacityEvent) Free() {
	this.ref.Free()
}

// Timestamp returns the value of property "AudioRenderCapacityEvent.timestamp".
//
// It returns ok=false if there is no such property.
func (this AudioRenderCapacityEvent) Timestamp() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioRenderCapacityEventTimestamp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AverageLoad returns the value of property "AudioRenderCapacityEvent.averageLoad".
//
// It returns ok=false if there is no such property.
func (this AudioRenderCapacityEvent) AverageLoad() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioRenderCapacityEventAverageLoad(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PeakLoad returns the value of property "AudioRenderCapacityEvent.peakLoad".
//
// It returns ok=false if there is no such property.
func (this AudioRenderCapacityEvent) PeakLoad() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioRenderCapacityEventPeakLoad(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UnderrunRatio returns the value of property "AudioRenderCapacityEvent.underrunRatio".
//
// It returns ok=false if there is no such property.
func (this AudioRenderCapacityEvent) UnderrunRatio() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioRenderCapacityEventUnderrunRatio(
		this.ref, js.Pointer(&ret),
	)
	return
}

type AudioScheduledSourceNode struct {
	AudioNode
}

func (this AudioScheduledSourceNode) Once() AudioScheduledSourceNode {
	this.ref.Once()
	return this
}

func (this AudioScheduledSourceNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this AudioScheduledSourceNode) FromRef(ref js.Ref) AudioScheduledSourceNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this AudioScheduledSourceNode) Free() {
	this.ref.Free()
}

// HasFuncStart returns true if the method "AudioScheduledSourceNode.start" exists.
func (this AudioScheduledSourceNode) HasFuncStart() bool {
	return js.True == bindings.HasFuncAudioScheduledSourceNodeStart(
		this.ref,
	)
}

// FuncStart returns the method "AudioScheduledSourceNode.start".
func (this AudioScheduledSourceNode) FuncStart() (fn js.Func[func(when float64)]) {
	bindings.FuncAudioScheduledSourceNodeStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "AudioScheduledSourceNode.start".
func (this AudioScheduledSourceNode) Start(when float64) (ret js.Void) {
	bindings.CallAudioScheduledSourceNodeStart(
		this.ref, js.Pointer(&ret),
		float64(when),
	)

	return
}

// TryStart calls the method "AudioScheduledSourceNode.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioScheduledSourceNode) TryStart(when float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioScheduledSourceNodeStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(when),
	)

	return
}

// HasFuncStart1 returns true if the method "AudioScheduledSourceNode.start" exists.
func (this AudioScheduledSourceNode) HasFuncStart1() bool {
	return js.True == bindings.HasFuncAudioScheduledSourceNodeStart1(
		this.ref,
	)
}

// FuncStart1 returns the method "AudioScheduledSourceNode.start".
func (this AudioScheduledSourceNode) FuncStart1() (fn js.Func[func()]) {
	bindings.FuncAudioScheduledSourceNodeStart1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start1 calls the method "AudioScheduledSourceNode.start".
func (this AudioScheduledSourceNode) Start1() (ret js.Void) {
	bindings.CallAudioScheduledSourceNodeStart1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart1 calls the method "AudioScheduledSourceNode.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioScheduledSourceNode) TryStart1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioScheduledSourceNodeStart1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStop returns true if the method "AudioScheduledSourceNode.stop" exists.
func (this AudioScheduledSourceNode) HasFuncStop() bool {
	return js.True == bindings.HasFuncAudioScheduledSourceNodeStop(
		this.ref,
	)
}

// FuncStop returns the method "AudioScheduledSourceNode.stop".
func (this AudioScheduledSourceNode) FuncStop() (fn js.Func[func(when float64)]) {
	bindings.FuncAudioScheduledSourceNodeStop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stop calls the method "AudioScheduledSourceNode.stop".
func (this AudioScheduledSourceNode) Stop(when float64) (ret js.Void) {
	bindings.CallAudioScheduledSourceNodeStop(
		this.ref, js.Pointer(&ret),
		float64(when),
	)

	return
}

// TryStop calls the method "AudioScheduledSourceNode.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioScheduledSourceNode) TryStop(when float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioScheduledSourceNodeStop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(when),
	)

	return
}

// HasFuncStop1 returns true if the method "AudioScheduledSourceNode.stop" exists.
func (this AudioScheduledSourceNode) HasFuncStop1() bool {
	return js.True == bindings.HasFuncAudioScheduledSourceNodeStop1(
		this.ref,
	)
}

// FuncStop1 returns the method "AudioScheduledSourceNode.stop".
func (this AudioScheduledSourceNode) FuncStop1() (fn js.Func[func()]) {
	bindings.FuncAudioScheduledSourceNodeStop1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stop1 calls the method "AudioScheduledSourceNode.stop".
func (this AudioScheduledSourceNode) Stop1() (ret js.Void) {
	bindings.CallAudioScheduledSourceNodeStop1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStop1 calls the method "AudioScheduledSourceNode.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioScheduledSourceNode) TryStop1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioScheduledSourceNodeStop1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type AudioWorkletProcessorConstructorFunc func(this js.Ref, options js.Object) js.Ref

func (fn AudioWorkletProcessorConstructorFunc) Register() js.Func[func(options js.Object) AudioWorkletProcessor] {
	return js.RegisterCallback[func(options js.Object) AudioWorkletProcessor](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AudioWorkletProcessorConstructorFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AudioWorkletProcessorConstructor[T any] struct {
	Fn  func(arg T, this js.Ref, options js.Object) js.Ref
	Arg T
}

func (cb *AudioWorkletProcessorConstructor[T]) Register() js.Func[func(options js.Object) AudioWorkletProcessor] {
	return js.RegisterCallback[func(options js.Object) AudioWorkletProcessor](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AudioWorkletProcessorConstructor[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AudioWorkletProcessor struct {
	ref js.Ref
}

func (this AudioWorkletProcessor) Once() AudioWorkletProcessor {
	this.ref.Once()
	return this
}

func (this AudioWorkletProcessor) Ref() js.Ref {
	return this.ref
}

func (this AudioWorkletProcessor) FromRef(ref js.Ref) AudioWorkletProcessor {
	this.ref = ref
	return this
}

func (this AudioWorkletProcessor) Free() {
	this.ref.Free()
}

// Port returns the value of property "AudioWorkletProcessor.port".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletProcessor) Port() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetAudioWorkletProcessorPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

type AudioWorkletGlobalScope struct {
	WorkletGlobalScope
}

func (this AudioWorkletGlobalScope) Once() AudioWorkletGlobalScope {
	this.ref.Once()
	return this
}

func (this AudioWorkletGlobalScope) Ref() js.Ref {
	return this.WorkletGlobalScope.Ref()
}

func (this AudioWorkletGlobalScope) FromRef(ref js.Ref) AudioWorkletGlobalScope {
	this.WorkletGlobalScope = this.WorkletGlobalScope.FromRef(ref)
	return this
}

func (this AudioWorkletGlobalScope) Free() {
	this.ref.Free()
}

// CurrentFrame returns the value of property "AudioWorkletGlobalScope.currentFrame".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletGlobalScope) CurrentFrame() (ret uint64, ok bool) {
	ok = js.True == bindings.GetAudioWorkletGlobalScopeCurrentFrame(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CurrentTime returns the value of property "AudioWorkletGlobalScope.currentTime".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletGlobalScope) CurrentTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioWorkletGlobalScopeCurrentTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SampleRate returns the value of property "AudioWorkletGlobalScope.sampleRate".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletGlobalScope) SampleRate() (ret float32, ok bool) {
	ok = js.True == bindings.GetAudioWorkletGlobalScopeSampleRate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Port returns the value of property "AudioWorkletGlobalScope.port".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletGlobalScope) Port() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetAudioWorkletGlobalScopePort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRegisterProcessor returns true if the method "AudioWorkletGlobalScope.registerProcessor" exists.
func (this AudioWorkletGlobalScope) HasFuncRegisterProcessor() bool {
	return js.True == bindings.HasFuncAudioWorkletGlobalScopeRegisterProcessor(
		this.ref,
	)
}

// FuncRegisterProcessor returns the method "AudioWorkletGlobalScope.registerProcessor".
func (this AudioWorkletGlobalScope) FuncRegisterProcessor() (fn js.Func[func(name js.String, processorCtor js.Func[func(options js.Object) AudioWorkletProcessor])]) {
	bindings.FuncAudioWorkletGlobalScopeRegisterProcessor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RegisterProcessor calls the method "AudioWorkletGlobalScope.registerProcessor".
func (this AudioWorkletGlobalScope) RegisterProcessor(name js.String, processorCtor js.Func[func(options js.Object) AudioWorkletProcessor]) (ret js.Void) {
	bindings.CallAudioWorkletGlobalScopeRegisterProcessor(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		processorCtor.Ref(),
	)

	return
}

// TryRegisterProcessor calls the method "AudioWorkletGlobalScope.registerProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioWorkletGlobalScope) TryRegisterProcessor(name js.String, processorCtor js.Func[func(options js.Object) AudioWorkletProcessor]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioWorkletGlobalScopeRegisterProcessor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		processorCtor.Ref(),
	)

	return
}

type AudioWorkletNodeOptions struct {
	// NumberOfInputs is "AudioWorkletNodeOptions.numberOfInputs"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_NumberOfInputs MUST be set to true to make this field effective.
	NumberOfInputs uint32
	// NumberOfOutputs is "AudioWorkletNodeOptions.numberOfOutputs"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_NumberOfOutputs MUST be set to true to make this field effective.
	NumberOfOutputs uint32
	// OutputChannelCount is "AudioWorkletNodeOptions.outputChannelCount"
	//
	// Optional
	OutputChannelCount js.Array[uint32]
	// ParameterData is "AudioWorkletNodeOptions.parameterData"
	//
	// Optional
	ParameterData js.Record[float64]
	// ProcessorOptions is "AudioWorkletNodeOptions.processorOptions"
	//
	// Optional
	ProcessorOptions js.Object
	// ChannelCount is "AudioWorkletNodeOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
	ChannelCount uint32
	// ChannelCountMode is "AudioWorkletNodeOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "AudioWorkletNodeOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_NumberOfInputs  bool // for NumberOfInputs.
	FFI_USE_NumberOfOutputs bool // for NumberOfOutputs.
	FFI_USE_ChannelCount    bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioWorkletNodeOptions with all fields set.
func (p AudioWorkletNodeOptions) FromRef(ref js.Ref) AudioWorkletNodeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioWorkletNodeOptions in the application heap.
func (p AudioWorkletNodeOptions) New() js.Ref {
	return bindings.AudioWorkletNodeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioWorkletNodeOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioWorkletNodeOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioWorkletNodeOptions) Update(ref js.Ref) {
	bindings.AudioWorkletNodeOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioWorkletNodeOptions) FreeMembers(recursive bool) {
	js.Free(
		p.OutputChannelCount.Ref(),
		p.ParameterData.Ref(),
		p.ProcessorOptions.Ref(),
	)
	p.OutputChannelCount = p.OutputChannelCount.FromRef(js.Undefined)
	p.ParameterData = p.ParameterData.FromRef(js.Undefined)
	p.ProcessorOptions = p.ProcessorOptions.FromRef(js.Undefined)
}

func NewAudioWorkletNode(context BaseAudioContext, name js.String, options AudioWorkletNodeOptions) (ret AudioWorkletNode) {
	ret.ref = bindings.NewAudioWorkletNodeByAudioWorkletNode(
		context.Ref(),
		name.Ref(),
		js.Pointer(&options))
	return
}

func NewAudioWorkletNodeByAudioWorkletNode1(context BaseAudioContext, name js.String) (ret AudioWorkletNode) {
	ret.ref = bindings.NewAudioWorkletNodeByAudioWorkletNode1(
		context.Ref(),
		name.Ref())
	return
}

type AudioWorkletNode struct {
	AudioNode
}

func (this AudioWorkletNode) Once() AudioWorkletNode {
	this.ref.Once()
	return this
}

func (this AudioWorkletNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this AudioWorkletNode) FromRef(ref js.Ref) AudioWorkletNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this AudioWorkletNode) Free() {
	this.ref.Free()
}

// Parameters returns the value of property "AudioWorkletNode.parameters".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletNode) Parameters() (ret AudioParamMap, ok bool) {
	ok = js.True == bindings.GetAudioWorkletNodeParameters(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Port returns the value of property "AudioWorkletNode.port".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletNode) Port() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetAudioWorkletNodePort(
		this.ref, js.Pointer(&ret),
	)
	return
}

type AudioWorkletProcessCallbackFunc func(this js.Ref, inputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], outputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], parameters js.Object) js.Ref

func (fn AudioWorkletProcessCallbackFunc) Register() js.Func[func(inputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], outputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], parameters js.Object) bool] {
	return js.RegisterCallback[func(inputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], outputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], parameters js.Object) bool](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AudioWorkletProcessCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]]{}.FromRef(args[0+1]),
		js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]]{}.FromRef(args[1+1]),
		js.Object{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AudioWorkletProcessCallback[T any] struct {
	Fn  func(arg T, this js.Ref, inputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], outputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], parameters js.Object) js.Ref
	Arg T
}

func (cb *AudioWorkletProcessCallback[T]) Register() js.Func[func(inputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], outputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], parameters js.Object) bool] {
	return js.RegisterCallback[func(inputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], outputs js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]], parameters js.Object) bool](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AudioWorkletProcessCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]]{}.FromRef(args[0+1]),
		js.FrozenArray[js.FrozenArray[js.TypedArray[float32]]]{}.FromRef(args[1+1]),
		js.Object{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AuthenticationExtensionsPRFValues struct {
	// First is "AuthenticationExtensionsPRFValues.first"
	//
	// Required
	First BufferSource
	// Second is "AuthenticationExtensionsPRFValues.second"
	//
	// Optional
	Second BufferSource

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsPRFValues with all fields set.
func (p AuthenticationExtensionsPRFValues) FromRef(ref js.Ref) AuthenticationExtensionsPRFValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsPRFValues in the application heap.
func (p AuthenticationExtensionsPRFValues) New() js.Ref {
	return bindings.AuthenticationExtensionsPRFValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsPRFValues) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsPRFValues) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsPRFValues) FreeMembers(recursive bool) {
	js.Free(
		p.First.Ref(),
		p.Second.Ref(),
	)
	p.First = p.First.FromRef(js.Undefined)
	p.Second = p.Second.FromRef(js.Undefined)
}

type AuthenticationExtensionsPRFInputs struct {
	// Eval is "AuthenticationExtensionsPRFInputs.eval"
	//
	// Optional
	//
	// NOTE: Eval.FFI_USE MUST be set to true to get Eval used.
	Eval AuthenticationExtensionsPRFValues
	// EvalByCredential is "AuthenticationExtensionsPRFInputs.evalByCredential"
	//
	// Optional
	EvalByCredential js.Record[AuthenticationExtensionsPRFValues]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsPRFInputs with all fields set.
func (p AuthenticationExtensionsPRFInputs) FromRef(ref js.Ref) AuthenticationExtensionsPRFInputs {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsPRFInputs in the application heap.
func (p AuthenticationExtensionsPRFInputs) New() js.Ref {
	return bindings.AuthenticationExtensionsPRFInputsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsPRFInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFInputsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsPRFInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFInputsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsPRFInputs) FreeMembers(recursive bool) {
	js.Free(
		p.EvalByCredential.Ref(),
	)
	p.EvalByCredential = p.EvalByCredential.FromRef(js.Undefined)
	if recursive {
		p.Eval.FreeMembers(true)
	}
}

type PaymentCurrencyAmount struct {
	// Currency is "PaymentCurrencyAmount.currency"
	//
	// Required
	Currency js.String
	// Value is "PaymentCurrencyAmount.value"
	//
	// Required
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentCurrencyAmount with all fields set.
func (p PaymentCurrencyAmount) FromRef(ref js.Ref) PaymentCurrencyAmount {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentCurrencyAmount in the application heap.
func (p PaymentCurrencyAmount) New() js.Ref {
	return bindings.PaymentCurrencyAmountJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PaymentCurrencyAmount) UpdateFrom(ref js.Ref) {
	bindings.PaymentCurrencyAmountJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PaymentCurrencyAmount) Update(ref js.Ref) {
	bindings.PaymentCurrencyAmountJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PaymentCurrencyAmount) FreeMembers(recursive bool) {
	js.Free(
		p.Currency.Ref(),
		p.Value.Ref(),
	)
	p.Currency = p.Currency.FromRef(js.Undefined)
	p.Value = p.Value.FromRef(js.Undefined)
}

type PaymentCredentialInstrument struct {
	// DisplayName is "PaymentCredentialInstrument.displayName"
	//
	// Required
	DisplayName js.String
	// Icon is "PaymentCredentialInstrument.icon"
	//
	// Required
	Icon js.String
	// IconMustBeShown is "PaymentCredentialInstrument.iconMustBeShown"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_IconMustBeShown MUST be set to true to make this field effective.
	IconMustBeShown bool

	FFI_USE_IconMustBeShown bool // for IconMustBeShown.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentCredentialInstrument with all fields set.
func (p PaymentCredentialInstrument) FromRef(ref js.Ref) PaymentCredentialInstrument {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentCredentialInstrument in the application heap.
func (p PaymentCredentialInstrument) New() js.Ref {
	return bindings.PaymentCredentialInstrumentJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PaymentCredentialInstrument) UpdateFrom(ref js.Ref) {
	bindings.PaymentCredentialInstrumentJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PaymentCredentialInstrument) Update(ref js.Ref) {
	bindings.PaymentCredentialInstrumentJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PaymentCredentialInstrument) FreeMembers(recursive bool) {
	js.Free(
		p.DisplayName.Ref(),
		p.Icon.Ref(),
	)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
	p.Icon = p.Icon.FromRef(js.Undefined)
}

type AuthenticationExtensionsPaymentInputs struct {
	// IsPayment is "AuthenticationExtensionsPaymentInputs.isPayment"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsPayment MUST be set to true to make this field effective.
	IsPayment bool
	// RpId is "AuthenticationExtensionsPaymentInputs.rpId"
	//
	// Optional
	RpId js.String
	// TopOrigin is "AuthenticationExtensionsPaymentInputs.topOrigin"
	//
	// Optional
	TopOrigin js.String
	// PayeeName is "AuthenticationExtensionsPaymentInputs.payeeName"
	//
	// Optional
	PayeeName js.String
	// PayeeOrigin is "AuthenticationExtensionsPaymentInputs.payeeOrigin"
	//
	// Optional
	PayeeOrigin js.String
	// Total is "AuthenticationExtensionsPaymentInputs.total"
	//
	// Optional
	//
	// NOTE: Total.FFI_USE MUST be set to true to get Total used.
	Total PaymentCurrencyAmount
	// Instrument is "AuthenticationExtensionsPaymentInputs.instrument"
	//
	// Optional
	//
	// NOTE: Instrument.FFI_USE MUST be set to true to get Instrument used.
	Instrument PaymentCredentialInstrument

	FFI_USE_IsPayment bool // for IsPayment.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsPaymentInputs with all fields set.
func (p AuthenticationExtensionsPaymentInputs) FromRef(ref js.Ref) AuthenticationExtensionsPaymentInputs {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsPaymentInputs in the application heap.
func (p AuthenticationExtensionsPaymentInputs) New() js.Ref {
	return bindings.AuthenticationExtensionsPaymentInputsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsPaymentInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsPaymentInputsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsPaymentInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsPaymentInputsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsPaymentInputs) FreeMembers(recursive bool) {
	js.Free(
		p.RpId.Ref(),
		p.TopOrigin.Ref(),
		p.PayeeName.Ref(),
		p.PayeeOrigin.Ref(),
	)
	p.RpId = p.RpId.FromRef(js.Undefined)
	p.TopOrigin = p.TopOrigin.FromRef(js.Undefined)
	p.PayeeName = p.PayeeName.FromRef(js.Undefined)
	p.PayeeOrigin = p.PayeeOrigin.FromRef(js.Undefined)
	if recursive {
		p.Total.FreeMembers(true)
		p.Instrument.FreeMembers(true)
	}
}

type HMACGetSecretInput struct {
	// Salt1 is "HMACGetSecretInput.salt1"
	//
	// Required
	Salt1 js.ArrayBuffer
	// Salt2 is "HMACGetSecretInput.salt2"
	//
	// Optional
	Salt2 js.ArrayBuffer

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HMACGetSecretInput with all fields set.
func (p HMACGetSecretInput) FromRef(ref js.Ref) HMACGetSecretInput {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HMACGetSecretInput in the application heap.
func (p HMACGetSecretInput) New() js.Ref {
	return bindings.HMACGetSecretInputJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HMACGetSecretInput) UpdateFrom(ref js.Ref) {
	bindings.HMACGetSecretInputJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HMACGetSecretInput) Update(ref js.Ref) {
	bindings.HMACGetSecretInputJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HMACGetSecretInput) FreeMembers(recursive bool) {
	js.Free(
		p.Salt1.Ref(),
		p.Salt2.Ref(),
	)
	p.Salt1 = p.Salt1.FromRef(js.Undefined)
	p.Salt2 = p.Salt2.FromRef(js.Undefined)
}

type AuthenticationExtensionsLargeBlobInputs struct {
	// Support is "AuthenticationExtensionsLargeBlobInputs.support"
	//
	// Optional
	Support js.String
	// Read is "AuthenticationExtensionsLargeBlobInputs.read"
	//
	// Optional
	//
	// NOTE: FFI_USE_Read MUST be set to true to make this field effective.
	Read bool
	// Write is "AuthenticationExtensionsLargeBlobInputs.write"
	//
	// Optional
	Write BufferSource

	FFI_USE_Read bool // for Read.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsLargeBlobInputs with all fields set.
func (p AuthenticationExtensionsLargeBlobInputs) FromRef(ref js.Ref) AuthenticationExtensionsLargeBlobInputs {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsLargeBlobInputs in the application heap.
func (p AuthenticationExtensionsLargeBlobInputs) New() js.Ref {
	return bindings.AuthenticationExtensionsLargeBlobInputsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsLargeBlobInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsLargeBlobInputsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsLargeBlobInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsLargeBlobInputsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsLargeBlobInputs) FreeMembers(recursive bool) {
	js.Free(
		p.Support.Ref(),
		p.Write.Ref(),
	)
	p.Support = p.Support.FromRef(js.Undefined)
	p.Write = p.Write.FromRef(js.Undefined)
}

type AuthenticationExtensionsDevicePublicKeyInputs struct {
	// Attestation is "AuthenticationExtensionsDevicePublicKeyInputs.attestation"
	//
	// Optional, defaults to "none".
	Attestation js.String
	// AttestationFormats is "AuthenticationExtensionsDevicePublicKeyInputs.attestationFormats"
	//
	// Optional, defaults to [].
	AttestationFormats js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsDevicePublicKeyInputs with all fields set.
func (p AuthenticationExtensionsDevicePublicKeyInputs) FromRef(ref js.Ref) AuthenticationExtensionsDevicePublicKeyInputs {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsDevicePublicKeyInputs in the application heap.
func (p AuthenticationExtensionsDevicePublicKeyInputs) New() js.Ref {
	return bindings.AuthenticationExtensionsDevicePublicKeyInputsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsDevicePublicKeyInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsDevicePublicKeyInputsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsDevicePublicKeyInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsDevicePublicKeyInputsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsDevicePublicKeyInputs) FreeMembers(recursive bool) {
	js.Free(
		p.Attestation.Ref(),
		p.AttestationFormats.Ref(),
	)
	p.Attestation = p.Attestation.FromRef(js.Undefined)
	p.AttestationFormats = p.AttestationFormats.FromRef(js.Undefined)
}

type AuthenticationExtensionsClientInputs struct {
	// MinPinLength is "AuthenticationExtensionsClientInputs.minPinLength"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinPinLength MUST be set to true to make this field effective.
	MinPinLength bool
	// Prf is "AuthenticationExtensionsClientInputs.prf"
	//
	// Optional
	//
	// NOTE: Prf.FFI_USE MUST be set to true to get Prf used.
	Prf AuthenticationExtensionsPRFInputs
	// Appid is "AuthenticationExtensionsClientInputs.appid"
	//
	// Optional
	Appid js.String
	// AppidExclude is "AuthenticationExtensionsClientInputs.appidExclude"
	//
	// Optional
	AppidExclude js.String
	// GetCredBlob is "AuthenticationExtensionsClientInputs.getCredBlob"
	//
	// Optional
	//
	// NOTE: FFI_USE_GetCredBlob MUST be set to true to make this field effective.
	GetCredBlob bool
	// HmacCreateSecret is "AuthenticationExtensionsClientInputs.hmacCreateSecret"
	//
	// Optional
	//
	// NOTE: FFI_USE_HmacCreateSecret MUST be set to true to make this field effective.
	HmacCreateSecret bool
	// CredProps is "AuthenticationExtensionsClientInputs.credProps"
	//
	// Optional
	//
	// NOTE: FFI_USE_CredProps MUST be set to true to make this field effective.
	CredProps bool
	// Payment is "AuthenticationExtensionsClientInputs.payment"
	//
	// Optional
	//
	// NOTE: Payment.FFI_USE MUST be set to true to get Payment used.
	Payment AuthenticationExtensionsPaymentInputs
	// HmacGetSecret is "AuthenticationExtensionsClientInputs.hmacGetSecret"
	//
	// Optional
	//
	// NOTE: HmacGetSecret.FFI_USE MUST be set to true to get HmacGetSecret used.
	HmacGetSecret HMACGetSecretInput
	// LargeBlob is "AuthenticationExtensionsClientInputs.largeBlob"
	//
	// Optional
	//
	// NOTE: LargeBlob.FFI_USE MUST be set to true to get LargeBlob used.
	LargeBlob AuthenticationExtensionsLargeBlobInputs
	// CredBlob is "AuthenticationExtensionsClientInputs.credBlob"
	//
	// Optional
	CredBlob js.ArrayBuffer
	// CredentialProtectionPolicy is "AuthenticationExtensionsClientInputs.credentialProtectionPolicy"
	//
	// Optional
	CredentialProtectionPolicy js.String
	// EnforceCredentialProtectionPolicy is "AuthenticationExtensionsClientInputs.enforceCredentialProtectionPolicy"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_EnforceCredentialProtectionPolicy MUST be set to true to make this field effective.
	EnforceCredentialProtectionPolicy bool
	// Uvm is "AuthenticationExtensionsClientInputs.uvm"
	//
	// Optional
	//
	// NOTE: FFI_USE_Uvm MUST be set to true to make this field effective.
	Uvm bool
	// DevicePubKey is "AuthenticationExtensionsClientInputs.devicePubKey"
	//
	// Optional
	//
	// NOTE: DevicePubKey.FFI_USE MUST be set to true to get DevicePubKey used.
	DevicePubKey AuthenticationExtensionsDevicePublicKeyInputs

	FFI_USE_MinPinLength                      bool // for MinPinLength.
	FFI_USE_GetCredBlob                       bool // for GetCredBlob.
	FFI_USE_HmacCreateSecret                  bool // for HmacCreateSecret.
	FFI_USE_CredProps                         bool // for CredProps.
	FFI_USE_EnforceCredentialProtectionPolicy bool // for EnforceCredentialProtectionPolicy.
	FFI_USE_Uvm                               bool // for Uvm.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsClientInputs with all fields set.
func (p AuthenticationExtensionsClientInputs) FromRef(ref js.Ref) AuthenticationExtensionsClientInputs {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsClientInputs in the application heap.
func (p AuthenticationExtensionsClientInputs) New() js.Ref {
	return bindings.AuthenticationExtensionsClientInputsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsClientInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsClientInputsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsClientInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsClientInputsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsClientInputs) FreeMembers(recursive bool) {
	js.Free(
		p.Appid.Ref(),
		p.AppidExclude.Ref(),
		p.CredBlob.Ref(),
		p.CredentialProtectionPolicy.Ref(),
	)
	p.Appid = p.Appid.FromRef(js.Undefined)
	p.AppidExclude = p.AppidExclude.FromRef(js.Undefined)
	p.CredBlob = p.CredBlob.FromRef(js.Undefined)
	p.CredentialProtectionPolicy = p.CredentialProtectionPolicy.FromRef(js.Undefined)
	if recursive {
		p.Prf.FreeMembers(true)
		p.Payment.FreeMembers(true)
		p.HmacGetSecret.FreeMembers(true)
		p.LargeBlob.FreeMembers(true)
		p.DevicePubKey.FreeMembers(true)
	}
}

type AuthenticationExtensionsClientInputsJSON struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsClientInputsJSON with all fields set.
func (p AuthenticationExtensionsClientInputsJSON) FromRef(ref js.Ref) AuthenticationExtensionsClientInputsJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsClientInputsJSON in the application heap.
func (p AuthenticationExtensionsClientInputsJSON) New() js.Ref {
	return bindings.AuthenticationExtensionsClientInputsJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsClientInputsJSON) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsClientInputsJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsClientInputsJSON) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsClientInputsJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsClientInputsJSON) FreeMembers(recursive bool) {
}

type AuthenticationExtensionsDevicePublicKeyOutputs struct {
	// Signature is "AuthenticationExtensionsDevicePublicKeyOutputs.signature"
	//
	// Optional
	Signature js.ArrayBuffer

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsDevicePublicKeyOutputs with all fields set.
func (p AuthenticationExtensionsDevicePublicKeyOutputs) FromRef(ref js.Ref) AuthenticationExtensionsDevicePublicKeyOutputs {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsDevicePublicKeyOutputs in the application heap.
func (p AuthenticationExtensionsDevicePublicKeyOutputs) New() js.Ref {
	return bindings.AuthenticationExtensionsDevicePublicKeyOutputsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsDevicePublicKeyOutputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsDevicePublicKeyOutputsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsDevicePublicKeyOutputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsDevicePublicKeyOutputsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsDevicePublicKeyOutputs) FreeMembers(recursive bool) {
	js.Free(
		p.Signature.Ref(),
	)
	p.Signature = p.Signature.FromRef(js.Undefined)
}

type UvmEntry = js.Array[uint32]

type UvmEntries = js.Array[UvmEntry]

type CredentialPropertiesOutput struct {
	// Rk is "CredentialPropertiesOutput.rk"
	//
	// Optional
	//
	// NOTE: FFI_USE_Rk MUST be set to true to make this field effective.
	Rk bool

	FFI_USE_Rk bool // for Rk.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CredentialPropertiesOutput with all fields set.
func (p CredentialPropertiesOutput) FromRef(ref js.Ref) CredentialPropertiesOutput {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CredentialPropertiesOutput in the application heap.
func (p CredentialPropertiesOutput) New() js.Ref {
	return bindings.CredentialPropertiesOutputJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CredentialPropertiesOutput) UpdateFrom(ref js.Ref) {
	bindings.CredentialPropertiesOutputJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CredentialPropertiesOutput) Update(ref js.Ref) {
	bindings.CredentialPropertiesOutputJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CredentialPropertiesOutput) FreeMembers(recursive bool) {
}

type AuthenticationExtensionsLargeBlobOutputs struct {
	// Supported is "AuthenticationExtensionsLargeBlobOutputs.supported"
	//
	// Optional
	//
	// NOTE: FFI_USE_Supported MUST be set to true to make this field effective.
	Supported bool
	// Blob is "AuthenticationExtensionsLargeBlobOutputs.blob"
	//
	// Optional
	Blob js.ArrayBuffer
	// Written is "AuthenticationExtensionsLargeBlobOutputs.written"
	//
	// Optional
	//
	// NOTE: FFI_USE_Written MUST be set to true to make this field effective.
	Written bool

	FFI_USE_Supported bool // for Supported.
	FFI_USE_Written   bool // for Written.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsLargeBlobOutputs with all fields set.
func (p AuthenticationExtensionsLargeBlobOutputs) FromRef(ref js.Ref) AuthenticationExtensionsLargeBlobOutputs {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsLargeBlobOutputs in the application heap.
func (p AuthenticationExtensionsLargeBlobOutputs) New() js.Ref {
	return bindings.AuthenticationExtensionsLargeBlobOutputsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsLargeBlobOutputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsLargeBlobOutputsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsLargeBlobOutputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsLargeBlobOutputsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsLargeBlobOutputs) FreeMembers(recursive bool) {
	js.Free(
		p.Blob.Ref(),
	)
	p.Blob = p.Blob.FromRef(js.Undefined)
}

type HMACGetSecretOutput struct {
	// Output1 is "HMACGetSecretOutput.output1"
	//
	// Required
	Output1 js.ArrayBuffer
	// Output2 is "HMACGetSecretOutput.output2"
	//
	// Optional
	Output2 js.ArrayBuffer

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HMACGetSecretOutput with all fields set.
func (p HMACGetSecretOutput) FromRef(ref js.Ref) HMACGetSecretOutput {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HMACGetSecretOutput in the application heap.
func (p HMACGetSecretOutput) New() js.Ref {
	return bindings.HMACGetSecretOutputJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HMACGetSecretOutput) UpdateFrom(ref js.Ref) {
	bindings.HMACGetSecretOutputJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HMACGetSecretOutput) Update(ref js.Ref) {
	bindings.HMACGetSecretOutputJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HMACGetSecretOutput) FreeMembers(recursive bool) {
	js.Free(
		p.Output1.Ref(),
		p.Output2.Ref(),
	)
	p.Output1 = p.Output1.FromRef(js.Undefined)
	p.Output2 = p.Output2.FromRef(js.Undefined)
}

type AuthenticationExtensionsPRFOutputs struct {
	// Enabled is "AuthenticationExtensionsPRFOutputs.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Results is "AuthenticationExtensionsPRFOutputs.results"
	//
	// Optional
	//
	// NOTE: Results.FFI_USE MUST be set to true to get Results used.
	Results AuthenticationExtensionsPRFValues

	FFI_USE_Enabled bool // for Enabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsPRFOutputs with all fields set.
func (p AuthenticationExtensionsPRFOutputs) FromRef(ref js.Ref) AuthenticationExtensionsPRFOutputs {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsPRFOutputs in the application heap.
func (p AuthenticationExtensionsPRFOutputs) New() js.Ref {
	return bindings.AuthenticationExtensionsPRFOutputsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsPRFOutputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFOutputsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsPRFOutputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFOutputsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsPRFOutputs) FreeMembers(recursive bool) {
	if recursive {
		p.Results.FreeMembers(true)
	}
}

type AuthenticationExtensionsClientOutputs struct {
	// DevicePubKey is "AuthenticationExtensionsClientOutputs.devicePubKey"
	//
	// Optional
	//
	// NOTE: DevicePubKey.FFI_USE MUST be set to true to get DevicePubKey used.
	DevicePubKey AuthenticationExtensionsDevicePublicKeyOutputs
	// Appid is "AuthenticationExtensionsClientOutputs.appid"
	//
	// Optional
	//
	// NOTE: FFI_USE_Appid MUST be set to true to make this field effective.
	Appid bool
	// AppidExclude is "AuthenticationExtensionsClientOutputs.appidExclude"
	//
	// Optional
	//
	// NOTE: FFI_USE_AppidExclude MUST be set to true to make this field effective.
	AppidExclude bool
	// Uvm is "AuthenticationExtensionsClientOutputs.uvm"
	//
	// Optional
	Uvm UvmEntries
	// CredProps is "AuthenticationExtensionsClientOutputs.credProps"
	//
	// Optional
	//
	// NOTE: CredProps.FFI_USE MUST be set to true to get CredProps used.
	CredProps CredentialPropertiesOutput
	// LargeBlob is "AuthenticationExtensionsClientOutputs.largeBlob"
	//
	// Optional
	//
	// NOTE: LargeBlob.FFI_USE MUST be set to true to get LargeBlob used.
	LargeBlob AuthenticationExtensionsLargeBlobOutputs
	// HmacGetSecret is "AuthenticationExtensionsClientOutputs.hmacGetSecret"
	//
	// Optional
	//
	// NOTE: HmacGetSecret.FFI_USE MUST be set to true to get HmacGetSecret used.
	HmacGetSecret HMACGetSecretOutput
	// Prf is "AuthenticationExtensionsClientOutputs.prf"
	//
	// Optional
	//
	// NOTE: Prf.FFI_USE MUST be set to true to get Prf used.
	Prf AuthenticationExtensionsPRFOutputs
	// HmacCreateSecret is "AuthenticationExtensionsClientOutputs.hmacCreateSecret"
	//
	// Optional
	//
	// NOTE: FFI_USE_HmacCreateSecret MUST be set to true to make this field effective.
	HmacCreateSecret bool

	FFI_USE_Appid            bool // for Appid.
	FFI_USE_AppidExclude     bool // for AppidExclude.
	FFI_USE_HmacCreateSecret bool // for HmacCreateSecret.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsClientOutputs with all fields set.
func (p AuthenticationExtensionsClientOutputs) FromRef(ref js.Ref) AuthenticationExtensionsClientOutputs {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsClientOutputs in the application heap.
func (p AuthenticationExtensionsClientOutputs) New() js.Ref {
	return bindings.AuthenticationExtensionsClientOutputsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsClientOutputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsClientOutputsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsClientOutputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsClientOutputsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsClientOutputs) FreeMembers(recursive bool) {
	js.Free(
		p.Uvm.Ref(),
	)
	p.Uvm = p.Uvm.FromRef(js.Undefined)
	if recursive {
		p.DevicePubKey.FreeMembers(true)
		p.CredProps.FreeMembers(true)
		p.LargeBlob.FreeMembers(true)
		p.HmacGetSecret.FreeMembers(true)
		p.Prf.FreeMembers(true)
	}
}

type AuthenticationExtensionsClientOutputsJSON struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationExtensionsClientOutputsJSON with all fields set.
func (p AuthenticationExtensionsClientOutputsJSON) FromRef(ref js.Ref) AuthenticationExtensionsClientOutputsJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationExtensionsClientOutputsJSON in the application heap.
func (p AuthenticationExtensionsClientOutputsJSON) New() js.Ref {
	return bindings.AuthenticationExtensionsClientOutputsJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationExtensionsClientOutputsJSON) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsClientOutputsJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationExtensionsClientOutputsJSON) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsClientOutputsJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationExtensionsClientOutputsJSON) FreeMembers(recursive bool) {
}

type Base64URLString = js.String

type AuthenticatorAssertionResponseJSON struct {
	// ClientDataJSON is "AuthenticatorAssertionResponseJSON.clientDataJSON"
	//
	// Required
	ClientDataJSON Base64URLString
	// AuthenticatorData is "AuthenticatorAssertionResponseJSON.authenticatorData"
	//
	// Required
	AuthenticatorData Base64URLString
	// Signature is "AuthenticatorAssertionResponseJSON.signature"
	//
	// Required
	Signature Base64URLString
	// UserHandle is "AuthenticatorAssertionResponseJSON.userHandle"
	//
	// Optional
	UserHandle Base64URLString
	// AttestationObject is "AuthenticatorAssertionResponseJSON.attestationObject"
	//
	// Optional
	AttestationObject Base64URLString

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticatorAssertionResponseJSON with all fields set.
func (p AuthenticatorAssertionResponseJSON) FromRef(ref js.Ref) AuthenticatorAssertionResponseJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticatorAssertionResponseJSON in the application heap.
func (p AuthenticatorAssertionResponseJSON) New() js.Ref {
	return bindings.AuthenticatorAssertionResponseJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticatorAssertionResponseJSON) UpdateFrom(ref js.Ref) {
	bindings.AuthenticatorAssertionResponseJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticatorAssertionResponseJSON) Update(ref js.Ref) {
	bindings.AuthenticatorAssertionResponseJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticatorAssertionResponseJSON) FreeMembers(recursive bool) {
	js.Free(
		p.ClientDataJSON.Ref(),
		p.AuthenticatorData.Ref(),
		p.Signature.Ref(),
		p.UserHandle.Ref(),
		p.AttestationObject.Ref(),
	)
	p.ClientDataJSON = p.ClientDataJSON.FromRef(js.Undefined)
	p.AuthenticatorData = p.AuthenticatorData.FromRef(js.Undefined)
	p.Signature = p.Signature.FromRef(js.Undefined)
	p.UserHandle = p.UserHandle.FromRef(js.Undefined)
	p.AttestationObject = p.AttestationObject.FromRef(js.Undefined)
}

type AuthenticationResponseJSON struct {
	// Id is "AuthenticationResponseJSON.id"
	//
	// Required
	Id Base64URLString
	// RawId is "AuthenticationResponseJSON.rawId"
	//
	// Required
	RawId Base64URLString
	// Response is "AuthenticationResponseJSON.response"
	//
	// Required
	//
	// NOTE: Response.FFI_USE MUST be set to true to get Response used.
	Response AuthenticatorAssertionResponseJSON
	// AuthenticatorAttachment is "AuthenticationResponseJSON.authenticatorAttachment"
	//
	// Optional
	AuthenticatorAttachment js.String
	// ClientExtensionResults is "AuthenticationResponseJSON.clientExtensionResults"
	//
	// Required
	//
	// NOTE: ClientExtensionResults.FFI_USE MUST be set to true to get ClientExtensionResults used.
	ClientExtensionResults AuthenticationExtensionsClientOutputsJSON
	// Type is "AuthenticationResponseJSON.type"
	//
	// Required
	Type js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuthenticationResponseJSON with all fields set.
func (p AuthenticationResponseJSON) FromRef(ref js.Ref) AuthenticationResponseJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuthenticationResponseJSON in the application heap.
func (p AuthenticationResponseJSON) New() js.Ref {
	return bindings.AuthenticationResponseJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AuthenticationResponseJSON) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationResponseJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuthenticationResponseJSON) Update(ref js.Ref) {
	bindings.AuthenticationResponseJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuthenticationResponseJSON) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.RawId.Ref(),
		p.AuthenticatorAttachment.Ref(),
		p.Type.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.RawId = p.RawId.FromRef(js.Undefined)
	p.AuthenticatorAttachment = p.AuthenticatorAttachment.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
	if recursive {
		p.Response.FreeMembers(true)
		p.ClientExtensionResults.FreeMembers(true)
	}
}

type AuthenticatorAssertionResponse struct {
	AuthenticatorResponse
}

func (this AuthenticatorAssertionResponse) Once() AuthenticatorAssertionResponse {
	this.ref.Once()
	return this
}

func (this AuthenticatorAssertionResponse) Ref() js.Ref {
	return this.AuthenticatorResponse.Ref()
}

func (this AuthenticatorAssertionResponse) FromRef(ref js.Ref) AuthenticatorAssertionResponse {
	this.AuthenticatorResponse = this.AuthenticatorResponse.FromRef(ref)
	return this
}

func (this AuthenticatorAssertionResponse) Free() {
	this.ref.Free()
}

// AuthenticatorData returns the value of property "AuthenticatorAssertionResponse.authenticatorData".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorAssertionResponse) AuthenticatorData() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorAssertionResponseAuthenticatorData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Signature returns the value of property "AuthenticatorAssertionResponse.signature".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorAssertionResponse) Signature() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorAssertionResponseSignature(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UserHandle returns the value of property "AuthenticatorAssertionResponse.userHandle".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorAssertionResponse) UserHandle() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorAssertionResponseUserHandle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AttestationObject returns the value of property "AuthenticatorAssertionResponse.attestationObject".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorAssertionResponse) AttestationObject() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorAssertionResponseAttestationObject(
		this.ref, js.Pointer(&ret),
	)
	return
}

type AuthenticatorAttachment uint32

const (
	_ AuthenticatorAttachment = iota

	AuthenticatorAttachment_PLATFORM
	AuthenticatorAttachment_CROSS_PLATFORM
)

func (AuthenticatorAttachment) FromRef(str js.Ref) AuthenticatorAttachment {
	return AuthenticatorAttachment(bindings.ConstOfAuthenticatorAttachment(str))
}

func (x AuthenticatorAttachment) String() (string, bool) {
	switch x {
	case AuthenticatorAttachment_PLATFORM:
		return "platform", true
	case AuthenticatorAttachment_CROSS_PLATFORM:
		return "cross-platform", true
	default:
		return "", false
	}
}
