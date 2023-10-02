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
// The returned bool will be false if there is no such property.
func (this GPURenderPassEncoder) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPURenderPassEncoderLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPURenderPassEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPURenderPassEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPURenderPassEncoderLabel(
		this.Ref(),
		val.Ref(),
	)
}

// SetViewport calls the method "GPURenderPassEncoder.setViewport".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetViewport(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetViewport(
		this.Ref(), js.Pointer(&_ok),
		float32(x),
		float32(y),
		float32(width),
		float32(height),
		float32(minDepth),
		float32(maxDepth),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetViewportFunc returns the method "GPURenderPassEncoder.setViewport".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetViewportFunc() (fn js.Func[func(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetViewportFunc(
			this.Ref(),
		),
	)
}

// SetScissorRect calls the method "GPURenderPassEncoder.setScissorRect".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetScissorRect(x GPUIntegerCoordinate, y GPUIntegerCoordinate, width GPUIntegerCoordinate, height GPUIntegerCoordinate) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetScissorRect(
		this.Ref(), js.Pointer(&_ok),
		uint32(x),
		uint32(y),
		uint32(width),
		uint32(height),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetScissorRectFunc returns the method "GPURenderPassEncoder.setScissorRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetScissorRectFunc() (fn js.Func[func(x GPUIntegerCoordinate, y GPUIntegerCoordinate, width GPUIntegerCoordinate, height GPUIntegerCoordinate)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetScissorRectFunc(
			this.Ref(),
		),
	)
}

// SetBlendConstant calls the method "GPURenderPassEncoder.setBlendConstant".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetBlendConstant(color GPUColor) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetBlendConstant(
		this.Ref(), js.Pointer(&_ok),
		color.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBlendConstantFunc returns the method "GPURenderPassEncoder.setBlendConstant".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetBlendConstantFunc() (fn js.Func[func(color GPUColor)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetBlendConstantFunc(
			this.Ref(),
		),
	)
}

// SetStencilReference calls the method "GPURenderPassEncoder.setStencilReference".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetStencilReference(reference GPUStencilValue) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetStencilReference(
		this.Ref(), js.Pointer(&_ok),
		uint32(reference),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetStencilReferenceFunc returns the method "GPURenderPassEncoder.setStencilReference".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetStencilReferenceFunc() (fn js.Func[func(reference GPUStencilValue)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetStencilReferenceFunc(
			this.Ref(),
		),
	)
}

// BeginOcclusionQuery calls the method "GPURenderPassEncoder.beginOcclusionQuery".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) BeginOcclusionQuery(queryIndex GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderBeginOcclusionQuery(
		this.Ref(), js.Pointer(&_ok),
		uint32(queryIndex),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeginOcclusionQueryFunc returns the method "GPURenderPassEncoder.beginOcclusionQuery".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) BeginOcclusionQueryFunc() (fn js.Func[func(queryIndex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderBeginOcclusionQueryFunc(
			this.Ref(),
		),
	)
}

// EndOcclusionQuery calls the method "GPURenderPassEncoder.endOcclusionQuery".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) EndOcclusionQuery() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderEndOcclusionQuery(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EndOcclusionQueryFunc returns the method "GPURenderPassEncoder.endOcclusionQuery".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) EndOcclusionQueryFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderEndOcclusionQueryFunc(
			this.Ref(),
		),
	)
}

// ExecuteBundles calls the method "GPURenderPassEncoder.executeBundles".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) ExecuteBundles(bundles js.Array[GPURenderBundle]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderExecuteBundles(
		this.Ref(), js.Pointer(&_ok),
		bundles.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ExecuteBundlesFunc returns the method "GPURenderPassEncoder.executeBundles".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) ExecuteBundlesFunc() (fn js.Func[func(bundles js.Array[GPURenderBundle])]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderExecuteBundlesFunc(
			this.Ref(),
		),
	)
}

// End calls the method "GPURenderPassEncoder.end".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) End() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderEnd(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EndFunc returns the method "GPURenderPassEncoder.end".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) EndFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderEndFunc(
			this.Ref(),
		),
	)
}

// SetPipeline calls the method "GPURenderPassEncoder.setPipeline".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetPipeline(pipeline GPURenderPipeline) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetPipeline(
		this.Ref(), js.Pointer(&_ok),
		pipeline.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPipelineFunc returns the method "GPURenderPassEncoder.setPipeline".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetPipelineFunc() (fn js.Func[func(pipeline GPURenderPipeline)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetPipelineFunc(
			this.Ref(),
		),
	)
}

// SetIndexBuffer calls the method "GPURenderPassEncoder.setIndexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetIndexBuffer(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetIndexBuffer(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
		float64(size),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetIndexBufferFunc returns the method "GPURenderPassEncoder.setIndexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetIndexBufferFunc() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetIndexBufferFunc(
			this.Ref(),
		),
	)
}

// SetIndexBuffer1 calls the method "GPURenderPassEncoder.setIndexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetIndexBuffer1(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetIndexBuffer1(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetIndexBuffer1Func returns the method "GPURenderPassEncoder.setIndexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetIndexBuffer1Func() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetIndexBuffer1Func(
			this.Ref(),
		),
	)
}

// SetIndexBuffer2 calls the method "GPURenderPassEncoder.setIndexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetIndexBuffer2(buffer GPUBuffer, indexFormat GPUIndexFormat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetIndexBuffer2(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		uint32(indexFormat),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetIndexBuffer2Func returns the method "GPURenderPassEncoder.setIndexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetIndexBuffer2Func() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetIndexBuffer2Func(
			this.Ref(),
		),
	)
}

// SetVertexBuffer calls the method "GPURenderPassEncoder.setVertexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetVertexBuffer(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetVertexBuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetVertexBufferFunc returns the method "GPURenderPassEncoder.setVertexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetVertexBufferFunc() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetVertexBufferFunc(
			this.Ref(),
		),
	)
}

// SetVertexBuffer1 calls the method "GPURenderPassEncoder.setVertexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetVertexBuffer1(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetVertexBuffer1(
		this.Ref(), js.Pointer(&_ok),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetVertexBuffer1Func returns the method "GPURenderPassEncoder.setVertexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetVertexBuffer1Func() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetVertexBuffer1Func(
			this.Ref(),
		),
	)
}

// SetVertexBuffer2 calls the method "GPURenderPassEncoder.setVertexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetVertexBuffer2(slot GPUIndex32, buffer GPUBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetVertexBuffer2(
		this.Ref(), js.Pointer(&_ok),
		uint32(slot),
		buffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetVertexBuffer2Func returns the method "GPURenderPassEncoder.setVertexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetVertexBuffer2Func() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetVertexBuffer2Func(
			this.Ref(),
		),
	)
}

// Draw calls the method "GPURenderPassEncoder.draw".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) Draw(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDraw(
		this.Ref(), js.Pointer(&_ok),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
		uint32(firstInstance),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawFunc returns the method "GPURenderPassEncoder.draw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) DrawFunc() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawFunc(
			this.Ref(),
		),
	)
}

// Draw1 calls the method "GPURenderPassEncoder.draw".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) Draw1(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDraw1(
		this.Ref(), js.Pointer(&_ok),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Draw1Func returns the method "GPURenderPassEncoder.draw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) Draw1Func() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDraw1Func(
			this.Ref(),
		),
	)
}

// Draw2 calls the method "GPURenderPassEncoder.draw".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) Draw2(vertexCount GPUSize32, instanceCount GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDraw2(
		this.Ref(), js.Pointer(&_ok),
		uint32(vertexCount),
		uint32(instanceCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Draw2Func returns the method "GPURenderPassEncoder.draw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) Draw2Func() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDraw2Func(
			this.Ref(),
		),
	)
}

// Draw3 calls the method "GPURenderPassEncoder.draw".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) Draw3(vertexCount GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDraw3(
		this.Ref(), js.Pointer(&_ok),
		uint32(vertexCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Draw3Func returns the method "GPURenderPassEncoder.draw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) Draw3Func() (fn js.Func[func(vertexCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDraw3Func(
			this.Ref(),
		),
	)
}

