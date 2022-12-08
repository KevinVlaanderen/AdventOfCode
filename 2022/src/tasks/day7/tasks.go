package day7

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"path"
	"sort"
	"strconv"
	"strings"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	_, allSizes := run(data)

	for _, size := range allSizes {
		if size <= 100000 {
			result.Value += size
		}
	}

	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	totalSize, allSizes := run(data)
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

func run(data []string) (totalSize int, allSizes map[string]int) {
	fileSystem := NewFileSystem()
	shell := NewShell()

	inputs := Parse(data)

	for _, input := range inputs {
		switch input.InputType {
		case COMMAND:
			shell.Execute(input.data)
		case CONTENT:
			parts := strings.Split(input.data, " ")
			switch parts[0] {
			case "dir":
				fileSystem.CreateDir(path.Join(shell.currentPath, parts[1]))
			default:
				size, _ := strconv.Atoi(parts[0])
				fileSystem.CreateFile(path.Join(shell.currentPath, parts[1]), size)
			}
		}
	}

	return fileSystem.Size("/")
}
