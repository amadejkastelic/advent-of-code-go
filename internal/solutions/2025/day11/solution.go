package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/amadejkastelic/advent-of-code-go/internal/graph"
)

var (
	inputPath = flag.String("input", "", "Path to input file")
	part      = flag.Int("part", 0, "Part of the puzzle to solve (1 or 2). If 0, solve both parts.")
)

func GraphFromString(input string) *graph.Graph {
	g := graph.NewGraph()

	for line := range strings.SplitSeq(strings.Trim(input, "\n"), "\n") {
		parts := strings.Fields(line)
		for _, part := range parts[1:] {
			g.AddEdge(parts[0][:len(parts[0])-1], part)
		}
	}

	return g
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

	g := GraphFromString(string(puzzle))

	if *part == 1 || *part == 0 {
		startNode, ok := g.GetNode("you")
		if ok {
			start := time.Now()
			fmt.Printf("Part 1: %v, took: %v\n", solvePart1(startNode, "out"), time.Since(start))
		}
	}
	if *part == 2 || *part == 0 {
		startNode, ok := g.GetNode("svr")
		if ok {
			start := time.Now()
			fmt.Printf("Part 2: %v, took: %v\n", solvePart2(startNode, "out"), time.Since(start))
		}
	}
}

func solvePart1(start *graph.Node, target string) int {
	res := 0
	start.TraverseFunc(func(n *graph.Node) {
		if n.ID == target {
			res++
		}
	})
	return res
}

type State struct {
	nodeID string
	hasDac bool
	hasFft bool
}

func solvePart2(start *graph.Node, target string) any {
	return dfs(start, target, false, false, make(map[string]bool), make(map[State]int))
}

func dfs(node *graph.Node, target string, hasDac, hasFft bool, visited map[string]bool, cache map[State]int) int {
	if node.ID == "dac" {
		hasDac = true
	}
	if node.ID == "fft" {
		hasFft = true
	}

	state := State{node.ID, hasDac, hasFft}

	if count, found := cache[state]; found {
		return count
	}

	if node.ID == target {
		if hasDac && hasFft {
			return 1
		}
		return 0
	}

	visited[node.ID] = true

	count := 0
	for _, neighbor := range node.Adjacent {
		if !visited[neighbor.ID] {
			count += dfs(neighbor, target, hasDac, hasFft, visited, cache)
		}
	}

	delete(visited, node.ID)

	cache[state] = count
	return count
}
