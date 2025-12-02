package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
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

type ProductID struct {
	Value int
}

func validate(s string) bool {
	length := len(s)

	if length%2 != 0 {
		return true
	}

	if s[:length/2] == s[length/2:] {
		return false
	}

	return true
}

func Parse(s string) (*ProductID, error) {
	if !validate(s) {
		return nil, fmt.Errorf("invalid id %s", s)
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &ProductID{Value: val}, nil
}

func validateV2(s string) bool {
	if !validate(s) {
		return false
	}

	if len(s) <= 1 {
		return true
	}

	runes := []rune(s)
	sequence := make([]rune, 0, len(runes))
	left := 0
	for i := range len(runes) {
		c := runes[i]

		if left == 0 && (len(sequence) == 0 || sequence[0] != c) {
			sequence = append(sequence, c)
			continue
		}

		if runes[left] == c && len(sequence) == i-left {
			left = i
		}

		if i == len(runes)-1 && !slices.Equal(sequence, runes[left:]) {
			return true
		}

		if c != runes[i-left] {
			return true
		}
	}

	return len(sequence) == len(runes)
}

func ParseV2(s string) (*ProductID, error) {
	if !validateV2(s) {
		return nil, fmt.Errorf("invalid id %s", s)
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return &ProductID{Value: val}, nil
}

func solvePart1(puzzle string) any {
	result := 0

	for r := range strings.SplitSeq(strings.ReplaceAll(puzzle, "\n", ""), ",") {
		nums := strings.SplitN(r, "-", 2)

		low, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		high, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		for i := low; i <= high; i++ {
			if _, err := Parse(strconv.Itoa(i)); err != nil {
				result += i
			}
		}
	}

	return result
}

func solvePart2(puzzle string) any {
	result := 0

	for r := range strings.SplitSeq(strings.ReplaceAll(puzzle, "\n", ""), ",") {
		nums := strings.SplitN(r, "-", 2)

		low, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		high, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		for i := low; i <= high; i++ {
			if _, err := ParseV2(strconv.Itoa(i)); err != nil {
				result += i
			}
		}
	}

	return result
}
