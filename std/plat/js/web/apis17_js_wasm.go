// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type GPURenderPassEncoder struct {
	ref js.Ref
}

func (this GPURenderPassEncoder) Once() GPURenderPassEncoder {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPURenderPassEncoder.label".
//
// It returns ok=false if there is no such property.
func (this GPURenderPassEncoder) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPURenderPassEncoderLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPURenderPassEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPURenderPassEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPURenderPassEncoderLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncSetViewport returns true if the method "GPURenderPassEncoder.setViewport" exists.
func (this GPURenderPassEncoder) HasFuncSetViewport() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetViewport(
		this.ref,
	)
}

// FuncSetViewport returns the method "GPURenderPassEncoder.setViewport".
func (this GPURenderPassEncoder) FuncSetViewport() (fn js.Func[func(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32)]) {
	bindings.FuncGPURenderPassEncoderSetViewport(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetViewport calls the method "GPURenderPassEncoder.setViewport".
func (this GPURenderPassEncoder) SetViewport(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetViewport(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(x),
		float32(y),
		float32(width),
		float32(height),
		float32(minDepth),
		float32(maxDepth),
	)

	return
}

// HasFuncSetScissorRect returns true if the method "GPURenderPassEncoder.setScissorRect" exists.
func (this GPURenderPassEncoder) HasFuncSetScissorRect() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetScissorRect(
		this.ref,
	)
}

// FuncSetScissorRect returns the method "GPURenderPassEncoder.setScissorRect".
func (this GPURenderPassEncoder) FuncSetScissorRect() (fn js.Func[func(x GPUIntegerCoordinate, y GPUIntegerCoordinate, width GPUIntegerCoordinate, height GPUIntegerCoordinate)]) {
	bindings.FuncGPURenderPassEncoderSetScissorRect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetScissorRect calls the method "GPURenderPassEncoder.setScissorRect".
func (this GPURenderPassEncoder) SetScissorRect(x GPUIntegerCoordinate, y GPUIntegerCoordinate, width GPUIntegerCoordinate, height GPUIntegerCoordinate) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetScissorRect(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(x),
		uint32(y),
		uint32(width),
		uint32(height),
	)

	return
}

// HasFuncSetBlendConstant returns true if the method "GPURenderPassEncoder.setBlendConstant" exists.
func (this GPURenderPassEncoder) HasFuncSetBlendConstant() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetBlendConstant(
		this.ref,
	)
}

// FuncSetBlendConstant returns the method "GPURenderPassEncoder.setBlendConstant".
func (this GPURenderPassEncoder) FuncSetBlendConstant() (fn js.Func[func(color GPUColor)]) {
	bindings.FuncGPURenderPassEncoderSetBlendConstant(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBlendConstant calls the method "GPURenderPassEncoder.setBlendConstant".
func (this GPURenderPassEncoder) SetBlendConstant(color GPUColor) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetBlendConstant(
		this.ref, js.Pointer(&ret),
		color.Ref(),
	)

	return
}

// TrySetBlendConstant calls the method "GPURenderPassEncoder.setBlendConstant"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetBlendConstant(color GPUColor) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetBlendConstant(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		color.Ref(),
	)

	return
}

// HasFuncSetStencilReference returns true if the method "GPURenderPassEncoder.setStencilReference" exists.
func (this GPURenderPassEncoder) HasFuncSetStencilReference() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetStencilReference(
		this.ref,
	)
}

// FuncSetStencilReference returns the method "GPURenderPassEncoder.setStencilReference".
func (this GPURenderPassEncoder) FuncSetStencilReference() (fn js.Func[func(reference GPUStencilValue)]) {
	bindings.FuncGPURenderPassEncoderSetStencilReference(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetStencilReference calls the method "GPURenderPassEncoder.setStencilReference".
func (this GPURenderPassEncoder) SetStencilReference(reference GPUStencilValue) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetStencilReference(
		this.ref, js.Pointer(&ret),
		uint32(reference),
	)

	return
}

// TrySetStencilReference calls the method "GPURenderPassEncoder.setStencilReference"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetStencilReference(reference GPUStencilValue) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetStencilReference(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(reference),
	)

	return
}

// HasFuncBeginOcclusionQuery returns true if the method "GPURenderPassEncoder.beginOcclusionQuery" exists.
func (this GPURenderPassEncoder) HasFuncBeginOcclusionQuery() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderBeginOcclusionQuery(
		this.ref,
	)
}

// FuncBeginOcclusionQuery returns the method "GPURenderPassEncoder.beginOcclusionQuery".
func (this GPURenderPassEncoder) FuncBeginOcclusionQuery() (fn js.Func[func(queryIndex GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderBeginOcclusionQuery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginOcclusionQuery calls the method "GPURenderPassEncoder.beginOcclusionQuery".
func (this GPURenderPassEncoder) BeginOcclusionQuery(queryIndex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderBeginOcclusionQuery(
		this.ref, js.Pointer(&ret),
		uint32(queryIndex),
	)

	return
}

// TryBeginOcclusionQuery calls the method "GPURenderPassEncoder.beginOcclusionQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryBeginOcclusionQuery(queryIndex GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderBeginOcclusionQuery(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(queryIndex),
	)

	return
}

// HasFuncEndOcclusionQuery returns true if the method "GPURenderPassEncoder.endOcclusionQuery" exists.
func (this GPURenderPassEncoder) HasFuncEndOcclusionQuery() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderEndOcclusionQuery(
		this.ref,
	)
}

// FuncEndOcclusionQuery returns the method "GPURenderPassEncoder.endOcclusionQuery".
func (this GPURenderPassEncoder) FuncEndOcclusionQuery() (fn js.Func[func()]) {
	bindings.FuncGPURenderPassEncoderEndOcclusionQuery(
		this.ref, js.Pointer(&fn),
	)
	return
}

// EndOcclusionQuery calls the method "GPURenderPassEncoder.endOcclusionQuery".
func (this GPURenderPassEncoder) EndOcclusionQuery() (ret js.Void) {
	bindings.CallGPURenderPassEncoderEndOcclusionQuery(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEndOcclusionQuery calls the method "GPURenderPassEncoder.endOcclusionQuery"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryEndOcclusionQuery() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderEndOcclusionQuery(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncExecuteBundles returns true if the method "GPURenderPassEncoder.executeBundles" exists.
func (this GPURenderPassEncoder) HasFuncExecuteBundles() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderExecuteBundles(
		this.ref,
	)
}

// FuncExecuteBundles returns the method "GPURenderPassEncoder.executeBundles".
func (this GPURenderPassEncoder) FuncExecuteBundles() (fn js.Func[func(bundles js.Array[GPURenderBundle])]) {
	bindings.FuncGPURenderPassEncoderExecuteBundles(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ExecuteBundles calls the method "GPURenderPassEncoder.executeBundles".
func (this GPURenderPassEncoder) ExecuteBundles(bundles js.Array[GPURenderBundle]) (ret js.Void) {
	bindings.CallGPURenderPassEncoderExecuteBundles(
		this.ref, js.Pointer(&ret),
		bundles.Ref(),
	)

	return
}

// TryExecuteBundles calls the method "GPURenderPassEncoder.executeBundles"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryExecuteBundles(bundles js.Array[GPURenderBundle]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderExecuteBundles(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		bundles.Ref(),
	)

	return
}

// HasFuncEnd returns true if the method "GPURenderPassEncoder.end" exists.
func (this GPURenderPassEncoder) HasFuncEnd() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderEnd(
		this.ref,
	)
}

// FuncEnd returns the method "GPURenderPassEncoder.end".
func (this GPURenderPassEncoder) FuncEnd() (fn js.Func[func()]) {
	bindings.FuncGPURenderPassEncoderEnd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// End calls the method "GPURenderPassEncoder.end".
func (this GPURenderPassEncoder) End() (ret js.Void) {
	bindings.CallGPURenderPassEncoderEnd(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEnd calls the method "GPURenderPassEncoder.end"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryEnd() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderEnd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetPipeline returns true if the method "GPURenderPassEncoder.setPipeline" exists.
func (this GPURenderPassEncoder) HasFuncSetPipeline() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetPipeline(
		this.ref,
	)
}

// FuncSetPipeline returns the method "GPURenderPassEncoder.setPipeline".
func (this GPURenderPassEncoder) FuncSetPipeline() (fn js.Func[func(pipeline GPURenderPipeline)]) {
	bindings.FuncGPURenderPassEncoderSetPipeline(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPipeline calls the method "GPURenderPassEncoder.setPipeline".
func (this GPURenderPassEncoder) SetPipeline(pipeline GPURenderPipeline) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetPipeline(
		this.ref, js.Pointer(&ret),
		pipeline.Ref(),
	)

	return
}

// TrySetPipeline calls the method "GPURenderPassEncoder.setPipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TrySetPipeline(pipeline GPURenderPipeline) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderSetPipeline(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		pipeline.Ref(),
	)

	return
}

// HasFuncSetIndexBuffer returns true if the method "GPURenderPassEncoder.setIndexBuffer" exists.
func (this GPURenderPassEncoder) HasFuncSetIndexBuffer() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetIndexBuffer(
		this.ref,
	)
}

// FuncSetIndexBuffer returns the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) FuncSetIndexBuffer() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64)]) {
	bindings.FuncGPURenderPassEncoderSetIndexBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetIndexBuffer calls the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) SetIndexBuffer(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetIndexBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
		float64(size),
	)

	return
}

// HasFuncSetIndexBuffer1 returns true if the method "GPURenderPassEncoder.setIndexBuffer" exists.
func (this GPURenderPassEncoder) HasFuncSetIndexBuffer1() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetIndexBuffer1(
		this.ref,
	)
}

// FuncSetIndexBuffer1 returns the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) FuncSetIndexBuffer1() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64)]) {
	bindings.FuncGPURenderPassEncoderSetIndexBuffer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetIndexBuffer1 calls the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) SetIndexBuffer1(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetIndexBuffer1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
	)

	return
}

// HasFuncSetIndexBuffer2 returns true if the method "GPURenderPassEncoder.setIndexBuffer" exists.
func (this GPURenderPassEncoder) HasFuncSetIndexBuffer2() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetIndexBuffer2(
		this.ref,
	)
}

// FuncSetIndexBuffer2 returns the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) FuncSetIndexBuffer2() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat)]) {
	bindings.FuncGPURenderPassEncoderSetIndexBuffer2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetIndexBuffer2 calls the method "GPURenderPassEncoder.setIndexBuffer".
func (this GPURenderPassEncoder) SetIndexBuffer2(buffer GPUBuffer, indexFormat GPUIndexFormat) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetIndexBuffer2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
	)

	return
}

// HasFuncSetVertexBuffer returns true if the method "GPURenderPassEncoder.setVertexBuffer" exists.
func (this GPURenderPassEncoder) HasFuncSetVertexBuffer() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetVertexBuffer(
		this.ref,
	)
}

// FuncSetVertexBuffer returns the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) FuncSetVertexBuffer() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64)]) {
	bindings.FuncGPURenderPassEncoderSetVertexBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetVertexBuffer calls the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) SetVertexBuffer(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetVertexBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// HasFuncSetVertexBuffer1 returns true if the method "GPURenderPassEncoder.setVertexBuffer" exists.
func (this GPURenderPassEncoder) HasFuncSetVertexBuffer1() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetVertexBuffer1(
		this.ref,
	)
}

// FuncSetVertexBuffer1 returns the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) FuncSetVertexBuffer1() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64)]) {
	bindings.FuncGPURenderPassEncoderSetVertexBuffer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetVertexBuffer1 calls the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) SetVertexBuffer1(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetVertexBuffer1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
	)

	return
}

// HasFuncSetVertexBuffer2 returns true if the method "GPURenderPassEncoder.setVertexBuffer" exists.
func (this GPURenderPassEncoder) HasFuncSetVertexBuffer2() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetVertexBuffer2(
		this.ref,
	)
}

// FuncSetVertexBuffer2 returns the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) FuncSetVertexBuffer2() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer)]) {
	bindings.FuncGPURenderPassEncoderSetVertexBuffer2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetVertexBuffer2 calls the method "GPURenderPassEncoder.setVertexBuffer".
func (this GPURenderPassEncoder) SetVertexBuffer2(slot GPUIndex32, buffer GPUBuffer) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetVertexBuffer2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
	)

	return
}

// HasFuncDraw returns true if the method "GPURenderPassEncoder.draw" exists.
func (this GPURenderPassEncoder) HasFuncDraw() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDraw(
		this.ref,
	)
}

// FuncDraw returns the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) FuncDraw() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderDraw(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Draw calls the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDraw(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
		uint32(firstInstance),
	)

	return
}

// HasFuncDraw1 returns true if the method "GPURenderPassEncoder.draw" exists.
func (this GPURenderPassEncoder) HasFuncDraw1() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDraw1(
		this.ref,
	)
}

// FuncDraw1 returns the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) FuncDraw1() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderDraw1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Draw1 calls the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw1(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDraw1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
	)

	return
}

// HasFuncDraw2 returns true if the method "GPURenderPassEncoder.draw" exists.
func (this GPURenderPassEncoder) HasFuncDraw2() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDraw2(
		this.ref,
	)
}

// FuncDraw2 returns the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) FuncDraw2() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderDraw2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Draw2 calls the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw2(vertexCount GPUSize32, instanceCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDraw2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
	)

	return
}

// HasFuncDraw3 returns true if the method "GPURenderPassEncoder.draw" exists.
func (this GPURenderPassEncoder) HasFuncDraw3() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDraw3(
		this.ref,
	)
}

// FuncDraw3 returns the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) FuncDraw3() (fn js.Func[func(vertexCount GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderDraw3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Draw3 calls the method "GPURenderPassEncoder.draw".
func (this GPURenderPassEncoder) Draw3(vertexCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDraw3(
		this.ref, js.Pointer(&ret),
		uint32(vertexCount),
	)

	return
}

