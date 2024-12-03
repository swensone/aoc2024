#!/usr/bin/env bash

# take the first script argument as the day to init and make sure it is numeric
DAY=$1
if ! [[ $DAY =~ ^[0-9]+$ ]]; then
  echo "usage: $0 <day>"
  exit 1
fi

# create the directory for the day with format "day##"
DAY_DIR="$(printf "day%02d" "${DAY}")"
mkdir -p "${DAY_DIR}"

# create a gitignore in the directory with the executable name
echo "${DAY_DIR}" > "${DAY_DIR}/.gitignore"

# create a boilerplate main.go in the directory
cat <<EOF > "${DAY_DIR}/main.go"
package main

import (
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

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	log.Println(line)
	// }

	// data, err := io.ReadAll(f)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
}
EOF
