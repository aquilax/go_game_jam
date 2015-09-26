package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

const (
	boardWidth      = 8
	boardHeight     = 8
	answersPerLevel = 10
	playerLives     = 3
	startLevel      = 1

	squareWidth   = 7
	borderWidth   = 1
	offsetX       = 1
	squareOffsetX = 1

	squareHeight  = 3
	borderHeight  = 1
	offsetY       = 1
	squareOffsetY = 1

	playerChar = '옷'
)

func buildLevel(game *tl.Game, gameLevel, score int) {
	screen := game.Screen()
	// Add title
	screen.AddEntity(tl.NewText(1, 0, " Number crusader! ", tl.ColorBlack, tl.ColorGreen))

	// Add the numbers
	level := tl.NewBaseLevel(tl.Cell{})
	// TODO: Remove this abomination
	level.AddEntity(tl.NewRectangle(1, 1, 65, 33, tl.ColorGreen))
	for i := 2; i < 63; i = i + 8 {
		for j := 2; j < 31; j = j + 4 {
			level.AddEntity(tl.NewRectangle(i, j, 7, 3, tl.ColorBlue))
		}
	}
	board := NewBoard(gameLevel)
	bcl := NewProblemList(gameLevel, answersPerLevel, boardWidth*boardHeight)
	for y := 0; y < boardWidth; y++ {
		for x := 0; x < boardHeight; x++ {
			rc := NewSquare(x, y, bcl[x*boardWidth+y])
			(*board)[x][y] = rc
			level.AddEntity(rc)
		}
	}
	player := NewPlayer(game, board)
	level.AddEntity(player)
	level.AddEntity(player.status)
	screen.SetLevel(level)
}

func main() {
	game := tl.NewGame()
	game.SetDebugOn(true)
	rand.Seed(time.Now().UTC().UnixNano())
	buildLevel(game, 1, 0)
	game.Start()
}
