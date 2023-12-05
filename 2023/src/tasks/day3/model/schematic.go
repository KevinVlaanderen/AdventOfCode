package model

import (
	"2023/src/framework/geometry"
	"regexp"
)

var numberPattern = regexp.MustCompile(`\d+`)
var symbolPattern = regexp.MustCompile(`[^\d\\.]`)

type Schematic struct {
	lines   int
	Numbers []Number
	Symbols []Symbol
}

func NewSchematic() Schematic {
	return Schematic{
		Numbers: make([]Number, 0),
		Symbols: make([]Symbol, 0),
	}
}

func (s *Schematic) Add(line string) {
	s.Numbers = append(s.Numbers, ExtractNumbers(line, s.lines)...)
	s.Symbols = append(s.Symbols, ExtractSymbols(line, s.lines)...)
	s.lines++
}

func (s *Schematic) SymbolAt(location geometry.Point) (Symbol, bool) {
	for _, symbol := range s.Symbols {
		if symbol.Position == location {
			return symbol, true
		}
	}
	return Symbol{}, false
}
