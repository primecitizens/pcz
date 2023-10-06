// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type AudioEncoderInit struct {
	// Output is "AudioEncoderInit.output"
	//
	// Required
	Output js.Func[func(output EncodedAudioChunk, metadata EncodedAudioChunkMetadata)]
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
func (p AudioEncoderInit) UpdateFrom(ref js.Ref) {
	bindings.AudioEncoderInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioEncoderInit) Update(ref js.Ref) {
	bindings.AudioEncoderInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p FlacEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.FlacEncoderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FlacEncoderConfig) Update(ref js.Ref) {
	bindings.FlacEncoderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p OpusEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.OpusEncoderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p OpusEncoderConfig) Update(ref js.Ref) {
	bindings.OpusEncoderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AudioEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.AudioEncoderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioEncoderConfig) Update(ref js.Ref) {
	bindings.AudioEncoderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AudioEncoderSupport) UpdateFrom(ref js.Ref) {
	bindings.AudioEncoderSupportJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioEncoderSupport) Update(ref js.Ref) {
	bindings.AudioEncoderSupportJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// State returns the value of property "AudioEncoder.state".
//
// It returns ok=false if there is no such property.
func (this AudioEncoder) State() (ret CodecState, ok bool) {
	ok = js.True == bindings.GetAudioEncoderState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EncodeQueueSize returns the value of property "AudioEncoder.encodeQueueSize".
//
// It returns ok=false if there is no such property.
func (this AudioEncoder) EncodeQueueSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioEncoderEncodeQueueSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasConfigure returns true if the method "AudioEncoder.configure" exists.
func (this AudioEncoder) HasConfigure() bool {
	return js.True == bindings.HasAudioEncoderConfigure(
		this.Ref(),
	)
}

// ConfigureFunc returns the method "AudioEncoder.configure".
func (this AudioEncoder) ConfigureFunc() (fn js.Func[func(config AudioEncoderConfig)]) {
	return fn.FromRef(
		bindings.AudioEncoderConfigureFunc(
			this.Ref(),
		),
	)
}

// Configure calls the method "AudioEncoder.configure".
func (this AudioEncoder) Configure(config AudioEncoderConfig) (ret js.Void) {
	bindings.CallAudioEncoderConfigure(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryConfigure calls the method "AudioEncoder.configure"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioEncoder) TryConfigure(config AudioEncoderConfig) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderConfigure(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

// HasEncode returns true if the method "AudioEncoder.encode" exists.
func (this AudioEncoder) HasEncode() bool {
	return js.True == bindings.HasAudioEncoderEncode(
		this.Ref(),
	)
}

// EncodeFunc returns the method "AudioEncoder.encode".
func (this AudioEncoder) EncodeFunc() (fn js.Func[func(data AudioData)]) {
	return fn.FromRef(
		bindings.AudioEncoderEncodeFunc(
			this.Ref(),
		),
	)
}

// Encode calls the method "AudioEncoder.encode".
func (this AudioEncoder) Encode(data AudioData) (ret js.Void) {
	bindings.CallAudioEncoderEncode(
		this.Ref(), js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryEncode calls the method "AudioEncoder.encode"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioEncoder) TryEncode(data AudioData) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderEncode(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFlush returns true if the method "AudioEncoder.flush" exists.
func (this AudioEncoder) HasFlush() bool {
	return js.True == bindings.HasAudioEncoderFlush(
		this.Ref(),
	)
}

// FlushFunc returns the method "AudioEncoder.flush".
func (this AudioEncoder) FlushFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.AudioEncoderFlushFunc(
			this.Ref(),
		),
	)
}

// Flush calls the method "AudioEncoder.flush".
func (this AudioEncoder) Flush() (ret js.Promise[js.Void]) {
	bindings.CallAudioEncoderFlush(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFlush calls the method "AudioEncoder.flush"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioEncoder) TryFlush() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderFlush(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReset returns true if the method "AudioEncoder.reset" exists.
func (this AudioEncoder) HasReset() bool {
	return js.True == bindings.HasAudioEncoderReset(
		this.Ref(),
	)
}

// ResetFunc returns the method "AudioEncoder.reset".
func (this AudioEncoder) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioEncoderResetFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "AudioEncoder.reset".
func (this AudioEncoder) Reset() (ret js.Void) {
	bindings.CallAudioEncoderReset(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "AudioEncoder.reset"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioEncoder) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderReset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "AudioEncoder.close" exists.
func (this AudioEncoder) HasClose() bool {
	return js.True == bindings.HasAudioEncoderClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "AudioEncoder.close".
func (this AudioEncoder) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioEncoderCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "AudioEncoder.close".
func (this AudioEncoder) Close() (ret js.Void) {
	bindings.CallAudioEncoderClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "AudioEncoder.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioEncoder) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsConfigSupported returns true if the staticmethod "AudioEncoder.isConfigSupported" exists.
func (this AudioEncoder) HasIsConfigSupported() bool {
	return js.True == bindings.HasAudioEncoderIsConfigSupported(
		this.Ref(),
	)
}

// IsConfigSupportedFunc returns the staticmethod "AudioEncoder.isConfigSupported".
func (this AudioEncoder) IsConfigSupportedFunc() (fn js.Func[func(config AudioEncoderConfig) js.Promise[AudioEncoderSupport]]) {
	return fn.FromRef(
		bindings.AudioEncoderIsConfigSupportedFunc(
			this.Ref(),
		),
	)
}

// IsConfigSupported calls the staticmethod "AudioEncoder.isConfigSupported".
func (this AudioEncoder) IsConfigSupported(config AudioEncoderConfig) (ret js.Promise[AudioEncoderSupport]) {
	bindings.CallAudioEncoderIsConfigSupported(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&config),
	)

	return
}

// TryIsConfigSupported calls the staticmethod "AudioEncoder.isConfigSupported"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioEncoder) TryIsConfigSupported(config AudioEncoderConfig) (ret js.Promise[AudioEncoderSupport], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioEncoderIsConfigSupported(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&config),
	)

	return
}

type AudioNode struct {
	EventTarget
}

func (this AudioNode) Once() AudioNode {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Context returns the value of property "AudioNode.context".
//
// It returns ok=false if there is no such property.
func (this AudioNode) Context() (ret BaseAudioContext, ok bool) {
	ok = js.True == bindings.GetAudioNodeContext(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberOfInputs returns the value of property "AudioNode.numberOfInputs".
//
// It returns ok=false if there is no such property.
func (this AudioNode) NumberOfInputs() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioNodeNumberOfInputs(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NumberOfOutputs returns the value of property "AudioNode.numberOfOutputs".
//
// It returns ok=false if there is no such property.
func (this AudioNode) NumberOfOutputs() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioNodeNumberOfOutputs(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ChannelCount returns the value of property "AudioNode.channelCount".
//
// It returns ok=false if there is no such property.
func (this AudioNode) ChannelCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioNodeChannelCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetChannelCount sets the value of property "AudioNode.channelCount" to val.
//
// It returns false if the property cannot be set.
func (this AudioNode) SetChannelCount(val uint32) bool {
	return js.True == bindings.SetAudioNodeChannelCount(
		this.Ref(),
		uint32(val),
	)
}

// ChannelCountMode returns the value of property "AudioNode.channelCountMode".
//
// It returns ok=false if there is no such property.
func (this AudioNode) ChannelCountMode() (ret ChannelCountMode, ok bool) {
	ok = js.True == bindings.GetAudioNodeChannelCountMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetChannelCountMode sets the value of property "AudioNode.channelCountMode" to val.
//
// It returns false if the property cannot be set.
func (this AudioNode) SetChannelCountMode(val ChannelCountMode) bool {
	return js.True == bindings.SetAudioNodeChannelCountMode(
		this.Ref(),
		uint32(val),
	)
}

// ChannelInterpretation returns the value of property "AudioNode.channelInterpretation".
//
// It returns ok=false if there is no such property.
func (this AudioNode) ChannelInterpretation() (ret ChannelInterpretation, ok bool) {
	ok = js.True == bindings.GetAudioNodeChannelInterpretation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetChannelInterpretation sets the value of property "AudioNode.channelInterpretation" to val.
//
// It returns false if the property cannot be set.
func (this AudioNode) SetChannelInterpretation(val ChannelInterpretation) bool {
	return js.True == bindings.SetAudioNodeChannelInterpretation(
		this.Ref(),
		uint32(val),
	)
}

// HasConnect returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasConnect() bool {
	return js.True == bindings.HasAudioNodeConnect(
		this.Ref(),
	)
}

// ConnectFunc returns the method "AudioNode.connect".
func (this AudioNode) ConnectFunc() (fn js.Func[func(destinationNode AudioNode, output uint32, input uint32) AudioNode]) {
	return fn.FromRef(
		bindings.AudioNodeConnectFunc(
			this.Ref(),
		),
	)
}

// Connect calls the method "AudioNode.connect".
func (this AudioNode) Connect(destinationNode AudioNode, output uint32, input uint32) (ret AudioNode) {
	bindings.CallAudioNodeConnect(
		this.Ref(), js.Pointer(&ret),
		destinationNode.Ref(),
		uint32(output),
		uint32(input),
	)

	return
}

// TryConnect calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryConnect(destinationNode AudioNode, output uint32, input uint32) (ret AudioNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
		uint32(output),
		uint32(input),
	)

	return
}

// HasConnect1 returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasConnect1() bool {
	return js.True == bindings.HasAudioNodeConnect1(
		this.Ref(),
	)
}

// Connect1Func returns the method "AudioNode.connect".
func (this AudioNode) Connect1Func() (fn js.Func[func(destinationNode AudioNode, output uint32) AudioNode]) {
	return fn.FromRef(
		bindings.AudioNodeConnect1Func(
			this.Ref(),
		),
	)
}

// Connect1 calls the method "AudioNode.connect".
func (this AudioNode) Connect1(destinationNode AudioNode, output uint32) (ret AudioNode) {
	bindings.CallAudioNodeConnect1(
		this.Ref(), js.Pointer(&ret),
		destinationNode.Ref(),
		uint32(output),
	)

	return
}

// TryConnect1 calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryConnect1(destinationNode AudioNode, output uint32) (ret AudioNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
		uint32(output),
	)

	return
}

// HasConnect2 returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasConnect2() bool {
	return js.True == bindings.HasAudioNodeConnect2(
		this.Ref(),
	)
}

// Connect2Func returns the method "AudioNode.connect".
func (this AudioNode) Connect2Func() (fn js.Func[func(destinationNode AudioNode) AudioNode]) {
	return fn.FromRef(
		bindings.AudioNodeConnect2Func(
			this.Ref(),
		),
	)
}

// Connect2 calls the method "AudioNode.connect".
func (this AudioNode) Connect2(destinationNode AudioNode) (ret AudioNode) {
	bindings.CallAudioNodeConnect2(
		this.Ref(), js.Pointer(&ret),
		destinationNode.Ref(),
	)

	return
}

// TryConnect2 calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryConnect2(destinationNode AudioNode) (ret AudioNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
	)

	return
}

// HasConnect3 returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasConnect3() bool {
	return js.True == bindings.HasAudioNodeConnect3(
		this.Ref(),
	)
}

// Connect3Func returns the method "AudioNode.connect".
func (this AudioNode) Connect3Func() (fn js.Func[func(destinationParam AudioParam, output uint32)]) {
	return fn.FromRef(
		bindings.AudioNodeConnect3Func(
			this.Ref(),
		),
	)
}

// Connect3 calls the method "AudioNode.connect".
func (this AudioNode) Connect3(destinationParam AudioParam, output uint32) (ret js.Void) {
	bindings.CallAudioNodeConnect3(
		this.Ref(), js.Pointer(&ret),
		destinationParam.Ref(),
		uint32(output),
	)

	return
}

// TryConnect3 calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryConnect3(destinationParam AudioParam, output uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destinationParam.Ref(),
		uint32(output),
	)

	return
}

// HasConnect4 returns true if the method "AudioNode.connect" exists.
func (this AudioNode) HasConnect4() bool {
	return js.True == bindings.HasAudioNodeConnect4(
		this.Ref(),
	)
}

// Connect4Func returns the method "AudioNode.connect".
func (this AudioNode) Connect4Func() (fn js.Func[func(destinationParam AudioParam)]) {
	return fn.FromRef(
		bindings.AudioNodeConnect4Func(
			this.Ref(),
		),
	)
}

// Connect4 calls the method "AudioNode.connect".
func (this AudioNode) Connect4(destinationParam AudioParam) (ret js.Void) {
	bindings.CallAudioNodeConnect4(
		this.Ref(), js.Pointer(&ret),
		destinationParam.Ref(),
	)

	return
}

// TryConnect4 calls the method "AudioNode.connect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryConnect4(destinationParam AudioParam) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeConnect4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destinationParam.Ref(),
	)

	return
}

