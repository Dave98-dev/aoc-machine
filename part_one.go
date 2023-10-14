package main

import (
	"strconv"
	"strings"
)

func convertMemoryLine(line string) [4]int {

	one := int(line[9]) - '0'
	two := int(line[12]) - '0'
	three := int(line[15]) - '0'
	four := int(line[18]) - '0'

	res := [4]int{one, two, three, four}
	return res
}
func convertOperationLine(line string) [4]int {

	vars := strings.Split(line, " ")

	one, _ := strconv.Atoi(vars[0])
	two, _ := strconv.Atoi(vars[1])
	three, _ := strconv.Atoi(vars[2])
	four, _ := strconv.Atoi(vars[3])

	res := [4]int{one, two, three, four}
	return res
}

func solvePartOne(part1 string) map[int][]int {

	input := make([]machineInstructionExample, 0)

	part1Lines := strings.Split(part1, "\r\n")

	for i := 0; i < len(part1Lines); i += 4 {
		if len(part1Lines[i]) < 19 {
			break
		}

		inputRow := machineInstructionExample{
			before:    convertMemoryLine(part1Lines[i]),
			operation: convertOperationLine(part1Lines[i+1]),
			after:     convertMemoryLine(part1Lines[i+2]),
		}

		input = append(input, inputRow)
	}

	m := mapOpCodes(input)

	println(countOpCodes(input))

	return m
}
