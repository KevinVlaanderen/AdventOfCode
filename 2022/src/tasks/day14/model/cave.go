package model

import (
	"2022/src/framework/generators"
	"2022/src/framework/geometry"
	"2022/src/framework/math"
)

type Material int

const (
	ROCK Material = iota
	SAND
	AIR
)

type Cave struct {
	Area             geometry.Grid[Material]
	MinX, MaxX, MaxY int
	initialized      bool
}

func (c *Cave) AddRock(from geometry.Point, to geometry.Point) {
	if from.X == to.X {
		lowestY := math.MinInt(from.Y, to.Y)
		highestY := math.MaxInt(from.Y, to.Y)
		for _, y := range generators.Range(lowestY, math.AbsInt(highestY-lowestY)+1, 1) {
			c.Area[geometry.Point{X: from.X, Y: y}] = ROCK
		}
	} else {
		lowestX := math.MinInt(from.X, to.X)
		highestX := math.MaxInt(from.X, to.X)
		for _, x := range generators.Range(lowestX, math.AbsInt(highestX-lowestX)+1, 1) {
			c.Area[geometry.Point{X: x, Y: from.Y}] = ROCK
		}
	}
	minX := math.MinInt(from.X, to.X)
	if !c.initialized || minX < c.MinX {
		c.MinX = minX
	}
	maxX := math.MaxInt(from.X, to.X)
	if !c.initialized || maxX > c.MaxX {
		c.MaxX = maxX
	}
	maxY := math.MaxInt(from.Y, to.Y)
	if !c.initialized || maxY > c.MaxY {
		c.MaxY = maxY
	}
	c.initialized = true
}

func (c *Cave) DropSand(position geometry.Point) (comesToRest bool) {
	if filled, _ := c.filled(geometry.Point{X: position.X, Y: position.Y}); filled {
		return false
	}
	for {
		if position.Y >= c.MaxY {
			return false
		}
		if filled, _ := c.filled(geometry.Point{X: position.X, Y: position.Y + 1}); !filled {
			position.Y += 1
			continue
		} else if filled, _ := c.filled(geometry.Point{X: position.X - 1, Y: position.Y + 1}); !filled {
			position.X -= 1
			position.Y += 1
			continue
		} else if filled, _ := c.filled(geometry.Point{X: position.X + 1, Y: position.Y + 1}); !filled {
			position.X += 1
			position.Y += 1
			continue
		} else {
			break
		}
	}

	c.Area[position] = SAND
	return true
}

func (c *Cave) filled(position geometry.Point) (filled bool, material Material) {
	if material, filled := c.Area[position]; filled {
		return true, material
	} else {
		return false, AIR
	}
}