// TryDraw3 calls the method "GPURenderPassEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDraw3(vertexCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDraw3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
	)

	return
}

// HasFuncDrawIndexed returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasFuncDrawIndexed() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDrawIndexed(
		this.ref,
	)
}

// FuncDrawIndexed returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) FuncDrawIndexed() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderDrawIndexed(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
		uint32(firstInstance),
	)

	return
}

// HasFuncDrawIndexed1 returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasFuncDrawIndexed1() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDrawIndexed1(
		this.ref,
	)
}

// FuncDrawIndexed1 returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) FuncDrawIndexed1() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32)]) {
	bindings.FuncGPURenderPassEncoderDrawIndexed1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed1 calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed1(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
	)

	return
}

// HasFuncDrawIndexed2 returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasFuncDrawIndexed2() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDrawIndexed2(
		this.ref,
	)
}

// FuncDrawIndexed2 returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) FuncDrawIndexed2() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderDrawIndexed2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed2 calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed2(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
	)

	return
}

// HasFuncDrawIndexed3 returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasFuncDrawIndexed3() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDrawIndexed3(
		this.ref,
	)
}

// FuncDrawIndexed3 returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) FuncDrawIndexed3() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderDrawIndexed3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed3 calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed3(indexCount GPUSize32, instanceCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
	)

	return
}

// HasFuncDrawIndexed4 returns true if the method "GPURenderPassEncoder.drawIndexed" exists.
func (this GPURenderPassEncoder) HasFuncDrawIndexed4() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDrawIndexed4(
		this.ref,
	)
}

// FuncDrawIndexed4 returns the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) FuncDrawIndexed4() (fn js.Func[func(indexCount GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderDrawIndexed4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed4 calls the method "GPURenderPassEncoder.drawIndexed".
func (this GPURenderPassEncoder) DrawIndexed4(indexCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexed4(
		this.ref, js.Pointer(&ret),
		uint32(indexCount),
	)

	return
}

// TryDrawIndexed4 calls the method "GPURenderPassEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryDrawIndexed4(indexCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderDrawIndexed4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
	)

	return
}

// HasFuncDrawIndirect returns true if the method "GPURenderPassEncoder.drawIndirect" exists.
func (this GPURenderPassEncoder) HasFuncDrawIndirect() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDrawIndirect(
		this.ref,
	)
}

// FuncDrawIndirect returns the method "GPURenderPassEncoder.drawIndirect".
func (this GPURenderPassEncoder) FuncDrawIndirect() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	bindings.FuncGPURenderPassEncoderDrawIndirect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndirect calls the method "GPURenderPassEncoder.drawIndirect".
func (this GPURenderPassEncoder) DrawIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndirect(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasFuncDrawIndexedIndirect returns true if the method "GPURenderPassEncoder.drawIndexedIndirect" exists.
func (this GPURenderPassEncoder) HasFuncDrawIndexedIndirect() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderDrawIndexedIndirect(
		this.ref,
	)
}

// FuncDrawIndexedIndirect returns the method "GPURenderPassEncoder.drawIndexedIndirect".
func (this GPURenderPassEncoder) FuncDrawIndexedIndirect() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	bindings.FuncGPURenderPassEncoderDrawIndexedIndirect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexedIndirect calls the method "GPURenderPassEncoder.drawIndexedIndirect".
func (this GPURenderPassEncoder) DrawIndexedIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderPassEncoderDrawIndexedIndirect(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasFuncSetBindGroup returns true if the method "GPURenderPassEncoder.setBindGroup" exists.
func (this GPURenderPassEncoder) HasFuncSetBindGroup() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetBindGroup(
		this.ref,
	)
}

// FuncSetBindGroup returns the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) FuncSetBindGroup() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset])]) {
	bindings.FuncGPURenderPassEncoderSetBindGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBindGroup calls the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetBindGroup(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	return
}

// HasFuncSetBindGroup1 returns true if the method "GPURenderPassEncoder.setBindGroup" exists.
func (this GPURenderPassEncoder) HasFuncSetBindGroup1() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetBindGroup1(
		this.ref,
	)
}

// FuncSetBindGroup1 returns the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) FuncSetBindGroup1() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup)]) {
	bindings.FuncGPURenderPassEncoderSetBindGroup1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBindGroup1 calls the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) SetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetBindGroup1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
	)

	return
}

// HasFuncSetBindGroup2 returns true if the method "GPURenderPassEncoder.setBindGroup" exists.
func (this GPURenderPassEncoder) HasFuncSetBindGroup2() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderSetBindGroup2(
		this.ref,
	)
}

// FuncSetBindGroup2 returns the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) FuncSetBindGroup2() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32)]) {
	bindings.FuncGPURenderPassEncoderSetBindGroup2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBindGroup2 calls the method "GPURenderPassEncoder.setBindGroup".
func (this GPURenderPassEncoder) SetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (ret js.Void) {
	bindings.CallGPURenderPassEncoderSetBindGroup2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	return
}

// HasFuncPushDebugGroup returns true if the method "GPURenderPassEncoder.pushDebugGroup" exists.
func (this GPURenderPassEncoder) HasFuncPushDebugGroup() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderPushDebugGroup(
		this.ref,
	)
}

// FuncPushDebugGroup returns the method "GPURenderPassEncoder.pushDebugGroup".
func (this GPURenderPassEncoder) FuncPushDebugGroup() (fn js.Func[func(groupLabel js.String)]) {
	bindings.FuncGPURenderPassEncoderPushDebugGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PushDebugGroup calls the method "GPURenderPassEncoder.pushDebugGroup".
func (this GPURenderPassEncoder) PushDebugGroup(groupLabel js.String) (ret js.Void) {
	bindings.CallGPURenderPassEncoderPushDebugGroup(
		this.ref, js.Pointer(&ret),
		groupLabel.Ref(),
	)

	return
}

// TryPushDebugGroup calls the method "GPURenderPassEncoder.pushDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryPushDebugGroup(groupLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderPushDebugGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		groupLabel.Ref(),
	)

	return
}

// HasFuncPopDebugGroup returns true if the method "GPURenderPassEncoder.popDebugGroup" exists.
func (this GPURenderPassEncoder) HasFuncPopDebugGroup() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderPopDebugGroup(
		this.ref,
	)
}

// FuncPopDebugGroup returns the method "GPURenderPassEncoder.popDebugGroup".
func (this GPURenderPassEncoder) FuncPopDebugGroup() (fn js.Func[func()]) {
	bindings.FuncGPURenderPassEncoderPopDebugGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PopDebugGroup calls the method "GPURenderPassEncoder.popDebugGroup".
func (this GPURenderPassEncoder) PopDebugGroup() (ret js.Void) {
	bindings.CallGPURenderPassEncoderPopDebugGroup(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPopDebugGroup calls the method "GPURenderPassEncoder.popDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryPopDebugGroup() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderPopDebugGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInsertDebugMarker returns true if the method "GPURenderPassEncoder.insertDebugMarker" exists.
func (this GPURenderPassEncoder) HasFuncInsertDebugMarker() bool {
	return js.True == bindings.HasFuncGPURenderPassEncoderInsertDebugMarker(
		this.ref,
	)
}

// FuncInsertDebugMarker returns the method "GPURenderPassEncoder.insertDebugMarker".
func (this GPURenderPassEncoder) FuncInsertDebugMarker() (fn js.Func[func(markerLabel js.String)]) {
	bindings.FuncGPURenderPassEncoderInsertDebugMarker(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertDebugMarker calls the method "GPURenderPassEncoder.insertDebugMarker".
func (this GPURenderPassEncoder) InsertDebugMarker(markerLabel js.String) (ret js.Void) {
	bindings.CallGPURenderPassEncoderInsertDebugMarker(
		this.ref, js.Pointer(&ret),
		markerLabel.Ref(),
	)

	return
}

// TryInsertDebugMarker calls the method "GPURenderPassEncoder.insertDebugMarker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderPassEncoder) TryInsertDebugMarker(markerLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderPassEncoderInsertDebugMarker(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *GPURenderPassColorAttachment) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassColorAttachmentJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPURenderPassColorAttachment) Update(ref js.Ref) {
	bindings.GPURenderPassColorAttachmentJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPURenderPassColorAttachment) FreeMembers(recursive bool) {
	js.Free(
		p.View.Ref(),
		p.ResolveTarget.Ref(),
		p.ClearValue.Ref(),
	)
	p.View = p.View.FromRef(js.Undefined)
	p.ResolveTarget = p.ResolveTarget.FromRef(js.Undefined)
	p.ClearValue = p.ClearValue.FromRef(js.Undefined)
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
func (p *GPURenderPassDepthStencilAttachment) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassDepthStencilAttachmentJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPURenderPassDepthStencilAttachment) Update(ref js.Ref) {
	bindings.GPURenderPassDepthStencilAttachmentJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPURenderPassDepthStencilAttachment) FreeMembers(recursive bool) {
	js.Free(
		p.View.Ref(),
	)
	p.View = p.View.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "GPUQuerySet.type".
//
// It returns ok=false if there is no such property.
func (this GPUQuerySet) Type() (ret GPUQueryType, ok bool) {
	ok = js.True == bindings.GetGPUQuerySetType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Count returns the value of property "GPUQuerySet.count".
//
// It returns ok=false if there is no such property.
func (this GPUQuerySet) Count() (ret GPUSize32Out, ok bool) {
	ok = js.True == bindings.GetGPUQuerySetCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "GPUQuerySet.label".
//
// It returns ok=false if there is no such property.
func (this GPUQuerySet) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUQuerySetLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUQuerySet.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUQuerySet) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUQuerySetLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncDestroy returns true if the method "GPUQuerySet.destroy" exists.
func (this GPUQuerySet) HasFuncDestroy() bool {
	return js.True == bindings.HasFuncGPUQuerySetDestroy(
		this.ref,
	)
}

// FuncDestroy returns the method "GPUQuerySet.destroy".
func (this GPUQuerySet) FuncDestroy() (fn js.Func[func()]) {
	bindings.FuncGPUQuerySetDestroy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Destroy calls the method "GPUQuerySet.destroy".
func (this GPUQuerySet) Destroy() (ret js.Void) {
	bindings.CallGPUQuerySetDestroy(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDestroy calls the method "GPUQuerySet.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQuerySet) TryDestroy() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQuerySetDestroy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *GPURenderPassTimestampWrites) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassTimestampWritesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPURenderPassTimestampWrites) Update(ref js.Ref) {
	bindings.GPURenderPassTimestampWritesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPURenderPassTimestampWrites) FreeMembers(recursive bool) {
	js.Free(
		p.QuerySet.Ref(),
	)
	p.QuerySet = p.QuerySet.FromRef(js.Undefined)
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
func (p *GPURenderPassDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPURenderPassDescriptor) Update(ref js.Ref) {
	bindings.GPURenderPassDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPURenderPassDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.ColorAttachments.Ref(),
		p.OcclusionQuerySet.Ref(),
		p.Label.Ref(),
	)
	p.ColorAttachments = p.ColorAttachments.FromRef(js.Undefined)
	p.OcclusionQuerySet = p.OcclusionQuerySet.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
	if recursive {
		p.DepthStencilAttachment.FreeMembers(true)
		p.TimestampWrites.FreeMembers(true)
	}
}

type GPUComputePassEncoder struct {
	ref js.Ref
}

func (this GPUComputePassEncoder) Once() GPUComputePassEncoder {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUComputePassEncoder.label".
//
// It returns ok=false if there is no such property.
func (this GPUComputePassEncoder) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUComputePassEncoderLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUComputePassEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUComputePassEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUComputePassEncoderLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncSetPipeline returns true if the method "GPUComputePassEncoder.setPipeline" exists.
func (this GPUComputePassEncoder) HasFuncSetPipeline() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderSetPipeline(
		this.ref,
	)
}

// FuncSetPipeline returns the method "GPUComputePassEncoder.setPipeline".
func (this GPUComputePassEncoder) FuncSetPipeline() (fn js.Func[func(pipeline GPUComputePipeline)]) {
	bindings.FuncGPUComputePassEncoderSetPipeline(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPipeline calls the method "GPUComputePassEncoder.setPipeline".
func (this GPUComputePassEncoder) SetPipeline(pipeline GPUComputePipeline) (ret js.Void) {
	bindings.CallGPUComputePassEncoderSetPipeline(
		this.ref, js.Pointer(&ret),
		pipeline.Ref(),
	)

	return
}

// TrySetPipeline calls the method "GPUComputePassEncoder.setPipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TrySetPipeline(pipeline GPUComputePipeline) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderSetPipeline(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		pipeline.Ref(),
	)

	return
}

// HasFuncDispatchWorkgroups returns true if the method "GPUComputePassEncoder.dispatchWorkgroups" exists.
func (this GPUComputePassEncoder) HasFuncDispatchWorkgroups() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderDispatchWorkgroups(
		this.ref,
	)
}

// FuncDispatchWorkgroups returns the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) FuncDispatchWorkgroups() (fn js.Func[func(workgroupCountX GPUSize32, workgroupCountY GPUSize32, workgroupCountZ GPUSize32)]) {
	bindings.FuncGPUComputePassEncoderDispatchWorkgroups(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DispatchWorkgroups calls the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) DispatchWorkgroups(workgroupCountX GPUSize32, workgroupCountY GPUSize32, workgroupCountZ GPUSize32) (ret js.Void) {
	bindings.CallGPUComputePassEncoderDispatchWorkgroups(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(workgroupCountX),
		uint32(workgroupCountY),
		uint32(workgroupCountZ),
	)

	return
}

// HasFuncDispatchWorkgroups1 returns true if the method "GPUComputePassEncoder.dispatchWorkgroups" exists.
func (this GPUComputePassEncoder) HasFuncDispatchWorkgroups1() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderDispatchWorkgroups1(
		this.ref,
	)
}

// FuncDispatchWorkgroups1 returns the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) FuncDispatchWorkgroups1() (fn js.Func[func(workgroupCountX GPUSize32, workgroupCountY GPUSize32)]) {
	bindings.FuncGPUComputePassEncoderDispatchWorkgroups1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DispatchWorkgroups1 calls the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) DispatchWorkgroups1(workgroupCountX GPUSize32, workgroupCountY GPUSize32) (ret js.Void) {
	bindings.CallGPUComputePassEncoderDispatchWorkgroups1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(workgroupCountX),
		uint32(workgroupCountY),
	)

	return
}

// HasFuncDispatchWorkgroups2 returns true if the method "GPUComputePassEncoder.dispatchWorkgroups" exists.
func (this GPUComputePassEncoder) HasFuncDispatchWorkgroups2() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderDispatchWorkgroups2(
		this.ref,
	)
}

// FuncDispatchWorkgroups2 returns the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) FuncDispatchWorkgroups2() (fn js.Func[func(workgroupCountX GPUSize32)]) {
	bindings.FuncGPUComputePassEncoderDispatchWorkgroups2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DispatchWorkgroups2 calls the method "GPUComputePassEncoder.dispatchWorkgroups".
func (this GPUComputePassEncoder) DispatchWorkgroups2(workgroupCountX GPUSize32) (ret js.Void) {
	bindings.CallGPUComputePassEncoderDispatchWorkgroups2(
		this.ref, js.Pointer(&ret),
		uint32(workgroupCountX),
	)

	return
}

// TryDispatchWorkgroups2 calls the method "GPUComputePassEncoder.dispatchWorkgroups"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryDispatchWorkgroups2(workgroupCountX GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderDispatchWorkgroups2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(workgroupCountX),
	)

	return
}

// HasFuncDispatchWorkgroupsIndirect returns true if the method "GPUComputePassEncoder.dispatchWorkgroupsIndirect" exists.
func (this GPUComputePassEncoder) HasFuncDispatchWorkgroupsIndirect() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderDispatchWorkgroupsIndirect(
		this.ref,
	)
}

