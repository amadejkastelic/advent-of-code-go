package main_test

import (
	"testing"

	main "github.com/amadejkastelic/advent-of-code-go/internal/solutions/2015/day5"
	"github.com/stretchr/testify/assert"
)

func TestIsNice(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "ugknbfddgicrmopn",
			expected: true,
		},
		{
			input:    "aaa",
			expected: true,
		},
		{
			input:    "jchzalrnumimnmhp",
			expected: false,
		},
		{
			input:    "haegwjzuvuyypxyu",
			expected: false,
		},
		{
			input:    "dvszwmarrgswjxmb",
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, main.IsNice(tc.input))
		})
	}
}

func TestIsNiceV2(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "qjhvhtzxzqqjkmpb",
			expected: true,
		},
		{
			input:    "xxyxx",
			expected: true,
		},
		{
			input:    "uurcxstgmygtbstg",
			expected: false,
		},
		{
			input:    "ieodomkazucvgmuy",
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, main.IsNiceV2(tc.input))
		})
	}
}
