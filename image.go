package ebitenplus

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
)

// DecodeImage calls image.Decode and ebiten.NewImageFromImage to get an ebiten.Image from r.
func DecodeImage(r io.Reader) (*ebiten.Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
