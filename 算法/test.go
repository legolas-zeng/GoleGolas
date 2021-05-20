package main

import "fmt"

type user struct {
	name string
	age  int
}

//func (u *user)Call(){
//	fmt.Println("hahaha")
//}
//
//func main(){
//	use := user{"admin",12}
//	do(&use)
//}
//
//func do(u *user){
//	var a uint8 = 8
//	//b := uint8(8)
//	c := a * 32
//	fmt.Println(a*32, c)
//	fmt.Println(reflect.TypeOf(c))
//	d := a * 31
//	fmt.Println(a*31, d)
//	fmt.Println(reflect.TypeOf(d))
//
//	getType := reflect.TypeOf(u)
//	getvalue := reflect.ValueOf(u)
//	//fmt.Println(getType.Name())
//	//fmt.Println(getType.NumField())
//	//fmt.Println(getType.Field(0).Name)
//	//fmt.Println(getvalue)
//
//	fmt.Println(getType.NumMethod())
//	for i := 0; i < getType.NumMethod(); i++ {
//		m := getType.Method(i)
//		fmt.Println(m.Name,m.Type)
//		//使用方法
//		//getaddrValue := reflect.ValueOf(u)
//		getvalue.MethodByName(m.Name).Call(nil)
//	}
//
//}

//func fibcci(n int) int{
//	if n <=1 {
//		return 1
//	}else {
//		return fibcci(n-1)+fibcci(n-2)
//	}
//}
//
//func main(){
//	for i:=0;i<10;i++{
//		fmt.Println(fibcci(i))
//	}
//}

func fibcci() func() int {
	x, y := 0, 1
	var count = 0
	return func() int {
		if count < 1 {
			count++
			return 1
		} else {
			x, y = y, x+y
			return y
		}
	}
}
func main() {
	f := fibcci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//type user struct {
//	name string
//	age int
//}
//
//func (u user)Call(){
//	fmt.Println("hahahahha")
//}
//
//func main(){
//	u :=user{
//		name: "aaa",
//		age:  12,
//	}
//	getTpye := reflect.TypeOf(u)
//	getValue:= reflect.ValueOf(u)
//
//	fmt.Println(getTpye,getValue)
//	fmt.Println(getTpye.NumMethod())
//	for i := 0;i< getTpye.NumMethod();i++{
//
//		m:=getTpye.Method(i)
//		getValue.MethodByName(m.Name).Call(nil)
//	}
//
//
//}

//func do(ch chan string)  {
//	for i:=0;i<10;i++{
//		ch<-"dog"
//	}
//	close(ch)
//}
//
//func main(){
//	ch:=make(chan string)
//	go do(ch)
//	for n := range ch{
//		fmt.Println(n)
//	}
//
//}

//func do(arr []int) []int{
//	for i:=0;i<len(arr)-1;i++{
//		for j:=0;j<len(arr)-1-i;j++ {
//			if arr[j] > arr[j+1] {
//			arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//	return arr
//}
//
//func main(){
//	arr:=[]int{3,6,4,5,1,2,7,9}
//	fmt.Println(do(arr))
//	//fmt.Println(QuickSort(arr))
//
//}
