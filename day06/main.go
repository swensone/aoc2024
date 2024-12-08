package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/swensone/aoc2024/common/pkg/config"
	"github.com/swensone/aoc2024/common/pkg/cslices"
	"github.com/swensone/aoc2024/day06/pathfinder"
)

func main() {
	cfg := config.Parse()

	f, err := os.Open(cfg.Input)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	// part 1: check for how many spaces we visited
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	p := pathfinder.New(bufio.NewScanner(bytes.NewReader(data)), "", cfg.Debug)
	initialPos := fmt.Sprintf("%d,%d", p.PositionX, p.PositionY)
	visited, _ := p.FindPath()
	fmt.Printf("visited: %d\n", len(visited))

	// part 2: look for loops
	visited = cslices.RemoveElement(visited, initialPos)
	loops := 0
	for _, v := range visited {
		fmt.Printf("testing position %s\n", v)
		lp := pathfinder.New(bufio.NewScanner(bytes.NewReader(data)), v, cfg.Debug)
		visited, looped := lp.FindPath()
		if looped {
			loops++
			fmt.Printf("obstacle %s: found loop, loops %d, visited %d\n", v, loops, len(visited))
		} else {
			fmt.Printf("obstacle %s: no loops, loops %d, visited %d\n", v, loops, len(visited))
		}
	}
	fmt.Printf("loops: %d\n", loops)
}
