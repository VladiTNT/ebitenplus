package ui

import (
	"image/color"

	"github.com/VladiTNT/ebitenplus"
	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	Field *Field

	oldCursor            ebitenplus.Vec
	bottomLeftCornerSens float64
	maxWidth             float64
	maxHeight            float64
}

func NewWindow(box ebitenplus.Rect) *Window {
	return &Window{
		Field: NewField(box),

		oldCursor:            ebitenplus.Ivec(ebiten.CursorPosition()),
		bottomLeftCornerSens: 30,
		maxWidth:             50,
		maxHeight:            50,
	}
}

func (w *Window) Update(cursor ebitenplus.Vec, m ebiten.MouseButton) {
	w.Field.Update(w.oldCursor, m)

	w.oldCursor = cursor
}

func (w *Window) Draw(screen *ebiten.Image, width float32, outline, fill color.Color) {
	w.Field.DrawRect(screen, width, outline, fill)
}
