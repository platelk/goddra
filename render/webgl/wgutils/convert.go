package wgutils

import (
	"goddra/render/webgl"
	"syscall/js"
)

func ToFloat32Array(data []float32) webgl.Float32Array {
	farr := js.Global().Get("Float32Array").New(len(data))
	for i, v := range data {
		farr.SetIndex(i, v)
	}
	return webgl.Float32Array(farr)
}
