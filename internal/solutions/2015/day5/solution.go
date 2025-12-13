package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
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

	ss := strings.Split(strings.Trim(string(puzzle), "\n"), "\n")

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(ss), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(ss), time.Since(start))
	}
}

var (
	vowelsRe    = regexp.MustCompile(`[aeiou].*[aeiou].*[aeiou]`)
	doubleRe    = regexp.MustCompile(`aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz`)
	forbiddenRe = regexp.MustCompile(`ab|cd|pq|xy`)
)

func IsNice(s string) bool {
	return len(s) >= 2 &&
		vowelsRe.MatchString(s) &&
		doubleRe.MatchString(s) &&
		!forbiddenRe.MatchString(s)
}

func solvePart1(ss []string) int {
	res := 0

	for _, s := range ss {
		if IsNice(s) {
			res++
		}
	}

	return res
}

func hasXYX(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func hasPairTwice(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			return true
		}
	}
	return false
}

func IsNiceV2(s string) bool {
	return len(s) >= 3 && hasPairTwice(s) && hasXYX(s)
}

func solvePart2(ss []string) int {
	res := 0

	for _, s := range ss {
		if IsNiceV2(s) {
			res++
		}
	}

	return res
}
