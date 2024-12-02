package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := ""
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.StringVar(&input, "input", "input.txt", "input file")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err.Error())
	}

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	safe := 0
	dampenedSafe := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		numbers := toNumberSlice(parts)
		inputSafe := false
		if isSafe(numbers) {
			inputSafe = true
			safe++
			dampenedSafe++
		}

		dampenedIsSafe := false
		if !inputSafe {
			for i := 0; i < len(numbers) && !dampenedIsSafe; i++ {
				dampened := removeAt(numbers, i)
				if isSafe(dampened) {
					fmt.Printf("vector %v safe after removing %d at index %d, dampened vector %v\n", numbers, numbers[i], i, dampened)
					dampenedIsSafe = true
				}
			}
			if dampenedIsSafe {
				dampenedSafe++
			}
		}
	}

	fmt.Printf("safe: %d\n", safe)
	fmt.Printf("dampened safe: %d\n", dampenedSafe)
}

func removeAt(s []int, i int) []int {
	data := deepClone(s)
	return append(data[:i], data[i+1:]...)
}

func toNumberSlice(s []string) []int {
	var numbers []int

	for _, n := range s {
		num, _ := strconv.Atoi(n)
		numbers = append(numbers, num)
	}

	return numbers
}

func isSafe(s []int) bool {
	data := deepClone(s)
	if slices.IsSorted(data) {
		return varianceSafe(data)
	}
	slices.Reverse(data)
	if slices.IsSorted(data) {
		return varianceSafe(data)
	}
	return false
}

func isSafeDampened(s []int) bool {
	if isSafe(s) {
		return true
	}

	for i := 0; i < len(s); i++ {
		dampened := removeAt(s, i)
		result := isSafe(dampened)
		if result {
			return true
		}
	}

	return false
}

func deepClone(s []int) []int {
	clone := make([]int, len(s))
	copy(clone, s)
	return clone
}

func varianceSafe(s []int) bool {
	minVariance := 1
	maxVariance := 3
	for i := 0; i < len(s)-1; i++ {
		s2 := abs(s[i+1])
		s1 := abs(s[i])
		variance := max(s2, s1) - min(s2, s1)
		if variance < minVariance || variance > maxVariance {
			return false
		}
	}

	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
