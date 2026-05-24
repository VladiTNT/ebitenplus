package ebitenplus

import "github.com/hajimehoshi/ebiten/v2"

// A scene is an isolated environment with Update and Draw functions (much like the ebiten.Game
// interface) that has a Jump method which jumps to the next scene. When making games you are going
// to have to separate your menu screens from your levels / actual game world.
//
// Check out the examples directory to see how you are supposed to compose scenes and run
// cleanup / initialization logic.
type Scene interface {
	Draw(screen *ebiten.Image)
	Update() error
	Jump() Scene
}
