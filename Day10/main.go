package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := "input.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	registerX := 1
	clock := 0
	sum := 0
	tick := func() {
		if clock%40-registerX < 2 && clock%40-registerX > -2 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		clock++
		if clock%40 == 0 {
			fmt.Print("\n")
		}
		if clock%40 == 20 {
			sum += registerX * clock
		}
	}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		tick()
		if len(input) > 1 {
			tick()
			operand, _ := strconv.Atoi(input[1])
			registerX += operand
		}
	}
	fmt.Printf("\nSignal: %d\n", sum)
}
