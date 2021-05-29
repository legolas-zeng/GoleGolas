package main

import "fmt"

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println(l1, l2)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var newlist *ListNode
	if l1.Val < l2.Val {
		newlist = l1
		newlist.Next = mergeTwoLists(l1.Next, l2)
	} else {
		newlist = l2
		newlist.Next = mergeTwoLists(l1, l2.Next)
	}
	return newlist
}

func main() {
	head1 := new(ListNode)
	head1.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 4

	head2 := new(ListNode)
	head2.Val = 1
	ln4 := new(ListNode)
	ln4.Val = 3
	ln5 := new(ListNode)
	ln5.Val = 4
	head1.Next = ln2
	ln2.Next = ln3

	head2.Next = ln4
	ln4.Next = ln5

	//pre := mergeTwoLists(head1,head2)
	mergeTwoLists(head1, head2)
	//fmt.Println(pre)
}
