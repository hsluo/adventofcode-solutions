package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input5")
	var count1, count2 int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if nice1(scanner.Text()) {
			count1 += 1
		}
		if nice2(scanner.Text()) {
			count2 += 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(count1, count2)
}

func nice1(input string) bool {
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

func nice2(input string) bool {
	var prop1, prop2 bool

	for i := 0; i < len(input)-1; i++ {
		pair := string([]byte{input[i], input[i+1]})
		if strings.Count(input, pair) > 1 {
			prop1 = true
			break
		}
	}

	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			prop2 = true
			break
		}
	}
	return prop1 && prop2
}
