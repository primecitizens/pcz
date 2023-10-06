// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type GPURenderPassEncoder struct {
	ref js.Ref
}

func (this GPURenderPassEncoder) Once() GPURenderPassEncoder {
	this.Ref().Once()
	return this
}

func (this GPURenderPassEncoder) Ref() js.Ref {
	return this.ref
}

func (this GPURenderPassEncoder) FromRef(ref js.Ref) GPURenderPassEncoder {
	this.ref = ref
	return this
}

func (this GPURenderPassEncoder) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPURenderPassEncoder.label".
//
// It returns ok=false if there is no such property.
func (this GPURenderPassEncoder) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPURenderPassEncoderLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPURenderPassEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPURenderPassEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPURenderPassEncoderLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasSetViewport returns true if the method "GPURenderPassEncoder.setViewport" exists.
func (this GPURenderPassEncoder) HasSetViewport() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetViewport(
		this.Ref(),
	)
}

// SetViewportFunc returns the method "GPURenderPassEncoder.setViewport".
func (this GPURenderPassEncoder) SetViewportFunc() (fn js.Func[func(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetViewportFunc(
			this.Ref(),
		),
	)
}

// SetViewport calls the method "GPURenderPassEncoder.setViewport".
func (this GPURenderPassEncoder) SetViewport(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetViewport(
		this.Ref(), js.Pointer(&ret),
		float32(x),
		float32(y),
		float32(width),
		float32(height),
		float32(minDepth),
		float32(maxDepth),
	)

	return
}

// TrySetViewport calls the method "GPURenderPassEncoder.setViewport"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetViewport(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetViewport(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(x),
		float32(y),
		float32(width),
		float32(height),
		float32(minDepth),
		float32(maxDepth),
	)

	return
}

// HasSetScissorRect returns true if the method "GPURenderPassEncoder.setScissorRect" exists.
func (this GPURenderPassEncoder) HasSetScissorRect() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetScissorRect(
		this.Ref(),
	)
}

// SetScissorRectFunc returns the method "GPURenderPassEncoder.setScissorRect".
func (this GPURenderPassEncoder) SetScissorRectFunc() (fn js.Func[func(x GPUIntegerCoordinate, y GPUIntegerCoordinate, width GPUIntegerCoordinate, height GPUIntegerCoordinate)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetScissorRectFunc(
			this.Ref(),
		),
	)
}

// SetScissorRect calls the method "GPURenderPassEncoder.setScissorRect".
func (this GPURenderPassEncoder) SetScissorRect(x GPUIntegerCoordinate, y GPUIntegerCoordinate, width GPUIntegerCoordinate, height GPUIntegerCoordinate) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetScissorRect(
		this.Ref(), js.Pointer(&ret),
		uint32(x),
		uint32(y),
		uint32(width),
		uint32(height),
	)

	return
}

// TrySetScissorRect calls the method "GPURenderPassEncoder.setScissorRect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetScissorRect(x GPUIntegerCoordinate, y GPUIntegerCoordinate, width GPUIntegerCoordinate, height GPUIntegerCoordinate) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetScissorRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(x),
		uint32(y),
		uint32(width),
		uint32(height),
	)

	return
}

// HasSetBlendConstant returns true if the method "GPURenderPassEncoder.setBlendConstant" exists.
func (this GPURenderPassEncoder) HasSetBlendConstant() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetBlendConstant(
		this.Ref(),
	)
}

// SetBlendConstantFunc returns the method "GPURenderPassEncoder.setBlendConstant".
func (this GPURenderPassEncoder) SetBlendConstantFunc() (fn js.Func[func(color GPUColor)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetBlendConstantFunc(
			this.Ref(),
		),
	)
}

// SetBlendConstant calls the method "GPURenderPassEncoder.setBlendConstant".
func (this GPURenderPassEncoder) SetBlendConstant(color GPUColor) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetBlendConstant(
		this.Ref(), js.Pointer(&ret),
		color.Ref(),
	)

	return
}

// TrySetBlendConstant calls the method "GPURenderPassEncoder.setBlendConstant"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetBlendConstant(color GPUColor) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetBlendConstant(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		color.Ref(),
	)

	return
}

// HasSetStencilReference returns true if the method "GPURenderPassEncoder.setStencilReference" exists.
func (this GPURenderPassEncoder) HasSetStencilReference() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetStencilReference(
		this.Ref(),
	)
}

// SetStencilReferenceFunc returns the method "GPURenderPassEncoder.setStencilReference".
func (this GPURenderPassEncoder) SetStencilReferenceFunc() (fn js.Func[func(reference GPUStencilValue)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetStencilReferenceFunc(
			this.Ref(),
		),
	)
}

// SetStencilReference calls the method "GPURenderPassEncoder.setStencilReference".
func (this GPURenderPassEncoder) SetStencilReference(reference GPUStencilValue) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetStencilReference(
		this.Ref(), js.Pointer(&ret),
		uint32(reference),
	)

	return
}

// TrySetStencilReference calls the method "GPURenderPassEncoder.setStencilReference"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetStencilReference(reference GPUStencilValue) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetStencilReference(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(reference),
	)

	return
}

// HasBeginOcclusionQuery returns true if the method "GPURenderPassEncoder.beginOcclusionQuery" exists.
func (this GPURenderPassEncoder) HasBeginOcclusionQuery() bool {
	return js.True == bindings.HasGPURenderPassEncoderBeginOcclusionQuery(
		this.Ref(),
	)
}

// BeginOcclusionQueryFunc returns the method "GPURenderPassEncoder.beginOcclusionQuery".
func (this GPURenderPassEncoder) BeginOcclusionQueryFunc() (fn js.Func[func(queryIndex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderBeginOcclusionQueryFunc(
			this.Ref(),
		),
	)
}

// BeginOcclusionQuery calls the method "GPURenderPassEncoder.beginOcclusionQuery".
func (this GPURenderPassEncoder) BeginOcclusionQuery(queryIndex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderBeginOcclusionQuery(
		this.Ref(), js.Pointer(&ret),
		uint32(queryIndex),
	)

	return
}

// TryBeginOcclusionQuery calls the method "GPURenderPassEncoder.beginOcclusionQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryBeginOcclusionQuery(queryIndex GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderBeginOcclusionQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(queryIndex),
	)

	return
}

// HasEndOcclusionQuery returns true if the method "GPURenderPassEncoder.endOcclusionQuery" exists.
func (this GPURenderPassEncoder) HasEndOcclusionQuery() bool {
	return js.True == bindings.HasGPURenderPassEncoderEndOcclusionQuery(
		this.Ref(),
	)
}

// EndOcclusionQueryFunc returns the method "GPURenderPassEncoder.endOcclusionQuery".
func (this GPURenderPassEncoder) EndOcclusionQueryFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderEndOcclusionQueryFunc(
			this.Ref(),
		),
	)
}

// EndOcclusionQuery calls the method "GPURenderPassEncoder.endOcclusionQuery".
func (this GPURenderPassEncoder) EndOcclusionQuery() (ret js.Void) {
	bindings.CallGPURenderPassEncoderEndOcclusionQuery(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEndOcclusionQuery calls the method "GPURenderPassEncoder.endOcclusionQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryEndOcclusionQuery() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderEndOcclusionQuery(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasExecuteBundles returns true if the method "GPURenderPassEncoder.executeBundles" exists.
func (this GPURenderPassEncoder) HasExecuteBundles() bool {
	return js.True == bindings.HasGPURenderPassEncoderExecuteBundles(
		this.Ref(),
	)
}

// ExecuteBundlesFunc returns the method "GPURenderPassEncoder.executeBundles".
func (this GPURenderPassEncoder) ExecuteBundlesFunc() (fn js.Func[func(bundles js.Array[GPURenderBundle])]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderExecuteBundlesFunc(
			this.Ref(),
		),
	)
}

// ExecuteBundles calls the method "GPURenderPassEncoder.executeBundles".
func (this GPURenderPassEncoder) ExecuteBundles(bundles js.Array[GPURenderBundle]) (ret js.Void) {
	bindings.CallGPURenderPassEncoderExecuteBundles(
		this.Ref(), js.Pointer(&ret),
		bundles.Ref(),
	)

	return
}

// TryExecuteBundles calls the method "GPURenderPassEncoder.executeBundles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryExecuteBundles(bundles js.Array[GPURenderBundle]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderExecuteBundles(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		bundles.Ref(),
	)

	return
}

// HasEnd returns true if the method "GPURenderPassEncoder.end" exists.
func (this GPURenderPassEncoder) HasEnd() bool {
	return js.True == bindings.HasGPURenderPassEncoderEnd(
		this.Ref(),
	)
}

// EndFunc returns the method "GPURenderPassEncoder.end".
func (this GPURenderPassEncoder) EndFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderEndFunc(
			this.Ref(),
		),
	)
}

// End calls the method "GPURenderPassEncoder.end".
func (this GPURenderPassEncoder) End() (ret js.Void) {
	bindings.CallGPURenderPassEncoderEnd(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEnd calls the method "GPURenderPassEncoder.end"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryEnd() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderEnd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetPipeline returns true if the method "GPURenderPassEncoder.setPipeline" exists.
func (this GPURenderPassEncoder) HasSetPipeline() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetPipeline(
		this.Ref(),
	)
}

// SetPipelineFunc returns the method "GPURenderPassEncoder.setPipeline".
func (this GPURenderPassEncoder) SetPipelineFunc() (fn js.Func[func(pipeline GPURenderPipeline)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetPipelineFunc(
			this.Ref(),
		),
	)
}

// SetPipeline calls the method "GPURenderPassEncoder.setPipeline".
func (this GPURenderPassEncoder) SetPipeline(pipeline GPURenderPipeline) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetPipeline(
		this.Ref(), js.Pointer(&ret),
		pipeline.Ref(),
	)

	return
}

// TrySetPipeline calls the method "GPURenderPassEncoder.setPipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetPipeline(pipeline GPURenderPipeline) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetPipeline(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		pipeline.Ref(),
	)

	return
}

// HasSetIndexBuffer returns true if the method "GPURenderPassEncoder.setIndexBuffer" exists.
func (this GPURenderPassEncoder) HasSetIndexBuffer() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetIndexBuffer(
		this.Ref(),
	)
}

// SetIndexBufferFunc returns the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) SetIndexBufferFunc() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetIndexBufferFunc(
			this.Ref(),
		),
	)
}

// SetIndexBuffer calls the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) SetIndexBuffer(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetIndexBuffer(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
		float64(size),
	)

	return
}

// TrySetIndexBuffer calls the method "GPURenderPassEncoder.setIndexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetIndexBuffer(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetIndexBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
		float64(size),
	)

	return
}

// HasSetIndexBuffer1 returns true if the method "GPURenderPassEncoder.setIndexBuffer" exists.
func (this GPURenderPassEncoder) HasSetIndexBuffer1() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetIndexBuffer1(
		this.Ref(),
	)
}

// SetIndexBuffer1Func returns the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) SetIndexBuffer1Func() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetIndexBuffer1Func(
			this.Ref(),
		),
	)
}

// SetIndexBuffer1 calls the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) SetIndexBuffer1(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetIndexBuffer1(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
	)

	return
}

// TrySetIndexBuffer1 calls the method "GPURenderPassEncoder.setIndexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetIndexBuffer1(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetIndexBuffer1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
	)

	return
}

// HasSetIndexBuffer2 returns true if the method "GPURenderPassEncoder.setIndexBuffer" exists.
func (this GPURenderPassEncoder) HasSetIndexBuffer2() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetIndexBuffer2(
		this.Ref(),
	)
}

// SetIndexBuffer2Func returns the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) SetIndexBuffer2Func() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetIndexBuffer2Func(
			this.Ref(),
		),
	)
}

// SetIndexBuffer2 calls the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) SetIndexBuffer2(buffer GPUBuffer, indexFormat GPUIndexFormat) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetIndexBuffer2(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		uint32(indexFormat),
	)

	return
}

// TrySetIndexBuffer2 calls the method "GPURenderPassEncoder.setIndexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetIndexBuffer2(buffer GPUBuffer, indexFormat GPUIndexFormat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetIndexBuffer2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
	)

	return
}

// HasSetVertexBuffer returns true if the method "GPURenderPassEncoder.setVertexBuffer" exists.
func (this GPURenderPassEncoder) HasSetVertexBuffer() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetVertexBuffer(
		this.Ref(),
	)
}

// SetVertexBufferFunc returns the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) SetVertexBufferFunc() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetVertexBufferFunc(
			this.Ref(),
		),
	)
}

// SetVertexBuffer calls the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) SetVertexBuffer(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetVertexBuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// TrySetVertexBuffer calls the method "GPURenderPassEncoder.setVertexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetVertexBuffer(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetVertexBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// HasSetVertexBuffer1 returns true if the method "GPURenderPassEncoder.setVertexBuffer" exists.
func (this GPURenderPassEncoder) HasSetVertexBuffer1() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetVertexBuffer1(
		this.Ref(),
	)
}

// SetVertexBuffer1Func returns the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) SetVertexBuffer1Func() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetVertexBuffer1Func(
			this.Ref(),
		),
	)
}

// SetVertexBuffer1 calls the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) SetVertexBuffer1(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetVertexBuffer1(
		this.Ref(), js.Pointer(&ret),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
	)

	return
}

// TrySetVertexBuffer1 calls the method "GPURenderPassEncoder.setVertexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetVertexBuffer1(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetVertexBuffer1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
	)

	return
}

// HasSetVertexBuffer2 returns true if the method "GPURenderPassEncoder.setVertexBuffer" exists.
func (this GPURenderPassEncoder) HasSetVertexBuffer2() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetVertexBuffer2(
		this.Ref(),
	)
}

// SetVertexBuffer2Func returns the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) SetVertexBuffer2Func() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetVertexBuffer2Func(
			this.Ref(),
		),
	)
}

// SetVertexBuffer2 calls the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) SetVertexBuffer2(slot GPUIndex32, buffer GPUBuffer) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetVertexBuffer2(
		this.Ref(), js.Pointer(&ret),
		uint32(slot),
		buffer.Ref(),
	)

	return
}

// TrySetVertexBuffer2 calls the method "GPURenderPassEncoder.setVertexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetVertexBuffer2(slot GPUIndex32, buffer GPUBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetVertexBuffer2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
	)

	return
}

// HasDraw returns true if the method "GPURenderPassEncoder.draw" exists.
func (this GPURenderPassEncoder) HasDraw() bool {
	return js.True == bindings.HasGPURenderPassEncoderDraw(
		this.Ref(),
	)
}

// DrawFunc returns the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) DrawFunc() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawFunc(
			this.Ref(),
		),
	)
}

// Draw calls the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDraw(
		this.Ref(), js.Pointer(&ret),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
		uint32(firstInstance),
	)

	return
}

// TryDraw calls the method "GPURenderPassEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDraw(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDraw(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
		uint32(firstInstance),
	)

	return
}

// HasDraw1 returns true if the method "GPURenderPassEncoder.draw" exists.
func (this GPURenderPassEncoder) HasDraw1() bool {
	return js.True == bindings.HasGPURenderPassEncoderDraw1(
		this.Ref(),
	)
}

// Draw1Func returns the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw1Func() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDraw1Func(
			this.Ref(),
		),
	)
}

// Draw1 calls the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw1(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDraw1(
		this.Ref(), js.Pointer(&ret),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
	)

	return
}

// TryDraw1 calls the method "GPURenderPassEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDraw1(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDraw1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
	)

	return
}

// HasDraw2 returns true if the method "GPURenderPassEncoder.draw" exists.
func (this GPURenderPassEncoder) HasDraw2() bool {
	return js.True == bindings.HasGPURenderPassEncoderDraw2(
		this.Ref(),
	)
}

// Draw2Func returns the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw2Func() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDraw2Func(
			this.Ref(),
		),
	)
}

// Draw2 calls the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw2(vertexCount GPUSize32, instanceCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDraw2(
		this.Ref(), js.Pointer(&ret),
		uint32(vertexCount),
		uint32(instanceCount),
	)

	return
}

// TryDraw2 calls the method "GPURenderPassEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDraw2(vertexCount GPUSize32, instanceCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDraw2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
	)

	return
}

// HasDraw3 returns true if the method "GPURenderPassEncoder.draw" exists.
func (this GPURenderPassEncoder) HasDraw3() bool {
	return js.True == bindings.HasGPURenderPassEncoderDraw3(
		this.Ref(),
	)
}

// Draw3Func returns the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw3Func() (fn js.Func[func(vertexCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDraw3Func(
			this.Ref(),
		),
	)
}

// Draw3 calls the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw3(vertexCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDraw3(
		this.Ref(), js.Pointer(&ret),
		uint32(vertexCount),
	)

	return
}

// TryDraw3 calls the method "GPURenderPassEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDraw3(vertexCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDraw3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
	)

	return
}

// HasDrawIndexed returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasDrawIndexed() bool {
	return js.True == bindings.HasGPURenderPassEncoderDrawIndexed(
		this.Ref(),
	)
}

// DrawIndexedFunc returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexedFunc() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexedFunc(
			this.Ref(),
		),
	)
}

// DrawIndexed calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
		uint32(firstInstance),
	)

	return
}

// TryDrawIndexed calls the method "GPURenderPassEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDrawIndexed(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDrawIndexed(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
		uint32(firstInstance),
	)

	return
}

// HasDrawIndexed1 returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasDrawIndexed1() bool {
	return js.True == bindings.HasGPURenderPassEncoderDrawIndexed1(
		this.Ref(),
	)
}

// DrawIndexed1Func returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed1Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexed1Func(
			this.Ref(),
		),
	)
}

// DrawIndexed1 calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed1(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed1(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
	)

	return
}

// TryDrawIndexed1 calls the method "GPURenderPassEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDrawIndexed1(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDrawIndexed1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
	)

	return
}

// HasDrawIndexed2 returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasDrawIndexed2() bool {
	return js.True == bindings.HasGPURenderPassEncoderDrawIndexed2(
		this.Ref(),
	)
}

// DrawIndexed2Func returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed2Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexed2Func(
			this.Ref(),
		),
	)
}

// DrawIndexed2 calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed2(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed2(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
	)

	return
}

// TryDrawIndexed2 calls the method "GPURenderPassEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDrawIndexed2(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDrawIndexed2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
	)

	return
}

// HasDrawIndexed3 returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasDrawIndexed3() bool {
	return js.True == bindings.HasGPURenderPassEncoderDrawIndexed3(
		this.Ref(),
	)
}

