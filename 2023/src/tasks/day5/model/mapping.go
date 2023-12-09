package model

import (
	"2023/src/framework/number"
	"golang.org/x/exp/slices"
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
	slices.SortStableFunc(mapping.ranges, func(a, b Range) int {
		return a.sourceStart - b.sourceStart
	})
	return source, target, mapping
}

func (c Mapping) MapToValue(value int) int {
	for _, r := range c.ranges {
		if newValue, ok := r.MapToNewValue(value); ok {
			return newValue
		}
	}
	return value
}

func (m Mapping) MapToRanges(seedRange SeedRange) []SeedRange {
	start, end := seedRange.Start, seedRange.End
	newSeedRanges := make([]SeedRange, 0)

	for _, targetRange := range m.ranges {
		if start <= targetRange.sourceEnd && end >= targetRange.sourceStart {
			if start < targetRange.sourceStart {
				newSeedRanges = append(newSeedRanges, SeedRange{
					Start: start,
					End:   targetRange.sourceStart - 1,
				})
			}

			offset := targetRange.destinationStart - targetRange.sourceStart
			var newStart, newEnd int
			if start >= targetRange.sourceStart {
				newStart = start
			} else {
				newStart = targetRange.sourceStart
			}
			if end <= targetRange.sourceEnd {
				newEnd = end
			} else {
				newEnd = targetRange.sourceEnd
			}

			newSeedRanges = append(newSeedRanges, SeedRange{
				Start: newStart + offset,
				End:   newEnd + offset,
			})
			start = newEnd + 1
		}
	}

	if start <= end {
		newSeedRanges = append(newSeedRanges, SeedRange{
			Start: start,
			End:   end,
		})
	}
	return newSeedRanges
}
