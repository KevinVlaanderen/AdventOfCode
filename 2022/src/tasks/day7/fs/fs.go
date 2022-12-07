package fs

import "strings"

type FileSystem struct {
	root *Directory
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		root: &Directory{},
	}
}

func (fs *FileSystem) ResolveDir(path string) *Directory {
	currentDir := fs.root

	for _, part := range strings.Split(path[1:], "/") {
		if part != "" {
			currentDir = currentDir.GetDir(part)
		}
	}

	return currentDir
}
