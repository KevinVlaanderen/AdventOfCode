package geo3d

import math2 "math"

type Point3D struct {
	X, Y, Z float64
}

func Distance3D(p1, p2 Point3D) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z
	return math2.Sqrt(dx*dx + dy*dy + dz*dz)
}
