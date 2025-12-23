package framework

import "iter"

func RangeSlice(start, count, step int) []int {
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = start
		start += step
	}
	return result
}

func RangeChan(start, count, step int) <-chan int {
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

func RangeYield(start, count, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < count; i++ {
			if !yield(start) {
				return
			}
			start += step
		}
	}
}
