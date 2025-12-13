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

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(string(puzzle)), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(string(puzzle)), time.Since(start))
	}
}

var dirMap = map[rune]utils.Direction{
	'^': utils.DirectionUp,
	'v': utils.DirectionDown,
	'<': utils.DirectionLeft,
	'>': utils.DirectionRight,
}

func solvePart1(puzzle string) any {
	cur := utils.Point2D{X: 0, Y: 0}
	visited := map[utils.Point2D]bool{cur: true}

	for _, dir := range strings.Trim(puzzle, "\n") {
		cur = *cur.Move(dirMap[dir])
		visited[cur] = true
	}

	return len(visited)
}

func solvePart2(puzzle string) any {
	santa := utils.Point2D{X: 0, Y: 0}
	robot := utils.Point2D{X: 0, Y: 0}
	visited := map[utils.Point2D]bool{santa: true}

	for i, dir := range strings.Trim(puzzle, "\n") {
		if i%2 == 0 {
			santa = *santa.Move(dirMap[dir])
			visited[santa] = true
		} else {
			robot = *robot.Move(dirMap[dir])
			visited[robot] = true
		}
	}

	return len(visited)
}