// HasDisconnect returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasDisconnect() bool {
	return js.True == bindings.HasAudioNodeDisconnect(
		this.Ref(),
	)
}

// DisconnectFunc returns the method "AudioNode.disconnect".
func (this AudioNode) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioNodeDisconnectFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect() (ret js.Void) {
	bindings.CallAudioNodeDisconnect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDisconnect1 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasDisconnect1() bool {
	return js.True == bindings.HasAudioNodeDisconnect1(
		this.Ref(),
	)
}

// Disconnect1Func returns the method "AudioNode.disconnect".
func (this AudioNode) Disconnect1Func() (fn js.Func[func(output uint32)]) {
	return fn.FromRef(
		bindings.AudioNodeDisconnect1Func(
			this.Ref(),
		),
	)
}

// Disconnect1 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect1(output uint32) (ret js.Void) {
	bindings.CallAudioNodeDisconnect1(
		this.Ref(), js.Pointer(&ret),
		uint32(output),
	)

	return
}

// TryDisconnect1 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryDisconnect1(output uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(output),
	)

	return
}

// HasDisconnect2 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasDisconnect2() bool {
	return js.True == bindings.HasAudioNodeDisconnect2(
		this.Ref(),
	)
}

// Disconnect2Func returns the method "AudioNode.disconnect".
func (this AudioNode) Disconnect2Func() (fn js.Func[func(destinationNode AudioNode)]) {
	return fn.FromRef(
		bindings.AudioNodeDisconnect2Func(
			this.Ref(),
		),
	)
}

