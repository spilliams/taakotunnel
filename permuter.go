package main

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var rotaPermutationBase = 4

func permute(tileModel, rotaModel string) (string, string, error) {
	tilePermutationBase := len(tileModel)

	rotaInt, _ := strconv.ParseInt(rotaModel, rotaPermutationBase, 64)
	newRotaInt := rotaInt + 1
	newRotaModel := strconv.FormatInt(newRotaInt, rotaPermutationBase)
	// pad rota model
	for len(newRotaModel) < tilePermutationBase {
		newRotaModel = fmt.Sprintf("0%v", newRotaModel)
	}
	if len(newRotaModel) > tilePermutationBase {
		newRotaModel = "0"
		for len(newRotaModel) < tilePermutationBase {
			newRotaModel = fmt.Sprintf("0%v", newRotaModel)
		}
		tileModel, e := permuteTiles(tileModel)
		return tileModel, newRotaModel, e
	}
	return tileModel, newRotaModel, nil
}

func permuteTiles(tileModel string) (string, error) {
	tileModel = incrementTileModel(tileModel)
	valid := false
	var e error
	for !valid && e == nil {
		valid, e = validateTileModel(tileModel)
		if e != nil {
			return tileModel, e
		}
		tileModel = incrementTileModel(tileModel)
	}
	return tileModel, nil
}

func incrementTileModel(tileModel string) string {
	tilePermutationBase := len(tileModel)
	// log.Debugf("increment tile model %v", tileModel)
	// log.Debugf("subtracted 1: %v", tileModel)
	tileInt, e := strconv.ParseInt(tileModel, tilePermutationBase, 64)
	check(e)
	// log.Debugf("converted to int in base %v: %v", tilePermutationBase, tileInt)
	tileInt++
	// log.Debugf("incremented: %v", tileInt)
	tileModel = strconv.FormatInt(tileInt, tilePermutationBase)
	// log.Debugf("converted to string: %v", tileModel)
	for len(tileModel) < tilePermutationBase {
		tileModel = fmt.Sprintf("0%v", tileModel)
	}
	// log.Debugf("padded string: %v", tileModel)
	// log.Debugf("added 1: %v", tileModel)
	return tileModel
}

func validateTileModel(tileModel string) (bool, error) {
	tilePermutationBase := len(tileModel)
	log.Debugf("validateTileModel %v in base %v?", tileModel, tilePermutationBase)
	if len(tileModel) != tilePermutationBase {
		// log.Debugf("  no, tileModel incorrect length")
		return false, fmt.Errorf("tileModel incorrect length")
	}
	for i := 1; i <= tilePermutationBase; i++ {
		// make sure tileModel has char
		lookfor := strconv.Itoa(i)
		found := false
		for j := 0; j < tilePermutationBase; j++ {
			if tileModel[j:j+1] == lookfor {
				found = true
			}
		}
		if !found {
			// log.Debugf("  no, couldn't find %v in %v", lookfor, tileModel)
			return false, nil
		}
	}
	return true, nil
}
