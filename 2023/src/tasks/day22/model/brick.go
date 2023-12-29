package model

import (
	"2023/src/framework/geometry"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

type Brick struct {
	Endpoints   lo.Tuple2[geometry.Voxel, geometry.Voxel]
	Dimension   Dimension
	Supports    mapset.Set[*Brick]
	SupportedBy mapset.Set[*Brick]
}

func NewBrick(a, b geometry.Voxel) *Brick {
	var dimension Dimension
	switch {
	case a.X != b.X:
		dimension = X
	case a.Y != b.Y:
		dimension = Y
	default:
		dimension = Z
	}

	if a.X > b.X || a.Y > b.Y || a.Z > b.Z {
		a, b = b, a
	}
	return &Brick{
		Endpoints:   lo.Tuple2[geometry.Voxel, geometry.Voxel]{A: a, B: b},
		Dimension:   dimension,
		Supports:    mapset.NewSet[*Brick](),
		SupportedBy: mapset.NewSet[*Brick](),
	}
}

func (b *Brick) MoveToZ(z int) {
	deltaZ := b.Endpoints.B.Z - b.Endpoints.A.Z
	b.Endpoints.A.Z = z
	b.Endpoints.B.Z = z + deltaZ
}

func (b *Brick) FillXYPoints(buffer []geometry.Point) []geometry.Point {
	buffer = buffer[:0]
	switch b.Dimension {
	case X:
		y := b.Endpoints.A.Y
		for x := b.Endpoints.A.X; x <= b.Endpoints.B.X; x++ {
			buffer = append(buffer, geometry.Point{X: x, Y: y})
		}
	case Y:
		x := b.Endpoints.A.X
		for y := b.Endpoints.A.Y; y <= b.Endpoints.B.Y; y++ {
			buffer = append(buffer, geometry.Point{X: x, Y: y})
		}
	case Z:
		buffer = append(buffer, geometry.Point{X: b.Endpoints.A.X, Y: b.Endpoints.A.Y})
	}

	return buffer
}

type ByHeight []*Brick

func (b ByHeight) Len() int           { return len(b) }
func (b ByHeight) Less(i, j int) bool { return b[i].Endpoints.A.Z < b[j].Endpoints.A.Z }
func (b ByHeight) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type Dimension int

const (
	X Dimension = iota
	Y
	Z
)
