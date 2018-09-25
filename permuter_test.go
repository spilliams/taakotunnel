package main

import (
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
	tm = incrementTileModel(tm)
	if tm != ""
}

func TestTilePermuter(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	tilePermutationBase = 4
	maxRotationModel = 0
	p := "1234"
	ps := []string{
		"1234",
		"1243",
		"1324",
		"1342",
		"1423",
		"1432",
		"2134",
		"2143",
		"2314",
		"2341",
		"2413",
		"2431",
		"3124",
		"3142",
		"3214",
		"3241",
		"3412",
		"3421",
		"4123",
		"4132",
		"4213",
		"4231",
		"4312",
		"4321",
	}
	valid := []string{}

	for validTileModel(p) {
		valid = append(valid, p)
		p, _ = permute(p, "0000")
	}

	if len(valid) != len(ps) {
		t.Fatalf("didn't find all the valid permutations. Stopped on %v", p)
	}
}
