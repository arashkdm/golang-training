package main

import "fmt"

func main() {
	//exec1()
	//exec2()
	//exec3()
	//exec4()
	//exec5()
	exec6()
}

const (
	_ = iota
	a = 365
	b = a * iota
	c = a * iota
	d = a * iota
)
func exec6() {
	fmt.Println(a,b,c,d)
}

/*
func exec5() {
	a := `here is something
	as
	a
	raw string
	literal
	"you see"
	another thing`
	fmt.Println(a)
}

var number int = 37
func exec4() {
	fmt.Printf("%d\t\t%b\t\t%#x\n", number, number, number)
	number = number << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", number, number, number)
}

const (
	a int = 5
	b  = 6
)
func exec3() {
	fmt.Println(a,b)
}

func exec2() {
	a := 5
	b := 8
	an1 := (a == b)
	an2 := (a <= b)
	an3 := (a >= b)
	an4 := (a != b)
	an5 := (a < b)
	an6 := (a > b)
	fmt.Println(an1, an2, an3, an4, an5, an6)
}

func exec1() {
	number := 37
	fmt.Printf("%d\t%b\t\t%#x\n", number, number, number)
}

 */