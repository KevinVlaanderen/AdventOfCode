package fs

import (
	"fmt"
)

type Directory struct {
	directories map[string]*Directory
	files       map[string]*File
}

func (d *Directory) AddDir(name string) {
	if d.directories == nil {
		d.directories = map[string]*Directory{}
	}
	d.directories[name] = &Directory{}
}

func (d *Directory) AddFile(name string, size int) {
	if d.files == nil {
		d.files = map[string]*File{}
	}
	d.files[name] = &File{
		size: size,
	}
}

func (d *Directory) GetDir(name string) *Directory {
	return d.directories[name]
}

func (d *Directory) GetFile(name string) *File {
	return d.files[name]
}

func (d *Directory) GetSize(cache map[string]int) int {
	var dirsSize, filesSize int

	for _, file := range d.files {
		filesSize += file.GetSize()
	}

	for _, dir := range d.directories {
		dirSize := dir.GetSize(cache)
		cache[fmt.Sprint(*&dir)] = dirSize
		dirsSize += dirSize
	}

	return dirsSize + filesSize
}
