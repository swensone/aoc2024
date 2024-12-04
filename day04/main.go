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
			// check horizontal
			if ws.Index[indexKey(ws.Word[0], i, j)] && ws.Index[indexKey(ws.Word[1], i, j+1)] && ws.Index[indexKey(ws.Word[2], i, j+2)] && ws.Index[indexKey(ws.Word[3], i, j+3)] {
				found++
			}

			// check reverse horizontal
			if ws.Index[indexKey(ws.Word[3], i, j)] && ws.Index[indexKey(ws.Word[2], i, j+1)] && ws.Index[indexKey(ws.Word[1], i, j+2)] && ws.Index[indexKey(ws.Word[0], i, j+3)] {
				found++
			}

			// check vertical
			if ws.Index[indexKey(ws.Word[0], i, j)] && ws.Index[indexKey(ws.Word[1], i+1, j)] && ws.Index[indexKey(ws.Word[2], i+2, j)] && ws.Index[indexKey(ws.Word[3], i+3, j)] {
				found++
			}

			// check reverse vertical
			if ws.Index[indexKey(ws.Word[3], i, j)] && ws.Index[indexKey(ws.Word[2], i+1, j)] && ws.Index[indexKey(ws.Word[1], i+2, j)] && ws.Index[indexKey(ws.Word[0], i+3, j)] {
				found++
			}

			// check diagonal down
			if ws.Index[indexKey(ws.Word[0], i, j)] && ws.Index[indexKey(ws.Word[1], i+1, j+1)] && ws.Index[indexKey(ws.Word[2], i+2, j+2)] && ws.Index[indexKey(ws.Word[3], i+3, j+3)] {
				found++
			}

			// check reverse diagonal down
			if ws.Index[indexKey(ws.Word[3], i, j)] && ws.Index[indexKey(ws.Word[2], i+1, j+1)] && ws.Index[indexKey(ws.Word[1], i+2, j+2)] && ws.Index[indexKey(ws.Word[0], i+3, j+3)] {
				found++
			}

			// check diagonal up
			if ws.Index[indexKey(ws.Word[0], i, j)] && ws.Index[indexKey(ws.Word[1], i+1, j-1)] && ws.Index[indexKey(ws.Word[2], i+2, j-2)] && ws.Index[indexKey(ws.Word[3], i+3, j-3)] {
				found++
			}

			// check reverse diagonal up
			if ws.Index[indexKey(ws.Word[3], i, j)] && ws.Index[indexKey(ws.Word[2], i+1, j-1)] && ws.Index[indexKey(ws.Word[1], i+2, j-2)] && ws.Index[indexKey(ws.Word[0], i+3, j-3)] {
				found++
			}
		}
	}

	return found
}

func indexKey(letter byte, line, column int) string {
	return fmt.Sprintf("%c:%d:%d", letter, line, column)
}
