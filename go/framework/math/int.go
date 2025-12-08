package math

import (
	"errors"
	"math"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

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

func PowInt(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func Length(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func Length2(number int) (int, error) {
	if number == 0 {
		return 1, nil
	}

	floor := math.Floor(math.Log10(float64(number)))
	if floor >= math.MaxInt64 || floor < math.MinInt64 {
		return 0, errors.New("number out of range")
	}

	return int(floor) + 1, nil
}

var numberPattern = regexp.MustCompile(`-?\d+`)

func ExtractNumbers(line string) []int {
	return lo.Map(numberPattern.FindAllString(line, -1), func(item string, index int) int {
		if number, err := strconv.Atoi(item); err != nil {
			panic(err)
		} else {
			return number
		}
	})
}

func Factors(a int) (result []int) {
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			result = append(result, i)
		}
	}
	return
}

func Digits(n int) []int {
	var slc []int
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}

	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}

	return slc
}
