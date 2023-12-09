package day3

import (
	"2023/src/framework"
	"2023/src/tasks/day3/model"
	"golang.org/x/exp/slices"
)

func Task1(filePath string) (result framework.Result[int]) {
	schematic := <-framework.ReadStream(filePath, createParser())

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

func Task2(filePath string) (result framework.Result[int]) {
	schematic := <-framework.ReadStream(filePath, createParser())

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

func createParser() framework.Parser[model.Schematic] {
	schematic := model.NewSchematic()

	return func(line string) (result model.Schematic, hasResult bool, err error) {
		if line == "" {
			return schematic, true, nil
		}

		schematic.Add(line)
		return schematic, false, nil
	}
}
