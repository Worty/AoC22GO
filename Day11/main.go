package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var lcm int = 1

func main() {
	file := "input.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	monkeys := []Monkey{}
	for scanner.Scan() {
		scanner.Scan()

		itemline := scanner.Text()[18:]
		numsline := strings.Split(itemline, ", ")
		items := readItems(numsline)

		fmt.Println(items)
		scanner.Scan()

		operation := getOperation(scanner.Text()[13:])
		scanner.Scan()

		test := scanner.Text()[8:]
		scanner.Scan()
		truel := scanner.Text()
		scanner.Scan()
		falsel := scanner.Text()
		testf := getThrowFunc(test, truel, falsel)
		scanner.Scan()

		monkeys = append(monkeys, Monkey{ItemList: items, Inspect: operation, ThrowTo: testf})
	}
	for i := 1; i <= 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			for len(monkeys[j].ItemList) > 0 {
				monkeys[j].inspected++
				item := monkeys[j].ItemList[0]
				monkeys[j].ItemList = monkeys[j].ItemList[1:]
				worrylevel := monkeys[j].Inspect(item) % lcm // / 3 for part 1
				throwto := monkeys[j].ThrowTo(worrylevel)
				monkeys[throwto].ItemList = append(monkeys[throwto].ItemList, worrylevel)
			}
		}
	}
	var inspected []int
	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("Monkey %d inspected %d items\n", i, monkeys[i].inspected)
		inspected = append(inspected, monkeys[i].inspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	fmt.Printf("Result: %d\n", inspected[0]*inspected[1])

}
func getThrowFunc(test, truel, falsel string) func(int) int {
	a := strings.Split(test, " ")
	n, err := strconv.Atoi(a[2])
	if err != nil {
		panic(err)
	}
	lcm *= n
	truec := strings.Split(truel, " ")
	truemoneky, _ := strconv.Atoi(truec[len(truec)-1])
	falsec := strings.Split(falsel, " ")
	falsemonkey, _ := strconv.Atoi(falsec[len(falsec)-1])
	return func(item int) int {
		if item%int(n) == 0 {
			return truemoneky
		} else {
			return falsemonkey
		}
	}
}

func readItems(numsline []string) (items []int) {
	for i := 0; i < len(numsline); i++ {
		num, err := strconv.Atoi(numsline[i])
		if err != nil {
			panic(err)
		}
		items = append(items, int(num))
	}
	return

}
func getOperation(input string) func(int) int {
	operatorindex := strings.Index(input, "+")
	if operatorindex != -1 {
		if input[operatorindex+2] != 'o' {
			n, _ := strconv.Atoi(input[operatorindex+2:])
			return func(old int) int {
				return old + int(n)
			}
		} else {
			return func(old int) int {
				return old + old
			}
		}
	}
	operatorindex = strings.Index(input, "*")
	if operatorindex != -1 {
		if input[operatorindex+2] != 'o' {
			n, _ := strconv.Atoi(input[operatorindex+2:])
			return func(old int) int {
				return old * int(n)
			}
		} else {
			return func(old int) int {
				return old * old
			}
		}
	}
	return func(old int) int {
		panic("no operation")
	}

}

type Monkey struct {
	ItemList  []int
	Inspect   func(int) int
	ThrowTo   func(int) int
	inspected int
}
