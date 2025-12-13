package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve_2015_Day2_Part1(t *testing.T) {
	testCases := []struct {
		present  *Present
		expected int64
	}{
		{
			present:  &Present{2, 3, 4},
			expected: 58,
		},
		{
			present:  &Present{1, 1, 10},
			expected: 43,
		},
		{
			present:  &Present{10, 1, 1},
			expected: 43,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.present.String(), func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart1([]*Present{tc.present}))
		})
	}
}

func TestSolve_2015_Day2_Part2(t *testing.T) {
	testCases := []struct {
		present  *Present
		expected int64
	}{
		{
			present:  &Present{2, 3, 4},
			expected: 34,
		},
		{
			present:  &Present{1, 1, 10},
			expected: 14,
		},
		{
			present:  &Present{10, 1, 1},
			expected: 14,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.present.String(), func(t *testing.T) {
			assert.Equal(t, tc.expected, solvePart2([]*Present{tc.present}))
		})
	}
}
