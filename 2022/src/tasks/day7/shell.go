package day7

import (
	"2022/src/tasks/day7/fs"
	"path"
	"strings"
)

type Shell struct {
	fileSystem  *fs.FileSystem
	currentPath string
}

func NewShell(fileSystem *fs.FileSystem) *Shell {
	return &Shell{
		fileSystem:  fileSystem,
		currentPath: "/",
	}
}

func (s *Shell) Execute(command string) {
	parts := strings.Split(command, " ")
	switch parts[0] {
	case "cd":
		s.currentPath = path.Join(s.currentPath, parts[1])
	case "ls":
	}
}

func (s *Shell) CurrentDir() *fs.Directory {
	return s.fileSystem.ResolveDir(s.currentPath)
}
