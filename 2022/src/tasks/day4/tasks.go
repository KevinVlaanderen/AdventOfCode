package day4

import (
	"2022/src/framework/tasks"
	"2022/src/tasks/day4/model"
	"fmt"
	"strings"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	pairs := tasks.ReadStream(filePath, parser)

	for pair := range pairs {
		if pair.Assignment1.Contains(pair.Assignment2) || pair.Assignment2.Contains(pair.Assignment1) {
			result.Value++
		}
	}

	return
}

func Task2(filePath string) (result tasks.TaskResult[int]) {
	pairs := tasks.ReadStream(filePath, parser)

	for pair := range pairs {
		if pair.Assignment1.OverlapsWith(pair.Assignment2) {
			result.Value++
		}
	}

	return
}

func parser(line string) (result model.Pair, hasResult bool, err error) {
	if line == "" {
		return result, false, nil
	}
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return result, false, fmt.Errorf("invalid input: %v", line)
	}

	var assignment1, assignment2 model.Assignment
	if assignment1, err = model.NewAssignment(parts[0]); err != nil {
		return result, false, fmt.Errorf("failed to parse range for assignment 1: %v", err)
	}
	if assignment2, err = model.NewAssignment(parts[1]); err != nil {
		return result, false, fmt.Errorf("failed to parse range for assignment 2: %v", err)
	}

	return model.Pair{Assignment1: assignment1, Assignment2: assignment2}, true, nil
}