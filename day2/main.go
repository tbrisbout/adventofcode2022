package main

import (
	"fmt"
	"strings"

	"github.com/tbrisbout/adventofcode2022/pkg"
)

var scores = map[string]int{"X": 1, "Y": 2, "Z": 3, "win": 6, "draw": 3, "loss": 0}
var mapping = map[string]string{"A": "X", "B": "Y", "C": "Z"}

func computeRound(round string) int {
	moves := strings.Split(round, " ")
	a, b := moves[0], moves[1]

	res := "loss"
	if b == mapping[a] {
		res = "draw"
	} else if a == "A" && b == "Y" || a == "B" && b == "Z" || a == "C" && b == "X" {
		res = "win"
	}

	return scores[b] + scores[res]
}

func computeScore(input string) int {
	rounds := strings.Split(strings.TrimSpace(input), "\n")

	return pkg.Sum(pkg.Map(rounds, computeRound))
}

func main() {
	fmt.Println("Part1:", computeScore(mainInput))
}
