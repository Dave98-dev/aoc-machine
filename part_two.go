package main

import (
	"aocmachine/helpers"
	"strings"
)

func solvePartTwo(part2 string) [4]int {

	part2Lines := strings.Split(part2, "\r\n")

	memory := [4]int{0, 0, 0, 0}

	for _, line := range part2Lines {
		if len(line) == 0 {
			continue
		}

		operation := convertOperationLine(line)
		memory = helpers.ExecuteInstructionWithAlias(memory, operation)
	}

	return memory
}