// DrawIndexed calls the method "GPURenderPassEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) DrawIndexed(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDrawIndexed(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
		uint32(firstInstance),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexedFunc returns the method "GPURenderPassEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) DrawIndexedFunc() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexedFunc(
			this.Ref(),
		),
	)
}

// DrawIndexed1 calls the method "GPURenderPassEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) DrawIndexed1(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDrawIndexed1(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexed1Func returns the method "GPURenderPassEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) DrawIndexed1Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexed1Func(
			this.Ref(),
		),
	)
}

// DrawIndexed2 calls the method "GPURenderPassEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) DrawIndexed2(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDrawIndexed2(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexed2Func returns the method "GPURenderPassEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) DrawIndexed2Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexed2Func(
			this.Ref(),
		),
	)
}

// DrawIndexed3 calls the method "GPURenderPassEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) DrawIndexed3(indexCount GPUSize32, instanceCount GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDrawIndexed3(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
		uint32(instanceCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexed3Func returns the method "GPURenderPassEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) DrawIndexed3Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexed3Func(
			this.Ref(),
		),
	)
}

// DrawIndexed4 calls the method "GPURenderPassEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) DrawIndexed4(indexCount GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDrawIndexed4(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexed4Func returns the method "GPURenderPassEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) DrawIndexed4Func() (fn js.Func[func(indexCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexed4Func(
			this.Ref(),
		),
	)
}

// DrawIndirect calls the method "GPURenderPassEncoder.drawIndirect".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) DrawIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDrawIndirect(
		this.Ref(), js.Pointer(&_ok),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndirectFunc returns the method "GPURenderPassEncoder.drawIndirect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) DrawIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndirectFunc(
			this.Ref(),
		),
	)
}

// DrawIndexedIndirect calls the method "GPURenderPassEncoder.drawIndexedIndirect".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) DrawIndexedIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderDrawIndexedIndirect(
		this.Ref(), js.Pointer(&_ok),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexedIndirectFunc returns the method "GPURenderPassEncoder.drawIndexedIndirect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) DrawIndexedIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderDrawIndexedIndirectFunc(
			this.Ref(),
		),
	)
}

// SetBindGroup calls the method "GPURenderPassEncoder.setBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetBindGroup(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBindGroupFunc returns the method "GPURenderPassEncoder.setBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetBindGroupFunc() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset])]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetBindGroupFunc(
			this.Ref(),
		),
	)
}

// SetBindGroup1 calls the method "GPURenderPassEncoder.setBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetBindGroup1(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		bindGroup.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBindGroup1Func returns the method "GPURenderPassEncoder.setBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetBindGroup1Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetBindGroup1Func(
			this.Ref(),
		),
	)
}

// SetBindGroup2 calls the method "GPURenderPassEncoder.setBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) SetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderSetBindGroup2(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBindGroup2Func returns the method "GPURenderPassEncoder.setBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) SetBindGroup2Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderSetBindGroup2Func(
			this.Ref(),
		),
	)
}

// PushDebugGroup calls the method "GPURenderPassEncoder.pushDebugGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) PushDebugGroup(groupLabel js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&_ok),
		groupLabel.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PushDebugGroupFunc returns the method "GPURenderPassEncoder.pushDebugGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) PushDebugGroupFunc() (fn js.Func[func(groupLabel js.String)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderPushDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PopDebugGroup calls the method "GPURenderPassEncoder.popDebugGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) PopDebugGroup() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PopDebugGroupFunc returns the method "GPURenderPassEncoder.popDebugGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) PopDebugGroupFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderPopDebugGroupFunc(
			this.Ref(),
		),
	)
}

// InsertDebugMarker calls the method "GPURenderPassEncoder.insertDebugMarker".
//
// The returned bool will be false if there is no such method.
func (this GPURenderPassEncoder) InsertDebugMarker(markerLabel js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderPassEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&_ok),
		markerLabel.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertDebugMarkerFunc returns the method "GPURenderPassEncoder.insertDebugMarker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderPassEncoder) InsertDebugMarkerFunc() (fn js.Func[func(markerLabel js.String)]) {
	return fn.FromRef(
		bindings.GPURenderPassEncoderInsertDebugMarkerFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this GPUQuerySet) Type() (GPUQueryType, bool) {
	var _ok bool
	_ret := bindings.GetGPUQuerySetType(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUQueryType(_ret), _ok
}

// Count returns the value of property "GPUQuerySet.count".
//
// The returned bool will be false if there is no such property.
func (this GPUQuerySet) Count() (GPUSize32Out, bool) {
	var _ok bool
	_ret := bindings.GetGPUQuerySetCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUSize32Out(_ret), _ok
}

// Label returns the value of property "GPUQuerySet.label".
//
// The returned bool will be false if there is no such property.
func (this GPUQuerySet) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUQuerySetLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUQuerySet.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUQuerySet) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUQuerySetLabel(
		this.Ref(),
		val.Ref(),
	)
}

// Destroy calls the method "GPUQuerySet.destroy".
//
// The returned bool will be false if there is no such method.
func (this GPUQuerySet) Destroy() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUQuerySetDestroy(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DestroyFunc returns the method "GPUQuerySet.destroy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUQuerySet) DestroyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUQuerySetDestroyFunc(
			this.Ref(),
		),
	)
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
	DepthStencilAttachment GPURenderPassDepthStencilAttachment
	// OcclusionQuerySet is "GPURenderPassDescriptor.occlusionQuerySet"
	//
	// Optional
	OcclusionQuerySet GPUQuerySet
	// TimestampWrites is "GPURenderPassDescriptor.timestampWrites"
	//
	// Optional
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
// The returned bool will be false if there is no such property.
func (this GPUComputePassEncoder) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUComputePassEncoderLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUComputePassEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUComputePassEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUComputePassEncoderLabel(
		this.Ref(),
		val.Ref(),
	)
}

// SetPipeline calls the method "GPUComputePassEncoder.setPipeline".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) SetPipeline(pipeline GPUComputePipeline) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderSetPipeline(
		this.Ref(), js.Pointer(&_ok),
		pipeline.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPipelineFunc returns the method "GPUComputePassEncoder.setPipeline".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) SetPipelineFunc() (fn js.Func[func(pipeline GPUComputePipeline)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderSetPipelineFunc(
			this.Ref(),
		),
	)
}

// DispatchWorkgroups calls the method "GPUComputePassEncoder.dispatchWorkgroups".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) DispatchWorkgroups(workgroupCountX GPUSize32, workgroupCountY GPUSize32, workgroupCountZ GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderDispatchWorkgroups(
		this.Ref(), js.Pointer(&_ok),
		uint32(workgroupCountX),
		uint32(workgroupCountY),
		uint32(workgroupCountZ),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DispatchWorkgroupsFunc returns the method "GPUComputePassEncoder.dispatchWorkgroups".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) DispatchWorkgroupsFunc() (fn js.Func[func(workgroupCountX GPUSize32, workgroupCountY GPUSize32, workgroupCountZ GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderDispatchWorkgroupsFunc(
			this.Ref(),
		),
	)
}

// DispatchWorkgroups1 calls the method "GPUComputePassEncoder.dispatchWorkgroups".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) DispatchWorkgroups1(workgroupCountX GPUSize32, workgroupCountY GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderDispatchWorkgroups1(
		this.Ref(), js.Pointer(&_ok),
		uint32(workgroupCountX),
		uint32(workgroupCountY),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DispatchWorkgroups1Func returns the method "GPUComputePassEncoder.dispatchWorkgroups".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) DispatchWorkgroups1Func() (fn js.Func[func(workgroupCountX GPUSize32, workgroupCountY GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderDispatchWorkgroups1Func(
			this.Ref(),
		),
	)
}

// DispatchWorkgroups2 calls the method "GPUComputePassEncoder.dispatchWorkgroups".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) DispatchWorkgroups2(workgroupCountX GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderDispatchWorkgroups2(
		this.Ref(), js.Pointer(&_ok),
		uint32(workgroupCountX),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DispatchWorkgroups2Func returns the method "GPUComputePassEncoder.dispatchWorkgroups".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) DispatchWorkgroups2Func() (fn js.Func[func(workgroupCountX GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderDispatchWorkgroups2Func(
			this.Ref(),
		),
	)
}

// DispatchWorkgroupsIndirect calls the method "GPUComputePassEncoder.dispatchWorkgroupsIndirect".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderDispatchWorkgroupsIndirect(
		this.Ref(), js.Pointer(&_ok),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DispatchWorkgroupsIndirectFunc returns the method "GPUComputePassEncoder.dispatchWorkgroupsIndirect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) DispatchWorkgroupsIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderDispatchWorkgroupsIndirectFunc(
			this.Ref(),
		),
	)
}

