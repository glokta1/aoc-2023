package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/glokta1/aoc-2023/cio"
)

var m = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"zero":  "0",
}

func getMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func generateRegexPattern() string {
	pattern := "?=\\d"
	nums := getMapKeys(m)
	for _, numString := range nums {
		pattern += fmt.Sprintf("|%s", numString)
	}

	return pattern
}

func main() {
	patternString := generateRegexPattern()
	lines, err := cio.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for idx, line := range lines {
		fmt.Println(idx, line)
		pattern, err := regexp.Compile(patternString)
		if err != nil {
			log.Fatal(err)
		}

		matches := pattern.FindAllString(line, -1)
		first, last := matches[0], matches[len(matches)-1]
		if val, ok := m[first]; ok {
			first = val
		}
		if val, ok := m[last]; ok {
			last = val
		}

		res, err := strconv.Atoi(first + last)
		if err != nil {
			fmt.Println(err)
		}

		sum += res
		fmt.Println(first, last, res, sum)
	}

	// fmt.Println(sum)
}
