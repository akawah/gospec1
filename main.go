package main

func main() {

	/////////////////////////////////// 3

	// --

	/////////////////////////////////// 2.5
	// if
	// var val int
	// fmt.Scan(&val)

	// if val%2 == 0 {
	// 	fmt.Println("Чётное")
	// } else {
	// 	fmt.Println("Нечётное")
	// }

	// if num := 10; num%2 == 0 {
	// 	fmt.Println("EVEN")
	// } else {
	// 	fmt.Println("ODD")
	// }

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
