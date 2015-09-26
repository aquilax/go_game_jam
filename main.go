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

	cellWidth   = 7
	borderWidth = 1
	offsetX     = 1

	cellHeight   = 3
	borderHeight = 1
	offsetY      = 1
)

// BoardCell contains number pair for solving
type BoardCell struct {
	n1 int
	n2 int
}

// BoardCellList contains list of number pairs
type BoardCellList []BoardCell

// Board contains all number pairs for the level
type Board [][]BoardCell

type Player struct {
	game   *tl.Game
	boardX int
	boardY int
	entity *tl.Entity
}

// String returns textual representation of the number pair
func (bc *BoardCell) String() string {
	return fmt.Sprintf("%2d+%-2d", bc.n1, bc.n2)
}

func NewPlayer(game *tl.Game) Player {
	player := Player{
		game,
		0,
		0,
		tl.NewEntity(1, 1, 1, 1),
	}
	player.entity.SetCell(0, 0, &tl.Cell{Bg: tl.ColorRed, Ch: '옷'})
	player.entity.SetPosition(player.getPosition())
	return player
}

func (player *Player) Draw(screen *tl.Screen) {
	player.entity.Draw(screen)
}

func (player *Player) getPosition() (int, int) {
	x := player.boardX*(cellWidth+borderWidth) + offsetX + borderWidth
	y := player.boardY*(cellHeight+borderHeight) + offsetY + borderHeight
	return x, y
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			if player.boardX < boardWidth-1 {
				player.boardX += 1
			}
			break
		case tl.KeyArrowLeft:
			if player.boardX > 0 {
				player.boardX -= 1
			}
			break
		case tl.KeyArrowUp:
			if player.boardY > 0 {
				player.boardY -= 1
			}
			break
		case tl.KeyArrowDown:
			if player.boardY < boardHeight-1 {
				player.boardY += 1
			}
			break
		}
		player.game.Log("BoardX=%d\tBoardY=%d", player.boardX, player.boardY)
		player.entity.SetPosition(player.getPosition())
	}
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
	// Add chrome
	scoretext := tl.NewText(50, 0, "Score: "+strconv.Itoa(score),
		tl.ColorWhite, tl.ColorBlack)
	game.Screen().AddEntity(tl.NewText(1, 0, " Number crusader! ", tl.ColorBlack, tl.ColorGreen))
	game.Screen().AddEntity(scoretext)

	// Create the level background
	level.AddEntity(tl.NewRectangle(1, 1, 65, 33, tl.ColorGreen))
	for i := 2; i < 63; i = i + 8 {
		for j := 2; j < 31; j = j + 4 {
			level.AddEntity(tl.NewRectangle(i, j, 7, 3, tl.ColorBlue))
		}
	}
	// Add the numbers
	board := NewBoard(gameLevel)
	for i := range board {
		for j := range board[i] {
			x := i*8 + 3
			y := j*4 + 3
			game.Screen().AddEntity(tl.NewText(x, y, board[i][j].String(), tl.ColorWhite, tl.ColorBlue))
		}
	}
	player := NewPlayer(game)
	level.AddEntity(&player)

}

func main() {
	game := tl.NewGame()
	game.SetDebugOn(true)
	buildLevel(game, 1, 0)
	game.Start()
}
