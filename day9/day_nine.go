package main

import (
	"regexp"
	"first/advent_of_code"
	"strconv"
	"fmt"
	"sort"
	"time"
)

var marbleCircle = []int{0}
var currentMarbleValue int = 0
var currentMarbleIndex int = 0
var turnsMade int = 0
var playerCount int = 9
var lastPlayerNum int = 0
var lastMarbleValue int = 0
var playerScoreMap = map[int]int{}

func bumpLastPlayerNum() {
	if lastPlayerNum < playerCount {
		lastPlayerNum++
	} else {
		lastPlayerNum = 1
	}
}

func makeScoreTurn() bool {
	score := 0
	scoreExtra := 0
	score += currentMarbleValue
	currentMarbleIndex -= 7
	if currentMarbleIndex < 0 {
		currentMarbleIndex = len(marbleCircle) + currentMarbleIndex
	}
	marbleCircle, scoreExtra, _ = advent_of_code.PopFromIntSlice(marbleCircle, currentMarbleIndex)
	playerScoreMap[lastPlayerNum] += score + scoreExtra
	bumpLastPlayerNum()
	if currentMarbleValue < lastMarbleValue	{
		return true
	} else {
		return false
	}
}

func makeTurn() bool {
	currentMarbleValue++
	turnsMade++
	if currentMarbleValue % 23 == 0 {
		return makeScoreTurn()
	}
	if currentMarbleIndex + 2 == len(marbleCircle) {
		marbleCircle = append(marbleCircle, currentMarbleValue)
		currentMarbleIndex = len(marbleCircle) - 1
	} else {
		marbleCircle, currentMarbleIndex = advent_of_code.InsertIntoIntSliceCircular(marbleCircle, currentMarbleValue, currentMarbleIndex + 2)
	}
	bumpLastPlayerNum()
	if currentMarbleValue < lastMarbleValue	{
		return true
	} else {
		return false
	}
}

func loadGameSettings() (playerCount int, lastMarbleValue int) {
	input := advent_of_code.LoadInputTxt(9)
	settingsRegex, _ := regexp.Compile(`(\d+) players;[^\d]+(\d+)`)
	res := settingsRegex.FindAllStringSubmatch(input, -1)
	count, _ := strconv.ParseInt(res[0][1], 10, 32)
	value, _ := strconv.ParseInt(res[0][2], 10, 32)
	return int(count), int(value)
}

func initPlayerScoreMap(players int) {
	for player := 0; player < players; player++ {
		playerScoreMap[player] = 0
	}
}

func getHighScore() int {
	results := []int{}
	for _, score := range playerScoreMap {
		results = append(results, score)
	}
	sort.Ints(results)
	return results[len(results) - 1]
}


func main() {
	startTime := time.Now()

	playerCount, lastMarbleValue = loadGameSettings()
	lastMarbleValue = lastMarbleValue * 100
	initPlayerScoreMap(playerCount)
	percentDone := 0
	part1Result := 0
	for makeTurn() == true {
		if currentMarbleValue % (lastMarbleValue/100) == 0 {
			if part1Result == 0 {
				part1Result = getHighScore()
				fmt.Println("Part 1 result: ", part1Result)
			}
			percentDone++
			fmt.Println(percentDone, "% done in ", time.Since(startTime))
		}

	}
	part2Result := getHighScore()
	fmt.Println("Part 2 result: ", part2Result)
	duration := time.Since(startTime)
	fmt.Println("\n\nScript took ", duration)
}
