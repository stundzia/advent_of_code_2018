package main

import (
	"fmt"
	"time"
)

var gridSerialNumber int = 7857

func getHundredthDigit(num int) int {
	if num < 100 {
		return 0
	}
	return num / 100 % 10
}

func getCellPowerLevel(xyCoords []int) int {
	rackID := xyCoords[0] + 10
	powerLevel := rackID * xyCoords[1] + gridSerialNumber
	powerLevel *= rackID
	powerLevel = getHundredthDigit(powerLevel)
	return powerLevel - 5
}

func createCellPowerLevelGrid() [][]int {
	grid := [][]int{}
	for y := 0; y <= 300; y++ {
		row := []int{}
		for x := 0; x <= 300; x++ {
			row = append(row, getCellPowerLevel([]int{x,y}))
		}
		grid = append(grid, row)
	}
	return grid
}

func getGridSectionPowerLevel(xyStart []int, grid [][]int, sectionSize []int) int {
	res := 0
	for x := xyStart[0]; x < xyStart[0] + sectionSize[0]; x++ {
		for y := xyStart[1]; y < xyStart[1] + sectionSize[1]; y++ {
			res += grid[y][x]
		}
	}
	return res
}

func getHighestGridSectionPowerLevel(grid [][]int, sectionSize int) (powerMax int, coords []int) {
	biggestPower := 0
	biggestPowerCoords := []int{0,0}
	power := 0
	for x := 1; x < 300 - sectionSize; x++ {
		for y := 1; y < 300 - sectionSize; y++ {
			power = getGridSectionPowerLevel([]int{x,y}, grid, []int{sectionSize,sectionSize})
			if power > biggestPower {
				biggestPower = power
				biggestPowerCoords = []int{x,y}
			}
		}
	}
	return biggestPower, biggestPowerCoords
}

func main() {
	startTime := time.Now()

	grid := createCellPowerLevelGrid()
	biggestPower := 0
	biggestPowerCoords := []int{0,0}
	power := 0
	for x := 1; x < 298; x++ {
		for y := 1; y < 298; y++ {
			power = getGridSectionPowerLevel([]int{x,y}, grid, []int{3,3})
			if power > biggestPower {
				biggestPower = power
				biggestPowerCoords = []int{x,y}
			}
		}
	}
	fmt.Println(biggestPower, " : ", biggestPowerCoords)

	// Part 2
	biggestPower2 := 0
	biggestPower2Coords := []int{0,0}
	power2 := 0
	coords := []int{0,0}
	bestSize := 0
	for size := 1; size < 300; size++ {
		power2, coords = getHighestGridSectionPowerLevel(grid, size)
		if power2 > biggestPower2 {
			biggestPower2 = power2
			biggestPower2Coords = coords
			bestSize = size
		}
	}

	fmt.Println(biggestPower2, " : ", biggestPower2Coords, " : ", bestSize)

	duration := time.Since(startTime)
	fmt.Println("\n\nScript took ", duration)
}
