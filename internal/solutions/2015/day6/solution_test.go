package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_2015_Day6_Part1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "turn on 0,0 through 999,999",
			expected: 1000000,
		},
		{
			input:    "toggle 0,0 through 999,0",
			expected: 1000,
		},
		{
			input:    "turn off 499,499 through 500,500",
			expected: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart1([]string{tc.input}))
		})
	}
}

func TestSolve_2015_Day6_Part2(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "turn on 0,0 through 0,0",
			expected: 1,
		},
		{
			input:    "toggle 0,0 through 999,999",
			expected: 2000000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart2([]string{tc.input}))
		})
	}
}
