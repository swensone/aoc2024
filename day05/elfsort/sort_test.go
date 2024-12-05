package elfsort

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type elfSortTestSuite struct {
	suite.Suite
}

func TestElfSortTestSuite(t *testing.T) {
	suite.Run(t, &elfSortTestSuite{})
}

func (s *elfSortTestSuite) TestSort() {
	tests := []struct {
		name     string
		data     []string
		expected []string
	}{
		{
			name:     "test 1",
			data:     []string{"75", "47", "61", "53", "29"},
			expected: []string{"75", "47", "61", "53", "29"},
		},
		{
			name:     "test 2",
			data:     []string{"97", "61", "53", "29", "13"},
			expected: []string{"97", "61", "53", "29", "13"},
		},
		{
			name:     "test 3",
			data:     []string{"75", "29", "13"},
			expected: []string{"75", "29", "13"},
		},
		{
			name:     "test 4",
			data:     []string{"75", "97", "47", "61", "53"},
			expected: []string{"97", "75", "47", "61", "53"},
		},
		{
			name:     "test 5",
			data:     []string{"61", "13", "29"},
			expected: []string{"61", "29", "13"},
		},
		{
			name:     "test 6",
			data:     []string{"97", "13", "75", "29", "47"},
			expected: []string{"97", "75", "47", "29", "13"},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			es := initElfSort()
			es.Sort(tt.data)
			s.Equal(tt.expected, tt.data)
		})
	}
}

func (s *elfSortTestSuite) TestIsSorted() {
	tests := []struct {
		name     string
		data     []string
		expected bool
	}{
		{
			name:     "test 1",
			data:     []string{"75", "47", "61", "53", "29"},
			expected: true,
		},
		{
			name:     "test 2",
			data:     []string{"97", "61", "53", "29", "13"},
			expected: true,
		},
		{
			name:     "test 3",
			data:     []string{"75", "29", "13"},
			expected: true,
		},
		{
			name:     "test 4",
			data:     []string{"75", "97", "47", "61", "53"},
			expected: false,
		},
		{
			name:     "test 5",
			data:     []string{"61", "13", "29"},
			expected: false,
		},
		{
			name:     "test 6",
			data:     []string{"97", "13", "75", "29", "47"},
			expected: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			es := initElfSort()
			actual := es.IsSorted(tt.data)
			s.Equal(tt.expected, actual)
		})
	}
}

func initElfSort() *ElfSort {
	es := New()
	es.AddRule("47|53")
	es.AddRule("97|13")
	es.AddRule("97|61")
	es.AddRule("97|47")
	es.AddRule("75|29")
	es.AddRule("61|13")
	es.AddRule("75|53")
	es.AddRule("29|13")
	es.AddRule("97|29")
	es.AddRule("53|29")
	es.AddRule("61|53")
	es.AddRule("97|53")
	es.AddRule("61|29")
	es.AddRule("47|13")
	es.AddRule("75|47")
	es.AddRule("97|75")
	es.AddRule("47|61")
	es.AddRule("75|61")
	es.AddRule("47|29")
	es.AddRule("75|13")
	es.AddRule("53|13")
	return es
}
