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

	instructions := strings.Split(strings.Trim(string(puzzle), "\n"), "\n")

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(instructions), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(instructions), time.Since(start))
	}
}

func solvePart1(instructions []string) int {
	status := make(map[utils.Point2D]bool)

	for _, instruction := range instructions {
		fields := strings.Fields(instruction)

		var action string
		var from, to []string

		switch fields[0] {
		case "turn":
			from = strings.SplitN(fields[2], ",", 2)
			to = strings.SplitN(fields[4], ",", 2)
			action = fields[1]
		case "toggle":
			from = strings.SplitN(fields[1], ",", 2)
			to = strings.SplitN(fields[3], ",", 2)
			action = fields[0]
		default:
			panic("what?")
		}

		for y := utils.ValOrPanic(strconv.Atoi(from[1])); y <= utils.ValOrPanic(strconv.Atoi(to[1])); y++ {
			for x := utils.ValOrPanic(strconv.Atoi(from[0])); x <= utils.ValOrPanic(strconv.Atoi(to[0])); x++ {
				switch action {
				case "on":
					status[utils.Point2D{X: x, Y: y}] = true
				case "off":
					status[utils.Point2D{X: x, Y: y}] = false
				case "toggle":
					status[utils.Point2D{X: x, Y: y}] = !status[utils.Point2D{X: x, Y: y}]
				}
			}
		}
	}

	res := 0
	for _, point := range status {
		if point {
			res++
		}
	}

	return res
}

func solvePart2(instructions []string) any {
	status := make(map[utils.Point2D]int)

	for _, instruction := range instructions {
		fields := strings.Fields(instruction)

		var action string
		var from, to []string

		switch fields[0] {
		case "turn":
			from = strings.SplitN(fields[2], ",", 2)
			to = strings.SplitN(fields[4], ",", 2)
			action = fields[1]
		case "toggle":
			from = strings.SplitN(fields[1], ",", 2)
			to = strings.SplitN(fields[3], ",", 2)
			action = fields[0]
		default:
			panic("what?")
		}

		for y := utils.ValOrPanic(strconv.Atoi(from[1])); y <= utils.ValOrPanic(strconv.Atoi(to[1])); y++ {
			for x := utils.ValOrPanic(strconv.Atoi(from[0])); x <= utils.ValOrPanic(strconv.Atoi(to[0])); x++ {
				switch action {
				case "on":
					status[utils.Point2D{X: x, Y: y}]++
				case "off":
					status[utils.Point2D{X: x, Y: y}] = max(0, status[utils.Point2D{X: x, Y: y}]-1)
				case "toggle":
					status[utils.Point2D{X: x, Y: y}] += 2
				}
			}
		}
	}

	res := 0
	for _, point := range status {
		res += point
	}

	return res
}
