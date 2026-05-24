package main

import (
"image"
"log"
"math"

"github.com/hajimehoshi/ebiten/v2"
)

const (
screenWidth  = 600
screenHeight = 400
)

type Game struct {
ticks int
}

func (g *Game) Update() error {
g.ticks++
return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))

t := float64(g.ticks) * 0.05
soundFreq1 := 0.05
soundFreq2 := 0.08
centerY := float64(screenHeight) / 2
amplitude := float64(screenHeight) / 4

for x := 0; x < screenWidth; x++ {
waveY := math.Sin(float64(x)*soundFreq1+t) + math.Sin(float64(x)*soundFreq2-t*0.5)
targetY := int(centerY + waveY*amplitude)

if targetY >= 0 && targetY < screenHeight {
r := uint8(127 + 127*math.Sin(float64(x)*0.01+t))
g := uint8(127 + 127*math.Sin(float64(x)*0.015+t))
b := uint8(127 + 127*math.Cos(float64(x)*0.02+t))

pos := (targetY*screenWidth + x) * 4
img.Pix[pos] = r
img.Pix[pos+1] = g
img.Pix[pos+2] = b
img.Pix[pos+3] = 255
}
}

screen.WritePixels(img.Pix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
return screenWidth, screenHeight
}

func main() {
ebiten.SetWindowSize(screenWidth, screenHeight)
ebiten.SetWindowTitle("Light & Sound Simulator (100% Pure Go)")

game := &Game{}
if err := ebiten.RunGame(game); err != nil {
log.Fatal(err)
}
}
