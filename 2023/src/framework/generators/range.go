package generators

func Range(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

type RangeChan chan int

func (r RangeChan) Next() *int {
	c, ok := <-r
	if !ok {
		return nil
	}
	return &c
}

func RangeGen(start, count, step int) RangeChan {
	c := make(chan int)
	go func() {
		for {
			if count == 0 {
				close(c)
				return
			}
			c <- start
			start += step
			count--
		}
	}()
	return c
}
