package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const puzzle = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestSolve_2025_Day1(t *testing.T) {
	assert.Equal(t, 3, solvePart1(puzzle))
	assert.Equal(t, 6, solvePart2(puzzle))
}
