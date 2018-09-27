/*
This puzzle took the form of a 3x3 grid, with various tunnel entrances around
the outer edge, along with 9 tiles with tunnels on them.
*/

package main

import (
	"fmt"

	"github.com/spilliams/taakotunnel/model"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	b, e := model.NewSmallBoard()
	check(e)

	count := 0

	var won bool

	for b.TileModel() < b.EndTileModel() {
		won, e = b.IsSolved()
		check(e)
		if won {
			break
		}
		count++
		tileModel, rotaModel, e := permute(b.TileModel(), b.RotaModel())
		check(e)
		b.SetTileModel(tileModel)
		b.SetRotaModel(rotaModel)
	}
	won, e = b.IsSolved()
	check(e)
	if won {
		fmt.Printf("Hooray, we have a solution! after %v tries!\n", count)
		set, e := b.MakeTileSet()
		check(e)
		fmt.Println(printTileSet(set, false))
	}
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
