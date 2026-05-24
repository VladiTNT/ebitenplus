package ebitenplus

// Rect is a rectangle with Min as the top left corner and Max as the bottom right corner.
// It's the same as the standard image.Rectangle but with float64 values.
type Rect struct {
	Min, Max Vec
}

// Rrect makes a Rect.
func Rrect(x0, y0, x1, y1 float64) Rect {
	return Rect{Vvec(x0, y0), Vvec(x1, y1)}
}

// String returns a string representation of r.
func (r Rect) String() string {
	return r.Min.String() + "-" + r.Max.String()
}

// Width returns r's width.
func (r Rect) Width() float64 {
	return r.Max.X - r.Min.X
}

// Height returns r's height.
func (r Rect) Height() float64 {
	return r.Max.Y - r.Min.Y
}

// Size returns r's width and height as a Vec.
func (r Rect) Size() Vec {
	return Vvec(r.Width(), r.Height())
}

// Add returns the r translated by v.
func (r Rect) Add(v Vec) Rect {
	return Rrect(r.Min.X+v.X, r.Min.Y+v.Y, r.Max.X+v.X, r.Max.Y+v.Y)
}

// Sub returns the rectangle r translated -v.
func (r Rect) Sub(v Vec) Rect {
	return Rrect(r.Min.X-v.X, r.Min.Y-v.Y, r.Max.X-v.X, r.Max.Y-v.Y)
}

// Inset returns the rectangle r inset by n, which may be negative. If either
// of r's dimensions is less than 2*n then an empty rectangle near the center
// of r will be returned.
func (r Rect) Inset(f float64) Rect {
	if r.Width() < 2*f {
		r.Min.X = (r.Min.X + r.Max.X) / 2
		r.Max.X = r.Min.X
	} else {
		r.Min.X += f
		r.Max.X -= f
	}
	if r.Height() < 2*f {
		r.Min.Y = (r.Min.Y + r.Max.Y) / 2
		r.Max.Y = r.Min.Y
	} else {
		r.Min.Y += f
		r.Max.Y -= f
	}
	return r
}

// Intersect returns the largest rectangle contained by both r and s. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (r Rect) Intersect(s Rect) Rect {
	if r.Min.X < s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y < s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X > s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y > s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	// Letting r0 and s0 be the values of r and s at the time that the method
	// is called, this next line is equivalent to:
	//
	// if max(r0.Min.X, s0.Min.X) >= min(r0.Max.X, s0.Max.X) || likewiseForY { etc }
	if r.Empty() {
		return Rect{}
	}
	return r
}

// Union returns the smallest rectangle that contains both r and s.
func (r Rect) Union(s Rect) Rect {
	if r.Empty() {
		return s
	}
	if s.Empty() {
		return r
	}
	if r.Min.X > s.Min.X {
		r.Min.X = s.Min.X
	}
	if r.Min.Y > s.Min.Y {
		r.Min.Y = s.Min.Y
	}
	if r.Max.X < s.Max.X {
		r.Max.X = s.Max.X
	}
	if r.Max.Y < s.Max.Y {
		r.Max.Y = s.Max.Y
	}
	return r
}

// Empty reports whether the rectangle contains no points.
func (r Rect) Empty() bool {
	return r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y
}

// Eq reports whether r and s contain the same set of points. All empty
// rectangles are considered equal.
func (r Rect) Eq(s Rect) bool {
	return r == s || r.Empty() && s.Empty()
}

// Overlaps reports whether r and s have a non-empty intersection.
func (r Rect) Overlaps(s Rect) bool {
	return !r.Empty() && !s.Empty() &&
		r.Min.X < s.Max.X && s.Min.X < r.Max.X &&
		r.Min.Y < s.Max.Y && s.Min.Y < r.Max.Y
}

// In reports whether every point in r is in s.
func (r Rect) In(s Rect) bool {
	if r.Empty() {
		return true
	}
	// Note that r.Max is an exclusive bound for r, so that r.In(s)
	// does not require that r.Max.In(s).
	return s.Min.X <= r.Min.X && r.Max.X <= s.Max.X &&
		s.Min.Y <= r.Min.Y && r.Max.Y <= s.Max.Y
}

// Canon returns the canonical version of r. The returned rectangle has minimum
// and maximum coordinates swapped if necessary so that it is well-formed.
func (r Rect) Canon() Rect {
	if r.Max.X < r.Min.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}
	if r.Max.Y < r.Min.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
	return r
}
