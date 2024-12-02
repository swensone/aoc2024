package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/swensone/aoc2024/common/pkg/cmath"
	"github.com/swensone/aoc2024/common/pkg/config"
	"github.com/swensone/aoc2024/common/pkg/cslices"
)

func main() {
	cfg := config.Parse()

	f, err := os.Open(cfg.Input)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	safe := 0
	dampenedSafe := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		numbers, err := cslices.ToIntSlice(strings.Split(scanner.Text(), " "))
		if err != nil {
			log.Fatal(err.Error())
		}

		valuesSafe := false
		if isSafe(numbers) {
			valuesSafe = true
			safe++
		}

		if !valuesSafe && isSafeDampened(numbers) {
			dampenedSafe++
		}
	}

	fmt.Printf("safe: %d\n", safe)
	fmt.Printf("safe with dampening: %d\n", safe+dampenedSafe)
}

func isSafe(s []int) bool {
	if slices.IsSortedFunc(s, func(a, b int) int {
		return a - b
	}) {
		return varianceSafe(s)
	}
	if slices.IsSortedFunc(s, func(a, b int) int {
		return b - a
	}) {
		return varianceSafe(s)
	}
	return false
}

func isSafeDampened(s []int) bool {
	for i := 0; i < len(s); i++ {
		dampened := cslices.RemoveAt(s, i)
		if isSafe(dampened) {
			fmt.Printf("vector %v safe after removing %d at index %d, dampened vector %v\n", s, s[i], i, dampened)
			return true
		}
	}

	return false
}

func varianceSafe(s []int) bool {
	minVariance := 1
	maxVariance := 3
	for i := 0; i < len(s)-1; i++ {
		s2 := cmath.Abs(s[i+1])
		s1 := cmath.Abs(s[i])
		variance := max(s2, s1) - min(s2, s1)
		if variance < minVariance || variance > maxVariance {
			return false
		}
	}

	return true
}
