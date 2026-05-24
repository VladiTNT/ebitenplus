package main

import (
	"fmt"

	"github.com/VladiTNT/ebitenplus"
	"github.com/VladiTNT/ebitenplus/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Follow string

const (
	BasicFollow  Follow = "BasicFollow"
	FollowAhead  Follow = "FollowAhead"
	FollowCursor Follow = "FollowCursor"
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Sprite1      *ebiten.Image
	Pos          ebitenplus.Vec
	Sprite2      *ebiten.Image
	Cam          *ebitenplus.Camera
	CameraFollow Follow
}

func NewGame() *Game {
	g := new(Game)

	al := ebitenplus.NewAssetLoader(internal.TestAssets, func(err error) { panic(err) })

	g.Sprite1 = al.LoadImage("assets/Sprite-0001.png")
	g.Sprite2 = al.LoadImage("assets/Sprite-0002.png")

	g.Pos = ebitenplus.Vvec(0, 0)

	g.ScreenWidth = 320
	g.ScreenHeight = 240

	g.Cam = ebitenplus.NewCamera(g.Pos, ebitenplus.Vvec(-16, -16), &g.ScreenWidth, &g.ScreenHeight)

	g.CameraFollow = BasicFollow

	return g
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		g.CameraFollow = BasicFollow
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		g.CameraFollow = FollowAhead
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.CameraFollow = FollowCursor
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Pos.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Pos.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Pos.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Pos.X -= 2
	}

	switch g.CameraFollow {
	case BasicFollow:
		g.Cam.Follow(g.Pos, 0.05)
	case FollowAhead:
		g.Cam.FollowAhead(g.Pos, ebitenplus.Vvec(0, 0), 0.05, 20)
	case FollowCursor:
		g.Cam.FollowCursor(g.Pos, 0.03, 25)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op1 := new(ebiten.DrawImageOptions)
	op1.GeoM.Translate(g.Cam.Apply(g.Pos).Coords())

	op2 := new(ebiten.DrawImageOptions)
	op2.GeoM.Translate(g.Cam.Apply(ebitenplus.Vvec(60, 60)).Coords())

	screen.DrawImage(g.Sprite2, op2)
	screen.DrawImage(g.Sprite1, op1)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FollowMode: %s", g.CameraFollow), 0, 0)
	ebitenutil.DebugPrintAt(screen, "I - BasicFollow, O - FollowAhead, P - FollowCursor", 0, 12)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PlayerX: %f, PlayerY: %f", g.Pos.X, g.Pos.Y), 0, 24)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("CamX: %f, CamY: %f", g.Cam.Pos.X, g.Cam.Pos.Y), 0, 36)
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func main() {
	ebiten.SetWindowTitle("Camera example")
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
