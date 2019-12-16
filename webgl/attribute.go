package webgl

import "syscall/js"

type Attribute struct {
	value js.Value
	prog *Program
	wg *WebGL
	name string
}

func (attr *Attribute) SetUniform(values ...float32) {
	switch len(values) {
	case 2:
		attr.Uniform2f(values[0], values[1])
	case 4:
		attr.Uniform4f(values[0], values[1], values[2], values[3])
	}
}

func (attr *Attribute) Uniform4f(p1, p2, p3, p4 float32) {
	attr.wg.gl.Call("uniform4f", attr.value, p1, p2, p3, p4)
}

func (attr *Attribute) Uniform2f(p1, p2 float32) {
	attr.wg.gl.Call("uniform2f", attr.value, p1, p2)
}