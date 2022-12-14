package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	smallestColumn = 0
	biggestColumn  = 0
	smallestDepth  = 0
	biggestDepth   = 0
	field          = make(Field, 0)
)

func main() {
	file := "input.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var inputs [][]int
	for scanner.Scan() {
		line := scanner.Text()
		//498,4 -> 498,6 -> 496,6
		in := strings.ReplaceAll(line, " -> ", ",")
		var input []int
		i := 0
		for _, v := range strings.Split(in, ",") {
			num, _ := strconv.Atoi(v)
			input = append(input, num)
			if i%2 == 0 {
				if num > biggestColumn {
					biggestColumn = num
				}
				if num < smallestColumn || smallestColumn == 0 {
					smallestColumn = num
				}
			} else {
				if num > biggestDepth {
					biggestDepth = num
				}
			}
			i++
		}
		fmt.Printf("input: %v\n", input)
		inputs = append(inputs, input)
	}
	for i := 0; i < biggestColumn-smallestColumn+1; i++ {
		field = append(field, make([]Material, biggestDepth-smallestDepth+1))
	}
	PrintField()
	for _, input := range inputs {
		field.generateRocks(input)
		PrintField()
	}
}

type Material int

type Field [][]Material

const (
	AIR  Material = iota
	ROCK Material = iota
	SAND Material = iota
)

func PrintField() {
	fmt.Printf("Field:\n%s\n", field)
}

func (f Field) String() string {
	var s string
	for depth := 0; depth < len(f[0]); depth++ {
		s += fmt.Sprintf("%03d", depth)
		for column := 0; column < len(f); column++ {
			switch f[column][depth] {
			case ROCK:
				s += "#"
			case SAND:
				s += "o"
			case AIR:
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (f Field) generateRocks(input []int) {
	var prevColumn, prevRow int
	for i := 0; i < len(input); i += 2 {
		column := input[i] - smallestColumn
		row := input[i+1] - smallestDepth
		f[column][row] = ROCK
		if column == prevColumn { // nach unten
			diff := row - prevRow
			if diff < 0 {
				for j := diff; j < 0; j++ {
					f[column][prevRow+j] = ROCK
				}
			} else {
				for j := 0; j < diff; j++ {
					f[column][prevRow+j] = ROCK
				}
			}
		} else if row == prevRow { // zur seite
			diff := column - prevColumn
			if diff < 0 {
				for j := diff; j < 0; j++ {
					f[prevColumn+j][row] = ROCK
				}
			} else {
				for j := 0; j < diff; j++ {
					f[prevColumn+j][row] = ROCK
				}
			}
		}
		prevColumn = column
		prevRow = row
	}
}
