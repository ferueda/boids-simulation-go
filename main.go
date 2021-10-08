package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	birdColor = color.RGBA{R: 40, G: 40, B: 43, A: 255}
	bgColor   = color.RGBA{R: 69, G: 123, B: 157, A: 255}
	boids     [boidCount]*Boid
	boidMap   [screenWidth + 1][screenHeight + 1]int
	rwLock    = sync.RWMutex{}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), birdColor)
		screen.Set(int(boid.position.x-1), int(boid.position.y), birdColor)
		screen.Set(int(boid.position.x), int(boid.position.y-1), birdColor)
		screen.Set(int(boid.position.x), int(boid.position.y+1), birdColor)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	// Fills our 2d map (boidMap) with -1
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids Simulation in Go")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
