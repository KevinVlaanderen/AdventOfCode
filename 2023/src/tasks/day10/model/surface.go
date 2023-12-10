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
			surface.Grid.Add(current, NewPipe(r))
		}
		current.X++
	}

	surface.Grid.Add(surface.start, surface.calculateStartPipe())

	return surface
}

func (s Surface) calculateStartPipe() Pipe {
	var top, right, bottom, left bool
	var neighbourPoint geometry.Point

	neighbourPoint = geometry.Point{X: s.start.X, Y: s.start.Y - 1}
	if neighbour, ok := s.Grid.Get(neighbourPoint); ok && neighbour.Bottom() {
		top = true
	}
	neighbourPoint = geometry.Point{X: s.start.X + 1, Y: s.start.Y}
	if neighbour, ok := s.Grid.Get(neighbourPoint); ok && neighbour.Left() {
		right = true
	}
	neighbourPoint = geometry.Point{X: s.start.X, Y: s.start.Y + 1}
	if neighbour, ok := s.Grid.Get(neighbourPoint); ok && neighbour.Top() {
		bottom = true
	}
	neighbourPoint = geometry.Point{X: s.start.X - 1, Y: s.start.Y}
	if neighbour, ok := s.Grid.Get(neighbourPoint); ok && neighbour.Right() {
		left = true
	}

	switch {
	case top && bottom:
		return NewPipe('|')
	case left && right:
		return NewPipe('-')
	case top && right:
		return NewPipe('L')
	case bottom && right:
		return NewPipe('F')
	case top && left:
		return NewPipe('J')
	case bottom && left:
		return NewPipe('7')
	}

	panic("invalid starting point")
}

func (s Surface) FindLoop() []Segment {
	startPipe, _ := s.Grid.Get(s.start)
	loop := []Segment{{
		Pipe:  &startPipe,
		Point: &s.start,
	}}

	currentPoint, previousPoint := s.start, geometry.Point{X: -1, Y: -1}
	for {
		neighbourPoint := s.Grid.NeighboursBy(currentPoint, func(neighbourPoint geometry.Point) bool {
			if currentPoint.X != neighbourPoint.X && currentPoint.Y != neighbourPoint.Y {
				return false
			}

			current, ok := s.Grid.Get(currentPoint)
			if !ok {
				return false
			}
			neighbour, ok := s.Grid.Get(neighbourPoint)
			if !ok {
				return false
			}

			var side Direction
			switch {
			case neighbourPoint.Y < currentPoint.Y:
				side = Top
			case neighbourPoint.X > currentPoint.X:
				side = Right
			case neighbourPoint.Y > currentPoint.Y:
				side = Bottom
			case neighbourPoint.X < currentPoint.X:
				side = Left
			}

			return current.ConnectsTo(neighbour, side) && neighbourPoint != previousPoint
		})[0]
		if neighbourPoint == s.start {
			break
		} else {
			neighbourPipe, ok := s.Grid.Get(neighbourPoint)
			if !ok {
				panic("neighbour pipe not found")
			}

			loop = append(loop, Segment{
				Pipe:  &neighbourPipe,
				Point: &neighbourPoint,
			})
			currentPoint, previousPoint = neighbourPoint, currentPoint
		}
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
	if _, ok := s.Grid.Get(point); ok && lo.ContainsBy(loop, func(item Segment) bool {
		return *item.Point == point
	}) {
		return
	}

	if lo.Contains(*found, point) {
		return
	}

	*found = append(*found, point)

	lo.ForEach(s.Grid.Neighbours(point), func(neighbour geometry.Point, index int) {
		s.findPointsInsideLoop(loop, neighbour, found)
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
