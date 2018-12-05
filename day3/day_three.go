package main

import (
	"regexp"
	"fmt"
	"first/advent_of_code"
	"strings"
	"strconv"
	"time"
)

type coordsStruct struct {
	id int
	xBegin int
	xEnd int
	yBegin int
	yEnd int
}

var coordsUsageMap = map[string]int{}

func loadCoordsStrings() []string {
	return strings.Split(advent_of_code.LoadInputTxt(3), "\n")
}

func parseCoordsStringToStruct(coordsString string) (coordsStructResult coordsStruct) {
	regex, _ := regexp.Compile(`#(\d+)\s+@\s+(\d+),(\d+):\s(\d+)x(\d+)`)
	res := regex.FindAllStringSubmatch(coordsString, -1)
	if res != nil && len(res[0]) == 6 {
		id, _ := strconv.ParseInt(res[0][1], 10, 32)
		xBegin, _ := strconv.ParseInt(res[0][2], 10, 32)
		yBegin, _ := strconv.ParseInt(res[0][3], 10, 32)
		xSize, _ := strconv.ParseInt(res[0][4], 10, 32)
		ySize, _ := strconv.ParseInt(res[0][5], 10, 32)
		xEnd := xBegin + xSize
		yEnd := yBegin + ySize
		// To make start inclusive.
		xBegin++
		yBegin++
		return coordsStruct{
			int(id),
			int(xBegin),
			int(xEnd),
			int(yBegin),
			int(yEnd),
		}
	}
	return
}


func getUsedCoordsSliceFromStruct(coords coordsStruct) []string {
	usedCoordsStrings := []string{}
	for x := coords.xBegin; x <= coords.xEnd; x++ {
		for y := coords.yBegin; y <= coords.yEnd; y++ {
			usedCoordsStrings = append(usedCoordsStrings, fmt.Sprintf("%d.%d", x, y))
		}
	}
	return usedCoordsStrings
}

func countCoordsUsageFromSlice(coordsSlice []string) {
	for _, coords := range coordsSlice {
		_, coordsExist := coordsUsageMap[coords]
		if !coordsExist {
			coordsUsageMap[coords] = 1
		} else {
			coordsUsageMap[coords]++
		}
	}
}

func main() {
	startTime := time.Now()
	coordsStrings := loadCoordsStrings()
	coordsStructs := []coordsStruct{}
	for _, coords := range coordsStrings {
		coordsStructs = append(coordsStructs, parseCoordsStringToStruct(coords))
	}
	for _, coordsStruct := range coordsStructs {
		coordsSlice := getUsedCoordsSliceFromStruct(coordsStruct)
		countCoordsUsageFromSlice(coordsSlice)
	}
	overusedCount := 0
	for _, count := range coordsUsageMap {
		if count > 1 {
			overusedCount++
		}
	}
	fmt.Printf("\nNumber of square inches that overlap (Part 1): %d", overusedCount)

	// Part 2
	var winningCoords coordsStruct
	for _, coordsStruct := range coordsStructs {
		coordsSlice := getUsedCoordsSliceFromStruct(coordsStruct)
		good := 1
		for _, coords := range coordsSlice {
			if coordsUsageMap[coords] != 1 {
				good = 0
				break
			}
		}
		if good == 1 {
			winningCoords = coordsStruct
		}
	}
	fmt.Printf("\nID of plot that does not overlap (Part 2): %d", winningCoords.id)
	duration := time.Since(startTime)
	fmt.Println("\n\nScript took ", duration)
}
