package main

import "fmt"

func mapOpCodes(example []machineInstructionExample) map[int]([]int) {
	m := make(map[int]([]int))

	for _, exampleRow := range example {
		originalOpCode := exampleRow.operation[0]
		val := m[originalOpCode]

		possibilities := make([]int, 0)

		for i := 0; i <= 15; i++ {
			result := executeInstruction(exampleRow.before, [4]int{i, exampleRow.operation[1], exampleRow.operation[2], exampleRow.operation[3]})

			if result == exampleRow.after {
				possibilities = append(possibilities, i)
			}
		}

		unionRes := union(possibilities, val)

		if len(unionRes) == 0 {
			fmt.Println(exampleRow.before)
			fmt.Println(exampleRow.operation)
			fmt.Println(exampleRow.after)
			fmt.Println(possibilities)
			fmt.Println(val)
			fmt.Println(m)
			panic("valore insulso")
		}

		m[originalOpCode] = unionRes
	}

	return m
}

func countOpCodes(example []machineInstructionExample) int {
	finalCount := 0

	for _, exampleRow := range example {
		possibilities := 0

		for i := 0; i <= 15; i++ {
			result := executeInstruction(exampleRow.before, [4]int{i, exampleRow.operation[1], exampleRow.operation[2], exampleRow.operation[3]})

			if result == exampleRow.after {
				possibilities++
			}
		}

		if possibilities >= 3 {
			finalCount++
		}
	}

	return finalCount
}

func union(one []int, two []int) []int {

	if len(two) == 0 {
		return one
	}

	res := make([]int, 0)
	for _, i := range one {
		for _, j := range two {
			if i == j {
				res = append(res, i)
			}
		}
	}
	return res
}
