package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input2")
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		gift := scanner.Text()
		sum += calc(gift)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}

func calc(gift string) int {
	edges := strings.Split(gift, "x")
	s := 0
	min := math.MaxUint32
	for i := 0; i < len(edges)-1; i++ {
		for j := i + 1; j < len(edges); j++ {
			a, _ := strconv.Atoi(edges[i])
			b, _ := strconv.Atoi(edges[j])
			sur := a * b
			if sur < min {
				min = sur
			}
			s += sur * 2
		}
	}
	s += min
	return s
}
