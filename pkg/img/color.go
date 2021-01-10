package img

import "github.com/sliceking/raytracing_oneweekend/pkg/geom"

// Color represents RGB color value
type Color struct {
	geom.Vec
}

// NewColor creates a new color from 3 floats
func NewColor(e0, e1, e2 float64) (c Color) {
	c.E[0] = e0
	c.E[1] = e1
	c.E[2] = e2
	return
}

// R returns the first element red
func (c Color) R() float64 {
	return c.E[0]
}

// G returns the second element green
func (c Color) G() float64 {
	return c.E[1]
}

// B returns the third element blue
func (c Color) B() float64 {
	return c.E[2]
}
