package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	input := "bgvyzdsv"
	for i := 0; ; i++ {
		if hash(input, i) {
			fmt.Println(i)
			break
		}
	}
}

func hash(input string, number int) bool {
	data := []byte(fmt.Sprintf("%s%d", input, number))
	h := fmt.Sprintf("%x", md5.Sum(data))
	return strings.HasPrefix(h, "00000")
}
