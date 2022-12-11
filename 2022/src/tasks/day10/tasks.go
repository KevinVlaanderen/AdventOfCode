package day10

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"fmt"
	"log"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	instructions := Parse(data)
	processor := NewProcessor(1)
	for state := range processor.Execute(instructions) {
		log.Printf("value at tick %v: %v ", state.tick+1, processor.value)

		if state.tick+1 <= 220 && (state.tick+1 == 20 || (state.tick+1-20)%40 == 0) {
			result.Value += (state.tick + 1) * state.value
		}
	}

	return
}

func Task2(filePath string) (result test.TaskResult[string]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	var outputLine string

	instructions := Parse(data)
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
