package day9

import (
	"aoc/2022/tasks/day9/model"
	"aoc/framework"
	"go/types"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	instructions := parse(data)

	rope := model.NewRope(2)
	visitedGrid := model.Grid{rope.TailPosition(): true}

	for _, instruction := range instructions {
		for range framework.Range(0, instruction.Steps, 1) {
			rope.Move(instruction.Direction)
			visitedGrid[rope.TailPosition()] = true
		}
	}

	result.Value = len(visitedGrid)

	// visitedGrid.PrintVisited()

	return
}

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	instructions := parse(data)

	rope := model.NewRope(10)
	visitedGrid := model.Grid{rope.TailPosition(): true}

	for _, instruction := range instructions {
		for range framework.Range(0, instruction.Steps, 1) {
			rope.Move(instruction.Direction)
			visitedGrid[rope.TailPosition()] = true
		}
	}

	result.Value = len(visitedGrid)

	// visitedGrid.PrintVisited()

	return
}

var linePattern = regexp.MustCompile(`^([LRUD]) (\d+)$`)

func parse(data string) []model.Instruction {
	return lo.Map(framework.Lines(data), func(line string, index int) model.Instruction {
		matches := linePattern.FindStringSubmatch(line)
		direction := model.ToDirection(matches[1])
		number, _ := strconv.Atoi(matches[2])

		return model.Instruction{
			Direction: direction,
			Steps:     number,
		}
	})
}
