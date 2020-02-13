package webglrender

import (
	"fmt"
	"goddra/geom"
	"goddra/graphic/color"
	"goddra/render/webglrender/shaders"
	"goddra/webgl"
	"goddra/webgl/wgutils"
	"syscall/js"
)

type WebGLRender struct {
	gl              *webgl.WebGL
	backgroundColor color.Color
	renderProg      map[string]func(...interface{}) error
	renderPipeline  map[string][]interface{}
}

func NewWebGLRender(canvasEl js.Value) (*WebGLRender, error) {
	wgr := &WebGLRender{}
	gl, err := webgl.NewWebGL(canvasEl)
	if err != nil {
		return nil, fmt.Errorf("can't instanciate WebGL for renderer: %w", err)
	}
	wgr.gl = gl
	wgr.backgroundColor = color.New(1.0, 1.0, 1.0, 1.0)
	wgr.renderProg = make(map[string]func(...interface{})error)
	wgr.renderPipeline = make(map[string][]interface{})
	err = wgr.initTriangleColorProg()
	if err != nil {
		return nil, err
	}
	return wgr, nil
}

func (wgr *WebGLRender) initTriangleColorProg() error {
	vshader, err := wgutils.CreateShader(wgr.gl, wgr.gl.Types().VertexShader, shaders.TriangleVertexShaderSource)
	if err != nil {
		return err
	}
	fshader, err := wgutils.CreateShader(wgr.gl, wgr.gl.Types().FragmentShader, shaders.TriangleFragmentShaderSource)
	if err != nil {
		return err
	}
	prog, err := wgutils.CreateProgram(wgr.gl, vshader, fshader)
	if err != nil {
		return err
	}
	aPosLoc := prog.GetAttribute("a_position")
	colorLoc := prog.GetUniformLoc("u_color")
	resLoc := prog.GetUniformLoc("u_resolution")
	transLoc := prog.GetUniformLoc("u_translation")
	posBuff := wgr.gl.CreateBuffer()
	vao := wgr.gl.CreateVertexArray()


	wgr.renderProg["triangle_color"] = func(i ...interface{}) error {
		// -- Setup vertex
		posBuff.BindToArrayBuffer()
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
		// -- end setup
		wgr.gl.UseProgram(prog)
		wgr.gl.BindVertexArray(vao)
		for _, ii := range i {
			t, ok := ii.(geom.Shape)
			if !ok {
				return fmt.Errorf("expect (graphic.Shape) but got %t", i)
			}

			colorLoc.SetUniform(0.8, 0.8, 0.8, 0.8)
			resLoc.SetUniform(float32(wgr.gl.Width()), float32(wgr.gl.Height()))
			transLoc.SetUniform(t.Position().Values()...)
			points := t.Points()
			posBuff.Float32ArrayData(wgutils.ToFloat32Array(points), wgr.gl.Types().StaticDraw)
			fmt.Println("draw")
			wgr.gl.DrawArrays(wgr.gl.Types().Triangles, 0, len(points)/t.Dim())
		}
		return nil
	}
	return nil
}

func (wgr *WebGLRender) DrawTriangle(t *geom.Triangle) error {
	wgr.renderPipeline["triangle"] = append(wgr.renderPipeline["triangle"], t)
	return nil
}

func (wgr *WebGLRender) DrawRectangle(t *geom.Rectangle) error {
	wgr.renderPipeline["triangle"] = append(wgr.renderPipeline["triangle"], t)
	return nil
}
func (wgr *WebGLRender) Resize() {
	wgr.gl.Viewport(0, 0, wgr.gl.Width(), wgr.gl.Height())
}

func (wgr *WebGLRender) Clear() {
	wgr.gl.ClearColor(wgr.backgroundColor.R(), wgr.backgroundColor.G(), wgr.backgroundColor.B(), wgr.backgroundColor.A())
	wgr.gl.Clear(wgr.gl.Types().ColorBufferBit)
}

func (wgr *WebGLRender) BackgroundColor(color color.Color) {
	wgr.backgroundColor = color
}

func (wgr *WebGLRender) Render() error {
	wgr.Resize()
	wgr.Clear()
	for k, v := range wgr.renderPipeline {
		if err := wgr.renderProg[k](v...); err != nil {
			return err
		}
		wgr.renderPipeline[k] = []interface{}{}
	}
	return nil
}
