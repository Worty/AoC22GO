package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var trees [][]int
var size int

func main() {
	file := "input.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var thisrow []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			thisrow = append(thisrow, num)
		}
		trees = append(trees, thisrow)
	}
	size = len(trees)
	part1and2()
}

func part1and2() {
	visible := ((size - 2) * 4) + 4
	maxScore := 0
	fmt.Printf("Read %dx%d trees, visible outside: %d\n", size, size, visible)
	for row := 1; row < size-1; row++ {
		for col := 1; col < size-1; col++ {
			visibleLeft, distanceLeft := checkToLeft(row, col)
			visibleTop, distanceTop := checkToTop(row, col)
			visibleBottom, distanceBottom := checkToBottom(row, col)
			visibleRight, distanceRight := checkToRight(row, col)
			if visibleLeft || visibleTop || visibleBottom || visibleRight {
				visible++
			}
			score := distanceLeft * distanceTop * distanceBottom * distanceRight
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Printf("Part 1 Visible: %d\n", visible)
	fmt.Printf("Part 2 Score is: %d\n", maxScore)
}

func checkToLeft(row, col int) (visible bool, distance int) {
	val := trees[row][col]
	for i := col - 1; i >= 0; i-- {
		distance++
		if trees[row][i] >= val {
			return false, distance
		}
	}
	return true, distance
}

func checkToRight(row, col int) (visible bool, distance int) {
	val := trees[row][col]
	for i := col + 1; i < len(trees[0]); i++ {
		distance++
		if trees[row][i] >= val {
			return false, distance
		}
	}
	return true, distance
}

func checkToTop(row, col int) (visible bool, distance int) {
	val := trees[row][col]
	for i := row - 1; i >= 0; i-- {
		distance++
		if trees[i][col] >= val {
			return false, distance
		}
	}
	return true, distance
}

func checkToBottom(row, col int) (visible bool, distance int) {
	val := trees[row][col]
	for i := row + 1; i < len(trees); i++ {
		distance++
		if trees[i][col] >= val {
			return false, distance
		}
	}
	return true, distance
}
