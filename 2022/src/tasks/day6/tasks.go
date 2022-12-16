package day6

import (
	"2022/src/framework/tasks"
	"2022/src/tasks/day6/model"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	line := <- tasks.ReadStream(filePath, parser)

	marker := line.FindMarker(4)

	return tasks.TaskResult[int]{Value: marker}
}

func Task2(filePath string) (result tasks.TaskResult[int]) {
	line := <- tasks.ReadStream(filePath, parser)

	marker := line.FindMarker(14)

	return tasks.TaskResult[int]{Value: marker}
}

func parser(line string) (result model.Signal, hasResult bool, err error) {
	if line == "" {
		return result, false, nil
	}
	return model.Signal(line), true, nil
}