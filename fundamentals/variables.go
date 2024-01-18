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

}
