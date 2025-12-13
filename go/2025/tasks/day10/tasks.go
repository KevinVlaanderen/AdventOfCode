package day10

import (
	"aoc/framework/tasks"
	"go/types"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	machines := Parse(data)

	result.Value = lo.SumBy(machines, func(machine MachineDescription) int {
		presses, _ := pressLightButtonsUntilValid(machine.buttons, machine.targetLights, 0, 0, 0, NewCache())
		return presses
	})

	return
}

func pressLightButtonsUntilValid(buttons []int, target, lights, presses, best int, cache *Cache) (int, bool) {
	if cache == nil {
		return 0, false
	}

	cache.lightsSeen[lights] = struct{}{}

	endSuccess := false
	endResult := 0

	for _, button := range buttons {
		newPresses := presses + 1
		newLights := lights ^ button

		if _, ok := cache.lightsSeen[newLights]; ok {
			continue
		}

		found := false
		result := 0

		if best > 0 && newPresses >= best {
			continue
		} else if newLights == target {
			found = true
			result = newPresses
		} else if cached, ok := cache.pressesToSuccess[newLights]; ok {
			found = true
			result = cached + newPresses
		} else {
			result, found = pressLightButtonsUntilValid(buttons, target, newLights, newPresses, best, cache)
			if found {
				cache.pressesToSuccess[newLights] = result - newPresses
			}
		}

		if found {
			endSuccess = true
			if best == 0 || result < best {
				best = result
			}
			if endResult == 0 || result < endResult {
				endResult = result
			}
		}
	}

	delete(cache.lightsSeen, lights)

	return endResult, endSuccess
}

type Cache struct {
	lightsSeen       map[int]struct{}
	pressesToSuccess map[int]int
}

func NewCache() *Cache {
	return &Cache{
		lightsSeen:       make(map[int]struct{}),
		pressesToSuccess: make(map[int]int),
	}
}
