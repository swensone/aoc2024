package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/swensone/aoc2024/common/pkg/config"
)

type operation int

const (
	plus = iota
	times
	concat
)

var debug bool

func Loc(x, y int) *Location {
	return &Location{X: x, Y: y}
}

type Location struct {
	X int
	Y int
}

func (s *Location) Equals(l *Location) bool {
	if s.X == l.X && s.Y == l.Y {
		return true
	}
	return false
}

func (s *Location) ToString() string {
	return strconv.Itoa(s.X) + "," + strconv.Itoa(s.Y)
}

func (s *Location) Antinodes(l *Location, maxX, maxY int) []Location {
	res := []Location{}

	antinode1 := Location{
		X: s.X + l.X,
		Y: s.Y + l.Y,
	}

	if antinode1.X < maxX && antinode1.Y < maxY {
		res = append(res, antinode1)
	}

	antinode2 := Location{
		X: s.X - l.X,
		Y: s.Y - l.Y,
	}

	if antinode2.X >= 0 && antinode2.Y >= 0 {
		res = append(res, antinode2)
	}

	return res
}

func main() {
	cfg := config.Parse()
	debug = cfg.Debug

	f, err := os.Open(cfg.Input)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	res := 0
	maxX := 0
	maxY := 0
	scanner := bufio.NewScanner(f)
	antennas := make(map[rune][]Location)
	for scanner.Scan() {
		text := scanner.Text()
		if maxX == 0 {
			maxX = len(text)
		}

		for i, c := range text {
			if antennas[c] == nil {
				antennas[c] = []Location{}
			}
			antennas[c] = append(antennas[c], Loc{i, maxy})
		}
		maxy++
	}

	antinodes := map[string]bool{}
	for c, locations := range antennas {
		for _, l1 := range locations {
			for _, l2 := range locations {
				if !l1.Equals(l2) {
					ans := l1.Antinodes(l2)
					for _, an := range ans {
						antinodes[ans.ToString()] = true
					}
				}
			}
		}
	}
	
	fmt.Printf("result: %d\n", len(antinodes))
}

func dprint(f string, a ...any) {
	if debug {
		fmt.Printf(f, a...)
	}
}