// FuncDispatchWorkgroupsIndirect returns the method "GPUComputePassEncoder.dispatchWorkgroupsIndirect".
func (this GPUComputePassEncoder) FuncDispatchWorkgroupsIndirect() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	bindings.FuncGPUComputePassEncoderDispatchWorkgroupsIndirect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DispatchWorkgroupsIndirect calls the method "GPUComputePassEncoder.dispatchWorkgroupsIndirect".
func (this GPUComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPUComputePassEncoderDispatchWorkgroupsIndirect(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasFuncEnd returns true if the method "GPUComputePassEncoder.end" exists.
func (this GPUComputePassEncoder) HasFuncEnd() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderEnd(
		this.ref,
	)
}

// FuncEnd returns the method "GPUComputePassEncoder.end".
func (this GPUComputePassEncoder) FuncEnd() (fn js.Func[func()]) {
	bindings.FuncGPUComputePassEncoderEnd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// End calls the method "GPUComputePassEncoder.end".
func (this GPUComputePassEncoder) End() (ret js.Void) {
	bindings.CallGPUComputePassEncoderEnd(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEnd calls the method "GPUComputePassEncoder.end"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryEnd() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderEnd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPushDebugGroup returns true if the method "GPUComputePassEncoder.pushDebugGroup" exists.
func (this GPUComputePassEncoder) HasFuncPushDebugGroup() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderPushDebugGroup(
		this.ref,
	)
}

// FuncPushDebugGroup returns the method "GPUComputePassEncoder.pushDebugGroup".
func (this GPUComputePassEncoder) FuncPushDebugGroup() (fn js.Func[func(groupLabel js.String)]) {
	bindings.FuncGPUComputePassEncoderPushDebugGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PushDebugGroup calls the method "GPUComputePassEncoder.pushDebugGroup".
func (this GPUComputePassEncoder) PushDebugGroup(groupLabel js.String) (ret js.Void) {
	bindings.CallGPUComputePassEncoderPushDebugGroup(
		this.ref, js.Pointer(&ret),
		groupLabel.Ref(),
	)

	return
}

// TryPushDebugGroup calls the method "GPUComputePassEncoder.pushDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryPushDebugGroup(groupLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderPushDebugGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		groupLabel.Ref(),
	)

	return
}

// HasFuncPopDebugGroup returns true if the method "GPUComputePassEncoder.popDebugGroup" exists.
func (this GPUComputePassEncoder) HasFuncPopDebugGroup() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderPopDebugGroup(
		this.ref,
	)
}

// FuncPopDebugGroup returns the method "GPUComputePassEncoder.popDebugGroup".
func (this GPUComputePassEncoder) FuncPopDebugGroup() (fn js.Func[func()]) {
	bindings.FuncGPUComputePassEncoderPopDebugGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PopDebugGroup calls the method "GPUComputePassEncoder.popDebugGroup".
func (this GPUComputePassEncoder) PopDebugGroup() (ret js.Void) {
	bindings.CallGPUComputePassEncoderPopDebugGroup(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPopDebugGroup calls the method "GPUComputePassEncoder.popDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryPopDebugGroup() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderPopDebugGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInsertDebugMarker returns true if the method "GPUComputePassEncoder.insertDebugMarker" exists.
func (this GPUComputePassEncoder) HasFuncInsertDebugMarker() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderInsertDebugMarker(
		this.ref,
	)
}

// FuncInsertDebugMarker returns the method "GPUComputePassEncoder.insertDebugMarker".
func (this GPUComputePassEncoder) FuncInsertDebugMarker() (fn js.Func[func(markerLabel js.String)]) {
	bindings.FuncGPUComputePassEncoderInsertDebugMarker(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertDebugMarker calls the method "GPUComputePassEncoder.insertDebugMarker".
func (this GPUComputePassEncoder) InsertDebugMarker(markerLabel js.String) (ret js.Void) {
	bindings.CallGPUComputePassEncoderInsertDebugMarker(
		this.ref, js.Pointer(&ret),
		markerLabel.Ref(),
	)

	return
}

// TryInsertDebugMarker calls the method "GPUComputePassEncoder.insertDebugMarker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUComputePassEncoder) TryInsertDebugMarker(markerLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUComputePassEncoderInsertDebugMarker(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		markerLabel.Ref(),
	)

	return
}

// HasFuncSetBindGroup returns true if the method "GPUComputePassEncoder.setBindGroup" exists.
func (this GPUComputePassEncoder) HasFuncSetBindGroup() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderSetBindGroup(
		this.ref,
	)
}

// FuncSetBindGroup returns the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) FuncSetBindGroup() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset])]) {
	bindings.FuncGPUComputePassEncoderSetBindGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBindGroup calls the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (ret js.Void) {
	bindings.CallGPUComputePassEncoderSetBindGroup(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	return
}

// HasFuncSetBindGroup1 returns true if the method "GPUComputePassEncoder.setBindGroup" exists.
func (this GPUComputePassEncoder) HasFuncSetBindGroup1() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderSetBindGroup1(
		this.ref,
	)
}

// FuncSetBindGroup1 returns the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) FuncSetBindGroup1() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup)]) {
	bindings.FuncGPUComputePassEncoderSetBindGroup1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBindGroup1 calls the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) SetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (ret js.Void) {
	bindings.CallGPUComputePassEncoderSetBindGroup1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
	)

	return
}

// HasFuncSetBindGroup2 returns true if the method "GPUComputePassEncoder.setBindGroup" exists.
func (this GPUComputePassEncoder) HasFuncSetBindGroup2() bool {
	return js.True == bindings.HasFuncGPUComputePassEncoderSetBindGroup2(
		this.ref,
	)
}

// FuncSetBindGroup2 returns the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) FuncSetBindGroup2() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32)]) {
	bindings.FuncGPUComputePassEncoderSetBindGroup2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBindGroup2 calls the method "GPUComputePassEncoder.setBindGroup".
