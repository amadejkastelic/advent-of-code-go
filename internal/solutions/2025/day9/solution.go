package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/amadejkastelic/advent-of-code-go/internal/mathutils"
	"github.com/amadejkastelic/advent-of-code-go/internal/sliceutils"
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

	points := []*Point{}
	for line := range strings.SplitSeq(strings.Trim(string(puzzle), "\n"), "\n") {
		points = append(points, PointFromString(line))
	}

	if *part == 1 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 1: %v, took: %v\n", solvePart1(points), time.Since(start))
	}
	if *part == 2 || *part == 0 {
		start := time.Now()
		fmt.Printf("Part 2: %v, took: %v\n", solvePart2(points, &Polygon{Vertices: points}), time.Since(start))
	}
}

type Point struct {
	X, Y int
}

func (p *Point) Area(other *Point) int {
	if p.X == other.X {
		return mathutils.Abs(p.Y - other.Y)
	}

	if p.Y == other.Y {
		return mathutils.Abs(p.X - other.X)
	}

	return (1 + mathutils.Abs(p.X-other.X)) * (1 + mathutils.Abs(p.Y-other.Y))
}

func PointFromString(s string) *Point {
	var p Point
	if _, err := fmt.Sscanf(s, "%d,%d", &p.X, &p.Y); err != nil {
		panic(err)
	}
	return &p
}

func solvePart1(points []*Point) int {
	maxArea := 0

	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			maxArea = max(maxArea, p1.Area(points[j]))
		}
	}

	return maxArea
}

type Polygon struct {
	Vertices []*Point
}

func (p *Polygon) Contains(point *Point) bool {
	count := 0
	n := len(p.Vertices)

	for i := range n {
		v1 := p.Vertices[i]
		v2 := p.Vertices[(i+1)%n]

		if (v1.Y == v2.Y && point.Y == v1.Y && point.X >= min(v1.X, v2.X) && point.X <= max(v1.X, v2.X)) ||
			(v1.X == v2.X && point.X == v1.X && point.Y >= min(v1.Y, v2.Y) && point.Y <= max(v1.Y, v2.Y)) {
			return true
		}

		if (v1.Y > point.Y) != (v2.Y > point.Y) {
			slope := float64(v2.X-v1.X) / float64(v2.Y-v1.Y)
			xIntersect := slope*float64(point.Y-v1.Y) + float64(v1.X)
			if float64(point.X) < xIntersect {
				count++
			}
		}
	}

	return count%2 == 1
}

func solvePart2(points []*Point, polygon *Polygon) any {
	maxArea := 0

	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]

			if p1.X == p2.X || p1.Y == p2.Y {
				continue
			}

			corners := []*Point{
				{X: p1.X, Y: p1.Y},
				{X: p1.X, Y: p2.Y},
				{X: p2.X, Y: p1.Y},
				{X: p2.X, Y: p2.Y},
			}

			hasEdges := sliceutils.All(corners, polygon.Contains)
			if !hasEdges {
				continue
			}

			edges := []*Point{}
			for x := min(p1.X, p2.X); x <= max(p1.X, p2.X); x++ {
				edges = append(edges, &Point{X: x, Y: p1.Y})
				edges = append(edges, &Point{X: x, Y: p2.Y})
			}
			for y := min(p1.Y, p2.Y); y <= max(p1.Y, p2.Y); y++ {
				edges = append(edges, &Point{X: p1.X, Y: y})
				edges = append(edges, &Point{X: p2.X, Y: y})
			}

			if hasEdges && sliceutils.All(edges, polygon.Contains) {
				maxArea = max(maxArea, p1.Area(p2))
			}
		}
	}

	return maxArea
}
