package generators

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
