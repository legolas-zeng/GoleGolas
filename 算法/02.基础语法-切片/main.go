package main

import "fmt"

var sli []int

//func main() {
//	sli = make([]int, 2)
//	sli[1] = 1
//
//	sli2 := []int{1, 2, 3, 4, 5, 6}
//
//	sli3 := append(sli2, 7, 8, 9)
//	sli4 := make([]int, 10)
//	copy(sli4, sli3)
//	sli5 := sli4
//	sli5[9] = 100
//	fmt.Println(sli, sli2, sli3, sli4,sli5)
//
//	fmt.Println(sli3[:3])
//	fmt.Println(sli3[3:])
//}

//////////////复制/////////////////////
func main() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	fmt.Println(str1)
	fmt.Println(str2)
	str2[1] = "new"
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str2)
}

//TODO copy复制会比等号复制慢。但是copy复制为值复制，改变原切片的值不会影响新切片。而等号复制为指针复制，改变原切片或新切片都会对另一个产生影响
///////////////////////////////////
