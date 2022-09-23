package main

import "fmt"

func main() {
	fmt.Println("lets discuss pointer in go")

	x := 7
	println(&x)
	/*
		it's explain how value of x is store in memory location and print some jibbries value
	*/
	a := 7
	b := &a
	fmt.Println(a, b)
	*b = 8
	fmt.Println(a, *b)

	// concept of nil
	var and *int
	fmt.Println(and)
	// if you passed any variable but don't pass any value its give you <nil>
}
