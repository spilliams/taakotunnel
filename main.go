/*
This puzzle took the form of a 3x3 grid, with various tunnel entrances around
the outer edge, along with 9 tiles with tunnels on them.
*/

package main

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// type solution struct {
// 	tiles     string
// 	rotations string
// }

func main() {
	// log.SetLevel(log.DebugLevel)

	boardTunnels := map[int]int{
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
	tiles := makeTiles()

	tileModel := "12345678"
	rotaModel := "00000000"
	endTileModel := "87654321"
	count := 0

	for tileModel < endTileModel {
		log.Debugf("Trying %v %v", tileModel, rotaModel)
		set, e := makeTileSet(tiles, tileModel, rotaModel)
		check(e)

		if checkBoard(boardTunnels, set) {
			fmt.Printf("Hooray, we have a solution! after %v tries!\n", count)
			fmt.Printf("Solution: %v %v", tileModel, endTileModel)
		}
		count++
		if count%10000000 == 0 {
			log.Infof("Tried %v combinations. Current tile set %v\n", count, printTileSet(set, true))
		}
		tileModel, rotaModel = permute(tileModel, rotaModel)
	}
	log.Debugf("Trying %v %v", tileModel, rotaModel)
	set, e := makeTileSet(tiles, tileModel, rotaModel)
	check(e)

	if checkBoard(boardTunnels, set) {
		fmt.Printf("Hooray, we have a solution! after %v tries!\n", count)
		fmt.Printf("Solution: %v %v", tileModel, endTileModel)
	}
	count++
	if count%10000000 == 0 {
		log.Infof("Tried %v combinations. Current model %v %v\n", count, tileModel, rotaModel)
	}
	tileModel, rotaModel = permute(tileModel, rotaModel)
}

func makeTileSet(tiles []*Tile, tileModel string, rotaModel string) ([][]*Tile, error) {
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
	for i := 0; i < len(rotaModel); i++ {
		rotateTo64, _ := strconv.ParseInt(rotaModel[i:i+1], 10, 0)
		e = tiles[i].RotateTopTo(int(rotateTo64))
		if e != nil {
			return nil, e
		}
	}

	tileSet := [][]*Tile{
		{
			tiles[0],
			tiles[1],
			tiles[2],
		}, {
			tiles[3],
			centerTile,
			tiles[4],
		}, {
			tiles[5],
			tiles[6],
			tiles[7],
		},
	}

	return tileSet, nil
}

func printTileSet(tileSet [][]*Tile, oneLine bool) string {
	p := ""
	for _, tileRow := range tileSet {
		for _, tile := range tileRow {
			p = p + fmt.Sprintf("%v%v ", tile.letter, tile.rotation)
		}
		if !oneLine {
			p = p + "\n"
		}
	}
	return p
}

func makeTiles() []*Tile {
	tile1, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelRightTop},
		{In: TunnelTopRight, Out: TunnelBottomRight},
		{In: TunnelBottomLeft, Out: TunnelLeftTop},
		{In: TunnelLeftBottom, Out: TunnelRightBottom},
	}, "a")
	check(e)
	tile2, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomLeft},
		{In: TunnelTopRight, Out: TunnelRightTop},
		{In: TunnelRightBottom, Out: TunnelLeftTop},
		{In: TunnelBottomRight, Out: TunnelLeftBottom},
	}, "b")
	check(e)
	tile3, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelLeftTop},
		{In: TunnelRightTop, Out: TunnelLeftBottom},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
	}, "c")
	check(e)
	tile4, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelRightTop},
		{In: TunnelTopRight, Out: TunnelBottomRight},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
		{In: TunnelLeftBottom, Out: TunnelLeftTop},
	}, "e")
	check(e)
	tile5, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelRightTop},
		{In: TunnelRightBottom, Out: TunnelLeftBottom},
		{In: TunnelBottomLeft, Out: TunnelLeftTop},
	}, "f")
	check(e)
	tile6, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelBottomRight},
		{In: TunnelTopRight, Out: TunnelLeftBottom},
		{In: TunnelRightTop, Out: TunnelLeftTop},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
	}, "g")
	check(e)
	tile7, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelTopRight},
		{In: TunnelRightTop, Out: TunnelLeftTop},
		{In: TunnelRightBottom, Out: TunnelBottomLeft},
		{In: TunnelBottomRight, Out: TunnelLeftBottom},
	}, "h")
	check(e)
	tile8, e := NewTile([]Tunnel{
		{In: TunnelTopLeft, Out: TunnelLeftTop},
		{In: TunnelTopRight, Out: TunnelRightBottom},
		{In: TunnelRightTop, Out: TunnelBottomRight},
		{In: TunnelBottomLeft, Out: TunnelLeftBottom},
	}, "i")
	check(e)
	tiles := []*Tile{
		tile1,
		tile2,
		tile3,
		tile4,
		tile5,
		tile6,
		tile7,
		tile8,
	}
	return tiles
}

