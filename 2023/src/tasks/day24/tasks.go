package day24

import (
	"2023/src/framework"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/spatial/r3"
	"math"
	"regexp"
	"strconv"
)

func Task1(data string, param lo.Tuple2[int, int]) (result framework.Result[int]) {
	hailstones := parse(data, true)
	minWindow, maxWindow := float64(param.A), float64(param.B)

	lo.ForEach(hailstones, func(current Hailstone, index int) {
		if index >= len(hailstones)-2 {
			return
		}

		for i := index + 1; i < len(hailstones); i++ {
			next := hailstones[i]

			if intersection, ok, s, t := current.intersectsAt(next); !ok || s < 0 || t < 0 {
				continue
			} else if intersection.X >= minWindow && intersection.X <= maxWindow && intersection.Y >= minWindow && intersection.Y <= maxWindow {
				result.Value++
			}
		}
	})

	return
}

var hailstonePattern = regexp.MustCompile(`(?m)^(-?\d+),\s+(-?\d+),\s+(-?\d+)\s+@\s+(-?\d+),\s+(-?\d+),\s+(-?\d+)$`)

func parse(data string, twoDimensions bool) []Hailstone {
	matches := hailstonePattern.FindAllStringSubmatch(data, -1)
	return lo.Map(matches, func(match []string, index int) Hailstone {
		px, _ := strconv.Atoi(match[1])
		py, _ := strconv.Atoi(match[2])
		pz, _ := strconv.Atoi(match[3])
		vx, _ := strconv.Atoi(match[4])
		vy, _ := strconv.Atoi(match[5])
		vz, _ := strconv.Atoi(match[6])

		if twoDimensions {
			pz, vz = 0, 0
		}

		return Hailstone{
			r3.Vec{X: float64(px), Y: float64(py), Z: float64(pz)},
			r3.Vec{X: float64(vx), Y: float64(vy), Z: float64(vz)}}
	})
}

type Hailstone struct {
	position, direction r3.Vec
}

func (h Hailstone) intersectsWith(other Hailstone) bool {
	return r3.Cos(h.direction, other.direction) == 0
}

func (h Hailstone) intersectsAt(other Hailstone) (r3.Vec, bool, float64, float64) {
	lineVec3 := r3.Sub(other.position, h.position)
	crossVec1and2 := r3.Cross(h.direction, other.direction)
	crossVec3and2 := r3.Cross(lineVec3, other.direction)
	crossVec3and1 := r3.Cross(lineVec3, h.direction)
	planarFactor := r3.Dot(lineVec3, crossVec1and2)

	sqrtCrossVec1and2 := r3.Norm2(crossVec1and2)

	if math.Abs(planarFactor) < 0.0001 && sqrtCrossVec1and2 > 0.0001 {
		s := r3.Dot(crossVec3and2, crossVec1and2) / sqrtCrossVec1and2
		t := r3.Dot(crossVec3and1, crossVec1and2) / sqrtCrossVec1and2
		return r3.Add(h.position, r3.Scale(s, h.direction)), true, s, t
	} else {
		return r3.Vec{}, false, 0, 0
	}
}