// End calls the method "GPUComputePassEncoder.end".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) End() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderEnd(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EndFunc returns the method "GPUComputePassEncoder.end".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) EndFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderEndFunc(
			this.Ref(),
		),
	)
}

// PushDebugGroup calls the method "GPUComputePassEncoder.pushDebugGroup".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) PushDebugGroup(groupLabel js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&_ok),
		groupLabel.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PushDebugGroupFunc returns the method "GPUComputePassEncoder.pushDebugGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) PushDebugGroupFunc() (fn js.Func[func(groupLabel js.String)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderPushDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PopDebugGroup calls the method "GPUComputePassEncoder.popDebugGroup".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) PopDebugGroup() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PopDebugGroupFunc returns the method "GPUComputePassEncoder.popDebugGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) PopDebugGroupFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderPopDebugGroupFunc(
			this.Ref(),
		),
	)
}

// InsertDebugMarker calls the method "GPUComputePassEncoder.insertDebugMarker".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) InsertDebugMarker(markerLabel js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&_ok),
		markerLabel.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertDebugMarkerFunc returns the method "GPUComputePassEncoder.insertDebugMarker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) InsertDebugMarkerFunc() (fn js.Func[func(markerLabel js.String)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderInsertDebugMarkerFunc(
			this.Ref(),
		),
	)
}

// SetBindGroup calls the method "GPUComputePassEncoder.setBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderSetBindGroup(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBindGroupFunc returns the method "GPUComputePassEncoder.setBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) SetBindGroupFunc() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset])]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderSetBindGroupFunc(
			this.Ref(),
		),
	)
}

// SetBindGroup1 calls the method "GPUComputePassEncoder.setBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) SetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderSetBindGroup1(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		bindGroup.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBindGroup1Func returns the method "GPUComputePassEncoder.setBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) SetBindGroup1Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderSetBindGroup1Func(
			this.Ref(),
		),
	)
}

// SetBindGroup2 calls the method "GPUComputePassEncoder.setBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPUComputePassEncoder) SetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUComputePassEncoderSetBindGroup2(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBindGroup2Func returns the method "GPUComputePassEncoder.setBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUComputePassEncoder) SetBindGroup2Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUComputePassEncoderSetBindGroup2Func(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this GPUCommandBuffer) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUCommandBufferLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUCommandBuffer.label" to val.
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
// The returned bool will be false if there is no such property.
func (this GPUCommandEncoder) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUCommandEncoderLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUCommandEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUCommandEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUCommandEncoderLabel(
		this.Ref(),
		val.Ref(),
	)
}

// BeginRenderPass calls the method "GPUCommandEncoder.beginRenderPass".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) BeginRenderPass(descriptor GPURenderPassDescriptor) (GPURenderPassEncoder, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderBeginRenderPass(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPURenderPassEncoder{}.FromRef(_ret), _ok
}

// BeginRenderPassFunc returns the method "GPUCommandEncoder.beginRenderPass".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) BeginRenderPassFunc() (fn js.Func[func(descriptor GPURenderPassDescriptor) GPURenderPassEncoder]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderBeginRenderPassFunc(
			this.Ref(),
		),
	)
}

// BeginComputePass calls the method "GPUCommandEncoder.beginComputePass".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) BeginComputePass(descriptor GPUComputePassDescriptor) (GPUComputePassEncoder, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderBeginComputePass(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUComputePassEncoder{}.FromRef(_ret), _ok
}

// BeginComputePassFunc returns the method "GPUCommandEncoder.beginComputePass".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) BeginComputePassFunc() (fn js.Func[func(descriptor GPUComputePassDescriptor) GPUComputePassEncoder]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderBeginComputePassFunc(
			this.Ref(),
		),
	)
}

// BeginComputePass1 calls the method "GPUCommandEncoder.beginComputePass".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) BeginComputePass1() (GPUComputePassEncoder, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderBeginComputePass1(
		this.Ref(), js.Pointer(&_ok),
	)

	return GPUComputePassEncoder{}.FromRef(_ret), _ok
}

// BeginComputePass1Func returns the method "GPUCommandEncoder.beginComputePass".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) BeginComputePass1Func() (fn js.Func[func() GPUComputePassEncoder]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderBeginComputePass1Func(
			this.Ref(),
		),
	)
}

// CopyBufferToBuffer calls the method "GPUCommandEncoder.copyBufferToBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) CopyBufferToBuffer(source GPUBuffer, sourceOffset GPUSize64, destination GPUBuffer, destinationOffset GPUSize64, size GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderCopyBufferToBuffer(
		this.Ref(), js.Pointer(&_ok),
		source.Ref(),
		float64(sourceOffset),
		destination.Ref(),
		float64(destinationOffset),
		float64(size),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyBufferToBufferFunc returns the method "GPUCommandEncoder.copyBufferToBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) CopyBufferToBufferFunc() (fn js.Func[func(source GPUBuffer, sourceOffset GPUSize64, destination GPUBuffer, destinationOffset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderCopyBufferToBufferFunc(
			this.Ref(),
		),
	)
}

// CopyBufferToTexture calls the method "GPUCommandEncoder.copyBufferToTexture".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) CopyBufferToTexture(source GPUImageCopyBuffer, destination GPUImageCopyTexture, copySize GPUExtent3D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderCopyBufferToTexture(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyBufferToTextureFunc returns the method "GPUCommandEncoder.copyBufferToTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) CopyBufferToTextureFunc() (fn js.Func[func(source GPUImageCopyBuffer, destination GPUImageCopyTexture, copySize GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderCopyBufferToTextureFunc(
			this.Ref(),
		),
	)
}

// CopyTextureToBuffer calls the method "GPUCommandEncoder.copyTextureToBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) CopyTextureToBuffer(source GPUImageCopyTexture, destination GPUImageCopyBuffer, copySize GPUExtent3D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderCopyTextureToBuffer(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyTextureToBufferFunc returns the method "GPUCommandEncoder.copyTextureToBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) CopyTextureToBufferFunc() (fn js.Func[func(source GPUImageCopyTexture, destination GPUImageCopyBuffer, copySize GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderCopyTextureToBufferFunc(
			this.Ref(),
		),
	)
}

// CopyTextureToTexture calls the method "GPUCommandEncoder.copyTextureToTexture".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) CopyTextureToTexture(source GPUImageCopyTexture, destination GPUImageCopyTexture, copySize GPUExtent3D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderCopyTextureToTexture(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyTextureToTextureFunc returns the method "GPUCommandEncoder.copyTextureToTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) CopyTextureToTextureFunc() (fn js.Func[func(source GPUImageCopyTexture, destination GPUImageCopyTexture, copySize GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderCopyTextureToTextureFunc(
			this.Ref(),
		),
	)
}

// ClearBuffer calls the method "GPUCommandEncoder.clearBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) ClearBuffer(buffer GPUBuffer, offset GPUSize64, size GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderClearBuffer(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBufferFunc returns the method "GPUCommandEncoder.clearBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) ClearBufferFunc() (fn js.Func[func(buffer GPUBuffer, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderClearBufferFunc(
			this.Ref(),
		),
	)
}

// ClearBuffer1 calls the method "GPUCommandEncoder.clearBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) ClearBuffer1(buffer GPUBuffer, offset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderClearBuffer1(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBuffer1Func returns the method "GPUCommandEncoder.clearBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) ClearBuffer1Func() (fn js.Func[func(buffer GPUBuffer, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderClearBuffer1Func(
			this.Ref(),
		),
	)
}

