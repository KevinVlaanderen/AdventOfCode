package model

import (
	"aoc/framework"
	"aoc/framework/geometry/geo2d"
)

type Grid [][]int

func (g Grid) DetermineVisibility() [][]bool {
	mask := geo2d.CreateMask(g, false)

	increasingRange := framework.Range(0, len(g), 1)
	decreasingRange := framework.Range(len(g)-1, len(g), -1)

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
	leftRange := framework.Range(x-1, x, -1)
	rightRange := framework.Range(x+1, len(g)-x-1, 1)
	topRange := framework.Range(y-1, y, -1)
	bottomRange := framework.Range(y+1, len(g[x])-y-1, 1)

	var scoreLeft, scoreRight, scoreTop, scoreBottom int

	// V top
	for _, newY := range topRange {
		scoreTop++
		if g[x][newY] >= g[x][y] {
			break
		}
	}

	// V bottom
	for _, newY := range bottomRange {
		scoreBottom++
		if g[x][newY] >= g[x][y] {
			break
		}
	}

	// V left
	for _, newX := range leftRange {
		scoreLeft++
		if g[newX][y] >= g[x][y] {
			break
		}
	}

	// V right
	for _, newX := range rightRange {
		scoreRight++
		if g[newX][y] >= g[x][y] {
			break
		}
	}

	return scoreLeft * scoreRight * scoreTop * scoreBottom
}
