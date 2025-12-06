package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const puzzle = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +
`

func TestSolve_2025_Day3(t *testing.T) {
	assert.Equal(t, int64(4277556), solvePart1(puzzle))
	//assert.Equal(t, int64(3263827), solvePart2(puzzle))
}
