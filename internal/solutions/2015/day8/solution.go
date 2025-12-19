package main

import (
	"flag"
	"fmt"
	"os"
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

	ss := utils.SplitLines(string(puzzle))

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(ss), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(ss), time.Since(start))
	}
}

func solvePart1(ss []string) int {
	res := 0

	for _, s := range ss {
		l := 0

		for i := 1; i < len(s)-1; i++ {
			if s[i] == '\\' {
				switch s[i+1] {
				case '"', '\\':
					i++
					l++
				case 'x':
					i += 3
					l++
				}
			} else {
				l++
			}
		}

		res += len(s) - l
	}

	return res
}

func solvePart2(ss []string) int {
	res := 0

	for _, s := range ss {
		escaped := strings.Builder{}

		for _, c := range s {
			var str string

			switch c {
			case '"':
				str = `\"`
			case '\\':
				str = `\\`
			default:
				str = string(c)
			}

			if _, err := escaped.WriteString(str); err != nil {
				panic(err)
			}
		}

		res += escaped.Len() + 2 - len(s)
	}

	return res
}
