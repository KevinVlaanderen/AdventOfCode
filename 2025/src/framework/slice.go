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
