//Author: zwa
//Date: 2020/7/21

package main

import "fmt"

var (
	Test_fun = func(n1 int, n2 int) int {
		return n1 - n2
	}
)

func main() {
	val1 := Test_fun(9, 7)

	fmt.Println("val1=", val1)
}
