package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var z int8 = -128 //-128 to 127

func main() {
	x = 42
	y = 42.13212
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
}
