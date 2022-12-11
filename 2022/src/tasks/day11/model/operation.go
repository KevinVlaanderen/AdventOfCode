package model

import (
	"fmt"
	"regexp"
	"strconv"
)

type Expression string

type Operator int

const (
	ADD Operator = iota
	SUBTRACT
	DIVIDE
	MULTIPLY
)

var expressionPattern = regexp.MustCompile(`(.*) ([+-/*]) (.*)`)

func (o Expression) Perform(context map[string]int) (result int, err error) {
	matches := expressionPattern.FindStringSubmatch(string(o))

	var leftOperand, rightOperand int

	if left, err := strconv.Atoi(matches[1]); err != nil {
		leftOperand = context[matches[1]]
	} else {
		leftOperand = left
	}
	if right, err := strconv.Atoi(matches[3]); err != nil {
		rightOperand = context[matches[3]]
	} else {
		rightOperand = right
	}

	operator, _ := parseOperator(matches[2])

	switch operator {
	case ADD:
		result = leftOperand + rightOperand
	case SUBTRACT:
		result = leftOperand - rightOperand
	case DIVIDE:
		result = leftOperand / rightOperand
	case MULTIPLY:
		result = leftOperand * rightOperand
	}

	return result, fmt.Errorf("error in expression: %v", o)
}

func parseOperator(character string) (operator Operator, err error) {
	switch character {
	case "+":
		operator = ADD
	case "-":
		operator = SUBTRACT
	case "/":
		operator = DIVIDE
	case "*":
		operator = MULTIPLY
	default:
		err = fmt.Errorf("could not parse operator %v", character)
	}
	return
}
