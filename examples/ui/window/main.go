package main

import (
	"github.com/VladiTNT/ebitenplus"
	"github.com/VladiTNT/ebitenplus/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	W *ui.Window
}

func (g *Game) Update() error {
	g.W.Update(ebitenplus.Ivec(ebiten.CursorPosition()), ebiten.MouseButton0)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.W.Draw(screen, 2, ui.LIGHT_GRAY, ui.GRAY)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	g := new(Game)
	g.W = ui.NewWindow(ebitenplus.Rrect(20, 20, 120, 120))

	ebiten.SetWindowTitle("Window example")
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
