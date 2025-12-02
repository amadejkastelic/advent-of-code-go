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

func getDigitsIntoArray(n int, digits *[10]int) int {
	if n == 0 {
		digits[0] = 0
		return 1
	}

	length := 0
	temp := n
	for temp > 0 {
		length++
		temp /= 10
	}

	for i := length - 1; i >= 0; i-- {
		digits[i] = n % 10
		n /= 10
	}

	return length
}

func validateInt(n int) bool {
	var digits [10]int
	length := getDigitsIntoArray(n, &digits)

	if length%2 != 0 {
		return true
	}

	half := length / 2
	for i := range half {
		if digits[i] != digits[half+i] {
			return true
		}
	}

	return false
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

func validateV2Int(n int) bool {
	var digitsArray [10]int
	length := getDigitsIntoArray(n, &digitsArray)
	digits := digitsArray[:length]

	if !validateIntFromDigitsSlice(digits) {
		return false
	}

	if length <= 1 {
		return true
	}

	sequence := make([]int, 0, length)
	left := 0
	for i := range length {
		c := digits[i]

		if left == 0 && (len(sequence) == 0 || sequence[0] != c) {
			sequence = append(sequence, c)
			continue
		}

		if digits[left] == c && len(sequence) == i-left {
			left = i
		}

		if i == length-1 && !slices.Equal(sequence, digits[left:]) {
			return true
		}

		if c != digits[i-left] {
			return true
		}
	}

	return len(sequence) == length
}

func validateIntFromDigitsSlice(digits []int) bool {
	length := len(digits)

	if length%2 != 0 {
		return true
	}

	half := length / 2
	for i := range half {
		if digits[i] != digits[half+i] {
			return true
		}
	}

	return false
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
			if !validateInt(i) {
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
			if !validateV2Int(i) {
				result += i
			}
		}
	}

	return result
}
