package model

import (
	"aoc/framework/geometry/geo2d"
)

type Pipe struct {
	Type       PipeType
	PartOfLoop bool
}

func NewPipeFromRune(data rune) *Pipe {
	var pipeType PipeType
	var found bool

	if pipeType, found = pipeTypes[data]; !found {
		panic("unknown character")
	}

	return &Pipe{Type: pipeType}
}

func NewPipeFromDirections(top, right, bottom, left bool) *Pipe {
	switch {
	case top && bottom:
		return &Pipe{Type: TopBottom}
	case left && right:
		return &Pipe{Type: LeftRight}
	case top && right:
		return &Pipe{Type: TopRight}
	case bottom && right:
		return &Pipe{Type: BottomRight}
	case top && left:
		return &Pipe{Type: TopLeft}
	case bottom && left:
		return &Pipe{Type: BottomLeft}
	default:
		panic("cannot create pipe from directions")
	}
}

func (p Pipe) Top() bool {
	return p.Type == TopBottom || p.Type == TopRight || p.Type == TopLeft
}

func (p Pipe) Right() bool {
	return p.Type == LeftRight || p.Type == TopRight || p.Type == BottomRight
}

func (p Pipe) Bottom() bool {
	return p.Type == TopBottom || p.Type == BottomRight || p.Type == BottomLeft
}

func (p Pipe) Left() bool {
	return p.Type == LeftRight || p.Type == TopLeft || p.Type == BottomLeft
}

func (p Pipe) ConnectsTo(other Pipe, side geo2d.Orientation) bool {
	if (side == geo2d.North && !p.Top()) ||
		(side == geo2d.East && !p.Right()) ||
		(side == geo2d.South && !p.Bottom()) ||
		(side == geo2d.West && !p.Left()) {
		return false
	}
	oppositeSide := geo2d.OppositeOrientation[side]
	if (oppositeSide == geo2d.North && !other.Top()) ||
		(oppositeSide == geo2d.East && !other.Right()) ||
		(oppositeSide == geo2d.South && !other.Bottom()) ||
		(oppositeSide == geo2d.West && !other.Left()) {
		return false
	}
	return true
}

func (p Pipe) Rotation(comingFrom geo2d.Orientation) geo2d.Rotation {
	switch {
	case comingFrom == geo2d.North && p.Type == TopLeft:
		return geo2d.CW
	case comingFrom == geo2d.North && p.Type == TopRight:
		return geo2d.CCW
	case comingFrom == geo2d.East && p.Type == TopRight:
		return geo2d.CW
	case comingFrom == geo2d.East && p.Type == BottomRight:
		return geo2d.CCW
	case comingFrom == geo2d.South && p.Type == BottomRight:
		return geo2d.CW
	case comingFrom == geo2d.South && p.Type == BottomLeft:
		return geo2d.CCW
	case comingFrom == geo2d.West && p.Type == BottomLeft:
		return geo2d.CW
	case comingFrom == geo2d.West && p.Type == TopLeft:
		return geo2d.CCW
	}
	return geo2d.Straight
}

func (p Pipe) EndpointDelta(comingFrom geo2d.Orientation) (int, int) {
	switch {
	case comingFrom == geo2d.North && p.Type == TopLeft:
		return -1, 0
	case comingFrom == geo2d.North && p.Type == TopRight:
		return 1, 0
	case comingFrom == geo2d.North && p.Type == TopBottom:
		return 0, 1
	case comingFrom == geo2d.East && p.Type == TopRight:
		return 0, -1
	case comingFrom == geo2d.East && p.Type == BottomRight:
		return 0, 1
	case comingFrom == geo2d.East && p.Type == LeftRight:
		return -1, 0
	case comingFrom == geo2d.South && p.Type == BottomRight:
		return 1, 0
	case comingFrom == geo2d.South && p.Type == BottomLeft:
		return -1, 0
	case comingFrom == geo2d.South && p.Type == TopBottom:
		return 0, -1
	case comingFrom == geo2d.West && p.Type == BottomLeft:
		return 0, 1
	case comingFrom == geo2d.West && p.Type == TopLeft:
		return 0, -1
	case comingFrom == geo2d.West && p.Type == LeftRight:
		return 1, 0
	}
	panic("cannot determine delta")
}

func (p Pipe) OtherSide(comingFrom geo2d.Orientation) geo2d.Orientation {
	switch {
	case comingFrom == geo2d.North && p.Type == TopLeft:
		return geo2d.West
	case comingFrom == geo2d.North && p.Type == TopRight:
		return geo2d.East
	case comingFrom == geo2d.North && p.Type == TopBottom:
		return geo2d.South
	case comingFrom == geo2d.East && p.Type == TopRight:
		return geo2d.North
	case comingFrom == geo2d.East && p.Type == BottomRight:
		return geo2d.South
	case comingFrom == geo2d.East && p.Type == LeftRight:
		return geo2d.West
	case comingFrom == geo2d.South && p.Type == BottomRight:
		return geo2d.East
	case comingFrom == geo2d.South && p.Type == BottomLeft:
		return geo2d.West
	case comingFrom == geo2d.South && p.Type == TopBottom:
		return geo2d.North
	case comingFrom == geo2d.West && p.Type == BottomLeft:
		return geo2d.South
	case comingFrom == geo2d.West && p.Type == TopLeft:
		return geo2d.North
	case comingFrom == geo2d.West && p.Type == LeftRight:
		return geo2d.East
	}
	panic("cannot determine other side")
}

type PipeType uint8

const (
	TopBottom PipeType = iota
	LeftRight
	TopRight
	BottomRight
	TopLeft
	BottomLeft
)

var pipeTypes = map[rune]PipeType{
	'|': TopBottom,
	'-': LeftRight,
	'L': TopRight,
	'F': BottomRight,
	'J': TopLeft,
	'7': BottomLeft,
}
