package model

import (
	"2022/src/framework/math"
	"fmt"
)

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

func ToDirection(value string) Direction {
	switch value {
	case "L":
		return LEFT
	case "R":
		return RIGHT
	case "U":
		return UP
	case "D":
		return DOWN
	default:
		panic(fmt.Sprintf("invalid direction: %v", value))
	}
}

type Rope struct {
	knots []knot
}

func NewRope(length int) *Rope {
	return &Rope{knots: make([]knot, length)}
}

func (r *Rope) Move(direction Direction) {
	r.head().move(direction)

	for index := range r.knots[1:] {
		previous := &r.knots[index]
		current := &r.knots[index+1]

		current.moveToward(previous)
	}
}

func (r *Rope) TailPosition() Position {
	return r.tail().Position
}

func (r *Rope) head() *knot {
	return &r.knots[0]
}

func (r *Rope) tail() *knot {
	return &r.knots[len(r.knots)-1]
}

type knot struct {
	Position
}

func (k *knot) distanceTo(other *knot) int {
	return math.MaxInt(math.AbsInt(k.X-other.X), math.AbsInt(k.Y-other.Y))
}

func (k *knot) alignedWith(other *knot) bool {
	return k.X == other.X || k.Y == other.Y
}

func (k *knot) move(direction Direction) {
	switch direction {
	case LEFT:
		k.X -= 1
	case RIGHT:
		k.X += 1
	case UP:
		k.Y += 1
	case DOWN:
		k.Y -= 1
	}
}

func (k *knot) moveToward(other *knot) {
	if k.distanceTo(other) <= 1 {
		return
	}

	if other.X > k.X {
		k.move(RIGHT)
	} else if other.X < k.X {
		k.move(LEFT)
	}
	if other.Y > k.Y {
		k.move(UP)
	} else if other.Y < k.Y {
		k.move(DOWN)
	}
}
