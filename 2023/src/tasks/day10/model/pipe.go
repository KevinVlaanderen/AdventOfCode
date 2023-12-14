package model

import "2023/src/framework/geometry"

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

func (p Pipe) ConnectsTo(other Pipe, side geometry.Orientation) bool {
	if (side == geometry.North && !p.Top()) ||
		(side == geometry.East && !p.Right()) ||
		(side == geometry.South && !p.Bottom()) ||
		(side == geometry.West && !p.Left()) {
		return false
	}
	oppositeSide := geometry.OppositeOrientation[side]
	if (oppositeSide == geometry.North && !other.Top()) ||
		(oppositeSide == geometry.East && !other.Right()) ||
		(oppositeSide == geometry.South && !other.Bottom()) ||
		(oppositeSide == geometry.West && !other.Left()) {
		return false
	}
	return true
}

func (p Pipe) Rotation(comingFrom geometry.Orientation) geometry.Rotation {
	switch {
	case comingFrom == geometry.North && p.Type == TopLeft:
		return geometry.CW
	case comingFrom == geometry.North && p.Type == TopRight:
		return geometry.CCW
	case comingFrom == geometry.East && p.Type == TopRight:
		return geometry.CW
	case comingFrom == geometry.East && p.Type == BottomRight:
		return geometry.CCW
	case comingFrom == geometry.South && p.Type == BottomRight:
		return geometry.CW
	case comingFrom == geometry.South && p.Type == BottomLeft:
		return geometry.CCW
	case comingFrom == geometry.West && p.Type == BottomLeft:
		return geometry.CW
	case comingFrom == geometry.West && p.Type == TopLeft:
		return geometry.CCW
	}
	return geometry.Straight
}

func (p Pipe) EndpointDelta(comingFrom geometry.Orientation) (int, int) {
	switch {
	case comingFrom == geometry.North && p.Type == TopLeft:
		return -1, 0
	case comingFrom == geometry.North && p.Type == TopRight:
		return 1, 0
	case comingFrom == geometry.North && p.Type == TopBottom:
		return 0, 1
	case comingFrom == geometry.East && p.Type == TopRight:
		return 0, -1
	case comingFrom == geometry.East && p.Type == BottomRight:
		return 0, 1
	case comingFrom == geometry.East && p.Type == LeftRight:
		return -1, 0
	case comingFrom == geometry.South && p.Type == BottomRight:
		return 1, 0
	case comingFrom == geometry.South && p.Type == BottomLeft:
		return -1, 0
	case comingFrom == geometry.South && p.Type == TopBottom:
		return 0, -1
	case comingFrom == geometry.West && p.Type == BottomLeft:
		return 0, 1
	case comingFrom == geometry.West && p.Type == TopLeft:
		return 0, -1
	case comingFrom == geometry.West && p.Type == LeftRight:
		return 1, 0
	}
	panic("cannot determine delta")
}

func (p Pipe) OtherSide(comingFrom geometry.Orientation) geometry.Orientation {
	switch {
	case comingFrom == geometry.North && p.Type == TopLeft:
		return geometry.West
	case comingFrom == geometry.North && p.Type == TopRight:
		return geometry.East
	case comingFrom == geometry.North && p.Type == TopBottom:
		return geometry.South
	case comingFrom == geometry.East && p.Type == TopRight:
		return geometry.North
	case comingFrom == geometry.East && p.Type == BottomRight:
		return geometry.South
	case comingFrom == geometry.East && p.Type == LeftRight:
		return geometry.West
	case comingFrom == geometry.South && p.Type == BottomRight:
		return geometry.East
	case comingFrom == geometry.South && p.Type == BottomLeft:
		return geometry.West
	case comingFrom == geometry.South && p.Type == TopBottom:
		return geometry.North
	case comingFrom == geometry.West && p.Type == BottomLeft:
		return geometry.South
	case comingFrom == geometry.West && p.Type == TopLeft:
		return geometry.North
	case comingFrom == geometry.West && p.Type == LeftRight:
		return geometry.East
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
