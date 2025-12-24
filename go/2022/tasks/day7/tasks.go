//go:generate go run aoc/cmd/generate-tests

package day7

import (
	"aoc/2022/tasks/day7/model"
	"aoc/framework/tasks"
	"go/types"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	inputs := parse(data)

	fileSystem := NewFileSystem()
	shell := NewShell()

	for _, input := range inputs {
		switch input.InputType {
		case model.COMMAND:
			shell.Execute(input.Data)
		case model.CONTENT:
			parts := strings.Split(input.Data, " ")
			switch parts[0] {
			case "dir":
				fileSystem.CreateDir(path.Join(shell.currentPath, parts[1]))
			default:
				size, _ := strconv.Atoi(parts[0])
				fileSystem.CreateFile(path.Join(shell.currentPath, parts[1]), size)
			}
		}
	}

	_, allSizes := fileSystem.Size("/")

	for _, size := range allSizes {
		if size <= 100000 {
			result.Value += size
		}
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	inputs := parse(data)

	fileSystem := NewFileSystem()
	shell := NewShell()

	for _, input := range inputs {
		switch input.InputType {
		case model.COMMAND:
			shell.Execute(input.Data)
		case model.CONTENT:
			parts := strings.Split(input.Data, " ")
			switch parts[0] {
			case "dir":
				fileSystem.CreateDir(path.Join(shell.currentPath, parts[1]))
			default:
				size, _ := strconv.Atoi(parts[0])
				fileSystem.CreateFile(path.Join(shell.currentPath, parts[1]), size)
			}
		}
	}

	totalSize, allSizes := fileSystem.Size("/")

	remainingSize := 70000000 - totalSize
	requiredSize := 30000000 - remainingSize

	var eligible []int
	for _, size := range allSizes {
		if size >= requiredSize {
			eligible = append(eligible, size)
		}
	}

	sort.Sort(sort.IntSlice(eligible))
	result.Value = eligible[0]

	return
}

func parse(data string) []model.Input {
	return lo.Map(tasks.Lines(data), func(line string, index int) model.Input {
		if line[0] == '$' {
			return model.Input{
				InputType: model.COMMAND,
				Data:      line[2:],
			}
		} else {
			return model.Input{
				InputType: model.CONTENT,
				Data:      line,
			}
		}
	})
}
