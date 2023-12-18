package model

import (
	"2023/src/framework/geometry"
	"2023/src/framework/geometry/grid"
	"github.com/samber/lo"
)

type Surface struct {
	start geometry.Point
	Grid  grid.SparseGrid[*Pipe]
}

func NewSurface(data string) Surface {
	surface := Surface{
		Grid: grid.NewSparseGrid[*Pipe](),
	}

	current := geometry.Point{}
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

	if neighbour, ok := s.Grid.Get(s.start.Neighbour(geometry.North)); ok && neighbour.Bottom() {
		top = true
	}
	if neighbour, ok := s.Grid.Get(s.start.Neighbour(geometry.East)); ok && neighbour.Left() {
		right = true
	}
	if neighbour, ok := s.Grid.Get(s.start.Neighbour(geometry.South)); ok && neighbour.Top() {
		bottom = true
	}
	if neighbour, ok := s.Grid.Get(s.start.Neighbour(geometry.West)); ok && neighbour.Right() {
		left = true
	}

	return NewPipeFromDirections(top, right, bottom, left)
}

func (s Surface) pickStartingDirection() geometry.Orientation {
	start, found := s.Grid.Get(s.start)
	if !found {
		panic("starting pipe not found")
	}
	switch {
	case start.Top():
		return geometry.North
	case start.Right():
		return geometry.East
	case start.Bottom():
		return geometry.South
	case start.Left():
		return geometry.West
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
		neighbourPoint := geometry.Point{X: currentPoint.X + dx, Y: currentPoint.Y + dy}

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
		currentPoint, currentPipe, comingFrom = neighbourPoint, neighbourPipe, geometry.OppositeOrientation[currentPipe.OtherSide(comingFrom)]
	}
	return loop
}

func (s Surface) FindAllPointsInsideLoop(loop []Segment) []geometry.Point {
	loopRotation := s.calculateLoopRotation(loop)

	found := make([]geometry.Point, 0)

	for index := 0; index < len(loop); index++ {
		currentIndex, previousIndex := index, index-1
		if previousIndex == -1 {
			previousIndex = len(loop) - 1
		}

		current, previous := loop[currentIndex], loop[previousIndex]
		direction := previous.DirectionOf(current)

		switch {
		case direction == geometry.North && current.Pipe.Type == TopBottom && loopRotation == geometry.CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
		case direction == geometry.North && current.Pipe.Type == TopBottom && loopRotation == geometry.CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		case direction == geometry.North && current.Pipe.Type == BottomLeft && loopRotation == geometry.CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
		case direction == geometry.North && current.Pipe.Type == BottomRight && loopRotation == geometry.CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)

		case direction == geometry.East && current.Pipe.Type == LeftRight && loopRotation == geometry.CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
		case direction == geometry.East && current.Pipe.Type == LeftRight && loopRotation == geometry.CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
		case direction == geometry.East && current.Pipe.Type == TopLeft && loopRotation == geometry.CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
		case direction == geometry.East && current.Pipe.Type == BottomLeft && loopRotation == geometry.CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)

		case direction == geometry.South && current.Pipe.Type == TopBottom && loopRotation == geometry.CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		case direction == geometry.South && current.Pipe.Type == TopBottom && loopRotation == geometry.CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
		case direction == geometry.South && current.Pipe.Type == TopLeft && loopRotation == geometry.CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
		case direction == geometry.South && current.Pipe.Type == TopRight && loopRotation == geometry.CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)

		case direction == geometry.West && current.Pipe.Type == LeftRight && loopRotation == geometry.CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
		case direction == geometry.West && current.Pipe.Type == LeftRight && loopRotation == geometry.CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
		case direction == geometry.West && current.Pipe.Type == BottomRight && loopRotation == geometry.CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		case direction == geometry.West && current.Pipe.Type == TopRight && loopRotation == geometry.CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		}
	}
	return found
}

func (s Surface) findPointsInsideLoop(loop []Segment, point geometry.Point, found *[]geometry.Point) {
	if pipe, ok := s.Grid.Get(point); (ok && pipe.PartOfLoop) || lo.Contains(*found, point) {
		return
	}

	*found = append(*found, point)

	lo.ForEach(point.Neighbors(geometry.All), func(neighbourPoint geometry.Point, index int) {
		s.findPointsInsideLoop(loop, neighbourPoint, found)
	})
}

func (s Surface) calculateLoopRotation(loop []Segment) geometry.Rotation {
	var nCW, nCCW int
	for index := 0; index < len(loop); index++ {
		currentIndex, nextIndex := index, index+1
		if currentIndex == len(loop)-1 {
			nextIndex = 0
		}

		current, next := loop[currentIndex], loop[nextIndex]
		comingFrom := geometry.OppositeOrientation[current.DirectionOf(next)]

		switch next.Pipe.Rotation(comingFrom) {
		case geometry.CW:
			nCW++
		case geometry.CCW:
			nCCW++
		default:
		}
	}

	if nCW > nCCW {
		return geometry.CW
	} else {
		return geometry.CCW
	}
}
