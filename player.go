package main

import (
	tl "github.com/JoelOtter/termloop"
)

type PlayerState int

const (
	stateAlive PlayerState = iota
	stateDead
)

type Player struct {
	lives  int
	score  int
	boardX int
	boardY int
	entity *tl.Entity
	game   *Game
	state  PlayerState
}

func NewPlayer(game *Game) *Player {
	player := Player{
		game:   game,
		entity: tl.NewEntity(1, 1, 1, 1),
	}
	player.Init()
	player.entity.SetCell(0, 0, &tl.Cell{Bg: tl.ColorRed, Ch: playerChar})
	return &player
}

func (player *Player) Draw(screen *tl.Screen) {
	player.entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		if player.state == stateDead {
			player.game.restartGame()
			return
		}
		switch event.Key {
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
			if (*player.game.board)[player.boardY][player.boardX].Hit() {
				player.score++
			} else {
				player.game.Kill()
			}
			player.game.updateStatus()
			if player.game.board.isLevelComplete() {
				player.boardX = 0
				player.boardY = 0
				player.entity.SetPosition(player.getPosition())
				player.game.nextLevel()
			}
			break
		}
		if player.game.isCaptured() {
			player.game.Kill()
		}
		player.entity.SetPosition(player.getPosition())
	}
}

func (player *Player) Init() {
	player.lives = playerLives
	player.score = 0
	player.boardX = 0
	player.boardY = 0
	player.state = stateAlive
	player.entity.SetPosition(player.getPosition())
}

func (player *Player) getPosition() (int, int) {
	x := player.boardX*(squareWidth+borderWidth) + offsetX + borderWidth
	y := player.boardY*(squareHeight+borderHeight) + offsetY + borderHeight
	return x, y
}
