package main

import "fmt"

//接收者(Receiver)有两种，一种是值接收者，一种是指针接收者
//关于这两种接收者，其实GO在编译的时候有一个隐式转换，将其转换为正确的接收者类型

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}

func (t *T) M2() {
	t.Name = "name2"
}

func main() {
	t1 := T{Name: "t1"}
	t1.M1()
	fmt.Println(t1.Name)
	t1.M2()
	fmt.Println(t1.Name)
}
