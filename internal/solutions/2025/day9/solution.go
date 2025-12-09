package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
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

func polygonEdgeCrossesRectangle(rectP1, rectP2, segP1, segP2 *Point) bool {
	minX, maxX := min(rectP1.X, rectP2.X), max(rectP1.X, rectP2.X)
	minY, maxY := min(rectP1.Y, rectP2.Y), max(rectP1.Y, rectP2.Y)

	if (segP1.X < minX && segP2.X < minX) || (segP1.X > maxX && segP2.X > maxX) {
		return false
	}
	if (segP1.Y < minY && segP2.Y < minY) || (segP1.Y > maxY && segP2.Y > maxY) {
		return false
	}

	onBoundary := func(p *Point) bool {
		return (p.X == minX || p.X == maxX) && p.Y >= minY && p.Y <= maxY ||
			(p.Y == minY || p.Y == maxY) && p.X >= minX && p.X <= maxX
	}

	if onBoundary(segP1) && onBoundary(segP2) {
		if (segP1.X == segP2.X && (segP1.X == minX || segP1.X == maxX)) ||
			(segP1.Y == segP2.Y && (segP1.Y == minY || segP1.Y == maxY)) {
			return false
		}
	}

	rectEdges := [][2]*Point{
		{{X: minX, Y: minY}, {X: maxX, Y: minY}}, // bottom
		{{X: maxX, Y: minY}, {X: maxX, Y: maxY}}, // right
		{{X: maxX, Y: maxY}, {X: minX, Y: maxY}}, // top
		{{X: minX, Y: maxY}, {X: minX, Y: minY}}, // left
	}

	for _, edge := range rectEdges {
		if crosses(segP1, segP2, edge[0], edge[1]) {
			return true
		}
	}

	return false
}

func crosses(p1, p2, p3, p4 *Point) bool {
	ccw := func(a, b, c *Point) int {
		val := (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
		if val > 0 {
			return 1
		}
		if val < 0 {
			return -1
		}
		return 0
	}

	d1 := ccw(p3, p4, p1)
	d2 := ccw(p3, p4, p2)
	d3 := ccw(p1, p2, p3)
	d4 := ccw(p1, p2, p4)

	return ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0))
}

type Pair struct {
	p1, p2 *Point
	area   int
}

func NewPair(p1, p2 *Point) Pair {
	return Pair{p1: p1, p2: p2, area: p1.Area(p2)}
}

func solvePart2(points []*Point, polygon *Polygon) int {
	areas := make([]Pair, 0, len(points)*(len(points)-1)/2)
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			areas = append(areas, NewPair(p1, points[j]))
		}
	}

	slices.SortFunc(areas, func(a, b Pair) int {
		return b.area - a.area
	})

	for _, area := range areas {
		p1, p2 := area.p1, area.p2

		if p1.X == p2.X || p1.Y == p2.Y {
			continue
		}

		corners := []*Point{
			{X: p1.X, Y: p1.Y},
			{X: p1.X, Y: p2.Y},
			{X: p2.X, Y: p1.Y},
			{X: p2.X, Y: p2.Y},
		}

		if !sliceutils.All(corners, polygon.Contains) {
			continue
		}

		valid := true
		for k := 0; k < len(polygon.Vertices); k++ {
			next := (k + 1) % len(polygon.Vertices)
			if polygonEdgeCrossesRectangle(p1, p2, polygon.Vertices[k], polygon.Vertices[next]) {
				valid = false
				break
			}
		}

		if valid {
			return area.area
		}
	}

	return 0
}
