package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input5")
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if nice(scanner.Text()) {
			count += 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}

func nice(input string) bool {
	bad := []string{"ab", "cd", "pq", "xy"}
	for i := range bad {
		if strings.Contains(input, bad[i]) {
			return false
		}
	}

	vowels := []string{"a", "e", "i", "o", "u"}
	count := 0
	for i := range vowels {
		count += strings.Count(input, vowels[i])
	}
	if count < 3 {
		return false
	}

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}
	return false
}