func (this GPUComputePassEncoder) SetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (ret js.Void) {
	bindings.CallGPUComputePassEncoderSetBindGroup2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *GPUComputePassTimestampWrites) UpdateFrom(ref js.Ref) {
	bindings.GPUComputePassTimestampWritesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUComputePassTimestampWrites) Update(ref js.Ref) {
	bindings.GPUComputePassTimestampWritesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUComputePassTimestampWrites) FreeMembers(recursive bool) {
	js.Free(
		p.QuerySet.Ref(),
	)
	p.QuerySet = p.QuerySet.FromRef(js.Undefined)
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
func (p *GPUComputePassDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUComputePassDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUComputePassDescriptor) Update(ref js.Ref) {
	bindings.GPUComputePassDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUComputePassDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
	if recursive {
		p.TimestampWrites.FreeMembers(true)
	}
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
func (p *GPUImageCopyBuffer) UpdateFrom(ref js.Ref) {
	bindings.GPUImageCopyBufferJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUImageCopyBuffer) Update(ref js.Ref) {
	bindings.GPUImageCopyBufferJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUImageCopyBuffer) FreeMembers(recursive bool) {
	js.Free(
		p.Buffer.Ref(),
	)
	p.Buffer = p.Buffer.FromRef(js.Undefined)
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
func (p *GPUOrigin3DDict) UpdateFrom(ref js.Ref) {
	bindings.GPUOrigin3DDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUOrigin3DDict) Update(ref js.Ref) {
	bindings.GPUOrigin3DDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUOrigin3DDict) FreeMembers(recursive bool) {
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
func (p *GPUImageCopyTexture) UpdateFrom(ref js.Ref) {
	bindings.GPUImageCopyTextureJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUImageCopyTexture) Update(ref js.Ref) {
	bindings.GPUImageCopyTextureJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUImageCopyTexture) FreeMembers(recursive bool) {
	js.Free(
		p.Texture.Ref(),
		p.Origin.Ref(),
	)
	p.Texture = p.Texture.FromRef(js.Undefined)
	p.Origin = p.Origin.FromRef(js.Undefined)
}

type GPUCommandBuffer struct {
	ref js.Ref
}

func (this GPUCommandBuffer) Once() GPUCommandBuffer {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUCommandBuffer.label".
//
// It returns ok=false if there is no such property.
func (this GPUCommandBuffer) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUCommandBufferLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUCommandBuffer.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUCommandBuffer) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUCommandBufferLabel(
		this.ref,
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
func (p *GPUCommandBufferDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUCommandBufferDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUCommandBufferDescriptor) Update(ref js.Ref) {
	bindings.GPUCommandBufferDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUCommandBufferDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
}

type GPUCommandEncoder struct {
	ref js.Ref
}

func (this GPUCommandEncoder) Once() GPUCommandEncoder {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUCommandEncoder.label".
//
// It returns ok=false if there is no such property.
func (this GPUCommandEncoder) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUCommandEncoderLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUCommandEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUCommandEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUCommandEncoderLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncBeginRenderPass returns true if the method "GPUCommandEncoder.beginRenderPass" exists.
func (this GPUCommandEncoder) HasFuncBeginRenderPass() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderBeginRenderPass(
		this.ref,
	)
}

// FuncBeginRenderPass returns the method "GPUCommandEncoder.beginRenderPass".
func (this GPUCommandEncoder) FuncBeginRenderPass() (fn js.Func[func(descriptor GPURenderPassDescriptor) GPURenderPassEncoder]) {
	bindings.FuncGPUCommandEncoderBeginRenderPass(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginRenderPass calls the method "GPUCommandEncoder.beginRenderPass".
func (this GPUCommandEncoder) BeginRenderPass(descriptor GPURenderPassDescriptor) (ret GPURenderPassEncoder) {
	bindings.CallGPUCommandEncoderBeginRenderPass(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryBeginRenderPass calls the method "GPUCommandEncoder.beginRenderPass"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryBeginRenderPass(descriptor GPURenderPassDescriptor) (ret GPURenderPassEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderBeginRenderPass(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncBeginComputePass returns true if the method "GPUCommandEncoder.beginComputePass" exists.
func (this GPUCommandEncoder) HasFuncBeginComputePass() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderBeginComputePass(
		this.ref,
	)
}

// FuncBeginComputePass returns the method "GPUCommandEncoder.beginComputePass".
func (this GPUCommandEncoder) FuncBeginComputePass() (fn js.Func[func(descriptor GPUComputePassDescriptor) GPUComputePassEncoder]) {
	bindings.FuncGPUCommandEncoderBeginComputePass(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginComputePass calls the method "GPUCommandEncoder.beginComputePass".
func (this GPUCommandEncoder) BeginComputePass(descriptor GPUComputePassDescriptor) (ret GPUComputePassEncoder) {
	bindings.CallGPUCommandEncoderBeginComputePass(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryBeginComputePass calls the method "GPUCommandEncoder.beginComputePass"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryBeginComputePass(descriptor GPUComputePassDescriptor) (ret GPUComputePassEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderBeginComputePass(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncBeginComputePass1 returns true if the method "GPUCommandEncoder.beginComputePass" exists.
func (this GPUCommandEncoder) HasFuncBeginComputePass1() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderBeginComputePass1(
		this.ref,
	)
}

// FuncBeginComputePass1 returns the method "GPUCommandEncoder.beginComputePass".
func (this GPUCommandEncoder) FuncBeginComputePass1() (fn js.Func[func() GPUComputePassEncoder]) {
	bindings.FuncGPUCommandEncoderBeginComputePass1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// BeginComputePass1 calls the method "GPUCommandEncoder.beginComputePass".
func (this GPUCommandEncoder) BeginComputePass1() (ret GPUComputePassEncoder) {
	bindings.CallGPUCommandEncoderBeginComputePass1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBeginComputePass1 calls the method "GPUCommandEncoder.beginComputePass"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryBeginComputePass1() (ret GPUComputePassEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderBeginComputePass1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCopyBufferToBuffer returns true if the method "GPUCommandEncoder.copyBufferToBuffer" exists.
func (this GPUCommandEncoder) HasFuncCopyBufferToBuffer() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderCopyBufferToBuffer(
		this.ref,
	)
}

// FuncCopyBufferToBuffer returns the method "GPUCommandEncoder.copyBufferToBuffer".
func (this GPUCommandEncoder) FuncCopyBufferToBuffer() (fn js.Func[func(source GPUBuffer, sourceOffset GPUSize64, destination GPUBuffer, destinationOffset GPUSize64, size GPUSize64)]) {
	bindings.FuncGPUCommandEncoderCopyBufferToBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyBufferToBuffer calls the method "GPUCommandEncoder.copyBufferToBuffer".
func (this GPUCommandEncoder) CopyBufferToBuffer(source GPUBuffer, sourceOffset GPUSize64, destination GPUBuffer, destinationOffset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPUCommandEncoderCopyBufferToBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		source.Ref(),
		float64(sourceOffset),
		destination.Ref(),
		float64(destinationOffset),
		float64(size),
	)

	return
}

// HasFuncCopyBufferToTexture returns true if the method "GPUCommandEncoder.copyBufferToTexture" exists.
func (this GPUCommandEncoder) HasFuncCopyBufferToTexture() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderCopyBufferToTexture(
		this.ref,
	)
}

// FuncCopyBufferToTexture returns the method "GPUCommandEncoder.copyBufferToTexture".
func (this GPUCommandEncoder) FuncCopyBufferToTexture() (fn js.Func[func(source GPUImageCopyBuffer, destination GPUImageCopyTexture, copySize GPUExtent3D)]) {
	bindings.FuncGPUCommandEncoderCopyBufferToTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyBufferToTexture calls the method "GPUCommandEncoder.copyBufferToTexture".
func (this GPUCommandEncoder) CopyBufferToTexture(source GPUImageCopyBuffer, destination GPUImageCopyTexture, copySize GPUExtent3D) (ret js.Void) {
	bindings.CallGPUCommandEncoderCopyBufferToTexture(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// HasFuncCopyTextureToBuffer returns true if the method "GPUCommandEncoder.copyTextureToBuffer" exists.
func (this GPUCommandEncoder) HasFuncCopyTextureToBuffer() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderCopyTextureToBuffer(
		this.ref,
	)
}

// FuncCopyTextureToBuffer returns the method "GPUCommandEncoder.copyTextureToBuffer".
func (this GPUCommandEncoder) FuncCopyTextureToBuffer() (fn js.Func[func(source GPUImageCopyTexture, destination GPUImageCopyBuffer, copySize GPUExtent3D)]) {
	bindings.FuncGPUCommandEncoderCopyTextureToBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTextureToBuffer calls the method "GPUCommandEncoder.copyTextureToBuffer".
func (this GPUCommandEncoder) CopyTextureToBuffer(source GPUImageCopyTexture, destination GPUImageCopyBuffer, copySize GPUExtent3D) (ret js.Void) {
	bindings.CallGPUCommandEncoderCopyTextureToBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// HasFuncCopyTextureToTexture returns true if the method "GPUCommandEncoder.copyTextureToTexture" exists.
func (this GPUCommandEncoder) HasFuncCopyTextureToTexture() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderCopyTextureToTexture(
		this.ref,
	)
}

// FuncCopyTextureToTexture returns the method "GPUCommandEncoder.copyTextureToTexture".
func (this GPUCommandEncoder) FuncCopyTextureToTexture() (fn js.Func[func(source GPUImageCopyTexture, destination GPUImageCopyTexture, copySize GPUExtent3D)]) {
	bindings.FuncGPUCommandEncoderCopyTextureToTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyTextureToTexture calls the method "GPUCommandEncoder.copyTextureToTexture".
func (this GPUCommandEncoder) CopyTextureToTexture(source GPUImageCopyTexture, destination GPUImageCopyTexture, copySize GPUExtent3D) (ret js.Void) {
	bindings.CallGPUCommandEncoderCopyTextureToTexture(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	return
}

// HasFuncClearBuffer returns true if the method "GPUCommandEncoder.clearBuffer" exists.
func (this GPUCommandEncoder) HasFuncClearBuffer() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderClearBuffer(
		this.ref,
	)
}

// FuncClearBuffer returns the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) FuncClearBuffer() (fn js.Func[func(buffer GPUBuffer, offset GPUSize64, size GPUSize64)]) {
	bindings.FuncGPUCommandEncoderClearBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBuffer calls the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) ClearBuffer(buffer GPUBuffer, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPUCommandEncoderClearBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// HasFuncClearBuffer1 returns true if the method "GPUCommandEncoder.clearBuffer" exists.
func (this GPUCommandEncoder) HasFuncClearBuffer1() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderClearBuffer1(
		this.ref,
	)
}

// FuncClearBuffer1 returns the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) FuncClearBuffer1() (fn js.Func[func(buffer GPUBuffer, offset GPUSize64)]) {
	bindings.FuncGPUCommandEncoderClearBuffer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBuffer1 calls the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) ClearBuffer1(buffer GPUBuffer, offset GPUSize64) (ret js.Void) {
	bindings.CallGPUCommandEncoderClearBuffer1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(offset),
	)

	return
}

// HasFuncClearBuffer2 returns true if the method "GPUCommandEncoder.clearBuffer" exists.
func (this GPUCommandEncoder) HasFuncClearBuffer2() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderClearBuffer2(
		this.ref,
	)
}

// FuncClearBuffer2 returns the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) FuncClearBuffer2() (fn js.Func[func(buffer GPUBuffer)]) {
	bindings.FuncGPUCommandEncoderClearBuffer2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ClearBuffer2 calls the method "GPUCommandEncoder.clearBuffer".
func (this GPUCommandEncoder) ClearBuffer2(buffer GPUBuffer) (ret js.Void) {
	bindings.CallGPUCommandEncoderClearBuffer2(
		this.ref, js.Pointer(&ret),
		buffer.Ref(),
	)

	return
}

// TryClearBuffer2 calls the method "GPUCommandEncoder.clearBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryClearBuffer2(buffer GPUBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderClearBuffer2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
	)

	return
}

// HasFuncWriteTimestamp returns true if the method "GPUCommandEncoder.writeTimestamp" exists.
func (this GPUCommandEncoder) HasFuncWriteTimestamp() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderWriteTimestamp(
		this.ref,
	)
}

// FuncWriteTimestamp returns the method "GPUCommandEncoder.writeTimestamp".
func (this GPUCommandEncoder) FuncWriteTimestamp() (fn js.Func[func(querySet GPUQuerySet, queryIndex GPUSize32)]) {
	bindings.FuncGPUCommandEncoderWriteTimestamp(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteTimestamp calls the method "GPUCommandEncoder.writeTimestamp".
func (this GPUCommandEncoder) WriteTimestamp(querySet GPUQuerySet, queryIndex GPUSize32) (ret js.Void) {
	bindings.CallGPUCommandEncoderWriteTimestamp(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		querySet.Ref(),
		uint32(queryIndex),
	)

	return
}

// HasFuncResolveQuerySet returns true if the method "GPUCommandEncoder.resolveQuerySet" exists.
func (this GPUCommandEncoder) HasFuncResolveQuerySet() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderResolveQuerySet(
		this.ref,
	)
}

// FuncResolveQuerySet returns the method "GPUCommandEncoder.resolveQuerySet".
func (this GPUCommandEncoder) FuncResolveQuerySet() (fn js.Func[func(querySet GPUQuerySet, firstQuery GPUSize32, queryCount GPUSize32, destination GPUBuffer, destinationOffset GPUSize64)]) {
	bindings.FuncGPUCommandEncoderResolveQuerySet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ResolveQuerySet calls the method "GPUCommandEncoder.resolveQuerySet".
func (this GPUCommandEncoder) ResolveQuerySet(querySet GPUQuerySet, firstQuery GPUSize32, queryCount GPUSize32, destination GPUBuffer, destinationOffset GPUSize64) (ret js.Void) {
	bindings.CallGPUCommandEncoderResolveQuerySet(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		querySet.Ref(),
		uint32(firstQuery),
		uint32(queryCount),
		destination.Ref(),
		float64(destinationOffset),
	)

	return
}

// HasFuncFinish returns true if the method "GPUCommandEncoder.finish" exists.
func (this GPUCommandEncoder) HasFuncFinish() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderFinish(
		this.ref,
	)
}

// FuncFinish returns the method "GPUCommandEncoder.finish".
func (this GPUCommandEncoder) FuncFinish() (fn js.Func[func(descriptor GPUCommandBufferDescriptor) GPUCommandBuffer]) {
	bindings.FuncGPUCommandEncoderFinish(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Finish calls the method "GPUCommandEncoder.finish".
func (this GPUCommandEncoder) Finish(descriptor GPUCommandBufferDescriptor) (ret GPUCommandBuffer) {
	bindings.CallGPUCommandEncoderFinish(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryFinish calls the method "GPUCommandEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryFinish(descriptor GPUCommandBufferDescriptor) (ret GPUCommandBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderFinish(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncFinish1 returns true if the method "GPUCommandEncoder.finish" exists.
func (this GPUCommandEncoder) HasFuncFinish1() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderFinish1(
		this.ref,
	)
}

// FuncFinish1 returns the method "GPUCommandEncoder.finish".
func (this GPUCommandEncoder) FuncFinish1() (fn js.Func[func() GPUCommandBuffer]) {
	bindings.FuncGPUCommandEncoderFinish1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Finish1 calls the method "GPUCommandEncoder.finish".
func (this GPUCommandEncoder) Finish1() (ret GPUCommandBuffer) {
	bindings.CallGPUCommandEncoderFinish1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFinish1 calls the method "GPUCommandEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryFinish1() (ret GPUCommandBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderFinish1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPushDebugGroup returns true if the method "GPUCommandEncoder.pushDebugGroup" exists.
func (this GPUCommandEncoder) HasFuncPushDebugGroup() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderPushDebugGroup(
		this.ref,
	)
}

// FuncPushDebugGroup returns the method "GPUCommandEncoder.pushDebugGroup".
func (this GPUCommandEncoder) FuncPushDebugGroup() (fn js.Func[func(groupLabel js.String)]) {
	bindings.FuncGPUCommandEncoderPushDebugGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PushDebugGroup calls the method "GPUCommandEncoder.pushDebugGroup".
func (this GPUCommandEncoder) PushDebugGroup(groupLabel js.String) (ret js.Void) {
	bindings.CallGPUCommandEncoderPushDebugGroup(
		this.ref, js.Pointer(&ret),
		groupLabel.Ref(),
	)

	return
}

// TryPushDebugGroup calls the method "GPUCommandEncoder.pushDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryPushDebugGroup(groupLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderPushDebugGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		groupLabel.Ref(),
	)

	return
}

// HasFuncPopDebugGroup returns true if the method "GPUCommandEncoder.popDebugGroup" exists.
func (this GPUCommandEncoder) HasFuncPopDebugGroup() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderPopDebugGroup(
		this.ref,
	)
}

// FuncPopDebugGroup returns the method "GPUCommandEncoder.popDebugGroup".
func (this GPUCommandEncoder) FuncPopDebugGroup() (fn js.Func[func()]) {
	bindings.FuncGPUCommandEncoderPopDebugGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PopDebugGroup calls the method "GPUCommandEncoder.popDebugGroup".
func (this GPUCommandEncoder) PopDebugGroup() (ret js.Void) {
	bindings.CallGPUCommandEncoderPopDebugGroup(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPopDebugGroup calls the method "GPUCommandEncoder.popDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryPopDebugGroup() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderPopDebugGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInsertDebugMarker returns true if the method "GPUCommandEncoder.insertDebugMarker" exists.
func (this GPUCommandEncoder) HasFuncInsertDebugMarker() bool {
	return js.True == bindings.HasFuncGPUCommandEncoderInsertDebugMarker(
		this.ref,
	)
}

// FuncInsertDebugMarker returns the method "GPUCommandEncoder.insertDebugMarker".
func (this GPUCommandEncoder) FuncInsertDebugMarker() (fn js.Func[func(markerLabel js.String)]) {
	bindings.FuncGPUCommandEncoderInsertDebugMarker(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertDebugMarker calls the method "GPUCommandEncoder.insertDebugMarker".
func (this GPUCommandEncoder) InsertDebugMarker(markerLabel js.String) (ret js.Void) {
	bindings.CallGPUCommandEncoderInsertDebugMarker(
		this.ref, js.Pointer(&ret),
		markerLabel.Ref(),
	)

	return
}

// TryInsertDebugMarker calls the method "GPUCommandEncoder.insertDebugMarker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCommandEncoder) TryInsertDebugMarker(markerLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCommandEncoderInsertDebugMarker(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *GPUCommandEncoderDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUCommandEncoderDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUCommandEncoderDescriptor) Update(ref js.Ref) {
	bindings.GPUCommandEncoderDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUCommandEncoderDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
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
func (p *GPURenderBundleDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPURenderBundleDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPURenderBundleDescriptor) Update(ref js.Ref) {
	bindings.GPURenderBundleDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPURenderBundleDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
}

type GPURenderBundleEncoder struct {
	ref js.Ref
}

func (this GPURenderBundleEncoder) Once() GPURenderBundleEncoder {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPURenderBundleEncoder.label".
//
// It returns ok=false if there is no such property.
func (this GPURenderBundleEncoder) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPURenderBundleEncoderLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPURenderBundleEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPURenderBundleEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPURenderBundleEncoderLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncFinish returns true if the method "GPURenderBundleEncoder.finish" exists.
func (this GPURenderBundleEncoder) HasFuncFinish() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderFinish(
		this.ref,
	)
}

// FuncFinish returns the method "GPURenderBundleEncoder.finish".
func (this GPURenderBundleEncoder) FuncFinish() (fn js.Func[func(descriptor GPURenderBundleDescriptor) GPURenderBundle]) {
	bindings.FuncGPURenderBundleEncoderFinish(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Finish calls the method "GPURenderBundleEncoder.finish".
func (this GPURenderBundleEncoder) Finish(descriptor GPURenderBundleDescriptor) (ret GPURenderBundle) {
	bindings.CallGPURenderBundleEncoderFinish(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryFinish calls the method "GPURenderBundleEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryFinish(descriptor GPURenderBundleDescriptor) (ret GPURenderBundle, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderFinish(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncFinish1 returns true if the method "GPURenderBundleEncoder.finish" exists.
func (this GPURenderBundleEncoder) HasFuncFinish1() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderFinish1(
		this.ref,
	)
}

// FuncFinish1 returns the method "GPURenderBundleEncoder.finish".
func (this GPURenderBundleEncoder) FuncFinish1() (fn js.Func[func() GPURenderBundle]) {
	bindings.FuncGPURenderBundleEncoderFinish1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Finish1 calls the method "GPURenderBundleEncoder.finish".
func (this GPURenderBundleEncoder) Finish1() (ret GPURenderBundle) {
	bindings.CallGPURenderBundleEncoderFinish1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryFinish1 calls the method "GPURenderBundleEncoder.finish"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryFinish1() (ret GPURenderBundle, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderFinish1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetPipeline returns true if the method "GPURenderBundleEncoder.setPipeline" exists.
func (this GPURenderBundleEncoder) HasFuncSetPipeline() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetPipeline(
		this.ref,
	)
}

// FuncSetPipeline returns the method "GPURenderBundleEncoder.setPipeline".
func (this GPURenderBundleEncoder) FuncSetPipeline() (fn js.Func[func(pipeline GPURenderPipeline)]) {
	bindings.FuncGPURenderBundleEncoderSetPipeline(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPipeline calls the method "GPURenderBundleEncoder.setPipeline".
func (this GPURenderBundleEncoder) SetPipeline(pipeline GPURenderPipeline) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetPipeline(
		this.ref, js.Pointer(&ret),
		pipeline.Ref(),
	)

	return
}

// TrySetPipeline calls the method "GPURenderBundleEncoder.setPipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TrySetPipeline(pipeline GPURenderPipeline) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderSetPipeline(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		pipeline.Ref(),
	)

	return
}

// HasFuncSetIndexBuffer returns true if the method "GPURenderBundleEncoder.setIndexBuffer" exists.
func (this GPURenderBundleEncoder) HasFuncSetIndexBuffer() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetIndexBuffer(
		this.ref,
	)
}

// FuncSetIndexBuffer returns the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) FuncSetIndexBuffer() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64)]) {
	bindings.FuncGPURenderBundleEncoderSetIndexBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetIndexBuffer calls the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) SetIndexBuffer(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetIndexBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
		float64(size),
	)

	return
}

// HasFuncSetIndexBuffer1 returns true if the method "GPURenderBundleEncoder.setIndexBuffer" exists.
func (this GPURenderBundleEncoder) HasFuncSetIndexBuffer1() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetIndexBuffer1(
		this.ref,
	)
}

// FuncSetIndexBuffer1 returns the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) FuncSetIndexBuffer1() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64)]) {
	bindings.FuncGPURenderBundleEncoderSetIndexBuffer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetIndexBuffer1 calls the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) SetIndexBuffer1(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetIndexBuffer1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
	)

	return
}

// HasFuncSetIndexBuffer2 returns true if the method "GPURenderBundleEncoder.setIndexBuffer" exists.
func (this GPURenderBundleEncoder) HasFuncSetIndexBuffer2() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetIndexBuffer2(
		this.ref,
	)
}

// FuncSetIndexBuffer2 returns the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) FuncSetIndexBuffer2() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat)]) {
	bindings.FuncGPURenderBundleEncoderSetIndexBuffer2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetIndexBuffer2 calls the method "GPURenderBundleEncoder.setIndexBuffer".
func (this GPURenderBundleEncoder) SetIndexBuffer2(buffer GPUBuffer, indexFormat GPUIndexFormat) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetIndexBuffer2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		uint32(indexFormat),
	)

	return
}

// HasFuncSetVertexBuffer returns true if the method "GPURenderBundleEncoder.setVertexBuffer" exists.
func (this GPURenderBundleEncoder) HasFuncSetVertexBuffer() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetVertexBuffer(
		this.ref,
	)
}

// FuncSetVertexBuffer returns the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) FuncSetVertexBuffer() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64)]) {
	bindings.FuncGPURenderBundleEncoderSetVertexBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetVertexBuffer calls the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) SetVertexBuffer(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetVertexBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	return
}

// HasFuncSetVertexBuffer1 returns true if the method "GPURenderBundleEncoder.setVertexBuffer" exists.
func (this GPURenderBundleEncoder) HasFuncSetVertexBuffer1() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetVertexBuffer1(
		this.ref,
	)
}

// FuncSetVertexBuffer1 returns the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) FuncSetVertexBuffer1() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64)]) {
	bindings.FuncGPURenderBundleEncoderSetVertexBuffer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetVertexBuffer1 calls the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) SetVertexBuffer1(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetVertexBuffer1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
	)

	return
}

// HasFuncSetVertexBuffer2 returns true if the method "GPURenderBundleEncoder.setVertexBuffer" exists.
func (this GPURenderBundleEncoder) HasFuncSetVertexBuffer2() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetVertexBuffer2(
		this.ref,
	)
}

// FuncSetVertexBuffer2 returns the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) FuncSetVertexBuffer2() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer)]) {
	bindings.FuncGPURenderBundleEncoderSetVertexBuffer2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetVertexBuffer2 calls the method "GPURenderBundleEncoder.setVertexBuffer".
func (this GPURenderBundleEncoder) SetVertexBuffer2(slot GPUIndex32, buffer GPUBuffer) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetVertexBuffer2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(slot),
		buffer.Ref(),
	)

	return
}

// HasFuncDraw returns true if the method "GPURenderBundleEncoder.draw" exists.
func (this GPURenderBundleEncoder) HasFuncDraw() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDraw(
		this.ref,
	)
}

// FuncDraw returns the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) FuncDraw() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32)]) {
	bindings.FuncGPURenderBundleEncoderDraw(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Draw calls the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDraw(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
		uint32(firstInstance),
	)

	return
}

// HasFuncDraw1 returns true if the method "GPURenderBundleEncoder.draw" exists.
func (this GPURenderBundleEncoder) HasFuncDraw1() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDraw1(
		this.ref,
	)
}

// FuncDraw1 returns the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) FuncDraw1() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32)]) {
	bindings.FuncGPURenderBundleEncoderDraw1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Draw1 calls the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw1(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDraw1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
	)

	return
}

// HasFuncDraw2 returns true if the method "GPURenderBundleEncoder.draw" exists.
func (this GPURenderBundleEncoder) HasFuncDraw2() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDraw2(
		this.ref,
	)
}

// FuncDraw2 returns the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) FuncDraw2() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32)]) {
	bindings.FuncGPURenderBundleEncoderDraw2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Draw2 calls the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw2(vertexCount GPUSize32, instanceCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDraw2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
		uint32(instanceCount),
	)

	return
}

// HasFuncDraw3 returns true if the method "GPURenderBundleEncoder.draw" exists.
func (this GPURenderBundleEncoder) HasFuncDraw3() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDraw3(
		this.ref,
	)
}

// FuncDraw3 returns the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) FuncDraw3() (fn js.Func[func(vertexCount GPUSize32)]) {
	bindings.FuncGPURenderBundleEncoderDraw3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Draw3 calls the method "GPURenderBundleEncoder.draw".
func (this GPURenderBundleEncoder) Draw3(vertexCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDraw3(
		this.ref, js.Pointer(&ret),
		uint32(vertexCount),
	)

	return
}

// TryDraw3 calls the method "GPURenderBundleEncoder.draw"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDraw3(vertexCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDraw3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(vertexCount),
	)

	return
}

