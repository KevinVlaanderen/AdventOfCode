package model

import (
	"2022/src/framework/generators"
	"fmt"
	"os"
)

type Grid map[Position]bool

func (g Grid) PrintVisited() {
	var minX, maxX, minY, maxY int
	for position := range g {
		if position.X < minX {
			minX = position.X
		} else if position.X > maxX {
			maxX = position.X
		}
		if position.Y < minY {
			minY = position.Y
		} else if position.Y > maxY {
			maxY = position.Y
		}
	}

	grid := make([][]bool, maxY-minY+1)
	for y := range grid {
		grid[y] = make([]bool, maxX-minX+1)
	}

	for pos := range g {
		grid[pos.Y-minY][pos.X-minX] = true
	}

	for _, gridY := range generators.Range(len(grid)-1, len(grid), -1) {
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