package model

import (
	"2022/src/framework"
	"strconv"
	"strings"
)

type Material int

const (
	ROCK Material = iota
	SAND
	AIR
)

type Cave struct {
	Area             map[framework.Point]Material
	MinX, MaxX, MaxY int
	initialized      bool
}

func NewCave(rockData []string) *Cave {
	cave := Cave{Area: map[framework.Point]Material{}}

	for _, line := range rockData {
		lineParts := strings.Split(line, " -> ")

		var positions []framework.Point
		for _, linePart := range lineParts {
			rockParts := strings.Split(linePart, ",")
			x, _ := strconv.Atoi(rockParts[0])
			y, _ := strconv.Atoi(rockParts[1])
			positions = append(positions, framework.Point{X: x, Y: y})
		}

		for index := range positions {
			if index == len(positions)-1 {
				break
			}

			cave.AddRock(positions[index], positions[index+1])
		}
	}

	return &cave
}

func (c *Cave) AddRock(from framework.Point, to framework.Point) {
	if from.X == to.X {
		lowestY := framework.MinInt(from.Y, to.Y)
		highestY := framework.MaxInt(from.Y, to.Y)
		for _, y := range framework.Range(lowestY, framework.AbsInt(highestY-lowestY)+1, 1) {
			c.Area[framework.Point{X: from.X, Y: y}] = ROCK
		}
	} else {
		lowestX := framework.MinInt(from.X, to.X)
		highestX := framework.MaxInt(from.X, to.X)
		for _, x := range framework.Range(lowestX, framework.AbsInt(highestX-lowestX)+1, 1) {
			c.Area[framework.Point{X: x, Y: from.Y}] = ROCK
		}
	}
	minX := framework.MinInt(from.X, to.X)
	if !c.initialized || minX < c.MinX {
		c.MinX = minX
	}
	maxX := framework.MaxInt(from.X, to.X)
	if !c.initialized || maxX > c.MaxX {
		c.MaxX = maxX
	}
	maxY := framework.MaxInt(from.Y, to.Y)
	if !c.initialized || maxY > c.MaxY {
		c.MaxY = maxY
	}
	c.initialized = true
}

func (c *Cave) DropSand(position framework.Point) (comesToRest bool) {
	if filled, _ := c.filled(framework.Point{X: position.X, Y: position.Y}); filled {
		return false
	}
	for {
		if position.Y >= c.MaxY {
			return false
		}
		if filled, _ := c.filled(framework.Point{X: position.X, Y: position.Y + 1}); !filled {
			position.Y += 1
			continue
		} else if filled, _ := c.filled(framework.Point{X: position.X - 1, Y: position.Y + 1}); !filled {
			position.X -= 1
			position.Y += 1
			continue
		} else if filled, _ := c.filled(framework.Point{X: position.X + 1, Y: position.Y + 1}); !filled {
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

func (c *Cave) filled(position framework.Point) (filled bool, material Material) {
	if material, filled := c.Area[position]; filled {
		return true, material
	} else {
		return false, AIR
	}
}
