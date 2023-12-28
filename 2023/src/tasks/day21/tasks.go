package day21

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
	"math"
)

func Task1(data string, param int) (result framework.Result[int]) {
	tiles, start := parse(data)
	g := createGraph(tiles)

	found := mapset.NewSet[Tile](start)
	search := traverse.BreadthFirst{}
	search.Walk(g, start, func(n graph.Node, d int) bool {
		if (d % 2) == (param % 2) {
			found.Add(n.(Tile))
		} else if d > param {
			return true
		}
		return false
	})

	result.Value = found.Cardinality()

	drawMap(tiles, found.ToSlice())

	return
}

func Task2(data string, param int) (result framework.Result[int]) {
	tiles, start := parse(data)
	g := createGraph(tiles)

	nTotal := lo.SumBy(tiles, func(item []Tile) int {
		return len(item)
	})

	found := make(map[Tile]int)
	search := traverse.BreadthFirst{}
	search.Walk(g, start, func(n graph.Node, d int) bool {
		if (d % 2) == (param % 2) {
			if _, ok := found[n.(Tile)]; !ok {
				found[n.(Tile)] = d
			}
		} else if d > param {
			return true
		}
		return false
	})

	//irregularFound := make(map[Tile]int)
	//search = traverse.BreadthFirst{}
	//search.Walk(g, start, func(n graph.Node, d int) bool {
	//	if (d % 2) != (param % 2) {
	//		if _, ok := irregularFound[n.(Tile)]; !ok {
	//			irregularFound[n.(Tile)] = d
	//		}
	//	} else if d > param {
	//		return true
	//	}
	//	return false
	//})

	nFound := len(found)
	nFoundAlt := nTotal - nFound

	_, height := len(tiles[0]), len(tiles)
	nExtraBlocksInOneDirection := (param - (height / 2)) / height
	nStepsRemaining := (param - (height / 2)) % height

	fmt.Printf("Extra needed: %v, Steps remaining: %v\n", nExtraBlocksInOneDirection, nStepsRemaining)

	//nFullBlocks := (nExtraBlocksInOneDirection*2 + 1) * 2
	//nExtraFullBlocks := nFullBlocks - 1

	//nFullBlocksAlt := nExtraFullBlocks / 3 * 2

	totalExtraPerQuadrant := (nExtraBlocksInOneDirection / 2) * (1 + nExtraBlocksInOneDirection)
	var extraIrregularPerQuadrant int
	if totalExtraPerQuadrant%2 == 0 {
		extraIrregularPerQuadrant = int(math.Ceil(float64(totalExtraPerQuadrant/2)) + 1)
	} else {
		extraIrregularPerQuadrant = int(math.Ceil(float64(totalExtraPerQuadrant/2)) - 1)
	}
	extraRegularPerQuadrant := totalExtraPerQuadrant - extraIrregularPerQuadrant

	countFullRegular := extraRegularPerQuadrant*4 + 1
	countFullIrregular := extraIrregularPerQuadrant * 4

	result.Value = nFound*countFullRegular + nFoundAlt*countFullIrregular

	drawMap(tiles, lo.Keys(found))

	return
}

func drawMap(tiles [][]Tile, found []Tile) {
	for _, row := range tiles {
		for _, tile := range row {
			switch {
			case lo.Contains(found, tile):
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
	distance int
}

func parse(data string) ([][]Tile, Tile) {
	lines := framework.Lines(data)

	var start Tile
	tiles := lo.Map(lines, func(line string, y int) []Tile {
		return lo.Map([]rune(line), func(char rune, x int) Tile {
			switch char {
			case '#':
				return Tile{Point: geometry.Point{X: x, Y: y}, tileType: Rock}
			case 'S':
				start = Tile{Point: geometry.Point{X: x, Y: y}, tileType: Garden, start: true}
				return start
			default:
				return Tile{Point: geometry.Point{X: x, Y: y}, tileType: Garden}
			}
		})
	})

	return tiles, start
}

func createGraph(tiles [][]Tile) graph.Graph {
	g := simple.NewUndirectedGraph()

	width, height := len(tiles[0]), len(tiles)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			currentTile := tiles[y][x]
			g.AddNode(currentTile)

			if x >= 1 {
				previousTile := tiles[y][x-1]
				if currentTile.tileType == Garden && previousTile.tileType == Garden {
					g.SetEdge(simple.Edge{F: currentTile, T: previousTile})
				}
			}
			if y >= 1 {
				previousTile := tiles[y-1][x]
				if currentTile.tileType == Garden && previousTile.tileType == Garden {
					g.SetEdge(simple.Edge{F: currentTile, T: previousTile})
				}
			}
		}

		//g.SetEdge(simple.Edge{F: tiles[y][0], T: tiles[y][width-1]})
	}

	for x := 0; x < width; x++ {
		//g.SetEdge(simple.Edge{F: tiles[0][x], T: tiles[height-1][x]})
	}

	return g
}

type TileType int

const (
	Garden TileType = iota
	Rock
)
