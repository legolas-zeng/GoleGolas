package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre = head
	len := length(head)
	fmt.Println("长度：", len)
	last := len - n
	if last == 0 {
		//如果是删除的头节点，就返回下一个节点
		return head.Next
	}
	for i := 0; i < last-1; i++ {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return pre
}

func length(head *ListNode) int {
	len := 0
	for head != nil {
		len++
		head = head.Next
	}
	return len
}

/////////////////////  双指针   ////////////////////////////
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	//设置两个指针
	fast := head
	slow := head
	//第一个指针前移n个位置，使它指向第n+1个节点
	//第二个指针没动，这时两个指针相差n个位置
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	//如果第一个指针为空，走到头了，就把第二个指针返回
	if fast == nil {
		return slow.Next
	}
	fmt.Println()
	//当第一个指针到达链表尾部时，
	//第二个指针刚好到达倒数第n个位置
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//跳过倒数第N的节点
	slow.Next = slow.Next.Next
	return head
}
