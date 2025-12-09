package day6

import (
	"aoc/framework/tasks"
	"fmt"
	"go/types"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Problem struct {
	numbers   []int
	operation Operation
}

type Operation rune

const (
	Add      Operation = '+'
	Multiply           = '*'
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	problems := parse(data)

	for _, problem := range problems {
		result.Value += lo.Reduce(problem.numbers, reduceOperation(problem.operation), 0)
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	input := tasks.CharLines(data)
	numbers := input[:len(input)-1]
	operations := input[len(input)-1]

	widths := lo.Map(input, func(line []rune, index int) int {
		return len(line)
	})
	maxWidth := lo.Max(widths)
	operationsWidth := len(operations)

	var currentOperation Operation
	currentNumbers := make([]int, 0)

	for i := 0; i < maxWidth; i++ {
		if i < operationsWidth && operations[i] == '+' {
			currentOperation = Add
			currentNumbers = make([]int, 0)
		} else if i < operationsWidth && operations[i] == '*' {
			currentOperation = Multiply
			currentNumbers = make([]int, 0)
		}

		digits := make([]rune, 0)
		for _, numberLine := range numbers {
			if numberLine[i] != ' ' {
				digits = append(digits, numberLine[i])
			}
		}

		if len(digits) > 0 {
			number, _ := strconv.Atoi(string(digits))
			currentNumbers = append(currentNumbers, number)
		} else {
			result.Value += lo.Reduce(currentNumbers, reduceOperation(currentOperation), 0)
		}
	}

	result.Value += lo.Reduce(currentNumbers, reduceOperation(currentOperation), 0)

	return
}

func reduceOperation(currentOperation Operation) func(agg int, number int, index int) int {
	return func(agg int, number int, index int) int {
		if index == 0 {
			return number
		}
		if currentOperation == Add {
			return agg + number
		} else if currentOperation == Multiply {
			return agg * number
		} else {
			panic("Invalid operation")
		}
	}
}

func parse(data string) (problems []Problem) {
	lines := tasks.Lines(data)

	numbers := lo.Map(lines[:len(lines)-1], func(line string, index int) []int {
		return lo.Map(strings.Fields(line), func(number string, index int) int {
			n, _ := strconv.Atoi(number)
			return n
		})
	})
	operations := lo.Map(strings.Fields(lines[len(lines)-1]), func(operation string, index int) Operation {
		switch operation {
		case "+":
			return Add
		case "*":
			return Multiply
		default:
			panic(fmt.Sprintf("unknown operation '%s'", operation))
		}
	})

	for i := 0; i < len(numbers[0]); i++ {
		currentNumbers := make([]int, len(numbers))
		for j := 0; j < len(numbers); j++ {
			currentNumbers[j] = numbers[j][i]
		}
		problems = append(problems, Problem{
			numbers:   currentNumbers,
			operation: operations[i],
		})
	}
	return
}
