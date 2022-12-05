package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file := "input.txt"
	part1(file)
	part2(file)
}

func part1(file string) {
	// day3
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	points := 0
	for scanner.Scan() {
		line := scanner.Text()
		firsthalf := line[0 : len(line)/2]
		secondhalf := line[len(line)/2:]
		common := getCommonCharsInStrings(firsthalf, secondhalf)
		sort.Slice(common, func(i, j int) bool {
			return runeToCharCode(common[i]) < runeToCharCode(common[j])
		})
		points += runeToCharCode(common[0]) // lowest char
	}
	fmt.Printf("Part1: %d\n", points)
}

func part2(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	points := 0
	for scanner.Scan() {
		backpacks := []string{}
		for i := 0; i < 3; i++ {
			line := scanner.Text()
			backpacks = append(backpacks, line)
			if i != 2 {
				scanner.Scan()
			}
		}

		firstcom := getCommonCharsInStrings(backpacks[0], backpacks[1])
		common := getCommonCharsInStrings(string(firstcom), backpacks[2])
		x := map[rune]int{}
		for _, c := range common {
			x[c]++
		}
		var biggestrune rune
		var biggestcount int
		for k, v := range x {
			if v > biggestcount {
				biggestcount = v
				biggestrune = k
			}
		}
		points += runeToCharCode(biggestrune)
	}
	fmt.Printf("Part2: %d\n", points)
}
func runeToCharCode(r rune) int {
	if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	}
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	return 0
}

func getCommonCharsInStrings(first string, second string) []rune {
	common := []rune{}
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			if first[i] == second[j] {
				common = append(common, rune(first[i]))
			}
		}
	}
	return common
}
