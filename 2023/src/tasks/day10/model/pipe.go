package model

type Pipe struct {
	Type PipeType
}

func NewPipe(data rune) Pipe {
	var pipeType PipeType
	var found bool

	if pipeType, found = pipeTypes[data]; !found {
		panic("unknown character")
	}

	return Pipe{Type: pipeType}
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

func (p Pipe) ConnectsTo(other Pipe, side Direction) bool {
	if (side == Top && !p.Top()) ||
		(side == Right && !p.Right()) ||
		(side == Bottom && !p.Bottom()) ||
		(side == Left && !p.Left()) {
		return false
	}
	oppositeSide := OppositeDirection[side]
	if (oppositeSide == Top && !other.Top()) ||
		(oppositeSide == Right && !other.Right()) ||
		(oppositeSide == Bottom && !other.Bottom()) ||
		(oppositeSide == Left && !other.Left()) {
		return false
	}
	return true
}

func (p Pipe) Rotation(comingFrom Direction) Rotation {
	switch {
	case comingFrom == Top && p.Type == TopLeft:
		return CW
	case comingFrom == Top && p.Type == TopRight:
		return CCW
	case comingFrom == Right && p.Type == TopRight:
		return CW
	case comingFrom == Right && p.Type == BottomRight:
		return CCW
	case comingFrom == Bottom && p.Type == BottomRight:
		return CW
	case comingFrom == Bottom && p.Type == BottomLeft:
		return CCW
	case comingFrom == Left && p.Type == BottomRight:
		return CW
	case comingFrom == Left && p.Type == TopLeft:
		return CCW
	}
	return Straight
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
