package model

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/geometry/geo2d/grid"
	"aoc/framework/tasks"

	"github.com/samber/lo"
)

type Cave struct {
	Grid grid.Grid[Tile]
}

func NewCave(data string) Cave {
	tiles := grid.NewSparseGrid[Tile]()
	for y, line := range tasks.Lines(data) {
		for x, char := range line {
			switch char {
			case '/':
				tiles.Set(geo2d.Point{X: x, Y: y}, Tile{MirrorRight})
			case '\\':
				tiles.Set(geo2d.Point{X: x, Y: y}, Tile{MirrorLeft})
			case '|':
				tiles.Set(geo2d.Point{X: x, Y: y}, Tile{SplitterVertical})
			case '-':
				tiles.Set(geo2d.Point{X: x, Y: y}, Tile{SplitterHorizontal})
			}
		}
	}

	return Cave{tiles}
}

func (c Cave) CountEnergized(position geo2d.Point, orientation geo2d.Orientation) int {
	steps := map[Step]bool{}
	c.followPath(c.Grid, position, orientation, &steps)

	found := lo.Associate(lo.Keys(steps), func(item Step) (geo2d.Point, bool) {
		return item.Position, true
	})

	return len(found)
}

func (c Cave) Boundaries() (int, int, int, int) {
	return c.Grid.Bounds()
}

func (c Cave) followPath(grid grid.Grid[Tile], current geo2d.Point, orientation geo2d.Orientation, steps *map[Step]bool) {
	minX, minY, maxX, maxY := grid.Bounds()

	for {
		if current.X < minX || current.X > maxX || current.Y < minY || current.Y > maxY {
			break
		}

		step := Step{current, orientation}
		if _, found := (*steps)[step]; found {
			break
		}

		(*steps)[step] = true

		if tile, found := grid.Get(current); !found {
			current = current.Neighbour(orientation)
		} else {
			switch tile.TileType {
			case MirrorLeft:
				switch orientation {
				case geo2d.North:
					orientation = geo2d.West
				case geo2d.East:
					orientation = geo2d.South
				case geo2d.South:
					orientation = geo2d.East
				case geo2d.West:
					orientation = geo2d.North
				default:
					panic("invalid orientation")
				}
				current = current.Neighbour(orientation)
			case MirrorRight:
				switch orientation {
				case geo2d.North:
					orientation = geo2d.East
				case geo2d.East:
					orientation = geo2d.North
				case geo2d.South:
					orientation = geo2d.West
				case geo2d.West:
					orientation = geo2d.South
				default:
					panic("invalid orientation")
				}
				current = current.Neighbour(orientation)
			case SplitterHorizontal:
				if orientation == geo2d.East || orientation == geo2d.West {
					current = current.Neighbour(orientation)
					continue
				}
				c.followPath(grid, current.Neighbour(geo2d.East), geo2d.East, steps)
				c.followPath(grid, current.Neighbour(geo2d.West), geo2d.West, steps)
				return
			case SplitterVertical:
				if orientation == geo2d.North || orientation == geo2d.South {
					current = current.Neighbour(orientation)
					continue
				}
				c.followPath(grid, current.Neighbour(geo2d.North), geo2d.North, steps)
				c.followPath(grid, current.Neighbour(geo2d.South), geo2d.South, steps)
				return
			}
		}
	}
	return
}
