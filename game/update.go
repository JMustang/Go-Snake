package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.reset()
		}
		return nil
	}

	g.handleInput()

	if time.Since(g.lastUpdate) < GameSpeed {
		return nil
	}
	g.lastUpdate = time.Now()

	g.updateSnake()
	return nil
}

func (g *Game) handleInput() {
	newDir := g.direction
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.direction != DirDown {
		newDir = DirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.direction != DirUp {
		newDir = DirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.direction != DirLeft {
		newDir = DirRight
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.direction != DirRight {
		newDir = DirLeft
	}
	g.direction = newDir
}
