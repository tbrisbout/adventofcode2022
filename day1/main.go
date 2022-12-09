package main

import (
	"fmt"
	"strings"

	"github.com/tbrisbout/adventofcode2022/pkg"
)

func countCalories(input string) int {
	elfBags := strings.Split(input, "\n\n")

	most := 0
	for _, bag := range elfBags {
		ingredients := strings.Split(bag, "\n")

		sum := pkg.Sum(pkg.Map(ingredients, pkg.MustAtoi))

		most = pkg.Max(sum, most)
	}

	return most
}

func main() {
	fmt.Println("Part 1:", countCalories(mainInput))
}
