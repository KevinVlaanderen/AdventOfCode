package day23

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"2023/src/framework/geometry/grid"
	"github.com/samber/lo"
	"go/types"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	g := parse(data)

	startingPoints := []PointWithHeading{{geometry.Point{X: 1}, geometry.South}}
	startingPoints = append(startingPoints, findSlopes(&g)...)

	segments := lo.Associate(startingPoints, func(startingPoint PointWithHeading) (geometry.Point, Segment) {
		return calculateSegment(&g, startingPoint, true, func(terrain Terrain, heading geometry.Orientation) bool {
			return !(heading == geometry.North && terrain&SlopeSouth != 0) &&
				!(heading == geometry.East && terrain&SlopeWest != 0) &&
				!(heading == geometry.South && terrain&SlopeNorth != 0) &&
				!(heading == geometry.West && terrain&SlopeEast != 0)
		})
	})

	startingSegment := segments[geometry.Point{X: 1}]

	result.Value = lo.Max(calculatePathLengths(&startingSegment, segments)) - 1

	return
}

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	g := parse(data)

	startingPoints := []PointWithHeading{{geometry.Point{X: 1}, geometry.South}}
	startingPoints = append(startingPoints, findStartingPoints(&g)...)

	segments := lo.Associate(startingPoints, func(startingPoint PointWithHeading) (geometry.Point, Segment) {
		return calculateSegment(&g, startingPoint, false, func(terrain Terrain, heading geometry.Orientation) bool { return true })
	})

	startingSegment := segments[geometry.Point{X: 1}]

	seenEndPoints := make([]geometry.Point, 0)
	result.Value = lo.Max(calculatePathLengths2(&startingSegment, segments, seenEndPoints)) - 1

	return
}

func findSlopes(g *grid.SparseGrid[Terrain]) []PointWithHeading {
	slopes := make([]PointWithHeading, 0)

	lo.ForEach(g.Keys(), func(point geometry.Point, index int) {
		if terrain, found := g.Get(point); !found {
			panic("terrain not found")
		} else if terrain&Slope != 0 {
			switch {
			case terrain&SlopeNorth != 0:
				slopes = append(slopes, PointWithHeading{point, geometry.North})
			case terrain&SlopeEast != 0:
				slopes = append(slopes, PointWithHeading{point, geometry.East})
			case terrain&SlopeSouth != 0:
				slopes = append(slopes, PointWithHeading{point, geometry.South})
			case terrain&SlopeWest != 0:
				slopes = append(slopes, PointWithHeading{point, geometry.West})
			}
		}
	})
	return slopes
}

func findStartingPoints(g *grid.SparseGrid[Terrain]) []PointWithHeading {
	startingPoints := make([]PointWithHeading, 0)

	lo.ForEach(g.Keys(), func(point geometry.Point, index int) {
		neighbours := make([]PointWithHeading, 0, 4)

		for _, offset := range point.NeighbourOffsets(geometry.Orthogonal) {
			next := geometry.Point{X: point.X + offset.X, Y: point.Y + offset.Y}
			if _, found := g.Get(next); found {
				heading, _ := point.OrientationOf(next)
				neighbours = append(neighbours, PointWithHeading{next, heading})
			}
		}

		if len(neighbours) > 2 {
			startingPoints = append(startingPoints, neighbours...)
		}
	})
	return startingPoints
}

func calculatePathLengths(start *Segment, segments map[geometry.Point]Segment) []int {
	if len(start.Next) == 0 {
		return []int{start.Length}
	}

	return lo.FlatMap(start.Next, func(nextPoint geometry.Point, index int) []int {
		nextSegment := segments[nextPoint]
		return lo.Map(calculatePathLengths(&nextSegment, segments), func(length int, index int) int {
			return length + start.Length
		})
	})
}

func calculatePathLengths2(start *Segment, segments map[geometry.Point]Segment, seenEndpoints []geometry.Point) []int {
	if len(start.Next) == 0 {
		return []int{start.Length}
	}

	seenEndpoints = append(seenEndpoints, start.End)
	currentSeenEndpoint := seenEndpoints

	return lo.FlatMap(lo.Filter(lo.Map(start.Next, func(nextPoint geometry.Point, index int) Segment {
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

func calculateSegment(g *grid.SparseGrid[Terrain], startingPoint PointWithHeading, slippery bool, canMoveTo func(Terrain, geometry.Orientation) bool) (geometry.Point, Segment) {
	length := 1
	current := startingPoint.Point
	heading := startingPoint.Direction

	for {
		var validNextSteps []geometry.Point

		if length == 1 {
			validNextSteps = []geometry.Point{current.Neighbour(startingPoint.Direction)}
		} else {
			for _, offset := range current.NeighbourOffsets(geometry.Orthogonal) {
				next := geometry.Point{X: current.X + offset.X, Y: current.Y + offset.Y}
				nextHeading, _ := current.OrientationOf(next)
				if terrain, found := g.Get(next); found && nextHeading != geometry.OppositeOrientation[heading] && canMoveTo(terrain, nextHeading) {
					validNextSteps = append(validNextSteps, next)
				}
			}
		}

		switch len(validNextSteps) {
		case 0:
			return startingPoint.Point, Segment{startingPoint.Point, current, length, []geometry.Point{}}
		case 1:
			next := validNextSteps[0]
			if terrain, found := g.Get(next); found && slippery && terrain&Slope != 0 {
				return startingPoint.Point, Segment{startingPoint.Point, current, length, []geometry.Point{next}}
			}
			heading, _ = current.OrientationOf(next)
			current = next
			length++
		default:
			return startingPoint.Point, Segment{startingPoint.Point, current, length, validNextSteps}
		}
	}
}

func parse(data string) grid.SparseGrid[Terrain] {
	g := grid.NewSparseGrid[Terrain]()
	lines := framework.Lines(data)
	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '.':
				g.Add(geometry.Point{X: x, Y: y}, Path)
			case '^':
				g.Add(geometry.Point{X: x, Y: y}, Path|Slope|SlopeNorth)
			case '>':
				g.Add(geometry.Point{X: x, Y: y}, Path|Slope|SlopeEast)
			case 'v':
				g.Add(geometry.Point{X: x, Y: y}, Path|Slope|SlopeSouth)
			case '<':
				g.Add(geometry.Point{X: x, Y: y}, Path|Slope|SlopeWest)
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
	Point     geometry.Point
	Direction geometry.Orientation
}

type Segment struct {
	Start, End geometry.Point
	Length     int
	Next       []geometry.Point
}

type Junction struct {
	Point      geometry.Point
	Neighbours map[geometry.Orientation]geometry.Point
}
