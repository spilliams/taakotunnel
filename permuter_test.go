package main

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestStringAdd(t *testing.T) {
	tests := []struct {
		in  string
		add int64
		out string
	}{
		{in: "1234", add: -1, out: "0123"},
		{in: "3120", add: 1, out: "4231"},
		{in: "0123", add: 1, out: "1234"},
		{in: "1234", add: -1, out: "0123"},
	}
	for _, tc := range tests {
		actual := stringAdd(tc.in, tc.add)
		if actual != tc.out {
			t.Fatalf("tried to add %v to %v, got %v instead of %v", tc.add, tc.in, actual, tc.out)
		}
	}
}

func TestIncrementTileModel(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	tilePermutationBase = 4
	tm := "1234"
	increments := []string{
		"1241",
		"1242",
		"1243",
		"1244",
		"1311",
		"1312",
		"1313",
		"1314",
		"1321",
		"1322",
		"1323",
		"1324",
		"1331",
		"1332",
		"1333",
		"1334",
		"1341",
		"1342",
		"1343",
		"1344",
		"1411",
		"1412",
		"1413",
		"1414",
		"1421",
		"1422",
		"1423",
		"1424",
		"1431",
		"1432",
		"1433",
		"1434",
		"1441",
		"1442",
		"1443",
		"1444",
		"2111",
	}
	for _, shouldbe := range increments {
		origTm := tm
		tm = incrementTileModel(tm)
		if tm != shouldbe {
			t.Fatalf("tile increment from %v should be %v, actual %v", origTm, shouldbe, tm)
		}
	}
}

func TestValidTileModel(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	tilePermutationBase = 4
	ps := map[string]bool{
		"1234": true,
		"1111": false,
		"234":  false,
		"4321": true,
		"4000": false,
	}
	for p, shouldBe := range ps {
		if validTileModel(p) != shouldBe {
			t.Fatalf("%v should be %v valid, is %v valid", p, shouldBe, validTileModel(p))
		}
	}
}

func TestTilePermuterBase3(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	fmt.Printf("\n\n\n")
	log.Info("Starting TestTilePermuterBase3")

	tilePermutationBase = 3
	p := "123"
	ps := []string{
		"132",
		"213",
		"231",
		"312",
		"321",
	}

	for _, shouldBe := range ps {
		log.Debugf("permutation %v", p)
		origP := p
		p, _ = permute(p, "333")
		if p != shouldBe {
			t.Fatalf("permutation of %v should be %v, actual %v", origP, shouldBe, p)
		}
		fmt.Println()
	}
}

func TestTilePermuterBase4(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	fmt.Printf("\n\n\n")
	log.Info("Starting TestTilePermuterBase4")

	tilePermutationBase = 4
	p := "1234"
	ps := []string{
		"1243", "1324", "1342", "1423", "1432", "2134", "2143", "2314", "2341", "2413", "2431", "3124", "3142", "3214", "3241", "3412", "3421",
	}

	for _, shouldBe := range ps {
		log.Debugf("permutation %v", p)
		origP := p
		p, _ = permute(p, "3333")
		if p != shouldBe {
			t.Fatalf("permutation of %v should be %v, actual %v", origP, shouldBe, p)
		}
		fmt.Println()
	}
}

func TestWholePermuterBase3(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	fmt.Printf("\n\n\n")
	log.Info("Starting TestWholePermuterBase3")

	tilePermutationBase = 3
	type permutation struct {
		t string
		r string
	}
	p := permutation{t: "123", r: "000"}
	ps := []permutation{
		{t: "123", r: "001"},
		{t: "123", r: "002"},
		{t: "123", r: "003"},
		{t: "123", r: "010"},
		{t: "123", r: "011"},
		{t: "123", r: "012"},
		{t: "123", r: "013"},
		{t: "123", r: "020"},
		{t: "123", r: "021"},
		{t: "123", r: "022"},
		{t: "123", r: "023"},
		{t: "123", r: "030"},
		{t: "123", r: "031"},
		{t: "123", r: "032"},
		{t: "123", r: "033"},
		{t: "123", r: "100"},
		{t: "123", r: "101"},
		{t: "123", r: "102"},
		{t: "123", r: "103"},
		{t: "123", r: "110"},
		{t: "123", r: "111"},
		{t: "123", r: "112"},
		{t: "123", r: "113"},
		{t: "123", r: "120"},
		{t: "123", r: "121"},
		{t: "123", r: "122"},
		{t: "123", r: "123"},
		{t: "123", r: "130"},
		{t: "123", r: "131"},
		{t: "123", r: "132"},
		{t: "123", r: "133"},
		{t: "123", r: "200"},
		{t: "123", r: "201"},
		{t: "123", r: "202"},
		{t: "123", r: "203"},
		{t: "123", r: "210"},
		{t: "123", r: "211"},
		{t: "123", r: "212"},
		{t: "123", r: "213"},
		{t: "123", r: "220"},
		{t: "123", r: "221"},
		{t: "123", r: "222"},
		{t: "123", r: "223"},
		{t: "123", r: "230"},
		{t: "123", r: "231"},
		{t: "123", r: "232"},
		{t: "123", r: "233"},
		{t: "123", r: "300"},
		{t: "123", r: "301"},
		{t: "123", r: "302"},
		{t: "123", r: "303"},
		{t: "123", r: "310"},
		{t: "123", r: "311"},
		{t: "123", r: "312"},
		{t: "123", r: "313"},
		{t: "123", r: "320"},
		{t: "123", r: "321"},
		{t: "123", r: "322"},
		{t: "123", r: "323"},
		{t: "123", r: "330"},
		{t: "123", r: "331"},
		{t: "123", r: "332"},
		{t: "123", r: "333"},
		{t: "132", r: "000"},
	}

	for _, shouldBe := range ps {
		log.Debugf("permutation %v", p)
		origP := p
		p.t, p.r = permute(p.t, p.r)
		if p != shouldBe {
			t.Fatalf("permutation of %v should be %v, actual %v", origP, shouldBe, p)
		}
		fmt.Println()
	}
}
