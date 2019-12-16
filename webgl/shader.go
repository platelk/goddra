package webgl

import "syscall/js"

type Shader struct {
	source string
	value js.Value
	wg *WebGL
}

func (s *Shader) GetParameter(parameter GLType) js.Value {
	return s.wg.GetShaderParameter(s, parameter)
}

func (s *Shader) GetCompileStatus() bool {
	return s.wg.GetShaderCompileStatus(s)
}

func (s *Shader) GetInfoLog() string {
	return s.wg.GetShaderInfoLog(s)
}

func (s *Shader) Source(source string) {
	s.wg.ShaderSource(s, source)
}

func (s *Shader) Compile() {
	s.wg.CompileShader(s)
}
