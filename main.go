package main

import "fmt"

func main() {

	/////////////////////////////////// 5

	name := "Hello world!"
	// fmt.Println(name)
	fmt.Println(name)
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Println()

	word := "Привет, мир!"
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i])
	}
	fmt.Println()

	runeSlice := []rune(word)
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c ", runeSlice[i])
	}
	fmt.Println()
	fmt.Printf("%c\n", runeSlice[2])

	for id, runeVal := range runeSlice {
		fmt.Printf("%d - %c\n", id, runeVal)
	}

	fmt.Println(string(runeSlice))
	fmt.Println(string(word))
	wSl := string(word)

	fmt.Println(wSl[3])
	/*
			var name string
			input := bufio.NewScanner(os.Stdin)
			if input.Scan() {
				name = input.Text()
			}
			if name == "" {
				fmt.Println("empty")
				return
			}
			fmt.Println(name)

		numStr := "500"
		num, _ := strconv.ParseInt(numStr, 10, 64)

		fmt.Println(num)
	*/
	/*


		for i := 0; i < len(word); i++ {
			fmt.Printf("%c ", word[i])
		}
		fmt.Println()

		runeSlice := []rune(word)
		for i := 0; i < len(runeSlice); i++ {
			fmt.Printf("%c ", runeSlice[i])
		}
		fmt.Println()
	*/
	/////////////////////////////////// 4
	/*
		originArr := [...]int{30, 40, 50, 60, 70, 80}
		firstSlice := originArr[1:4]
		for i, _ := range firstSlice {
			firstSlice[i]++
		}
		fmt.Println("OriginArr:", originArr)
		fmt.Println("FirstSlice:", firstSlice)

		fSlice := originArr[:]
		sSlice := originArr[2:5]

		fmt.Println("Befor modification: Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)
		fSlice[3]++
		sSlice[1]++
		fmt.Println("After modification: Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)
	*/

	/*
		arr := [...]float64{12.54, 23, 11.1, 19.43, 83.00}
		for i := 0; i < len(arr); i++ {
			fmt.Printf("El %d = %.2f\n", i, arr[i])
		}

		var sum float64
		for _, val := range arr {
			sum += val
		}
		fmt.Println("Сумма эл-тов:", sum)

		var arr3 [5]int
		fmt.Println("Это массив:", arr3)
	*/
	/////////////////////////////////// 2.5
	// if
	/*
		var val int
		fmt.Scan(&val)

		if val%2 == 0 {
			fmt.Println("Чётное")
		} else {
			fmt.Println("Нечётное")
		}

		if num := 10; num%2 == 0 {
			fmt.Println("EVEN")
		} else {
			fmt.Println("ODD")
		}
	*/
	//
	//
	//
	// coob, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	// handle error
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }
	// fmt.Println(coob * coob * coob)

	/*

		///////////////////////////// 2
		// bool
		var bul bool
		fmt.Println(bul)

		bul1, bul2 := true, true

		fmt.Println(bul1 && bul2)

		// int8 int16 int32 int64 int
		// uint8 uint16 uint32 uint64 uint

		var a int = 32
		b := 93

		fmt.Printf("Sum = %d\n", a+b)
		fmt.Printf("%T\n", a)
		fmt.Printf("%T size of %d bytes\n", a, unsafe.Sizeof(a))
		fmt.Printf("%T size of %d bytes\n", b, unsafe.Sizeof(b))

		var d32 int32 = 12
		var d64 int64 = 16
		fmt.Printf("Sum = %d\n", int64(d32)+d64)

		// float32 float64
		flo1, flo2 := 3.65, 83.55
		fmt.Printf("%T, %T, sum = %.2f\n", flo1, flo2, flo1+flo2)

		// complex
		c1 := complex(5, 7)
		c2 := 12 + 32i
		fmt.Println(c1 + c2)

		// strings
		name := "Вася"
		lastname := "Пупкин"
		fmt.Println(name + " " + lastname)
		fmt.Println(len(name))
		fmt.Println(utf8.RuneCountInString(name))

		// rune -> alias int32
		var qRune rune = 'Q'
		fmt.Printf("Rune %c\n", qRune)

		fmt.Println(strings.Compare("abcd", "abcd")) // 0
		fmt.Println(strings.Compare("abcd", "defg")) // -1
		fmt.Println(strings.Compare("defg", "abcd")) // 1

		// byte -> alias int8
		var iByte byte = 11
		fmt.Println("Byte:", iByte)






	*/

	///////////////////////////// 1
	/*	fmt.Println("Hello")
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
		fmt.Print("Скажи возраст: ")
		fmt.Scan(&age2)
		fmt.Print("Скажи имя: ")
		fmt.Scan(&name2)

		fmt.Scan(&age2, &name2)

		fmt.Printf("Моё имя: %s, возраст: %d\n", name2, age2)
	*/
}
