package model

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type BigBoard struct {
	TileBoard
}

func NewBigBoard() (*BigBoard, error) {
	b := BigBoard{}
	b.tileModel = "01234567"
	b.rotaModel = "00000000"
	b.endTileModel = "76543210"
	e := b.makeTiles()
	return &b, e
}

func (b *BigBoard) EdgeTunnels() map[int]int {
	return map[int]int{
		1:  20, // I
		2:  12, // II
		3:  16, // III
		4:  13, // IV
		5:  7,  // V
		6:  9,  // VI
		8:  23, // VII
		10: 15, // VIII
		11: 14, // IX
		17: 19, // X
		18: 21, // XI
		22: 24, // XII
	}
}

func (b *BigBoard) MakeTileSet() ([][]*Tile, error) {
	centerTile, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelLeftTop},
		{In: TunnelRightTop, Out: TunnelRightBottom},
		{In: TunnelBottomLeft, Out: TunnelLeftBottom},
	}, "d")
	if e != nil {
		return nil, e
	}
	centerTile.RotateTopTo(DirectionSouth)

	// rotate tiles
	for i := 0; i < len(b.rotaModel); i++ {
		rotateTo64, _ := strconv.ParseInt(b.rotaModel[i:i+1], 10, 0)
		e = b.tiles[i].RotateTopTo(int(rotateTo64))
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
		tileIndexInts[i] = int(in - 1)
	}

	tileSet := [][]*Tile{
		{
			b.tiles[tileIndexInts[0]],
			b.tiles[tileIndexInts[1]],
			b.tiles[tileIndexInts[2]],
		}, {
			b.tiles[tileIndexInts[3]],
			centerTile,
			b.tiles[tileIndexInts[4]],
		}, {
			b.tiles[tileIndexInts[5]],
			b.tiles[tileIndexInts[6]],
			b.tiles[tileIndexInts[7]],
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
	}, "e")
	if e != nil {
		return e
	}
	tile4, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelRightTop},
		{In: TunnelRightBottom, Out: TunnelLeftBottom},
		{In: TunnelBottomLeft, Out: TunnelLeftTop},
	}, "f")
	if e != nil {
		return e
	}
	tile5, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelLeftBottom},
		{In: TunnelRightTop, Out: TunnelLeftTop},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
	}, "g")
	if e != nil {
		return e
	}
	tile6, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelTopRight},
		{In: TunnelRightTop, Out: TunnelLeftTop},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
		{In: TunnelBottomRight, Out: TunnelLeftBottom},
	}, "h")
	if e != nil {
		return e
	}
	tile7, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelLeftTop},
		{In: TunnelTopRight, Out: TunnelRightBottom},
		{In: TunnelRightTop, Out: TunnelBottomRight},
		{In: TunnelBottomLeft, Out: TunnelLeftBottom},
	}, "i")
	if e != nil {
		return e
	}
	b.tiles = []*Tile{
		tile0,
		tile1,
		tile2,
		tile3,
		tile4,
		tile5,
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

	for startingTunnel, endingTunnel := range b.EdgeTunnels() {
		log.Debugf("Checking that tunnel %v connects to %v", startingTunnel, endingTunnel)
		loc := boardMap[startingTunnel]
		log.Debugf("  starting at %v", loc)
		for !loc.end {
			loc, e = follow(loc, tileSet)
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

func follow(loc Location, tileSet [][]*Tile) (Location, error) {
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
	if newLoc.row < 0 || newLoc.row > 2 || newLoc.col < 0 || newLoc.col > 2 {
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
