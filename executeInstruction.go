package main

func executeInstruction(memory [4]int, operation [4]int) [4]int {
	opcode, A, B, C := operation[0], operation[1], operation[2], operation[3]
	switch opcode {
	case 0:
		return addr(memory, A, B, C)
	case 1:
		return addi(memory, A, B, C)
	case 2:
		return mulr(memory, A, B, C)
	case 3:
		return muli(memory, A, B, C)
	case 4:
		return banr(memory, A, B, C)
	case 5:
		return bani(memory, A, B, C)
	case 6:
		return setr(memory, A, B, C)
	case 7:
		return seti(memory, A, B, C)
	case 8:
		return gtir(memory, A, B, C)
	case 9:
		return gtri(memory, A, B, C)
	case 10:
		return gtrr(memory, A, B, C)
	case 11:
		return eqir(memory, A, B, C)
	case 12:
		return eqri(memory, A, B, C)
	case 13:
		return eqrr(memory, A, B, C)
	case 14:
		return bori(memory, A, B, C)
	case 15:
		return borr(memory, A, B, C)
	default:
		panic("codice non trovato")
	}

}

func executeInstructionWithAlias(memory [4]int, operation [4]int) [4]int {
	switch operation[0] {
	case 0:
		operation[0] = 2
		return executeInstruction(memory, operation)
	case 1:
		operation[0] = 12
		return executeInstruction(memory, operation)
	case 2:
		operation[0] = 6
		return executeInstruction(memory, operation)
	case 3:
		operation[0] = 13
		return executeInstruction(memory, operation)
	case 4:
		operation[0] = 10
		return executeInstruction(memory, operation)
	case 5:
		operation[0] = 3
		return executeInstruction(memory, operation)
	case 6:
		operation[0] = 15
		return executeInstruction(memory, operation)
	case 7:
		operation[0] = 5
		return executeInstruction(memory, operation)
	case 8:
		operation[0] = 0
		return executeInstruction(memory, operation)
	case 9:
		operation[0] = 4
		return executeInstruction(memory, operation)
	case 10:
		operation[0] = 11
		return executeInstruction(memory, operation)
	case 11:
		operation[0] = 8
		return executeInstruction(memory, operation)
	case 12:
		operation[0] = 1
		return executeInstruction(memory, operation)
	case 13:
		operation[0] = 9
		return executeInstruction(memory, operation)
	case 14:
		operation[0] = 7
		return executeInstruction(memory, operation)
	case 15:
		operation[0] = 14
		return executeInstruction(memory, operation)
	default:
		panic("codice non trovato")
	}

}
