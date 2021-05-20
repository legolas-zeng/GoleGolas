package main

import "fmt"

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

func main() {
	arr := []int{1, -2, 0, -1, -2, 3, 4}
	arrlist := [][]int{}
	for i := 0; i < len(arr)-2; i++ {
		//fmt.Println("第一位数为----------------",arr[i])
		for j := i + 1; j < len(arr)-1; j++ {
			//fmt.Println("第二位数为+++++++++++++++",arr[j])
			for k := j + 1; k < len(arr); k++ {
				tmparr := []int{}
				if arr[i]+arr[j]+arr[k] == 0 {
					tmparr = append(tmparr, arr[i], arr[j], arr[k])
					arrlist = append(arrlist, tmparr)
				}
			}
		}
	}
	//fmt.Println(arrlist)
	eqnumber(arrlist)
}

func eqnumber(arrlist [][]int) {
	for i := 0; i < len(arrlist); i++ {
		for j := i + 1; j < len(arrlist); j++ {
			fmt.Println(arrlist[i][0], arrlist[i][1], arrlist[i][2], "vs", arrlist[j][0], arrlist[j][1], arrlist[j][2])
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					if arrlist[i][m] == arrlist[j][n] {
						fmt.Println("元素相同")

					}
				}
			}
		}
	}
}
