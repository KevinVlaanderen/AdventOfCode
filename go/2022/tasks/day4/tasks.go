//go:generate go run aoc/cmd/generate-tests

package day4

import (
	"aoc/2022/tasks/day4/model"
	"aoc/framework/tasks"
	"go/types"
	"strings"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	pairs := parse(data)

	for _, pair := range pairs {
		if pair.Assignment1.Contains(pair.Assignment2) || pair.Assignment2.Contains(pair.Assignment1) {
			result.Value++
		}
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	pairs := parse(data)

	for _, pair := range pairs {
		if pair.Assignment1.OverlapsWith(pair.Assignment2) {
			result.Value++
		}
	}

	return
}

func parse(data string) []model.Pair {
	return lo.Map(tasks.Lines(data), func(line string, index int) model.Pair {
		parts := strings.Split(line, ",")

		assignment1, _ := model.NewAssignment(parts[0])
		assignment2, _ := model.NewAssignment(parts[1])

		return model.Pair{Assignment1: assignment1, Assignment2: assignment2}
	})
}
