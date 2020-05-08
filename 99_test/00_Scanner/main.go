package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readLines("./capitals.txt")
}

func readLines(path string) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
