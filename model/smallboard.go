package model

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

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
	log.Debugf("making tile set from %v, %v", b.tileModel, b.rotaModel)
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
		1:  Location{row: 0, col: 0, tunnel: TunnelTopLeft},
		2:  Location{row: 0, col: 0, tunnel: TunnelTopRight},
		3:  Location{row: 0, col: 1, tunnel: TunnelTopLeft},
		4:  Location{row: 0, col: 1, tunnel: TunnelTopRight},
		5:  Location{row: 0, col: 1, tunnel: TunnelRightTop},
		6:  Location{row: 0, col: 1, tunnel: TunnelRightBottom},
		7:  Location{row: 1, col: 1, tunnel: TunnelRightTop},
		8:  Location{row: 1, col: 1, tunnel: TunnelRightBottom},
		9:  Location{row: 1, col: 1, tunnel: TunnelBottomRight},
		10: Location{row: 1, col: 1, tunnel: TunnelBottomLeft},
		11: Location{row: 1, col: 0, tunnel: TunnelBottomRight},
		12: Location{row: 1, col: 0, tunnel: TunnelBottomLeft},
		13: Location{row: 1, col: 0, tunnel: TunnelLeftBottom},
		14: Location{row: 1, col: 0, tunnel: TunnelLeftTop},
		15: Location{row: 0, col: 0, tunnel: TunnelLeftBottom},
		16: Location{row: 0, col: 0, tunnel: TunnelLeftTop},
	}
}

// TODO: these two functions are the same between BigBoard and SmallBoard. DRY them out!

func (b *SmallBoard) IsSolved() (bool, error) {
	tileSet, e := b.MakeTileSet()
	if e != nil {
		return false, e
	}
	boardMap := b.TunnelMap()

	for startingTunnel, endingTunnel := range b.EdgeTunnels() {
		log.Debugf("Checking that tunnel %v connects to %v", startingTunnel, endingTunnel)
		loc := boardMap[startingTunnel]
		log.Debugf("  starting at %v", loc)
		for !loc.end {
			loc, e = b.follow(loc, tileSet)
			if e != nil {
				return false, e
			}
			log.Debugf("  followed to %v", loc)
		}
		if boardMap[endingTunnel].row != loc.row || boardMap[endingTunnel].col != loc.col || boardMap[endingTunnel].tunnel != loc.tunnel {
			log.Debug("  Failed!")
			return false, nil
		}
	}
	return true, nil
}

func (b *SmallBoard) follow(loc Location, tileSet [][]*Tile) (Location, error) {
	// which tile are we entering?
	log.Debugf("    entering tile [%v,%v] from %v", loc.row, loc.col, loc.tunnel)
	thisTile := tileSet[loc.row][loc.col]
	// how does that tile route us?
	outlet, e := thisTile.Follow(loc.tunnel)
	if e != nil {
		return Location{}, e
	}
	log.Debugf("    that tunnel goes to %v", outlet)
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
	log.Debugf("    next tile is [%v,%v]", newLoc.row, newLoc.col)

	// is there a tile to go to there?
	if newLoc.row < 0 || newLoc.row > 1 || newLoc.col < 0 || newLoc.col > 1 {
		// no? new location is the same as old location, but with the tunnel updated
		loc.tunnel = outlet
		loc.end = true
		log.Debug("    went off the board! returning old location with outlet tunnel")
		return loc, nil
	}
	// which tunnel are we entering on the new tile?
	inlet, e := NextTunnel(outlet)
	if e != nil {
		return Location{}, e
	}
	log.Debugf("    still on the board. outlet tunnel translates to inlet tunnel %v", inlet)
	newLoc.tunnel = inlet

	return newLoc, nil
}
