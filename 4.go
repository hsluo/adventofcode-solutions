package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	input := "bgvyzdsv"
	i := 0
	i = find(i, input, "00000")
	fmt.Println(i)
	i = find(i, input, "000000")
	fmt.Println(i)
}

func find(start int, input, prefix string) int {
	for {
		h := hash(input, start)
		if strings.HasPrefix(h, prefix) {
			return start
		}
		start++
	}
}

func hash(input string, number int) string {
	data := []byte(fmt.Sprintf("%s%d", input, number))
	return fmt.Sprintf("%x", md5.Sum(data))
}
