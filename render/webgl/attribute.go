package webgl

import "syscall/js"

type Attribute struct {
	value js.Value
	prog *Program
	wg *WebGL
	name string
}
