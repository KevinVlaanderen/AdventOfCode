package model

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
	case comingFrom == Left && p.Type == BottomLeft:
		return CW
	case comingFrom == Left && p.Type == TopLeft:
		return CCW
	}
	return Straight
}

func (p Pipe) EndpointDelta(comingFrom Direction) (int, int) {
	switch {
	case comingFrom == Top && p.Type == TopLeft:
		return -1, 0
	case comingFrom == Top && p.Type == TopRight:
		return 1, 0
	case comingFrom == Top && p.Type == TopBottom:
		return 0, 1
	case comingFrom == Right && p.Type == TopRight:
		return 0, -1
	case comingFrom == Right && p.Type == BottomRight:
		return 0, 1
	case comingFrom == Right && p.Type == LeftRight:
		return -1, 0
	case comingFrom == Bottom && p.Type == BottomRight:
		return 1, 0
	case comingFrom == Bottom && p.Type == BottomLeft:
		return -1, 0
	case comingFrom == Bottom && p.Type == TopBottom:
		return 0, -1
	case comingFrom == Left && p.Type == BottomLeft:
		return 0, 1
	case comingFrom == Left && p.Type == TopLeft:
		return 0, -1
	case comingFrom == Left && p.Type == LeftRight:
		return 1, 0
	}
	panic("cannot determine delta")
}

func (p Pipe) OtherSide(comingFrom Direction) Direction {
	switch {
	case comingFrom == Top && p.Type == TopLeft:
		return Left
	case comingFrom == Top && p.Type == TopRight:
		return Right
	case comingFrom == Top && p.Type == TopBottom:
		return Bottom
	case comingFrom == Right && p.Type == TopRight:
		return Top
	case comingFrom == Right && p.Type == BottomRight:
		return Bottom
	case comingFrom == Right && p.Type == LeftRight:
		return Left
	case comingFrom == Bottom && p.Type == BottomRight:
		return Right
	case comingFrom == Bottom && p.Type == BottomLeft:
		return Left
	case comingFrom == Bottom && p.Type == TopBottom:
		return Top
	case comingFrom == Left && p.Type == BottomLeft:
		return Bottom
	case comingFrom == Left && p.Type == TopLeft:
		return Top
	case comingFrom == Left && p.Type == LeftRight:
		return Right
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
