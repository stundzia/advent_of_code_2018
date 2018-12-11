package advent_of_code

import (
	"io/ioutil"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("Fatal Exception `%s` : %s", msg, err)
		panic(err)
	}
}

func LoadInputTxt(day int) string {
	fname := fmt.Sprintf("./advent_of_code/day%d/input.txt", day)
	absPath, _ := filepath.Abs(fname)
	dat, err := ioutil.ReadFile(absPath)
	if err != nil {
		failOnError(err, "Could not load file")
	}
	return string(dat)
}

func LoadInputAsStringSlice(day int) []string {
	input := LoadInputTxt(day)
	return strings.Split(input, "\n")
}

func SortRuneSlice(runeSlice []rune) []rune {
	res := []rune{}
	for len(runeSlice) > 0 {
		var lastItem rune = -1
		var lastItemIndex int = -1
		for i, item := range runeSlice {
			if lastItem == -1 {
				lastItem = item
				lastItemIndex = i
				continue
			}
			if item <= lastItem {
				lastItem = item
				lastItemIndex = i
			}
		}
		res = append(res, lastItem)
		if len(runeSlice) == 1 {
			runeSlice = []rune{}
		} else {
			runeSlice = append(runeSlice[:lastItemIndex], runeSlice[lastItemIndex+1:]...)
		}
	}
	return res
}

func DeleteItemFromRuneSlice(slice []rune, itemToPop rune)  (resSlice []rune, foundAndDeleted bool) {
	for index, item := range slice {
		if item == itemToPop {
			slice = append(slice[:index], slice[index+1:]...)
			return slice, true
		}
	}
	return slice, false
}

func InsertIntoIntSliceCircular(slice []int, value int, index int) (resSlice []int, actualIndex int) {
	for index < 0 {
		index = len(slice) + index + 1
	}
	if index >= len(slice) {
		index = index % len(slice)
	}
	slice = append(slice, 0)
	copy(slice[index + 1:], slice[index:])
	slice[index] = value
	return slice, index
}

func PopFromIntSlice(slice []int, index int) (resSlice []int, res int, actualIndex int) {
	for index < 0 {
		index = len(slice) + index + 1
	}
	val := slice[index]
	copy(slice[index:], slice[index + 1:])
	slice[len(slice) - 1] = 0
	slice = slice[:len(slice) - 1]
	return slice, val, index
}


func StringSliceContains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}
