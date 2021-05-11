package main

import "fmt"

type UserAges struct {
	ages map[string]int
	name string
}

type animal struct {
	name string
	addr string
}

var c animal

func main() {
	a := new(UserAges)
	a.ages = map[string]int{"123": 1}

	b := animal{name: "dog", addr: "123"}
	c.name = "cat"
	c.addr = "321"
	fmt.Println(a, b, c)
}
