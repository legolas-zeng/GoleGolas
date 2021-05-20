package main

import (
	"fmt"
	"reflect"
)

type User interface {
	callname()
	//calladdr()
}

type Name struct {
	firstname string
	lastname  string
}

type Addr struct {
	country string
	city    string
}

func (name Name) callname() {
	name.lastname = "zhang"
	name.firstname = "san"
	fmt.Println(name)
}

func (addr Addr) callname() {
	fmt.Println("china")
}

func do() User {
	name := new(Name)
	return name
}

//func main(){
//	//var u User
//	u := new(Name)
//	u.callname()
//}

/////////////////////类型断言////////////////////
//func main() {
//	var demo interface{} = "Golang梦工厂"
//	// 类型断言
//	str,err := demo.(int)
//	if !err{
//		fmt.Println("failed")
//	}
//	fmt.Printf("value: %v", str)
//}
/////////////////////////////////////////////////

/////////////////////类型转换////////////////////
func echoArray(a interface{}) {
	//对于函数内部，该变量仍然为interface{}类型,所以会报错
	//for _,v:=range a {
	//	fmt.Print(v," ")
	//}
	//fmt.Println()
	//return
	//////////////////修改之后/////////////////////////
	//	b,_:=a.([]int)//通过断言实现类型转换
	//	for _,v:=range b {
	//		fmt.Print(v," ")
	//	}
	//	fmt.Println()
	//	return
	//////////////////修改之后/////////////////////////
	switch b := a.(type) {
	case []int:
		fmt.Println(b)
	}
	return
}

//func main(){
//	a:=[]int{2,1,3,5,4}
//	echoArray(a)
//}
/////////////////////////////////////////////////

//////////////////interface内部结构/////////////////////////
type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	fmt.Println(reflect.ValueOf(live()))
	fmt.Println(live())
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

/////////////////////////////////////////////////
