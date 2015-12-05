package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("input3")
	input, _ := ioutil.ReadAll(f)
	part1(input)
	part2(input)
}

func part1(input []byte) {
	visited := make(map[string]struct{})
	var x, y int
	visited[key(x, y)] = struct{}{}
	for i := range input {
		x, y = next(input[i], x, y)
		visited[key(x, y)] = struct{}{}
	}
	count := 0
	for _ = range visited {
		count += 1
	}
	fmt.Println(count)
}

func part2(input []byte) {
	visited := make(map[string]struct{})
	visited["0,0"] = struct{}{}
	var x1, y1, x2, y2 int
	l := len(input)
	for i := 0; i < l; i += 2 {
		x1, y1 = next(input[i], x1, y1)
		visited[key(x1, y1)] = struct{}{}
		if i+1 < l {
			x2, y2 = next(input[i+1], x2, y2)
			visited[key(x2, y2)] = struct{}{}
		}
	}
	count := 0
	for _ = range visited {
		count += 1
	}
	fmt.Println(count)
}

func next(ins byte, x, y int) (int, int) {
	switch ins {
	case '<':
		x -= 1
	case '^':
		y += 1
	case '>':
		x += 1
	case 'v':
		y -= 1
	}
	return x, y
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
