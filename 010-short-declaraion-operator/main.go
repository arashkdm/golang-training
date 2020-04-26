package main

import "fmt"

var a = 50
var b int
var c string = `James said, "hahahaha"`

func main() {
	x := 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 99
	fmt.Println(x)
	y := 100 + 42
	fmt.Println(y)
	z := "Bond"
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	// d := 80 error because is unused
	y = 94 // won't error because is only a change

	foo()
}

func foo() {
	fmt.Println("again: ", a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
