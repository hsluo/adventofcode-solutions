package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("input1")
	input, _ := ioutil.ReadAll(f)
	result := 0
	index := 0
	for i := range input {
		if input[i] == '(' {
			result++
		} else if input[i] == ')' {
			result--
			if result < 0 && index == 0 {
				index = i + 1
			}
		}
	}
	fmt.Println(result, index)
}
