// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

// Code adapted from
// https://github.com/webgpu/webgpu-samples/blob/main/src/sample/rotatingCube/main.ts
//
// Copyright 2019 WebGPU Samples Contributors
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	_ "embed" // for go:embed

	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/math"
	"github.com/primecitizens/std/math/matrix"
	"github.com/primecitizens/std/plat/js/web"
	"github.com/primecitizens/std/ui/html"
)

var (
	webgpu, hasWebGPU = navigator.Gpu()

	strMain = js.NewString("main")
)

var (
	//go:embed vertex.wgsl
	vertexWGSL string

	//go:embed vertexPosColor.wgsl
	vertexPosColorWGSL string
)

type ExampleWebGPU struct {
	initialized bool
	running     bool

	animationID uint32

	dev     web.GPUDevice
	q       web.GPUQueue
	context web.GPUCanvasContext

	projectionMatrix matrix.Mat4[float32]

	pipeline web.GPURenderPipeline

	colorAttachmentRef js.Ref
	colorAttachment    web.GPURenderPassColorAttachment

	renderPassDescriptor web.GPURenderPassDescriptor

	verticesBuffer web.GPUBuffer

	uniformBindGroup web.GPUBindGroup
	uniformBuffer    web.GPUBuffer

	frameHandler    js.Func[func(time web.DOMHighResTimeStamp)]
	callbackContext web.FrameRequestCallback[*ExampleWebGPU]

	modelViewProjectionMatrix js.TypedArray[float32]

	fpsLine web.HTMLElement
	fps     web.HTMLSpanElement
	canvas  web.HTMLCanvasElement

	frames        int
	lastFPSUpdate web.DOMHighResTimeStamp
	fpsBuilder    html.Builder
}

func (x *ExampleWebGPU) Run() {
	if !hasWebGPU {
		displayBuffer.Discard().TextBlock("<p>", "</p>",
			"Sorry, WebGPU is not supported in this browser 😭",
		)
		return
	}

	if x.running {
		x.animate()
		return
	}

	if !x.initialized {
		if !x.init() {
			return
		}

		x.initialized = true
	}

	x.running = true
	displayArea.AppendChild(x.fpsLine.Node)
	displayArea.AppendChild(x.canvas.Node)
	x.animate()
}

func (x *ExampleWebGPU) Stop() {
	if x.running {
		x.frames = 0
		x.running = false
		window.CancelAnimationFrame(x.animationID)
	}
}

