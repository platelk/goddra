package geom

type Rectangle struct {
	position Vec2
	coord    Coord2D
	point    []float32
}

func NewRectangle(width, height float32) *Rectangle {
	var p []float32
	for _, point := range []Vec2{{0, 0}, {width, 0}, {0, height}, {0, height}, {width, 0}, {width, height}} {
		p = append(p, point...)
	}
	return &Rectangle{
		coord:    Coord2D{{0, 0}, {width, 0}, {width, height}, {0, height}},
		point:    p,
		position: NewVec2(0, 0),
	}
}

func (r *Rectangle) SetPosition(p Vec) {
	values := p.Values()
	r.position = NewVec2(values[0], values[1])
}

func (r *Rectangle) Position() Vec {
	return r.position
}

func (r *Rectangle) Points() []float32 {
	return r.point
}

func (r *Rectangle) Dim() int {
	return 2
}


