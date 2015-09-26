package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	level  int
	lives  int
	score  int
	boardX int
	boardY int
	entity *tl.Entity
	game   *Game
}

func NewPlayer() *Player {
	player := Player{
		lives:  playerLives,
		score:  0,
		boardX: 0,
		boardY: 0,
		entity: tl.NewEntity(1, 1, 1, 1),
	}
	player.entity.SetCell(0, 0, &tl.Cell{Bg: tl.ColorRed, Ch: playerChar})
	player.entity.SetPosition(player.getPosition())
	return &player
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
			if (*player.game.board)[player.boardX][player.boardY].Hit() {
				player.score++
			} else {
				player.lives--
			}
			player.game.updateStatus()
			if player.game.board.isLevelComplete() {
				player.game.nextLevel()
			}
			break
		}
		player.entity.SetPosition(player.getPosition())
	}
}

func (player *Player) getPosition() (int, int) {
	x := player.boardX*(squareWidth+borderWidth) + offsetX + borderWidth
	y := player.boardY*(squareHeight+borderHeight) + offsetY + borderHeight
	return x, y
}

func (player *Player) setGame(game *Game) {
	player.game = game
}
