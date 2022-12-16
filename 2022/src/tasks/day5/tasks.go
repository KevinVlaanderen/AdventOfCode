package day5

import (
	"2022/src/framework/generators"
	"2022/src/framework/tasks"
	"2022/src/tasks/day5/model"
	"2022/src/tasks/day5/mover"
	"regexp"
	"strconv"
)

type Mover interface {
	Move(n int, from int, to int)
}

func Task1(filePath string) (result tasks.TaskResult[string]) {
	storage := <- tasks.ReadStream(filePath, createStorageParser())
	instructions := tasks.ReadStream(filePath, createInstructionsParser())

	crateMover := &mover.CrateMover9000{Storage: storage}

	for instruction := range instructions {
		crateMover.Move(instruction.N, instruction.From, instruction.To)
	}

	for _, stack := range storage.Stacks {
		result.Value += stack.Top().Label
	}

	return
}

func Task2(filePath string) (result tasks.TaskResult[string]) {
	storage := <- tasks.ReadStream(filePath, createStorageParser())
	instructions := tasks.ReadStream(filePath, createInstructionsParser())

	crateMover := &mover.CrateMover9001{Storage: storage}

	for instruction := range instructions {
		crateMover.Move(instruction.N, instruction.From, instruction.To)
	}

	for _, stack := range storage.Stacks {
		result.Value += stack.Top().Label
	}

	return
}

var numberPattern = regexp.MustCompile(`\d+`)
var cratePattern = regexp.MustCompile(`\[(\w+)]`)

func createStorageParser() func(line string) (result *model.Storage, hasResult bool, err error) {
	var lines []string
	var done bool

	return func(line string) (result *model.Storage, hasResult bool, err error) {
		if done {
			return result, false, nil
		} else if line != "" {
			lines = append(lines, line)
			return result, false, nil
		}

		done = true

		numberLine := lines[len(lines)-1]
		stackData := lines[0 : len(lines)-1]
		numberMatches := numberPattern.FindAllString(numberLine, -1)
		numStacks := len(numberMatches)

		stacks := make([]model.Stack, numStacks)

		for _, lineIndex := range generators.Range(len(stackData)-1, len(stackData), -1) {
			for _, stackIndex := range generators.Range(0, numStacks, 1) {
				crateIndex := stackIndex * 4
				crateString := stackData[lineIndex][crateIndex : crateIndex+3]

				crateMatches := cratePattern.FindStringSubmatch(crateString)
				if len(crateMatches) < 2 {
					continue
				}

				stacks[stackIndex].Crates = append(stacks[stackIndex].Crates, model.Crate{Label: crateMatches[1]})
			}
		}

		storage := &model.Storage{Stacks: stacks}
		return storage, true, nil
	}
}

var instructionPattern = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func createInstructionsParser() func(line string) (result model.Instruction, hasResult bool, err error) {
	var start bool

	return func(line string) (result model.Instruction, hasResult bool, err error) {
		if line == "" && !start {
			start = true
			return result, false, nil
		} else if line == "" || !start {
			return result, false, nil
		}

		matches := instructionPattern.FindStringSubmatch(line)
		n, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])
		instruction := model.Instruction{
			N:    n,
			From: from - 1,
			To:   to - 1,
		}
		return instruction, true, nil
	}
}