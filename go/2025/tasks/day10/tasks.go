package day10

import (
	"aoc/framework/tasks"
	"go/types"

	"github.com/oleiade/lane/v2"
	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	machines := Parse(data)

	result.Value = lo.SumBy(machines, func(machine MachineDescription) int {
		return pressLightButtonsUntilValid(machine.buttons, machine.targetLights)
	})

	return
}

func pressLightButtonsUntilValid(buttons []int, target int) int {
	best := 0
	lightsSeen := make(map[int]struct{})

	type State struct {
		lights  int
		presses int
	}

	queue := lane.NewQueue(State{lights: 0, presses: 0})

	for queue.Size() > 0 {
		state, _ := queue.Dequeue()

		if _, ok := lightsSeen[state.lights]; ok {
			continue
		}

		lightsSeen[state.lights] = struct{}{}

		if state.lights == target {
			if best == 0 || state.presses < best {
				best = state.presses
			}
			continue
		} else if best > 0 && state.presses > best {
			continue
		}

		for _, button := range buttons {
			newPresses := state.presses + 1
			newLights := state.lights ^ button

			queue.Enqueue(State{lights: newLights, presses: newPresses})
		}
	}

	return best
}
