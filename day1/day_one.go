package main

import (
	"first/advent_of_code"
	"fmt"
	"strings"
	"strconv"
	"time"
)
var frequencies = map[int]bool {0:true}

func applyChange(frequency *int, changeStr string) {
	changeParsed, _ := strconv.ParseInt(changeStr, 10, 32)
	intchangeParsed := int(changeParsed)
	*frequency = *frequency + intchangeParsed
}

func applyChangeReturnDuplicate(frequency *int, changeStr string) int {
	changeParsed, _ := strconv.ParseInt(changeStr, 10, 32)
	intchangeParsed := int(changeParsed)
	*frequency = *frequency + intchangeParsed
	if _, ok := frequencies[*frequency]; !ok {
		frequencies[*frequency] = true
	} else {
		return *frequency
	}
	return 0
}

func loadFrequencyChange() []string {
	res := advent_of_code.LoadInputTxt(1)
	freqChanges := strings.Split(res, "\n")
	return freqChanges
}


func main() {
	startTime := time.Now()
	frequency := 0
	loadFrequencyChange()
	changes := loadFrequencyChange()
	//PART 1
	for _, change := range changes {
		applyChange(&frequency, change)
	}
	fmt.Printf("Frequency after changes result (Part 1): %d\n", frequency)
	//PART 2
	frequency = 0
	res := 0
	iterCount := 0
	for res == 0 {
		iterCount++
		for _, change := range changes {
			res = applyChangeReturnDuplicate(&frequency, change)
			if res != 0 {
				break
			}
		}
	}
	fmt.Printf("Duplicate frequency result (Part 2): %d\n", res)
	fmt.Printf("Iteration count: %d", iterCount)
	duration := time.Since(startTime)
	fmt.Println("\n\nScript took ", duration)
}
