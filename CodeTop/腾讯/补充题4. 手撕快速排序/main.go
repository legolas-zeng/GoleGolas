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
func sortArray3(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	flag := nums[0]
	left, right := 0, len(nums)-1
	for i := 1; i <= right; {
		if nums[i] > flag {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			fmt.Println(nums)
		} else {
			nums[i], nums[left] = nums[left], nums[i]
			i++
			left++
			fmt.Println(nums)
		}
	}
	sortArray3(nums[:left])
	sortArray3(nums[left+1:])
	return nums
}
func main() {
	nums := []int{5, 1, 6, 2, 0, 0}
	fmt.Println(sortArray3(nums))
}
