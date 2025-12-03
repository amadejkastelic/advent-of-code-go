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

func sliceToInt64(digits []int) int64 {
	res := int64(0)
	for _, d := range digits {
		res = res*10 + int64(d)
	}
	return res
}

type Bank struct {
	batteries []int
}

func BankFromString(s string) *Bank {
	bats := make([]int, 0, len(s))
	for _, bat := range s {
		bats = append(bats, int(bat-'0'))
	}
	return &Bank{batteries: bats}
}

func (b *Bank) MaxJoltage(n int) int64 {
	joltage := make([]int, n)

	for i, battery := range b.batteries {
		for j, batJol := range joltage {
			if i-j >= len(b.batteries)-n+1 {
				continue
			}

			if battery > batJol {
				joltage[j] = battery
				for k := j + 1; k < len(joltage); k++ {
					joltage[k] = 0
				}
				break
			}
		}
	}

	return sliceToInt64(joltage)
}

func solvePart1(puzzle string) any {
	res := int64(0)
	for bank := range strings.SplitSeq(strings.Trim(puzzle, "\n"), "\n") {
		res += BankFromString(bank).MaxJoltage(2)
	}

	return res
}

func solvePart2(puzzle string) any {
	res := int64(0)
	for bank := range strings.SplitSeq(strings.Trim(puzzle, "\n"), "\n") {
		res += BankFromString(bank).MaxJoltage(12)
	}

	return res
}
