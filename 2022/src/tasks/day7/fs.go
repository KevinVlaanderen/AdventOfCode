package day7

import (
	"path"
	"strings"
)

type FileSystem struct {
	root *Directory
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		root: &Directory{},
	}
}

func (fs *FileSystem) CreateDir(p string) {
	dir := path.Dir(p)
	base := path.Base(p)
	directory := fs.resolveDir(dir)
	if directory.directories == nil {
		directory.directories = make(map[string]*Directory)
	} else if directory.directories[base] != nil {
		return
	}
	directory.directories[base] = &Directory{}
}

func (fs *FileSystem) CreateFile(p string, size int) {
	dir := path.Dir(p)
	base := path.Base(p)
	directory := fs.resolveDir(dir)
	if directory.files == nil {
		directory.files = make(map[string]*File)
	} else if directory.files[base] != nil {
		return
	}
	directory.files[base] = &File{size: size}
}

func (fs *FileSystem) Directories(p string) []string {
	directory := fs.resolveDir(p)
	var dirNames []string
	for name := range directory.directories {
		dirNames = append(dirNames, name)
	}
	return dirNames
}

func (fs *FileSystem) Files(p string) []string {
	directory := fs.resolveDir(p)
	var fileNames []string
	for name := range directory.files {
		fileNames = append(fileNames, name)
	}
	return fileNames
}

func (fs *FileSystem) Size(p string) (int, map[string]int) {
	cache := make(map[string]int)
	rootSize := fs.recursiveSize(p, cache)
	cache[p] = rootSize
	return rootSize, cache
}

func (fs *FileSystem) recursiveSize(dir string, cache map[string]int) int {
	directory := fs.resolveDir(dir)

	var filesSize, dirsSize int
	for _, file := range directory.files {
		filesSize += file.size
	}

	for name := range directory.directories {
		childPath := path.Join(dir, name)
		dirSize := fs.recursiveSize(childPath, cache)
		cache[childPath] = dirSize
		dirsSize += dirSize
	}

	return filesSize + dirsSize
}

func (fs *FileSystem) resolveDir(path string) *Directory {
	currentDir := fs.root

	for _, part := range strings.Split(path[1:], "/") {
		if part != "" {
			currentDir = currentDir.directories[part]
		}
	}

	return currentDir
}

type Directory struct {
	directories map[string]*Directory
	files       map[string]*File
}

type File struct {
	size int
}
