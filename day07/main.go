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

func main() {
	cfg := config.Parse()

	f, err := os.Open(cfg.Input)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		total, vals := parseLine(text)
		fmt.Println("total: %d, vals: %+v\n", total, vals)
	}
}

func parseLine(line string) (int, []int) {
	split1 := strings.Split(line, ": ")
	total, _ := strconv.Atoi(split1[0])
	vals := []int{}
	for _, v := range strings.Split(split1[1], " ") {
		val, _ := strconv.Atoi(v)
		vals = append(vals, val)
	}
	return total, vals
}