package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("input3")
	bytes, _ := ioutil.ReadAll(f)
	input := string(bytes)
	houses := make(map[string]struct{})
	x, y := 0, 0
	houses[key(x, y)] = struct{}{}
	for i := range input {
		switch input[i] {
		case '<':
			x -= 1
		case '^':
			y += 1
		case '>':
			x += 1
		case 'v':
			y -= 1
		}
		houses[key(x, y)] = struct{}{}
	}
	count := 0
	for _ = range houses {
		count += 1
	}
	fmt.Println(count)
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
