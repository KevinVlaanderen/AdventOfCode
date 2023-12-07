package model

import (
	"regexp"
)

type Almanac struct {
	Targets  map[string]string
	Mappings map[string]Mapping
}

func (a Almanac) MapToNextStep(source string, value int) int {
	target, ok := a.Targets[source]
	if ok {
		newValue := a.Mappings[source].MapAcrossRanges(value)
		return a.MapToNextStep(target, newValue)
	} else {
		return value
	}
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
