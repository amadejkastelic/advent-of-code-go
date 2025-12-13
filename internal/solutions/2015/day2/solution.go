package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
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

	presents := sliceutils.Map(strings.Split(strings.Trim(string(puzzle), "\n"), "\n"), PresentFromString)

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(presents), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(presents), time.Since(start))
	}
}

type Present struct {
	l, w, h int64
}

func PresentFromString(s string) *Present {
	dimensions := strings.SplitN(s, "x", 3)
	return &Present{
		l: utils.ValOrPanic(strconv.ParseInt(dimensions[0], 10, 64)),
		w: utils.ValOrPanic(strconv.ParseInt(dimensions[1], 10, 64)),
		h: utils.ValOrPanic(strconv.ParseInt(dimensions[2], 10, 64)),
	}
}

func (p *Present) RequiredWrappingPaper() int64 {
	lw := p.l * p.w
	wh := p.w * p.h
	hl := p.h * p.l
	return 2*lw + 2*wh + 2*hl + min(lw, wh, hl)
}

func (p *Present) RequiredRibbon() int64 {
	dims := []int64{p.l, p.w, p.h}
	slices.Sort(dims)

	return 2*dims[0] + 2*dims[1] + p.l*p.w*p.h
}

func (p *Present) String() string {
	return fmt.Sprintf("%dx%dx%d", p.l, p.w, p.h)
}

func solvePart1(presents []*Present) int64 {
	res := int64(0)
	for _, present := range presents {
		res += present.RequiredWrappingPaper()
	}
	return res
}

func solvePart2(presents []*Present) any {
	res := int64(0)
	for _, present := range presents {
		res += present.RequiredRibbon()
	}
	return res
}
