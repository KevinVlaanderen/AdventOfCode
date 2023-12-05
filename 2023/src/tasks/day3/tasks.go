package day3

import (
	"2023/src/framework/task"
	"2023/src/tasks/day3/model"
)

func Task1(filePath string) (result task.Result[int]) {
	schematic := <-task.ReadStream(filePath, createParser())

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

func createParser() task.Parser[model.Schematic] {
	schematic := model.NewSchematic()

	return func(line string) (result model.Schematic, hasResult bool, err error) {
		if line == "" {
			return schematic, true, nil
		}

		schematic.Add(line)
		return schematic, false, nil
	}
}