// DrawIndexed3Func returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed3Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexed3Func(
			this.Ref(),
		),
	)
}

// DrawIndexed3 calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed3(indexCount GPUSize32, instanceCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed3(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
		uint32(instanceCount),
	)

	return
}

// TryDrawIndexed3 calls the method "GPURenderPassEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDrawIndexed3(indexCount GPUSize32, instanceCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDrawIndexed3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
	)

	return
}

// HasDrawIndexed4 returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasDrawIndexed4() bool {
	return js.True == bindings.HasGPURenderPassEncoderDrawIndexed4(
		this.Ref(),
	)
}

// DrawIndexed4Func returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed4Func() (fn js.Func[func(indexCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexed4Func(
			this.Ref(),
		),
	)
}

// DrawIndexed4 calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed4(indexCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed4(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
	)

	return
}

// TryDrawIndexed4 calls the method "GPURenderPassEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDrawIndexed4(indexCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDrawIndexed4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
	)

	return
}

// HasDrawIndirect returns true if the method "GPURenderPassEncoder.drawIndirect" exists.
func (this GPURenderPassEncoder) HasDrawIndirect() bool {
	return js.True == bindings.HasGPURenderPassEncoderDrawIndirect(
		this.Ref(),
	)
}

// DrawIndirectFunc returns the method "GPURenderPassEncoder.drawIndirect".
func (this GPURenderPassEncoder) DrawIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndirectFunc(
			this.Ref(),
		),
	)
}

// DrawIndirect calls the method "GPURenderPassEncoder.drawIndirect".
func (this GPURenderPassEncoder) DrawIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndirect(
		this.Ref(), js.Pointer(&ret),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// TryDrawIndirect calls the method "GPURenderPassEncoder.drawIndirect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDrawIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDrawIndirect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasDrawIndexedIndirect returns true if the method "GPURenderPassEncoder.drawIndexedIndirect" exists.
func (this GPURenderPassEncoder) HasDrawIndexedIndirect() bool {
	return js.True == bindings.HasGPURenderPassEncoderDrawIndexedIndirect(
		this.Ref(),
	)
}

// DrawIndexedIndirectFunc returns the method "GPURenderPassEncoder.drawIndexedIndirect".
func (this GPURenderPassEncoder) DrawIndexedIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexedIndirectFunc(
			this.Ref(),
		),
	)
}

// DrawIndexedIndirect calls the method "GPURenderPassEncoder.drawIndexedIndirect".
func (this GPURenderPassEncoder) DrawIndexedIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexedIndirect(
		this.Ref(), js.Pointer(&ret),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// TryDrawIndexedIndirect calls the method "GPURenderPassEncoder.drawIndexedIndirect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDrawIndexedIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDrawIndexedIndirect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasSetBindGroup returns true if the method "GPURenderPassEncoder.setBindGroup" exists.
func (this GPURenderPassEncoder) HasSetBindGroup() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetBindGroup(
		this.Ref(),
	)
}

// SetBindGroupFunc returns the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) SetBindGroupFunc() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset])]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetBindGroupFunc(
			this.Ref(),
		),
	)
}

// SetBindGroup calls the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetBindGroup(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	return
}

// TrySetBindGroup calls the method "GPURenderPassEncoder.setBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetBindGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	return
}

// HasSetBindGroup1 returns true if the method "GPURenderPassEncoder.setBindGroup" exists.
func (this GPURenderPassEncoder) HasSetBindGroup1() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetBindGroup1(
		this.Ref(),
	)
}

// SetBindGroup1Func returns the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) SetBindGroup1Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetBindGroup1Func(
			this.Ref(),
		),
	)
}

// SetBindGroup1 calls the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) SetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetBindGroup1(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		bindGroup.Ref(),
	)

	return
}

// TrySetBindGroup1 calls the method "GPURenderPassEncoder.setBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetBindGroup1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
	)

	return
}

// HasSetBindGroup2 returns true if the method "GPURenderPassEncoder.setBindGroup" exists.
func (this GPURenderPassEncoder) HasSetBindGroup2() bool {
	return js.True == bindings.HasGPURenderPassEncoderSetBindGroup2(
		this.Ref(),
	)
}

// SetBindGroup2Func returns the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) SetBindGroup2Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetBindGroup2Func(
			this.Ref(),
		),
	)
}

// SetBindGroup2 calls the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) SetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetBindGroup2(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	return
}

// TrySetBindGroup2 calls the method "GPURenderPassEncoder.setBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetBindGroup2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	return
}

// HasPushDebugGroup returns true if the method "GPURenderPassEncoder.pushDebugGroup" exists.
func (this GPURenderPassEncoder) HasPushDebugGroup() bool {
	return js.True == bindings.HasGPURenderPassEncoderPushDebugGroup(
		this.Ref(),
	)
}

// PushDebugGroupFunc returns the method "GPURenderPassEncoder.pushDebugGroup".
func (this GPURenderPassEncoder) PushDebugGroupFunc() (fn js.Func[func(groupLabel js.String)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderPushDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PushDebugGroup calls the method "GPURenderPassEncoder.pushDebugGroup".
func (this GPURenderPassEncoder) PushDebugGroup(groupLabel js.String) (ret js.Void) {
	bindings.CallGPURenderPassEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&ret),
		groupLabel.Ref(),
	)

	return
}

// TryPushDebugGroup calls the method "GPURenderPassEncoder.pushDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryPushDebugGroup(groupLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		groupLabel.Ref(),
	)

	return
}

// HasPopDebugGroup returns true if the method "GPURenderPassEncoder.popDebugGroup" exists.
func (this GPURenderPassEncoder) HasPopDebugGroup() bool {
	return js.True == bindings.HasGPURenderPassEncoderPopDebugGroup(
		this.Ref(),
	)
}

// PopDebugGroupFunc returns the method "GPURenderPassEncoder.popDebugGroup".
func (this GPURenderPassEncoder) PopDebugGroupFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderPopDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PopDebugGroup calls the method "GPURenderPassEncoder.popDebugGroup".
func (this GPURenderPassEncoder) PopDebugGroup() (ret js.Void) {
	bindings.CallGPURenderPassEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPopDebugGroup calls the method "GPURenderPassEncoder.popDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryPopDebugGroup() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInsertDebugMarker returns true if the method "GPURenderPassEncoder.insertDebugMarker" exists.
func (this GPURenderPassEncoder) HasInsertDebugMarker() bool {
	return js.True == bindings.HasGPURenderPassEncoderInsertDebugMarker(
		this.Ref(),
	)
}

// InsertDebugMarkerFunc returns the method "GPURenderPassEncoder.insertDebugMarker".
func (this GPURenderPassEncoder) InsertDebugMarkerFunc() (fn js.Func[func(markerLabel js.String)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderInsertDebugMarkerFunc(
			this.Ref(),
		),
	)
}

// InsertDebugMarker calls the method "GPURenderPassEncoder.insertDebugMarker".
func (this GPURenderPassEncoder) InsertDebugMarker(markerLabel js.String) (ret js.Void) {
	bindings.CallGPURenderPassEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&ret),
		markerLabel.Ref(),
	)

	return
}

// TryInsertDebugMarker calls the method "GPURenderPassEncoder.insertDebugMarker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryInsertDebugMarker(markerLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		markerLabel.Ref(),
	)

	return
}

type GPULoadOp uint32

const (
	_ GPULoadOp = iota

	GPULoadOp_LOAD
	GPULoadOp_CLEAR
)

func (GPULoadOp) FromRef(str js.Ref) GPULoadOp {
	return GPULoadOp(bindings.ConstOfGPULoadOp(str))
}

func (x GPULoadOp) String() (string, bool) {
	switch x {
	case GPULoadOp_LOAD:
		return "load", true
	case GPULoadOp_CLEAR:
		return "clear", true
	default:
		return "", false
	}
}

type GPUStoreOp uint32

const (
	_ GPUStoreOp = iota

	GPUStoreOp_STORE
	GPUStoreOp_DISCARD
)

func (GPUStoreOp) FromRef(str js.Ref) GPUStoreOp {
	return GPUStoreOp(bindings.ConstOfGPUStoreOp(str))
}

func (x GPUStoreOp) String() (string, bool) {
	switch x {
	case GPUStoreOp_STORE:
		return "store", true
	case GPUStoreOp_DISCARD:
		return "discard", true
	default:
		return "", false
	}
}

type GPURenderPassColorAttachment struct {
	// View is "GPURenderPassColorAttachment.view"
	//
	// Required
	View GPUTextureView
	// ResolveTarget is "GPURenderPassColorAttachment.resolveTarget"
	//
	// Optional
	ResolveTarget GPUTextureView
	// ClearValue is "GPURenderPassColorAttachment.clearValue"
	//
	// Optional
	ClearValue GPUColor
	// LoadOp is "GPURenderPassColorAttachment.loadOp"
	//
	// Required
	LoadOp GPULoadOp
	// StoreOp is "GPURenderPassColorAttachment.storeOp"
	//
	// Required
	StoreOp GPUStoreOp

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURenderPassColorAttachment with all fields set.
func (p GPURenderPassColorAttachment) FromRef(ref js.Ref) GPURenderPassColorAttachment {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPURenderPassColorAttachment in the application heap.
func (p GPURenderPassColorAttachment) New() js.Ref {
	return bindings.GPURenderPassColorAttachmentJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPURenderPassColorAttachment) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassColorAttachmentJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPURenderPassColorAttachment) Update(ref js.Ref) {
	bindings.GPURenderPassColorAttachmentJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPURenderPassDepthStencilAttachment struct {
	// View is "GPURenderPassDepthStencilAttachment.view"
	//
	// Required
	View GPUTextureView
	// DepthClearValue is "GPURenderPassDepthStencilAttachment.depthClearValue"
	//
	// Optional
	//
	// NOTE: FFI_USE_DepthClearValue MUST be set to true to make this field effective.
	DepthClearValue float32
	// DepthLoadOp is "GPURenderPassDepthStencilAttachment.depthLoadOp"
	//
	// Optional
	DepthLoadOp GPULoadOp
	// DepthStoreOp is "GPURenderPassDepthStencilAttachment.depthStoreOp"
	//
	// Optional
	DepthStoreOp GPUStoreOp
	// DepthReadOnly is "GPURenderPassDepthStencilAttachment.depthReadOnly"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_DepthReadOnly MUST be set to true to make this field effective.
	DepthReadOnly bool
	// StencilClearValue is "GPURenderPassDepthStencilAttachment.stencilClearValue"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_StencilClearValue MUST be set to true to make this field effective.
	StencilClearValue GPUStencilValue
	// StencilLoadOp is "GPURenderPassDepthStencilAttachment.stencilLoadOp"
	//
	// Optional
	StencilLoadOp GPULoadOp
	// StencilStoreOp is "GPURenderPassDepthStencilAttachment.stencilStoreOp"
	//
	// Optional
	StencilStoreOp GPUStoreOp
	// StencilReadOnly is "GPURenderPassDepthStencilAttachment.stencilReadOnly"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_StencilReadOnly MUST be set to true to make this field effective.
	StencilReadOnly bool

	FFI_USE_DepthClearValue   bool // for DepthClearValue.
	FFI_USE_DepthReadOnly     bool // for DepthReadOnly.
	FFI_USE_StencilClearValue bool // for StencilClearValue.
	FFI_USE_StencilReadOnly   bool // for StencilReadOnly.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURenderPassDepthStencilAttachment with all fields set.
func (p GPURenderPassDepthStencilAttachment) FromRef(ref js.Ref) GPURenderPassDepthStencilAttachment {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPURenderPassDepthStencilAttachment in the application heap.
func (p GPURenderPassDepthStencilAttachment) New() js.Ref {
	return bindings.GPURenderPassDepthStencilAttachmentJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPURenderPassDepthStencilAttachment) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassDepthStencilAttachmentJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPURenderPassDepthStencilAttachment) Update(ref js.Ref) {
	bindings.GPURenderPassDepthStencilAttachmentJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUQueryType uint32

const (
	_ GPUQueryType = iota

	GPUQueryType_OCCLUSION
	GPUQueryType_TIMESTAMP
)

func (GPUQueryType) FromRef(str js.Ref) GPUQueryType {
	return GPUQueryType(bindings.ConstOfGPUQueryType(str))
}

func (x GPUQueryType) String() (string, bool) {
	switch x {
	case GPUQueryType_OCCLUSION:
		return "occlusion", true
	case GPUQueryType_TIMESTAMP:
		return "timestamp", true
	default:
		return "", false
	}
}

type GPUQuerySet struct {
	ref js.Ref
}

func (this GPUQuerySet) Once() GPUQuerySet {
	this.Ref().Once()
	return this
}

func (this GPUQuerySet) Ref() js.Ref {
	return this.ref
}

func (this GPUQuerySet) FromRef(ref js.Ref) GPUQuerySet {
	this.ref = ref
	return this
}

func (this GPUQuerySet) Free() {
	this.Ref().Free()
}

// Type returns the value of property "GPUQuerySet.type".
//
// It returns ok=false if there is no such property.
func (this GPUQuerySet) Type() (ret GPUQueryType, ok bool) {
	ok = js.True == bindings.GetGPUQuerySetType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Count returns the value of property "GPUQuerySet.count".
//
// It returns ok=false if there is no such property.
func (this GPUQuerySet) Count() (ret GPUSize32Out, ok bool) {
	ok = js.True == bindings.GetGPUQuerySetCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "GPUQuerySet.label".
//
// It returns ok=false if there is no such property.
func (this GPUQuerySet) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUQuerySetLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUQuerySet.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUQuerySet) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUQuerySetLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasDestroy returns true if the method "GPUQuerySet.destroy" exists.
func (this GPUQuerySet) HasDestroy() bool {
	return js.True == bindings.HasGPUQuerySetDestroy(
		this.Ref(),
	)
}

// DestroyFunc returns the method "GPUQuerySet.destroy".
func (this GPUQuerySet) DestroyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUQuerySetDestroyFunc(
			this.Ref(),
		),
	)
}

// Destroy calls the method "GPUQuerySet.destroy".
func (this GPUQuerySet) Destroy() (ret js.Void) {
	bindings.CallGPUQuerySetDestroy(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDestroy calls the method "GPUQuerySet.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQuerySet) TryDestroy() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQuerySetDestroy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type GPURenderPassTimestampWrites struct {
	// QuerySet is "GPURenderPassTimestampWrites.querySet"
	//
	// Required
	QuerySet GPUQuerySet
	// BeginningOfPassWriteIndex is "GPURenderPassTimestampWrites.beginningOfPassWriteIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_BeginningOfPassWriteIndex MUST be set to true to make this field effective.
	BeginningOfPassWriteIndex GPUSize32
	// EndOfPassWriteIndex is "GPURenderPassTimestampWrites.endOfPassWriteIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_EndOfPassWriteIndex MUST be set to true to make this field effective.
	EndOfPassWriteIndex GPUSize32

	FFI_USE_BeginningOfPassWriteIndex bool // for BeginningOfPassWriteIndex.
	FFI_USE_EndOfPassWriteIndex       bool // for EndOfPassWriteIndex.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURenderPassTimestampWrites with all fields set.
func (p GPURenderPassTimestampWrites) FromRef(ref js.Ref) GPURenderPassTimestampWrites {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPURenderPassTimestampWrites in the application heap.
func (p GPURenderPassTimestampWrites) New() js.Ref {
	return bindings.GPURenderPassTimestampWritesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPURenderPassTimestampWrites) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassTimestampWritesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPURenderPassTimestampWrites) Update(ref js.Ref) {
	bindings.GPURenderPassTimestampWritesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPURenderPassDescriptor struct {
	// ColorAttachments is "GPURenderPassDescriptor.colorAttachments"
	//
	// Required
	ColorAttachments js.Array[GPURenderPassColorAttachment]
	// DepthStencilAttachment is "GPURenderPassDescriptor.depthStencilAttachment"
	//
	// Optional
	//
	// NOTE: DepthStencilAttachment.FFI_USE MUST be set to true to get DepthStencilAttachment used.
	DepthStencilAttachment GPURenderPassDepthStencilAttachment
	// OcclusionQuerySet is "GPURenderPassDescriptor.occlusionQuerySet"
	//
	// Optional
	OcclusionQuerySet GPUQuerySet
	// TimestampWrites is "GPURenderPassDescriptor.timestampWrites"
	//
	// Optional
	//
	// NOTE: TimestampWrites.FFI_USE MUST be set to true to get TimestampWrites used.
	TimestampWrites GPURenderPassTimestampWrites
	// MaxDrawCount is "GPURenderPassDescriptor.maxDrawCount"
	//
	// Optional, defaults to 50000000.
	//
	// NOTE: FFI_USE_MaxDrawCount MUST be set to true to make this field effective.
	MaxDrawCount GPUSize64
	// Label is "GPURenderPassDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE_MaxDrawCount bool // for MaxDrawCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURenderPassDescriptor with all fields set.
func (p GPURenderPassDescriptor) FromRef(ref js.Ref) GPURenderPassDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPURenderPassDescriptor in the application heap.
func (p GPURenderPassDescriptor) New() js.Ref {
	return bindings.GPURenderPassDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPURenderPassDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPURenderPassDescriptor) Update(ref js.Ref) {
	bindings.GPURenderPassDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUComputePassEncoder struct {
	ref js.Ref
}

func (this GPUComputePassEncoder) Once() GPUComputePassEncoder {
	this.Ref().Once()
	return this
}

func (this GPUComputePassEncoder) Ref() js.Ref {
	return this.ref
}

func (this GPUComputePassEncoder) FromRef(ref js.Ref) GPUComputePassEncoder {
	this.ref = ref
	return this
}

func (this GPUComputePassEncoder) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUComputePassEncoder.label".
//
// It returns ok=false if there is no such property.
func (this GPUComputePassEncoder) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUComputePassEncoderLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUComputePassEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUComputePassEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUComputePassEncoderLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasSetPipeline returns true if the method "GPUComputePassEncoder.setPipeline" exists.
func (this GPUComputePassEncoder) HasSetPipeline() bool {
	return js.True == bindings.HasGPUComputePassEncoderSetPipeline(
		this.Ref(),
	)
}

// SetPipelineFunc returns the method "GPUComputePassEncoder.setPipeline".
func (this GPUComputePassEncoder) SetPipelineFunc() (fn js.Func[func(pipeline GPUComputePipeline)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderSetPipelineFunc(
			this.Ref(),
		),
	)
}

// SetPipeline calls the method "GPUComputePassEncoder.setPipeline".
func (this GPUComputePassEncoder) SetPipeline(pipeline GPUComputePipeline) (ret js.Void) {
	bindings.CallGPUComputePassEncoderSetPipeline(
		this.Ref(), js.Pointer(&ret),
		pipeline.Ref(),
	)

	return
}

// TrySetPipeline calls the method "GPUComputePassEncoder.setPipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TrySetPipeline(pipeline GPUComputePipeline) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderSetPipeline(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		pipeline.Ref(),
	)

	return
}

// HasDispatchWorkgroups returns true if the method "GPUComputePassEncoder.dispatchWorkgroups" exists.
func (this GPUComputePassEncoder) HasDispatchWorkgroups() bool {
	return js.True == bindings.HasGPUComputePassEncoderDispatchWorkgroups(
		this.Ref(),
	)
}

// DispatchWorkgroupsFunc returns the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) DispatchWorkgroupsFunc() (fn js.Func[func(workgroupCountX GPUSize32, workgroupCountY GPUSize32, workgroupCountZ GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderDispatchWorkgroupsFunc(
			this.Ref(),
		),
	)
}

// DispatchWorkgroups calls the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) DispatchWorkgroups(workgroupCountX GPUSize32, workgroupCountY GPUSize32, workgroupCountZ GPUSize32) (ret js.Void) {
	bindings.CallGPUComputePassEncoderDispatchWorkgroups(
		this.Ref(), js.Pointer(&ret),
		uint32(workgroupCountX),
		uint32(workgroupCountY),
		uint32(workgroupCountZ),
	)

	return
}

// TryDispatchWorkgroups calls the method "GPUComputePassEncoder.dispatchWorkgroups"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryDispatchWorkgroups(workgroupCountX GPUSize32, workgroupCountY GPUSize32, workgroupCountZ GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderDispatchWorkgroups(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(workgroupCountX),
		uint32(workgroupCountY),
		uint32(workgroupCountZ),
	)

	return
}

