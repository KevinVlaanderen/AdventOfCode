package model

type SeedRange struct {
	Start, End int
}

func NewSeedRange(start, length int) SeedRange {
	return SeedRange{
		Start: start,
		End:   start + length - 1,
	}
}
