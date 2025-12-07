package day6

import (
	"2025/src/framework"
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

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	problems := parse(data)

	for _, problem := range problems {
		result.Value += lo.Reduce(problem.numbers, func(agg int, number int, index int) int {
			if index == 0 {
				return number
			}
			if problem.operation == Add {
				return agg + number
			} else if problem.operation == Multiply {
				return agg * number
			} else {
				panic("Invalid operation")
			}
		}, 0)
	}

	return
}

func parse(data string) (problems []Problem) {
	lines := framework.Lines(data)

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