// HasDispatchWorkgroups1 returns true if the method "GPUComputePassEncoder.dispatchWorkgroups" exists.
func (this GPUComputePassEncoder) HasDispatchWorkgroups1() bool {
	return js.True == bindings.HasGPUComputePassEncoderDispatchWorkgroups1(
		this.Ref(),
	)
}

// DispatchWorkgroups1Func returns the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) DispatchWorkgroups1Func() (fn js.Func[func(workgroupCountX GPUSize32, workgroupCountY GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderDispatchWorkgroups1Func(
			this.Ref(),
		),
	)
}

// DispatchWorkgroups1 calls the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) DispatchWorkgroups1(workgroupCountX GPUSize32, workgroupCountY GPUSize32) (ret js.Void) {
	bindings.CallGPUComputePassEncoderDispatchWorkgroups1(
		this.Ref(), js.Pointer(&ret),
		uint32(workgroupCountX),
		uint32(workgroupCountY),
	)

	return
}

// TryDispatchWorkgroups1 calls the method "GPUComputePassEncoder.dispatchWorkgroups"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryDispatchWorkgroups1(workgroupCountX GPUSize32, workgroupCountY GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderDispatchWorkgroups1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(workgroupCountX),
		uint32(workgroupCountY),
	)

	return
}

// HasDispatchWorkgroups2 returns true if the method "GPUComputePassEncoder.dispatchWorkgroups" exists.
func (this GPUComputePassEncoder) HasDispatchWorkgroups2() bool {
	return js.True == bindings.HasGPUComputePassEncoderDispatchWorkgroups2(
		this.Ref(),
	)
}

// DispatchWorkgroups2Func returns the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) DispatchWorkgroups2Func() (fn js.Func[func(workgroupCountX GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderDispatchWorkgroups2Func(
			this.Ref(),
		),
	)
}

// DispatchWorkgroups2 calls the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) DispatchWorkgroups2(workgroupCountX GPUSize32) (ret js.Void) {
	bindings.CallGPUComputePassEncoderDispatchWorkgroups2(
		this.Ref(), js.Pointer(&ret),
		uint32(workgroupCountX),
	)

	return
}

// TryDispatchWorkgroups2 calls the method "GPUComputePassEncoder.dispatchWorkgroups"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryDispatchWorkgroups2(workgroupCountX GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderDispatchWorkgroups2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(workgroupCountX),
	)

	return
}

// HasDispatchWorkgroupsIndirect returns true if the method "GPUComputePassEncoder.dispatchWorkgroupsIndirect" exists.
func (this GPUComputePassEncoder) HasDispatchWorkgroupsIndirect() bool {
	return js.True == bindings.HasGPUComputePassEncoderDispatchWorkgroupsIndirect(
		this.Ref(),
	)
}

// DispatchWorkgroupsIndirectFunc returns the method "GPUComputePassEncoder.dispatchWorkgroupsIndirect".
func (this GPUComputePassEncoder) DispatchWorkgroupsIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderDispatchWorkgroupsIndirectFunc(
			this.Ref(),
		),
	)
}

// DispatchWorkgroupsIndirect calls the method "GPUComputePassEncoder.dispatchWorkgroupsIndirect".
func (this GPUComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPUComputePassEncoderDispatchWorkgroupsIndirect(
		this.Ref(), js.Pointer(&ret),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// TryDispatchWorkgroupsIndirect calls the method "GPUComputePassEncoder.dispatchWorkgroupsIndirect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryDispatchWorkgroupsIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderDispatchWorkgroupsIndirect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasEnd returns true if the method "GPUComputePassEncoder.end" exists.
func (this GPUComputePassEncoder) HasEnd() bool {
	return js.True == bindings.HasGPUComputePassEncoderEnd(
		this.Ref(),
	)
}

// EndFunc returns the method "GPUComputePassEncoder.end".
func (this GPUComputePassEncoder) EndFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderEndFunc(
			this.Ref(),
		),
	)
}

// End calls the method "GPUComputePassEncoder.end".
func (this GPUComputePassEncoder) End() (ret js.Void) {
	bindings.CallGPUComputePassEncoderEnd(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEnd calls the method "GPUComputePassEncoder.end"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryEnd() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderEnd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPushDebugGroup returns true if the method "GPUComputePassEncoder.pushDebugGroup" exists.
func (this GPUComputePassEncoder) HasPushDebugGroup() bool {
	return js.True == bindings.HasGPUComputePassEncoderPushDebugGroup(
		this.Ref(),
	)
}

// PushDebugGroupFunc returns the method "GPUComputePassEncoder.pushDebugGroup".
func (this GPUComputePassEncoder) PushDebugGroupFunc() (fn js.Func[func(groupLabel js.String)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderPushDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PushDebugGroup calls the method "GPUComputePassEncoder.pushDebugGroup".
func (this GPUComputePassEncoder) PushDebugGroup(groupLabel js.String) (ret js.Void) {
	bindings.CallGPUComputePassEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&ret),
		groupLabel.Ref(),
	)

	return
}

// TryPushDebugGroup calls the method "GPUComputePassEncoder.pushDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryPushDebugGroup(groupLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		groupLabel.Ref(),
	)

	return
}

// HasPopDebugGroup returns true if the method "GPUComputePassEncoder.popDebugGroup" exists.
func (this GPUComputePassEncoder) HasPopDebugGroup() bool {
	return js.True == bindings.HasGPUComputePassEncoderPopDebugGroup(
		this.Ref(),
	)
}

// PopDebugGroupFunc returns the method "GPUComputePassEncoder.popDebugGroup".
func (this GPUComputePassEncoder) PopDebugGroupFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderPopDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PopDebugGroup calls the method "GPUComputePassEncoder.popDebugGroup".
func (this GPUComputePassEncoder) PopDebugGroup() (ret js.Void) {
	bindings.CallGPUComputePassEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPopDebugGroup calls the method "GPUComputePassEncoder.popDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryPopDebugGroup() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInsertDebugMarker returns true if the method "GPUComputePassEncoder.insertDebugMarker" exists.
func (this GPUComputePassEncoder) HasInsertDebugMarker() bool {
	return js.True == bindings.HasGPUComputePassEncoderInsertDebugMarker(
		this.Ref(),
	)
}

// InsertDebugMarkerFunc returns the method "GPUComputePassEncoder.insertDebugMarker".
func (this GPUComputePassEncoder) InsertDebugMarkerFunc() (fn js.Func[func(markerLabel js.String)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderInsertDebugMarkerFunc(
			this.Ref(),
		),
	)
}

// InsertDebugMarker calls the method "GPUComputePassEncoder.insertDebugMarker".
func (this GPUComputePassEncoder) InsertDebugMarker(markerLabel js.String) (ret js.Void) {
	bindings.CallGPUComputePassEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&ret),
		markerLabel.Ref(),
	)

	return
}

// TryInsertDebugMarker calls the method "GPUComputePassEncoder.insertDebugMarker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryInsertDebugMarker(markerLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		markerLabel.Ref(),
	)

	return
}

// HasSetBindGroup returns true if the method "GPUComputePassEncoder.setBindGroup" exists.
func (this GPUComputePassEncoder) HasSetBindGroup() bool {
	return js.True == bindings.HasGPUComputePassEncoderSetBindGroup(
		this.Ref(),
	)
}

// SetBindGroupFunc returns the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) SetBindGroupFunc() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset])]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderSetBindGroupFunc(
			this.Ref(),
		),
	)
}

// SetBindGroup calls the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (ret js.Void) {
	bindings.CallGPUComputePassEncoderSetBindGroup(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	return
}

// TrySetBindGroup calls the method "GPUComputePassEncoder.setBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TrySetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderSetBindGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	return
}

// HasSetBindGroup1 returns true if the method "GPUComputePassEncoder.setBindGroup" exists.
func (this GPUComputePassEncoder) HasSetBindGroup1() bool {
	return js.True == bindings.HasGPUComputePassEncoderSetBindGroup1(
		this.Ref(),
	)
}

// SetBindGroup1Func returns the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) SetBindGroup1Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderSetBindGroup1Func(
			this.Ref(),
		),
	)
}

// SetBindGroup1 calls the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) SetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (ret js.Void) {
	bindings.CallGPUComputePassEncoderSetBindGroup1(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		bindGroup.Ref(),
	)

	return
}

// TrySetBindGroup1 calls the method "GPUComputePassEncoder.setBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TrySetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderSetBindGroup1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
	)

	return
}

// HasSetBindGroup2 returns true if the method "GPUComputePassEncoder.setBindGroup" exists.
func (this GPUComputePassEncoder) HasSetBindGroup2() bool {
	return js.True == bindings.HasGPUComputePassEncoderSetBindGroup2(
		this.Ref(),
	)
}

// SetBindGroup2Func returns the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) SetBindGroup2Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderSetBindGroup2Func(
			this.Ref(),
		),
	)
}

// SetBindGroup2 calls the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) SetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (ret js.Void) {
	bindings.CallGPUComputePassEncoderSetBindGroup2(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	return
}

// TrySetBindGroup2 calls the method "GPUComputePassEncoder.setBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TrySetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderSetBindGroup2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	return
}

