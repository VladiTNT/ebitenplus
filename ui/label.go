package ui

import (
	"github.com/VladiTNT/ebitenplus"
	"github.com/hajimehoshi/ebiten/v2"
)

type Label struct {
	Text string
	Pos  ebitenplus.Vec
}

func NewLabel(text string, pos ebitenplus.Vec) *Label {
	return &Label{
		Text: text,
		Pos:  pos,
	}
}

func (l *Label) Draw(screen *ebiten.Image, printer ebitenplus.Printer) {
	printer.Print(screen, l.Text, l.Pos)
}
