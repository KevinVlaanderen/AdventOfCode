package day2

import (
	"2022/src/framework"
	"fmt"
	"strings"
)

type Result int

const (
	Draw Result = iota
	Win
	Loss
)

type Hand int

const (
	Rock Hand = iota
	Paper
	Scissors
)

var (
	opponentHandMap = map[string]Hand{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}
	playerHandMap = map[string]Hand{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}
)

func Task1(filePath string) (result framework.TaskResult) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		return framework.TaskResult{Error: err}
	}

	for _, line := range data {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return framework.TaskResult{Error: fmt.Errorf("invalid input: %v", line)}
		}

		opponentHand, ok := parseHand(parts[0], opponentHandMap)
		if !ok {
			return framework.TaskResult{Error: fmt.Errorf("invalid value: %v", parts[0])}
		}
		playerHand, ok := parseHand(parts[1], playerHandMap)
		if !ok {
			return framework.TaskResult{Error: fmt.Errorf("invalid value: %v", parts[1])}
		}

		roundResult, ok := determineResult(opponentHand, playerHand)
		if !ok {
			return framework.TaskResult{Error: fmt.Errorf("invalid result: %v vs %v", opponentHand, playerHand)}
		}

		result.Value += calculateScore(roundResult, playerHand)
	}

	return
}

func Task2(filePath string) (result framework.TaskResult) {
	return
}

func determineResult(a Hand, b Hand) (Result, bool) {
	switch {
	case a == b:
		return Draw, true
	case a == Rock && b == Paper:
		return Win, true
	case a == Rock && b == Scissors:
		return Loss, true
	case a == Paper && b == Rock:
		return Loss, true
	case a == Paper && b == Scissors:
		return Win, true
	case a == Scissors && b == Rock:
		return Win, true
	case a == Scissors && b == Paper:
		return Loss, true
	}
	return Draw, false
}

func calculateScore(result Result, hand Hand) (score int) {
	switch result {
	case Loss:
	case Draw:
		score += 3
	case Win:
		score += 6
	}
	switch hand {
	case Rock:
		score += 1
	case Paper:
		score += 2
	case Scissors:
		score += 3
	}
	return
}

func parseHand(symbol string, mapping map[string]Hand) (Hand, bool) {
	c, ok := mapping[symbol]
	return c, ok
}
