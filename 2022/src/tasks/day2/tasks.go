package day2

import (
	"2022/src/framework/tasks"
	"2022/src/tasks/day2/model"
	"fmt"
	"strings"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	rounds := tasks.ReadStream(filePath, task1Parser)

	for round := range rounds {
		if score, err := round.ScoreFor(model.Player2); err != nil {
			result.Error = err
			return
		} else {
			result.Value += score
		}
	}

	return
}

func Task2(filePath string) (result tasks.TaskResult[int]) {
	predictions := tasks.ReadStream(filePath, task2Parser)

	for prediction := range predictions {
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

func task1Parser(line string) (result model.Round, hasResult bool, err error) {
	if line == "" {
		return result, false, nil
	}
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return result, false, fmt.Errorf("invalid input: %v", line)
	}
	round := model.NewRound(
		model.ParseHandFor(model.Player1, parts[0]),
		model.ParseHandFor(model.Player2, parts[1]),
	)

	return round, true, nil
}

func task2Parser(line string) (result model.Prediction, hasResult bool, err error) {
	if line == "" {
		return result, false, nil
	}

	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return result, false, fmt.Errorf("invalid input: %v", line)
	}

	roundPrediction := model.NewPrediction(
		model.ParseHandFor(model.Player1, parts[0]),
		model.ParseDesiredResult(parts[1]),
	)

	return roundPrediction, true, nil
}