// HasFuncDrawIndexed returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasFuncDrawIndexed() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDrawIndexed(
		this.ref,
	)
}

// FuncDrawIndexed returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) FuncDrawIndexed() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32)]) {
	bindings.FuncGPURenderBundleEncoderDrawIndexed(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
		uint32(firstInstance),
	)

	return
}

// HasFuncDrawIndexed1 returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasFuncDrawIndexed1() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDrawIndexed1(
		this.ref,
	)
}

// FuncDrawIndexed1 returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) FuncDrawIndexed1() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32)]) {
	bindings.FuncGPURenderBundleEncoderDrawIndexed1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed1 calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed1(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
	)

	return
}

// HasFuncDrawIndexed2 returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasFuncDrawIndexed2() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDrawIndexed2(
		this.ref,
	)
}

// FuncDrawIndexed2 returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) FuncDrawIndexed2() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32)]) {
	bindings.FuncGPURenderBundleEncoderDrawIndexed2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed2 calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed2(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
	)

	return
}

// HasFuncDrawIndexed3 returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasFuncDrawIndexed3() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDrawIndexed3(
		this.ref,
	)
}

// FuncDrawIndexed3 returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) FuncDrawIndexed3() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32)]) {
	bindings.FuncGPURenderBundleEncoderDrawIndexed3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed3 calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed3(indexCount GPUSize32, instanceCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed3(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
		uint32(instanceCount),
	)

	return
}

// HasFuncDrawIndexed4 returns true if the method "GPURenderBundleEncoder.drawIndexed" exists.
func (this GPURenderBundleEncoder) HasFuncDrawIndexed4() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDrawIndexed4(
		this.ref,
	)
}

// FuncDrawIndexed4 returns the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) FuncDrawIndexed4() (fn js.Func[func(indexCount GPUSize32)]) {
	bindings.FuncGPURenderBundleEncoderDrawIndexed4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexed4 calls the method "GPURenderBundleEncoder.drawIndexed".
func (this GPURenderBundleEncoder) DrawIndexed4(indexCount GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexed4(
		this.ref, js.Pointer(&ret),
		uint32(indexCount),
	)

	return
}

// TryDrawIndexed4 calls the method "GPURenderBundleEncoder.drawIndexed"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryDrawIndexed4(indexCount GPUSize32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderDrawIndexed4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(indexCount),
	)

	return
}

// HasFuncDrawIndirect returns true if the method "GPURenderBundleEncoder.drawIndirect" exists.
func (this GPURenderBundleEncoder) HasFuncDrawIndirect() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDrawIndirect(
		this.ref,
	)
}

// FuncDrawIndirect returns the method "GPURenderBundleEncoder.drawIndirect".
func (this GPURenderBundleEncoder) FuncDrawIndirect() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	bindings.FuncGPURenderBundleEncoderDrawIndirect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndirect calls the method "GPURenderBundleEncoder.drawIndirect".
func (this GPURenderBundleEncoder) DrawIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndirect(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasFuncDrawIndexedIndirect returns true if the method "GPURenderBundleEncoder.drawIndexedIndirect" exists.
func (this GPURenderBundleEncoder) HasFuncDrawIndexedIndirect() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderDrawIndexedIndirect(
		this.ref,
	)
}

// FuncDrawIndexedIndirect returns the method "GPURenderBundleEncoder.drawIndexedIndirect".
func (this GPURenderBundleEncoder) FuncDrawIndexedIndirect() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	bindings.FuncGPURenderBundleEncoderDrawIndexedIndirect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawIndexedIndirect calls the method "GPURenderBundleEncoder.drawIndexedIndirect".
func (this GPURenderBundleEncoder) DrawIndexedIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderDrawIndexedIndirect(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	return
}

// HasFuncSetBindGroup returns true if the method "GPURenderBundleEncoder.setBindGroup" exists.
func (this GPURenderBundleEncoder) HasFuncSetBindGroup() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetBindGroup(
		this.ref,
	)
}

// FuncSetBindGroup returns the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) FuncSetBindGroup() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset])]) {
	bindings.FuncGPURenderBundleEncoderSetBindGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBindGroup calls the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetBindGroup(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	return
}

// HasFuncSetBindGroup1 returns true if the method "GPURenderBundleEncoder.setBindGroup" exists.
func (this GPURenderBundleEncoder) HasFuncSetBindGroup1() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetBindGroup1(
		this.ref,
	)
}

// FuncSetBindGroup1 returns the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) FuncSetBindGroup1() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup)]) {
	bindings.FuncGPURenderBundleEncoderSetBindGroup1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBindGroup1 calls the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) SetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetBindGroup1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
	)

	return
}

// HasFuncSetBindGroup2 returns true if the method "GPURenderBundleEncoder.setBindGroup" exists.
func (this GPURenderBundleEncoder) HasFuncSetBindGroup2() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderSetBindGroup2(
		this.ref,
	)
}

// FuncSetBindGroup2 returns the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) FuncSetBindGroup2() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32)]) {
	bindings.FuncGPURenderBundleEncoderSetBindGroup2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetBindGroup2 calls the method "GPURenderBundleEncoder.setBindGroup".
func (this GPURenderBundleEncoder) SetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderSetBindGroup2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	return
}

// HasFuncPushDebugGroup returns true if the method "GPURenderBundleEncoder.pushDebugGroup" exists.
func (this GPURenderBundleEncoder) HasFuncPushDebugGroup() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderPushDebugGroup(
		this.ref,
	)
}

// FuncPushDebugGroup returns the method "GPURenderBundleEncoder.pushDebugGroup".
func (this GPURenderBundleEncoder) FuncPushDebugGroup() (fn js.Func[func(groupLabel js.String)]) {
	bindings.FuncGPURenderBundleEncoderPushDebugGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PushDebugGroup calls the method "GPURenderBundleEncoder.pushDebugGroup".
func (this GPURenderBundleEncoder) PushDebugGroup(groupLabel js.String) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderPushDebugGroup(
		this.ref, js.Pointer(&ret),
		groupLabel.Ref(),
	)

	return
}

// TryPushDebugGroup calls the method "GPURenderBundleEncoder.pushDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryPushDebugGroup(groupLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderPushDebugGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		groupLabel.Ref(),
	)

	return
}

