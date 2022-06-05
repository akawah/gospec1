package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Scanln(&m)
	var prime bool = true

	for i := 1; i <= m/2; i++ {
		if m%i == 0 {
			fmt.Printf("%d ", i)
			if i != 1 {
				prime = false
			}
		}
	}
	fmt.Printf("%d\n", m)
	if prime && m != 1 {
		fmt.Println("ACHTUNG")
	}
}

// m, _ := strconv.Atoi(os.Args[1])
// https://contest.yandex.ru/contest/25606/problems/   ready
// https://contest.yandex.ru/contest/25622/problems/G/
// https://contest.yandex.ru/contest/19034/problems/ финальное задание
/*

 */
