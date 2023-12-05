package model

import "2023/src/framework/geometry"

type Symbol struct {
	Value    rune
	Position geometry.Point
}

func ExtractSymbols(line string, y int) (symbols []Symbol) {
	symbolMatches := symbolPattern.FindAllStringIndex(line, -1)
	for _, symbolMatch := range symbolMatches {
		x := symbolMatch[0]
		symbols = append(symbols, Symbol{
			Value:    rune(line[x]),
			Position: geometry.Point{X: x, Y: y},
		})
	}
	return
}
