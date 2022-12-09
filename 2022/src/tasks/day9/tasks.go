package day9

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"regexp"
	"strconv"
)

type Position struct {
	x, y int
}

var linePattern = regexp.MustCompile(`^([LRUD]) (\d+)$`)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	var head, tail Position

	visited := map[Position]bool{tail: true}

	for _, line := range data {
		matches := linePattern.FindStringSubmatch(line)
		direction := matches[1]
		number, _ := strconv.Atoi(matches[2])

		for range framework.Range(0, number, 1) {
			switch direction {
			case "L":
				head.x -= 1
			case "R":
				head.x += 1
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			}

			if framework.MaxInt(framework.AbsInt(head.x-tail.x), framework.AbsInt(head.y-tail.y)) <= 1 {
				continue
			}

			switch framework.AbsInt(head.x-tail.x)+framework.AbsInt(head.y-tail.y) == 2 {
			case true:
				tail.x += (head.x - tail.x) / 2
				tail.y += (head.y - tail.y) / 2
			case false:
				if framework.AbsInt(head.x-tail.x) > 1 {
					tail.x += (head.x - tail.x) / 2
					tail.y = head.y
				} else {
					tail.x = head.x
					tail.y += (head.y - tail.y) / 2
				}
			}

			visited[tail] = true
		}
	}

	result.Value = len(visited)

	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	_, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	return
}
