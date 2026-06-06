package ebitenplus

import (
	"strconv"
)

// AABB is the clasic AABB collision check logic that checks if the point (x, y) is in the rectangle
// ((X0, Y0), (X1, Y1)).
func AABB(x, y, X0, X1, Y0, Y1 float64) bool {
	return x >= X0 && x <= X1 && y >= Y0 && y <= Y1
}

// A Vec is an X, Y float64 coordinate pair. The axes increase right and down.
type Vec struct {
	X, Y float64
}

// Ivec returns a Vec from the given ints.
func Ivec(x, y int) Vec {
	return Vec{float64(x), float64(y)}
}

// Vvec is a shorthand for Vec{x, y}.
func Vvec(x, y float64) Vec {
	return Vec{x, y}
}

// String returns a string representation of v like "(3.14, 2.71)".
func (v Vec) String() string {
	return "(" + strconv.FormatFloat(v.X, 'f', -1, 64) + ", " + strconv.FormatFloat(v.Y, 'f', -1, 64) + ")"
}

// Coords returns the X and Y coordinate pair.
func (v Vec) Coords() (float64, float64) {
	return v.X, v.Y
}

// Add returns the vector v+q.
func (v Vec) Add(q Vec) Vec {
	return Vvec(v.X+q.X, v.Y+q.Y)
}

// Div returns the vector v/k.
func (v Vec) Div(k float64) Vec {
	return Vvec(v.X/k, v.Y/k)
}

// Eq reports whether v and q are equal.
func (v Vec) Eq(q Vec) bool {
	return v.X == q.X && v.Y == q.Y
}

// In reports whether v is in r.
func (v Vec) In(r Rect) bool {
	return v.X >= r.Min.X && v.X <= r.Max.X && v.Y >= r.Min.Y && v.Y <= r.Max.Y
}

// Mul returns the vector v*k.
func (v Vec) Mul(k float64) Vec {
	return Vvec(v.X*k, v.Y*k)
}

// Sub returns the vector v-q.
func (v Vec) Sub(q Vec) Vec {
	return Vvec(v.X-q.X, v.Y-q.Y)
}
