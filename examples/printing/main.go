package main

import (
	"github.com/VladiTNT/ebitenplus"
	"github.com/VladiTNT/ebitenplus/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)

type Game struct {
	Printer *ebitenplus.Printer
}

func NewGame() *Game {
	al := ebitenplus.NewAssetLoader(internal.TestAssets, func(err error) { panic(err) })
	return &Game{
		Printer: ebitenplus.NewPrinter(
			al.LoadFont("assets/SpaceMono.ttf"),
			text.DirectionLeftToRight,
			24, language.AmericanEnglish,
		),
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Printer.Print(screen, "Hello!!!", ebitenplus.Vvec(0, 0))
	g.Printer.PrintCol(screen, "You are a nerd!!!", ebitenplus.Vvec(0, 24), ebitenplus.GREEN)

	op := new(text.DrawOptions)
	op.GeoM.Translate(0, 48)
	op.ColorScale.ScaleWithColor(ebitenplus.RED)

	ebitenplus.PrintString(
		screen,
		"From PrintString",
		g.Printer.Font,
		text.DirectionLeftToRight,
		12, language.AmericanEnglish, op,
	)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 640, 360
}

func main() {
	ebiten.SetWindowTitle("Printing example")
	ebiten.SetWindowSize(640, 360)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
