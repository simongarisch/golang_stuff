package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct{}

func (g *Game) Update() error {
	// update the logical state
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// render the screen
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenwidth, screenheight int) {
	// return the game screen size
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
