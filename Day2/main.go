package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := "input.txt"
	part1(file)
	part2(file)
}

func part1(file string) {
	// day2
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	points := 0
	for scanner.Scan() {
		line := scanner.Text()
		opponentMove := variableToEnum(string(line[0]))
		corretMove := rightMove(opponentMove)
		shouldMove := string(line[2])
		if (shouldMove) == corretMove {
			points += WON
		} else if (shouldMove) == convertOpponentMoveToMyMove(opponentMove) {
			points += DRAW
		}

		switch shouldMove {
		case MYROCK:
			points += 1
		case MYPAPER:
			points += 2
		case MYSISSORS:
			points += 3
		}
	}
	fmt.Printf("Part1: Points: %d\n", points)
}

func part2(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	points := 0
	for scanner.Scan() {
		line := scanner.Text()
		opponentMove := variableToEnum(string(line[0]))
		shouldWin := string(line[2])
		if (shouldWin) == SHOULDLOSE {
			points += LOSE
			points += pointsForWrongMove(opponentMove)
		} else if (shouldWin) == SHOULDWIN {
			points += pointsForRightMove(opponentMove)
			points += WON
		} else if (shouldWin) == SHOULDDRAW {
			points += DRAW
			points += pointsForMyMove(opponentMove)
		}
	}
	fmt.Printf("Part2: Points: %d\n", points)
}

const (
	ROCK       string = "A"
	PAPER      string = "B"
	SISSORS    string = "C"
	MYROCK     string = "X"
	MYPAPER    string = "Y"
	MYSISSORS  string = "Z"
	SHOULDLOSE string = "X"
	SHOULDDRAW string = "Y"
	SHOULDWIN  string = "Z"
)

const (
	MYROCKPOINTS    = 1
	MYPAPERPOINTS   = 2
	MYSISSORSPOINTS = 3
)

const (
	LOSE = 0
	DRAW = 3
	WON  = 6
)

func rightMove(opponentMove string) string {
	switch opponentMove {
	case ROCK:
		return MYPAPER
	case PAPER:
		return MYSISSORS
	case SISSORS:
		return MYROCK
	}
	return "0"
}

func convertOpponentMoveToMyMove(opponentMove string) string {
	switch opponentMove {
	case ROCK:
		return MYROCK
	case PAPER:
		return MYPAPER
	case SISSORS:
		return MYSISSORS
	}
	return "0"
}

func variableToEnum(opponentMove string) string {
	switch opponentMove {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SISSORS
	}
	return "0"
}

func pointsForMyMove(move string) int {
	switch move {
	case ROCK:
		return MYROCKPOINTS
	case PAPER:
		return MYPAPERPOINTS
	case SISSORS:
		return MYSISSORSPOINTS
	}
	return 0
}

func pointsForRightMove(opponentMove string) int {
	switch opponentMove {
	case ROCK:
		return MYPAPERPOINTS
	case PAPER:
		return MYSISSORSPOINTS
	case SISSORS:
		return MYROCKPOINTS
	}
	return 0
}

func pointsForWrongMove(opponentMove string) int {
	switch opponentMove {
	case ROCK:
		return MYSISSORSPOINTS
	case PAPER:
		return MYROCKPOINTS
	case SISSORS:
		return MYPAPERPOINTS
	}
	return 0
}
