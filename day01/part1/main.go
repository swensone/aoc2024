package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	perr(err)
	defer f.Close()

	var list1, list2 []int

	fmt.Println("reading the file contents")
	// read the file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")

		num1, err := strconv.Atoi(numbers[0])
		perr(err)
		num2, err := strconv.Atoi(numbers[1])
		perr(err)

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	fmt.Println("sorting the lists")
	slices.Sort(list1)
	slices.Sort(list2)

	distance := 0
	fmt.Println("determining distance")
	for i := 0; i < len(list1); i++ {
		dist := list1[i] - list2[i]
		if dist < 0 {
			dist = -dist
		}
		distance += dist
	}

	fmt.Println("distance: ", distance)
}

func perr(err error) {
	if err != nil {
		panic(err)
	}
}
