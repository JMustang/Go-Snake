package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	for _, p := range g.snake {
		vector.DrawFilledRect(
			screen,
			float32(p.X*GridSize),
			float32(p.Y*GridSize),
			GridSize,
			GridSize,
			ColorSnake,
			true,
		)
	}

	vector.DrawFilledRect(
		screen,
		float32(g.food.X*GridSize),
		float32(g.food.Y*GridSize),
		GridSize,
		GridSize,
		ColorFood,
		true,
	)

	scoreText := fmt.Sprintf("Score: %d", g.score)
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(ColorText)
	text.Draw(screen, scoreText, g.fontFace, op)

	if g.gameOver {
		gameOverText := "Game Over! Press SPACE to restart"
		w, h := text.Measure(gameOverText, g.fontFace, g.fontFace.Size)
		op := &text.DrawOptions{}
		op.GeoM.Translate(ScreenWidth/2-w/2, ScreenHeight/2-h/2)
		op.ColorScale.ScaleWithColor(ColorText)
		text.Draw(screen, gameOverText, g.fontFace, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
