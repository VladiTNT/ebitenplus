package ebitenplus

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)

// PrintString is a very low level function for text rendering. You shouldn't really use it directly.
// It's meant more to be used with a wrapper kinda like the Printer struct.
func PrintString(
	screen *ebiten.Image, str string,
	font *text.GoTextFaceSource, direction text.Direction, fontSize float64, lang language.Tag,
	opts *text.DrawOptions,
) {
	text.Draw(screen, str, &text.GoTextFace{
		Source:    font,
		Direction: direction,
		Size:      fontSize,
		Language:  lang,
	}, opts)
}

// A Printer is basically just a wrapper around the PrintString function for easier use.
type Printer struct {
	Font      *text.GoTextFaceSource
	Direction text.Direction
	FontSize  float64
	Lang      language.Tag

	LineSpacingMultiplier float64
}

// NewPrinter is a constructor. It makes a Printer with the given values.
func NewPrinter(font *text.GoTextFaceSource, dir text.Direction, fontSize float64, lang language.Tag) *Printer {
	return &Printer{
		Font:      font,
		Direction: dir,
		FontSize:  fontSize,
		Lang:      lang,

		LineSpacingMultiplier: 1.2,
	}
}

// Prints a string with the given options.
func (pp *Printer) PrintWithOpts(screen *ebiten.Image, str string, opts *text.DrawOptions) {
	PrintString(screen, str, pp.Font, pp.Direction, pp.FontSize, pp.Lang, opts)
}

// Print just prints the string at the given position.
func (pp *Printer) Print(screen *ebiten.Image, str string, pos Vec) {
	op := new(text.DrawOptions)
	op.LineSpacing = pp.LineSpacingMultiplier * pp.FontSize
	op.GeoM.Translate(pos.Coords())

	pp.PrintWithOpts(screen, str, op)
}

// PrintCol prints the string with the given color.
func (pp *Printer) PrintCol(screen *ebiten.Image, str string, pos Vec, col color.Color) {
	op := new(text.DrawOptions)
	op.LineSpacing = pp.LineSpacingMultiplier * pp.FontSize
	op.GeoM.Translate(pos.Coords())
	op.ColorScale.ScaleWithColor(col)

	pp.PrintWithOpts(screen, str, op)
}

// PrintSize prints the string but it over rides the size values.
func (pp *Printer) PrintSize(screen *ebiten.Image, str string, pos Vec, fontSize float64) {
	op := new(text.DrawOptions)
	op.LineSpacing = pp.LineSpacingMultiplier * fontSize
	op.GeoM.Translate(pos.Coords())

	PrintString(screen, str, pp.Font, pp.Direction, fontSize, pp.Lang, op)
}
