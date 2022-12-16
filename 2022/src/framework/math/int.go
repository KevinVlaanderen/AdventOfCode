package math

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func AbsInt(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
