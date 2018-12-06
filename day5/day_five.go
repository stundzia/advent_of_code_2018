package main

import (
	"fmt"
	"time"
	"first/advent_of_code"
)


func getPolymerString() string {
	return advent_of_code.LoadInputTxt(5)
}

func doReaction(polymerElementSlice []rune, polyPairToDrop []rune) (resSlice []rune) {
	var resultSlice = []rune{}
	operationCount := 0
	for i := 0; i < len(polymerElementSlice); i++ {
		if (polymerElementSlice[i] == polyPairToDrop[0]) || (polymerElementSlice[i] == polyPairToDrop[1]) {
			continue
		}
		if i + 1 == len(polymerElementSlice) {
			resultSlice = append(resultSlice, polymerElementSlice[i])
			break
		}
		diff := rune(0)
		if polymerElementSlice[i] > polymerElementSlice[i + 1] {
			diff = polymerElementSlice[i] - polymerElementSlice[i + 1]
		} else {
			diff = polymerElementSlice[i + 1] - polymerElementSlice[i]
		}
		if diff == 32 {
			i++
			operationCount++
			continue
		}
		resultSlice = append(resultSlice, polymerElementSlice[i])
	}
	if operationCount == 0 {
		return polymerElementSlice
	} else {
		return doReaction(resultSlice, polyPairToDrop)
	}
}



func main() {
	startTime := time.Now()

	polyStringSlice := []rune(getPolymerString())
	polyStringSlice = doReaction(polyStringSlice, []rune("00"))
	result1 := len(polyStringSlice)
	fmt.Printf("\nPart 1 result: %d", result1)

	//Part 2
	shortestLenPoly := result1
	shortestLenPolyDrop := ""
	for i := 65; i <= 90; i++ {
		resPolyString := doReaction([]rune(getPolymerString()), []rune{rune(i), rune(i+32)})
		if len(resPolyString) < shortestLenPoly {
			shortestLenPoly = len(resPolyString)
			shortestLenPolyDrop = string([]rune{rune(i), rune(i+32)})
		}
	}
	fmt.Printf("\n\nResult part 2: length: %d  dropped poly: %s", shortestLenPoly, shortestLenPolyDrop)

	duration := time.Since(startTime)
	fmt.Println("\n\nScript took ", duration)
}
