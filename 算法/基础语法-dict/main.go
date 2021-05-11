package main

import "fmt"

type animal struct {
	id   int
	name string
}

var dict map[int]int

var dict4 map[int]*animal

func main() {
	dict = make(map[int]int)
	dict[1] = 1
	dict2 := make(map[int]string)
	dict2[2] = "123"
	dict3 := map[int]string{
		1: "string",
	}
	dict3[2] = "123"

	dict4 = make(map[int]*animal)
	an := &animal{id: 1, name: "dog"}
	dict4[1] = an

	fmt.Println(dict, dict2, dict3, dict4)
}
