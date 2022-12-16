package day10

import (
	"2022/src/framework/tasks"
	"2022/src/tasks/day10/model"
	"fmt"
	"log"
	"strings"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	instructions := tasks.Read(filePath, parser)

	processor := NewProcessor(1)
	for state := range processor.Execute(instructions) {
		log.Printf("value at tick %v: %v ", state.tick+1, processor.value)

		if state.tick+1 <= 220 && (state.tick+1 == 20 || (state.tick+1-20)%40 == 0) {
			result.Value += (state.tick + 1) * state.value
		}
	}

	return
}

func Task2(filePath string) (result tasks.TaskResult[string]) {
	instructions := tasks.Read(filePath, parser)

	var outputLine string

	processor := NewProcessor(1)
	for state := range processor.Execute(instructions) {
		column := state.tick % 40

		if column >= state.value-1 && column <= state.value+1 {
			outputLine += "#"
		} else {
			outputLine += "."
		}

		if column == 39 {
			fmt.Println(outputLine)
			outputLine = ""
		}
	}

	result.Value = "RBPARAGF"

	return
}

func parser(line string) (result model.Instruction, hasResult bool, err error) {
	if line == "" {
		return result, false, nil
	}

	parts := strings.Split(line, " ")
	instructionType, _ := model.MapWordToInstructionType(parts[0])
	instruction := model.Instruction{Type: instructionType}

	switch parts[0] {
	case "noop":
	case "addx":
		instruction.Data = parts[1:]
	}

	return instruction, true, nil
}