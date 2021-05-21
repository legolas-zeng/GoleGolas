package main

import "fmt"

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

////////////////////暴力解法////////////////////////
//func do(arr []int,target int) []int{
//	//tmp:=[]int{}
//	for i:=0;i<len(arr)-1;i++{
//		for j:= i+1 ;j<len(arr);j++{
//			if arr[i]+arr[j]==target{
//				//tmp=append(tmp,i,j)
//				return []int{i,j}
//			}
//		}
//	}
//	//return tmp
//	return nil
//}
//
//
func main() {
	arr := []int{3, 2, 4}
	target := 6
	fmt.Println(do(arr, target))
}

/////////////////////哈希表///////////////////////////
func do(arr []int, target int) []int {
	hashmap := map[int]int{}
	for i, j := range arr {
		if n, ok := hashmap[target-j]; ok {
			// 这里通过用差值来判断这个相等于差值的key是否存在，key1+key2=target,如果key1存在，那么这时候来了key2，map[key1]=map[target-key2],判断key1是否存在
			fmt.Println(j, target-j)
			return []int{i, n}
		}
		hashmap[j] = i
	}
	fmt.Println(hashmap)
	return nil
}
