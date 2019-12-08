package webgl

import (
	"fmt"
	"syscall/js"
)

var ErrNotCompatibleBrowser = fmt.Errorf("browser don't support webglrender")

type WebGL struct {
	width, height int
	canvasEl      js.Value
	types         *GLTypes
	gl            js.Value
}

// NewWebGL will instantiate a webglrender binding for Go
func NewWebGL(canvasEl js.Value) (*WebGL, error) {
	width := canvasEl.Get("clientWidth").Int()
	height := canvasEl.Get("clientHeight").Int()
	canvasEl.Set("width", width)
	canvasEl.Set("height", height)
	gl := canvasEl.Call("getContext", "webgl2")
	// once again
	if gl == js.Undefined() {
		return nil, ErrNotCompatibleBrowser
	}
	fmt.Println(width, height)
	wg := &WebGL{
		canvasEl: canvasEl,
		width:    width,
		height:   height,
		gl:       gl,
	}
	wg.types = NewGLTypes(wg)
	wg.Resize()
	return wg, nil
}

func (wg *WebGL) Resize() {
	width := wg.canvasEl.Get("clientWidth").Int()
	height := wg.canvasEl.Get("clientHeight").Int()
	if wg.height != height {
		wg.canvasEl.Set("height", height)
	}
	if wg.width != width {
		wg.canvasEl.Set("width", width)
	}
}

func (wg *WebGL) Viewport(x1, y1, x2, y2 int) {
	wg.gl.Call("viewport", x1, y1, x2, y2)
}

func (wg *WebGL) ClearColor(r, g, b, a float32) {
	wg.gl.Call("clearColor", r, g, b, a)
}

func (wg *WebGL) Clear(t GLType) {
	wg.gl.Call("clear", t)
}

func (wg *WebGL) UseProgram(p *Program) {
	wg.gl.Call("useProgram", p.value)
}

func (wg *WebGL) Enable(t GLType) {
	wg.gl.Call("enable", t)
}

func (wg *WebGL) DrawArrays(drawType GLType, from, to int) {
	wg.gl.Call("drawArrays", drawType, from, to)
}

// =============
// Shader
// =============

func (wg *WebGL) CreateShader(shaderType GLType) *Shader {
	shader := wg.gl.Call("createShader", shaderType)
	return &Shader{value: shader, wg: wg}
}

func (wg *WebGL) CreateVertexShader() *Shader {
	return wg.CreateShader(wg.types.VertexShader)
}

func (wg *WebGL) CreateFragmentShader() *Shader {
	return wg.CreateShader(wg.types.FragmentShader)
}

func (wg *WebGL) ShaderSource(shader *Shader, source string) {
	wg.gl.Call("shaderSource", shader.value, source)
}

func (wg *WebGL) CompileShader(shader *Shader) {
	wg.gl.Call("compileShader", shader.value)
}

func (wg *WebGL) GetShaderParameter(shader *Shader, parameter GLType) js.Value {
	return wg.gl.Call("getShaderParameter", shader.value, js.Value(parameter))
}

func (wg *WebGL) GetShaderCompileStatus(shader *Shader) bool {
	return wg.GetShaderParameter(shader, wg.types.CompileStatus).Bool()
}

func (wg *WebGL) GetShaderInfoLog(shader *Shader) string {
	return wg.gl.Call("getShaderInfoLog", shader.value).String()
}

func (wg *WebGL) DeleteShader(shader *Shader) {
	wg.gl.Call("deleteShader", shader.value)
}

// =============
// Program
// =============

func (wg *WebGL) CreateProgram() *Program {
	program := wg.gl.Call("createProgram")
	return &Program{
		wg:    wg,
		value: program,
	}
}

func (wg *WebGL) AttachShader(program *Program, shader *Shader) {
	wg.gl.Call("attachShader", program.value, shader.value)
}

func (wg *WebGL) LinkProgram(program *Program) {
	wg.gl.Call("linkProgram", program.value)
}

func (wg *WebGL) DeleteProgram(program *Program) {
	wg.gl.Call("deleteProgram", program.value)
}

func (wg *WebGL) GetProgramParameter(program *Program, parameters GLType) js.Value {
	return wg.gl.Call("getProgramParameter", program.value, parameters)
}

func (wg *WebGL) GetProgramLinkStatus(program *Program) bool {
	return wg.gl.Call("getProgramParameter", program.value, wg.types.LinkStatus).Bool()
}

func (wg *WebGL) GetProgramInfoLog(program *Program) string {
	return wg.gl.Call("getProgramInfoLog", program.value).String()
}

func (wg *WebGL) GetAttribLocation(program *Program, attrib string) *Attribute {
	attribute := wg.gl.Call("getAttribLocation", program.value, attrib)
	return &Attribute{prog: program, wg: wg, value: attribute, name: attrib}
}

// =============
// Buffer
// =============

func (wg *WebGL) CreateBuffer() *Buffer {
	return &Buffer{value: wg.gl.Call("createBuffer"), wg: wg}
}

func (wg *WebGL) BindBuffer(t GLType, b *Buffer) {
	wg.gl.Call("bindBuffer", t, b.value)
}

func (wg *WebGL) BufferData(t GLType, data interface{}, d GLType) {
	wg.gl.Call("bufferData", t, data, d)
}

// =============
// Vertex Array
// =============

func (wg *WebGL) CreateVertexArray() *VertexArray {
	v := wg.gl.Call("createVertexArray")
	return &VertexArray{wg: wg, value: v}
}

func (wg *WebGL) BindVertexArray(v *VertexArray) {
	wg.gl.Call("bindVertexArray", v.value)
}

func (wg *WebGL) EnableVertexAttribArray(a *Attribute) {
	wg.gl.Call("enableVertexAttribArray", a.value)
}

func (wg *WebGL) VertexAttribPointer(a *Attribute, size int, ty GLType, normalize bool, stride, offset int) {
	wg.gl.Call("vertexAttribPointer",
		a.value, size,
		ty, normalize, stride, offset)
}
