package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var instructions = map[string]string{
	"x": "123",
	"y": "456",
	"d": "x AND y",
	"e": "x OR y",
	"f": "x LSHIFT 2",
	"g": "y RSHIFT 2",
	"h": "NOT x",
	"i": "NOT y",
	"a": "NOT b",
	"b": "d",
}

func TestSolve_2015_Day7(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		assert.Equal(t, 65079, solvePart1(instructions, "i"))
	})
	t.Run("Part 2", func(t *testing.T) {
		assert.Equal(t, 72, solvePart2(instructions))
	})
}
