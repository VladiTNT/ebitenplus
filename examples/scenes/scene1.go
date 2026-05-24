package main

import (
	"github.com/VladiTNT/ebitenplus"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Each scene must have some sort of scope value to help the Jump method know when it should jump to
// the next scene. Usually a boolean is good but if your scene can lead to multiple other scenes, an
// enum might be better.
type Scene1 struct {
	Alive bool
	G     *Game

	Sprite *ebiten.Image
	Pos    ebitenplus.Vec
}

func NewScene1(g *Game) *Scene1 {
	// Scene constructor
	return &Scene1{
		Alive: true,
		G:     g,

		Sprite: g.AssetLoader.LoadImage("assets/Sprite-0001.png"),
		Pos:    ebitenplus.Vvec(0, 0),
	}
}

func (s1 *Scene1) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		s1.Alive = false
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		s1.Pos.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		s1.Pos.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		s1.Pos.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		s1.Pos.X -= 2
	}

	return nil
}

func (s1 *Scene1) Draw(screen *ebiten.Image) {
	op := new(ebiten.DrawImageOptions)
	op.GeoM.Translate(s1.Pos.Coords())
	screen.DrawImage(s1.Sprite, op)

	ebitenutil.DebugPrintAt(screen, "Press <TAB> to switch", 0, 0)
	ebitenutil.DebugPrintAt(screen, "Press <ArrowKeys> to move", 0, 12)
	ebitenutil.DebugPrintAt(screen, "Press <ESC> to exit", 0, 24)
}

// Jump returns nil if it's not time to switch scenes yet. When it's time to change, Jump should first run
// the cleanup logic and then return the constructor of the next scene. That's basically the whole
// pattern.
func (s1 *Scene1) Jump() ebitenplus.Scene {
	if s1.Alive {
		return nil
	} else {
		// Scene destructor
		s1.Sprite.Deallocate()

		return NewScene2(s1.G)
	}
}
