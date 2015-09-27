package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Square struct {
	boardX  int
	boardY  int
	problem *Problem
	visible bool
	canvas  []tl.Cell
	entity  *tl.Entity
}

// Board contains all number pairs for the level
type Board [][]*Square

func NewSquare(boardX, boardY int, pr *Problem) *Square {
	str := []rune(pr.String())
	c := make([]tl.Cell, len(str))
	for i := range c {
		c[i] = tl.Cell{Ch: str[i]}
	}
	sq := Square{
		boardX:  boardX,
		boardY:  boardY,
		problem: pr,
		visible: true,
		canvas:  c,
		entity:  tl.NewEntity(1, 1, squareWidth, squareHeight),
	}
	sq.entity.SetPosition(sq.getPosition())
	return &sq
}

func (sq *Square) getPosition() (int, int) {
	x := offsetX + borderWidth + (sq.boardX * squareWidth) + sq.boardX*squareOffsetX + 1
	y := offsetY + borderHeight + (sq.boardY * squareHeight) + sq.boardY*squareOffsetY + 1
	return x, y
}

func (sq *Square) Tick(event tl.Event) {}

func (sq *Square) Draw(screen *tl.Screen) {
	if sq.visible {
		x, y := sq.getPosition()
		for i := 0; i < 5; i++ {
			screen.RenderCell(x+i, y, &sq.canvas[i])
		}
	}
}

func (sq *Square) Hit() bool {
	if sq.problem.valid && sq.visible {
		sq.visible = false
		return true
	}
	return false
}

// NewBoard generates new level board
func NewBoard() *Board {
	board := make(Board, boardHeight)
	for i := range board {
		board[i] = make([]*Square, boardWidth)
	}
	return &board
}

func (b *Board) isLevelComplete() bool {
	board := *b
	for y := range board {
		for x := range board {
			if board[x][y].problem.valid && board[x][y].visible {
				return false
			}
		}
	}
	return true
}

func (b *Board) populateBoard(gameLevel, answersPerLevel int, level tl.Level) {
	pl := NewProblemList(gameLevel, answersPerLevel, boardWidth*boardHeight)
	i := 0
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			rc := NewSquare(x, y, pl[i])
			(*b)[y][x] = rc
			level.AddEntity(rc)
			i++
		}
	}
}
