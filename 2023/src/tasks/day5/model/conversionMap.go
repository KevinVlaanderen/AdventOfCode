package model

type ConversionMap struct {
	destination int
	source      int
	length      int
}

func (c ConversionMap) Map(value int) (int, bool) {
	if value >= c.source && value < c.source+c.length {
		return c.destination + (value - c.source), true
	}
	return 0, false
}
