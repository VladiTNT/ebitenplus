package main

import (
	"github.com/VladiTNT/ebitenplus"
	"github.com/VladiTNT/ebitenplus/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	AudioContext *audio.Context
	AssetLoader  *ebitenplus.AssetLoader

	Sprite            *ebiten.Image
	CoinPlayer        *audio.Player
	BackOnTrackPlayer *audio.Player
}

func NewGame() *Game {
	audioContext := audio.NewContext(48000)
	assetLoader := ebitenplus.NewAssetLoader(internal.TestAssets, func(err error) { panic(err) })

	sprite := assetLoader.LoadImage("assets/Sprite-0001.png")

	coinBytes := assetLoader.LoadRawAudio("assets/coin.wav", ebitenplus.WAV)
	coinPlayer := audioContext.NewPlayerF32FromBytes(coinBytes)

	coinPlayer.SetVolume(0.2)

	backOnTrackStream, _ := assetLoader.LoadStream("assets/BackOnTrack.mp3", ebitenplus.MP3)

	backOnTrackLoop := audio.NewInfiniteLoopF32(backOnTrackStream, backOnTrackStream.Length())

	backOnTrackPlayer, err := audioContext.NewPlayerF32(backOnTrackLoop)
	if err != nil {
		panic(err)
	}

	backOnTrackPlayer.SetVolume(0.2)

	return &Game{
		AudioContext: audioContext,
		AssetLoader:  assetLoader,

		Sprite:            sprite,
		CoinPlayer:        coinPlayer,
		BackOnTrackPlayer: backOnTrackPlayer,
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.CoinPlayer.Rewind()
		g.CoinPlayer.Play()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && !g.BackOnTrackPlayer.IsPlaying() {
		g.BackOnTrackPlayer.Play()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(40, 40)
	screen.DrawImage(g.Sprite, op)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowTitle("Asset loader example")
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
