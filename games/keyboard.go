package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/keyboard/keyboard"
	rkeyabord "github.com/hajimehoshi/ebiten/examples/resources/images/keyboard"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var keyboardImage *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(rkeyabord.Keyboard_png))
	if err != nil {
		log.Fatal(err)
	}

	keyboardImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	pressed []ebiten.Key
}

func (g *Game) Update() error {
	g.pressed = nil
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			g.pressed = append(g.pressed, k)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const (
		offsetX = 24
		offsetY = 40
	)

	// Draw the base (grayed) keyboard image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(offsetX, offsetY)
	op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	screen.DrawImage(keyboardImage, op)

	// Draw the highlighted keys.
	op = &ebiten.DrawImageOptions{}
	for _, p := range g.pressed {
		op.GeoM.Reset()
		r, ok := keyboard.KeyRect(p)
		if !ok {
			continue
		}
		op.GeoM.Translate(float64(r.Min.X), float64(r.Min.Y))
		op.GeoM.Translate(offsetX, offsetY)
		screen.DrawImage(keyboardImage.SubImage(r).(*ebiten.Image), op)
	}

	keyStrs := []string{}
	for _, p := range g.pressed {
		keyStrs = append(keyStrs, p.String())
	}
	ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Keyboard (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
