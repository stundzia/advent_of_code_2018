package main

import (
	"first/advent_of_code"
	"strings"
	"regexp"
	"strconv"
	"fmt"
	"time"
)

func parseCoordinatesStringToSlice(coordsString string) []int {
	coordsRegex, _ := regexp.Compile(`(\d+),\s(\d+)`)
	res := coordsRegex.FindAllStringSubmatch(coordsString, -1)
	xCoord, _ := strconv.ParseInt(res[0][1], 10, 32)
	yCoord, _ := strconv.ParseInt(res[0][2], 10, 32)
	return []int{int(xCoord), int(yCoord)}
}

func loadCoordinatesSlice() [][]int {
	coordsStrings := strings.Split(advent_of_code.LoadInputTxt(6), "\n")
	var res = [][]int{}
	for _, coordsString := range coordsStrings {
		res = append(res, parseCoordinatesStringToSlice(coordsString))
	}
	return res
}

func coordsSliceMapWithIds(coordsSlice [][]int) map[string][]int {
	var coordsMap = map[string][]int{}
	for _, coords := range coordsSlice {
		coordsMap[fmt.Sprintf("%d.%d", coords[0], coords[1])] = coords
	}
	return coordsMap
}

func findHighestXY(coordsSlices [][]int) []int {
	highestX := 0
	highestY := 0
	for _, coordsSlice := range coordsSlices {
		if coordsSlice[0] > highestX {
			highestX = coordsSlice[0]
		}
		if coordsSlice[1] > highestY {
			highestY = coordsSlice[1]
		}
	}
	return []int{highestX, highestY}
}


func getDistanceBetweenCoordinates(coordsA []int, coordsB []int) (distance int){
	// May seem a bit silly, but the other way to do this would be to use math.Abs,
	// which expects and returns float64, so conversion to and from would create more
	// overhead, as well as sticking to a heavier type.
	distanceX := 0
	distanceY := 0
	if coordsA[0] > coordsB[0] {
		distanceX = coordsA[0] - coordsB[0]
	} else {
		distanceX = coordsB[0] - coordsA[0]
	}
	if coordsA[1] > coordsB[1] {
		distanceY = coordsA[1] - coordsB[1]
	} else {
		distanceY = coordsB[1] - coordsA[1]
	}
	return distanceX + distanceY
}

func findClosest(coords []int, coordsMap map[string][]int) string {
	var closest = []string{}
	closestDistance := -1
	distance := 0
	for id, coordsDanger := range coordsMap {
		distance = getDistanceBetweenCoordinates(coords, coordsDanger)
		if distance < closestDistance || closestDistance == -1 {
			closestDistance = distance
			closest = []string{id}
		} else if distance == closestDistance {
			closest = append(closest, id)
		}
	}
	if len(closest) == 1 {
		return closest[0]
	} else {
		return "NONE"
	}
}

func findDistanceSum(coords []int, coordsMap map[string][]int) int {
	distance := 0
	totalDistance := 0
	for _, coordsDanger := range coordsMap {
		distance = getDistanceBetweenCoordinates(coords, coordsDanger)
		totalDistance += distance
		if totalDistance > 10050 {
			break
		}
	}
	return totalDistance
}

func mapClosest(xyLimit []int, coordsMap map[string][]int) (areaPMap map[string]int, infiniteAreas []string) {
	var areaPointMap = map[string]int{}
	id := ""
	var lines = [][]string{}
	line := []string{}
	mod := 0
	for y := -mod; y < xyLimit[1] +mod; y++ {
		line = []string{}
		for x := -mod; x < xyLimit[0] + mod; x++ {
			id = findClosest([]int{x,y}, coordsMap)
			line = append(line, fmt.Sprintf("[%s]", id))
			if _, ok := areaPointMap[id];!ok {
				areaPointMap[id] = 1
			} else {
				areaPointMap[id]++
			}
			if x == 0 || x == xyLimit[0] || y == 0 || y == xyLimit[1] {
				infiniteAreas = append(infiniteAreas, id)
			}
		}
		lines = append(lines, line)
	}
	return areaPointMap, infiniteAreas
}

func countSafeSpots(xyLimit []int, coordsMap map[string][]int) int {
	// How much to extend the whole map by in all directions.
	mod := 0
	safeSpotCount := 0
	for y := -mod; y < xyLimit[1]+mod; y++ {
		for x := -mod; x < xyLimit[0]+mod; x++ {
			if findDistanceSum([]int{x, y}, coordsMap) < 10000 {
				safeSpotCount++
			}
		}
	}
	return safeSpotCount
}

func main() {
	startTime := time.Now()

	coordsSlices := loadCoordinatesSlice()
	fmt.Println(coordsSlices)
	fmt.Println(findHighestXY(coordsSlices))
	coordsIdMap := coordsSliceMapWithIds(coordsSlices)
	xyLimit := findHighestXY(coordsSlices)
	res, infiniteAreas := mapClosest(xyLimit, coordsIdMap)
	fmt.Println(res)
	biggestAreaSize := 0
	biggestId := ""
	for id, areaSize := range res {
		if areaSize > biggestAreaSize && !advent_of_code.StringSliceContains(infiniteAreas, id) &&  id != "NONE" {
			biggestAreaSize = areaSize
			biggestId = id
		}
	}

	// Part 1
	fmt.Printf("\n\nBiggest size: %d  (id: %s)", biggestAreaSize, biggestId)

	//Part 2
	safeSpots := countSafeSpots(xyLimit, coordsIdMap)
	fmt.Printf("\n\nSafe Spot count: %d", safeSpots)

	duration := time.Since(startTime)
	fmt.Println("\n\nScript took ", duration)
}
