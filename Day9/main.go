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

	numoftails := 9 //part 2
	rope := make([]Tail, 0, numoftails+1)
	rope = append(rope, Tail{Pos: Coordinate{0, 0}}) //head
	for i := 1; i <= numoftails; i++ {
		rope = append(rope, Tail{Visited: make(map[Coordinate]bool)}) // tails
		rope[i].SetPos(Coordinate{0, 0})
	}

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		op := split[0]
		opcount, _ := strconv.Atoi(split[1])
		for i := 0; i < opcount; i++ {
			rope[0].Pos.Move(op)
			for j := 1; j <= numoftails; j++ {
				rope[j].SetPos(moveTail(rope[j-1].Pos, rope[j].Pos))
			}
		}
	}

	for i := 1; i <= numoftails; i++ {
		fmt.Printf("rope[%d] Visited: %d\n", i, len(rope[i].Visited))
	}
}

type Tail struct {
	Pos     Coordinate
	Visited map[Coordinate]bool
}

func (t *Tail) SetPos(pos Coordinate) {
	t.Pos = pos
	t.Visited[t.Pos] = true
}

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) Move(dir string) {
	switch dir {
	case "U":
		c.Y++
	case "D":
		c.Y--
	case "L":
		c.X--
	case "R":
		c.X++
	default:
		panic("Unknown direction")
	}
}

func moveTail(head, tail Coordinate) Coordinate {
	diffX := head.X - tail.X
	diffY := head.Y - tail.Y
	dirX := 0
	dirY := 0
	if diffX > 0 {
		dirX = 1
	}
	if diffX < 0 {
		dirX = -1
	}
	if diffY > 0 {
		dirY = 1
	}
	if diffY < 0 {
		dirY = -1
	}
	if Abs(diffX) > 1 || Abs(diffY) > 1 {
		return Coordinate{tail.X + dirX, tail.Y + dirY}
	}
	return tail
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
