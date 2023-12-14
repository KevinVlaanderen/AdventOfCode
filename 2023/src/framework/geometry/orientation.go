package geometry

type Orientation uint8

const (
	North Orientation = iota
	East
	South
	West
)

var OppositeOrientation = map[Orientation]Orientation{North: South, East: West, South: North, West: East}
