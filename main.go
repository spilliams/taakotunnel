/*
This puzzle took the form of a 3x3 grid, with various tunnel entrances around
the outer edge, along with 9 tiles with tunnels on them.
*/

package main

import (
	"fmt"
	"strconv"

	"github.com/spilliams/taakotunnel/model"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	b, e := model.NewBigBoard()
	check(e)

	count := 0
	var won bool
	var set [][]*model.Tile

	// b.SetTileModel("2104635")
	// b.SetRotaModel("3201221")

	for true {
		// log.Debugf("%v: %v %v", count, b.TileModel(), b.RotaModel())
		won, e = IsSolved(b)
		check(e)
		if won {
			break
		}
		tileModel, rotaModel, e := permute(b.TileModel(), b.RotaModel())
		if e != nil {
			log.Infof("error from permuter: %v (%v)", e, tileModel)
			break
		}
		b.SetTileModel(tileModel)
		b.SetRotaModel(rotaModel)
		count++
		if count%1000000 == 0 {
			set, e = b.MakeTileSet()
			countS := strconv.FormatInt(int64(count), 10)
			for len(countS) < 10 {
				countS = fmt.Sprintf(" %v", countS)
			}
			log.Infof("%v:\t%v\t%v %v", countS, printTileSet(set, true), b.TileModel(), b.RotaModel())
		}
	}
	if won {
		fmt.Printf("Hooray, we have a solution! after %v tries!\n", count)
		set, e = b.MakeTileSet()
		check(e)
		fmt.Println(printTileSet(set, false))
	} else {
		fmt.Printf("Failed to find a solution in %v tries\n", count)
	}
}

func checkBase4() error {
	type permutation struct {
		t string
		r string
	}
	tileModel := "0123"
	rotaModel := "0000"
	tries := []permutation{{t: tileModel, r: rotaModel}}

	for true {
		fmt.Printf("%v %v\n", tileModel, rotaModel)
		var e error
		tileModel, rotaModel, e = permute(tileModel, rotaModel)
		if e != nil {
			return e
		}
		tries = append(tries, permutation{t: tileModel, r: rotaModel})
		if len(tries)%1000 == 0 {
			fmt.Println(len(tries))
		}
	}

	fmt.Printf("%v attempts", len(tries))
	fmt.Println(tries)
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printTileSet(tileSet [][]*model.Tile, oneLine bool) string {
	p := ""
	for _, tileRow := range tileSet {
		for _, tile := range tileRow {
			p = p + fmt.Sprintf("%v%v ", tile.Letter(), tile.Rotation())
		}
		if !oneLine {
			p = p + "\n"
		}
	}
	return p
}

func IsSolved(b model.Board) (bool, error) {
	numOK, e := numOK(b)
	// if numOK == 12 {
	// 	printTunnels(b, numOK)
	// }
	if numOK == len(b.EdgeTunnels()) {
		return true, e
	}
	return false, e
}

// func printTunnels(b model.Board, num int) {
// 	log.Infof("OK Tunnels for %v %v", b.TileModel(), b.RotaModel())
// 	t := b.EdgeTunnels()
// 	for i := 0; i < num; i++ {
// 		log.Infof("  %v", t[i])
// 	}
// }

func numOK(b model.Board) (int, error) {
	numOK := 0
	tileSet, e := b.MakeTileSet()
	if e != nil {
		return numOK, e
	}
	boardMap := b.TunnelMap()
	// log.Debugf("Checking is board solved for %v %v", b.TileModel(), b.RotaModel())
	for _, tunnel := range b.EdgeTunnels() {
		// log.Debugf("  Tunnel %v:", tunnel)
		loc := boardMap[tunnel.In]
		for !loc.End {
			loc, e = follow(b, loc, tileSet)
			if e != nil {
				// log.Debugf("  Error %v", e)
				return numOK, e
			}
		}
		if boardMap[tunnel.Out].Row == loc.Row && boardMap[tunnel.Out].Col == loc.Col && boardMap[tunnel.Out].Tunnel == loc.Tunnel {
			// log.Debug("  Good!")
			numOK++
		} else {
			// log.Debugf("  Expected %v Actual %v", boardMap[tunnel.Out], loc)
		}
	}
	return numOK, nil
}

func follow(b model.Board, loc model.Location, tileSet [][]*model.Tile) (model.Location, error) {
	// which tile are we entering?
	thisTile := tileSet[loc.Row][loc.Col]
	// log.Debugf("    entering tile [%v,%v] (%v) from %v", loc.Row, loc.Col, thisTile.Letter(), loc.Tunnel)
	// based on the tile's rotation, which of the tile's tunnels is this?
	tileInlet := (loc.Tunnel+7+(thisTile.Rotation()*2))%8 + 1
	// how does that tile route us?
	tileOutlet, e := thisTile.Follow(tileInlet)
	if e != nil {
		return model.Location{}, e
	}
	// log.Debugf("    wrt the tile, that's %v to %v", tileInlet, tileOutlet)
	outlet := (tileOutlet+7-(thisTile.Rotation()*2))%8 + 1
	// log.Debugf("    that tunnel goes to %v", outlet)
	// which tile is this outlet pointing to?
	newLoc := model.Location{Row: loc.Row, Col: loc.Col}
	if outlet < model.TunnelRightTop {
		newLoc.Row = newLoc.Row - 1
	} else if outlet < model.TunnelBottomRight {
		newLoc.Col = newLoc.Col + 1
	} else if outlet < model.TunnelLeftBottom {
		newLoc.Row = newLoc.Row + 1
	} else {
		newLoc.Col = newLoc.Col - 1
	}
	// log.Debugf("    next tile is [%v,%v]", newLoc.Row, newLoc.Col)

	// is there a tile to go to there?
	if newLoc.Row < 0 || newLoc.Row >= b.Size() || newLoc.Col < 0 || newLoc.Col >= b.Size() {
		// no? new location is the same as old location, but with the tunnel updated
		loc.Tunnel = outlet
		loc.End = true
		// log.Debug("    went off the board! returning old location with outlet tunnel")
		return loc, nil
	}
	// which tunnel are we entering on the new tile?
	inlet, e := model.NextTunnel(outlet)
	if e != nil {
		return model.Location{}, e
	}
	// log.Debugf("    still on the board. outlet tunnel translates to inlet tunnel %v", inlet)
	newLoc.Tunnel = inlet

	return newLoc, nil
}
