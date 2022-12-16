package model

import (
	"2022/src/framework/generators"
	"2022/src/framework/geometry"
)

type Grid [][]int

func (g Grid) DetermineVisibility() [][]bool {
	mask := geometry.CreateMask(g, false)

	increasingRange := generators.Range(0, len(g), 1)
	decreasingRange := generators.Range(len(g)-1, len(g), -1)

	// Top to bottom
	for _, x := range increasingRange {
		max := -1
		for _, y := range increasingRange {
			if g[x][y] > max {
				mask[x][y] = true
				max = g[x][y]
			}
		}
	}

	// Bottom to top
	for _, x := range increasingRange {
		max := -1
		for _, y := range decreasingRange {
			if g[x][y] > max {
				mask[x][y] = true
				max = g[x][y]
			}
		}
	}

	// Left to right
	for _, y := range increasingRange {
		max := -1
		for _, x := range increasingRange {
			if g[x][y] > max {
				mask[x][y] = true
				max = g[x][y]
			}
		}
	}

	// Right to left
	for _, y := range increasingRange {
		max := -1
		for _, x := range decreasingRange {
			if g[x][y] > max {
				mask[x][y] = true
				max = g[x][y]
			}
		}
	}

	return mask
}

func (g Grid) DetermineScore(x int, y int) int {
	leftRange := generators.Range(x-1, x, -1)
	rightRange := generators.Range(x+1, len(g)-x-1, 1)
	topRange := generators.Range(y-1, y, -1)
	bottomRange := generators.Range(y+1, len(g[x])-y-1, 1)

	var scoreLeft, scoreRight, scoreTop, scoreBottom int

	// To top
	for _, newY := range topRange {
		scoreTop++
		if g[x][newY] >= g[x][y] {
			break
		}
	}

	// To bottom
	for _, newY := range bottomRange {
		scoreBottom++
		if g[x][newY] >= g[x][y] {
			break
		}
	}

	// To left
	for _, newX := range leftRange {
		scoreLeft++
		if g[newX][y] >= g[x][y] {
			break
		}
	}

	// To right
	for _, newX := range rightRange {
		scoreRight++
		if g[newX][y] >= g[x][y] {
			break
		}
	}

	return scoreLeft * scoreRight * scoreTop * scoreBottom
}

