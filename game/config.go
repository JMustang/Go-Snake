package game

import (
	"image/color"
	"time"
)

const (
	GameSpeed    = time.Second / 6
	ScreenWidth  = 640
	ScreenHeight = 480
	GridSize     = 20
	InitialSize  = 1
)

var (
	ColorSnake = color.White
	ColorFood  = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	ColorText  = color.White
)
