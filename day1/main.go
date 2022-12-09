package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countCalories(input string) int {
	elfBags := strings.Split(input, "\n\n")

	most := 0
	for _, bag := range elfBags {
		ingredients := strings.Split(bag, "\n")
		sum := 0

		for _, ingredient := range ingredients {
			calories, _ := strconv.Atoi(ingredient)
			sum += calories
		}

		if sum > most {
			most = sum
		}
	}

	return most
}

func main() {
	fmt.Println("Part 1:", countCalories(mainInput))
}
