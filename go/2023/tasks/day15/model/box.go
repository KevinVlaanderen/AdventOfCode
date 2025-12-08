package model

import "github.com/elliotchance/orderedmap/v2"

type Box struct {
	Lenses *orderedmap.OrderedMap[string, uint8]
}
