package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	perr(err)
	defer f.Close()

	var list1 []int
	map2 := map[int]int{}

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
		map2[num2]++
	}

	similarity := 0
	fmt.Println("determining similarity")
	for i := 0; i < len(list1); i++ {
		rightcount := map2[list1[i]]
		similarity += list1[i] * rightcount
	}

	fmt.Println("similarity: ", similarity)
}

func perr(err error) {
	if err != nil {
		panic(err)
	}
}
