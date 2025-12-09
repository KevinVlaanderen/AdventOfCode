package model

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/math"
	"log"
	"strconv"
)

type Number struct {
	Value int
	Area  geo2d.Area
}

func (n Number) Length() int {
	return math.Length(n.Value)
}

func ExtractNumbers(line string, y int) (numbers []Number) {
	numberMatches := numberPattern.FindAllStringIndex(line, -1)
	for _, numberMatch := range numberMatches {
		start, end := numberMatch[0], numberMatch[1]
		value, err := strconv.Atoi(line[start:end])
		if err != nil {
			log.Panic(err)
		}
		numbers = append(numbers, Number{
			Value: value,
			Area:  geo2d.Area{From: geo2d.Point{X: start, Y: y}, To: geo2d.Point{X: end - 1, Y: y}},
		})
	}
	return
}