// Disconnect2 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect2(destinationNode AudioNode) (ret js.Void) {
	bindings.CallAudioNodeDisconnect2(
		this.Ref(), js.Pointer(&ret),
		destinationNode.Ref(),
	)

	return
}

// TryDisconnect2 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryDisconnect2(destinationNode AudioNode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
	)

	return
}

// HasDisconnect3 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasDisconnect3() bool {
	return js.True == bindings.HasAudioNodeDisconnect3(
		this.Ref(),
	)
}

// Disconnect3Func returns the method "AudioNode.disconnect".
func (this AudioNode) Disconnect3Func() (fn js.Func[func(destinationNode AudioNode, output uint32)]) {
	return fn.FromRef(
		bindings.AudioNodeDisconnect3Func(
			this.Ref(),
		),
	)
}

// Disconnect3 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect3(destinationNode AudioNode, output uint32) (ret js.Void) {
	bindings.CallAudioNodeDisconnect3(
		this.Ref(), js.Pointer(&ret),
		destinationNode.Ref(),
		uint32(output),
	)

	return
}

// TryDisconnect3 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryDisconnect3(destinationNode AudioNode, output uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
		uint32(output),
	)

	return
}

// HasDisconnect4 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasDisconnect4() bool {
	return js.True == bindings.HasAudioNodeDisconnect4(
		this.Ref(),
	)
}

