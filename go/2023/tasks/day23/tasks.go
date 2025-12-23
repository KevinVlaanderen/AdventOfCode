//go:generate go run aoc/cmd/generate-tests

package day23

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/geometry/geo2d/grid"
	"aoc/framework/tasks"
	"go/types"

	"github.com/samber/lo"
)

// Task1 type:mock 	file:data	expected:94
// Task1 type:real 	file:day23 	expected:2034
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	g := parse(data)

	startingPoints := []PointWithHeading{{geo2d.Point{X: 1}, geo2d.South}}
	startingPoints = append(startingPoints, findSlopes(g)...)

	segments := lo.Associate(startingPoints, func(startingPoint PointWithHeading) (geo2d.Point, Segment) {
		return calculateSegment(g, startingPoint, true, func(terrain Terrain, heading geo2d.Orientation) bool {
			return !(heading == geo2d.North && terrain&SlopeSouth != 0) &&
				!(heading == geo2d.East && terrain&SlopeWest != 0) &&
				!(heading == geo2d.South && terrain&SlopeNorth != 0) &&
				!(heading == geo2d.West && terrain&SlopeEast != 0)
		})
	})

	startingSegment := segments[geo2d.Point{X: 1}]

	result.Value = lo.Max(calculatePathLengths(&startingSegment, segments)) - 1

	return
}

// Task2 type:mock 	file:data	expected:154
// Task2 type:real 	file:day23 	expected:6302
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	g := parse(data)

	startingPoints := []PointWithHeading{{geo2d.Point{X: 1}, geo2d.South}}
	startingPoints = append(startingPoints, findStartingPoints(g)...)

	segments := lo.Associate(startingPoints, func(startingPoint PointWithHeading) (geo2d.Point, Segment) {
		return calculateSegment(g, startingPoint, false, func(terrain Terrain, heading geo2d.Orientation) bool { return true })
	})

	startingSegment := segments[geo2d.Point{X: 1}]

	seenEndPoints := make([]geo2d.Point, 0)
	result.Value = lo.Max(calculatePathLengths2(&startingSegment, segments, seenEndPoints)) - 1

	return
}

func findSlopes(g grid.Grid[Terrain]) []PointWithHeading {
	slopes := make([]PointWithHeading, 0)

	for point := range g.PointsSet() {
		if terrain, found := g.Get(point); !found {
			panic("terrain not found")
		} else if terrain&Slope != 0 {
			switch {
			case terrain&SlopeNorth != 0:
				slopes = append(slopes, PointWithHeading{point, geo2d.North})
			case terrain&SlopeEast != 0:
				slopes = append(slopes, PointWithHeading{point, geo2d.East})
			case terrain&SlopeSouth != 0:
				slopes = append(slopes, PointWithHeading{point, geo2d.South})
			case terrain&SlopeWest != 0:
				slopes = append(slopes, PointWithHeading{point, geo2d.West})
			}
		}
	}
	return slopes
}

func findStartingPoints(g grid.Grid[Terrain]) []PointWithHeading {
	startingPoints := make([]PointWithHeading, 0)

	for point := range g.PointsSet() {
		neighbours := make([]PointWithHeading, 0, 4)

		for _, offset := range point.NeighbourOffsets(geo2d.Orthogonal) {
			next := geo2d.Point{X: point.X + offset.X, Y: point.Y + offset.Y}
			if _, found := g.Get(next); found {
				heading, _ := point.OrientationOf(next)
				neighbours = append(neighbours, PointWithHeading{next, heading})
			}
		}

		if len(neighbours) > 2 {
			startingPoints = append(startingPoints, neighbours...)
		}
	}
	return startingPoints
}

func calculatePathLengths(start *Segment, segments map[geo2d.Point]Segment) []int {
	if len(start.Next) == 0 {
		return []int{start.Length}
	}

	return lo.FlatMap(start.Next, func(nextPoint geo2d.Point, index int) []int {
		nextSegment := segments[nextPoint]
		return lo.Map(calculatePathLengths(&nextSegment, segments), func(length int, index int) int {
			return length + start.Length
		})
	})
}

