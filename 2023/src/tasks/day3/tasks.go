package day3

import (
	"2023/src/framework"
	"2023/src/tasks/day3/model"
	"golang.org/x/exp/slices"
)

func Task1(data string) (result framework.Result[int]) {
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

func Task2(data string) (result framework.Result[int]) {
	schematic := parse(data)

	for _, symbol := range schematic.Symbols {
		if symbol.Value != '*' {
			continue
		}

		numbers := make([]model.Number, 0)

		for _, coordinate := range symbol.Position.Neighbors() {
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

	for _, line := range framework.Lines(data) {
		schematic.Add(line)
	}

	return schematic
}
