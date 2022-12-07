package day7

import (
	"2022/src/framework"
	"2022/src/tasks/day7/fs"
	"fmt"
	"strconv"
	"strings"
)

func Task1(filePath string) (result framework.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	fileSystem := fs.NewFileSystem()
	shell := NewShell(fileSystem)

	inputs := Parse(data)

	for _, input := range inputs {
		switch input.InputType {
		case COMMAND:
			shell.Execute(input.data)
		case CONTENT:
			parts := strings.Split(input.data, " ")
			switch parts[0] {
			case "dir":
				shell.CurrentDir().AddDir(parts[1])
			default:
				size, _ := strconv.Atoi(parts[0])
				shell.CurrentDir().AddFile(parts[1], size)
			}
		}
	}

	cache := make(map[string]int)
	rootSize := fileSystem.ResolveDir("/").GetSize(cache)

	fmt.Print(rootSize)

	for _, size := range cache {
		if size <= 100000 {
			result.Value += size
		}
	}

	return
}

func Task2(filePath string) (result framework.TaskResult[int]) {
	_, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	return
}
