package model

type Tile struct {
	TileType TileType
}

type TileType uint8

const (
	MirrorLeft TileType = iota
	MirrorRight
	SplitterHorizontal
	SplitterVertical
)
