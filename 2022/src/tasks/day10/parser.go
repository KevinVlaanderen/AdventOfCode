package day10

import (
	"strings"
)

type InstructionType int

const (
	NOOP InstructionType = iota
	ADDX
)

type Instruction struct {
	Type InstructionType
	Data []string
}

func Parse(data []string) (instructions []Instruction) {
	for _, line := range data {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		instructionType, _ := mapWordToInstructionType(parts[0])
		instruction := Instruction{Type: instructionType}

		switch parts[0] {
		case "noop":
		case "addx":
			instruction.Data = parts[1:]
		}

		instructions = append(instructions, instruction)
	}
	return
}

func mapWordToInstructionType(word string) (instructionType InstructionType, ok bool) {
	switch word {
	case "noop":
		return NOOP, true
	case "addx":
		return ADDX, true
	}
	return instructionType, false
}
