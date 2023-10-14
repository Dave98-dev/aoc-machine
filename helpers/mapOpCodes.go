package helpers

import (
	model "aocmachine/model"
	"fmt"
)

func MapOpCodes(example []model.MachineInstructionExample) map[int]([]int) {
	m := make(map[int]([]int))

	for _, exampleRow := range example {
		originalOpCode := exampleRow.Operation[0]
		val := m[originalOpCode]

		possibilities := make([]int, 0)

		for i := 0; i <= 15; i++ {
			result := executeInstruction(exampleRow.Before, [4]int{i, exampleRow.Operation[1], exampleRow.Operation[2], exampleRow.Operation[3]})

			if result == exampleRow.After {
				possibilities = append(possibilities, i)
			}
		}

		unionRes := union(possibilities, val)

		if len(unionRes) == 0 {
			fmt.Println(exampleRow.Before)
			fmt.Println(exampleRow.Operation)
			fmt.Println(exampleRow.After)
			fmt.Println(possibilities)
			fmt.Println(val)
			fmt.Println(m)
			panic("valore insulso")
		}

		m[originalOpCode] = unionRes
	}

	return m
}

func CountOpCodes(example []model.MachineInstructionExample) int {
	finalCount := 0

	for _, exampleRow := range example {
		possibilities := 0

		for i := 0; i <= 15; i++ {
			result := executeInstruction(exampleRow.Before, [4]int{i, exampleRow.Operation[1], exampleRow.Operation[2], exampleRow.Operation[3]})

			if result == exampleRow.After {
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
