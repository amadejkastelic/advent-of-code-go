package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_2015_Day3_Part1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    ">",
			expected: 2,
		},
		{
			input:    "^>v<",
			expected: 4,
		},
		{
			input:    "^v^v^v^v^v",
			expected: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart1(tc.input))
		})
	}
}

func TestSolve_2015_Day3_Part2(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "^v",
			expected: 3,
		},
		{
			input:    "^>v<",
			expected: 3,
		},
		{
			input:    "^v^v^v^v^v",
			expected: 11,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart2(tc.input))
		})
	}
}
