package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/tbrisbout/adventofcode2022/pkg"
)

func findDuplicate(in string) rune {
	l := len(in) / 2
	left, right := in[:l], in[l:]

	for _, r := range right {
		if strings.ContainsRune(left, r) {
			return r
		}
	}

	return '0'
}

func charCodeToPriority(r rune) (p int) {
	if unicode.IsLower(r) {
		// 96 is charcode of 'a' - 1
		p = int(r) - 96
	} else {
		// 64 is charcode of 'A' - 1
		// 26 is the alphabet len
		p = int(r) - 64 + 26
	}
	return p
}

func sumDuplicates(input string) (sum int) {
	for _, r := range strings.Fields(input) {
		code := findDuplicate(r)
		sum += charCodeToPriority(code)
	}

	return sum
}

func sumDuplicatesInGroups(input string) (sum int) {
	groups := pkg.ChunkBy(strings.Fields(input), 3)

	for _, g := range groups {
		for _, r := range g[2] {
			if strings.ContainsRune(g[0], r) && strings.ContainsRune(g[1], r) {
				sum += charCodeToPriority(r)
				break
			}
		}
	}

	return sum
}

func main() {
	fmt.Println("Part1:", sumDuplicates(mainInput))
	fmt.Println("Part2:", sumDuplicatesInGroups(mainInput))
}
