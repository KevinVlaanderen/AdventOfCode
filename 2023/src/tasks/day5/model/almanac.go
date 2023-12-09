package model

import (
	"github.com/samber/lo"
	"regexp"
)

type Almanac struct {
	Targets  map[string]string
	Mappings map[string]Mapping
}

var namePattern = regexp.MustCompile(`(.*)-to-(.*) map:`)

func NewAlmanac(blocks [][]string) Almanac {
	almanac := Almanac{
		Targets:  map[string]string{},
		Mappings: make(map[string]Mapping),
	}

	for _, block := range blocks {
		source, target, mapping := NewMapping(block)
		almanac.Targets[source] = target
		almanac.Mappings[source] = mapping
	}
	return almanac
}

func (a Almanac) MapToValue(value int) int {
	source := "seed"

	var mapping Mapping
	var ok bool

	for {
		if mapping, ok = a.Mappings[source]; !ok {
			break
		}

		value = mapping.MapToValue(value)

		if source, ok = a.Targets[source]; !ok {
			return value
		}
	}
	return value
}

func (a Almanac) MapToRanges(seedRange SeedRange) []SeedRange {
	source := "seed"

	newRanges := []SeedRange{seedRange}

	var mapping Mapping
	var ok bool

	for {
		if mapping, ok = a.Mappings[source]; !ok {
			break
		}

		newRanges = lo.FlatMap(newRanges, func(item SeedRange, index int) []SeedRange {
			return mapping.MapToRanges(item)
		})

		if source, ok = a.Targets[source]; !ok {
			break
		}
	}
	return newRanges
}
