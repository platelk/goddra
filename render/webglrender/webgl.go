package webglrender

import (
	"fmt"
	"goddra/render/webgl"
	"syscall/js"
)

var ErrNotCompatibleBrowser = fmt.Errorf("browser don't support webglrender")

// WebGL is go binding on WebGL
type WebGL struct {
	types  *webgl.GLTypes
	gl     js.Value
	width  int
	height int
}

// NewWebGL will instantiate a webglrender binding for Go
//func NewWebGL(canvasEl js.Value) (*WebGL, error) {
//	width := canvasEl.Get("clientWidth").Int()
//	height := canvasEl.Get("clientHeight").Int()
//	canvasEl.Set("width", width)
//	canvasEl.Set("height", height)
//	gl := canvasEl.Call("getContext", "webgl2")
//	// once again
//	if gl == js.Undefined() {
//		return nil, ErrNotCompatibleBrowser
//	}
//	fmt.Println(width, height)
//	return &WebGL{
//		width:  width,
//		height: height,
//		gl:     gl,
//		types:  webgl.NewGLTypes(gl),
//	}, nil
//}
//
//func (wg *WebGL) DrawTriangle() error {
//	gl := wg.gl
//
//	fmt.Println("create vert shader...")
//	vertex, err := wg.createShader(wg.types.VertexShader, vertexShaderSource)
//	if err != nil {
//		return err
//	}
//	fmt.Println("create frag shader...")
//	frag, err := wg.createShader(wg.types.FragmentShader, fragmentShaderSource)
//	if err != nil {
//		return err
//	}
//	fmt.Println("create program...")
//	prog, err := wg.createProgram(vertex, frag)
//	if err != nil {
//		return err
//	}
//	// look up where the vertex data needs to go.
//	aPosLoc := gl.Call("getAttribLocation", prog, "a_position")
//	// Create a buffer and put three 2d clip space points in it
//	posBuffer := gl.Call("createBuffer")
//	// Bind it to ARRAY_BUFFER (think of it as ARRAY_BUFFER = positionBuffer)
//	gl.Call("bindBuffer", wg.types.ArrayBuffer, posBuffer)
//	positions := []float32{
//		0, 0,
//		0, 0.5,
//		0.7, 0,
//	}
//
//	farr := js.Global().Get("Float32Array").New(6)
//	for i, v := range positions {
//		farr.SetIndex(i, v)
//	}
//	gl.Call("bufferData", wg.types.ArrayBuffer, farr, wg.types.StaticDraw)
//	// Create a vertex array object (attribute state)
//	vao := gl.Call("createVertexArray")
//	// and make it the one we're currently working with
//	gl.Call("bindVertexArray", vao)
//	// Turn on the attribute
//	gl.Call("enableVertexAttribArray", aPosLoc)
//	// Tell the attribute how to get data out of positionBuffer (ARRAY_BUFFER)
//	size := 2          // 2 components per iteration
//	ty := wg.types.Float      // the data is 32bit floats
//	normalize := false // don't normalize the data
//	stride := 0        // 0 = move forward size * sizeof(type) each iteration to get the next position
//	offset := 0        // start at the beginning of the buffer
//	gl.Call("vertexAttribPointer",
//		aPosLoc, size,
//	ty, normalize, stride, offset)
//
//	//js.Global().Get("webglUtils").Call("resizeCanvasToDisplaySize", gl.Get("canvas"))
//
//	// Tell WebGL how to convert from clip space to pixels
//	gl.Call("viewport", 0, 0, wg.width, wg.height)
//	// Clear the canvas
//	gl.Call("clearColor", 0.8, 0.8, 0.8, 0.9)
//	gl.Call("clear", wg.types.ColorBufferBit)
//
//	gl.Call("useProgram", prog)
//	gl.Call("bindVertexArray", vao)
//	// Enable the depth test
//	gl.Call("enable", wg.types.DepthTest)
//	gl.Call("drawArrays", wg.types.Triangles, 0, 3)
//	return fmt.Errorf("end of program")
//}
//
//
//func (wg *WebGL) createShader(shaderType js.Value, source string) (js.Value, error) {
//	gl := wg.gl
//	shader := gl.Call("createShader", shaderType)
//	gl.Call("shaderSource", shader, source)
//	gl.Call("compileShader", shader)
//	success := gl.Call("getShaderParameter", shader, wg.types.CompileStatus).Bool()
//	if success {
//		return shader, nil
//	}
//	info := gl.Call("getShaderInfoLog", shader).String()
//	gl.Call("deleteShader", shader)
//	return js.Null(), fmt.Errorf("%s", info)
//}
//
//func (wg *WebGL) createProgram(vertexShader, fragmentShader js.Value) (js.Value, error) {
//	gl := wg.gl
//
//	gl.Call("attachShader", program, vertexShader)
//	gl.Call("attachShader", program, fragmentShader)
//
//	gl.Call("linkProgram", program)
//
//	success := gl.Call("getProgramParameter", program, wg.types.LinkStatus).Bool()
//	if success {
//		return program, nil
//	}
//	info := gl.Call("getProgramInfoLog", program).String()
//	gl.Call("deleteProgram", program)
//	return js.Null(), fmt.Errorf("%s", info)
//}
//
