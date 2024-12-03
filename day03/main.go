package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/swensone/aoc2024/common/pkg/config"
)

const (
	doToken      = "do()"
	doTokenLen   = 4
	dontToken    = "don't()"
	dontTokenLen = 7
	mulToken     = "mul("
	mulTokenLen  = 4
)

func main() {
	cfg := config.Parse()

	f, err := os.Open(cfg.Input)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	var (
		text             = string(data)
		sum, filteredsum = 0, 0
		enable           = true
	)

	nextToken, containsTokens := nextTokenIndex(text)
	for containsTokens {
		text = text[nextToken:]
		if strings.HasPrefix(text, doToken) {
			text = text[doTokenLen:]
			enable = true
		} else if strings.HasPrefix(text, dontToken) {
			text = text[dontTokenLen:]
			enable = false
		} else {
			text = text[mulTokenLen:]
			val, err := handleMulToken(text)
			if err == nil {
				sum += val
				if enable {
					filteredsum += val
				}
			}
		}

		nextToken, containsTokens = nextTokenIndex(text)
	}

	fmt.Printf("sum: %d\n", sum)
	fmt.Printf("filteredsum: %d\n", filteredsum)
}

func nextTokenIndex(text string) (int, bool) {
	nextIdx := len(text)

	mulIdx := strings.Index(text, mulToken)
	doIdx := strings.Index(text, doToken)
	dontIdx := strings.Index(text, dontToken)
	contains := false

	if mulIdx != -1 {
		nextIdx = min(nextIdx, mulIdx)
		contains = true
	}
	if doIdx != -1 {
		nextIdx = min(nextIdx, doIdx)
		contains = true
	}
	if dontIdx != -1 {
		nextIdx = min(nextIdx, dontIdx)
		contains = true
	}

	return nextIdx, contains
}

func handleMulToken(text string) (int, error) {
	num1End := strings.Index(text, ",")
	num2Start := num1End + 1
	num2End := max(strings.Index(text, ")"), num2Start)

	v1, err := strconv.Atoi(text[:num1End])
	if err != nil {
		return 0, err
	}

	v2, err := strconv.Atoi(text[num2Start:num2End])
	if err != nil {
		return 0, err
	}

	return v1 * v2, nil
}
