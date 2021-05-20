package main

import "fmt"

var (
	v1 = 0
	v2 = 0
)

func fibs(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibs(n-1) + fibs(n-2)
	}
	return res
}

func main() {
	for i := 0; i < 10; i++ {
		result := fibs(i)
		fmt.Println(result)
		//fmt.Println(fibci())
	}
}

func fibci() int {
	if v1 == 0 {
		v1 = 1
	} else {
		v1, v2 = v1+v2, v1
	}
	return v1
}
