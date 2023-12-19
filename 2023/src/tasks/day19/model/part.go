package model

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
	"strings"
)

var partPattern = regexp.MustCompile(`(?m)^\{(.+)}$`)

func ParseParts(data string) []Part {
	partMatches := partPattern.FindAllStringSubmatch(data, -1)
	parts := make([]Part, len(partMatches))
	for partIndex, partMatch := range partMatches {
		ratingStrings := strings.Split(partMatch[1], ",")
		ratings := make(map[Category]int)
		for _, ratingString := range ratingStrings {
			category := Category(ratingString[0])
			value, err := strconv.Atoi(ratingString[2:])
			if err != nil {
				panic(err)
			}
			ratings[category] = value
		}

		parts[partIndex] = Part{ratings}
	}
	return parts
}

type Part struct {
	Ratings map[Category]int
}

func (p Part) Rating() int {
	return lo.Sum(lo.Values(p.Ratings))
}

type Category string
