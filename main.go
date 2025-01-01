package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/exp/rand"
)

// Game constants
const (
	gameSpeed    = time.Second / 6
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 20
	initialSize  = 1
)

// Colors
var (
	colorSnake = color.White
	colorFood  = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	colorText  = color.White
)

// Direction represents a 2D vector for movement
type Direction struct {
	x, y int
}

var (
	dirUp    = Direction{x: 0, y: -1}
	dirDown  = Direction{x: 0, y: 1}
	dirLeft  = Direction{x: -1, y: 0}
	dirRight = Direction{x: 1, y: 0}
)

// Point represents a position on the game grid
type Point struct {
	x, y int
}

// Game holds the game state
type Game struct {
	snake      []Point
	direction  Direction
	lastUpdate time.Time
	food       Point
	gameOver   bool
	score      int
	fontFace   *text.GoTextFace
}

// NewGame creates and initializes a new game instance
func NewGame(fontSource *text.GoTextFaceSource) *Game {
	g := &Game{
		snake: make([]Point, initialSize),
		fontFace: &text.GoTextFace{
			Source: fontSource,
			Size:   30,
		},
	}

	// Initialize snake at the center of the screen
	g.snake[0] = Point{
		x: screenWidth / gridSize / 2,
		y: screenHeight / gridSize / 2,
	}
	g.direction = dirRight
	g.spawnFood()

	return g
}

// Update handles game logic updates
func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.reset()
		}
		return nil
	}

	g.handleInput()

	if time.Since(g.lastUpdate) < gameSpeed {
		return nil
	}
	g.lastUpdate = time.Now()

	g.updateSnake()
	return nil
}

// handleInput processes keyboard input
func (g *Game) handleInput() {
	newDir := g.direction
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.direction != dirDown {
		newDir = dirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.direction != dirUp {
		newDir = dirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.direction != dirLeft {
		newDir = dirRight
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.direction != dirRight {
		newDir = dirLeft
	}
	g.direction = newDir
}

// updateSnake moves the snake and handles collisions
func (g *Game) updateSnake() {
	head := g.snake[0]
	newHead := Point{
		x: head.x + g.direction.x,
		y: head.y + g.direction.y,
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

// isCollision checks if a point collides with walls or snake body
func (g *Game) isCollision(p Point) bool {
	// Wall collision
	if p.x < 0 || p.y < 0 || p.x >= screenWidth/gridSize || p.y >= screenHeight/gridSize {
		return true
	}

	// Self collision
	for _, sp := range g.snake[1:] {
		if sp == p {
			return true
		}
	}

	return false
}

// spawnFood places food at a random empty position
func (g *Game) spawnFood() {
	for {
		g.food = Point{
			x: rand.Intn(screenWidth / gridSize),
			y: rand.Intn(screenHeight / gridSize),
		}
		// Make sure food doesn't spawn on snake
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

// reset resets the game state for a new game
func (g *Game) reset() {
	g.snake = make([]Point, initialSize)
	g.snake[0] = Point{
		x: screenWidth / gridSize / 2,
		y: screenHeight / gridSize / 2,
	}
	g.direction = dirRight
	g.gameOver = false
	g.score = 0
	g.spawnFood()
}

// Draw renders the game state
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw snake
	for _, p := range g.snake {
		vector.DrawFilledRect(
			screen,
			float32(p.x*gridSize),
			float32(p.y*gridSize),
			gridSize,
			gridSize,
			colorSnake,
			true,
		)
	}

	// Draw food
	vector.DrawFilledRect(
		screen,
		float32(g.food.x*gridSize),
		float32(g.food.y*gridSize),
		gridSize,
		gridSize,
		colorFood,
		true,
	)

	// Draw score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(colorText)
	text.Draw(screen, scoreText, g.fontFace, op)

	// Draw game over screen
	if g.gameOver {
		gameOverText := "Game Over! Press SPACE to restart"
		w, h := text.Measure(gameOverText, g.fontFace, g.fontFace.Size)
		op := &text.DrawOptions{}
		op.GeoM.Translate(screenWidth/2-w/2, screenHeight/2-h/2)
		op.ColorScale.ScaleWithColor(colorText)
		text.Draw(screen, gameOverText, g.fontFace, op)
	}
}

// Layout implements ebiten.Game interface
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(NewGame(fontSource)); err != nil {
		log.Fatal(err)
	}
}
