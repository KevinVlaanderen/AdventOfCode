package framework

func Range(start, count, step int) []int {
	result := make([]int, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, start)
		start += step
	}
	return result
}

func RangeGen(start, count, step int) <-chan int {
	c := make(chan int, count)
	go func() {
		defer close(c)
		for i := 0; i < count; i++ {
			c <- start
			start += step
		}
	}()
	return c
}
