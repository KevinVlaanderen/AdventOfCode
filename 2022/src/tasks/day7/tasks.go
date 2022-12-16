package day7

import (
	"2022/src/framework/tasks"
	"2022/src/tasks/day7/model"
	"path"
	"sort"
	"strconv"
	"strings"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	inputs := tasks.ReadStream(filePath, parser)

	fileSystem := NewFileSystem()
	shell := NewShell()

	for input := range inputs {
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

func Task2(filePath string) (result tasks.TaskResult[int]) {
	inputs := tasks.ReadStream(filePath, parser)

	fileSystem := NewFileSystem()
	shell := NewShell()

	for input := range inputs {
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

func parser(line string) (result model.Input, hasResult bool, err error) {
	if line == "" {
		return result, false, nil
	}

	if line[0] == '$' {
		return model.Input{
			InputType: model.COMMAND,
			Data:      line[2:],
		}, true, nil
	} else {
		return model.Input{
			InputType: model.CONTENT,
			Data:      line,
		}, true, nil
	}
}
