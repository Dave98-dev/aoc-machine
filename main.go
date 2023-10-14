package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func UNUSED(x ...interface{}) {}

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	inputString := string(dat)
	res := strings.Split(inputString, "---")
	part1, part2 := res[0], res[1]

	fmt.Println(solvePartOne(part1))
	fmt.Println(solvePartTwo(part2))
}
