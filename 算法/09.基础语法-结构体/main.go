package main

import "fmt"

//////////////////开闭原则////////////////////////

//我们要写一个类,Banker银行业务员
type Banker struct {
}

//存款业务
func (this *Banker) Save() {
	fmt.Println("进行了 存款业务...")
}

//转账业务
func (this *Banker) Transfer() {
	fmt.Println("进行了 转账业务...")
}

//支付业务
func (this *Banker) Pay() {
	fmt.Println("进行了 支付业务...")
}

//func main() {
//	banker := &Banker{}
//	banker.Save()
//	banker.Transfer()
//	banker.Pay()
//}
/////////////////////////////////////////////////

//////////////////结构方法/////////////////////////

//type T struct{}
//
//func NewT() *T {
//	return &T{}
//}
//
//func (t *T) callFuncA(x, y string)  {
//	fmt.Println("12345678")
//}
//
//func main() {
//	NewT().callFuncA("炸", "煎鱼")
//	a :=new(T)
//	a.callFuncA("炸", "煎鱼")
//}
/////////////////////////////////////////////////

type Student struct {
	Name string
}

func main() {
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
}
