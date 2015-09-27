package main

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

type Game struct {
	level  int
	game   *tl.Game
	board  *Board
	player *Player
	foes   []*Foe
	status *tl.Text
}

func NewGame() *Game {
	game := &Game{
		level:  startLevel,
		game:   tl.NewGame(),
		board:  NewBoard(),
		status: tl.NewText(19, 0, "", tl.ColorWhite, tl.ColorBlack),
	}
	game.player = NewPlayer(game)
	game.updateStatus()
	return game
}

func (g *Game) Run() {
	rand.Seed(time.Now().UTC().UnixNano())
	g.addChrome()
	g.buildLevel(1)
	g.game.Start()
}

func (g *Game) buildLevel(gameLevel int) {
	level := tl.NewBaseLevel(tl.Cell{})
	// TODO: Remove this abomination
	width := boardWidth*squareWidth + (boardWidth+1)*borderWidth
	height := boardHeight*squareHeight + (boardHeight+1)*borderHeight
	level.AddEntity(tl.NewRectangle(1, 1, width, height, tl.ColorGreen))
	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			x := offsetX + borderWidth + (j * squareWidth) + j*borderWidth
			y := offsetY + borderHeight + (i * squareHeight) + i*borderHeight
			level.AddEntity(tl.NewRectangle(x, y, squareWidth, squareHeight, tl.ColorBlue))
		}
	}
	g.board.populateBoard(gameLevel, answersPerLevel, level)
	level.AddEntity(g.player)
	// Add Foes
	foes := int(gameLevel/10) + 2
	g.foes = g.foes[:0]
	var foe *Foe
	for i := 0; i < foes; i++ {
		foe = NewFoe(g)
		g.foes = append(g.foes, foe)
		level.AddEntity(foe)
	}
	g.game.Screen().SetLevel(level)
	g.updateStatus()
}

func (g *Game) addChrome() {
	screen := g.game.Screen()
	screen.AddEntity(tl.NewText(offsetX, 0, " Number crusher! ", tl.ColorBlack, tl.ColorGreen))
	x := 2*offsetX + (boardWidth * squareWidth) + (boardWidth * borderWidth) + borderWidth
	rules := tl.NewEntityFromCanvas(x, offsetY, tl.CanvasFromString(rules))
	screen.AddEntity(rules)
	screen.AddEntity(g.status)
}

func (g *Game) updateStatus() {
	statusText := fmt.Sprintf("Lvl %2d | ❤ %d | Score %06d", g.level, g.player.lives, g.player.score)
	g.status.SetText(statusText)
}

func (g *Game) nextLevel() {
	g.level++
	g.buildLevel(g.level)
}

func (g *Game) restartGame() {
	g.level = 1
	g.player.Init()
	g.buildLevel(g.level)
}

func (g *Game) gameOver() {
	g.game.Screen().Level().AddEntity(tl.NewText(28, 17, " GAME OVER ", tl.ColorBlack, tl.ColorRed))
}

func (g *Game) isCaptured() bool {
	px := g.player.boardX
	py := g.player.boardY
	for i := range g.foes {
		if px == g.foes[i].boardX && py == g.foes[i].boardY {
			return true
		}
	}
	return false
}

func (g *Game) Kill() {
	g.player.lives--
	g.updateStatus()
	g.player.boardX = 0
	g.player.boardY = 0
	for i := range g.foes {
		g.foes[i].boardX = boardWidth - 1
		g.foes[i].boardY = boardHeight - 1
		g.foes[i].entity.SetPosition(g.foes[i].getPosition())
	}
}
