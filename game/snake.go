package game

import "golang.org/x/exp/rand"

func (g *Game) updateSnake() {
	head := g.snake[0]
	newHead := Point{
		X: head.X + g.direction.X,
		Y: head.Y + g.direction.Y,
	}

	if g.isCollision(newHead) {
		g.gameOver = true
		return
	}

	if newHead == g.food {
		g.snake = append([]Point{newHead}, g.snake...)
		g.score++
		g.spawnFood()
	} else {
		g.snake = append([]Point{newHead}, g.snake[:len(g.snake)-1]...)
	}
}

func (g *Game) isCollision(p Point) bool {
	if p.X < 0 || p.Y < 0 || p.X >= ScreenWidth/GridSize || p.Y >= ScreenHeight/GridSize {
		return true
	}

	for _, sp := range g.snake[1:] {
		if sp == p {
			return true
		}
	}

	return false
}

func (g *Game) spawnFood() {
	for {
		g.food = Point{
			X: rand.Intn(ScreenWidth / GridSize),
			Y: rand.Intn(ScreenHeight / GridSize),
		}
		collision := false
		for _, p := range g.snake {
			if p == g.food {
				collision = true
				break
			}
		}
		if !collision {
			break
		}
	}
}

func (g *Game) reset() {
	g.snake = make([]Point, InitialSize)
	g.snake[0] = Point{
		X: ScreenWidth / GridSize / 2,
		Y: ScreenHeight / GridSize / 2,
	}
	g.direction = DirRight
	g.gameOver = false
	g.score = 0
	g.spawnFood()
}
