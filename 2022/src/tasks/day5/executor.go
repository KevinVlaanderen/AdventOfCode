package day5

import "2022/src/tasks/day5/model"

type Executor struct {
	mover Mover
}

func (e Executor) ExecuteInstructions(instructions *[]model.Instruction) {
	for _, instruction := range *instructions {
		e.mover.Move(instruction.N, instruction.From, instruction.To)
	}
}
