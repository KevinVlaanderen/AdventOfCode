package day1

import (
	"2023/src/framework/parser"
	"2023/src/framework/task"
	"strconv"
	"unicode"
)

func Task1(filePath string) (result task.TaskResult[int]) {
	for line := range task.ReadStream(filePath, parser.Lines) {
		result.Value += calculateCalibrationValue(line)
	}
	return
}

func calculateCalibrationValue(line string) int {
	firstDigit, _ := findFirstDigit(line)
	lastDigit, _ := findFirstDigit(reverse(line))

	if value, err := strconv.Atoi(string(firstDigit) + string(lastDigit)); err != nil {
		panic(err)
	} else {
		return value
	}
}

func findFirstDigit(input string) (char rune, index int) {
	for index, char = range input {
		if unicode.IsDigit(char) {
			return
		}
	}
	return
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