// ClearBuffer2 calls the method "GPUCommandEncoder.clearBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) ClearBuffer2(buffer GPUBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderClearBuffer2(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearBuffer2Func returns the method "GPUCommandEncoder.clearBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) ClearBuffer2Func() (fn js.Func[func(buffer GPUBuffer)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderClearBuffer2Func(
			this.Ref(),
		),
	)
}

// WriteTimestamp calls the method "GPUCommandEncoder.writeTimestamp".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) WriteTimestamp(querySet GPUQuerySet, queryIndex GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderWriteTimestamp(
		this.Ref(), js.Pointer(&_ok),
		querySet.Ref(),
		uint32(queryIndex),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WriteTimestampFunc returns the method "GPUCommandEncoder.writeTimestamp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) WriteTimestampFunc() (fn js.Func[func(querySet GPUQuerySet, queryIndex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderWriteTimestampFunc(
			this.Ref(),
		),
	)
}

// ResolveQuerySet calls the method "GPUCommandEncoder.resolveQuerySet".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) ResolveQuerySet(querySet GPUQuerySet, firstQuery GPUSize32, queryCount GPUSize32, destination GPUBuffer, destinationOffset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderResolveQuerySet(
		this.Ref(), js.Pointer(&_ok),
		querySet.Ref(),
		uint32(firstQuery),
		uint32(queryCount),
		destination.Ref(),
		float64(destinationOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResolveQuerySetFunc returns the method "GPUCommandEncoder.resolveQuerySet".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) ResolveQuerySetFunc() (fn js.Func[func(querySet GPUQuerySet, firstQuery GPUSize32, queryCount GPUSize32, destination GPUBuffer, destinationOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderResolveQuerySetFunc(
			this.Ref(),
		),
	)
}

// Finish calls the method "GPUCommandEncoder.finish".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) Finish(descriptor GPUCommandBufferDescriptor) (GPUCommandBuffer, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderFinish(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUCommandBuffer{}.FromRef(_ret), _ok
}

// FinishFunc returns the method "GPUCommandEncoder.finish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) FinishFunc() (fn js.Func[func(descriptor GPUCommandBufferDescriptor) GPUCommandBuffer]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderFinishFunc(
			this.Ref(),
		),
	)
}

// Finish1 calls the method "GPUCommandEncoder.finish".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) Finish1() (GPUCommandBuffer, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderFinish1(
		this.Ref(), js.Pointer(&_ok),
	)

	return GPUCommandBuffer{}.FromRef(_ret), _ok
}

// Finish1Func returns the method "GPUCommandEncoder.finish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) Finish1Func() (fn js.Func[func() GPUCommandBuffer]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderFinish1Func(
			this.Ref(),
		),
	)
}

// PushDebugGroup calls the method "GPUCommandEncoder.pushDebugGroup".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) PushDebugGroup(groupLabel js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&_ok),
		groupLabel.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PushDebugGroupFunc returns the method "GPUCommandEncoder.pushDebugGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) PushDebugGroupFunc() (fn js.Func[func(groupLabel js.String)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderPushDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PopDebugGroup calls the method "GPUCommandEncoder.popDebugGroup".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) PopDebugGroup() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PopDebugGroupFunc returns the method "GPUCommandEncoder.popDebugGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) PopDebugGroupFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderPopDebugGroupFunc(
			this.Ref(),
		),
	)
}

// InsertDebugMarker calls the method "GPUCommandEncoder.insertDebugMarker".
//
// The returned bool will be false if there is no such method.
func (this GPUCommandEncoder) InsertDebugMarker(markerLabel js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCommandEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&_ok),
		markerLabel.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertDebugMarkerFunc returns the method "GPUCommandEncoder.insertDebugMarker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCommandEncoder) InsertDebugMarkerFunc() (fn js.Func[func(markerLabel js.String)]) {
	return fn.FromRef(
		bindings.GPUCommandEncoderInsertDebugMarkerFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this GPURenderBundleEncoder) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPURenderBundleEncoderLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPURenderBundleEncoder.label" to val.
//
// It returns false if the property cannot be set.
func (this GPURenderBundleEncoder) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPURenderBundleEncoderLabel(
		this.Ref(),
		val.Ref(),
	)
}

// Finish calls the method "GPURenderBundleEncoder.finish".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) Finish(descriptor GPURenderBundleDescriptor) (GPURenderBundle, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderFinish(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPURenderBundle{}.FromRef(_ret), _ok
}

// FinishFunc returns the method "GPURenderBundleEncoder.finish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) FinishFunc() (fn js.Func[func(descriptor GPURenderBundleDescriptor) GPURenderBundle]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderFinishFunc(
			this.Ref(),
		),
	)
}

// Finish1 calls the method "GPURenderBundleEncoder.finish".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) Finish1() (GPURenderBundle, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderFinish1(
		this.Ref(), js.Pointer(&_ok),
	)

	return GPURenderBundle{}.FromRef(_ret), _ok
}

// Finish1Func returns the method "GPURenderBundleEncoder.finish".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) Finish1Func() (fn js.Func[func() GPURenderBundle]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderFinish1Func(
			this.Ref(),
		),
	)
}

// SetPipeline calls the method "GPURenderBundleEncoder.setPipeline".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetPipeline(pipeline GPURenderPipeline) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetPipeline(
		this.Ref(), js.Pointer(&_ok),
		pipeline.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPipelineFunc returns the method "GPURenderBundleEncoder.setPipeline".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetPipelineFunc() (fn js.Func[func(pipeline GPURenderPipeline)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetPipelineFunc(
			this.Ref(),
		),
	)
}

// SetIndexBuffer calls the method "GPURenderBundleEncoder.setIndexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetIndexBuffer(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetIndexBuffer(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
		float64(size),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetIndexBufferFunc returns the method "GPURenderBundleEncoder.setIndexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetIndexBufferFunc() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetIndexBufferFunc(
			this.Ref(),
		),
	)
}

// SetIndexBuffer1 calls the method "GPURenderBundleEncoder.setIndexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetIndexBuffer1(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetIndexBuffer1(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		uint32(indexFormat),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetIndexBuffer1Func returns the method "GPURenderBundleEncoder.setIndexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetIndexBuffer1Func() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetIndexBuffer1Func(
			this.Ref(),
		),
	)
}

// SetIndexBuffer2 calls the method "GPURenderBundleEncoder.setIndexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetIndexBuffer2(buffer GPUBuffer, indexFormat GPUIndexFormat) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetIndexBuffer2(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		uint32(indexFormat),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetIndexBuffer2Func returns the method "GPURenderBundleEncoder.setIndexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetIndexBuffer2Func() (fn js.Func[func(buffer GPUBuffer, indexFormat GPUIndexFormat)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetIndexBuffer2Func(
			this.Ref(),
		),
	)
}

// SetVertexBuffer calls the method "GPURenderBundleEncoder.setVertexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetVertexBuffer(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetVertexBuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
		float64(size),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetVertexBufferFunc returns the method "GPURenderBundleEncoder.setVertexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetVertexBufferFunc() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetVertexBufferFunc(
			this.Ref(),
		),
	)
}

// SetVertexBuffer1 calls the method "GPURenderBundleEncoder.setVertexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetVertexBuffer1(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetVertexBuffer1(
		this.Ref(), js.Pointer(&_ok),
		uint32(slot),
		buffer.Ref(),
		float64(offset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetVertexBuffer1Func returns the method "GPURenderBundleEncoder.setVertexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetVertexBuffer1Func() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer, offset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetVertexBuffer1Func(
			this.Ref(),
		),
	)
}

// SetVertexBuffer2 calls the method "GPURenderBundleEncoder.setVertexBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetVertexBuffer2(slot GPUIndex32, buffer GPUBuffer) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetVertexBuffer2(
		this.Ref(), js.Pointer(&_ok),
		uint32(slot),
		buffer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetVertexBuffer2Func returns the method "GPURenderBundleEncoder.setVertexBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetVertexBuffer2Func() (fn js.Func[func(slot GPUIndex32, buffer GPUBuffer)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetVertexBuffer2Func(
			this.Ref(),
		),
	)
}

