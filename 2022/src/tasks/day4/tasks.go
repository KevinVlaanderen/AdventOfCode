package day4

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"fmt"
	"strconv"
	"strings"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	for _, line := range data {
		parts := strings.Split(line, ",")
		var assignment1, assignment2 Assignment
		if assignment1, err = NewAssignment(parts[0]); err != nil {
			result.Error = fmt.Errorf("failed to parse range for assignment 1: %v", err)
			return
		}
		if assignment2, err = NewAssignment(parts[1]); err != nil {
			result.Error = fmt.Errorf("failed to parse range for assignment 2: %v", err)
			return
		}

		if assignment1.Contains(assignment2) || assignment2.Contains(assignment1) {
			result.Value++
		}
	}

	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	for _, line := range data {
		parts := strings.Split(line, ",")
		var assignment1, assignment2 Assignment
		if assignment1, err = NewAssignment(parts[0]); err != nil {
			result.Error = fmt.Errorf("failed to parse range for assignment 1: %v", err)
			return
		}
		if assignment2, err = NewAssignment(parts[1]); err != nil {
			result.Error = fmt.Errorf("failed to parse range for assignment 2: %v", err)
			return
		}

		if assignment1.OverlapsWith(assignment2) {
			result.Value++
		}
	}

	return
}

type Assignment struct {
	low  int
	high int
}

func NewAssignment(input string) (assignment Assignment, err error) {
	parts := strings.Split(input, "-")
	if assignment.low, err = strconv.Atoi(parts[0]); err != nil {
		return
	}
	if assignment.high, err = strconv.Atoi(parts[1]); err != nil {
		return
	}

	return
}

func (a Assignment) Contains(b Assignment) bool {
	return a.low <= b.low && a.high >= b.high
}

func (a Assignment) OverlapsWith(b Assignment) bool {
	range1 := makeRange(a.low, a.high)
	range2 := makeRange(b.low, b.high)

	if intersection, err := framework.Intersection(range1, range2); err != nil {
		return false
	} else {
		return len(intersection) > 0
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
