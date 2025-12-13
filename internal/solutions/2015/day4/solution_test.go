package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_2015_Day4_Part1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcdef",
			expected: 609043,
		},
		{
			input:    "pqrstuv",
			expected: 1048970,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart1(tc.input))
		})
	}
}

func TestSolve_2015_Day4_Part2(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcdef",
			expected: 6742839,
		},
		{
			input:    "pqrstuv",
			expected: 5714438,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart2(tc.input))
		})
	}
}
