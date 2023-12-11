package model

import (
	"2023/src/framework/geometry"
	"github.com/samber/lo"
)

type Surface struct {
	start geometry.Point
	Grid  geometry.Grid[Pipe]
}

func NewSurface(data string) Surface {
	surface := Surface{
		Grid: geometry.NewGrid[Pipe](),
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

	if neighbour, ok := s.Grid.Get(s.start.Up()); ok && neighbour.Bottom() {
		top = true
	}
	if neighbour, ok := s.Grid.Get(s.start.Right()); ok && neighbour.Left() {
		right = true
	}
	if neighbour, ok := s.Grid.Get(s.start.Down()); ok && neighbour.Top() {
		bottom = true
	}
	if neighbour, ok := s.Grid.Get(s.start.Left()); ok && neighbour.Right() {
		left = true
	}

	return NewPipeFromDirections(top, right, bottom, left)
}

func (s Surface) pickStartingDirection() Direction {
	start, found := s.Grid.Get(s.start)
	if !found {
		panic("starting pipe not found")
	}
	switch {
	case start.Top():
		return Top
	case start.Right():
		return Right
	case start.Bottom():
		return Bottom
	case start.Left():
		return Left
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
		currentPoint, currentPipe, comingFrom = neighbourPoint, neighbourPipe, OppositeDirection[currentPipe.OtherSide(comingFrom)]
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
		case direction == Top && current.Pipe.Type == TopBottom && loopRotation == CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
		case direction == Top && current.Pipe.Type == TopBottom && loopRotation == CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		case direction == Top && current.Pipe.Type == BottomLeft && loopRotation == CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
		case direction == Top && current.Pipe.Type == BottomRight && loopRotation == CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)

		case direction == Right && current.Pipe.Type == LeftRight && loopRotation == CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
		case direction == Right && current.Pipe.Type == LeftRight && loopRotation == CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
		case direction == Right && current.Pipe.Type == TopLeft && loopRotation == CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
		case direction == Right && current.Pipe.Type == BottomLeft && loopRotation == CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)

		case direction == Bottom && current.Pipe.Type == TopBottom && loopRotation == CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		case direction == Bottom && current.Pipe.Type == TopBottom && loopRotation == CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
		case direction == Bottom && current.Pipe.Type == TopLeft && loopRotation == CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X + 1, Y: loop[index].Point.Y}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
		case direction == Bottom && current.Pipe.Type == TopRight && loopRotation == CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)

		case direction == Left && current.Pipe.Type == LeftRight && loopRotation == CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
		case direction == Left && current.Pipe.Type == LeftRight && loopRotation == CCW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y + 1}, &found)
		case direction == Left && current.Pipe.Type == BottomRight && loopRotation == CW:
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X, Y: loop[index].Point.Y - 1}, &found)
			s.findPointsInsideLoop(loop, geometry.Point{X: loop[index].Point.X - 1, Y: loop[index].Point.Y}, &found)
		case direction == Left && current.Pipe.Type == TopRight && loopRotation == CCW:
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

	lo.ForEach(s.Grid.Neighbours(point, false), func(neighbourPoint geometry.Point, index int) {
		s.findPointsInsideLoop(loop, neighbourPoint, found)
	})
}

func (s Surface) calculateLoopRotation(loop []Segment) Rotation {
	var nCW, nCCW int
	for index := 0; index < len(loop); index++ {
		currentIndex, nextIndex := index, index+1
		if currentIndex == len(loop)-1 {
			nextIndex = 0
		}

		current, next := loop[currentIndex], loop[nextIndex]
		comingFrom := OppositeDirection[current.DirectionOf(next)]

		switch next.Pipe.Rotation(comingFrom) {
		case CW:
			nCW++
		case CCW:
			nCCW++
		default:
		}
	}

	if nCW > nCCW {
		return CW
	} else {
		return CCW
	}
}
