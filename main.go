package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	dirUp    = Point{x: 0, y: -1}
	dirDown  = Point{x: 0, y: 1}
	dirLeft  = Point{x: -1, y: 0}
	dirRight = Point{x: 1, y: 0}
)

const (
	gameSpeed    = time.Second / 6
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 20
)

type Point struct {
	x, y int
}

type Game struct {
	snake      []Point
	direction  Point
	lastUpdate time.Time
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.direction = dirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.direction = dirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.direction = dirRight
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.direction = dirLeft
	}

	if time.Since(g.lastUpdate) < gameSpeed {
		return nil
	}
	g.lastUpdate = time.Now()

	g.updateSnake(&g.snake, g.direction)
	return nil
}

func (g *Game) updateSnake(snake *[]Point, direction Point) {
	head := (*snake)[0]

	newHead := Point{
		x: head.x + direction.x,
		y: head.y + direction.y,
	}

	*snake = append(
		[]Point{newHead},
		(*snake)[:len(*snake)-1]...,
	)
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, p := range g.snake {
		vector.DrawFilledRect(
			screen,
			float32(p.x*gridSize),
			float32(p.y*gridSize),
			gridSize,
			gridSize,
			color.White,
			true,
		)
	}
}

func (g *Game) Layout(
	outsideWidth,
	outsideHeight int,
) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{
		snake: []Point{{
			x: screenWidth / gridSize / 2,
			y: screenHeight / gridSize / 2,
		}},
		direction: Point{x: 1, y: 0},
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Another Snake Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
