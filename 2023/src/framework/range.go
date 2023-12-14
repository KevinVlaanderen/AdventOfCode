package framework

func Range(start, count, step int) (result []int) {
	for i := 0; i < count; i++ {
		result = append(result, start)
		start += step
	}
	return
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
