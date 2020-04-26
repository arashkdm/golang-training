package main

import "fmt"

var a = 50
var b int

func main() {
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 42
	fmt.Println(y)
	z := "Bond"
	fmt.Println(z)
	fmt.Println(a)

	foo()
}

func foo() {
	fmt.Println("again: ", a)
	fmt.Println(b)
}
