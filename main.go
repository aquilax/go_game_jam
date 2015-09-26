package main

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"strconv"
)

const (
	boardWidth      = 8
	boardHeight     = 8
	answersPerLevel = 10
)

type BoardCell struct {
	n1 int
	n2 int
}

type BoardCellList []BoardCell

type Board [][]BoardCell

func (bc *BoardCell) String() string {
	return fmt.Sprintf("%2d+%-2d", bc.n1, bc.n2)
}

// NewBoardCell generates new pair of numbers for the board
func NewBoardCell(level int, isHit bool) BoardCell {
	sum := level + 1
	n1 := rand.Intn(sum)
	n2 := sum - n1
	if !isHit {
		n2++
	}
	return BoardCell{n1, n2}
}

// NewBoardCellList generates all number pairs for a level
func NewBoardCellList(level, valid, size int) BoardCellList {
	bcl := make(BoardCellList, size)
	isHit := true
	for i := range bcl {
		if i > valid {
			isHit = false
		}
		bcl[i] = NewBoardCell(level, isHit)
	}
	for i := range bcl {
		j := rand.Intn(i + 1)
		bcl[i], bcl[j] = bcl[j], bcl[i]
	}
	return bcl
}

// NewBoard generates new level board
func NewBoard(level int) Board {
	board := make(Board, boardWidth)
	bcl := NewBoardCellList(level, answersPerLevel, boardWidth*boardHeight)
	for i := range board {
		board[i] = make([]BoardCell, boardHeight)
		for j := range board {
			board[i][j] = bcl[i+j]
		}
	}
	return board
}

func buildLevel(game *tl.Game, gameLevel, score int) {
	level := tl.NewBaseLevel(tl.Cell{})
	game.Screen().SetLevel(level)
	scoretext := tl.NewText(50, 0, "Score: "+strconv.Itoa(score),
		tl.ColorWhite, tl.ColorBlack)
	game.Screen().AddEntity(tl.NewText(1, 0, " Number crusader! ", tl.ColorBlack, tl.ColorGreen))
	game.Screen().AddEntity(scoretext)

	level.AddEntity(tl.NewRectangle(1, 1, 65, 33, tl.ColorGreen))
	for i := 2; i < 63; i = i + 8 {
		for j := 2; j < 31; j = j + 4 {
			level.AddEntity(tl.NewRectangle(i, j, 7, 3, tl.ColorBlue))
		}
	}
	board := NewBoard(gameLevel)
	for i := range board {
		for j := range board[i] {
			x := i*8 + 3
			y := j*4 + 3
			game.Screen().AddEntity(tl.NewText(x, y, board[i][j].String(), tl.ColorWhite, tl.ColorBlue))
		}
	}
}

func main() {
	game := tl.NewGame()
	game.SetDebugOn(true)
	buildLevel(game, 1, 0)
	game.Start()
}
