package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file := "input.txt"
	part1(file)
}

const (
	Red   = "\033[31m"
	White = "\033[0m"
	//offset = 4 // 4 for part 1, 14 for part 2
	offset = 14
)

func part1(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line)-offset; i++ {
			seq := line[i : i+offset]
			if isMarker(seq) {
				fmt.Printf("%s%s%s%s%s after %d\n", line[:i], Red, seq, White, line[i+offset:], i+offset)
				break
			}
		}
	}
}

func isMarker(line string) (isMarker bool) {
	isMarker = true
	for i, char := range line {
		x := strings.IndexRune(line, char)
		if x != i && x != -1 {
			isMarker = false
		}
	}
	return
}
