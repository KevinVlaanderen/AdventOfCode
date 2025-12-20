package day3

import (
	"aoc/2023/tasks/day3/model"
	"aoc/framework/geometry/geo2d"
	"aoc/framework/tasks"
	"go/types"

	"golang.org/x/exp/slices"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	schematic := parse(data)

	for _, number := range schematic.Numbers {
	Lookup:
		for _, coordinate := range number.Area.Neighbors() {
			if _, found := schematic.SymbolAt(coordinate); found {
				result.Value += number.Value
				break Lookup
			}
		}
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	schematic := parse(data)

	for _, symbol := range schematic.Symbols {
		if symbol.Value != '*' {
			continue
		}

		numbers := make([]model.Number, 0)

		for _, coordinate := range symbol.Position.Neighbors(geo2d.All) {
			if number, found := schematic.NumberAt(coordinate); found && !slices.Contains(numbers, number) {
				numbers = append(numbers, number)
			}
		}

		if len(numbers) == 2 {
			result.Value += numbers[0].Value * numbers[1].Value
		}
	}

	return
}

func parse(data string) model.Schematic {
	schematic := model.NewSchematic()

	for _, line := range tasks.Lines(data) {
		schematic.Add(line)
	}

	return schematic
}
