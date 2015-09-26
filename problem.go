package main

import (
	"fmt"
	"math/rand"
)

// Problem contains number pair for solving
type Problem struct {
	n1    int
	n2    int
	valid bool
}

// ProblemList contains list of number pairs
type ProblemList []*Problem

// String returns textual representation of the number pair
func (pr *Problem) String() string {
	return fmt.Sprintf("%2d+%-2d", pr.n1, pr.n2)
}

// NewProblem generates new pair of numbers for the board
func NewProblem(level int, isValid bool) *Problem {
	sum := level
	n1 := rand.Intn(sum)
	n2 := sum - n1
	if !isValid {
		n2++
	}
	r := rand.Intn(10)
	if r > 4 {
		n1, n2 = n2, n1
	}
	return &Problem{n1, n2, isValid}
}

// NewProblemList generates all number pairs for a level
func NewProblemList(level, valid, size int) ProblemList {
	pl := make(ProblemList, size)
	isHit := true
	for i := range pl {
		if i >= valid {
			isHit = false
		}
		pl[i] = NewProblem(level, isHit)
	}
	for i := range pl {
		j := rand.Intn(i + 1)
		pl[i], pl[j] = pl[j], pl[i]
	}
	return pl
}
