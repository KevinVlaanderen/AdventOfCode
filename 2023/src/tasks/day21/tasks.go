package day21

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

func Task1(data string, param int) (result framework.Result[int]) {
	tiles, start := parse(data)
	g := createGraph(tiles)

	found := mapset.NewSet[geometry.Point](start)
	search := traverse.BreadthFirst{}
	search.Walk(g, start, func(n graph.Node, d int) bool {
		if d == param || (d%2) == (param%2) {
			found.Add(n.(geometry.Point))
		} else if d > param {
			return true
		}
		return false
	})

	result.Value = found.Cardinality()

	drawMap(tiles, found)

	return
}

func drawMap(tiles [][]Tile, found mapset.Set[geometry.Point]) {
	for _, row := range tiles {
		for _, tile := range row {
			switch {
			case found.Contains(tile.Point):
				print("O")
			case tile.start:
				print("S")
			case tile.tileType == Rock:
				print("#")
			case tile.tileType == Garden:
				print(".")
			}
		}
		print("\n")
	}
	println()
}

type Tile struct {
	geometry.Point
	tileType TileType
	start    bool
}

func parse(data string) ([][]Tile, geometry.Point) {
	lines := framework.Lines(data)

	var start geometry.Point
	tiles := lo.Map(lines, func(line string, y int) []Tile {
		return lo.Map([]rune(line), func(char rune, x int) Tile {
			switch char {
			case '#':
				return Tile{geometry.Point{X: x, Y: y}, Rock, false}
			case 'S':
				start = geometry.Point{X: x, Y: y}
				return Tile{start, Garden, true}
			default:
				return Tile{geometry.Point{X: x, Y: y}, Garden, false}
			}
		})
	})

	return tiles, start
}

func createGraph(tiles [][]Tile) graph.Graph {
	g := simple.NewUndirectedGraph()

	for y := 0; y < len(tiles); y++ {
		for x := 0; x < len(tiles[0]); x++ {
			currentTile := tiles[y][x]
			g.AddNode(currentTile.Point)

			if x >= 1 {
				previousTile := tiles[y][x-1]
				if currentTile.tileType == Garden && previousTile.tileType == Garden {
					g.SetEdge(simple.Edge{F: currentTile.Point, T: previousTile.Point})
				}
			}
			if y >= 1 {
				previousTile := tiles[y-1][x]
				if currentTile.tileType == Garden && previousTile.tileType == Garden {
					g.SetEdge(simple.Edge{F: currentTile.Point, T: previousTile.Point})
				}
			}
		}
	}

	return g
}

type TileType int

const (
	Garden TileType = iota
	Rock
)
