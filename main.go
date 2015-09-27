package main

const (
	boardWidth      = 8
	boardHeight     = 8
	answersPerLevel = 10
	playerLives     = 3
	startLevel      = 1

	squareWidth   = 7
	borderWidth   = 1
	offsetX       = 1
	squareOffsetX = 1

	squareHeight  = 3
	borderHeight  = 1
	offsetY       = 1
	squareOffsetY = 1

	playerChar = '옷'
	foeChar    = '@'
)

func main() {
	NewGame().Run()
}
