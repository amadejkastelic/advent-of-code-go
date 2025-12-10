package sliceutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amadejkastelic/advent-of-code-go/internal/sliceutils"
)

func TestCombinations(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		k        int
		expected [][]int
	}{
		{
			name:  "Basic case",
			input: []int{1, 2, 3},
			k:     2,
			expected: [][]int{
				{1, 2},
				{1, 3},
				{2, 3},
			},
		},
		{
			name:  "k equals input length",
			input: []int{1, 2, 3},
			k:     3,
			expected: [][]int{
				{1, 2, 3},
			},
		},
		{
			name:     "k is zero",
			input:    []int{1, 2, 3},
			k:        0,
			expected: [][]int{{}},
		},
		{
			name:     "Empty input",
			input:    []int{},
			k:        2,
			expected: [][]int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sliceutils.Combinations(tc.input, tc.k)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestCombinationsWithReplacement(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		k        int
		expected [][]int
	}{
		{
			name:  "Basic case",
			input: []int{1, 2},
			k:     2,
			expected: [][]int{
				{1, 1},
				{1, 2},
				{2, 2},
			},
		},
		{
			name:  "k equals input length",
			input: []int{1, 2},
			k:     2,
			expected: [][]int{
				{1, 1},
				{1, 2},
				{2, 2},
			},
		},
		{
			name:     "k is zero",
			input:    []int{1, 2},
			k:        0,
			expected: [][]int{{}},
		},
		{
			name:     "Empty input",
			input:    []int{},
			k:        2,
			expected: [][]int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := sliceutils.CombinationsWithReplacement(tc.input, tc.k)
			assert.Equal(t, tc.expected, result)
		})
	}
}
