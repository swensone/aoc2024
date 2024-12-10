package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"

	"github.com/swensone/aoc2024/common/pkg/config"
)

type operation int

const (
	plus = iota
	times
	concat
)

var debug bool

func Loc(x, y int) Location {
	return Location{X: x, Y: y}
}

type Location struct {
	X int
	Y int
}

func (s *Location) Equals(l Location) bool {
	if s.X == l.X && s.Y == l.Y {
		return true
	}
	return false
}

func (s *Location) ToString() string {
	return strconv.Itoa(s.X) + "," + strconv.Itoa(s.Y)
}

func (s *Location) Antinodes(l Location, maxX, maxY int) []Location {
	res := []Location{}

	diffx := s.X - l.X
	diffy := s.Y - l.Y

	antinode1 := Location{
		X: s.X + diffx,
		Y: s.Y + diffy,
	}

	if antinode1.X < maxX && antinode1.X >= 0 && antinode1.Y < maxY && antinode1.Y >= 0 {
		res = append(res, antinode1)
	}

	antinode2 := Location{
		X: l.X + diffx,
		Y: l.Y + diffy,
	}

	if antinode2.X < maxX && antinode2.X >= 0 && antinode2.Y < maxY && antinode2.Y >= 0 {
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

	maxX := 0
	maxY := 0
	scanner := bufio.NewScanner(f)
	antennas := make(map[rune][]Location)
	antmap := map[string]string{}
	for scanner.Scan() {
		text := scanner.Text()
		if maxX == 0 {
			maxX = len(text)
		}

		for i, c := range text {
			if isalphanum(c) {
				if antennas[c] == nil {
					antennas[c] = []Location{}
				}
				l := Loc(i, maxY)
				antmap[l.ToString()] = string(c)
				antennas[c] = append(antennas[c], l)
			}
		}
		maxY++
	}

	antinodes := map[string]bool{}
	for _, locations := range antennas {
		for _, l1 := range locations {
			for _, l2 := range locations {
				if !l1.Equals(l2) {
					ans := l1.Antinodes(l2, maxX, maxY)
					for _, an := range ans {
						antinodes[an.ToString()] = true
					}
				}
			}
		}
	}
	fmt.Printf("maxX: %d, maxY: %d\n", maxX, maxY)
	val := printMap(antmap, antinodes, maxX, maxY)
	fmt.Printf("print res: %d\n", val)

	res := 0
	for y := range maxY {
		for x := range maxX {
			l := Loc(x, y)
			if antinodes[l.ToString()] {
				res++
			}
		}
	}
	
	fmt.Printf("result: %d\n", res)
}

func dprint(f string, a ...any) {
	if debug {
		fmt.Printf(f, a...)
	}
}

func isalphanum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func printMap(antennas map[string]string, antinodes map[string]bool, maxX, maxY int) int {
	res := 0
	for y := range maxY {
		for x := range maxX {
			l := Loc(x, y)
			ant, ok := antennas[l.ToString()]
			if ok {
				print(ant)
				continue
			}
			if antinodes[l.ToString()] {
				res++
				print("#")
				continue
			}
			print(".")
		}
		fmt.Println()
	}
	return res
}