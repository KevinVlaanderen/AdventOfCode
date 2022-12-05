package day5

import (
	"2022/src/framework"
)

func Task1(filePath string) (result framework.TaskResult[string]) {
	_, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	return
}

func Task2(filePath string) (result framework.TaskResult[string]) {
	_, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	return
}
