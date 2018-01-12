package core

// Color is an RGB color
type Color struct {
	R float32
	G float32
	B float32
}

func NewColor(r, g, b float32) Color {
	return Color{R: r, G: g, B: b}
}
