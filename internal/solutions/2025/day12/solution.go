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

	presents := make([]int, 0)
	regions := make([]*Region, 0)
	for section := range strings.SplitSeq(strings.Trim(string(puzzle), "\n"), "\n\n") {
		if strings.Contains(section, "x") {
			for r := range strings.SplitSeq(strings.Trim(section, "\n"), "\n") {
				regions = append(regions, RegionFromString(r))
			}
			continue
		}

		present := 0
		for line := range strings.SplitSeq(strings.Trim(section, "\n"), "\n") {
			for _, char := range line {
				if char == '#' {
					present++
				}
			}
		}
		presents = append(presents, present)
	}

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(regions, presents), time.Since(start))
	}
}

type Region struct {
	x, y     int
	presents []int
}

func RegionFromString(s string) *Region {
	parts := strings.Fields(s)

	size := strings.Split(strings.Trim(parts[0], ":"), "x")

	region := Region{
		x:        utils.ValOrPanic(strconv.Atoi(size[0])),
		y:        utils.ValOrPanic(strconv.Atoi(size[1])),
		presents: make([]int, 0, len(parts)-1),
	}

	for _, present := range parts[1:] {
		region.presents = append(region.presents, utils.ValOrPanic(strconv.Atoi(present)))
	}

	return &region
}

func solvePart1(regions []*Region, presents []int) any {
	res := 0

	for _, region := range regions {
		cnt := 0
		for i, p := range region.presents {
			cnt += presents[i] * p
		}

		if region.x*region.y > cnt {
			res++
		}
	}

	return res
}
