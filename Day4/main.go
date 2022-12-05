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
	part1(file)
}

func part1(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	countfullyoverlap := 0
	countoverlap := 0
	for scanner.Scan() {
		line := scanner.Text()

		a, b := getRange(strings.Split(line, ","))
		if getFullyContaines(a, b) || getFullyContaines(b, a) {
			countfullyoverlap++
		}
		if len(getIntersection(a, b)) > 0 {
			countoverlap++
		}
	}
	fmt.Printf("Fully overlap: %d\n", countfullyoverlap)
	fmt.Printf("Overlap: %d\n", countoverlap)
}

func getFullyContaines(a []int, b []int) bool {
	// bad runtime because a and b are sorted
	tmp := make(map[int]bool)
	for _, v := range a {
		tmp[v] = true
	}
	for _, v := range b {
		if !tmp[v] {
			return false
		}
	}
	return true
}

func getIntersection(a []int, b []int) []int {
	var ret []int
	for _, v := range a {
		for _, v2 := range b {
			if v == v2 {
				ret = append(ret, v)
			} else if v < v2 { // speedup because its sorted
				break
			}
		}
	}
	return ret
}

func getRange(input []string) (retl []int, retr []int) {
	// input: ["51-88","52-87"]
	splitLeft := strings.Split(input[0], "-")
	splitRight := strings.Split(input[1], "-")

	left := convert(splitLeft)
	right := convert(splitRight)
	for i := left[0]; i <= left[1]; i++ {
		retl = append(retl, i)
	}
	for i := right[0]; i <= right[1]; i++ {
		retr = append(retr, i)
	}
	return
}

func convert(input []string) []int {
	var ret []int
	from, _ := strconv.Atoi(input[0])
	to, _ := strconv.Atoi(input[1])

	ret = append(ret, from, to)
	return ret
}
