package main

import (
	"fmt"
	"strings"
	"unicode"

	aoc "github.com/glokta1/aoc-2023/internal"
)

var m = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

// iterate through string
// if digit, store in a var, continue
// else check if substring starting from current position contains one of the "number words" and store its integer mapping in a var
func findNumbersInString(s string) int {
	cnt := 0
	var first, last int
	for i, ch := range s {
		if unicode.IsDigit(ch) {
			if cnt == 0 {
				// convert rune digit to int value
				// NOTE: plain ch - '0' gives me the rune representation, need to explicitly convert
				first = int(ch - '0')
				cnt++
			}
			last = int(ch - '0')
		}

		for k, v := range m {
			if strings.HasPrefix(s[i:], k) {
				if cnt == 0 {
					first = v
					cnt++
				}

				last = v
			}
		}
	}

	return first*10 + last
}

func main() {
	sum := 0
	for i, line := range aoc.ReadLines("input.txt") {
		fmt.Println(i, line)
		sum += findNumbersInString(line)
	}

	fmt.Println("sum: ", sum)
}
