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

// [P]     [C]         [M]
// [D]     [P] [B]     [V] [S]
// [Q] [V] [R] [V]     [G] [B]
// [R] [W] [G] [J]     [T] [M]     [V]
// [V] [Q] [Q] [F] [C] [N] [V]     [W]
// [B] [Z] [Z] [H] [L] [P] [L] [J] [N]
// [H] [D] [L] [D] [W] [R] [R] [P] [C]
// [F] [L] [H] [R] [Z] [J] [J] [D] [D]
//
//	1   2   3   4   5   6   7   8   9
var stacks = []stack{
	{"F", "H", "B", "V", "R", "Q", "D", "D", "P"},
	{"L", "D", "Z", "Q", "W", "V"},
	{"H", "L", "Z", "Q", "G", "R", "P", "C"},
	{"R", "D", "H", "F", "J", "V", "B"},
	{"Z", "W", "L", "C"},
	{"J", "R", "P", "N", "T", "G", "V", "M"},
	{"J", "R", "L", "V", "M", "B", "S"},
	{"D", "P", "J"},
	{"D", "C", "N", "W", "V"},
}
var stackspart2 = []stack{
	{"F", "H", "B", "V", "R", "Q", "D", "D", "P"},
	{"L", "D", "Z", "Q", "W", "V"},
	{"H", "L", "Z", "Q", "G", "R", "P", "C"},
	{"R", "D", "H", "F", "J", "V", "B"},
	{"Z", "W", "L", "C"},
	{"J", "R", "P", "N", "T", "G", "V", "M"},
	{"J", "R", "L", "V", "M", "B", "S"},
	{"D", "P", "J"},
	{"D", "C", "N", "W", "V"},
}

func part1(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		op := strings.Split(line, " ")
		count, _ := strconv.Atoi(op[1])
		src, _ := strconv.Atoi(op[3])
		dst, _ := strconv.Atoi(op[5])
		for i := 0; i < count; i++ {
			move9000(src-1, dst-1)
		}
		move9001(src-1, dst-1, count)

	}
	for _, s := range stacks {
		fmt.Printf("%s", s[len(s)-1])
	}
	fmt.Println("")
	for _, s := range stackspart2 {
		fmt.Printf("%s", s[len(s)-1])
	}
	fmt.Println("")

}
func move9001(src, dst, num int) {
	var s []string
	stackspart2[src], s = stackspart2[src].PopMany(num)
	stackspart2[dst] = stackspart2[dst].Push(s...)
}

func move9000(src, dst int) {
	var s string
	stacks[src], s = stacks[src].Pop()
	stacks[dst] = stacks[dst].Push(s)
}

type stack []string

func (s stack) Push(v ...string) stack {
	return append(s, v...)
}

func (s stack) Pop() (stack, string) {
	l := len(s) - 1
	return s[:l], s[l]
}

func (s stack) PopMany(n int) (stack, []string) {
	l := len(s)
	last := l - n
	return s[:last], s[last:]
}
