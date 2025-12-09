package model

import (
	"aoc/framework/geometry/geo2d"
)

type Symbol struct {
	Value    rune
	Position geo2d.Point
}

func ExtractSymbols(line string, y int) (symbols []Symbol) {
	symbolMatches := symbolPattern.FindAllStringIndex(line, -1)
	for _, symbolMatch := range symbolMatches {
		x := symbolMatch[0]
		symbols = append(symbols, Symbol{
			Value:    rune(line[x]),
			Position: geo2d.Point{X: x, Y: y},
		})
	}
	return
}
