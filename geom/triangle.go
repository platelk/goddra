package geom

type Triangle struct {
	position Vec2
	coord    Coord2D
	point    []float32
}

func NewTriangle(p1, p2, p3 Vec2) *Triangle {
	var p []float32
	for _, point := range []Vec2{p1, p2, p3} {
		p = append(p, point...)
	}
	return &Triangle{
		coord:    Coord2D{p1, p2, p3},
		point:    p,
		position: NewVec2(0, 0),
	}
}

func (t *Triangle) Position() Vec {
	return t.position
}

func (t *Triangle) SetPosition(p Vec) {
	values := p.Values()
	t.position = NewVec2(values[0], values[1])
}

func (t *Triangle) Points() []float32 {
	return t.point
}

func (t *Triangle) Dim() int {
	return 2
}


