package model

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"github.com/samber/lo"
)

type Cave geometry.Grid[Tile]

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

func (c Cave) CountEnergized(position geometry.Point, orientation geometry.Orientation) int {
	steps := map[Step]bool{}
	grid := geometry.Grid[Tile](c)
	c.followPath(&grid, position, orientation, &steps)

	found := lo.Associate(lo.Keys(steps), func(item Step) (geometry.Point, bool) {
		return item.Position, true
	})

	return len(found)
}

func (c Cave) Boundaries() (int, int, int, int) {
	grid := geometry.Grid[Tile](c)
	return grid.Boundaries()
}

func (c Cave) followPath(grid *geometry.Grid[Tile], current geometry.Point, orientation geometry.Orientation, steps *map[Step]bool) {
	xMin, xMax, yMin, yMax := grid.Boundaries()

	for {
		if current.X < xMin || current.X > xMax || current.Y < yMin || current.Y > yMax {
			break
		}

		step := Step{current, orientation}
		if _, found := (*steps)[step]; found {
			break
		}

		(*steps)[step] = true

		if tile, found := grid.Get(current); !found {
			current = nextPosition(current, orientation)
		} else {
			switch tile.TileType {
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
				current = nextPosition(current, orientation)
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
				current = nextPosition(current, orientation)
			case SplitterHorizontal:
				if orientation == geometry.East || orientation == geometry.West {
					current = nextPosition(current, orientation)
					continue
				}
				c.followPath(grid, nextPosition(current, geometry.East), geometry.East, steps)
				c.followPath(grid, nextPosition(current, geometry.West), geometry.West, steps)
				return
			case SplitterVertical:
				if orientation == geometry.North || orientation == geometry.South {
					current = nextPosition(current, orientation)
					continue
				}
				c.followPath(grid, nextPosition(current, geometry.North), geometry.North, steps)
				c.followPath(grid, nextPosition(current, geometry.South), geometry.South, steps)
				return
			}
		}
	}
	return
}

func nextPosition(position geometry.Point, orientation geometry.Orientation) geometry.Point {
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
