package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/amadejkastelic/advent-of-code-go/internal/utils"
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
	s1, s2 := []int{}, []int{}
	for line := range strings.SplitSeq(puzzle, "\n") {
		if line == "" {
			continue
		}

		nums := strings.Split(line, "   ")
		s1 = append(s1, utils.ValOrPanic(strconv.Atoi(nums[0])))
		s2 = append(s2, utils.ValOrPanic(strconv.Atoi(nums[1])))
	}

	slices.Sort(s1)
	slices.Sort(s2)

	result := 0
	for i, num := range s1 {
		result += int(math.Abs(float64(num - s2[i])))
	}

	return result
}

func solvePart2(puzzle string) any {
	s1, m2 := []int{}, map[int]int{}
	for line := range strings.SplitSeq(puzzle, "\n") {
		if line == "" {
			continue
		}

		nums := strings.Split(line, "   ")
		s1 = append(s1, utils.ValOrPanic(strconv.Atoi(nums[0])))
		m2[utils.ValOrPanic(strconv.Atoi(nums[1]))]++
	}

	result := 0
	for _, num := range s1 {
		result += num * m2[num]
	}

	return result
}
