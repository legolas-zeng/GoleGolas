package main

import "fmt"

func fibcci() func() int {
	x, y := 0, 1
	f := func() int {
		x, y = y, x+y
		return y
	}
	//return func() int {
	//	x,y=y,x+y
	//	return y
	//
	//}
	return f
}

func main() {
	f := fibcci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", f())
	}
	B()
	C()
}

func A(i int) {
	i++
	fmt.Println(i)
}

func B() {
	f1 := A
	f1(1)
}
func C() {
	f2 := A
	f2(1)
}
