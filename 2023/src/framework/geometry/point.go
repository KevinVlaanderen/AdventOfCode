package geometry

import (
	math2 "2023/src/framework/math"
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"hash/fnv"
)

type Point struct {
	X, Y int
}

func (p Point) ID() int64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(p.Hash()))
	return int64(h.Sum64() >> 1)
}

func (p Point) Neighbors(mode NeighbourMode) (points []Point) {
	var orientations []Orientation
	switch mode {
	case All:
		orientations = allNeighbourOrientations
	case Orthogonal:
		orientations = orthogonalNeighbourOrientations
	}
	return lo.Map(orientations, func(orientation Orientation, index int) Point {
		return p.Neighbour(orientation)
	})
}

func (p Point) Neighbour(orientation Orientation) Point {
	if orientation < 0 || int(orientation) >= len(allNeighbourOffsets) {
		panic("unknown orientation")
	}
	offset := allNeighbourOffsets[orientation]
	return Point{p.X + offset.X, p.Y + offset.Y}
}

func (p Point) NeighbourOffsets(mode NeighbourMode) []Point {
	switch mode {
	case All:
		return allNeighbourOffsets
	case Orthogonal:
		return orthogonalNeighbourOffsets
	}
	panic("unknown mode")
}

func (p Point) OrientationOf(other Point) (Orientation, bool) {
	x := other.X - p.X
	y := other.Y - p.Y
	if x != 0 {
		x /= math2.AbsInt(x)
	}
	if y != 0 {
		y /= math2.AbsInt(y)
	}
	offset := Point{x, y}
	if index := slices.Index(allNeighbourOffsets, offset); index >= 0 {
		return allNeighbourOrientations[index], true
	}
	return 0, false
}

func (p Point) Hash() string {
	return fmt.Sprint(p)
}

var allNeighbourOrientations = []Orientation{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}
var orthogonalNeighbourOrientations = []Orientation{North, East, South, West}

var allNeighbourOffsets = []Point{
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

var orthogonalNeighbourOffsets = []Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type NeighbourMode uint8

const (
	All NeighbourMode = iota
	Orthogonal
)
