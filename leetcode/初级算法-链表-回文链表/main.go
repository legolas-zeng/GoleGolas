package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vallist := []int{}
	for head != nil {
		//把链表转换成切片
		vallist = append(vallist, head.Val)
		head = head.Next
	}
	n := len(vallist)
	//使用双指针法判断是否为回文
	for i, v := range vallist[:n/2] {
		if v != vallist[n-1-i] {
			return false
		}
	}
	return true
}
