package day5

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"2022/src/tasks/day5/model"
	"2022/src/tasks/day5/mover"
)

type Mover interface {
	Move(n int, from int, to int)
}

func Task1(filePath string) (result test.TaskResult[string]) {
	data, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	storage := model.NewStorage(data[0])

	crateMover := mover.CrateMover9000{Storage: storage}
	executor := Executor{mover: &crateMover}

	instructions := model.NewInstructions(data[1])
	executor.ExecuteInstructions(instructions)

	for _, stack := range storage.Stacks {
		result.Value += stack.Top().Label
	}

	return
}

func Task2(filePath string) (result test.TaskResult[string]) {
	data, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	storage := model.NewStorage(data[0])

	crateMover := mover.CrateMover9001{Storage: storage}
	executor := Executor{mover: &crateMover}

	instructions := model.NewInstructions(data[1])
	executor.ExecuteInstructions(instructions)

	for _, stack := range storage.Stacks {
		result.Value += stack.Top().Label
	}

	return
}
