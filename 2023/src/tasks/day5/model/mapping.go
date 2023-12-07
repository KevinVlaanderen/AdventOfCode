package model

import (
	"2023/src/framework/number"
)

type Mapping struct {
	ranges []Range
}

func NewMapping(block []string) (source string, target string, mapping Mapping) {
	mapping = Mapping{
		ranges: make([]Range, 0),
	}

	for _, line := range block {
		if source == "" {
			nameMatch := namePattern.FindStringSubmatch(line)
			source = nameMatch[1]
			target = nameMatch[2]
		} else {
			numbers := number.ExtractNumbers(line)
			mapping.ranges = append(mapping.ranges, NewRange(numbers[0], numbers[1], numbers[2]))
		}
	}
	return source, target, mapping
}

func (c Mapping) MapAcrossRanges(value int) int {
	for _, r := range c.ranges {
		if newValue, ok := r.MapToNewValue(value); ok {
			return newValue
		}
	}
	return value
}