func (x *ExampleWebGPU) init() bool {
	adapter, err, ok := js.Must(webgpu.RequestAdapter1()).Await(true)
	if !ok {
		displayBuffer.Discard().
			TextBlock("<p>", "</p>",
				"no gpu adapter available 😭, please see console log for error information.",
			).Flush(false)
		console.Log(err)
		return false
	}

	x.dev, err, ok = js.Must(adapter.Once().RequestDevice1()).Await(true)
	if !ok {
		displayBuffer.Discard().
			TextBlock("<p>", "</p>",
				"no gpu device available 😭, please see console log for error information.",
			).Flush(false)
		console.Log(err)
		return false
	}

	x.callbackContext = web.FrameRequestCallback[*ExampleWebGPU]{
		Fn:  (*ExampleWebGPU).updateFrame,
		Arg: x,
	}
	x.frameHandler = x.callbackContext.Register()

	x.fpsLine = web.HTMLElement{}.FromRef(js.Must(document.CreateElement1(js.NewString("p").Once())).Ref())
	x.fps = web.NewHTMLSpanElement()
	x.canvas = web.NewHTMLCanvasElement()

	x.fpsLine.SetInnerText(js.NewString("fps: ").Once())
	x.fpsLine.AppendChild(x.fps.Node)
	x.fpsBuilder.Reset(x.fps.Ref()).Text("0").Flush(false)

	x.context = web.GPUCanvasContext{}.FromRef(
		js.Must(x.canvas.GetContext1(js.NewString("webgpu").Once())).Ref(),
	)
	if x.context.Ref().Undefined() {
		displayBuffer.Discard().
			TextBlock("<p>", "</p>",
				"webgpu context not available in canvas 😭",
			).Flush(false)
		return false
	}

	devicePixelRatio := js.Must(window.DevicePixelRatio())
	if devicePixelRatio == 0 {
		devicePixelRatio = 1
	}

	// TODO: adjust width and height automatically
	width := uint32(800)  // uint32(float64(js.Must(displayArea.ClientWidth())) * devicePixelRatio)
	height := uint32(600) // uint32(float64(js.Must(window.InnerHeight())-js.Must(x.fpsLine.ClientHeight())) * devicePixelRatio)
	x.canvas.SetWidth(width)
	x.canvas.SetHeight(height)

	presentationFormat := js.Must(webgpu.GetPreferredCanvasFormat())

	js.Must(x.context.Configure(web.GPUCanvasConfiguration{
		Device:    x.dev,
		Format:    presentationFormat,
		AlphaMode: web.GPUCanvasAlphaMode_PREMULTIPLIED,
	}))

	cubeVertexArray := js.NewTypedArrayOf[float32](
		// float4 position, float4 color, float2 uv,
		1, -1, 1, 1, 1, 0, 1, 1, 0, 1,
		-1, -1, 1, 1, 0, 0, 1, 1, 1, 1,
		-1, -1, -1, 1, 0, 0, 0, 1, 1, 0,
		1, -1, -1, 1, 1, 0, 0, 1, 0, 0,
		1, -1, 1, 1, 1, 0, 1, 1, 0, 1,
		-1, -1, -1, 1, 0, 0, 0, 1, 1, 0,

		1, 1, 1, 1, 1, 1, 1, 1, 0, 1,
		1, -1, 1, 1, 1, 0, 1, 1, 1, 1,
		1, -1, -1, 1, 1, 0, 0, 1, 1, 0,
		1, 1, -1, 1, 1, 1, 0, 1, 0, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 0, 1,
		1, -1, -1, 1, 1, 0, 0, 1, 1, 0,

		-1, 1, 1, 1, 0, 1, 1, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, -1, 1, 1, 1, 0, 1, 1, 0,
		-1, 1, -1, 1, 0, 1, 0, 1, 0, 0,
		-1, 1, 1, 1, 0, 1, 1, 1, 0, 1,
		1, 1, -1, 1, 1, 1, 0, 1, 1, 0,

		-1, -1, 1, 1, 0, 0, 1, 1, 0, 1,
		-1, 1, 1, 1, 0, 1, 1, 1, 1, 1,
		-1, 1, -1, 1, 0, 1, 0, 1, 1, 0,
		-1, -1, -1, 1, 0, 0, 0, 1, 0, 0,
		-1, -1, 1, 1, 0, 0, 1, 1, 0, 1,
		-1, 1, -1, 1, 0, 1, 0, 1, 1, 0,

		1, 1, 1, 1, 1, 1, 1, 1, 0, 1,
		-1, 1, 1, 1, 0, 1, 1, 1, 1, 1,
		-1, -1, 1, 1, 0, 0, 1, 1, 1, 0,
		-1, -1, 1, 1, 0, 0, 1, 1, 1, 0,
		1, -1, 1, 1, 1, 0, 1, 1, 0, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 0, 1,

		1, -1, -1, 1, 1, 0, 0, 1, 0, 1,
		-1, -1, -1, 1, 0, 0, 0, 1, 1, 1,
		-1, 1, -1, 1, 0, 1, 0, 1, 1, 0,
		1, 1, -1, 1, 1, 1, 0, 1, 0, 0,
		1, -1, -1, 1, 1, 0, 0, 1, 0, 1,
		-1, 1, -1, 1, 0, 1, 0, 1, 1, 0,
	)
	x.verticesBuffer = js.Must(x.dev.CreateBuffer(web.GPUBufferDescriptor{
		Size:  web.GPUSize64(cubeVertexArray.ByteLength()),
		Usage: web.GPUBufferUsageFlags(web.GPUBufferUsage_VERTEX),

		FFI_USE_MappedAtCreation: true,
		MappedAtCreation:         true,
	}))

	js.TypedArray[float32]{}.FromArrayBuffer(true, js.Must(x.verticesBuffer.GetMappedRange2())).
		Once().Set(cubeVertexArray)
	x.verticesBuffer.Unmap()

	const (
		cubeVertexSize     = 4 * 10 // Byte size of one cube vertex.
		cubePositionOffset = 0
		cubeColorOffset    = 4 * 4 // Byte offset of cube vertex color attribute.
		cubeUVOffset       = 4 * 8
		cubeVertexCount    = 36
	)

	x.pipeline = js.Must(x.dev.CreateRenderPipeline(web.GPURenderPipelineDescriptor{
		Layout: web.OneOf_GPUPipelineLayout_GPUAutoLayoutMode{}.FromRef(
			js.NewString(js.Must(web.GPUAutoLayoutMode_AUTO.String())).Ref(),
		),
		Vertex: web.GPUVertexState{
			FFI_USE: true,

			Module: js.Must(
				x.dev.CreateShaderModule(web.GPUShaderModuleDescriptor{
					Code: js.NewString(vertexWGSL).Once(),
				}),
			),
			EntryPoint: strMain,
			Buffers: js.NewArrayOf[web.GPUVertexBufferLayout](true,
				web.GPUVertexBufferLayout{
					ArrayStride: cubeVertexSize,
					Attributes: js.NewArrayOf[web.GPUVertexAttribute](true,
						web.GPUVertexAttribute{
							Format:         web.GPUVertexFormat_FLOAT_32X4,
							Offset:         cubePositionOffset,
							ShaderLocation: 0,
						}.New(),
						web.GPUVertexAttribute{
							Format:         web.GPUVertexFormat_FLOAT_32X2,
							Offset:         cubeUVOffset,
							ShaderLocation: 1,
						}.New(),
					),
				}.New(),
			),
		},
		Fragment: web.GPUFragmentState{
			FFI_USE: true,

			Module: js.Must(x.dev.CreateShaderModule(web.GPUShaderModuleDescriptor{
				Code: js.NewString(vertexPosColorWGSL).Once(),
			})),
			EntryPoint: strMain,
			Targets: js.NewArrayOf[web.GPUColorTargetState](true,
				web.GPUColorTargetState{
					Format: presentationFormat,
				}.New(),
			),
		},
		Primitive: web.GPUPrimitiveState{
			FFI_USE: true,

			Topology: web.GPUPrimitiveTopology_TRIANGLE_LIST,
			CullMode: web.GPUCullMode_BACK,
		},
		DepthStencil: web.GPUDepthStencilState{
			FFI_USE: true,

			Format:            web.GPUTextureFormat_DEPTH_24PLUS,
			DepthWriteEnabled: true,
			DepthCompare:      web.GPUCompareFunction_LESS,
		},
	}))

	depthTexture := js.Must(x.dev.CreateTexture(web.GPUTextureDescriptor{
		Size: web.GPUExtent3D{}.FromRef(
			web.GPUExtent3DDict{
				Width:          web.GPUIntegerCoordinate(width),
				FFI_USE_Height: true,
				Height:         web.GPUIntegerCoordinate(height),
			}.New().Once(),
		),
		Format: web.GPUTextureFormat_DEPTH_24PLUS,
		Usage:  web.GPUTextureUsageFlags(web.GPUTextureUsage_RENDER_ATTACHMENT),
	}))

	const uniformBufferSize = 4 * 16 // 4x4 matrix
	x.uniformBuffer = js.Must(x.dev.CreateBuffer(web.GPUBufferDescriptor{
		Size: uniformBufferSize,
		Usage: web.GPUBufferUsageFlags(
			web.GPUBufferUsage_UNIFORM | web.GPUBufferUsage_COPY_DST,
		),
	}))

	x.uniformBindGroup = js.Must(x.dev.CreateBindGroup(web.GPUBindGroupDescriptor{
		Layout: js.Must(x.pipeline.GetBindGroupLayout(0)),
		Entries: js.NewArrayOf[web.GPUBindGroupEntry](true,
			web.GPUBindGroupEntry{
				Binding: 0,
				Resource: web.OneOf_GPUSampler_GPUTextureView_GPUBufferBinding_GPUExternalTexture{}.
					FromRef(
						web.GPUBufferBinding{
							Buffer: x.uniformBuffer,
						}.New().Once(),
					),
			}.New(),
		),
	}))

	x.colorAttachment = web.GPURenderPassColorAttachment{
		View:          web.GPUTextureView{},
		ResolveTarget: web.GPUTextureView{},
		ClearValue: web.OneOf_ArrayFloat64_GPUColorDict{}.FromRef(
			web.GPUColorDict{
				R: 0.5,
				G: 0.5,
				B: 0.5,
				A: 1.0,
			}.New(),
		),
		LoadOp:  web.GPULoadOp_CLEAR,
		StoreOp: web.GPUStoreOp_STORE,
	}

	x.colorAttachmentRef = x.colorAttachment.New()

	x.renderPassDescriptor = web.GPURenderPassDescriptor{
		ColorAttachments: js.NewArrayOf[web.GPURenderPassColorAttachment](false, x.colorAttachmentRef),
		DepthStencilAttachment: web.GPURenderPassDepthStencilAttachment{
			FFI_USE: true,

			View:                    js.Must(depthTexture.CreateView1()),
			FFI_USE_DepthClearValue: true,
			DepthClearValue:         1.0,
			DepthLoadOp:             web.GPULoadOp_CLEAR,
			DepthStoreOp:            web.GPUStoreOp_STORE,
		},
	}

	matrix.SetPerspective4(
		&x.projectionMatrix,
		(2*math.Pi)/5,
		float64(width)/float64(height),
		1,
		100.0,
	)

	x.modelViewProjectionMatrix = js.NewTypedArray[float32](16)
	x.q = js.Must(x.dev.Queue())

	return true
}

