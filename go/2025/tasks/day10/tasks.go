package day10

import (
	"aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	machineDescriptions := Parse(data)

	result.Value = lo.SumBy(machineDescriptions, func(machineDescription MachineDescription) int {
		_, presses := pressLightButtonsUntilValid(&machineDescription, 0, 0, NewCache())
		return presses
	})

	return
}

func pressLightButtonsUntilValid(description *MachineDescription, lights, presses int, cache *Cache) (bool, int) {
	if cache == nil {
		return false, 0
	}

	cache.lightsSeen[lights] = struct{}{}

	endSuccess := false
	endPresses := 0

	for button := 0; button < len(description.ButtonWiringSchematics); button++ {
		newPresses := presses + 1
		newLights := lights
		for _, lightIndex := range description.ButtonWiringSchematics[button] {
			newLights ^= math.PowInt(2, lightIndex)
		}

		if _, ok := cache.lightsSeen[newLights]; ok {
			continue
		}

		found := false
		result := 0

		if cache.best > 0 && newPresses >= cache.best {
			continue
		} else if newLights == description.IndicatorLightDiagram {
			found = true
			result = newPresses
		} else if cached, ok := cache.pressesToSuccess[newLights]; ok {
			found = true
			result = cached + newPresses
		} else {
			found, result = pressLightButtonsUntilValid(description, newLights, newPresses, cache)
			if found {
				cache.pressesToSuccess[newLights] = result - newPresses
			}
		}

		if found {
			endSuccess = true
			if endPresses == 0 || result < endPresses {
				endPresses = result
			}
			if cache.best == 0 || result < cache.best {
				cache.best = result
			}
		}

	}

	delete(cache.lightsSeen, lights)

	return endSuccess, endPresses
}

type Cache struct {
	best             int
	lightsSeen       map[int]struct{}
	pressesToSuccess map[int]int
}

func NewCache() *Cache {
	return &Cache{
		lightsSeen:       make(map[int]struct{}),
		pressesToSuccess: make(map[int]int),
	}
}
