package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
)

func main() {
	file := "input.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	part1inOrder := 0
	pair := 1
	listOfInput := [][]any{}

	for scanner.Scan() {
		first := unmarshal(scanner.Text())
		listOfInput = append(listOfInput, first)
		scanner.Scan()
		second := unmarshal(scanner.Text())
		listOfInput = append(listOfInput, second)
		scanner.Scan()
		ord := compare(first, second)
		fmt.Printf("Pair %d: inOrder: %d\n\n", pair, ord)
		if ord == Less {
			part1inOrder += pair
		}
		pair++
	}
	fmt.Printf("Part 1: %d\n", part1inOrder)

	var div1 []any
	var div2 []any
	json.Unmarshal([]byte("[[2]]"), &div1)
	json.Unmarshal([]byte("[[6]]"), &div2)
	listOfInput = append(listOfInput, div1, div2)

	sort.Slice(listOfInput, func(i, j int) bool {
		return compare(listOfInput[i], listOfInput[j]) == Less
	})

	for i, v := range listOfInput {
		fmt.Printf("%d: %v\n", i+1, v)
	}
	// go run .  | grep -e " \[\[2\]\]" -e " \[\[6\]\]"
}

type Order int

const (
	Less    Order = iota
	Equal   Order = iota
	Greater Order = iota
)

func unmarshal(input string) []any {
	var v []any
	json.Unmarshal([]byte(input), &v)
	return v
}

func compare(a, b []any) Order {
	fmt.Printf("Compare %v vs %v\n", a, b)
	for i := 0; i < len(a) && i < len(b); i++ {
		left := a[i]
		right := b[i]
		if isNum(left) && isNum(right) {
			lFloat := left.(float64)
			rFloat := right.(float64)
			if lFloat < rFloat {
				return Less
			}
			if lFloat > rFloat {
				return Greater
			}
			if lFloat == rFloat {
				continue
			}
		} else {
			ret := compare(getCorrectAny(left), getCorrectAny(right))
			if ret != Equal {
				return ret
			}
		}
	}

	if len(a) < len(b) {
		return Less
	}
	if len(a) > len(b) {
		return Greater
	}

	return Equal
}

func isNum(a any) bool {
	return reflect.TypeOf(a).Kind() == reflect.Float64
}

func getCorrectAny(a any) []any {
	if isNum(a) {
		return []any{a}
	}
	return a.([]any)
}
