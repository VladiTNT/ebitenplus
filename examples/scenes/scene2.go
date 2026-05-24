package main

import (
	"github.com/VladiTNT/ebitenplus"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Scene2 struct {
	Alive bool
	G     *Game

	AudioPlayer *audio.Player
}

func NewScene2(g *Game) *Scene2 {
	coinSfx := g.AssetLoader.LoadRawAudio("assets/coin.wav", ebitenplus.WAV)
	player := g.AudioContext.NewPlayerF32FromBytes(coinSfx)
	player.SetVolume(0.2)

	// Scene constructor
	return &Scene2{
		Alive: true,
		G:     g,

		AudioPlayer: player,
	}
}

func (s2 *Scene2) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		s2.Alive = false
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s2.AudioPlayer.Rewind()
		s2.AudioPlayer.Play()
	}

	return nil
}

func (s2 *Scene2) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Press <TAB> to switch", 0, 0)
	ebitenutil.DebugPrintAt(screen, "Press <ArrowUp> to play coin sound", 0, 12)
}

func (s2 *Scene2) Jump() ebitenplus.Scene {
	if s2.Alive {
		return nil
	} else {
		// Scene destructor
		s2.AudioPlayer.Close()

		return NewScene1(s2.G)
	}
}
