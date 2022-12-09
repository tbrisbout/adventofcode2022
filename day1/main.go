package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tbrisbout/adventofcode2022/pkg"
)

func countNTopHolders(input string, n int) int {
	chunks := strings.Split(input, "\n\n")

	l := []int{}
	for _, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		sum := pkg.Sum(pkg.Map(lines, pkg.MustAtoi))
		l = append(l, sum)
	}

	sort.Ints(l)

	return pkg.Sum(pkg.TakeRight(l, n))
}

func main() {
	fmt.Println("Part 1:", countNTopHolders(mainInput, 1))
	fmt.Println("Part 2:", countNTopHolders(mainInput, 3))
}
