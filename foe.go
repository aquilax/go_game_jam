package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)

type Foe struct {
	boardX int
	boardY int
	game   *Game
	entity *tl.Entity
	speed  int
	frame  int
}

func NewFoe(game *Game) *Foe {
	foe := Foe{
		game:   game,
		entity: tl.NewEntity(1, 1, 1, 1),
	}
	foe.Init()
	foe.entity.SetCell(0, 0, &tl.Cell{Bg: tl.ColorMagenta, Ch: foeChar})
	return &foe
}

func (foe *Foe) Draw(screen *tl.Screen) {
	foe.speed = int(foe.game.game.Screen().TimeDelta() * float64(400000/foe.game.level))
	if foe.frame > foe.speed {
		foe.boardX, foe.boardY = foe.newPosition(foe.game.player.boardX, foe.game.player.boardY, foe.boardX, foe.boardY)
		foe.entity.SetPosition(foe.getPosition())
		if foe.game.isCaptured() {
			foe.game.Kill()
		}
		foe.frame = 0
	}
	foe.frame++

	foe.entity.Draw(screen)
}

func (foe *Foe) Tick(event tl.Event) {
}

func (foe *Foe) newPosition(playerX, playerY, x, y int) (int, int) {
	move := rand.Intn(2) - 1
	if move != 0 {
		if rand.Intn(100) > 50 {
			newX := x + move
			if newX >= 0 && newX < boardWidth {
				return newX, y
			}
			return x - move, y
		} else {
			newY := y + move
			if newY >= 0 && newY < boardHeight {
				return x, newY
			}
			return x, y - move

		}
	}
	return x, y
}

func (foe *Foe) getPosition() (int, int) {
	x := foe.boardX*(squareWidth+borderWidth) + offsetX + borderWidth + squareWidth - 1
	y := foe.boardY*(squareHeight+borderHeight) + offsetY + borderHeight + squareHeight - 1
	return x, y
}

func (foe *Foe) Init() {
	foe.boardX = boardHeight - 1
	foe.boardY = boardWidth - 1
	foe.frame = 0
	foe.entity.SetPosition(foe.getPosition())
}
