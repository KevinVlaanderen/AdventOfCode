//go:generate go run aoc/cmd/generate-tests

package day2

import (
	"aoc/2022/tasks/day2/model"
	"aoc/framework/tasks"
	"fmt"
	"go/types"
	"strings"

	"github.com/samber/lo"
)

// Task1 type:mock 	file:data	expected:15
// Task1 type:real 	file:day2 	expected:11475
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	rounds := parse1(data)

	for _, round := range rounds {
		if score, err := round.ScoreFor(model.Player2); err != nil {
			result.Error = err
			return
		} else {
			result.Value += score
		}
	}

	return
}

// Task2 type:mock 	file:data	expected:12
// Task2 type:real 	file:day2 	expected:16862
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	predictions := parse2(data)

	for _, prediction := range predictions {
		if prediction.Outcome == model.InvalidResult {
			result.Error = fmt.Errorf("invalid desired result: %v", prediction.Outcome)
			return
		}

		round := model.NewRound(prediction.Player1.Hand, prediction.RequiredHand())

		if score, err := round.ScoreFor(model.Player2); err != nil {
			result.Error = err
			return
		} else {
			result.Value += score
		}
	}

	return
}

func parse1(data string) []model.Round {
	return lo.Map(tasks.Lines(data), func(line string, index int) model.Round {
		parts := strings.Split(line, " ")
		return model.NewRound(
			model.ParseHandFor(model.Player1, parts[0]),
			model.ParseHandFor(model.Player2, parts[1]),
		)
	})
}

func parse2(data string) []model.Prediction {
	return lo.Map(tasks.Lines(data), func(line string, index int) model.Prediction {
		parts := strings.Split(line, " ")
		return model.NewPrediction(
			model.ParseHandFor(model.Player1, parts[0]),
			model.ParseDesiredResult(parts[1]),
		)
	})
}
