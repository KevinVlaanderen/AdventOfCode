package day2

import (
	"2022/src/framework"
	"2022/src/tasks/day2/model"
	"fmt"
	"strings"
)

func Task1(filePath string) (result framework.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	for _, line := range data {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			result.Error = fmt.Errorf("invalid input: %v", line)
			return
		}

		round := model.NewRound(
			model.ParseHandFor(model.Player1, parts[0]),
			model.ParseHandFor(model.Player2, parts[1]),
		)

		if score, err := round.ScoreFor(model.Player2); err != nil {
			result.Error = err
			return
		} else {
			result.Value += score
		}
	}

	return
}

func Task2(filePath string) (result framework.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	for _, line := range data {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			result.Error = fmt.Errorf("invalid input: %v", line)
			return
		}

		desiredResult := getDesiredResult(parts[1])
		if desiredResult == invalid {
			result.Error = fmt.Errorf("invalid desired result: %v", parts[1])
			return
		}

		player1Hand := model.ParseHandFor(model.Player1, parts[0])
		player2Hand := requiredResultTo(desiredResult, player1Hand)

		round := model.NewRound(player1Hand, player2Hand)

		if score, err := round.ScoreFor(model.Player2); err != nil {
			result.Error = err
			return
		} else {
			result.Value += score
		}
	}

	return
}

type desiredResult int

const (
	invalid desiredResult = iota
	draw
	win
	loss
)

func requiredResultTo(result desiredResult, against model.Hand) model.Hand {
	switch result {
	case win:
		return against.LosesAgainst()
	case loss:
		return against.WinsAgainst()
	case draw:
		return against
	}
	return model.Invalid
}

var resultMap = map[string]desiredResult{
	"X": loss,
	"Y": draw,
	"Z": win,
}

func getDesiredResult(symbol string) desiredResult {
	if c, ok := resultMap[symbol]; ok {
		return c
	} else {
		return invalid
	}
}
