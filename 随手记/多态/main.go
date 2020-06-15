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
	var a annimal
	a = HelloKitty{}
	test(a)
	a = husky{}
	test(a)
}