type GPUComputePassTimestampWrites struct {
	// QuerySet is "GPUComputePassTimestampWrites.querySet"
	//
	// Required
	QuerySet GPUQuerySet
	// BeginningOfPassWriteIndex is "GPUComputePassTimestampWrites.beginningOfPassWriteIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_BeginningOfPassWriteIndex MUST be set to true to make this field effective.
	BeginningOfPassWriteIndex GPUSize32
	// EndOfPassWriteIndex is "GPUComputePassTimestampWrites.endOfPassWriteIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_EndOfPassWriteIndex MUST be set to true to make this field effective.
	EndOfPassWriteIndex GPUSize32

	FFI_USE_BeginningOfPassWriteIndex bool // for BeginningOfPassWriteIndex.
	FFI_USE_EndOfPassWriteIndex       bool // for EndOfPassWriteIndex.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUComputePassTimestampWrites with all fields set.
func (p GPUComputePassTimestampWrites) FromRef(ref js.Ref) GPUComputePassTimestampWrites {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUComputePassTimestampWrites in the application heap.
func (p GPUComputePassTimestampWrites) New() js.Ref {
	return bindings.GPUComputePassTimestampWritesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUComputePassTimestampWrites) UpdateFrom(ref js.Ref) {
	bindings.GPUComputePassTimestampWritesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUComputePassTimestampWrites) Update(ref js.Ref) {
	bindings.GPUComputePassTimestampWritesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUComputePassDescriptor struct {
	// TimestampWrites is "GPUComputePassDescriptor.timestampWrites"
	//
	// Optional
	//
	// NOTE: TimestampWrites.FFI_USE MUST be set to true to get TimestampWrites used.
	TimestampWrites GPUComputePassTimestampWrites
	// Label is "GPUComputePassDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUComputePassDescriptor with all fields set.
func (p GPUComputePassDescriptor) FromRef(ref js.Ref) GPUComputePassDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUComputePassDescriptor in the application heap.
func (p GPUComputePassDescriptor) New() js.Ref {
	return bindings.GPUComputePassDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUComputePassDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUComputePassDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUComputePassDescriptor) Update(ref js.Ref) {
	bindings.GPUComputePassDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUImageCopyBuffer struct {
	// Buffer is "GPUImageCopyBuffer.buffer"
	//
	// Required
	Buffer GPUBuffer
	// Offset is "GPUImageCopyBuffer.offset"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset GPUSize64
	// BytesPerRow is "GPUImageCopyBuffer.bytesPerRow"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesPerRow MUST be set to true to make this field effective.
	BytesPerRow GPUSize32
	// RowsPerImage is "GPUImageCopyBuffer.rowsPerImage"
	//
	// Optional
	//
	// NOTE: FFI_USE_RowsPerImage MUST be set to true to make this field effective.
	RowsPerImage GPUSize32

	FFI_USE_Offset       bool // for Offset.
	FFI_USE_BytesPerRow  bool // for BytesPerRow.
	FFI_USE_RowsPerImage bool // for RowsPerImage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUImageCopyBuffer with all fields set.
func (p GPUImageCopyBuffer) FromRef(ref js.Ref) GPUImageCopyBuffer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUImageCopyBuffer in the application heap.
func (p GPUImageCopyBuffer) New() js.Ref {
	return bindings.GPUImageCopyBufferJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUImageCopyBuffer) UpdateFrom(ref js.Ref) {
	bindings.GPUImageCopyBufferJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUImageCopyBuffer) Update(ref js.Ref) {
	bindings.GPUImageCopyBufferJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUOrigin3DDict struct {
	// X is "GPUOrigin3DDict.x"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X GPUIntegerCoordinate
	// Y is "GPUOrigin3DDict.y"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y GPUIntegerCoordinate
	// Z is "GPUOrigin3DDict.z"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Z MUST be set to true to make this field effective.
	Z GPUIntegerCoordinate

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.
	FFI_USE_Z bool // for Z.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUOrigin3DDict with all fields set.
func (p GPUOrigin3DDict) FromRef(ref js.Ref) GPUOrigin3DDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUOrigin3DDict in the application heap.
func (p GPUOrigin3DDict) New() js.Ref {
	return bindings.GPUOrigin3DDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUOrigin3DDict) UpdateFrom(ref js.Ref) {
	bindings.GPUOrigin3DDictJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUOrigin3DDict) Update(ref js.Ref) {
	bindings.GPUOrigin3DDictJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_ArrayGPUIntegerCoordinate_GPUOrigin3DDict struct {
	ref js.Ref
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin3DDict) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin3DDict) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin3DDict) FromRef(ref js.Ref) OneOf_ArrayGPUIntegerCoordinate_GPUOrigin3DDict {
	return OneOf_ArrayGPUIntegerCoordinate_GPUOrigin3DDict{
		ref: ref,
	}
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin3DDict) ArrayGPUIntegerCoordinate() js.Array[GPUIntegerCoordinate] {
	return js.Array[GPUIntegerCoordinate]{}.FromRef(x.ref)
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin3DDict) GPUOrigin3DDict() GPUOrigin3DDict {
	var ret GPUOrigin3DDict
	ret.UpdateFrom(x.ref)
	return ret
}

type GPUOrigin3D = OneOf_ArrayGPUIntegerCoordinate_GPUOrigin3DDict

type GPUImageCopyTexture struct {
	// Texture is "GPUImageCopyTexture.texture"
	//
	// Required
	Texture GPUTexture
	// MipLevel is "GPUImageCopyTexture.mipLevel"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MipLevel MUST be set to true to make this field effective.
	MipLevel GPUIntegerCoordinate
	// Origin is "GPUImageCopyTexture.origin"
	//
	// Optional, defaults to {}.
	Origin GPUOrigin3D
	// Aspect is "GPUImageCopyTexture.aspect"
	//
	// Optional, defaults to "all".
	Aspect GPUTextureAspect

	FFI_USE_MipLevel bool // for MipLevel.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUImageCopyTexture with all fields set.
func (p GPUImageCopyTexture) FromRef(ref js.Ref) GPUImageCopyTexture {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUImageCopyTexture in the application heap.
func (p GPUImageCopyTexture) New() js.Ref {
	return bindings.GPUImageCopyTextureJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUImageCopyTexture) UpdateFrom(ref js.Ref) {
	bindings.GPUImageCopyTextureJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUImageCopyTexture) Update(ref js.Ref) {
	bindings.GPUImageCopyTextureJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUCommandBuffer struct {
	ref js.Ref
}

func (this GPUCommandBuffer) Once() GPUCommandBuffer {
	this.Ref().Once()
	return this
}

func (this GPUCommandBuffer) Ref() js.Ref {
	return this.ref
}

func (this GPUCommandBuffer) FromRef(ref js.Ref) GPUCommandBuffer {
	this.ref = ref
	return this
}

func (this GPUCommandBuffer) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUCommandBuffer.label".
//
// It returns ok=false if there is no such property.
func (this GPUCommandBuffer) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUCommandBufferLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUCommandBuffer.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUCommandBuffer) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUCommandBufferLabel(
		this.Ref(),
		val.Ref(),
	)
}

type GPUCommandBufferDescriptor struct {
	// Label is "GPUCommandBufferDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUCommandBufferDescriptor with all fields set.
func (p GPUCommandBufferDescriptor) FromRef(ref js.Ref) GPUCommandBufferDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUCommandBufferDescriptor in the application heap.
func (p GPUCommandBufferDescriptor) New() js.Ref {
	return bindings.GPUCommandBufferDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUCommandBufferDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUCommandBufferDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUCommandBufferDescriptor) Update(ref js.Ref) {
	bindings.GPUCommandBufferDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUCommandEncoder struct {
	ref js.Ref
}

func (this GPUCommandEncoder) Once() GPUCommandEncoder {
	this.Ref().Once()
	return this
}

func (this GPUCommandEncoder) Ref() js.Ref {
	return this.ref
}

func (this GPUCommandEncoder) FromRef(ref js.Ref) GPUCommandEncoder {
	this.ref = ref
	return this
}

func (this GPUCommandEncoder) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUCommandEncoder.label".
//
// It returns ok=false if there is no such property.
func (this GPUCommandEncoder) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUCommandEncoderLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUCommandEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUCommandEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUCommandEncoderLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasBeginRenderPass returns true if the method "GPUCommandEncoder.beginRenderPass" exists.
func (this GPUCommandEncoder) HasBeginRenderPass() bool {
	return js.True == bindings.HasGPUCommandEncoderBeginRenderPass(
		this.Ref(),
	)
}

// BeginRenderPassFunc returns the method "GPUCommandEncoder.beginRenderPass".
func (this GPUCommandEncoder) BeginRenderPassFunc() (fn js.Func[func(descriptor GPURenderPassDescriptor) GPURenderPassEncoder]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderBeginRenderPassFunc(
			this.Ref(),
		),
	)
}

// BeginRenderPass calls the method "GPUCommandEncoder.beginRenderPass".
func (this GPUCommandEncoder) BeginRenderPass(descriptor GPURenderPassDescriptor) (ret GPURenderPassEncoder) {
	bindings.CallGPUCommandEncoderBeginRenderPass(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryBeginRenderPass calls the method "GPUCommandEncoder.beginRenderPass"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryBeginRenderPass(descriptor GPURenderPassDescriptor) (ret GPURenderPassEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderBeginRenderPass(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasBeginComputePass returns true if the method "GPUCommandEncoder.beginComputePass" exists.
func (this GPUCommandEncoder) HasBeginComputePass() bool {
	return js.True == bindings.HasGPUCommandEncoderBeginComputePass(
		this.Ref(),
	)
}

// BeginComputePassFunc returns the method "GPUCommandEncoder.beginComputePass".
func (this GPUCommandEncoder) BeginComputePassFunc() (fn js.Func[func(descriptor GPUComputePassDescriptor) GPUComputePassEncoder]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderBeginComputePassFunc(
			this.Ref(),
		),
	)
}

// BeginComputePass calls the method "GPUCommandEncoder.beginComputePass".
func (this GPUCommandEncoder) BeginComputePass(descriptor GPUComputePassDescriptor) (ret GPUComputePassEncoder) {
	bindings.CallGPUCommandEncoderBeginComputePass(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryBeginComputePass calls the method "GPUCommandEncoder.beginComputePass"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryBeginComputePass(descriptor GPUComputePassDescriptor) (ret GPUComputePassEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderBeginComputePass(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasBeginComputePass1 returns true if the method "GPUCommandEncoder.beginComputePass" exists.
func (this GPUCommandEncoder) HasBeginComputePass1() bool {
	return js.True == bindings.HasGPUCommandEncoderBeginComputePass1(
		this.Ref(),
	)
}

// BeginComputePass1Func returns the method "GPUCommandEncoder.beginComputePass".
func (this GPUCommandEncoder) BeginComputePass1Func() (fn js.Func[func() GPUComputePassEncoder]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderBeginComputePass1Func(
			this.Ref(),
		),
	)
}

// BeginComputePass1 calls the method "GPUCommandEncoder.beginComputePass".
func (this GPUCommandEncoder) BeginComputePass1() (ret GPUComputePassEncoder) {
	bindings.CallGPUCommandEncoderBeginComputePass1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBeginComputePass1 calls the method "GPUCommandEncoder.beginComputePass"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryBeginComputePass1() (ret GPUComputePassEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderBeginComputePass1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCopyBufferToBuffer returns true if the method "GPUCommandEncoder.copyBufferToBuffer" exists.
func (this GPUCommandEncoder) HasCopyBufferToBuffer() bool {
	return js.True == bindings.HasGPUCommandEncoderCopyBufferToBuffer(
		this.Ref(),
	)
}

// CopyBufferToBufferFunc returns the method "GPUCommandEncoder.copyBufferToBuffer".
func (this GPUCommandEncoder) CopyBufferToBufferFunc() (fn js.Func[func(source GPUBuffer, sourceOffset GPUSize64, destination GPUBuffer, destinationOffset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderCopyBufferToBufferFunc(
			this.Ref(),
		),
	)
}

// CopyBufferToBuffer calls the method "GPUCommandEncoder.copyBufferToBuffer".
func (this GPUCommandEncoder) CopyBufferToBuffer(source GPUBuffer, sourceOffset GPUSize64, destination GPUBuffer, destinationOffset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPUCommandEncoderCopyBufferToBuffer(
		this.Ref(), js.Pointer(&ret),
		source.Ref(),
		float64(sourceOffset),
		destination.Ref(),
		float64(destinationOffset),
		float64(size),
	)

	return
}

// TryCopyBufferToBuffer calls the method "GPUCommandEncoder.copyBufferToBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryCopyBufferToBuffer(source GPUBuffer, sourceOffset GPUSize64, destination GPUBuffer, destinationOffset GPUSize64, size GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderCopyBufferToBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
		float64(sourceOffset),
		destination.Ref(),
		float64(destinationOffset),
		float64(size),
	)

	return
}

// HasCopyBufferToTexture returns true if the method "GPUCommandEncoder.copyBufferToTexture" exists.
func (this GPUCommandEncoder) HasCopyBufferToTexture() bool {
	return js.True == bindings.HasGPUCommandEncoderCopyBufferToTexture(
		this.Ref(),
	)
}

// CopyBufferToTextureFunc returns the method "GPUCommandEncoder.copyBufferToTexture".
func (this GPUCommandEncoder) CopyBufferToTextureFunc() (fn js.Func[func(source GPUImageCopyBuffer, destination GPUImageCopyTexture, copySize GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderCopyBufferToTextureFunc(
			this.Ref(),
		),
	)
}

// CopyBufferToTexture calls the method "GPUCommandEncoder.copyBufferToTexture".
func (this GPUCommandEncoder) CopyBufferToTexture(source GPUImageCopyBuffer, destination GPUImageCopyTexture, copySize GPUExtent3D) (ret js.Void) {
	bindings.CallGPUCommandEncoderCopyBufferToTexture(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// TryCopyBufferToTexture calls the method "GPUCommandEncoder.copyBufferToTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryCopyBufferToTexture(source GPUImageCopyBuffer, destination GPUImageCopyTexture, copySize GPUExtent3D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderCopyBufferToTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// HasCopyTextureToBuffer returns true if the method "GPUCommandEncoder.copyTextureToBuffer" exists.
func (this GPUCommandEncoder) HasCopyTextureToBuffer() bool {
	return js.True == bindings.HasGPUCommandEncoderCopyTextureToBuffer(
		this.Ref(),
	)
}

// CopyTextureToBufferFunc returns the method "GPUCommandEncoder.copyTextureToBuffer".
func (this GPUCommandEncoder) CopyTextureToBufferFunc() (fn js.Func[func(source GPUImageCopyTexture, destination GPUImageCopyBuffer, copySize GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderCopyTextureToBufferFunc(
			this.Ref(),
		),
	)
}

// CopyTextureToBuffer calls the method "GPUCommandEncoder.copyTextureToBuffer".
func (this GPUCommandEncoder) CopyTextureToBuffer(source GPUImageCopyTexture, destination GPUImageCopyBuffer, copySize GPUExtent3D) (ret js.Void) {
	bindings.CallGPUCommandEncoderCopyTextureToBuffer(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// TryCopyTextureToBuffer calls the method "GPUCommandEncoder.copyTextureToBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryCopyTextureToBuffer(source GPUImageCopyTexture, destination GPUImageCopyBuffer, copySize GPUExtent3D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderCopyTextureToBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// HasCopyTextureToTexture returns true if the method "GPUCommandEncoder.copyTextureToTexture" exists.
func (this GPUCommandEncoder) HasCopyTextureToTexture() bool {
	return js.True == bindings.HasGPUCommandEncoderCopyTextureToTexture(
		this.Ref(),
	)
}

// CopyTextureToTextureFunc returns the method "GPUCommandEncoder.copyTextureToTexture".
func (this GPUCommandEncoder) CopyTextureToTextureFunc() (fn js.Func[func(source GPUImageCopyTexture, destination GPUImageCopyTexture, copySize GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderCopyTextureToTextureFunc(
			this.Ref(),
		),
	)
}

// CopyTextureToTexture calls the method "GPUCommandEncoder.copyTextureToTexture".
func (this GPUCommandEncoder) CopyTextureToTexture(source GPUImageCopyTexture, destination GPUImageCopyTexture, copySize GPUExtent3D) (ret js.Void) {
	bindings.CallGPUCommandEncoderCopyTextureToTexture(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// TryCopyTextureToTexture calls the method "GPUCommandEncoder.copyTextureToTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryCopyTextureToTexture(source GPUImageCopyTexture, destination GPUImageCopyTexture, copySize GPUExtent3D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderCopyTextureToTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// HasClearBuffer returns true if the method "GPUCommandEncoder.clearBuffer" exists.
func (this GPUCommandEncoder) HasClearBuffer() bool {
	return js.True == bindings.HasGPUCommandEncoderClearBuffer(
		this.Ref(),
	)
}

// ClearBufferFunc returns the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) ClearBufferFunc() (fn js.Func[func(buffer GPUBuffer, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderClearBufferFunc(
			this.Ref(),
		),
	)
}

// ClearBuffer calls the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) ClearBuffer(buffer GPUBuffer, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPUCommandEncoderClearBuffer(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// TryClearBuffer calls the method "GPUCommandEncoder.clearBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryClearBuffer(buffer GPUBuffer, offset GPUSize64, size GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderClearBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// HasClearBuffer1 returns true if the method "GPUCommandEncoder.clearBuffer" exists.
func (this GPUCommandEncoder) HasClearBuffer1() bool {
	return js.True == bindings.HasGPUCommandEncoderClearBuffer1(
		this.Ref(),
	)
}

// ClearBuffer1Func returns the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) ClearBuffer1Func() (fn js.Func[func(buffer GPUBuffer, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderClearBuffer1Func(
			this.Ref(),
		),
	)
}

// ClearBuffer1 calls the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) ClearBuffer1(buffer GPUBuffer, offset GPUSize64) (ret js.Void) {
	bindings.CallGPUCommandEncoderClearBuffer1(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		float64(offset),
	)

	return
}

// TryClearBuffer1 calls the method "GPUCommandEncoder.clearBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryClearBuffer1(buffer GPUBuffer, offset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderClearBuffer1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(offset),
	)

	return
}

// HasClearBuffer2 returns true if the method "GPUCommandEncoder.clearBuffer" exists.
func (this GPUCommandEncoder) HasClearBuffer2() bool {
	return js.True == bindings.HasGPUCommandEncoderClearBuffer2(
		this.Ref(),
	)
}

// ClearBuffer2Func returns the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) ClearBuffer2Func() (fn js.Func[func(buffer GPUBuffer)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderClearBuffer2Func(
			this.Ref(),
		),
	)
}

// ClearBuffer2 calls the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) ClearBuffer2(buffer GPUBuffer) (ret js.Void) {
	bindings.CallGPUCommandEncoderClearBuffer2(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryClearBuffer2 calls the method "GPUCommandEncoder.clearBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryClearBuffer2(buffer GPUBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderClearBuffer2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasWriteTimestamp returns true if the method "GPUCommandEncoder.writeTimestamp" exists.
func (this GPUCommandEncoder) HasWriteTimestamp() bool {
	return js.True == bindings.HasGPUCommandEncoderWriteTimestamp(
		this.Ref(),
	)
}

// WriteTimestampFunc returns the method "GPUCommandEncoder.writeTimestamp".
func (this GPUCommandEncoder) WriteTimestampFunc() (fn js.Func[func(querySet GPUQuerySet, queryIndex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderWriteTimestampFunc(
			this.Ref(),
		),
	)
}

// WriteTimestamp calls the method "GPUCommandEncoder.writeTimestamp".
func (this GPUCommandEncoder) WriteTimestamp(querySet GPUQuerySet, queryIndex GPUSize32) (ret js.Void) {
	bindings.CallGPUCommandEncoderWriteTimestamp(
		this.Ref(), js.Pointer(&ret),
		querySet.Ref(),
		uint32(queryIndex),
	)

	return
}

// TryWriteTimestamp calls the method "GPUCommandEncoder.writeTimestamp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryWriteTimestamp(querySet GPUQuerySet, queryIndex GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderWriteTimestamp(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		querySet.Ref(),
		uint32(queryIndex),
	)

	return
}

// HasResolveQuerySet returns true if the method "GPUCommandEncoder.resolveQuerySet" exists.
func (this GPUCommandEncoder) HasResolveQuerySet() bool {
	return js.True == bindings.HasGPUCommandEncoderResolveQuerySet(
		this.Ref(),
	)
}

// ResolveQuerySetFunc returns the method "GPUCommandEncoder.resolveQuerySet".
func (this GPUCommandEncoder) ResolveQuerySetFunc() (fn js.Func[func(querySet GPUQuerySet, firstQuery GPUSize32, queryCount GPUSize32, destination GPUBuffer, destinationOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderResolveQuerySetFunc(
			this.Ref(),
		),
	)
}

// ResolveQuerySet calls the method "GPUCommandEncoder.resolveQuerySet".
func (this GPUCommandEncoder) ResolveQuerySet(querySet GPUQuerySet, firstQuery GPUSize32, queryCount GPUSize32, destination GPUBuffer, destinationOffset GPUSize64) (ret js.Void) {
	bindings.CallGPUCommandEncoderResolveQuerySet(
		this.Ref(), js.Pointer(&ret),
		querySet.Ref(),
		uint32(firstQuery),
		uint32(queryCount),
		destination.Ref(),
		float64(destinationOffset),
	)

	return
}

// TryResolveQuerySet calls the method "GPUCommandEncoder.resolveQuerySet"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryResolveQuerySet(querySet GPUQuerySet, firstQuery GPUSize32, queryCount GPUSize32, destination GPUBuffer, destinationOffset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderResolveQuerySet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		querySet.Ref(),
		uint32(firstQuery),
		uint32(queryCount),
		destination.Ref(),
		float64(destinationOffset),
	)

	return
}

// HasFinish returns true if the method "GPUCommandEncoder.finish" exists.
func (this GPUCommandEncoder) HasFinish() bool {
	return js.True == bindings.HasGPUCommandEncoderFinish(
		this.Ref(),
	)
}

// FinishFunc returns the method "GPUCommandEncoder.finish".
func (this GPUCommandEncoder) FinishFunc() (fn js.Func[func(descriptor GPUCommandBufferDescriptor) GPUCommandBuffer]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderFinishFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "GPUCommandEncoder.finish".
func (this GPUCommandEncoder) Finish(descriptor GPUCommandBufferDescriptor) (ret GPUCommandBuffer) {
	bindings.CallGPUCommandEncoderFinish(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryFinish calls the method "GPUCommandEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryFinish(descriptor GPUCommandBufferDescriptor) (ret GPUCommandBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderFinish(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFinish1 returns true if the method "GPUCommandEncoder.finish" exists.
func (this GPUCommandEncoder) HasFinish1() bool {
	return js.True == bindings.HasGPUCommandEncoderFinish1(
		this.Ref(),
	)
}

// Finish1Func returns the method "GPUCommandEncoder.finish".
func (this GPUCommandEncoder) Finish1Func() (fn js.Func[func() GPUCommandBuffer]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderFinish1Func(
			this.Ref(),
		),
	)
}

// Finish1 calls the method "GPUCommandEncoder.finish".
func (this GPUCommandEncoder) Finish1() (ret GPUCommandBuffer) {
	bindings.CallGPUCommandEncoderFinish1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFinish1 calls the method "GPUCommandEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryFinish1() (ret GPUCommandBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderFinish1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPushDebugGroup returns true if the method "GPUCommandEncoder.pushDebugGroup" exists.
func (this GPUCommandEncoder) HasPushDebugGroup() bool {
	return js.True == bindings.HasGPUCommandEncoderPushDebugGroup(
		this.Ref(),
	)
}

// PushDebugGroupFunc returns the method "GPUCommandEncoder.pushDebugGroup".
func (this GPUCommandEncoder) PushDebugGroupFunc() (fn js.Func[func(groupLabel js.String)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderPushDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PushDebugGroup calls the method "GPUCommandEncoder.pushDebugGroup".
func (this GPUCommandEncoder) PushDebugGroup(groupLabel js.String) (ret js.Void) {
	bindings.CallGPUCommandEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&ret),
		groupLabel.Ref(),
	)

	return
}

// TryPushDebugGroup calls the method "GPUCommandEncoder.pushDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryPushDebugGroup(groupLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		groupLabel.Ref(),
	)

	return
}

// HasPopDebugGroup returns true if the method "GPUCommandEncoder.popDebugGroup" exists.
func (this GPUCommandEncoder) HasPopDebugGroup() bool {
	return js.True == bindings.HasGPUCommandEncoderPopDebugGroup(
		this.Ref(),
	)
}

// PopDebugGroupFunc returns the method "GPUCommandEncoder.popDebugGroup".
func (this GPUCommandEncoder) PopDebugGroupFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderPopDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PopDebugGroup calls the method "GPUCommandEncoder.popDebugGroup".
func (this GPUCommandEncoder) PopDebugGroup() (ret js.Void) {
	bindings.CallGPUCommandEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPopDebugGroup calls the method "GPUCommandEncoder.popDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryPopDebugGroup() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInsertDebugMarker returns true if the method "GPUCommandEncoder.insertDebugMarker" exists.
func (this GPUCommandEncoder) HasInsertDebugMarker() bool {
	return js.True == bindings.HasGPUCommandEncoderInsertDebugMarker(
		this.Ref(),
	)
}

// InsertDebugMarkerFunc returns the method "GPUCommandEncoder.insertDebugMarker".
func (this GPUCommandEncoder) InsertDebugMarkerFunc() (fn js.Func[func(markerLabel js.String)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderInsertDebugMarkerFunc(
			this.Ref(),
		),
	)
}

// InsertDebugMarker calls the method "GPUCommandEncoder.insertDebugMarker".
func (this GPUCommandEncoder) InsertDebugMarker(markerLabel js.String) (ret js.Void) {
	bindings.CallGPUCommandEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&ret),
		markerLabel.Ref(),
	)

	return
}

// TryInsertDebugMarker calls the method "GPUCommandEncoder.insertDebugMarker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryInsertDebugMarker(markerLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		markerLabel.Ref(),
	)

	return
}

type GPUCommandEncoderDescriptor struct {
	// Label is "GPUCommandEncoderDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUCommandEncoderDescriptor with all fields set.
func (p GPUCommandEncoderDescriptor) FromRef(ref js.Ref) GPUCommandEncoderDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUCommandEncoderDescriptor in the application heap.
func (p GPUCommandEncoderDescriptor) New() js.Ref {
	return bindings.GPUCommandEncoderDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUCommandEncoderDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUCommandEncoderDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUCommandEncoderDescriptor) Update(ref js.Ref) {
	bindings.GPUCommandEncoderDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPURenderBundleDescriptor struct {
	// Label is "GPURenderBundleDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURenderBundleDescriptor with all fields set.
func (p GPURenderBundleDescriptor) FromRef(ref js.Ref) GPURenderBundleDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPURenderBundleDescriptor in the application heap.
func (p GPURenderBundleDescriptor) New() js.Ref {
	return bindings.GPURenderBundleDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPURenderBundleDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPURenderBundleDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPURenderBundleDescriptor) Update(ref js.Ref) {
	bindings.GPURenderBundleDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPURenderBundleEncoder struct {
	ref js.Ref
}

func (this GPURenderBundleEncoder) Once() GPURenderBundleEncoder {
	this.Ref().Once()
	return this
}

func (this GPURenderBundleEncoder) Ref() js.Ref {
	return this.ref
}

func (this GPURenderBundleEncoder) FromRef(ref js.Ref) GPURenderBundleEncoder {
	this.ref = ref
	return this
}

func (this GPURenderBundleEncoder) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPURenderBundleEncoder.label".
//
// It returns ok=false if there is no such property.
func (this GPURenderBundleEncoder) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPURenderBundleEncoderLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPURenderBundleEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPURenderBundleEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPURenderBundleEncoderLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasFinish returns true if the method "GPURenderBundleEncoder.finish" exists.
func (this GPURenderBundleEncoder) HasFinish() bool {
	return js.True == bindings.HasGPURenderBundleEncoderFinish(
		this.Ref(),
	)
}

// FinishFunc returns the method "GPURenderBundleEncoder.finish".
func (this GPURenderBundleEncoder) FinishFunc() (fn js.Func[func(descriptor GPURenderBundleDescriptor) GPURenderBundle]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderFinishFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "GPURenderBundleEncoder.finish".
func (this GPURenderBundleEncoder) Finish(descriptor GPURenderBundleDescriptor) (ret GPURenderBundle) {
	bindings.CallGPURenderBundleEncoderFinish(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryFinish calls the method "GPURenderBundleEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryFinish(descriptor GPURenderBundleDescriptor) (ret GPURenderBundle, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderFinish(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFinish1 returns true if the method "GPURenderBundleEncoder.finish" exists.
func (this GPURenderBundleEncoder) HasFinish1() bool {
	return js.True == bindings.HasGPURenderBundleEncoderFinish1(
		this.Ref(),
	)
}

// Finish1Func returns the method "GPURenderBundleEncoder.finish".
func (this GPURenderBundleEncoder) Finish1Func() (fn js.Func[func() GPURenderBundle]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderFinish1Func(
			this.Ref(),
		),
	)
}

// Finish1 calls the method "GPURenderBundleEncoder.finish".
func (this GPURenderBundleEncoder) Finish1() (ret GPURenderBundle) {
	bindings.CallGPURenderBundleEncoderFinish1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFinish1 calls the method "GPURenderBundleEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryFinish1() (ret GPURenderBundle, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderFinish1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetPipeline returns true if the method "GPURenderBundleEncoder.setPipeline" exists.
func (this GPURenderBundleEncoder) HasSetPipeline() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetPipeline(
		this.Ref(),
	)
}

// SetPipelineFunc returns the method "GPURenderBundleEncoder.setPipeline".
func (this GPURenderBundleEncoder) SetPipelineFunc() (fn js.Func[func(pipeline GPURenderPipeline)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetPipelineFunc(
			this.Ref(),
		),
	)
}

// SetPipeline calls the method "GPURenderBundleEncoder.setPipeline".
func (this GPURenderBundleEncoder) SetPipeline(pipeline GPURenderPipeline) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetPipeline(
		this.Ref(), js.Pointer(&ret),
		pipeline.Ref(),
	)

	return
}

// TrySetPipeline calls the method "GPURenderBundleEncoder.setPipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetPipeline(pipeline GPURenderPipeline) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetPipeline(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		pipeline.Ref(),
	)

	return
}

// HasSetIndexBuffer returns true if the method "GPURenderBundleEncoder.setIndexBuffer" exists.
func (this GPURenderBundleEncoder) HasSetIndexBuffer() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetIndexBuffer(
		this.Ref(),
	)
}

// SetIndexBufferFunc returns the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) SetIndexBufferFunc() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetIndexBufferFunc(
			this.Ref(),
		),
	)
}

// SetIndexBuffer calls the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) SetIndexBuffer(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetIndexBuffer(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
		float64(size),
	)

	return
}

// TrySetIndexBuffer calls the method "GPURenderBundleEncoder.setIndexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetIndexBuffer(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetIndexBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
		float64(size),
	)

	return
}

// HasSetIndexBuffer1 returns true if the method "GPURenderBundleEncoder.setIndexBuffer" exists.
func (this GPURenderBundleEncoder) HasSetIndexBuffer1() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetIndexBuffer1(
		this.Ref(),
	)
}

// SetIndexBuffer1Func returns the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) SetIndexBuffer1Func() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetIndexBuffer1Func(
			this.Ref(),
		),
	)
}

// SetIndexBuffer1 calls the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) SetIndexBuffer1(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetIndexBuffer1(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
	)

	return
}

