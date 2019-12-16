package webgl

import (
	"syscall/js"
)

type Buffer struct {
	value js.Value
	wg    *WebGL
	t     GLType
}

func (b *Buffer) Bind(t GLType) {
	b.t = t
	b.wg.BindBuffer(t, b)
}

func (b *Buffer) BindToArrayBuffer() {
	b.Bind(b.wg.types.ArrayBuffer)
}

func (b *Buffer) Float32ArrayData(data Float32Array, g GLType) {
	b.wg.BufferData(b.t, js.Value(data), g)
}