// Disconnect4Func returns the method "AudioNode.disconnect".
func (this AudioNode) Disconnect4Func() (fn js.Func[func(destinationNode AudioNode, output uint32, input uint32)]) {
	return fn.FromRef(
		bindings.AudioNodeDisconnect4Func(
			this.Ref(),
		),
	)
}

// Disconnect4 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect4(destinationNode AudioNode, output uint32, input uint32) (ret js.Void) {
	bindings.CallAudioNodeDisconnect4(
		this.Ref(), js.Pointer(&ret),
		destinationNode.Ref(),
		uint32(output),
		uint32(input),
	)

	return
}

// TryDisconnect4 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryDisconnect4(destinationNode AudioNode, output uint32, input uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destinationNode.Ref(),
		uint32(output),
		uint32(input),
	)

	return
}

// HasDisconnect5 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasDisconnect5() bool {
	return js.True == bindings.HasAudioNodeDisconnect5(
		this.Ref(),
	)
}

// Disconnect5Func returns the method "AudioNode.disconnect".
func (this AudioNode) Disconnect5Func() (fn js.Func[func(destinationParam AudioParam)]) {
	return fn.FromRef(
		bindings.AudioNodeDisconnect5Func(
			this.Ref(),
		),
	)
}

// Disconnect5 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect5(destinationParam AudioParam) (ret js.Void) {
	bindings.CallAudioNodeDisconnect5(
		this.Ref(), js.Pointer(&ret),
		destinationParam.Ref(),
	)

	return
}

