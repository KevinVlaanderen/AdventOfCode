package day15

import (
	"2023/src/framework"
	"2023/src/tasks/day15/model"
	"github.com/elliotchance/orderedmap/v2"
	"github.com/samber/lo"
	"go/types"
	"strings"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	words := strings.Split(framework.Lines(data)[0], ",")
	hashCalculator := createHashCalculator()

	for _, word := range words {
		result.Value += int(hashCalculator(word))
	}

	return
}

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	hashCalculator := createHashCalculator()
	instructions := lo.Map(strings.Split(framework.Lines(data)[0], ","), func(word string, index int) model.Instruction {
		return model.NewInstruction(word, hashCalculator)
	})

	boxes := make([]model.Box, 256)
	for i := range boxes {
		boxes[i].Lenses = orderedmap.NewOrderedMap[string, uint8]()
	}

	for _, instruction := range instructions {
		performInstruction(instruction, &boxes)
	}

	for boxIndex, box := range boxes {
		for lensIndex, lensKey := range box.Lenses.Keys() {
			var focalLength uint8
			var found bool
			if focalLength, found = box.Lenses.Get(lensKey); !found {
				panic("could not get lens")
			}
			result.Value += (1 + boxIndex) * (1 + lensIndex) * int(focalLength)
		}
	}

	return
}

func performInstruction(instruction model.Instruction, boxes *[]model.Box) {
	switch instruction.Operation {
	case model.Remove:
		if _, found := (*boxes)[instruction.Box].Lenses.Get(instruction.Label); found {
			(*boxes)[instruction.Box].Lenses.Delete(instruction.Label)
		}
	case model.Set:
		(*boxes)[instruction.Box].Lenses.Set(instruction.Label, instruction.FocalLength)
	}
}

func createHashCalculator() func(word string) uint8 {
	var text strings.Builder
	return func(word string) uint8 {
		text.Reset()

		var current uint8
		for _, char := range word {
			current += uint8(char)
			current *= 17
			current = uint8(uint(current) % 256)
		}
		return current
	}
}
