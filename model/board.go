package model

type Board interface {
	TileModel() string
	SetTileModel(string)
	EndTileModel() string
	RotaModel() string
	SetRotaModel(string)
	EdgeTunnels() []Tunnel
	MakeTileSet() ([][]*Tile, error)
	TunnelMap() map[int]Location
	IsSolved() bool
	Size() int
}

type TileBoard struct {
	tileModel    string
	rotaModel    string
	endTileModel string
	tiles        []*Tile
}

func (b *TileBoard) TileModel() string {
	return b.tileModel
}

func (b *TileBoard) SetTileModel(newTileModel string) {
	b.tileModel = newTileModel
}

func (b *TileBoard) EndTileModel() string {
	return b.endTileModel
}

func (b *TileBoard) RotaModel() string {
	return b.rotaModel
}

func (b *TileBoard) SetRotaModel(newRotaModel string) {
	b.rotaModel = newRotaModel
}
