package main

import (
	"fmt"
)

func main() {
	for {
		var s string
		fmt.Scanln(&s)
		if len(s) == 0 {
			return
		}
		fmt.Println(s)
	}
}

// https://contest.yandex.ru/contest/25606/problems/   ready
// https://contest.yandex.ru/contest/25622/problems/G/
// https://contest.yandex.ru/contest/19034/problems/ финальное задание
/*

 */
