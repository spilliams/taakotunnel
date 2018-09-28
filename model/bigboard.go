package model

import (
	"strconv"
	"strings"
	// log "github.com/sirupsen/logrus"
)

type BigBoard struct {
	TileBoard
	centerTile   *Tile
	widewideTile *Tile
}

func NewBigBoard() (*BigBoard, error) {
	b := BigBoard{}
	b.tileModel = "0123456"
	b.rotaModel = "0000000"
	e := b.makeTiles()

	var centerTile, widewideTile *Tile

	centerTile, e = NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelLeftTop},
		{In: TunnelRightTop, Out: TunnelRightBottom},
		{In: TunnelBottomLeft, Out: TunnelLeftBottom},
	}, "C")
	if e != nil {
		return nil, e
	}
	centerTile.RotateTopTo(DirectionSouth)
	b.centerTile = centerTile

	widewideTile, e = NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelLeftBottom},
		{In: TunnelRightTop, Out: TunnelLeftTop},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
	}, "W")
	if e != nil {
		return nil, e
	}
	widewideTile.RotateTopTo(DirectionSouth)
	b.widewideTile = widewideTile

	return &b, e
}

func (b *BigBoard) Size() int {
	return 3
}

func (b *BigBoard) EdgeTunnels() []Tunnel {
	return []Tunnel{
		{In: 5, Out: 7},   // V, 1 tile apart
		{In: 11, Out: 14}, // IX, 1 tile apart
		{In: 17, Out: 19}, // X, 1 tile apart
		{In: 6, Out: 9},   // VI, 2 tiles apart
		{In: 18, Out: 21}, // XI, 2 tiles apart
		{In: 22, Out: 24}, // XII, 2 tiles apart
		{In: 1, Out: 20},  // I, 3 tiles apart
		{In: 3, Out: 16},  // III, 3 tiles apart
		{In: 8, Out: 23},  // VII, 3 tiles apart
		{In: 10, Out: 15}, // VIII, 3 tiles apart
		{In: 4, Out: 13},  // IV, 4 tiles apart
		{In: 2, Out: 12},  // II, 4 tiles apart
	}
}

func (b *BigBoard) MakeTileSet() ([][]*Tile, error) {
	// get the tiles in the order of the tileModel
	tileIndexStrings := strings.Split(b.tileModel, "")
	tileIndexInts := make([]int, len(tileIndexStrings))
	for i, s := range tileIndexStrings {
		in, e := strconv.ParseInt(s, 10, 0)
		if e != nil {
			return nil, e
		}
		tileIndexInts[i] = int(in)
	}
	tilesInOrder := make([]*Tile, len(b.tileModel))
	for i, t := range tileIndexInts {
		tilesInOrder[i] = b.tiles[t]
	}

	// rotate tiles
	for i := 0; i < len(b.rotaModel); i++ {
		rotateTo64, _ := strconv.ParseInt(b.rotaModel[i:i+1], 10, 0)
		e := tilesInOrder[i].RotateTopTo(int(rotateTo64))
		if e != nil {
			return nil, e
		}
	}

	tileSet := [][]*Tile{
		{
			tilesInOrder[0],
			tilesInOrder[1],
			tilesInOrder[2],
		}, {
			tilesInOrder[3],
			b.centerTile,
			tilesInOrder[4],
		}, {
			tilesInOrder[5],
			tilesInOrder[6],
			b.widewideTile,
		},
	}

	return tileSet, nil
}

func (b *BigBoard) makeTiles() error {
	tile0, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelRightTop},
		{In: TunnelTopRight, Out: TunnelBottomRight},
		{In: TunnelBottomLeft, Out: TunnelLeftTop},
		{In: TunnelLeftBottom, Out: TunnelRightBottom},
	}, "a")
	if e != nil {
		return e
	}
	tile1, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomLeft},
		{In: TunnelTopRight, Out: TunnelRightTop},
		{In: TunnelRightBottom, Out: TunnelLeftTop},
		{In: TunnelBottomRight, Out: TunnelLeftBottom},
	}, "b")
	if e != nil {
		return e
	}
	tile2, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelLeftTop},
		{In: TunnelRightTop, Out: TunnelLeftBottom},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
	}, "c")
	if e != nil {
		return e
	}
	tile3, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelRightTop},
		{In: TunnelTopRight, Out: TunnelBottomRight},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
		{In: TunnelLeftBottom, Out: TunnelLeftTop},
	}, "d")
	if e != nil {
		return e
	}
	tile4, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelRightTop},
		{In: TunnelRightBottom, Out: TunnelLeftBottom},
		{In: TunnelBottomLeft, Out: TunnelLeftTop},
	}, "e")
	if e != nil {
		return e
	}
	tile6, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelTopRight},
		{In: TunnelRightTop, Out: TunnelLeftTop},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
		{In: TunnelBottomRight, Out: TunnelLeftBottom},
	}, "f")
	if e != nil {
		return e
	}
	tile7, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelLeftTop},
		{In: TunnelTopRight, Out: TunnelRightBottom},
		{In: TunnelRightTop, Out: TunnelBottomRight},
		{In: TunnelBottomLeft, Out: TunnelLeftBottom},
	}, "g")
	if e != nil {
		return e
	}
	b.tiles = []*Tile{
		tile0,
		tile1,
		tile2,
		tile3,
		tile4,
		tile6,
		tile7,
	}
	return nil
}

// returns a map from tunnel number to location{tile, tiletunnel}
func (b *BigBoard) TunnelMap() map[int]Location {
	return map[int]Location{
		1:  Location{Row: 0, Col: 0, Tunnel: TunnelTopLeft},
		2:  Location{Row: 0, Col: 0, Tunnel: TunnelTopRight},
		3:  Location{Row: 0, Col: 1, Tunnel: TunnelTopLeft},
		4:  Location{Row: 0, Col: 1, Tunnel: TunnelTopRight},
		5:  Location{Row: 0, Col: 2, Tunnel: TunnelTopLeft},
		6:  Location{Row: 0, Col: 2, Tunnel: TunnelTopRight},
		7:  Location{Row: 0, Col: 2, Tunnel: TunnelRightTop},
		8:  Location{Row: 0, Col: 2, Tunnel: TunnelRightBottom},
		9:  Location{Row: 1, Col: 2, Tunnel: TunnelRightTop},
		10: Location{Row: 1, Col: 2, Tunnel: TunnelRightBottom},
		11: Location{Row: 2, Col: 2, Tunnel: TunnelRightTop},
		12: Location{Row: 2, Col: 2, Tunnel: TunnelRightBottom},
		13: Location{Row: 2, Col: 2, Tunnel: TunnelBottomRight},
		14: Location{Row: 2, Col: 2, Tunnel: TunnelBottomLeft},
		15: Location{Row: 2, Col: 1, Tunnel: TunnelBottomRight},
		16: Location{Row: 2, Col: 1, Tunnel: TunnelBottomLeft},
		17: Location{Row: 2, Col: 0, Tunnel: TunnelBottomRight},
		18: Location{Row: 2, Col: 0, Tunnel: TunnelBottomLeft},
		19: Location{Row: 2, Col: 0, Tunnel: TunnelLeftBottom},
		20: Location{Row: 2, Col: 0, Tunnel: TunnelLeftTop},
		21: Location{Row: 1, Col: 0, Tunnel: TunnelLeftBottom},
		22: Location{Row: 1, Col: 0, Tunnel: TunnelLeftTop},
		23: Location{Row: 0, Col: 0, Tunnel: TunnelLeftBottom},
		24: Location{Row: 0, Col: 0, Tunnel: TunnelLeftTop},
	}
}
