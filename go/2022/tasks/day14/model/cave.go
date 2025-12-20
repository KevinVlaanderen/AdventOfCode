package model

import (
	"aoc/framework"
	"aoc/framework/geometry/geo2d"
	"aoc/framework/geometry/geo2d/grid"
	"aoc/framework/math"
)

type Material int

const (
	AIR Material = iota
	ROCK
	SAND
)

type Cave struct {
	Area             grid.Grid[Material]
	MinX, MaxX, MaxY int
	initialized      bool
}

func (c *Cave) AddRock(from geo2d.Point, to geo2d.Point) {
	if from.X == to.X {
		lowestY := math.MinInt(from.Y, to.Y)
		highestY := math.MaxInt(from.Y, to.Y)
		for y := range framework.RangeGen(lowestY, math.AbsInt(highestY-lowestY)+1, 1) {
			if ok := c.Area.Set(geo2d.Point{X: from.X, Y: y}, ROCK); !ok {
				panic("out of bounds")
			}
		}
	} else {
		lowestX := math.MinInt(from.X, to.X)
		highestX := math.MaxInt(from.X, to.X)
		for x := range framework.RangeGen(lowestX, math.AbsInt(highestX-lowestX)+1, 1) {
			if ok := c.Area.Set(geo2d.Point{X: x, Y: from.Y}, ROCK); !ok {
				panic("out of bounds")
			}
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

func (c *Cave) DropSand(position geo2d.Point) (comesToRest bool) {
	if filled := c.filled(geo2d.Point{X: position.X, Y: position.Y}); filled {
		return false
	}
	for {
		if position.Y >= c.MaxY {
			return false
		}
		if filled := c.filled(geo2d.Point{X: position.X, Y: position.Y + 1}); !filled {
			position.Y += 1
			continue
		} else if filled = c.filled(geo2d.Point{X: position.X - 1, Y: position.Y + 1}); !filled {
			position.X -= 1
			position.Y += 1
			continue
		} else if filled = c.filled(geo2d.Point{X: position.X + 1, Y: position.Y + 1}); !filled {
			position.X += 1
			position.Y += 1
			continue
		} else {
			break
		}
	}

	if ok := c.Area.Set(position, SAND); !ok {
		panic("out of bounds")
	}

	return true
}

func (c *Cave) filled(position geo2d.Point) bool {
	material, found := c.Area.Get(position)
	return found && material != AIR
}
