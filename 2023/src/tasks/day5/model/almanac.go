package model

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds          []int
	Mappings       map[string]string
	ConversionMaps map[string][]ConversionMap
}

func (a Almanac) MapToNext(source string, value int) int {
	target, ok := a.Mappings[source]
	if ok {
		newValue := a.MapValue(source, value)
		return a.MapToNext(target, newValue)
	} else {
		return value
	}
}

func (a Almanac) MapValue(source string, value int) int {
	for _, conversionMap := range a.ConversionMaps[source] {
		if newValue, ok := conversionMap.Map(value); ok {
			return newValue
		}
	}
	return value
}

var seedsPattern = regexp.MustCompile(`seeds: ([\d\s]+)`)
var namePattern = regexp.MustCompile(`(.*)-to-(.*) map:`)

func NewAlmanac(blocks [][]string) Almanac {
	almanac := Almanac{
		Mappings:       map[string]string{},
		ConversionMaps: make(map[string][]ConversionMap),
	}

	seedsMatch := seedsPattern.FindStringSubmatch(blocks[0][0])
	almanac.Seeds = stringsToInts(strings.Split(seedsMatch[1], " "))

	for _, block := range blocks[1:] {
		var currentSource string
		conversionMaps := make([]ConversionMap, 0)

		for _, line := range block {
			if currentSource == "" {
				nameMatch := namePattern.FindStringSubmatch(line)
				almanac.Mappings[nameMatch[1]] = nameMatch[2]
				currentSource = nameMatch[1]
			} else {
				numbers := stringsToInts(strings.Split(line, " "))
				conversionMaps = append(conversionMaps, ConversionMap{
					destination: numbers[0],
					source:      numbers[1],
					length:      numbers[2],
				})
			}
		}
		almanac.ConversionMaps[currentSource] = conversionMaps
	}
	return almanac
}

func stringsToInts(s []string) []int {
	return lo.Map(s, func(item string, index int) int {
		if number, err := strconv.Atoi(item); err != nil {
			panic(err)
		} else {
			return number
		}
	})
}
