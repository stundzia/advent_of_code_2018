package main

import (
	"first/advent_of_code"
	"strings"
	"regexp"
	"fmt"
	"time"
)

var lettersUsed = map[rune]bool{}


func loadStepDirections() []string {
	return strings.Split(advent_of_code.LoadInputTxt(7), "\n")
}

func parseStepDirection(stepDirection string) (step rune, requires rune) {
	stepReg, _ := regexp.Compile(`Step ([A-Z]{1}).{29}\s([A-Z]{1})`)
	res := stepReg.FindAllStringSubmatch(stepDirection, -1)
	step_str := res[0][2]
	stepRune := rune(step_str[0])
	requires_str := res[0][1]
	requiresRune := rune(requires_str[0])
	return stepRune, requiresRune
}

func mapStepRequirements(directions []string) map[rune][]rune {
	var requirementsMap = map[rune][]rune{}
	for _, direction := range directions {
		step, requires := parseStepDirection(direction)
		if _, ok := requirementsMap[step];!ok {
			requirementsMap[step] = []rune{requires}
		} else {
			requirementsMap[step] = append(requirementsMap[step], requires)
		}
		if _, ok := lettersUsed[step]; !ok {
			lettersUsed[step] = true
		}
		if _, ok := lettersUsed[requires]; !ok {
			lettersUsed[requires] = true
		}
	}
	return requirementsMap
}


func getStartingLetters(requirementsMap map[rune][]rune) []rune {
	startLetters := []rune{}
	for letter, _ := range lettersUsed {
		if _, ok := requirementsMap[letter]; !ok {
			startLetters = append(startLetters, letter)
		}
	}
	return advent_of_code.SortRuneSlice(startLetters)
}

func getAvailableAction(requirementsMap map[rune][]rune) []rune {
	availableActions := []rune{}
	for step, requirements := range requirementsMap {
		if len(requirements) == 0 {
			availableActions = append(availableActions, step)
		}
	}
	return []rune{advent_of_code.SortRuneSlice(availableActions)[0]}
}

func getAvailableActions(requirementsMap map[rune][]rune) []rune {
	availableActions := []rune{}
	for step, requirements := range requirementsMap {
		if len(requirements) == 0 {
			availableActions = append(availableActions, step)
		}
	}
	return advent_of_code.SortRuneSlice(availableActions)
}

func clearCompletedRequirements(requirementsMap map[rune][]rune, actionsDone []rune) map[rune][]rune {
	for _, action := range actionsDone {
		delete(requirementsMap, action)
		for step := range requirementsMap {
			requirementsMap[step], _ = advent_of_code.DeleteItemFromRuneSlice(requirementsMap[step], action)
			}
	}
	return requirementsMap
}

func getStepTime(step rune) int {
	r := int(step) - 4
	return r
}

func partOne() string {
	directions := loadStepDirections()
	requirementsMap := mapStepRequirements(directions)
	actionOrder := []rune{}
	startingLetters := getStartingLetters(requirementsMap)
	for _, letter := range startingLetters {
		requirementsMap[letter] = []rune{}
	}
	for len(requirementsMap) > 0 {
		actions := getAvailableAction(requirementsMap)
		for _, action := range actions {
			actionOrder = append(actionOrder, action)
		}
		requirementsMap = clearCompletedRequirements(requirementsMap, actions)
	}
	return string(actionOrder)
}

func passSecond(
	workers map[int]int,
	workerTasks map[int]rune,
	actions []rune,
	) (resworkers map[int]int, resworkerTasks map[int]rune, doneTasks []rune) {
		doneTasks = []rune{}
		for worker, timeLeft := range workers {
			if timeLeft <= 1 {
				if workerTasks[worker] != 0 {
					doneTasks = append(doneTasks, workerTasks[worker])
					workerTasks[worker] = 0
				}
				if len(actions) > 0 {
					workerTasks[worker] = actions[0]
					workers[worker] = getStepTime(actions[0])
					actions, _ = advent_of_code.DeleteItemFromRuneSlice(actions, actions[0])
				}
			} else {
				workers[worker]--
			}
		}
	return workers, workerTasks, advent_of_code.SortRuneSlice(doneTasks)
}

func partTwo() int {
	directions := loadStepDirections()
	requirementsMap := mapStepRequirements(directions)
	startingLetters := getStartingLetters(requirementsMap)
	for _, letter := range startingLetters {
		getStepTime(letter)
		requirementsMap[letter] = []rune{}
	}
	var workers = map[int]int{1:0,2:0,3:0,4:0,5:0}
	var workerTasks = map[int]rune{1:0,2:0,3:0,4:0,5:0}
	var doneTasks []rune
	var actionsToDo []rune
	res := 0
	resString := []rune{}

	for len(requirementsMap) > 0 {
		for worker, timeleft := range workers {
			if timeleft == 1 && workerTasks[worker] != 0 {
				clearCompletedRequirements(requirementsMap, []rune{workerTasks[worker]})
			}
		}
		actionsToDo = getAvailableActions(requirementsMap)
		for _, task := range workerTasks {
			actionsToDo, _ = advent_of_code.DeleteItemFromRuneSlice(actionsToDo, task)
		}
		workers, workerTasks, doneTasks = passSecond(workers, workerTasks, actionsToDo)
		res++
		for _, dt := range doneTasks {
			resString = append(resString, dt)
		}
		requirementsMap = clearCompletedRequirements(requirementsMap, doneTasks)
	}

	return res - 1

}

func main() {
	startTime := time.Now()

	fmt.Println("Part 1 result: ", partOne())
	fmt.Println("Part 2 result: ", partTwo())

	duration := time.Since(startTime)
	fmt.Println("\n\nScript took ", duration)
}
