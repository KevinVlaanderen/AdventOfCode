//go:generate go run aoc/cmd/generate-tests

package day5

import (
	"aoc/2022/tasks/day5/model"
	"aoc/2022/tasks/day5/mover"
	"aoc/framework"
	"aoc/framework/tasks"
	"go/types"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

type Mover interface {
	Move(n int, from int, to int)
}

func Task1(data string, _ types.Nil) (result tasks.Result[string]) {
	storage, instructions := parse(data)

	crateMover := &mover.CrateMover9000{Storage: storage}

	for _, instruction := range instructions {
		crateMover.Move(instruction.N, instruction.From, instruction.To)
	}

	for _, stack := range storage.Stacks {
		result.Value += stack.Top().Label
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[string]) {
	storage, instructions := parse(data)

	crateMover := &mover.CrateMover9001{Storage: storage}

	for _, instruction := range instructions {
		crateMover.Move(instruction.N, instruction.From, instruction.To)
	}

	for _, stack := range storage.Stacks {
		result.Value += stack.Top().Label
	}

	return
}

var numberPattern = regexp.MustCompile(`\d+`)
var cratePattern = regexp.MustCompile(`\[(\w+)]`)

func parse(data string) (*model.Storage, []model.Instruction) {
	blocks := tasks.LineBlocks(data)

	storage := parseStorage(blocks[0])
	instructions := parseInstructions(blocks[1])

	return storage, instructions
}

func parseStorage(block []string) *model.Storage {
	numberLine := block[len(block)-1]
	stackData := block[0 : len(block)-1]
	numberMatches := numberPattern.FindAllString(numberLine, -1)
	numStacks := len(numberMatches)

	stacks := make([]model.Stack, numStacks)

	for _, lineIndex := range framework.RangeSlice(len(stackData)-1, len(stackData), -1) {
		for _, stackIndex := range framework.RangeSlice(0, numStacks, 1) {
			crateIndex := stackIndex * 4
			crateString := stackData[lineIndex][crateIndex : crateIndex+3]

			crateMatches := cratePattern.FindStringSubmatch(crateString)
			if len(crateMatches) < 2 {
				continue
			}

			stacks[stackIndex].Crates = append(stacks[stackIndex].Crates, model.Crate{Label: crateMatches[1]})
		}
	}

	return &model.Storage{Stacks: stacks}
}

var instructionPattern = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func parseInstructions(block []string) []model.Instruction {
	return lo.Map(block, func(line string, index int) model.Instruction {
		matches := instructionPattern.FindStringSubmatch(line)
		n, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])
		instruction := model.Instruction{
			N:    n,
			From: from - 1,
			To:   to - 1,
		}
		return instruction
	})
}
