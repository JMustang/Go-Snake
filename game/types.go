package game

type Direction struct {
	X, Y int
}

type Point struct {
	X, Y int
}

var (
	DirUp    = Direction{X: 0, Y: -1}
	DirDown  = Direction{X: 0, Y: 1}
	DirLeft  = Direction{X: -1, Y: 0}
	DirRight = Direction{X: 1, Y: 0}
)