// TrySetIndexBuffer1 calls the method "GPURenderBundleEncoder.setIndexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetIndexBuffer1(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetIndexBuffer1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
	)

	return
}

// HasSetIndexBuffer2 returns true if the method "GPURenderBundleEncoder.setIndexBuffer" exists.
func (this GPURenderBundleEncoder) HasSetIndexBuffer2() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetIndexBuffer2(
		this.Ref(),
	)
}

// SetIndexBuffer2Func returns the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) SetIndexBuffer2Func() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetIndexBuffer2Func(
			this.Ref(),
		),
	)
}

// SetIndexBuffer2 calls the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) SetIndexBuffer2(buffer GPUBuffer, indexFormat GPUIndexFormat) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetIndexBuffer2(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		uint32(indexFormat),
	)

	return
}

// TrySetIndexBuffer2 calls the method "GPURenderBundleEncoder.setIndexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetIndexBuffer2(buffer GPUBuffer, indexFormat GPUIndexFormat) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetIndexBuffer2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
	)

	return
}

// HasSetVertexBuffer returns true if the method "GPURenderBundleEncoder.setVertexBuffer" exists.
func (this GPURenderBundleEncoder) HasSetVertexBuffer() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetVertexBuffer(
		this.Ref(),
	)
}

// SetVertexBufferFunc returns the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) SetVertexBufferFunc() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetVertexBufferFunc(
			this.Ref(),
		),
	)
}

// SetVertexBuffer calls the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) SetVertexBuffer(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetVertexBuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// TrySetVertexBuffer calls the method "GPURenderBundleEncoder.setVertexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetVertexBuffer(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetVertexBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// HasSetVertexBuffer1 returns true if the method "GPURenderBundleEncoder.setVertexBuffer" exists.
func (this GPURenderBundleEncoder) HasSetVertexBuffer1() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetVertexBuffer1(
		this.Ref(),
	)
}

// SetVertexBuffer1Func returns the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) SetVertexBuffer1Func() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetVertexBuffer1Func(
			this.Ref(),
		),
	)
}

// SetVertexBuffer1 calls the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) SetVertexBuffer1(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetVertexBuffer1(
		this.Ref(), js.Pointer(&ret),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
	)

	return
}

// TrySetVertexBuffer1 calls the method "GPURenderBundleEncoder.setVertexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetVertexBuffer1(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetVertexBuffer1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
	)

	return
}

// HasSetVertexBuffer2 returns true if the method "GPURenderBundleEncoder.setVertexBuffer" exists.
func (this GPURenderBundleEncoder) HasSetVertexBuffer2() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetVertexBuffer2(
		this.Ref(),
	)
}

// SetVertexBuffer2Func returns the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) SetVertexBuffer2Func() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetVertexBuffer2Func(
			this.Ref(),
		),
	)
}

// SetVertexBuffer2 calls the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) SetVertexBuffer2(slot GPUIndex32, buffer GPUBuffer) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetVertexBuffer2(
		this.Ref(), js.Pointer(&ret),
		uint32(slot),
		buffer.Ref(),
	)

	return
}

// TrySetVertexBuffer2 calls the method "GPURenderBundleEncoder.setVertexBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetVertexBuffer2(slot GPUIndex32, buffer GPUBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetVertexBuffer2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
	)

	return
}

// HasDraw returns true if the method "GPURenderBundleEncoder.draw" exists.
func (this GPURenderBundleEncoder) HasDraw() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDraw(
		this.Ref(),
	)
}

// DrawFunc returns the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) DrawFunc() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawFunc(
			this.Ref(),
		),
	)
}

// Draw calls the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDraw(
		this.Ref(), js.Pointer(&ret),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
		uint32(firstInstance),
	)

	return
}

// TryDraw calls the method "GPURenderBundleEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDraw(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDraw(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
		uint32(firstInstance),
	)

	return
}

// HasDraw1 returns true if the method "GPURenderBundleEncoder.draw" exists.
func (this GPURenderBundleEncoder) HasDraw1() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDraw1(
		this.Ref(),
	)
}

// Draw1Func returns the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw1Func() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDraw1Func(
			this.Ref(),
		),
	)
}

// Draw1 calls the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw1(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDraw1(
		this.Ref(), js.Pointer(&ret),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
	)

	return
}

// TryDraw1 calls the method "GPURenderBundleEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDraw1(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDraw1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
	)

	return
}

// HasDraw2 returns true if the method "GPURenderBundleEncoder.draw" exists.
func (this GPURenderBundleEncoder) HasDraw2() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDraw2(
		this.Ref(),
	)
}

// Draw2Func returns the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw2Func() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDraw2Func(
			this.Ref(),
		),
	)
}

// Draw2 calls the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw2(vertexCount GPUSize32, instanceCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDraw2(
		this.Ref(), js.Pointer(&ret),
		uint32(vertexCount),
		uint32(instanceCount),
	)

	return
}

// TryDraw2 calls the method "GPURenderBundleEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDraw2(vertexCount GPUSize32, instanceCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDraw2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
	)

	return
}

// HasDraw3 returns true if the method "GPURenderBundleEncoder.draw" exists.
func (this GPURenderBundleEncoder) HasDraw3() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDraw3(
		this.Ref(),
	)
}

// Draw3Func returns the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw3Func() (fn js.Func[func(vertexCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDraw3Func(
			this.Ref(),
		),
	)
}

// Draw3 calls the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw3(vertexCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDraw3(
		this.Ref(), js.Pointer(&ret),
		uint32(vertexCount),
	)

	return
}

// TryDraw3 calls the method "GPURenderBundleEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDraw3(vertexCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDraw3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
	)

	return
}

// HasDrawIndexed returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasDrawIndexed() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDrawIndexed(
		this.Ref(),
	)
}

// DrawIndexedFunc returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexedFunc() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexedFunc(
			this.Ref(),
		),
	)
}

// DrawIndexed calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
		uint32(firstInstance),
	)

	return
}

// TryDrawIndexed calls the method "GPURenderBundleEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDrawIndexed(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDrawIndexed(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
		uint32(firstInstance),
	)

	return
}

// HasDrawIndexed1 returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasDrawIndexed1() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDrawIndexed1(
		this.Ref(),
	)
}

// DrawIndexed1Func returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed1Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexed1Func(
			this.Ref(),
		),
	)
}

// DrawIndexed1 calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed1(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed1(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
	)

	return
}

// TryDrawIndexed1 calls the method "GPURenderBundleEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDrawIndexed1(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDrawIndexed1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
	)

	return
}

// HasDrawIndexed2 returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasDrawIndexed2() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDrawIndexed2(
		this.Ref(),
	)
}

// DrawIndexed2Func returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed2Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexed2Func(
			this.Ref(),
		),
	)
}

// DrawIndexed2 calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed2(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed2(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
	)

	return
}

// TryDrawIndexed2 calls the method "GPURenderBundleEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDrawIndexed2(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDrawIndexed2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
	)

	return
}

// HasDrawIndexed3 returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasDrawIndexed3() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDrawIndexed3(
		this.Ref(),
	)
}

// DrawIndexed3Func returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed3Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexed3Func(
			this.Ref(),
		),
	)
}

// DrawIndexed3 calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed3(indexCount GPUSize32, instanceCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed3(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
		uint32(instanceCount),
	)

	return
}

// TryDrawIndexed3 calls the method "GPURenderBundleEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDrawIndexed3(indexCount GPUSize32, instanceCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDrawIndexed3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
	)

	return
}

// HasDrawIndexed4 returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasDrawIndexed4() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDrawIndexed4(
		this.Ref(),
	)
}

// DrawIndexed4Func returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed4Func() (fn js.Func[func(indexCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexed4Func(
			this.Ref(),
		),
	)
}

// DrawIndexed4 calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed4(indexCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed4(
		this.Ref(), js.Pointer(&ret),
		uint32(indexCount),
	)

	return
}

// TryDrawIndexed4 calls the method "GPURenderBundleEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDrawIndexed4(indexCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDrawIndexed4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
	)

	return
}

// HasDrawIndirect returns true if the method "GPURenderBundleEncoder.drawIndirect" exists.
func (this GPURenderBundleEncoder) HasDrawIndirect() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDrawIndirect(
		this.Ref(),
	)
}

// DrawIndirectFunc returns the method "GPURenderBundleEncoder.drawIndirect".
func (this GPURenderBundleEncoder) DrawIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndirectFunc(
			this.Ref(),
		),
	)
}

// DrawIndirect calls the method "GPURenderBundleEncoder.drawIndirect".
func (this GPURenderBundleEncoder) DrawIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndirect(
		this.Ref(), js.Pointer(&ret),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// TryDrawIndirect calls the method "GPURenderBundleEncoder.drawIndirect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDrawIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDrawIndirect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasDrawIndexedIndirect returns true if the method "GPURenderBundleEncoder.drawIndexedIndirect" exists.
func (this GPURenderBundleEncoder) HasDrawIndexedIndirect() bool {
	return js.True == bindings.HasGPURenderBundleEncoderDrawIndexedIndirect(
		this.Ref(),
	)
}

// DrawIndexedIndirectFunc returns the method "GPURenderBundleEncoder.drawIndexedIndirect".
func (this GPURenderBundleEncoder) DrawIndexedIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexedIndirectFunc(
			this.Ref(),
		),
	)
}

// DrawIndexedIndirect calls the method "GPURenderBundleEncoder.drawIndexedIndirect".
func (this GPURenderBundleEncoder) DrawIndexedIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexedIndirect(
		this.Ref(), js.Pointer(&ret),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// TryDrawIndexedIndirect calls the method "GPURenderBundleEncoder.drawIndexedIndirect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDrawIndexedIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDrawIndexedIndirect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasSetBindGroup returns true if the method "GPURenderBundleEncoder.setBindGroup" exists.
func (this GPURenderBundleEncoder) HasSetBindGroup() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetBindGroup(
		this.Ref(),
	)
}

// SetBindGroupFunc returns the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) SetBindGroupFunc() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset])]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetBindGroupFunc(
			this.Ref(),
		),
	)
}

// SetBindGroup calls the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetBindGroup(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	return
}

// TrySetBindGroup calls the method "GPURenderBundleEncoder.setBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetBindGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	return
}

// HasSetBindGroup1 returns true if the method "GPURenderBundleEncoder.setBindGroup" exists.
func (this GPURenderBundleEncoder) HasSetBindGroup1() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetBindGroup1(
		this.Ref(),
	)
}

// SetBindGroup1Func returns the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) SetBindGroup1Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetBindGroup1Func(
			this.Ref(),
		),
	)
}

// SetBindGroup1 calls the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) SetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetBindGroup1(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		bindGroup.Ref(),
	)

	return
}

// TrySetBindGroup1 calls the method "GPURenderBundleEncoder.setBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetBindGroup1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
	)

	return
}

// HasSetBindGroup2 returns true if the method "GPURenderBundleEncoder.setBindGroup" exists.
func (this GPURenderBundleEncoder) HasSetBindGroup2() bool {
	return js.True == bindings.HasGPURenderBundleEncoderSetBindGroup2(
		this.Ref(),
	)
}

// SetBindGroup2Func returns the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) SetBindGroup2Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetBindGroup2Func(
			this.Ref(),
		),
	)
}

// SetBindGroup2 calls the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) SetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetBindGroup2(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	return
}

// TrySetBindGroup2 calls the method "GPURenderBundleEncoder.setBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetBindGroup2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	return
}

// HasPushDebugGroup returns true if the method "GPURenderBundleEncoder.pushDebugGroup" exists.
func (this GPURenderBundleEncoder) HasPushDebugGroup() bool {
	return js.True == bindings.HasGPURenderBundleEncoderPushDebugGroup(
		this.Ref(),
	)
}

// PushDebugGroupFunc returns the method "GPURenderBundleEncoder.pushDebugGroup".
func (this GPURenderBundleEncoder) PushDebugGroupFunc() (fn js.Func[func(groupLabel js.String)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderPushDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PushDebugGroup calls the method "GPURenderBundleEncoder.pushDebugGroup".
func (this GPURenderBundleEncoder) PushDebugGroup(groupLabel js.String) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&ret),
		groupLabel.Ref(),
	)

	return
}

// TryPushDebugGroup calls the method "GPURenderBundleEncoder.pushDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryPushDebugGroup(groupLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		groupLabel.Ref(),
	)

	return
}

