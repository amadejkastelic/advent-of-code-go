package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var points = []*Point{
	{7, 1},
	{11, 1},
	{11, 7},
	{9, 7},
	{9, 5},
	{2, 5},
	{2, 3},
	{7, 3},
}

func TestSolve_2025_Day3(t *testing.T) {
	assert.Equal(t, 50, solvePart1(points))
	assert.Equal(t, 24, solvePart2(points, &Polygon{Vertices: points}))
}
