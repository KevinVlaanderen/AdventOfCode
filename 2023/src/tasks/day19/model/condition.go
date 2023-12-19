package model

import (
	"fmt"
	"strconv"
)

func ParseCondition(data string) Condition {
	category := Category(data[0])
	n, err := strconv.Atoi(data[2:])
	if err != nil {
		panic(err)
	}
	return Condition{category, Operator(data[1]), n}

}

type Condition struct {
	Category Category
	Operator Operator
	N        int
}

func (c Condition) Opposite() *Condition {
	switch c.Operator {
	case LessThan:
		return &Condition{c.Category, GreaterThan, c.N - 1}
	case GreaterThan:
		return &Condition{c.Category, LessThan, c.N + 1}
	default:
		panic("unknown operator")
	}
}

func (c Condition) String() string {
	return fmt.Sprintf("%v %v %v", c.Category, c.Operator, c.N)
}

type Operator string

const (
	LessThan    Operator = "<"
	GreaterThan          = ">"
)
