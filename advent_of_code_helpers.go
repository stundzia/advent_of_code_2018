package advent_of_code

import (
	"io/ioutil"
	"fmt"
	"log"
	"path/filepath"
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
