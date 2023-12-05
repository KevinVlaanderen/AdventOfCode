package model

import (
	"2023/src/framework/geometry"
	"2023/src/framework/number"
	"log"
	"strconv"
)

type Number struct {
	Value int
	Area  geometry.Area
}

func (n Number) Length() int {
	return number.Length(n.Value)
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
			Area:  geometry.Area{From: geometry.Point{X: start, Y: y}, To: geometry.Point{X: end - 1, Y: y}},
		})
	}
	return
}
