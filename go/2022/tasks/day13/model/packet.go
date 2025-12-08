package model

type Packet = []any

func Compare(a any, b any) int {
	aSlice, aIsList := a.([]any)
	bSlice, bIsList := b.([]any)

	if !aIsList && !bIsList {
		return int(a.(float64) - b.(float64))
	} else {
		if aIsList && !bIsList {
			bSlice = []any{b}
		} else if !aIsList && bIsList {
			aSlice = []any{a}
		}

		if len(aSlice) == 0 && len(bSlice) == 0 {
			return 0
		} else if len(aSlice) == 0 {
			return -1
		} else if len(bSlice) == 0 {
			return 1
		}

		for index := range aSlice {
			if index >= len(bSlice) {
				return 1
			}

			itemA := aSlice[index]
			itemB := bSlice[index]

			switch result := Compare(itemA, itemB); {
			case result > 0:
				return 1
			case result < 0:
				return -1
			}
		}

		if len(bSlice) > len(aSlice) {
			return -1
		}
	}

	return 0
}