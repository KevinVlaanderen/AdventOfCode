package model

import (
	"2022/src/framework/generators"
	"2022/src/framework/sets"
	"strconv"
	"strings"
)

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
	range1 := generators.Range(a.low, a.high-a.low+1, 1)
	range2 := generators.Range(b.low, b.high-b.low+1, 1)

	if intersection, err := sets.Intersection(range1, range2); err != nil {
		return false
	} else {
		return len(intersection) > 0
	}
}
