package model

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/geometry/geo2d/grid"

	"github.com/samber/lo"
)

type Surface struct {
	start geo2d.Point
	Grid  grid.SparseGrid[*Pipe]
}

func NewSurface(data string) Surface {
	surface := Surface{
		Grid: grid.NewSparseGrid[*Pipe](),
	}

	current := geo2d.Point{}
	for _, r := range data {
		switch r {
		case '\n':
			current.X = 0
			current.Y++
			continue
		case 'S':
			surface.start = current
		case '.':
		default:
			surface.Grid.Add(current, NewPipeFromRune(r))
		}
		current.X++
	}

	surface.Grid.Add(surface.start, surface.calculateStartPipe())

	return surface
}

func (s Surface) calculateStartPipe() *Pipe {
	var top, right, bottom, left bool

	if neighbour, ok := s.Grid.Get(s.start.Neighbour(geo2d.North)); ok && neighbour.Bottom() {
		top = true
	}
	if neighbour, ok := s.Grid.Get(s.start.Neighbour(geo2d.East)); ok && neighbour.Left() {
		right = true
	}
	if neighbour, ok := s.Grid.Get(s.start.Neighbour(geo2d.South)); ok && neighbour.Top() {
		bottom = true
	}
	if neighbour, ok := s.Grid.Get(s.start.Neighbour(geo2d.West)); ok && neighbour.Right() {
		left = true
	}

	return NewPipeFromDirections(top, right, bottom, left)
}

func (s Surface) pickStartingDirection() geo2d.Orientation {
	start, found := s.Grid.Get(s.start)
	if !found {
		panic("starting pipe not found")
	}
	switch {
	case start.Top():
		return geo2d.North
	case start.Right():
		return geo2d.East
	case start.Bottom():
		return geo2d.South
	case start.Left():
		return geo2d.West
	}
	panic("starting pipe invalid")
}

func (s Surface) FindLoop() []Segment {
	startPipe, _ := s.Grid.Get(s.start)

	loop := []Segment{{
		Pipe:  startPipe,
		Point: &s.start,
	}}
	comingFrom := s.pickStartingDirection()

	currentPoint := s.start
	currentPipe := startPipe

	for {
		var neighbourPipe *Pipe

		dx, dy := currentPipe.EndpointDelta(comingFrom)
		neighbourPoint := geo2d.Point{X: currentPoint.X + dx, Y: currentPoint.Y + dy}

		if neighbourPoint == s.start {
			break
		}

		if neighbour, ok := s.Grid.Get(neighbourPoint); !ok {
			panic("neighbour not found")
		} else {
			neighbourPipe = neighbour
		}

		loop = append(loop, Segment{
			Pipe:  neighbourPipe,
			Point: &neighbourPoint,
		})
		currentPoint, currentPipe, comingFrom = neighbourPoint, neighbourPipe, geo2d.OppositeOrientation[currentPipe.OtherSide(comingFrom)]
	}
	return loop
}

func (s Surface) FindAllPointsInsideLoop(loop []Segment) []geo2d.Point {
	loopRotation := s.calculateLoopRotation(loop)

	found := make([]geo2d.Point, 0)

	for index := 0; index < len(loop); index++ {
		currentIndex, previousIndex := index, index-1
		if previousIndex == -1 {
			previousIndex = len(loop) - 1
		}

		current, previous := loop[currentIndex], loop[previousIndex]
		direction := previous.DirectionOf(current)

		switch {
		case direction == geo2d.North && current.Pipe.Type == TopBottom && loopRotation == geo2d.CW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
		case direction == geo2d.North && current.Pipe.Type == TopBottom && loopRotation == geo2d.CCW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		case direction == geo2d.North && current.Pipe.Type == BottomLeft && loopRotation == geo2d.CW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
		case direction == geo2d.North && current.Pipe.Type == BottomRight && loopRotation == geo2d.CCW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)

		case direction == geo2d.East && current.Pipe.Type == LeftRight && loopRotation == geo2d.CW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
		case direction == geo2d.East && current.Pipe.Type == LeftRight && loopRotation == geo2d.CCW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
		case direction == geo2d.East && current.Pipe.Type == TopLeft && loopRotation == geo2d.CW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
		case direction == geo2d.East && current.Pipe.Type == BottomLeft && loopRotation == geo2d.CCW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)

		case direction == geo2d.South && current.Pipe.Type == TopBottom && loopRotation == geo2d.CW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		case direction == geo2d.South && current.Pipe.Type == TopBottom && loopRotation == geo2d.CCW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
		case direction == geo2d.South && current.Pipe.Type == TopLeft && loopRotation == geo2d.CCW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
		case direction == geo2d.South && current.Pipe.Type == TopRight && loopRotation == geo2d.CW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)

		case direction == geo2d.West && current.Pipe.Type == LeftRight && loopRotation == geo2d.CW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
		case direction == geo2d.West && current.Pipe.Type == LeftRight && loopRotation == geo2d.CCW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
		case direction == geo2d.West && current.Pipe.Type == BottomRight && loopRotation == geo2d.CW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		case direction == geo2d.West && current.Pipe.Type == TopRight && loopRotation == geo2d.CCW:
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
			s.findPointsInsideLoop(loop, geo2d.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		}
	}
	return found
}

func (s Surface) findPointsInsideLoop(loop []Segment, point geo2d.Point, found *[]geo2d.Point) {
	if pipe, ok := s.Grid.Get(point); (ok && pipe.PartOfLoop) || lo.Contains(*found, point) {
		return
	}

	*found = append(*found, point)

	lo.ForEach(point.Neighbors(geo2d.All), func(neighbourPoint geo2d.Point, index int) {
		s.findPointsInsideLoop(loop, neighbourPoint, found)
	})
}

func (s Surface) calculateLoopRotation(loop []Segment) geo2d.Rotation {
	var nCW, nCCW int
	for index := 0; index < len(loop); index++ {
		currentIndex, nextIndex := index, index+1
		if currentIndex == len(loop)-1 {
			nextIndex = 0
		}

		current, next := loop[currentIndex], loop[nextIndex]
		comingFrom := geo2d.OppositeOrientation[current.DirectionOf(next)]

		switch next.Pipe.Rotation(comingFrom) {
		case geo2d.CW:
			nCW++
		case geo2d.CCW:
			nCCW++
		default:
		}
	}

	if nCW > nCCW {
		return geo2d.CW
	} else {
		return geo2d.CCW
	}
}
