package model

type InstructionType int

const (
	NOOP InstructionType = iota
	ADDX
)

type Instruction struct {
	Type InstructionType
	Data []string
}

func MapWordToInstructionType(word string) (instructionType InstructionType, ok bool) {
	switch word {
	case "noop":
		return NOOP, true
	case "addx":
		return ADDX, true
	}
	return instructionType, false
}
