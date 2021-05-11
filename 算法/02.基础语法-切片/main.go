package main

import "fmt"

var sli []int

func main() {
	sli = make([]int, 2)
	sli[1] = 1

	sli2 := []int{1, 2, 3, 4, 5, 6}

	sli3 := append(sli2, 7, 8, 9)
	sli4 := make([]int, 0)
	copy(sli3, sli4)
	fmt.Println(sli, sli2, sli3, sli4)

	fmt.Println(sli3[:3])
	fmt.Println(sli3[3:])
}
