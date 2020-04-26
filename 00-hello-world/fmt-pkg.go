package main

import "fmt"

var y = 42

func main() {
	// print standard stout
	fmt.Print(y)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\n%b\n%x\n", y, y, y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// print in a string
	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y)
	fmt.Println(s)
	fmt.Printf("%v\n", y)

	// print in a file
}
