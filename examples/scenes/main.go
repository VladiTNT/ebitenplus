package main

import (
	"github.com/VladiTNT/ebitenplus"
	"github.com/VladiTNT/ebitenplus/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

// When working with scenes, the Game struct acts as a global manager for the most important things for
// your game (ex: screen dimensions, asset loader, audio context, etc.). You also need for it to keep
// track of the current scene (duhhh).
type Game struct {
	ScreenWidth  int
	ScreenHeight int

	AssetLoader  *ebitenplus.AssetLoader
	AudioContext *audio.Context

	CurrentScene ebitenplus.Scene
}

func NewGame() *Game {
	g := new(Game)

	g.ScreenWidth = 320
	g.ScreenHeight = 240
	g.AssetLoader = ebitenplus.NewAssetLoader(internal.TestAssets, func(err error) { panic(err) })
	g.AudioContext = audio.NewContext(48000)
	g.CurrentScene = NewScene1(g)

	return g
}

func (g *Game) Update() error {
	err := g.CurrentScene.Update()

	if scene := g.CurrentScene.Jump(); scene != nil {
		g.CurrentScene = scene
	}

	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.CurrentScene.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func main() {
	ebiten.SetWindowTitle("Multiple scenes example")
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
