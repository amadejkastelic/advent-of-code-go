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

func parseInput(inp string) ([][]rune, *Point) {
	var startP *Point
	area := make([][]rune, 0)
	for y, line := range strings.Split(strings.Trim(inp, "\n"), "\n") {
		area = append(area, []rune(line))
		if startP == nil {
			for x, ch := range line {
				if ch == 'S' {
					startP = &Point{x, y}
					break
				}
			}
		}
	}
	return area, startP
}

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
		area, startP := parseInput(string(puzzle))
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(area, startP), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		area, startP := parseInput(string(puzzle))
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(area, startP), time.Since(start))
	}
}

var directions = []Point{
	{-1, 1},
	{1, 1},
}

type Point struct {
	X int
	Y int
}

func solvePart1(area [][]rune, start *Point) any {
	result := 0

	q := utils.Queue[*Point]{}

	q.Push(start)

	for q.Len() > 0 {
		curr := q.Pop()

		if curr.Y >= len(area)-1 {
			continue
		}

		switch area[curr.Y+1][curr.X] {
		case '.':
			area[curr.Y+1][curr.X] = '|'
			next := &Point{curr.X, curr.Y + 1}
			q.Push(next)
		case '|':
			continue
		case '^':
			result++
			for _, dir := range directions {
				if dir.X+curr.X < 0 || dir.X+curr.X >= len(area[0]) {
					continue
				}
				if area[curr.Y+1][curr.X+dir.X] == '.' {
					area[curr.Y+1][curr.X+dir.X] = '|'
					next := &Point{curr.X + dir.X, curr.Y + 1}
					q.Push(next)
				}
			}
		}
	}

	return result
}

func solvePart2(area [][]rune, start *Point) any {
	result := 0

	q := utils.Queue[*Point]{}

	q.Push(start)

	for q.Len() > 0 {
		curr := q.Pop()

		if curr.Y >= len(area)-1 {
			continue
		}

		switch area[curr.Y+1][curr.X] {
		case '.':
			next := &Point{curr.X, curr.Y + 1}
			q.Push(next)
		case '^':
			for _, dir := range directions {
				if dir.X+curr.X < 0 || dir.X+curr.X >= len(area[0]) {
					continue
				}
				next := &Point{curr.X + dir.X, curr.Y + 1}
				q.Push(next)
				result++
			}
		}
	}

	return result
}
