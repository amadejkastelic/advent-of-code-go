package main

import (
	"flag"
	"fmt"
	"os"
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

	ranges := make([]Range, 0)
	ids := make([]int, 0)
	for line := range strings.SplitSeq(strings.Trim(string(puzzle), "\n"), "\n") {
		if strings.Contains(line, "-") {
			r := strings.SplitN(line, "-", 2)
			ranges = append(ranges, Range{
				start: utils.ValOrPanic(strconv.Atoi(r[0])),
				end:   utils.ValOrPanic(strconv.Atoi(r[1])),
			})
			continue
		}
		if line != "" && line != "\n" {
			ids = append(ids, utils.ValOrPanic(strconv.Atoi(line)))
		}
	}

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(ranges, ids), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(ranges), time.Since(start))
	}
}

type Range struct {
	start int
	end   int
}

func (r *Range) Contains(value int) bool {
	return value >= r.start && value <= r.end
}

func (r *Range) Length() int {
	return r.end - r.start + 1
}

func (r *Range) LengthExclusive(others []Range) int {
	res := 0

	for _, other := range others {
		if other.end < r.start || other.start > r.end {
			continue
		}

		if other.start <= r.start && other.end >= r.start {
			if other.end >= r.end {
				return 0
			}
			r.start = other.end + 1
			continue
		}

		if other.start <= r.end && other.end >= r.end {
			if other.start <= r.start {
				return 0
			}
			r.end = other.start - 1
			continue
		}

		if other.start > r.start && other.end < r.end {
			res += other.start - r.start
			r.start = other.end + 1
		}
	}

	return r.Length()
}

func solvePart1(ranges []Range, ids []int) int {
	res := 0

	for _, id := range ids {
		for _, r := range ranges {
			if r.Contains(id) {
				res++
				break
			}
		}
	}

	return res
}

func solvePart2(ranges []Range) int {
	res := 0

	for i, r := range ranges {
		res += r.LengthExclusive(ranges[:i])
	}

	return res
}
