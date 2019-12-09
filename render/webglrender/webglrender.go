package webglrender

import (
	"fmt"
	"goddra/render/webgl"
	"goddra/render/webgl/wgutils"
	"syscall/js"
)

type WebGLRender struct {
	gl *webgl.WebGL
}

func NewWebGLRender(canvasEl js.Value) (*WebGLRender, error) {
	wgr := &WebGLRender{}
	gl, err := webgl.NewWebGL(canvasEl)
	if err != nil {
		return nil, fmt.Errorf("can't instanciate WebGL for renderer: %w", err)
	}
	wgr.gl = gl
	return wgr, nil
}

func (wgr *WebGLRender) DrawTriangle() error {
	vshader, err := wgutils.CreateShader(wgr.gl, wgr.gl.Types().VertexShader, triangleVertexShaderSource)
	if err != nil {
		return err
	}
	fshader, err := wgutils.CreateShader(wgr.gl, wgr.gl.Types().FragmentShader, triangleFragmentShaderSource)
	if err != nil {
		return err
	}
	prog, err := wgutils.CreateProgram(wgr.gl, vshader, fshader)
	if err != nil {
		return err
	}
	aPosLoc := prog.GetAttribute("a_position")
	posBuff := wgr.gl.CreateBuffer()
	posBuff.BindToArrayBuffer()
	posBuff.Float32ArrayData(wgutils.ToFloat32Array([]float32{
		0, 0,
		0, 0.5,
		0.7, 0,
	}), wgr.gl.Types().StaticDraw)
	vao := wgr.gl.CreateVertexArray()
	vao.Bind()
	wgr.gl.EnableVertexAttribArray(aPosLoc)
	// Tell the attribute how to get data out of positionBuffer (ARRAY_BUFFER)
	size := 2          // 2 components per iteration
	ty := wgr.gl.Types().Float      // the data is 32bit floats
	normalize := false // don't normalize the data
	stride := 0        // 0 = move forward size * sizeof(type) each iteration to get the next position
	offset := 0        // start at the beginning of the buffer
	wgr.gl.VertexAttribPointer(
		aPosLoc, size,
		ty, normalize, stride, offset)
	wgr.gl.Viewport(0, 0, wgr.gl.Width(), wgr.gl.Height())
	wgr.gl.ClearColor(0.8, 0.8, 0.8, 0.9)
	wgr.gl.Clear(wgr.gl.Types().ColorBufferBit)

	wgr.gl.UseProgram(prog)
	wgr.gl.BindVertexArray(vao)
	wgr.gl.Enable(wgr.gl.Types().DepthTest)
	wgr.gl.DrawArrays(wgr.gl.Types().Triangles, 0, 3)
	return nil
}

func (wgr *WebGLRender) Render() {

}

const triangleVertexShaderSource = `#version 300 es

// an attribute is an input (in) to a vertex shader.
// It will receive data from a buffer
in vec4 a_position;

// all shaders have a main function
void main() {

  // gl_Position is a special variable a vertex shader
  // is responsible for setting
  gl_Position = a_position;
}
`

const triangleFragmentShaderSource = `#version 300 es

// fragment shaders don't have a default precision so we need
// to pick one. mediump is a good default. It means "medium precision"
precision mediump float;

// we need to declare an output for the fragment shader
out vec4 outColor;

void main() {
  // Just set the output to a constant redish-purple
  outColor = vec4(1, 0, 0.5, 1);
}
`