package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/swensone/aoc2024/common/pkg/config"
	"github.com/swensone/aoc2024/day05/elfsort"
)

func main() {
	cfg := config.Parse()

	f, err := os.Open(cfg.Input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	es := elfsort.New()
	correctSum, incorrectSum := 0, 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			es.AddRule(line)
		} else if strings.Contains(line, ",") {
			update := strings.Split(line, ",")
			if es.IsSorted(update) {
				midpoint, err := strconv.Atoi(update[len(update)/2])
				if err != nil {
					log.Fatal(err)
				}
				correctSum += midpoint
			} else {
				es.Sort(update)
				midpoint, err := strconv.Atoi(update[len(update)/2])
				if err != nil {
					log.Fatal(err)
				}
				incorrectSum += midpoint
			}
		}
	}
	fmt.Printf("correctly ordered sum:   %d\n", correctSum)
	fmt.Printf("incorrectly ordered sum: %d\n", incorrectSum)
}
