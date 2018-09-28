package main

import (
	"fmt"
	"strconv"
	// log "github.com/sirupsen/logrus"
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
		tileModel, e := permuteTiles(tileModel, tilePermutationBase)
		return tileModel, newRotaModel, e
	}
	return tileModel, newRotaModel, nil
}

func permuteTiles(tileModel string, base int) (string, error) {
	valid := false
	var e error
	for !valid && e == nil {
		tileModel = incrementTileModel(tileModel, base)
		valid, e = validateTileModel(tileModel, base)
		if e != nil {
			return tileModel, e
		}
	}
	return tileModel, nil
}

func incrementTileModel(tileModel string, base int) string {
	// log.Debugf("increment tile model %v", tileModel)
	tileInt, e := strconv.ParseInt(tileModel, base, 64)
	check(e)
	// log.Debugf("converted to int in base %v: %v", base, tileInt)
	tileInt++
	// log.Debugf("incremented: %v", tileInt)
	tileModel = strconv.FormatInt(tileInt, base)
	// log.Debugf("converted to string: %v", tileModel)
	for len(tileModel) < base {
		tileModel = fmt.Sprintf("0%v", tileModel)
	}
	// log.Debugf("padded string: %v", tileModel)
	return tileModel
}

func validateTileModel(tileModel string, base int) (bool, error) {
	// log.Debugf("validateTileModel %v in base %v?", tileModel, base)
	if len(tileModel) != base {
		// log.Debugf("  no, tileModel incorrect length")
		return false, fmt.Errorf("tileModel incorrect length")
	}
	for i := 0; i < base; i++ {
		// make sure tileModel has char
		lookfor := strconv.Itoa(i)
		found := false
		for j := 0; j < base; j++ {
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
