package ebitenplus

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// ParseSpriteSheet takes in a sprite sheet and returns an array of all of the individual sprites.
// Btw spN means number of sprites to parse and sheetWidth uses sprite lengths as the unit.
func ParseSpriteSheet(src *ebiten.Image, spN, spLen, spHeight, sheetWidth int) []*ebiten.Image {
	var sprites []*ebiten.Image

	for i := range spN {
		col := i % sheetWidth
		row := i / sheetWidth

		x0 := col * spLen
		y0 := row * spHeight
		x1 := x0 + spLen
		y1 := y0 + spHeight

		rect := image.Rect(x0, y0, x1, y1)

		subImg := src.SubImage(rect).(*ebiten.Image)
		sprites = append(sprites, subImg)
	}

	return sprites
}
