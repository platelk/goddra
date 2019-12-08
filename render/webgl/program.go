package webgl

import "syscall/js"

type Program struct {
	wg *WebGL
	value js.Value
}

func (p *Program) AttachShader(vertex, fragment *Shader) {
	p.wg.AttachShader(p, vertex)
	p.wg.AttachShader(p, fragment)
}

func (p *Program) GetParameter(parameter GLType) js.Value {
	return p.wg.GetProgramParameter(p, parameter)
}

func (p *Program) GetLinkStatus() bool {
	return p.wg.GetProgramLinkStatus(p)
}

func (p *Program) GetInfoLog() string {
	return p.wg.GetProgramInfoLog(p)
}

func (p *Program) Link() {
	p.wg.LinkProgram(p)
}

func (p *Program) GetAttribute(attributeName string) *Attribute {
	return p.wg.GetAttribLocation(p, attributeName)
}