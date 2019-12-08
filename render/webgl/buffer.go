package webgl

import "syscall/js"

type Buffer struct {
	value js.Value
	wg *WebGL
	t GLType
}

func (b *Buffer) Bind(t GLType) {
	b.t = t
	b.wg.BindBuffer(t, b)
}

func (b *Buffer) Data(data interface{}, g GLType) {
	b.wg.BufferData(b.t, data, g)
}
