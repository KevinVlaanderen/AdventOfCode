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

	startingPoints := []StartingPoint{{geometry.Point{X: 1}, geometry.South}}

	lo.ForEach(g.Keys(), func(point geometry.Point, index int) {
		if terrain, found := g.Get(point); !found {
			panic("terrain not found")
		} else if terrain&Slope != 0 {
			switch {
			case terrain&SlopeNorth != 0:
				startingPoints = append(startingPoints, StartingPoint{point, geometry.North})
			case terrain&SlopeEast != 0:
				startingPoints = append(startingPoints, StartingPoint{point, geometry.East})
			case terrain&SlopeSouth != 0:
				startingPoints = append(startingPoints, StartingPoint{point, geometry.South})
			case terrain&SlopeWest != 0:
				startingPoints = append(startingPoints, StartingPoint{point, geometry.West})
			}
		}
	})

	segments := lo.Associate(startingPoints, func(startingPoint StartingPoint) (geometry.Point, Segment) {
		return calculateSegment(&g, startingPoint)
	})

	startingSegment := segments[geometry.Point{X: 1}]

	result.Value = lo.Max(calculateLengths(&startingSegment, segments)) - 1

	return
}

func calculateLengths(start *Segment, segments map[geometry.Point]Segment) []int {
	if len(start.Next) == 0 {
		return []int{start.Length}
	}

	return lo.FlatMap(start.Next, func(nextPoint geometry.Point, index int) []int {
		nextSegment := segments[nextPoint]
		return lo.Map(calculateLengths(&nextSegment, segments), func(length int, index int) int {
			return length + start.Length
		})
	})
}

func calculateSegment(g *grid.SparseGrid[Terrain], startingPoint StartingPoint) (geometry.Point, Segment) {
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
				if nextHeading == geometry.OppositeOrientation[heading] {
					continue
				}
				if terrain, found := g.Get(next); found {
					if nextHeading == geometry.North && terrain&SlopeSouth != 0 ||
						nextHeading == geometry.East && terrain&SlopeWest != 0 ||
						nextHeading == geometry.South && terrain&SlopeNorth != 0 ||
						nextHeading == geometry.West && terrain&SlopeEast != 0 {
						continue
					}
					validNextSteps = append(validNextSteps, next)
				}
			}
		}

		switch len(validNextSteps) {
		case 0:
			return startingPoint.Point, Segment{startingPoint.Point, current, length, []geometry.Point{}}
		case 1:
			next := validNextSteps[0]
			if terrain, found := g.Get(next); found && terrain&Slope != 0 {
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
				g.Add(geometry.Point{X: x, Y: y}, Slope|SlopeNorth)
			case '>':
				g.Add(geometry.Point{X: x, Y: y}, Slope|SlopeEast)
			case 'v':
				g.Add(geometry.Point{X: x, Y: y}, Slope|SlopeSouth)
			case '<':
				g.Add(geometry.Point{X: x, Y: y}, Slope|SlopeWest)
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

type StartingPoint struct {
	Point     geometry.Point
	Direction geometry.Orientation
}

type Segment struct {
	Start, End geometry.Point
	Length     int
	Next       []geometry.Point
}
