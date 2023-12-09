package aoc

import (
	"os"
	"strings"
)

func ReadLines(filepath string) ([]string, error) {
	fd, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(fd), "\n")

	return lines, nil
}
