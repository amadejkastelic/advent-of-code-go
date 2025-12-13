package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	inputPath = flag.String("input", "", "Path to input file")
	part      = flag.Int("part", 0, "Part of the puzzle to solve (1 or 2). If 0, solve both parts.")
)

func main() {
	flag.Parse()

	if *inputPath == "" {
		panic("Input file path is required")
	}

	puzzle, err := os.ReadFile(*inputPath)
	if err != nil {
		panic(err)
	}

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(string(puzzle)), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(string(puzzle)), time.Since(start))
	}
}

func solvePart1(puzzle string) any {
	res := 0

	for line := range strings.SplitSeq(strings.Trim(puzzle, "\n"), "\n") {
		for _, c := range line {
			if c == '(' {
				res++
			} else {
				res--
			}
		}
	}

	return res
}

func solvePart2(puzzle string) any {
	pos := 0

	for line := range strings.SplitSeq(strings.Trim(puzzle, "\n"), "\n") {
		for i, c := range line {
			if c == '(' {
				pos++
			} else {
				pos--
			}

			if pos == -1 {
				return i + 1
			}
		}
	}

	return 0
}
