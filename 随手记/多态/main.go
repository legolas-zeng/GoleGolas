//Author: zwa
//Date: 2020/6/15

package main

import "fmt"

type annimal interface {
	eat()
	sleep()
	run()
}

type cat interface {
	annimal
	Climb()
}
type dog interface {
	annimal
}

type HelloKitty struct {
	cat
}
type husky struct {
	dog
}

func (h HelloKitty) eat() {
	fmt.Println("eat cake!!")
}

func (h husky) eat() {
	fmt.Println("eat bone!!")
}

func test(a annimal) {
	a.eat()
}

func main() {
	//var a annimal
	//a = HelloKitty{}
	//test(a)
	//a = husky{}
	//test(a)
	var a annimal
	a = HelloKitty{}
	var b annimal
	b = husky{}

	var animals [2]annimal = [...]annimal{a, b}

	for _, v := range animals {
		if data, ok := v.(husky); ok {
			data.eat()
			fmt.Println("this is wangcai : ")
		}
		if data, ok := v.(HelloKitty); ok {
			data.eat()
			fmt.Println("this is HelloKitty : ")
		}
	}

}
