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

	inp := inputTo2DRuneArr(string(puzzle))

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(inp), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(inp), time.Since(start))
	}
}

func inputTo2DRuneArr(inp string) [][]rune {
	res := make([][]rune, 0)

	for line := range strings.SplitSeq(strings.Trim(inp, "\n"), "\n") {
		res = append(res, []rune(line))
	}

	return res
}

const paperRoll = '@'

type Point struct {
	x int
	y int
}

var directions = []Point{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func canAccessRoll(area [][]rune, x int, y int) bool {
	if area[y][x] != paperRoll {
		return false
	}

	cnt := 0
	for _, point := range directions {
		if x+point.x < 0 || x+point.x >= len(area[0]) {
			continue
		}
		if y+point.y < 0 || y+point.y >= len(area) {
			continue
		}

		if area[y+point.y][x+point.x] == paperRoll {
			cnt++
		}
	}

	return cnt < 4
}

func solvePart1(area [][]rune) any {
	res := 0

	for y, line := range area {
		for x := range line {
			if canAccessRoll(area, x, y) {
				res++
			}
		}
	}

	return res
}

func solvePart2(area [][]rune) any {
	res := 0

	oldRes := 0
	for {
		oldRes = res
		for y, line := range area {
			for x := range line {
				if canAccessRoll(area, x, y) {
					area[y][x] = '.'
					res++
				}
			}
		}
		if oldRes == res {
			break
		}
	}

	return res
}