// HasPopDebugGroup returns true if the method "GPURenderBundleEncoder.popDebugGroup" exists.
func (this GPURenderBundleEncoder) HasPopDebugGroup() bool {
	return js.True == bindings.HasGPURenderBundleEncoderPopDebugGroup(
		this.Ref(),
	)
}

// PopDebugGroupFunc returns the method "GPURenderBundleEncoder.popDebugGroup".
func (this GPURenderBundleEncoder) PopDebugGroupFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderPopDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PopDebugGroup calls the method "GPURenderBundleEncoder.popDebugGroup".
func (this GPURenderBundleEncoder) PopDebugGroup() (ret js.Void) {
	bindings.CallGPURenderBundleEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPopDebugGroup calls the method "GPURenderBundleEncoder.popDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryPopDebugGroup() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInsertDebugMarker returns true if the method "GPURenderBundleEncoder.insertDebugMarker" exists.
func (this GPURenderBundleEncoder) HasInsertDebugMarker() bool {
	return js.True == bindings.HasGPURenderBundleEncoderInsertDebugMarker(
		this.Ref(),
	)
}

// InsertDebugMarkerFunc returns the method "GPURenderBundleEncoder.insertDebugMarker".
func (this GPURenderBundleEncoder) InsertDebugMarkerFunc() (fn js.Func[func(markerLabel js.String)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderInsertDebugMarkerFunc(
			this.Ref(),
		),
	)
}

// InsertDebugMarker calls the method "GPURenderBundleEncoder.insertDebugMarker".
func (this GPURenderBundleEncoder) InsertDebugMarker(markerLabel js.String) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&ret),
		markerLabel.Ref(),
	)

	return
}

// TryInsertDebugMarker calls the method "GPURenderBundleEncoder.insertDebugMarker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryInsertDebugMarker(markerLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		markerLabel.Ref(),
	)

	return
}

