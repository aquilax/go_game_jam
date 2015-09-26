package main

import (
	"fmt"
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

	cellWidth   = 7
	borderWidth = 1
	offsetX     = 1
	cellOffsetX = 1

	cellHeight   = 3
	borderHeight = 1
	offsetY      = 1
	cellOffsetY  = 1

	playerChar = '옷'
)

// BoardCell contains number pair for solving
type BoardCell struct {
	n1    int
	n2    int
	valid bool
}

// BoardCellList contains list of number pairs
type BoardCellList []BoardCell

// Board contains all number pairs for the level

type Player struct {
	level  int
	lives  int
	score  int
	game   *tl.Game
	boardX int
	boardY int
	board  *Board
	entity *tl.Entity
	status *tl.Text
}

type RenderCell struct {
	boardX  int
	boardY  int
	bc      BoardCell
	visible bool
	canvas  []tl.Cell
	entity  *tl.Entity
}

type Board [][]*RenderCell

// String returns textual representation of the number pair
func (bc *BoardCell) String() string {
	return fmt.Sprintf("%2d+%-2d", bc.n1, bc.n2)
}

func NewRenderCell(boardX, boardY int, bc BoardCell) *RenderCell {
	str := []rune(bc.String())
	c := make([]tl.Cell, len(str))
	for i := range c {
		c[i] = tl.Cell{Ch: str[i]}
	}
	rc := RenderCell{
		boardX,
		boardY,
		bc,
		true,
		c,
		tl.NewEntity(1, 1, cellWidth, cellHeight),
	}
	rc.entity.SetPosition(rc.getPosition())
	return &rc
}

func (rc *RenderCell) getPosition() (int, int) {
	x := offsetX + borderWidth + (rc.boardX * cellWidth) + rc.boardX*cellOffsetX + 1
	y := offsetY + borderHeight + (rc.boardY * cellHeight) + rc.boardY*cellOffsetY + 1
	return x, y
}

func (rc *RenderCell) Tick(event tl.Event) {}

func (rc *RenderCell) Draw(screen *tl.Screen) {
	if rc.visible {
		x, y := rc.getPosition()
		for i := 0; i < 5; i++ {
			screen.RenderCell(x+i, y, &rc.canvas[i])
		}
	}
}

func (rc *RenderCell) Hit() bool {
	if rc.bc.valid {
		rc.visible = false
		return true
	}
	return false
}

func NewPlayer(game *tl.Game, board *Board) Player {
	player := Player{
		startLevel,
		playerLives,
		0,
		game,
		0,
		0,
		board,
		tl.NewEntity(1, 1, 1, 1),
		tl.NewText(20, 0, "", tl.ColorWhite, tl.ColorBlack),
	}
	player.updateStatus()
	player.entity.SetCell(0, 0, &tl.Cell{Bg: tl.ColorRed, Ch: playerChar})
	player.entity.SetPosition(player.getPosition())
	return player
}

func (player *Player) Draw(screen *tl.Screen) {
	player.entity.Draw(screen)
}

func (player *Player) updateStatus() {
	statusText := fmt.Sprintf("Lvl %2d | ❤ %d | Score %06d", player.level, player.lives, player.score)
	player.status.SetText(statusText)
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
		case tl.KeySpace:
			if (*player.board)[player.boardX][player.boardY].Hit() {
				player.score++
			} else {
				player.lives--
			}
			player.updateStatus()
			break
		}

		player.game.Log("BoardX=%d\tBoardY=%d", player.boardX, player.boardY)
		player.entity.SetPosition(player.getPosition())
	}
}

// NewBoardCell generates new pair of numbers for the board
func NewBoardCell(level int, isValid bool) BoardCell {
	sum := level
	n1 := rand.Intn(sum)
	n2 := sum - n1
	if !isValid {
		n2++
	}
	return BoardCell{n1, n2, isValid}
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
	for i := range board {
		board[i] = make([]*RenderCell, boardHeight)
	}
	return board
}

func buildLevel(game *tl.Game, gameLevel, score int) {
	level := tl.NewBaseLevel(tl.Cell{})
	game.Screen().SetLevel(level)
	// Add title
	game.Screen().AddEntity(tl.NewText(1, 0, " Number crusader! ", tl.ColorBlack, tl.ColorGreen))

	// Create the level background
	level.AddEntity(tl.NewRectangle(1, 1, 65, 33, tl.ColorGreen))
	for i := 2; i < 63; i = i + 8 {
		for j := 2; j < 31; j = j + 4 {
			level.AddEntity(tl.NewRectangle(i, j, 7, 3, tl.ColorBlue))
		}
	}
	// Add the numbers
	board := NewBoard(gameLevel)
	bcl := NewBoardCellList(gameLevel, answersPerLevel, boardWidth*boardHeight)
	for x := 0; x < boardWidth; x++ {
		for y := 0; y < boardHeight; y++ {
			rc := NewRenderCell(x, y, bcl[x+y])
			board[x][y] = rc
			game.Screen().AddEntity(rc)
		}
	}
	player := NewPlayer(game, &board)
	level.AddEntity(&player)
	level.AddEntity(player.status)
}

func main() {
	game := tl.NewGame()
	game.SetDebugOn(true)
	rand.Seed(time.Now().UTC().UnixNano())
	buildLevel(game, 1, 0)
	game.Start()
}
