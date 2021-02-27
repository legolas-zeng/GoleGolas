//Author: zwa
//Date: 2020/12/31

package main

import "fmt"

func main() {
	n := 0
	f := func() int {
		n += 1
		return n
	}
	fmt.Println(f()) // 别忘记括号，不加括号相当于地址
	fmt.Println(f())
}