func (x *ExampleWebGPU) getTransformationMatrix(now float64) js.TypedArray[float32] {
	var (
		viewMat       matrix.Mat4[float32]
		projectionMat = x.projectionMatrix
	)

	now /= 700
	viewMat.SetIdentity().
		Translate(matrix.Vec3[float32]{0, 0, -4}).
		Rotate(1, matrix.Vec3[float32]{float32(math.Sin(now)), float32(math.Cos(now)), 0})

	return x.modelViewProjectionMatrix.MustFill(0, projectionMat.Mul(viewMat)[:]...)
}

func (x *ExampleWebGPU) updateFrame(this js.Ref, time web.DOMHighResTimeStamp) js.Ref {
	if !x.running {
		return js.Undefined
	}

	if time-x.lastFPSUpdate >= 1000 {
		x.fpsBuilder.Int(int64(x.frames), 10).Flush(false)
		x.frames = 0
		x.lastFPSUpdate = time
	} else {
		x.frames++
	}

	txMat := x.getTransformationMatrix(float64(time))

	x.q.WriteBuffer(
		x.uniformBuffer,
		0,
		web.AllowSharedBufferSource{}.FromRef(txMat.Buffer(false)),
		web.GPUSize64(txMat.ByteOffset()),
		web.GPUSize64(txMat.ByteLength()),
	)
	x.colorAttachment.View = js.Must(js.Must(x.context.GetCurrentTexture()).Once().CreateView1()).Once()
	x.colorAttachment.Update(x.colorAttachmentRef)

	cmdEnc := js.Must(x.dev.CreateCommandEncoder1())
	enc := js.Must(cmdEnc.BeginRenderPass(x.renderPassDescriptor))
	enc.SetPipeline(x.pipeline)
	enc.SetBindGroup1(0, x.uniformBindGroup)
	enc.SetVertexBuffer2(0, x.verticesBuffer)
	enc.Draw3(36)
	enc.End()
	x.q.Submit(js.NewArrayOf[web.GPUCommandBuffer](true, js.Must(cmdEnc.Finish1()).Ref()).Once())

	x.animate()

	return js.Undefined
}

func (x *ExampleWebGPU) animate() {
	x.animationID = js.Must(window.RequestAnimationFrame(x.frameHandler))
}
