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

	nextidx, containsTokens := nextIndex(text)
	for containsTokens {
		text = text[nextidx:]
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

		nextidx, containsTokens = nextIndex(text)
	}

	fmt.Printf("sum: %d\n", sum)
	fmt.Printf("filteredsum: %d\n", filteredsum)
}

func nextIndex(text string) (int, bool) {
	nextIndex := len(text)
	mulidx := strings.Index(text, mulToken)
	doidx := strings.Index(text, doToken)
	dontidx := strings.Index(text, dontToken)
	contains := false
	if mulidx != -1 {
		nextIndex = min(nextIndex, mulidx)
		contains = true
	}
	if doidx != -1 {
		nextIndex = min(nextIndex, doidx)
		contains = true
	}
	if dontidx != -1 {
		nextIndex = min(nextIndex, dontidx)
		contains = true
	}

	return nextIndex, contains
}

func handleMulToken(text string) (int, error) {
	num1 := text[:strings.Index(text, ",")]
	temptext := text[strings.Index(text, ",")+1:]
	num2 := temptext[:strings.Index(temptext, ")")]

	v1, err := strconv.Atoi(num1)
	if err != nil {
		return 0, err
	}

	v2, err := strconv.Atoi(num2)
	if err != nil {
		return 0, err
	}

	return v1 * v2, nil
}
