package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var root = fsObject{name: "/", size: 0, children: map[string]*fsObject{}, dir: true, parent: nil}

func main() {
	file := "input.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var current *fsObject
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if words[0] == "$" {
			if words[1] == "cd" {
				target := words[2]
				if target == ".." {
					current = current.parent
				} else if target == "/" {
					current = &root
				} else {
					current = current.getChild(target)
				}
			}
		} else if len(words) == 2 {
			if words[0] == "dir" {
				dir := words[1]
				new := fsObject{name: dir, parent: current, size: 0, children: map[string]*fsObject{}, dir: true}
				current.addChild(&new)
			} else {
				size, _ := strconv.Atoi(words[0])
				file := fsObject{name: words[1], parent: current, size: size, dir: false}
				current.addChild(&file)
			}
		}
	}
	root.calcSize()
	//root.Print()
	fmt.Printf("Size / : %d\n", root.size)
	part1()
	part2()
}

func part1() {
	save := make([]*fsObject, 0)
	target := 100000
	root.findFolderSmallerThan(&save, target)
	var sum int
	for _, c := range save {
		sum += c.size
	}
	fmt.Printf("Sum of all under %d: %d\n", target, sum)
}
func part2() {
	save := make([]*fsObject, 0)
	total := 70000000
	free := total - root.size
	target := 30000000 - free
	root.findFolderBiggerThan(&save, target)
	var sum int
	for _, c := range save {
		sum += c.size
	}
	sort.Slice(save, func(i, j int) bool {
		return save[i].size < save[j].size
	})
	fmt.Printf("Smallest above %d: %s %d\n", target, save[0].name, save[0].size)
}

func (f *fsObject) findFolderBiggerThan(save *[]*fsObject, size int) {
	for _, c := range f.children {
		if c.isDir() {
			if c.size > size {
				*save = append(*save, c)
			}
			c.findFolderBiggerThan(save, size)
		}
	}
}

func (f *fsObject) findFolderSmallerThan(save *[]*fsObject, size int) {
	for _, c := range f.children {
		if c.isDir() {
			if c.size < size {
				*save = append(*save, c)
			}
			c.findFolderSmallerThan(save, size)
		}
	}
}

type fsObject struct {
	name     string
	size     int
	dir      bool
	parent   *fsObject
	children map[string]*fsObject
}

func (f *fsObject) addChild(newchild *fsObject) {
	if f.isDir() && !f.isChildExistant(newchild.name) {
		f.children[newchild.name] = newchild
	}
}

func (f *fsObject) getChild(name string) *fsObject {
	if v, exists := f.children[name]; exists {
		return v
	}
	return nil
}

func (f *fsObject) Print() {
	if f.isDir() {
		for _, c := range f.children {
			c.Print()
		}
	} else {
		fmt.Printf("%s : %d\n", f.getFullPath(), f.size)
	}
}

func (f *fsObject) getFullPath() string {
	for f.parent != nil {
		return f.parent.getFullPath() + "/" + f.name
	}
	return ""
}

func (f *fsObject) isChildExistant(name string) bool {
	_, exists := f.children[name]
	return exists
}

func (f *fsObject) isDir() bool {
	return f.dir
}

func (f *fsObject) calcSize() int {
	if f.isDir() {
		for _, c := range f.children {
			f.size += c.calcSize()
		}
	}
	return f.size
}
