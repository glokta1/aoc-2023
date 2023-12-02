package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/glokta1/aoc-2023/cio"
)

func main() {
	lines, err := cio.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for i, line := range lines {
		cnt, first, last := 0, string(line[0]), string(line[0])
		for _, char := range line {
			if char >= '0' && char <= '9' {
				cnt++
				if cnt == 1 {
					first = string(char)
				}

				last = string(char)
			}
		}

		res, _ := strconv.Atoi(first + last)
		sum += res

		fmt.Println(i, first, last, sum)
	}

	fmt.Println(sum)

}
