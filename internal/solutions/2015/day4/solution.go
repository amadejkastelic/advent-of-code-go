package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/amadejkastelic/advent-of-code-go/internal/hashutils"
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

	input, err := os.ReadFile(*inputPath)
	if err != nil {
		panic(err)
	}

	password := strings.Trim(string(input), "\n")

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(string(password)), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(string(password)), time.Since(start))
	}
}

func mine(password string, numZeroes int) int {
	wantPrefix := strings.Builder{}
	for range numZeroes {
		wantPrefix.WriteRune('0')
	}
	wantPrefixStr := wantPrefix.String()

	i := 0
	for {
		hash := hashutils.MD5Hash(password + strconv.Itoa(i))
		if strings.HasPrefix(hash, wantPrefixStr) {
			return i
		}
		i++
	}
}

func solvePart1(password string) int {
	return mine(password, 5)
}

func solvePart2(password string) int {
	return mine(password, 6)
}
