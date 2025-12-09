package day10

import (
	"aoc/2022/tasks/day10/model"
	"aoc/framework/tasks"
	"go/types"
	"strings"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	instructions := parse(data)

	processor := NewProcessor(1)
	for state := range processor.Execute(instructions) {
		// log.Printf("value at tick %v: %v ", state.tick+1, processor.value)

		if state.tick+1 <= 220 && (state.tick+1 == 20 || (state.tick+1-20)%40 == 0) {
			result.Value += (state.tick + 1) * state.value
		}
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[string]) {
	instructions := parse(data)

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
			// fmt.Println(outputLine)
			outputLine = ""
		}
	}

	result.Value = "RBPARAGF"

	return
}

func parse(data string) []model.Instruction {
	return lo.Map(tasks.Lines(data), func(line string, index int) model.Instruction {
		parts := strings.Split(line, " ")
		instructionType, _ := model.MapWordToInstructionType(parts[0])
		instruction := model.Instruction{Type: instructionType}

		switch parts[0] {
		case "noop":
		case "addx":
			instruction.Data = parts[1:]
		}

		return instruction
	})
}