type location struct {
	row    int // 0-2
	col    int // 0-2
	tunnel int // 0-8
	end    bool
}

func checkBoard(boardTunnels map[int]int, tileSet [][]*Tile) bool {
	boardMap := makeBoardMap(tileSet)

	for startingTunnel, endingTunnel := range boardTunnels {
		log.Debugf("Checking that tunnel %v connects to %v", startingTunnel, endingTunnel)
		loc := boardMap[startingTunnel]
		log.Debugf("  starting at %v", loc)
		for !loc.end {
			loc = follow(loc, tileSet)
			log.Debugf("  followed to %v", loc)
		}
		if boardMap[endingTunnel].row != loc.row || boardMap[endingTunnel].col != loc.col || boardMap[endingTunnel].tunnel != loc.tunnel {
			log.Debug("  Failed!")
			return false
		}
	}
	return true
}

// returns a map from tunnel number to location{tile, tiletunnel}
func makeBoardMap(tileSet [][]*Tile) map[int]location {
	return map[int]location{
		1:  location{row: 0, col: 0, tunnel: TunnelTopLeft},
		2:  location{row: 0, col: 0, tunnel: TunnelTopRight},
		3:  location{row: 0, col: 1, tunnel: TunnelTopLeft},
		4:  location{row: 0, col: 1, tunnel: TunnelTopRight},
		5:  location{row: 0, col: 2, tunnel: TunnelTopLeft},
		6:  location{row: 0, col: 2, tunnel: TunnelTopRight},
		7:  location{row: 0, col: 2, tunnel: TunnelRightTop},
		8:  location{row: 0, col: 2, tunnel: TunnelRightBottom},
		9:  location{row: 1, col: 2, tunnel: TunnelRightTop},
		10: location{row: 1, col: 2, tunnel: TunnelRightBottom},
		11: location{row: 2, col: 2, tunnel: TunnelRightTop},
		12: location{row: 2, col: 2, tunnel: TunnelRightBottom},
		13: location{row: 2, col: 2, tunnel: TunnelBottomRight},
		14: location{row: 2, col: 2, tunnel: TunnelBottomLeft},
		15: location{row: 2, col: 1, tunnel: TunnelBottomRight},
		16: location{row: 2, col: 1, tunnel: TunnelBottomLeft},
		17: location{row: 2, col: 0, tunnel: TunnelBottomRight},
		18: location{row: 2, col: 0, tunnel: TunnelBottomLeft},
		19: location{row: 2, col: 0, tunnel: TunnelLeftBottom},
		20: location{row: 2, col: 0, tunnel: TunnelLeftTop},
		21: location{row: 1, col: 0, tunnel: TunnelLeftBottom},
		22: location{row: 1, col: 0, tunnel: TunnelLeftTop},
		23: location{row: 0, col: 0, tunnel: TunnelLeftBottom},
		24: location{row: 0, col: 0, tunnel: TunnelLeftTop},
	}
}

func follow(loc location, tileSet [][]*Tile) location {
	// which tile are we entering?
	log.Debugf("    entering tile [%v,%v] from %v", loc.row, loc.col, loc.tunnel)
	thisTile := tileSet[loc.row][loc.col]
	// how does that tile route us?
	outlet, e := thisTile.Follow(loc.tunnel)
	check(e)
	log.Debugf("    that tunnel goes to %v", outlet)
	// which tile is this outlet pointing to?
	newLoc := location{row: loc.row, col: loc.col}
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
		return loc
	}
	// which tunnel are we entering on the new tile?
	inlet, e := NextTunnel(outlet)
	check(e)
	log.Debugf("    still on the board. outlet tunnel translates to inlet tunnel %v", inlet)
	newLoc.tunnel = inlet

	return newLoc
}
