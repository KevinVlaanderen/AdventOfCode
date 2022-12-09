package day9

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	rope := NewRope(2)
	visited := map[Position]bool{rope.TailPosition(): true}

	for _, instruction := range ParseInstructions(data) {
		for range framework.Range(0, instruction.steps, 1) {
			rope.Move(instruction.direction)
			visited[rope.TailPosition()] = true
		}
	}

	result.Value = len(visited)

	if false {
		printVisited(visited)
	}

	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	rope := NewRope(10)
	visited := map[Position]bool{rope.TailPosition(): true}

	for _, instruction := range ParseInstructions(data) {
		for range framework.Range(0, instruction.steps, 1) {
			rope.Move(instruction.direction)
			visited[rope.TailPosition()] = true
		}
	}

	result.Value = len(visited)

	if false {
		printVisited(visited)
	}

	return
}

func printVisited(visited map[Position]bool) {
	var minX, maxX, minY, maxY int
	for position := range visited {
		if position.x < minX {
			minX = position.x
		} else if position.x > maxX {
			maxX = position.x
		}
		if position.y < minY {
			minY = position.y
		} else if position.y > maxY {
			maxY = position.y
		}
	}

	grid := make([][]bool, maxY-minY+1)
	for y := range grid {
		grid[y] = make([]bool, maxX-minX+1)
	}

	for pos := range visited {
		grid[pos.y-minY][pos.x-minX] = true
	}

	for _, gridY := range framework.Range(len(grid)-1, len(grid), -1) {
		var line string
		for _, gridX := range grid[gridY] {
			value := "."
			if gridX {
				value = "#"
			}
			line += value
		}
		_, _ = fmt.Fprintln(os.Stdout, line)
	}
}

type Instruction struct {
	direction Direction
	steps     int
}

var linePattern = regexp.MustCompile(`^([LRUD]) (\d+)$`)

func ParseInstructions(data []string) (instructions []Instruction) {
	for _, line := range data {
		matches := linePattern.FindStringSubmatch(line)
		direction := ToDirection(matches[1])
		number, _ := strconv.Atoi(matches[2])

		instructions = append(instructions, Instruction{
			direction: direction,
			steps:     number,
		})
	}
	return
}
