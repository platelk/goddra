package color

type Color [4]float32

var Grey = New(0.8, 0.8, 0.8, 1.0)

func New(r, g, b, a float32) Color {
	return Color{r, g, b, a}
}

func (c Color) R() float32 {
	return c[0]
}

func (c Color) G() float32 {
	return c[1]
}

func (c Color) B() float32 {
	return c[2]
}

func (c Color) A() float32 {
	return c[3]
}
