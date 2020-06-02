package main

import "fmt"

type Person interface {
	GetAge() int
	GetName() string
}

type Student struct {
	age  int
	name string
}

func (s Student) GetAge() int {
	return s.age
}

func (s Student) GetName() string {
	return s.name
}

func main() {
	var p Person = Student{20, "Elon"}

	fmt.Println("This person name is", p.GetName())
	fmt.Println("This person age is", p.GetAge())
}
