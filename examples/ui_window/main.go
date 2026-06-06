package main

import (
	"fmt"

	"github.com/VladiTNT/ebitenplus"
	"github.com/VladiTNT/ebitenplus/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	F *ui.Field
}

func NewGame() *Game {
	return &Game{
		F: ui.NewField(ebitenplus.Rrect(100, 100, 160, 160)),
	}
}

func (g *Game) Update() error {
	g.F.Update(ebitenplus.Ivec(ebiten.CursorPosition()), ebiten.MouseButton0)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.F.DrawRect(screen, 2, ui.AQUA, ui.GREEN)

	s := fmt.Sprintf("Hover: %t\nClicked: %t\nHeld: %t", g.F.IsHover, g.F.Click, g.F.IsHeld)
	ebitenutil.DebugPrintAt(screen, s, 0, 0)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowTitle("Ui window example")
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
