package model

import (
	"strconv"
)

func ParseCondition(data string) Condition {
	category := Category(data[0])
	n, err := strconv.Atoi(data[2:])
	if err != nil {
		panic(err)
	}
	switch data[1] {
	case '<':
		return LessThan(category, n)
	case '>':
		return GreaterThan(category, n)
	default:
		panic("unknown operator")
	}
}

type Condition func(state *State) bool

func Always() Condition {
	return func(state *State) bool {
		return true
	}
}

func LessThan(category Category, n int) Condition {
	return func(state *State) bool {
		return state.Part.Ratings[category] < n
	}
}

func GreaterThan(category Category, n int) Condition {
	return func(state *State) bool {
		return state.Part.Ratings[category] > n
	}
}
