package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_2015_Day8_Part1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    `""`,
			expected: 2,
		},
		{
			input:    `"abc"`,
			expected: 2,
		},
		{
			input:    `"aaa\"aaa"`,
			expected: 3,
		},
		{
			input:    `"\x27"`,
			expected: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart1([]string{tc.input}))
		})
	}
}

func TestSolve_2015_Day8_Part2(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    `""`,
			expected: 4,
		},
		{
			input:    `"abc"`,
			expected: 4,
		},
		{
			input:    `"aaa\"aaa"`,
			expected: 6,
		},
		{
			input:    `"\x27"`,
			expected: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart2([]string{tc.input}))
		})
	}
}
