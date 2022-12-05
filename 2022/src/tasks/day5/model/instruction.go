package model

import (
	"regexp"
	"strconv"
)

var instructionPattern = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

type Instruction struct {
	N    int
	From int
	To   int
}

func NewInstruction(line string) Instruction {
	matches := instructionPattern.FindStringSubmatch(line)
	n, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])
	return Instruction{
		N:    n,
		From: from - 1,
		To:   to - 1,
	}
}

func NewInstructions(lines []string) *[]Instruction {
	createdInstructions := make([]Instruction, 0)
	for _, instructionLine := range lines {
		createdInstructions = append(createdInstructions, NewInstruction(instructionLine))
	}
	return &createdInstructions
}
