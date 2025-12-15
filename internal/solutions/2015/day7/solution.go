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

	instructions := make(map[string]string, 0)
	for _, instruction := range utils.SplitLines(string(puzzle)) {
		lr := strings.SplitN(instruction, " -> ", 2)
		instructions[lr[1]] = lr[0]
	}

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(instructions, "a"), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(instructions), time.Since(start))
	}
}

var memo = map[string]uint16{}

func evaluate(instructions map[string]string, wire string) uint16 {
	if val, ok := memo[wire]; ok {
		return val
	}

	op := strings.Fields(strings.TrimSpace(instructions[wire]))
	if len(op) == 1 {
		if num, err := strconv.ParseUint(op[0], 10, 16); err == nil {
			memo[wire] = uint16(num)
		} else {
			memo[wire] = evaluate(instructions, op[0])
		}
		return memo[wire]
	}

	if len(op) == 2 && strings.HasPrefix(op[0], "NOT") {
		memo[wire] = ^evaluate(instructions, op[1])
		return memo[wire]
	}

	var l, r uint16
	if left, err := strconv.ParseUint(op[0], 10, 16); err == nil {
		l = uint16(left)
	} else {
		l = evaluate(instructions, op[0])
	}
	if right, err := strconv.ParseUint(op[2], 10, 16); err == nil {
		r = uint16(right)
	} else {
		r = evaluate(instructions, op[2])
	}

	switch op[1] {
	case "AND":
		memo[wire] = l & r
	case "OR":
		memo[wire] = l | r
	case "LSHIFT":
		memo[wire] = l << r
	case "RSHIFT":
		memo[wire] = l >> r
	}

	return memo[wire]
}

func solvePart1(instructions map[string]string, wire string) int {
	memo = make(map[string]uint16, 0)
	return int(evaluate(instructions, wire))
}

func solvePart2(instructions map[string]string) any {
	instructions["b"] = strconv.Itoa(solvePart1(instructions, "a"))
	return solvePart1(instructions, "a")
}