// Draw calls the method "GPURenderBundleEncoder.draw".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) Draw(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDraw(
		this.Ref(), js.Pointer(&_ok),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
		uint32(firstInstance),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawFunc returns the method "GPURenderBundleEncoder.draw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) DrawFunc() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32, firstInstance GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawFunc(
			this.Ref(),
		),
	)
}

// Draw1 calls the method "GPURenderBundleEncoder.draw".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) Draw1(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDraw1(
		this.Ref(), js.Pointer(&_ok),
		uint32(vertexCount),
		uint32(instanceCount),
		uint32(firstVertex),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Draw1Func returns the method "GPURenderBundleEncoder.draw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) Draw1Func() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32, firstVertex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDraw1Func(
			this.Ref(),
		),
	)
}

// Draw2 calls the method "GPURenderBundleEncoder.draw".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) Draw2(vertexCount GPUSize32, instanceCount GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDraw2(
		this.Ref(), js.Pointer(&_ok),
		uint32(vertexCount),
		uint32(instanceCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Draw2Func returns the method "GPURenderBundleEncoder.draw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) Draw2Func() (fn js.Func[func(vertexCount GPUSize32, instanceCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDraw2Func(
			this.Ref(),
		),
	)
}

// Draw3 calls the method "GPURenderBundleEncoder.draw".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) Draw3(vertexCount GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDraw3(
		this.Ref(), js.Pointer(&_ok),
		uint32(vertexCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Draw3Func returns the method "GPURenderBundleEncoder.draw".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) Draw3Func() (fn js.Func[func(vertexCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDraw3Func(
			this.Ref(),
		),
	)
}

// DrawIndexed calls the method "GPURenderBundleEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexed(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDrawIndexed(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
		uint32(firstInstance),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexedFunc returns the method "GPURenderBundleEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexedFunc() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32, firstInstance GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexedFunc(
			this.Ref(),
		),
	)
}

// DrawIndexed1 calls the method "GPURenderBundleEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexed1(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDrawIndexed1(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
		int32(baseVertex),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexed1Func returns the method "GPURenderBundleEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexed1Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32, baseVertex GPUSignedOffset32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexed1Func(
			this.Ref(),
		),
	)
}

// DrawIndexed2 calls the method "GPURenderBundleEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexed2(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDrawIndexed2(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
		uint32(instanceCount),
		uint32(firstIndex),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexed2Func returns the method "GPURenderBundleEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexed2Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32, firstIndex GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexed2Func(
			this.Ref(),
		),
	)
}

// DrawIndexed3 calls the method "GPURenderBundleEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexed3(indexCount GPUSize32, instanceCount GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDrawIndexed3(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
		uint32(instanceCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexed3Func returns the method "GPURenderBundleEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexed3Func() (fn js.Func[func(indexCount GPUSize32, instanceCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexed3Func(
			this.Ref(),
		),
	)
}

// DrawIndexed4 calls the method "GPURenderBundleEncoder.drawIndexed".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexed4(indexCount GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDrawIndexed4(
		this.Ref(), js.Pointer(&_ok),
		uint32(indexCount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexed4Func returns the method "GPURenderBundleEncoder.drawIndexed".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexed4Func() (fn js.Func[func(indexCount GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexed4Func(
			this.Ref(),
		),
	)
}

// DrawIndirect calls the method "GPURenderBundleEncoder.drawIndirect".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) DrawIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDrawIndirect(
		this.Ref(), js.Pointer(&_ok),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndirectFunc returns the method "GPURenderBundleEncoder.drawIndirect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) DrawIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndirectFunc(
			this.Ref(),
		),
	)
}

// DrawIndexedIndirect calls the method "GPURenderBundleEncoder.drawIndexedIndirect".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexedIndirect(indirectBuffer GPUBuffer, indirectOffset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderDrawIndexedIndirect(
		this.Ref(), js.Pointer(&_ok),
		indirectBuffer.Ref(),
		float64(indirectOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawIndexedIndirectFunc returns the method "GPURenderBundleEncoder.drawIndexedIndirect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) DrawIndexedIndirectFunc() (fn js.Func[func(indirectBuffer GPUBuffer, indirectOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderDrawIndexedIndirectFunc(
			this.Ref(),
		),
	)
}

// SetBindGroup calls the method "GPURenderBundleEncoder.setBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetBindGroup(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetBindGroup(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsets.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBindGroupFunc returns the method "GPURenderBundleEncoder.setBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetBindGroupFunc() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsets js.Array[GPUBufferDynamicOffset])]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetBindGroupFunc(
			this.Ref(),
		),
	)
}

// SetBindGroup1 calls the method "GPURenderBundleEncoder.setBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetBindGroup1(index GPUIndex32, bindGroup GPUBindGroup) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetBindGroup1(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		bindGroup.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBindGroup1Func returns the method "GPURenderBundleEncoder.setBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetBindGroup1Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetBindGroup1Func(
			this.Ref(),
		),
	)
}

// SetBindGroup2 calls the method "GPURenderBundleEncoder.setBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) SetBindGroup2(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderSetBindGroup2(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		bindGroup.Ref(),
		dynamicOffsetsData.Ref(),
		float64(dynamicOffsetsDataStart),
		uint32(dynamicOffsetsDataLength),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetBindGroup2Func returns the method "GPURenderBundleEncoder.setBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) SetBindGroup2Func() (fn js.Func[func(index GPUIndex32, bindGroup GPUBindGroup, dynamicOffsetsData js.TypedArray[uint32], dynamicOffsetsDataStart GPUSize64, dynamicOffsetsDataLength GPUSize32)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderSetBindGroup2Func(
			this.Ref(),
		),
	)
}

// PushDebugGroup calls the method "GPURenderBundleEncoder.pushDebugGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) PushDebugGroup(groupLabel js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderPushDebugGroup(
		this.Ref(), js.Pointer(&_ok),
		groupLabel.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PushDebugGroupFunc returns the method "GPURenderBundleEncoder.pushDebugGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) PushDebugGroupFunc() (fn js.Func[func(groupLabel js.String)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderPushDebugGroupFunc(
			this.Ref(),
		),
	)
}

// PopDebugGroup calls the method "GPURenderBundleEncoder.popDebugGroup".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) PopDebugGroup() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderPopDebugGroup(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PopDebugGroupFunc returns the method "GPURenderBundleEncoder.popDebugGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) PopDebugGroupFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderPopDebugGroupFunc(
			this.Ref(),
		),
	)
}

