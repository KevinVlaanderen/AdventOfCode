package day9

import (
	"2022/src/framework"
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

type Position struct {
	x, y int
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
	return framework.MaxInt(framework.AbsInt(k.x-other.x), framework.AbsInt(k.y-other.y))
}

func (k *knot) alignedWith(other *knot) bool {
	return k.x == other.x || k.y == other.y
}

func (k *knot) move(direction Direction) {
	switch direction {
	case LEFT:
		k.x -= 1
	case RIGHT:
		k.x += 1
	case UP:
		k.y += 1
	case DOWN:
		k.y -= 1
	}
}

func (k *knot) moveToward(other *knot) {
	if k.distanceTo(other) <= 1 {
		return
	}

	if other.x > k.x {
		k.move(RIGHT)
	} else if other.x < k.x {
		k.move(LEFT)
	}
	if other.y > k.y {
		k.move(UP)
	} else if other.y < k.y {
		k.move(DOWN)
	}
}
