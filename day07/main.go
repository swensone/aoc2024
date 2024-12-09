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

type operation int

const (
	plus = iota
	times
	concat
)

var debug bool

func main() {
	cfg := config.Parse()
	debug = cfg.Debug

	f, err := os.Open(cfg.Input)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	validcnt := 0
	totalvalid := 0
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		total, vals := parseLine(text)
		dprint("total: %d, vals: %v\n", total, vals)
		cnt := (valid(total, plus, 0, vals) + valid(total, times, 0, vals) + valid(total, concat, 0, vals))
		if cnt > 0 {
			sum += total
			totalvalid += cnt
			validcnt++
			dprint("cnt: %d, validcnt: %d, totalvalid: %d, sum: %d\n", cnt, validcnt, totalvalid, sum)
		}
	}
	fmt.Printf("valid: %d, total valid: %d sum: %d\n", validcnt, totalvalid, sum)
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

func valid(product int, op operation, res int, vals []int) int {
	if op == plus {
		res += vals[0]
	} else if op == times {
		res *= vals[0]
	} else {
		res, _ = strconv.Atoi(strconv.Itoa(res) + strconv.Itoa(vals[0]))
	}

	if len(vals) == 1 {
		if product == res {
			return 1
		}
		return 0
	}

	return valid(product, plus, res, vals[1:]) + valid(product, times, res, vals[1:]) + valid(product, concat, res, vals[1:])
}

func dprint(f string, a ...any) {
	if debug {
		fmt.Printf(f, a...)
	}
}