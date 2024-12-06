package pathfinder

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type findPathSuite struct {
	suite.Suite
}

func TestFindPathSuite(t *testing.T) {
	suite.Run(t, &findPathSuite{})
}

var testinput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func (s *findPathSuite) TestFindPath() {
	tests := []struct {
		name     string
		data     string
		expected int
	}{
		{
			name:     "test 1",
			data:     testinput,
			expected: 41,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			// create a bufio.Scanner from the test input
			p := New(bufio.NewScanner(strings.NewReader(tt.data)), true)
			s.Equal(tt.expected, p.FindPath())
		})
	}
}
