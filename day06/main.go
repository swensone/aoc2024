package main

import (
	"bufio"
	"fmt"
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
	scanner := bufio.NewScanner(f)
	p := pathfinder.New(scanner, "", cfg.Debug)
	initialPos := fmt.Sprintf("%d:%d", p.PositionX, p.PositionY)
	visited, _ := p.FindPath()
	fmt.Printf("visited: %d\n", len(visited))

	// part 2: look for loops
	visited = cslices.RemoveElement(visited, initialPos)
	loops := 0
	for _, v := range visited {
		lp := pathfinder.New(scanner, v, cfg.Debug)
		_, looped := lp.FindPath()
		if looped {
			loops++
			fmt.Printf("found loop at %s, loops %d\n", v, loops)
		}
	}
	fmt.Printf("loops: %d\n", loops)
}
