package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const puzzle = `987654321111111
811111111111119
234234234234278
818181911112111
`

func TestSolve_2025_Day3(t *testing.T) {
	assert.Equal(t, int64(357), solvePart1(puzzle))
	assert.Equal(t, int64(3121910778619), solvePart2(puzzle))
}
