package framework

func RepeatSlice[T any](slice []T, count int) []T {
	size := len(slice)
	var bigSlice = make([]T, size*count)
	copy(bigSlice, slice)
	for j := 1; j < count; j++ {
		copy(bigSlice[size*j:], bigSlice[:size])
	}
	return bigSlice
}

func Partition[T any](slice []T, size int) (slices [][]T) {
	sliceSize := len(slice)
	var j int
	for i := 0; i < sliceSize; i += size {
		j += size
		if j > sliceSize {
			j = sliceSize
		}
		// do what do you want to with the sub-slice, here just printing the sub-slices
		slices = append(slices, slice[i:j])
	}
	return
}
