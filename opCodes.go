package main

func addr(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = outputMemory[A] + outputMemory[B]

	return outputMemory
}

func addi(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = outputMemory[A] + B

	return outputMemory
}

func mulr(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = outputMemory[A] * outputMemory[B]

	return outputMemory
}

func muli(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = outputMemory[A] * B

	return outputMemory
}

func banr(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = outputMemory[A] & outputMemory[B]

	return outputMemory
}

func bani(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = outputMemory[A] & B

	return outputMemory
}

func borr(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = outputMemory[A] | outputMemory[B]

	return outputMemory
}

func bori(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = outputMemory[A] | B

	return outputMemory
}

func setr(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = outputMemory[A]

	return outputMemory
}

func seti(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	outputMemory[C] = A

	return outputMemory
}

func gtir(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	if A > outputMemory[B] {
		outputMemory[C] = 1
	} else {
		outputMemory[C] = 0
	}

	return outputMemory
}

func gtri(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	if outputMemory[A] > B {
		outputMemory[C] = 1
	} else {
		outputMemory[C] = 0
	}

	return outputMemory
}
func gtrr(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	if outputMemory[A] > outputMemory[B] {
		outputMemory[C] = 1
	} else {
		outputMemory[C] = 0
	}

	return outputMemory
}

func eqir(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	if A == outputMemory[B] {
		outputMemory[C] = 1
	} else {
		outputMemory[C] = 0
	}

	return outputMemory
}

func eqri(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	if outputMemory[A] == B {
		outputMemory[C] = 1
	} else {
		outputMemory[C] = 0
	}

	return outputMemory
}
func eqrr(memory [4]int, A, B, C int) [4]int {
	outputMemory := memory

	if outputMemory[A] == outputMemory[B] {
		outputMemory[C] = 1
	} else {
		outputMemory[C] = 0
	}

	return outputMemory
}
