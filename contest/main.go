package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var str string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		str = strings.TrimSpace(input.Text())
	}
	// fmt.Println('д')
	runeStr := []rune(str)
	fmt.Println(runeStr)
	fmt.Println(runeStr[0])
	fmt.Println(runeStr[utf8.RuneCountInString(str)-1])
//11. Строки не изменяемые
// fullName[0] = "Q"
//12. А слайсы изменяемые :)
fullNameSlice := []rune(fullName)
fullNameSlice[0] = 'F'
fullName = string(fullName
}

/*
4
Сходить в зал
Собрать портфель
Погладить вещи
Почитать книгу
2
4
1
*/
// https://contest.yandex.ru/contest/25606/problems/   ready
// https://contest.yandex.ru/contest/25622/problems/
// https://contest.yandex.ru/contest/19034/problems/ финальное задание
/*

 */