// InsertDebugMarker calls the method "GPURenderBundleEncoder.insertDebugMarker".
//
// The returned bool will be false if there is no such method.
func (this GPURenderBundleEncoder) InsertDebugMarker(markerLabel js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPURenderBundleEncoderInsertDebugMarker(
		this.Ref(), js.Pointer(&_ok),
		markerLabel.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InsertDebugMarkerFunc returns the method "GPURenderBundleEncoder.insertDebugMarker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPURenderBundleEncoder) InsertDebugMarkerFunc() (fn js.Func[func(markerLabel js.String)]) {
	return fn.FromRef(
		bindings.GPURenderBundleEncoderInsertDebugMarkerFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this GPUError) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUErrorMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxTextureDimension1D() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxTextureDimension1D(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxTextureDimension2D returns the value of property "GPUSupportedLimits.maxTextureDimension2D".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxTextureDimension2D() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxTextureDimension2D(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxTextureDimension3D returns the value of property "GPUSupportedLimits.maxTextureDimension3D".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxTextureDimension3D() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxTextureDimension3D(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxTextureArrayLayers returns the value of property "GPUSupportedLimits.maxTextureArrayLayers".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxTextureArrayLayers() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxTextureArrayLayers(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxBindGroups returns the value of property "GPUSupportedLimits.maxBindGroups".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxBindGroups() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxBindGroups(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxBindGroupsPlusVertexBuffers returns the value of property "GPUSupportedLimits.maxBindGroupsPlusVertexBuffers".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxBindGroupsPlusVertexBuffers() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxBindGroupsPlusVertexBuffers(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxBindingsPerBindGroup returns the value of property "GPUSupportedLimits.maxBindingsPerBindGroup".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxBindingsPerBindGroup() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxBindingsPerBindGroup(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxDynamicUniformBuffersPerPipelineLayout returns the value of property "GPUSupportedLimits.maxDynamicUniformBuffersPerPipelineLayout".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxDynamicUniformBuffersPerPipelineLayout() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxDynamicUniformBuffersPerPipelineLayout(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxDynamicStorageBuffersPerPipelineLayout returns the value of property "GPUSupportedLimits.maxDynamicStorageBuffersPerPipelineLayout".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxDynamicStorageBuffersPerPipelineLayout() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxDynamicStorageBuffersPerPipelineLayout(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxSampledTexturesPerShaderStage returns the value of property "GPUSupportedLimits.maxSampledTexturesPerShaderStage".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxSampledTexturesPerShaderStage() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxSampledTexturesPerShaderStage(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxSamplersPerShaderStage returns the value of property "GPUSupportedLimits.maxSamplersPerShaderStage".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxSamplersPerShaderStage() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxSamplersPerShaderStage(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxStorageBuffersPerShaderStage returns the value of property "GPUSupportedLimits.maxStorageBuffersPerShaderStage".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxStorageBuffersPerShaderStage() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxStorageBuffersPerShaderStage(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxStorageTexturesPerShaderStage returns the value of property "GPUSupportedLimits.maxStorageTexturesPerShaderStage".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxStorageTexturesPerShaderStage() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxStorageTexturesPerShaderStage(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxUniformBuffersPerShaderStage returns the value of property "GPUSupportedLimits.maxUniformBuffersPerShaderStage".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxUniformBuffersPerShaderStage() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxUniformBuffersPerShaderStage(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxUniformBufferBindingSize returns the value of property "GPUSupportedLimits.maxUniformBufferBindingSize".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxUniformBufferBindingSize() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxUniformBufferBindingSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// MaxStorageBufferBindingSize returns the value of property "GPUSupportedLimits.maxStorageBufferBindingSize".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxStorageBufferBindingSize() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxStorageBufferBindingSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// MinUniformBufferOffsetAlignment returns the value of property "GPUSupportedLimits.minUniformBufferOffsetAlignment".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MinUniformBufferOffsetAlignment() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMinUniformBufferOffsetAlignment(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MinStorageBufferOffsetAlignment returns the value of property "GPUSupportedLimits.minStorageBufferOffsetAlignment".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MinStorageBufferOffsetAlignment() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMinStorageBufferOffsetAlignment(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxVertexBuffers returns the value of property "GPUSupportedLimits.maxVertexBuffers".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxVertexBuffers() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxVertexBuffers(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxBufferSize returns the value of property "GPUSupportedLimits.maxBufferSize".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxBufferSize() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxBufferSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// MaxVertexAttributes returns the value of property "GPUSupportedLimits.maxVertexAttributes".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxVertexAttributes() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxVertexAttributes(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxVertexBufferArrayStride returns the value of property "GPUSupportedLimits.maxVertexBufferArrayStride".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxVertexBufferArrayStride() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxVertexBufferArrayStride(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxInterStageShaderComponents returns the value of property "GPUSupportedLimits.maxInterStageShaderComponents".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxInterStageShaderComponents() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxInterStageShaderComponents(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxInterStageShaderVariables returns the value of property "GPUSupportedLimits.maxInterStageShaderVariables".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxInterStageShaderVariables() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxInterStageShaderVariables(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxColorAttachments returns the value of property "GPUSupportedLimits.maxColorAttachments".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxColorAttachments() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxColorAttachments(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxColorAttachmentBytesPerSample returns the value of property "GPUSupportedLimits.maxColorAttachmentBytesPerSample".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxColorAttachmentBytesPerSample() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxColorAttachmentBytesPerSample(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxComputeWorkgroupStorageSize returns the value of property "GPUSupportedLimits.maxComputeWorkgroupStorageSize".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupStorageSize() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxComputeWorkgroupStorageSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxComputeInvocationsPerWorkgroup returns the value of property "GPUSupportedLimits.maxComputeInvocationsPerWorkgroup".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxComputeInvocationsPerWorkgroup() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxComputeInvocationsPerWorkgroup(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxComputeWorkgroupSizeX returns the value of property "GPUSupportedLimits.maxComputeWorkgroupSizeX".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupSizeX() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxComputeWorkgroupSizeX(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxComputeWorkgroupSizeY returns the value of property "GPUSupportedLimits.maxComputeWorkgroupSizeY".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupSizeY() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxComputeWorkgroupSizeY(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxComputeWorkgroupSizeZ returns the value of property "GPUSupportedLimits.maxComputeWorkgroupSizeZ".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupSizeZ() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxComputeWorkgroupSizeZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxComputeWorkgroupsPerDimension returns the value of property "GPUSupportedLimits.maxComputeWorkgroupsPerDimension".
//
// The returned bool will be false if there is no such property.
func (this GPUSupportedLimits) MaxComputeWorkgroupsPerDimension() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGPUSupportedLimitsMaxComputeWorkgroupsPerDimension(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this GPUQueue) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUQueueLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUQueue.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUQueue) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUQueueLabel(
		this.Ref(),
		val.Ref(),
	)
}

// Submit calls the method "GPUQueue.submit".
//
// The returned bool will be false if there is no such method.
func (this GPUQueue) Submit(commandBuffers js.Array[GPUCommandBuffer]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUQueueSubmit(
		this.Ref(), js.Pointer(&_ok),
		commandBuffers.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SubmitFunc returns the method "GPUQueue.submit".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUQueue) SubmitFunc() (fn js.Func[func(commandBuffers js.Array[GPUCommandBuffer])]) {
	return fn.FromRef(
		bindings.GPUQueueSubmitFunc(
			this.Ref(),
		),
	)
}

// OnSubmittedWorkDone calls the method "GPUQueue.onSubmittedWorkDone".
//
// The returned bool will be false if there is no such method.
func (this GPUQueue) OnSubmittedWorkDone() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallGPUQueueOnSubmittedWorkDone(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// OnSubmittedWorkDoneFunc returns the method "GPUQueue.onSubmittedWorkDone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUQueue) OnSubmittedWorkDoneFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.GPUQueueOnSubmittedWorkDoneFunc(
			this.Ref(),
		),
	)
}

// WriteBuffer calls the method "GPUQueue.writeBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPUQueue) WriteBuffer(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64, size GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUQueueWriteBuffer(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
		float64(dataOffset),
		float64(size),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WriteBufferFunc returns the method "GPUQueue.writeBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUQueue) WriteBufferFunc() (fn js.Func[func(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64, size GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUQueueWriteBufferFunc(
			this.Ref(),
		),
	)
}

// WriteBuffer1 calls the method "GPUQueue.writeBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPUQueue) WriteBuffer1(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUQueueWriteBuffer1(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
		float64(dataOffset),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WriteBuffer1Func returns the method "GPUQueue.writeBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUQueue) WriteBuffer1Func() (fn js.Func[func(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource, dataOffset GPUSize64)]) {
	return fn.FromRef(
		bindings.GPUQueueWriteBuffer1Func(
			this.Ref(),
		),
	)
}

// WriteBuffer2 calls the method "GPUQueue.writeBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPUQueue) WriteBuffer2(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUQueueWriteBuffer2(
		this.Ref(), js.Pointer(&_ok),
		buffer.Ref(),
		float64(bufferOffset),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WriteBuffer2Func returns the method "GPUQueue.writeBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUQueue) WriteBuffer2Func() (fn js.Func[func(buffer GPUBuffer, bufferOffset GPUSize64, data AllowSharedBufferSource)]) {
	return fn.FromRef(
		bindings.GPUQueueWriteBuffer2Func(
			this.Ref(),
		),
	)
}

// WriteTexture calls the method "GPUQueue.writeTexture".
//
// The returned bool will be false if there is no such method.
func (this GPUQueue) WriteTexture(destination GPUImageCopyTexture, data AllowSharedBufferSource, dataLayout GPUImageDataLayout, size GPUExtent3D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUQueueWriteTexture(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&destination),
		data.Ref(),
		js.Pointer(&dataLayout),
		size.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WriteTextureFunc returns the method "GPUQueue.writeTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUQueue) WriteTextureFunc() (fn js.Func[func(destination GPUImageCopyTexture, data AllowSharedBufferSource, dataLayout GPUImageDataLayout, size GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUQueueWriteTextureFunc(
			this.Ref(),
		),
	)
}

// CopyExternalImageToTexture calls the method "GPUQueue.copyExternalImageToTexture".
//
// The returned bool will be false if there is no such method.
func (this GPUQueue) CopyExternalImageToTexture(source GPUImageCopyExternalImage, destination GPUImageCopyTextureTagged, copySize GPUExtent3D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUQueueCopyExternalImageToTexture(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&source),
		js.Pointer(&destination),
		copySize.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CopyExternalImageToTextureFunc returns the method "GPUQueue.copyExternalImageToTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUQueue) CopyExternalImageToTextureFunc() (fn js.Func[func(source GPUImageCopyExternalImage, destination GPUImageCopyTextureTagged, copySize GPUExtent3D)]) {
	return fn.FromRef(
		bindings.GPUQueueCopyExternalImageToTextureFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this GPUDeviceLostInfo) Reason() (GPUDeviceLostReason, bool) {
	var _ok bool
	_ret := bindings.GetGPUDeviceLostInfoReason(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUDeviceLostReason(_ret), _ok
}

// Message returns the value of property "GPUDeviceLostInfo.message".
//
// The returned bool will be false if there is no such property.
func (this GPUDeviceLostInfo) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUDeviceLostInfoMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this GPUDevice) Features() (GPUSupportedFeatures, bool) {
	var _ok bool
	_ret := bindings.GetGPUDeviceFeatures(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUSupportedFeatures{}.FromRef(_ret), _ok
}

// Limits returns the value of property "GPUDevice.limits".
//
// The returned bool will be false if there is no such property.
func (this GPUDevice) Limits() (GPUSupportedLimits, bool) {
	var _ok bool
	_ret := bindings.GetGPUDeviceLimits(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUSupportedLimits{}.FromRef(_ret), _ok
}

// Queue returns the value of property "GPUDevice.queue".
//
// The returned bool will be false if there is no such property.
func (this GPUDevice) Queue() (GPUQueue, bool) {
	var _ok bool
	_ret := bindings.GetGPUDeviceQueue(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUQueue{}.FromRef(_ret), _ok
}

// Lost returns the value of property "GPUDevice.lost".
//
// The returned bool will be false if there is no such property.
func (this GPUDevice) Lost() (js.Promise[GPUDeviceLostInfo], bool) {
	var _ok bool
	_ret := bindings.GetGPUDeviceLost(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[GPUDeviceLostInfo]{}.FromRef(_ret), _ok
}

// Label returns the value of property "GPUDevice.label".
//
// The returned bool will be false if there is no such property.
func (this GPUDevice) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGPUDeviceLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "GPUDevice.label" to val.
//
// It returns false if the property cannot be set.
func (this GPUDevice) SetLabel(val js.String) bool {
	return js.True == bindings.SetGPUDeviceLabel(
		this.Ref(),
		val.Ref(),
	)
}

// Destroy calls the method "GPUDevice.destroy".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) Destroy() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceDestroy(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DestroyFunc returns the method "GPUDevice.destroy".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) DestroyFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUDeviceDestroyFunc(
			this.Ref(),
		),
	)
}

// CreateBuffer calls the method "GPUDevice.createBuffer".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateBuffer(descriptor GPUBufferDescriptor) (GPUBuffer, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateBuffer(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUBuffer{}.FromRef(_ret), _ok
}

// CreateBufferFunc returns the method "GPUDevice.createBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateBufferFunc() (fn js.Func[func(descriptor GPUBufferDescriptor) GPUBuffer]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateBufferFunc(
			this.Ref(),
		),
	)
}

// CreateTexture calls the method "GPUDevice.createTexture".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateTexture(descriptor GPUTextureDescriptor) (GPUTexture, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateTexture(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUTexture{}.FromRef(_ret), _ok
}

// CreateTextureFunc returns the method "GPUDevice.createTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateTextureFunc() (fn js.Func[func(descriptor GPUTextureDescriptor) GPUTexture]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateTextureFunc(
			this.Ref(),
		),
	)
}

