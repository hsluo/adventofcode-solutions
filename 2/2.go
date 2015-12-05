package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input2")
	scanner := bufio.NewScanner(f)
	var sum, l int
	for scanner.Scan() {
		gift := scanner.Text()
		edges := parse(gift)
		sum += wrap(edges)
		l += ribbon(edges)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum, l)
}

func parse(gift string) []int {
	splits := strings.Split(gift, "x")
	edges := make([]int, len(splits))
	for i := range splits {
		e, _ := strconv.Atoi(splits[i])
		edges[i] = e
	}
	sort.Ints(edges)
	return edges
}

func wrap(edges []int) int {
	var s int
	min := math.MaxUint32
	for i := 0; i < len(edges)-1; i++ {
		for j := i + 1; j < len(edges); j++ {
			sur := edges[i] * edges[j]
			if sur < min {
				min = sur
			}
			s += sur * 2
		}
	}
	s += min
	return s
}

func ribbon(edges []int) int {
	return (edges[0]+edges[1])*2 + (edges[0] * edges[1] * edges[2])
}
