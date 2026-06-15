package main

import (
	"time"

	"github.com/VladiTNT/ebitenplus"
	"github.com/VladiTNT/ebitenplus/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	A *ebitenplus.Animation
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := new(ebiten.DrawImageOptions)
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(60, 60)
	screen.DrawImage(g.A.Play(), op)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	g := new(Game)
	al := ebitenplus.NewAssetLoader(internal.TestAssets, func(err error) { panic(err) })
	g.A = ebitenplus.NewAnimation(100*time.Millisecond, al.LoadSpriteSheet("assets/SpriteSheet.png", 16, 16, 16, 4))

	ebiten.SetWindowTitle("Animation example")
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
