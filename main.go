package main

import (
	tl "github.com/JoelOtter/termloop"
	"strconv"
)

type Board struct {
	level  int
	entity *tl.Entity
}

func (board *Board) Draw(screen *tl.Screen) {
	board.entity.Draw(screen)
}

func (board *Board) Tick(event tl.Event) {
}

func NewBoard(level int) *Board {
	board := &Board{
		level,
		tl.NewEntity(1, 1, 63, 63),
	}
	return board
}

func buildLevel(game *tl.Game, score int) {
	level := tl.NewBaseLevel(tl.Cell{})
	game.Screen().SetLevel(level)
	scoretext := tl.NewText(50, 0, "Score: "+strconv.Itoa(score),
		tl.ColorWhite, tl.ColorBlack)
	game.Screen().AddEntity(tl.NewText(1, 0, "Number crusader!", tl.ColorRed, tl.ColorBlack))
	game.Screen().AddEntity(scoretext)

	level.AddEntity(tl.NewRectangle(1, 1, 65, 33, tl.ColorGreen))
	for i := 2; i < 63; i = i + 8 {
		for j := 2; j < 31; j = j + 4 {
			level.AddEntity(tl.NewRectangle(i, j, 7, 3, tl.ColorBlue))
		}
	}
}

func main() {
	game := tl.NewGame()
	game.SetDebugOn(true)
	buildLevel(game, 0)
	game.Start()
}
