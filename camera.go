package ebitenplus

import "github.com/hajimehoshi/ebiten/v2"

// Camera helps you with translations when rendering your sprites.
//
// The position vector is the point that the camera is centered on relative to the other sprites.
// The offset vector offset's the translation by it's value.
// Camera also takes in pointers to the screen's width and height, these are so that it can be centered
// properly and change the zooming of the camera automatically if you want to implement different screen
// sizes.
type Camera struct {
	Pos    Vec
	Offset Vec
	SW, SH *int
}

// NewCamera is a constructor. It makes a camera with the given values.
func NewCamera(pos, off Vec, sw, sh *int) *Camera {
	return &Camera{
		Pos:    pos,
		Offset: off,
		SW:     sw, SH: sh,
	}
}

// SetPos set's the camera's position to that of the given vector.
func (c *Camera) SetPos(v Vec) {
	c.Pos = v
}

// SetOffset set's the camera's offset vector to the given one.
func (c *Camera) SetOffset(v Vec) {
	c.Offset = v
}

// Apply translates v relative to the screen's width and height as well as the camera's position
// and offset.
func (c *Camera) Apply(v Vec) Vec {
	x := v.X - c.Pos.X + c.Offset.X + float64(*c.SW)/2
	y := v.Y - c.Pos.Y + c.Offset.Y + float64(*c.SH)/2
	return Vvec(x, y)
}

// Follow is a basic follow function. It makes the camera smoothly get closer to v.
//
// For an example if you have a player that moves around with it's position being the vector v,
// Follow should be called in the Update() function after the player's movement logic has been
// updated, the result will be a camera that follow's the player smoothly.
//
// The value delta is the speed at which the camera follows v. Slower (smaller value) is smoother.
//
// All the other 'Follow' functions work mostly the same but they do other cool stuff too.
func (c *Camera) Follow(v Vec, delta float64) {
	c.Pos.X += (v.X - c.Pos.X) * delta
	c.Pos.Y += (v.Y - c.Pos.Y) * delta
}

// FollowAhead works like Follow but it follows a target determined by v and it's velocity vel.
// The value maxDist is the rough maximum distance that the target will be ahead.
func (c *Camera) FollowAhead(v, vel Vec, delta, maxDist float64) {
	targetX := v.X + (vel.X * maxDist)
	targetY := v.Y + (vel.Y * maxDist)

	c.Pos.X += (targetX - c.Pos.X) * delta
	c.Pos.Y += (targetY - c.Pos.Y) * delta
}

// FollowCursor works like Follow but it follows v with an offset determined by the cursor position.
//
// The value maxDist is the radius around v after which the camera won't follow the cursor.
func (c *Camera) FollowCursor(v Vec, delta, maxDist float64) {
	x, y := ebiten.CursorPosition()

	targetX := v.X + float64(x) - float64(*c.SW)/2
	targetY := v.Y + float64(y) - float64(*c.SH)/2

	offX := targetX - v.X
	offY := targetY - v.Y

	if offX > maxDist {
		offX = maxDist
	}
	if offX < maxDist {
		offX = -maxDist
	}
	if offY > maxDist {
		offY = maxDist
	}
	if offY < maxDist {
		offY = -maxDist
	}

	c.Pos.X += (v.X + offX - c.Pos.X) * delta
	c.Pos.Y += (v.Y + offY - c.Pos.Y) * delta
}
