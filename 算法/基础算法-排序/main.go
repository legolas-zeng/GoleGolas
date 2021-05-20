package main

import "fmt"

//func do (arr []int)[]int{
//	for i:=0;i<len(arr)-1;i++{
//		for j:=0;j<len(arr)-i-1;j++{
//			if arr[j+1] > arr[j]{
//				arr[j],arr[j+1]=arr[j+1],arr[j]
//				fmt.Println(arr)
//			}
//		}
//	}
//
//	return arr
//}
//
//func  main()  {
//	arr :=[]int{1,2,3,4,5,9,6,8,7,0}
//	fmt.Println(do(arr))
//}

func QuickSort(intList []int) {
	flag := intList[0]
	left, right := 0, len(intList)-1
	for i := 1; i <= right; {
		if intList[i] > flag {
			intList[i], intList[right] = intList[right], intList[i]
			right--
		} else {
			intList[i], intList[left] = intList[left], intList[i]
			i++
			left++
		}
	}
	QuickSort(intList[:left])
	QuickSort(intList[left+1:])
}

func main() {
	list := []int{5, 6, 4, 3, 7, 99, 33, 8}
	QuickSort(list)
	fmt.Println(list)
}
