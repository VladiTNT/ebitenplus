package ebitenplus

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Field is an input field. By itself it only updates the proprieties IsHover, Click and IsHeld with it's Update method.
// It is meant as a base for other UI components. It's more or less just a boring button.
type Field struct {
	// Box is the space in which the field checks for input.
	Box Rect

	// IsHover will be true when the mouse is hovering over the button.
	IsHover bool
	// Click will be true on the exact tick when the field was clicked.
	Click bool
	// IsHeld will be true as long as the field is being pressed. IsHeld resets if the mouse leaves the space.
	IsHeld bool
}

// NewField is a constructor. It makes a field in the given space.
func NewField(box Rect) *Field {
	return &Field{
		Box: box,

		IsHover: false,
		Click:   false,
		IsHeld:  false,
	}
}

// Update update's the field's properties with the given cursor position and mouse button.
func (f *Field) Update(cursor Vec, m ebiten.MouseButton) {
	// Check if button is hovered
	if cursor.In(f.Box) {
		f.IsHover = true
	} else {
		f.IsHover = false
	}

	// Check if button is clicked
	if f.IsHover {
		// Clicked on this tick
		if inpututil.IsMouseButtonJustPressed(m) {
			f.Click = true
			// Click check sets held on the same frame
			f.IsHeld = true
		} else {
			f.Click = false
		}

		// Held check
		if f.IsHeld && !inpututil.IsMouseButtonJustReleased(m) {
			f.IsHeld = true
		} else {
			f.IsHeld = false
		}
	} else {
		f.IsHeld = false
	}
}

// DrawRect is mostly a debug method. It draws the field's box with the given values.
func (f *Field) DrawRect(screen *ebiten.Image, width float32, outline, fill color.Color) {
	x, y, w, h := float32(f.Box.X), float32(f.Box.Y), float32(f.Box.W), float32(f.Box.H)

	vector.FillRect(screen, x, y, w, h, fill, false)
	vector.StrokeRect(screen, x, y, w, h, width, outline, false)
}
