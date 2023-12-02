package main

import (
	"fmt"
	"log"

	"github.com/glokta1/aoc-2023/cio"
)

func main() {
	lines, err := cio.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		for l, r := 0, len(line)-1; l <= r; l = l + 1 {
			if line[l] >= '0' && line[l] <= '9' {
				fmt.Printf("Found number! %c", line[l])
			}
		}
		fmt.Println()
	}

}
