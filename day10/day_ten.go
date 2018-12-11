package main

import (
	"first/advent_of_code"
	"regexp"
	"fmt"
	"strconv"
	"time"
)

type StarStruct struct {
	PositionXY []int
	VelocityXY []int
}


func (ss *StarStruct) passSecond() {
	ss.PositionXY[0] += ss.VelocityXY[0]
	ss.PositionXY[1] += ss.VelocityXY[1]
}

func (ss *StarStruct) undoSecond() {
	ss.PositionXY[0] -= ss.VelocityXY[0]
	ss.PositionXY[1] -= ss.VelocityXY[1]
}

func loadStarData() []string {
	return advent_of_code.LoadInputAsStringSlice(10)
}

func parseStarEntry(entry string) StarStruct {
	starEntryRegex, _ := regexp.Compile(`position=<\s?(\-?\d+),\s+?(\-?\d+)>\svelocity=<\s?(\-?\d+),\s+(\-?\d+)`)
	res := starEntryRegex.FindAllStringSubmatch(entry, -1)
	positionX, _ := strconv.ParseInt(res[0][1], 10, 32)
	positionY, _ := strconv.ParseInt(res[0][2], 10, 32)
	velocityX, _ := strconv.ParseInt(res[0][3], 10, 32)
	velocityY, _ := strconv.ParseInt(res[0][4], 10, 32)
	return StarStruct{
		PositionXY: []int{int(positionX), int(positionY)},
		VelocityXY: []int{int(velocityX), int(velocityY)},
	}
}

func getCoordsRange(starStructs []StarStruct) (xRange []int, yRange []int) {
	minX := starStructs[0].PositionXY[0]
	maxX := starStructs[0].PositionXY[0]
	minY := starStructs[0].PositionXY[1]
	maxY := starStructs[0].PositionXY[1]
	for _, ss := range starStructs {
		if ss.PositionXY[0] < minX {
			minX = ss.PositionXY[0]
		} else if ss.PositionXY[0] > maxX {
			maxX = ss.PositionXY[0]
		}
		if ss.PositionXY[1] < minY {
			minY = ss.PositionXY[1]
		} else if ss.PositionXY[1] > maxY {
			maxY = ss.PositionXY[1]
		}
	}
	return []int{minX - 10, maxX + 10}, []int{minY - 10, maxY + 10}
}

func getParsedStarData() []StarStruct {
	starEntries := []StarStruct{}
	for _, entry := range loadStarData() {
		starEntries = append(starEntries, parseStarEntry(entry))
	}
	return starEntries
}

func createBlankMatrix(starStructs []StarStruct) [][]string {
	xRange, yRange := getCoordsRange(starStructs)
	res := [][]string{}
	line := []string{}
	for y:= yRange[0]; y < yRange[1]; y++ {
		for x := xRange[0]; x < xRange[1]; x++ {
			line = append(line, ".")
		}
		res = append(res, line)
		line = []string{}
	}
	return res
}

func placeStarsOnMatrix(starStructs []StarStruct, matrix [][]string) [][]string {
	xRange, yRange := getCoordsRange(starStructs)
	for _, ss := range starStructs {
		matrix[ss.PositionXY[1] - xRange[0]][ss.PositionXY[0] - yRange[0]] = "#"
	}
	return matrix
}

func getRangesDiffSum(starStructs []StarStruct) int {
	xRange, yRange := getCoordsRange(starStructs)
	return (xRange[1] - xRange[0]) + (yRange[1] - yRange[0])
}

func main() {
	startTime := time.Now()

	rr := getParsedStarData()

	//
	prevRangesDiff := getRangesDiffSum(rr)
	secondsPassed := 0
	for getRangesDiffSum(rr) <= prevRangesDiff {
		prevRangesDiff = getRangesDiffSum(rr)
		for _, ss := range rr {
			ss.passSecond()
		}
		secondsPassed++
	}

	// Undo last move.
	for _, ss := range rr {
		ss.undoSecond()
	}
	secondsPassed--

	fmt.Println("Part 1: ")
	matrix := createBlankMatrix(rr)
	matrix = placeStarsOnMatrix(rr, matrix)
	for _, line := range matrix {
		fmt.Println(line)
	}

	fmt.Println("Part 2 result: ", secondsPassed)

	duration := time.Since(startTime)
	fmt.Println("\n\nScript took ", duration)
}
