package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre = head
	len := length(head)
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
