package day10

import (
	"2022/src/tasks/day10/model"
	"strconv"
)

type State struct {
	tick  int
	value int
}

type Task struct {
	cost  int
	ticks int
	run   func(state *State)
}

func createNoopTask() Task {
	return Task{
		cost: 1,
		run:  func(state *State) {},
	}
}

func createAddXTask(value int) Task {
	return Task{
		cost: 2,
		run: func(state *State) {
			state.value += value
		},
	}
}

type Processor struct {
	State
}

func NewProcessor(value int) *Processor {
	return &Processor{State{value: value}}
}

func (p *Processor) Execute(instructions []model.Instruction) <-chan State {
	stateChannel := make(chan State)

	if len(instructions) == 0 {
		close(stateChannel)
		return stateChannel
	}

	var instructionIndex int
	currentInstruction := instructions[instructionIndex]
	currentTask := p.createTask(currentInstruction)

	stop := make(chan struct{})
	clock := p.clock(stop)

	go func() {
		for tick := range clock {
			p.tick = tick
			stateChannel <- p.State

			currentTask.ticks++

			if currentTask.ticks < currentTask.cost {
				continue
			}

			currentTask.run(&p.State)
			instructionIndex++
			if instructionIndex < len(instructions) {
				currentInstruction = instructions[instructionIndex]
				currentTask = p.createTask(currentInstruction)
			} else {
				close(stop)
				break
			}
		}

		close(stateChannel)
	}()

	return stateChannel
}

func (p *Processor) createTask(instruction model.Instruction) Task {
	var task Task
	switch instruction.Type {
	case model.NOOP:
		task = createNoopTask()
	case model.ADDX:
		param, _ := strconv.Atoi(instruction.Data[0])
		task = createAddXTask(param)
	}
	return task
}

func (p *Processor) clock(stop <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		i := 0
		for true {
			select {
			case ch <- i:
				i++
			case <-stop:
				return
			}
		}
	}()
	return ch
}
