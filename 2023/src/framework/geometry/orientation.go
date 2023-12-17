package geometry

type Orientation uint8

func (o Orientation) Name() string {
	return orientationNames[o]
}

const (
	North Orientation = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

var OppositeOrientation = map[Orientation]Orientation{
	North:     South,
	NorthEast: SouthWest,
	East:      West,
	SouthEast: NorthWest,
	South:     North,
	SouthWest: NorthEast,
	West:      East,
	NorthWest: SouthEast,
}

var orientationNames = []string{
	"north",
	"north-east",
	"east",
	"south-east",
	"south",
	"south-west",
	"west",
	"north-west",
}
