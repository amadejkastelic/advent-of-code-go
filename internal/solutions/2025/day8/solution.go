package main

import (
	"cmp"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

var (
	inputPath = flag.String("input", "", "Path to input file")
	part      = flag.Int("part", 0, "Part of the puzzle to solve (1 or 2). If 0, solve both parts.")
)

func parseInput(input string) []*Point {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	points := make([]*Point, 0, len(lines))
	for _, line := range lines {
		points = append(points, PointFromString(line))
	}
	return points
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

	points := parseInput(string(puzzle))
	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(points, 1000), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(points), time.Since(start))
	}
}

type Point struct {
	X, Y, Z int
}

func PointFromString(s string) *Point {
	var x, y, z int
	if _, err := fmt.Sscanf(s, "%d,%d,%d", &x, &y, &z); err != nil {
		panic(err)
	}
	return &Point{X: x, Y: y, Z: z}
}

func (p *Point) StraightLineDistance(o *Point) int {
	dx := p.X - o.X
	dy := p.Y - o.Y
	dz := p.Z - o.Z
	return dx*dx + dy*dy + dz*dz
}

type Junction struct {
	P1       *Point
	P2       *Point
	Distance int
}

type Circuit struct {
	points []*Point
}

func (c *Circuit) Size() int {
	return len(c.points)
}

func (c *Circuit) AddPoint(p *Point) {
	c.points = append(c.points, p)
}

func solvePart1(points []*Point, numPairs int) any {
	n := len(points)
	junctions := make([]Junction, 0, n*(n-1)/2)
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			junctions = append(junctions, Junction{p1, p2, p1.StraightLineDistance(p2)})
		}
	}

	slices.SortFunc(junctions, func(a, b Junction) int {
		return cmp.Compare(a.Distance, b.Distance)
	})

	allCircuits := make([]*Circuit, 0, n/2)
	circuits := make(map[*Point]*Circuit, n)
	for _, junction := range junctions[:numPairs] {
		p1Circuit, ok1 := circuits[junction.P1]
		p2Circuit, ok2 := circuits[junction.P2]

		if ok1 && ok2 {
			if p1Circuit == p2Circuit {
				continue
			}

			p1Circuit.points = append(p1Circuit.points, p2Circuit.points...)
			for _, p := range p2Circuit.points {
				circuits[p] = p1Circuit
			}
			p2Circuit.points = nil
			continue
		} else if !ok1 && !ok2 {
			circuit := &Circuit{[]*Point{junction.P1, junction.P2}}
			allCircuits = append(allCircuits, circuit)
			circuits[junction.P1] = circuit
			circuits[junction.P2] = circuit
		} else if ok1 {
			circuits[junction.P2] = p1Circuit
			p1Circuit.AddPoint(junction.P2)
		} else if ok2 {
			circuits[junction.P1] = p2Circuit
			p2Circuit.AddPoint(junction.P1)
		}
	}

	slices.SortFunc(allCircuits, func(c1, c2 *Circuit) int {
		return c2.Size() - c1.Size()
	})
	res := 1
	for _, circuit := range allCircuits[:3] {
		res *= circuit.Size()
	}

	return res
}

func solvePart2(points []*Point) any {
	n := len(points)
	junctions := make([]Junction, 0, n*(n-1)/2)
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			junctions = append(junctions, Junction{p1, p2, p1.StraightLineDistance(p2)})
		}
	}

	slices.SortFunc(junctions, func(a, b Junction) int {
		return cmp.Compare(a.Distance, b.Distance)
	})

	circuits := make(map[*Point]*Circuit)
	res := 0
	for _, junction := range junctions {
		p1Circuit, ok1 := circuits[junction.P1]
		p2Circuit, ok2 := circuits[junction.P2]

		if ok1 && ok2 {
			if p1Circuit == p2Circuit {
				continue
			}

			p1Circuit.points = append(p1Circuit.points, p2Circuit.points...)
			for _, p := range p2Circuit.points {
				circuits[p] = p1Circuit
			}
			p2Circuit.points = nil
		} else if !ok1 && !ok2 {
			circuit := &Circuit{[]*Point{junction.P1, junction.P2}}
			circuits[junction.P1] = circuit
			circuits[junction.P2] = circuit
		} else if ok1 {
			circuits[junction.P2] = p1Circuit
			p1Circuit.AddPoint(junction.P2)
		} else if ok2 {
			circuits[junction.P1] = p2Circuit
			p2Circuit.AddPoint(junction.P1)
		}

		res = junction.P1.X * junction.P2.X
	}

	return res
}
