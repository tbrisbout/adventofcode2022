package main

import (
	"fmt"
	"regexp"

	"github.com/tbrisbout/adventofcode2022/pkg"
)

var rangePattern = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func countContainedRanges(input string) (sum int) {
	for _, ranges := range rangePattern.FindAllStringSubmatch(input, -1) {
		n := pkg.Map(ranges[1:], pkg.MustAtoi)

		if n[0] <= n[2] && n[1] >= n[3] || n[0] >= n[2] && n[1] <= n[3] {
			sum++
		}
	}

	return sum
}

func between(n, lower, upper int) bool {
	return lower <= n && n <= upper
}

func countOverlappingRanges(input string) (sum int) {
	for _, ranges := range rangePattern.FindAllStringSubmatch(input, -1) {
		n := pkg.Map(ranges[1:], pkg.MustAtoi)

		a1, a2, b1, b2 := n[0], n[1], n[2], n[3]
		if between(b1, a1, a2) || between(b2, a1, a2) || between(a1, b1, b2) || between(a2, b1, b2) {
			sum++
		}
	}

	return sum
}

func main() {
	fmt.Println("Part1:", countContainedRanges(mainInput))
	fmt.Println("Part1:", countOverlappingRanges(mainInput))
}
