package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/amadejkastelic/advent-of-code-go/internal/sliceutils"
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

type Sheet struct {
	numbers []int
	op      string
	result  int64
}

func NewSheet(op string) *Sheet {
	return &Sheet{op: op}
}

func (s *Sheet) AddNumber(num int) {
	s.numbers = append(s.numbers, num)
	switch s.op {
	case "*":
		if s.result == 0 {
			s.result = int64(num)
		} else {
			s.result *= int64(num)
		}
	case "+":
		s.result += int64(num)
	}
}

func (s *Sheet) Result() int64 {
	if s.result != 0 {
		return s.result
	}

	initial := 0
	if s.op == "*" {
		initial = 1
	}
	return int64(sliceutils.Reduce(s.numbers, func(a, b int) int {
		if s.op == "*" {
			return a * b
		}
		return a + b
	}, initial))
}

func sum(sheets []*Sheet) int64 {
	res := int64(0)
	for _, sheet := range sheets {
		res += sheet.Result()
	}
	return res
}

func solvePart1(puzzle string) any {
	sheets := make([]*Sheet, 0)

	lines := strings.Split(strings.Trim(puzzle, "\n"), "\n")

	for op := range strings.FieldsSeq(lines[len(lines)-1]) {
		sheets = append(sheets, NewSheet(op))
	}

	for i := len(lines) - 2; i >= 0; i-- {
		for j, num := range strings.Fields(lines[i]) {
			sheets[j].AddNumber(utils.ValOrPanic(strconv.Atoi(num)))
		}
	}

	return sum(sheets)
}

func solvePart2(puzzle string) any {
	sheets := make([]*Sheet, 0)

	lines := strings.Split(strings.Trim(puzzle, "\n"), "\n")

	for op := range strings.FieldsSeq(lines[len(lines)-1]) {
		sheets = append(sheets, NewSheet(op))
	}

	for i := 0; i < len(lines)-1; i++ {
		j := 0
		k := 0
		addedNum := false
		for _, ch := range lines[i] {
			if ch == ' ' && addedNum {
				j++
				k = 0
				addedNum = false
				continue
			} else if ch != ' ' && len(sheets[j].numbers) <= k {
				sheets[j].numbers = append(sheets[j].numbers, int(ch-'0'))
				addedNum = true
			} else if ch != ' ' {
				sheets[j].numbers[k] = sheets[j].numbers[k]*10 + int(ch-'0')
				addedNum = true
			}
			k++
		}
	}

	return sum(sheets)
}