// CreateSampler calls the method "GPUDevice.createSampler".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateSampler(descriptor GPUSamplerDescriptor) (GPUSampler, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateSampler(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUSampler{}.FromRef(_ret), _ok
}

// CreateSamplerFunc returns the method "GPUDevice.createSampler".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateSamplerFunc() (fn js.Func[func(descriptor GPUSamplerDescriptor) GPUSampler]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateSamplerFunc(
			this.Ref(),
		),
	)
}

// CreateSampler1 calls the method "GPUDevice.createSampler".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateSampler1() (GPUSampler, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateSampler1(
		this.Ref(), js.Pointer(&_ok),
	)

	return GPUSampler{}.FromRef(_ret), _ok
}

// CreateSampler1Func returns the method "GPUDevice.createSampler".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateSampler1Func() (fn js.Func[func() GPUSampler]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateSampler1Func(
			this.Ref(),
		),
	)
}

// ImportExternalTexture calls the method "GPUDevice.importExternalTexture".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) ImportExternalTexture(descriptor GPUExternalTextureDescriptor) (GPUExternalTexture, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceImportExternalTexture(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUExternalTexture{}.FromRef(_ret), _ok
}

// ImportExternalTextureFunc returns the method "GPUDevice.importExternalTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) ImportExternalTextureFunc() (fn js.Func[func(descriptor GPUExternalTextureDescriptor) GPUExternalTexture]) {
	return fn.FromRef(
		bindings.GPUDeviceImportExternalTextureFunc(
			this.Ref(),
		),
	)
}

// CreateBindGroupLayout calls the method "GPUDevice.createBindGroupLayout".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateBindGroupLayout(descriptor GPUBindGroupLayoutDescriptor) (GPUBindGroupLayout, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateBindGroupLayout(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUBindGroupLayout{}.FromRef(_ret), _ok
}

// CreateBindGroupLayoutFunc returns the method "GPUDevice.createBindGroupLayout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateBindGroupLayoutFunc() (fn js.Func[func(descriptor GPUBindGroupLayoutDescriptor) GPUBindGroupLayout]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateBindGroupLayoutFunc(
			this.Ref(),
		),
	)
}

// CreatePipelineLayout calls the method "GPUDevice.createPipelineLayout".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreatePipelineLayout(descriptor GPUPipelineLayoutDescriptor) (GPUPipelineLayout, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreatePipelineLayout(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUPipelineLayout{}.FromRef(_ret), _ok
}

// CreatePipelineLayoutFunc returns the method "GPUDevice.createPipelineLayout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreatePipelineLayoutFunc() (fn js.Func[func(descriptor GPUPipelineLayoutDescriptor) GPUPipelineLayout]) {
	return fn.FromRef(
		bindings.GPUDeviceCreatePipelineLayoutFunc(
			this.Ref(),
		),
	)
}

// CreateBindGroup calls the method "GPUDevice.createBindGroup".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateBindGroup(descriptor GPUBindGroupDescriptor) (GPUBindGroup, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateBindGroup(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUBindGroup{}.FromRef(_ret), _ok
}

// CreateBindGroupFunc returns the method "GPUDevice.createBindGroup".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateBindGroupFunc() (fn js.Func[func(descriptor GPUBindGroupDescriptor) GPUBindGroup]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateBindGroupFunc(
			this.Ref(),
		),
	)
}

// CreateShaderModule calls the method "GPUDevice.createShaderModule".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateShaderModule(descriptor GPUShaderModuleDescriptor) (GPUShaderModule, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateShaderModule(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUShaderModule{}.FromRef(_ret), _ok
}

// CreateShaderModuleFunc returns the method "GPUDevice.createShaderModule".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateShaderModuleFunc() (fn js.Func[func(descriptor GPUShaderModuleDescriptor) GPUShaderModule]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateShaderModuleFunc(
			this.Ref(),
		),
	)
}

