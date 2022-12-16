package day9

import (
	"2022/src/framework/generators"
	"2022/src/framework/tasks"
	"2022/src/tasks/day9/model"
	"regexp"
	"strconv"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	instructions := tasks.ReadStream(filePath, parser)

	rope := model.NewRope(2)
	visitedGrid := model.Grid{rope.TailPosition(): true}

	for instruction := range instructions {
		for range generators.Range(0, instruction.Steps, 1) {
			rope.Move(instruction.Direction)
			visitedGrid[rope.TailPosition()] = true
		}
	}

	result.Value = len(visitedGrid)

	visitedGrid.PrintVisited()

	return
}

func Task2(filePath string) (result tasks.TaskResult[int]) {
	instructions := tasks.ReadStream(filePath, parser)

	rope := model.NewRope(10)
	visitedGrid := model.Grid{rope.TailPosition(): true}

	for instruction := range instructions {
		for range generators.Range(0, instruction.Steps, 1) {
			rope.Move(instruction.Direction)
			visitedGrid[rope.TailPosition()] = true
		}
	}

	result.Value = len(visitedGrid)

	visitedGrid.PrintVisited()

	return
}

var linePattern = regexp.MustCompile(`^([LRUD]) (\d+)$`)

func parser(line string) (result model.Instruction, hasResult bool, err error) {
	if line == "" {
		return result, false, nil
	}

	matches := linePattern.FindStringSubmatch(line)
	direction := model.ToDirection(matches[1])
	number, _ := strconv.Atoi(matches[2])

	return model.Instruction{
		Direction: direction,
		Steps:     number,
	}, true, nil
}