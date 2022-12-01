package tasks

import (
	"bufio"
	"strconv"
	"strings"
)

type Day1 struct {
}

func (d Day1) task1(data string) int {
	scanner := bufio.NewScanner(strings.NewReader(data))

	current, max := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			if current > max {
				max = current
			}
			current = 0
		} else {
			current += number
		}
	}

	return max
}
