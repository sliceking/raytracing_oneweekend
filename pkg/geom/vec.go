package geom

import (
	"fmt"
	"io"
	"math"
)

// Vec represents a 3-element vector
type Vec struct {
	E [3]float64
}

// NewVec creates a new vec from 3 floats
func NewVec(e0, e1, e2 float64) (v Vec) {
	v.E[0] = e0
	v.E[1] = e1
	v.E[2] = e2
	return
}

// X returns the first element
func (v Vec) X() float64 {
	return v.E[0]
}

// Y returns the second element
func (v Vec) Y() float64 {
	return v.E[1]
}

// Z returns the third element
func (v Vec) Z() float64 {
	return v.E[2]
}

// Inv returns this vectors inverse as a new vector
func (v Vec) Inv() Vec {
	return NewVec(-v.E[0], -v.E[1], -v.E[2])
}

// Len returns the vector's length
func (v Vec) Len() float64 {
	return math.Sqrt(v.LenSq())
}

// LenSq returns the square of the vectors length
func (v Vec) LenSq() float64 {
	return v.E[0]*v.E[0] + v.E[1]*v.E[1] + v.E[2]*v.E[2]
}

// ToUnit coverts a vector to a unit vector
func (v Vec) ToUnit() (u Unit) {
	k := 1.0 / v.Len()
	u.E[0] = v.E[0] * k
	u.E[1] = v.E[1] * k
	u.E[2] = v.E[2] * k
	return
}

// IStream streams in space-separated vector values from a Reader
func (v Vec) IStream(r io.Reader) error {
	_, err := fmt.Fscan(r, v.E[0], v.E[1], v.E[2])
	return err
}

// OStream writes space-separated vector values to a Writer
func (v Vec) OStream(w io.Writer) error {
	_, err := fmt.Fprint(w, v.E[0], v.E[1], v.E[2])
	return err
}

// Plus returns the sum of two vectors
func (v Vec) Plus(v2 Vec) Vec {
	return NewVec(v.E[0]+v2.E[0], v.E[1]+v2.E[1], v.E[2]+v2.E[2])
}

// Minus returns the difference of two vectors
func (v Vec) Minus(v2 Vec) Vec {
	return NewVec(v.E[0]-v2.E[0], v.E[1]-v2.E[1], v.E[2]-v2.E[2])
}

// Times returns the product of two vectors
func (v Vec) Times(v2 Vec) Vec {
	return NewVec(v.E[0]*v2.E[0], v.E[1]*v2.E[1], v.E[2]*v2.E[2])
}

// Div returns the division of two vectors
func (v Vec) Div(v2 Vec) Vec {
	return NewVec(v.E[0]/v2.E[0], v.E[1]/v2.E[1], v.E[2]/v2.E[2])
}

// Scaled returns a vector scaled by a scalar.
func (v Vec) Scaled(n float64) Vec {
	return NewVec(v.E[0]*n, v.E[1]*n, v.E[2]*n)
}

// Cross returns the cross product of two vectors.
func (v Vec) Cross(v2 Vec) Vec {
	return NewVec(
		v.E[1]*v2.E[2]-v.E[2]*v2.E[1],
		v.E[2]*v2.E[0]-v.E[0]*v2.E[2],
		v.E[0]*v2.E[1]-v.E[1]*v2.E[0],
	)
}