// HasFuncPopDebugGroup returns true if the method "GPURenderBundleEncoder.popDebugGroup" exists.
func (this GPURenderBundleEncoder) HasFuncPopDebugGroup() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderPopDebugGroup(
		this.ref,
	)
}

// FuncPopDebugGroup returns the method "GPURenderBundleEncoder.popDebugGroup".
func (this GPURenderBundleEncoder) FuncPopDebugGroup() (fn js.Func[func()]) {
	bindings.FuncGPURenderBundleEncoderPopDebugGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PopDebugGroup calls the method "GPURenderBundleEncoder.popDebugGroup".
func (this GPURenderBundleEncoder) PopDebugGroup() (ret js.Void) {
	bindings.CallGPURenderBundleEncoderPopDebugGroup(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPopDebugGroup calls the method "GPURenderBundleEncoder.popDebugGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryPopDebugGroup() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderPopDebugGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInsertDebugMarker returns true if the method "GPURenderBundleEncoder.insertDebugMarker" exists.
func (this GPURenderBundleEncoder) HasFuncInsertDebugMarker() bool {
	return js.True == bindings.HasFuncGPURenderBundleEncoderInsertDebugMarker(
		this.ref,
	)
}

// FuncInsertDebugMarker returns the method "GPURenderBundleEncoder.insertDebugMarker".
func (this GPURenderBundleEncoder) FuncInsertDebugMarker() (fn js.Func[func(markerLabel js.String)]) {
	bindings.FuncGPURenderBundleEncoderInsertDebugMarker(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertDebugMarker calls the method "GPURenderBundleEncoder.insertDebugMarker".
func (this GPURenderBundleEncoder) InsertDebugMarker(markerLabel js.String) (ret js.Void) {
	bindings.CallGPURenderBundleEncoderInsertDebugMarker(
		this.ref, js.Pointer(&ret),
		markerLabel.Ref(),
	)

	return
}

// TryInsertDebugMarker calls the method "GPURenderBundleEncoder.insertDebugMarker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPURenderBundleEncoder) TryInsertDebugMarker(markerLabel js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPURenderBundleEncoderInsertDebugMarker(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *GPURenderBundleEncoderDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPURenderBundleEncoderDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPURenderBundleEncoderDescriptor) Update(ref js.Ref) {
	bindings.GPURenderBundleEncoderDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPURenderBundleEncoderDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.ColorFormats.Ref(),
		p.Label.Ref(),
	)
	p.ColorFormats = p.ColorFormats.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
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
func (p *GPUQuerySetDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GPUQuerySetDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUQuerySetDescriptor) Update(ref js.Ref) {
	bindings.GPUQuerySetDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUQuerySetDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Message returns the value of property "GPUError.message".
//
// It returns ok=false if there is no such property.
func (this GPUError) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUErrorMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GPUSupportedFeatures struct {
	ref js.Ref
}

func (this GPUSupportedFeatures) Once() GPUSupportedFeatures {
	this.ref.Once()
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
	this.ref.Free()
}

type GPUSupportedLimits struct {
	ref js.Ref
}

func (this GPUSupportedLimits) Once() GPUSupportedLimits {
	this.ref.Once()
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
	this.ref.Free()
}

// MaxTextureDimension1D returns the value of property "GPUSupportedLimits.maxTextureDimension1D".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxTextureDimension1D() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxTextureDimension1D(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxTextureDimension2D returns the value of property "GPUSupportedLimits.maxTextureDimension2D".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxTextureDimension2D() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxTextureDimension2D(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxTextureDimension3D returns the value of property "GPUSupportedLimits.maxTextureDimension3D".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxTextureDimension3D() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxTextureDimension3D(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxTextureArrayLayers returns the value of property "GPUSupportedLimits.maxTextureArrayLayers".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxTextureArrayLayers() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxTextureArrayLayers(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxBindGroups returns the value of property "GPUSupportedLimits.maxBindGroups".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxBindGroups() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxBindGroups(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxBindGroupsPlusVertexBuffers returns the value of property "GPUSupportedLimits.maxBindGroupsPlusVertexBuffers".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxBindGroupsPlusVertexBuffers() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxBindGroupsPlusVertexBuffers(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxBindingsPerBindGroup returns the value of property "GPUSupportedLimits.maxBindingsPerBindGroup".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxBindingsPerBindGroup() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxBindingsPerBindGroup(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxDynamicUniformBuffersPerPipelineLayout returns the value of property "GPUSupportedLimits.maxDynamicUniformBuffersPerPipelineLayout".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxDynamicUniformBuffersPerPipelineLayout() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxDynamicUniformBuffersPerPipelineLayout(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxDynamicStorageBuffersPerPipelineLayout returns the value of property "GPUSupportedLimits.maxDynamicStorageBuffersPerPipelineLayout".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxDynamicStorageBuffersPerPipelineLayout() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxDynamicStorageBuffersPerPipelineLayout(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxSampledTexturesPerShaderStage returns the value of property "GPUSupportedLimits.maxSampledTexturesPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxSampledTexturesPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxSampledTexturesPerShaderStage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxSamplersPerShaderStage returns the value of property "GPUSupportedLimits.maxSamplersPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxSamplersPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxSamplersPerShaderStage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxStorageBuffersPerShaderStage returns the value of property "GPUSupportedLimits.maxStorageBuffersPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxStorageBuffersPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxStorageBuffersPerShaderStage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxStorageTexturesPerShaderStage returns the value of property "GPUSupportedLimits.maxStorageTexturesPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxStorageTexturesPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxStorageTexturesPerShaderStage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxUniformBuffersPerShaderStage returns the value of property "GPUSupportedLimits.maxUniformBuffersPerShaderStage".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxUniformBuffersPerShaderStage() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxUniformBuffersPerShaderStage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxUniformBufferBindingSize returns the value of property "GPUSupportedLimits.maxUniformBufferBindingSize".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxUniformBufferBindingSize() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxUniformBufferBindingSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxStorageBufferBindingSize returns the value of property "GPUSupportedLimits.maxStorageBufferBindingSize".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxStorageBufferBindingSize() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxStorageBufferBindingSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MinUniformBufferOffsetAlignment returns the value of property "GPUSupportedLimits.minUniformBufferOffsetAlignment".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MinUniformBufferOffsetAlignment() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMinUniformBufferOffsetAlignment(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MinStorageBufferOffsetAlignment returns the value of property "GPUSupportedLimits.minStorageBufferOffsetAlignment".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MinStorageBufferOffsetAlignment() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMinStorageBufferOffsetAlignment(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxVertexBuffers returns the value of property "GPUSupportedLimits.maxVertexBuffers".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxVertexBuffers() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxVertexBuffers(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxBufferSize returns the value of property "GPUSupportedLimits.maxBufferSize".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxBufferSize() (ret uint64, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxBufferSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxVertexAttributes returns the value of property "GPUSupportedLimits.maxVertexAttributes".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxVertexAttributes() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxVertexAttributes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxVertexBufferArrayStride returns the value of property "GPUSupportedLimits.maxVertexBufferArrayStride".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxVertexBufferArrayStride() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxVertexBufferArrayStride(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxInterStageShaderComponents returns the value of property "GPUSupportedLimits.maxInterStageShaderComponents".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxInterStageShaderComponents() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxInterStageShaderComponents(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxInterStageShaderVariables returns the value of property "GPUSupportedLimits.maxInterStageShaderVariables".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxInterStageShaderVariables() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxInterStageShaderVariables(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxColorAttachments returns the value of property "GPUSupportedLimits.maxColorAttachments".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxColorAttachments() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxColorAttachments(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxColorAttachmentBytesPerSample returns the value of property "GPUSupportedLimits.maxColorAttachmentBytesPerSample".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxColorAttachmentBytesPerSample() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxColorAttachmentBytesPerSample(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupStorageSize returns the value of property "GPUSupportedLimits.maxComputeWorkgroupStorageSize".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupStorageSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupStorageSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxComputeInvocationsPerWorkgroup returns the value of property "GPUSupportedLimits.maxComputeInvocationsPerWorkgroup".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeInvocationsPerWorkgroup() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeInvocationsPerWorkgroup(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupSizeX returns the value of property "GPUSupportedLimits.maxComputeWorkgroupSizeX".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupSizeX() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupSizeX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupSizeY returns the value of property "GPUSupportedLimits.maxComputeWorkgroupSizeY".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupSizeY() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupSizeY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupSizeZ returns the value of property "GPUSupportedLimits.maxComputeWorkgroupSizeZ".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupSizeZ() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupSizeZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxComputeWorkgroupsPerDimension returns the value of property "GPUSupportedLimits.maxComputeWorkgroupsPerDimension".
//
// It returns ok=false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupsPerDimension() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGPUSupportedLimitsMaxComputeWorkgroupsPerDimension(
		this.ref, js.Pointer(&ret),
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
func (p *GPUImageDataLayout) UpdateFrom(ref js.Ref) {
	bindings.GPUImageDataLayoutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUImageDataLayout) Update(ref js.Ref) {
	bindings.GPUImageDataLayoutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUImageDataLayout) FreeMembers(recursive bool) {
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
func (p *GPUOrigin2DDict) UpdateFrom(ref js.Ref) {
	bindings.GPUOrigin2DDictJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUOrigin2DDict) Update(ref js.Ref) {
	bindings.GPUOrigin2DDictJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUOrigin2DDict) FreeMembers(recursive bool) {
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
func (p *GPUImageCopyExternalImage) UpdateFrom(ref js.Ref) {
	bindings.GPUImageCopyExternalImageJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUImageCopyExternalImage) Update(ref js.Ref) {
	bindings.GPUImageCopyExternalImageJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUImageCopyExternalImage) FreeMembers(recursive bool) {
	js.Free(
		p.Source.Ref(),
		p.Origin.Ref(),
	)
	p.Source = p.Source.FromRef(js.Undefined)
	p.Origin = p.Origin.FromRef(js.Undefined)
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
func (p *GPUImageCopyTextureTagged) UpdateFrom(ref js.Ref) {
	bindings.GPUImageCopyTextureTaggedJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUImageCopyTextureTagged) Update(ref js.Ref) {
	bindings.GPUImageCopyTextureTaggedJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUImageCopyTextureTagged) FreeMembers(recursive bool) {
	js.Free(
		p.Texture.Ref(),
		p.Origin.Ref(),
	)
	p.Texture = p.Texture.FromRef(js.Undefined)
	p.Origin = p.Origin.FromRef(js.Undefined)
}

type GPUQueue struct {
	ref js.Ref
}

func (this GPUQueue) Once() GPUQueue {
	this.ref.Once()
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
	this.ref.Free()
}

// Label returns the value of property "GPUQueue.label".
//
// It returns ok=false if there is no such property.
func (this GPUQueue) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUQueueLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUQueue.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUQueue) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUQueueLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncSubmit returns true if the method "GPUQueue.submit" exists.
func (this GPUQueue) HasFuncSubmit() bool {
	return js.True == bindings.HasFuncGPUQueueSubmit(
		this.ref,
	)
}

// FuncSubmit returns the method "GPUQueue.submit".
func (this GPUQueue) FuncSubmit() (fn js.Func[func(commandBuffers js.Array[GPUCommandBuffer])]) {
	bindings.FuncGPUQueueSubmit(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Submit calls the method "GPUQueue.submit".
func (this GPUQueue) Submit(commandBuffers js.Array[GPUCommandBuffer]) (ret js.Void) {
	bindings.CallGPUQueueSubmit(
		this.ref, js.Pointer(&ret),
		commandBuffers.Ref(),
	)

	return
}

// TrySubmit calls the method "GPUQueue.submit"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQueue) TrySubmit(commandBuffers js.Array[GPUCommandBuffer]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQueueSubmit(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		commandBuffers.Ref(),
	)

	return
}

// HasFuncOnSubmittedWorkDone returns true if the method "GPUQueue.onSubmittedWorkDone" exists.
func (this GPUQueue) HasFuncOnSubmittedWorkDone() bool {
	return js.True == bindings.HasFuncGPUQueueOnSubmittedWorkDone(
		this.ref,
	)
}

// FuncOnSubmittedWorkDone returns the method "GPUQueue.onSubmittedWorkDone".
func (this GPUQueue) FuncOnSubmittedWorkDone() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncGPUQueueOnSubmittedWorkDone(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OnSubmittedWorkDone calls the method "GPUQueue.onSubmittedWorkDone".
func (this GPUQueue) OnSubmittedWorkDone() (ret js.Promise[js.Void]) {
	bindings.CallGPUQueueOnSubmittedWorkDone(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOnSubmittedWorkDone calls the method "GPUQueue.onSubmittedWorkDone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUQueue) TryOnSubmittedWorkDone() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUQueueOnSubmittedWorkDone(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncWriteBuffer returns true if the method "GPUQueue.writeBuffer" exists.
func (this GPUQueue) HasFuncWriteBuffer() bool {
	return js.True == bindings.HasFuncGPUQueueWriteBuffer(
		this.ref,
	)
}

// FuncWriteBuffer returns the method "GPUQueue.writeBuffer".
func (this GPUQueue) FuncWriteBuffer() (fn js.Func[func(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64, size GPUSize64)]) {
	bindings.FuncGPUQueueWriteBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteBuffer calls the method "GPUQueue.writeBuffer".
func (this GPUQueue) WriteBuffer(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64, size GPUSize64) (ret js.Void) {
	bindings.CallGPUQueueWriteBuffer(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
		float64(dataOffset),
		float64(size),
	)

	return
}

// HasFuncWriteBuffer1 returns true if the method "GPUQueue.writeBuffer" exists.
func (this GPUQueue) HasFuncWriteBuffer1() bool {
	return js.True == bindings.HasFuncGPUQueueWriteBuffer1(
		this.ref,
	)
}

// FuncWriteBuffer1 returns the method "GPUQueue.writeBuffer".
func (this GPUQueue) FuncWriteBuffer1() (fn js.Func[func(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64)]) {
	bindings.FuncGPUQueueWriteBuffer1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteBuffer1 calls the method "GPUQueue.writeBuffer".
func (this GPUQueue) WriteBuffer1(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64) (ret js.Void) {
	bindings.CallGPUQueueWriteBuffer1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
		float64(dataOffset),
	)

	return
}

// HasFuncWriteBuffer2 returns true if the method "GPUQueue.writeBuffer" exists.
func (this GPUQueue) HasFuncWriteBuffer2() bool {
	return js.True == bindings.HasFuncGPUQueueWriteBuffer2(
		this.ref,
	)
}

// FuncWriteBuffer2 returns the method "GPUQueue.writeBuffer".
func (this GPUQueue) FuncWriteBuffer2() (fn js.Func[func(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource)]) {
	bindings.FuncGPUQueueWriteBuffer2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteBuffer2 calls the method "GPUQueue.writeBuffer".
func (this GPUQueue) WriteBuffer2(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource) (ret js.Void) {
	bindings.CallGPUQueueWriteBuffer2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
	)

	return
}

// HasFuncWriteTexture returns true if the method "GPUQueue.writeTexture" exists.
func (this GPUQueue) HasFuncWriteTexture() bool {
	return js.True == bindings.HasFuncGPUQueueWriteTexture(
		this.ref,
	)
}

// FuncWriteTexture returns the method "GPUQueue.writeTexture".
func (this GPUQueue) FuncWriteTexture() (fn js.Func[func(destination GPUImageCopyTexture, data AllowSharedBufferSource, dataLayout GPUImageDataLayout, size GPUExtent3D)]) {
	bindings.FuncGPUQueueWriteTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WriteTexture calls the method "GPUQueue.writeTexture".
func (this GPUQueue) WriteTexture(destination GPUImageCopyTexture, data AllowSharedBufferSource, dataLayout GPUImageDataLayout, size GPUExtent3D) (ret js.Void) {
	bindings.CallGPUQueueWriteTexture(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&destination),
		data.Ref(),
		js.Pointer(&dataLayout),
		size.Ref(),
	)

	return
}

// HasFuncCopyExternalImageToTexture returns true if the method "GPUQueue.copyExternalImageToTexture" exists.
func (this GPUQueue) HasFuncCopyExternalImageToTexture() bool {
	return js.True == bindings.HasFuncGPUQueueCopyExternalImageToTexture(
		this.ref,
	)
}

// FuncCopyExternalImageToTexture returns the method "GPUQueue.copyExternalImageToTexture".
func (this GPUQueue) FuncCopyExternalImageToTexture() (fn js.Func[func(source GPUImageCopyExternalImage, destination GPUImageCopyTextureTagged, copySize GPUExtent3D)]) {
	bindings.FuncGPUQueueCopyExternalImageToTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CopyExternalImageToTexture calls the method "GPUQueue.copyExternalImageToTexture".
func (this GPUQueue) CopyExternalImageToTexture(source GPUImageCopyExternalImage, destination GPUImageCopyTextureTagged, copySize GPUExtent3D) (ret js.Void) {
	bindings.CallGPUQueueCopyExternalImageToTexture(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Reason returns the value of property "GPUDeviceLostInfo.reason".
//
// It returns ok=false if there is no such property.
func (this GPUDeviceLostInfo) Reason() (ret GPUDeviceLostReason, ok bool) {
	ok = js.True == bindings.GetGPUDeviceLostInfoReason(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "GPUDeviceLostInfo.message".
//
// It returns ok=false if there is no such property.
func (this GPUDeviceLostInfo) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUDeviceLostInfoMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GPUDevice struct {
	EventTarget
}

func (this GPUDevice) Once() GPUDevice {
	this.ref.Once()
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
	this.ref.Free()
}

// Features returns the value of property "GPUDevice.features".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Features() (ret GPUSupportedFeatures, ok bool) {
	ok = js.True == bindings.GetGPUDeviceFeatures(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Limits returns the value of property "GPUDevice.limits".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Limits() (ret GPUSupportedLimits, ok bool) {
	ok = js.True == bindings.GetGPUDeviceLimits(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Queue returns the value of property "GPUDevice.queue".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Queue() (ret GPUQueue, ok bool) {
	ok = js.True == bindings.GetGPUDeviceQueue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Lost returns the value of property "GPUDevice.lost".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Lost() (ret js.Promise[GPUDeviceLostInfo], ok bool) {
	ok = js.True == bindings.GetGPUDeviceLost(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "GPUDevice.label".
//
// It returns ok=false if there is no such property.
func (this GPUDevice) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGPUDeviceLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "GPUDevice.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUDevice) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUDeviceLabel(
		this.ref,
		val.Ref(),
	)
}

// HasFuncDestroy returns true if the method "GPUDevice.destroy" exists.
func (this GPUDevice) HasFuncDestroy() bool {
	return js.True == bindings.HasFuncGPUDeviceDestroy(
		this.ref,
	)
}

// FuncDestroy returns the method "GPUDevice.destroy".
func (this GPUDevice) FuncDestroy() (fn js.Func[func()]) {
	bindings.FuncGPUDeviceDestroy(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Destroy calls the method "GPUDevice.destroy".
func (this GPUDevice) Destroy() (ret js.Void) {
	bindings.CallGPUDeviceDestroy(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDestroy calls the method "GPUDevice.destroy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryDestroy() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceDestroy(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateBuffer returns true if the method "GPUDevice.createBuffer" exists.
func (this GPUDevice) HasFuncCreateBuffer() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateBuffer(
		this.ref,
	)
}

// FuncCreateBuffer returns the method "GPUDevice.createBuffer".
func (this GPUDevice) FuncCreateBuffer() (fn js.Func[func(descriptor GPUBufferDescriptor) GPUBuffer]) {
	bindings.FuncGPUDeviceCreateBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBuffer calls the method "GPUDevice.createBuffer".
func (this GPUDevice) CreateBuffer(descriptor GPUBufferDescriptor) (ret GPUBuffer) {
	bindings.CallGPUDeviceCreateBuffer(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateBuffer calls the method "GPUDevice.createBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateBuffer(descriptor GPUBufferDescriptor) (ret GPUBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateTexture returns true if the method "GPUDevice.createTexture" exists.
func (this GPUDevice) HasFuncCreateTexture() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateTexture(
		this.ref,
	)
}

// FuncCreateTexture returns the method "GPUDevice.createTexture".
func (this GPUDevice) FuncCreateTexture() (fn js.Func[func(descriptor GPUTextureDescriptor) GPUTexture]) {
	bindings.FuncGPUDeviceCreateTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTexture calls the method "GPUDevice.createTexture".
func (this GPUDevice) CreateTexture(descriptor GPUTextureDescriptor) (ret GPUTexture) {
	bindings.CallGPUDeviceCreateTexture(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateTexture calls the method "GPUDevice.createTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateTexture(descriptor GPUTextureDescriptor) (ret GPUTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateSampler returns true if the method "GPUDevice.createSampler" exists.
func (this GPUDevice) HasFuncCreateSampler() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateSampler(
		this.ref,
	)
}

// FuncCreateSampler returns the method "GPUDevice.createSampler".
func (this GPUDevice) FuncCreateSampler() (fn js.Func[func(descriptor GPUSamplerDescriptor) GPUSampler]) {
	bindings.FuncGPUDeviceCreateSampler(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSampler calls the method "GPUDevice.createSampler".
func (this GPUDevice) CreateSampler(descriptor GPUSamplerDescriptor) (ret GPUSampler) {
	bindings.CallGPUDeviceCreateSampler(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateSampler calls the method "GPUDevice.createSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateSampler(descriptor GPUSamplerDescriptor) (ret GPUSampler, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateSampler(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateSampler1 returns true if the method "GPUDevice.createSampler" exists.
func (this GPUDevice) HasFuncCreateSampler1() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateSampler1(
		this.ref,
	)
}

// FuncCreateSampler1 returns the method "GPUDevice.createSampler".
func (this GPUDevice) FuncCreateSampler1() (fn js.Func[func() GPUSampler]) {
	bindings.FuncGPUDeviceCreateSampler1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateSampler1 calls the method "GPUDevice.createSampler".
func (this GPUDevice) CreateSampler1() (ret GPUSampler) {
	bindings.CallGPUDeviceCreateSampler1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateSampler1 calls the method "GPUDevice.createSampler"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateSampler1() (ret GPUSampler, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateSampler1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncImportExternalTexture returns true if the method "GPUDevice.importExternalTexture" exists.
func (this GPUDevice) HasFuncImportExternalTexture() bool {
	return js.True == bindings.HasFuncGPUDeviceImportExternalTexture(
		this.ref,
	)
}

// FuncImportExternalTexture returns the method "GPUDevice.importExternalTexture".
func (this GPUDevice) FuncImportExternalTexture() (fn js.Func[func(descriptor GPUExternalTextureDescriptor) GPUExternalTexture]) {
	bindings.FuncGPUDeviceImportExternalTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ImportExternalTexture calls the method "GPUDevice.importExternalTexture".
func (this GPUDevice) ImportExternalTexture(descriptor GPUExternalTextureDescriptor) (ret GPUExternalTexture) {
	bindings.CallGPUDeviceImportExternalTexture(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryImportExternalTexture calls the method "GPUDevice.importExternalTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryImportExternalTexture(descriptor GPUExternalTextureDescriptor) (ret GPUExternalTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceImportExternalTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateBindGroupLayout returns true if the method "GPUDevice.createBindGroupLayout" exists.
func (this GPUDevice) HasFuncCreateBindGroupLayout() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateBindGroupLayout(
		this.ref,
	)
}

// FuncCreateBindGroupLayout returns the method "GPUDevice.createBindGroupLayout".
func (this GPUDevice) FuncCreateBindGroupLayout() (fn js.Func[func(descriptor GPUBindGroupLayoutDescriptor) GPUBindGroupLayout]) {
	bindings.FuncGPUDeviceCreateBindGroupLayout(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBindGroupLayout calls the method "GPUDevice.createBindGroupLayout".
func (this GPUDevice) CreateBindGroupLayout(descriptor GPUBindGroupLayoutDescriptor) (ret GPUBindGroupLayout) {
	bindings.CallGPUDeviceCreateBindGroupLayout(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateBindGroupLayout calls the method "GPUDevice.createBindGroupLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateBindGroupLayout(descriptor GPUBindGroupLayoutDescriptor) (ret GPUBindGroupLayout, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateBindGroupLayout(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreatePipelineLayout returns true if the method "GPUDevice.createPipelineLayout" exists.
func (this GPUDevice) HasFuncCreatePipelineLayout() bool {
	return js.True == bindings.HasFuncGPUDeviceCreatePipelineLayout(
		this.ref,
	)
}

// FuncCreatePipelineLayout returns the method "GPUDevice.createPipelineLayout".
func (this GPUDevice) FuncCreatePipelineLayout() (fn js.Func[func(descriptor GPUPipelineLayoutDescriptor) GPUPipelineLayout]) {
	bindings.FuncGPUDeviceCreatePipelineLayout(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreatePipelineLayout calls the method "GPUDevice.createPipelineLayout".
func (this GPUDevice) CreatePipelineLayout(descriptor GPUPipelineLayoutDescriptor) (ret GPUPipelineLayout) {
	bindings.CallGPUDeviceCreatePipelineLayout(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreatePipelineLayout calls the method "GPUDevice.createPipelineLayout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreatePipelineLayout(descriptor GPUPipelineLayoutDescriptor) (ret GPUPipelineLayout, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreatePipelineLayout(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateBindGroup returns true if the method "GPUDevice.createBindGroup" exists.
func (this GPUDevice) HasFuncCreateBindGroup() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateBindGroup(
		this.ref,
	)
}

// FuncCreateBindGroup returns the method "GPUDevice.createBindGroup".
func (this GPUDevice) FuncCreateBindGroup() (fn js.Func[func(descriptor GPUBindGroupDescriptor) GPUBindGroup]) {
	bindings.FuncGPUDeviceCreateBindGroup(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBindGroup calls the method "GPUDevice.createBindGroup".
func (this GPUDevice) CreateBindGroup(descriptor GPUBindGroupDescriptor) (ret GPUBindGroup) {
	bindings.CallGPUDeviceCreateBindGroup(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateBindGroup calls the method "GPUDevice.createBindGroup"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateBindGroup(descriptor GPUBindGroupDescriptor) (ret GPUBindGroup, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateBindGroup(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateShaderModule returns true if the method "GPUDevice.createShaderModule" exists.
func (this GPUDevice) HasFuncCreateShaderModule() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateShaderModule(
		this.ref,
	)
}

// FuncCreateShaderModule returns the method "GPUDevice.createShaderModule".
func (this GPUDevice) FuncCreateShaderModule() (fn js.Func[func(descriptor GPUShaderModuleDescriptor) GPUShaderModule]) {
	bindings.FuncGPUDeviceCreateShaderModule(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateShaderModule calls the method "GPUDevice.createShaderModule".
func (this GPUDevice) CreateShaderModule(descriptor GPUShaderModuleDescriptor) (ret GPUShaderModule) {
	bindings.CallGPUDeviceCreateShaderModule(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateShaderModule calls the method "GPUDevice.createShaderModule"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateShaderModule(descriptor GPUShaderModuleDescriptor) (ret GPUShaderModule, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateShaderModule(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateComputePipeline returns true if the method "GPUDevice.createComputePipeline" exists.
func (this GPUDevice) HasFuncCreateComputePipeline() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateComputePipeline(
		this.ref,
	)
}

// FuncCreateComputePipeline returns the method "GPUDevice.createComputePipeline".
func (this GPUDevice) FuncCreateComputePipeline() (fn js.Func[func(descriptor GPUComputePipelineDescriptor) GPUComputePipeline]) {
	bindings.FuncGPUDeviceCreateComputePipeline(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateComputePipeline calls the method "GPUDevice.createComputePipeline".
func (this GPUDevice) CreateComputePipeline(descriptor GPUComputePipelineDescriptor) (ret GPUComputePipeline) {
	bindings.CallGPUDeviceCreateComputePipeline(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateComputePipeline calls the method "GPUDevice.createComputePipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateComputePipeline(descriptor GPUComputePipelineDescriptor) (ret GPUComputePipeline, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateComputePipeline(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateRenderPipeline returns true if the method "GPUDevice.createRenderPipeline" exists.
func (this GPUDevice) HasFuncCreateRenderPipeline() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateRenderPipeline(
		this.ref,
	)
}

// FuncCreateRenderPipeline returns the method "GPUDevice.createRenderPipeline".
func (this GPUDevice) FuncCreateRenderPipeline() (fn js.Func[func(descriptor GPURenderPipelineDescriptor) GPURenderPipeline]) {
	bindings.FuncGPUDeviceCreateRenderPipeline(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateRenderPipeline calls the method "GPUDevice.createRenderPipeline".
func (this GPUDevice) CreateRenderPipeline(descriptor GPURenderPipelineDescriptor) (ret GPURenderPipeline) {
	bindings.CallGPUDeviceCreateRenderPipeline(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateRenderPipeline calls the method "GPUDevice.createRenderPipeline"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateRenderPipeline(descriptor GPURenderPipelineDescriptor) (ret GPURenderPipeline, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateRenderPipeline(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateComputePipelineAsync returns true if the method "GPUDevice.createComputePipelineAsync" exists.
func (this GPUDevice) HasFuncCreateComputePipelineAsync() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateComputePipelineAsync(
		this.ref,
	)
}

// FuncCreateComputePipelineAsync returns the method "GPUDevice.createComputePipelineAsync".
func (this GPUDevice) FuncCreateComputePipelineAsync() (fn js.Func[func(descriptor GPUComputePipelineDescriptor) js.Promise[GPUComputePipeline]]) {
	bindings.FuncGPUDeviceCreateComputePipelineAsync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateComputePipelineAsync calls the method "GPUDevice.createComputePipelineAsync".
func (this GPUDevice) CreateComputePipelineAsync(descriptor GPUComputePipelineDescriptor) (ret js.Promise[GPUComputePipeline]) {
	bindings.CallGPUDeviceCreateComputePipelineAsync(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateComputePipelineAsync calls the method "GPUDevice.createComputePipelineAsync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateComputePipelineAsync(descriptor GPUComputePipelineDescriptor) (ret js.Promise[GPUComputePipeline], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateComputePipelineAsync(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateRenderPipelineAsync returns true if the method "GPUDevice.createRenderPipelineAsync" exists.
func (this GPUDevice) HasFuncCreateRenderPipelineAsync() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateRenderPipelineAsync(
		this.ref,
	)
}

// FuncCreateRenderPipelineAsync returns the method "GPUDevice.createRenderPipelineAsync".
func (this GPUDevice) FuncCreateRenderPipelineAsync() (fn js.Func[func(descriptor GPURenderPipelineDescriptor) js.Promise[GPURenderPipeline]]) {
	bindings.FuncGPUDeviceCreateRenderPipelineAsync(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateRenderPipelineAsync calls the method "GPUDevice.createRenderPipelineAsync".
func (this GPUDevice) CreateRenderPipelineAsync(descriptor GPURenderPipelineDescriptor) (ret js.Promise[GPURenderPipeline]) {
	bindings.CallGPUDeviceCreateRenderPipelineAsync(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateRenderPipelineAsync calls the method "GPUDevice.createRenderPipelineAsync"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateRenderPipelineAsync(descriptor GPURenderPipelineDescriptor) (ret js.Promise[GPURenderPipeline], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateRenderPipelineAsync(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateCommandEncoder returns true if the method "GPUDevice.createCommandEncoder" exists.
func (this GPUDevice) HasFuncCreateCommandEncoder() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateCommandEncoder(
		this.ref,
	)
}

// FuncCreateCommandEncoder returns the method "GPUDevice.createCommandEncoder".
func (this GPUDevice) FuncCreateCommandEncoder() (fn js.Func[func(descriptor GPUCommandEncoderDescriptor) GPUCommandEncoder]) {
	bindings.FuncGPUDeviceCreateCommandEncoder(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCommandEncoder calls the method "GPUDevice.createCommandEncoder".
func (this GPUDevice) CreateCommandEncoder(descriptor GPUCommandEncoderDescriptor) (ret GPUCommandEncoder) {
	bindings.CallGPUDeviceCreateCommandEncoder(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateCommandEncoder calls the method "GPUDevice.createCommandEncoder"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateCommandEncoder(descriptor GPUCommandEncoderDescriptor) (ret GPUCommandEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateCommandEncoder(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateCommandEncoder1 returns true if the method "GPUDevice.createCommandEncoder" exists.
func (this GPUDevice) HasFuncCreateCommandEncoder1() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateCommandEncoder1(
		this.ref,
	)
}

// FuncCreateCommandEncoder1 returns the method "GPUDevice.createCommandEncoder".
func (this GPUDevice) FuncCreateCommandEncoder1() (fn js.Func[func() GPUCommandEncoder]) {
	bindings.FuncGPUDeviceCreateCommandEncoder1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCommandEncoder1 calls the method "GPUDevice.createCommandEncoder".
func (this GPUDevice) CreateCommandEncoder1() (ret GPUCommandEncoder) {
	bindings.CallGPUDeviceCreateCommandEncoder1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateCommandEncoder1 calls the method "GPUDevice.createCommandEncoder"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateCommandEncoder1() (ret GPUCommandEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateCommandEncoder1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateRenderBundleEncoder returns true if the method "GPUDevice.createRenderBundleEncoder" exists.
func (this GPUDevice) HasFuncCreateRenderBundleEncoder() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateRenderBundleEncoder(
		this.ref,
	)
}

// FuncCreateRenderBundleEncoder returns the method "GPUDevice.createRenderBundleEncoder".
func (this GPUDevice) FuncCreateRenderBundleEncoder() (fn js.Func[func(descriptor GPURenderBundleEncoderDescriptor) GPURenderBundleEncoder]) {
	bindings.FuncGPUDeviceCreateRenderBundleEncoder(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateRenderBundleEncoder calls the method "GPUDevice.createRenderBundleEncoder".
func (this GPUDevice) CreateRenderBundleEncoder(descriptor GPURenderBundleEncoderDescriptor) (ret GPURenderBundleEncoder) {
	bindings.CallGPUDeviceCreateRenderBundleEncoder(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateRenderBundleEncoder calls the method "GPUDevice.createRenderBundleEncoder"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateRenderBundleEncoder(descriptor GPURenderBundleEncoderDescriptor) (ret GPURenderBundleEncoder, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateRenderBundleEncoder(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncCreateQuerySet returns true if the method "GPUDevice.createQuerySet" exists.
func (this GPUDevice) HasFuncCreateQuerySet() bool {
	return js.True == bindings.HasFuncGPUDeviceCreateQuerySet(
		this.ref,
	)
}

// FuncCreateQuerySet returns the method "GPUDevice.createQuerySet".
func (this GPUDevice) FuncCreateQuerySet() (fn js.Func[func(descriptor GPUQuerySetDescriptor) GPUQuerySet]) {
	bindings.FuncGPUDeviceCreateQuerySet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateQuerySet calls the method "GPUDevice.createQuerySet".
func (this GPUDevice) CreateQuerySet(descriptor GPUQuerySetDescriptor) (ret GPUQuerySet) {
	bindings.CallGPUDeviceCreateQuerySet(
		this.ref, js.Pointer(&ret),
		js.Pointer(&descriptor),
	)

	return
}

// TryCreateQuerySet calls the method "GPUDevice.createQuerySet"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryCreateQuerySet(descriptor GPUQuerySetDescriptor) (ret GPUQuerySet, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDeviceCreateQuerySet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&descriptor),
	)

	return
}

// HasFuncPushErrorScope returns true if the method "GPUDevice.pushErrorScope" exists.
func (this GPUDevice) HasFuncPushErrorScope() bool {
	return js.True == bindings.HasFuncGPUDevicePushErrorScope(
		this.ref,
	)
}

// FuncPushErrorScope returns the method "GPUDevice.pushErrorScope".
func (this GPUDevice) FuncPushErrorScope() (fn js.Func[func(filter GPUErrorFilter)]) {
	bindings.FuncGPUDevicePushErrorScope(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PushErrorScope calls the method "GPUDevice.pushErrorScope".
func (this GPUDevice) PushErrorScope(filter GPUErrorFilter) (ret js.Void) {
	bindings.CallGPUDevicePushErrorScope(
		this.ref, js.Pointer(&ret),
		uint32(filter),
	)

	return
}

// TryPushErrorScope calls the method "GPUDevice.pushErrorScope"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryPushErrorScope(filter GPUErrorFilter) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDevicePushErrorScope(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(filter),
	)

	return
}

// HasFuncPopErrorScope returns true if the method "GPUDevice.popErrorScope" exists.
func (this GPUDevice) HasFuncPopErrorScope() bool {
	return js.True == bindings.HasFuncGPUDevicePopErrorScope(
		this.ref,
	)
}

// FuncPopErrorScope returns the method "GPUDevice.popErrorScope".
func (this GPUDevice) FuncPopErrorScope() (fn js.Func[func() js.Promise[GPUError]]) {
	bindings.FuncGPUDevicePopErrorScope(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PopErrorScope calls the method "GPUDevice.popErrorScope".
func (this GPUDevice) PopErrorScope() (ret js.Promise[GPUError]) {
	bindings.CallGPUDevicePopErrorScope(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPopErrorScope calls the method "GPUDevice.popErrorScope"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUDevice) TryPopErrorScope() (ret js.Promise[GPUError], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUDevicePopErrorScope(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *GPUCanvasConfiguration) UpdateFrom(ref js.Ref) {
	bindings.GPUCanvasConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUCanvasConfiguration) Update(ref js.Ref) {
	bindings.GPUCanvasConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUCanvasConfiguration) FreeMembers(recursive bool) {
	js.Free(
		p.Device.Ref(),
		p.ViewFormats.Ref(),
	)
	p.Device = p.Device.FromRef(js.Undefined)
	p.ViewFormats = p.ViewFormats.FromRef(js.Undefined)
}

type GPUCanvasContext struct {
	ref js.Ref
}

func (this GPUCanvasContext) Once() GPUCanvasContext {
	this.ref.Once()
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
	this.ref.Free()
}

// Canvas returns the value of property "GPUCanvasContext.canvas".
//
// It returns ok=false if there is no such property.
func (this GPUCanvasContext) Canvas() (ret OneOf_HTMLCanvasElement_OffscreenCanvas, ok bool) {
	ok = js.True == bindings.GetGPUCanvasContextCanvas(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncConfigure returns true if the method "GPUCanvasContext.configure" exists.
func (this GPUCanvasContext) HasFuncConfigure() bool {
	return js.True == bindings.HasFuncGPUCanvasContextConfigure(
		this.ref,
	)
}

// FuncConfigure returns the method "GPUCanvasContext.configure".
func (this GPUCanvasContext) FuncConfigure() (fn js.Func[func(configuration GPUCanvasConfiguration)]) {
	bindings.FuncGPUCanvasContextConfigure(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Configure calls the method "GPUCanvasContext.configure".
func (this GPUCanvasContext) Configure(configuration GPUCanvasConfiguration) (ret js.Void) {
	bindings.CallGPUCanvasContextConfigure(
		this.ref, js.Pointer(&ret),
		js.Pointer(&configuration),
	)

	return
}

// TryConfigure calls the method "GPUCanvasContext.configure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCanvasContext) TryConfigure(configuration GPUCanvasConfiguration) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCanvasContextConfigure(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&configuration),
	)

	return
}

// HasFuncUnconfigure returns true if the method "GPUCanvasContext.unconfigure" exists.
func (this GPUCanvasContext) HasFuncUnconfigure() bool {
	return js.True == bindings.HasFuncGPUCanvasContextUnconfigure(
		this.ref,
	)
}

// FuncUnconfigure returns the method "GPUCanvasContext.unconfigure".
func (this GPUCanvasContext) FuncUnconfigure() (fn js.Func[func()]) {
	bindings.FuncGPUCanvasContextUnconfigure(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Unconfigure calls the method "GPUCanvasContext.unconfigure".
func (this GPUCanvasContext) Unconfigure() (ret js.Void) {
	bindings.CallGPUCanvasContextUnconfigure(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryUnconfigure calls the method "GPUCanvasContext.unconfigure"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCanvasContext) TryUnconfigure() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCanvasContextUnconfigure(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetCurrentTexture returns true if the method "GPUCanvasContext.getCurrentTexture" exists.
func (this GPUCanvasContext) HasFuncGetCurrentTexture() bool {
	return js.True == bindings.HasFuncGPUCanvasContextGetCurrentTexture(
		this.ref,
	)
}

// FuncGetCurrentTexture returns the method "GPUCanvasContext.getCurrentTexture".
func (this GPUCanvasContext) FuncGetCurrentTexture() (fn js.Func[func() GPUTexture]) {
	bindings.FuncGPUCanvasContextGetCurrentTexture(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetCurrentTexture calls the method "GPUCanvasContext.getCurrentTexture".
func (this GPUCanvasContext) GetCurrentTexture() (ret GPUTexture) {
	bindings.CallGPUCanvasContextGetCurrentTexture(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetCurrentTexture calls the method "GPUCanvasContext.getCurrentTexture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GPUCanvasContext) TryGetCurrentTexture() (ret GPUTexture, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGPUCanvasContextGetCurrentTexture(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}
