package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const puzzle = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestSolve_2025_Day3(t *testing.T) {
	inp := inputTo2DRuneArr(puzzle)
	assert.Equal(t, 13, solvePart1(inp))
	assert.Equal(t, 43, solvePart2(inp))
}