// CreateComputePipeline calls the method "GPUDevice.createComputePipeline".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateComputePipeline(descriptor GPUComputePipelineDescriptor) (GPUComputePipeline, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateComputePipeline(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUComputePipeline{}.FromRef(_ret), _ok
}

// CreateComputePipelineFunc returns the method "GPUDevice.createComputePipeline".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateComputePipelineFunc() (fn js.Func[func(descriptor GPUComputePipelineDescriptor) GPUComputePipeline]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateComputePipelineFunc(
			this.Ref(),
		),
	)
}

// CreateRenderPipeline calls the method "GPUDevice.createRenderPipeline".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateRenderPipeline(descriptor GPURenderPipelineDescriptor) (GPURenderPipeline, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateRenderPipeline(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPURenderPipeline{}.FromRef(_ret), _ok
}

// CreateRenderPipelineFunc returns the method "GPUDevice.createRenderPipeline".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateRenderPipelineFunc() (fn js.Func[func(descriptor GPURenderPipelineDescriptor) GPURenderPipeline]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateRenderPipelineFunc(
			this.Ref(),
		),
	)
}

// CreateComputePipelineAsync calls the method "GPUDevice.createComputePipelineAsync".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateComputePipelineAsync(descriptor GPUComputePipelineDescriptor) (js.Promise[GPUComputePipeline], bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateComputePipelineAsync(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return js.Promise[GPUComputePipeline]{}.FromRef(_ret), _ok
}

// CreateComputePipelineAsyncFunc returns the method "GPUDevice.createComputePipelineAsync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateComputePipelineAsyncFunc() (fn js.Func[func(descriptor GPUComputePipelineDescriptor) js.Promise[GPUComputePipeline]]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateComputePipelineAsyncFunc(
			this.Ref(),
		),
	)
}

// CreateRenderPipelineAsync calls the method "GPUDevice.createRenderPipelineAsync".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateRenderPipelineAsync(descriptor GPURenderPipelineDescriptor) (js.Promise[GPURenderPipeline], bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateRenderPipelineAsync(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return js.Promise[GPURenderPipeline]{}.FromRef(_ret), _ok
}

// CreateRenderPipelineAsyncFunc returns the method "GPUDevice.createRenderPipelineAsync".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateRenderPipelineAsyncFunc() (fn js.Func[func(descriptor GPURenderPipelineDescriptor) js.Promise[GPURenderPipeline]]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateRenderPipelineAsyncFunc(
			this.Ref(),
		),
	)
}

// CreateCommandEncoder calls the method "GPUDevice.createCommandEncoder".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateCommandEncoder(descriptor GPUCommandEncoderDescriptor) (GPUCommandEncoder, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateCommandEncoder(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUCommandEncoder{}.FromRef(_ret), _ok
}

// CreateCommandEncoderFunc returns the method "GPUDevice.createCommandEncoder".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateCommandEncoderFunc() (fn js.Func[func(descriptor GPUCommandEncoderDescriptor) GPUCommandEncoder]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateCommandEncoderFunc(
			this.Ref(),
		),
	)
}

// CreateCommandEncoder1 calls the method "GPUDevice.createCommandEncoder".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateCommandEncoder1() (GPUCommandEncoder, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateCommandEncoder1(
		this.Ref(), js.Pointer(&_ok),
	)

	return GPUCommandEncoder{}.FromRef(_ret), _ok
}

// CreateCommandEncoder1Func returns the method "GPUDevice.createCommandEncoder".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateCommandEncoder1Func() (fn js.Func[func() GPUCommandEncoder]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateCommandEncoder1Func(
			this.Ref(),
		),
	)
}

// CreateRenderBundleEncoder calls the method "GPUDevice.createRenderBundleEncoder".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateRenderBundleEncoder(descriptor GPURenderBundleEncoderDescriptor) (GPURenderBundleEncoder, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateRenderBundleEncoder(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPURenderBundleEncoder{}.FromRef(_ret), _ok
}

// CreateRenderBundleEncoderFunc returns the method "GPUDevice.createRenderBundleEncoder".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateRenderBundleEncoderFunc() (fn js.Func[func(descriptor GPURenderBundleEncoderDescriptor) GPURenderBundleEncoder]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateRenderBundleEncoderFunc(
			this.Ref(),
		),
	)
}

// CreateQuerySet calls the method "GPUDevice.createQuerySet".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) CreateQuerySet(descriptor GPUQuerySetDescriptor) (GPUQuerySet, bool) {
	var _ok bool
	_ret := bindings.CallGPUDeviceCreateQuerySet(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&descriptor),
	)

	return GPUQuerySet{}.FromRef(_ret), _ok
}

// CreateQuerySetFunc returns the method "GPUDevice.createQuerySet".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) CreateQuerySetFunc() (fn js.Func[func(descriptor GPUQuerySetDescriptor) GPUQuerySet]) {
	return fn.FromRef(
		bindings.GPUDeviceCreateQuerySetFunc(
			this.Ref(),
		),
	)
}

// PushErrorScope calls the method "GPUDevice.pushErrorScope".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) PushErrorScope(filter GPUErrorFilter) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUDevicePushErrorScope(
		this.Ref(), js.Pointer(&_ok),
		uint32(filter),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PushErrorScopeFunc returns the method "GPUDevice.pushErrorScope".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) PushErrorScopeFunc() (fn js.Func[func(filter GPUErrorFilter)]) {
	return fn.FromRef(
		bindings.GPUDevicePushErrorScopeFunc(
			this.Ref(),
		),
	)
}

// PopErrorScope calls the method "GPUDevice.popErrorScope".
//
// The returned bool will be false if there is no such method.
func (this GPUDevice) PopErrorScope() (js.Promise[GPUError], bool) {
	var _ok bool
	_ret := bindings.CallGPUDevicePopErrorScope(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[GPUError]{}.FromRef(_ret), _ok
}

// PopErrorScopeFunc returns the method "GPUDevice.popErrorScope".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUDevice) PopErrorScopeFunc() (fn js.Func[func() js.Promise[GPUError]]) {
	return fn.FromRef(
		bindings.GPUDevicePopErrorScopeFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this GPUCanvasContext) Canvas() (OneOf_HTMLCanvasElement_OffscreenCanvas, bool) {
	var _ok bool
	_ret := bindings.GetGPUCanvasContextCanvas(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_HTMLCanvasElement_OffscreenCanvas{}.FromRef(_ret), _ok
}

// Configure calls the method "GPUCanvasContext.configure".
//
// The returned bool will be false if there is no such method.
func (this GPUCanvasContext) Configure(configuration GPUCanvasConfiguration) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCanvasContextConfigure(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&configuration),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ConfigureFunc returns the method "GPUCanvasContext.configure".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCanvasContext) ConfigureFunc() (fn js.Func[func(configuration GPUCanvasConfiguration)]) {
	return fn.FromRef(
		bindings.GPUCanvasContextConfigureFunc(
			this.Ref(),
		),
	)
}

// Unconfigure calls the method "GPUCanvasContext.unconfigure".
//
// The returned bool will be false if there is no such method.
func (this GPUCanvasContext) Unconfigure() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallGPUCanvasContextUnconfigure(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnconfigureFunc returns the method "GPUCanvasContext.unconfigure".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCanvasContext) UnconfigureFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.GPUCanvasContextUnconfigureFunc(
			this.Ref(),
		),
	)
}

// GetCurrentTexture calls the method "GPUCanvasContext.getCurrentTexture".
//
// The returned bool will be false if there is no such method.
func (this GPUCanvasContext) GetCurrentTexture() (GPUTexture, bool) {
	var _ok bool
	_ret := bindings.CallGPUCanvasContextGetCurrentTexture(
		this.Ref(), js.Pointer(&_ok),
	)

	return GPUTexture{}.FromRef(_ret), _ok
}

// GetCurrentTextureFunc returns the method "GPUCanvasContext.getCurrentTexture".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GPUCanvasContext) GetCurrentTextureFunc() (fn js.Func[func() GPUTexture]) {
	return fn.FromRef(
		bindings.GPUCanvasContextGetCurrentTextureFunc(
			this.Ref(),
		),
	)
}
