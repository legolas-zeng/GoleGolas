//Author: zwa
//Date: 2020/6/15

package main

import (
	"fmt"
)

type Biology interface {
	sayhi()
}
type Man struct {
	name string
	age  int
}

type Monster struct {
	name string
	age  int
}

func (this *Man) sayhi() { // 实现抽象方法1
	// fmt.Printf("Man[%s, %d] sayhi\n", this.name, this.age)
}

func (this *Monster) sayhi() { // 实现抽象方法1
	fmt.Printf("Monster[%s, %d] sayhi\n", this.name, this.age)
}

func WhoSayHi(i Biology) {
	i.sayhi()
}

func main() {
	man := &Man{"我是人", 100}
	monster := &Monster{"妖怪", 1000}
	WhoSayHi(man)
	WhoSayHi(monster)
}
