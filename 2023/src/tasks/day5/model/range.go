package model

type Range struct {
	sourceStart, sourceEnd           int
	destinationStart, destinationEnd int
	offset                           int
}

func NewRange(destination, source, length int) Range {
	return Range{
		sourceStart:      source,
		sourceEnd:        source + length - 1,
		destinationStart: destination,
		destinationEnd:   destination + length - 1,
		offset:           destination - source,
	}
}

func (r Range) InRange(value int) bool {
	return value >= r.sourceStart && value <= r.sourceEnd
}

func (r Range) MapToNewValue(value int) (int, bool) {
	if r.InRange(value) {
		return value + r.offset, true
	}
	return 0, false
}
