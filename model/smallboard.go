package model

type SmallBoard struct {
	TileBoard
}

func NewSmallBoard() (*SmallBoard, error) {
	b := SmallBoard{}
	b.tileModel = "0123"
	b.rotaModel = "0000"
	b.endTileModel = "3210"
	e := b.makeTiles()
	return &b, e
}

func (b *SmallBoard) makeTiles() error {
	tile0, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomLeft},
		{In: TunnelTopRight, Out: TunnelRightTop},
		{In: TunnelRightBottom, Out: TunnelBottomRight},
		{In: TunnelLeftBottom, Out: TunnelLeftTop},
	}, "a")
	if e != nil {
		return e
	}
	tile1, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelLeftTop},
		{In: TunnelTopRight, Out: TunnelBottomRight},
		{In: TunnelRightTop, Out: TunnelLeftBottom},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
	}, "b")
	if e != nil {
		return e
	}
	tile2, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomLeft},
		{In: TunnelTopRight, Out: TunnelRightTop},
		{In: TunnelRightBottom, Out: TunnelLeftBottom},
		{In: TunnelBottomRight, Out: TunnelLeftTop},
	}, "c")
	if e != nil {
		return e
	}
	tile3, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelTopRight},
		{In: TunnelRightTop, Out: TunnelLeftTop},
		{In: TunnelRightBottom, Out: TunnelLeftBottom},
		{In: TunnelBottomRight, Out: TunnelBottomLeft},
	}, "d")
	if e != nil {
		return e
	}
	b.tiles = []*Tile{tile3, tile2, tile1, tile0}
	return nil
}
