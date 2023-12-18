package aoc

import (
	"fmt"
	"os"
	"strings"
)

func ReadLines(filepath string) []string {
	fd, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return []string{""}
	}

	lines := strings.Split(string(fd), "\n")

	return lines
}
