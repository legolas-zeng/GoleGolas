package main

import "fmt"

type mokey struct {
	name string
}

type youngmokey struct {
	mokey
}

type bird interface {
	fly()
}

func (ym *youngmokey) fly() {
	fmt.Println("飞飞飞")
}

func (m *mokey) do() {
	fmt.Println("wangwangwang")
}

func main() {
	littlem := youngmokey{
		mokey{
			"dog",
		},
	}
	littlem.do()
	littlem.fly()
}
