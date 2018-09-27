package main

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestIncrementTileModel(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	tm := "0123"
	increments := []string{
		"0130",
		"0131",
		"0132",
		"0133",
		"0200",
		"0201",
		"0202",
		"0203",
		"0210",
		"0211",
		"0212",
		"0213",
		"0220",
		"0221",
		"0222",
		"0223",
		"0230",
		"0231",
		"0232",
		"0233",
		"0300",
		"0301",
		"0302",
		"0303",
		"0310",
		"0311",
		"0312",
		"0313",
		"0320",
		"0321",
		"0322",
		"0323",
		"0330",
		"0331",
		"0332",
		"0333",
		"1000",
	}
	for _, shouldbe := range increments {
		origTm := tm
		tm = incrementTileModel(tm, 4)
		if tm != shouldbe {
			t.Fatalf("tile increment from %v should be %v, actual %v", origTm, shouldbe, tm)
		}
	}
}

func TestValidateTileModel(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	ps := map[string]bool{
		"0123": true,
		"0000": false,
		"3210": true,
		"3000": false,
		"0133": false,
	}
	for p, shouldBe := range ps {
		valid, e := validateTileModel(p, 4)
		if e != nil {
			t.Fatal(e)
		}
		if valid != shouldBe {
			t.Fatalf("%v should be %v valid, is %v valid", p, shouldBe, valid)
		}
	}
}

func TestTilePermuterBase3(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	fmt.Printf("\n\n\n")
	log.Info("Starting TestTilePermuterBase3")

	p := "012"
	ps := []string{
		"021",
		"102",
		"120",
		"201",
		"210",
	}

	var e error
	for _, shouldBe := range ps {
		log.Debugf("testing permutation %v", p)
		origP := p
		p, _, e = permute(p, "333")
		if e != nil {
			t.Fatal(e)
		}
		log.Debugf("test permutation came back %v", p)
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

	p := "0123"
	ps := []string{"0132", "0213", "0231", "0312", "0321", "1023", "1032", "1203", "1230", "1302", "1320", "2013", "2031", "2103", "2130", "2301", "2310", "3012", "3021", "3102", "3120", "3201", "3210"}

	var e error
	for _, shouldBe := range ps {
		log.Debugf("testing permutation %v", p)
		origP := p
		p, _, e = permute(p, "3333")
		if e != nil {
			t.Fatal(e)
		}
		log.Debugf("test permutation came back %v", p)
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

	type permutation struct {
		t string
		r string
	}
	p := permutation{t: "012", r: "000"}
	ps := []permutation{
		{t: "012", r: "001"},
		{t: "012", r: "002"},
		{t: "012", r: "003"},
		{t: "012", r: "010"},
		{t: "012", r: "011"},
		{t: "012", r: "012"},
		{t: "012", r: "013"},
		{t: "012", r: "020"},
		{t: "012", r: "021"},
		{t: "012", r: "022"},
		{t: "012", r: "023"},
		{t: "012", r: "030"},
		{t: "012", r: "031"},
		{t: "012", r: "032"},
		{t: "012", r: "033"},
		{t: "012", r: "100"},
		{t: "012", r: "101"},
		{t: "012", r: "102"},
		{t: "012", r: "103"},
		{t: "012", r: "110"},
		{t: "012", r: "111"},
		{t: "012", r: "112"},
		{t: "012", r: "113"},
		{t: "012", r: "120"},
		{t: "012", r: "121"},
		{t: "012", r: "122"},
		{t: "012", r: "123"},
		{t: "012", r: "130"},
		{t: "012", r: "131"},
		{t: "012", r: "132"},
		{t: "012", r: "133"},
		{t: "012", r: "200"},
		{t: "012", r: "201"},
		{t: "012", r: "202"},
		{t: "012", r: "203"},
		{t: "012", r: "210"},
		{t: "012", r: "211"},
		{t: "012", r: "212"},
		{t: "012", r: "213"},
		{t: "012", r: "220"},
		{t: "012", r: "221"},
		{t: "012", r: "222"},
		{t: "012", r: "223"},
		{t: "012", r: "230"},
		{t: "012", r: "231"},
		{t: "012", r: "232"},
		{t: "012", r: "233"},
		{t: "012", r: "300"},
		{t: "012", r: "301"},
		{t: "012", r: "302"},
		{t: "012", r: "303"},
		{t: "012", r: "310"},
		{t: "012", r: "311"},
		{t: "012", r: "312"},
		{t: "012", r: "313"},
		{t: "012", r: "320"},
		{t: "012", r: "321"},
		{t: "012", r: "322"},
		{t: "012", r: "323"},
		{t: "012", r: "330"},
		{t: "012", r: "331"},
		{t: "012", r: "332"},
		{t: "012", r: "333"},
		{t: "021", r: "000"},
	}

	for _, shouldBe := range ps {
		log.Debugf("permutation %v", p)
		origP := p
		var e error
		p.t, p.r, e = permute(p.t, p.r)
		if e != nil {
			t.Fatal(e)
		}
		if p != shouldBe {
			t.Fatalf("permutation of %v should be %v, actual %v", origP, shouldBe, p)
		}
		fmt.Println()
	}
}
