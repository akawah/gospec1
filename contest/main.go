package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)

	newStr := strings.ReplaceAll(str, "р", " ")

	arrStr := strings.Fields(newStr)

	max := 0

	for _, el := range arrStr {
		el := []rune(el)
		if len(el) > max {
			max = len(el)
		}
	}
	fmt.Println(max)
}
