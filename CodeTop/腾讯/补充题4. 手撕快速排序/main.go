package main

import "fmt"

//给你一个整数数组 nums，请你将该数组升序排列。

/////////////////////// 冒泡排序法 /////////////////////////
func sortArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
	return nums
}

///////////////////// 桶排序法 ///////////////////////////////
func sortArray2(nums []int) []int {
	s := [101]int{}
	for i := 0; i < len(nums); i++ {
		s[nums[i]+50]++
	}
	idx := 0
	fmt.Println(s)
	for i := 0; i < 101; i++ {
		for s[i] > 0 {
			nums[idx] = i - 50
			idx++
			s[i]--
		}
	}
	return nums
}

///////////////////// 快速排序法 ///////////////////////////////

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	sortArray2(nums)
}
