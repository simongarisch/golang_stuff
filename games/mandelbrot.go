package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 640
	maxIt        = 128
)

var (
	offscreen    = ebiten.NewImage(screenWidth, screenHeight)
	offscreenPix []byte
	palette      [maxIt]byte
)

func color(it int) (r, g, b byte) {
	if it == maxIt {
		return 0xff, 0xff, 0xff
	}
	c := palette[it]
	return c, c, c
}

func updateOffscreen(centerX, centerY, size float64) {
	for j := 0; j < screenHeight; j++ {
		for i := 0; i < screenHeight; i++ {
			x := float64(i)*size/screenWidth - size/2 + centerX
			y := (screenHeight-float64(j))*size/screenHeight - size/2 + centerY
			c := complex(x, y)
			z := complex(0, 0)
			it := 0
			for ; it < maxIt; it++ {
				z = z*z + c
				if real(z)*real(z)+imag(z)*imag(z) > 4 {
					break
				}
			}
			r, g, b := color(it)
			p := 4 * (i + j*screenWidth)
			offscreenPix[p] = r
			offscreenPix[p+1] = g
			offscreenPix[p+2] = b
			offscreenPix[p+3] = 0xff
		}
	}
	offscreen.ReplacePixels(offscreenPix)
}

func init() {
	offscreenPix = make([]byte, screenWidth*screenHeight*4)
	for i := range palette {
		palette[i] = byte(math.Sqrt(float64(i)/float64(len(palette))) * 0x80)
	}
	// Now it is not feasible to call updateOffscreen every frame due to performance.
	updateOffscreen(-0.75, 0.25, 2)
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(offscreen, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Mandelbrot (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
