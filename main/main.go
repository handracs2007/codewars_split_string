package main

import (
	"fmt"
)

func Solution(str string) []string {
	size := len(str) / 2

	if len(str)%2 != 0 {
		size++
	}

	a := make([]string, size)
	si := 0

	for i := 0; i < len(str); i++ {
		if i > 0 && i%2 == 0 {
			si++
		}

		a[si] += string(str[i])
	}

	if len(str)%2 != 0 {
		a[si] += "_"
	}

	return a
}

func main() {
	fmt.Println(Solution("abc"))
	fmt.Println(Solution("abcdef"))
}
