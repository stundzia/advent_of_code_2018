package main

import (
	"first/advent_of_code"
	"strings"
	"fmt"
	"errors"
	"time"
)

var similarIdMap = map[string][]string{}

func getCountsFromIdString(id string) (appearTwice int, appearThrice int) {
	c := make([]int, 0, len(id))
	m := make(map[int]int)
	for _, val := range id {
		if _, ok := m[int(val)]; !ok {
			m[int(val)] = 1
			c = append(c, int(val))
		} else {
			m[int(val)]++
		}
	}
	threes := 0
	twos := 0
	for _, val := range m {
		if val == 2 {
			twos = 1
		}
		if val == 3 {
			threes = 1
		}
	}
	return twos, threes
}

func compareIds(id1 []rune, id2 []rune) bool {
	idLen := len(id1)
	mismatch := 0
	for i := 0; i < idLen && i < len(id2); i++ {
		if mismatch == 2 {
			break
		}
		if id1[i] == id2[i] {
			continue
		} else {
			mismatch++
		}
	}
	if mismatch == 1 {
		return true
	}
	return false
}

func getSimilarIds(id string, idsList []string) {
	splitId := []rune(id)
	var similar bool
	for _, idCompare := range idsList {
		splitCompareId := []rune(idCompare)
		similar = compareIds(splitId, splitCompareId)
		if similar == true {
			_, idExists := similarIdMap[id]
			_, compareIdExists := similarIdMap[idCompare]
			if !compareIdExists {
				if !idExists {
					similarIdMap[id] = []string{idCompare}
				} else {
					similarIdMap[id] = append(similarIdMap[id], idCompare)
				}
			}
		}
	}
}


func getSimilarIdSliceFromMap(similarIdMap map[string][]string) []string {
	similarIdSlice := []string{}
	for id, similarIds := range similarIdMap {
		similarIdSlice = append(similarIdSlice, id)
		for _, similarId := range similarIds {
			similarIdSlice = append(similarIdSlice, similarId)
		}
	}
	return similarIdSlice
}

func getIdsSlice() []string {
	input := advent_of_code.LoadInputTxt(2)
	ids := strings.Split(input, "\n")
	return ids
}

func getSharedSubstringFromStringSlice(stringSlice []string) string {
	res := []rune{}
	for i, c := range []rune(stringSlice[0]) {
		for _, stringInSlice := range stringSlice[1:] {
			stringChars := []rune(stringInSlice)
			if stringChars[i] != c {
				break
			}
			res = append(res, c)
		}
	}
	return string(res)
}

func main() {
	startTime := time.Now()
	ids := getIdsSlice()
	doubleAppearers := 0
	tripleAppearers := 0
	for _, id := range ids {
		double, triple := getCountsFromIdString(id)
		doubleAppearers += double
		tripleAppearers += triple
	}
	checksum := doubleAppearers * tripleAppearers
	fmt.Println("Result checksum (Part 1): ", checksum)
	for _, id := range ids {
		getSimilarIds(id, ids)
	}
	if len(similarIdMap) != 1 {
		panic(errors.New("invalid data supplied, result cannot be determined"))
	}
	similarIdSlice := getSimilarIdSliceFromMap(similarIdMap)
	fmt.Println("Common chars in similar ids result (Part 2): ")
	fmt.Println(getSharedSubstringFromStringSlice(similarIdSlice))
	duration := time.Since(startTime)
	fmt.Println("\nScript took ", duration)
}
