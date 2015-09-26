package main

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
)

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
			if player.board.isLevelComplete() {
				// TODO: New Level
			}
			break
		}
		player.game.Log("BoardX=%d\tBoardY=%d", player.boardX, player.boardY)
		player.entity.SetPosition(player.getPosition())
	}
}

func (player *Player) updateStatus() {
	statusText := fmt.Sprintf("Lvl %2d | ❤ %d | Score %06d", player.level, player.lives, player.score)
	player.status.SetText(statusText)
}

func (player *Player) getPosition() (int, int) {
	x := player.boardX*(squareWidth+borderWidth) + offsetX + borderWidth
	y := player.boardY*(squareHeight+borderHeight) + offsetY + borderHeight
	return x, y
}
