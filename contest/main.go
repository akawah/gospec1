package main

import (
	"fmt"
)

func main() {
	var n, sum int

	fmt.Scanln(&n)
	slArr := []int{}

	for i := 0; i < n; i++ {
		var el int
		fmt.Scanln(&el)
		slArr = append(slArr, el)
	}

	for i := 0; i < n; i++ {
		if i%2 != 0 {
			slArr[i] = -slArr[i]
		}
		sum += slArr[i]
	}
	fmt.Println(sum)
}

/*
myWords = append(myWords, anotherSlice...)
1
 1 1
1 2 1
*/
// https://contest.yandex.ru/contest/25606/problems/   ready
// https://contest.yandex.ru/contest/25622/problems/G/
// https://contest.yandex.ru/contest/19034/problems/ финальное задание
/*

 */
