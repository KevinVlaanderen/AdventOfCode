package model

type Direction uint8

const (
	Top Direction = iota
	Right
	Bottom
	Left
)

var OppositeDirection = map[Direction]Direction{Top: Bottom, Right: Left, Bottom: Top, Left: Right}
