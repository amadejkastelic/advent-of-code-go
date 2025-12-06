package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const puzzle = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `

func TestSolve_2025_Day3(t *testing.T) {
	assert.Equal(t, int64(4277556), solvePart1(puzzle))
	assert.Equal(t, int64(3263827), solvePart2(puzzle))
}

func TestSolve_2025_Day3_EdgeCases(t *testing.T) {
	assert.Equal(t, int64(3253600), solvePart2(` 51
387
215
*  `))
	assert.Equal(t, int64(8544), solvePart2(`123
 45
  6
*  `))
	assert.Equal(t, int64(625+3253600), solvePart2(`328  51
64  387
98  215
+   *  `))
	assert.Equal(t, int64(6073), solvePart2(`   5
   8
 899
2344
+   `))
}

func TestSolve_2025_Day3_EdgeCases_2(t *testing.T) {
	assert.Equal(t, int64(824051026), solvePart2(`133 64  362
876 329 644
82  317 581
87  875   3
*   +   *  `))
}
