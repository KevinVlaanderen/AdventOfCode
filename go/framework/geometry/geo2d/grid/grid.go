package grid

import "aoc/framework/geometry/geo2d"

type Grid[T comparable] interface {
	Get(p geo2d.Point) (value T, ok bool)
	Set(p geo2d.Point, value T) bool
	Clear(key geo2d.Point) bool
	MinX() int
	MinY() int
	MaxX() int
	MaxY() int
	Bounds() (minX int, minY int, maxX int, maxY int)
	InBounds(point geo2d.Point) bool
	Points() <-chan geo2d.Point
	PointsSet() <-chan geo2d.Point
	DrawPointGrid(mapping map[T]rune, fallback rune)
	DrawPointGridBy(mapping func(value T, isSet bool, x int, y int) rune)
}
