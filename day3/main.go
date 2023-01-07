package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/tbrisbout/adventofcode2022/pkg"
)

func sumDuplicates(groups [][]string) (sum int) {
	for _, group := range groups {
		code, _ := findDuplicate(group)
		sum += charCodeToPriority(code)
	}
	return sum
}

func findDuplicate(lines []string) (rune, bool) {
	if len(lines[0]) == 0 {
		return 0, false
	}

	for _, r := range lines[0] {
		found := true
		for _, line := range lines[1:] {
			if !strings.ContainsRune(line, r) {
				found = false
				break
			}
		}
		if found {
			return r, true
		}
	}

	return 0, false
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

func splitByHalf(input string) (ret [][]string) {
	for _, line := range strings.Fields(input) {
		ret = append(ret, half(line))
	}

	return ret
}

func half(in string) []string {
	l := len(in) / 2
	return []string{in[:l], in[l:]}
}

func splitByChunk(input string) [][]string {
	return pkg.ChunkBy(strings.Fields(input), 3)
}

func main() {
	fmt.Println("Part1:", sumDuplicates(splitByHalf(mainInput)))
	fmt.Println("Part2:", sumDuplicates(splitByChunk(mainInput)))
}
