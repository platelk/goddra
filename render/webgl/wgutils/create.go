package wgutils

import (
	"fmt"
	"goddra/render/webgl"
)

func CreateProgram(gl *webgl.WebGL, vertexShader, fragShader *webgl.Shader) (*webgl.Program, error) {
	prog := gl.CreateProgram()
	prog.AttachShader(vertexShader, fragShader)
	prog.Link()
	if prog.GetLinkStatus() {
		return prog, nil
	}
	info := prog.GetInfoLog()
	gl.DeleteProgram(prog)
	return nil, fmt.Errorf("can't compile program: %s", info)
}

func CreateShader(gl *webgl.WebGL, shaderType webgl.GLType, source string) (*webgl.Shader, error) {
	shader := gl.CreateShader(shaderType)
	shader.Source(source)
	if shader.GetCompileStatus() {
		return shader, nil
	}
	info := shader.GetInfoLog()
	gl.DeleteShader(shader)
	return nil, fmt.Errorf("can't compile shader: %s", info)
}
