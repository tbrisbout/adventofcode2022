package main

import (
	"fmt"
	"strings"
	"unicode"
)

func findDuplicate(in string) rune {
	seen := make(map[rune]bool)
	l := len(in) / 2
	left, right := in[:l], in[l:]
	for _, r := range left {
		seen[r] = true
	}

	for _, r := range right {
		if seen[r] {
			return r
		}
	}

	return '0'
}

func sumDuplicates(input string) int {
	rucksacks := strings.Split(strings.TrimSpace(input), "\n")
	sum := 0
	for _, r := range rucksacks {
		code := findDuplicate(r)
		if unicode.IsLower(code) {
			sum += int(code) - 96
		} else {
			sum += int(code) - 64 + 26
		}
	}

	return sum
}

func main() {
	fmt.Println("Part1:", sumDuplicates(mainInput))
	// fmt.Println("Part2:", sumDuplicates(mainInput))
}
