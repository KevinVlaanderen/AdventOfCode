package day5

import (
	"2022/src/framework"
	"2022/src/tasks/day5/model"
)

func Task1(filePath string) (result framework.TaskResult[string]) {
	data, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	storage := model.NewStorage(data[0])

	for _, instructionLine := range data[1] {
		instruction := model.NewInstruction(instructionLine)
		storage.ExecuteInstruction(instruction)
	}

	for _, stack := range storage.Stacks {
		result.Value += stack.Top().Label
	}

	return
}

func Task2(filePath string) (result framework.TaskResult[string]) {
	_, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	return
}