// TryDisconnect5 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryDisconnect5(destinationParam AudioParam) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect5(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		destinationParam.Ref(),
	)

	return
}

// HasDisconnect6 returns true if the method "AudioNode.disconnect" exists.
func (this AudioNode) HasDisconnect6() bool {
	return js.True == bindings.HasAudioNodeDisconnect6(
		this.Ref(),
	)
}

// Disconnect6Func returns the method "AudioNode.disconnect".
func (this AudioNode) Disconnect6Func() (fn js.Func[func(destinationParam AudioParam, output uint32)]) {
	return fn.FromRef(
		bindings.AudioNodeDisconnect6Func(
			this.Ref(),
		),
	)
}

// Disconnect6 calls the method "AudioNode.disconnect".
func (this AudioNode) Disconnect6(destinationParam AudioParam, output uint32) (ret js.Void) {
	bindings.CallAudioNodeDisconnect6(
		this.Ref(), js.Pointer(&ret),
		destinationParam.Ref(),
		uint32(output),
	)

	return
}

// TryDisconnect6 calls the method "AudioNode.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioNode) TryDisconnect6(destinationParam AudioParam, output uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioNodeDisconnect6(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p AudioOutputOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioOutputOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioOutputOptions) Update(ref js.Ref) {
	bindings.AudioOutputOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AudioParamDescriptor) UpdateFrom(ref js.Ref) {
	bindings.AudioParamDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioParamDescriptor) Update(ref js.Ref) {
	bindings.AudioParamDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AudioParamMap struct {
	ref js.Ref
}

func (this AudioParamMap) Once() AudioParamMap {
	this.Ref().Once()
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
	this.Ref().Free()
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
func (p AudioProcessingEventInit) UpdateFrom(ref js.Ref) {
	bindings.AudioProcessingEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioProcessingEventInit) Update(ref js.Ref) {
	bindings.AudioProcessingEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// PlaybackTime returns the value of property "AudioProcessingEvent.playbackTime".
//
// It returns ok=false if there is no such property.
func (this AudioProcessingEvent) PlaybackTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioProcessingEventPlaybackTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InputBuffer returns the value of property "AudioProcessingEvent.inputBuffer".
//
// It returns ok=false if there is no such property.
func (this AudioProcessingEvent) InputBuffer() (ret AudioBuffer, ok bool) {
	ok = js.True == bindings.GetAudioProcessingEventInputBuffer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OutputBuffer returns the value of property "AudioProcessingEvent.outputBuffer".
//
// It returns ok=false if there is no such property.
func (this AudioProcessingEvent) OutputBuffer() (ret AudioBuffer, ok bool) {
	ok = js.True == bindings.GetAudioProcessingEventOutputBuffer(
		this.Ref(), js.Pointer(&ret),
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
func (p AudioRenderCapacityEventInit) UpdateFrom(ref js.Ref) {
	bindings.AudioRenderCapacityEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioRenderCapacityEventInit) Update(ref js.Ref) {
	bindings.AudioRenderCapacityEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Timestamp returns the value of property "AudioRenderCapacityEvent.timestamp".
//
// It returns ok=false if there is no such property.
func (this AudioRenderCapacityEvent) Timestamp() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioRenderCapacityEventTimestamp(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AverageLoad returns the value of property "AudioRenderCapacityEvent.averageLoad".
//
// It returns ok=false if there is no such property.
func (this AudioRenderCapacityEvent) AverageLoad() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioRenderCapacityEventAverageLoad(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PeakLoad returns the value of property "AudioRenderCapacityEvent.peakLoad".
//
// It returns ok=false if there is no such property.
func (this AudioRenderCapacityEvent) PeakLoad() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioRenderCapacityEventPeakLoad(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UnderrunRatio returns the value of property "AudioRenderCapacityEvent.underrunRatio".
//
// It returns ok=false if there is no such property.
func (this AudioRenderCapacityEvent) UnderrunRatio() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioRenderCapacityEventUnderrunRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type AudioScheduledSourceNode struct {
	AudioNode
}

func (this AudioScheduledSourceNode) Once() AudioScheduledSourceNode {
	this.Ref().Once()
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
	this.Ref().Free()
}

// HasStart returns true if the method "AudioScheduledSourceNode.start" exists.
func (this AudioScheduledSourceNode) HasStart() bool {
	return js.True == bindings.HasAudioScheduledSourceNodeStart(
		this.Ref(),
	)
}

// StartFunc returns the method "AudioScheduledSourceNode.start".
func (this AudioScheduledSourceNode) StartFunc() (fn js.Func[func(when float64)]) {
	return fn.FromRef(
		bindings.AudioScheduledSourceNodeStartFunc(
			this.Ref(),
		),
	)
}

// Start calls the method "AudioScheduledSourceNode.start".
func (this AudioScheduledSourceNode) Start(when float64) (ret js.Void) {
	bindings.CallAudioScheduledSourceNodeStart(
		this.Ref(), js.Pointer(&ret),
		float64(when),
	)

	return
}

// TryStart calls the method "AudioScheduledSourceNode.start"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioScheduledSourceNode) TryStart(when float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioScheduledSourceNodeStart(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(when),
	)

	return
}

// HasStart1 returns true if the method "AudioScheduledSourceNode.start" exists.
func (this AudioScheduledSourceNode) HasStart1() bool {
	return js.True == bindings.HasAudioScheduledSourceNodeStart1(
		this.Ref(),
	)
}

// Start1Func returns the method "AudioScheduledSourceNode.start".
func (this AudioScheduledSourceNode) Start1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioScheduledSourceNodeStart1Func(
			this.Ref(),
		),
	)
}

// Start1 calls the method "AudioScheduledSourceNode.start".
func (this AudioScheduledSourceNode) Start1() (ret js.Void) {
	bindings.CallAudioScheduledSourceNodeStart1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStart1 calls the method "AudioScheduledSourceNode.start"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioScheduledSourceNode) TryStart1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioScheduledSourceNodeStart1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStop returns true if the method "AudioScheduledSourceNode.stop" exists.
func (this AudioScheduledSourceNode) HasStop() bool {
	return js.True == bindings.HasAudioScheduledSourceNodeStop(
		this.Ref(),
	)
}

// StopFunc returns the method "AudioScheduledSourceNode.stop".
func (this AudioScheduledSourceNode) StopFunc() (fn js.Func[func(when float64)]) {
	return fn.FromRef(
		bindings.AudioScheduledSourceNodeStopFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "AudioScheduledSourceNode.stop".
func (this AudioScheduledSourceNode) Stop(when float64) (ret js.Void) {
	bindings.CallAudioScheduledSourceNodeStop(
		this.Ref(), js.Pointer(&ret),
		float64(when),
	)

	return
}

// TryStop calls the method "AudioScheduledSourceNode.stop"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioScheduledSourceNode) TryStop(when float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioScheduledSourceNodeStop(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(when),
	)

	return
}

// HasStop1 returns true if the method "AudioScheduledSourceNode.stop" exists.
func (this AudioScheduledSourceNode) HasStop1() bool {
	return js.True == bindings.HasAudioScheduledSourceNodeStop1(
		this.Ref(),
	)
}

// Stop1Func returns the method "AudioScheduledSourceNode.stop".
func (this AudioScheduledSourceNode) Stop1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AudioScheduledSourceNodeStop1Func(
			this.Ref(),
		),
	)
}

// Stop1 calls the method "AudioScheduledSourceNode.stop".
func (this AudioScheduledSourceNode) Stop1() (ret js.Void) {
	bindings.CallAudioScheduledSourceNodeStop1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStop1 calls the method "AudioScheduledSourceNode.stop"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioScheduledSourceNode) TryStop1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioScheduledSourceNodeStop1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
		assert.Throw("invalid", "callback", "invocation")
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Port returns the value of property "AudioWorkletProcessor.port".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletProcessor) Port() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetAudioWorkletProcessorPort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type AudioWorkletGlobalScope struct {
	WorkletGlobalScope
}

func (this AudioWorkletGlobalScope) Once() AudioWorkletGlobalScope {
	this.Ref().Once()
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
	this.Ref().Free()
}

// CurrentFrame returns the value of property "AudioWorkletGlobalScope.currentFrame".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletGlobalScope) CurrentFrame() (ret uint64, ok bool) {
	ok = js.True == bindings.GetAudioWorkletGlobalScopeCurrentFrame(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentTime returns the value of property "AudioWorkletGlobalScope.currentTime".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletGlobalScope) CurrentTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetAudioWorkletGlobalScopeCurrentTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SampleRate returns the value of property "AudioWorkletGlobalScope.sampleRate".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletGlobalScope) SampleRate() (ret float32, ok bool) {
	ok = js.True == bindings.GetAudioWorkletGlobalScopeSampleRate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Port returns the value of property "AudioWorkletGlobalScope.port".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletGlobalScope) Port() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetAudioWorkletGlobalScopePort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRegisterProcessor returns true if the method "AudioWorkletGlobalScope.registerProcessor" exists.
func (this AudioWorkletGlobalScope) HasRegisterProcessor() bool {
	return js.True == bindings.HasAudioWorkletGlobalScopeRegisterProcessor(
		this.Ref(),
	)
}

// RegisterProcessorFunc returns the method "AudioWorkletGlobalScope.registerProcessor".
func (this AudioWorkletGlobalScope) RegisterProcessorFunc() (fn js.Func[func(name js.String, processorCtor js.Func[func(options js.Object) AudioWorkletProcessor])]) {
	return fn.FromRef(
		bindings.AudioWorkletGlobalScopeRegisterProcessorFunc(
			this.Ref(),
		),
	)
}

// RegisterProcessor calls the method "AudioWorkletGlobalScope.registerProcessor".
func (this AudioWorkletGlobalScope) RegisterProcessor(name js.String, processorCtor js.Func[func(options js.Object) AudioWorkletProcessor]) (ret js.Void) {
	bindings.CallAudioWorkletGlobalScopeRegisterProcessor(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		processorCtor.Ref(),
	)

	return
}

// TryRegisterProcessor calls the method "AudioWorkletGlobalScope.registerProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioWorkletGlobalScope) TryRegisterProcessor(name js.String, processorCtor js.Func[func(options js.Object) AudioWorkletProcessor]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioWorkletGlobalScopeRegisterProcessor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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
func (p AudioWorkletNodeOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioWorkletNodeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioWorkletNodeOptions) Update(ref js.Ref) {
	bindings.AudioWorkletNodeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	this.Ref().Once()
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
	this.Ref().Free()
}

// Parameters returns the value of property "AudioWorkletNode.parameters".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletNode) Parameters() (ret AudioParamMap, ok bool) {
	ok = js.True == bindings.GetAudioWorkletNodeParameters(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Port returns the value of property "AudioWorkletNode.port".
//
// It returns ok=false if there is no such property.
func (this AudioWorkletNode) Port() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetAudioWorkletNodePort(
		this.Ref(), js.Pointer(&ret),
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
		assert.Throw("invalid", "callback", "invocation")
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
func (p AuthenticationExtensionsPRFValues) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsPRFValues) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsPRFInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFInputsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsPRFInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFInputsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p PaymentCurrencyAmount) UpdateFrom(ref js.Ref) {
	bindings.PaymentCurrencyAmountJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentCurrencyAmount) Update(ref js.Ref) {
	bindings.PaymentCurrencyAmountJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p PaymentCredentialInstrument) UpdateFrom(ref js.Ref) {
	bindings.PaymentCredentialInstrumentJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentCredentialInstrument) Update(ref js.Ref) {
	bindings.PaymentCredentialInstrumentJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsPaymentInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsPaymentInputsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsPaymentInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsPaymentInputsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p HMACGetSecretInput) UpdateFrom(ref js.Ref) {
	bindings.HMACGetSecretInputJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HMACGetSecretInput) Update(ref js.Ref) {
	bindings.HMACGetSecretInputJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsLargeBlobInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsLargeBlobInputsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsLargeBlobInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsLargeBlobInputsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsDevicePublicKeyInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsDevicePublicKeyInputsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsDevicePublicKeyInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsDevicePublicKeyInputsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsClientInputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsClientInputsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsClientInputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsClientInputsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsClientInputsJSON) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsClientInputsJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsClientInputsJSON) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsClientInputsJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsDevicePublicKeyOutputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsDevicePublicKeyOutputsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsDevicePublicKeyOutputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsDevicePublicKeyOutputsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p CredentialPropertiesOutput) UpdateFrom(ref js.Ref) {
	bindings.CredentialPropertiesOutputJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CredentialPropertiesOutput) Update(ref js.Ref) {
	bindings.CredentialPropertiesOutputJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsLargeBlobOutputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsLargeBlobOutputsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsLargeBlobOutputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsLargeBlobOutputsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p HMACGetSecretOutput) UpdateFrom(ref js.Ref) {
	bindings.HMACGetSecretOutputJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HMACGetSecretOutput) Update(ref js.Ref) {
	bindings.HMACGetSecretOutputJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsPRFOutputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFOutputsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsPRFOutputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsPRFOutputsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsClientOutputs) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsClientOutputsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsClientOutputs) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsClientOutputsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationExtensionsClientOutputsJSON) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationExtensionsClientOutputsJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationExtensionsClientOutputsJSON) Update(ref js.Ref) {
	bindings.AuthenticationExtensionsClientOutputsJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticatorAssertionResponseJSON) UpdateFrom(ref js.Ref) {
	bindings.AuthenticatorAssertionResponseJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticatorAssertionResponseJSON) Update(ref js.Ref) {
	bindings.AuthenticatorAssertionResponseJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
func (p AuthenticationResponseJSON) UpdateFrom(ref js.Ref) {
	bindings.AuthenticationResponseJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuthenticationResponseJSON) Update(ref js.Ref) {
	bindings.AuthenticationResponseJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AuthenticatorAssertionResponse struct {
	AuthenticatorResponse
}

func (this AuthenticatorAssertionResponse) Once() AuthenticatorAssertionResponse {
	this.Ref().Once()
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
	this.Ref().Free()
}

// AuthenticatorData returns the value of property "AuthenticatorAssertionResponse.authenticatorData".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorAssertionResponse) AuthenticatorData() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorAssertionResponseAuthenticatorData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Signature returns the value of property "AuthenticatorAssertionResponse.signature".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorAssertionResponse) Signature() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorAssertionResponseSignature(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UserHandle returns the value of property "AuthenticatorAssertionResponse.userHandle".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorAssertionResponse) UserHandle() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorAssertionResponseUserHandle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AttestationObject returns the value of property "AuthenticatorAssertionResponse.attestationObject".
//
// It returns ok=false if there is no such property.
func (this AuthenticatorAssertionResponse) AttestationObject() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetAuthenticatorAssertionResponseAttestationObject(
		this.Ref(), js.Pointer(&ret),
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