func calculatePathLengths2(start *Segment, segments map[geo2d.Point]Segment, seenEndpoints []geo2d.Point) []int {
	if len(start.Next) == 0 {
		return []int{start.Length}
	}

	seenEndpoints = append(seenEndpoints, start.End)
	currentSeenEndpoint := seenEndpoints

	return lo.FlatMap(lo.Filter(lo.Map(start.Next, func(nextPoint geo2d.Point, index int) Segment {
		return segments[nextPoint]
	}), func(nextSegment Segment, index int) bool {
		return !lo.Contains(seenEndpoints, nextSegment.End)
	}), func(nextSegment Segment, index int) []int {
		seenEndpoints = currentSeenEndpoint
		return lo.Map(calculatePathLengths2(&nextSegment, segments, seenEndpoints), func(length int, index int) int {
			return length + start.Length
		})
	})
}

func calculateSegment(g grid.Grid[Terrain], startingPoint PointWithHeading, slippery bool, canMoveTo func(Terrain, geo2d.Orientation) bool) (geo2d.Point, Segment) {
	length := 1
	current := startingPoint.Point
	heading := startingPoint.Direction

	for {
		var validNextSteps []geo2d.Point

		if length == 1 {
			validNextSteps = []geo2d.Point{current.Neighbour(startingPoint.Direction)}
		} else {
			for _, offset := range current.NeighbourOffsets(geo2d.Orthogonal) {
				next := geo2d.Point{X: current.X + offset.X, Y: current.Y + offset.Y}
				nextHeading, _ := current.OrientationOf(next)
				if terrain, found := g.Get(next); found && nextHeading != geo2d.OppositeOrientation[heading] && canMoveTo(terrain, nextHeading) {
					validNextSteps = append(validNextSteps, next)
				}
			}
		}

		switch len(validNextSteps) {
		case 0:
			return startingPoint.Point, Segment{startingPoint.Point, current, length, []geo2d.Point{}}
		case 1:
			next := validNextSteps[0]
			if terrain, found := g.Get(next); found && slippery && terrain&Slope != 0 {
				return startingPoint.Point, Segment{startingPoint.Point, current, length, []geo2d.Point{next}}
			}
			heading, _ = current.OrientationOf(next)
			current = next
			length++
		default:
			return startingPoint.Point, Segment{startingPoint.Point, current, length, validNextSteps}
		}
	}
}

func parse(data string) grid.Grid[Terrain] {
	g := grid.NewSparseGrid[Terrain]()
	lines := tasks.Lines(data)
	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '.':
				g.Set(geo2d.Point{X: x, Y: y}, Path)
			case '^':
				g.Set(geo2d.Point{X: x, Y: y}, Path|Slope|SlopeNorth)
			case '>':
				g.Set(geo2d.Point{X: x, Y: y}, Path|Slope|SlopeEast)
			case 'v':
				g.Set(geo2d.Point{X: x, Y: y}, Path|Slope|SlopeSouth)
			case '<':
				g.Set(geo2d.Point{X: x, Y: y}, Path|Slope|SlopeWest)
			}
		}
	}
	return g
}

type Terrain int

const (
	Path       Terrain = 1 << 0
	Slope              = 1 << 1
	SlopeNorth         = 1 << 2
	SlopeEast          = 1 << 3
	SlopeSouth         = 1 << 4
	SlopeWest          = 1 << 5
)

type PointWithHeading struct {
	Point     geo2d.Point
	Direction geo2d.Orientation
}

type Segment struct {
	Start, End geo2d.Point
	Length     int
	Next       []geo2d.Point
}

type Junction struct {
	Point      geo2d.Point
	Neighbours map[geo2d.Orientation]geo2d.Point
}
