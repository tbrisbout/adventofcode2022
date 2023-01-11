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

func main() {
	fmt.Println("Part1:", countContainedRanges(mainInput))
}
