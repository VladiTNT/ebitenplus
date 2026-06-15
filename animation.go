package ebitenplus

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Animation is a struct that contains several frames that get played in sequence at the given interval. The field
// Index holds the current index in the array of frames. This animation loop is simple by design as I think it's
// practical to build your own higher level animation system for complex characters.
type Animation struct {
	Frames   []*ebiten.Image
	Interval time.Duration
	Index    int

	lastFrame time.Time
}

// NewAnimation is a constructor, it makes an animation with the given interval and sprites.
func NewAnimation(inerval time.Duration, frames []*ebiten.Image) *Animation {
	return &Animation{
		Frames:   frames,
		Interval: inerval,
		Index:    0,

		lastFrame: time.Now(),
	}
}

// Play returns the current frame in the loop, it's important to know that Play also updates the state of the
// animation, so it should be placed in the draw loop.
func (a *Animation) Play() *ebiten.Image {
	if time.Since(a.lastFrame) >= a.Interval {
		if a.Index >= len(a.Frames)-1 {
			a.Index = 0
		} else {
			a.Index++
		}
		a.lastFrame = time.Now()
	}
	return a.Frames[a.Index]
}
