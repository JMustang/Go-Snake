package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	snake      []Point
	direction  Direction
	lastUpdate time.Time
	food       Point
	gameOver   bool
	score      int
	fontFace   *text.GoTextFace
}

func New(fontSource *text.GoTextFaceSource) *Game {
	g := &Game{
		snake: make([]Point, InitialSize),
		fontFace: &text.GoTextFace{
			Source: fontSource,
			Size:   30,
		},
	}

	g.snake[0] = Point{
		X: ScreenWidth / GridSize / 2,
		Y: ScreenHeight / GridSize / 2,
	}
	g.direction = DirRight
	g.spawnFood()

	return g
}
