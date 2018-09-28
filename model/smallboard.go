package model

import (
	"strconv"
	"strings"
)

type SmallBoard struct {
	TileBoard
}

func NewSmallBoard() (*SmallBoard, error) {
	b := SmallBoard{}
	b.tileModel = "0123"
	b.rotaModel = "0000"
	e := b.makeTiles()
	return &b, e
}

func (b *SmallBoard) Size() int {
	return 2
}

func (b *SmallBoard) EdgeTunnels() []Tunnel {
	return []Tunnel{
		{In: 1, Out: 12},
		{In: 2, Out: 3},
		{In: 4, Out: 6},
		{In: 5, Out: 7},
		{In: 8, Out: 13},
		{In: 9, Out: 10},
		{In: 11, Out: 14},
		{In: 15, Out: 16},
	}
}

func (b *SmallBoard) MakeTileSet() ([][]*Tile, error) {
	// log.Debugf("making tile set from %v, %v", b.tileModel, b.rotaModel)
	// rotate tiles
	for i := 0; i < len(b.rotaModel); i++ {
		rotateTo64, _ := strconv.ParseInt(b.rotaModel[i:i+1], 10, 0)
		e := b.tiles[i].RotateTopTo(int(rotateTo64))
		if e != nil {
			return nil, e
		}
	}

	tileIndexStrings := strings.Split(b.tileModel, "")
	tileIndexInts := make([]int, len(tileIndexStrings))
	for i, s := range tileIndexStrings {
		in, e := strconv.ParseInt(s, 10, 0)
		if e != nil {
			return nil, e
		}
		tileIndexInts[i] = int(in)
	}

	tileSet := [][]*Tile{
		{
			b.tiles[tileIndexInts[0]],
			b.tiles[tileIndexInts[1]],
		},
		{
			b.tiles[tileIndexInts[2]],
			b.tiles[tileIndexInts[3]],
		},
	}
	return tileSet, nil
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
	b.tiles = []*Tile{tile0, tile1, tile3, tile2}
	return nil
}

func (s *SmallBoard) TunnelMap() map[int]Location {
	return map[int]Location{
		1:  Location{Row: 0, Col: 0, Tunnel: TunnelTopLeft},
		2:  Location{Row: 0, Col: 0, Tunnel: TunnelTopRight},
		3:  Location{Row: 0, Col: 1, Tunnel: TunnelTopLeft},
		4:  Location{Row: 0, Col: 1, Tunnel: TunnelTopRight},
		5:  Location{Row: 0, Col: 1, Tunnel: TunnelRightTop},
		6:  Location{Row: 0, Col: 1, Tunnel: TunnelRightBottom},
		7:  Location{Row: 1, Col: 1, Tunnel: TunnelRightTop},
		8:  Location{Row: 1, Col: 1, Tunnel: TunnelRightBottom},
		9:  Location{Row: 1, Col: 1, Tunnel: TunnelBottomRight},
		10: Location{Row: 1, Col: 1, Tunnel: TunnelBottomLeft},
		11: Location{Row: 1, Col: 0, Tunnel: TunnelBottomRight},
		12: Location{Row: 1, Col: 0, Tunnel: TunnelBottomLeft},
		13: Location{Row: 1, Col: 0, Tunnel: TunnelLeftBottom},
		14: Location{Row: 1, Col: 0, Tunnel: TunnelLeftTop},
		15: Location{Row: 0, Col: 0, Tunnel: TunnelLeftBottom},
		16: Location{Row: 0, Col: 0, Tunnel: TunnelLeftTop},
	}
}
