package Day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func day1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	if err != nil {
		panic(err)
	}
	elves := []int{}
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == "\n" {
			elves = append(elves, counter)
			counter = 0
			continue
		}
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		counter += int(num)
	}

	sort.Ints(elves)
	fmt.Printf("Max: %d\n", elves[len(elves)-1])
	fmt.Printf("Top 3: %v\n", elves[len(elves)-3:])
	fmt.Printf("Sum: %d\n", manySum(elves[len(elves)-3:]...))
}

func manySum(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}
