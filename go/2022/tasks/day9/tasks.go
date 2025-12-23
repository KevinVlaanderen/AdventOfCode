//go:generate go run aoc/cmd/generate-tests

package day9

import (
	"aoc/2022/tasks/day9/model"
	"aoc/framework"
	"aoc/framework/tasks"
	"go/types"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

// Task1 type:mock 	file:data1	expected:13
// Task1 type:real 	file:day9 	expected:6470

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
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

	return
}

// Task2 type:mock 	file:data1	expected:1
// Task2 type:mock 	file:data2	expected:36
// Task2 type:real 	file:day9 	expected:2658
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
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

	return
}

var linePattern = regexp.MustCompile(`^([LRUD]) (\d+)$`)

func parse(data string) []model.Instruction {
	return lo.Map(tasks.Lines(data), func(line string, index int) model.Instruction {
		matches := linePattern.FindStringSubmatch(line)
		direction := model.ToDirection(matches[1])
		number, _ := strconv.Atoi(matches[2])

		return model.Instruction{
			Direction: direction,
			Steps:     number,
		}
	})
}
