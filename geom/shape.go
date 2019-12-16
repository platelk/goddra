package geom

type Shape interface {
	Points() []float32
	Position() Vec
	SetPosition(p Vec)
	Dim() int
}
