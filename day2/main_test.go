package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type isSafeTestSuite struct {
	suite.Suite
}

func TestIsSafeTestSuite(t *testing.T) {
	suite.Run(t, &isSafeTestSuite{})
}

func (s *isSafeTestSuite) TestIsSafe() {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "decreasing, variance decreasing by 1 to 2",
			input:    []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			name:     "increasing, variance increasing by greater than 3",
			input:    []int{1, 2, 7, 8, 9},
			expected: false,
		},
		{
			name:     "decreasing, variance decreasing by greater than 3",
			input:    []int{9, 7, 6, 2, 1},
			expected: false,
		},
		{
			name:     "unsorted",
			input:    []int{1, 3, 2, 4, 5},
			expected: false,
		},
		{
			name:     "no decrease between two numbers",
			input:    []int{8, 6, 4, 4, 1},
			expected: false,
		},
		{
			name:     "increasing, variance 1 to 3",
			input:    []int{1, 3, 6, 7, 9},
			expected: true,
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			result := isSafe(test.input)
			fmt.Printf("test: %q, result: %v, correct: %t\n", test.name, result, test.expected == result)
			s.Equal(test.expected, result)
		})
	}
}

func (s *isSafeTestSuite) TestIsSafeDampened() {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "safe without removing any level",
			input:    []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			name:     "unsafe 1 regardless of which level is removed",
			input:    []int{1, 2, 7, 8, 9},
			expected: false,
		},
		{
			name:     "unsafe 2 regardless of which level is removed",
			input:    []int{9, 7, 6, 2, 1},
			expected: false,
		},
		{
			name:     "safe by removing the second level, 3",
			input:    []int{1, 3, 2, 4, 5},
			expected: true,
		},
		{
			name:     "safe by removing the third level, 4",
			input:    []int{8, 6, 4, 4, 1},
			expected: true,
		},
		{
			name:     "safe without removing any level",
			input:    []int{1, 3, 6, 7, 9},
			expected: true,
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			result := isSafeDampened(test.input)
			fmt.Printf("test: %q, result: %v, correct: %t\n", test.name, result, test.expected == result)
			s.Equal(test.expected, result)
		})
	}
}
