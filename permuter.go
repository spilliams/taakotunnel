package main

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

// TODO: The middle tile is always the same, and always rotated the same.

var tilePermutationBase = 9
var maxRotationModel, _ = strconv.ParseInt("333333333", 4, 64)

func permute(tileModel, rotaModel string) (string, string) {
	rotaInt, _ := strconv.ParseInt(rotaModel, 4, 64)
	newTileModel := tileModel
	newRotaInt := rotaInt
	if rotaInt < maxRotationModel {
		newRotaInt = permuteRotation(rotaInt)
	} else {
		newTileModel = permuteTiles(tileModel)
		newRotaInt = 0
	}
	newRotaModel := strconv.FormatInt(newRotaInt, 4)
	return newTileModel, newRotaModel
}

func permuteRotation(rotaInt int64) int64 {
	return rotaInt + 1
}

func permuteTiles(tileModel string) string {
	tileModel = incrementTileModel(tileModel)
	for !validTileModel(tileModel) {
		tileModel = incrementTileModel(tileModel)
	}
	return tileModel
}

func incrementTileModel(tileModel string) string {
	tileModel = stringAdd(tileModel, -1)
	tileInt, e := strconv.ParseInt(tileModel, tilePermutationBase, 64)
	check(e)
	tileInt++
	tileModel = strconv.FormatInt(tileInt, tilePermutationBase)
	return stringAdd(tileModel, 1)
}

func stringAdd(st string, add int64) string {
	out := ""
	for i := 0; i < len(st); i++ {
		in, e := strconv.ParseInt(st[i:i+1], 10, 64)
		check(e)
		in += add
		out += strconv.FormatInt(in, 10)
	}
	return out
}

func validTileModel(tileModel string) bool {
	log.Debugf("validTileModel %v in base %v?", tileModel, tilePermutationBase)
	if len(tileModel) != tilePermutationBase {
		log.Debugf("  no, tileModel incorrect length")
		return false
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
			log.Debugf("  no, couldn't find %v in %v", lookfor, tileModel)
		}
	}
	return true
}
