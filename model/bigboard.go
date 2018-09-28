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
	b.endTileModel = "6543210"
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
		1:  Location{row: 0, col: 0, tunnel: TunnelTopLeft},
		2:  Location{row: 0, col: 0, tunnel: TunnelTopRight},
		3:  Location{row: 0, col: 1, tunnel: TunnelTopLeft},
		4:  Location{row: 0, col: 1, tunnel: TunnelTopRight},
		5:  Location{row: 0, col: 2, tunnel: TunnelTopLeft},
		6:  Location{row: 0, col: 2, tunnel: TunnelTopRight},
		7:  Location{row: 0, col: 2, tunnel: TunnelRightTop},
		8:  Location{row: 0, col: 2, tunnel: TunnelRightBottom},
		9:  Location{row: 1, col: 2, tunnel: TunnelRightTop},
		10: Location{row: 1, col: 2, tunnel: TunnelRightBottom},
		11: Location{row: 2, col: 2, tunnel: TunnelRightTop},
		12: Location{row: 2, col: 2, tunnel: TunnelRightBottom},
		13: Location{row: 2, col: 2, tunnel: TunnelBottomRight},
		14: Location{row: 2, col: 2, tunnel: TunnelBottomLeft},
		15: Location{row: 2, col: 1, tunnel: TunnelBottomRight},
		16: Location{row: 2, col: 1, tunnel: TunnelBottomLeft},
		17: Location{row: 2, col: 0, tunnel: TunnelBottomRight},
		18: Location{row: 2, col: 0, tunnel: TunnelBottomLeft},
		19: Location{row: 2, col: 0, tunnel: TunnelLeftBottom},
		20: Location{row: 2, col: 0, tunnel: TunnelLeftTop},
		21: Location{row: 1, col: 0, tunnel: TunnelLeftBottom},
		22: Location{row: 1, col: 0, tunnel: TunnelLeftTop},
		23: Location{row: 0, col: 0, tunnel: TunnelLeftBottom},
		24: Location{row: 0, col: 0, tunnel: TunnelLeftTop},
	}
}

func (b *BigBoard) IsSolved() (bool, error) {
	tileSet, e := b.MakeTileSet()
	if e != nil {
		return false, e
	}
	boardMap := b.TunnelMap()

	for _, tunnel := range b.EdgeTunnels() {
		// log.Debugf("Checking that tunnel %v connects to %v", tunnel.In, tunnel.Out)
		loc := boardMap[tunnel.In]
		// log.Debugf("  starting at %v", loc)
		for !loc.end {
			loc, e = b.follow(loc, tileSet)
			if e != nil {
				return false, e
			}
			// log.Debugf("  followed to %v", loc)
		}
		if boardMap[tunnel.Out].row != loc.row || boardMap[tunnel.Out].col != loc.col || boardMap[tunnel.Out].tunnel != loc.tunnel {
			// log.Debug("  Failed!")
			return false, nil
		}
	}
	return true, nil
}

func (b *BigBoard) follow(loc Location, tileSet [][]*Tile) (Location, error) {
	// which tile are we entering?
	// log.Debugf("    entering tile [%v,%v] from %v", loc.row, loc.col, loc.tunnel)
	thisTile := tileSet[loc.row][loc.col]
	// how does that tile route us?
	outlet, e := thisTile.Follow(loc.tunnel)
	if e != nil {
		return Location{}, e
	}
	// log.Debugf("    that tunnel goes to %v", outlet)
	// which tile is this outlet pointing to?
	newLoc := Location{row: loc.row, col: loc.col}
	if outlet < TunnelRightTop {
		newLoc.row = newLoc.row - 1
	} else if outlet < TunnelBottomRight {
		newLoc.col = newLoc.col + 1
	} else if outlet < TunnelLeftBottom {
		newLoc.row = newLoc.row + 1
	} else {
		newLoc.col = newLoc.col - 1
	}
	// log.Debugf("    next tile is [%v,%v]", newLoc.row, newLoc.col)
	// is there a tile to go to there?
	if newLoc.row < 0 || newLoc.row > 2 || newLoc.col < 0 || newLoc.col > 2 {
		// no? new location is the same as old location, but with the tunnel updated
		loc.tunnel = outlet
		loc.end = true
		// log.Debug("    went off the board! returning old location with outlet tunnel")
		return loc, nil
	}
	// which tunnel are we entering on the new tile?
	inlet, e := NextTunnel(outlet)
	if e != nil {
		return Location{}, e
	}
	// log.Debugf("    still on the board. outlet tunnel translates to inlet tunnel %v", inlet)
	newLoc.tunnel = inlet

	return newLoc, nil
}