type GPURenderBundleEncoderDescriptor struct {
	// DepthReadOnly is "GPURenderBundleEncoderDescriptor.depthReadOnly"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_DepthReadOnly MUST be set to true to make this field effective.
	DepthReadOnly bool
	// StencilReadOnly is "GPURenderBundleEncoderDescriptor.stencilReadOnly"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_StencilReadOnly MUST be set to true to make this field effective.
	StencilReadOnly bool
	// ColorFormats is "GPURenderBundleEncoderDescriptor.colorFormats"
	//
	// Required
	ColorFormats js.Array[GPUTextureFormat]
	// DepthStencilFormat is "GPURenderBundleEncoderDescriptor.depthStencilFormat"
	//
	// Optional
	DepthStencilFormat GPUTextureFormat
	// SampleCount is "GPURenderBundleEncoderDescriptor.sampleCount"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_SampleCount MUST be set to true to make this field effective.
	SampleCount GPUSize32
	// Label is "GPURenderBundleEncoderDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE_DepthReadOnly   bool // for DepthReadOnly.
	FFI_USE_StencilReadOnly bool // for StencilReadOnly.
	FFI_USE_SampleCount     bool // for SampleCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURenderBundleEncoderDescriptor with all fields set.
func (p GPURenderBundleEncoderDescriptor) FromRef(ref js.Ref) GPURenderBundleEncoderDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPURenderBundleEncoderDescriptor in the application heap.
func (p GPURenderBundleEncoderDescriptor) New() js.Ref {
	return bindings.GPURenderBundleEncoderDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPURenderBundleEncoderDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPURenderBundleEncoderDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPURenderBundleEncoderDescriptor) Update(ref js.Ref) {
	bindings.GPURenderBundleEncoderDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUQuerySetDescriptor struct {
	// Type is "GPUQuerySetDescriptor.type"
	//
	// Required
	Type GPUQueryType
	// Count is "GPUQuerySetDescriptor.count"
	//
	// Required
	Count GPUSize32
	// Label is "GPUQuerySetDescriptor.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUQuerySetDescriptor with all fields set.
func (p GPUQuerySetDescriptor) FromRef(ref js.Ref) GPUQuerySetDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUQuerySetDescriptor in the application heap.
func (p GPUQuerySetDescriptor) New() js.Ref {
	return bindings.GPUQuerySetDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUQuerySetDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUQuerySetDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUQuerySetDescriptor) Update(ref js.Ref) {
	bindings.GPUQuerySetDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUErrorFilter uint32

const (
	_ GPUErrorFilter = iota

	GPUErrorFilter_VALIDATION
	GPUErrorFilter_OUT_OF_MEMORY
	GPUErrorFilter_INTERNAL
)

func (GPUErrorFilter) FromRef(str js.Ref) GPUErrorFilter {
	return GPUErrorFilter(bindings.ConstOfGPUErrorFilter(str))
}

func (x GPUErrorFilter) String() (string, bool) {
	switch x {
	case GPUErrorFilter_VALIDATION:
		return "validation", true
	case GPUErrorFilter_OUT_OF_MEMORY:
		return "out-of-memory", true
	case GPUErrorFilter_INTERNAL:
		return "internal", true
	default:
		return "", false
	}
}

type GPUError struct {
	ref js.Ref
}

func (this GPUError) Once() GPUError {
	this.Ref().Once()
	return this
}

func (this GPUError) Ref() js.Ref {
	return this.ref
}

func (this GPUError) FromRef(ref js.Ref) GPUError {
	this.ref = ref
	return this
}

func (this GPUError) Free() {
	this.Ref().Free()
}

// Message returns the value of property "GPUError.message".
//
// It returns ok=false if there is no such property.
func (this GPUError) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUErrorMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type GPUSupportedFeatures struct {
	ref js.Ref
}

func (this GPUSupportedFeatures) Once() GPUSupportedFeatures {
	this.Ref().Once()
	return this
}

func (this GPUSupportedFeatures) Ref() js.Ref {
	return this.ref
}

func (this GPUSupportedFeatures) FromRef(ref js.Ref) GPUSupportedFeatures {
	this.ref = ref
	return this
}

func (this GPUSupportedFeatures) Free() {
	this.Ref().Free()
}

type GPUSupportedLimits struct {
	ref js.Ref
}

func (this GPUSupportedLimits) Once() GPUSupportedLimits {
	this.Ref().Once()
	return this
}

func (this GPUSupportedLimits) Ref() js.Ref {
	return this.ref
}

func (this GPUSupportedLimits) FromRef(ref js.Ref) GPUSupportedLimits {
	this.ref = ref
	return this
}

func (this GPUSupportedLimits) Free() {
	this.Ref().Free()
}

// MaxTextureDimension1D returns the value of property "GPUSupportedLimits.maxTextureDimension1D".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxTextureDimension1D() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxTextureDimension1D(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxTextureDimension2D returns the value of property "GPUSupportedLimits.maxTextureDimension2D".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxTextureDimension2D() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxTextureDimension2D(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxTextureDimension3D returns the value of property "GPUSupportedLimits.maxTextureDimension3D".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxTextureDimension3D() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxTextureDimension3D(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxTextureArrayLayers returns the value of property "GPUSupportedLimits.maxTextureArrayLayers".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxTextureArrayLayers() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxTextureArrayLayers(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxBindGroups returns the value of property "GPUSupportedLimits.maxBindGroups".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxBindGroups() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxBindGroups(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxBindGroupsPlusVertexBuffers returns the value of property "GPUSupportedLimits.maxBindGroupsPlusVertexBuffers".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxBindGroupsPlusVertexBuffers() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxBindGroupsPlusVertexBuffers(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxBindingsPerBindGroup returns the value of property "GPUSupportedLimits.maxBindingsPerBindGroup".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxBindingsPerBindGroup() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxBindingsPerBindGroup(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxDynamicUniformBuffersPerPipelineLayout returns the value of property "GPUSupportedLimits.maxDynamicUniformBuffersPerPipelineLayout".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxDynamicUniformBuffersPerPipelineLayout() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxDynamicUniformBuffersPerPipelineLayout(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxDynamicStorageBuffersPerPipelineLayout returns the value of property "GPUSupportedLimits.maxDynamicStorageBuffersPerPipelineLayout".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxDynamicStorageBuffersPerPipelineLayout() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxDynamicStorageBuffersPerPipelineLayout(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxSampledTexturesPerShaderStage returns the value of property "GPUSupportedLimits.maxSampledTexturesPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxSampledTexturesPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxSampledTexturesPerShaderStage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxSamplersPerShaderStage returns the value of property "GPUSupportedLimits.maxSamplersPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxSamplersPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxSamplersPerShaderStage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxStorageBuffersPerShaderStage returns the value of property "GPUSupportedLimits.maxStorageBuffersPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxStorageBuffersPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxStorageBuffersPerShaderStage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxStorageTexturesPerShaderStage returns the value of property "GPUSupportedLimits.maxStorageTexturesPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxStorageTexturesPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxStorageTexturesPerShaderStage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxUniformBuffersPerShaderStage returns the value of property "GPUSupportedLimits.maxUniformBuffersPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxUniformBuffersPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxUniformBuffersPerShaderStage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxUniformBufferBindingSize returns the value of property "GPUSupportedLimits.maxUniformBufferBindingSize".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxUniformBufferBindingSize() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxUniformBufferBindingSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxStorageBufferBindingSize returns the value of property "GPUSupportedLimits.maxStorageBufferBindingSize".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxStorageBufferBindingSize() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxStorageBufferBindingSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MinUniformBufferOffsetAlignment returns the value of property "GPUSupportedLimits.minUniformBufferOffsetAlignment".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MinUniformBufferOffsetAlignment() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMinUniformBufferOffsetAlignment(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MinStorageBufferOffsetAlignment returns the value of property "GPUSupportedLimits.minStorageBufferOffsetAlignment".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MinStorageBufferOffsetAlignment() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMinStorageBufferOffsetAlignment(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxVertexBuffers returns the value of property "GPUSupportedLimits.maxVertexBuffers".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxVertexBuffers() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxVertexBuffers(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxBufferSize returns the value of property "GPUSupportedLimits.maxBufferSize".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxBufferSize() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxBufferSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxVertexAttributes returns the value of property "GPUSupportedLimits.maxVertexAttributes".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxVertexAttributes() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxVertexAttributes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxVertexBufferArrayStride returns the value of property "GPUSupportedLimits.maxVertexBufferArrayStride".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxVertexBufferArrayStride() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxVertexBufferArrayStride(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxInterStageShaderComponents returns the value of property "GPUSupportedLimits.maxInterStageShaderComponents".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxInterStageShaderComponents() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxInterStageShaderComponents(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxInterStageShaderVariables returns the value of property "GPUSupportedLimits.maxInterStageShaderVariables".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxInterStageShaderVariables() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxInterStageShaderVariables(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxColorAttachments returns the value of property "GPUSupportedLimits.maxColorAttachments".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxColorAttachments() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxColorAttachments(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxColorAttachmentBytesPerSample returns the value of property "GPUSupportedLimits.maxColorAttachmentBytesPerSample".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxColorAttachmentBytesPerSample() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxColorAttachmentBytesPerSample(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupStorageSize returns the value of property "GPUSupportedLimits.maxComputeWorkgroupStorageSize".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupStorageSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupStorageSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxComputeInvocationsPerWorkgroup returns the value of property "GPUSupportedLimits.maxComputeInvocationsPerWorkgroup".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeInvocationsPerWorkgroup() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeInvocationsPerWorkgroup(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupSizeX returns the value of property "GPUSupportedLimits.maxComputeWorkgroupSizeX".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupSizeX() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupSizeX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupSizeY returns the value of property "GPUSupportedLimits.maxComputeWorkgroupSizeY".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupSizeY() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupSizeY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupSizeZ returns the value of property "GPUSupportedLimits.maxComputeWorkgroupSizeZ".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupSizeZ() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupSizeZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupsPerDimension returns the value of property "GPUSupportedLimits.maxComputeWorkgroupsPerDimension".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupsPerDimension() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupsPerDimension(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type GPUImageDataLayout struct {
	// Offset is "GPUImageDataLayout.offset"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Offset MUST be set to true to make this field effective.
	Offset GPUSize64
	// BytesPerRow is "GPUImageDataLayout.bytesPerRow"
	//
	// Optional
	//
	// NOTE: FFI_USE_BytesPerRow MUST be set to true to make this field effective.
	BytesPerRow GPUSize32
	// RowsPerImage is "GPUImageDataLayout.rowsPerImage"
	//
	// Optional
	//
	// NOTE: FFI_USE_RowsPerImage MUST be set to true to make this field effective.
	RowsPerImage GPUSize32

	FFI_USE_Offset       bool // for Offset.
	FFI_USE_BytesPerRow  bool // for BytesPerRow.
	FFI_USE_RowsPerImage bool // for RowsPerImage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUImageDataLayout with all fields set.
func (p GPUImageDataLayout) FromRef(ref js.Ref) GPUImageDataLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUImageDataLayout in the application heap.
func (p GPUImageDataLayout) New() js.Ref {
	return bindings.GPUImageDataLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUImageDataLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUImageDataLayoutJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUImageDataLayout) Update(ref js.Ref) {
	bindings.GPUImageDataLayoutJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas struct {
	ref js.Ref
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) Free() {
	x.ref.Free()
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) FromRef(ref js.Ref) OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas {
	return OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas{
		ref: ref,
	}
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) ImageBitmap() ImageBitmap {
	return ImageBitmap{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) ImageData() ImageData {
	return ImageData{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) HTMLImageElement() HTMLImageElement {
	return HTMLImageElement{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) HTMLVideoElement() HTMLVideoElement {
	return HTMLVideoElement{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) VideoFrame() VideoFrame {
	return VideoFrame{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) HTMLCanvasElement() HTMLCanvasElement {
	return HTMLCanvasElement{}.FromRef(x.ref)
}

func (x OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas) OffscreenCanvas() OffscreenCanvas {
	return OffscreenCanvas{}.FromRef(x.ref)
}

type GPUImageCopyExternalImageSource = OneOf_ImageBitmap_ImageData_HTMLImageElement_HTMLVideoElement_VideoFrame_HTMLCanvasElement_OffscreenCanvas

type GPUOrigin2DDict struct {
	// X is "GPUOrigin2DDict.x"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_X MUST be set to true to make this field effective.
	X GPUIntegerCoordinate
	// Y is "GPUOrigin2DDict.y"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Y MUST be set to true to make this field effective.
	Y GPUIntegerCoordinate

	FFI_USE_X bool // for X.
	FFI_USE_Y bool // for Y.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUOrigin2DDict with all fields set.
func (p GPUOrigin2DDict) FromRef(ref js.Ref) GPUOrigin2DDict {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUOrigin2DDict in the application heap.
func (p GPUOrigin2DDict) New() js.Ref {
	return bindings.GPUOrigin2DDictJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUOrigin2DDict) UpdateFrom(ref js.Ref) {
	bindings.GPUOrigin2DDictJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUOrigin2DDict) Update(ref js.Ref) {
	bindings.GPUOrigin2DDictJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_ArrayGPUIntegerCoordinate_GPUOrigin2DDict struct {
	ref js.Ref
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin2DDict) Ref() js.Ref {
	return x.ref
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin2DDict) Free() {
	x.ref.Free()
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin2DDict) FromRef(ref js.Ref) OneOf_ArrayGPUIntegerCoordinate_GPUOrigin2DDict {
	return OneOf_ArrayGPUIntegerCoordinate_GPUOrigin2DDict{
		ref: ref,
	}
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin2DDict) ArrayGPUIntegerCoordinate() js.Array[GPUIntegerCoordinate] {
	return js.Array[GPUIntegerCoordinate]{}.FromRef(x.ref)
}

func (x OneOf_ArrayGPUIntegerCoordinate_GPUOrigin2DDict) GPUOrigin2DDict() GPUOrigin2DDict {
	var ret GPUOrigin2DDict
	ret.UpdateFrom(x.ref)
	return ret
}

type GPUOrigin2D = OneOf_ArrayGPUIntegerCoordinate_GPUOrigin2DDict

type GPUImageCopyExternalImage struct {
	// Source is "GPUImageCopyExternalImage.source"
	//
	// Required
	Source GPUImageCopyExternalImageSource
	// Origin is "GPUImageCopyExternalImage.origin"
	//
	// Optional, defaults to {}.
	Origin GPUOrigin2D
	// FlipY is "GPUImageCopyExternalImage.flipY"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_FlipY MUST be set to true to make this field effective.
	FlipY bool

	FFI_USE_FlipY bool // for FlipY.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUImageCopyExternalImage with all fields set.
func (p GPUImageCopyExternalImage) FromRef(ref js.Ref) GPUImageCopyExternalImage {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUImageCopyExternalImage in the application heap.
func (p GPUImageCopyExternalImage) New() js.Ref {
	return bindings.GPUImageCopyExternalImageJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUImageCopyExternalImage) UpdateFrom(ref js.Ref) {
	bindings.GPUImageCopyExternalImageJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUImageCopyExternalImage) Update(ref js.Ref) {
	bindings.GPUImageCopyExternalImageJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUImageCopyTextureTagged struct {
	// ColorSpace is "GPUImageCopyTextureTagged.colorSpace"
	//
	// Optional, defaults to "srgb".
	ColorSpace PredefinedColorSpace
	// PremultipliedAlpha is "GPUImageCopyTextureTagged.premultipliedAlpha"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PremultipliedAlpha MUST be set to true to make this field effective.
	PremultipliedAlpha bool
	// Texture is "GPUImageCopyTextureTagged.texture"
	//
	// Required
	Texture GPUTexture
	// MipLevel is "GPUImageCopyTextureTagged.mipLevel"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MipLevel MUST be set to true to make this field effective.
	MipLevel GPUIntegerCoordinate
	// Origin is "GPUImageCopyTextureTagged.origin"
	//
	// Optional, defaults to {}.
	Origin GPUOrigin3D
	// Aspect is "GPUImageCopyTextureTagged.aspect"
	//
	// Optional, defaults to "all".
	Aspect GPUTextureAspect

	FFI_USE_PremultipliedAlpha bool // for PremultipliedAlpha.
	FFI_USE_MipLevel           bool // for MipLevel.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUImageCopyTextureTagged with all fields set.
func (p GPUImageCopyTextureTagged) FromRef(ref js.Ref) GPUImageCopyTextureTagged {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUImageCopyTextureTagged in the application heap.
func (p GPUImageCopyTextureTagged) New() js.Ref {
	return bindings.GPUImageCopyTextureTaggedJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUImageCopyTextureTagged) UpdateFrom(ref js.Ref) {
	bindings.GPUImageCopyTextureTaggedJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUImageCopyTextureTagged) Update(ref js.Ref) {
	bindings.GPUImageCopyTextureTaggedJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUQueue struct {
	ref js.Ref
}

func (this GPUQueue) Once() GPUQueue {
	this.Ref().Once()
	return this
}

func (this GPUQueue) Ref() js.Ref {
	return this.ref
}

func (this GPUQueue) FromRef(ref js.Ref) GPUQueue {
	this.ref = ref
	return this
}

func (this GPUQueue) Free() {
	this.Ref().Free()
}

// Label returns the value of property "GPUQueue.label".
//
// It returns ok=false if there is no such property.
func (this GPUQueue) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUQueueLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUQueue.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUQueue) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUQueueLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasSubmit returns true if the method "GPUQueue.submit" exists.
func (this GPUQueue) HasSubmit() bool {
	return js.True == bindings.HasGPUQueueSubmit(
		this.Ref(),
	)
}

// SubmitFunc returns the method "GPUQueue.submit".
func (this GPUQueue) SubmitFunc() (fn js.Func[func(commandBuffers js.Array[GPUCommandBuffer])]) {
	return fn.FromRef(
		bindings.GPUQueueSubmitFunc(
			this.Ref(),
		),
	)
}

// Submit calls the method "GPUQueue.submit".
func (this GPUQueue) Submit(commandBuffers js.Array[GPUCommandBuffer]) (ret js.Void) {
	bindings.CallGPUQueueSubmit(
		this.Ref(), js.Pointer(&ret),
		commandBuffers.Ref(),
	)

	return
}

// TrySubmit calls the method "GPUQueue.submit"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQueue) TrySubmit(commandBuffers js.Array[GPUCommandBuffer]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQueueSubmit(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		commandBuffers.Ref(),
	)

	return
}

// HasOnSubmittedWorkDone returns true if the method "GPUQueue.onSubmittedWorkDone" exists.
func (this GPUQueue) HasOnSubmittedWorkDone() bool {
	return js.True == bindings.HasGPUQueueOnSubmittedWorkDone(
		this.Ref(),
	)
}

// OnSubmittedWorkDoneFunc returns the method "GPUQueue.onSubmittedWorkDone".
func (this GPUQueue) OnSubmittedWorkDoneFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.GPUQueueOnSubmittedWorkDoneFunc(
			this.Ref(),
		),
	)
}

// OnSubmittedWorkDone calls the method "GPUQueue.onSubmittedWorkDone".
func (this GPUQueue) OnSubmittedWorkDone() (ret js.Promise[js.Void]) {
	bindings.CallGPUQueueOnSubmittedWorkDone(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryOnSubmittedWorkDone calls the method "GPUQueue.onSubmittedWorkDone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQueue) TryOnSubmittedWorkDone() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQueueOnSubmittedWorkDone(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasWriteBuffer returns true if the method "GPUQueue.writeBuffer" exists.
func (this GPUQueue) HasWriteBuffer() bool {
	return js.True == bindings.HasGPUQueueWriteBuffer(
		this.Ref(),
	)
}

// WriteBufferFunc returns the method "GPUQueue.writeBuffer".
func (this GPUQueue) WriteBufferFunc() (fn js.Func[func(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUQueueWriteBufferFunc(
			this.Ref(),
		),
	)
}

// WriteBuffer calls the method "GPUQueue.writeBuffer".
func (this GPUQueue) WriteBuffer(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPUQueueWriteBuffer(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
		float64(dataOffset),
		float64(size),
	)

	return
}

// TryWriteBuffer calls the method "GPUQueue.writeBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQueue) TryWriteBuffer(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64, size GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQueueWriteBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
		float64(dataOffset),
		float64(size),
	)

	return
}

// HasWriteBuffer1 returns true if the method "GPUQueue.writeBuffer" exists.
func (this GPUQueue) HasWriteBuffer1() bool {
	return js.True == bindings.HasGPUQueueWriteBuffer1(
		this.Ref(),
	)
}

// WriteBuffer1Func returns the method "GPUQueue.writeBuffer".
func (this GPUQueue) WriteBuffer1Func() (fn js.Func[func(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUQueueWriteBuffer1Func(
			this.Ref(),
		),
	)
}

// WriteBuffer1 calls the method "GPUQueue.writeBuffer".
func (this GPUQueue) WriteBuffer1(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64) (ret js.Void) {
	bindings.CallGPUQueueWriteBuffer1(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
		float64(dataOffset),
	)

	return
}

// TryWriteBuffer1 calls the method "GPUQueue.writeBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQueue) TryWriteBuffer1(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQueueWriteBuffer1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
		float64(dataOffset),
	)

	return
}

// HasWriteBuffer2 returns true if the method "GPUQueue.writeBuffer" exists.
func (this GPUQueue) HasWriteBuffer2() bool {
	return js.True == bindings.HasGPUQueueWriteBuffer2(
		this.Ref(),
	)
}

// WriteBuffer2Func returns the method "GPUQueue.writeBuffer".
func (this GPUQueue) WriteBuffer2Func() (fn js.Func[func(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource)]) {
	return fn.FromRef(
		bindings.GPUQueueWriteBuffer2Func(
			this.Ref(),
		),
	)
}

// WriteBuffer2 calls the method "GPUQueue.writeBuffer".
func (this GPUQueue) WriteBuffer2(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource) (ret js.Void) {
	bindings.CallGPUQueueWriteBuffer2(
		this.Ref(), js.Pointer(&ret),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
	)

	return
}

// TryWriteBuffer2 calls the method "GPUQueue.writeBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQueue) TryWriteBuffer2(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQueueWriteBuffer2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
	)

	return
}

// HasWriteTexture returns true if the method "GPUQueue.writeTexture" exists.
func (this GPUQueue) HasWriteTexture() bool {
	return js.True == bindings.HasGPUQueueWriteTexture(
		this.Ref(),
	)
}

// WriteTextureFunc returns the method "GPUQueue.writeTexture".
func (this GPUQueue) WriteTextureFunc() (fn js.Func[func(destination GPUImageCopyTexture, data AllowSharedBufferSource, dataLayout GPUImageDataLayout, size GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUQueueWriteTextureFunc(
			this.Ref(),
		),
	)
}

// WriteTexture calls the method "GPUQueue.writeTexture".
func (this GPUQueue) WriteTexture(destination GPUImageCopyTexture, data AllowSharedBufferSource, dataLayout GPUImageDataLayout, size GPUExtent3D) (ret js.Void) {
	bindings.CallGPUQueueWriteTexture(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&destination),
		data.Ref(),
		js.Pointer(&dataLayout),
		size.Ref(),
	)

	return
}

// TryWriteTexture calls the method "GPUQueue.writeTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQueue) TryWriteTexture(destination GPUImageCopyTexture, data AllowSharedBufferSource, dataLayout GPUImageDataLayout, size GPUExtent3D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQueueWriteTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&destination),
		data.Ref(),
		js.Pointer(&dataLayout),
		size.Ref(),
	)

	return
}

// HasCopyExternalImageToTexture returns true if the method "GPUQueue.copyExternalImageToTexture" exists.
func (this GPUQueue) HasCopyExternalImageToTexture() bool {
	return js.True == bindings.HasGPUQueueCopyExternalImageToTexture(
		this.Ref(),
	)
}

// CopyExternalImageToTextureFunc returns the method "GPUQueue.copyExternalImageToTexture".
func (this GPUQueue) CopyExternalImageToTextureFunc() (fn js.Func[func(source GPUImageCopyExternalImage, destination GPUImageCopyTextureTagged, copySize GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUQueueCopyExternalImageToTextureFunc(
			this.Ref(),
		),
	)
}

// CopyExternalImageToTexture calls the method "GPUQueue.copyExternalImageToTexture".
func (this GPUQueue) CopyExternalImageToTexture(source GPUImageCopyExternalImage, destination GPUImageCopyTextureTagged, copySize GPUExtent3D) (ret js.Void) {
	bindings.CallGPUQueueCopyExternalImageToTexture(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// TryCopyExternalImageToTexture calls the method "GPUQueue.copyExternalImageToTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQueue) TryCopyExternalImageToTexture(source GPUImageCopyExternalImage, destination GPUImageCopyTextureTagged, copySize GPUExtent3D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQueueCopyExternalImageToTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

type GPUDeviceLostReason uint32

const (
	_ GPUDeviceLostReason = iota

	GPUDeviceLostReason_UNKNOWN
	GPUDeviceLostReason_DESTROYED
)

func (GPUDeviceLostReason) FromRef(str js.Ref) GPUDeviceLostReason {
	return GPUDeviceLostReason(bindings.ConstOfGPUDeviceLostReason(str))
}

func (x GPUDeviceLostReason) String() (string, bool) {
	switch x {
	case GPUDeviceLostReason_UNKNOWN:
		return "unknown", true
	case GPUDeviceLostReason_DESTROYED:
		return "destroyed", true
	default:
		return "", false
	}
}

type GPUDeviceLostInfo struct {
	ref js.Ref
}

func (this GPUDeviceLostInfo) Once() GPUDeviceLostInfo {
	this.Ref().Once()
	return this
}

func (this GPUDeviceLostInfo) Ref() js.Ref {
	return this.ref
}

func (this GPUDeviceLostInfo) FromRef(ref js.Ref) GPUDeviceLostInfo {
	this.ref = ref
	return this
}

func (this GPUDeviceLostInfo) Free() {
	this.Ref().Free()
}

// Reason returns the value of property "GPUDeviceLostInfo.reason".
//
// It returns ok=false if there is no such property.
func (this GPUDeviceLostInfo) Reason() (ret GPUDeviceLostReason, ok bool) {
	ok = js.True == bindings.GetGPUDeviceLostInfoReason(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "GPUDeviceLostInfo.message".
//
// It returns ok=false if there is no such property.
func (this GPUDeviceLostInfo) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUDeviceLostInfoMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type GPUDevice struct {
	EventTarget
}

func (this GPUDevice) Once() GPUDevice {
	this.Ref().Once()
	return this
}

func (this GPUDevice) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this GPUDevice) FromRef(ref js.Ref) GPUDevice {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this GPUDevice) Free() {
	this.Ref().Free()
}

// Features returns the value of property "GPUDevice.features".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Features() (ret GPUSupportedFeatures, ok bool) {
	ok = js.True == bindings.GetGPUDeviceFeatures(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Limits returns the value of property "GPUDevice.limits".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Limits() (ret GPUSupportedLimits, ok bool) {
	ok = js.True == bindings.GetGPUDeviceLimits(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Queue returns the value of property "GPUDevice.queue".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Queue() (ret GPUQueue, ok bool) {
	ok = js.True == bindings.GetGPUDeviceQueue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Lost returns the value of property "GPUDevice.lost".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Lost() (ret js.Promise[GPUDeviceLostInfo], ok bool) {
	ok = js.True == bindings.GetGPUDeviceLost(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "GPUDevice.label".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUDeviceLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUDevice.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUDevice) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUDeviceLabel(
		this.Ref(),
		val.Ref(),
	)
}

// HasDestroy returns true if the method "GPUDevice.destroy" exists.
func (this GPUDevice) HasDestroy() bool {
	return js.True == bindings.HasGPUDeviceDestroy(
		this.Ref(),
	)
}

// DestroyFunc returns the method "GPUDevice.destroy".
func (this GPUDevice) DestroyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUDeviceDestroyFunc(
			this.Ref(),
		),
	)
}

// Destroy calls the method "GPUDevice.destroy".
func (this GPUDevice) Destroy() (ret js.Void) {
	bindings.CallGPUDeviceDestroy(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDestroy calls the method "GPUDevice.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryDestroy() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceDestroy(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateBuffer returns true if the method "GPUDevice.createBuffer" exists.
func (this GPUDevice) HasCreateBuffer() bool {
	return js.True == bindings.HasGPUDeviceCreateBuffer(
		this.Ref(),
	)
}

// CreateBufferFunc returns the method "GPUDevice.createBuffer".
func (this GPUDevice) CreateBufferFunc() (fn js.Func[func(descriptor GPUBufferDescriptor) GPUBuffer]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateBufferFunc(
			this.Ref(),
		),
	)
}

// CreateBuffer calls the method "GPUDevice.createBuffer".
func (this GPUDevice) CreateBuffer(descriptor GPUBufferDescriptor) (ret GPUBuffer) {
	bindings.CallGPUDeviceCreateBuffer(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateBuffer calls the method "GPUDevice.createBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateBuffer(descriptor GPUBufferDescriptor) (ret GPUBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateTexture returns true if the method "GPUDevice.createTexture" exists.
func (this GPUDevice) HasCreateTexture() bool {
	return js.True == bindings.HasGPUDeviceCreateTexture(
		this.Ref(),
	)
}

// CreateTextureFunc returns the method "GPUDevice.createTexture".
func (this GPUDevice) CreateTextureFunc() (fn js.Func[func(descriptor GPUTextureDescriptor) GPUTexture]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateTextureFunc(
			this.Ref(),
		),
	)
}

// CreateTexture calls the method "GPUDevice.createTexture".
func (this GPUDevice) CreateTexture(descriptor GPUTextureDescriptor) (ret GPUTexture) {
	bindings.CallGPUDeviceCreateTexture(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateTexture calls the method "GPUDevice.createTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateTexture(descriptor GPUTextureDescriptor) (ret GPUTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateSampler returns true if the method "GPUDevice.createSampler" exists.
func (this GPUDevice) HasCreateSampler() bool {
	return js.True == bindings.HasGPUDeviceCreateSampler(
		this.Ref(),
	)
}

// CreateSamplerFunc returns the method "GPUDevice.createSampler".
func (this GPUDevice) CreateSamplerFunc() (fn js.Func[func(descriptor GPUSamplerDescriptor) GPUSampler]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateSamplerFunc(
			this.Ref(),
		),
	)
}

// CreateSampler calls the method "GPUDevice.createSampler".
func (this GPUDevice) CreateSampler(descriptor GPUSamplerDescriptor) (ret GPUSampler) {
	bindings.CallGPUDeviceCreateSampler(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateSampler calls the method "GPUDevice.createSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateSampler(descriptor GPUSamplerDescriptor) (ret GPUSampler, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateSampler(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateSampler1 returns true if the method "GPUDevice.createSampler" exists.
func (this GPUDevice) HasCreateSampler1() bool {
	return js.True == bindings.HasGPUDeviceCreateSampler1(
		this.Ref(),
	)
}

// CreateSampler1Func returns the method "GPUDevice.createSampler".
func (this GPUDevice) CreateSampler1Func() (fn js.Func[func() GPUSampler]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateSampler1Func(
			this.Ref(),
		),
	)
}

// CreateSampler1 calls the method "GPUDevice.createSampler".
func (this GPUDevice) CreateSampler1() (ret GPUSampler) {
	bindings.CallGPUDeviceCreateSampler1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateSampler1 calls the method "GPUDevice.createSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateSampler1() (ret GPUSampler, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateSampler1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasImportExternalTexture returns true if the method "GPUDevice.importExternalTexture" exists.
func (this GPUDevice) HasImportExternalTexture() bool {
	return js.True == bindings.HasGPUDeviceImportExternalTexture(
		this.Ref(),
	)
}

// ImportExternalTextureFunc returns the method "GPUDevice.importExternalTexture".
func (this GPUDevice) ImportExternalTextureFunc() (fn js.Func[func(descriptor GPUExternalTextureDescriptor) GPUExternalTexture]) {
	return fn.FromRef(
		bindings.GPUDeviceImportExternalTextureFunc(
			this.Ref(),
		),
	)
}

// ImportExternalTexture calls the method "GPUDevice.importExternalTexture".
func (this GPUDevice) ImportExternalTexture(descriptor GPUExternalTextureDescriptor) (ret GPUExternalTexture) {
	bindings.CallGPUDeviceImportExternalTexture(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryImportExternalTexture calls the method "GPUDevice.importExternalTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryImportExternalTexture(descriptor GPUExternalTextureDescriptor) (ret GPUExternalTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceImportExternalTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateBindGroupLayout returns true if the method "GPUDevice.createBindGroupLayout" exists.
func (this GPUDevice) HasCreateBindGroupLayout() bool {
	return js.True == bindings.HasGPUDeviceCreateBindGroupLayout(
		this.Ref(),
	)
}

// CreateBindGroupLayoutFunc returns the method "GPUDevice.createBindGroupLayout".
func (this GPUDevice) CreateBindGroupLayoutFunc() (fn js.Func[func(descriptor GPUBindGroupLayoutDescriptor) GPUBindGroupLayout]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateBindGroupLayoutFunc(
			this.Ref(),
		),
	)
}

// CreateBindGroupLayout calls the method "GPUDevice.createBindGroupLayout".
func (this GPUDevice) CreateBindGroupLayout(descriptor GPUBindGroupLayoutDescriptor) (ret GPUBindGroupLayout) {
	bindings.CallGPUDeviceCreateBindGroupLayout(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateBindGroupLayout calls the method "GPUDevice.createBindGroupLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateBindGroupLayout(descriptor GPUBindGroupLayoutDescriptor) (ret GPUBindGroupLayout, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateBindGroupLayout(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreatePipelineLayout returns true if the method "GPUDevice.createPipelineLayout" exists.
func (this GPUDevice) HasCreatePipelineLayout() bool {
	return js.True == bindings.HasGPUDeviceCreatePipelineLayout(
		this.Ref(),
	)
}

// CreatePipelineLayoutFunc returns the method "GPUDevice.createPipelineLayout".
func (this GPUDevice) CreatePipelineLayoutFunc() (fn js.Func[func(descriptor GPUPipelineLayoutDescriptor) GPUPipelineLayout]) {
	return fn.FromRef(
		bindings.GPUDeviceCreatePipelineLayoutFunc(
			this.Ref(),
		),
	)
}

// CreatePipelineLayout calls the method "GPUDevice.createPipelineLayout".
func (this GPUDevice) CreatePipelineLayout(descriptor GPUPipelineLayoutDescriptor) (ret GPUPipelineLayout) {
	bindings.CallGPUDeviceCreatePipelineLayout(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreatePipelineLayout calls the method "GPUDevice.createPipelineLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreatePipelineLayout(descriptor GPUPipelineLayoutDescriptor) (ret GPUPipelineLayout, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreatePipelineLayout(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateBindGroup returns true if the method "GPUDevice.createBindGroup" exists.
func (this GPUDevice) HasCreateBindGroup() bool {
	return js.True == bindings.HasGPUDeviceCreateBindGroup(
		this.Ref(),
	)
}

// CreateBindGroupFunc returns the method "GPUDevice.createBindGroup".
func (this GPUDevice) CreateBindGroupFunc() (fn js.Func[func(descriptor GPUBindGroupDescriptor) GPUBindGroup]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateBindGroupFunc(
			this.Ref(),
		),
	)
}

// CreateBindGroup calls the method "GPUDevice.createBindGroup".
func (this GPUDevice) CreateBindGroup(descriptor GPUBindGroupDescriptor) (ret GPUBindGroup) {
	bindings.CallGPUDeviceCreateBindGroup(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateBindGroup calls the method "GPUDevice.createBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateBindGroup(descriptor GPUBindGroupDescriptor) (ret GPUBindGroup, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateBindGroup(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateShaderModule returns true if the method "GPUDevice.createShaderModule" exists.
func (this GPUDevice) HasCreateShaderModule() bool {
	return js.True == bindings.HasGPUDeviceCreateShaderModule(
		this.Ref(),
	)
}

// CreateShaderModuleFunc returns the method "GPUDevice.createShaderModule".
func (this GPUDevice) CreateShaderModuleFunc() (fn js.Func[func(descriptor GPUShaderModuleDescriptor) GPUShaderModule]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateShaderModuleFunc(
			this.Ref(),
		),
	)
}

// CreateShaderModule calls the method "GPUDevice.createShaderModule".
func (this GPUDevice) CreateShaderModule(descriptor GPUShaderModuleDescriptor) (ret GPUShaderModule) {
	bindings.CallGPUDeviceCreateShaderModule(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateShaderModule calls the method "GPUDevice.createShaderModule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateShaderModule(descriptor GPUShaderModuleDescriptor) (ret GPUShaderModule, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateShaderModule(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateComputePipeline returns true if the method "GPUDevice.createComputePipeline" exists.
func (this GPUDevice) HasCreateComputePipeline() bool {
	return js.True == bindings.HasGPUDeviceCreateComputePipeline(
		this.Ref(),
	)
}

// CreateComputePipelineFunc returns the method "GPUDevice.createComputePipeline".
func (this GPUDevice) CreateComputePipelineFunc() (fn js.Func[func(descriptor GPUComputePipelineDescriptor) GPUComputePipeline]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateComputePipelineFunc(
			this.Ref(),
		),
	)
}

// CreateComputePipeline calls the method "GPUDevice.createComputePipeline".
func (this GPUDevice) CreateComputePipeline(descriptor GPUComputePipelineDescriptor) (ret GPUComputePipeline) {
	bindings.CallGPUDeviceCreateComputePipeline(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateComputePipeline calls the method "GPUDevice.createComputePipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateComputePipeline(descriptor GPUComputePipelineDescriptor) (ret GPUComputePipeline, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateComputePipeline(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateRenderPipeline returns true if the method "GPUDevice.createRenderPipeline" exists.
func (this GPUDevice) HasCreateRenderPipeline() bool {
	return js.True == bindings.HasGPUDeviceCreateRenderPipeline(
		this.Ref(),
	)
}

// CreateRenderPipelineFunc returns the method "GPUDevice.createRenderPipeline".
func (this GPUDevice) CreateRenderPipelineFunc() (fn js.Func[func(descriptor GPURenderPipelineDescriptor) GPURenderPipeline]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateRenderPipelineFunc(
			this.Ref(),
		),
	)
}

// CreateRenderPipeline calls the method "GPUDevice.createRenderPipeline".
func (this GPUDevice) CreateRenderPipeline(descriptor GPURenderPipelineDescriptor) (ret GPURenderPipeline) {
	bindings.CallGPUDeviceCreateRenderPipeline(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateRenderPipeline calls the method "GPUDevice.createRenderPipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateRenderPipeline(descriptor GPURenderPipelineDescriptor) (ret GPURenderPipeline, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateRenderPipeline(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateComputePipelineAsync returns true if the method "GPUDevice.createComputePipelineAsync" exists.
func (this GPUDevice) HasCreateComputePipelineAsync() bool {
	return js.True == bindings.HasGPUDeviceCreateComputePipelineAsync(
		this.Ref(),
	)
}

// CreateComputePipelineAsyncFunc returns the method "GPUDevice.createComputePipelineAsync".
func (this GPUDevice) CreateComputePipelineAsyncFunc() (fn js.Func[func(descriptor GPUComputePipelineDescriptor) js.Promise[GPUComputePipeline]]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateComputePipelineAsyncFunc(
			this.Ref(),
		),
	)
}

// CreateComputePipelineAsync calls the method "GPUDevice.createComputePipelineAsync".
func (this GPUDevice) CreateComputePipelineAsync(descriptor GPUComputePipelineDescriptor) (ret js.Promise[GPUComputePipeline]) {
	bindings.CallGPUDeviceCreateComputePipelineAsync(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateComputePipelineAsync calls the method "GPUDevice.createComputePipelineAsync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateComputePipelineAsync(descriptor GPUComputePipelineDescriptor) (ret js.Promise[GPUComputePipeline], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateComputePipelineAsync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateRenderPipelineAsync returns true if the method "GPUDevice.createRenderPipelineAsync" exists.
func (this GPUDevice) HasCreateRenderPipelineAsync() bool {
	return js.True == bindings.HasGPUDeviceCreateRenderPipelineAsync(
		this.Ref(),
	)
}

// CreateRenderPipelineAsyncFunc returns the method "GPUDevice.createRenderPipelineAsync".
func (this GPUDevice) CreateRenderPipelineAsyncFunc() (fn js.Func[func(descriptor GPURenderPipelineDescriptor) js.Promise[GPURenderPipeline]]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateRenderPipelineAsyncFunc(
			this.Ref(),
		),
	)
}

// CreateRenderPipelineAsync calls the method "GPUDevice.createRenderPipelineAsync".
func (this GPUDevice) CreateRenderPipelineAsync(descriptor GPURenderPipelineDescriptor) (ret js.Promise[GPURenderPipeline]) {
	bindings.CallGPUDeviceCreateRenderPipelineAsync(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateRenderPipelineAsync calls the method "GPUDevice.createRenderPipelineAsync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateRenderPipelineAsync(descriptor GPURenderPipelineDescriptor) (ret js.Promise[GPURenderPipeline], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateRenderPipelineAsync(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateCommandEncoder returns true if the method "GPUDevice.createCommandEncoder" exists.
func (this GPUDevice) HasCreateCommandEncoder() bool {
	return js.True == bindings.HasGPUDeviceCreateCommandEncoder(
		this.Ref(),
	)
}

// CreateCommandEncoderFunc returns the method "GPUDevice.createCommandEncoder".
func (this GPUDevice) CreateCommandEncoderFunc() (fn js.Func[func(descriptor GPUCommandEncoderDescriptor) GPUCommandEncoder]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateCommandEncoderFunc(
			this.Ref(),
		),
	)
}

// CreateCommandEncoder calls the method "GPUDevice.createCommandEncoder".
func (this GPUDevice) CreateCommandEncoder(descriptor GPUCommandEncoderDescriptor) (ret GPUCommandEncoder) {
	bindings.CallGPUDeviceCreateCommandEncoder(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateCommandEncoder calls the method "GPUDevice.createCommandEncoder"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateCommandEncoder(descriptor GPUCommandEncoderDescriptor) (ret GPUCommandEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateCommandEncoder(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateCommandEncoder1 returns true if the method "GPUDevice.createCommandEncoder" exists.
func (this GPUDevice) HasCreateCommandEncoder1() bool {
	return js.True == bindings.HasGPUDeviceCreateCommandEncoder1(
		this.Ref(),
	)
}

// CreateCommandEncoder1Func returns the method "GPUDevice.createCommandEncoder".
func (this GPUDevice) CreateCommandEncoder1Func() (fn js.Func[func() GPUCommandEncoder]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateCommandEncoder1Func(
			this.Ref(),
		),
	)
}

// CreateCommandEncoder1 calls the method "GPUDevice.createCommandEncoder".
func (this GPUDevice) CreateCommandEncoder1() (ret GPUCommandEncoder) {
	bindings.CallGPUDeviceCreateCommandEncoder1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateCommandEncoder1 calls the method "GPUDevice.createCommandEncoder"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateCommandEncoder1() (ret GPUCommandEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateCommandEncoder1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateRenderBundleEncoder returns true if the method "GPUDevice.createRenderBundleEncoder" exists.
func (this GPUDevice) HasCreateRenderBundleEncoder() bool {
	return js.True == bindings.HasGPUDeviceCreateRenderBundleEncoder(
		this.Ref(),
	)
}

// CreateRenderBundleEncoderFunc returns the method "GPUDevice.createRenderBundleEncoder".
func (this GPUDevice) CreateRenderBundleEncoderFunc() (fn js.Func[func(descriptor GPURenderBundleEncoderDescriptor) GPURenderBundleEncoder]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateRenderBundleEncoderFunc(
			this.Ref(),
		),
	)
}

// CreateRenderBundleEncoder calls the method "GPUDevice.createRenderBundleEncoder".
func (this GPUDevice) CreateRenderBundleEncoder(descriptor GPURenderBundleEncoderDescriptor) (ret GPURenderBundleEncoder) {
	bindings.CallGPUDeviceCreateRenderBundleEncoder(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateRenderBundleEncoder calls the method "GPUDevice.createRenderBundleEncoder"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateRenderBundleEncoder(descriptor GPURenderBundleEncoderDescriptor) (ret GPURenderBundleEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateRenderBundleEncoder(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasCreateQuerySet returns true if the method "GPUDevice.createQuerySet" exists.
func (this GPUDevice) HasCreateQuerySet() bool {
	return js.True == bindings.HasGPUDeviceCreateQuerySet(
		this.Ref(),
	)
}

// CreateQuerySetFunc returns the method "GPUDevice.createQuerySet".
func (this GPUDevice) CreateQuerySetFunc() (fn js.Func[func(descriptor GPUQuerySetDescriptor) GPUQuerySet]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateQuerySetFunc(
			this.Ref(),
		),
	)
}

// CreateQuerySet calls the method "GPUDevice.createQuerySet".
func (this GPUDevice) CreateQuerySet(descriptor GPUQuerySetDescriptor) (ret GPUQuerySet) {
	bindings.CallGPUDeviceCreateQuerySet(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateQuerySet calls the method "GPUDevice.createQuerySet"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateQuerySet(descriptor GPUQuerySetDescriptor) (ret GPUQuerySet, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateQuerySet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasPushErrorScope returns true if the method "GPUDevice.pushErrorScope" exists.
func (this GPUDevice) HasPushErrorScope() bool {
	return js.True == bindings.HasGPUDevicePushErrorScope(
		this.Ref(),
	)
}

// PushErrorScopeFunc returns the method "GPUDevice.pushErrorScope".
func (this GPUDevice) PushErrorScopeFunc() (fn js.Func[func(filter GPUErrorFilter)]) {
	return fn.FromRef(
		bindings.GPUDevicePushErrorScopeFunc(
			this.Ref(),
		),
	)
}

// PushErrorScope calls the method "GPUDevice.pushErrorScope".
func (this GPUDevice) PushErrorScope(filter GPUErrorFilter) (ret js.Void) {
	bindings.CallGPUDevicePushErrorScope(
		this.Ref(), js.Pointer(&ret),
		uint32(filter),
	)

	return
}

// TryPushErrorScope calls the method "GPUDevice.pushErrorScope"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryPushErrorScope(filter GPUErrorFilter) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDevicePushErrorScope(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(filter),
	)

	return
}

// HasPopErrorScope returns true if the method "GPUDevice.popErrorScope" exists.
func (this GPUDevice) HasPopErrorScope() bool {
	return js.True == bindings.HasGPUDevicePopErrorScope(
		this.Ref(),
	)
}

// PopErrorScopeFunc returns the method "GPUDevice.popErrorScope".
func (this GPUDevice) PopErrorScopeFunc() (fn js.Func[func() js.Promise[GPUError]]) {
	return fn.FromRef(
		bindings.GPUDevicePopErrorScopeFunc(
			this.Ref(),
		),
	)
}

// PopErrorScope calls the method "GPUDevice.popErrorScope".
func (this GPUDevice) PopErrorScope() (ret js.Promise[GPUError]) {
	bindings.CallGPUDevicePopErrorScope(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPopErrorScope calls the method "GPUDevice.popErrorScope"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryPopErrorScope() (ret js.Promise[GPUError], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDevicePopErrorScope(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type GPUCanvasAlphaMode uint32

const (
	_ GPUCanvasAlphaMode = iota

	GPUCanvasAlphaMode_OPAQUE
	GPUCanvasAlphaMode_PREMULTIPLIED
)

func (GPUCanvasAlphaMode) FromRef(str js.Ref) GPUCanvasAlphaMode {
	return GPUCanvasAlphaMode(bindings.ConstOfGPUCanvasAlphaMode(str))
}

func (x GPUCanvasAlphaMode) String() (string, bool) {
	switch x {
	case GPUCanvasAlphaMode_OPAQUE:
		return "opaque", true
	case GPUCanvasAlphaMode_PREMULTIPLIED:
		return "premultiplied", true
	default:
		return "", false
	}
}

type GPUCanvasConfiguration struct {
	// Device is "GPUCanvasConfiguration.device"
	//
	// Required
	Device GPUDevice
	// Format is "GPUCanvasConfiguration.format"
	//
	// Required
	Format GPUTextureFormat
	// Usage is "GPUCanvasConfiguration.usage"
	//
	// Optional, defaults to 0x10.
	//
	// NOTE: FFI_USE_Usage MUST be set to true to make this field effective.
	Usage GPUTextureUsageFlags
	// ViewFormats is "GPUCanvasConfiguration.viewFormats"
	//
	// Optional, defaults to [].
	ViewFormats js.Array[GPUTextureFormat]
	// ColorSpace is "GPUCanvasConfiguration.colorSpace"
	//
	// Optional, defaults to "srgb".
	ColorSpace PredefinedColorSpace
	// AlphaMode is "GPUCanvasConfiguration.alphaMode"
	//
	// Optional, defaults to "opaque".
	AlphaMode GPUCanvasAlphaMode

	FFI_USE_Usage bool // for Usage.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUCanvasConfiguration with all fields set.
func (p GPUCanvasConfiguration) FromRef(ref js.Ref) GPUCanvasConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUCanvasConfiguration in the application heap.
func (p GPUCanvasConfiguration) New() js.Ref {
	return bindings.GPUCanvasConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUCanvasConfiguration) UpdateFrom(ref js.Ref) {
	bindings.GPUCanvasConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUCanvasConfiguration) Update(ref js.Ref) {
	bindings.GPUCanvasConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GPUCanvasContext struct {
	ref js.Ref
}

func (this GPUCanvasContext) Once() GPUCanvasContext {
	this.Ref().Once()
	return this
}

func (this GPUCanvasContext) Ref() js.Ref {
	return this.ref
}

func (this GPUCanvasContext) FromRef(ref js.Ref) GPUCanvasContext {
	this.ref = ref
	return this
}

func (this GPUCanvasContext) Free() {
	this.Ref().Free()
}

// Canvas returns the value of property "GPUCanvasContext.canvas".
//
// It returns ok=false if there is no such property.
func (this GPUCanvasContext) Canvas() (ret OneOf_HTMLCanvasElement_OffscreenCanvas, ok bool) {
	ok = js.True == bindings.GetGPUCanvasContextCanvas(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasConfigure returns true if the method "GPUCanvasContext.configure" exists.
func (this GPUCanvasContext) HasConfigure() bool {
	return js.True == bindings.HasGPUCanvasContextConfigure(
		this.Ref(),
	)
}

// ConfigureFunc returns the method "GPUCanvasContext.configure".
func (this GPUCanvasContext) ConfigureFunc() (fn js.Func[func(configuration GPUCanvasConfiguration)]) {
	return fn.FromRef(
		bindings.GPUCanvasContextConfigureFunc(
			this.Ref(),
		),
	)
}

// Configure calls the method "GPUCanvasContext.configure".
func (this GPUCanvasContext) Configure(configuration GPUCanvasConfiguration) (ret js.Void) {
	bindings.CallGPUCanvasContextConfigure(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&configuration),
	)

	return
}

// TryConfigure calls the method "GPUCanvasContext.configure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCanvasContext) TryConfigure(configuration GPUCanvasConfiguration) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCanvasContextConfigure(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&configuration),
	)

	return
}

// HasUnconfigure returns true if the method "GPUCanvasContext.unconfigure" exists.
func (this GPUCanvasContext) HasUnconfigure() bool {
	return js.True == bindings.HasGPUCanvasContextUnconfigure(
		this.Ref(),
	)
}

// UnconfigureFunc returns the method "GPUCanvasContext.unconfigure".
func (this GPUCanvasContext) UnconfigureFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUCanvasContextUnconfigureFunc(
			this.Ref(),
		),
	)
}

// Unconfigure calls the method "GPUCanvasContext.unconfigure".
func (this GPUCanvasContext) Unconfigure() (ret js.Void) {
	bindings.CallGPUCanvasContextUnconfigure(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryUnconfigure calls the method "GPUCanvasContext.unconfigure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCanvasContext) TryUnconfigure() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCanvasContextUnconfigure(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetCurrentTexture returns true if the method "GPUCanvasContext.getCurrentTexture" exists.
func (this GPUCanvasContext) HasGetCurrentTexture() bool {
	return js.True == bindings.HasGPUCanvasContextGetCurrentTexture(
		this.Ref(),
	)
}

// GetCurrentTextureFunc returns the method "GPUCanvasContext.getCurrentTexture".
func (this GPUCanvasContext) GetCurrentTextureFunc() (fn js.Func[func() GPUTexture]) {
	return fn.FromRef(
		bindings.GPUCanvasContextGetCurrentTextureFunc(
			this.Ref(),
		),
	)
}

// GetCurrentTexture calls the method "GPUCanvasContext.getCurrentTexture".
func (this GPUCanvasContext) GetCurrentTexture() (ret GPUTexture) {
	bindings.CallGPUCanvasContextGetCurrentTexture(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetCurrentTexture calls the method "GPUCanvasContext.getCurrentTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCanvasContext) TryGetCurrentTexture() (ret GPUTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCanvasContextGetCurrentTexture(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
