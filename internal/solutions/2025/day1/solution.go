package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/amadejkastelic/advent-of-code-go/internal/mathutils"
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

type Dial struct {
	Direction int
	Max       int
	Min       int
}

func NewDial(start, max, min int) *Dial {
	return &Dial{
		Direction: start,
		Max:       max,
		Min:       min,
	}
}

func (d *Dial) Rotate(dir string, clicks int) int {
	if dir == "L" {
		start := d.Direction
		d.Direction -= clicks

		if d.Direction < d.Min {
			val := d.Direction
			d.Direction = mathutils.Mod(d.Direction, d.Max)

			res := int(math.Abs(float64(val / d.Max)))
			if start != 0 {
				return res + 1
			}
			return res
		}
	} else {
		d.Direction += clicks
		if d.Direction >= d.Max {
			res := d.Direction / d.Max
			d.Direction %= d.Max
			return res
		}
	}

	if d.Direction == 0 {
		return 1
	}

	return 0
}

func (d *Dial) IsAtZero() bool {
	return d.Direction == 0
}

func solvePart1(puzzle string) any {
	result := 0
	dial := NewDial(50, 100, 0)

	for instruction := range strings.SplitSeq(strings.Trim(puzzle, "\n"), "\n") {
		direction := string(instruction[0])

		clicks, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}

		dial.Rotate(direction, clicks)

		if dial.IsAtZero() {
			result++
		}
	}

	return result
}

func solvePart2(puzzle string) any {
	result := 0
	dial := NewDial(50, 100, 0)

	for instruction := range strings.SplitSeq(strings.Trim(puzzle, "\n"), "\n") {
		direction := string(instruction[0])

		clicks, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}

		result += dial.Rotate(direction, clicks)
	}

	return result
}
