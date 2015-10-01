package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

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

	rules = `Rules:
옷 - you
@ - math police

Every level has math
problems to solve. Your
task is to find the
equations with result
equal to the level
number.

Each wrong try costs
you one live. Each level
gives you new live, and
bonus points equal to
the remaining lives.
Avoid the math police.

Control:
▲ - move up
▼ - move down
◀ - move left
▶ - move right
Space - select solution
Esc - exit game`
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	NewGame().Run()
}
