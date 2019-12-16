package geom

type Vec interface {
	Values() []float32
	Dim() int
}

type Vec3 []float32

func NewVec3(x, y, z float32) Vec3 {
	return Vec3{x, y, z}
}

func (p Vec3) Values() []float32 {
	return p
}

func (p Vec3) Dim() int {
	return 3
}

type Vec2 Vec3

func NewVec2(x, y float32) Vec2 {
	return Vec2{x, y}
}

func (p Vec2) Dim() int {
	return 2
}

func (p Vec2) Values() []float32 {
	return p[:2]
}
