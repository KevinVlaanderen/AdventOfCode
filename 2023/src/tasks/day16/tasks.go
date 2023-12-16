package day16

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/samber/lo"
)

func Task1(data string) (result framework.Result[int]) {
	cave := NewCave(data)

	currentPosition := geometry.Point{}
	currentDirection := geometry.East

	steps := map[Step]bool{}

	cave.FindPaths(currentPosition, currentDirection, &steps)

	found := hashset.New()
	lo.ForEach(lo.Keys(steps), func(item Step, index int) {
		found.Add(item.position)
	})
	cave.Draw(found)

	result.Value = found.Size()

	return
}

type Cave geometry.Grid[Tile]

func (c Cave) Draw(foundTiles *hashset.Set) {
	grid := geometry.Grid[Tile](c)
	grid.DrawPointGrid(func(value *Tile, x, y int) (rune, bool) {
		if foundTiles.Contains(geometry.Point{X: x, Y: y}) {
			return '#', true
		} else {
			return '.', true
		}
	}, nil)
	println()
}

func (c Cave) FindPaths(current geometry.Point, orientation geometry.Orientation, steps *map[Step]bool) {
	grid := geometry.Grid[Tile](c)
	xMin, xMax, yMin, yMax := grid.Boundaries()

	for {
		step := Step{current, orientation}

		if current.X < xMin || current.X > xMax || current.Y < yMin || current.Y > yMax {
			break
		} else if _, found := (*steps)[step]; found {
			break
		}

		(*steps)[step] = true

		if tile, found := grid.Get(current); !found {
			current = NextPosition(current, orientation)
		} else {
			switch tile.tileType {
			case MirrorLeft:
				switch orientation {
				case geometry.North:
					orientation = geometry.West
				case geometry.East:
					orientation = geometry.South
				case geometry.South:
					orientation = geometry.East
				case geometry.West:
					orientation = geometry.North
				}
				current = NextPosition(current, orientation)
			case MirrorRight:
				switch orientation {
				case geometry.North:
					orientation = geometry.East
				case geometry.East:
					orientation = geometry.North
				case geometry.South:
					orientation = geometry.West
				case geometry.West:
					orientation = geometry.South
				}
				current = NextPosition(current, orientation)
			case SplitterHorizontal:
				if orientation == geometry.East || orientation == geometry.West {
					current = NextPosition(current, orientation)
					continue
				}
				c.FindPaths(NextPosition(current, geometry.East), geometry.East, steps)
				c.FindPaths(NextPosition(current, geometry.West), geometry.West, steps)
				return
			case SplitterVertical:
				if orientation == geometry.North || orientation == geometry.South {
					current = NextPosition(current, orientation)
					continue
				}
				c.FindPaths(NextPosition(current, geometry.North), geometry.North, steps)
				c.FindPaths(NextPosition(current, geometry.South), geometry.South, steps)
				return
			}
		}
	}
	return
}

func NewCave(data string) Cave {
	tiles := geometry.NewGrid[Tile]()
	for y, line := range framework.Lines(data) {
		for x, char := range line {
			switch char {
			case '/':
				tiles.Add(geometry.Point{X: x, Y: y}, &Tile{MirrorRight})
			case '\\':
				tiles.Add(geometry.Point{X: x, Y: y}, &Tile{MirrorLeft})
			case '|':
				tiles.Add(geometry.Point{X: x, Y: y}, &Tile{SplitterVertical})
			case '-':
				tiles.Add(geometry.Point{X: x, Y: y}, &Tile{SplitterHorizontal})
			}
		}
	}

	return Cave(tiles)
}

type Step struct {
	position    geometry.Point
	orientation geometry.Orientation
}

type Tile struct {
	tileType TileType
}

type TileType uint8

const (
	MirrorLeft TileType = iota
	MirrorRight
	SplitterHorizontal
	SplitterVertical
)

func NextPosition(position geometry.Point, orientation geometry.Orientation) geometry.Point {
	switch orientation {
	case geometry.North:
		position.Y--
	case geometry.East:
		position.X++
	case geometry.South:
		position.Y++
	case geometry.West:
		position.X--
	}
	return position
}
