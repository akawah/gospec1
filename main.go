package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello")
	fmt.Println("world")
	fmt.Print("Hello")
	a, b := fmt.Print("world")

	name := "Вах"
	fmt.Printf("\nHello %s\n", name)
	fmt.Println(a, b)

	var age int32 = 4
	fmt.Printf("Переменная age: %d\n", age)

	var c float64 = 3.3

	fmt.Printf("Переменная c: %f\n", c)

	var Var1, Var2 int = 7, 9
	fmt.Printf("Переменная 1: %d\nПеременная 2: %d\n", Var1, Var2)

	var (
		personName string = "Bob"
		personAge         = 42
		personUID  int
	)

	fmt.Printf("Name: %s\nAge %d\nUID: %d\n", personName, personAge, personUID)

	width, height := 23.6, 19.7

	fmt.Printf("Меньшая величина из высоты и ширины: %.2f\n", math.Min(width, height))

	var (
		age2  int
		name2 string
	)
	/* 	fmt.Print("Скажи возраст: ")
	   	fmt.Scan(&age2)
	   	fmt.Print("Скажи имя: ")
	   	fmt.Scan(&name2) */

	fmt.Scan(&age2, &name2)

	fmt.Printf("Моё имя: %s, возраст: %d\n", name2, age2)

}
