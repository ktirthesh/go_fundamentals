package main

import "fmt"

func variables() {
	fmt.Println("print the variable variable")
	var a int = 20
	var b float32 = 4.234
	const pi float64 = 3.14

	x, y := 14, 15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("learn the new variable arithmatics")
	fmt.Println("x + y", x+y)
	fmt.Println("x - y", x-y)
	fmt.Println("x * y", x*y)
	fmt.Println("x / y", x/y)

	fmt.Println("x mod y", x%y)

	isbool := true
	hate := false

	fmt.Println("learn the reletional operator")
	fmt.Println("&&", isbool && hate)
	fmt.Println("||", isbool || hate)
	fmt.Println("not", !isbool)

}
