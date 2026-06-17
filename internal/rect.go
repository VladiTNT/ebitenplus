package internal

import (
	"fmt"
	"math"
)

// Rect is a rectangle with the point defined by X and Y as it's top left corner and W and H as it's width and height.
type Rect struct {
	X, Y, W, H float64
}

// Rrect makes a rect with the given values.
func Rrect(x, y, w, h float64) Rect {
	return Rect{x, y, w, h}
}

// Vrect makes a rect with the given vector and width and height values.
func Vrect(pos Vec, w, h float64) Rect {
	return Rect{pos.X, pos.Y, w, h}
}

// Returns a string representation of r like "(20, 30) - W: 10 - H: 20".
func (r Rect) String() string {
	return fmt.Sprintf("(%f, %f) - W: %f - H: %f", r.X, r.Y, r.W, r.H)
}

// Size returns r's width and height as a Vec.
func (r Rect) Size() Vec {
	return Vvec(r.W, r.H)
}

// MaxX returns r.X + r.W
func (r Rect) MaxX() float64 {
	return r.X + r.W
}

// MaxY returns r.Y + r.H
func (r Rect) MaxY() float64 {
	return r.Y + r.H
}

// Max returns the bottom right corner as a Vec.
func (r Rect) Max() Vec {
	return Vvec(r.MaxX(), r.MaxY())
}

// Empty reports whether there are any points inside of r.
func (r Rect) Empty() bool {
	return r.W <= 0 || r.H <= 0
}

// Eq reports if the rect's are equal.
func (r Rect) Eq(s Rect) bool {
	return r == s || r.Empty() && s.Empty()
}

// Add translates r's base position by v.
func (r Rect) Add(v Vec) Rect {
	return Rrect(r.X+v.X, r.Y+v.Y, r.W, r.H)
}

// Sub subtracts v out of r's base position
func (r Rect) Sub(v Vec) Rect {
	return Rrect(r.X-v.X, r.Y-v.Y, r.W, r.H)
}

// Inset returns the rectangle r inset by n, which may be negative. If either
// of r's dimensions is less than 2*n then an empty rectangle near the center
// of r will be returned.
func (r Rect) Inset(n float64) Rect {
	if r.W < 2*n {
		r.X += r.W / 2
		r.W = 0
	} else {
		r.X += n
		r.W -= 2 * n
	}
	if r.H < 2*n {
		r.Y += r.H / 2
		r.H = 0
	} else {
		r.Y += n
		r.H -= 2 * n
	}
	return r
}

// Intersect returns the largest rectangle contained by both r and s. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (r Rect) Intersect(s Rect) Rect {
	minX := max(r.X, s.X)
	minY := max(r.Y, s.Y)
	maxX := min(r.X+r.W, s.X+s.W)
	maxY := min(r.Y+r.H, s.Y+s.H)

	// If the intersection is valid
	if minX >= maxX || minY >= maxY {
		return Rect{}
	}

	return Rrect(minX, minY, maxX-minX, maxY-minY)
}

// Union returns the smallest rectangle that contains both r and s.
func (r Rect) Union(s Rect) Rect {
	if r.Empty() {
		return s
	}
	if s.Empty() {
		return r
	}

	minX := min(r.X, s.X)
	minY := min(r.Y, s.Y)
	maxX := max(r.X+r.W, s.X+s.W)
	maxY := max(r.Y+r.H, s.Y+s.H)

	return Rrect(minX, minY, maxX-minX, maxY-minY)
}

// Overlaps reports whether r and s have a non-empty intersection.
func (r Rect) Overlaps(s Rect) bool {
	return !r.Empty() && !s.Empty() &&
		r.X < s.MaxX() && s.X < r.MaxX() &&
		r.Y < s.MaxY() && s.Y < r.MaxY()
}

// In reports whether every point in r is in s.
func (r Rect) In(s Rect) bool {
	if r.Empty() {
		return false
	}

	return s.X <= r.X && r.MaxX() <= s.MaxX() &&
		s.Y <= r.Y && r.MaxY() <= s.MaxY()
}

// Canon returns r but with the absolute values of r.W and r.H.
func (r Rect) Canon() Rect {
	return Rrect(r.X, r.Y, math.Abs(r.W), math.Abs(r.H))
}
