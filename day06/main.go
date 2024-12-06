package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/swensone/aoc2024/common/pkg/config"
	"github.com/swensone/aoc2024/day06/pathfinder"
)

func main() {
	cfg := config.Parse()

	f, err := os.Open(cfg.Input)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	p := pathfinder.New(scanner, cfg.Debug)
	fmt.Println(p.FindPath())
}
