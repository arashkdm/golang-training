package main

import "fmt"

// for exec2:
//var x int
//var y string
//var z bool

// for exec3:
//var x int = 42
//var y string = "James Bond"
//var z bool = true

// for exec4:
//type archioux int
//var x archioux

//for exec5:
type archioux int
var x archioux
var y int

func main() {
	//exec1()
	//exec2()
	//exec3()
	//exec4()
	exec5()
}

func exec5() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

//func exec4() {
//	fmt.Println(x)
//	fmt.Printf("%T\n", x)
//	x = 42
//	fmt.Println(x)
//}

//func exec3() {
//		s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
//		fmt.Println(s)
//}

//func exec2() {
//	// zero values
//	fmt.Println(x)
//	fmt.Println(y)
//	fmt.Println(z)
//}

//func exec1() {
//	x := 42
//	y := "james Bond"
//	z := true
//	fmt.Println(x, y, z)
//	fmt.Println(x)
//	fmt.Println(y)
//	fmt.Println(z)
//}