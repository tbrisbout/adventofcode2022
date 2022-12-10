package main

import (
	"fmt"
	"strings"

	"github.com/tbrisbout/adventofcode2022/pkg"
)

var (
	moves = []string{"Rock", "Paper", "Scissor"}

	moveMap   = map[string]string{"A": "Rock", "B": "Paper", "C": "Scissor", "X": "Rock", "Y": "Paper", "Z": "Scissor"}
	endingMap = map[string]string{"X": "loss", "Y": "draw", "Z": "win"}

	moveScore   = map[string]int{"Rock": 1, "Paper": 2, "Scissor": 3}
	endingScore = map[string]int{"loss": 0, "draw": 3, "win": 6}
)

func computeRound(round string) int {
	a, b := parseRound(round)
	a, b = moveMap[a], moveMap[b]

	res := "loss"
	if b == a {
		res = "draw"
	} else if pkg.IsPrev(moves, a, b) {
		res = "win"
	}

	return moveScore[b] + endingScore[res]
}

func computeScore(input string) int {
	rounds := strings.Split(strings.TrimSpace(input), "\n")

	return pkg.Sum(pkg.Map(rounds, computeRound))
}

func computeRoundWithEnding(round string) int {
	a, ending := parseRound(round)
	a, ending = moveMap[a], endingMap[ending]

	b := a // draw
	switch ending {
	case "loss":
		b = pkg.Prev(moves, a)
	case "win":
		b = pkg.Next(moves, a)
	}

	return moveScore[b] + endingScore[ending]
}

func computeScoreWithEnding(input string) int {
	rounds := strings.Split(strings.TrimSpace(input), "\n")

	return pkg.Sum(pkg.Map(rounds, computeRoundWithEnding))
}

func parseRound(round string) (a, b string) {
	m := strings.Split(round, " ")
	return m[0], m[1]
}

func main() {
	fmt.Println("Part1:", computeScore(mainInput))
	fmt.Println("Part2:", computeScoreWithEnding(mainInput))
}
