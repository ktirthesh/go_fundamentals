package main

import "fmt"

func pointereampple() {
	x := 10

	fmt.Println("pointer")
	fmt.Println("value", x)
	fmt.Println("pointer", &x)
	// changevalue(x)
	// fmt.Println("after change value nothing happen beacuse sending as value ", x)
	changevalue(&x)
	fmt.Println("after change bhy passing refernece to variable", x)

}

func changevalue(x *int) {
	*x = 7

}
