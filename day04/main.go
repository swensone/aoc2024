package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
	ws := NewWordSearch("XMAS", scanner)

	fmt.Println("finding words")
	count := ws.FindWords()
	fmt.Printf("found: %d\n", count)
}

type WordSearch struct {
	Word    string
	Letters map[string]int
	Index   map[string]bool
	Columns int
	Lines   int
}

func NewWordSearch(word string, scanner *bufio.Scanner) *WordSearch {
	letters := make(map[string]int)
	for i, l := range word {
		letters[string(l)] = i
	}
	ws := &WordSearch{Word: word, Letters: letters}
	ws.BuildIndex(scanner)
	return ws
}

func (ws *WordSearch) BuildIndex(scanner *bufio.Scanner) {
	ws.Index = map[string]bool{}

	ws.Lines = 0
	for scanner.Scan() {
		line := scanner.Text()
		if ws.Columns == 0 {
			ws.Columns = len(line)
		}

		for columnNum, l := range line {
			ws.Index[indexKey(byte(l), ws.Lines, columnNum)] = true
		}

		ws.Lines++
	}
}

func (ws *WordSearch) FindWords() int {
	found := 0
	for i := 0; i < ws.Lines; i++ {
		for j := 0; j < ws.Columns; j++ {
			// M . S
			// . A .
			// M . S
			if ws.Index[indexKey('M', i, j)] && ws.Index[indexKey('S', i, j+2)] &&
				ws.Index[indexKey('A', i+1, j+1)] &&
				ws.Index[indexKey('M', i+2, j)] && ws.Index[indexKey('S', i+2, j+2)] {
				found++
			}

			// M . M
			// . A .
			// S . S
			if ws.Index[indexKey('M', i, j)] && ws.Index[indexKey('M', i, j+2)] &&
				ws.Index[indexKey('A', i+1, j+1)] &&
				ws.Index[indexKey('S', i+2, j)] && ws.Index[indexKey('S', i+2, j+2)] {
				found++
			}

			// S . M
			// . A .
			// S . M
			if ws.Index[indexKey('S', i, j)] && ws.Index[indexKey('M', i, j+2)] &&
				ws.Index[indexKey('A', i+1, j+1)] &&
				ws.Index[indexKey('S', i+2, j)] && ws.Index[indexKey('M', i+2, j+2)] {
				found++
			}

			// S . S
			// . A .
			// M . M
			if ws.Index[indexKey('S', i, j)] && ws.Index[indexKey('S', i, j+2)] &&
				ws.Index[indexKey('A', i+1, j+1)] &&
				ws.Index[indexKey('M', i+2, j)] && ws.Index[indexKey('M', i+2, j+2)] {
				found++
			}
		}
	}

	return found
}

func indexKey(letter byte, line, column int) string {
	return fmt.Sprintf("%c:%d:%d", letter, line, column)
}
