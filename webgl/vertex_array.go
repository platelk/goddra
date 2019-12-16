package webgl

import "syscall/js"

type VertexArray struct {
	wg *WebGL
	value js.Value
}

func (va *VertexArray) Bind() {
	va.wg.BindVertexArray(va)
}