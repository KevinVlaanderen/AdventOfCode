package day7

import (
	"path"
	"strings"
)

type Shell struct {
	currentPath string
}

func NewShell() *Shell {
	return &Shell{
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

func (s *Shell) CurrentPath() string {
	return s.currentPath
}
