package webgl

import "syscall/js"

type GLType js.Value

// GLTypes provides WebGL bindings.
type GLTypes struct {
	StaticDraw         GLType
	ArrayBuffer        GLType
	ElementArrayBuffer GLType
	VertexShader       GLType
	FragmentShader     GLType
	Float              GLType
	DepthTest          GLType
	ColorBufferBit     GLType
	DepthBufferBit     GLType
	Triangles          GLType
	UnsignedShort      GLType
	LEqual             GLType
	LineLoop           GLType
	CompileStatus      GLType
	LinkStatus      	GLType
}

// New grabs the WebGL bindings from a GL context.
func NewGLTypes(wg *WebGL) *GLTypes {
	gl := wg.gl
	types := &GLTypes{}

	types.StaticDraw = GLType(gl.Get("STATIC_DRAW"))
	types.ArrayBuffer = GLType(gl.Get("ARRAY_BUFFER"))
	types.ElementArrayBuffer = GLType(gl.Get("ELEMENT_ARRAY_BUFFER"))
	types.VertexShader = GLType(gl.Get("VERTEX_SHADER"))
	types.FragmentShader = GLType(gl.Get("FRAGMENT_SHADER"))
	types.Float = GLType(gl.Get("FLOAT"))
	types.DepthTest = GLType(gl.Get("DEPTH_TEST"))
	types.ColorBufferBit = GLType(gl.Get("COLOR_BUFFER_BIT"))
	types.Triangles = GLType(gl.Get("TRIANGLES"))
	types.UnsignedShort = GLType(gl.Get("UNSIGNED_SHORT"))
	types.LEqual = GLType(gl.Get("LEQUAL"))
	types.DepthBufferBit = GLType(gl.Get("DEPTH_BUFFER_BIT"))
	types.LineLoop = GLType(gl.Get("LINE_LOOP"))
	types.CompileStatus = GLType(gl.Get("COMPILE_STATUS"))
	types.LinkStatus = GLType(gl.Get("LINK_STATUS"))

	return types
}
