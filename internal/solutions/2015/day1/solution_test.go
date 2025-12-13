package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_2015_Day1_Part1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "(())",
			expected: 0,
		},
		{
			input:    "()()",
			expected: 0,
		},
		{
			input:    "(((",
			expected: 3,
		},
		{
			input:    "(()(()(",
			expected: 3,
		},
		{
			input:    "))(((((",
			expected: 3,
		},
		{
			input:    "())",
			expected: -1,
		},
		{
			input:    "))(",
			expected: -1,
		},
		{
			input:    ")))",
			expected: -3,
		},
		{
			input:    ")())())",
			expected: -3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart1(tc.input))
		})
	}
}

func TestSolve_2015_Day1_Part2(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    ")",
			expected: 1,
		},
		{
			input:    "()())",
			expected: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart2(tc.input))
		})
	}
}
