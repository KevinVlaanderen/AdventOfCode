package fs

type File struct {
	size int
}

func (f *File) GetSize() int {
	return f.size
}
