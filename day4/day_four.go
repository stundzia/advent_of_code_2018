package main

import (
	"first/advent_of_code"
	"strings"
	"regexp"
	"fmt"
	"time"
	"sort"
	"strconv"
)

type shiftEntryStruct struct {
	GuardId string
	ActionTime time.Time
	Action string
}

var guardIds = map[string]bool{}
var guardSleepMap = map[string][]int{}
var shiftEntryStructs = map[string][]shiftEntryStruct{}

func parseShiftEntry(shiftEntry string, guardId string) shiftEntryStruct {
	shiftRegex, _ := regexp.Compile(`\[(\d+\-\d+\-\d+\s\d+:\d+)\]\s(\w{5}\s\w{2,6}|\w{5}\s#(\d+))`)
	res := shiftRegex.FindAllStringSubmatch(shiftEntry, -1)
	date := res[0][1]
	action := res[0][2]
	layout := "2006-01-02 15:04"
	t, _ := time.Parse(layout, date)
	if res[0][3] != "" {
		action = "begins shift"
		guardId = res[0][3]
	}
	if _, ok := guardIds[guardId]; !ok {
		guardIds[guardId] = true
	}
	return shiftEntryStruct{
		guardId,
		t,
		action,
	}
}

func parseShift(shiftEntries []shiftEntryStruct) {
	asleep := []int{}
	awake := []int{}
	guardID := ""
	for _, entry := range shiftEntries {
		if guardID == "" {
			guardID = entry.GuardId
		}
		guardID = entry.GuardId
		if entry.Action == "falls asleep" {
			asleep = append(asleep, entry.ActionTime.Minute())
		}
		if entry.Action == "wakes up" {
			awake = append(awake, entry.ActionTime.Minute())
		}
	}
	for i := 0; i < len(asleep); i++ {
		for m := asleep[i]; m < awake[i]; m++ {
			guardSleepMap[guardID] = append(guardSleepMap[guardID], m)
		}
	}
}

func parseGuardEntries(id string) {
	entries := []shiftEntryStruct{}
	for _, entry := range shiftEntryStructs[id] {
		if len(entries) != 0 && entry.Action == "begins shift" {
			parseShift(entries)
			entries = []shiftEntryStruct{}
		}
		entries = append(entries, entry)
	}
}

func makeGuardSleepMap() {
	for gId, _ := range guardIds {
		guardSleepMap[gId] = []int{}
		parseGuardEntries(gId)
	}
}

func getShiftEntrySlice() []string {
	res := strings.Split(advent_of_code.LoadInputTxt(4), "\n")
	sort.Strings(res)
	return res
}

func findBiggestSleepyHead() (guardId string) {
	biggestLen := 0
	biggestSleeperId := ""
	for guardId, sleepMinutes := range guardSleepMap {
		if len(sleepMinutes) > biggestLen {
			biggestLen = len(sleepMinutes)
			biggestSleeperId = guardId
		}
	}
	return biggestSleeperId
}

func findGuardsSleepiestMinute(guardId string) (sleepiestMinute int, count int) {
	var sleepyMinuteMap = map[int]int{}
	for _, minute := range guardSleepMap[guardId] {
		if _, ok := sleepyMinuteMap[minute]; !ok {
			sleepyMinuteMap[minute] = 1
		} else {
			sleepyMinuteMap[minute]++
		}
	}
	mostCommonMinute := 0
	mostCommonMinuteCount := 0
	for minute, count := range sleepyMinuteMap {
		if count > mostCommonMinuteCount {
			mostCommonMinute = minute
			mostCommonMinuteCount = count
		}
	}
	return mostCommonMinute, mostCommonMinuteCount
}


func findSleeperOfSleepiestMinute() (resGuardId int, minute int, count int){
	sleepiestMinute := 0
	sleepiestMinuteCount := 0
	sleepiestMinuteGuardId := ""
	for guardId, _ := range guardIds {
		minute, count := findGuardsSleepiestMinute(guardId)
		if count > sleepiestMinuteCount {
			sleepiestMinute = minute
			sleepiestMinuteCount = count
			sleepiestMinuteGuardId = guardId
		}
	}
	guardIdInt64, _ := strconv.ParseInt(sleepiestMinuteGuardId, 10, 32)
	resGuardId = int(guardIdInt64)
	return resGuardId, sleepiestMinute, sleepiestMinuteCount
}

func main() {
	startTime := time.Now()
	entries := getShiftEntrySlice()
	guardId := "0"
	for _, entry := range entries {
		entryStruct := parseShiftEntry(entry, guardId)
		guardId = entryStruct.GuardId
		if _, ok := shiftEntryStructs[guardId]; !ok {
			shiftEntryStructs[guardId] = []shiftEntryStruct{entryStruct}
		} else {
			shiftEntryStructs[guardId] = append(shiftEntryStructs[guardId], entryStruct)
		}
	}
	makeGuardSleepMap()

	sleepyGuardId := findBiggestSleepyHead()
	sleepiestMinute, _ := findGuardsSleepiestMinute(sleepyGuardId)
	resGuardIdInt, _ := strconv.ParseInt(sleepyGuardId, 10, 32)
	resultPart1 := int(resGuardIdInt) * sleepiestMinute

	//Part 1 result:
	fmt.Printf("\n\nPart 1 result: %d", resultPart1)

	//Part 2
	sleepiestMinuteGuard, sleepiestMinuteOverall, _ := findSleeperOfSleepiestMinute()
	resultPart2 := sleepiestMinuteGuard * sleepiestMinuteOverall
	fmt.Printf("\n\nPart 2 result: %d", resultPart2)

	duration := time.Since(startTime)
	fmt.Println("\n\nScript took ", duration)
}
