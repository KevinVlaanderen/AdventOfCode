package geometry

func CreateMask[I any, O any](data [][]I, initial O) [][]O {
	sizeA := len(data)
	mask := make([][]O, sizeA)
	for a, aValues := range data {
		sizeB := len(data)
		aNew := make([]O, sizeB)
		for b := range aValues {
			aNew[b] = initial
		}
		mask[a] = aNew
	}
	return mask
}
