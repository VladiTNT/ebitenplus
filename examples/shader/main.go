package main

import (
	"github.com/VladiTNT/ebitenplus"
	"github.com/VladiTNT/ebitenplus/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Shader *ebiten.Shader
	Img    *ebiten.Image
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := g.Img.Bounds().Dx(), g.Img.Bounds().Dy()

	op := new(ebiten.DrawRectShaderOptions)

	op.GeoM.Translate(40, 40)

	op.Images[0] = g.Img

	screen.DrawRectShader(w, h, g.Shader, op)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	g := new(Game)
	al := ebitenplus.NewAssetLoader(internal.TestAssets, func(err error) { panic(err) })
	g.Img = al.LoadImage("assets/Sprite-0001.png")
	g.Shader = al.LoadShader("assets/shader.kage")

	ebiten.SetWindowTitle("Shader example")
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
