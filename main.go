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

	for true {
		// log.Debugf("%v: %v %v", count, b.TileModel(), b.RotaModel())
		won, e = b.IsSolved()
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
